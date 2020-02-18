// Code generated by 'gen/eventhandlergen.go'  DO NOT EDIT.
// IT SHOULD NOT BE EDITED BY HAND AS ANY CHANGES MAY BE OVERWRITTEN
// Please reference 'gen/filterergen_template.go' for more details
// File was generated at 2020-02-18 16:55:27.422852 +0000 UTC
package filterer

import (
	log "github.com/golang/glog"
	"runtime"
	"sync"

	"github.com/Jeffail/tunny"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	specs "github.com/joincivil/civil-events-crawler/pkg/contractspecs"
	commongen "github.com/joincivil/civil-events-crawler/pkg/generated/common"
	"github.com/joincivil/civil-events-crawler/pkg/model"

	"github.com/joincivil/go-common/pkg/generated/contract"
	ctime "github.com/joincivil/go-common/pkg/time"
)

func NewGovernmentContractFilterers(contractAddress common.Address) *GovernmentContractFilterers {
	contractFilterers := &GovernmentContractFilterers{
		contractAddress:   contractAddress,
		eventTypes:        commongen.EventTypesGovernmentContract(),
		eventToStartBlock: make(map[string]uint64),
		lastEvents:        make([]*model.Event, 0),
	}
	for _, eventType := range contractFilterers.eventTypes {
		contractFilterers.eventToStartBlock[eventType] = 0
	}
	return contractFilterers
}

type GovernmentContractFilterers struct {
	contractAddress   common.Address
	contract          *contract.GovernmentContract
	eventTypes        []string
	eventToStartBlock map[string]uint64
	lastEvents        []*model.Event
	lastEventsMutex   sync.Mutex
	pastEventsMutex   sync.Mutex
}

func (f *GovernmentContractFilterers) ContractName() string {
	return "GovernmentContract"
}

func (f *GovernmentContractFilterers) ContractAddress() common.Address {
	return f.contractAddress
}

func (f *GovernmentContractFilterers) StartFilterers(client bind.ContractBackend,
	pastEvents []*model.Event, nonSubOnly bool) ([]*model.Event, error) {
	return f.StartGovernmentContractFilterers(client, pastEvents, nonSubOnly)
}

func (f *GovernmentContractFilterers) EventTypes() []string {
	return f.eventTypes
}

func (f *GovernmentContractFilterers) UpdateStartBlock(eventType string, startBlock uint64) {
	f.eventToStartBlock[eventType] = startBlock
}

func (f *GovernmentContractFilterers) LastEvents() []*model.Event {
	return f.lastEvents
}

