#!/bin/bash
set -e
source integration_tests/cli/common.sh

account=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 6 | head -n 1)
container="validator-demo"
node="node-demo"
chain_id="zblchain"
ip="192.167.10.6"
node0="tcp://192.167.10.2:26657"
passphrase="test1234"

if docker container ls  | grep -q $container; then
  if docker container inspect $container | grep -q '"Status": "running"'; then
    echo "Stopping container"
    docker container kill $container
  fi

  echo "Removing container"
  docker container rm "$container"
fi

docker run -d --name $container --ip $ip -p "26664-26665:26656-26657" --network zb-ledger_localnet -i zbledger

echo "Generate keys for $account"
docker exec $container /bin/sh -c "echo $passphrase | zblcli keys add $account"
address=$(docker exec $container zblcli keys show $account -a)
pubkey=$(docker exec $container zblcli keys show $account -p)

echo "Create account for $account and Assign NodeAdmin role"
echo $passphrase | zblcli tx auth propose-add-account --address="$address" --pubkey="$pubkey" --roles="NodeAdmin" --from jack --yes
echo $passphrase | zblcli tx auth approve-add-account --address="$address" --from alice --yes

echo "$account Preapare Node configuration files"
docker exec $container zbld init $node --chain-id $chain_id
docker cp ./localnet/node0/config/genesis.json $container:/root/.zbld/config
peers=$(cat localnet/node0/config/config.toml | grep -o -E "persistent_peers = \".*\"")
docker exec $container sed -i "s/persistent_peers = \"\"/$peers/g" /root/.zbld/config/config.toml
docker exec $container sed -i 's/laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/g' /root/.zbld/config/config.toml

echo "$account Configure CLI"
docker exec $container /bin/sh -c "
  zblcli config chain-id zblchain &&
  zblcli config output json &&
  zblcli config indent true &&
  zblcli config trust-node false &&
  zblcli config node $node0"

echo "$account Add Node \"$node\" to validator set"
vaddress=$(docker exec $container zbld tendermint show-address)
vpubkey=$(docker exec $container zbld tendermint show-validator)
result=$(docker exec $container /bin/sh -c "echo test1234 | zblcli tx validator add-node --validator-address=$vaddress --validator-pubkey=$vpubkey --name=$node --from=$account --yes")
check_response "$result" "\"success\": true"
echo "$result"

echo "$account Start Node \"$node\""
docker exec -d $container zbld start
sleep 10

echo "Check node \"$node\" is in the validator set"
result=$(zblcli query validator all-nodes)
check_response "$result" "\"name\": \"$node\""
check_response "$result" "\"validator_address\": \"$vaddress\""
check_response "$result" "\"validator_pubkey\": \"$vpubkey\""
echo "$result"

echo "Connect CLI to node \"$node\" and check status"
zblcli config node "tcp://$ip:26657"
result=$(zblcli status)
check_response "$result" "\"moniker\": \"$node\""
echo "$result"

echo "Sent transactions using node \"$node\""
create_new_account vendor_account "Vendor"

echo "Publish Model"
vid=$RANDOM
pid=$RANDOM
name="Device #1"
echo "Add Model with VID: $vid PID: $pid"
result=$(echo "test1234" | zblcli tx modelinfo add-model --vid=$vid --pid=$pid --name="$name" --description="Device Description" --sku="SKU12FS" --firmware-version="1.0" --hardware-version="2.0" --tis-or-trp-testing-completed=true --from "$vendor_account" --yes)
check_response "$result" "\"success\": true"
echo "$result"

sleep 5

echo "Connect CLI to node \"node0\""
zblcli config node $node0
result=$(zblcli status)
check_response "$result" "\"moniker\": \"node0\""
echo "$result"

echo "Query Model using node0 node"
echo "Get Model with VID: $vid PID: $pid"
result=$(zblcli query modelinfo model --vid=$vid --pid=$pid)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"name\": \"$name\""
echo "$result"

docker rm -f $container
