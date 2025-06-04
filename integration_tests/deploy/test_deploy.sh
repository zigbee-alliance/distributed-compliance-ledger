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

# TODO
# - firewall check: install ufw (might be a part of dockerfile) and configure to allow ledger connections

DCL_USER="dcl"
DCL_USER_HOME="/var/lib/dcl"
DCL_DIR="$DCL_USER_HOME/.dcl"
GENESIS_FILE="$DCL_DIR/config/genesis.json"

GVN_NAME="test_deploy_gvn"
VN_NAME="test_deploy_vn"

CHAIN_ID=dcl_test_deploy

cleanup() {
    make test_deploy_env_clean
}
trap cleanup EXIT


function docker_exec {
    docker exec -u "$DCL_USER" "$@"
}


test_divider

echo "PROVISIONING"
GOPATH=${GOPATH:-${HOME}/go}
GOBIN=${GOBIN:-${GOPATH}/bin}

mkdir -p "$GOBIN"

docker build -f integration_tests/deploy/Dockerfile-build -t dcl-deploy-build .
docker container create --name dcl-deploy-build-inst dcl-deploy-build
docker cp dcl-deploy-build-inst:/go/bin/dcld "$GOBIN"/
docker cp dcl-deploy-build-inst:/go/bin/cosmovisor "$GOBIN"/
docker rm dcl-deploy-build-inst

make test_deploy_env_build

GVN_IP="$(docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $GVN_NAME)"
VN_IP="$(docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $VN_NAME)"

test_divider

echo "CLUSTER NODES PREPARATION"
for node in "$GVN_NAME" "$VN_NAME"; do
    echo "$node: install cosmovisor"
    docker cp "$GOBIN"/cosmovisor "$node":/usr/bin

    # TODO firewall routine (requires ufw installed)

    echo "$node: upload release artifacts"
    docker cp deployment/preupgrade.sh "$node":"$DCL_USER_HOME"
    docker cp deployment/dcld_manager.sh "$node":"$DCL_USER_HOME"
    docker cp deployment/cosmovisor.service "$node":"$DCL_USER_HOME"
    docker cp deployment/cosmovisor.conf "$node":"$DCL_USER_HOME"
    docker cp "$GOBIN"/dcld "$node":"$DCL_USER_HOME"
    docker cp deployment/scripts/run_dcl_node "$node":"$DCL_USER_HOME"
    docker cp deployment/scripts/test_peers_conn "$node":"$DCL_USER_HOME"

    docker_exec "$node" ./dcld version

    echo "$node: init"
    docker_exec "$node" bash -c "./dcld init $node --chain-id $CHAIN_ID 2>init.stderr"

    echo "$node: Create NodeAdmin and Trustee keys"
    docker_exec "$node" ./dcld config keyring-backend test
    docker_exec "$node" ./dcld keys add "${node}_admin"
    docker_exec "$node" ./dcld keys add "${node}_tr"
done

echo "Generating persistent peers list"
GVN_ID="$(docker_exec "$GVN_NAME" grep -o 'node_id.*$' init.stderr  | awk -F'"' '{print $3}')"
VN_ID="$(docker_exec "$VN_NAME" grep -o 'node_id.*$' init.stderr  | awk -F'"' '{print $3}')"

PERSISTENT_PEERS="$GVN_ID@$GVN_IP:26656,$VN_ID@$VN_IP:26656"

for node in "$GVN_NAME" "$VN_NAME"; do
    docker_exec "$node" bash -c "echo '$PERSISTENT_PEERS' >persistent_peers.txt"
done


test_divider

echo "GENESIS NODE RUN"
docker_exec -e KEYRING_BACKEND=test "$GVN_NAME"  \
    ./run_dcl_node -t genesis -c "$CHAIN_ID" \
    --gen-key-name "${GVN_NAME}_admin" --gen-key-name-trustee "${GVN_NAME}_tr" \
    "$GVN_NAME"
docker_exec "$GVN_NAME" systemctl status cosmovisor

# we should be fine with shell limits here
genesis_data="$(docker_exec "$GVN_NAME" cat "$GENESIS_FILE")"

test_divider

echo "VALIDATOR NODE RUN"

echo "VN: Set genesis data"
docker_exec "$VN_NAME" bash -c "echo '$genesis_data' >genesis.json"

echo "VN: Verify connections to other nodes"
docker_exec "$VN_NAME" ./test_peers_conn

echo "VN: run"
docker_exec -e KEYRING_BACKEND=test "$VN_NAME" ./run_dcl_node -c "$CHAIN_ID" "$VN_NAME"
docker_exec "$VN_NAME" systemctl status cosmovisor

test_divider

echo "VALIDATOR NODE REGISTRATION"

echo "$GVN_NAME: approve \"$VN_NAME\" to validator set"
vn_addr=$(docker_exec "$VN_NAME" ./dcld tendermint show-address)
vn_pubkey=$(docker_exec "$VN_NAME" ./dcld tendermint show-validator)
vn_admin_name="${VN_NAME}_admin"
vn_admin_addr="$(docker_exec "$VN_NAME" ./dcld keys show "$vn_admin_name" -a)"
vn_admin_pubkey="$(docker_exec "$VN_NAME" ./dcld keys show "$vn_admin_name" -p)"

# ensure that the genesis node is ready
wait_for_height 2 15 normal "tcp://$GVN_IP:26657"

result="$(docker_exec "$GVN_NAME" ./dcld tx auth propose-add-account --address "$vn_admin_addr" --pubkey "$vn_admin_pubkey" --roles="NodeAdmin" --from "${GVN_NAME}_tr" --yes)"
result=$(get_txn_result "$result")
#dcld tx auth approve-add-account --address="$vn_admin_addr" --from alice --yes

echo "$GVN_NAME: Add Node \"$VN_NAME\" to validator set"

# ensure that the validator node is ready
wait_for_height 4 30 normal "tcp://$VN_IP:26657"
result="$(docker_exec "$VN_NAME" ./dcld tx validator add-node --pubkey="$vn_pubkey" --moniker="$VN_NAME" --from="$vn_admin_name" --yes)"
result=$(get_txn_result "$result")

sleep 10
echo "Check node \"$VN_NAME\" is in the validator set"
result=$(docker_exec "$GVN_NAME" ./dcld query validator all-nodes)
check_response "$result" "\"moniker\": \"$VN_NAME\""
check_response "$result" "\"pubKey\":$vn_pubkey" raw

echo "PASSED"
