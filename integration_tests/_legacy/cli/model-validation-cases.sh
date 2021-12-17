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

# FIXME issue 99: enable once implemented
exit 0

set -euo pipefail
source integration_tests/cli/common.sh

# Preparation of Actors

vid_1=$RANDOM
vendor_account_1=vendor_account_$vid_1
echo "Create Vendor account $vendor_account_1 with vid $vid_1"
create_new_vendor_account $vendor_account_1 $vid_1

pid_1=$RANDOM
pid_2=$RANDOM
pid_3=$RANDOM


test_divider

# Create a new model with minimum fields
echo "Add Model with minimum required fields with VID: $vid_1 PID: $pid_1"
result=$(echo "test1234" | dcld tx model add-model --vid=$vid_1 --pid=$pid_1 --deviceTypeID=1 --productName=TestProduct --productLabel="Test Product" --partNumber=1 --from=$vendor_account_1 --yes)
check_response "$result" "\"success\": true"

test_divider

# Query the model created above to see if it is added
echo "Query the model created above to see if it is added"
result=$(dcld query model get-model --vid=$vid_1 --pid=$pid_1)
check_response_and_report "$result" "\"vid\": $vid_1"
check_response_and_report "$result" "\"pid\": $pid_1"
check_response_and_report "$result" "\"productName\": \"TestProduct\""
check_response_and_report "$result" "\"partNumber\": \"1\""
check_response_and_report "$result" "\"commissioningCustomFlow\": 0"

test_divider

# Create a new model with all fields
echo "Add Model with all fields with VID: $vid_1 PID: $pid_2"
result=$(echo "test1234" | dcld tx model add-model --vid=$vid_1 --pid=$pid_2 --deviceTypeID=2 --productName="Test Product with All Fields" --productLabel="Test Product with All fields" \
--partNumber="23.456" --commissioningCustomFlow=1 --commissioningCustomFlowURL="https://customflow.url.info" \
--commissioningModeInitialStepsHint=1  --commissioningModeInitialStepsInstruction="Initial Instructions" \
--commissioningModeSecondaryStepsHint=2 --commissioningModeSecondaryStepsInstruction="Secondary Steps Instruction" \
--userManualURL="https://usermanual.url" --productURL="https://product.url.info" --supportURL="https://support.url.info"   --from=$vendor_account_1 --yes)
check_response "$result" "\"success\": true"

test_divider

# Query the model created above to see if it is added
echo "Query the model created above to see if it is added"
result=$(dcld query model get-model --vid=$vid_1 --pid=$pid_2)
check_response_and_report "$result" "\"vid\": $vid_1"
check_response_and_report "$result" "\"pid\": $pid_2"
check_response_and_report "$result" "\"deviceTypeID\": 2"
check_response_and_report "$result" "\"productName\": \"Test Product with All Fields\""
check_response_and_report "$result" "\"productLabel\": \"Test Product with All fields\""
check_response_and_report "$result" "\"partNumber\": \"23.456\""
check_response_and_report "$result" "\"commissioningCustomFlow\": 1"
check_response_and_report "$result" "\"commissioningCustomFlowURL\": \"https://customflow.url.info\""
check_response_and_report "$result" "\"commissioningModeInitialStepsHint\": 1"
check_response_and_report "$result" "\"commissioningModeInitialStepsInstruction\": \"Initial Instructions\""
check_response_and_report "$result" "\"commissioningModeSecondaryStepsHint\": 2"
check_response_and_report "$result" "\"commissioningModeSecondaryStepsInstruction\": \"Secondary Steps Instruction\""
check_response_and_report "$result" "\"userManualURL\": \"https://usermanual.url\""
check_response_and_report "$result" "\"supportURL\": \"https://support.url.info\""
check_response_and_report "$result" "\"productURL\": \"https://product.url.info\""


test_divider

