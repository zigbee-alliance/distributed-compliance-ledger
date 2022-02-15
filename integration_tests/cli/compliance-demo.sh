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

echo "Create CertificationCenter account"
create_new_account zb_account "CertificationCenter"

test_divider

echo "Create other CertificationCenter account"
create_new_account second_zb_account "CertificationCenter"

test_divider

# Body

pid=$RANDOM
sv=$RANDOM
svs=$RANDOM
certification_date="2020-01-01T00:00:01Z"
zigbee_certification_type="zigbee"
matter_certification_type="matter"

echo "Certify unknown Model with VID: $vid PID: $pid  SV: ${sv} with zigbee certification"
result=$(echo "$passphrase" | dcld tx compliance certify-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$zigbee_certification_type" --certificationDate="$certification_date" --from $zb_account --yes)
echo "$result"
check_response "$result" "\"code\": 517"
check_response "$result" "No model version"

test_divider

echo "Add Model with VID: $vid PID: $pid"
result=$(echo "$passphrase" | dcld tx model add-model --vid=$vid --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0  --from $vendor_account --yes)
echo $result
check_response "$result" "\"code\": 0"

test_divider

echo "Certify unknown Model Version with VID: $vid PID: $pid  SV: ${sv} with zigbee certification"
result=$(echo "$passphrase" | dcld tx compliance certify-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$zigbee_certification_type" --certificationDate="$certification_date" --from $zb_account --yes)
echo "$result"
check_response "$result" "\"code\": 517"
check_response "$result" "No model version"

test_divider

echo "Add Model Version with VID: $vid PID: $pid SV: $sv SoftwareVersionString:$svs"
result=$(echo '$passphrase' | dcld tx model add-model-version --cdVersionNumber=1 --maxApplicableSoftwareVersion=10 --minApplicableSoftwareVersion=1 --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --from=$vendor_account --yes)
echo $result
check_response "$result" "\"code\": 0"

test_divider

echo "Get Compliance Info with VID: ${vid} PID: ${pid} SV: ${sv} before compliance record was created"
result=$(dcld query compliance compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType="zigbee")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"pid\": $pid"
response_does_not_contain "$result" "\"vid\": $vid"
echo $result

test_divider

echo "Get Certified Model with VID: ${vid} PID: ${pid} SV: ${sv} before compliance record was created"
result=$(dcld query compliance certified-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType="zigbee")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"pid\": $pid"
response_does_not_contain "$result" "\"vid\": $vid"
echo $result

test_divider

echo "Get Revoked Model with VID: ${vid} PID: ${pid} SV: ${sv} before compliance record was created"
result=$(dcld query compliance revoked-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType="zigbee")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"pid\": $pid"
response_does_not_contain "$result" "\"vid\": $vid"
echo "$result"

test_divider

echo "Get Provisional Model with VID: ${vid} PID: ${pid} SV: ${sv} before compliance record was created"
result=$(dcld query compliance provisional-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType="zigbee")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"pid\": $pid"
response_does_not_contain "$result" "\"vid\": $vid"
echo "$result"

test_divider

echo "Get All Compliance Info empty"
result=$(dcld query compliance all-compliance-info)
check_response "$result" "\[\]"
response_does_not_contain "$result" "\"pid\": $pid"
response_does_not_contain "$result" "\"vid\": $vid"
echo "$result"

test_divider

echo "Get All Certified Models empty"
result=$(dcld query compliance all-certified-models)
check_response "$result" "\[\]"
response_does_not_contain "$result" "\"pid\": $pid"
response_does_not_contain "$result" "\"vid\": $vid"
echo "$result"

test_divider

echo "Get All Revoked Models empty"
result=$(dcld query compliance all-revoked-models)
check_response "$result" "\[\]"
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

invalid_svs=$RANDOM
echo "Certify Model with VID: $vid PID: $pid  SV: ${sv} with zigbee certification and invalid SoftwareVersionString: $invalid_svs"
result=$(echo "$passphrase" | dcld tx compliance certify-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$invalid_svs --certificationType="$zigbee_certification_type" --certificationDate="$certification_date" --from $zb_account --yes)
check_response "$result" "\"code\": 306"
check_response "$result" "ledger does not have matching softwareVersionString=$invalid_svs: model version does not match"

test_divider

echo "Certify Model with VID: $vid PID: $pid  SV: ${sv} with zigbee certification"
result=$(echo "$passphrase" | dcld tx compliance certify-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$zigbee_certification_type" --certificationDate="$certification_date" --from $zb_account --yes)
echo "$result"
check_response "$result" "\"code\": 0"

test_divider

echo "Certify Model with VID: $vid PID: $pid  SV: ${sv} with matter certification"
echo "dcld tx compliance certify-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$matter_certification_type" --certificationDate="$certification_date" --from $zb_account --yes"
result=$(echo "$passphrase" | dcld tx compliance certify-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$matter_certification_type" --certificationDate="$certification_date" --from $zb_account --yes)
echo "$result"
check_response "$result" "\"code\": 0"

test_divider

echo "ReCertify Model with VID: $vid PID: $pid  SV: ${sv} by different account"
zigbee_certification_type="zigbee"
result=$(echo "$passphrase" | dcld tx compliance certify-model --vid=$vid --pid=$pid  --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$zigbee_certification_type" --certificationDate="$certification_date" --from $second_zb_account --yes)
check_response "$result" "\"code\": 303"
check_response "$result" "already certified on the ledger"
echo "$result"

test_divider

echo "ReCertify Model with VID: $vid PID: $pid  SV: ${sv} by same account"
zigbee_certification_type="zigbee"
result=$(echo "$passphrase" | dcld tx compliance certify-model --vid=$vid --pid=$pid  --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$zigbee_certification_type" --certificationDate="$certification_date" --from $zb_account --yes)
check_response "$result" "\"code\": 303"
check_response "$result" "already certified on the ledger"
echo "$result"

test_divider

echo "Get Certified Model with VID: ${vid} PID: ${pid} SV: ${sv} for matter certification"
result=$(dcld query compliance certified-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$matter_certification_type)
check_response "$result" "\"value\": true"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"vid\": $vid"
echo "$result"

test_divider

echo "Get Certified Model with VID: ${vid} PID: ${pid} SV: ${sv} for zigbee certification"
result=$(dcld query compliance certified-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$zigbee_certification_type)
check_response "$result" "\"value\": true"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"vid\": $vid"
echo "$result"

test_divider

echo "Get Revoked Model with VID: ${vid} PID: ${pid} SV: ${sv} for matter certification"
result=$(dcld query compliance revoked-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$matter_certification_type)
check_response "$result" "\"value\": false"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"vid\": $vid"
echo "$result"

test_divider

echo "Get Revoked Model with VID: ${vid} PID: ${pid} SV: ${sv} for zigbee certification"
result=$(dcld query compliance revoked-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$zigbee_certification_type)
check_response "$result" "\"value\": false"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"vid\": $vid"
echo "$result"

test_divider

echo "Get Provisional Model with VID: ${vid} PID: ${pid} SV: ${sv} for matter certification"
result=$(dcld query compliance provisional-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$matter_certification_type)
check_response "$result" "\"value\": false"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"vid\": $vid"
echo "$result"

test_divider

echo "Get Provisional Model with VID: ${vid} PID: ${pid} SV: ${sv} for zigbee certification"
result=$(dcld query compliance provisional-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$zigbee_certification_type)
check_response "$result" "\"value\": false"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"vid\": $vid"
echo "$result"

test_divider

echo "Get Compliance Info for Model with VID: ${vid} PID: ${pid} SV: ${sv} for $zigbee_certification_type"
result=$(dcld query compliance compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$zigbee_certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"software_version_certification_status\": 2"
check_response "$result" "\"date\": \"$certification_date\""
check_response "$result" "\"certification_type\": \"$zigbee_certification_type\""
echo "$result"

test_divider

echo "Get Compliance Info for Model with VID: ${vid} PID: ${pid} SV: ${sv} for $matter_certification_type"
result=$(dcld query compliance compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$matter_certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"software_version_certification_status\": 2"
check_response "$result" "\"date\": \"$certification_date\""
check_response "$result" "\"certification_type\": \"$matter_certification_type\""
echo "$result"

test_divider

echo "Get All Certified Models"
result=$(dcld query compliance all-certified-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"certification_type\": \"$zigbee_certification_type\""
check_response "$result" "\"certification_type\": \"$matter_certification_type\""
echo "$result"

test_divider

echo "Get All Revoked Models empty"
result=$(dcld query compliance all-revoked-models)
check_response "$result" "\[\]"
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

echo "Get All Compliance Info Recordss"
result=$(dcld query compliance all-compliance-info)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"certification_type\": \"$zigbee_certification_type\""
check_response "$result" "\"certification_type\": \"$matter_certification_type\""
check_response "$result" "\"date\": \"$certification_date\""
echo "$result"

test_divider

revocation_reason="some reason"

echo "Revoke Certification for Model with VID: $vid PID: $pid SV: ${sv} from the past"
revocation_date_past="2020-01-01T00:00:00Z"
result=$(echo "$passphrase" | dcld tx compliance revoke-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$zigbee_certification_type" --revocationDate="$revocation_date_past" --reason "$revocation_reason" --from $zb_account --yes)
check_response "$result" "\"code\": 302"
check_response "$result" "must be after"
echo "$result"

test_divider


echo "Revoke Certification for Model with VID: $vid PID: $pid SV: ${sv} "
revocation_date="2020-02-02T02:20:20Z"
result=$(echo "$passphrase" | dcld tx compliance revoke-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$zigbee_certification_type" --revocationDate="$revocation_date" --reason "$revocation_reason" --from $zb_account --yes)
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "Get Compliance Info for Model with VID: ${vid} PID: ${pid} SV: ${sv} "
result=$(dcld query compliance compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$zigbee_certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"software_version_certification_status\": 3"
check_response "$result" "\"date\": \"$revocation_date\""
check_response "$result" "\"reason\": \"$revocation_reason\""
check_response "$result" "\"certification_type\": \"$zigbee_certification_type\""
check_response "$result" "\"history\""
echo "$result"

test_divider

echo "Get Revoked Model with VID: ${vid} PID: ${pid} SV: ${sv} "
result=$(dcld query compliance revoked-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$zigbee_certification_type)
check_response "$result" "\"value\": true"
echo "$result"

test_divider

echo "Get Certified Model with VID: ${vid} PID: ${pid} SV: ${sv} "
result=$(dcld query compliance certified-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$zigbee_certification_type)
check_response "$result" "\"value\": false"
echo "$result"

test_divider

echo "Get Provisional Model with VID: ${vid} PID: ${pid} SV: ${sv} for zigbee certification"
result=$(dcld query compliance provisional-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$zigbee_certification_type)
check_response "$result" "\"value\": false"
echo "$result"

test_divider

echo "Get All Revoked Models"
result=$(dcld query compliance all-revoked-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"certification_type\": \"$zigbee_certification_type\""
echo "$result"

test_divider

echo "Get All Certified Models"
result=$(dcld query compliance all-certified-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"certification_type\": \"$matter_certification_type\""
response_does_not_contain "$result" "\"certification_type\": \"$zigbee_certification_type\""
echo "$result"

test_divider


echo "Again Certify Model with VID: $vid PID: $pid SV: ${sv}"
certification_date="2020-03-03T00:00:00Z"
result=$(echo "$passphrase" | dcld tx compliance certify-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$zigbee_certification_type" --certificationDate="$certification_date" --from $zb_account --yes)
check_response "$result" "\"code\": 0"
echo "$result"

echo "Get Compliance Info for Model with VID: ${vid} PID: ${pid} SV: ${sv}"
result=$(dcld query compliance compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$zigbee_certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"software_version_certification_status\": 2"
check_response "$result" "\"date\": \"$certification_date\""
check_response "$result" "\"certification_type\": \"zigbee\""
echo "$result"

echo "Get Certified Model with VID: ${vid} PID: ${pid} SV: ${sv}"
result=$(dcld query compliance certified-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$zigbee_certification_type)
check_response "$result" "\"value\": true"
echo "$result"

echo "Get Revoked Model with VID: ${vid} PID: ${pid} SV: ${sv}"
result=$(dcld query compliance revoked-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$zigbee_certification_type)
check_response "$result" "\"value\": false"
echo "$result"

echo "Get All Compliance Infos"
result=$(dcld query compliance all-compliance-info)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"software_version_certification_status\": 2"
echo "$result"

echo "Get All Revoked Models"
result=$(dcld query compliance all-revoked-models)
check_response "$result" "\[\]"
response_does_not_contain "$result" "\"pid\": $pid"
response_does_not_contain "$result" "\"vid\": $vid"
echo "$result"

test_divider

echo "Get All Certified Models"
result=$(dcld query compliance all-certified-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"certification_type\": \"$matter_certification_type\""
check_response "$result" "\"certification_type\": \"$zigbee_certification_type\""
echo "$result"

test_divider