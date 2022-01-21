set -euo pipefail
source integration_tests/cli/common.sh

random_string user
cmd="(echo $passphrase; echo $passphrase) | dcld keys add $user"
result="$(bash -c "$cmd")"
user_address=$(echo $passphrase | dcld keys show $user -a)
user_pubkey=$(echo $passphrase | dcld keys show $user -p)

random_string user2
cmd="(echo $passphrase; echo $passphrase) | dcld keys add $user2"
result="$(bash -c "$cmd")"
user_address2=$(echo $passphrase | dcld keys show $user2 -a)

# 1. check non-existent values when no entry added via light client
echo "check non-existent values when no entry added via light client"

test_divider

# connect to light client proxy
dcld config node tcp://localhost:26620
sleep 10

echo "Query non existant account"
result=$(execute_with_retry "dcld query auth account --address=$user_address")
echo "$result"
check_response "$result" "Not Found"

test_divider

echo "Query non existant proposed account"
result=$(execute_with_retry "dcld query auth proposed-account --address=$user_address")
echo "$result"
check_response "$result" "Not Found"

test_divider

echo "Query non existant proposed account to revoke"
result=$(execute_with_retry "dcld query auth proposed-account-to-revoke --address=$user_address")
echo "$result"
check_response "$result" "Not Found"

test_divider

# 2. list queries should return a warning via light client

echo "list queries should return a warning via light client"

test_divider

echo "Request all accounts"
result=$(execute_with_retry "dcld query auth all-accounts")
echo "$result"
check_response "$result" "List queries don't work with a Light Client Proxy"

test_divider

echo "Request all proposed accounts"
result=$(execute_with_retry "dcld query auth all-proposed-accounts")
echo "$result"
check_response "$result" "List queries don't work with a Light Client Proxy"

test_divider

echo "Request all proposed accounts to revoke"
result=$(execute_with_retry "dcld query auth all-proposed-accounts-to-revoke")
echo "$result"
check_response "$result" "List queries don't work with a Light Client Proxy"

test_divider

# 3. write entries

echo "write entries"

test_divider

# write requests can be sent via Full Node only
dcld config node tcp://localhost:26657 

vid=$RANDOM

echo "Jack proposes account for $user"
result=$(echo $passphrase | dcld tx auth propose-add-account --address="$user_address" --pubkey="$user_pubkey" --roles="Vendor" --vid="$vid" --from jack --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Alice approves account for \"$user\""
result=$(echo $passphrase | dcld tx auth approve-add-account --address="$user_address" --from alice --yes)
check_response "$result" "\"code\": 0"

test_divider


# 4. check existent values via light client

echo "check existent values via light client"

test_divider

# connect to light client proxy
dcld config node tcp://localhost:26620
sleep 5

echo "Get $user account"
result=$(execute_with_retry "dcld query auth account --address=$user_address")
check_response "$result" "\"address\": \"$user_address\""

test_divider


# 5. check non-existent values when entry added via light client

echo "check non-existent values when entry added via light client"

test_divider

echo "Query non existant account"
result=$(execute_with_retry "dcld query auth account --address=$user_address2")
echo "$result"
check_response "$result" "Not Found"

test_divider

echo "Query non existant proposed account"
result=$(execute_with_retry "dcld query auth proposed-account --address=$user_address2")
echo "$result"
check_response "$result" "Not Found"

test_divider

echo "Query non existant proposed account to revoke"
result=$(execute_with_retry "dcld query auth proposed-account-to-revoke --address=$user_address2")
echo "$result"
check_response "$result" "Not Found"

test_divider

# 6. try to write via light client proxy

echo "try to write via light client proxy"

test_divider

echo "Add vendorinfo"
result=$(echo $passphrase | dcld tx auth propose-add-account --address="$user_address" --pubkey="$user_pubkey" --roles="Vendor" --vid="$vid" --from $user_address --yes)
echo "$result"
check_response "$result" "Write requests don't work with a Light Client Proxy"


test_divider
