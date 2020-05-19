#!/bin/bash
set -e
source integration_tests/cli/common.sh

# Preparation of Actors

echo "Create Vendor account"
create_new_account vendor_account "Vendor"

echo "Create TestHouse account"
create_new_account test_house_account "TestHouse"

echo "Create second TestHouse account"
create_new_account second_test_house_account "TestHouse"

# Body

vid=$RANDOM
pid=$RANDOM
echo "Add Model with VID: $vid PID: $pid"
result=$(echo "test1234" | zblcli tx modelinfo add-model --vid=$vid --pid=$pid --name="Device #1" --description="Device Description" --sku="SKU12FS" --firmware-version="1.0" --hardware-version="2.0" --tis-or-trp-testing-completed=true --from $vendor_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Add Testing Result for Model VID: $vid PID: $pid"
testing_result="http://first.place.com"
test_date="2020-01-01T00:00:00Z"
result=$(echo "test1234" | zblcli tx compliancetest add-test-result --vid=$vid --pid=$pid --test-result="$testing_result" --test-date="$test_date" --from $test_house_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Add Second Testing Result for Model VID: $vid PID: $pid"
second_testing_result="http://second.place.com"
second_test_date="2020-04-04T10:00:00Z"
result=$(echo "test1234" | zblcli tx compliancetest add-test-result --vid=$vid --pid=$pid --test-result="$second_testing_result" --test-date=$second_test_date --from $second_test_house_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get Testing Result for Model with VID: ${vid} PID: ${pid}"
result=$(zblcli query compliancetest test-result --vid=$vid --pid=$pid)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"test_result\": \"$testing_result\""
check_response "$result" "\"test_date\": \"$test_date\""
check_response "$result" "\"test_result\": \"$second_testing_result\""
check_response "$result" "\"test_date\": \"$second_test_date\""
echo "$result"

vid=$RANDOM
pid=$RANDOM
echo "Add Model with VID: $vid PID: $pid"
result=$(echo "test1234" | zblcli tx modelinfo add-model --vid=$vid --pid=$pid --name="Device #1" --description="Device Description" --sku="SKU12FS" --firmware-version="1.0" --hardware-version="2.0" --tis-or-trp-testing-completed=true --from $vendor_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Add Testing Result for Model VID: $vid PID: $pid"
testing_result="blob string"
test_date="2020-11-24T10:00:00Z"
result=$(echo "test1234" | zblcli tx compliancetest add-test-result --vid=$vid --pid=$pid --test-result="$testing_result" --test-date="$test_date" --from $test_house_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get Testing Result for Model with VID: ${vid} PID: ${pid}"
result=$(zblcli query compliancetest test-result --vid=$vid --pid=$pid)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"test_result\": \"$testing_result\""
check_response "$result" "\"test_date\": \"$test_date\""
echo "$result"