# Create a new model with mandatory and some non mandatory fields
echo "Add Model with mandatory and some non mandatory fields with VID: $vid_1 PID: $pid_3"
result=$(echo "test1234" | dcld tx model add-model --vid=$vid_1 --pid=$pid_3 --deviceTypeID=2 --productName="Test Product with All Fields" --productLabel="Test Product with All fields" \
--partNumber="23.456" --commissioningCustomFlow=1 --commissioningCustomFlowURL="https://customflow.url.info" \
--commissioningModeInitialStepsHint=1  --commissioningModeInitialStepsInstruction="Initial Instructions" \
--commissioningModeSecondaryStepsHint=2 --commissioningModeSecondaryStepsInstruction="Secondary Steps Instruction" \
--from=$vendor_account_1 --yes)
check_response "$result" "\"success\": true"

test_divider

# Query the model created above to see if it is added
echo "Query the model created above to see if it is added"
result=$(dcld query model get-model --vid=$vid_1 --pid=$pid_3)
check_response_and_report "$result" "\"vid\": $vid_1"
check_response_and_report "$result" "\"pid\": $pid_3"
check_response_and_report "$result" "\"deviceTypeID\": 2"
check_response_and_report "$result" "\"productName\": \"Test Product with All Fields\""
check_response_and_report "$result" "\"productLabel\": \"Test Product with All fields\""
check_response_and_report "$result" "\"partNumber\": \"23.456\""
check_response_and_report "$result" "\"commissioningCustomFlow\": 1"
check_response_and_report "$result" "\"commissioningCustomFlowURL\": \"https://customflow.url.info\""
check_response_and_report "$result" "\"commissioningModeInitialStepsHint\": 1"
check_response_and_report "$result" "\"commissioningModeInitialStepsInstruction\": \"Initial Instructions\""
check_response_and_report "$result" "\"commissioningModeSecondaryStepsHint\": 2"
check_response_and_report "$result" "\"commissioningModeSecondaryStepsInstruction\": \"Secondary Steps Instruction\""
response_does_not_contain "$result" "\"userManualURL\""
response_does_not_contain "$result" "\"supportURL\""

test_divider

# Update model with mutable fields and make sure they are updated properly
echo "Update model with mutable fields and make sure they are updated properly VID: $vid_1 PID: $pid_1"
result=$(echo "test1234" | dcld tx model update-model --vid=$vid_1 --pid=$pid_1 --productName="Updated Product Name" --productLabel="Updated Test Product" --partNumber="2" --from=$vendor_account_1 --yes)
check_response "$result" "\"success\": true"

test_divider

# Query the model updated above to see if it is updated
echo "Query the model updated above to see if it is updated"
result=$(dcld query model get-model --vid=$vid_1 --pid=$pid_1)
check_response_and_report "$result" "\"vid\": $vid_1"
check_response_and_report "$result" "\"pid\": $pid_1"
check_response_and_report "$result" "\"productName\": \"Updated Product Name\""
check_response_and_report "$result" "\"partNumber\": \"2\""
check_response_and_report "$result" "\"productLabel\": \"Updated Test Product\""
check_response_and_report "$result" "\"commissioningCustomFlow\": 0" # default value set when this model was created

test_divider
# Update model with just one mutable fields and make sure they are updated properly
echo "Update model with just one mutable field and make sure they are updated properly VID: $vid_1 PID: $pid_1"
result=$(echo "test1234" | dcld tx model update-model --vid=$vid_1 --pid=$pid_1  --productLabel="Updated Test Product V2" --from=$vendor_account_1 --yes)
check_response "$result" "\"success\": true"

test_divider
# Query the model updated above to see if it is updated
echo "Query the model updated above to see if it is updated"
result=$(dcld query model get-model --vid=$vid_1 --pid=$pid_1)
check_response_and_report "$result" "\"vid\": $vid_1"
check_response_and_report "$result" "\"pid\": $pid_1"
check_response_and_report "$result" "\"productName\": \"Updated Product Name\""
check_response_and_report "$result" "\"partNumber\": \"2\""
check_response_and_report "$result" "\"productLabel\": \"Updated Test Product V2\""
check_response_and_report "$result" "\"commissioningCustomFlow\": 0" # default value set when this model was created

