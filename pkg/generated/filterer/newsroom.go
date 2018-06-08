// Code generated by 'gen/watchergen'  DO NOT EDIT.
// IT SHOULD NOT BE EDITED BY HAND AS ANY CHANGES MAY BE OVERWRITTEN
// Please reference 'listener/watchergen' for more details
// File was generated at 2018-06-08 21:13:58.322726475 +0000 UTC
package filterer

import (
	"fmt"
	log "github.com/golang/glog"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/joincivil/civil-events-crawler/pkg/generated/contract"
	"github.com/joincivil/civil-events-crawler/pkg/model"

	"math/big"
)

type NewsroomContractFilterers struct{}

func (r *NewsroomContractFilterers) ContractName() string {
	return "NewsroomContract"
}

func (r *NewsroomContractFilterers) StartFilterers(client bind.ContractBackend, contractAddress common.Address,
	pastEvents *[]model.CivilEvent, startBlock uint64) error {
	return r.StartNewsroomContractFilterers(client, contractAddress, pastEvents, startBlock)
}

// StartNewsroomContractFilterers retrieves events for NewsroomContract
func (r *NewsroomContractFilterers) StartNewsroomContractFilterers(client bind.ContractBackend,
	contractAddress common.Address, pastEvents *[]model.CivilEvent, startBlock uint64) error {
	contract, err := contract.NewNewsroomContract(contractAddress, client)
	if err != nil {
		log.Errorf("Error initializing StartNewsroomContract: err: %v", err)
		return err
	}

	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	err = startFilterContentPublished(opts, contract, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving ContentPublished: err: %v", err)
	}

	err = startFilterNameChanged(opts, contract, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving NameChanged: err: %v", err)
	}

	err = startFilterOwnershipTransferred(opts, contract, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving OwnershipTransferred: err: %v", err)
	}

	err = startFilterRevisionSigned(opts, contract, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving RevisionSigned: err: %v", err)
	}

	err = startFilterRevisionUpdated(opts, contract, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving RevisionUpdated: err: %v", err)
	}

	err = startFilterRoleAdded(opts, contract, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving RoleAdded: err: %v", err)
	}

	err = startFilterRoleRemoved(opts, contract, pastEvents)
	if err != nil {
		return fmt.Errorf("Error retrieving RoleRemoved: err: %v", err)
	}

	return nil
}

func startFilterContentPublished(opts *bind.FilterOpts, _contract *contract.NewsroomContract, pastEvents *[]model.CivilEvent) error {
	itr, err := _contract.FilterContentPublished(
		opts,
		[]common.Address{},
		[]*big.Int{},
	)
	if err != nil {
		log.Errorf("Error getting event ContentPublished: %v", err)
		return err
	}
	nextEvent := itr.Next()
	for nextEvent {
		civilEvent := model.NewCivilEvent("ContentPublished", itr.Event)
		*pastEvents = append(*pastEvents, *civilEvent)
		nextEvent = itr.Next()
	}
	return nil
}

func startFilterNameChanged(opts *bind.FilterOpts, _contract *contract.NewsroomContract, pastEvents *[]model.CivilEvent) error {
	itr, err := _contract.FilterNameChanged(
		opts,
	)
	if err != nil {
		log.Errorf("Error getting event NameChanged: %v", err)
		return err
	}
	nextEvent := itr.Next()
	for nextEvent {
		civilEvent := model.NewCivilEvent("NameChanged", itr.Event)
		*pastEvents = append(*pastEvents, *civilEvent)
		nextEvent = itr.Next()
	}
	return nil
}

func startFilterOwnershipTransferred(opts *bind.FilterOpts, _contract *contract.NewsroomContract, pastEvents *[]model.CivilEvent) error {
	itr, err := _contract.FilterOwnershipTransferred(
		opts,
		[]common.Address{},
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error getting event OwnershipTransferred: %v", err)
		return err
	}
	nextEvent := itr.Next()
	for nextEvent {
		civilEvent := model.NewCivilEvent("OwnershipTransferred", itr.Event)
		*pastEvents = append(*pastEvents, *civilEvent)
		nextEvent = itr.Next()
	}
	return nil
}

func startFilterRevisionSigned(opts *bind.FilterOpts, _contract *contract.NewsroomContract, pastEvents *[]model.CivilEvent) error {
	itr, err := _contract.FilterRevisionSigned(
		opts,
		[]*big.Int{},
		[]*big.Int{},
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error getting event RevisionSigned: %v", err)
		return err
	}
	nextEvent := itr.Next()
	for nextEvent {
		civilEvent := model.NewCivilEvent("RevisionSigned", itr.Event)
		*pastEvents = append(*pastEvents, *civilEvent)
		nextEvent = itr.Next()
	}
	return nil
}

func startFilterRevisionUpdated(opts *bind.FilterOpts, _contract *contract.NewsroomContract, pastEvents *[]model.CivilEvent) error {
	itr, err := _contract.FilterRevisionUpdated(
		opts,
		[]common.Address{},
		[]*big.Int{},
		[]*big.Int{},
	)
	if err != nil {
		log.Errorf("Error getting event RevisionUpdated: %v", err)
		return err
	}
	nextEvent := itr.Next()
	for nextEvent {
		civilEvent := model.NewCivilEvent("RevisionUpdated", itr.Event)
		*pastEvents = append(*pastEvents, *civilEvent)
		nextEvent = itr.Next()
	}
	return nil
}

func startFilterRoleAdded(opts *bind.FilterOpts, _contract *contract.NewsroomContract, pastEvents *[]model.CivilEvent) error {
	itr, err := _contract.FilterRoleAdded(
		opts,
		[]common.Address{},
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error getting event RoleAdded: %v", err)
		return err
	}
	nextEvent := itr.Next()
	for nextEvent {
		civilEvent := model.NewCivilEvent("RoleAdded", itr.Event)
		*pastEvents = append(*pastEvents, *civilEvent)
		nextEvent = itr.Next()
	}
	return nil
}

func startFilterRoleRemoved(opts *bind.FilterOpts, _contract *contract.NewsroomContract, pastEvents *[]model.CivilEvent) error {
	itr, err := _contract.FilterRoleRemoved(
		opts,
		[]common.Address{},
		[]common.Address{},
	)
	if err != nil {
		log.Errorf("Error getting event RoleRemoved: %v", err)
		return err
	}
	nextEvent := itr.Next()
	for nextEvent {
		civilEvent := model.NewCivilEvent("RoleRemoved", itr.Event)
		*pastEvents = append(*pastEvents, *civilEvent)
		nextEvent = itr.Next()
	}
	return nil
}
