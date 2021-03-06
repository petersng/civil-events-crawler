POSTGRES_DATA_DIR=postgresdata
POSTGRES_DOCKER_IMAGE=circleci/postgres:9.6-alpine
POSTGRES_PORT=5432
POSTGRES_DB_NAME=civil_crawler
POSTGRES_USER=docker
POSTGRES_PSWD=docker

PUBSUB_SIM_DOCKER_IMAGE=kinok/google-pubsub-emulator:latest

GOVERSION=go1.12.7

GOCMD=go
GOGEN=$(GOCMD) generate
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOCOVER=$(GOCMD) tool cover

## Check to see if these commands are installed
GO:=$(shell command -v go 2> /dev/null)
DOCKER:=$(shell command -v docker 2> /dev/null)
APT:=$(shell command -v apt-get 2> /dev/null)
GOVERCURRENT=$(shell go version |awk {'print $$3'})

## List of expected dirs for generated code
GENERATED_DIR=pkg/generated
GENERATED_WATCHER_DIR=$(GENERATED_DIR)/watcher
GENERATED_FILTERER_DIR=$(GENERATED_DIR)/filterer
GENERATED_COMMON_DIR=$(GENERATED_DIR)/common
GENERATED_HANDLER_LIST_DIR=$(GENERATED_DIR)/handlerlist

## Civil specific commands
EVENTHANDLER_GEN_MAIN=cmd/eventhandlergen/main.go
HANDLERLIST_GEN_MAIN=cmd/handlerlistgen/main.go

# curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin vX.Y.Z
GOLANGCILINT_URL=https://install.goreleaser.com/github.com/golangci/golangci-lint.sh
GOLANGCILINT_VERSION_TAG=v1.16.0

## Reliant on go and $GOPATH being set.
.PHONY: check-go-env
check-go-env:
ifndef GO
	$(error go command is not installed or in PATH)
endif
ifndef GOPATH
	$(error GOPATH is not set)
endif
ifneq ($(GOVERCURRENT), $(GOVERSION))
	$(error Incorrect go version, needs $(GOVERSION))
endif

## NOTE: If installing on a Mac, use Docker for Mac, not Docker toolkit
## https://www.docker.com/docker-mac
.PHONY: check-docker-env
check-docker-env:
ifndef DOCKER
	$(error docker command is not installed or in PATH)
endif

.PHONY: install-linter
install-linter: check-go-env ## Installs linter
	@curl -sfL $(GOLANGCILINT_URL) | sh -s -- -b $(shell go env GOPATH)/bin $(GOLANGCILINT_VERSION_TAG)
ifdef APT
	@sudo apt-get install golang-race-detector-runtime || true
endif

.PHONY: install-cover
install-cover: check-go-env ## Installs code coverage tool
	@$(GOGET) -u golang.org/x/tools/cmd/cover

.PHONY: setup
setup: check-go-env install-linter install-cover ## Sets up the tooling.

.PHONY: postgres-setup-launch
postgres-setup-launch:
ifeq ("$(wildcard $(POSTGRES_DATA_DIR))", "")
	mkdir -p $(POSTGRES_DATA_DIR)
	docker run \
		-v $$PWD/$(POSTGRES_DATA_DIR):/tmp/$(POSTGRES_DATA_DIR) -i -t $(POSTGRES_DOCKER_IMAGE) \
		/bin/bash -c "cp -rp /var/lib/postgresql /tmp/$(POSTGRES_DATA_DIR)"
endif
	docker run -e "POSTGRES_USER="$(POSTGRES_USER) -e "POSTGRES_PASSWORD"=$(POSTGRES_PSWD) -e "POSTGRES_DB"=$(POSTGRES_DB_NAME) \
	    -v $$PWD/$(POSTGRES_DATA_DIR)/postgresql:/var/lib/postgresql -d -p $(POSTGRES_PORT):$(POSTGRES_PORT) \
		$(POSTGRES_DOCKER_IMAGE);

.PHONY: postgres-check-available
postgres-check-available:
	@for i in `seq 1 10`; \
	do \
		nc -z localhost 5432 2> /dev/null && exit 0; \
		sleep 3; \
	done; \
	exit 1;

.PHONY: postgres-start
postgres-start: check-docker-env postgres-setup-launch postgres-check-available ## Starts up a development PostgreSQL server
	@echo "Postgresql launched and available"

