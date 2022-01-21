set -euo pipefail
source integration_tests/cli/common.sh

# 1. check non-existent values when no entry added via light client
echo "check non-existent values when no entry added via light client"

test_divider

# connect to light client proxy
dcld config node tcp://localhost:26620
sleep 10

vid=$RANDOM
pid=$RANDOM
sv=$RANDOM
svs=$RANDOM

echo "Query non existant compliancetest"
result=$(execute_with_retry "dcld query compliancetest test-result --vid=$vid --pid=$pid --softwareVersion=$sv")
echo "$result"
check_response "$result" "Not Found"

test_divider

# 2. list queries should return a warning via light client

echo "list queries should return a warning via light client"

test_divider

echo "Request all compliancetests"
result=$(execute_with_retry "dcld query compliancetest all-test-results")
echo "$result"
check_response "$result" "List queries don't work with a Light Client Proxy"

test_divider

# 3. write entries

echo "write entries"

test_divider

# write requests can be sent via Full Node only
dcld config node tcp://localhost:26657 

vid=$RANDOM
pid=$RANDOM
sv=$RANDOM
svs=$RANDOM
testing_result="http://first.place.com"
test_date="2020-11-24T10:00:00Z"

vendor_account=vendor_account_$vid
echo "Create Vendor account - $vendor_account"
create_new_vendor_account $vendor_account $vid

test_divider

echo "Create TestHouse account"
create_new_account test_house_account "TestHouse"

test_divider

echo "Add Model and a New Model Version with VID: $vid PID: $pid SV: $sv"
create_model_and_version $vid $pid $sv $svs $vendor_account

test_divider

echo "Add Testing Result for Model VID: $vid PID: $pid SV: $sv"
result=$(echo "$passphrase" | dcld tx compliancetest add-test-result --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --test-result="$testing_result" --test-date="$test_date" --from $test_house_account --yes)
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

# 4. check existent values via light client

echo "check existent values via light client"

test_divider

# connect to light client proxy
dcld config node tcp://localhost:26620
sleep 5

echo "Get compliancetest"
result=$(execute_with_retry "dcld query compliancetest test-result --vid=$vid --pid=$pid --softwareVersion=$sv")
echo "$result"
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"software_version\": $sv"
check_response "$result" "\"software_version_string\": \"$svs\""
check_response "$result" "\"test_result\": \"$testing_result\""
check_response "$result" "\"test_date\": \"$test_date\""
check_response "$result" "\"test_result\": \"$testing_result\""
check_response "$result" "\"test_date\": \"$test_date\""

test_divider


# 5. check non-existent values when entry added via light client

echo "check non-existent values when entry added via light client"

test_divider

vid=$RANDOM
pid=$RANDOM
sv=$RANDOM
svs=$RANDOM

echo "Query non existant compliancetest"
result=$(execute_with_retry "dcld query compliancetest test-result --vid=$vid --pid=$pid --softwareVersion=$sv")
echo "$result"
check_response "$result" "Not Found"

test_divider

# 6. try to write via light client proxy

echo "try to write via light client proxy"

test_divider

echo "Add compliance test"
result=$(echo '$passphrase' | dcld tx compliancetest add-test-result --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --test-result="$testing_result" --test-date="$test_date" --from $test_house_account --yes)
echo "$result"
check_response "$result" "Write requests don't work with a Light Client Proxy"


test_divider
