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
svs=$RANDOM

echo "Query non existent complianceinfo"
result=$(execute_with_retry "dcld query compliance compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType="zigbee"")
echo "$result"
check_response "$result" "Not Found"

test_divider

echo "Query non existent certified"
result=$(execute_with_retry "dcld query compliance certified-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType="zigbee"")
echo "$result"
check_response "$result" "Not Found"

test_divider

echo "Query non existent revoked"
result=$(execute_with_retry "dcld query compliance revoked-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType="zigbee"")
echo "$result"
check_response "$result" "Not Found"

test_divider

echo "Query non existent provision"
result=$(execute_with_retry "dcld query compliance provisional-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType="zigbee"")
echo "$result"
check_response "$result" "Not Found"

test_divider

# 2. list queries should return a warning via light client

echo "list queries should return a warning via light client"

test_divider

echo "Request all compliance infos"
result=$(execute_with_retry "dcld query compliance all-compliance-info")
echo "$result"
check_response "$result" "List queries don't work with a Light Client Proxy"

test_divider

echo "Request all certified"
result=$(execute_with_retry "dcld query compliance all-certified-models")
echo "$result"
check_response "$result" "List queries don't work with a Light Client Proxy"

test_divider

echo "Request all revoked"
result=$(execute_with_retry "dcld query compliance all-revoked-models")
echo "$result"
check_response "$result" "List queries don't work with a Light Client Proxy"

test_divider

echo "Request all provisional"
result=$(execute_with_retry "dcld query compliance all-provisional-models")
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
certification_date="2020-01-01T00:00:01Z"

vendor_account=vendor_account_$vid
echo "Create Vendor account - $vendor_account"
create_new_vendor_account $vendor_account $vid

test_divider

echo "Create CertificationCenter account"
create_new_account zb_account "CertificationCenter"

test_divider

echo "Add Model and a New Model Version with VID: $vid PID: $pid SV: $sv"
create_model_and_version $vid $pid $sv $svs $vendor_account

test_divider

echo "Certify Model with VID: $vid PID: $pid  SV: ${sv} with zigbee certification"
result=$(echo "$passphrase" | dcld tx compliance certify-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --certificationType="zigbee" --certificationDate="$certification_date" --from $zb_account --yes)
echo "$result"
check_response "$result" "\"code\": 0"

test_divider

# 4. check existent values via light client

echo "check existent values via light client"

test_divider

# connect to light client proxy
dcld config node tcp://localhost:26620
sleep 5

echo "Query complianceinfo"
result=$(execute_with_retry "dcld query compliance compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType="zigbee"")
echo "$result"
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"software_version_certification_status\": 2"
check_response "$result" "\"date\": \"$certification_date\""
check_response "$result" "\"certification_type\": \"zigbee\""

test_divider

echo "Query certified"
result=$(execute_with_retry "dcld query compliance certified-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType="zigbee"")
echo "$result"
check_response "$result" "\"value\": true"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"vid\": $vid"

test_divider

echo "Query revoked"
result=$(execute_with_retry "dcld query compliance revoked-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType="zigbee"")
echo "$result"
check_response "$result" "\"value\": false"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"vid\": $vid"

test_divider

echo "Query provision"
result=$(execute_with_retry "dcld query compliance provisional-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType="zigbee"")
echo "$result"
check_response "$result" "\"value\": false"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"vid\": $vid"

test_divider


# 5. check non-existent values via light client when entry added

echo "check non-existent values via light client when entry added"

test_divider

vid=$RANDOM
pid=$RANDOM
sv=$RANDOM
svs=$RANDOM


echo "Query non existent complianceinfo"
result=$(execute_with_retry "dcld query compliance compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType="zigbee"")
echo "$result"
check_response "$result" "Not Found"

test_divider

echo "Query non existent certified"
result=$(execute_with_retry "dcld query compliance certified-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType="zigbee"")
echo "$result"
check_response "$result" "Not Found"

test_divider

echo "Query non existent revoked"
result=$(execute_with_retry "dcld query compliance revoked-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType="zigbee"")
echo "$result"
check_response "$result" "Not Found"

test_divider

echo "Query non existent provision"
result=$(execute_with_retry "dcld query compliance provisional-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType="zigbee"")
echo "$result"
check_response "$result" "Not Found"

test_divider

# 6. try to write via light client proxy

echo "try to write via light client proxy"

test_divider

echo "Add compliance info"
result=$(echo "$passphrase" | dcld tx compliance certify-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --certificationType="zigbee" --certificationDate="$certification_date" --from $zb_account --yes)
echo "$result"
check_response "$result" "Write requests don't work with a Light Client Proxy"


test_divider
