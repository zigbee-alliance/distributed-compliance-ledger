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

echo "Create CertificationCenter account"
create_new_account zb_account "CertificationCenter"

# Create a new model version
echo "Add Model with VID: $vid PID: $pid"
result=$(echo "test1234" | dcld tx model add-model --vid=$vid --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel="Test Product" --partNumber=1 --commissioningCustomFlow=0 --enhancedSetupFlowOptions=0 --from=$vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

sv=$RANDOM
schema_version_0=0

echo "Certify  Model with VID: $vid PID: $pid  SV: ${sv} with zigbee certification"
result=$(echo "test1234" | dcld tx compliance certify-model --vid=$vid --pid=$pid --softwareVersion=$sv --cdVersionNumber=1 --softwareVersionString=1 --cdCertificateId=12345678910abcdefgh --certificationType=zigbee --specificationVersion=1 --certificationDate="2020-01-01T00:00:01Z" --from $zb_account --yes)
result=$(get_txn_result "$result")
echo "$result"
check_response "$result" "\"code\": 0"

echo "Create a Device Model Version with VID: $vid PID: $pid SV: $sv"
result=$(echo 'test1234' | dcld tx model add-model-version --cdVersionNumber=1 --maxApplicableSoftwareVersion=10 --minApplicableSoftwareVersion=1 --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=1   --schemaVersion=$schema_version_0 --from=$vendor_account --yes)
result=$(get_txn_result "$result")
echo "$result"
check_response "$result" "\"code\": 0"

test_divider

# Query the model version 
echo "Query Device Model Version with VID: $vid PID: $pid SV: $sv"
result=$(dcld query model model-version --vid=$vid --pid=$pid --softwareVersion=$sv)
echo "$result"
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"softwareVersion\": $sv"
check_response "$result" "\"softwareVersionString\": \"1\""
check_response "$result" "\"cdVersionNumber\": 1"
check_response "$result" "\"softwareVersionValid\": true"
check_response "$result" "\"minApplicableSoftwareVersion\": 1"
check_response "$result" "\"maxApplicableSoftwareVersion\": 10"
check_response "$result" "\"schemaVersion\": $schema_version_0"

test_divider

# Query all model versions
echo "Query all model versions with VID: $vid PID: $pid "
result=$(dcld query model all-model-versions --vid=$vid --pid=$pid)
echo "$result"
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"softwareVersions\""
check_response "$result" "$sv"

test_divider



# Query non existent model version
echo "Query Device Model Version with VID: $vid PID: $pid SV: 123456"
result=$(dcld query model model-version --vid=$vid --pid=$pid --softwareVersion=123456)
check_response "$result" "Not Found"

test_divider

# Query non existent model versions
vid1=$RANDOM
pid1=$RANDOM
echo "Query all Device Model Versions with VID: $vid1 PID: $pid1"
result=$(dcld query model all-model-versions --vid=$vid1 --pid=$pid1)
check_response "$result" "Not Found"

test_divider

# Update the existing model version
echo "Update Device Model Version with VID: $vid PID: $pid SV: $sv"
result=$(echo 'test1234' | dcld tx model update-model-version --vid=$vid --pid=$pid --minApplicableSoftwareVersion=2 --maxApplicableSoftwareVersion=10 --softwareVersion=$sv --softwareVersionValid=false  --schemaVersion=$schema_version_0 --from=$vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

# Query Updated model version
echo "Query updated Device Model Version with VID: $vid PID: $pid SV: $sv"
result=$(dcld query model model-version --vid=$vid --pid=$pid --softwareVersion=$sv)
echo "$result"
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"softwareVersion\": $sv"
check_response "$result" "\"softwareVersionString\": \"1\""
check_response "$result" "\"cdVersionNumber\": 1"
check_response "$result" "\"softwareVersionValid\": false"
check_response "$result" "\"minApplicableSoftwareVersion\": 2"
check_response "$result" "\"maxApplicableSoftwareVersion\": 10"
check_response "$result" "\"schemaVersion\": $schema_version_0"

test_divider

# Add second model version
sv2=$RANDOM
echo "Create a Second Device Model Version with VID: $vid PID: $pid SV: $sv2"
result=$(echo 'test1234' | dcld tx model add-model-version --cdVersionNumber=1 --maxApplicableSoftwareVersion=10 --minApplicableSoftwareVersion=1 --vid=$vid --pid=$pid --softwareVersion=$sv2 --softwareVersionString=1 --from=$vendor_account --yes)
result=$(get_txn_result "$result")
echo "$result"
check_response "$result" "\"code\": 0"

test_divider

# Query all model versions
echo "Query all model versions with VID: $vid PID: $pid "
result=$(dcld query model all-model-versions --vid=$vid --pid=$pid)
echo "$result"
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"softwareVersions\""
check_response "$result" "$sv"
check_response "$result" "$sv2"

test_divider


# Create model version with vid belonging to another vendor
echo "Create a Device Model Version with VID: $vid PID: $pid SV: $sv from a different vendor account"
newvid=$RANDOM
different_vendor_account=vendor_account_$newvid
create_new_vendor_account $different_vendor_account $newvid
result=$(echo 'test1234' | dcld tx model add-model-version --cdVersionNumber=1 --maxApplicableSoftwareVersion=10 --minApplicableSoftwareVersion=1 --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=1 --from=$different_vendor_account --yes)
result=$(get_txn_result "$result")
check_response_and_report "$result" "transaction should be signed by a vendor account containing the vendorID $vid"

test_divider

# Update model version with vid belonging to another vendor
echo "Update a Device Model Version with VID: $vid PID: $pid SV: $sv from a different vendor account"
result=$(echo 'test1234' | dcld tx model update-model-version --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionValid=false --from=$different_vendor_account --yes)
result=$(get_txn_result "$result")
check_response_and_report "$result" "transaction should be signed by a vendor account containing the vendorID $vid"

# Delete corresponding compliance info
echo "Delete compliance info vid=$vid pid=$pid softwareVersion=$sv certificationType=zigbee"
result=$(echo 'test1234' | dcld tx compliance delete-compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=zigbee --from=$zb_account --yes)
result=$(get_txn_result "$result")

# Delete existing model version
echo "Delete a Device Model Version with VID: $vid PID: $pid SV: $sv"
result=$(echo 'test1234' | dcld tx model delete-model-version --vid=$vid --pid=$pid --softwareVersion=$sv --from=$vendor_account --yes)
result=$(get_txn_result "$result")
echo "$result"
check_response "$result" "\"code\": 0"

# Query deleted model version
echo "Query deleted Device Model Version with VID: $vid PID: $pid SV: $sv"
result=$(dcld query model model-version --vid=$vid --pid=$pid --softwareVersion=$sv)
echo "$result"
check_response "$result" "Not Found"

test_divider

echo "VendorAdmin can add/update/delete model versions for any vendor"
create_new_account vendor_admin_account "VendorAdmin"
vid3=$RANDOM
pid3=$RANDOM
sv3=$RANDOM

echo "Add Model with VID: $vid3 PID: $pid3"
result=$(echo "test1234" | dcld tx model add-model --vid=$vid3 --pid=$pid3 --deviceTypeID=1 --productName=TestProduct --productLabel="Test Product" --partNumber=1 --commissioningCustomFlow=0 --enhancedSetupFlowOptions=0 --from=$vendor_admin_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "VendorAdmin adds a model version for VID: $vid3 PID: $pid3 SV: $sv3"
result=$(echo 'test1234' | dcld tx model add-model-version --cdVersionNumber=1 --maxApplicableSoftwareVersion=10 --minApplicableSoftwareVersion=1 --vid=$vid3 --pid=$pid3 --softwareVersion=$sv3 --softwareVersionString=1 --from=$vendor_admin_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"
echo "$result"

echo "Query Device Model Version with VID: $vid3 PID: $pid3 SV: $sv3"
result=$(dcld query model model-version --vid=$vid3 --pid=$pid3 --softwareVersion=$sv3)
check_response "$result" "\"vid\": $vid3"
check_response "$result" "\"softwareVersion\": $sv3"
echo "$result"

echo "VendorAdmin updates a model version for VID: $vid3 PID: $pid3 SV: $sv3"
result=$(echo 'test1234' | dcld tx model update-model-version --vid=$vid3 --pid=$pid3 --softwareVersion=$sv3 --softwareVersionValid=false --from=$vendor_admin_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"
echo "$result"

echo "Query updated Device Model Version with VID: $vid3 PID: $pid3 SV: $sv3"
result=$(dcld query model model-version --vid=$vid3 --pid=$pid3 --softwareVersion=$sv3)
check_response "$result" "\"softwareVersionValid\": false"
echo "$result"

echo "VendorAdmin deletes a model version for VID: $vid3 PID: $pid3 SV: $sv3"
result=$(echo 'test1234' | dcld tx model delete-model-version --vid=$vid3 --pid=$pid3 --softwareVersion=$sv3 --from=$vendor_admin_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"
echo "$result"

echo "Query deleted Device Model Version"
result=$(dcld query model model-version --vid=$vid3 --pid=$pid3 --softwareVersion=$sv3)
check_response "$result" "Not Found"
echo "$result"