test_divider
# Update model with all possible mutable fields and make sure they are updated properly
echo "Update model with just all mutable fields and make sure they are updated properly VID: $vid_1 PID: $pid_1"
result=$(echo "test1234" | dcld tx model update-model --vid=$vid_1 --pid=$pid_1 --productName="Updated Product Name V3" \
--partNumber="V3" --commissioningCustomFlowURL="https://updated.url.info" \
--productLabel="Updated Test Product V3" --commissioningModeInitialStepsInstruction="Instructions updated v3" \
--commissioningModeSecondaryStepsInstruction="Secondary Instructions v3" --userManualURL="https://userManual.info/v3" \
--supportURL="https://support.url.info/v3" --productURL="https://product.landingpage.url" --from=$vendor_account_1 --yes)
check_response "$result" "\"success\": true"

test_divider
# Query the model updated above to see if it is updated
echo "Query the model updated above to see if it is updated"
result=$(dcld query model get-model --vid=$vid_1 --pid=$pid_1)
check_response_and_report "$result" "\"vid\": $vid_1"
check_response_and_report "$result" "\"pid\": $pid_1"
check_response_and_report "$result" "\"productName\": \"Updated Product Name V3\""
check_response_and_report "$result" "\"partNumber\": \"V3\""
check_response_and_report "$result" "\"productLabel\": \"Updated Test Product V3\""
check_response_and_report "$result" "\"commissioningCustomFlow\": 0" # default value set when this model was created
check_response_and_report "$result" "\"commissioningCustomFlowURL\": \"https://updated.url.info\""
check_response_and_report "$result" "\"commissioningModeInitialStepsInstruction\": \"Instructions updated v3\""
check_response_and_report "$result" "\"commissioningModeSecondaryStepsInstruction\": \"Secondary Instructions v3\""
check_response_and_report "$result" "\"userManualURL\": \"https://userManual.info/v3\""
check_response_and_report "$result" "\"supportURL\": \"https://support.url.info/v3\""
check_response_and_report "$result" "\"productURL\": \"https://product.landingpage.url\""


test_divider
# Update model with just one mutable fields and make sure they are updated properly
echo "Update model with just one mutable field and make sure all other mutated fields are still the same for VID: $vid_1 PID: $pid_1"
result=$(echo "test1234" | dcld tx model update-model --vid=$vid_1 --pid=$pid_1  --productLabel="Updated Test Product V4" --from=$vendor_account_1 --yes)
check_response "$result" "\"success\": true"

test_divider
# Query the model updated above to see if it is updated
echo "Query the model updated above to see if it is updated"
result=$(dcld query model get-model --vid=$vid_1 --pid=$pid_1)
check_response_and_report "$result" "\"vid\": $vid_1"
check_response_and_report "$result" "\"pid\": $pid_1"
check_response_and_report "$result" "\"productName\": \"Updated Product Name V3\""
check_response_and_report "$result" "\"partNumber\": \"V3\""
check_response_and_report "$result" "\"productLabel\": \"Updated Test Product V4\""
check_response_and_report "$result" "\"commissioningCustomFlow\": 0" # default value set when this model was created
check_response_and_report "$result" "\"commissioningCustomFlowURL\": \"https://updated.url.info\""
check_response_and_report "$result" "\"commissioningModeInitialStepsInstruction\": \"Instructions updated v3\""
check_response_and_report "$result" "\"commissioningModeSecondaryStepsInstruction\": \"Secondary Instructions v3\""
check_response_and_report "$result" "\"userManualURL\": \"https://userManual.info/v3\""
check_response_and_report "$result" "\"supportURL\": \"https://support.url.info/v3\""
check_response_and_report "$result" "\"productURL\": \"https://product.landingpage.url\""

test_divider
# Update model with just one mutable fields and make sure they are updated properly
echo "Update model with no fields are still the same for VID: $vid_1 PID: $pid_1"
result=$(echo "test1234" | dcld tx model update-model --vid=$vid_1 --pid=$pid_1 --from=$vendor_account_1 --yes)
check_response "$result" "\"success\": true"