// StartGovernmentContractFilterers retrieves events for GovernmentContract
func (f *GovernmentContractFilterers) StartGovernmentContractFilterers(client bind.ContractBackend,
	pastEvents []*model.Event, nonSubOnly bool) ([]*model.Event, error) {
	contract, err := contract.NewGovernmentContract(f.contractAddress, client)
	if err != nil {
		log.Errorf("Error initializing StartGovernmentContract: err: %v", err)
		return pastEvents, err
	}
	f.contract = contract

	workerMultiplier := 1
	numWorkers := runtime.NumCPU() * workerMultiplier
	pool := tunny.NewFunc(numWorkers, func(payload interface{}) interface{} {
		f := payload.(func())
		f()
		return nil
	})
	defer pool.Close()

	wg := sync.WaitGroup{}
	resultsChan := make(chan []*model.Event)
	done := make(chan struct{})
	filtsRun := 0

	if !specs.IsEventDisabled("GovernmentContract", "AppellateSet") && (!nonSubOnly || !specs.IsListenerEnabledForEvent("GovernmentContract", "AppellateSet")) {
		wg.Add(1)
		go func() {
			filterFunc := func() {
				startBlock := f.eventToStartBlock["AppellateSet"]
				pevents, e := f.startFilterAppellateSet(startBlock, []*model.Event{})
				if e != nil {
					log.Errorf("Error retrieving AppellateSet: err: %v", e)
					return
				}
				if len(pevents) > 0 {
					f.lastEventsMutex.Lock()
					f.lastEvents = append(f.lastEvents, pevents[len(pevents)-1])
					f.lastEventsMutex.Unlock()
					resultsChan <- pevents
				}
			}
			pool.Process(filterFunc)
			wg.Done()
		}()
		filtsRun++
	}

	if !specs.IsEventDisabled("GovernmentContract", "GovtReparameterizationProposal") && (!nonSubOnly || !specs.IsListenerEnabledForEvent("GovernmentContract", "GovtReparameterizationProposal")) {
		wg.Add(1)
		go func() {
			filterFunc := func() {
				startBlock := f.eventToStartBlock["GovtReparameterizationProposal"]
				pevents, e := f.startFilterGovtReparameterizationProposal(startBlock, []*model.Event{})
				if e != nil {
					log.Errorf("Error retrieving GovtReparameterizationProposal: err: %v", e)
					return
				}
				if len(pevents) > 0 {
					f.lastEventsMutex.Lock()
					f.lastEvents = append(f.lastEvents, pevents[len(pevents)-1])
					f.lastEventsMutex.Unlock()
					resultsChan <- pevents
				}
			}
			pool.Process(filterFunc)
			wg.Done()
		}()
		filtsRun++
	}

	if !specs.IsEventDisabled("GovernmentContract", "NewConstSet") && (!nonSubOnly || !specs.IsListenerEnabledForEvent("GovernmentContract", "NewConstSet")) {
		wg.Add(1)
		go func() {
			filterFunc := func() {
				startBlock := f.eventToStartBlock["NewConstSet"]
				pevents, e := f.startFilterNewConstSet(startBlock, []*model.Event{})
				if e != nil {
					log.Errorf("Error retrieving NewConstSet: err: %v", e)
					return
				}
				if len(pevents) > 0 {
					f.lastEventsMutex.Lock()
					f.lastEvents = append(f.lastEvents, pevents[len(pevents)-1])
					f.lastEventsMutex.Unlock()
					resultsChan <- pevents
				}
			}
			pool.Process(filterFunc)
			wg.Done()
		}()
		filtsRun++
	}

	if !specs.IsEventDisabled("GovernmentContract", "ParameterSet") && (!nonSubOnly || !specs.IsListenerEnabledForEvent("GovernmentContract", "ParameterSet")) {
		wg.Add(1)
		go func() {
			filterFunc := func() {
				startBlock := f.eventToStartBlock["ParameterSet"]
				pevents, e := f.startFilterParameterSet(startBlock, []*model.Event{})
				if e != nil {
					log.Errorf("Error retrieving ParameterSet: err: %v", e)
					return
				}
				if len(pevents) > 0 {
					f.lastEventsMutex.Lock()
					f.lastEvents = append(f.lastEvents, pevents[len(pevents)-1])
					f.lastEventsMutex.Unlock()
					resultsChan <- pevents
				}
			}
			pool.Process(filterFunc)
			wg.Done()
		}()
		filtsRun++
	}

	if !specs.IsEventDisabled("GovernmentContract", "ProposalExpired") && (!nonSubOnly || !specs.IsListenerEnabledForEvent("GovernmentContract", "ProposalExpired")) {
		wg.Add(1)
		go func() {
			filterFunc := func() {
				startBlock := f.eventToStartBlock["ProposalExpired"]
				pevents, e := f.startFilterProposalExpired(startBlock, []*model.Event{})
				if e != nil {
					log.Errorf("Error retrieving ProposalExpired: err: %v", e)
					return
				}
				if len(pevents) > 0 {
					f.lastEventsMutex.Lock()
					f.lastEvents = append(f.lastEvents, pevents[len(pevents)-1])
					f.lastEventsMutex.Unlock()
					resultsChan <- pevents
				}
			}
			pool.Process(filterFunc)
			wg.Done()
		}()
		filtsRun++
	}

	if !specs.IsEventDisabled("GovernmentContract", "ProposalFailed") && (!nonSubOnly || !specs.IsListenerEnabledForEvent("GovernmentContract", "ProposalFailed")) {
		wg.Add(1)
		go func() {
			filterFunc := func() {
				startBlock := f.eventToStartBlock["ProposalFailed"]
				pevents, e := f.startFilterProposalFailed(startBlock, []*model.Event{})
				if e != nil {
					log.Errorf("Error retrieving ProposalFailed: err: %v", e)
					return
				}
				if len(pevents) > 0 {
					f.lastEventsMutex.Lock()
					f.lastEvents = append(f.lastEvents, pevents[len(pevents)-1])
					f.lastEventsMutex.Unlock()
					resultsChan <- pevents
				}
			}
			pool.Process(filterFunc)
			wg.Done()
		}()
		filtsRun++
	}

	if !specs.IsEventDisabled("GovernmentContract", "ProposalPassed") && (!nonSubOnly || !specs.IsListenerEnabledForEvent("GovernmentContract", "ProposalPassed")) {
		wg.Add(1)
		go func() {
			filterFunc := func() {
				startBlock := f.eventToStartBlock["ProposalPassed"]
				pevents, e := f.startFilterProposalPassed(startBlock, []*model.Event{})
				if e != nil {
					log.Errorf("Error retrieving ProposalPassed: err: %v", e)
					return
				}
				if len(pevents) > 0 {
					f.lastEventsMutex.Lock()
					f.lastEvents = append(f.lastEvents, pevents[len(pevents)-1])
					f.lastEventsMutex.Unlock()
					resultsChan <- pevents
				}
			}
			pool.Process(filterFunc)
			wg.Done()
		}()
		filtsRun++
	}

	go func() {
		wg.Wait()
		close(done)
		log.Info("Filtering routines complete")
	}()

Loop:
	for {
		select {
		case <-done:
			break Loop
		case pevents := <-resultsChan:
			f.pastEventsMutex.Lock()
			pastEvents = append(pastEvents, pevents...)
			f.pastEventsMutex.Unlock()
		}
	}
	log.Infof("Total filterers run: %v, events found: %v", filtsRun, len(pastEvents))
	return pastEvents, nil
}

