// Code generated by 'gen/eventhandlergen.go'  DO NOT EDIT.
// IT SHOULD NOT BE EDITED BY HAND AS ANY CHANGES MAY BE OVERWRITTEN
// Please reference 'gen/filterergen_template.go' for more details
// File was generated at 2020-02-18 16:55:29.970844 +0000 UTC
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

	"math/big"
)

func NewMultiSigWalletContractFilterers(contractAddress common.Address) *MultiSigWalletContractFilterers {
	contractFilterers := &MultiSigWalletContractFilterers{
		contractAddress:   contractAddress,
		eventTypes:        commongen.EventTypesMultiSigWalletContract(),
		eventToStartBlock: make(map[string]uint64),
		lastEvents:        make([]*model.Event, 0),
	}
	for _, eventType := range contractFilterers.eventTypes {
		contractFilterers.eventToStartBlock[eventType] = 0
	}
	return contractFilterers
}

type MultiSigWalletContractFilterers struct {
	contractAddress   common.Address
	contract          *contract.MultiSigWalletContract
	eventTypes        []string
	eventToStartBlock map[string]uint64
	lastEvents        []*model.Event
	lastEventsMutex   sync.Mutex
	pastEventsMutex   sync.Mutex
}

func (f *MultiSigWalletContractFilterers) ContractName() string {
	return "MultiSigWalletContract"
}

func (f *MultiSigWalletContractFilterers) ContractAddress() common.Address {
	return f.contractAddress
}

func (f *MultiSigWalletContractFilterers) StartFilterers(client bind.ContractBackend,
	pastEvents []*model.Event, nonSubOnly bool) ([]*model.Event, error) {
	return f.StartMultiSigWalletContractFilterers(client, pastEvents, nonSubOnly)
}

func (f *MultiSigWalletContractFilterers) EventTypes() []string {
	return f.eventTypes
}

func (f *MultiSigWalletContractFilterers) UpdateStartBlock(eventType string, startBlock uint64) {
	f.eventToStartBlock[eventType] = startBlock
}

func (f *MultiSigWalletContractFilterers) LastEvents() []*model.Event {
	return f.lastEvents
}

