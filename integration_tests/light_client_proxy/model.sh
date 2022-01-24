set -euo pipefail
source integration_tests/cli/common.sh

# 1. check non-existent values via light client when no entry added
echo "check non-existent values via light client when no entry added"

test_divider

# connect to light client proxy
dcld config node tcp://localhost:26620
sleep 10

vid=$RANDOM
pid=$RANDOM
sv=$RANDOM

echo "Query non existent model"
result=$(execute_with_retry "dcld query model get-model --vid=$vid --pid=$pid")
echo "$result"
check_response "$result" "Not Found"

test_divider

echo "Query non existent Vendor Models"
result=$(execute_with_retry "dcld query model vendor-models --vid=$vid")
echo "$result"
check_response "$result" "Not Found"

test_divider

echo "Query non existent model version"
result=$(execute_with_retry "dcld query model get-model-version --vid=$vid --pid=$pid --softwareVersion=$sv")
echo "$result"
check_response "$result" "Not Found"

test_divider

echo "Query non existent model versions"
result=$(execute_with_retry "dcld query model all-model-versions --vid=$vid --pid=$pid")
echo "$result"
check_response "$result" "Not Found"

test_divider

# 2. list queries should return a warning via light client

echo "list queries should return a warning via light client"

test_divider

echo "Request all models"
result=$(execute_with_retry "dcld query model all-models")
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

vendor_account=vendor_account_$vid
echo "Create Vendor account - $vendor_account"
create_new_vendor_account $vendor_account $vid

test_divider

productLabel="Device #1"
echo "Add Model with VID: $vid PID: $pid"
result=$(echo "test1234" | dcld tx model add-model --vid=$vid --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel="$productLabel" --partNumber=1 --commissioningCustomFlow=0 --from=$vendor_account --yes)
echo "$result"
check_response "$result" "\"code\": 0"

test_divider

echo "Create a Device Model Version with VID: $vid PID: $pid SV: $sv"
result=$(echo 'test1234' | dcld tx model add-model-version --cdVersionNumber=1 --maxApplicableSoftwareVersion=10 --minApplicableSoftwareVersion=1 --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=1 --from=$vendor_account --yes)
echo "$result"
check_response "$result" "\"code\": 0"

test_divider

# 4. check existent values via light client

echo "check existent values via light client"

test_divider

# connect to light client proxy
dcld config node tcp://localhost:26620
sleep 5

echo "Get Model with VID: $vid PID: $pid"
result=$(execute_with_retry "dcld query model get-model --vid=$vid --pid=$pid")
echo "$result"
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"product_label\": \"$productLabel\""

test_divider

echo "Query Device Model Version with VID: $vid PID: $pid SV: $sv"
result=$(execute_with_retry "dcld query model get-model-version --vid=$vid --pid=$pid --softwareVersion=$sv")
echo "$result"
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"software_version\": $sv"
check_response "$result" "\"software_version_string\": \"1\""
check_response "$result" "\"cd_version_number\": 1"
check_response "$result" "\"software_version_valid\": true"
check_response "$result" "\"min_applicable_software_version\": 1"
check_response "$result" "\"max_applicable_software_version\": 10"

test_divider

echo "Query Vendor Models"
result=$(execute_with_retry "dcld query model vendor-models --vid=$vid")
echo "$result"
check_response "$result" "\"pid\": $pid"

test_divider

echo "Query model versions"
result=$(execute_with_retry "dcld query model all-model-versions --vid=$vid --pid=$pid")
echo "$result"
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"software_versions\""
check_response "$result" "$sv"

test_divider

# 5. check non-existent values via light client when entry added

echo "check non-existent values via light client when entry added"

test_divider

vid=$RANDOM
pid=$RANDOM
sv=$RANDOM

echo "Query non existent model"
result=$(execute_with_retry "dcld query model get-model --vid=$vid --pid=$pid")
echo "$result"
check_response "$result" "Not Found"

test_divider

echo "Query non existent model version"
result=$(execute_with_retry "dcld query model get-model-version --vid=$vid --pid=$pid --softwareVersion=$sv")
echo "$result"
check_response "$result" "Not Found"

test_divider

echo "Query non existent Vendor Models"
result=$(execute_with_retry "dcld query model vendor-models --vid=$vid")
echo "$result"
check_response "$result" "Not Found"

test_divider

echo "Query non existent model versions"
result=$(execute_with_retry "dcld query model all-model-versions --vid=$vid --pid=$pid")
echo "$result"
check_response "$result" "Not Found"

test_divider

# 6. try to write via light client proxy

echo "try to write via light client proxy"

test_divider

echo "Add Model with VID: $vid PID: $pid"
result=$(echo "test1234" | dcld tx model add-model --vid=$vid --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel="$productLabel" --partNumber=1 --commissioningCustomFlow=0 --from=$vendor_account --yes)
echo "$result"
check_response "$result" "Write requests don't work with a Light Client Proxy"


test_divider

echo "Create a Device Model Version with VID: $vid PID: $pid SV: $sv"
result=$(echo 'test1234' | dcld tx model add-model-version --cdVersionNumber=1 --maxApplicableSoftwareVersion=10 --minApplicableSoftwareVersion=1 --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=1 --from=$vendor_account --yes)
echo "$result"
check_response "$result" "Write requests don't work with a Light Client Proxy"

test_divider