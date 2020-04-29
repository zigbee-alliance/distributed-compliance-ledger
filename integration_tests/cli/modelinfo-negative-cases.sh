#!/bin/bash
set -e
source integration_tests/cli/common.sh

account=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 6 | head -n 1)
create_account_with_name $account

# Ledger side errors

vid=$RANDOM
pid=$RANDOM

echo "Add Model with VID: $vid PID: $pid: Not Trustee"
result=$(echo "test1234" | zblcli tx modelinfo add-model $vid $pid "Name 1" "Device Description" "SKU12FS" "1.0" "2.0" true --from $account --yes)
check_response "$result" "\"success\": false"
check_response "$result" "\"code\": 4"
echo "$result"

echo "Jack assigns Vendor role to $account"
result=$(echo "test1234" | zblcli tx authz assign-role $(zblcli keys show "$account" -a) "Vendor" --from jack --yes)

echo "Add Model with VID: $vid PID: $pid: Twice"
result=$(echo "test1234" | zblcli tx modelinfo add-model $vid $pid "Name 1" "Device Description" "SKU12FS" "1.0" "2.0" true --from $account --yes)
result=$(echo "test1234" | zblcli tx modelinfo add-model $vid $pid "Name 1" "Device Description" "SKU12FS" "1.0" "2.0" true --from $account --yes)
check_response "$result" "\"success\": false"
check_response "$result" "\"code\": 101"
echo "$result"

# CLI side errors

echo "Add Model with VID: $vid PID: $pid: Unknown account"
result=$(zblcli tx modelinfo add-model $vid $pid "Name 1" "Device Description" "SKU12FS" "1.0" "2.0" true --from "Unknown"  2>&1) || true
check_response "$result" "Key Unknown not found"

echo "Add model with invalid VID/PID"
for i in "0" "string"
do
  result=$(echo "test1234" | zblcli tx modelinfo add-model $i $pid "Name" "Device Description" "SKU12FS" "1.0" "2.0" true --from $address --yes 2>&1) || true
  check_response "$result" "Invalid VID"

  result=$(echo "test1234" | zblcli tx modelinfo add-model $vid $i "Name" "Device Description" "SKU12FS" "1.0" "2.0" true --from $address --yes 2>&1) || true
  check_response "$result" "Invalid PID"
done

echo "Add model with epmty name"
result=$(echo "test1234" | zblcli tx modelinfo add-model $vid $pid "" "Device Description" "SKU12FS" "1.0" "2.0" true --from $address --yes 2>&1) || true
check_response "$result" "Code: 6"
check_response "$result" "Invalid Name"

echo "Add model with epmty description"
result=$(echo "test1234" | zblcli tx modelinfo add-model $vid $pid "Name" "" "SKU12FS" "1.0" "2.0" true --from $address --yes 2>&1) || true
check_response "$result" "Code: 6"
check_response "$result" "Invalid Description"

echo "Add model with epmty SKU"
result=$(echo "test1234" | zblcli tx modelinfo add-model $vid $pid "Name" "Description" "" "1.0" "2.0" true --from $address --yes 2>&1) || true
check_response "$result" "Code: 6"
check_response "$result" "Invalid SKU"

echo "Add model with epmty Firmwere Version"
result=$(echo "test1234" | zblcli tx modelinfo add-model $vid $pid "Name" "Description" "SKU12FS" "" "2.0" true --from $address --yes 2>&1) || true
check_response "$result" "Code: 6"
check_response "$result" "Invalid FirmwareVersion"

echo "Add model with epmty Hardware Version"
result=$(echo "test1234" | zblcli tx modelinfo add-model $vid $pid "Name" "Description" "SKU12FS" "1.0" "" true --from $address --yes 2>&1) || true
check_response "$result" "Code: 6"
check_response "$result" "Invalid HardwareVersion"

echo "Add model with Invalid TIS flag"
result=$(echo "test1234" | zblcli tx modelinfo add-model $vid $pid "Name" "Description" "SKU12FS" "1.0" "2.0" "string" --from $address --yes 2>&1) || true
check_response "$result" "Code: 6"
check_response "$result" "Invalid Tis-or-trp-testing-completed"

echo "Add model without --from flag"
result=$(echo "test1234" | zblcli tx modelinfo add-model $vid $pid "Name" "Device Description" "SKU12FS" "1.0" "2.0" true --yes 2>&1) || true
check_response "$result" "Invalid Signer"

echo "Add model with not enought parameters"
result=$(echo "test1234" | zblcli tx modelinfo add-model $vid $pid "" "Device Description" "SKU12FS" "1.0" "2.0" --yes 2>&1) || true
check_response "$result" "ERROR: accepts 8 arg(s), received 7"