func (f *GovernmentContractFilterers) startFilterAppellateSet(startBlock uint64, pastEvents []*model.Event) ([]*model.Event, error) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	var itr *contract.GovernmentContractAppellateSetIterator
	var err error
	maxRetries := 3
	retry := 0
	for {
		itr, err = f.contract.FilterAppellateSet(
			opts,
		)
		if err == nil {
			break
		}
		if retry >= maxRetries {
			log.Errorf("Failed filter: GovernmentContract AppellateSet for %v: err: %v", f.contractAddress.Hex(), err)
			return pastEvents, err
		}
		log.Infof("Retrying filter: GovernmentContract AppellateSet for %v: err: %v", f.contractAddress.Hex(), err)
		retry++
	}
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("AppellateSet", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	return pastEvents, nil
}

func (f *GovernmentContractFilterers) startFilterGovtReparameterizationProposal(startBlock uint64, pastEvents []*model.Event) ([]*model.Event, error) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	var itr *contract.GovernmentContractGovtReparameterizationProposalIterator
	var err error
	maxRetries := 3
	retry := 0
	for {
		itr, err = f.contract.FilterGovtReparameterizationProposal(
			opts,
		)
		if err == nil {
			break
		}
		if retry >= maxRetries {
			log.Errorf("Failed filter: GovernmentContract GovtReparameterizationProposal for %v: err: %v", f.contractAddress.Hex(), err)
			return pastEvents, err
		}
		log.Infof("Retrying filter: GovernmentContract GovtReparameterizationProposal for %v: err: %v", f.contractAddress.Hex(), err)
		retry++
	}
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("GovtReparameterizationProposal", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	return pastEvents, nil
}

