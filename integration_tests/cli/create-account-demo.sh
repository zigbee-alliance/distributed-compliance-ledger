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

echo "Get all accouts. No $user account in the list because it is in a pending state"
result=$(zblcli query auth all-accounts)
response_does_not_contain "$result" "\"address\": \"$user_address\""
echo "$result"

echo "Get all proposed accouts. $user account in the list"
result=$(zblcli query auth all-proposed-accounts)
check_response "$result" "\"address\": \"$user_address\""
echo "$result"

echo "Alice approve account for \"$user\""
result=$(echo $passphrase | zblcli tx auth approve-add-account --address="$user_address" --from alice --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get all proposed accouts. No $user account in the list because it received enought approvals"
result=$(zblcli query auth all-proposed-accounts)
response_does_not_contain "$result" "\"address\": \"$user_address\""
echo "$result"

echo "Get all accouts. $user account in the list"
result=$(zblcli query auth all-accounts)
check_response "$result" "\"address\": \"$user_address\""
echo "$result"

echo "Get $user accout"
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
