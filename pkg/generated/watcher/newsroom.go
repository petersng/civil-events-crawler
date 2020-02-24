// Code generated by 'gen/watchergen'  DO NOT EDIT.
// IT SHOULD NOT BE EDITED BY HAND AS ANY CHANGES MAY BE OVERWRITTEN
// Please reference 'listener/watchergen' for more details
// File was generated at 2020-02-18 16:45:00.651718 +0000 UTC
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

	"math/big"
)

func NewNewsroomContractWatchers(contractAddress common.Address) *NewsroomContractWatchers {
	return &NewsroomContractWatchers{
		contractAddress: contractAddress,
	}
}

type NewsroomContractWatchers struct {
	errors          chan error
	contractAddress common.Address
	contract        *contract.NewsroomContract
	activeSubs      []utils.WatcherSubscription
}

func (w *NewsroomContractWatchers) ContractAddress() common.Address {
	return w.contractAddress
}

func (w *NewsroomContractWatchers) ContractName() string {
	return "NewsroomContract"
}

func (w *NewsroomContractWatchers) cancelFunc(cancelFn context.CancelFunc, killCancel <-chan struct{}) {
}

func (w *NewsroomContractWatchers) StopWatchers(unsub bool) error {
	if unsub {
		for _, sub := range w.activeSubs {
			sub.Unsubscribe()
		}
	}
	w.activeSubs = nil
	w.contract = nil
	return nil
}

func (w *NewsroomContractWatchers) StartWatchers(client bind.ContractBackend,
	eventRecvChan chan *model.Event, errs chan error) ([]utils.WatcherSubscription, error) {
	return w.StartNewsroomContractWatchers(client, eventRecvChan, errs)
}

// StartNewsroomContractWatchers starts up the event watchers for NewsroomContract
func (w *NewsroomContractWatchers) StartNewsroomContractWatchers(client bind.ContractBackend,
	eventRecvChan chan *model.Event, errs chan error) ([]utils.WatcherSubscription, error) {
	w.errors = errs
	contract, err := contract.NewNewsroomContract(w.contractAddress, client)
	if err != nil {
		log.Errorf("Error initializing StartNewsroomContract: err: %v", err)
		return nil, errors.Wrap(err, "error initializing StartNewsroomContract")
	}
	w.contract = contract

	var sub utils.WatcherSubscription
	subs := []utils.WatcherSubscription{}

	if specs.IsListenerEnabledForEvent("NewsroomContract", "ContentPublished") {
		sub, err = w.startWatchContentPublished(eventRecvChan)
		if err != nil {
			return nil, errors.WithMessage(err, "error starting startContentPublished")
		}
		subs = append(subs, sub)
	}

	if specs.IsListenerEnabledForEvent("NewsroomContract", "NameChanged") {
		sub, err = w.startWatchNameChanged(eventRecvChan)
		if err != nil {
			return nil, errors.WithMessage(err, "error starting startNameChanged")
		}
		subs = append(subs, sub)
	}

	if specs.IsListenerEnabledForEvent("NewsroomContract", "OwnershipRenounced") {
		sub, err = w.startWatchOwnershipRenounced(eventRecvChan)
		if err != nil {
			return nil, errors.WithMessage(err, "error starting startOwnershipRenounced")
		}
		subs = append(subs, sub)
	}

	if specs.IsListenerEnabledForEvent("NewsroomContract", "OwnershipTransferred") {
		sub, err = w.startWatchOwnershipTransferred(eventRecvChan)
		if err != nil {
			return nil, errors.WithMessage(err, "error starting startOwnershipTransferred")
		}
		subs = append(subs, sub)
	}

	if specs.IsListenerEnabledForEvent("NewsroomContract", "RevisionSigned") {
		sub, err = w.startWatchRevisionSigned(eventRecvChan)
		if err != nil {
			return nil, errors.WithMessage(err, "error starting startRevisionSigned")
		}
		subs = append(subs, sub)
	}

	if specs.IsListenerEnabledForEvent("NewsroomContract", "RevisionUpdated") {
		sub, err = w.startWatchRevisionUpdated(eventRecvChan)
		if err != nil {
			return nil, errors.WithMessage(err, "error starting startRevisionUpdated")
		}
		subs = append(subs, sub)
	}

	if specs.IsListenerEnabledForEvent("NewsroomContract", "RoleAdded") {
		sub, err = w.startWatchRoleAdded(eventRecvChan)
		if err != nil {
			return nil, errors.WithMessage(err, "error starting startRoleAdded")
		}
		subs = append(subs, sub)
	}

	if specs.IsListenerEnabledForEvent("NewsroomContract", "RoleRemoved") {
		sub, err = w.startWatchRoleRemoved(eventRecvChan)
		if err != nil {
			return nil, errors.WithMessage(err, "error starting startRoleRemoved")
		}
		subs = append(subs, sub)
	}

	w.activeSubs = subs
	return subs, nil
}

