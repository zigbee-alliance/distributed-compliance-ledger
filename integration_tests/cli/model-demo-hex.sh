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

# Check add model with fieds VID/PID in hex format and 
# get model with fields VID/PID in hex format

# Preperation of Actors

vid_in_hex_format=0xA13
pid_in_hex_format=0xA11
vid=2579
pid=2577

vendor_account=vendor_account_$vid_in_hex_format
echo "Create Vendor account - $vendor_account"
create_new_vendor_account $vendor_account $vid_in_hex_format

test_divider

# Body

echo "Query non existent model"
result=$(dcld query model get-model --vid=$vid_in_hex_format --pid=$pid_in_hex_format)
check_response "$result" "Not Found"
echo "$result"

test_divider

echo "Query non existent Vendor Models"
result=$(dcld query model vendor-models --vid=$vid_in_hex_format)
check_response "$result" "Not Found"
echo "$result"

test_divider

echo "Request all models must be empty"
result=$(dcld query model all-models)
check_response "$result" "\[\]"
echo "$result"

test_divider

productLabel="Device #1"
echo "Add Model with VID: $vid_in_hex_format PID: $pid_in_hex_format"
result=$(echo "test1234" | dcld tx model add-model --vid=$vid_in_hex_format --pid=$pid_in_hex_format --deviceTypeID=1 --productName=TestProduct --productLabel="$productLabel" --partNumber=1 --commissioningCustomFlow=0 --enhancedSetupFlowOptions=1 --from=$vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "Get Model with VID: $vid_in_hex_format PID: $pid_in_hex_format"
result=$(dcld query model get-model --vid=$vid_in_hex_format --pid=$pid_in_hex_format)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"productLabel\": \"$productLabel\""
echo "$result"

test_divider

echo "Get all models"
result=$(dcld query model all-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
echo "$result"

test_divider

echo "Get Vendor Models with VID: ${vid_in_hex_format}"
result=$(dcld query model vendor-models --vid=$vid_in_hex_format)
check_response "$result" "\"pid\": $pid"
echo "$result"

test_divider

echo "Update Model with VID: ${vid_in_hex_format} PID: ${pid_in_hex_format} with new description"
description="New Device Description"
result=$(echo "test1234" | dcld tx model update-model --vid=$vid_in_hex_format --pid=$pid_in_hex_format --enhancedSetupFlowOptions=2 --from $vendor_account --yes --productLabel "$description")
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "Get Model with VID: ${vid_in_hex_format} PID: ${pid_in_hex_format}"
result=$(dcld query model get-model --vid=$vid_in_hex_format --pid=$pid_in_hex_format)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"productLabel\": \"$description\""
echo "$result"

test_divider

echo "Update Model with VID: ${vid_in_hex_format} PID: ${pid_in_hex_format} modifying supportURL"
supportURL="https://newsupporturl.test"
result=$(echo "test1234" | dcld tx model update-model --vid=$vid_in_hex_format --pid=$pid_in_hex_format --enhancedSetupFlowOptions=2 --from $vendor_account --yes --supportURL "$supportURL")
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "Get Model with VID: ${vid_in_hex_format} PID: ${pid_in_hex_format}"
result=$(dcld query model get-model --vid=$vid_in_hex_format --pid=$pid_in_hex_format)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"supportUrl\": \"$supportURL\""
echo "$result"

test_divider

echo "Delete Model with VID: ${vid_in_hex_format} PID: ${pid_in_hex_format}"
result=$(dcld tx model delete-model --vid=$vid_in_hex_format --pid=$pid_in_hex_format --from=$vendor_account --yes)
result=$(get_txn_result "$result")
echo "$result"

test_divider

echo "Query non existent model"
result=$(dcld query model get-model --vid=$vid_in_hex_format --pid=$pid_in_hex_format)
check_response "$result" "Not Found"
echo "$result"

test_divider