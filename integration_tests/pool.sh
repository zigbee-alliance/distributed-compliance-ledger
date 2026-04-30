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

SED_EXT=
if [ "$(uname)" == "Darwin" ]; then
    SED_EXT="''"
fi

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
  DETAILED_OUTPUT_TARGET="${DETAILED_OUTPUT_TARGET:-/dev/stdout}"

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
  DETAILED_OUTPUT_TARGET="${DETAILED_OUTPUT_TARGET:-/dev/stdout}"
  log "Cleaning up pool"
  log "-> Stopping pool & Removing configurations" >${DETAILED_OUTPUT_TARGET}
  make localnet_clean
}