.PHONY: postgres-stop
postgres-stop: check-docker-env ## Stops the development PostgreSQL server
	@docker stop `docker ps -q --filter "ancestor=$(POSTGRES_DOCKER_IMAGE)"`
	@echo 'Postgres stopped'

.PHONY: pubsub-setup-launch
pubsub-setup-launch:
	@docker run -it -d -p 8042:8042 $(PUBSUB_SIM_DOCKER_IMAGE)

.PHONY: pubsub-start
pubsub-start: check-docker-env pubsub-setup-launch ## Starts up the pubsub simulator
	@echo 'Google pubsub simulator up'

.PHONY: pubsub-stop
pubsub-stop: check-docker-env ## Stops the pubsub simulator
	@docker stop `docker ps -q --filter "ancestor=$(PUBSUB_SIM_DOCKER_IMAGE)"`
	@echo 'Google pubsub simulator down'

## golangci-lint config in .golangci.yml
.PHONY: lint
lint: check-go-env ## Runs linting.
	@golangci-lint run ./...

.PHONY: generate-civil
generate-civil: generate-civil-watchers generate-civil-filterers generate-civil-common generate-civil-handler-lists ## Runs all the civil code generation

.PHONY: generate-civil-watchers
generate-civil-watchers: check-go-env ## Runs watchergen to generate contract Watch* wrapper code for Civil.
	@mkdir -p $(GENERATED_WATCHER_DIR)
	@$(GORUN) $(EVENTHANDLER_GEN_MAIN) civiltcr watcher watcher > ./$(GENERATED_WATCHER_DIR)/civiltcr.go
	@$(GORUN) $(EVENTHANDLER_GEN_MAIN) newsroom watcher watcher > ./$(GENERATED_WATCHER_DIR)/newsroom.go
	@$(GORUN) $(EVENTHANDLER_GEN_MAIN) civilplcrvoting watcher watcher > ./$(GENERATED_WATCHER_DIR)/civilplcrvoting.go
	@$(GORUN) $(EVENTHANDLER_GEN_MAIN) cvltoken watcher watcher > ./$(GENERATED_WATCHER_DIR)/cvltoken.go
	@$(GORUN) $(EVENTHANDLER_GEN_MAIN) civilparameterizer watcher watcher > ./$(GENERATED_WATCHER_DIR)/civilparameterizer.go
	@$(GORUN) $(EVENTHANDLER_GEN_MAIN) civilgovernment watcher watcher > ./$(GENERATED_WATCHER_DIR)/civilgovernment.go
	@$(GORUN) $(EVENTHANDLER_GEN_MAIN) multisigwallet watcher watcher > ./$(GENERATED_WATCHER_DIR)/multisigwallet.go
	@$(GORUN) $(EVENTHANDLER_GEN_MAIN) newsroomfactory watcher watcher > ./$(GENERATED_WATCHER_DIR)/newsroomfactory.go
	@$(GORUN) $(EVENTHANDLER_GEN_MAIN) multisigwalletfactory watcher watcher > ./$(GENERATED_WATCHER_DIR)/multisigwalletfactory.go

.PHONY: generate-civil-filterers
generate-civil-filterers: check-go-env ## Runs filterergen to generate contract Filter* wrapper code for Civil.
	@mkdir -p $(GENERATED_FILTERER_DIR)
	@$(GORUN) $(EVENTHANDLER_GEN_MAIN) civiltcr filterer filterer > ./$(GENERATED_FILTERER_DIR)/civiltcr.go
	@$(GORUN) $(EVENTHANDLER_GEN_MAIN) newsroom filterer filterer > ./$(GENERATED_FILTERER_DIR)/newsroom.go
	@$(GORUN) $(EVENTHANDLER_GEN_MAIN) civilplcrvoting filterer filterer > ./$(GENERATED_FILTERER_DIR)/civilplcrvoting.go
	@$(GORUN) $(EVENTHANDLER_GEN_MAIN) cvltoken filterer filterer > ./$(GENERATED_FILTERER_DIR)/cvltoken.go
	@$(GORUN) $(EVENTHANDLER_GEN_MAIN) civilparameterizer filterer filterer > ./$(GENERATED_FILTERER_DIR)/civilparameterizer.go
	@$(GORUN) $(EVENTHANDLER_GEN_MAIN) civilgovernment filterer filterer > ./$(GENERATED_FILTERER_DIR)/civilgovernment.go
	@$(GORUN) $(EVENTHANDLER_GEN_MAIN) multisigwallet filterer filterer > ./$(GENERATED_FILTERER_DIR)/multisigwallet.go
	@$(GORUN) $(EVENTHANDLER_GEN_MAIN) newsroomfactory filterer filterer > ./$(GENERATED_FILTERER_DIR)/newsroomfactory.go
	@$(GORUN) $(EVENTHANDLER_GEN_MAIN) multisigwalletfactory filterer filterer > ./$(GENERATED_FILTERER_DIR)/multisigwalletfactory.go

