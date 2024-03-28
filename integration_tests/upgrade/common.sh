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

set -euo pipefail

# common constants
node_p2p_port=26670
node_client_port=26671
chain_id="dclchain"
node0conn="tcp://192.167.10.2:26657"
docker_network="distributed-compliance-ledger_localnet"
passphrase="test1234"
LOCALNET_DIR=".localnet"

# RED=`tput setaf 1`
# GREEN=`tput setaf 2`
# RESET=`tput sgr0`
GREEN=""
RED=""
RESET=""

random_string() {
  local __resultvar="$1"
  local length=${2:-6} # Default is 6
  # Newer mac might have shasum instead of sha1sum
  if  command -v shasum &> /dev/null
    then
      eval $__resultvar="'$(date +%s.%N | shasum | fold -w ${length} | head -n 1)'"
    else
      eval $__resultvar="'$(date +%s.%N | sha1sum | fold -w ${length} | head -n 1)'"
  fi
}

DEF_OUTPUT_MODE=json


# json: pretty (indented) json
# raw or otherwise: raw
_check_response() {
    local _result="$1"
    local _expected_string="$2"
    local _mode="${3:-$DEF_OUTPUT_MODE}"

    if [[ "$_mode" == "json" ]]; then
        if [[ -n "$(echo "$_result" | jq | grep "$_expected_string" 2>/dev/null)" ]]; then
            echo true
            return
        fi
    else
        if [[ -n "$(echo "$_result" | grep "$_expected_string" 2>/dev/null)" ]]; then
            echo true
            return
        fi
    fi

    echo false
}

check_response() {
    local _result="$1"
    local _expected_string="$2"
    local _mode="${3:-$DEF_OUTPUT_MODE}"

    if [[ "$(_check_response "$_result" "$_expected_string" "$_mode")" != true ]]; then
        echo "${GREEN}ERROR:${RESET} command failed. The expected string: '$_expected_string' not found in the result: $_result"
        exit 1
    fi
}

check_response_and_report() {
    local _result="$1"
    local _expected_string="$2"
    local _mode="${3:-$DEF_OUTPUT_MODE}"

    check_response "$_result" "$_expected_string" "$_mode"
    echo "${GREEN}SUCCESS: ${RESET} Result contains expected substring: '$_expected_string'"
}

response_does_not_contain() {
    local _result="$1"
    local _unexpected_string="$2"
    local _mode="${3:-$DEF_OUTPUT_MODE}"

    if [[ "$(_check_response "$_result" "$_unexpected_string" "$_mode")" == true ]]; then
        echo "ERROR: command failed. The unexpected string: '$_unexpected_string' found in the result: $_result"
        exit 1
    fi

    echo "${GREEN}SUCCESS: ${RESET}Result does not contain unexpected substring: '$_unexpected_string'"
}

create_new_account(){
  local __resultvar="$1"
  random_string name
  eval $__resultvar="'$name'"

  local roles="$2"
  local _dcld="${3:-dcld}"
  echo "Account name: $name"

  echo "Generate key for $name"
  (echo $passphrase; echo $passphrase) | $_dcld keys add "$name"

  address=$(echo $passphrase | $_dcld keys show $name -a)
  pubkey=$(echo $passphrase | $_dcld keys show $name -p)

  echo "Jack proposes account for \"$name\" with roles: \"$roles\""
  result=$(echo $passphrase | $_dcld tx auth propose-add-account --address="$address" --pubkey="$pubkey" --roles=$roles --from jack --yes)
  check_response "$result" "\"code\": 0"
  echo "$result"

  echo "Alice approves account for \"$name\" with roles: \"$roles\""
  result=$(echo $passphrase | $_dcld tx auth approve-add-account --address="$address" --from alice --yes)
  check_response "$result" "\"code\": 0"
  echo "$result"
}

