#!/bin/bash
set -e
source integration_tests/cli/common.sh

echo "Get key info for Jack"
result=$(zblcli keys show jack)
check_response "$result" "\"name\": \"jack\""
echo "$result"

echo "Assign Vendor role to Jack"
result=$(echo "test1234" | zblcli tx authz assign-role $(zblcli keys show jack -a) "Vendor" --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

vid1=$RANDOM
pid1=$RANDOM
echo "Jack adds Model with VID: $vid1 PID: $pid1. Using default Broadcast Mode: block"
result=$(echo "test1234" | zblcli tx modelinfo add-model $vid1 $pid1 "Device #1" "Device Description" "SKU12FS" "1.0" "2.0" true --from tony --yes)
check_response "$result" "\"gas_used\""
check_response "$result" "\"txhash\""
check_response "$result" "\"raw_log\""
check_response "$result" "\"height\""
response_does_not_contain "$result" "\"height\": \"0\""
echo "$result"

vid2=$RANDOM
pid2=$RANDOM
echo "Jack adds Model with VID: $vid2 PID: $pid2. Using async Broadcast Mode"
result=$(echo "test1234" | zblcli tx modelinfo add-model $vid2 $pid2 "Device #2" "Device Description" "SKU12FS" "1.0" "2.0" true --from tony --yes --broadcast-mode "async")
check_response "$result" "\"txhash\""
check_response "$result" "\"height\": \"0\""
response_does_not_contain "$result" "\"gas_used\""
response_does_not_contain "$result" "\"raw_log\""
echo "$result"

sleep 6

vid3=$RANDOM
pid3=$RANDOM
echo "Jack adds Model with VID: $vid3 PID: $pid3. Using sync Broadcast Mode"
result=$(echo "test1234" | zblcli tx modelinfo add-model $vid3 $pid3 "Device #2" "Device Description" "SKU12FS" "1.0" "2.0" true --from tony --yes --broadcast-mode "sync")
check_response "$result" "\"txhash\""
check_response "$result" "\"raw_log\""
check_response "$result" "\"height\": \"0\""
response_does_not_contain "$result" "\"gas_used\""
echo "$result"

sleep 6

echo "Get Model with VID: $vid1 PID: $pid1"
result=$(zblcli query modelinfo model $vid1 $pid1 --prev-height)
check_response "$result" "\"vid\": $vid1"
check_response "$result" "\"pid\": $pid1"
echo "$result"

echo "Get Model with VID: $vid2 PID: $pid2"
result=$(zblcli query modelinfo model $vid2 $pid2 --prev-height)
check_response "$result" "\"vid\": $vid2"
check_response "$result" "\"pid\": $pid2"
echo "$result"

echo "Get Model with VID: $vid3 PID: $pid3"
result=$(zblcli query modelinfo model $vid3 $pid3 --prev-height)
check_response "$result" "\"vid\": $vid3"
check_response "$result" "\"pid\": $pid3"
echo "$result"