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

LOCALNET_DIR=".localnet"
DCL_DIR="/root/.dcl"

random_string account
container="validator-demo"
node_name="node-demo"
chain_id="dclchain"
ip="192.167.10.6"
node0="tcp://192.167.10.2:26657"
passphrase="test1234"
docker_network="distributed-compliance-ledger_localnet"

cleanup() {
    if docker container ls -a | grep -q $container; then
      if docker container inspect $container | grep -q '"Status": "running"'; then
        echo "Stopping container"
        docker container kill $container
      fi

      echo "Removing container"
      docker container rm "$container"
    fi
}
trap cleanup EXIT

cleanup

docker run -d --name $container --ip $ip -p "26664-26665:26656-26657" --network $docker_network -i dcledger

test_divider

echo "$account Configure CLI"
docker exec $container /bin/sh -c "
  dcld config chain-id dclchain &&
  dcld config output json &&
  dcld config node $node0 &&
  dcld config keyring-backend file &&
  dcld config broadcast-mode block"

# TODO issue 99: check the replacement for the setting
#  dcld config indent true &&
#  dcld config trust-node false &&


test_divider

echo "$account Prepare Node configuration files"
docker exec $container dcld init $node_name --chain-id $chain_id
docker cp "$LOCALNET_DIR/node0/config/genesis.json" $container:$DCL_DIR/config
peers="$(cat "$LOCALNET_DIR/node0/config/config.toml" | grep -o -E "persistent_peers = \".*\"")"
docker exec $container sed -i "s/persistent_peers = \"\"/$peers/g" $DCL_DIR/config/config.toml
docker exec $container sed -i 's/laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/g' $DCL_DIR/config/config.toml

test_divider

echo "Generate keys for $account"
cmd="(echo $passphrase; echo $passphrase) | dcld keys add $account"
docker exec $container /bin/sh -c "$cmd"

address="$(docker exec $container /bin/sh -c "echo $passphrase | dcld keys show $account -a")"
pubkey="$(docker exec $container /bin/sh -c "echo $passphrase | dcld keys show $account -p")"
echo "Create account for $account and Assign NodeAdmin role"
echo $passphrase | dcld tx auth propose-add-account --address="$address" --pubkey="$pubkey" --roles="NodeAdmin" --from jack --yes
echo $passphrase | dcld tx auth approve-add-account --address="$address" --from alice --yes

test_divider

echo "$account Add Node \"$node_name\" to validator set"
vaddress=$(docker exec $container dcld tendermint show-address)
vpubkey=$(docker exec $container dcld tendermint show-validator)

! read -r -d '' _script << EOF
    set -eu; echo test1234 | dcld tx validator add-node --pubkey='$vpubkey' --name="$node_name" --from="$account" --yes
EOF
result="$(docker exec "$container" /bin/sh -c "echo test1234 | dcld tx validator add-node --pubkey='$vpubkey' --name="$node_name" --from="$account" --yes")"
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "$account Start Node \"$node_name\""
docker exec -d $container dcld start
sleep 10

test_divider

echo "Check node \"$node_name\" is in the validator set"
result=$(dcld query validator all-nodes)
check_response "$result" "\"name\": \"$node_name\""
#check_response "$result" "\"validator_address\": \"$vaddress\""
check_response "$result" "\"pubKey\":$vpubkey" raw
echo "$result"

test_divider

echo "Connect CLI to node \"$node_name\" and check status"
dcld config node "tcp://localhost:26665"
result=$(dcld status)
check_response "$result" "\"moniker\": \"$node_name\""
echo "$result"

test_divider

echo "Sent transactions using node \"$node_name\""
vid=$RANDOM
pid=$RANDOM
vendor_account=vendor_account_$vid
create_new_vendor_account $vendor_account $vid

test_divider

# FIXME issue 99: enable once implemented
exit 0

echo "Publish Model"
pid=$RANDOM
productName="TestingProductLabel"
echo "Add Model with VID: $vid PID: $pid"
result=$(echo 'test1234' | dcld tx model add-model --vid=$vid --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0 --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

sleep 5

echo "Connect CLI to node \"node0\""
dcld config node "tcp://localhost:26657"
result=$(dcld status)
check_response "$result" "\"moniker\": \"node0\""
echo "$result"

test_divider

echo "Query Model using node0 node"
echo "Get Model with VID: $vid PID: $pid"
result=$(dcld query model get-model --vid=$vid --pid=$pid)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"productLabel\": \"$productName\""
echo "$result"

cleanup
