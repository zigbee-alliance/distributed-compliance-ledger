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
DCL_USER_HOME="/var/lib/dcl"
DCL_DIR="$DCL_USER_HOME/.dcl"

random_string account
container="validator-demo"
node_name="node-demo"
node_p2p_port=26670
node_client_port=26671
chain_id="dclchain"
ip="192.167.10.6"
node0conn="tcp://192.167.10.2:26657"
passphrase="test1234"
docker_network="distributed-compliance-ledger_localnet"

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
trap cleanup EXIT

cleanup

docker build -f Dockerfile-build -t dcld-build .
docker container create --name dcld-build-inst dcld-build
docker cp dcld-build-inst:/go/bin/dcld ./
docker rm dcld-build-inst

docker run -d --name $container --ip $ip -p "$node_p2p_port-$node_client_port:26656-26657" --network $docker_network -i dcledger

docker cp ./dcld "$container":"$DCL_USER_HOME"/
rm -f ./dcld

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
alice_address="$(dcld keys show alice -a)"
bob_address="$(dcld keys show bob -a)"
jack_address="$(dcld keys show jack -a)"
echo "Create account for $account and Assign NodeAdmin role"
echo $passphrase | dcld tx auth propose-add-account --address="$address" --pubkey="$pubkey" --roles="NodeAdmin" --from jack --yes
echo $passphrase | dcld tx auth approve-add-account --address="$address" --from alice --yes


test_divider
vaddress=$(docker exec $container ./dcld tendermint show-address)
vpubkey=$(docker exec $container ./dcld tendermint show-validator)

echo "Check pool response for yet unknown node \"$node_name\""
result=$(dcld query validator node --address "$address")
check_response "$result" "Not Found"
echo "$result"
result=$(dcld query validator last-power --address "$address")
check_response "$result" "Not Found"
echo "$result"

echo "$account Add Node \"$node_name\" to validator set"

! read -r -d '' _script << EOF
    set -eu; echo test1234 | dcld tx validator add-node --pubkey='$vpubkey' --moniker="$node_name" --from="$account" --yes
EOF
result="$(docker exec "$container" /bin/sh -c "echo test1234 | ./dcld tx validator add-node --pubkey='$vpubkey' --moniker="$node_name" --from="$account" --yes")"
check_response "$result" "\"code\": 0"
echo "$result"


test_divider


echo "Locating the app to $DCL_DIR/cosmovisor/genesis/bin directory"
docker exec $container mkdir -p "$DCL_DIR"/cosmovisor/genesis/bin
docker exec $container cp -f ./dcld "$DCL_DIR"/cosmovisor/genesis/bin/

echo "$account Start Node \"$node_name\""
docker exec -d $container cosmovisor start
sleep 10


test_divider

echo "Check node \"$node_name\" is in the validator set"
result=$(dcld query validator all-nodes)
check_response "$result" "\"moniker\": \"$node_name\""
#check_response "$result" "\"validator_address\": \"$vaddress\""
check_response "$result" "\"pubKey\":$vpubkey" raw
echo "$result"


test_divider

result=$(dcld query validator node --address "$address")
validator_address=$(echo "$result" | jq -r '.owner')
echo "$result"

test_divider

echo "Connect CLI to node \"$node_name\" and check status"
dcld config node "tcp://localhost:$node_client_port"
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

