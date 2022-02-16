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

echo "Create CertificationCenter account"
create_new_account zb_account "CertificationCenter"

echo "Create other CertificationCenter account"
create_new_account second_zb_account "CertificationCenter"

certification_type="zigbee"
certification_type_matter="matter"

# Body

pid=$RANDOM
sv=$RANDOM
svs=$RANDOM

echo "Add Model and a New Model Version with VID: $vid PID: $pid SV: $sv"
create_model_and_version $vid $pid $sv $svs $vendor_account

test_divider

echo "Revoke Certification for uncertificate Model with VID: $vid PID: $pid"
revocation_date="2020-02-02T02:20:20Z"
revocation_reason="some reason"
result=$(echo "$passphrase" | dcld tx compliance revoke-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$certification_type" --revocationDate="$revocation_date" --reason "$revocation_reason" --from $zb_account --yes)
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "ReRevoke Model with VID: $vid PID: $pid  SV: ${sv} by different account"
result=$(echo "$passphrase" | dcld tx compliance revoke-model --vid=$vid --pid=$pid  --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$certification_type" --revocationDate="$revocation_date" --from $second_zb_account --yes)
check_response "$result" "\"code\": 304"
check_response "$result" "already revoked on the ledger"
echo "$result"

test_divider

echo "ReRevoke Model with VID: $vid PID: $pid  SV: ${sv} by same account"
result=$(echo "$passphrase" | dcld tx compliance revoke-model --vid=$vid --pid=$pid  --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$certification_type" --revocationDate="$revocation_date" --from $zb_account --yes)
check_response "$result" "\"code\": 304"
check_response "$result" "already revoked on the ledger"
echo "$result"

test_divider

echo "Revoke Certification for uncertificate Model with VID: $vid PID: $pid for matter"
revocation_date="2020-02-02T02:20:20Z"
revocation_reason="some reason"
result=$(echo "$passphrase" | dcld tx compliance revoke-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$certification_type_matter" --revocationDate="$revocation_date" --reason "$revocation_reason" --from $zb_account --yes)
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "Get Certified Model with VID: ${vid} PID: ${pid} before compliance record was created for ZB"
result=$(dcld query compliance certified-model --vid=$vid --pid=$pid  --softwareVersion=$sv --certificationType="$certification_type")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"pid\": $pid"
response_does_not_contain "$result" "\"vid\": $vid"
response_does_not_contain "$result" "\"certification_type\": \"$certification_type\""
echo "$result"

test_divider

echo "Get Certified Model with VID: ${vid} PID: ${pid} before compliance record was created for Matter"
result=$(dcld query compliance certified-model --vid=$vid --pid=$pid  --softwareVersion=$sv --certificationType="$certification_type_matter")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"pid\": $pid"
response_does_not_contain "$result" "\"vid\": $vid"
response_does_not_contain "$result" "\"certification_type\": \"$certification_type_matter\""
echo "$result"

test_divider


echo "Get Revoked Model with VID: ${vid} PID: ${pid} and SV: ${sv} before compliance record was created for ZB"
result=$(dcld query compliance revoked-model --vid=$vid --pid=$pid  --softwareVersion=$sv --certificationType="$certification_type")
check_response "$result" "\"value\": true"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"certification_type\": \"$certification_type\""
echo "$result"

test_divider

echo "Get Revoked Model with VID: ${vid} PID: ${pid} and SV: ${sv} before compliance record was created for Matter"
result=$(dcld query compliance revoked-model --vid=$vid --pid=$pid  --softwareVersion=$sv --certificationType="$certification_type_matter")
check_response "$result" "\"value\": true"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"certification_type\": \"$certification_type_matter\""
echo "$result"

test_divider


echo "Get Compliance Info for Model with VID: ${vid} PID: ${pid} SV: ${sv} for ZB"
result=$(dcld query compliance compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"software_version_certification_status\": 3"
check_response "$result" "\"date\": \"$revocation_date\""
check_response "$result" "\"reason\": \"$revocation_reason\""
check_response "$result" "\"certification_type\": \"$certification_type\""
check_response "$result" "\"history\""
echo "$result"

test_divider

echo "Get Compliance Info for Model with VID: ${vid} PID: ${pid} SV: ${sv} for Matter"
result=$(dcld query compliance compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$certification_type_matter)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"software_version_certification_status\": 3"
check_response "$result" "\"date\": \"$revocation_date\""
check_response "$result" "\"reason\": \"$revocation_reason\""
check_response "$result" "\"certification_type\": \"$certification_type_matter\""
check_response "$result" "\"history\""
echo "$result"

test_divider

echo "Get All Compliance Infos"
result=$(dcld query compliance all-compliance-info)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"software_version_certification_status\": 3"
check_response "$result" "\"date\": \"$revocation_date\""
check_response "$result" "\"certification_type\": \"$certification_type\""
check_response "$result" "\"certification_type\": \"$certification_type_matter\""
check_response "$result" "\"reason\": \"$revocation_reason\""
echo "$result"

test_divider

echo "Get All Revoked Models"
result=$(dcld query compliance all-revoked-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"certification_type\": \"$certification_type\""
check_response "$result" "\"certification_type\": \"$certification_type_matter\""
echo "$result"

test_divider

echo "Get All Certified Models"
result=$(dcld query compliance all-certified-models)
check_response "$result" "\[\]"
response_does_not_contain "$result" "\"pid\": $pid"
response_does_not_contain "$result" "\"vid\": $vid"
echo "$result"

test_divider

echo "Get Provisional Model with VID: ${vid} PID: ${pid} SV: ${sv} for zigbee certification"
result=$(dcld query compliance provisional-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$certification_type)
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"pid\": $pid"
response_does_not_contain "$result" "\"vid\": $vid"
echo "$result"

test_divider

echo "Get All Provisional Models empty"
result=$(dcld query compliance all-provisional-models)
check_response "$result" "\[\]"
response_does_not_contain "$result" "\"pid\": $pid"
response_does_not_contain "$result" "\"vid\": $vid"
echo "$result"

test_divider

certification_reason="some reason 2"

echo "Certify revoked Model with VID: $vid PID: $pid from the past"
certification_date_past="2020-02-02T02:20:19Z"
result=$(echo "$passphrase" | dcld tx compliance certify-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$certification_type" --certificationDate="$certification_date_past" --reason "$certification_reason" --from $zb_account --yes)
check_response "$result" "\"code\": 302"
check_response "$result" "must be after"
echo "$result"

test_divider

echo "Certify revoked Model with VID: $vid PID: $pid for ZB"
certification_date="2020-02-02T02:20:21Z"
result=$(echo "$passphrase" | dcld tx compliance certify-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$certification_type" --certificationDate="$certification_date" --reason "$certification_reason" --from $zb_account --yes)
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "Get Revoked Model with VID: ${vid} PID: ${pid} for ZB"
result=$(dcld query compliance revoked-model --vid=$vid --pid=$pid  --softwareVersion=$sv --certificationType="$certification_type")
check_response "$result" "\"value\": false"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"certification_type\": \"$certification_type\""
echo "$result"

test_divider

echo "Get Revoked Model with VID: ${vid} PID: ${pid} for Matter"
result=$(dcld query compliance revoked-model --vid=$vid --pid=$pid  --softwareVersion=$sv --certificationType="$certification_type_matter")
check_response "$result" "\"value\": true"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"certification_type\": \"$certification_type_matter\""
echo "$result"

test_divider


echo "Get Certified Model with VID: ${vid} PID: ${pid} and SV: ${sv} for ZB"
result=$(dcld query compliance certified-model --vid=$vid --pid=$pid  --softwareVersion=$sv --certificationType="$certification_type")
check_response "$result" "\"value\": true"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"certification_type\": \"$certification_type\""
echo "$result"

test_divider

echo "Get Certified Model with VID: ${vid} PID: ${pid} and SV: ${sv} for Matter"
result=$(dcld query compliance certified-model --vid=$vid --pid=$pid  --softwareVersion=$sv --certificationType="$certification_type_matter")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"pid\": $pid"
response_does_not_contain "$result" "\"vid\": $vid"
response_does_not_contain "$result" "\"certification_type\": \"$certification_type_matter\""
echo "$result"

test_divider

echo "Get Compliance Info for Model with VID: ${vid} PID: ${pid} SV: ${sv} "
result=$(dcld query compliance compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"software_version_certification_status\": 2"
check_response "$result" "\"date\": \"$certification_date\""
check_response "$result" "\"reason\": \"$certification_reason\""
check_response "$result" "\"certification_type\": \"$certification_type\""
check_response "$result" "\"history\""
echo "$result"

test_divider

echo "Get All Compliance Infos"
result=$(dcld query compliance all-compliance-info)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"software_version_certification_status\": 2"
check_response "$result" "\"software_version_certification_status\": 3"
check_response "$result" "\"date\": \"$certification_date\""
check_response "$result" "\"date\": \"$revocation_date\""
check_response "$result" "\"reason\": \"$certification_reason\""
check_response "$result" "\"reason\": \"$revocation_reason\""
check_response "$result" "\"certification_type\": \"$certification_type\""
check_response "$result" "\"certification_type\": \"$certification_type_matter\""
echo "$result"

test_divider

echo "Get All Revoked Models"
result=$(dcld query compliance all-revoked-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"certification_type\": \"$certification_type_matter\""
response_does_not_contain "$result" "\"certification_type\": \"$certification_type\""
echo "$result"

test_divider

echo "Get All Certified Models"
result=$(dcld query compliance all-certified-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"certification_type\": \"$certification_type\""
response_does_not_contain "$result" "\"certification_type\": \"$certification_type_matter\""
echo "$result"

test_divider

echo "Get Provisional Model with VID: ${vid} PID: ${pid} SV: ${sv} for zigbee certification"
result=$(dcld query compliance provisional-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$certification_type)
check_response "$result" "\"value\": false"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"vid\": $vid"
echo "$result"

test_divider

echo "Get All Provisional Models empty"
result=$(dcld query compliance all-provisional-models)
check_response "$result" "\[\]"
response_does_not_contain "$result" "\"pid\": $pid"
response_does_not_contain "$result" "\"vid\": $vid"
echo "$result"

test_divider