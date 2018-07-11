package postgres // import "github.com/joincivil/civil-events-crawler/pkg/persistence/postgres"

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"math/big"
	"strings"
)

// EventsTableSchema returns the schema to create this table
func EventsTableSchema() string {
	schema := `
		CREATE TABLE IF NOT EXISTS events(
			id SERIAL PRIMARY KEY,
			event_type TEXT,
			hash TEXT UNIQUE,
			contract_address TEXT,
			contract_name TEXT,
			timestamp INT,
			payload JSONB,
			log_payload JSONB
		);
	`
	return schema
}

// EventPayloadMap is the jsonb payload
type EventPayloadMap map[string]interface{}

// Value is the value interface implemented for the sql driver
func (ep EventPayloadMap) Value() (driver.Value, error) {
	return json.Marshal(ep)
}

// Scan is the scan interface implemented for the sql driver
func (ep *EventPayloadMap) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion .([]byte) failed")
	}
	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*ep, ok = i.(map[string]interface{})
	if !ok {
		return errors.New("type assertion .(map[string]interface{}) failed")
	}

	return nil
}

// CivilEvent is the model for events table in DB
type CivilEvent struct {
	EventType       string          `db:"event_type"`
	EventHash       string          `db:"hash"`
	ContractAddress string          `db:"contract_address"`
	ContractName    string          `db:"contract_name"`
	Timestamp       int             `db:"timestamp"`
	EventPayload    EventPayloadMap `db:"payload"`
	LogPayload      EventPayloadMap `db:"log_payload"`
}

// NewCivilEvent constructs a civil event for DB from a model.civilevent
// Rename this to NewDBEventFromCivilEvent
func NewCivilEvent(civilEvent *model.CivilEvent) (*CivilEvent, error) {
	dbCivilEvent := &CivilEvent{}
	dbCivilEvent.EventType = civilEvent.EventType()
	dbCivilEvent.EventHash = civilEvent.Hash()
	dbCivilEvent.ContractName = civilEvent.ContractName()
	dbCivilEvent.ContractAddress = civilEvent.ContractAddress().Hex()
	dbCivilEvent.Timestamp = civilEvent.Timestamp()
	dbCivilEvent.EventPayload = make(EventPayloadMap)
	dbCivilEvent.LogPayload = make(EventPayloadMap)
	err := dbCivilEvent.parseEventPayload(civilEvent)
	if err != nil {
		return nil, err
	}
	return dbCivilEvent, nil
}

// parseEventPayload() parses and converts payloads from civilevent to store in DB:
func (c *CivilEvent) parseEventPayload(civilEvent *model.CivilEvent) error {
	_abi, err := model.AbiJSON(c.ContractName)
	if err != nil {
		return err
	}
	err = c.EventDataToDB(civilEvent.EventPayload(), _abi)
	if err != nil {
		return err
	}
	c.EventLogDataToDB(civilEvent.LogPayload())
	return nil
}

// EventDataToDB converts event types so they can be stored in the DB
func (c *CivilEvent) EventDataToDB(civilEvent map[string]interface{}, _abi abi.ABI) error {
	// loop through abi, and for each val in map, have a way to convert it
	for _, input := range _abi.Events["_"+c.EventType].Inputs {
		eventFieldName := strings.Title(input.Name)
		eventField := civilEvent[eventFieldName]
		switch input.Type.String() {
		case "address":
			c.EventPayload[eventFieldName] = eventField.(common.Address).Hex()
		case "uint256":
			// NOTE: converting all *big.Int to int64. assuming for now that numbers will fall into int64 range.
			c.EventPayload[eventFieldName] = eventField.(*big.Int).Int64()
		case "string":
			c.EventPayload[eventFieldName] = eventField.(string)
		case "default":
			return fmt.Errorf("unsupported type")

		}
	}
	return nil
}