// StartMultiSigWalletContractFilterers retrieves events for MultiSigWalletContract
func (f *MultiSigWalletContractFilterers) StartMultiSigWalletContractFilterers(client bind.ContractBackend,
	pastEvents []*model.Event, nonSubOnly bool) ([]*model.Event, error) {
	contract, err := contract.NewMultiSigWalletContract(f.contractAddress, client)
	if err != nil {
		log.Errorf("Error initializing StartMultiSigWalletContract: err: %v", err)
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

	if !specs.IsEventDisabled("MultiSigWalletContract", "Confirmation") && (!nonSubOnly || !specs.IsListenerEnabledForEvent("MultiSigWalletContract", "Confirmation")) {
		wg.Add(1)
		go func() {
			filterFunc := func() {
				startBlock := f.eventToStartBlock["Confirmation"]
				pevents, e := f.startFilterConfirmation(startBlock, []*model.Event{})
				if e != nil {
					log.Errorf("Error retrieving Confirmation: err: %v", e)
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

	if !specs.IsEventDisabled("MultiSigWalletContract", "Deposit") && (!nonSubOnly || !specs.IsListenerEnabledForEvent("MultiSigWalletContract", "Deposit")) {
		wg.Add(1)
		go func() {
			filterFunc := func() {
				startBlock := f.eventToStartBlock["Deposit"]
				pevents, e := f.startFilterDeposit(startBlock, []*model.Event{})
				if e != nil {
					log.Errorf("Error retrieving Deposit: err: %v", e)
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

	if !specs.IsEventDisabled("MultiSigWalletContract", "Execution") && (!nonSubOnly || !specs.IsListenerEnabledForEvent("MultiSigWalletContract", "Execution")) {
		wg.Add(1)
		go func() {
			filterFunc := func() {
				startBlock := f.eventToStartBlock["Execution"]
				pevents, e := f.startFilterExecution(startBlock, []*model.Event{})
				if e != nil {
					log.Errorf("Error retrieving Execution: err: %v", e)
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

	if !specs.IsEventDisabled("MultiSigWalletContract", "ExecutionFailure") && (!nonSubOnly || !specs.IsListenerEnabledForEvent("MultiSigWalletContract", "ExecutionFailure")) {
		wg.Add(1)
		go func() {
			filterFunc := func() {
				startBlock := f.eventToStartBlock["ExecutionFailure"]
				pevents, e := f.startFilterExecutionFailure(startBlock, []*model.Event{})
				if e != nil {
					log.Errorf("Error retrieving ExecutionFailure: err: %v", e)
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

	if !specs.IsEventDisabled("MultiSigWalletContract", "OwnerAddition") && (!nonSubOnly || !specs.IsListenerEnabledForEvent("MultiSigWalletContract", "OwnerAddition")) {
		wg.Add(1)
		go func() {
			filterFunc := func() {
				startBlock := f.eventToStartBlock["OwnerAddition"]
				pevents, e := f.startFilterOwnerAddition(startBlock, []*model.Event{})
				if e != nil {
					log.Errorf("Error retrieving OwnerAddition: err: %v", e)
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

	if !specs.IsEventDisabled("MultiSigWalletContract", "OwnerRemoval") && (!nonSubOnly || !specs.IsListenerEnabledForEvent("MultiSigWalletContract", "OwnerRemoval")) {
		wg.Add(1)
		go func() {
			filterFunc := func() {
				startBlock := f.eventToStartBlock["OwnerRemoval"]
				pevents, e := f.startFilterOwnerRemoval(startBlock, []*model.Event{})
				if e != nil {
					log.Errorf("Error retrieving OwnerRemoval: err: %v", e)
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

	if !specs.IsEventDisabled("MultiSigWalletContract", "RequirementChange") && (!nonSubOnly || !specs.IsListenerEnabledForEvent("MultiSigWalletContract", "RequirementChange")) {
		wg.Add(1)
		go func() {
			filterFunc := func() {
				startBlock := f.eventToStartBlock["RequirementChange"]
				pevents, e := f.startFilterRequirementChange(startBlock, []*model.Event{})
				if e != nil {
					log.Errorf("Error retrieving RequirementChange: err: %v", e)
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

	if !specs.IsEventDisabled("MultiSigWalletContract", "Revocation") && (!nonSubOnly || !specs.IsListenerEnabledForEvent("MultiSigWalletContract", "Revocation")) {
		wg.Add(1)
		go func() {
			filterFunc := func() {
				startBlock := f.eventToStartBlock["Revocation"]
				pevents, e := f.startFilterRevocation(startBlock, []*model.Event{})
				if e != nil {
					log.Errorf("Error retrieving Revocation: err: %v", e)
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

	if !specs.IsEventDisabled("MultiSigWalletContract", "Submission") && (!nonSubOnly || !specs.IsListenerEnabledForEvent("MultiSigWalletContract", "Submission")) {
		wg.Add(1)
		go func() {
			filterFunc := func() {
				startBlock := f.eventToStartBlock["Submission"]
				pevents, e := f.startFilterSubmission(startBlock, []*model.Event{})
				if e != nil {
					log.Errorf("Error retrieving Submission: err: %v", e)
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

func (f *MultiSigWalletContractFilterers) startFilterConfirmation(startBlock uint64, pastEvents []*model.Event) ([]*model.Event, error) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	var itr *contract.MultiSigWalletContractConfirmationIterator
	var err error
	maxRetries := 3
	retry := 0
	for {
		itr, err = f.contract.FilterConfirmation(
			opts,
			[]common.Address{},
			[]*big.Int{},
		)
		if err == nil {
			break
		}
		if retry >= maxRetries {
			log.Errorf("Failed filter: MultiSigWalletContract Confirmation for %v: err: %v", f.contractAddress.Hex(), err)
			return pastEvents, err
		}
		log.Infof("Retrying filter: MultiSigWalletContract Confirmation for %v: err: %v", f.contractAddress.Hex(), err)
		retry++
	}
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("Confirmation", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	return pastEvents, nil
}

func (f *MultiSigWalletContractFilterers) startFilterDeposit(startBlock uint64, pastEvents []*model.Event) ([]*model.Event, error) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	var itr *contract.MultiSigWalletContractDepositIterator
	var err error
	maxRetries := 3
	retry := 0
	for {
		itr, err = f.contract.FilterDeposit(
			opts,
			[]common.Address{},
		)
		if err == nil {
			break
		}
		if retry >= maxRetries {
			log.Errorf("Failed filter: MultiSigWalletContract Deposit for %v: err: %v", f.contractAddress.Hex(), err)
			return pastEvents, err
		}
		log.Infof("Retrying filter: MultiSigWalletContract Deposit for %v: err: %v", f.contractAddress.Hex(), err)
		retry++
	}
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("Deposit", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	return pastEvents, nil
}

func (f *MultiSigWalletContractFilterers) startFilterExecution(startBlock uint64, pastEvents []*model.Event) ([]*model.Event, error) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	var itr *contract.MultiSigWalletContractExecutionIterator
	var err error
	maxRetries := 3
	retry := 0
	for {
		itr, err = f.contract.FilterExecution(
			opts,
			[]*big.Int{},
		)
		if err == nil {
			break
		}
		if retry >= maxRetries {
			log.Errorf("Failed filter: MultiSigWalletContract Execution for %v: err: %v", f.contractAddress.Hex(), err)
			return pastEvents, err
		}
		log.Infof("Retrying filter: MultiSigWalletContract Execution for %v: err: %v", f.contractAddress.Hex(), err)
		retry++
	}
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("Execution", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	return pastEvents, nil
}

func (f *MultiSigWalletContractFilterers) startFilterExecutionFailure(startBlock uint64, pastEvents []*model.Event) ([]*model.Event, error) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	var itr *contract.MultiSigWalletContractExecutionFailureIterator
	var err error
	maxRetries := 3
	retry := 0
	for {
		itr, err = f.contract.FilterExecutionFailure(
			opts,
			[]*big.Int{},
		)
		if err == nil {
			break
		}
		if retry >= maxRetries {
			log.Errorf("Failed filter: MultiSigWalletContract ExecutionFailure for %v: err: %v", f.contractAddress.Hex(), err)
			return pastEvents, err
		}
		log.Infof("Retrying filter: MultiSigWalletContract ExecutionFailure for %v: err: %v", f.contractAddress.Hex(), err)
		retry++
	}
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("ExecutionFailure", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	return pastEvents, nil
}

func (f *MultiSigWalletContractFilterers) startFilterOwnerAddition(startBlock uint64, pastEvents []*model.Event) ([]*model.Event, error) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	var itr *contract.MultiSigWalletContractOwnerAdditionIterator
	var err error
	maxRetries := 3
	retry := 0
	for {
		itr, err = f.contract.FilterOwnerAddition(
			opts,
			[]common.Address{},
		)
		if err == nil {
			break
		}
		if retry >= maxRetries {
			log.Errorf("Failed filter: MultiSigWalletContract OwnerAddition for %v: err: %v", f.contractAddress.Hex(), err)
			return pastEvents, err
		}
		log.Infof("Retrying filter: MultiSigWalletContract OwnerAddition for %v: err: %v", f.contractAddress.Hex(), err)
		retry++
	}
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("OwnerAddition", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	return pastEvents, nil
}

func (f *MultiSigWalletContractFilterers) startFilterOwnerRemoval(startBlock uint64, pastEvents []*model.Event) ([]*model.Event, error) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	var itr *contract.MultiSigWalletContractOwnerRemovalIterator
	var err error
	maxRetries := 3
	retry := 0
	for {
		itr, err = f.contract.FilterOwnerRemoval(
			opts,
			[]common.Address{},
		)
		if err == nil {
			break
		}
		if retry >= maxRetries {
			log.Errorf("Failed filter: MultiSigWalletContract OwnerRemoval for %v: err: %v", f.contractAddress.Hex(), err)
			return pastEvents, err
		}
		log.Infof("Retrying filter: MultiSigWalletContract OwnerRemoval for %v: err: %v", f.contractAddress.Hex(), err)
		retry++
	}
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("OwnerRemoval", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	return pastEvents, nil
}

func (f *MultiSigWalletContractFilterers) startFilterRequirementChange(startBlock uint64, pastEvents []*model.Event) ([]*model.Event, error) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	var itr *contract.MultiSigWalletContractRequirementChangeIterator
	var err error
	maxRetries := 3
	retry := 0
	for {
		itr, err = f.contract.FilterRequirementChange(
			opts,
		)
		if err == nil {
			break
		}
		if retry >= maxRetries {
			log.Errorf("Failed filter: MultiSigWalletContract RequirementChange for %v: err: %v", f.contractAddress.Hex(), err)
			return pastEvents, err
		}
		log.Infof("Retrying filter: MultiSigWalletContract RequirementChange for %v: err: %v", f.contractAddress.Hex(), err)
		retry++
	}
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("RequirementChange", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	return pastEvents, nil
}

func (f *MultiSigWalletContractFilterers) startFilterRevocation(startBlock uint64, pastEvents []*model.Event) ([]*model.Event, error) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	var itr *contract.MultiSigWalletContractRevocationIterator
	var err error
	maxRetries := 3
	retry := 0
	for {
		itr, err = f.contract.FilterRevocation(
			opts,
			[]common.Address{},
			[]*big.Int{},
		)
		if err == nil {
			break
		}
		if retry >= maxRetries {
			log.Errorf("Failed filter: MultiSigWalletContract Revocation for %v: err: %v", f.contractAddress.Hex(), err)
			return pastEvents, err
		}
		log.Infof("Retrying filter: MultiSigWalletContract Revocation for %v: err: %v", f.contractAddress.Hex(), err)
		retry++
	}
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("Revocation", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	return pastEvents, nil
}

func (f *MultiSigWalletContractFilterers) startFilterSubmission(startBlock uint64, pastEvents []*model.Event) ([]*model.Event, error) {
	var opts = &bind.FilterOpts{
		Start: startBlock,
	}

	var itr *contract.MultiSigWalletContractSubmissionIterator
	var err error
	maxRetries := 3
	retry := 0
	for {
		itr, err = f.contract.FilterSubmission(
			opts,
			[]*big.Int{},
		)
		if err == nil {
			break
		}
		if retry >= maxRetries {
			log.Errorf("Failed filter: MultiSigWalletContract Submission for %v: err: %v", f.contractAddress.Hex(), err)
			return pastEvents, err
		}
		log.Infof("Retrying filter: MultiSigWalletContract Submission for %v: err: %v", f.contractAddress.Hex(), err)
		retry++
	}
	nextEvent := itr.Next()
	for nextEvent {
		modelEvent, err := model.NewEventFromContractEvent("Submission", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
		if err != nil {
			log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
			continue
		}
		pastEvents = append(pastEvents, modelEvent)
		nextEvent = itr.Next()
	}
	return pastEvents, nil
}
