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

# Preparation

# constants
trustee_account_1="jack"
trustee_account_2="alice"
vendor_account="van"

plan_name="test-upgrade"

original_model_module_version=1
updated_model_module_version=2

vid=1
pid=2
original_product_label="TestProductLabel"
updated_product_label="TestProductLabel_UPDATED"

echo "Create Vendor account $vendor_account"
create_new_vendor_account $vendor_account $vid

test_divider

# Body

get_height current_height
echo "Current height is $current_height"

plan_height=$(expr $current_height \* 4)

test_divider

echo "Propose upgrade $plan_name at height $plan_height"
result=$(echo $passphrase | dcld tx dclupgrade propose-upgrade --name $plan_name --upgrade-height $plan_height --from $trustee_account_1 --yes)
echo "$result"
check_response "$result" "\"code\": 0"

test_divider

echo "Approve upgrade $plan_name"
result=$(echo $passphrase | dcld tx dclupgrade approve-upgrade $plan_name --from $trustee_account_2 --yes)
echo "$result"
check_response "$result" "\"code\": 0"

test_divider

echo "Verify that upgrade $plan_name has been scheduled at height $plan_height"
result=$(dcld query upgrade plan)
echo "$result"
check_response "$result" "\"name\": \"$plan_name\""
check_response "$result" "\"height\": \"$plan_height\""

test_divider

echo "Add model with vid=$vid, pid=$pid and productLabel=$original_product_label"
result=$(echo $passphrase | dcld tx model add-model --vid=$vid --pid=$pid --deviceTypeID=1 --productName=TestProduct \
--productLabel=$original_product_label --partNumber=1 --commissioningCustomFlow=0 --from=$vendor_account --yes)
echo "$result"
check_response "$result" "\"code\": 0"

test_divider

echo "Verify that model with vid=$vid, pid=$pid has been added with original product label $original_product_label"
result=$(dcld query model get-model --vid=$vid --pid=$pid)
echo "$result"
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"product_label\": \"$original_product_label\""

test_divider

echo "Verify that version of model module is currently $original_model_module_version"
result=$(dcld query upgrade module_versions model)
echo "$result"
check_response "$result" "\"version\": \"$original_model_module_version\""

test_divider

echo "Wait for block height to become greater than upgrade $plan_name plan height"
wait_for_height $(expr $plan_height + 1) 300 outage-safe

test_divider

echo "Verify that no upgrade has been scheduled anymore"
result=$(dcld query upgrade plan 2>&1) || true
check_response_and_report "$result" "no upgrade scheduled" raw

test_divider

echo "Verify that version of model module is currently $updated_model_module_version"
result=$(dcld query upgrade module_versions model)
echo "$result"
check_response "$result" "\"version\": \"$updated_model_module_version\""

test_divider

echo "Verify that model with vid=$vid, pid=$pid has been updated so that product label is now $updated_product_label"
result=$(dcld query model get-model --vid=$vid --pid=$pid)
echo "$result"
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"product_label\": \"$updated_product_label\""

echo "PASSED"
