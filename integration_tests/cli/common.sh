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

check_env() {
    jq --version || (echo "jq tool is not found" && exit 1)
}

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

  echo "Account name: $name"

  echo "Generate key for $name"
  (echo $passphrase; echo $passphrase) | dcld keys add "$name"

  address=$(echo $passphrase | dcld keys show $name -a)
  pubkey=$(echo $passphrase | dcld keys show $name -p)

  echo "Jack proposes account for \"$name\" with roles: \"$roles\""
  result=$(echo $passphrase | dcld tx auth propose-add-account --address="$address" --pubkey="$pubkey" --roles=$roles --from jack --yes)
  check_response "$result" "\"code\": 0"
  echo "$result"

  echo "Alice approves account for \"$name\" with roles: \"$roles\""
  result=$(echo $passphrase | dcld tx auth approve-add-account --address="$address" --from alice --yes)
  check_response "$result" "\"code\": 0"
  echo "$result"
}

create_new_vendor_account(){

  local _name="$1"
  local _vid="$2"

  echo $passphrase | dcld keys add "$_name"
  address=$(echo $passphrase | dcld keys show $_name -a)
  pubkey=$(echo $passphrase | dcld keys show $_name -p)

  echo "Jack proposes account for \"$_name\" with Vendor role"
  result=$(echo $passphrase | dcld tx auth propose-add-account --address="$address" --pubkey="$pubkey" --roles=Vendor --vid=$_vid --from jack --yes)
  check_response "$result" "\"code\": 0"

  echo "Alice approves account for \"$_name\" with Vendor role"
  result=$(echo $passphrase | dcld tx auth approve-add-account --address="$address" --from alice --yes)
  check_response "$result" "\"code\": 0"

}

create_model_and_version() {
  local _vid="$1"
  local _pid="$2"
  local _softwareVersion="$3"
  local _softwareVersionString="$4"
  local _user_address="$5"
  result=$(echo "$passphrase" | dcld tx model add-model --vid=$_vid --pid=$_pid --deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0 --from=$_user_address --yes)
  check_response "$result" "\"code\": 0"
  result=$(echo "$passphrase" | dcld tx model add-model-version --cdVersionNumber=1 --maxApplicableSoftwareVersion=10 --minApplicableSoftwareVersion=1 --vid=$_vid --pid=$_pid --softwareVersion=$_softwareVersion --softwareVersionString=$_softwareVersionString --from=$_user_address --yes)
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

# TODO: see https://github.com/zigbee-alliance/distributed-compliance-ledger/issues/203
execute_with_retry() {
  local _command=${1}

  local _result=$($_command)

  for i in {1..10}; do
    if [[ "$(_check_response "$_result" "EOF" "raw")" == true ]]; then
      #echo "EOF detected, re-trying"
      _result=$($_command)
    else
      break  
    fi
  done

  echo "$_result"
}
