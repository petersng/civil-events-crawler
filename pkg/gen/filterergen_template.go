// Package gen contains all the components for code generation.
package gen

const filtererTmpl = `
// Code generated by 'gen/eventhandlergen.go'  DO NOT EDIT.
// IT SHOULD NOT BE EDITED BY HAND AS ANY CHANGES MAY BE OVERWRITTEN
// Please reference 'gen/filterergen_template.go' for more details
// File was generated at {{.GenTime}}
package {{.PackageName}}

import (
    log "github.com/golang/glog"
    "sync"
    "runtime"

    "github.com/Jeffail/tunny"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"


    specs "github.com/joincivil/civil-events-crawler/pkg/contractspecs"
    "github.com/joincivil/civil-events-crawler/pkg/model"
    commongen "github.com/joincivil/civil-events-crawler/pkg/generated/common"

    ctime "github.com/joincivil/go-common/pkg/time"
{{if .ContractImportPath -}}
    "{{.ContractImportPath}}"
{{- end}}
{{if .AdditionalImports -}}
{{- range .AdditionalImports}}
    "{{.}}"
{{- end}}
{{- end}}
)

func New{{.ContractTypeName}}Filterers(contractAddress common.Address) *{{.ContractTypeName}}Filterers {
    contractFilterers := &{{.ContractTypeName}}Filterers{
        contractAddress: contractAddress,
        eventTypes: commongen.EventTypes{{.ContractTypeName}}(),
        eventToStartBlock: make(map[string]uint64),
        lastEvents: make([]*model.Event, 0),
    }
    for _, eventType := range contractFilterers.eventTypes {
        contractFilterers.eventToStartBlock[eventType] = {{.DefaultStartBlock}}
    }
    return contractFilterers
}

type {{.ContractTypeName}}Filterers struct {
    contractAddress common.Address
    contract *{{.ContractTypePackage}}.{{.ContractTypeName}}
    eventTypes []string
    eventToStartBlock map[string]uint64
    lastEvents  []*model.Event
    lastEventsMutex sync.Mutex
    pastEventsMutex sync.Mutex
}

func (f *{{.ContractTypeName}}Filterers) ContractName() string {
    return "{{.ContractTypeName}}"
}

func (f *{{.ContractTypeName}}Filterers) ContractAddress() common.Address {
    return f.contractAddress
}

func (f *{{.ContractTypeName}}Filterers) StartFilterers(client bind.ContractBackend,
    pastEvents []*model.Event, nonSubOnly bool) ([]*model.Event, error) {
    return f.Start{{.ContractTypeName}}Filterers(client, pastEvents, nonSubOnly)
}

func (f *{{.ContractTypeName}}Filterers) EventTypes() []string {
    return f.eventTypes
}

func (f *{{.ContractTypeName}}Filterers) UpdateStartBlock(eventType string, startBlock uint64) {
    f.eventToStartBlock[eventType] = startBlock
}

func (f *{{.ContractTypeName}}Filterers) LastEvents() []*model.Event {
    return f.lastEvents
}

// Start{{.ContractTypeName}}Filterers retrieves events for {{.ContractTypeName}}
func (f *{{.ContractTypeName}}Filterers) Start{{.ContractTypeName}}Filterers(client bind.ContractBackend,
    pastEvents []*model.Event, nonSubOnly bool) ([]*model.Event, error) {
    contract, err := {{.ContractTypePackage}}.New{{.ContractTypeName}}(f.contractAddress, client)
    if err != nil {
        log.Errorf("Error initializing Start{{.ContractTypeName}}: err: %v", err)
        return pastEvents, err
    }
    f.contract = contract

    workerMultiplier := 1
    numWorkers := runtime.NumCPU() * workerMultiplier
    log.Infof("Filter worker #: %v", numWorkers)
    pool := tunny.NewFunc(numWorkers, func(payload interface{}) interface{} {
        f := payload.(func())
        f()
        return nil
    })
    defer pool.Close()

    wg := sync.WaitGroup{}
    resultsChan := make(chan []*model.Event)
    done := make(chan bool)
    filtsRun := 0


{{if .EventHandlers -}}
{{- range .EventHandlers}}

    if !specs.IsEventDisabled("{{$.ContractTypeName}}", "{{.EventMethod}}") && (!nonSubOnly || !specs.IsListenerEnabledForEvent("{{$.ContractTypeName}}", "{{.EventMethod}}")){
        wg.Add(1)
        go func() {
            filterFunc := func() {
                startBlock := f.eventToStartBlock["{{.EventMethod}}"]
                pevents, e := f.startFilter{{.EventMethod}}(startBlock, []*model.Event{})
                if e != nil {
                    log.Errorf("Error retrieving {{.EventMethod}}: err: %v", e)
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


{{- end}}
{{- end}}

    go func() {
        wg.Wait()
        done <- true
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

{{if .EventHandlers -}}
{{- range .EventHandlers}}

func (f *{{$.ContractTypeName}}Filterers) startFilter{{.EventMethod}}(startBlock uint64, pastEvents []*model.Event) ([]*model.Event, error) {
    var opts = &bind.FilterOpts{
        Start: startBlock,
    }

    log.Infof("Filtering {{$.ContractTypeName}} {{.EventMethod}} for %v at block %v", f.contractAddress.Hex(), startBlock)
	var itr *contract.{{$.ContractTypeName}}{{.EventMethod}}Iterator
	var err error
	maxRetries := 3
    retry := 0
    for {
        itr, err = f.contract.Filter{{.EventMethod}}(
            opts,
        {{- if .ParamValues -}}
        {{range .ParamValues}}
            []{{.Type}}{},
        {{- end}}
        {{end}}
        )
		if err == nil {
            log.Infof("Done filter: {{$.ContractTypeName}} {{.EventMethod}} for %v", f.contractAddress.Hex())
			break
		}
		if retry >= maxRetries {
            log.Errorf("Failed filter: {{$.ContractTypeName}} {{.EventMethod}} for %v: err: %v", f.contractAddress.Hex(), err)
			return pastEvents, err
		}
        log.Infof("Retrying filter: {{$.ContractTypeName}} {{.EventMethod}} for %v: err: %v", f.contractAddress.Hex(), err)
		retry++
    }
    beforeCount := len(pastEvents)
    nextEvent := itr.Next()
    for nextEvent {
        modelEvent, err := model.NewEventFromContractEvent("{{.EventMethod}}", f.ContractName(), f.contractAddress, itr.Event, ctime.CurrentEpochSecsInInt64(), model.Filterer)
        if err != nil {
            log.Errorf("Error creating new event: event: %v, err: %v", itr.Event, err)
            continue
        }
        pastEvents = append(pastEvents, modelEvent)
        nextEvent = itr.Next()
    }
    numEventsAdded := len(pastEvents) - beforeCount
    log.Infof("{{$.ContractTypeName}} {{.EventMethod}} added: %v", numEventsAdded)
    return pastEvents, nil
}

{{- end}}
{{- end}}
`
