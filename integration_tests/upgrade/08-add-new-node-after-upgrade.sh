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
source integration_tests/cli/common.sh

localnet_dir=".localnet"
dcl_user_home="/var/lib/dcl"
DCL_DIR="$dcl_user_home/.dcl"

node_p2p_port=26570
node_client_port=26571
chain_id="dclchain"
ip="192.167.10.28"
docker_network="distributed-compliance-ledger_localnet"
binary_version="0.12.0"

MASTER_UPGRADE_DOCKERFILE="./integration_tests/upgrade/Dockerfile-build-master"
MASTER_UPGRADE_IMAGE="dcld-build-master"
MASTER_UPGRADE_CONTAINER_NAME="$MASTER_UPGRADE_IMAGE-inst"

function check_expected_catching_up_status_for_interval {
    local expected_status="$1"
    local overall_ping_time_sec="${2:-100}"
    local sleep_time_sec="${3:-1}"
    local seconds=0

    while [ $seconds -lt $overall_ping_time_sec ]; do

        if [ $( docker container ls -a | grep "$NEW_OBSERVER_CONTAINER_NAME" | wc -l ) -eq 0 ]; then
            echo "! docker container inspect"
            break
        fi

        if ! docker container inspect "$NEW_OBSERVER_CONTAINER_NAME" | grep -q '"Status": "running"'; then
            break
        fi

        local dcld_status=$(docker exec --user root "$NEW_OBSERVER_CONTAINER_NAME" dcld status 2>&1)

        status_substring="\"catching_up\":$expected_status"
        if [[ $dcld_status == *"$status_substring"* ]]; then
            echo -e "dcld status:\n$dcld_status"
            return 0
        fi

        sleep $sleep_time_sec
        local seconds=$((seconds+1))
    done

    return 1
}

function check_expected_version_for_interval {
    local expected_version="$1"
    local overall_ping_time_sec="${2:-10}"
    local sleep_time_sec="${3:-1}"
    local seconds=0

    while [ $seconds -lt $overall_ping_time_sec ]; do

        if [ $( docker container ls -a | grep "$NEW_OBSERVER_CONTAINER_NAME" | wc -l ) -eq 0 ]; then
            break
        fi

        if ! docker container inspect "$NEW_OBSERVER_CONTAINER_NAME" | grep -q '"Status": "running"'; then
            break
        fi

        local dcld_version=$(docker exec "$NEW_OBSERVER_CONTAINER_NAME" dcld version 2>&1)
        echo "dcld_version = $dcld_version"

        if [ "$dcld_version" == "$expected_version" ]; then
            echo "dcld_version = $dcld_version"
            return 0
        fi
        
        sleep $sleep_time_sec
        local seconds=$((seconds+1))
    done

    return 1
}

cleanup_container $NEW_OBSERVER_CONTAINER_NAME

echo "1. run $NEW_OBSERVER_CONTAINER_NAME container"
docker run -d --name "$NEW_OBSERVER_CONTAINER_NAME" --ip $ip -p "$node_p2p_port-$node_client_port:26656-26657" --network $docker_network -i dcledger

test_divider

echo "2. install dcld v$binary_version to $NEW_OBSERVER_CONTAINER_NAME"
wget "https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v$binary_version/dcld"
chmod ugo+x dcld
docker cp ./dcld "$NEW_OBSERVER_CONTAINER_NAME":"$dcl_user_home"/
rm -f ./dcld

test_divider

echo "3. Set up configuration files for $NEW_OBSERVER_CONTAINER_NAME"
docker exec "$NEW_OBSERVER_CONTAINER_NAME" ./dcld init "$NEW_OBSERVER_CONTAINER_NAME" --chain-id $chain_id
docker cp "$localnet_dir/node0/config/genesis.json" $NEW_OBSERVER_CONTAINER_NAME:$DCL_DIR/config
peers="$(cat "$localnet_dir/node0/config/config.toml" | grep -o -E "persistent_peers = \".*\"")"
docker exec "$NEW_OBSERVER_CONTAINER_NAME" sed -i "s/persistent_peers = \"\"/$peers/g" $DCL_DIR/config/config.toml
docker exec "$NEW_OBSERVER_CONTAINER_NAME" sed -i 's/laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/g' $DCL_DIR/config/config.toml

test_divider

echo "4. Locate the app to $DCL_DIR/cosmovisor/genesis/bin directory in $NEW_OBSERVER_CONTAINER_NAME"
docker exec "$NEW_OBSERVER_CONTAINER_NAME" mkdir -p "$DCL_DIR"/cosmovisor/genesis/bin
docker exec "$NEW_OBSERVER_CONTAINER_NAME" cp -f ./dcld "$DCL_DIR"/cosmovisor/genesis/bin/

test_divider

echo "5. Set up master upgrade for $NEW_OBSERVER_CONTAINER_NAME"
docker cp "$MASTER_UPGRADE_CONTAINER_NAME":/go/bin/dcld ./dcld_master
docker cp ./dcld_master "$NEW_OBSERVER_CONTAINER_NAME":"$DCL_DIR"/dcld
master_upgrade_plan_name="$(docker run "$MASTER_UPGRADE_IMAGE" /bin/sh -c "cd /go/src/distributed-compliance-ledger && git rev-parse --short HEAD")"
docker exec "$NEW_OBSERVER_CONTAINER_NAME" /bin/sh -c "cosmovisor add-upgrade "$master_upgrade_plan_name" "$DCL_DIR"/dcld"
docker rm "$MASTER_UPGRADE_CONTAINER_NAME"

test_divider

echo "6. Start Node \"$NEW_OBSERVER_CONTAINER_NAME\""
docker exec -d "$NEW_OBSERVER_CONTAINER_NAME" sh -c "/var/lib/dcl/./node_helper.sh | tee /proc/1/fd/1"
docker logs -f "$NEW_OBSERVER_CONTAINER_NAME" &

test_divider

echo "7. Check dcld version == $binary_version in $NEW_OBSERVER_CONTAINER_NAME"

check_expected_version_for_interval "$binary_version"

is_version_correct=$?

if [ $is_version_correct == 1 ] ; then
    echo "installed dcld version does not match dcld mainnet version"
    exit 1
fi

test_divider

sleep_time_sec=1
overall_ping_time_sec=700

echo "8. Check node $NEW_OBSERVER_CONTAINER_NAME for START catching up process pinging it every $sleep_time_sec second for $overall_ping_time_sec seconds"

check_expected_catching_up_status_for_interval true $overall_ping_time_sec $sleep_time_sec
is_catching_up=$?

if [ $is_catching_up == 1 ] ; then
    echo "Catch-up procedure does not started"
    exit 1
fi

test_divider

echo "9. Check node $NEW_OBSERVER_CONTAINER_NAME for FINISH catching up process pinging it every $sleep_time_sec second for $overall_ping_time_sec seconds"

check_expected_catching_up_status_for_interval false $overall_ping_time_sec $sleep_time_sec
is_not_catching_up=$?

if [ $is_not_catching_up == 1 ] ; then
    echo "Catch-up procedure does not finished"
    exit 1
fi

test_divider

echo "10. Check node $NEW_OBSERVER_CONTAINER_NAME dcld updated to version $master_upgrade_plan_name"

check_expected_version_for_interval "$master_upgrade_plan_name"

is_version_correct=$?

if [ $is_version_correct == 1 ] ; then
    echo "installed dcld version does not match dcld expected version"
    exit 1
fi

echo "PASSED"