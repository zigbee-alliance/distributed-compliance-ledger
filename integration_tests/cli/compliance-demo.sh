#!/bin/bash
set -e
source integration_tests/cli/common.sh

echo "Assign Vendor role to Jack"
result=$(echo "test1234" | zblcli tx authz assign-role $(zblcli keys show jack -a) "Vendor" --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

sleep 6

vid=$RANDOM
pid=$RANDOM
echo "Add Model with VID: $vid PID: $pid"
result=$(echo "test1234" | zblcli tx modelinfo add-model $vid $pid "Device #1" "Device Description" "SKU12FS" "1.0" "2.0" true --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

sleep 6

echo "Assign TestHouse role to Jack"
result=$(echo "test1234" | zblcli tx authz assign-role $(zblcli keys show jack -a) "TestHouse" --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

sleep 6

echo "Add Testing Result for Model VID: $vid PID: $pid"
testing_result="http://first.place.com"
result=$(echo "test1234" | zblcli tx compliancetest add-test-result $vid $pid "$testing_result" --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

sleep 6

echo "Assign ZBCertificationCenter role to Jack"
result=$(echo "test1234" | zblcli tx authz assign-role $(zblcli keys show jack -a) "ZBCertificationCenter" --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

sleep 6

echo "Certify Model with VID: $vid PID: $pid"
certification_date="2020-01-01T00:00:00Z"
result=$(echo "test1234" | zblcli tx compliance certify-model $vid $pid "$certification_date" --certification-type "zb" --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

sleep 6

echo "Get Certified Model with VID: ${vid} PID: ${pid}"
result=$(zblcli query compliance certified-model $vid $pid)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"certification_type\": \"zb\""
echo "$result"

sleep 6

echo "Get All Certified Models"
result=$(zblcli query compliance all-certified-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"certification_type\": \"zb\""
echo "$result"