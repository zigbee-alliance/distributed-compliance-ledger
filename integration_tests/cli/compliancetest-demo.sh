#!/bin/bash
# Copyright 2020 DSR Corporation
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -euo pipefail
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
result=$(echo "test1234" | dclcli tx modelinfo add-model --vid=$vid --pid=$pid --name="Device #1" --description="Device Description" --sku="SKU12FS" --software-version="10123" --software-version-string="1.0b123"  --hardware-version="5123" --hardware-version-string="5.1.23"  --cd-version-number="32" --from $vendor_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Add Testing Result for Model VID: $vid PID: $pid"
testing_result="http://first.place.com"
test_date="2020-01-01T00:00:00Z"
result=$(echo "test1234" | dclcli tx compliancetest add-test-result --vid=$vid --pid=$pid --test-result="$testing_result" --test-date="$test_date" --from $test_house_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Add Second Testing Result for Model VID: $vid PID: $pid"
second_testing_result="http://second.place.com"
second_test_date="2020-04-04T10:00:00Z"
result=$(echo "test1234" | dclcli tx compliancetest add-test-result --vid=$vid --pid=$pid --test-result="$second_testing_result" --test-date=$second_test_date --from $second_test_house_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get Testing Result for Model with VID: ${vid} PID: ${pid}"
result=$(dclcli query compliancetest test-result --vid=$vid --pid=$pid)
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
result=$(echo "test1234" | dclcli tx modelinfo add-model --vid=$vid --pid=$pid --name="Device #1" --description="Device Description" --sku="SKU12FS" --software-version="10123" --software-version-string="1.0b123"  --hardware-version="5123" --hardware-version-string="5.1.23"  --cd-version-number="32" --from $vendor_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Add Testing Result for Model VID: $vid PID: $pid"
testing_result="blob string"
test_date="2020-11-24T10:00:00Z"
result=$(echo "test1234" | dclcli tx compliancetest add-test-result --vid=$vid --pid=$pid --test-result="$testing_result" --test-date="$test_date" --from $test_house_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get Testing Result for Model with VID: ${vid} PID: ${pid}"
result=$(dclcli query compliancetest test-result --vid=$vid --pid=$pid)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"test_result\": \"$testing_result\""
check_response "$result" "\"test_date\": \"$test_date\""
echo "$result"
