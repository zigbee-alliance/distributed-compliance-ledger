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

log() {
  echo "${LOG_PREFIX}$1"
}

setup_environment() {
  log "Setting up environment"

  log "-> Compiling dcl binaries"
  make install
}

setup_test() {
  log "Setting up test environment:"

  log "-> Generating network configuration"
  make localnet_init

  log "-> Running pool"
  make localnet_start
}

cleanup_test() {
  log "Cleaning up test:"

  log "-> Removing configurations"
  rm -rf ~/.dclcli
  rm -rf ~/.dcld
  rm -rf localnet
}


setup_environment
cleanup_test

SHELL_TESTS=$(find integration_tests/cli -type f -not -name "common.sh")

for SHELL_TEST in ${SHELL_TESTS}; do
  setup_test

  log "Running $SHELL_TEST"
  sh "$SHELL_TEST"

  cleanup_test
done
