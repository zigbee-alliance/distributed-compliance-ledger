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

echo "Create Vendor account"
create_new_account vendor_account "Vendor"

echo "Create TestHouse account"
create_new_account test_house_account "TestHouse"

echo "Create ZBCertificationCenter account"
create_new_account zb_account "ZBCertificationCenter"

echo "Create other ZBCertificationCenter account"
create_new_account second_zb_account "ZBCertificationCenter"

# Body

vid=$RANDOM
pid=$RANDOM
echo "Add Model with VID: $vid PID: $pid"
result=$(echo "test1234" | dclcli tx modelinfo add-model --vid=$vid --pid=$pid --name="Device #1" --description="Device Description" --sku="SKU12FS" --softwareVersion="10123" --softwareVersionString="1.0b123"  --hardwareVersion="5123" --hardwareVersionString="5.1.23"  --cdVersionNumber="32" --from $vendor_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Add Testing Result for Model VID: $vid PID: $pid"
testing_result="http://first.place.com"
test_date="2020-11-24T10:00:00Z"
result=$(echo "test1234" | dclcli tx compliancetest add-test-result --vid=$vid --pid=$pid --test-result="$testing_result" --test-date="$test_date" --from $test_house_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get Certified Model with VID: ${vid} PID: ${pid} before compliance record was created"
result=$(dclcli query compliance certified-model --vid=$vid --pid=$pid --certification-type="zb")
check_response "$result" "\"value\": false"
echo "$result"

echo "Get Revoked Model with VID: ${vid} PID: ${pid} before compliance record was created"
result=$(dclcli query compliance revoked-model --vid=$vid --pid=$pid --certification-type="zb")
check_response "$result" "\"value\": false"
echo "$result"

echo "Certify Model with VID: $vid PID: $pid"
certification_date="2020-01-01T00:00:00Z"
certification_type="zb"
result=$(echo "test1234" | dclcli tx compliance certify-model --vid=$vid --pid=$pid --certification-type="$certification_type" --certification-date="$certification_date" --from $zb_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Certify Model with VID: $vid PID: $pid"
certification_date="2020-01-01T00:00:00Z"
certification_type="zb"
result=$(echo "test1234" | dclcli tx compliance certify-model --vid=$vid --pid=$pid --certification-type="$certification_type" --certification-date="$certification_date" --from $second_zb_account --yes)
check_response "$result" "\"success\": false"
echo "$result"

echo "Get Certified Model with VID: ${vid} PID: ${pid}"
result=$(dclcli query compliance certified-model --vid=$vid --pid=$pid --certification-type=$certification_type)
check_response "$result" "\"value\": true"
echo "$result"

echo "Get Compliance Info for Model with VID: ${vid} PID: ${pid}"
result=$(dclcli query compliance compliance-info --vid=$vid --pid=$pid --certification-type=$certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"state\": \"certified\""
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

echo "Revoke Certification for Model with VID: $vid PID: $pid"
revocation_date="2020-02-02T02:20:20Z"
revocation_reason="some reason"
result=$(echo "test1234" | dclcli tx compliance revoke-model --vid=$vid --pid=$pid --certification-type="$certification_type" --revocation-date="$revocation_date" --reason "$revocation_reason" --from $zb_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get Compliance Info for Model with VID: ${vid} PID: ${pid}"
result=$(dclcli query compliance compliance-info --vid=$vid --pid=$pid --certification-type=$certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"state\": \"revoked\""
check_response "$result" "\"date\": \"$revocation_date\""
check_response "$result" "\"reason\": \"$revocation_reason\""
check_response "$result" "\"certification_type\": \"$certification_type\""
check_response "$result" "\"history\""
echo "$result"

echo "Get Revoked Model with VID: ${vid} PID: ${pid}"
result=$(dclcli query compliance revoked-model --vid=$vid --pid=$pid --certification-type=$certification_type)
check_response "$result" "\"value\": true"
echo "$result"

echo "Get Certified Model with VID: ${vid} PID: ${pid}"
result=$(dclcli query compliance certified-model --vid=$vid --pid=$pid --certification-type=$certification_type)
check_response "$result" "\"value\": false"
echo "$result"

echo "Get All Revoked Models"
result=$(dclcli query compliance all-revoked-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"certification_type\": \"$certification_type\""
echo "$result"

echo "Again Certify Model with VID: $vid PID: $pid"
certification_date="2020-03-03T00:00:00Z"
result=$(echo "test1234" | dclcli tx compliance certify-model --vid=$vid --pid=$pid --certification-type="$certification_type" --certification-date="$certification_date" --from $zb_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get Compliance Info for Model with VID: ${vid} PID: ${pid}"
result=$(dclcli query compliance compliance-info --vid=$vid --pid=$pid --certification-type=$certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"state\": \"certified\""
check_response "$result" "\"date\": \"$certification_date\""
check_response "$result" "\"certification_type\": \"zb\""
echo "$result"

echo "Get Certified Model with VID: ${vid} PID: ${pid}"
result=$(dclcli query compliance certified-model --vid=$vid --pid=$pid --certification-type=$certification_type)
check_response "$result" "\"value\": true"
echo "$result"

echo "Get Revoked Model with VID: ${vid} PID: ${pid}"
result=$(dclcli query compliance revoked-model --vid=$vid --pid=$pid --certification-type=$certification_type)
check_response "$result" "\"value\": false"
echo "$result"

echo "Get All Compliance Infos"
result=$(dclcli query compliance all-compliance-info-records)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"state\": \"certified\""
echo "$result"