create_new_vendor_account(){

  local _name="$1"
  local _vid="$2"
  local _dcld="${3:-dcld}"

  echo $passphrase | $_dcld keys add "$_name"
  _address=$(echo $passphrase | $_dcld keys show $_name -a)
  _pubkey=$(echo $passphrase | $_dcld keys show $_name -p)

  result=$(echo $passphrase | $_dcld tx auth propose-add-account --address="$_address" --pubkey="$_pubkey" --roles=Vendor --vid=$_vid --from jack --yes)
  check_response "$result" "\"code\": 0"

}


test_divider() {
  echo ""
  echo "--------------------------"
  echo ""
}

get_height() {
  local __resultvar="$1"
  eval $__resultvar="'$(dcld status | jq | grep latest_block_height | awk -F'"' '{print $4}')'"
}

wait_for_height() {
  local target_height="${1:-1}" # Default is 1
  local wait_time="${2:-10}"    # In seconds, default - 10
  local mode="${3:-normal}"     # normal or outage-safe
  local node="${4:-""}"

  local _output=${DETAILED_OUTPUT_TARGET:-/dev/stdout}

  local waited=0
  local wait_interval=1

  if [[ -n "$node" ]]; then
      node="--node $node"
  fi

  while true; do
    sleep "${wait_interval}"
    waited=$((waited + wait_interval))

    if [[ "$mode" == "outage-safe" ]]; then
      current_height="$(dcld status $node 2>/dev/null | jq | grep latest_block_height | awk -F'"' '{print $4}')" || true
    else
      current_height="$(dcld status $node | jq | grep latest_block_height | awk -F'"' '{print $4}')"

      if [[ -z "$current_height" ]]; then
        echo "No height found in status"
        exit 1
      fi
    fi

    if [[ -n "$current_height" ]] && ((current_height >= target_height)); then
      echo "Height $target_height is reached in $waited seconds" &>${_output}
      break
    fi

    if ((waited > wait_time)); then
      echo "Height $target_height is not reached in $wait_time seconds"
      exit 1
    fi

    echo "Waiting for height: $target_height... Current height: ${current_height:-unavailable}, " \
      "wait time: $waited, time limit: $wait_time." &>${_output}
  done
}

get_txn_result() {
  local _broadcast_result=${1}
  local _txHash=$(echo "$_broadcast_result" | jq -r '.txhash')
  local _command="dcld query tx $_txHash"
  local _result=$($_command 2>&1)

  for i in {1..20}; do
    if [[ "$(_check_response "$_result" "not found" "raw")" == true ]]; then
      sleep 2
      _result=$($_command 2>&1)
    else
      break
    fi
  done

  echo "$_result"
}

execute_with_retry() {
  local _command=${1}
  local _error=${2:-"EOF"}
  local _result=$($_command)

  for i in {1..20}; do
    if [[ "$(_check_response "$_result" $_error "raw")" == true ]]; then
      #echo "EOF detected, re-trying"
      sleep 2
      _result=$($_command)
    else
      break
    fi
  done

  echo "$_result"
}

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

start_pool() {
  log "Setting up pool"

  log "-> Generating network configuration" >${DETAILED_OUTPUT_TARGET}
  make localnet_init_latest_stable_release MAINNET_STABLE_VERSION=$binary_version_old &>${DETAILED_OUTPUT_TARGET}

  patch_consensus_config

  log "-> Running pool" >${DETAILED_OUTPUT_TARGET}
  make localnet_start &>${DETAILED_OUTPUT_TARGET}

  log "-> Waiting for the second block (needed to request proofs)" >${DETAILED_OUTPUT_TARGET}
  execute_with_retry "dcld status" "connection"
  wait_for_height 2 20
}

cleanup() {
  if docker container ls -a | grep -q $container; then
    if docker container inspect $container | grep -q '"Status": "running"'; then
      echo "Stopping container"
      docker container kill $container
    fi

    echo "Removing container"
    docker container rm -f "$container"
  fi
}

