PACKAGES = $(shell go list ./... | grep -v '/integration_tests')

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=DcLedger \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=dcld \
	-X github.com/cosmos/cosmos-sdk/version.ClientName=dclcli \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) 

BUILD_FLAGS := -ldflags '$(ldflags)'
OUTPUT_DIR ?= build

LOCALNET_DIR ?= localnet

LICENSE_TYPE = "apache"
COPYRIGHT_YEAR = "2020"
COPYRIGHT_HOLDER = "DSR Corporation"
LICENSED_FILES = $(shell find . -type f -not -wholename '*/.*')

all: install

build: go.sum
	go build -mod=readonly $(BUILD_FLAGS) -o $(OUTPUT_DIR)/dcld ./cmd/dcld
	go build -mod=readonly $(BUILD_FLAGS) -o $(OUTPUT_DIR)/dclcli ./cmd/dclcli

install: go.sum
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/dcld
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/dclcli

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
	docker-compose up -d

localnet_stop:
	docker-compose down

localnet_clean: localnet_stop
	rm -rf $(LOCALNET_DIR)

.PHONY: all build install test lint clean image localnet_init localnet_start localnet_stop localnet_clean license license-check