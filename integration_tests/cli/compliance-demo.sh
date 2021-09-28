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

test_divider

echo "Create TestHouse account"
create_new_account test_house_account "TestHouse"

test_divider

echo "Create ZBCertificationCenter account"
create_new_account zb_account "ZBCertificationCenter"

test_divider

echo "Create other ZBCertificationCenter account"
create_new_account second_zb_account "ZBCertificationCenter"

test_divider

# Body


pid=$RANDOM
echo "Add Model with VID: $vid PID: $pid"

result=$(echo "test1234" | dclcli tx model add-model --vid=$vid --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0  --from $vendor_account --yes)
echo $result
check_response "$result" "\"success\": true"

test_divider

sv=$RANDOM
echo "Add Model Version with VID: $vid PID: $pid SV: $sv"
result=$(echo 'test1234' | dclcli tx modelversion add-model-version --cdVersionNumber=1 --maxApplicableSoftwareVersion=10 --minApplicableSoftwareVersion=1 --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=1 --from=$vendor_account --yes)
echo $result
check_response "$result" "\"success\": true"

test_divider

echo "Add Testing Result for Model VID: $vid PID: $pid SV: $sv"
testing_result="http://first.place.com"
test_date="2020-11-24T10:00:00Z"
result=$(echo "test1234" | dclcli tx compliancetest add-test-result --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString="1.0" --test-result="$testing_result" --test-date="$test_date" --from $test_house_account --yes)
echo $result
check_response "$result" "\"success\": true"

test_divider

echo "Get Certified Model with VID: ${vid} PID: ${pid} SV: ${sv} before compliance record was created"
result=$(dclcli query compliance certified-model --vid=$vid --pid=$pid --softwareVersion=$sv --certification-type="zb")
echo $result
check_response "$result" "\"value\": false"

test_divider

echo "Get Revoked Model with VID: ${vid} PID: ${pid} SV: ${sv} before compliance record was created"
result=$(dclcli query compliance revoked-model --vid=$vid --pid=$pid --softwareVersion=$sv --certification-type="zb")
echo "$result"
check_response "$result" "\"value\": false"

test_divider

echo "Certify Model with VID: $vid PID: $pid  SV: ${sv} "
certification_date="2020-01-01T00:00:00Z"
certification_type="zb"
result=$(echo "test1234" | dclcli tx compliance certify-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=1.0 --certification-type="$certification_type" --certification-date="$certification_date" --from $zb_account --yes)
echo "$result"
check_response "$result" "\"success\": true"

test_divider

echo "Certify Model with VID: $vid PID: $pid  SV: ${sv} "
certification_date="2020-01-01T00:00:00Z"
certification_type="zb"
result=$(echo "test1234" | dclcli tx compliance certify-model --vid=$vid --pid=$pid  --softwareVersion=$sv --softwareVersionString=1.0 --certification-type="$certification_type" --certification-date="$certification_date" --from $second_zb_account --yes)
check_response "$result" "\"success\": false"
echo "$result"

echo "Get Certified Model with VID: ${vid} PID: ${pid} SV: ${sv} "
result=$(dclcli query compliance certified-model --vid=$vid --pid=$pid --softwareVersion=$sv --certification-type=$certification_type)
check_response "$result" "\"value\": true"
echo "$result"

echo "Get Compliance Info for Model with VID: ${vid} PID: ${pid} SV: ${sv} "
result=$(dclcli query compliance compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certification-type=$certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"softwareVersionCertificationStatus\": 2"
check_response "$result" "\"date\": \"$certification_date\""
check_response "$result" "\"certification_type\": \"$certification_type\""
echo "$result"

echo "Get All Certified Models"
result=$(dclcli query compliance all-certified-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"certification_type\": \"$certification_type\""
echo "$result"

echo "Get All Compliance Info Recordss"
result=$(dclcli query compliance all-compliance-info-records)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"date\": \"$certification_date\""
echo "$result"

echo "Revoke Certification for Model with VID: $vid PID: $pid SV: ${sv} "
revocation_date="2020-02-02T02:20:20Z"
revocation_reason="some reason"
result=$(echo "test1234" | dclcli tx compliance revoke-model --vid=$vid --pid=$pid --softwareVersion=$sv --certification-type="$certification_type" --revocation-date="$revocation_date" --reason "$revocation_reason" --from $zb_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get Compliance Info for Model with VID: ${vid} PID: ${pid} SV: ${sv} "
result=$(dclcli query compliance compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certification-type=$certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"softwareVersionCertificationStatus\": 3"
check_response "$result" "\"date\": \"$revocation_date\""
check_response "$result" "\"reason\": \"$revocation_reason\""
check_response "$result" "\"certification_type\": \"$certification_type\""
check_response "$result" "\"history\""
echo "$result"

echo "Get Revoked Model with VID: ${vid} PID: ${pid} SV: ${sv} "
result=$(dclcli query compliance revoked-model --vid=$vid --pid=$pid --softwareVersion=$sv --certification-type=$certification_type)
check_response "$result" "\"value\": true"
echo "$result"

echo "Get Certified Model with VID: ${vid} PID: ${pid} SV: ${sv} "
result=$(dclcli query compliance certified-model --vid=$vid --pid=$pid --softwareVersion=$sv --certification-type=$certification_type)
check_response "$result" "\"value\": false"
echo "$result"

echo "Get All Revoked Models"
result=$(dclcli query compliance all-revoked-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"certification_type\": \"$certification_type\""
echo "$result"

echo "Again Certify Model with VID: $vid PID: $pid SV: ${sv}"
certification_date="2020-03-03T00:00:00Z"
result=$(echo "test1234" | dclcli tx compliance certify-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=1.0 --certification-type="$certification_type" --certification-date="$certification_date" --from $zb_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get Compliance Info for Model with VID: ${vid} PID: ${pid} SV: ${sv}"
result=$(dclcli query compliance compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certification-type=$certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"softwareVersionCertificationStatus\": 2"
check_response "$result" "\"date\": \"$certification_date\""
check_response "$result" "\"certification_type\": \"zb\""
echo "$result"

echo "Get Certified Model with VID: ${vid} PID: ${pid} SV: ${sv}"
result=$(dclcli query compliance certified-model --vid=$vid --pid=$pid --softwareVersion=$sv --certification-type=$certification_type)
check_response "$result" "\"value\": true"
echo "$result"

echo "Get Revoked Model with VID: ${vid} PID: ${pid} SV: ${sv}"
result=$(dclcli query compliance revoked-model --vid=$vid --pid=$pid --softwareVersion=$sv --certification-type=$certification_type)
check_response "$result" "\"value\": false"
echo "$result"

echo "Get All Compliance Infos"
result=$(dclcli query compliance all-compliance-info-records)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"softwareVersionCertificationStatus\": 2"
echo "$result"
