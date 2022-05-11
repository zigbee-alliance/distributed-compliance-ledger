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

vid=$RANDOM
pid=$RANDOM

echo "Get not yet approved $user account"
result=$(dcld query auth account --address=$user_address)
check_response "$result" "Not Found"

test_divider

echo "Get not yet proposed $user account to revoke"
result=$(dcld query auth proposed-account-to-revoke --address=$user_address)
check_response "$result" "Not Found"

test_divider

echo "Get not yet proposed $user account"
result=$(dcld query auth proposed-account --address=$user_address)
check_response "$result" "Not Found"

test_divider

echo "Get not yet revoked $user account"
result=$(dcld query auth revoked-account --address=$user_address)
check_response "$result" "Not Found"

test_divider

echo "Get all proposed accounts must be empty"
result=$(dcld query auth all-proposed-accounts)
check_response "$result" "\[\]"
response_does_not_contain "$result" "\"address\": \"$user_address\""

test_divider

echo "Get all proposed accounts to revoke must be empty"
result=$(dcld query auth all-proposed-accounts-to-revoke)
check_response "$result" "\[\]"
response_does_not_contain "$result" "\"address\": \"$user_address\""

test_divider

echo "Get all revoked accounts. No $user account in the list"
result=$(dcld query auth all-revoked-accounts)
check_response "$result" "\[\]"
response_does_not_contain "$result" "\"address\": \"$user_address\""

test_divider

