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

random_string account
container="validator-demo"
node="node-demo"
chain_id="dclchain"
ip="192.167.10.6"
node0="tcp://192.167.10.2:26657"
passphrase="test1234"
docker_network="distributed-compliance-ledger_localnet"
if docker container ls -a | grep -q $container; then
  if docker container inspect $container | grep -q '"Status": "running"'; then
    echo "Stopping container"
    docker container kill $container
  fi

  echo "Removing container"
  docker container rm "$container"
fi

docker run -d --name $container --ip $ip -p "26664-26665:26656-26657" --network $docker_network -i dcledger

echo "Generate keys for $account"
docker exec $container /bin/sh -c "echo $passphrase | dclcli keys add $account"

test_divider

address=$(docker exec $container dclcli keys show $account -a)
pubkey=$(docker exec $container dclcli keys show $account -p)
echo "Create account for $account and Assign NodeAdmin role"
echo $passphrase | dclcli tx auth propose-add-account --address="$address" --pubkey="$pubkey" --roles="NodeAdmin" --from jack --yes
echo $passphrase | dclcli tx auth approve-add-account --address="$address" --from alice --yes

test_divider

echo "$account Prepare Node configuration files"
docker exec $container dcld init $node --chain-id $chain_id
docker cp ./localnet/node0/config/genesis.json $container:/root/.dcld/config
peers=$(cat localnet/node0/config/config.toml | grep -o -E "persistent_peers = \".*\"")
docker exec $container sed -i "s/persistent_peers = \"\"/$peers/g" /root/.dcld/config/config.toml
docker exec $container sed -i 's/laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/g' /root/.dcld/config/config.toml

test_divider

echo "$account Configure CLI"
docker exec $container /bin/sh -c "
  dclcli config chain-id dclchain &&
  dclcli config output json &&
  dclcli config indent true &&
  dclcli config trust-node false &&
  dclcli config node $node0"

test_divider

echo "$account Add Node \"$node\" to validator set"
vaddress=$(docker exec $container dcld tendermint show-address)
vpubkey=$(docker exec $container dcld tendermint show-validator)
result=$(docker exec $container /bin/sh -c "echo test1234 | dclcli tx validator add-node --validator-address=$vaddress --validator-pubkey=$vpubkey --name=$node --from=$account --yes")
check_response "$result" "\"success\": true"
echo "$result"

test_divider

echo "$account Start Node \"$node\""
docker exec -d $container dcld start
sleep 10

test_divider

echo "Check node \"$node\" is in the validator set"
result=$(dclcli query validator all-nodes)
check_response "$result" "\"name\": \"$node\""
check_response "$result" "\"validator_address\": \"$vaddress\""
check_response "$result" "\"validator_pubkey\": \"$vpubkey\""
echo "$result"

test_divider

echo "Connect CLI to node \"$node\" and check status"
dclcli config node "tcp://localhost:26665"
result=$(dclcli status)
check_response "$result" "\"moniker\": \"$node\""
echo "$result"

test_divider

echo "Sent transactions using node \"$node\""
vid=$RANDOM
pid=$RANDOM
vendor_account=vendor_account_$vid
create_new_vendor_account $vendor_account $vid

test_divider

echo "Publish Model"
pid=$RANDOM
productName="TestingProductLabel"
echo "Add Model with VID: $vid PID: $pid"
result=$(echo 'test1234' | dclcli tx model add-model --vid=$vid --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0 --from=$vendor_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

test_divider

sleep 5

echo "Connect CLI to node \"node0\""
dclcli config node "tcp://localhost:26657"
result=$(dclcli status)
check_response "$result" "\"moniker\": \"node0\""
echo "$result"

test_divider

echo "Query Model using node0 node"
echo "Get Model with VID: $vid PID: $pid"
result=$(dclcli query model get-model --vid=$vid --pid=$pid)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"productLabel\": \"$productName\""
echo "$result"

docker rm -f $container