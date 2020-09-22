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

set -e
source integration_tests/cli_utils/common.sh

# Preparation of Actors

echo "Create Vendor account"
create_new_account vendor_account "Vendor"

# Body

vid1=$RANDOM
pid1=$RANDOM
echo "Jack adds Model with VID: $vid1 PID: $pid1. Using default Broadcast Mode: block"
result=$(echo "test1234" | dclcli tx modelinfo add-model --vid=$vid1 --pid=$pid1 --name="Device #1" --description="Device Description" --sku="SKU12FS" --firmware-version="1.0" --hardware-version="2.0" --tis-or-trp-testing-completed=true --from=$vendor_account --yes)
check_response "$result" "\"gas_used\""
check_response "$result" "\"txhash\""
check_response "$result" "\"raw_log\""
check_response "$result" "\"height\""
response_does_not_contain "$result" "\"height\": \"0\""
echo "$result"

vid2=$RANDOM
pid2=$RANDOM
echo "Jack adds Model with VID: $vid2 PID: $pid2. Using async Broadcast Mode"
result=$(echo "test1234" | dclcli tx modelinfo add-model --vid=$vid2 --pid=$pid2 --name="Device #2" --description="Device Description" --sku="SKU12FS" --firmware-version="1.0" --hardware-version="2.0" --tis-or-trp-testing-completed=true --from=$vendor_account --yes --broadcast-mode "async")
check_response "$result" "\"txhash\""
check_response "$result" "\"height\": \"0\""
response_does_not_contain "$result" "\"gas_used\""
response_does_not_contain "$result" "\"raw_log\""
echo "$result"

sleep 6

vid3=$RANDOM
pid3=$RANDOM
echo "Jack adds Model with VID: $vid3 PID: $pid3. Using sync Broadcast Mode"
result=$(echo "test1234" | dclcli tx modelinfo add-model --vid=$vid3 --pid=$pid3 --name="Device #2" --description="Device Description" --sku="SKU12FS" --firmware-version="1.0" --hardware-version="2.0" --tis-or-trp-testing-completed=true --from=$vendor_account --yes --broadcast-mode "sync")
check_response "$result" "\"txhash\""
check_response "$result" "\"raw_log\""
check_response "$result" "\"height\": \"0\""
response_does_not_contain "$result" "\"gas_used\""
echo "$result"

sleep 6

echo "Get Model with VID: $vid1 PID: $pid1"
result=$(dclcli query modelinfo model --vid=$vid1 --pid=$pid1 --prev-height)
check_response "$result" "\"vid\": $vid1"
check_response "$result" "\"pid\": $pid1"
echo "$result"

echo "Get Model with VID: $vid2 PID: $pid2"
result=$(dclcli query modelinfo model --vid=$vid2 --pid=$pid2 --prev-height)
check_response "$result" "\"vid\": $vid2"
check_response "$result" "\"pid\": $pid2"
echo "$result"

echo "Get Model with VID: $vid3 PID: $pid3"
result=$(dclcli query modelinfo model --vid=$vid3 --pid=$pid3 --prev-height)
check_response "$result" "\"vid\": $vid3"
check_response "$result" "\"pid\": $pid3"
echo "$result"
