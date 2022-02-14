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

# Body

pid=$RANDOM
sv=$RANDOM
svs=$RANDOM
certification_type_zb="zigbee"
certification_type_matter="matter"
provision_date="2020-02-02T02:20:20Z"
provision_reason="some reason"

test_divider

echo "Provision for uncertificate Model with VID: $vid PID: $pid for ZB"
result=$(echo "$passphrase" | dcld tx compliance provision-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$certification_type_zb" --provisionalDate="$provision_date" --reason "$provision_reason" --from $zb_account --yes)
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "Provision for uncertificate Model with VID: $vid PID: $pid for Matter"
result=$(echo "$passphrase" | dcld tx compliance provision-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$certification_type_matter" --provisionalDate="$provision_date" --reason "$provision_reason" --from $zb_account --yes)
check_response "$result" "\"code\": 0"
echo "$result"

test_divider


echo "ReProvision Model with VID: $vid PID: $pid  SV: ${sv} by different account"
result=$(echo "$passphrase" | dcld tx compliance provision-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$certification_type_zb" --provisionalDate="$provision_date" --reason "$provision_reason" --from $zb_account --yes)
check_response "$result" "\"code\": 305"
check_response "$result" "already in provisional state"
echo "$result"

test_divider

echo "ReProvision Model with VID: $vid PID: $pid  SV: ${sv} by same account"
result=$(echo "$passphrase" | dcld tx compliance provision-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$certification_type_zb" --provisionalDate="$provision_date" --reason "$provision_reason" --from $zb_account --yes)
check_response "$result" "\"code\": 305"
check_response "$result" "already in provisional state"
echo "$result"

test_divider

echo "Get Provisional Model with VID: ${vid} PID: ${pid} SV: ${sv} for zigbee certification"
result=$(dcld query compliance provisional-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$certification_type_zb)
check_response "$result" "\"value\": true"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"certification_type\": \"$certification_type_zb\""
echo "$result"

test_divider

echo "Get Provisional Model with VID: ${vid} PID: ${pid} SV: ${sv} for matter certification"
result=$(dcld query compliance provisional-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$certification_type_matter)
check_response "$result" "\"value\": true"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"certification_type\": \"$certification_type_matter\""
echo "$result"

test_divider

echo "Get Certified Model with VID: ${vid} PID: ${pid} before compliance record was created for ZB"
result=$(dcld query compliance certified-model --vid=$vid --pid=$pid  --softwareVersion=$sv --certificationType="$certification_type_zb")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"pid\": $pid"
response_does_not_contain "$result" "\"vid\": $vid"
response_does_not_contain "$result" "\"certification_type\": \"$certification_type_zb\""
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
result=$(dcld query compliance revoked-model --vid=$vid --pid=$pid  --softwareVersion=$sv --certificationType="$certification_type_zb")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"pid\": $pid"
response_does_not_contain "$result" "\"vid\": $vid"
response_does_not_contain "$result" "\"certification_type\": \"$certification_type_zb\""
echo "$result"

test_divider

echo "Get Revoked Model with VID: ${vid} PID: ${pid} and SV: ${sv} before compliance record was created for Matter"
result=$(dcld query compliance revoked-model --vid=$vid --pid=$pid  --softwareVersion=$sv --certificationType="$certification_type_matter")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"pid\": $pid"
response_does_not_contain "$result" "\"vid\": $vid"
response_does_not_contain "$result" "\"certification_type\": \"$certification_type_matter\""
echo "$result"

test_divider


echo "Get Compliance Info for Model with VID: ${vid} PID: ${pid} SV: ${sv} for ZB"
result=$(dcld query compliance compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$certification_type_zb)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"software_version_certification_status\": 1"
check_response "$result" "\"date\": \"$provision_date\""
check_response "$result" "\"reason\": \"$provision_reason\""
check_response "$result" "\"certification_type\": \"$certification_type_zb\""
check_response "$result" "\"history\""
echo "$result"

test_divider

echo "Get Compliance Info for Model with VID: ${vid} PID: ${pid} SV: ${sv} for Matter"
result=$(dcld query compliance compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$certification_type_matter)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"software_version_certification_status\": 1"
check_response "$result" "\"date\": \"$provision_date\""
check_response "$result" "\"reason\": \"$provision_reason\""
check_response "$result" "\"certification_type\": \"$certification_type_matter\""
check_response "$result" "\"history\""
echo "$result"

test_divider

echo "Get All Compliance Infos"
result=$(dcld query compliance all-compliance-info)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"software_version_certification_status\": 1"
check_response "$result" "\"date\": \"$provision_date\""
check_response "$result" "\"certification_type\": \"$certification_type_zb\""
check_response "$result" "\"certification_type\": \"$certification_type_matter\""
check_response "$result" "\"reason\": \"$provision_reason\""
echo "$result"

test_divider

echo "Get All Revoked Models"
result=$(dcld query compliance all-revoked-models)
check_response "$result" "\[\]"
response_does_not_contain "$result" "\"pid\": $pid"
response_does_not_contain "$result" "\"vid\": $vid"
echo "$result"

test_divider

echo "Get All Certified Models"
result=$(dcld query compliance all-certified-models)
check_response "$result" "\[\]"
response_does_not_contain "$result" "\"pid\": $pid"
response_does_not_contain "$result" "\"vid\": $vid"
echo "$result"

test_divider


echo "Get All Provisional Models"
result=$(dcld query compliance all-provisional-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"certification_type\": \"$certification_type_zb\""
check_response "$result" "\"certification_type\": \"$certification_type_matter\""
echo "$result"

test_divider

pid2=$RANDOM
sv2=$RANDOM
svs2=$RANDOM

echo "Add Model and a New Model Version with VID: $vid PID: $pid2 SV: $sv2"
create_model_and_version $vid $pid2 $sv2 $svs2 $vendor_account

test_divider

echo "Certify Model with VID: $vid PID: $pid2 for Matter"
certification_date="2021-02-02T02:20:19Z"
certification_reason="some reason 2"
result=$(echo "$passphrase" | dcld tx compliance certify-model --vid=$vid --pid=$pid2 --softwareVersion=$sv2 --softwareVersionString=$svs2 --certificationType="$certification_type_matter" --certificationDate="$certification_date" --reason "$certification_reason" --from $zb_account --yes)
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "Provision for already certified Model with VID: $vid PID: $pid2 for Matter"
result=$(echo "$passphrase" | dcld tx compliance provision-model --vid=$vid --pid=$pid2 --softwareVersion=$sv2 --softwareVersionString=$svs2 --certificationType="$certification_type_matter" --provisionalDate="$provision_date" --reason "$provision_reason" --from $zb_account --yes)
check_response "$result" "\"code\": 303"
check_response "$result" "already certified on the ledger"
echo "$result"

test_divider

pid3=$RANDOM
sv3=$RANDOM
svs3=$RANDOM

echo "Add Model and a New Model Version with VID: $vid PID: $pid3 SV: $sv3"
create_model_and_version $vid $pid3 $sv3 $svs3 $vendor_account

test_divider

echo "Revoke Certification for uncertificate Model with VID: $vid PID: $pid3"
revocation_date="2021-02-02T02:20:20Z"
revocation_reason="some reason 11"
result=$(echo "$passphrase" | dcld tx compliance revoke-model --vid=$vid --pid=$pid3 --softwareVersion=$sv3 --softwareVersionString=$svs3 --certificationType="$certification_type_zb" --revocationDate="$revocation_date" --reason "$revocation_reason" --from $zb_account --yes)
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "Provision for already revoked Model with VID: $vid PID: $pid3 for Matter"
result=$(echo "$passphrase" | dcld tx compliance provision-model --vid=$vid --pid=$pid3 --softwareVersion=$sv3 --softwareVersionString=$svs3 --certificationType="$certification_type_zb" --provisionalDate="$provision_date" --reason "$provision_reason" --from $zb_account --yes)
check_response "$result" "\"code\": 304"
check_response "$result" "already revoked on the ledger"
echo "$result"

test_divider

echo "Add Model and a New Model Version with VID: $vid PID: $pid SV: $sv"
create_model_and_version $vid $pid $sv $svs $vendor_account

test_divider

echo "Certify provisioned Model with VID: $vid PID: $pid"
certification_date="2021-02-02T02:20:19Z"
certification_reason="some reason 2"
result=$(echo "$passphrase" | dcld tx compliance certify-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$certification_type_zb" --certificationDate="$certification_date" --reason "$certification_reason" --from $zb_account --yes)
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "Revoke provisioned Model with VID: $vid PID: $pid"
revocation_date="2021-02-02T02:20:20Z"
revocation_reason="some reason 22"
result=$(echo "$passphrase" | dcld tx compliance revoke-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$certification_type_matter" --revocationDate="$revocation_date" --reason "$revocation_reason" --from $zb_account --yes)
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "Get Provisional Model with VID: ${vid} PID: ${pid} SV: ${sv} for zigbee certification"
result=$(dcld query compliance provisional-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$certification_type_zb)
check_response "$result" "\"value\": false"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"certification_type\": \"$certification_type_zb\""
echo "$result"

test_divider

echo "Get Provisional Model with VID: ${vid} PID: ${pid} SV: ${sv} for matter certification"
result=$(dcld query compliance provisional-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$certification_type_matter)
check_response "$result" "\"value\": false"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"certification_type\": \"$certification_type_matter\""
echo "$result"

test_divider

echo "Get Certified Model with VID: ${vid} PID: ${pid} for ZB"
result=$(dcld query compliance certified-model --vid=$vid --pid=$pid  --softwareVersion=$sv --certificationType="$certification_type_zb")
check_response "$result" "true"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"certification_type\": \"$certification_type_zb\""
echo "$result"

test_divider

echo "Get Certified Model with VID: ${vid} PID: ${pid} for Matter"
result=$(dcld query compliance certified-model --vid=$vid --pid=$pid  --softwareVersion=$sv --certificationType="$certification_type_matter")
check_response "$result" "false"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"certification_type\": \"$certification_type_matter\""
echo "$result"

test_divider


echo "Get Revoked Model with VID: ${vid} PID: ${pid} and SV: ${sv} for ZB"
result=$(dcld query compliance revoked-model --vid=$vid --pid=$pid  --softwareVersion=$sv --certificationType="$certification_type_zb")
check_response "$result" "false"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"certification_type\": \"$certification_type_zb\""
echo "$result"

test_divider

echo "Get Revoked Model with VID: ${vid} PID: ${pid} and SV: ${sv} for Matter"
result=$(dcld query compliance revoked-model --vid=$vid --pid=$pid  --softwareVersion=$sv --certificationType="$certification_type_matter")
check_response "$result" "true"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"certification_type\": \"$certification_type_matter\""
echo "$result"

test_divider


echo "Get Compliance Info for Model with VID: ${vid} PID: ${pid} SV: ${sv} for ZB"
result=$(dcld query compliance compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$certification_type_zb)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"software_version_certification_status\": 2"
check_response "$result" "\"date\": \"$certification_date\""
check_response "$result" "\"reason\": \"$certification_reason\""
check_response "$result" "\"certification_type\": \"$certification_type_zb\""
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
check_response "$result" "\"software_version_certification_status\": 2"
check_response "$result" "\"software_version_certification_status\": 3"
check_response "$result" "\"date\": \"$revocation_date\""
check_response "$result" "\"date\": \"$certification_date\""
check_response "$result" "\"certification_type\": \"$certification_type_zb\""
check_response "$result" "\"certification_type\": \"$certification_type_matter\""
check_response "$result" "\"reason\": \"$revocation_reason\""
check_response "$result" "\"reason\": \"$certification_reason\""
echo "$result"

test_divider

echo "Get All Revoked Models"
result=$(dcld query compliance all-revoked-models)
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"certification_type\": \"$certification_type_matter\""
echo "$result"

test_divider

echo "Get All Certified Models"
result=$(dcld query compliance all-certified-models)
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"certification_type\": \"$certification_type_zb\""
echo "$result"

test_divider


echo "Get All Provisional Models"
result=$(dcld query compliance all-provisional-models)
check_response "$result" "\[\]"
response_does_not_contain "$result" "\"pid\": $pid"
response_does_not_contain "$result" "\"vid\": $vid"
echo "$result"

test_divider
