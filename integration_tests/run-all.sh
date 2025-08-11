#!/bin/bash
# Copyright 2020 DSR Corporation
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Possible values: all (default) | cli | light | rest | upgrade | deploy | cli,light | cli,rest | light, rest | cli,light,rest | etc.
TESTS_TO_RUN=${1:-all}

SCRIPT_PATH="$(readlink -f "$0")"
BASEDIR="$(dirname "$SCRIPT_PATH")"
GOCOVER_ENABLED=false
DCL_TMP_GOCOVERDIR="/tmp/dcl_gocover"
DCL_TMP_VALIDATOR_DEMO_GOCOVERDIR="/tmp/dcl_validator_demo_gocover"

if [ "${2:-}" = "cover" ]; then
  export GOCOVER=1
  export GOCOVERDIR="$BASEDIR/gocover"
  GOCOVER_ENABLED=true
  rm -rf "$GOCOVERDIR"
  mkdir -p "$GOCOVERDIR"
fi

DETAILED_OUTPUT=true

LOCALNET_DIR=".localnet"

LOG_PREFIX="[run all] "
SED_EXT=
if [ "$(uname)" == "Darwin" ]; then
  # Mac OS X sed needs the file extension when -i flag is used. Keeping it empty as we don't need backupfile
  SED_EXT="''"
fi

if ${DETAILED_OUTPUT}; then
    DETAILED_OUTPUT_TARGET=/dev/stdout
else
    DETAILED_OUTPUT_TARGET=/dev/null
fi

cleanup() {
  make localnet_clean
}
trap cleanup EXIT

source integration_tests/cli/common.sh

log() {
  echo "${LOG_PREFIX}$1"
}

  # patch configs properly by having all values >= 1 sec, otherwise headers may start having time from the future and light client verification will fail
  # if we patch config to have new blocks created in less than 1 sec, the min time in a time header is still 1 sec.
  # So, new blocks started to be from the future.
patch_consensus_config() {
  local NODE_CONFIGS="$(find "$LOCALNET_DIR" -type f -name "config.toml" -wholename "*node*")"

  for NODE_CONFIG in ${NODE_CONFIGS}; do
    sed -i $SED_EXT 's/timeout_propose = "3s"/timeout_propose = "1s"/g' "${NODE_CONFIG}"
    #sed -i $SED_EXT 's/timeout_prevote = "1s"/timeout_prevote = "1s"/g' "${NODE_CONFIG}"
    #sed -i $SED_EXT 's/timeout_precommit = "1s"/timeout_precommit = "1s"/g' "${NODE_CONFIG}"
    sed -i $SED_EXT 's/timeout_commit = "5s"/timeout_commit = "1s"/g' "${NODE_CONFIG}"
  done
}

init_pool() {
  local _patch_config="${1:-yes}";
  local _localnet_init_target=${2:-localnet_init}
  local _binary_version=${3:-""}

  log "Setting up pool"

  if [ -n "$_binary_version" ]; then
    log "-> Generating network configuration with binary version=$_binary_version" >${DETAILED_OUTPUT_TARGET}
    make ${_localnet_init_target} MAINNET_STABLE_VERSION=$_binary_version &>${DETAILED_OUTPUT_TARGET}
  else
    log "-> Generating network configuration" >${DETAILED_OUTPUT_TARGET}
    make ${_localnet_init_target} &>${DETAILED_OUTPUT_TARGET}
  fi

  if [ "$_patch_config" = "yes" ];
  then
    patch_consensus_config
  fi;

  log "-> Running pool" >${DETAILED_OUTPUT_TARGET}
  make localnet_start &>${DETAILED_OUTPUT_TARGET}

  log "-> Waiting for the second block (needed to request proofs)" >${DETAILED_OUTPUT_TARGET}
  execute_with_retry "dcld status" "connection"
  wait_for_height 2 20
}

cleanup_pool() {
  log "Cleaning up pool"
  log "-> Stopping pool & Removing configurations" >${DETAILED_OUTPUT_TARGET}
  make localnet_clean
}

stop_rest_server() {
  log "Stopping cli in rest-server mode"
  killall dcld
}

