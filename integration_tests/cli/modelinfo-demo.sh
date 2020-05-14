#!/bin/bash
set -e
source integration_tests/cli/common.sh

echo "Get key info for Jack"
result=$(zblcli keys show jack)
check_response "$result" "\"name\": \"jack\""
echo "$result"

echo "Get account info for Jack"
result=$(zblcli query auth account --address=$(zblcli keys show jack -a))
check_response "$result" "\"account_number\":"
echo "$result"

echo "Assign Vendor role to Jack"
result=$(echo "test1234" | zblcli tx auth assign-role --address=$(zblcli keys show jack -a) --role="Vendor" --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

vid=$RANDOM
pid=$RANDOM
name="Device #1"
echo "Add Model with VID: $vid PID: $pid"
result=$(echo "test1234" | zblcli tx modelinfo add-model --vid=$vid --pid=$pid --name="$name" --description="Device Description" --sku="SKU12FS" --firmware-version="1.0" --hardware-version="2.0" --tis-or-trp-testing-completed=true --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get Model with VID: $vid PID: $pid"
result=$(zblcli query modelinfo model --vid=$vid --pid=$pid)
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
result=$(zblcli query modelinfo vendor-models --vid=$vid)
check_response "$result" "\"pid\": $pid"
echo "$result"

echo "Update Model with VID: ${vid} PID: ${pid}"
description="New Device Description"
result=$(echo "test1234" | zblcli tx modelinfo update-model --vid=$vid --pid=$pid --tis-or-trp-testing-completed=true --from jack --yes --description "$description")
check_response "$result" "\"success\": true"
echo "$result"

echo "Get Model with VID: ${vid} PID: ${pid}"
result=$(zblcli query modelinfo model --vid=$vid --pid=$pid)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"description\": \"$description\""
echo "$result"
