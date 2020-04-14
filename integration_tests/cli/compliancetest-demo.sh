#!/bin/bash
set -e
source integration_tests/cli/common.sh

echo "Assign Vendor role to Jack"
result=$(echo "test1234" | zblcli tx authz assign-role $(zblcli keys show jack -a) "Vendor" --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

sleep 5

vid=$RANDOM
pid=$RANDOM
echo "Add Model with VID: $vid PID: $pid"
result=$(echo "test1234" | zblcli tx modelinfo add-model $vid $pid "Device #1" "Device Description" "SKU12FS" "1.0" "2.0" true --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

sleep 5

echo "Assign TestHouse role to Jack"
result=$(echo "test1234" | zblcli tx authz assign-role $(zblcli keys show jack -a) "TestHouse" --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

sleep 5

echo "Add Testing Result for Model VID: $vid PID: $pid"
testing_result="http://place.com"
result=$(echo "test1234" | zblcli tx compliancetest add-test-result $vid $pid "$testing_result" --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

sleep 10

echo "Get Testing Result for Model with VID: ${vid} PID: ${pid}"
result=$(zblcli query compliancetest test-result $vid $pid)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"test_result\": \"$testing_result\""
echo "$result"

sleep 5

vid=$RANDOM
pid=$RANDOM
echo "Add Model with VID: $vid PID: $pid"
result=$(echo "test1234" | zblcli tx modelinfo add-model $vid $pid "Device #1" "Device Description" "SKU12FS" "1.0" "2.0" true --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

sleep 5

echo "Add Testing Result for Model VID: $vid PID: $pid"
testing_result="blob string"
result=$(echo "test1234" | zblcli tx compliancetest add-test-result $vid $pid "$testing_result" --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

sleep 10

echo "Get Testing Result for Model with VID: ${vid} PID: ${pid}"
result=$(zblcli query compliancetest test-result $vid $pid)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"test_result\": \"$testing_result\""
echo "$result"