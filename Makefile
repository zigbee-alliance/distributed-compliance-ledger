PACKAGES = $(shell go list ./... | grep -v '/integration_tests')

ifndef DCL_VERSION
DCL_VERSION := $(shell echo $(shell git describe --tags --always) | sed 's/^v//')
endif

ifndef DCL_COMMIT
DCL_COMMIT := $(shell git log -1 --format='%H')
endif

NAME ?= dcl
APPNAME ?= $(NAME)d
LEDGER_ENABLED ?= true

OUTPUT_DIR ?= build

### Process ld flags

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=DcLedger \
	-X github.com/cosmos/cosmos-sdk/version.AppName=$(APPNAME) \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(DCL_VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(DCL_COMMIT)

# DB backend selection
ifeq (cleveldb,$(findstring cleveldb,$(COSMOS_BUILD_OPTIONS)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif


### Process build tags
build_tags = netgo cgo
ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

ifeq (cleveldb,$(findstring cleveldb,$(COSMOS_BUILD_OPTIONS)))
  build_tags += gcc
endif

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))


ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))


### Resulting build flags
BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'

# Check for nostrip option
ifeq (,$(findstring nostrip,$(COSMOS_BUILD_OPTIONS)))
  BUILD_FLAGS += -trimpath
endif

# Check for debug option
ifeq (debug,$(findstring debug,$(COSMOS_BUILD_OPTIONS)))
  BUILD_FLAGS += -gcflags "all=-N -l"
endif


LICENSE_TYPE = "apache"
COPYRIGHT_YEAR = "2020"
COPYRIGHT_HOLDER = "DSR Corporation"
LICENSED_FILES = $(shell find . -type f -not -path '*/.*' -not -name '*.md' -not -name 'requirements.txt')

MK_TEST = "Makefile.test"
LOCALNET_TARGETS = image localnet_init localnet_init_latest_stable_release localnet_start localnet_stop localnet_clean localnet_export localnet_reset localnet_rebuild
TEST_DEPLOY_TARGETS = test_deploy_env_build test_deploy_env_clean
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
	addlicense -l ${LICENSE_TYPE} -y ${COPYRIGHT_YEAR} -c ${COPYRIGHT_HOLDER} -check ${LICENSED_FILES}

clean:
	rm -rf $(OUTPUT_DIR)

${TEST_TARGETS}:
	make -f ${MK_TEST} $@

.PHONY: all build install test lint clean \
		license license-check \
		${TEST_TARGETS}
