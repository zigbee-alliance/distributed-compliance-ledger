PACKAGES = $(shell go list ./... | grep -v '/integration_tests')

ifndef DCL_VERSION
DCL_VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
endif

ifndef DCL_COMMIT
DCL_COMMIT := $(shell git log -1 --format='%H')
endif

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=DcLedger \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=dcld \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(DCL_VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(DCL_COMMIT)

BUILD_FLAGS := -ldflags '$(ldflags)'
OUTPUT_DIR ?= build

LICENSE_TYPE = "apache"
COPYRIGHT_YEAR = "2022"
COPYRIGHT_HOLDER = "DSR Corporation"
# LICENSED_FILES = $(shell find . -type f -not -path '*/.*' -not -name '*.md' -not -name 'requirements.txt')
LICENSED_FILES = $(shell find . -type f -name '*.go' -not -name '*.pb.*')

MK_TEST = "Makefile.test"
LOCALNET_TARGETS = image localnet_init localnet_start localnet_stop localnet_clean localnet_export localnet_reset localnet_rebuild
TEST_DEPLOY_TARGETS = test_deploy_image test_deploy_env_build test_deploy_env_clean
TEST_TARGETS= ${LOCALNET_TARGETS} ${TEST_DEPLOY_TARGETS}

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
	addlicense -l ${LICENSE_TYPE} -check ${LICENSED_FILES}

clean:
	rm -rf $(OUTPUT_DIR)

${TEST_TARGETS}:
	make -f ${MK_TEST} $@

.PHONY: all build install test lint clean \
		license license-check \
		${TEST_TARGETS}
