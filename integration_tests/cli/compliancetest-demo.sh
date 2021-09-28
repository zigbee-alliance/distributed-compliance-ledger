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

vid=$RANDOM
vendor_account=vendor_account_$vid
echo "Create Vendor account - $vendor_account"
create_new_vendor_account $vendor_account $vid

echo "Create TestHouse account"
create_new_account test_house_account "TestHouse"

echo "Create second TestHouse account"
create_new_account second_test_house_account "TestHouse"

# Body

pid=$RANDOM
sv=$RANDOM
echo "Add Model and a New Model Version with VID: $vid PID: $pid SV: $sv"
create_model_and_version $vid $pid $sv $vendor_account

test_divider

echo "Add Testing Result for Model VID: $vid PID: $pid SV: $sv"
testing_result="http://first.place.com"
test_date="2020-01-01T00:00:00Z"
result=$(echo "test1234" | dclcli tx compliancetest add-test-result --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString="1.0" --test-result="$testing_result" --test-date="$test_date" --from $test_house_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Add Second Testing Result for Model VID: $vid PID: $pid SV: $sv"
second_testing_result="http://second.place.com"
second_test_date="2020-04-04T10:00:00Z"
result=$(echo "test1234" | dclcli tx compliancetest add-test-result --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString="1.0" --test-result="$second_testing_result" --test-date=$second_test_date --from $second_test_house_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get Testing Result for Model with VID: ${vid} PID: ${pid} SV: $sv"
result=$(dclcli query compliancetest test-result --vid=$vid --pid=$pid --softwareVersion=$sv)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"test_result\": \"$testing_result\""
check_response "$result" "\"test_date\": \"$test_date\""
check_response "$result" "\"test_result\": \"$second_testing_result\""
check_response "$result" "\"test_date\": \"$second_test_date\""
echo "$result"

test_divider

pid=$RANDOM
echo "Add Model and a New Model Version with VID: $vid PID: $pid SV: $sv"
create_model_and_version $vid $pid $sv $vendor_account

test_divider

echo "Add Testing Result for Model VID: $vid PID: $pid SV: $sv"
testing_result="blob string"
test_date="2020-11-24T10:00:00Z"
result=$(echo "test1234" | dclcli tx compliancetest add-test-result --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString="1.0" --test-result="$testing_result" --test-date="$test_date" --from $test_house_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

test_divider

echo "Get Testing Result for Model with VID: ${vid} PID: ${pid} SV:$sv"
result=$(dclcli query compliancetest test-result --vid=$vid --pid=$pid --softwareVersion=$sv)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"softwareVersion\": $sv"
check_response "$result" "\"test_result\": \"$testing_result\""
check_response "$result" "\"test_date\": \"$test_date\""
echo "$result"