container="validator-demo"
add_validator_node() {
  # FIXME: as it's called before upgrade, mainnet stable version of dcld needs to be used (not the latest master)
  # FIXME: check adding new node after upgrade as well
  random_string account
  address=""
  LOCALNET_DIR=".localnet"
  DCL_USER_HOME="/var/lib/dcl"
  DCL_DIR="$DCL_USER_HOME/.dcl"

  node_name="node-demo"
  node_p2p_port=26670
  node_client_port=26671
  chain_id="dclchain"
  ip="192.167.10.6"
  node0conn="tcp://192.167.10.2:26657"
  passphrase="test1234"
  docker_network="distributed-compliance-ledger_localnet"

  docker run -d --name $container --ip $ip -p "$node_p2p_port-$node_client_port:26656-26657" --network $docker_network -i dcledger

  docker cp "$DCLD_BIN_OLD" "$container":"$DCL_USER_HOME"/dcld

  test_divider

  echo "$account Configure CLI"
  docker exec $container /bin/sh -c "
    ./dcld config chain-id dclchain &&
    ./dcld config output json &&
    ./dcld config node $node0conn &&
    ./dcld config keyring-backend test &&
    ./dcld config broadcast-mode block"

  test_divider

  echo "$account Prepare Node configuration files"
  docker exec $container ./dcld init $node_name --chain-id $chain_id
  docker cp "$LOCALNET_DIR/node0/config/genesis.json" $container:$DCL_DIR/config
  peers="$(cat "$LOCALNET_DIR/node0/config/config.toml" | grep -o -E "persistent_peers = \".*\"")"
  docker exec $container sed -i "s/persistent_peers = \"\"/$peers/g" $DCL_DIR/config/config.toml
  docker exec $container sed -i 's/laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/g' $DCL_DIR/config/config.toml

  test_divider

  echo "Generate keys for $account"
  cmd="(echo $passphrase; echo $passphrase) | ./dcld keys add $account"
  docker exec $container /bin/sh -c "$cmd"

  address="$(docker exec $container /bin/sh -c "echo $passphrase | ./dcld keys show $account -a")"
  pubkey="$(docker exec $container /bin/sh -c "echo $passphrase | ./dcld keys show $account -p")"
  alice_address="$($DCLD_BIN_OLD keys show alice -a)"
  bob_address="$($DCLD_BIN_OLD keys show bob -a)"
  jack_address="$($DCLD_BIN_OLD keys show jack -a)"
  echo "Create account for $account and Assign NodeAdmin role"
  echo $passphrase | $DCLD_BIN_OLD tx auth propose-add-account --address="$address" --pubkey="$pubkey" --roles="NodeAdmin" --from jack --yes
  echo $passphrase | $DCLD_BIN_OLD tx auth approve-add-account --address="$address" --from alice --yes
  echo $passphrase | $DCLD_BIN_OLD tx auth approve-add-account --address="$address" --from bob --yes

  test_divider
  vaddress=$(docker exec $container ./dcld tendermint show-address)
  vpubkey=$(docker exec $container ./dcld tendermint show-validator)

  echo "Check pool response for yet unknown node \"$node_name\""
  result=$($DCLD_BIN_OLD query validator node --address "$address")
  check_response "$result" "Not Found"
  echo "$result"
  result=$($DCLD_BIN_OLD query validator last-power --address "$address")
  check_response "$result" "Not Found"
  echo "$result"

  echo "$account Add Node \"$node_name\" to validator set"
  cmd="$(docker exec "$container" /bin/sh -c "echo test1234 | ./dcld tx validator add-node --pubkey='$vpubkey' --moniker="$node_name" --from="$account" --yes")"
  result=$(execute_with_retry "$cmd" "not found")
  echo "$result"

  test_divider

  echo "Locating the app to $DCL_DIR/cosmovisor/genesis/bin directory"
  docker exec $container mkdir -p "$DCL_DIR"/cosmovisor/genesis/bin
  docker exec $container cp -f ./dcld "$DCL_DIR"/cosmovisor/genesis/bin/

  echo "$account Start Node \"$node_name\""
  docker exec -d $container cosmovisor run start
  sleep 10

  result=$($DCLD_BIN_OLD query validator node --address "$address")
  validator_address=$(echo "$result" | jq -r '.owner')
  echo "$result"
}