test_divider
# Query the model updated above to see if it is updated
echo "Query the model updated above to see if it is added"
result=$(dcld query model get-model --vid=$vid_1 --pid=$pid_1)
check_response_and_report "$result" "\"vid\": $vid_1"
check_response_and_report "$result" "\"pid\": $pid_1"
check_response_and_report "$result" "\"productName\": \"Updated Product Name V3\""
check_response_and_report "$result" "\"partNumber\": \"V3\""
check_response_and_report "$result" "\"productLabel\": \"Updated Test Product V4\""
check_response_and_report "$result" "\"commissioningCustomFlow\": 0" # default value set when this model was created
check_response_and_report "$result" "\"commissioningCustomFlowURL\": \"https://updated.url.info\""
check_response_and_report "$result" "\"commissioningModeInitialStepsInstruction\": \"Instructions updated v3\""
check_response_and_report "$result" "\"commissioningModeSecondaryStepsInstruction\": \"Secondary Instructions v3\""
check_response_and_report "$result" "\"userManualURL\": \"https://userManual.info/v3\""
check_response_and_report "$result" "\"supportURL\": \"https://support.url.info/v3\""
check_response_and_report "$result" "\"productURL\": \"https://product.landingpage.url\""


test_divider

################################################################################
### Model Version Validation
################################################################################

sv_1=$RANDOM
# Create a new model version
echo "Create a Device Model Version with minimum mandatory fields for VID: $vid_1 PID: $pid_1 SV: $sv_1"
result=$(echo 'test1234' | dcld tx model add-model-version --cdVersionNumber=1 --maxApplicableSoftwareVersion=20 --minApplicableSoftwareVersion=10 --vid=$vid_1 --pid=$pid_1 --softwareVersion=$sv_1 --softwareVersionString=1 --from=$vendor_account_1 --yes)
echo "$result"
check_response_and_report "$result" "\"success\": true"

test_divider

# Query the model version 
echo "Query Device Model Version with VID: $vid_1 PID: $pid_1 SV: $sv_1"
result=$(dcld query model get-model-version --vid=$vid_1 --pid=$pid_1 --softwareVersion=$sv_1)

check_response_and_report "$result" "\"vid\": $vid_1"
check_response_and_report "$result" "\"pid\": $pid_1"
check_response_and_report "$result" "\"softwareVersion\": $sv_1"
check_response_and_report "$result" "\"softwareVersionString\": \"1\""
check_response_and_report "$result" "\"CDVersionNumber\": 1"
check_response_and_report "$result" "\"softwareVersionValid\": true"
check_response_and_report "$result" "\"minApplicableSoftwareVersion\": 10"
check_response_and_report "$result" "\"maxApplicableSoftwareVersion\": 20"

test_divider

# Update the model version with only one mutable field and make sure all other fields are still the same
echo "Update Device Model Version with only one mutable field and make sure all other fields are still the same for VID: $vid_1 PID: $pid_1 SV: $sv_1"
result=$(echo "test1234" | dcld tx model update-model-version --vid=$vid_1 --pid=$pid_1 --softwareVersion=$sv_1 --softwareVersionValid=false --from=$vendor_account_1 --yes)
check_response_and_report "$result" "\"success\": true"

test_divider

# Query the updated model version 
echo "Query the updated model version VID: $vid_1 PID: $pid_1 SV: $sv_1"
result=$(dcld query model get-model-version --vid=$vid_1 --pid=$pid_1 --softwareVersion=$sv_1)

check_response_and_report "$result" "\"vid\": $vid_1"
check_response_and_report "$result" "\"pid\": $pid_1"
check_response_and_report "$result" "\"softwareVersion\": $sv_1"
check_response_and_report "$result" "\"softwareVersionString\": \"1\""
check_response_and_report "$result" "\"CDVersionNumber\": 1"
check_response_and_report "$result" "\"softwareVersionValid\": false"
check_response_and_report "$result" "\"minApplicableSoftwareVersion\": 10"
check_response_and_report "$result" "\"maxApplicableSoftwareVersion\": 20"

test_divider

# Update the model version with few mutable fields and make sure all other fields are still the same
echo "Update Device Model Version with few mutable fields and make sure all other fields are still the same for VID: $vid_1 PID: $pid_1 SV: $sv_1"
result=$(echo "test1234" | dcld tx model update-model-version --vid=$vid_1 --pid=$pid_1 --softwareVersion=$sv_1 --softwareVersionValid=true \
--releaseNotesURL="https://release.url.info" --minApplicableSoftwareVersion=2 --from=$vendor_account_1 --yes)
check_response_and_report "$result" "\"success\": true"

