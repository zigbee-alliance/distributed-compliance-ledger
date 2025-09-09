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

productLabel="Device #1"
enhancedSetupFlowOptions_1=1
enhancedSetupFlowTCUrl="https://example.org/file.txt"
enhancedSetupFlowTCRevision=1
enhancedSetupFlowTCDigest="MWRjNGE0NDA0MWRjYWYxMTU0NWI3NTQzZGZlOTQyZjQ3NDJmNTY4YmU2OGZlZTI3NTQ0MWIwOTJiYjYwZGVlZA=="
enhancedSetupFlowTCFileSize=1024
maintenanceUrl="https://example.org"
commissioningFallbackUrl="https://url.commissioningfallbackurl.dclmodel"
discoveryCapabilitiesBitmask=1
echo "Add Model with VID: $vid_with_pids PID: $pid"
result=$(echo "test1234" | dcld tx model add-model --vid=$vid_with_pids --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel="$productLabel" --partNumber=1 --commissioningCustomFlow=0 --enhancedSetupFlowOptions=$enhancedSetupFlowOptions_1 \
  --enhancedSetupFlowTCUrl=$enhancedSetupFlowTCUrl --enhancedSetupFlowTCRevision=$enhancedSetupFlowTCRevision --enhancedSetupFlowTCDigest=$enhancedSetupFlowTCDigest --enhancedSetupFlowTCFileSize=$enhancedSetupFlowTCFileSize --maintenanceUrl=$maintenanceUrl \
  --commissioningFallbackUrl=$commissioningFallbackUrl --discoveryCapabilitiesBitmask=$discoveryCapabilitiesBitmask --from=$vendor_account_with_pids --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "Get Model with VID: $vid PID: $pid"
result=$(dcld query model get-model --vid=$vid --pid=$pid)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"productLabel\": \"$productLabel\""
check_response "$result" "\"schemaVersion\": $schema_version_0"
check_response "$result" "\"enhancedSetupFlowOptions\": $enhancedSetupFlowOptions_0"
echo "$result"

echo "Get Model with VID: $vid_with_pids PID: $pid"
result=$(dcld query model get-model --vid=$vid_with_pids --pid=$pid)
check_response "$result" "\"vid\": $vid_with_pids"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"productLabel\": \"$productLabel\""
check_response "$result" "\"schemaVersion\": $schema_version_0"
check_response "$result" "\"enhancedSetupFlowOptions\": $enhancedSetupFlowOptions_1"
check_response "$result" "\"enhancedSetupFlowTCUrl\": \"$enhancedSetupFlowTCUrl\""
check_response "$result" "\"enhancedSetupFlowTCRevision\": $enhancedSetupFlowTCRevision"
check_response "$result" "\"enhancedSetupFlowTCDigest\": \"$enhancedSetupFlowTCDigest\""
check_response "$result" "\"enhancedSetupFlowTCFileSize\": $enhancedSetupFlowTCFileSize"
check_response "$result" "\"maintenanceUrl\": \"$maintenanceUrl\""
check_response "$result" "\"commissioningFallbackUrl\": \"$commissioningFallbackUrl\""
check_response "$result" "\"discoveryCapabilitiesBitmask\": $discoveryCapabilitiesBitmask"
echo "$result"

test_divider

sv=1
cd_version_num=10
echo "Create Model Versions with VID: $vid PID: $pid SoftwareVersion: $sv"
result=$(echo "test1234" | dcld tx model add-model-version --vid=$vid --pid=$pid --softwareVersion=$sv --minApplicableSoftwareVersion=1 --maxApplicableSoftwareVersion=15 --softwareVersionString=$sv --cdVersionNumber=$cd_version_num --from=$vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "Create Model Versions with VID: $vid_with_pids PID: $pid SoftwareVersion: $sv"
result=$(echo "test1234" | dcld tx model add-model-version --vid=$vid_with_pids --pid=$pid --softwareVersion=$sv --minApplicableSoftwareVersion=1 --maxApplicableSoftwareVersion=15 --softwareVersionString=$sv --cdVersionNumber=$cd_version_num --from=$vendor_account_with_pids --yes)
result=$(get_txn_result "$result")
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

echo "Update Model with VID: ${vid} PID: ${pid} with new description, commissioningModeInitialStepsHint, and icdUserActiveModeTriggerHint"
description="New Device Description"
newCommissioningModeInitialStepsHint=8
newCommissioningModeSecondaryStepsHint=9
newIcdUserActiveModeTriggerHint=7
enhancedSetupFlowOptions_2=2
result=$(echo "test1234" | dcld tx model update-model --vid=$vid --pid=$pid --from $vendor_account --yes --productLabel "$description" --schemaVersion=$schema_version_0 \
  --commissioningModeInitialStepsHint="$newCommissioningModeInitialStepsHint" --commissioningModeSecondaryStepsHint="$newCommissioningModeSecondaryStepsHint" \
  --icdUserActiveModeTriggerHint="$newIcdUserActiveModeTriggerHint" --enhancedSetupFlowOptions=$enhancedSetupFlowOptions_2)
echo "Update Model with VID: ${vid} PID: ${pid} with new description, commissioningModeInitialStepsHint, and factoryResetStepsHint"
description="New Device Description"
newCommissioningModeInitialStepsHint=8
newCommissioningModeSecondaryStepsHint=9
newFactoryResetStepsHint=6
enhancedSetupFlowOptions_2=2
result=$(echo "test1234" | dcld tx model update-model --vid=$vid --pid=$pid --from $vendor_account --yes --productLabel "$description" --schemaVersion=$schema_version_0 \
  --commissioningModeInitialStepsHint="$newCommissioningModeInitialStepsHint" --commissioningModeSecondaryStepsHint="$newCommissioningModeSecondaryStepsHint" \
  --factoryResetStepsHint="$newFactoryResetStepsHint" --enhancedSetupFlowOptions=$enhancedSetupFlowOptions_2)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

newEnhancedSetupFlowTCUrl="https://example.org/file2.txt"
newEnhancedSetupFlowTCRevision=2
newEnhancedSetupFlowTCDigest="MWRjM2E0MTA0MWRjYWYxMTU0NWI3NTQzZGZlOTQyZjQ3NDJmNTY4YmU2OGZlZTI3NTQ0MWIwOTJiYjYxZGVlZA=="
newEnhancedSetupFlowTCFileSize=2048
newMaintenanceUrl="https://example2.org"
newCommissioningFallbackUrl="https://url.commissioningfallbackurl2.dclmodel"
echo "Update Model with VID: ${vid_with_pids} PID: ${pid} with new description, enhancedSetupFlowTCUrl, enhancedSetupFlowTCRevision, enhancedSetupFlowTCDigest, enhancedSetupFlowTCFileSize and maintenanceUrl"
result=$(echo "test1234" | dcld tx model update-model --vid=$vid_with_pids --pid=$pid --from $vendor_account_with_pids --yes --productLabel "$description" --enhancedSetupFlowOptions=$enhancedSetupFlowOptions_1 \
    --enhancedSetupFlowTCUrl=$newEnhancedSetupFlowTCUrl --enhancedSetupFlowTCRevision=$newEnhancedSetupFlowTCRevision --enhancedSetupFlowTCDigest=$newEnhancedSetupFlowTCDigest --enhancedSetupFlowTCFileSize=$newEnhancedSetupFlowTCFileSize --maintenanceUrl=$newMaintenanceUrl --commissioningFallbackUrl=$newCommissioningFallbackUrl --from=$vendor_account_with_pids --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "Get Model with VID: ${vid} PID: ${pid}"
result=$(dcld query model get-model --vid=$vid --pid=$pid)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"productLabel\": \"$description\""
check_response "$result" "\"schemaVersion\": $schema_version_0"
check_response "$result" "\"commissioningModeInitialStepsHint\": $newCommissioningModeInitialStepsHint"
check_response "$result" "\"commissioningModeSecondaryStepsHint\": $newCommissioningModeSecondaryStepsHint"
check_response "$result" "\"icdUserActiveModeTriggerHint\": $newIcdUserActiveModeTriggerHint"
check_response "$result" "\"factoryResetStepsHint\": $newFactoryResetStepsHint"
check_response "$result" "\"enhancedSetupFlowOptions\": $enhancedSetupFlowOptions_2"
echo "$result"

echo "Get Model with VID: $vid_with_pids PID: $pid"
result=$(dcld query model get-model --vid=$vid_with_pids --pid=$pid)
check_response "$result" "\"vid\": $vid_with_pids"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"schemaVersion\": $schema_version_0"
check_response "$result" "\"enhancedSetupFlowOptions\": $enhancedSetupFlowOptions_1"
check_response "$result" "\"enhancedSetupFlowTCUrl\": \"$newEnhancedSetupFlowTCUrl\""
check_response "$result" "\"enhancedSetupFlowTCRevision\": $newEnhancedSetupFlowTCRevision"
check_response "$result" "\"enhancedSetupFlowTCDigest\": \"$newEnhancedSetupFlowTCDigest\""
check_response "$result" "\"enhancedSetupFlowTCFileSize\": $newEnhancedSetupFlowTCFileSize"
check_response "$result" "\"maintenanceUrl\": \"$newMaintenanceUrl\""
check_response "$result" "\"commissioningFallbackUrl\": \"$newCommissioningFallbackUrl\""
echo "$result"

test_divider

echo "Update Model with VID: ${vid} PID: ${pid} modifying supportURL"
supportURL="https://newsupporturl.test"
result=$(echo "test1234" | dcld tx model update-model --vid=$vid --pid=$pid --from $vendor_account --yes --supportURL "$supportURL" --enhancedSetupFlowOptions=$enhancedSetupFlowOptions_0)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "Get Model with VID: ${vid} PID: ${pid}"
result=$(dcld query model get-model --vid=$vid --pid=$pid)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"supportUrl\": \"$supportURL\""
check_response "$result" "\"commissioningModeInitialStepsHint\": $newCommissioningModeInitialStepsHint"
check_response "$result" "\"commissioningModeSecondaryStepsHint\": $newCommissioningModeSecondaryStepsHint"
check_response "$result" "\"icdUserActiveModeTriggerHint\": $newIcdUserActiveModeTriggerHint"
check_response "$result" "\"factoryResetStepsHint\": $newFactoryResetStepsHint"
echo "$result"

test_divider

echo "Delete Model with VID: ${vid} PID: ${pid}"
result=$(dcld tx model delete-model --vid=$vid --pid=$pid --from=$vendor_account --yes)
result=$(get_txn_result "$result")
echo "$result"

test_divider

echo "Delete Model with VID: ${vid_with_pids} PID: ${pid}"
result=$(dcld tx model delete-model --vid=$vid_with_pids --pid=$pid --from=$vendor_account_with_pids --yes)
result=$(get_txn_result "$result")
echo "$result"

test_divider

echo "Query non existent model"
result=$(dcld query model get-model --vid=$vid --pid=$pid)
check_response "$result" "Not Found"
echo "$result"

test_divider

echo "Query non existent model"
result=$(dcld query model get-model --vid=$vid_with_pids --pid=$pid)
check_response "$result" "Not Found"
echo "$result"

test_divider

echo "Query model versions for deleted model"
result=$(dcld query model model-version --vid=$vid --pid=$pid --softwareVersion=$sv)
check_response "$result" "Not Found"
echo "$result"

echo "Query model versions for deleted model"
result=$(dcld query model model-version --vid=$vid_with_pids --pid=$pid --softwareVersion=$sv)
check_response "$result" "Not Found"
echo "$result"