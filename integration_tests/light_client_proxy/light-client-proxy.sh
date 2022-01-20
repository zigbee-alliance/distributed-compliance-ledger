set -euo pipefail
source integration_tests/cli/common.sh

# 1. check non-existent values when no enrty added

# connect to light client proxy
dcld config node tcp://localhost:26620
sleep 10

vid=$RANDOM
pid=$RANDOM

echo "Query non existant model"
result=$(dcld query model get-model --vid=$vid --pid=$pid)
check_response "$result" "Not Found"
echo "$result"

test_divider

# 2. write entries

# write requests can be sent via Full Node only
dcld config node tcp://localhost:26657 

vid=$RANDOM
pid=$RANDOM

vendor_account=vendor_account_$vid
echo "Create Vendor account - $vendor_account"
create_new_vendor_account $vendor_account $vid

test_divider

productLabel="Device #1"
echo "Add Model with VID: $vid PID: $pid"
result=$(echo "test1234" | dcld tx model add-model --vid=$vid --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel="$productLabel" --partNumber=1 --commissioningCustomFlow=0 --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

# 3. check existent values

# connect to light client proxy
dcld config node tcp://localhost:26620
sleep 5

echo "Get Model with VID: $vid PID: $pid"
result=$(dcld query model get-model --vid=$vid --pid=$pid)
echo "$result"
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"product_label\": \"$productLabel\""

test_divider

# 4. check non-existent values when entry added
vid=$RANDOM
pid=$RANDOM

echo "Query non existant model"
result=$(dcld query model get-model --vid=$vid --pid=$pid)
check_response "$result" "Not Found"
echo "$result"

test_divider