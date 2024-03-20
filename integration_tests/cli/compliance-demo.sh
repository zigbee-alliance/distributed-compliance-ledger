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
cd_certificate_id="123"
cd_version_number=1
schema_version_0=0
schema_version_2=2
echo "Certify unknown Model with VID: $vid PID: $pid  SV: ${sv} with zigbee certification"
result=$(echo "$passphrase" | dcld tx compliance certify-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$zigbee_certification_type" --certificationDate="$certification_date" --cdCertificateId="$cd_certificate_id" --from $zb_account --yes)
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
result=$(echo "$passphrase" | dcld tx compliance certify-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$zigbee_certification_type" --certificationDate="$certification_date" --cdCertificateId="$cd_certificate_id" --from $zb_account --yes)
echo "$result"
check_response "$result" "\"code\": 517"
check_response "$result" "No model version"

test_divider

echo "Add Model Version with VID: $vid PID: $pid SV: $sv SoftwareVersionString:$svs"
result=$(echo '$passphrase' | dcld tx model add-model-version --cdVersionNumber=$cd_version_number --maxApplicableSoftwareVersion=10 --minApplicableSoftwareVersion=1 --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --from=$vendor_account --yes)
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

echo "Get Device Software Compliance with CDCertificateID: ${cd_certificate_id}"
result=$(dcld query compliance device-software-compliance --cdCertificateId="$cd_certificate_id")
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

echo "Get All Device Software Compliance empty"
result=$(dcld query compliance all-device-software-compliance)
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
result=$(echo "$passphrase" | dcld tx compliance certify-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$invalid_svs --certificationType="$zigbee_certification_type" --certificationDate="$certification_date" --cdCertificateId="$cd_certificate_id" --from $zb_account --yes)
check_response "$result" "\"code\": 306"
# check_response "$result" "failed to execute message; message index: 0: Model with vid=$vid, pid=$pid, softwareVersion=$svs present on the ledger does not have matching softwareVersionString=$invalid_svs: model version does not match"

test_divider

echo "Certify Model with VID: $vid PID: $pid  SV: ${sv} with zigbee certification and invalid CDVersionNumber"
result=$(echo "$passphrase" | dcld tx compliance certify-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$zigbee_certification_type" --certificationDate="$certification_date" --cdCertificateId="$cd_certificate_id" --from $zb_account --yes)
echo "$result"
check_response "$result" "\"code\": 306"
check_response "$result" "ledger does not have matching CDVersionNumber=0: model version does not match"

test_divider

echo "Certify Model with VID: $vid PID: $pid  SV: ${sv} with zigbee certification"
result=$(echo "$passphrase" | dcld tx compliance certify-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$zigbee_certification_type" --certificationDate="$certification_date" --cdCertificateId="$cd_certificate_id" --cdVersionNumber=$cd_version_number  --schemaVersion=$schema_version_2 --from $zb_account --yes)
echo "$result"
check_response "$result" "\"code\": 0"

test_divider

echo "Certify Model with VID: $vid PID: $pid  SV: ${sv} with matter certification"
echo "dcld tx compliance certify-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$matter_certification_type" --certificationDate="$certification_date" --from $zb_account --yes"
result=$(echo "$passphrase" | dcld tx compliance certify-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$matter_certification_type" --certificationDate="$certification_date" --cdCertificateId="$cd_certificate_id" --cdVersionNumber=$cd_version_number --from $zb_account --yes)
echo "$result"
check_response "$result" "\"code\": 0"

test_divider

echo "ReCertify Model with VID: $vid PID: $pid  SV: ${sv} by different account"
zigbee_certification_type="zigbee"
result=$(echo "$passphrase" | dcld tx compliance certify-model --vid=$vid --pid=$pid  --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$zigbee_certification_type" --certificationDate="$certification_date" --cdCertificateId="$cd_certificate_id" --cdVersionNumber=$cd_version_number --from $second_zb_account --yes)
check_response "$result" "\"code\": 303"
check_response "$result" "already certified on the ledger"
echo "$result"

test_divider

echo "ReCertify Model with VID: $vid PID: $pid  SV: ${sv} by same account"
zigbee_certification_type="zigbee"
result=$(echo "$passphrase" | dcld tx compliance certify-model --vid=$vid --pid=$pid  --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$zigbee_certification_type" --certificationDate="$certification_date" --cdCertificateId="$cd_certificate_id" --cdVersionNumber=$cd_version_number --from $zb_account --yes)
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
check_response "$result" "\"softwareVersionCertificationStatus\": 2"
check_response "$result" "\"cDCertificateId\": \"$cd_certificate_id\""
check_response "$result" "\"date\": \"$certification_date\""
check_response "$result" "\"certificationType\": \"$zigbee_certification_type\""
check_response "$result" "\"schemaVersion\": $schema_version_2"
echo "$result"

test_divider

echo "Get Compliance Info for Model with VID: ${vid} PID: ${pid} SV: ${sv} for $matter_certification_type"
result=$(dcld query compliance compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$matter_certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"softwareVersionCertificationStatus\": 2"
check_response "$result" "\"cDCertificateId\": \"$cd_certificate_id\""
check_response "$result" "\"date\": \"$certification_date\""
check_response "$result" "\"certificationType\": \"$matter_certification_type\""
check_response "$result" "\"schemaVersion\": $schema_version_0"
echo "$result"

test_divider

echo "Get Device Software Compliance for Model with CDCertificateID: ${cd_certificate_id}"
result=$(dcld query compliance device-software-compliance --cdCertificateId="$cd_certificate_id")
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"softwareVersionCertificationStatus\": 2"
check_response "$result" "\"cDCertificateId\": \"$cd_certificate_id\""
check_response "$result" "\"date\": \"$certification_date\""
check_response "$result" "\"certificationType\": \"$zigbee_certification_type\""
check_response "$result" "\"certificationType\": \"$matter_certification_type\""
echo "$result"

test_divider

echo "Get All Certified Models"
result=$(dcld query compliance all-certified-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"certificationType\": \"$zigbee_certification_type\""
check_response "$result" "\"certificationType\": \"$matter_certification_type\""
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

echo "Get All Compliance Info Records"
result=$(dcld query compliance all-compliance-info)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"certificationType\": \"$zigbee_certification_type\""
check_response "$result" "\"certificationType\": \"$matter_certification_type\""
check_response "$result" "\"date\": \"$certification_date\""
echo "$result"

test_divider

echo "Get All Device Software Compliance"
result=$(dcld query compliance all-device-software-compliance)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"certificationType\": \"$zigbee_certification_type\""
check_response "$result" "\"certificationType\": \"$matter_certification_type\""
check_response "$result" "\"date\": \"$certification_date\""
echo "$result"

test_divider

revocation_reason="some reason"

echo "Revoke Certification for Model with VID: $vid PID: $pid SV: ${sv} from the past"
revocation_date_past="2020-01-01T00:00:00Z"
result=$(echo "$passphrase" | dcld tx compliance revoke-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$zigbee_certification_type" --revocationDate="$revocation_date_past" --reason "$revocation_reason" --cdVersionNumber=$cd_version_number --from $zb_account --yes)
check_response "$result" "\"code\": 302"
check_response "$result" "must be after"
echo "$result"

test_divider


echo "Revoke Certification for Model with VID: $vid PID: $pid SV: ${sv} "
revocation_date="2020-02-02T02:20:20Z"
result=$(echo "$passphrase" | dcld tx compliance revoke-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$zigbee_certification_type" --revocationDate="$revocation_date" --reason "$revocation_reason" --cdVersionNumber=$cd_version_number --from $zb_account --yes)
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "Get Compliance Info for Model with VID: ${vid} PID: ${pid} SV: ${sv} "
result=$(dcld query compliance compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$zigbee_certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"softwareVersionCertificationStatus\": 3"
check_response "$result" "\"cDCertificateId\": \"$cd_certificate_id\""
check_response "$result" "\"date\": \"$revocation_date\""
check_response "$result" "\"reason\": \"$revocation_reason\""
check_response "$result" "\"certificationType\": \"$zigbee_certification_type\""
check_response "$result" "\"history\""
echo "$result"

test_divider

echo "Get Device Software Compliance for Model with CDCertificateID: ${cd_certificate_id}"
result=$(dcld query compliance device-software-compliance --cdCertificateId="$cd_certificate_id")
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"softwareVersionCertificationStatus\": 2"
check_response "$result" "\"cDCertificateId\": \"$cd_certificate_id\""
check_response "$result" "\"date\": \"$certification_date\""
check_response "$result" "\"certificationType\": \"$matter_certification_type\""
response_does_not_contain "$result" "\"certificationType\": \"$zigbee_certification_type\""
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
check_response "$result" "\"certificationType\": \"$zigbee_certification_type\""
echo "$result"

test_divider

echo "Get All Certified Models"
result=$(dcld query compliance all-certified-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"certificationType\": \"$matter_certification_type\""
response_does_not_contain "$result" "\"certificationType\": \"$zigbee_certification_type\""
echo "$result"

test_divider


echo "Again Certify Model with VID: $vid PID: $pid SV: ${sv}"
certification_date="2020-03-03T00:00:00Z"
result=$(echo "$passphrase" | dcld tx compliance certify-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$zigbee_certification_type" --certificationDate="$certification_date" --cdCertificateId="$cd_certificate_id" --cdVersionNumber=$cd_version_number --from $zb_account --yes)
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "Get Compliance Info for Model with VID: ${vid} PID: ${pid} SV: ${sv}"
result=$(dcld query compliance compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$zigbee_certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"softwareVersionCertificationStatus\": 2"
check_response "$result" "\"cDCertificateId\": \"$cd_certificate_id\""
check_response "$result" "\"date\": \"$certification_date\""
check_response "$result" "\"certificationType\": \"zigbee\""
echo "$result"

test_divider

echo "Get Device Software Compliance for Model with ${cd_certificate_id}"
result=$(dcld query compliance device-software-compliance --cdCertificateId="$cd_certificate_id")
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"softwareVersionCertificationStatus\": 2"
check_response "$result" "\"cDCertificateId\": \"$cd_certificate_id\""
check_response "$result" "\"date\": \"$certification_date\""
check_response "$result" "\"certificationType\": \"zigbee\""
check_response "$result" "\"certificationType\": \"matter\""
echo "$result"

test_divider

echo "Get Certified Model with VID: ${vid} PID: ${pid} SV: ${sv}"
result=$(dcld query compliance certified-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$zigbee_certification_type)
check_response "$result" "\"value\": true"
echo "$result"

test_divider

echo "Get Revoked Model with VID: ${vid} PID: ${pid} SV: ${sv}"
result=$(dcld query compliance revoked-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$zigbee_certification_type)
check_response "$result" "\"value\": false"
echo "$result"

test_divider

echo "Get All Compliance Infos"
result=$(dcld query compliance all-compliance-info)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"softwareVersionCertificationStatus\": 2"
echo "$result"

test_divider

echo "Get All Device Software Compliance"
result=$(dcld query compliance all-device-software-compliance)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"softwareVersionCertificationStatus\": 2"
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
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"certificationType\": \"$matter_certification_type\""
check_response "$result" "\"certificationType\": \"$zigbee_certification_type\""
echo "$result"

test_divider

###########################################################################################################################################
# PREPERATION
pid=$RANDOM
sv=$RANDOM
svs=$RANDOM

# ADD MODEL
echo "Add Model with VID: $vid PID: $pid"
result=$(echo "$passphrase" | dcld tx model add-model --vid=$vid --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0  --from $vendor_account --yes)
echo $result
check_response "$result" "\"code\": 0"

test_divider

# ADD MODEL VERSION
echo "Add Model Version with VID: $vid PID: $pid SV: $sv SoftwareVersionString:$svs"
result=$(echo '$passphrase' | dcld tx model add-model-version --cdVersionNumber=$cd_version_number --maxApplicableSoftwareVersion=10 --minApplicableSoftwareVersion=1 --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --from=$vendor_account --yes)
echo $result
check_response "$result" "\"code\": 0"

test_divider

# ADD CERTIFY MODEL WITH ALL OPTIONAL FIELDS
echo "Certify Model with VID: $vid PID: $pid SV: ${sv} with zigbee certification"
result=$(echo "$passphrase" | dcld tx compliance certify-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$zigbee_certification_type" --certificationDate="$certification_date" --cdCertificateId="$cd_certificate_id" --cdVersionNumber=$cd_version_number --programTypeVersion="1.0" --familyId="someFID" --supportedClusters="someClusters" --compliantPlatformUsed="WIFI" --compliantPlatformVersion="V1" --OSVersion="someV" --certificationRoute="Full" --programType="pType" --transport="someTransport" --parentChild="parent" --certificationIDOfSoftwareComponent="someIDOfSoftwareComponent" --from $zb_account --yes)
echo "$result"
check_response "$result" "\"code\": 0"

test_divider

# GET CERTIFIED MODEL
echo "Get Certified Model with VID: ${vid} PID: ${pid} SV: ${sv} for zigbee certification"
result=$(dcld query compliance certified-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$zigbee_certification_type)
check_response "$result" "\"value\": true"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"vid\": $vid"
echo "$result"

test_divider

# GET COMPLIANCE INFO
echo "Get Compliance Info for Model with VID: ${vid} PID: ${pid} SV: ${sv} for $zigbee_certification_type"
result=$(dcld query compliance compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$zigbee_certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"softwareVersionCertificationStatus\": 2"
check_response "$result" "\"date\": \"$certification_date\""
check_response "$result" "\"certificationType\": \"$zigbee_certification_type\""
check_response "$result" "\"programTypeVersion\": \"1.0\""
check_response "$result" "\"cDCertificateId\": \"$cd_certificate_id\""
check_response "$result" "\"familyId\": \"someFID\""
check_response "$result" "\"supportedClusters\": \"someClusters\""
check_response "$result" "\"compliantPlatformUsed\": \"WIFI\""
check_response "$result" "\"compliantPlatformVersion\": \"V1\""
check_response "$result" "\"OSVersion\": \"someV\""
check_response "$result" "\"certificationRoute\": \"Full\""
check_response "$result" "\"programType\": \"pType\""
check_response "$result" "\"transport\": \"someTransport\""
check_response "$result" "\"parentChild\": \"parent\""
check_response "$result" "\"certificationIdOfSoftwareComponent\": \"someIDOfSoftwareComponent\""
echo "$result"

test_divider

# GET DEVICE SOFTWARE COMPLIANCE
echo "Get Device Software Compliance for Model with CDCertificateID: ${cd_certificate_id}"
result=$(dcld query compliance device-software-compliance --cdCertificateId="$cd_certificate_id")
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"softwareVersionCertificationStatus\": 2"
check_response "$result" "\"date\": \"$certification_date\""
check_response "$result" "\"certificationType\": \"$zigbee_certification_type\""
check_response "$result" "\"programTypeVersion\": \"1.0\""
check_response "$result" "\"cDCertificateId\": \"$cd_certificate_id\""
check_response "$result" "\"familyId\": \"someFID\""
check_response "$result" "\"supportedClusters\": \"someClusters\""
check_response "$result" "\"compliantPlatformUsed\": \"WIFI\""
check_response "$result" "\"compliantPlatformVersion\": \"V1\""
check_response "$result" "\"OSVersion\": \"someV\""
check_response "$result" "\"certificationRoute\": \"Full\""
check_response "$result" "\"programType\": \"pType\""
check_response "$result" "\"transport\": \"someTransport\""
check_response "$result" "\"parentChild\": \"parent\""
check_response "$result" "\"certificationIdOfSoftwareComponent\": \"someIDOfSoftwareComponent\""
echo "$result"
###########################################################################################################################################

test_divider

new_reason="new_reason"
new_program_type="new_program_type"
new_parent_child="child"
new_transport="new_transport"

# UPDATE COMPLIANCE INFO BY CERTIFICATION CENTER ACCOUNT
echo "Update Compliance Info for Model with VID: ${vid} PID: ${pid} SV: ${sv} for $zigbee_certification_type with some optional fields set"
result=$(echo "$passphrase" | dcld tx compliance update-compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$zigbee_certification_type --reason=$new_reason --programType=$new_program_type --parentChild=$new_parent_child --transport=$new_transport --from=$zb_account --yes)
echo "$result"

result=$(dcld query compliance compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$zigbee_certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"softwareVersionCertificationStatus\": 2"
check_response "$result" "\"date\": \"$certification_date\""
check_response "$result" "\"certificationType\": \"$zigbee_certification_type\""
check_response "$result" "\"programTypeVersion\": \"1.0\""
check_response "$result" "\"cDCertificateId\": \"$cd_certificate_id\""
check_response "$result" "\"familyId\": \"someFID\""
check_response "$result" "\"supportedClusters\": \"someClusters\""
check_response "$result" "\"compliantPlatformUsed\": \"WIFI\""
check_response "$result" "\"compliantPlatformVersion\": \"V1\""
check_response "$result" "\"OSVersion\": \"someV\""
check_response "$result" "\"certificationRoute\": \"Full\""
check_response "$result" "\"programType\": \"$new_program_type\""
check_response "$result" "\"transport\": \"$new_transport\""
check_response "$result" "\"parentChild\": \"$new_parent_child\""
check_response "$result" "\"reason\": \"$new_reason\""
check_response "$result" "\"certificationIdOfSoftwareComponent\": \"someIDOfSoftwareComponent\""
echo "$result"

by_vendor_reason="by_vendor_reason"
by_vendor_program_type="by_vendor_program_type"
by_vendor_parent_child="parent"
by_vendor_transport="by_vendor_transport"

test_divider

# UPDATE COMPLIANCE INFO BY *NON CERTIFICATION CENTER ACCOUNT
echo "Update Compliance Info for Model with VID: ${vid} PID: ${pid} SV: ${sv} for $zigbee_certification_type by non Certification Center account"
result=$(echo "$passphrase" | dcld tx compliance update-compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$zigbee_certification_type --reason=$by_vendor_reason --programType=$by_vendor_program_type --parentChild=$by_vendor_parent_child --transport=$by_vendor_transport --from=$vendor_account --yes)
check_response "$result" "unauthorized"

echo "Compliance Info for Model with VID: ${vid} PID: ${pid} SV: ${sv} for $zigbee_certification_type is not updated by non CertificationCenter account"
result=$(dcld query compliance compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$zigbee_certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"softwareVersionCertificationStatus\": 2"
check_response "$result" "\"date\": \"$certification_date\""
check_response "$result" "\"certificationType\": \"$zigbee_certification_type\""
check_response "$result" "\"programTypeVersion\": \"1.0\""
check_response "$result" "\"cDCertificateId\": \"$cd_certificate_id\""
check_response "$result" "\"familyId\": \"someFID\""
check_response "$result" "\"supportedClusters\": \"someClusters\""
check_response "$result" "\"compliantPlatformUsed\": \"WIFI\""
check_response "$result" "\"compliantPlatformVersion\": \"V1\""
check_response "$result" "\"OSVersion\": \"someV\""
check_response "$result" "\"certificationRoute\": \"Full\""
check_response "$result" "\"programType\": \"$new_program_type\""
check_response "$result" "\"transport\": \"$new_transport\""
check_response "$result" "\"parentChild\": \"$new_parent_child\""
check_response "$result" "\"reason\": \"$new_reason\""
check_response "$result" "\"certificationIdOfSoftwareComponent\": \"someIDOfSoftwareComponent\""
echo "$result"

test_divider

# UPDATE COMPLIANCE INFO WITH NO OPTIONAL FIELDS SET
echo "Update Compliance Info for Model with VID: ${vid} PID: ${pid} SV: ${sv} for $zigbee_certification_type with no optional fields set"
result=$(echo "$passphrase" | dcld tx compliance update-compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$zigbee_certification_type  --from=$zb_account --yes)
echo "$result"

echo "Compliance Info for Model with VID: ${vid} PID: ${pid} SV: ${sv} for $zigbee_certification_type is not updated"
result=$(dcld query compliance compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$zigbee_certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"softwareVersionCertificationStatus\": 2"
check_response "$result" "\"date\": \"$certification_date\""
check_response "$result" "\"certificationType\": \"$zigbee_certification_type\""
check_response "$result" "\"programTypeVersion\": \"1.0\""
check_response "$result" "\"cDCertificateId\": \"$cd_certificate_id\""
check_response "$result" "\"familyId\": \"someFID\""
check_response "$result" "\"supportedClusters\": \"someClusters\""
check_response "$result" "\"compliantPlatformUsed\": \"WIFI\""
check_response "$result" "\"compliantPlatformVersion\": \"V1\""
check_response "$result" "\"OSVersion\": \"someV\""
check_response "$result" "\"certificationRoute\": \"Full\""
check_response "$result" "\"programType\": \"$new_program_type\""
check_response "$result" "\"transport\": \"$new_transport\""
check_response "$result" "\"parentChild\": \"$new_parent_child\""
check_response "$result" "\"reason\": \"$new_reason\""
check_response "$result" "\"certificationIdOfSoftwareComponent\": \"someIDOfSoftwareComponent\""
echo "$result"

test_divider

# UPDATE COMPLIANCE INFO WITH ALL OPTIONAL FIELDS SET
upd_cd_version_number="1"
upd_certification_date="2022-01-01T00:00:01Z"
upd_reason="brand_new_reason"
upd_cd_certificate_id="brand_new_ID"
upd_certification_route="brand_new_route"
upd_program_type="brand_new_program_type"
upd_program_type_version="brand_new_program_type_version"
upd_compliant_platform_used="brand_new_compliant_platform_used"
upd_compliant_platform_version="brand_new_compliant_platform_version"
upd_transport="brand_new_transport"
upd_familyID="brand_new_family_ID"
upd_supported_clusters="brand_new_clusters"
upd_os_version="brand_new_os_version"
upd_parent_child="parent"
upd_certification_id_of_software_component="brand_new_component"
schema_version_3=3

echo "Update Compliance Info for Model with VID: ${vid} PID: ${pid} SV: ${sv} for $zigbee_certification_type with all optional fields set"
result=$(echo "$passphrase" | dcld tx compliance update-compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$zigbee_certification_type --cdVersionNumber=$upd_cd_version_number --certificationDate=$upd_certification_date --reason=$upd_reason --cdCertificateId=$upd_cd_certificate_id --certificationRoute=$upd_certification_route --programType=$upd_program_type --programTypeVersion=$upd_program_type_version --compliantPlatformUsed=$upd_compliant_platform_used --compliantPlatformVersion=$upd_compliant_platform_version --transport=$upd_transport --familyId=$upd_familyID --supportedClusters=$upd_supported_clusters --OSVersion=$upd_os_version --parentChild=$upd_parent_child --certificationIDOfSoftwareComponent=$upd_certification_id_of_software_component --schemaVersion=$schema_version_3 --from=$zb_account --yes)
echo "$result"

echo "Compliance Info for Model with VID: ${vid} PID: ${pid} SV: ${sv} for $zigbee_certification_type all fields updated"
result=$(dcld query compliance compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$zigbee_certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"softwareVersion\": $sv"
check_response "$result" "\"certificationType\": \"$zigbee_certification_type\""
check_response "$result" "\"cDVersionNumber\": $(($upd_cd_version_number + 0))"
check_response "$result" "\"date\": \"$upd_certification_date\""
check_response "$result" "\"reason\": \"$upd_reason\""
check_response "$result" "\"cDCertificateId\": \"$upd_cd_certificate_id\""
check_response "$result" "\"certificationRoute\": \"$upd_certification_route\""
check_response "$result" "\"programType\": \"$upd_program_type\""
check_response "$result" "\"programTypeVersion\": \"$upd_program_type_version\""
check_response "$result" "\"compliantPlatformUsed\": \"$upd_compliant_platform_used\""
check_response "$result" "\"compliantPlatformVersion\": \"$upd_compliant_platform_version\""
check_response "$result" "\"transport\": \"$upd_transport\""
check_response "$result" "\"familyId\": \"$upd_familyID\""
check_response "$result" "\"supportedClusters\": \"$upd_supported_clusters\""
check_response "$result" "\"OSVersion\": \"$upd_os_version\""
check_response "$result" "\"parentChild\": \"$upd_parent_child\""
check_response "$result" "\"certificationIdOfSoftwareComponent\": \"$upd_certification_id_of_software_component\""
check_response "$result" "\"schemaVersion\": $schema_version_3"


test_divider

echo "Get Device Software Compliance for Model with the updated CDCertificateID: ${upd_cd_certificate_id}"
result=$(dcld query compliance device-software-compliance --cdCertificateId="$upd_cd_certificate_id")
echo $result
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"

# delete compliance info
echo "delete compliance info vid=$vid pid=$pid softwareVerion=$sv certificationType=$zigbee_certification_type"
result=$(echo "$passphrase" | dcld tx compliance delete-compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$zigbee_certification_type --from=$zb_account --yes)

test_divider

echo "Get Compliance Info with VID: ${vid} PID: ${pid} SV: ${sv} after deletion"
result=$(dcld query compliance compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$zigbee_certification_type)
check_response "$result" "Not Found"

test_divider

echo "Get Device Software Compliance for cdCertificateId: $cd_certificate_id after deletion"
result=$(dcld query compliance device-software-compliance --cdCertificateId="$cd_certificate_id")
response_does_not_contain "$result" "\"vid\":$vid,\"pid\":$pid,\"softwareVersion\":$sv,\"certificationType\":\"$zigbee_certification_type\""

test_divider

echo "Get Compliance Info with VID: ${vid} PID: ${pid} SV: ${sv} after deletion"
result=$(dcld query compliance certified-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$zigbee_certification_type)
check_response "$result" "Not Found"

test_divider

echo "PASSED"
