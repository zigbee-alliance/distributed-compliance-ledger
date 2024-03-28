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
cmd="(echo $passphrase; echo $passphrase) | dcld keys add $user"
result="$(bash -c "$cmd")"

test_divider

echo "Get key info for $user"
result=$(echo $passphrase | dcld keys show $user)
check_response "$result" "\"name\": \"$user\""

test_divider

user_address=$(echo $passphrase | dcld keys show $user -a)
user_pubkey=$(echo $passphrase | dcld keys show $user -p)

jack_address=$(echo $passphrase | dcld keys show jack -a)
alice_address=$(echo $passphrase | dcld keys show alice -a)
bob_address=$(echo $passphrase | dcld keys show bob -a)
anna_address=$(echo $passphrase | dcld keys show anna -a)

vid_in_hex_format=0xA13
pid_in_hex_format=0xA11
vid=2579
pid=2577

echo "Jack proposes account for $user"
result=$(echo $passphrase | dcld tx auth propose-add-account --info="Jack is proposing this account" --address="$user_address" --pubkey="$user_pubkey" --roles="Vendor" --vid="$vid_in_hex_format" --from jack --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Get all active accounts. $user account in the list because has enough approvals"
result=$(dcld query auth all-accounts)
check_response "$result" "\"address\": \"$user_address\""

test_divider

echo "Get an account for $user"
result=$(dcld query auth account --address=$user_address)
check_response "$result" "\"address\": \"$user_address\""
check_response_and_report "$result"  $jack_address "json"
check_response_and_report "$result"  '"info": "Jack is proposing this account"' "json"

test_divider

echo "Get an proposed account for $user is not found"
result=$(dcld query auth proposed-account --address=$user_address)
check_response "$result" "Not Found"

test_divider

echo "Get all proposed accounts. $user account is not in the list"
result=$(dcld query auth all-proposed-accounts)
check_response "$result" "\[\]"

test_divider

productName="Device #1"
echo "$user adds Model with VID: $vid_in_hex_format PID: $pid"
result=$(echo "test1234" | dcld tx model add-model --vid=$vid_in_hex_format --pid=$pid_in_hex_format --productName="$productName" --productLabel="Device Description"   --commissioningCustomFlow=0 --deviceTypeID=12 --partNumber=12 --from=$user_address --yes)
result=$(get_txn_result "$result")
check_response_and_report "$result" "\"code\": 0"

test_divider

vid_plus_one_in_hex_format=0xA14
vidPlusOne=$((vid_in_hex_format+1))
echo "$user adds Model with a VID: $vid_plus_one_in_hex_format PID: $pid_in_hex_format, This fails with Permission denied as the VID is not associated with this vendor account."
result=$(echo "test1234" | dcld tx model add-model --vid=$vid_plus_one_in_hex_format --pid=$pid_in_hex_format --productName="$productName" --productLabel="Device Description"   --commissioningCustomFlow=0 --deviceTypeID=12 --partNumber=12 --from=$user_address --yes 2>&1) || true
result=$(get_txn_result "$result")
check_response_and_report "$result" "transaction should be signed by a vendor account containing the vendorID $vidPlusOne"

test_divider

echo "$user updates Model with VID: $vid_in_hex_format PID: $pid_in_hex_format"
result=$(echo "test1234" | dcld tx model update-model --vid=$vid_in_hex_format --pid=$pid_in_hex_format --productName="$productName" --productLabel="Device Description" --partNumber=12 --from=$user_address --yes)
result=$(get_txn_result "$result")
check_response_and_report "$result" "\"code\": 0"

test_divider

echo "Get Model with VID: $vid_in_hex_format PID: $pid_in_hex_format"
result=$(dcld query model get-model --vid=$vid_in_hex_format --pid=$pid_in_hex_format)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"productName\": \"$productName\""

echo "PASSED"
