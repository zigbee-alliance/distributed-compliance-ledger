#!/bin/bash
set -e
source integration_tests/cli/common.sh

echo "Get key info for Jack"
result=$(zblcli keys show jack)
check_response "$result" "\"name\": \"jack\""
echo "$result"

echo "Get account info for Jack"
result=$(zblcli query account $(zblcli keys show jack -a))
check_response "$result" "\"account_number\":"
echo "$result"

echo "Assign Vendor role to Jack"
result=$(echo "test1234" | zblcli tx authz assign-role $(zblcli keys show jack -a) "Vendor" --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

sleep 5

vid=$RANDOM
pid=$RANDOM
name="Device #1"
echo "Add Model with VID: $vid PID: $pid"
result=$(echo "test1234" | zblcli tx modelinfo add-model $vid $pid "$name" "Device Description" "SKU12FS" "1.0" "2.0" true --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

sleep 5

echo "Get Model with VID: $vid PID: $pid"
result=$(zblcli query modelinfo model $vid $pid)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"name\": \"$name\""
echo "$result"

echo "Get all model infos"
result=$(zblcli query modelinfo all-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
echo "$result"

echo "Get all vendors"
result=$(zblcli query modelinfo vendors)
check_response "$result" "\"vid\": $vid"
echo "$result"

echo "Get Vendor Models with VID: ${vid}"
result=$(zblcli query modelinfo vendor-models $vid)
check_response "$result" "\"pid\": $pid"
echo "$result"

echo "Update Model with VID: ${vid} PID: ${pid}"
description="New Device Description"
result=$(echo "test1234" | zblcli tx modelinfo update-model $vid $pid true --from jack --yes --description "$description")
check_response "$result" "\"success\": true"
echo "$result"

sleep 5

echo "Get Model with VID: ${vid} PID: ${pid}"
result=$(zblcli query modelinfo model $vid $pid)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"description\": \"$description\""
echo "$result"
