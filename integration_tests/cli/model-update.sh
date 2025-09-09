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

((vid_with_pids=vid + 1))
pid_ranges="$pid-$pid"
vendor_account_with_pids=vendor_account_$vid_with_pids
echo "Create Vendor account - $vid_with_pids with ProductIDs - $pid_ranges"
create_new_vendor_account $vendor_account_with_pids $vid_with_pids $pid_ranges

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
enhancedSetupFlowOptions_0=0
schema_version_0=0

echo "Add Model with VID: $vid PID: $pid"
result=$(echo "test1234" | dcld tx model add-model --vid=$vid --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel="$productLabel" --partNumber=1 --commissioningCustomFlow=0 --enhancedSetupFlowOptions=$enhancedSetupFlowOptions_0 --schemaVersion=$schema_version_0 --from=$vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

# check default values for commissioningModeInitialStepsHint and commissioningModeSecondaryStepsHint and icdUserActiveModeTriggerHint
echo "Get Model with VID: $vid PID: $pid"
result=$(dcld query model get-model --vid=$vid --pid=$pid)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"productLabel\": \"$productLabel\""
check_response "$result" "\"schemaVersion\": $schema_version_0"
check_response "$result" "\"commissioningModeInitialStepsHint\": 1"
check_response "$result" "\"commissioningModeSecondaryStepsHint\": 1"
check_response "$result" "\"icdUserActiveModeTriggerHint\": 1"
check_response "$result" "\"enhancedSetupFlowOptions\": $enhancedSetupFlowOptions_0"
echo "$result"

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

echo "Update Model with VID: ${vid} PID: ${pid} with new description, commissioningModeInitialStepsHint and commissioningModeSecondaryStepsHint"
description="New Device Description"
commissioningModeInitialStepsHint=3
commissioningModeSecondaryStepsHint=4
icdUserActiveModeTriggerHint=5
enhancedSetupFlowOptions_2=2
result=$(echo "test1234" | dcld tx model update-model --vid=$vid --pid=$pid --from $vendor_account --yes --productLabel "$description" --schemaVersion=$schema_version_0 \
  --commissioningModeInitialStepsHint="$commissioningModeInitialStepsHint" --commissioningModeSecondaryStepsHint="$commissioningModeSecondaryStepsHint" \
  --icdUserActiveModeTriggerHint="$icdUserActiveModeTriggerHint" --enhancedSetupFlowOptions=$enhancedSetupFlowOptions_2)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

# check updated values for commissioningModeInitialStepsHint == 3 and commissioningModeSecondaryStepsHint == 4
echo "Get Model with VID: ${vid} PID: ${pid}"
result=$(dcld query model get-model --vid=$vid --pid=$pid)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"productLabel\": \"$description\""
check_response "$result" "\"schemaVersion\": $schema_version_0"
check_response "$result" "\"commissioningModeInitialStepsHint\": $commissioningModeInitialStepsHint"
check_response "$result" "\"commissioningModeSecondaryStepsHint\": $commissioningModeSecondaryStepsHint"
check_response "$result" "\"icdUserActiveModeTriggerHint\": $icdUserActiveModeTriggerHint"
check_response "$result" "\"enhancedSetupFlowOptions\": $enhancedSetupFlowOptions_2"
echo "$result"

test_divider

echo "Update Model with VID: ${vid} PID: ${pid} with new description only"
description="New Device Description 2"
result=$(echo "test1234" | dcld tx model update-model --vid=$vid --pid=$pid --from $vendor_account --yes --productLabel "$description" --schemaVersion=$schema_version_0 --enhancedSetupFlowOptions=$enhancedSetupFlowOptions_2)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

# check non-updated values for commissioningModeInitialStepsHint and commissioningModeSecondaryStepsHint and icdUserActiveModeTriggerHint
# (because the values have not been set)
echo "Get Model with VID: ${vid} PID: ${pid}"
result=$(dcld query model get-model --vid=$vid --pid=$pid)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"productLabel\": \"$description\""
check_response "$result" "\"schemaVersion\": $schema_version_0"
check_response "$result" "\"commissioningModeInitialStepsHint\": $commissioningModeInitialStepsHint"
check_response "$result" "\"commissioningModeSecondaryStepsHint\": $commissioningModeSecondaryStepsHint"
check_response "$result" "\"icdUserActiveModeTriggerHint\": $icdUserActiveModeTriggerHint"
check_response "$result" "\"enhancedSetupFlowOptions\": $enhancedSetupFlowOptions_2"
echo "$result"

test_divider

echo "Update Model with VID: ${vid} PID: ${pid} with new description, commissioningModeInitialStepsHint, commissioningModeSecondaryStepsHint, and icdUserActiveModeTriggerHint"
description="New Device Description 3"
result=$(echo "test1234" | dcld tx model update-model --vid=$vid --pid=$pid --from $vendor_account --yes --productLabel "$description" --schemaVersion=$schema_version_0 \
  --commissioningModeInitialStepsHint=0 --commissioningModeSecondaryStepsHint=0 --icdUserActiveModeTriggerHint=0 --enhancedSetupFlowOptions=$enhancedSetupFlowOptions_2)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

# check non-updated values for commissioningModeInitialStepsHint, commissioningModeSecondaryStepsHint, and icdUserActiveModeTriggerHint
# (because the values were set to 0)
echo "Get Model with VID: ${vid} PID: ${pid}"
result=$(dcld query model get-model --vid=$vid --pid=$pid)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"productLabel\": \"$description\""
check_response "$result" "\"schemaVersion\": $schema_version_0"
check_response "$result" "\"commissioningModeInitialStepsHint\": $commissioningModeInitialStepsHint"
check_response "$result" "\"commissioningModeSecondaryStepsHint\": $commissioningModeSecondaryStepsHint"
check_response "$result" "\"icdUserActiveModeTriggerHint\": $icdUserActiveModeTriggerHint"
check_response "$result" "\"enhancedSetupFlowOptions\": $enhancedSetupFlowOptions_2"
echo "$result"
