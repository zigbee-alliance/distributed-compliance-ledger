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

random_string user
echo "$user generates keys"
result=$(echo 'test1234' | dclcli keys add $user)

test_divider

echo "Get key info for $user"
result=$(dclcli keys show $user)
check_response "$result" "\"name\": \"$user\""

test_divider

user_address=$(dclcli keys show $user -a)
user_pubkey=$(dclcli keys show $user -p)
vid=$RANDOM
pid=$RANDOM

echo "Jack proposes account for $user"
result=$(echo $passphrase | dclcli tx auth propose-add-account --address="$user_address" --pubkey="$user_pubkey" --roles="Vendor" --vid="$vid" --from jack --yes)
check_response "$result" "\"success\": true"

test_divider

echo "Get all active accounts. No $user account in the list because not enough approvals received"
result=$(dclcli query auth all-accounts)
response_does_not_contain "$result" "\"address\": \"$user_address\""

test_divider

echo "Get all proposed accounts. $user account in the list"
result=$(dclcli query auth all-proposed-accounts)
check_response "$result" "\"address\": \"$user_address\""

test_divider

echo "Get all proposed accounts to revoke. No $user account in the list"
result=$(dclcli query auth all-proposed-accounts-to-revoke)
response_does_not_contain "$result" "\"address\": \"$user_address\""

test_divider

echo "Alice approves account for \"$user\""
result=$(echo $passphrase | dclcli tx auth approve-add-account --address="$user_address" --from alice --yes)
check_response "$result" "\"success\": true"

test_divider

echo "Get all accounts. $user account in the list because enough approvals received"
result=$(dclcli query auth all-accounts)
check_response "$result" "\"address\": \"$user_address\""

test_divider

echo "Get all proposed accounts. No $user account in the list anymore"
result=$(dclcli query auth all-proposed-accounts)
response_does_not_contain "$result" "\"address\": \"$user_address\""

test_divider

echo "Get all proposed accounts to revoke. No $user account in the list"
result=$(dclcli query auth all-proposed-accounts-to-revoke)
response_does_not_contain "$result" "\"address\": \"$user_address\""

test_divider

echo "Get $user account"
result=$(dclcli query auth account --address=$user_address)
check_response_and_report "$result" "\"address\": \"$user_address\""

test_divider

productName="Device #1"
echo "$user adds Model with VID: $vid PID: $pid"
result=$(echo "test1234" | dclcli tx model add-model --vid=$vid --pid=$pid --productName="$productName" --productLabel="Device Description"   --commissioningCustomFlow=0 --deviceTypeID=12 --partNumber=12 --from=$user_address --yes)
check_response_and_report "$result" "\"success\": true"

test_divider

vidPlusOne=$((vid+1))
echo "$user adds Model with a VID: $vidPlusOne PID: $pid, This fails with Permission denied as the VID is not associated with this vendor account."
result=$(echo "test1234" | dclcli tx model add-model --vid=$vidPlusOne --pid=$pid --productName="$productName" --productLabel="Device Description"   --commissioningCustomFlow=0 --deviceTypeID=12 --partNumber=12 --from=$user_address --yes 2>&1) || true
check_response_and_report "$result" "transaction should be signed by an vendor account containing the vendorId $vidPlusOne"

test_divider

echo "$user updates Model with VID: $vid PID: $pid"
result=$(echo "test1234" | dclcli tx model update-model --vid=$vid --pid=$pid --productName="$productName" --productLabel="Device Description" --partNumber=12 --from=$user_address --yes 2>&1) || true
check_response_and_report "$result" "\"success\": true"

test_divider

echo "Get Model with VID: $vid PID: $pid"
result=$(dclcli query model get-model --vid=$vid --pid=$pid)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"productName\": \"$productName\""

test_divider

echo "Alice proposes to revoke account for $user"
result=$(echo $passphrase | dclcli tx auth propose-revoke-account --address="$user_address" --from alice --yes)
check_response "$result" "\"success\": true"


test_divider

echo "Get all accounts. $user account still in the list because not enought approvals to revoke received"
result=$(dclcli query auth all-accounts)
check_response "$result" "\"address\": \"$user_address\""


test_divider

echo "Get all proposed accounts. No $user account in the list"
result=$(dclcli query auth all-proposed-accounts)
response_does_not_contain "$result" "\"address\": \"$user_address\""


test_divider

echo "Get all proposed accounts to revoke. $user account in the list"
result=$(dclcli query auth all-proposed-accounts-to-revoke)
check_response "$result" "\"address\": \"$user_address\""

test_divider

echo "Bob approves to revoke account for $user"
result=$(echo $passphrase | dclcli tx auth approve-revoke-account --address="$user_address" --from bob --yes)
check_response "$result" "\"success\": true"

test_divider

echo "Get all accounts. No $user account in the list anymore because enought approvals to revoke received"
result=$(dclcli query auth all-accounts)
response_does_not_contain "$result" "\"address\": \"$user_address\""

test_divider

echo "Get all proposed accounts. No $user account in the list"
result=$(dclcli query auth all-proposed-accounts)
response_does_not_contain "$result" "\"address\": \"$user_address\""

test_divider

echo "Get all proposed accounts to revoke. No $user account in the list anymore"
result=$(dclcli query auth all-proposed-accounts-to-revoke)
response_does_not_contain "$result" "\"address\": \"$user_address\""

test_divider

echo "Get $user account"
result=$(dclcli query auth account --address=$user_address 2>&1) || true
check_response_and_report "$result" "No account associated with the address"

test_divider

pid=$RANDOM
productName="Device #2"
echo "$user adds Model with VID: $vid PID: $pid"
result=$(echo "test1234" | dclcli tx model add-model --vid=$vid --pid=$pid --productName="$productName" --productLabel="Device Description" --commissioningCustomFlow=0 --deviceTypeID=12 --partNumber=12 --from=$user_address --yes 2>&1) || true
check_response_and_report "$result" "No account associated with the address"