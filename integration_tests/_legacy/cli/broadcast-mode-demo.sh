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

# FIXME issue 99: enable once implemented
exit 0

# Preparation of Actors
vid1=$RANDOM
pid1=$RANDOM
productName="Broadcast-Demo-Product"
vendor_account=vendor_account_$vid1

echo "Create Vendor account - $vendor_account"
create_new_vendor_account $vendor_account $vid1

test_divider

# Body

echo "Jack adds Model with VID: $vid1 PID: $pid1. Using default Broadcast Mode: block"
result=$(echo "test1234" | dcld tx model add-model --vid=$vid1 --pid=$pid1 --productName="$productName" --productLabel="Device Description"   --commissioningCustomFlow=0 --deviceTypeID=12 --partNumber=12  --from=$vendor_account --yes)
check_response "$result" "\"gas_used\""
check_response "$result" "\"txhash\""
check_response "$result" "\"raw_log\""
check_response "$result" "\"height\""
response_does_not_contain "$result" "\"height\": \"0\""

test_divider

pid2=$RANDOM
echo "Jack adds Model with VID: $vid1 PID: $pid2. Using async Broadcast Mode"
result=$(echo "test1234" | dcld tx model add-model --vid=$vid1 --pid=$pid2 --productName="$productName" --productLabel="Device Description"   --commissioningCustomFlow=0 --deviceTypeID=12 --partNumber=12  --from=$vendor_account --yes --broadcast-mode "async")
check_response "$result" "\"txhash\""
check_response "$result" "\"height\": \"0\""
response_does_not_contain "$result" "\"gas_used\""
response_does_not_contain "$result" "\"raw_log\""


test_divider

sleep 6
pid3=$RANDOM
echo "Jack adds Model with VID: $vid1 PID: $pid3. Using sync Broadcast Mode"
result=$(echo "test1234" | dcld tx model add-model --vid=$vid1 --pid=$pid3 --productName="$productName" --productLabel="Device Description"   --commissioningCustomFlow=0 --deviceTypeID=12 --partNumber=12  --from=$vendor_account --yes --broadcast-mode "sync")
check_response "$result" "\"txhash\""
check_response "$result" "\"raw_log\""
check_response "$result" "\"height\": \"0\""
response_does_not_contain "$result" "\"gas_used\""


test_divider

sleep 6
echo "Get Model with VID: $vid1 PID: $pid1"
result=$(dcld query model get-model --vid=$vid1 --pid=$pid1 )
check_response "$result" "\"vid\": $vid1"
check_response "$result" "\"pid\": $pid1"


echo "Get Model with VID: $vid1 PID: $pid2"
result=$(dcld query model get-model --vid=$vid1 --pid=$pid2 )
check_response "$result" "\"vid\": $vid1"
check_response "$result" "\"pid\": $pid2"


echo "Get Model with VID: $vid1 PID: $pid3"
result=$(dcld query model get-model --vid=$vid1 --pid=$pid3 )
check_response "$result" "\"vid\": $vid1"
check_response "$result" "\"pid\": $pid3"