collect_cover() {
  if "$GOCOVER_ENABLED"; then
    rm -rf "$DCL_TMP_GOCOVERDIR"
    mkdir -p "$DCL_TMP_GOCOVERDIR"
    local cover_dirs="$GOCOVERDIR"
    local nodes=
		for node in node0 node1 node2 node3 observer0; do
			if [ -d "$LOCALNET_DIR/$node/gocover" ]; then
        docker exec "$node" kill 1
        docker wait "$node" >/dev/null
        cover_dirs+=",$LOCALNET_DIR/$node/gocover"
        nodes+=" $node"
			fi
		done

    if [ -n "$nodes" ]; then
      log "Collecting coverage from$nodes"
      go tool covdata merge -i="$cover_dirs" -o="$DCL_TMP_GOCOVERDIR"
      rm -rf "$GOCOVERDIR"/*
      cp "$DCL_TMP_GOCOVERDIR"/* "$GOCOVERDIR"/
    fi
  fi
}

collect_validator_demo_cover() {
  if "$GOCOVER_ENABLED"; then
    if [ -d "$DCL_TMP_VALIDATOR_DEMO_GOCOVERDIR" ]; then
      log "Collecting coverage from validator-demo"
      rm -rf "$DCL_TMP_GOCOVERDIR"
      mkdir -p "$DCL_TMP_GOCOVERDIR"
      go tool covdata merge -i="$DCL_TMP_VALIDATOR_DEMO_GOCOVERDIR,$GOCOVERDIR" -o="$DCL_TMP_GOCOVERDIR"
      rm -rf "$GOCOVERDIR"/*
      cp "$DCL_TMP_GOCOVERDIR"/* "$GOCOVERDIR"/
      rm -rf "$DCL_TMP_VALIDATOR_DEMO_GOCOVERDIR"
    fi
  fi
}

# Global init
set -euo pipefail

log "Verifying the environment"
check_env

log "Compiling local binaries"
make install &>${DETAILED_OUTPUT_TARGET}

log "Building docker image"
make image &>${DETAILED_OUTPUT_TARGET}

cleanup_pool

# Upgrade procedure tests
if [[ $TESTS_TO_RUN =~ "all" || $TESTS_TO_RUN =~ "upgrade" ]]; then
  UPGRADE_SHELL_TEST="./integration_tests/upgrade/test-upgrade.sh"

  init_pool yes localnet_init_latest_stable_release "v0.12.0"

  log "*****************************************************************************************"
  log "Running $UPGRADE_SHELL_TEST"
  log "*****************************************************************************************"

  if bash "$UPGRADE_SHELL_TEST" &>${DETAILED_OUTPUT_TARGET}; then
    rm dcld_mainnet_stable
    log "$UPGRADE_SHELL_TEST finished successfully"
  else
    log "$UPGRADE_SHELL_TEST failed"
    exit 1
  fi

  collect_cover
  cleanup_pool
fi

# Deploy tests
if [[ $TESTS_TO_RUN =~ "all" || $TESTS_TO_RUN =~ "deploy" ]]; then
  DEPLOY_SHELL_TEST="./integration_tests/deploy/test_deploy.sh"
  if bash "$DEPLOY_SHELL_TEST" &>${DETAILED_OUTPUT_TARGET}; then
    log "$DEPLOY_SHELL_TEST finished successfully"
  else
    log "$DEPLOY_SHELL_TEST failed"
    exit 1
  fi
fi

# Cli shell tests
if [[ $TESTS_TO_RUN =~ "all" || $TESTS_TO_RUN =~ "cli" ]]; then
  CLI_SHELL_TESTS=$(find integration_tests/cli -type f -name '*.sh' -not -name "common.sh")

  for CLI_SHELL_TEST in ${CLI_SHELL_TESTS}; do
    init_pool

    log "*****************************************************************************************"
    log "Running $CLI_SHELL_TEST"
    log "*****************************************************************************************"

    if bash "$CLI_SHELL_TEST" &>${DETAILED_OUTPUT_TARGET}; then
      log "$CLI_SHELL_TEST finished successfully"
    else
      log "$CLI_SHELL_TEST failed"
      exit 1
    fi

    if [[ "$CLI_SHELL_TEST" == *"validator-demo.sh"* ]]; then
      collect_validator_demo_cover
    else
      collect_cover
    fi

    cleanup_pool
  done
fi

# Light Client Proxy Cli shell tests
if [[ $TESTS_TO_RUN =~ "all" || $TESTS_TO_RUN =~ "light" ]]; then
  CLI_SHELL_TESTS=$(find integration_tests/light_client_proxy -type f -name '*.sh' -not -name "common.sh")

  for CLI_SHELL_TEST in ${CLI_SHELL_TESTS}; do
    init_pool

    log "*****************************************************************************************"
    log "Running $CLI_SHELL_TEST"
    log "*****************************************************************************************"

    if bash "$CLI_SHELL_TEST" &>${DETAILED_OUTPUT_TARGET}; then
      log "$CLI_SHELL_TEST finished successfully"
    else
      log "$CLI_SHELL_TEST failed"
      exit 1
    fi

    collect_cover
    cleanup_pool
  done
fi

# Go rest tests
if [[ $TESTS_TO_RUN =~ "all" || $TESTS_TO_RUN =~ "rest" ]]; then
  GO_REST_TESTS="$(find integration_tests/grpc_rest -type f -name '*_test.go')"

  for GO_REST_TEST in ${GO_REST_TESTS}; do
    init_pool

    log "*****************************************************************************************"
    log "Running $GO_REST_TEST"
    log "*****************************************************************************************"

    # TODO issue 99: improve, that await helps with the cases of not ready connections to Cosmos endpoints
    sleep 5

    dcld config keyring-backend test
    if go test "$GO_REST_TEST" &>${DETAILED_OUTPUT_TARGET}; then
      log "$GO_REST_TEST finished successfully"
    else
      log "$GO_REST_TEST failed"
      exit 1
    fi

    collect_cover
    cleanup_pool
  done
fi