echo "Jack proposes account for $user"
result=$(echo $passphrase | dcld tx auth propose-add-account --info="Jack is proposing this account" --address="$user_address" --pubkey="$user_pubkey" --roles="Vendor" --vid="$vid" --from jack --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Get all active accounts. No $user account in the list because not enough approvals received"
result=$(dcld query auth all-accounts)
response_does_not_contain "$result" "\"address\": \"$user_address\""

test_divider

echo "Get an account for $user is not found"
result=$(dcld query auth account --address=$user_address)
check_response "$result" "Not Found"

echo "Get a proposed account for $user and confirm that the approval contains Jack's address"
result=$(dcld query auth proposed-account --address=$user_address)
check_response "$result" "\"address\": \"$user_address\""
check_response_and_report "$result"  $jack_address "json"
check_response_and_report "$result"  '"info": "Jack is proposing this account"' "json"
response_does_not_contain "$result"  $alice_address "json"


test_divider

echo "Get all proposed accounts. $user account in the list"
result=$(dcld query auth all-proposed-accounts)
check_response "$result" "\"address\": \"$user_address\""

test_divider

echo "Get all proposed accounts to revoke. No $user account in the list"
result=$(dcld query auth all-proposed-accounts-to-revoke)
check_response "$result" "\[\]"

test_divider

echo "Alice approves account for \"$user\""
result=$(echo $passphrase | dcld tx auth approve-add-account --address="$user_address" --info="Alice is approving this account" --from alice --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Alice cannot reject account for  \"$user\""
result=$(echo $passphrase | dcld tx auth reject-add-account --address="$user_address" --info="Alice is rejecting this account" --from alice --yes 2>&1 || true)
response_does_not_contain "$result" "\"code\": 0"

test_divider

echo "Get all accounts. $user account in the list because enough approvals received"
result=$(dcld query auth all-accounts)
check_response "$result" "\"address\": \"$user_address\""

test_divider

echo "Get $user account and confirm that alice and jack are set as approvers"
result=$(dcld query auth account --address=$user_address)
check_response_and_report "$result" $user_address "json"
check_response_and_report "$result"  $jack_address "json"
check_response_and_report "$result"  $alice_address "json"
check_response_and_report "$result"  '"info": "Alice is approving this account"' "json"
check_response_and_report "$result"  '"info": "Jack is proposing this account"' "json"

test_divider

echo "Get all proposed accounts. No $user account in the list anymore"
result=$(dcld query auth all-proposed-accounts)
check_response "$result" "\[\]"

test_divider

echo "Get all proposed accounts to revoke. No $user account in the list"
result=$(dcld query auth all-proposed-accounts-to-revoke)
check_response "$result" "\[\]"

test_divider

echo "Get $user account"
result=$(dcld query auth account --address=$user_address)
check_response_and_report "$result" "\"address\": \"$user_address\""

test_divider

echo "Get a proposed account for $user is not found"
result=$(dcld query auth proposed-account --address=$user_address)
check_response "$result" "Not Found"

test_divider

productName="Device #1"
echo "$user adds Model with VID: $vid PID: $pid"
result=$(echo "test1234" | dcld tx model add-model --vid=$vid --pid=$pid --productName="$productName" --productLabel="Device Description"   --commissioningCustomFlow=0 --deviceTypeID=12 --partNumber=12 --from=$user_address --yes)
check_response_and_report "$result" "\"code\": 0"

test_divider

vidPlusOne=$((vid+1))
echo "$user adds Model with a VID: $vidPlusOne PID: $pid, This fails with Permission denied as the VID is not associated with this vendor account."
result=$(echo "test1234" | dcld tx model add-model --vid=$vidPlusOne --pid=$pid --productName="$productName" --productLabel="Device Description"   --commissioningCustomFlow=0 --deviceTypeID=12 --partNumber=12 --from=$user_address --yes 2>&1) || true
check_response_and_report "$result" "transaction should be signed by a vendor account containing the vendorID $vidPlusOne"

test_divider

echo "$user updates Model with VID: $vid PID: $pid"
result=$(echo "test1234" | dcld tx model update-model --vid=$vid --pid=$pid --productName="$productName" --productLabel="Device Description" --partNumber=12 --from=$user_address --yes)
check_response_and_report "$result" "\"code\": 0"

test_divider

echo "Get Model with VID: $vid PID: $pid"
result=$(dcld query model get-model --vid=$vid --pid=$pid)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"productName\": \"$productName\""

test_divider

echo "Get $user account and confirm that alice and jack are set still present as approvers, also the sequence number should be 3 as we executed 3 txs"
result=$(dcld query auth account --address=$user_address)
check_response_and_report "$result" $user_address "json"
check_response_and_report "$result"  $jack_address "json"
check_response_and_report "$result"  $alice_address "json"
check_response_and_report "$result"  '"info": "Alice is approving this account"' "json"
check_response_and_report "$result"  '"info": "Jack is proposing this account"' "json"
check_response_and_report "$result"  '"sequence": "3"' "json"

test_divider

echo "Alice proposes to revoke account for $user"
result=$(echo $passphrase | dcld tx auth propose-revoke-account --address="$user_address" --info="Alice proposes to revoke account" --from alice --yes)
check_response "$result" "\"code\": 0"


test_divider

echo "Get all accounts. $user account still in the list because not enought approvals to revoke received"
result=$(dcld query auth all-accounts)
check_response "$result" "\"address\": \"$user_address\""


test_divider

echo "Get all proposed accounts. No $user account in the list"
result=$(dcld query auth all-proposed-accounts)
check_response "$result" "\[\]"


test_divider

echo "Get all proposed accounts to revoke. $user account in the list"
result=$(dcld query auth all-proposed-accounts-to-revoke)
check_response "$result" "\"address\": \"$user_address\""


test_divider

echo "Get a proposed account to revoke for $user"
result=$(dcld query auth proposed-account-to-revoke --address="$user_address")
check_response "$result" "\"address\": \"$user_address\""
check_response_and_report "$result"  $alice_address "json"
check_response_and_report "$result"  '"info": "Alice proposes to revoke account"' "json"

test_divider

echo "Get all revoked accounts. No $user account in the list"
result=$(dcld query auth all-revoked-accounts)
check_response "$result" "\[\]"
response_does_not_contain "$result" "\"address\": \"$user_address\""

test_divider

echo "Bob approves to revoke account for $user"
result=$(echo $passphrase | dcld tx auth approve-revoke-account --address="$user_address" --from bob --yes)
check_response "$result" "\"code\": 0"

test_divider

trustee_voting="TrusteeVoting"

echo "Get all revoked accounts. $user account in the list"
result=$(dcld query auth all-revoked-accounts)
check_response "$result" "\"address\": \"$user_address\""
check_response "$result" "\"reason\": \"$trustee_voting\""

test_divider

echo "Get all accounts. No $user account in the list anymore because enought approvals to revoke received"
result=$(dcld query auth all-accounts)
response_does_not_contain "$result" "\"address\": \"$user_address\""

test_divider

echo "Get all proposed accounts. No $user account in the list"
result=$(dcld query auth all-proposed-accounts)
check_response "$result" "\[\]"

test_divider

echo "Get all proposed accounts to revoke. No $user account in the list anymore"
result=$(dcld query auth all-proposed-accounts-to-revoke)
check_response "$result" "\[\]"

test_divider

echo "Get revoked account $user account"
result=$(dcld query auth revoked-account --address="$user_address")
check_response "$result" "\"address\": \"$user_address\""
check_response "$result" "\"reason\": \"$trustee_voting\""

test_divider

echo "Get a proposed account to revoke for $user is not found"
result=$(dcld query auth proposed-account-to-revoke --address="$user_address")
check_response "$result" "Not Found"

test_divider

echo "Get $user account"
result=$(dcld query auth account --address=$user_address)
check_response "$result" "Not Found"

test_divider

echo "Jack proposes account for $user"
result=$(echo $passphrase | dcld tx auth propose-add-account --info="Jack is proposing this account" --address="$user_address" --pubkey="$user_pubkey" --roles="Vendor" --vid="$vid" --from jack --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Get all revoked accounts. $user account in the list"
result=$(dcld query auth all-revoked-accounts)
check_response "$result" "\"address\": \"$user_address\""
check_response "$result" "\"reason\": \"$trustee_voting\""

test_divider

echo "Get all active accounts. No $user account in the list because not enough approvals received"
result=$(dcld query auth all-accounts)
response_does_not_contain "$result" "\"address\": \"$user_address\""

test_divider

echo "Get an account for $user is not found"
result=$(dcld query auth account --address=$user_address)
check_response "$result" "Not Found"

echo "Get a proposed account for $user and confirm that the approval contains Jack's address"
result=$(dcld query auth proposed-account --address=$user_address)
check_response "$result" "\"address\": \"$user_address\""
check_response_and_report "$result"  $jack_address "json"
check_response_and_report "$result"  '"info": "Jack is proposing this account"' "json"
response_does_not_contain "$result"  $alice_address "json"

test_divider

echo "Get all proposed accounts. $user account in the list"
result=$(dcld query auth all-proposed-accounts)
check_response "$result" "\"address\": \"$user_address\""

test_divider

echo "Get all proposed accounts to revoke. No $user account in the list"
result=$(dcld query auth all-proposed-accounts-to-revoke)
check_response "$result" "\[\]"

test_divider

echo "Alice approves account for \"$user\""
result=$(echo $passphrase | dcld tx auth approve-add-account --address="$user_address" --info="Alice is approving this account" --from alice --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Alice doesn't reject account for \"$user\", because Alice already approved account"
result=$(echo $passphrase | dcld tx auth reject-add-account --address="$user_address" --info="Alice is rejecting this account" --from alice --yes 2>&1 || true)
response_does_not_contain "$result" "\"code\": 0"

test_divider

echo "Get all revoked accounts. No $user account in the list"
result=$(dcld query auth all-revoked-accounts)
check_response "$result" "\[\]"
response_does_not_contain "$result" "\"address\": \"$user_address\""

test_divider

echo "Get all accounts. $user account in the list because enough approvals received"
result=$(dcld query auth all-accounts)
check_response "$result" "\"address\": \"$user_address\""

test_divider

echo "Get $user account and confirm that alice and jack are set as approvers"
result=$(dcld query auth account --address=$user_address)
check_response_and_report "$result" $user_address "json"
check_response_and_report "$result"  $jack_address "json"
check_response_and_report "$result"  $alice_address "json"
check_response_and_report "$result"  '"info": "Alice is approving this account"' "json"
check_response_and_report "$result"  '"info": "Jack is proposing this account"' "json"

test_divider

echo "Get all proposed accounts. No $user account in the list anymore"
result=$(dcld query auth all-proposed-accounts)
check_response "$result" "\[\]"

test_divider

echo "Get all proposed accounts to revoke. No $user account in the list"
result=$(dcld query auth all-proposed-accounts-to-revoke)
check_response "$result" "\[\]"

test_divider

echo "Get $user account"
result=$(dcld query auth account --address=$user_address)
check_response_and_report "$result" "\"address\": \"$user_address\""

test_divider

echo "Get a proposed account for $user is not found"
result=$(dcld query auth proposed-account --address=$user_address)
check_response "$result" "Not Found"


test_divider

echo "Alice proposes to revoke account for $user"
result=$(echo $passphrase | dcld tx auth propose-revoke-account --address="$user_address" --info="Alice proposes to revoke account" --from alice --yes)
check_response "$result" "\"code\": 0"


test_divider

echo "Get all accounts. $user account still in the list because not enought approvals to revoke received"
result=$(dcld query auth all-accounts)
check_response "$result" "\"address\": \"$user_address\""


test_divider

echo "Get all proposed accounts. No $user account in the list"
result=$(dcld query auth all-proposed-accounts)
check_response "$result" "\[\]"


test_divider

echo "Get all proposed accounts to revoke. $user account in the list"
result=$(dcld query auth all-proposed-accounts-to-revoke)
check_response "$result" "\"address\": \"$user_address\""


test_divider

echo "Get a proposed account to revoke for $user"
result=$(dcld query auth proposed-account-to-revoke --address="$user_address")
check_response "$result" "\"address\": \"$user_address\""
check_response_and_report "$result"  $alice_address "json"
check_response_and_report "$result"  '"info": "Alice proposes to revoke account"' "json"

test_divider

echo "Get all revoked accounts. No $user account in the list"
result=$(dcld query auth all-revoked-accounts)
check_response "$result" "\[\]"
response_does_not_contain "$result" "\"address\": \"$user_address\""

test_divider

echo "Bob approves to revoke account for $user"
result=$(echo $passphrase | dcld tx auth approve-revoke-account --address="$user_address" --from bob --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Get all revoked accounts. $user account in the list"
result=$(dcld query auth all-revoked-accounts)
check_response "$result" "\"address\": \"$user_address\""

test_divider

echo "Get all accounts. No $user account in the list anymore because enought approvals to revoke received"
result=$(dcld query auth all-accounts)
response_does_not_contain "$result" "\"address\": \"$user_address\""

test_divider

echo "Get all proposed accounts. No $user account in the list"
result=$(dcld query auth all-proposed-accounts)
check_response "$result" "\[\]"

test_divider

echo "Get all proposed accounts to revoke. No $user account in the list anymore"
result=$(dcld query auth all-proposed-accounts-to-revoke)
check_response "$result" "\[\]"

test_divider

echo "Get a proposed account to revoke for $user is not found"
result=$(dcld query auth proposed-account-to-revoke --address="$user_address")
check_response "$result" "Not Found"

test_divider

echo "Get $user account"
result=$(dcld query auth account --address=$user_address)
check_response "$result" "Not Found"

test_divider

echo "Jack proposes account for $user"
result=$(echo $passphrase | dcld tx auth propose-add-account --info="Jack is proposing this account" --address="$user_address" --pubkey="$user_pubkey" --roles="Vendor" --vid="$vid" --from jack --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Get a proposed account for $user and confirm that the approval contains Jack's address"
result=$(dcld query auth proposed-account --address=$user_address)
check_response "$result" "\"address\": \"$user_address\""
check_response_and_report "$result"  $jack_address "json"
check_response_and_report "$result"  '"info": "Jack is proposing this account"' "json"
response_does_not_contain "$result"  $alice_address "json"

test_divider

echo "Alice rejects account for \"$user\""
result=$(echo $passphrase | dcld tx auth reject-add-account --address="$user_address" --info="Alice is rejecting this account" --from alice --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Alice doesn't approve account for \"$user\", because Alice already rejected account"
result=$(echo $passphrase | dcld tx auth approve-add-account --address="$user_address" --info="Alice is approving this account" --from alice --yes 2>&1 || true)
response_does_not_contain "$result" "\"code\": 0"

test_divider

echo "Get all proposed accounts. $user account in the list"
result=$(dcld query auth all-proposed-accounts)
check_response "$result" "\"address\": \"$user_address\""

test_divider

echo "Get all accounts. $user account doesn't exist in the list, because doesn't enough approvals received"
result=$(dcld query auth all-accounts)
response_does_not_contain "$result" "\"address\": \"$user_address\""

test_divider

echo "Get all rejected accounts. $user account doesn't exist in the list, because doesn't enough rejected approvals received"
result=$(dcld query auth all-rejected-accounts)
check_response "$result" "\[\]"

test_divider

echo "Get all proposed accounts to revoke. No $user account in the list"
result=$(dcld query auth all-proposed-accounts-to-revoke)
check_response "$result" "\[\]"

test_divider

echo "Get all revoked accounts to revoke. No $user account in the list"
result=$(dcld query auth all-proposed-accounts-to-revoke)
response_does_not_contain "$result" "\"address\": \"$user_address\""

test_divider

echo "Get a proposed account for $user and confirm that the approval contains Jack's address"
result=$(dcld query auth proposed-account --address=$user_address)
check_response "$result" "\"address\": \"$user_address\""
check_response_and_report "$result"  $jack_address "json"
check_response_and_report "$result"  '"info": "Jack is proposing this account"' "json"
check_response_and_report "$result"  $alice_address "json"
check_response_and_report "$result"  '"info": "Alice is rejecting this account"' "json"

test_divider

echo "Get a account for $user is not found"
result=$(dcld query auth account --address=$user_address)
check_response "$result" "Not Found"

test_divider

echo "Get a proposed account to revoked for $user is not found"
result=$(dcld query auth proposed-account-to-revoke --address=$user_address)
check_response "$result" "Not Found"

test_divider

echo "Get a revoked account for $user is not found"
result=$(dcld query auth revoked-account --address=$user_address)
check_response "$result" "\"address\": \"$user_address\""

test_divider

echo "Bob rejects account for \"$user\""
result=$(echo $passphrase | dcld tx auth reject-add-account --address="$user_address" --info="Bob is rejecting this account" --from bob --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Bob cannot reject the same account \"$user\" for the second time"
result=$(echo $passphrase | dcld tx auth reject-add-account --address="$user_address" --info="Bob is rejecting this account" --from bob --yes 2>&1 || true)
response_does_not_contain "$result" "\"code\": 0"

test_divider

echo "Get all rejected accounts. $user account must be exist in the list, because enough rejected approvals received"
result=$(dcld query auth all-rejected-accounts)
check_response "$result" "\"address\": \"$user_address\""

test_divider

echo "Get all proposed accounts. $user account doesn't exist in the list"
result=$(dcld query auth all-proposed-accounts)
response_does_not_contain "$result" "\"address\": \"$user_address\""

test_divider

echo "Get all accounts. $user account doesn't exist in the list"
result=$(dcld query auth all-accounts)
response_does_not_contain "$result" "\"address\": \"$user_address\""

test_divider

echo "Get all proposed accounts to revoke. No $user account in the list"
result=$(dcld query auth all-proposed-accounts-to-revoke)
check_response "$result" "\[\]"

test_divider

echo "Get all revoked accounts to revoke. No $user account in the list"
result=$(dcld query auth all-proposed-accounts-to-revoke)
response_does_not_contain "$result" "\"address\": \"$user_address\""

test_divider

echo "Get a rejected account for $user. $user account must exist"
result=$(dcld query auth rejected-account --address=$user_address)
check_response "$result" "\"address\": \"$user_address\""
check_response_and_report "$result"  $jack_address "json"
check_response_and_report "$result"  '"info": "Jack is proposing this account"' "json"
check_response_and_report "$result"  $alice_address "json"
check_response_and_report "$result"  '"info": "Alice is rejecting this account"' "json"
check_response_and_report "$result"  $bob_address "json"
check_response_and_report "$result"  '"info": "Bob is rejecting this account"' "json"

test_divider

echo "Get a account for $user is not found"
result=$(dcld query auth account --address=$user_address)
check_response "$result" "Not Found"

test_divider

echo "Get a proposed account to revoked for $user is not found"
result=$(dcld query auth proposed-account-to-revoke --address=$user_address)
check_response "$result" "Not Found"

test_divider

echo "Get a revoked account for $user is not found"
result=$(dcld query auth revoked-account --address=$user_address)
check_response "$result" "\"address\": \"$user_address\""

test_divider

echo "Jack proposes account for $user"
result=$(echo $passphrase | dcld tx auth propose-add-account --info="Jack is proposing this account" --address="$user_address" --pubkey="$user_pubkey" --roles="Vendor" --vid="$vid" --from jack --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Get a proposed account for $user and confirm that the approval contains Jack's address"
result=$(dcld query auth proposed-account --address=$user_address)
check_response "$result" "\"address\": \"$user_address\""
check_response_and_report "$result"  $jack_address "json"
check_response_and_report "$result"  '"info": "Jack is proposing this account"' "json"
response_does_not_contain "$result"  $alice_address "json"

test_divider

echo "Get a rejected account for $user is not found"
result=$(dcld query auth rejected-account --address=$user_address)
check_response "$result" "Not Found"

test_divider

pid=$RANDOM
productName="Device #2"
echo "$user adds Model with VID: $vid PID: $pid"
result=$(echo "test1234" | dcld tx model add-model --vid=$vid --pid=$pid --productName="$productName" --productLabel="Device Description" --commissioningCustomFlow=0 --deviceTypeID=12 --partNumber=12 --from=$user_address --yes 2>&1) || true
check_response_and_report "$result" "key not found" raw

echo "PASSED"