test_divider

# Query the updated model version 
echo "Query the updated model version VID: $vid_1 PID: $pid_1 SV: $sv_1"
result=$(dcld query model get-model-version --vid=$vid_1 --pid=$pid_1 --softwareVersion=$sv_1)

check_response_and_report "$result" "\"vid\": $vid_1"
check_response_and_report "$result" "\"pid\": $pid_1"
check_response_and_report "$result" "\"softwareVersion\": $sv_1"
check_response_and_report "$result" "\"softwareVersionString\": \"1\""
check_response_and_report "$result" "\"CDVersionNumber\": 1"
check_response_and_report "$result" "\"softwareVersionValid\": true"
check_response_and_report "$result" "\"minApplicableSoftwareVersion\": 2"
check_response_and_report "$result" "\"maxApplicableSoftwareVersion\": 20"
check_response_and_report "$result" "\"releaseNotesURL\": \"https://release.url.info\""



sv_1=$RANDOM
echo "Create a Device Model Version with all fields for VID: $vid_1 PID: $pid_1 SV: $sv_1"
result=$(echo 'test1234' | dcld tx model add-model-version --vid=$vid_1 --pid=$pid_1 --softwareVersion=$sv_1 \
--softwareVersionString="1.0" --cdVersionNumber=21334 \
--firmwareDigests="123456789012345678901234567890123456789012345678901234567890123" \
--softwareVersionValid=true --otaURL="https://ota.url.info" --otaFileSize=123456789 \
--otaChecksum="123456789012345678901234567890123456789012345678901234567890123" --releaseNotesURL="https://release.notes.url.info" \
--otaChecksumType=1 --maxApplicableSoftwareVersion=32 --minApplicableSoftwareVersion=5   --from=$vendor_account_1 --yes)
echo "$result"
check_response_and_report "$result" "\"success\": true"

test_divider

# Query the model version 
echo "Query Device Model Version with VID: $vid_1 PID: $pid_1 SV: $sv_1"
result=$(dcld query model get-model-version --vid=$vid_1 --pid=$pid_1 --softwareVersion=$sv_1)

check_response_and_report "$result" "\"vid\": $vid_1"
check_response_and_report "$result" "\"pid\": $pid_1"
check_response_and_report "$result" "\"softwareVersion\": $sv_1"
check_response_and_report "$result" "\"softwareVersionString\": \"1.0\""
check_response_and_report "$result" "\"CDVersionNumber\": 21334"
check_response_and_report "$result" "\"firmwareDigests\": \"123456789012345678901234567890123456789012345678901234567890123\""
check_response_and_report "$result" "\"softwareVersionValid\": true"
check_response_and_report "$result" "\"otaURL\": \"https://ota.url.info\""
check_response_and_report "$result" "\"otaFileSize\": 123456789"
check_response_and_report "$result" "\"otaChecksum\": \"123456789012345678901234567890123456789012345678901234567890123\""
check_response_and_report "$result" "\"otaChecksumType\": 1"
check_response_and_report "$result" "\"releaseNotesURL\": \"https://release.notes.url.info\""
check_response_and_report "$result" "\"maxApplicableSoftwareVersion\": 32"
check_response_and_report "$result" "\"minApplicableSoftwareVersion\": 5"

test_divider

# Update the model version with minimum fields i.e. no update at all 
echo "Update Device Model Version with only one mutable field and make sure all other fields are still the same for VID: $vid_1 PID: $pid_1 SV: $sv_1"
result=$(echo "test1234" | dcld tx model update-model-version --vid=$vid_1 --pid=$pid_1 --softwareVersion=$sv_1 --from=$vendor_account_1 --yes)
check_response_and_report "$result" "\"success\": true"

test_divider

# Query the model version 
echo "Query Updated Model Version with VID: $vid_1 PID: $pid_1 SV: $sv_1"
result=$(dcld query model get-model-version --vid=$vid_1 --pid=$pid_1 --softwareVersion=$sv_1)

