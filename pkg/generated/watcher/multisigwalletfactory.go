// Code generated by 'gen/watchergen'  DO NOT EDIT.
// IT SHOULD NOT BE EDITED BY HAND AS ANY CHANGES MAY BE OVERWRITTEN
// Please reference 'listener/watchergen' for more details
// File was generated at 2020-02-18 16:45:16.216557 +0000 UTC
package watcher

import (
	// "fmt"
	"context"
	"time"

	"github.com/davecgh/go-spew/spew"
	log "github.com/golang/glog"
	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	specs "github.com/joincivil/civil-events-crawler/pkg/contractspecs"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/utils"

	"github.com/joincivil/go-common/pkg/generated/contract"
	ctime "github.com/joincivil/go-common/pkg/time"
)

func NewMultiSigWalletFactoryContractWatchers(contractAddress common.Address) *MultiSigWalletFactoryContractWatchers {
	return &MultiSigWalletFactoryContractWatchers{
		contractAddress: contractAddress,
	}
}

type MultiSigWalletFactoryContractWatchers struct {
	errors          chan error
	contractAddress common.Address
	contract        *contract.MultiSigWalletFactoryContract
	activeSubs      []utils.WatcherSubscription
}

func (w *MultiSigWalletFactoryContractWatchers) ContractAddress() common.Address {
	return w.contractAddress
}

func (w *MultiSigWalletFactoryContractWatchers) ContractName() string {
	return "MultiSigWalletFactoryContract"
}

func (w *MultiSigWalletFactoryContractWatchers) cancelFunc(cancelFn context.CancelFunc, killCancel <-chan struct{}) {
}

func (w *MultiSigWalletFactoryContractWatchers) StopWatchers(unsub bool) error {
	if unsub {
		for _, sub := range w.activeSubs {
			sub.Unsubscribe()
		}
	}
	w.activeSubs = nil
	w.contract = nil
	return nil
}

func (w *MultiSigWalletFactoryContractWatchers) StartWatchers(client bind.ContractBackend,
	eventRecvChan chan *model.Event, errs chan error) ([]utils.WatcherSubscription, error) {
	return w.StartMultiSigWalletFactoryContractWatchers(client, eventRecvChan, errs)
}

// StartMultiSigWalletFactoryContractWatchers starts up the event watchers for MultiSigWalletFactoryContract
func (w *MultiSigWalletFactoryContractWatchers) StartMultiSigWalletFactoryContractWatchers(client bind.ContractBackend,
	eventRecvChan chan *model.Event, errs chan error) ([]utils.WatcherSubscription, error) {
	w.errors = errs
	contract, err := contract.NewMultiSigWalletFactoryContract(w.contractAddress, client)
	if err != nil {
		log.Errorf("Error initializing StartMultiSigWalletFactoryContract: err: %v", err)
		return nil, errors.Wrap(err, "error initializing StartMultiSigWalletFactoryContract")
	}
	w.contract = contract

	var sub utils.WatcherSubscription
	subs := []utils.WatcherSubscription{}

	if specs.IsListenerEnabledForEvent("MultiSigWalletFactoryContract", "ContractInstantiation") {
		sub, err = w.startWatchContractInstantiation(eventRecvChan)
		if err != nil {
			return nil, errors.WithMessage(err, "error starting startContractInstantiation")
		}
		subs = append(subs, sub)
	}

	w.activeSubs = subs
	return subs, nil
}

func (w *MultiSigWalletFactoryContractWatchers) startWatchContractInstantiation(eventRecvChan chan *model.Event) (utils.WatcherSubscription, error) {
	killCancelTimeoutSecs := 10
	return utils.NewWatcherSubscription("WatchContractInstantiation", func(quit <-chan struct{}) error {
		startupFn := func() (utils.WatcherSubscription, chan *contract.MultiSigWalletFactoryContractContractInstantiation, error) {
			ctx := context.Background()
			ctx, cancelFn := context.WithCancel(ctx)
			opts := &bind.WatchOpts{Context: ctx}
			killCancel := make(chan struct{})
			// 10 sec timeout mechanism for starting up watcher
			go func(cancelFn context.CancelFunc, killCancel <-chan struct{}) {
				select {
				case <-time.After(time.Duration(killCancelTimeoutSecs) * time.Second):
					log.Errorf("WatchContractInstantiation start timeout, cancelling...")
					cancelFn()
				case <-killCancel:
				}
			}(cancelFn, killCancel)
			recvChan := make(chan *contract.MultiSigWalletFactoryContractContractInstantiation)
			sub, err := w.contract.WatchContractInstantiation(
				opts,
				recvChan,
			)
			close(killCancel)
			if err != nil {
				if sub != nil {
					log.Infof("startupFn: Unsubscribing WatchContractInstantiation: addr: %v", w.contractAddress.Hex())
					sub.Unsubscribe()
					sub = nil
				}
				return nil, nil, errors.Wrap(err, "startupFn: error starting WatchContractInstantiation")
			}
			return sub, recvChan, nil
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchContractInstantiation: addr: %v, %v", w.contractAddress.Hex(), err)
			if sub != nil {
				sub.Unsubscribe()
				close(recvChan)
				sub = nil
			}
			w.errors <- err
			return err
		}
		defer func() {
			sub.Unsubscribe()
			close(recvChan)
			sub = nil
		}()
		log.Infof("Starting up WatchContractInstantiation: addr: %v", w.contractAddress.Hex())
		for {
			select {
			case event := <-recvChan:
				if log.V(2) {
					log.Infof("Received event on WatchContractInstantiation: %v", spew.Sprintf("%#+v", event))
				} else {
					log.Info("Received event on WatchContractInstantiation")
				}
				modelEvent, err := model.NewEventFromContractEvent("ContractInstantiation", w.ContractName(), w.contractAddress, event, ctime.CurrentEpochSecsInInt64(), model.Watcher)
				if err != nil {
					log.Errorf("Error creating new event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- modelEvent:
					if log.V(2) {
						log.Infof("Sent event to eventRecvChan on WatchContractInstantiation: %v", spew.Sprintf("%#+v", event))
					} else {
						log.Info("Sent event to eventRecvChan on WatchContractInstantiation")
					}
				case err := <-sub.Err():
					log.Errorf("Error with WatchContractInstantiation: addr: %v, fatal (a): %v", w.contractAddress.Hex(), err)
					err = errors.Wrap(err, "error with WatchContractInstantiation")
					w.errors <- err
					return err
				case <-quit:
					log.Infof("Quit WatchContractInstantiation (a): addr: %v", w.contractAddress.Hex())
					return nil
				}
			case err := <-sub.Err():
				log.Errorf("Error with WatchContractInstantiation: addr: %v, fatal (b): %v", w.contractAddress.Hex(), err)
				err = errors.Wrap(err, "error with WatchContractInstantiation")
				w.errors <- err
				return err
			case <-quit:
				log.Infof("Quitting loop for WatchContractInstantiation: addr: %v", w.contractAddress.Hex())
				return nil
			}
		}
	}), nil
}
