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

source integration_tests/cli/common.sh
set +euo pipefail

localnet_dir=".localnet"
dcl_user_home="/var/lib/dcl"
DCL_DIR="$dcl_user_home/.dcl"

node_name="new-observer"
node_p2p_port=26570
node_client_port=26571
chain_id="dclchain"
ip="192.167.10.28"
docker_network="distributed-compliance-ledger_localnet"

function check_expected_catching_up_status_for_interval {
    local expected_status="$1"
    local overall_ping_time_sec="${2:-100}"
    local sleep_time_sec="${3:-1}"
    local seconds=0

    while [ $seconds -lt $overall_ping_time_sec ]; do
        local dcld_status=$(docker exec --user root $node_name dcld status 2>&1)
        
        status_substring="\"catching_up\":$expected_status"
        if [[ $dcld_status == *"$status_substring"* ]]; then
            echo -e "dcld status:\n$dcld_status"
            return 1
        fi
        
        sleep $sleep_time_sec
        local seconds=$((seconds+1))
    done

    return 0
}

cleanup() {
    if docker container ls -a | grep -q $node_name; then
      if docker container inspect $node_name | grep -q '"Status": "running"'; then
        echo "Stopping container"
        docker container kill $node_name
      fi

      echo "Removing container"
      docker container rm -f "$node_name"
    fi
}

# trap cleanup EXIT

check_adding_new_node() {
  local stable_binary_version="${1:-0.12.1}"
  local latest_binary_version="${2:-1.4.4}"

  echo "1. run $node_name container"
  docker run -d --name $node_name --ip $ip -p "$node_p2p_port-$node_client_port:26656-26657" --network $docker_network -i dcledger

  test_divider

  echo "2. install dcld v$stable_binary_version to $node_name"
  wget "https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v$stable_binary_version/dcld"
  chmod ugo+x dcld
  docker cp ./dcld "$node_name":"$dcl_user_home"/
  rm -f ./dcld

  test_divider

  echo "3. Set up configuration files for $node_name"
  docker exec $node_name ./dcld init $node_name --chain-id $chain_id
  docker cp "$localnet_dir/node0/config/genesis.json" $node_name:$DCL_DIR/config
  peers="$(cat "$localnet_dir/node0/config/config.toml" | grep -o -E "persistent_peers = \".*\"")"
  docker exec $node_name sed -i "s/persistent_peers = \"\"/$peers/g" $DCL_DIR/config/config.toml
  docker exec $node_name sed -i 's/laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/g' $DCL_DIR/config/config.toml

  test_divider

  echo "4. Locate the app to $DCL_DIR/cosmovisor/genesis/bin directory in $node_name"
  docker exec $node_name mkdir -p "$DCL_DIR"/cosmovisor/genesis/bin
  docker exec $node_name cp -f ./dcld "$DCL_DIR"/cosmovisor/genesis/bin/

  test_divider

  echo "5. Start Node \"$node_name\""
  docker exec -d $node_name /var/lib/dcl/./node_helper.sh

  test_divider

  sleep 1

  echo "6. Check dcld version == $stable_binary_version in $node_name"
  dcld_version=$(docker exec $node_name dcld version)
  echo "dcld_version = $dcld_version"
  if [ "$dcld_version" != $stable_binary_version ]; then
      echo "installed dcld version $dcld_version != dcld mainnet version $stable_binary_version"
      exit 1
  fi

  test_divider

  sleep_time_sec=1
  overall_ping_time_sec=700

  echo "7. Check node $node_name for START catching up process pinging it every $sleep_time_sec second for $overall_ping_time_sec seconds"

  check_expected_catching_up_status_for_interval true $overall_ping_time_sec $sleep_time_sec
  is_catching_up=$?

  if [ $is_catching_up == 0 ] ; then
      echo "Catch-up procedure does not started"
      exit 1
  fi

  test_divider

  echo "8. Check node $node_name for FINISH catching up process pinging it every $sleep_time_sec second for $overall_ping_time_sec seconds"

  check_expected_catching_up_status_for_interval false $overall_ping_time_sec $sleep_time_sec
  is_not_catching_up=$?

  if [ $is_not_catching_up == 0 ] ; then
      echo "Catch-up procedure does not finished"
      exit 1
  fi

  test_divider

  echo "9. Check node $node_name dcld updated to version $latest_binary_version"
  dcld_version=$(docker exec $node_name dcld version)
  echo "dcld_version = $dcld_version"
  if [ "$dcld_version" != "$latest_binary_version" ]; then
      echo "installed dcld version $dcld_version != dcld expected version $latest_binary_version"
      exit 1
  fi

  echo "PASSED"

  cleanup
}