.PHONY: generate-civil-common
generate-civil-common: check-go-env ## Runs commongen to generate common contract wrapper code for Civil.
	@mkdir -p $(GENERATED_COMMON_DIR)
	@$(GORUN) $(EVENTHANDLER_GEN_MAIN) civiltcr common common > ./$(GENERATED_COMMON_DIR)/civiltcr.go
	@$(GORUN) $(EVENTHANDLER_GEN_MAIN) newsroom common common > ./$(GENERATED_COMMON_DIR)/newsroom.go
	@$(GORUN) $(EVENTHANDLER_GEN_MAIN) civilplcrvoting common common > ./$(GENERATED_COMMON_DIR)/civilplcrvoting.go
	@$(GORUN) $(EVENTHANDLER_GEN_MAIN) cvltoken common common > ./$(GENERATED_COMMON_DIR)/cvltoken.go
	@$(GORUN) $(EVENTHANDLER_GEN_MAIN) civilparameterizer common common > ./$(GENERATED_COMMON_DIR)/civilparameterizer.go
	@$(GORUN) $(EVENTHANDLER_GEN_MAIN) civilgovernment common common > ./$(GENERATED_COMMON_DIR)/civilgovernment.go
	@$(GORUN) $(EVENTHANDLER_GEN_MAIN) multisigwallet common common > ./$(GENERATED_COMMON_DIR)/multisigwallet.go
	@$(GORUN) $(EVENTHANDLER_GEN_MAIN) newsroomfactory common common > ./$(GENERATED_COMMON_DIR)/newsroomfactory.go
	@$(GORUN) $(EVENTHANDLER_GEN_MAIN) multisigwalletfactory common common > ./$(GENERATED_COMMON_DIR)/multisigwalletfactory.go

.PHONY: generate-civil-handler-lists
generate-civil-handler-lists: check-go-env ## Runs handlerlistgen to generate handler list wrapper code for Civil.
	@mkdir -p $(GENERATED_HANDLER_LIST_DIR)
	@$(GORUN) $(HANDLERLIST_GEN_MAIN) handlerlist > ./$(GENERATED_HANDLER_LIST_DIR)/handlerlist.go

.PHONY: build
build: check-go-env ## Builds the code.
	$(GOBUILD) -o ./build/crawler cmd/crawler/main.go

.PHONY: test
test: check-go-env ## Runs unit tests and tests code coverage.
	@echo 'mode: atomic' > coverage.txt && $(GOTEST) -covermode=atomic -coverprofile=coverage.txt -v -race -timeout=5m ./...

.PHONY: test-integration
test-integration: check-go-env ## Runs tagged integration tests
	@echo 'mode: atomic' > coverage.txt && PUBSUB_EMULATOR_HOST=localhost:8042 $(GOTEST) -covermode=atomic -coverprofile=coverage.txt -v -race -timeout=5m -tags=integration ./...

.PHONY: test-integration-ci
test-integration-ci: check-go-env ## Runs tagged integration tests serially for low mem/low cpu CI env (set -p to 1)
	@echo 'mode: atomic' > coverage.txt && PUBSUB_EMULATOR_HOST=localhost:8042 $(GOTEST) -covermode=atomic -coverprofile=coverage.txt -v -p 1 -race -timeout=5m -tags=integration ./...

.PHONY: cover
cover: test ## Runs unit tests, code coverage, and runs HTML coverage tool.
	@$(GOCOVER) -html=coverage.txt

.PHONY: cover-integration
cover-integration: test-integration ## Runs unit tests, code coverage, and runs HTML coverage tool for integration
	@$(GOCOVER) -html=coverage.txt

.PHONY: clean
clean: ## go clean and clean up of artifacts.
	@$(GOCLEAN) ./... || true
	@rm coverage.txt || true

## Some magic from http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
