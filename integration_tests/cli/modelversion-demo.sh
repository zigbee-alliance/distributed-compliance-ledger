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


# Create a new model version

echo "Add Model with VID: $vid PID: $pid"
result=$(echo "test1234" | dcld tx model add-model --vid=$vid --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel="Test Product" --partNumber=1 --commissioningCustomFlow=0 --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"

test_divider

sv=$RANDOM
echo "Create a Device Model Version with VID: $vid PID: $pid SV: $sv"
result=$(echo 'test1234' | dcld tx model add-model-version --cdVersionNumber=1 --maxApplicableSoftwareVersion=10 --minApplicableSoftwareVersion=1 --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=1 --from=$vendor_account --yes)
echo "$result"
check_response "$result" "\"code\": 0"

test_divider

# Query the model version 
echo "Query Device Model Version with VID: $vid PID: $pid SV: $sv"
result=$(dcld query model get-model-version --vid=$vid --pid=$pid --softwareVersion=$sv)
echo "$result"

check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"software_version\": $sv"
check_response "$result" "\"software_version_string\": \"1\""
check_response "$result" "\"cd_version_number\": 1"
check_response "$result" "\"software_version_valid\": true"
check_response "$result" "\"min_applicable_software_version\": 1"
check_response "$result" "\"max_applicable_software_version\": 10"

test_divider

# Query non existant model version 
echo "Query Device Model Version with VID: $vid PID: $pid SV: 123456"
result=$(dcld query model get-model-version --vid=$vid --pid=$pid --softwareVersion=123456 2>&1) || true
check_response_and_report "$result" "code = InvalidArgument desc = not found" plain

test_divider

# Update the existing model version
echo "Update Device Model Version with VID: $vid PID: $pid SV: $sv"
result=$(echo 'test1234' | dcld tx model update-model-version --vid=$vid --pid=$pid --minApplicableSoftwareVersion=2 --maxApplicableSoftwareVersion=10 --softwareVersion=$sv --softwareVersionValid=false --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"

test_divider

# Query Updated model version
echo "Query updated Device Model Version with VID: $vid PID: $pid SV: $sv"
result=$(dcld query model get-model-version --vid=$vid --pid=$pid --softwareVersion=$sv)
echo "$result"
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"software_version\": $sv"
check_response "$result" "\"software_version_string\": \"1\""
check_response "$result" "\"cd_version_number\": 1"
check_response "$result" "\"software_version_valid\": false"
check_response "$result" "\"min_applicable_software_version\": 2"
check_response "$result" "\"max_applicable_software_version\": 10"

test_divider

# Create model version with vid belonging to another vendor
echo "Create a Device Model Version with VID: $vid PID: $pid SV: $sv from a different vendor account"
newvid=$RANDOM
different_vendor_account=vendor_account_$newvid
create_new_vendor_account $different_vendor_account $newvid
result=$(echo 'test1234' | dcld tx model add-model-version --cdVersionNumber=1 --maxApplicableSoftwareVersion=10 --minApplicableSoftwareVersion=1 --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=1 --from=$different_vendor_account --yes)
check_response_and_report "$result" "transaction should be signed by a vendor account containing the vendorID $vid"

test_divider

# Update model version with vid belonging to another vendor
echo "Update a Device Model Version with VID: $vid PID: $pid SV: $sv from a different vendor account"
result=$(echo 'test1234' | dcld tx model update-model-version --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionValid=false --from=$different_vendor_account --yes)
check_response_and_report "$result" "transaction should be signed by a vendor account containing the vendorID $vid"

