#!/bin/bash
set -e
source integration_tests/cli/common.sh

user=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 6 | head -n 1)
echo "$user generates keys"
result=$(echo 'test1234' | zblcli keys add $user)
echo "$result"

echo "Get key info for $user"
result=$(zblcli keys show $user)
check_response "$result" "\"name\": \"$user\""
echo "$result"

user_address=$(zblcli keys show $user -a)
user_pubkey=$(zblcli keys show $user -p)

echo "Jack proposes account for $user"
result=$(echo $passphrase | zblcli tx auth propose-add-account --address="$user_address" --pubkey="$user_pubkey" --roles="Vendor" --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get all active accounts. No $user account in the list because not enough approvals received"
result=$(zblcli query auth all-accounts)
response_does_not_contain "$result" "\"address\": \"$user_address\""
echo "$result"

echo "Get all proposed accounts. $user account in the list"
result=$(zblcli query auth all-proposed-accounts)
check_response "$result" "\"address\": \"$user_address\""
echo "$result"

echo "Get all proposed accounts to revoke. No $user account in the list"
result=$(zblcli query auth all-proposed-accounts-to-revoke)
response_does_not_contain "$result" "\"address\": \"$user_address\""
echo "$result"

echo "Alice approves account for \"$user\""
result=$(echo $passphrase | zblcli tx auth approve-add-account --address="$user_address" --from alice --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get all accounts. $user account in the list because enough approvals received"
result=$(zblcli query auth all-accounts)
check_response "$result" "\"address\": \"$user_address\""
echo "$result"

echo "Get all proposed accounts. No $user account in the list anymore"
result=$(zblcli query auth all-proposed-accounts)
response_does_not_contain "$result" "\"address\": \"$user_address\""
echo "$result"

echo "Get all proposed accounts to revoke. No $user account in the list"
result=$(zblcli query auth all-proposed-accounts-to-revoke)
response_does_not_contain "$result" "\"address\": \"$user_address\""
echo "$result"

echo "Get $user account"
result=$(zblcli query auth account --address=$user_address)
check_response "$result" "\"address\": \"$user_address\""
echo "$result"

vid=$RANDOM
pid=$RANDOM
name="Device #1"
echo "$user adds Model with VID: $vid PID: $pid"
result=$(echo "test1234" | zblcli tx modelinfo add-model --vid=$vid --pid=$pid --name="$name" --description="Device Description" --sku="SKU12FS" --firmware-version="1.0" --hardware-version="2.0" --tis-or-trp-testing-completed=true --from=$user_address --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get Model with VID: $vid PID: $pid"
result=$(zblcli query modelinfo model --vid=$vid --pid=$pid)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"name\": \"$name\""
echo "$result"

echo "Alice proposes to revoke account for $user"
result=$(echo $passphrase | zblcli tx auth propose-revoke-account --address="$user_address" --from alice --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get all accounts. $user account still in the list because not enought approvals to revoke received"
result=$(zblcli query auth all-accounts)
check_response "$result" "\"address\": \"$user_address\""
echo "$result"

echo "Get all proposed accounts. No $user account in the list"
result=$(zblcli query auth all-proposed-accounts)
response_does_not_contain "$result" "\"address\": \"$user_address\""
echo "$result"

echo "Get all proposed accounts to revoke. $user account in the list"
result=$(zblcli query auth all-proposed-accounts-to-revoke)
check_response "$result" "\"address\": \"$user_address\""
echo "$result"

echo "Bob approves to revoke account for $user"
result=$(echo $passphrase | zblcli tx auth approve-revoke-account --address="$user_address" --from bob --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get all accounts. No $user account in the list anymore because enought approvals to revoke received"
result=$(zblcli query auth all-accounts)
response_does_not_contain "$result" "\"address\": \"$user_address\""
echo "$result"

echo "Get all proposed accounts. No $user account in the list"
result=$(zblcli query auth all-proposed-accounts)
response_does_not_contain "$result" "\"address\": \"$user_address\""
echo "$result"

echo "Get all proposed accounts to revoke. No $user account in the list anymore"
result=$(zblcli query auth all-proposed-accounts-to-revoke)
response_does_not_contain "$result" "\"address\": \"$user_address\""
echo "$result"

echo "Get $user account"
result=$(zblcli query auth account --address=$user_address)
echo "$result"

vid=$RANDOM
pid=$RANDOM
name="Device #2"
echo "$user adds Model with VID: $vid PID: $pid"
result=$(echo "test1234" | zblcli tx modelinfo add-model --vid=$vid --pid=$pid --name="$name" --description="Device Description" --sku="SKU12FS" --firmware-version="1.0" --hardware-version="2.0" --tis-or-trp-testing-completed=true --from=$user_address --yes)
echo "$result"
