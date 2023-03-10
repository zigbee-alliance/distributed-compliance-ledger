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
result=$(echo $passphrase | dcld tx auth propose-add-account --info="Jack is proposing this account" --address="$user_address" --pubkey="$user_pubkey" --roles="NodeAdmin" --from jack --yes)
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

echo "Get $user account and confirm that alice and jack are set still present as approvers"
result=$(dcld query auth account --address=$user_address)
check_response_and_report "$result" $user_address "json"
check_response_and_report "$result"  $jack_address "json"
check_response_and_report "$result"  $alice_address "json"
check_response_and_report "$result"  '"info": "Alice is approving this account"' "json"
check_response_and_report "$result"  '"info": "Jack is proposing this account"' "json"

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
result=$(echo $passphrase | dcld tx auth propose-add-account --info="Jack is proposing this account" --address="$user_address" --pubkey="$user_pubkey" --roles="NodeAdmin" --from jack --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Jack can reject account for \"$user\" even if Jack already approved account"
result=$(echo $passphrase | dcld tx auth reject-add-account --address="$user_address" --info="Jack is rejecting this account" --from jack --yes 2>&1 || true)
check_response "$result" "\"code\": 0"

test_divider

echo "Jack re-approves account for \"$user\""
result=$(echo $passphrase | dcld tx auth approve-add-account --address="$user_address" --info="Jack is proposing this account" --from jack --yes)
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
result=$(echo $passphrase | dcld tx auth propose-add-account --info="Jack is proposing this account" --address="$user_address" --pubkey="$user_pubkey" --roles="NodeAdmin" --from jack --yes)
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

vid=$RANDOM
pid=$RANDOM
productName="Device #2"
echo "$user adds Model with VID: $vid PID: $pid"
result=$(echo "test1234" | dcld tx model add-model --vid=$vid --pid=$pid --productName="$productName" --productLabel="Device Description" --commissioningCustomFlow=0 --deviceTypeID=12 --partNumber=12 --from=$user_address --yes 2>&1) || true
check_response_and_report "$result" "key not found" raw


# These tests are needed so that we can check that when we add a Vendor account, we need 1/3 of the approvals.
test_divider

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

test_divider

echo "Jack proposes account for $user"
result=$(echo $passphrase | dcld tx auth propose-add-account --info="Jack is proposing this account" --address="$user_address" --pubkey="$user_pubkey" --roles="Vendor,NodeAdmin" --vid=$vid --from jack --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Get all active accounts. $user account is not in the list because has not enough approvals received"
result=$(dcld query auth all-accounts)
response_does_not_contain "$result" "\"address\": \"$user_address\""

test_divider

echo "Get an account for $user. $user not found"
result=$(dcld query auth account --address=$user_address)
check_response "$result" "Not Found"

test_divider

echo "Get all proposed accounts. $user account is in the list because has not enough approvals received"
result=$(dcld query auth all-proposed-accounts)
check_response "$result" "\"address\": \"$user_address\""

test_divider

echo "Get an proposed account for $user"
result=$(dcld query auth proposed-account --address=$user_address)
check_response "$result" "\"address\": \"$user_address\""
check_response_and_report "$result"  $jack_address "json"
check_response_and_report "$result"  '"info": "Jack is proposing this account"' "json"

test_divider

echo "Alice approves account for \"$user\""
result=$(echo $passphrase | dcld tx auth approve-add-account --address="$user_address" --info="Alice is approving this account" --from alice --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Get all active accounts. $user account is in the list because has enough approvals received"
result=$(dcld query auth all-accounts)
check_response "$result" "\"address\": \"$user_address\""

test_divider

echo "Get an account for $user"
result=$(dcld query auth account --address=$user_address)
check_response "$result" "\"address\": \"$user_address\""
check_response_and_report "$result"  $jack_address "json"
check_response_and_report "$result"  '"info": "Jack is proposing this account"' "json"
check_response_and_report "$result"  $alice_address "json"
check_response_and_report "$result"  '"info": "Alice is approving this account"' "json"

test_divider

echo "Get all proposed accounts. $user account is not in the list because has enough approvals received"
result=$(dcld query auth all-proposed-accounts)
response_does_not_contain "$result" "\"address\": \"$user_address\""

test_divider

echo "Get an proposed account for $user. $user not found"
result=$(dcld query auth proposed-account --address=$user_address)
check_response "$result" "Not Found"

test_divider

vid=$RANDOM
pid=$RANDOM

test_divider

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

test_divider

echo "Jack proposes account for $user"
result=$(echo $passphrase | dcld tx auth propose-add-account --info="Jack is proposing this account" --address="$user_address" --pubkey="$user_pubkey" --roles="Vendor" --vid=$vid --from jack --yes)
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

random_string new_trustee1
echo "$new_trustee1 generates keys"
cmd="(echo $passphrase; echo $passphrase) | dcld keys add $new_trustee1"
result="$(bash -c "$cmd")"

test_divider

echo "Get key info for $new_trustee1"
result=$(echo $passphrase | dcld keys show $new_trustee1)
check_response "$result" "\"name\": \"$new_trustee1\""

test_divider

new_trustee_address1=$(echo $passphrase | dcld keys show $new_trustee1 -a)
new_trustee_pubkey1=$(echo $passphrase | dcld keys show $new_trustee1 -p)

test_divider

echo "Jack proposes account for $new_trustee1"
result=$(echo $passphrase | dcld tx auth propose-add-account --info="Jack is proposing this account" --address="$new_trustee_address1" --pubkey="$new_trustee_pubkey1" --roles="Trustee" --from jack --yes)
check_response "$result" "\"code\": 0"

echo "Alice approves account for \"$new_trustee1\""
result=$(echo $passphrase | dcld tx auth approve-add-account --address="$new_trustee_address1" --info="Alice is approving this account" --from alice --yes)
check_response "$result" "\"code\": 0"

test_divider

random_string new_trustee2
echo "$new_trustee2 generates keys"
cmd="(echo $passphrase; echo $passphrase) | dcld keys add $new_trustee2"
result="$(bash -c "$cmd")"

test_divider

echo "Get key info for $new_trustee2"
result=$(echo $passphrase | dcld keys show $new_trustee2)
check_response "$result" "\"name\": \"$new_trustee2\""

test_divider

new_trustee_address2=$(echo $passphrase | dcld keys show $new_trustee2 -a)
new_trustee_pubkey2=$(echo $passphrase | dcld keys show $new_trustee2 -p)

test_divider

echo "Jack proposes account for $new_trustee2"
result=$(echo $passphrase | dcld tx auth propose-add-account --info="Jack is proposing this account" --address="$new_trustee_address2" --pubkey="$new_trustee_pubkey2" --roles="Trustee" --from jack --yes)
check_response "$result" "\"code\": 0"

echo "Alice approves account for \"$new_trustee2\""
result=$(echo $passphrase | dcld tx auth approve-add-account --address="$new_trustee_address2" --info="Alice is approving this account" --from alice --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Bob approves account for \"$new_trustee2\""
result=$(echo $passphrase | dcld tx auth approve-add-account --address="$new_trustee_address2" --info="Bob is approving this account" --from bob --yes)
check_response "$result" "\"code\": 0"

test_divider

vid=$RANDOM
pid=$RANDOM

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

test_divider

echo "Jack proposes account for $user"
result=$(echo $passphrase | dcld tx auth propose-add-account --info="Jack is proposing this account" --address="$user_address" --pubkey="$user_pubkey" --roles="Vendor" --vid=$vid --from jack --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Get all active accounts. $user account is not in the list because has not enough approvals received"
result=$(dcld query auth all-accounts)
response_does_not_contain "$result" "\"address\": \"$user_address\""

test_divider

echo "Get an account for $user. $user not found"
result=$(dcld query auth account --address=$user_address)
check_response "$result" "Not Found"

test_divider

echo "Get all proposed accounts. $user account is in the list because has not enough approvals received"
result=$(dcld query auth all-proposed-accounts)
check_response "$result" "\"address\": \"$user_address\""

test_divider

echo "Get an proposed account for $user"
result=$(dcld query auth proposed-account --address=$user_address)
check_response "$result" "\"address\": \"$user_address\""
check_response_and_report "$result"  $jack_address "json"
check_response_and_report "$result"  '"info": "Jack is proposing this account"' "json"

test_divider

echo "Alice approves account for \"$user\""
result=$(echo $passphrase | dcld tx auth approve-add-account --address="$user_address" --info="Alice is approving this account" --from alice --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Get all active accounts. $user account is in the list because has enough approvals received"
result=$(dcld query auth all-accounts)
check_response "$result" "\"address\": \"$user_address\""

test_divider

echo "Get an account for $user"
result=$(dcld query auth account --address=$user_address)
check_response "$result" "\"address\": \"$user_address\""
check_response_and_report "$result"  $jack_address "json"
check_response_and_report "$result"  '"info": "Jack is proposing this account"' "json"
check_response_and_report "$result"  $alice_address "json"
check_response_and_report "$result"  '"info": "Alice is approving this account"' "json"

test_divider

echo "Get all proposed accounts. $user account is not in the list because has enough approvals received"
result=$(dcld query auth all-proposed-accounts)
response_does_not_contain "$result" "\"address\": \"$user_address\""

test_divider

echo "Get an proposed account for $user. $user not found"
result=$(dcld query auth proposed-account --address=$user_address)
check_response "$result" "Not Found"

test_divider


###########################################################################################################################################
# THESE TEST ARE NEED TO CHECK THIS SITUATION:
# Example:
# There are 10 trustees in the network
# Account "A" is proposed to be added (requires at least 7 approvals. 2/3 of trustees)
# Account "A" receives 6 approvals (one more approval is required)
# Meanwhile 1 trustee account has been revoked and there are 9 trustees in the network
# Account "A" is still in pending approvals, even if it has 6 approvals which are enough after the revocation of a trustee
# Account "A" receives one more approval and now it has 7 approvals. 
# The account has not been added even if it has more than the required number of approvals (6 required)


# REVOKE ACCOUNT, WE NEED 4 TRUSTEE'S APPROVALS, BECAUSE WE HAVE 5 TRUSTEES
echo "Alice proposes to revoke account for $user"
result=$(echo $passphrase | dcld tx auth propose-revoke-account --address="$user_address" --info="Alice proposes to revoke account" --from alice --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Get a proposed account to revoke for $user"
result=$(dcld query auth proposed-account-to-revoke --address="$user_address")
check_response "$result" "\"address\": \"$user_address\""
check_response_and_report "$result"  $alice_address "json"
check_response_and_report "$result"  '"info": "Alice proposes to revoke account"' "json"

test_divider

echo "Bob approves to revoke account for $user"
result=$(echo $passphrase | dcld tx auth approve-revoke-account --address="$user_address" --from bob --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Get a proposed account to revoke for $user"
result=$(dcld query auth proposed-account-to-revoke --address="$user_address")
check_response "$result" "\"address\": \"$user_address\""
check_response_and_report "$result"  $alice_address "json"
check_response_and_report "$result"  '"info": "Alice proposes to revoke account"' "json"

test_divider

# REMOVE TRUSTEE ACCOUNT
echo "Alice proposes to revoke account for $new_trustee1"
result=$(echo $passphrase | dcld tx auth propose-revoke-account --address="$new_trustee_address1" --from alice --yes)
check_response "$result" "\"code\": 0"

echo "Bob approves to revoke account for $new_trustee1"
result=$(echo $passphrase | dcld tx auth approve-revoke-account --address="$new_trustee_address1" --from bob --yes)
check_response "$result" "\"code\": 0"

echo "Jack approves to revoke account for $new_trustee1"
result=$(echo $passphrase | dcld tx auth approve-revoke-account --address="$new_trustee_address1" --from jack --yes)
check_response "$result" "\"code\": 0"

echo "$new_trustee1 approves to revoke account for $new_trustee1"
result=$(echo $passphrase | dcld tx auth approve-revoke-account --address="$new_trustee_address1" --from $new_trustee1 --yes)
check_response "$result" "\"code\": 0"

echo "Get revoked account $new_trustee1"
result=$(dcld query auth revoked-account --address="$new_trustee_address1")
check_response "$result" "\"address\": \"$new_trustee_address1\""
check_response "$result" "\"reason\": \"$trustee_voting\""

test_divider

# REVOKE ACCOUNT
echo "Jack approves to revoke account for $user"
result=$(echo $passphrase | dcld tx auth approve-revoke-account --address="$user_address" --from jack --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Get revoked account $user account"
result=$(dcld query auth revoked-account --address="$user_address")
check_response "$result" "\"address\": \"$user_address\""
check_response "$result" "\"reason\": \"$trustee_voting\""

test_divider



# REJECT A NEW ACCOUNT, WE NEED 3 TRUSTEE'S REJECTS, BECAUSE WE HAVE 4 TRUSTEES
echo "Jack proposes account for $user"
result=$(echo $passphrase | dcld tx auth propose-add-account --info="Jack is proposing this account" --address="$user_address" --pubkey="$user_pubkey" --roles="Vendor,NodeAdmin" --vid=$vid --from jack --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Get a proposed account for $user and confirm that the approval contains Jack's address"
result=$(dcld query auth proposed-account --address=$user_address)
check_response "$result" "\"address\": \"$user_address\""
check_response_and_report "$result"  $jack_address "json"
check_response_and_report "$result"  '"info": "Jack is proposing this account"' "json"
response_does_not_contain "$result"  $alice_address "json"

test_divider

echo "Bob rejects account for \"$user\""
result=$(echo $passphrase | dcld tx auth reject-add-account --address="$user_address" --info="Bob is rejecting this account" --from bob --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Get a proposed account for $user and confirm that the approval contains Jack's address"
result=$(dcld query auth proposed-account --address=$user_address)
check_response "$result" "\"address\": \"$user_address\""
check_response_and_report "$result"  $jack_address "json"
check_response_and_report "$result"  '"info": "Jack is proposing this account"' "json"
check_response_and_report "$result"  $bob_address "json"
check_response_and_report "$result"  '"info": "Bob is rejecting this account"' "json"

test_divider

# WE REMOVE TRUSTEE ACCOUNT
echo "Alice proposes to revoke account for $new_trustee2"
result=$(echo $passphrase | dcld tx auth propose-revoke-account --address="$new_trustee_address2" --from alice --yes)
check_response "$result" "\"code\": 0"

echo "Bob approves to revoke account for $new_trustee2"
result=$(echo $passphrase | dcld tx auth approve-revoke-account --address="$new_trustee_address2" --from bob --yes)
check_response "$result" "\"code\": 0"

echo "Jack approves to revoke account for $new_trustee2"
result=$(echo $passphrase | dcld tx auth approve-revoke-account --address="$new_trustee_address2" --from jack --yes)
check_response "$result" "\"code\": 0"

echo "Get revoked account $new_trustee2"
result=$(dcld query auth revoked-account --address="$new_trustee_address2")
check_response "$result" "\"address\": \"$new_trustee_address2\""
check_response "$result" "\"reason\": \"$trustee_voting\""

test_divider

# REJECT A NEW ACCOUNT
echo "Alice rejects account for \"$user\""
result=$(echo $passphrase | dcld tx auth reject-add-account --address="$user_address" --info="Alice is rejecting this account" --from alice --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Get all rejected accounts. $user account is in the list because has enough rejects received"
result=$(dcld query auth all-rejected-accounts)
check_response "$result" "\"address\": \"$user_address\""

test_divider

echo "Get rejected account for $user"
result=$(dcld query auth rejected-account --address=$user_address)
check_response "$result" "\"address\": \"$user_address\""
check_response_and_report "$result"  $jack_address "json"
check_response_and_report "$result"  '"info": "Jack is proposing this account"' "json"
check_response_and_report "$result"  $bob_address "json"
check_response_and_report "$result"  '"info": "Bob is rejecting this account"' "json"
check_response_and_report "$result"  $alice_address "json"
check_response_and_report "$result"  '"info": "Alice is rejecting this account"' "json"

test_divider



# ADD A NEW TRUSTEE ACCOUNT
random_string new_trustee
echo "$new_trustee generates keys"
cmd="(echo $passphrase; echo $passphrase) | dcld keys add $new_trustee"
result="$(bash -c "$cmd")"

echo "Get key info for $new_trustee"
result=$(echo $passphrase | dcld keys show $new_trustee)
check_response "$result" "\"name\": \"$new_trustee\""

new_trustee_address=$(echo $passphrase | dcld keys show $new_trustee -a)
new_trustee_pubkey=$(echo $passphrase | dcld keys show $new_trustee -p)

echo "Jack proposes account for $new_trustee"
result=$(echo $passphrase | dcld tx auth propose-add-account --info="Jack is proposing this account" --address="$new_trustee_address" --pubkey="$new_trustee_pubkey" --roles="Trustee" --from jack --yes)
check_response "$result" "\"code\": 0"

echo "Alice approves account for \"$new_trustee\""
result=$(echo $passphrase | dcld tx auth approve-add-account --address="$new_trustee_address" --info="Alice is approving this account" --from alice --yes)
check_response "$result" "\"code\": 0"

test_divider

# ADD A NEW ACCOUNT, WE NEED 3 TRUSTEE'S APPROVALS, BECAUSE WE HAVE 4 TRUSTEES
echo "Jack proposes account for $user"
result=$(echo $passphrase | dcld tx auth propose-add-account --info="Jack is proposing this account" --address="$user_address" --pubkey="$user_pubkey" --roles="Vendor,NodeAdmin" --vid=$vid --from jack --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Get a proposed account for $user and confirm that the approval contains Jack's address"
result=$(dcld query auth proposed-account --address=$user_address)
check_response "$result" "\"address\": \"$user_address\""
check_response_and_report "$result"  $jack_address "json"
check_response_and_report "$result"  '"info": "Jack is proposing this account"' "json"

test_divider

echo "Bob approves account for \"$new_trustee\""
result=$(echo $passphrase | dcld tx auth approve-add-account --info="Bob is approving this account" --address="$user_address" --from bob --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Get a proposed account for $user and confirm that the approval contains Jack's address"
result=$(dcld query auth proposed-account --address=$user_address)
check_response "$result" "\"address\": \"$user_address\""
check_response_and_report "$result"  $jack_address "json"
check_response_and_report "$result"  '"info": "Jack is proposing this account"' "json"
check_response_and_report "$result"  $bob_address "json"
check_response_and_report "$result"  '"info": "Bob is approving this account"' "json"

test_divider

# WE REMOVE TRUSTEE ACCOUNT
echo "Alice proposes to revoke account for $new_trustee"
result=$(echo $passphrase | dcld tx auth propose-revoke-account --address="$new_trustee_address" --from alice --yes)
check_response "$result" "\"code\": 0"

echo "Bob approves to revoke account for $new_trustee"
result=$(echo $passphrase | dcld tx auth approve-revoke-account --address="$new_trustee_address" --from bob --yes)
check_response "$result" "\"code\": 0"

echo "Jack approves to revoke account for $new_trustee"
result=$(echo $passphrase | dcld tx auth approve-revoke-account --address="$new_trustee_address" --from jack --yes)
check_response "$result" "\"code\": 0"

echo "Get revoked account $new_trustee"
result=$(dcld query auth revoked-account --address="$new_trustee_address")
check_response "$result" "\"address\": \"$new_trustee_address\""
check_response "$result" "\"reason\": \"$trustee_voting\""

test_divider

# ADD A NEW ACCOUNT
echo "Alice approves account for \"$user\""
result=$(echo $passphrase | dcld tx auth approve-add-account --address="$user_address" --info="Alice is approving this account" --from alice --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Get all active accounts. $user account is in the list because has enough approvals received"
result=$(dcld query auth all-accounts)
check_response "$result" "\"address\": \"$user_address\""

test_divider

echo "Get an account for $user"
result=$(dcld query auth account --address=$user_address)
check_response "$result" "\"address\": \"$user_address\""
check_response_and_report "$result"  $jack_address "json"
check_response_and_report "$result"  '"info": "Jack is proposing this account"' "json"
check_response_and_report "$result"  $bob_address "json"
check_response_and_report "$result"  '"info": "Bob is approving this account"' "json"
check_response_and_report "$result"  $alice_address "json"
check_response_and_report "$result"  '"info": "Alice is approving this account"' "json"

test_divider

echo "Get an proposed account for $user. $user not found"
result=$(dcld query auth proposed-account --address=$user_address)
check_response "$result" "Not Found"

test_divider
###########################################################################################################################################


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


test_divider

echo "PASSED"
