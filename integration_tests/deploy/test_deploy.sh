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
TEST_NODE="test_deploy_node"

random_string account

cleanup() {
    make test_deploy_env_clean
    make localnet_clean
}
trap cleanup EXIT

test_divider

echo "Prepare the environment"

GOPATH=${GOPATH:-${HOME}/go}
GOBIN=${GOBIN:-${GOPATH}/bin}

mkdir -p "$GOBIN"

docker build -f integration_tests/deploy/Dockerfile-build -t dcl-deploy-build .
docker container create --name dcl-deploy-build-inst dcl-deploy-build
docker cp dcl-deploy-build-inst:/go/bin/dcld "$GOBIN"/
docker cp dcl-deploy-build-inst:/go/bin/cosmovisor "$GOBIN"/
docker rm dcl-deploy-build-inst

make localnet_rebuild localnet_start
make test_deploy_env_build
# ensure that the pool is ready
wait_for_height 2 20

docker cp "$GOBIN"/cosmovisor "$TEST_NODE":/usr/bin
docker cp "$LOCALNET_DIR"/genesis.json "$TEST_NODE":"$DCL_USER_HOME"
docker cp "$LOCALNET_DIR"/persistent_peers.txt "$TEST_NODE":"$DCL_USER_HOME"
docker cp "$GOBIN"/dcld "$TEST_NODE":"$DCL_USER_HOME"
docker cp deployment/scripts/run_dcl_node "$TEST_NODE":"$DCL_USER_HOME"
docker cp deployment/cosmovisor.service "$TEST_NODE":"$DCL_USER_HOME"

echo "Configure CLI"
docker exec -u "$DCL_USER" "$TEST_NODE" /bin/sh -c "
  ./dcld config chain-id $chain_id &&
  ./dcld config output json &&
  ./dcld config keyring-backend test &&
  ./dcld config broadcast-mode block"

echo "Configure and start new node"
docker exec -u "$DCL_USER" "$TEST_NODE" ./dcld init $TEST_NODE --chain-id  $chain_id
docker exec -u "$DCL_USER" "$TEST_NODE" ./run_dcl_node -u $DCL_USER -c $chain_id $TEST_NODE
docker exec "$TEST_NODE" systemctl status cosmovisor
vaddress=$(docker exec -u "$DCL_USER" "$TEST_NODE" ./dcld tendermint show-address)
vpubkey=$(docker exec -u "$DCL_USER" "$TEST_NODE" ./dcld tendermint show-validator)

echo "Create and register new NodeAdmin account"
docker exec -u "$DCL_USER" "$TEST_NODE" ./dcld keys add "$account"
address="$(docker exec -u "$DCL_USER" "$TEST_NODE" ./dcld keys show $account -a)"
pubkey="$(docker exec -u "$DCL_USER" "$TEST_NODE" ./dcld keys show $account -p)"
dcld tx auth propose-add-account --address="$address" --pubkey="$pubkey" --roles="NodeAdmin" --from jack --yes
dcld tx auth approve-add-account --address="$address" --from alice --yes

echo "$account Add Node \"$TEST_NODE\" to validator set"
docker exec -u "$DCL_USER" "$TEST_NODE" ./dcld tx validator add-node --pubkey="$vpubkey" --moniker="$TEST_NODE" --from="$account" --yes

echo "Check node \"$TEST_NODE\" is in the validator set"
result=$(dcld query validator all-nodes)
check_response "$result" "\"moniker\": \"$TEST_NODE\""
check_response "$result" "\"pubKey\":$vpubkey" raw

echo "PASSED"
