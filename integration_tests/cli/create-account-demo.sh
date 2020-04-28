#!/bin/bash
set -e
source integration_tests/cli/common.sh

echo "Get key info for Jack"
result=$(zblcli keys show jack)
check_response "$result" "\"name\": \"jack\""
echo "$result"

echo "Jack assigns Trustee role to Jack"
result=$(echo "test1234" | zblcli tx authz assign-role $(zblcli keys show jack -a) "Trustee" --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Tony generates keys"
result=$(echo 'test1234' | zblcli keys add tony)
echo "$result"

echo "Get key info for Tony"
result=$(zblcli keys show tony)
check_response "$result" "\"name\": \"tony\""
echo "$result"

tony_address=$(zblcli keys show tony -a)
tony_pubkey=$(zblcli keys show tony -p)

echo "Jack creates account for Tony"
result=$(echo "test1234" | zblcli tx authnext create-account "$tony_address" "$tony_pubkey" --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get all accouts"
result=$(zblcli query authnext account-headers)
check_response "$result" "\"address\": \"$tony_address\""
check_response "$result" "\"public_key\": \"$tony_pubkey\""
check_response "$result" "\"roles\": []"
echo "$result"

echo "Jack assigns Vendor role to Tony"
result=$(echo "test1234" | zblcli tx authz assign-role $tony_address "Vendor" --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get account roles for Tony"
result=$(zblcli query authz account-roles $tony_address)
check_response "$result" "\"Vendor\""
echo "$result"

vid=$RANDOM
pid=$RANDOM
name="Device #1"

echo "Tony adds Model with VID: $vid PID: $pid"
result=$(echo "test1234" | zblcli tx modelinfo add-model $vid $pid "$name" "Device Description" "SKU12FS" "1.0" "2.0" true --from tony --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get Model with VID: $vid PID: $pid"
result=$(zblcli query modelinfo model $vid $pid)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"name\": \"$name\""
echo "$result"