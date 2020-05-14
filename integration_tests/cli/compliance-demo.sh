#!/bin/bash
set -e
source integration_tests/cli/common.sh

# Preparation of Actors

echo "Assign Vendor role to Jack"
result=$(echo "test1234" | zblcli tx authz assign-role --address=$(zblcli keys show jack -a) --role="Vendor" --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

test_house_account=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 6 | head -n 1)
echo "Create TestHouse account with address: $test_house_account"
create_account_with_name $test_house_account

test_house_address=$(zblcli keys show "$test_house_account" -a)

result=$(echo "test1234" | zblcli tx authz assign-role --address=$test_house_address --role="TestHouse" --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

zb_account=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 6 | head -n 1)
echo "Create ZBCertificationCenter account with address: $zb_account"
create_account_with_name $zb_account

zb_address=$(zblcli keys show "$zb_account" -a)

result=$(echo "test1234" | zblcli tx authz assign-role --address=$zb_address --role="ZBCertificationCenter" --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

# Body

vid=$RANDOM
pid=$RANDOM
echo "Add Model with VID: $vid PID: $pid"
result=$(echo "test1234" | zblcli tx modelinfo add-model --vid=$vid --pid=$pid --name="Device #1" --description="Device Description" --sku="SKU12FS" --firmware-version="1.0" --hardware-version="2.0" --tis-or-trp-testing-completed=true --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Add Testing Result for Model VID: $vid PID: $pid"
testing_result="http://first.place.com"
test_date="2020-11-24T10:00:00Z"
result=$(echo "test1234" | zblcli tx compliancetest add-test-result --vid=$vid --pid=$pid --test-result="$testing_result" --test-date="$test_date" --from $test_house_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get Certified Model with VID: ${vid} PID: ${pid} before compliance record was created"
result=$(zblcli query compliance certified-model --vid=$vid --pid=$pid --certification-type="zb")
check_response "$result" "\"value\": false"
echo "$result"

echo "Get Revoked Model with VID: ${vid} PID: ${pid} before compliance record was created"
result=$(zblcli query compliance revoked-model --vid=$vid --pid=$pid --certification-type="zb")
check_response "$result" "\"value\": false"
echo "$result"

echo "Certify Model with VID: $vid PID: $pid"
certification_date="2020-01-01T00:00:00Z"
certification_type="zb"
result=$(echo "test1234" | zblcli tx compliance certify-model --vid=$vid --pid=$pid --certification-type="$certification_type" --certification-date="$certification_date" --from $zb_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get Certified Model with VID: ${vid} PID: ${pid}"
result=$(zblcli query compliance certified-model --vid=$vid --pid=$pid --certification-type=$certification_type)
check_response "$result" "\"value\": true"
echo "$result"

echo "Get Compliance Info for Model with VID: ${vid} PID: ${pid}"
result=$(zblcli query compliance compliance-info --vid=$vid --pid=$pid --certification-type=$certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"state\": \"certified\""
check_response "$result" "\"date\": \"$certification_date\""
check_response "$result" "\"certification_type\": \"$certification_type\""
echo "$result"

echo "Get All Certified Models"
result=$(zblcli query compliance all-certified-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"certification_type\": \"$certification_type\""
echo "$result"

echo "Get All Compliance Info Recordss"
result=$(zblcli query compliance all-compliance-info-records)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"date\": \"$certification_date\""
echo "$result"

echo "Revoke Certification for Model with VID: $vid PID: $pid"
revocation_date="2020-02-02T02:20:20Z"
revocation_reason="some reason"
result=$(echo "test1234" | zblcli tx compliance revoke-model --vid=$vid --pid=$pid --certification-type="$certification_type" --revocation-date="$revocation_date" --reason "$revocation_reason" --from $zb_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get Compliance Info for Model with VID: ${vid} PID: ${pid}"
result=$(zblcli query compliance compliance-info --vid=$vid --pid=$pid --certification-type=$certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"state\": \"revoked\""
check_response "$result" "\"date\": \"$revocation_date\""
check_response "$result" "\"reason\": \"$revocation_reason\""
check_response "$result" "\"certification_type\": \"$certification_type\""
check_response "$result" "\"history\""
echo "$result"

echo "Get Revoked Model with VID: ${vid} PID: ${pid}"
result=$(zblcli query compliance revoked-model --vid=$vid --pid=$pid --certification-type=$certification_type)
check_response "$result" "\"value\": true"
echo "$result"

echo "Get Certified Model with VID: ${vid} PID: ${pid}"
result=$(zblcli query compliance certified-model --vid=$vid --pid=$pid --certification-type=$certification_type)
check_response "$result" "\"value\": false"
echo "$result"

echo "Get All Revoked Models"
result=$(zblcli query compliance all-revoked-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"certification_type\": \"$certification_type\""
echo "$result"

echo "Again Certify Model with VID: $vid PID: $pid"
certification_date="2020-03-03T00:00:00Z"
result=$(echo "test1234" | zblcli tx compliance certify-model --vid=$vid --pid=$pid --certification-type="$certification_type" --certification-date="$certification_date" --from $zb_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get Compliance Info for Model with VID: ${vid} PID: ${pid}"
result=$(zblcli query compliance compliance-info --vid=$vid --pid=$pid --certification-type=$certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"state\": \"certified\""
check_response "$result" "\"date\": \"$certification_date\""
check_response "$result" "\"certification_type\": \"zb\""
echo "$result"

echo "Get Certified Model with VID: ${vid} PID: ${pid}"
result=$(zblcli query compliance certified-model --vid=$vid --pid=$pid --certification-type=$certification_type)
check_response "$result" "\"value\": true"
echo "$result"

echo "Get Revoked Model with VID: ${vid} PID: ${pid}"
result=$(zblcli query compliance revoked-model --vid=$vid --pid=$pid --certification-type=$certification_type)
check_response "$result" "\"value\": false"
echo "$result"

echo "Get All Compliance Infos"
result=$(zblcli query compliance all-compliance-info-records)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"state\": \"certified\""
echo "$result"