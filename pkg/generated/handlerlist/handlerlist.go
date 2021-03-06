// Code generated by 'gen/handlerlistgen'  DO NOT EDIT.
// IT SHOULD NOT BE EDITED BY HAND AS ANY CHANGES MAY BE OVERWRITTEN
// File was generated at 2019-12-19 21:32:36.705867 +0000 UTC
package handlerlist

import (
	log "github.com/golang/glog"

	"github.com/ethereum/go-ethereum/common"

	"github.com/joincivil/civil-events-crawler/pkg/generated/filterer"
	"github.com/joincivil/civil-events-crawler/pkg/generated/watcher"
	"github.com/joincivil/civil-events-crawler/pkg/model"
)

func ContractFilterers(nameToAddrs map[string][]common.Address) []model.ContractFilterers {
	filters := []model.ContractFilterers{}

	var addrs []common.Address
	var addr common.Address
	var ok bool

	addrs, ok = nameToAddrs["newsroom"]
	if ok {
		for _, addr = range addrs {
			filter := filterer.NewNewsroomContractFilterers(addr)
			filters = append(filters, filter)
			log.Info("Added NewsroomContract filterer")
		}
	}

	addrs, ok = nameToAddrs["newsroomfactory"]
	if ok {
		for _, addr = range addrs {
			filter := filterer.NewNewsroomFactoryFilterers(addr)
			filters = append(filters, filter)
			log.Info("Added NewsroomFactory filterer")
		}
	}

	addrs, ok = nameToAddrs["civiltcr"]
	if ok {
		for _, addr = range addrs {
			filter := filterer.NewCivilTCRContractFilterers(addr)
			filters = append(filters, filter)
			log.Info("Added CivilTCRContract filterer")
		}
	}

	addrs, ok = nameToAddrs["cvltoken"]
	if ok {
		for _, addr = range addrs {
			filter := filterer.NewCVLTokenContractFilterers(addr)
			filters = append(filters, filter)
			log.Info("Added CVLTokenContract filterer")
		}
	}

	addrs, ok = nameToAddrs["civilparameterizer"]
	if ok {
		for _, addr = range addrs {
			filter := filterer.NewParameterizerContractFilterers(addr)
			filters = append(filters, filter)
			log.Info("Added ParameterizerContract filterer")
		}
	}

	addrs, ok = nameToAddrs["civilgovernment"]
	if ok {
		for _, addr = range addrs {
			filter := filterer.NewGovernmentContractFilterers(addr)
			filters = append(filters, filter)
			log.Info("Added GovernmentContract filterer")
		}
	}

	addrs, ok = nameToAddrs["multisigwallet"]
	if ok {
		for _, addr = range addrs {
			filter := filterer.NewMultiSigWalletContractFilterers(addr)
			filters = append(filters, filter)
			log.Info("Added MultiSigWalletContract filterer")
		}
	}

	addrs, ok = nameToAddrs["multisigwalletfactory"]
	if ok {
		for _, addr = range addrs {
			filter := filterer.NewMultiSigWalletFactoryContractFilterers(addr)
			filters = append(filters, filter)
			log.Info("Added MultiSigWalletFactoryContract filterer")
		}
	}

	addrs, ok = nameToAddrs["civilplcrvoting"]
	if ok {
		for _, addr = range addrs {
			filter := filterer.NewCivilPLCRVotingContractFilterers(addr)
			filters = append(filters, filter)
			log.Info("Added CivilPLCRVotingContract filterer")
		}
	}

	return filters
}

func ContractWatchers(nameToAddrs map[string][]common.Address) []model.ContractWatchers {
	watchers := []model.ContractWatchers{}

	var addrs []common.Address
	var addr common.Address
	var ok bool

	addrs, ok = nameToAddrs["newsroom"]
	if ok {
		for _, addr = range addrs {
			watch := watcher.NewNewsroomContractWatchers(addr)
			watchers = append(watchers, watch)
			log.Info("Added NewsroomContract watcher")
		}
	}

	addrs, ok = nameToAddrs["newsroomfactory"]
	if ok {
		for _, addr = range addrs {
			watch := watcher.NewNewsroomFactoryWatchers(addr)
			watchers = append(watchers, watch)
			log.Info("Added NewsroomFactory watcher")
		}
	}

	addrs, ok = nameToAddrs["civiltcr"]
	if ok {
		for _, addr = range addrs {
			watch := watcher.NewCivilTCRContractWatchers(addr)
			watchers = append(watchers, watch)
			log.Info("Added CivilTCRContract watcher")
		}
	}

	addrs, ok = nameToAddrs["cvltoken"]
	if ok {
		for _, addr = range addrs {
			watch := watcher.NewCVLTokenContractWatchers(addr)
			watchers = append(watchers, watch)
			log.Info("Added CVLTokenContract watcher")
		}
	}

	addrs, ok = nameToAddrs["civilparameterizer"]
	if ok {
		for _, addr = range addrs {
			watch := watcher.NewParameterizerContractWatchers(addr)
			watchers = append(watchers, watch)
			log.Info("Added ParameterizerContract watcher")
		}
	}

	addrs, ok = nameToAddrs["civilgovernment"]
	if ok {
		for _, addr = range addrs {
			watch := watcher.NewGovernmentContractWatchers(addr)
			watchers = append(watchers, watch)
			log.Info("Added GovernmentContract watcher")
		}
	}

	addrs, ok = nameToAddrs["multisigwallet"]
	if ok {
		for _, addr = range addrs {
			watch := watcher.NewMultiSigWalletContractWatchers(addr)
			watchers = append(watchers, watch)
			log.Info("Added MultiSigWalletContract watcher")
		}
	}

	addrs, ok = nameToAddrs["multisigwalletfactory"]
	if ok {
		for _, addr = range addrs {
			watch := watcher.NewMultiSigWalletFactoryContractWatchers(addr)
			watchers = append(watchers, watch)
			log.Info("Added MultiSigWalletFactoryContract watcher")
		}
	}

	addrs, ok = nameToAddrs["civilplcrvoting"]
	if ok {
		for _, addr = range addrs {
			watch := watcher.NewCivilPLCRVotingContractWatchers(addr)
			watchers = append(watchers, watch)
			log.Info("Added CivilPLCRVotingContract watcher")
		}
	}

	return watchers
}