// DBToEventData converts the db event model to a model.CivilEvent
// NOTE: because this is stored in DB as a map[string]interface{}, Postgres converts some fields, see notes in function.
func (c *CivilEvent) DBToEventData() (*model.CivilEvent, error) {
	civilEvent := &model.CivilEvent{}
	_abi, err := model.AbiJSON(c.ContractName)
	if err != nil {
		return civilEvent, err
	}
	eventPayload := make(map[string]interface{})

	for _, input := range _abi.Events["_"+c.EventType].Inputs {
		eventFieldName := strings.Title(input.Name)
		eventField, ok := c.EventPayload[eventFieldName]
		if !ok {
			return civilEvent, fmt.Errorf("Cannot get %v field of DB CivilEvent", eventFieldName)
		}
		switch input.Type.String() {
		case "address":
			address, addressOk := eventField.(string)
			if !addressOk {
				return civilEvent, errors.New("Cannot cast DB contract address to string")
			}
			eventPayload[eventFieldName] = common.HexToAddress(address)
		case "uint256":
			// NOTE: Ints are stored in DB as float64
			num, numOk := eventField.(float64)
			if !numOk {
				return civilEvent, errors.New("Cannot cast DB int to float64")
			}
			eventPayload[eventFieldName] = big.NewInt(int64(num))
		case "string":
			str, stringOk := eventField.(string)
			if !stringOk {
				return civilEvent, errors.New("Cannot cast DB string val to string")
			}
			eventPayload[eventFieldName] = str
		default:
			return civilEvent, fmt.Errorf("unsupported type in %v field encountered in %v event", eventFieldName, c.EventHash)
		}
	}

	logPayload := c.DBToEventLogData()
	contractAddress := common.HexToAddress(c.ContractAddress)
	civilEvent, err = model.NewCivilEvent(c.EventType, c.ContractName, contractAddress, c.Timestamp, eventPayload,
		logPayload)

	return civilEvent, err

}

// EventLogDataToDB converts the raw log data to Postgresql types
func (c *CivilEvent) EventLogDataToDB(payload *types.Log) {
	c.LogPayload["Address"] = payload.Address.Hex()

	topics := make([]string, len(payload.Topics))
	for _, topic := range payload.Topics {
		topics = append(topics, topic.Hex())
	}
	c.LogPayload["Topics"] = topics

	c.LogPayload["Data"] = payload.Data

	c.LogPayload["BlockNumber"] = payload.BlockNumber
	c.LogPayload["TxHash"] = payload.TxHash.Hex()
	c.LogPayload["TxIndex"] = payload.TxIndex
	c.LogPayload["BlockHash"] = payload.BlockHash.Hex()
	c.LogPayload["Index"] = payload.Index
	c.LogPayload["Removed"] = payload.Removed

}

// DBToEventLogData converts the DB raw log payload back to types.Log
// NOTE: because this is stored in DB as a map[string]interface{}, Postgres converts some fields, see notes in function.
func (c *CivilEvent) DBToEventLogData() *types.Log {
	log := &types.Log{}
	log.Address = common.HexToAddress(c.LogPayload["Address"].(string))

	topics := c.LogPayload["Topics"].([]interface{})
	newTopics := make([]common.Hash, len(topics))
	for i, topic := range topics {
		topicString := topic.(string)
		newTopics[i] = common.HexToHash(topicString)
	}
	log.Topics = newTopics

	// NOTE: Data is stored in DB as a string
	log.Data = []byte(c.LogPayload["Data"].(string))

	// NOTE: BlockNumber is stored in DB as float64
	log.BlockNumber = uint64(c.LogPayload["BlockNumber"].(float64))

	log.TxHash = common.HexToHash(c.LogPayload["TxHash"].(string))

	// NOTE: TxIndex is stored in DB as float64
	log.TxIndex = uint(c.LogPayload["TxIndex"].(float64))

	log.BlockHash = common.HexToHash(c.LogPayload["BlockHash"].(string))

	// NOTE: Index is stored in DB as float64
	log.Index = uint(c.LogPayload["Index"].(float64))

	log.Removed = c.LogPayload["Removed"].(bool)
	return log
}