echo "Publish Model"
pid=$RANDOM
productName="TestingProductLabel"
echo "Add Model with VID: $vid PID: $pid"
result=$(echo 'test1234' | dcld tx model add-model --vid=$vid --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel="$productName" --partNumber=1 --commissioningCustomFlow=0 --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"
echo "$result"


test_divider

sleep 5

echo "Connect CLI to node \"node0\""
dcld config node "$node0conn"
node0id=$(docker exec node0 dcld tendermint show-node-id)
result=$(dcld status)
# FIXME issue 99: moniker is returned wrongly
# check_response "$result" "\"moniker\": \"node0\""
check_response "$result" "\"id\": \"$node0id\""
echo "$result"


test_divider

echo "Query Model using node0 node"
echo "Get Model with VID: $vid PID: $pid"
result=$(dcld query model get-model --vid=$vid --pid=$pid)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"productLabel\": \"$productName\""
echo "$result"


test_divider

# TEST PROPOSE AND REJECT DISABLE VALIDATOR
echo "jack (Trustee) propose disable validator"
result=$(dcld tx validator propose-disable-node --address="$validator_address" --from alice --yes)
check_response "$result" "\"code\": 0"

echo "jack (Trustee) rejects disable validator"
result=$(dcld tx validator reject-disable-node --address="$validator_address" --from alice --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Propose not found in proposed disable validator"
result=$(dcld query validator proposed-disable-node --address="$validator_address")
echo $result | jq
check_response "$result" "Not Found"

test_divider

echo "Upgrade not found in rejected upgrade"
result=$(dcld query validator rejected-disable-node --address="$validator_address")
echo $result | jq
check_response "$result" "Not Found"

test_divider

echo "Upgrade not found in approved upgrade query"
result=$(dcld query validator disabled-node --address="$validator_address")
echo $result | jq
check_response "$result" "Not Found"

echo "node admin disables validator"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld tx validator disable-node --from "$account" --yes")
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "node admin doesn't add a new validator with new pubkey, because node admin already has disabled validator"
result="$(docker exec "$container" /bin/sh -c "echo test1234 | ./dcld tx validator add-node --pubkey='$pubkey' --moniker="$node_name" --from="$account" --yes 2>&1 || true")"
response_does_not_contain "$result" "\"code\": 0"

test_divider

echo "Get a validator $address from disabled-validator query"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator disabled-node --address=$address")
check_response "$result" "\"approvals\": \[\]"
check_response "$result" "\"disabledByNodeAdmin\": true"
check_response "$result" "\"address\": \"$validator_address\""
echo "$result"

test_divider

echo "Get a validator $validator_address from disabled-validator query"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator disabled-node --address=$validator_address")
check_response "$result" "\"approvals\": \[\]"
check_response "$result" "\"disabledByNodeAdmin\": true"
check_response "$result" "\"address\": \"$validator_address\""
echo "$result"

test_divider


echo "node admin enables validator"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld tx validator enable-node --from "$account" --yes")
check_response "$result" "\"code\": 0"
echo "$result"


test_divider


echo "Get a validator $address from disabled-validator query"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator disabled-node --address="$address"")
check_response "$result" "Not Found"
echo "$result"

test_divider

echo "Get a validator $validator_address from disabled-validator query"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator disabled-node --address="$validator_address"")
check_response "$result" "Not Found"
echo "$result"

test_divider

echo "Get a validator $address from node query"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator node --address="$address"")
check_response "$result" "\"moniker\": \"$node_name\""
check_response "$result" "\"pubKey\":$vpubkey" raw
check_response "$result" "\"owner\": \"$validator_address\""
echo "$result"

test_divider

echo "Alice proposes to disable validator $address"
result=$(dcld tx validator propose-disable-node --address="$validator_address" --from alice --yes)
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "Get all proposed validators to disable. $address in the list"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator all-proposed-disable-nodes")
check_response "$result" "\"approvals\":\[{\"address\":\"$alice_address\"" raw
check_response "$result" "\"address\": \"$validator_address\""
echo "$result"

test_divider


echo "Get a proposed validator to disable $address"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator proposed-disable-node --address="$address"")
check_response "$result" "\"approvals\":\[{\"address\":\"$alice_address\"" raw
check_response "$result" "\"address\": \"$validator_address\""
echo "$result"

test_divider

echo "Get a proposed validator to disable $validator_address"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator proposed-disable-node --address="$validator_address"")
check_response "$result" "\"approvals\":\[{\"address\":\"$alice_address\"" raw
check_response "$result" "\"address\": \"$validator_address\""
echo "$result"

test_divider


echo "Bob approves to disable validator $address"
result=$(dcld tx validator approve-disable-node --address="$validator_address" --from bob --yes)
check_response "$result" "\"code\": 0"
echo "$result"

test_divider


echo "node admin doesn't add a new validator with new pubkey, because node admin already has disabled validator"
result="$(docker exec "$container" /bin/sh -c "echo test1234 | ./dcld tx validator add-node --pubkey='$pubkey' --moniker="$node_name" --from="$account" --yes 2>&1 || true")"
response_does_not_contain "$result" "\"code\": 0"

test_divider

echo "node admin enables validator"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld tx validator enable-node --from "$account" --yes")
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "Get a validator $address from disabled-validator query"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator disabled-node --address="$address"")
check_response "$result" "Not Found"
echo "$result"

test_divider

echo "Get a validator $validator_address from disabled-validator query"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator disabled-node --address="$validator_address"")
check_response "$result" "Not Found"
echo "$result"

test_divider

echo "Get a validator $address from node query"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator node --address="$address"")
check_response "$result" "\"moniker\": \"$node_name\""
check_response "$result" "\"pubKey\":$vpubkey" raw
check_response "$result" "\"owner\": \"$validator_address\""
echo "$result"

test_divider

echo "Get a validator $validator_address from node query"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator node --address="$validator_address"")
check_response "$result" "\"moniker\": \"$node_name\""
check_response "$result" "\"pubKey\":$vpubkey" raw
check_response "$result" "\"owner\": \"$validator_address\""
echo "$result"

test_divider

echo "Alice proposes to disable validator $address"
result=$(dcld tx validator propose-disable-node --address="$validator_address" --from alice --yes)
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "Get a proposed validator to disable $address"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator proposed-disable-node --address="$address"")
check_response "$result" "\"approvals\":\[{\"address\":\"$alice_address\"" raw
check_response "$result" "\"address\": \"$validator_address\""
echo "$result"

test_divider

echo "Get a proposed validator to disable $validator_address"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator proposed-disable-node --address="$validator_address"")
check_response "$result" "\"approvals\":\[{\"address\":\"$alice_address\"" raw
check_response "$result" "\"address\": \"$validator_address\""
echo "$result"

test_divider


echo "Bob approves to disable validator $address"
result=$(dcld tx validator approve-disable-node --address="$validator_address" --from bob --yes)
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "Bob cannot reject to disable validator $address, because Bob already rejected to disable validator"
result=$(dcld tx validator reject-disable-node --address="$validator_address" --from bob --yes)
response_does_not_contain "$result" "\"code\": 0"
echo "$result"

test_divider

echo "Get all proposed validators to disable. $address not in the list"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator all-proposed-disable-nodes")
response_does_not_contain "$result" "\"address\": \"$address\""
echo "$result"


test_divider


echo "Get a validator $address from disabled-validator query"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator disabled-node --address="$address"")
check_response "$result" "\"creator\": \"$alice_address\""
check_response "$result" "\"disabledByNodeAdmin\": false"
check_response "$result" "\"address\": \"$validator_address\""
echo "$result"

test_divider

echo "Get a validator $validator_address from disabled-validator query"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator disabled-node --address="$validator_address"")
check_response "$result" "\"creator\": \"$alice_address\""
check_response "$result" "\"disabledByNodeAdmin\": false"
check_response "$result" "\"address\": \"$validator_address\""
echo "$result"


test_divider


echo "Get a validator $address"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator node --address=$address")
check_response "$result" "\"jailed\": true"
echo "$result"

test_divider

echo "Get a validator $validator_address"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator node --address=$validator_address")
check_response "$result" "\"jailed\": true"
echo "$result"

test_divider

echo "node admin enables validator"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld tx validator enable-node --from "$account" --yes")
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

random_string new_trustee1
echo "$new_trustee1 generates keys"
cmd="(echo $passphrase; echo $passphrase) | dcld keys add $new_trustee1"
result="$(bash -c "$cmd")"

test_divider

echo "Get key info for $new_trustee1"
result=$(echo $passphrase | dcld keys show $new_trustee1)
check_response "$result" "\"name\": \"$new_trustee1\""

test_divider

new_trustee_address1=$(echo $passphrase | dcld keys show $new_trustee1 -a)
new_trustee_pubkey1=$(echo $passphrase | dcld keys show $new_trustee1 -p)

test_divider

echo "Jack proposes account for $new_trustee1"
result=$(echo $passphrase | dcld tx auth propose-add-account --info="Jack is proposing this account" --address="$new_trustee_address1" --pubkey="$new_trustee_pubkey1" --roles="Trustee" --from jack --yes)
check_response "$result" "\"code\": 0"

echo "Alice approves account for \"$new_trustee1\""
result=$(echo $passphrase | dcld tx auth approve-add-account --address="$new_trustee_address1" --info="Alice is approving this account" --from alice --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Alice proposes to disable validator $address"
result=$(dcld tx validator propose-disable-node --address="$validator_address" --from alice --yes)
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "Bob approves to disable validator $address"
result=$(dcld tx validator approve-disable-node --address="$validator_address" --from bob --yes)
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "Bob can revote to reject disable validator $address even if Bob already approves to disable validator"
result=$(dcld tx validator reject-disable-node --address="$validator_address" --from bob --yes)
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "Bob approves to disable validator $address"
result=$(dcld tx validator approve-disable-node --address="$validator_address" --from bob --yes)
check_response "$result" "\"code\": 0"
echo "$result"

echo "Alice proposes to revoke account for $new_trustee1"
result=$(echo $passphrase | dcld tx auth propose-revoke-account --address="$new_trustee_address1" --from alice --yes)
check_response "$result" "\"code\": 0"

echo "Bob approves to revoke account for $new_trustee1"
result=$(echo $passphrase | dcld tx auth approve-revoke-account --address="$new_trustee_address1" --from bob --yes)
check_response "$result" "\"code\": 0"

echo "Jack approves to revoke account for $new_trustee1"
result=$(echo $passphrase | dcld tx auth approve-revoke-account --address="$new_trustee_address1" --from jack --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Get all proposed validators to disable. $address in the list"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator all-proposed-disable-nodes")
check_response "$result" "\"approvals\":\[{\"address\":\"$alice_address\"" raw
check_response "$result" "\"address\": \"$validator_address\""
echo "$result"

test_divider

echo "Get a proposed validator to disable $address"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator proposed-disable-node --address="$address"")
check_response "$result" "\"approvals\":\[{\"address\":\"$alice_address\"" raw
check_response "$result" "\"address\": \"$validator_address\""
echo "$result"

test_divider

echo "Get a proposed validator to disable $validator_address"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator proposed-disable-node --address="$validator_address"")
check_response "$result" "\"approvals\":\[{\"address\":\"$alice_address\"" raw
check_response "$result" "\"address\": \"$validator_address\""
echo "$result"

test_divider

echo "Bob rejects to disable validator $address"
result=$(dcld tx validator reject-disable-node --address="$validator_address" --from bob --yes)
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "Bob cannot reject to disable validator $address, because Bob already rejected to disable validator"
result=$(dcld tx validator reject-disable-node --address="$validator_address" --from bob --yes)
response_does_not_contain "$result" "\"code\": 0"
echo "$result"

test_divider

echo "Get a validator $address from disabled-validator query"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator disabled-node --address="$address"")
response_does_not_contain "$result" "\"creator\": \"$alice_address\""
response_does_not_contain "$result" "\"disabledByNodeAdmin\": false"
response_does_not_contain "$result" "\"address\": \"$validator_address\""
echo "$result"

test_divider

echo "Get all proposed validators to disable. $address in the list"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator all-proposed-disable-nodes")
check_response "$result" "\"approvals\":\[{\"address\":\"$alice_address\"" raw
check_response "$result" "\"rejects\":\[{\"address\":\"$bob_address\"" raw
check_response "$result" "\"address\": \"$validator_address\""
echo "$result"

test_divider

echo "Get a proposed validator to disable $address"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator proposed-disable-node --address="$address"")
check_response "$result" "\"approvals\":\[{\"address\":\"$alice_address\"" raw
check_response "$result" "\"rejects\":\[{\"address\":\"$bob_address\"" raw
check_response "$result" "\"address\": \"$validator_address\""
echo "$result"

test_divider

echo "Get a proposed validator to disable $validator_address"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator proposed-disable-node --address="$validator_address"")
check_response "$result" "\"approvals\":\[{\"address\":\"$alice_address\"" raw
check_response "$result" "\"rejects\":\[{\"address\":\"$bob_address\"" raw
check_response "$result" "\"address\": \"$validator_address\""
echo "$result"

test_divider

echo "Jack rejects to disable validator $address"
result=$(dcld tx validator reject-disable-node --address="$validator_address" --from jack --yes)
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "Jack cannot reject to disable validator $address, because Jack already rejected to disable validator"
result=$(dcld tx validator reject-disable-node --address="$validator_address" --from jack --yes)
response_does_not_contain "$result" "\"code\": 0"
echo "$result"

test_divider

echo "Get all proposed validators to disable. $address in the list"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator all-proposed-disable-nodes")
response_does_not_contain "$result" "\"approvals\":\[{\"address\":\"$alice_address\"" raw
response_does_not_contain "$result" "\"rejects\":\[{\"address\":\"$bob_address\"" raw
response_does_not_contain "$result" "\"address\": \"$validator_address\""
echo "$result"

test_divider

echo "Get all rejected validators to disable. $address in the list"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator all-rejected-disable-nodes")
check_response "$result" "\"approvals\":\[{\"address\":\"$alice_address\"" raw
check_response "$result" "\"rejects\":\[{\"address\":\"$bob_address\"" raw
check_response "$result" "\"address\": \"$jack_address\""
check_response "$result" "\"address\": \"$validator_address\""
echo "$result"

test_divider

echo "Get a rejected validator to disable $address"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator rejected-disable-node --address="$address"")
check_response "$result" "\"approvals\":\[{\"address\":\"$alice_address\"" raw
check_response "$result" "\"rejects\":\[{\"address\":\"$bob_address\"" raw
check_response "$result" "\"address\": \"$jack_address\""
check_response "$result" "\"address\": \"$validator_address\""
echo "$result"

echo "Alice proposes to disable validator $address"
result=$(dcld tx validator propose-disable-node --address="$validator_address" --from alice --yes)
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "Get a proposed validator to disable $address"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator proposed-disable-node --address="$address"")
check_response "$result" "\"approvals\":\[{\"address\":\"$alice_address\"" raw
check_response "$result" "\"address\": \"$validator_address\""
echo "$result"

test_divider

echo "Get a proposed validator to disable $validator_address"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator proposed-disable-node --address="$validator_address"")
check_response "$result" "\"approvals\":\[{\"address\":\"$alice_address\"" raw
check_response "$result" "\"address\": \"$validator_address\""
echo "$result"

test_divider

echo "Get a rejected account for $address is not found"
result=$(dcld query validator rejected-disable-node --address=$address)
check_response "$result" "Not Found"

test_divider

echo "Get a rejected account for $validator_address is not found"
result=$(dcld query validator rejected-disable-node --address=$validator_address)
check_response "$result" "Not Found"

test_divider

echo "Alice proposes to revoke NodeAdmin $address"
result=$(dcld tx auth propose-revoke-account --address="$address" --from alice --yes)
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "Bob approves to revoke NodeAdmin $address"
result=$(dcld tx auth approve-revoke-account --address="$address" --from bob --yes)
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "Node Admin cannot enable validator"
result=$(dcld tx validator enable-node --from "$account" --yes 2>&1 || true)
check_response "$result" "key not found" raw

test_divider

echo "PASSED"

cleanup