func (w *NewsroomContractWatchers) startWatchContentPublished(eventRecvChan chan *model.Event) (utils.WatcherSubscription, error) {
	killCancelTimeoutSecs := 10
	return utils.NewWatcherSubscription("WatchContentPublished", func(quit <-chan struct{}) error {
		startupFn := func() (utils.WatcherSubscription, chan *contract.NewsroomContractContentPublished, error) {
			ctx := context.Background()
			ctx, cancelFn := context.WithCancel(ctx)
			opts := &bind.WatchOpts{Context: ctx}
			killCancel := make(chan struct{})
			// 10 sec timeout mechanism for starting up watcher
			go func(cancelFn context.CancelFunc, killCancel <-chan struct{}) {
				select {
				case <-time.After(time.Duration(killCancelTimeoutSecs) * time.Second):
					log.Errorf("WatchContentPublished start timeout, cancelling...")
					cancelFn()
				case <-killCancel:
				}
			}(cancelFn, killCancel)
			recvChan := make(chan *contract.NewsroomContractContentPublished)
			sub, err := w.contract.WatchContentPublished(
				opts,
				recvChan,
				[]common.Address{},
				[]*big.Int{},
			)
			close(killCancel)
			if err != nil {
				if sub != nil {
					log.Infof("startupFn: Unsubscribing WatchContentPublished: addr: %v", w.contractAddress.Hex())
					sub.Unsubscribe()
					sub = nil
				}
				return nil, nil, errors.Wrap(err, "startupFn: error starting WatchContentPublished")
			}
			return sub, recvChan, nil
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchContentPublished: addr: %v, %v", w.contractAddress.Hex(), err)
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
		log.Infof("Starting up WatchContentPublished: addr: %v", w.contractAddress.Hex())
		for {
			select {
			case event := <-recvChan:
				if log.V(2) {
					log.Infof("Received event on WatchContentPublished: %v", spew.Sprintf("%#+v", event))
				} else {
					log.Info("Received event on WatchContentPublished")
				}
				modelEvent, err := model.NewEventFromContractEvent("ContentPublished", w.ContractName(), w.contractAddress, event, ctime.CurrentEpochSecsInInt64(), model.Watcher)
				if err != nil {
					log.Errorf("Error creating new event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- modelEvent:
					if log.V(2) {
						log.Infof("Sent event to eventRecvChan on WatchContentPublished: %v", spew.Sprintf("%#+v", event))
					} else {
						log.Info("Sent event to eventRecvChan on WatchContentPublished")
					}
				case err := <-sub.Err():
					log.Errorf("Error with WatchContentPublished: addr: %v, fatal (a): %v", w.contractAddress.Hex(), err)
					err = errors.Wrap(err, "error with WatchContentPublished")
					w.errors <- err
					return err
				case <-quit:
					log.Infof("Quit WatchContentPublished (a): addr: %v", w.contractAddress.Hex())
					return nil
				}
			case err := <-sub.Err():
				log.Errorf("Error with WatchContentPublished: addr: %v, fatal (b): %v", w.contractAddress.Hex(), err)
				err = errors.Wrap(err, "error with WatchContentPublished")
				w.errors <- err
				return err
			case <-quit:
				log.Infof("Quitting loop for WatchContentPublished: addr: %v", w.contractAddress.Hex())
				return nil
			}
		}
	}), nil
}

func (w *NewsroomContractWatchers) startWatchNameChanged(eventRecvChan chan *model.Event) (utils.WatcherSubscription, error) {
	killCancelTimeoutSecs := 10
	return utils.NewWatcherSubscription("WatchNameChanged", func(quit <-chan struct{}) error {
		startupFn := func() (utils.WatcherSubscription, chan *contract.NewsroomContractNameChanged, error) {
			ctx := context.Background()
			ctx, cancelFn := context.WithCancel(ctx)
			opts := &bind.WatchOpts{Context: ctx}
			killCancel := make(chan struct{})
			// 10 sec timeout mechanism for starting up watcher
			go func(cancelFn context.CancelFunc, killCancel <-chan struct{}) {
				select {
				case <-time.After(time.Duration(killCancelTimeoutSecs) * time.Second):
					log.Errorf("WatchNameChanged start timeout, cancelling...")
					cancelFn()
				case <-killCancel:
				}
			}(cancelFn, killCancel)
			recvChan := make(chan *contract.NewsroomContractNameChanged)
			sub, err := w.contract.WatchNameChanged(
				opts,
				recvChan,
			)
			close(killCancel)
			if err != nil {
				if sub != nil {
					log.Infof("startupFn: Unsubscribing WatchNameChanged: addr: %v", w.contractAddress.Hex())
					sub.Unsubscribe()
					sub = nil
				}
				return nil, nil, errors.Wrap(err, "startupFn: error starting WatchNameChanged")
			}
			return sub, recvChan, nil
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchNameChanged: addr: %v, %v", w.contractAddress.Hex(), err)
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
		log.Infof("Starting up WatchNameChanged: addr: %v", w.contractAddress.Hex())
		for {
			select {
			case event := <-recvChan:
				if log.V(2) {
					log.Infof("Received event on WatchNameChanged: %v", spew.Sprintf("%#+v", event))
				} else {
					log.Info("Received event on WatchNameChanged")
				}
				modelEvent, err := model.NewEventFromContractEvent("NameChanged", w.ContractName(), w.contractAddress, event, ctime.CurrentEpochSecsInInt64(), model.Watcher)
				if err != nil {
					log.Errorf("Error creating new event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- modelEvent:
					if log.V(2) {
						log.Infof("Sent event to eventRecvChan on WatchNameChanged: %v", spew.Sprintf("%#+v", event))
					} else {
						log.Info("Sent event to eventRecvChan on WatchNameChanged")
					}
				case err := <-sub.Err():
					log.Errorf("Error with WatchNameChanged: addr: %v, fatal (a): %v", w.contractAddress.Hex(), err)
					err = errors.Wrap(err, "error with WatchNameChanged")
					w.errors <- err
					return err
				case <-quit:
					log.Infof("Quit WatchNameChanged (a): addr: %v", w.contractAddress.Hex())
					return nil
				}
			case err := <-sub.Err():
				log.Errorf("Error with WatchNameChanged: addr: %v, fatal (b): %v", w.contractAddress.Hex(), err)
				err = errors.Wrap(err, "error with WatchNameChanged")
				w.errors <- err
				return err
			case <-quit:
				log.Infof("Quitting loop for WatchNameChanged: addr: %v", w.contractAddress.Hex())
				return nil
			}
		}
	}), nil
}

func (w *NewsroomContractWatchers) startWatchOwnershipRenounced(eventRecvChan chan *model.Event) (utils.WatcherSubscription, error) {
	killCancelTimeoutSecs := 10
	return utils.NewWatcherSubscription("WatchOwnershipRenounced", func(quit <-chan struct{}) error {
		startupFn := func() (utils.WatcherSubscription, chan *contract.NewsroomContractOwnershipRenounced, error) {
			ctx := context.Background()
			ctx, cancelFn := context.WithCancel(ctx)
			opts := &bind.WatchOpts{Context: ctx}
			killCancel := make(chan struct{})
			// 10 sec timeout mechanism for starting up watcher
			go func(cancelFn context.CancelFunc, killCancel <-chan struct{}) {
				select {
				case <-time.After(time.Duration(killCancelTimeoutSecs) * time.Second):
					log.Errorf("WatchOwnershipRenounced start timeout, cancelling...")
					cancelFn()
				case <-killCancel:
				}
			}(cancelFn, killCancel)
			recvChan := make(chan *contract.NewsroomContractOwnershipRenounced)
			sub, err := w.contract.WatchOwnershipRenounced(
				opts,
				recvChan,
				[]common.Address{},
			)
			close(killCancel)
			if err != nil {
				if sub != nil {
					log.Infof("startupFn: Unsubscribing WatchOwnershipRenounced: addr: %v", w.contractAddress.Hex())
					sub.Unsubscribe()
					sub = nil
				}
				return nil, nil, errors.Wrap(err, "startupFn: error starting WatchOwnershipRenounced")
			}
			return sub, recvChan, nil
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchOwnershipRenounced: addr: %v, %v", w.contractAddress.Hex(), err)
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
		log.Infof("Starting up WatchOwnershipRenounced: addr: %v", w.contractAddress.Hex())
		for {
			select {
			case event := <-recvChan:
				if log.V(2) {
					log.Infof("Received event on WatchOwnershipRenounced: %v", spew.Sprintf("%#+v", event))
				} else {
					log.Info("Received event on WatchOwnershipRenounced")
				}
				modelEvent, err := model.NewEventFromContractEvent("OwnershipRenounced", w.ContractName(), w.contractAddress, event, ctime.CurrentEpochSecsInInt64(), model.Watcher)
				if err != nil {
					log.Errorf("Error creating new event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- modelEvent:
					if log.V(2) {
						log.Infof("Sent event to eventRecvChan on WatchOwnershipRenounced: %v", spew.Sprintf("%#+v", event))
					} else {
						log.Info("Sent event to eventRecvChan on WatchOwnershipRenounced")
					}
				case err := <-sub.Err():
					log.Errorf("Error with WatchOwnershipRenounced: addr: %v, fatal (a): %v", w.contractAddress.Hex(), err)
					err = errors.Wrap(err, "error with WatchOwnershipRenounced")
					w.errors <- err
					return err
				case <-quit:
					log.Infof("Quit WatchOwnershipRenounced (a): addr: %v", w.contractAddress.Hex())
					return nil
				}
			case err := <-sub.Err():
				log.Errorf("Error with WatchOwnershipRenounced: addr: %v, fatal (b): %v", w.contractAddress.Hex(), err)
				err = errors.Wrap(err, "error with WatchOwnershipRenounced")
				w.errors <- err
				return err
			case <-quit:
				log.Infof("Quitting loop for WatchOwnershipRenounced: addr: %v", w.contractAddress.Hex())
				return nil
			}
		}
	}), nil
}

func (w *NewsroomContractWatchers) startWatchOwnershipTransferred(eventRecvChan chan *model.Event) (utils.WatcherSubscription, error) {
	killCancelTimeoutSecs := 10
	return utils.NewWatcherSubscription("WatchOwnershipTransferred", func(quit <-chan struct{}) error {
		startupFn := func() (utils.WatcherSubscription, chan *contract.NewsroomContractOwnershipTransferred, error) {
			ctx := context.Background()
			ctx, cancelFn := context.WithCancel(ctx)
			opts := &bind.WatchOpts{Context: ctx}
			killCancel := make(chan struct{})
			// 10 sec timeout mechanism for starting up watcher
			go func(cancelFn context.CancelFunc, killCancel <-chan struct{}) {
				select {
				case <-time.After(time.Duration(killCancelTimeoutSecs) * time.Second):
					log.Errorf("WatchOwnershipTransferred start timeout, cancelling...")
					cancelFn()
				case <-killCancel:
				}
			}(cancelFn, killCancel)
			recvChan := make(chan *contract.NewsroomContractOwnershipTransferred)
			sub, err := w.contract.WatchOwnershipTransferred(
				opts,
				recvChan,
				[]common.Address{},
				[]common.Address{},
			)
			close(killCancel)
			if err != nil {
				if sub != nil {
					log.Infof("startupFn: Unsubscribing WatchOwnershipTransferred: addr: %v", w.contractAddress.Hex())
					sub.Unsubscribe()
					sub = nil
				}
				return nil, nil, errors.Wrap(err, "startupFn: error starting WatchOwnershipTransferred")
			}
			return sub, recvChan, nil
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchOwnershipTransferred: addr: %v, %v", w.contractAddress.Hex(), err)
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
		log.Infof("Starting up WatchOwnershipTransferred: addr: %v", w.contractAddress.Hex())
		for {
			select {
			case event := <-recvChan:
				if log.V(2) {
					log.Infof("Received event on WatchOwnershipTransferred: %v", spew.Sprintf("%#+v", event))
				} else {
					log.Info("Received event on WatchOwnershipTransferred")
				}
				modelEvent, err := model.NewEventFromContractEvent("OwnershipTransferred", w.ContractName(), w.contractAddress, event, ctime.CurrentEpochSecsInInt64(), model.Watcher)
				if err != nil {
					log.Errorf("Error creating new event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- modelEvent:
					if log.V(2) {
						log.Infof("Sent event to eventRecvChan on WatchOwnershipTransferred: %v", spew.Sprintf("%#+v", event))
					} else {
						log.Info("Sent event to eventRecvChan on WatchOwnershipTransferred")
					}
				case err := <-sub.Err():
					log.Errorf("Error with WatchOwnershipTransferred: addr: %v, fatal (a): %v", w.contractAddress.Hex(), err)
					err = errors.Wrap(err, "error with WatchOwnershipTransferred")
					w.errors <- err
					return err
				case <-quit:
					log.Infof("Quit WatchOwnershipTransferred (a): addr: %v", w.contractAddress.Hex())
					return nil
				}
			case err := <-sub.Err():
				log.Errorf("Error with WatchOwnershipTransferred: addr: %v, fatal (b): %v", w.contractAddress.Hex(), err)
				err = errors.Wrap(err, "error with WatchOwnershipTransferred")
				w.errors <- err
				return err
			case <-quit:
				log.Infof("Quitting loop for WatchOwnershipTransferred: addr: %v", w.contractAddress.Hex())
				return nil
			}
		}
	}), nil
}

func (w *NewsroomContractWatchers) startWatchRevisionSigned(eventRecvChan chan *model.Event) (utils.WatcherSubscription, error) {
	killCancelTimeoutSecs := 10
	return utils.NewWatcherSubscription("WatchRevisionSigned", func(quit <-chan struct{}) error {
		startupFn := func() (utils.WatcherSubscription, chan *contract.NewsroomContractRevisionSigned, error) {
			ctx := context.Background()
			ctx, cancelFn := context.WithCancel(ctx)
			opts := &bind.WatchOpts{Context: ctx}
			killCancel := make(chan struct{})
			// 10 sec timeout mechanism for starting up watcher
			go func(cancelFn context.CancelFunc, killCancel <-chan struct{}) {
				select {
				case <-time.After(time.Duration(killCancelTimeoutSecs) * time.Second):
					log.Errorf("WatchRevisionSigned start timeout, cancelling...")
					cancelFn()
				case <-killCancel:
				}
			}(cancelFn, killCancel)
			recvChan := make(chan *contract.NewsroomContractRevisionSigned)
			sub, err := w.contract.WatchRevisionSigned(
				opts,
				recvChan,
				[]*big.Int{},
				[]*big.Int{},
				[]common.Address{},
			)
			close(killCancel)
			if err != nil {
				if sub != nil {
					log.Infof("startupFn: Unsubscribing WatchRevisionSigned: addr: %v", w.contractAddress.Hex())
					sub.Unsubscribe()
					sub = nil
				}
				return nil, nil, errors.Wrap(err, "startupFn: error starting WatchRevisionSigned")
			}
			return sub, recvChan, nil
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchRevisionSigned: addr: %v, %v", w.contractAddress.Hex(), err)
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
		log.Infof("Starting up WatchRevisionSigned: addr: %v", w.contractAddress.Hex())
		for {
			select {
			case event := <-recvChan:
				if log.V(2) {
					log.Infof("Received event on WatchRevisionSigned: %v", spew.Sprintf("%#+v", event))
				} else {
					log.Info("Received event on WatchRevisionSigned")
				}
				modelEvent, err := model.NewEventFromContractEvent("RevisionSigned", w.ContractName(), w.contractAddress, event, ctime.CurrentEpochSecsInInt64(), model.Watcher)
				if err != nil {
					log.Errorf("Error creating new event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- modelEvent:
					if log.V(2) {
						log.Infof("Sent event to eventRecvChan on WatchRevisionSigned: %v", spew.Sprintf("%#+v", event))
					} else {
						log.Info("Sent event to eventRecvChan on WatchRevisionSigned")
					}
				case err := <-sub.Err():
					log.Errorf("Error with WatchRevisionSigned: addr: %v, fatal (a): %v", w.contractAddress.Hex(), err)
					err = errors.Wrap(err, "error with WatchRevisionSigned")
					w.errors <- err
					return err
				case <-quit:
					log.Infof("Quit WatchRevisionSigned (a): addr: %v", w.contractAddress.Hex())
					return nil
				}
			case err := <-sub.Err():
				log.Errorf("Error with WatchRevisionSigned: addr: %v, fatal (b): %v", w.contractAddress.Hex(), err)
				err = errors.Wrap(err, "error with WatchRevisionSigned")
				w.errors <- err
				return err
			case <-quit:
				log.Infof("Quitting loop for WatchRevisionSigned: addr: %v", w.contractAddress.Hex())
				return nil
			}
		}
	}), nil
}

func (w *NewsroomContractWatchers) startWatchRevisionUpdated(eventRecvChan chan *model.Event) (utils.WatcherSubscription, error) {
	killCancelTimeoutSecs := 10
	return utils.NewWatcherSubscription("WatchRevisionUpdated", func(quit <-chan struct{}) error {
		startupFn := func() (utils.WatcherSubscription, chan *contract.NewsroomContractRevisionUpdated, error) {
			ctx := context.Background()
			ctx, cancelFn := context.WithCancel(ctx)
			opts := &bind.WatchOpts{Context: ctx}
			killCancel := make(chan struct{})
			// 10 sec timeout mechanism for starting up watcher
			go func(cancelFn context.CancelFunc, killCancel <-chan struct{}) {
				select {
				case <-time.After(time.Duration(killCancelTimeoutSecs) * time.Second):
					log.Errorf("WatchRevisionUpdated start timeout, cancelling...")
					cancelFn()
				case <-killCancel:
				}
			}(cancelFn, killCancel)
			recvChan := make(chan *contract.NewsroomContractRevisionUpdated)
			sub, err := w.contract.WatchRevisionUpdated(
				opts,
				recvChan,
				[]common.Address{},
				[]*big.Int{},
				[]*big.Int{},
			)
			close(killCancel)
			if err != nil {
				if sub != nil {
					log.Infof("startupFn: Unsubscribing WatchRevisionUpdated: addr: %v", w.contractAddress.Hex())
					sub.Unsubscribe()
					sub = nil
				}
				return nil, nil, errors.Wrap(err, "startupFn: error starting WatchRevisionUpdated")
			}
			return sub, recvChan, nil
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchRevisionUpdated: addr: %v, %v", w.contractAddress.Hex(), err)
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
		log.Infof("Starting up WatchRevisionUpdated: addr: %v", w.contractAddress.Hex())
		for {
			select {
			case event := <-recvChan:
				if log.V(2) {
					log.Infof("Received event on WatchRevisionUpdated: %v", spew.Sprintf("%#+v", event))
				} else {
					log.Info("Received event on WatchRevisionUpdated")
				}
				modelEvent, err := model.NewEventFromContractEvent("RevisionUpdated", w.ContractName(), w.contractAddress, event, ctime.CurrentEpochSecsInInt64(), model.Watcher)
				if err != nil {
					log.Errorf("Error creating new event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- modelEvent:
					if log.V(2) {
						log.Infof("Sent event to eventRecvChan on WatchRevisionUpdated: %v", spew.Sprintf("%#+v", event))
					} else {
						log.Info("Sent event to eventRecvChan on WatchRevisionUpdated")
					}
				case err := <-sub.Err():
					log.Errorf("Error with WatchRevisionUpdated: addr: %v, fatal (a): %v", w.contractAddress.Hex(), err)
					err = errors.Wrap(err, "error with WatchRevisionUpdated")
					w.errors <- err
					return err
				case <-quit:
					log.Infof("Quit WatchRevisionUpdated (a): addr: %v", w.contractAddress.Hex())
					return nil
				}
			case err := <-sub.Err():
				log.Errorf("Error with WatchRevisionUpdated: addr: %v, fatal (b): %v", w.contractAddress.Hex(), err)
				err = errors.Wrap(err, "error with WatchRevisionUpdated")
				w.errors <- err
				return err
			case <-quit:
				log.Infof("Quitting loop for WatchRevisionUpdated: addr: %v", w.contractAddress.Hex())
				return nil
			}
		}
	}), nil
}

func (w *NewsroomContractWatchers) startWatchRoleAdded(eventRecvChan chan *model.Event) (utils.WatcherSubscription, error) {
	killCancelTimeoutSecs := 10
	return utils.NewWatcherSubscription("WatchRoleAdded", func(quit <-chan struct{}) error {
		startupFn := func() (utils.WatcherSubscription, chan *contract.NewsroomContractRoleAdded, error) {
			ctx := context.Background()
			ctx, cancelFn := context.WithCancel(ctx)
			opts := &bind.WatchOpts{Context: ctx}
			killCancel := make(chan struct{})
			// 10 sec timeout mechanism for starting up watcher
			go func(cancelFn context.CancelFunc, killCancel <-chan struct{}) {
				select {
				case <-time.After(time.Duration(killCancelTimeoutSecs) * time.Second):
					log.Errorf("WatchRoleAdded start timeout, cancelling...")
					cancelFn()
				case <-killCancel:
				}
			}(cancelFn, killCancel)
			recvChan := make(chan *contract.NewsroomContractRoleAdded)
			sub, err := w.contract.WatchRoleAdded(
				opts,
				recvChan,
				[]common.Address{},
				[]common.Address{},
			)
			close(killCancel)
			if err != nil {
				if sub != nil {
					log.Infof("startupFn: Unsubscribing WatchRoleAdded: addr: %v", w.contractAddress.Hex())
					sub.Unsubscribe()
					sub = nil
				}
				return nil, nil, errors.Wrap(err, "startupFn: error starting WatchRoleAdded")
			}
			return sub, recvChan, nil
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchRoleAdded: addr: %v, %v", w.contractAddress.Hex(), err)
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
		log.Infof("Starting up WatchRoleAdded: addr: %v", w.contractAddress.Hex())
		for {
			select {
			case event := <-recvChan:
				if log.V(2) {
					log.Infof("Received event on WatchRoleAdded: %v", spew.Sprintf("%#+v", event))
				} else {
					log.Info("Received event on WatchRoleAdded")
				}
				modelEvent, err := model.NewEventFromContractEvent("RoleAdded", w.ContractName(), w.contractAddress, event, ctime.CurrentEpochSecsInInt64(), model.Watcher)
				if err != nil {
					log.Errorf("Error creating new event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- modelEvent:
					if log.V(2) {
						log.Infof("Sent event to eventRecvChan on WatchRoleAdded: %v", spew.Sprintf("%#+v", event))
					} else {
						log.Info("Sent event to eventRecvChan on WatchRoleAdded")
					}
				case err := <-sub.Err():
					log.Errorf("Error with WatchRoleAdded: addr: %v, fatal (a): %v", w.contractAddress.Hex(), err)
					err = errors.Wrap(err, "error with WatchRoleAdded")
					w.errors <- err
					return err
				case <-quit:
					log.Infof("Quit WatchRoleAdded (a): addr: %v", w.contractAddress.Hex())
					return nil
				}
			case err := <-sub.Err():
				log.Errorf("Error with WatchRoleAdded: addr: %v, fatal (b): %v", w.contractAddress.Hex(), err)
				err = errors.Wrap(err, "error with WatchRoleAdded")
				w.errors <- err
				return err
			case <-quit:
				log.Infof("Quitting loop for WatchRoleAdded: addr: %v", w.contractAddress.Hex())
				return nil
			}
		}
	}), nil
}

func (w *NewsroomContractWatchers) startWatchRoleRemoved(eventRecvChan chan *model.Event) (utils.WatcherSubscription, error) {
	killCancelTimeoutSecs := 10
	return utils.NewWatcherSubscription("WatchRoleRemoved", func(quit <-chan struct{}) error {
		startupFn := func() (utils.WatcherSubscription, chan *contract.NewsroomContractRoleRemoved, error) {
			ctx := context.Background()
			ctx, cancelFn := context.WithCancel(ctx)
			opts := &bind.WatchOpts{Context: ctx}
			killCancel := make(chan struct{})
			// 10 sec timeout mechanism for starting up watcher
			go func(cancelFn context.CancelFunc, killCancel <-chan struct{}) {
				select {
				case <-time.After(time.Duration(killCancelTimeoutSecs) * time.Second):
					log.Errorf("WatchRoleRemoved start timeout, cancelling...")
					cancelFn()
				case <-killCancel:
				}
			}(cancelFn, killCancel)
			recvChan := make(chan *contract.NewsroomContractRoleRemoved)
			sub, err := w.contract.WatchRoleRemoved(
				opts,
				recvChan,
				[]common.Address{},
				[]common.Address{},
			)
			close(killCancel)
			if err != nil {
				if sub != nil {
					log.Infof("startupFn: Unsubscribing WatchRoleRemoved: addr: %v", w.contractAddress.Hex())
					sub.Unsubscribe()
					sub = nil
				}
				return nil, nil, errors.Wrap(err, "startupFn: error starting WatchRoleRemoved")
			}
			return sub, recvChan, nil
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting WatchRoleRemoved: addr: %v, %v", w.contractAddress.Hex(), err)
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
		log.Infof("Starting up WatchRoleRemoved: addr: %v", w.contractAddress.Hex())
		for {
			select {
			case event := <-recvChan:
				if log.V(2) {
					log.Infof("Received event on WatchRoleRemoved: %v", spew.Sprintf("%#+v", event))
				} else {
					log.Info("Received event on WatchRoleRemoved")
				}
				modelEvent, err := model.NewEventFromContractEvent("RoleRemoved", w.ContractName(), w.contractAddress, event, ctime.CurrentEpochSecsInInt64(), model.Watcher)
				if err != nil {
					log.Errorf("Error creating new event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- modelEvent:
					if log.V(2) {
						log.Infof("Sent event to eventRecvChan on WatchRoleRemoved: %v", spew.Sprintf("%#+v", event))
					} else {
						log.Info("Sent event to eventRecvChan on WatchRoleRemoved")
					}
				case err := <-sub.Err():
					log.Errorf("Error with WatchRoleRemoved: addr: %v, fatal (a): %v", w.contractAddress.Hex(), err)
					err = errors.Wrap(err, "error with WatchRoleRemoved")
					w.errors <- err
					return err
				case <-quit:
					log.Infof("Quit WatchRoleRemoved (a): addr: %v", w.contractAddress.Hex())
					return nil
				}
			case err := <-sub.Err():
				log.Errorf("Error with WatchRoleRemoved: addr: %v, fatal (b): %v", w.contractAddress.Hex(), err)
				err = errors.Wrap(err, "error with WatchRoleRemoved")
				w.errors <- err
				return err
			case <-quit:
				log.Infof("Quitting loop for WatchRoleRemoved: addr: %v", w.contractAddress.Hex())
				return nil
			}
		}
	}), nil
}
