#!/bin/bash

set -euo pipefail
source integration_tests/cli/common.sh

binary_version="v1.4.4"
local_build_bin="/build/dcld" # Path to locally built dcld v1.4.5
node_count=4

wget -O dcld_old "https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/$binary_version/dcld"
chmod ugo+x dcld_old

DCLD_BIN_OLD="./dcld_old"

check_pool_accepts_tx() {
  # Generate random test data for transaction
  random_string vendor_name
  random_string company_legal_name
  random_string company_preferred_name
  vendor_landing_page_url="https://example.com"
  vid=$((RANDOM % 10000 + 100))
  passphrase="test1234"

  tx_result=$(echo $passphrase | $DCLD_BIN_OLD tx vendorinfo add-vendor \
    --vid=$vid \
    --vendorName="$vendor_name" \
    --companyLegalName="$company_legal_name" \
    --companyPreferredName="$company_preferred_name" \
    --vendorLandingPageURL="$vendor_landing_page_url" \
    --from=jack --yes || true)
  echo "$tx_result"
  if [[ "$tx_result" == *"code\": 0"* ]]; then
    echo "Pool accepts transactions"
    return 0
  else
    echo "Pool does NOT accept transactions"
    return 1
  fi
}




# Add NodeAdmin profile and approve with trustees
random_string nodeadmin_account
passphrase="test1234"
echo $passphrase | $DCLD_BIN_OLD keys add $nodeadmin_account
nodeadmin_address=$(echo $passphrase | $DCLD_BIN_OLD keys show $nodeadmin_account -a)
nodeadmin_pubkey=$(echo $passphrase | $DCLD_BIN_OLD keys show $nodeadmin_account -p)

trustees=("jack" "alice" "bob")
echo $passphrase | $DCLD_BIN_OLD tx auth propose-add-account --address="$nodeadmin_address" --pubkey="$nodeadmin_pubkey" --roles="NodeAdmin" --from jack --yes
for trustee in "${trustees[@]}"; do
  echo $passphrase | $DCLD_BIN_OLD tx auth approve-add-account --address="$nodeadmin_address" --from $trustee --yes
done

# Query account to check validity
result=$($DCLD_BIN_OLD query auth account --address="$nodeadmin_address")
check_response "$result" "$nodeadmin_address"

# Try to add validator with wrong pubkey (NodeAdmin pubkey instead of validator pubkey)
random_string bad_validator_account
docker exec validator-1 /bin/sh -c "(echo $passphrase; echo $passphrase) | ./dcld keys add $bad_validator_account"
bad_validator_address=$(docker exec validator-1 /bin/sh -c "echo $passphrase | ./dcld keys show $bad_validator_account -a")
bad_validator_pubkey=$(docker exec validator-1 /bin/sh -c "echo $passphrase | ./dcld keys show $bad_validator_account -p")

# NodeAdmin tries to add validator with its own pubkey (should break consensus)
result=$(docker exec validator-1 /bin/sh -c "echo $passphrase | ./dcld tx validator add-node --pubkey='$nodeadmin_pubkey' --moniker='bad-node' --from='$nodeadmin_account' --yes")
echo "$result"

# Check that pool stopped accepting tx (simulate by sending tx and expecting failure)
if check_pool_accepts_tx; then
  echo "Pool still accepts transactions, test failed"
  exit 1
else
  echo "Pool stopped accepting transactions as expected"
fi

# Check logs for "CONSENSUS FAILURE"
for i in $(seq 0 $((node_count-1))); do
  name="node$i"
  log=$(docker logs $name 2>&1 | grep "CONSENSUS FAILURE" || true)
  if [[ -z "$log" ]]; then
    echo "CONSENSUS FAILURE not found in $name logs"
    exit 1
  fi
  echo "$name log: $log"
done

echo "Building new dcld binary v1.4.5 from source"
make build
cp ./build/dcld /build/dcld

# Upgrade all validator nodes to 1.4.5 (local build)
for i in $(seq 0 $((node_count-1))); do
  name="node$i"
  docker cp $local_build_bin $name:/var/lib/dcl/dcld
  docker exec $name pkill dcld || true
  docker exec -d $name cosmovisor run start
done

sleep 10

# Check that pool accepts transactions again
if check_pool_accepts_tx "$local_build_bin"; then
  echo "Pool accepts transactions after upgrade"
else
  echo "Pool does NOT accept transactions after upgrade, test failed"
  exit 1
fi

echo "Consensus failure patch test passed"