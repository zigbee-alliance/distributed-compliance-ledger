PACKAGES=$(shell go list ./... | grep -v '/simulation')

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=ZbLedger \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=zbld \
	-X github.com/cosmos/cosmos-sdk/version.ClientName=zblcli \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) 

BUILD_FLAGS := -ldflags '$(ldflags)'
OUTPUT_DIR?=build
LOCALNET_DIR?=localnet

all: install

build: go.sum
	go build -mod=readonly $(BUILD_FLAGS) -o $(OUTPUT_DIR)/zbld ./cmd/zbld
	go build -mod=readonly $(BUILD_FLAGS) -o $(OUTPUT_DIR)/zblcli ./cmd/zblcli

install: go.sum
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/zbld
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/zblcli

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	GO111MODULE=on go mod verify

test:
	go test -v $(PACKAGES)

lint:
	golangci-lint run ./... --timeout 5m0s

clean:
	rm -rf $(OUTPUT_DIR)

# Docker

image:
	docker build -t zbledger .

localnet_init:
	/bin/bash ./genlocalnetconfig.sh

localnet_start:
	docker-compose up -d

localnet_stop:
	docker-compose down

localnet_clean: localnet_stop
	rm -rf $(LOCALNET_DIR)

.PHONY: all build install test lint clean image localnet_init localnet_start localnet_stop localnet_clean