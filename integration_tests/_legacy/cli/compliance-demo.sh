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

# FIXME issue 99: enable once implemented
exit 0

# Preparation of Actors
vid=$RANDOM
vendor_account=vendor_account_$vid
echo "Create Vendor account - $vendor_account"
create_new_vendor_account $vendor_account $vid

test_divider

echo "Create TestHouse account"
create_new_account test_house_account "TestHouse"

test_divider

echo "Create CertificationCenter account"
create_new_account zb_account "CertificationCenter"

test_divider

echo "Create other CertificationCenter account"
create_new_account second_zb_account "CertificationCenter"

test_divider

# Body


pid=$RANDOM
echo "Add Model with VID: $vid PID: $pid"

result=$(echo "test1234" | dcld tx model add-model --vid=$vid --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0  --from $vendor_account --yes)
echo $result
check_response "$result" "\"success\": true"

test_divider

sv=$RANDOM
svs=$RANDOM
echo "Add Model Version with VID: $vid PID: $pid SV: $sv SoftwareVersionString:$svs"
result=$(echo 'test1234' | dcld tx model add-model-version --cdVersionNumber=1 --maxApplicableSoftwareVersion=10 --minApplicableSoftwareVersion=1 --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --from=$vendor_account --yes)
echo $result
check_response "$result" "\"success\": true"

test_divider

invalid_svs=$RANDOM
testing_result="http://first.place.com"
test_date="2020-11-24T10:00:00Z"
echo "Add Testing Result for Model VID: $vid PID: $pid SV: $sv and invalid SoftwareVersionString: $invalid_svs"
result=$(echo 'test1234' | dcld tx compliancetest add-test-result --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$invalid_svs --test-result="$testing_result" --test-date="$test_date" --from $test_house_account --yes)
check_response "$result" "\"success\": false"
check_response "$result" "ledger does not have  matching softwareVersionString=$invalid_svs"

test_divider

echo "Add Testing Result for Model VID: $vid PID: $pid SV: $sv SoftwareVersionString:$svs"
result=$(echo "test1234" | dcld tx compliancetest add-test-result --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --test-result="$testing_result" --test-date="$test_date" --from $test_house_account --yes)
echo $result
check_response "$result" "\"success\": true"

test_divider

echo "Get Certified Model with VID: ${vid} PID: ${pid} SV: ${sv} before compliance record was created"
result=$(dcld query compliance certified-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType="zigbee")
echo $result
check_response "$result" "\"value\": false"

test_divider

echo "Get Revoked Model with VID: ${vid} PID: ${pid} SV: ${sv} before compliance record was created"
result=$(dcld query compliance revoked-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType="zigbee")
echo "$result"
check_response "$result" "\"value\": false"

test_divider

echo "Certify Model with VID: $vid PID: $pid  SV: ${sv} with zigbee certification"
certification_date="2020-01-01T00:00:00Z"
zigbee_certification_type="zigbee"
matter_certification_type="matter"
result=$(echo "test1234" | dcld tx compliance certify-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$zigbee_certification_type" --certificationDate="$certification_date" --from $zb_account --yes)
echo "$result"
check_response "$result" "\"success\": true"

echo "Certify Model with VID: $vid PID: $pid  SV: ${sv} with matter certification"
echo "dcld tx compliance certify-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$matter_certification_type" --certificationDate="$certification_date" --from $zb_account --yes"
result=$(echo "test1234" | dcld tx compliance certify-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$matter_certification_type" --certificationDate="$certification_date" --from $zb_account --yes)
echo "$result"
check_response "$result" "\"success\": true"

test_divider

echo "ReCertify Model with VID: $vid PID: $pid  SV: ${sv} "
certification_date="2020-01-01T00:00:00Z"
zigbee_certification_type="zigbee"
result=$(echo "test1234" | dcld tx compliance certify-model --vid=$vid --pid=$pid  --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$zigbee_certification_type" --certificationDate="$certification_date" --from $second_zb_account --yes)
check_response "$result" "\"success\": false"
echo "$result"

echo "Get Certified Model with VID: ${vid} PID: ${pid} SV: ${sv} for matter certification"
result=$(dcld query compliance certified-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$matter_certification_type)
check_response "$result" "\"value\": true"
echo "$result"

echo "Get Certified Model with VID: ${vid} PID: ${pid} SV: ${sv} for zigbee certification"
result=$(dcld query compliance certified-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$zigbee_certification_type)
check_response "$result" "\"value\": true"
echo "$result"

echo "Get Compliance Info for Model with VID: ${vid} PID: ${pid} SV: ${sv} for $zigbee_certification_type"
result=$(dcld query compliance compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$zigbee_certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"softwareVersionCertificationStatus\": 2"
check_response "$result" "\"date\": \"$certification_date\""
check_response "$result" "\"certification_type\": \"$zigbee_certification_type\""
echo "$result"

echo "Get Compliance Info for Model with VID: ${vid} PID: ${pid} SV: ${sv} for $matter_certification_type"
result=$(dcld query compliance compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$matter_certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"softwareVersionCertificationStatus\": 2"
check_response "$result" "\"date\": \"$certification_date\""
check_response "$result" "\"certification_type\": \"$matter_certification_type\""
echo "$result"

echo "Get All Certified Models"
result=$(dcld query compliance all-certified-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"certification_type\": \"$zigbee_certification_type\""
check_response "$result" "\"certification_type\": \"$matter_certification_type\""
echo "$result"

echo "Get All Compliance Info Recordss"
result=$(dcld query compliance all-compliance-info-records)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"date\": \"$certification_date\""
echo "$result"

echo "Revoke Certification for Model with VID: $vid PID: $pid SV: ${sv} "
revocation_date="2020-02-02T02:20:20Z"
revocation_reason="some reason"
result=$(echo "test1234" | dcld tx compliance revoke-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType="$zigbee_certification_type" --revocationDate="$revocation_date" --reason "$revocation_reason" --from $zb_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get Compliance Info for Model with VID: ${vid} PID: ${pid} SV: ${sv} "
result=$(dcld query compliance compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$zigbee_certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"softwareVersionCertificationStatus\": 3"
check_response "$result" "\"date\": \"$revocation_date\""
check_response "$result" "\"reason\": \"$revocation_reason\""
check_response "$result" "\"certification_type\": \"$zigbee_certification_type\""
check_response "$result" "\"history\""
echo "$result"

echo "Get Revoked Model with VID: ${vid} PID: ${pid} SV: ${sv} "
result=$(dcld query compliance revoked-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$zigbee_certification_type)
check_response "$result" "\"value\": true"
echo "$result"

echo "Get Certified Model with VID: ${vid} PID: ${pid} SV: ${sv} "
result=$(dcld query compliance certified-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$zigbee_certification_type)
check_response "$result" "\"value\": false"
echo "$result"

echo "Get All Revoked Models"
result=$(dcld query compliance all-revoked-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"certification_type\": \"$zigbee_certification_type\""
echo "$result"

echo "Again Certify Model with VID: $vid PID: $pid SV: ${sv}"
certification_date="2020-03-03T00:00:00Z"
result=$(echo "test1234" | dcld tx compliance certify-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$zigbee_certification_type" --certificationDate="$certification_date" --from $zb_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get Compliance Info for Model with VID: ${vid} PID: ${pid} SV: ${sv}"
result=$(dcld query compliance compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$zigbee_certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"softwareVersionCertificationStatus\": 2"
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
result=$(dcld query compliance all-compliance-info-records)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"softwareVersionCertificationStatus\": 2"
echo "$result"