func (f *GovernmentContractFilterers) startFilterNewConstSet(startBlock uint64, pastEvents []*model.Event) ([]*model.Event, error) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	var itr *contract.GovernmentContractNewConstSetIterator
	var err error
	maxRetries := 3
	retry := 0
	for {
		itr, err = f.contract.FilterNewConstSet(
			opts,
		)
		if err == nil {
			break
		}
		if retry >= maxRetries {
			log.Errorf("Failed filter: GovernmentContract NewConstSet for %v: err: %v", f.contractAddress.Hex(), err)
			return pastEvents, err
		}
		log.Infof("Retrying filter: GovernmentContract NewConstSet for %v: err: %v", f.contractAddress.Hex(), err)
		retry++
	}
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("NewConstSet", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	return pastEvents, nil
}

func (f *GovernmentContractFilterers) startFilterParameterSet(startBlock uint64, pastEvents []*model.Event) ([]*model.Event, error) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	var itr *contract.GovernmentContractParameterSetIterator
	var err error
	maxRetries := 3
	retry := 0
	for {
		itr, err = f.contract.FilterParameterSet(
			opts,
		)
		if err == nil {
			break
		}
		if retry >= maxRetries {
			log.Errorf("Failed filter: GovernmentContract ParameterSet for %v: err: %v", f.contractAddress.Hex(), err)
			return pastEvents, err
		}
		log.Infof("Retrying filter: GovernmentContract ParameterSet for %v: err: %v", f.contractAddress.Hex(), err)
		retry++
	}
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("ParameterSet", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	return pastEvents, nil
}

func (f *GovernmentContractFilterers) startFilterProposalExpired(startBlock uint64, pastEvents []*model.Event) ([]*model.Event, error) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	var itr *contract.GovernmentContractProposalExpiredIterator
	var err error
	maxRetries := 3
	retry := 0
	for {
		itr, err = f.contract.FilterProposalExpired(
			opts,
		)
		if err == nil {
			break
		}
		if retry >= maxRetries {
			log.Errorf("Failed filter: GovernmentContract ProposalExpired for %v: err: %v", f.contractAddress.Hex(), err)
			return pastEvents, err
		}
		log.Infof("Retrying filter: GovernmentContract ProposalExpired for %v: err: %v", f.contractAddress.Hex(), err)
		retry++
	}
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("ProposalExpired", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	return pastEvents, nil
}

func (f *GovernmentContractFilterers) startFilterProposalFailed(startBlock uint64, pastEvents []*model.Event) ([]*model.Event, error) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	var itr *contract.GovernmentContractProposalFailedIterator
	var err error
	maxRetries := 3
	retry := 0
	for {
		itr, err = f.contract.FilterProposalFailed(
			opts,
		)
		if err == nil {
			break
		}
		if retry >= maxRetries {
			log.Errorf("Failed filter: GovernmentContract ProposalFailed for %v: err: %v", f.contractAddress.Hex(), err)
			return pastEvents, err
		}
		log.Infof("Retrying filter: GovernmentContract ProposalFailed for %v: err: %v", f.contractAddress.Hex(), err)
		retry++
	}
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("ProposalFailed", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	return pastEvents, nil
}

func (f *GovernmentContractFilterers) startFilterProposalPassed(startBlock uint64, pastEvents []*model.Event) ([]*model.Event, error) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	var itr *contract.GovernmentContractProposalPassedIterator
	var err error
	maxRetries := 3
	retry := 0
	for {
		itr, err = f.contract.FilterProposalPassed(
			opts,
		)
		if err == nil {
			break
		}
		if retry >= maxRetries {
			log.Errorf("Failed filter: GovernmentContract ProposalPassed for %v: err: %v", f.contractAddress.Hex(), err)
			return pastEvents, err
		}
		log.Infof("Retrying filter: GovernmentContract ProposalPassed for %v: err: %v", f.contractAddress.Hex(), err)
		retry++
	}
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("ProposalPassed", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	return pastEvents, nil
}