check_response_and_report "$result" "\"vid\": $vid_1"
check_response_and_report "$result" "\"pid\": $pid_1"
check_response_and_report "$result" "\"softwareVersion\": $sv_1"
check_response_and_report "$result" "\"softwareVersionString\": \"1.0\""
check_response_and_report "$result" "\"CDVersionNumber\": 21334"
check_response_and_report "$result" "\"firmwareDigests\": \"123456789012345678901234567890123456789012345678901234567890123\""
check_response_and_report "$result" "\"softwareVersionValid\": true"
check_response_and_report "$result" "\"otaURL\": \"https://ota.url.info\""
check_response_and_report "$result" "\"otaFileSize\": 123456789"
check_response_and_report "$result" "\"otaChecksum\": \"123456789012345678901234567890123456789012345678901234567890123\""
check_response_and_report "$result" "\"otaChecksumType\": 1"
check_response_and_report "$result" "\"releaseNotesURL\": \"https://release.notes.url.info\""
check_response_and_report "$result" "\"maxApplicableSoftwareVersion\": 32"
check_response_and_report "$result" "\"minApplicableSoftwareVersion\": 5"

test_divider

# Update the model version with only one mutable field and make sure all other fields are still the same
echo "Update Device Model Version with only one mutable field and make sure all other fields are still the same for VID: $vid_1 PID: $pid_1 SV: $sv_1"
result=$(echo "test1234" | dcld tx model update-model-version --vid=$vid_1 --pid=$pid_1 --softwareVersion=$sv_1 --softwareVersionValid=false --from=$vendor_account_1 --yes)
check_response_and_report "$result" "\"success\": true"

test_divider

# Query the model version 
echo "Query Updated Model Version with VID: $vid_1 PID: $pid_1 SV: $sv_1"
result=$(dcld query model get-model-version --vid=$vid_1 --pid=$pid_1 --softwareVersion=$sv_1)

check_response_and_report "$result" "\"vid\": $vid_1"
check_response_and_report "$result" "\"pid\": $pid_1"
check_response_and_report "$result" "\"softwareVersion\": $sv_1"
check_response_and_report "$result" "\"softwareVersionString\": \"1.0\""
check_response_and_report "$result" "\"CDVersionNumber\": 21334"
check_response_and_report "$result" "\"firmwareDigests\": \"123456789012345678901234567890123456789012345678901234567890123\""
check_response_and_report "$result" "\"softwareVersionValid\": false"
check_response_and_report "$result" "\"otaURL\": \"https://ota.url.info\""
check_response_and_report "$result" "\"otaFileSize\": 123456789"
check_response_and_report "$result" "\"otaChecksum\": \"123456789012345678901234567890123456789012345678901234567890123\""
check_response_and_report "$result" "\"otaChecksumType\": 1"
check_response_and_report "$result" "\"releaseNotesURL\": \"https://release.notes.url.info\""
check_response_and_report "$result" "\"maxApplicableSoftwareVersion\": 32"
check_response_and_report "$result" "\"minApplicableSoftwareVersion\": 5"

test_divider

echo "Update Device Model Version with all mutable fields for VID: $vid_1 PID: $pid_1 SV: $sv_1"
result=$(echo 'test1234' | dcld tx model update-model-version --vid=$vid_1 --pid=$pid_1 --softwareVersion=$sv_1 \
--softwareVersionValid=true --otaURL="https://updated.ota.url.info" --releaseNotesURL="https://updated.release.notes.url.info" \
--maxApplicableSoftwareVersion=25 --minApplicableSoftwareVersion=15   --from=$vendor_account_1 --yes)
echo "$result"
check_response_and_report "$result" "\"success\": true"

test_divider

# Query the model version 
echo "Query Updated Model Version with VID: $vid_1 PID: $pid_1 SV: $sv_1"
result=$(dcld query model get-model-version --vid=$vid_1 --pid=$pid_1 --softwareVersion=$sv_1)

