#!/bin/bash
set -e
source integration_tests/cli/common.sh

echo "Get key info for Jack"
result=$(zblcli keys show jack)
check_response "$result" "\"name\": \"jack\""
echo "$result"

echo "Jack assigns Trustee role to Jack"
result=$(echo "test1234" | zblcli tx auth assign-role --address=$(zblcli keys show jack -a) --role="Trustee" --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

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

echo "Jack creates account for $user"
result=$(echo "test1234" | zblcli tx auth create-account --address="$user_address" --pubkey="$user_pubkey" --roles="Vendor" --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get all accouts"
result=$(zblcli query auth accounts)
check_response "$result" "\"address\": \"$user_address\""
check_response "$result" "\"public_key\": \"$user_pubkey\""
echo "$result"

echo "Get $user accout"
check_response "$result" "\"address\": \"$user_address\""
check_response "$result" "\"public_key\": \"$user_pubkey\""
result=$(zblcli query auth account --address=$user_address)
check_response "$result" "\"roles\": []"
echo "$result"

echo "Jack assigns Vendor role to Tony"
result=$(echo "test1234" | zblcli tx auth assign-role --address=$tony_address --role="Vendor" --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get account roles for Tony"
result=$(zblcli query auth account-roles --address=$tony_address)
check_response "$result" "\"Vendor\""
echo "$result"

vid=$RANDOM
pid=$RANDOM
name="Device #1"

echo "$user adds Model with VID: $vid PID: $pid"
result=$(echo "test1234" | zblcli tx modelinfo add-model $vid $pid "$name" "Device Description" "SKU12FS" "1.0" "2.0" true --from $user --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get Model with VID: $vid PID: $pid"
result=$(zblcli query modelinfo model --vid=$vid --pid=$pid)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"name\": \"$name\""
echo "$result"
