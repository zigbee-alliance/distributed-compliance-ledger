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

binary_version_old="v1.4.4"
binary_version_new="v1.4.5"
node_count=4

DCLD_BIN_OLD="/tmp/dcld_bins/dcld_$binary_version_old"
DCLD_BIN_NEW="/tmp/dcld_bins/dcld_$binary_version_new"

check_pool_accepts_tx() {
  # Generate random test data for transaction
  random_string vendor_name
  random_string company_legal_name
  random_string company_preferred_name
  vendor_landing_page_url="https://example.com"
  passphrase="test1234"

  tx_result=$(echo $passphrase | $DCLD_BIN_OLD tx vendorinfo update-vendor \
    --vid=$vid \
    --vendorName="$vendor_name" \
    --companyLegalName="$company_legal_name" \
    --companyPreferredName="$company_preferred_name" \
    --vendorLandingPageURL="$vendor_landing_page_url" \
    --from=$vendor_account --yes || true)
  result=$(get_txn_result "$tx_result")
  if $(_check_response "$result" "\"code\": 0" ); then
    return 0
  else
    return 1
  fi
}


test_divider
if ! check_pool_accepts_tx; then
  echo "FAIL: Pool does NOT accept transactions"
  exit 1
fi
if ! check_pool_accepts_tx; then
  echo "FAIL: Pool does NOT accept transactions"
  exit 1
fi

echo "Add NodeAdmin profile and approve with trustees"
random_string nodeadmin_account
passphrase="test1234"
echo $passphrase | $DCLD_BIN_OLD keys add $nodeadmin_account
nodeadmin_address=$(echo $passphrase | $DCLD_BIN_OLD keys show $nodeadmin_account -a)
nodeadmin_pubkey=$(echo $passphrase | $DCLD_BIN_OLD keys show $nodeadmin_account -p)

result=$(echo $passphrase | $DCLD_BIN_OLD tx auth propose-add-account --address="$nodeadmin_address" --pubkey="$nodeadmin_pubkey" --roles="NodeAdmin" --from $trustee_account_1 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

for trustee in $trustee_account_2 $trustee_account_3 $trustee_account_4 $trustee_account_5; do
  result=$(echo $passphrase | $DCLD_BIN_OLD tx auth approve-add-account --address="$nodeadmin_address" --from $trustee --yes)
  result=$(get_txn_result "$result")
  echo "$trustee approved new NodeAdmin $nodeadmin_address"
done

echo "Query account to check validity"
result=$($DCLD_BIN_OLD query auth account --address="$nodeadmin_address")
check_response "$result" "$nodeadmin_address"

test_divider

echo "Try to add validator with a wrong pubkey (NodeAdmin pubkey instead of validator pubkey)"
echo "It should break consensus"
result=$(echo $passphrase | $DCLD_BIN_OLD  tx validator add-node --pubkey="$nodeadmin_pubkey" --moniker="bad-node" --from="$nodeadmin_account" --yes)

test_divider

if check_pool_accepts_tx; then
  echo "FAIL: Pool still accepts transactions, test failed"
  exit 1
else
  echo "Pool stopped accepting transactions as expected"
fi

test_divider

echo "Check logs for CONSENSUS FAILURE"
for i in $(seq 0 $((node_count-1))); do
  name="node$i"
  log=$(docker logs $name 2>&1 | grep "CONSENSUS FAILURE" | grep "failed to apply block" || true)
  if [[ -z "$log" ]]; then
    echo "FAIL: CONSENSUS FAILURE not found in $name logs"
    exit 1
  fi
  echo "$name log: $log"
done

get_height broken_height

test_divider

echo "Rollback and upgrade all validator nodes to 1.4.5"
for i in $(seq 0 $((node_count-1))); do
  test_divider

  name="node$i"
  echo "Rollback for $name"

  docker cp $DCLD_BIN_NEW $name:/var/lib/dcl/.dcl/cosmovisor/upgrades/v1.4.4/bin/dcld
  docker stop $name
  result=$($DCLD_BIN_NEW rollback --hard --home ./.localnet/$name)
  echo "$result"
  docker start $name
done

wait_for_height $(expr $broken_height + 2) 300 outage-safe

test_divider

echo "Check that pool accepts transactions again"
if check_pool_accepts_tx "$DCLD_BIN_NEW"; then
  echo "Pool accepts transactions after upgrade"
else
  echo "FAIL: Pool does NOT accept transactions after upgrade, test failed"
  echo $($DCLD_BIN_NEW status)
  exit 1
fi

test_divider

echo "Rollback and upgrade the last node"

docker cp $DCLD_BIN_NEW $VALIDATOR_DEMO_CONTAINER_NAME:/var/lib/dcl/.dcl/cosmovisor/upgrades/v1.4.4/bin/dcld
docker exec $VALIDATOR_DEMO_CONTAINER_NAME pkill cosmovisor
docker exec $VALIDATOR_DEMO_CONTAINER_NAME dcld rollback --hard
docker exec -d $VALIDATOR_DEMO_CONTAINER_NAME cosmovisor run start

wait_for_height $(expr $broken_height + 3) 300 outage-safe "tcp://localhost:$node_client_port"

echo "Node started successfully after rollback and upgrade"

test_divider
echo "Consensus failure patch test passed"