check_response_and_report "$result" "\"vid\": $vid_1"
check_response_and_report "$result" "\"pid\": $pid_1"
check_response_and_report "$result" "\"softwareVersion\": $sv_1"
check_response_and_report "$result" "\"softwareVersionString\": \"1.0\""
check_response_and_report "$result" "\"CDVersionNumber\": 21334"
check_response_and_report "$result" "\"firmwareDigests\": \"123456789012345678901234567890123456789012345678901234567890123\""
check_response_and_report "$result" "\"softwareVersionValid\": true"
check_response_and_report "$result" "\"otaURL\": \"https://updated.ota.url.info\""
check_response_and_report "$result" "\"otaFileSize\": 123456789"
check_response_and_report "$result" "\"otaChecksum\": \"123456789012345678901234567890123456789012345678901234567890123\""
check_response_and_report "$result" "\"otaChecksumType\": 1"
check_response_and_report "$result" "\"releaseNotesURL\": \"https://updated.release.notes.url.info\""
check_response_and_report "$result" "\"maxApplicableSoftwareVersion\": 25"
check_response_and_report "$result" "\"minApplicableSoftwareVersion\": 15"

sv_1=$RANDOM
echo "Create a Device Model Version with mandatory fields and some optional fields for VID: $vid_1 PID: $pid_1 SV: $sv_1"
result=$(echo 'test1234' | dcld tx model add-model-version --vid=$vid_1 --pid=$pid_1 --softwareVersion=$sv_1 \
--softwareVersionString="1.0" --cdVersionNumber=21334 \
--firmwareDigests="123456789012345678901234567890123456789012345678901234567890123" \
--softwareVersionValid=true --otaURL="https://ota.url.info" --otaFileSize=123456789 \
--otaChecksum="123456789012345678901234567890123456789012345678901234567890123" \
--otaChecksumType=1 --maxApplicableSoftwareVersion=32 --minApplicableSoftwareVersion=5   --from=$vendor_account_1 --yes)
echo "$result"
check_response_and_report "$result" "\"success\": true"

test_divider

# Query the model version 
echo "Query Device Model Version with VID: $vid_1 PID: $pid_1 SV: $sv_1"
result=$(dcld query model get-model-version --vid=$vid_1 --pid=$pid_1 --softwareVersion=$sv_1)

check_response_and_report "$result" "\"vid\": $vid_1"
check_response_and_report "$result" "\"pid\": $pid_1"
check_response_and_report "$result" "\"softwareVersion\": $sv_1"
check_response_and_report "$result" "\"softwareVersionString\": \"1.0\""
check_response_and_report "$result" "\"CDVersionNumber\": 21334"
check_response_and_report "$result" "\"firmwareDigests\": \"123456789012345678901234567890123456789012345678901234567890123\""
check_response_and_report "$result" "\"softwareVersionValid\": true"
check_response_and_report "$result" "\"otaURL\": \"https://ota.url.info\""
check_response_and_report "$result" "\"otaFileSize\": 123456789"
check_response_and_report "$result" "\"otaChecksum\": \"123456789012345678901234567890123456789012345678901234567890123\""
check_response_and_report "$result" "\"otaChecksumType\": 1"
check_response_and_report "$result" "\"maxApplicableSoftwareVersion\": 32"
check_response_and_report "$result" "\"minApplicableSoftwareVersion\": 5"
response_does_not_contain "$result" "\"releaseNotesURL\""

test_divider

# Update the model version with maxApplicableSoftwareVersion less then minApplicableSoftwareVersion 
echo "Update the model version with maxApplicableSoftwareVersion less then minApplicableSoftwareVersion and make sure we get error back VID: $vid_1 PID: $pid_1 SV: $sv_1"
result=$(echo "test1234" | dcld tx model update-model-version --vid=$vid_1 --pid=$pid_1 --softwareVersion=$sv_1 --maxApplicableSoftwareVersion=3 --from=$vendor_account_1 --yes)
check_response_and_report "$result" "\"success\": false"

test_divider

# Update the model version with minApplicableSoftwareVersion greater then maxApplicableSoftwareVersion 
echo "Update the model version with minApplicableSoftwareVersion greater then maxApplicableSoftwareVersion and make sure we get error back VID: $vid_1 PID: $pid_1 SV: $sv_1"
result=$(echo "test1234" | dcld tx model update-model-version --vid=$vid_1 --pid=$pid_1 --softwareVersion=$sv_1 --minApplicableSoftwareVersion=33 --from=$vendor_account_1 --yes)
check_response_and_report "$result" "\"success\": false"

test_divider
