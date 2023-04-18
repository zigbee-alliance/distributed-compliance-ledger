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
pid=$RANDOM
vendor_account=vendor_account_$vid
echo "Create Vendor account - $vendor_account"
create_new_vendor_account $vendor_account $vid

test_divider

# Body

echo "Query non existent model"
result=$(dcld query model get-model --vid=$vid --pid=$pid)
check_response "$result" "Not Found"
echo "$result"

test_divider

echo "Query non existent Vendor Models"
result=$(dcld query model vendor-models --vid=$vid)
check_response "$result" "Not Found"
echo "$result"

test_divider

echo "Request all models must be empty"
result=$(dcld query model all-models)
check_response "$result" "\[\]"
echo "$result"

test_divider

productLabel="Device #1"
echo "Add Model with VID: $vid PID: $pid"
result=$(echo "test1234" | dcld tx model add-model --vid=$vid --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel="$productLabel" --partNumber=1 --commissioningCustomFlow=0 --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "Get Model with VID: $vid PID: $pid"
result=$(dcld query model get-model --vid=$vid --pid=$pid)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"productLabel\": \"$productLabel\""
echo "$result"

test_divider

sv=1
cd_version_num=10
echo "Create Model Versions with VID: $vid PID: $pid SoftwareVersion: $sv"
result=$(echo "test1234" | dcld tx model add-model-version --vid=$vid --pid=$pid --softwareVersion=$sv --minApplicableSoftwareVersion=1 --maxApplicableSoftwareVersion=15 --softwareVersionString=$sv --cdVersionNumber=$cd_version_num --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "Get all models"
result=$(dcld query model all-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
echo "$result"

test_divider

echo "Get Vendor Models with VID: ${vid}"
result=$(dcld query model vendor-models --vid=$vid)
check_response "$result" "\"pid\": $pid"
echo "$result"

test_divider

echo "Update Model with VID: ${vid} PID: ${pid} with new description"
description="New Device Description"
result=$(echo "test1234" | dcld tx model update-model --vid=$vid --pid=$pid --from $vendor_account --yes --productLabel "$description")
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "Get Model with VID: ${vid} PID: ${pid}"
result=$(dcld query model get-model --vid=$vid --pid=$pid)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"productLabel\": \"$description\""
echo "$result"

test_divider

echo "Update Model with VID: ${vid} PID: ${pid} modifying supportURL"
supportURL="https://newsupporturl.test"
result=$(echo "test1234" | dcld tx model update-model --vid=$vid --pid=$pid --from $vendor_account --yes --supportURL "$supportURL")
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "Get Model with VID: ${vid} PID: ${pid}"
result=$(dcld query model get-model --vid=$vid --pid=$pid)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"supportUrl\": \"$supportURL\""
echo "$result"

test_divider

echo "Delete Model with VID: ${vid} PID: ${pid}"
result=$(dcld tx model delete-model --vid=$vid --pid=$pid --from=$vendor_account --yes)
echo "$result"

test_divider

echo "Query non existent model"
result=$(dcld query model get-model --vid=$vid --pid=$pid)
check_response "$result" "Not Found"
echo "$result"

test_divider

echo "Query model versions for deleted model"
result=$(dcld query model model-version --vid=$vid --pid=$pid --softwareVersion=$sv)
check_response "$result" "Not Found"
echo "$result"