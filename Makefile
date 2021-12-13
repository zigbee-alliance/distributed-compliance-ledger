PACKAGES = $(shell go list ./... | grep -v '/integration_tests')

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=DcLedger \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=dcld \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT)

BUILD_FLAGS := -ldflags '$(ldflags)'
OUTPUT_DIR ?= build

LOCALNET_DIR ?= .localnet

LICENSE_TYPE = "apache"
COPYRIGHT_YEAR = "2020"
COPYRIGHT_HOLDER = "DSR Corporation"
LICENSED_FILES = $(shell find . -type f -not -path '*/.*' -not -name '*.md' -not -name 'requirements.txt')

all: install

build: go.sum
	go build -mod=readonly $(BUILD_FLAGS) -o $(OUTPUT_DIR)/dcld ./cmd/dcld

install: go.sum
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/dcld

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	GO111MODULE=on go mod verify

test:
	go test -v $(PACKAGES)

lint:
	golangci-lint run ./... --timeout 5m0s

license:
	addlicense -l ${LICENSE_TYPE} -y ${COPYRIGHT_YEAR} -c ${COPYRIGHT_HOLDER} ${LICENSED_FILES}

license-check:
	addlicense -l ${LICENSE_TYPE} -y ${COPYRIGHT_YEAR} -c ${COPYRIGHT_HOLDER} -check ${LICENSED_FILES}

clean:
	rm -rf $(OUTPUT_DIR)

# Docker

image:
	docker build -t dcledger .

localnet_init:
	/bin/bash ./genlocalnetconfig.sh

localnet_start:
	@if [ -d "${LOCALNET_DIR}/observer0" ]; then\
		docker-compose --profile observers up -d;\
	else\
		docker-compose up -d;\
	fi

localnet_stop:
	docker-compose down

localnet_export: localnet_stop
	docker-compose run node0 dcld export --for-zero-height  >genesis.export.node0.json
	docker-compose run node1 dcld export --for-zero-height  >genesis.export.node1.json
	docker-compose run node2 dcld export --for-zero-height  >genesis.export.node2.json
	docker-compose run node3 dcld export --for-zero-height  >genesis.export.node3.json
	@if [ -d "${LOCALNET_DIR}/observer0" ]; then\
		docker-compose run observer0 dcld export --for-zero-height  >genesis.export.observer0.json;\
	fi


localnet_reset: localnet_stop
	docker-compose run node0 dcld unsafe-reset-all
	docker-compose run node1 dcld unsafe-reset-all
	docker-compose run node2 dcld unsafe-reset-all
	docker-compose run node3 dcld unsafe-reset-all
	@if [ -d "${LOCALNET_DIR}/observer0" ]; then\
		docker-compose run observer0 dcld unsafe-reset-all;\
	fi

localnet_clean: localnet_stop
	rm -rf $(LOCALNET_DIR)

.PHONY: all build install test lint clean image localnet_init localnet_start localnet_stop localnet_clean localnet_export localnet_reset license license-check
