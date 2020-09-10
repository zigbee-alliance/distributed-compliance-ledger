#!/usr/bin/bash
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

LOG_PREFIX="[run all] "
DETAILED_OUTPUT_TARGET=/dev/null # Use /dev/stdout for debug

log() {
  echo "${LOG_PREFIX}$1"
}

wait_for_height() {
  local target_height=${1:-1} # Default is 1
  local wait_time=${2:-10}    # In seconds, default - 10

  local waited=0
  local wait_interval=1

  while true; do
    sleep "${wait_interval}"
    waited=$((waited + wait_interval))

    current_height=$(dclcli status | grep height | sed 's/.*"\(.*\)".*/\1/')

    if ((current_height >= target_height)); then
      echo "Height $target_height is reached in $waited seconds" &>${DETAILED_OUTPUT_TARGET}
      break
    fi

    if ((waited > wait_time)); then
      echo "Height $target_height is not reached in $wait_time seconds" &>${DETAILED_OUTPUT_TARGET}
      exit 1
    fi

    echo "Waiting for height: $target_height... Current height: $current_height, " \
      "wait time: $waited, time limit: $wait_time." &>${DETAILED_OUTPUT_TARGET}
  done
}

patch_consensus_config() {
  local NODE_CONFIGS=$(find localnet -type f -name "config.toml" -wholename "*node*")

  for NODE_CONFIG in ${NODE_CONFIGS}; do
    sed -i 's/timeout_propose = "3s"/timeout_propose = "500ms"/g' "${NODE_CONFIG}"
    sed -i 's/timeout_prevote = "1s"/timeout_prevote = "500ms"/g' "${NODE_CONFIG}"
    sed -i 's/timeout_precommit = "1s"/timeout_precommit = "500ms"/g' "${NODE_CONFIG}"
    sed -i 's/timeout_commit = "5s"/timeout_commit = "500ms"/g' "${NODE_CONFIG}"
  done
}

init_pool() {
  log "Setting up pool"

  log "-> Generating network configuration" >${DETAILED_OUTPUT_TARGET}
  make localnet_init &>${DETAILED_OUTPUT_TARGET}
  patch_consensus_config

  log "-> Running pool" >${DETAILED_OUTPUT_TARGET}
  make localnet_start &>${DETAILED_OUTPUT_TARGET}

  log "-> Waiting for the second block (needed to request proofs)" >${DETAILED_OUTPUT_TARGET}
  wait_for_height 2 20
}

cleanup_pool() {
  log "Cleaning up pool"

  log "-> Stopping pool" >${DETAILED_OUTPUT_TARGET}
  make localnet_stop &>${DETAILED_OUTPUT_TARGET}

  log "-> Removing configurations" >${DETAILED_OUTPUT_TARGET}
  rm -rf ~/.dclcli
  rm -rf ~/.dcld
  sudo rm -rf localnet
}

run_rest_server() {
  log "Running cli in rest-server mode"
  dclcli rest-server --chain-id dclchain &>>rest-server.out &
}

stop_rest_server() {
  log "Stopping cli in rest-server mode"
  killall dclcli
}

# Global init
set -o errexit
set -o pipefail

make install &>${DETAILED_OUTPUT_TARGET}

# Cli shell tests
CLI_SHELL_TESTS=$(find integration_tests/cli -type f -not -name "common.sh")

for CLI_SHELL_TEST in ${CLI_SHELL_TESTS}; do
  init_pool

  log "Running $CLI_SHELL_TEST"

  if bash "$CLI_SHELL_TEST" &>"test.out"; then
    cat "test.out" >${DETAILED_OUTPUT_TARGET}
    log "$CLI_SHELL_TEST finished successfully"
  else
    cat "test.out"
    log "$CLI_SHELL_TEST falied"
    exit 1
  fi

  cleanup_pool
done

# Go rest tests
GO_REST_TESTS=$(find integration_tests/rest -type f)

for GO_REST_TEST in ${GO_REST_TESTS}; do
  init_pool
  run_rest_server

  log "Running $GO_REST_TEST"

  if go test "$GO_REST_TEST" &>"test.out"; then
    cat "test.out" >${DETAILED_OUTPUT_TARGET}
    log "$GO_REST_TEST finished successfully"
  else
    cat "test.out"
    log "$GO_REST_TEST falied"
    exit 1
  fi

  stop_rest_server
  cleanup_pool
done
