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

# Upgrade constants

plan_name="v1.4.4"
upgrade_checksum="sha256:e4031c6a77aa8e58add391be671a334613271bcf6e7f11d23b04a0881ece6958"
binary_version_old="v1.4.3"
binary_version_new="v1.4.4"

wget -O dcld_old "https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/$binary_version_old/dcld"
chmod ugo+x dcld_old

wget -O dcld_new "https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/$binary_version_new/dcld"
chmod ugo+x dcld_new

DCLD_BIN_OLD="./dcld_old"
DCLD_BIN_NEW="./dcld_new"
$DCLD_BIN_NEW config broadcast-mode sync
########################################################################################

# Upgrade to version 1.4

get_height current_height
echo "Current height is $current_height"

plan_height=$(expr $current_height \+ 20)

test_divider

echo "Propose upgrade $plan_name at height $plan_height"
echo "https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/$binary_version_new/dcld?checksum=$upgrade_checksum"
result=$(echo $passphrase | $DCLD_BIN_OLD tx dclupgrade propose-upgrade --name=$plan_name --upgrade-height=$plan_height --upgrade-info="{\"binaries\":{\"linux/amd64\":\"https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/$binary_version_new/dcld?checksum=$upgrade_checksum\"}}" --from $trustee_account_1 --yes)
echo "$result"
check_response "$result" "\"code\": 0"

test_divider

echo "Approve upgrade $plan_name"
result=$(echo $passphrase | $DCLD_BIN_OLD tx dclupgrade approve-upgrade --name $plan_name --from $trustee_account_2 --yes)
echo "$result"
check_response "$result" "\"code\": 0"

echo "Approve upgrade $plan_name"
result=$(echo $passphrase | $DCLD_BIN_OLD tx dclupgrade approve-upgrade --name $plan_name --from $trustee_account_3 --yes)
echo "$result"
check_response "$result" "\"code\": 0"

echo "Approve upgrade $plan_name"
result=$(echo $passphrase | $DCLD_BIN_OLD tx dclupgrade approve-upgrade --name $plan_name --from $trustee_account_4 --yes)
echo "$result"
check_response "$result" "\"code\": 0"

test_divider

echo "Wait for block height to become greater than upgrade $plan_name plan height"
wait_for_height $(expr $plan_height + 1) 300 outage-safe

test_divider

echo "Verify that no upgrade has been scheduled anymore"
result=$($DCLD_BIN_NEW query upgrade plan 2>&1) || true
check_response_and_report "$result" "no upgrade scheduled" raw

test_divider

echo "Verify that upgrade is applied"
result=$($DCLD_BIN_NEW query upgrade applied $plan_name)
echo "$result"

test_divider

########################################################################################

echo "Verify that old data is not corrupted"

test_divider

# VENDORINFO

echo "Verify if VendorInfo Record for VID: $vid is present or not"
result=$($DCLD_BIN_NEW query vendorinfo vendor --vid=$vid)
check_response "$result" "\"vendorID\": $vid"
check_response "$result" "\"companyLegalName\": \"$company_legal_name\""
check_response "$result" "\"vendorName\": \"$vendor_name\""
check_response "$result" "\"companyPreferredName\": \"$company_preferred_name_for_1_2\""
check_response "$result" "\"vendorLandingPageURL\": \"$vendor_landing_page_url_for_1_2\""

echo "Verify if VendorInfo Record for VID: $vid_for_1_2 is present or not"
result=$($DCLD_BIN_NEW query vendorinfo vendor --vid=$vid_for_1_2)
check_response "$result" "\"vendorID\": $vid_for_1_2"
check_response "$result" "\"companyLegalName\": \"$company_legal_name_for_1_2\""

echo "Verify if VendorInfo Record for VID: $vid_for_1_4_3 is present or not"
result=$($DCLD_BIN_NEW query vendorinfo vendor --vid=$vid_for_1_4_3)
check_response "$result" "\"vendorID\": $vid_for_1_4_3"
check_response "$result" "\"companyLegalName\": \"$company_legal_name_for_1_4_3\""

echo "Request all vendor infos"
result=$($DCLD_BIN_NEW query vendorinfo all-vendors)
check_response "$result" "\"vendorID\": $vid"
check_response "$result" "\"vendorID\": $vid_for_1_2"
check_response "$result" "\"vendorID\": $vid_for_1_4_3"
check_response "$result" "\"companyLegalName\": \"$company_legal_name\""
check_response "$result" "\"companyLegalName\": \"$company_legal_name_for_1_2\""
check_response "$result" "\"companyLegalName\": \"$company_legal_name_for_1_4_3\""
check_response "$result" "\"vendorName\": \"$vendor_name\""
check_response "$result" "\"vendorName\": \"$vendor_name_for_1_2\""
check_response "$result" "\"vendorName\": \"$vendor_name_for_1_4_3\""

test_divider

# MODEL

echo "Get Model with VID: $vid PID: $pid_1"
result=$($DCLD_BIN_NEW query model get-model --vid=$vid --pid=$pid_1)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_1"
check_response "$result" "\"productLabel\": \"$product_label\""

echo "Get Model with VID: $vid PID: $pid_2"
result=$($DCLD_BIN_NEW query model get-model --vid=$vid --pid=$pid_2)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_2"
check_response "$result" "\"productLabel\": \"$product_label_for_1_4_3\""
check_response "$result" "\"partNumber\": \"$part_number_for_1_4_3\""

echo "Get Model with VID: $vid_for_1_2 PID: $pid_1_for_1_2"
result=$($DCLD_BIN_NEW query model get-model --vid=$vid_for_1_2 --pid=$pid_1_for_1_2)
check_response "$result" "\"vid\": $vid_for_1_2"
check_response "$result" "\"pid\": $pid_1_for_1_2"
check_response "$result" "\"productLabel\": \"$product_label_for_1_2\""

echo "Get Model with VID: $vid_for_1_2 PID: $pid_2_for_1_2"
result=$($DCLD_BIN_NEW query model get-model --vid=$vid_for_1_2 --pid=$pid_2_for_1_2)
check_response "$result" "\"vid\": $vid_for_1_2"
check_response "$result" "\"pid\": $pid_2_for_1_2"
check_response "$result" "\"productLabel\": \"$product_label_for_1_2\""

echo "Get Model with VID: $vid_for_1_4_3 PID: $pid_2_for_1_4_3"
result=$($DCLD_BIN_NEW query model get-model --vid=$vid_for_1_4_3 --pid=$pid_2_for_1_4_3)
check_response "$result" "\"vid\": $vid_for_1_4_3"
check_response "$result" "\"pid\": $pid_2_for_1_4_3"
check_response "$result" "\"productLabel\": \"$product_label_for_1_4_3\""

echo "Get all models"
result=$($DCLD_BIN_NEW query model all-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_1"
check_response "$result" "\"pid\": $pid_2"
check_response "$result" "\"vid\": $vid_for_1_2"
check_response "$result" "\"pid\": $pid_1_for_1_2"
check_response "$result" "\"pid\": $pid_2_for_1_2"
check_response "$result" "\"vid\": $vid_for_1_4_3"
check_response "$result" "\"pid\": $pid_1_for_1_4_3"
check_response "$result" "\"pid\": $pid_2_for_1_4_3"

echo "Get Vendor Models with VID: ${vid}"
result=$($DCLD_BIN_NEW query model vendor-models --vid=$vid)
check_response "$result" "\"pid\": $pid_1"
check_response "$result" "\"pid\": $pid_2"

echo "Get Vendor Models with VID: ${vid_for_1_2}"
result=$($DCLD_BIN_NEW query model vendor-models --vid=$vid_for_1_2)
check_response "$result" "\"pid\": $pid_1_for_1_2"
check_response "$result" "\"pid\": $pid_2_for_1_2"

echo "Get Vendor Models with VID: ${vid_for_1_2}"
result=$($DCLD_BIN_NEW query model vendor-models --vid=$vid_for_1_4_3)
check_response "$result" "\"pid\": $pid_1_for_1_4_3"
check_response "$result" "\"pid\": $pid_2_for_1_4_3"

echo "Get model version VID: $vid PID: $pid_1"
result=$($DCLD_BIN_NEW query model model-version --vid=$vid --pid=$pid_1 --softwareVersion=$software_version)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_1"
check_response "$result" "\"softwareVersion\": $software_version"

echo "Get model version VID: $vid PID: $pid_2 updated"
result=$($DCLD_BIN_NEW query model model-version --vid=$vid --pid=$pid_2  --softwareVersion=$software_version)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_2"
check_response "$result" "\"softwareVersion\": $software_version"
check_response "$result" "\"minApplicableSoftwareVersion\": $min_applicable_software_version_for_1_4_3"
check_response "$result" "\"maxApplicableSoftwareVersion\": $max_applicable_software_version_for_1_4_3"

echo "Get model version VID: $vid_for_1_2 PID: $pid_1_for_1_2"
result=$($DCLD_BIN_NEW query model model-version --vid=$vid_for_1_2 --pid=$pid_1_for_1_2 --softwareVersion=$software_version_for_1_2)
check_response "$result" "\"vid\": $vid_for_1_2"
check_response "$result" "\"pid\": $pid_1_for_1_2"
check_response "$result" "\"softwareVersion\": $software_version_for_1_2"

echo "Get model version VID: $vid_for_1_2 PID: $pid_2_for_1_2"
result=$($DCLD_BIN_NEW query model model-version --vid=$vid_for_1_2 --pid=$pid_2_for_1_2 --softwareVersion=$software_version_for_1_2)
check_response "$result" "\"vid\": $vid_for_1_2"
check_response "$result" "\"pid\": $pid_2_for_1_2"
check_response "$result" "\"softwareVersion\": $software_version_for_1_2"

echo "Get model version VID: $vid_for_1_4_3 PID: $pid_2_for_1_4_3"
result=$($DCLD_BIN_NEW query model model-version --vid=$vid_for_1_4_3 --pid=$pid_2_for_1_4_3 --softwareVersion=$software_version_for_1_4_3)
check_response "$result" "\"vid\": $vid_for_1_4_3"
check_response "$result" "\"pid\": $pid_2_for_1_4_3"
check_response "$result" "\"softwareVersion\": $software_version_for_1_4_3"

test_divider

# COMPLIANCE

echo "Get certified model vid=$vid pid=$pid_1"
result=$($DCLD_BIN_NEW query compliance certified-model --vid=$vid --pid=$pid_1 --softwareVersion=$software_version --certificationType=$certification_type)
check_response "$result" "\"value\": true"
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_1"
check_response "$result" "\"softwareVersion\": $software_version"
check_response "$result" "\"certificationType\": \"$certification_type\""

echo "Get certified model vid=$vid_for_1_2 pid=$pid_1_for_1_2"
result=$($DCLD_BIN_NEW query compliance certified-model --vid=$vid_for_1_2 --pid=$pid_1_for_1_2 --softwareVersion=$software_version_for_1_2 --certificationType=$certification_type_for_1_2)
check_response "$result" "\"value\": true"
check_response "$result" "\"vid\": $vid_for_1_2"
check_response "$result" "\"pid\": $pid_1_for_1_2"
check_response "$result" "\"softwareVersion\": $software_version_for_1_2"
check_response "$result" "\"certificationType\": \"$certification_type_for_1_2\""

echo "Get certified model vid=$vid_for_1_4_3 pid=$pid_1_for_1_4_3"
result=$($DCLD_BIN_NEW query compliance certified-model --vid=$vid_for_1_4_3 --pid=$pid_1_for_1_4_3 --softwareVersion=$software_version_for_1_4_3 --certificationType=$certification_type_for_1_4_3)
check_response "$result" "\"value\": true"
check_response "$result" "\"vid\": $vid_for_1_4_3"
check_response "$result" "\"pid\": $pid_1_for_1_4_3"
check_response "$result" "\"softwareVersion\": $software_version_for_1_4_3"
check_response "$result" "\"certificationType\": \"$certification_type_for_1_4_3\""

echo "Get revoked Model with VID: $vid PID: $pid_2"
result=$($DCLD_BIN_NEW query compliance revoked-model --vid=$vid --pid=$pid_2 --softwareVersion=$software_version --certificationType=$certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_2"

echo "Get revoked Model with VID: $vid_for_1_2 PID: $pid_2_for_1_2"
result=$($DCLD_BIN_NEW query compliance revoked-model --vid=$vid_for_1_2 --pid=$pid_2_for_1_2 --softwareVersion=$software_version_for_1_2 --certificationType=$certification_type_for_1_2)
check_response "$result" "\"vid\": $vid_for_1_2"
check_response "$result" "\"pid\": $pid_2_for_1_2"

echo "Get revoked Model with VID: $vid_for_1_4_3 PID: $pid_2_for_1_4_3"
result=$($DCLD_BIN_NEW query compliance revoked-model --vid=$vid_for_1_4_3 --pid=$pid_2_for_1_4_3 --softwareVersion=$software_version_for_1_4_3 --certificationType=$certification_type_for_1_4_3)
check_response "$result" "\"vid\": $vid_for_1_4_3"
check_response "$result" "\"pid\": $pid_2_for_1_4_3"

echo "Get provisional model with VID: $vid PID: $pid_3"
result=$($DCLD_BIN_NEW query compliance provisional-model --vid=$vid --pid=$pid_3 --softwareVersion=$software_version --certificationType=$certification_type)
check_response "$result" "\"value\": true"
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_3"

echo "Get provisional model with VID: $vid_for_1_2 PID: $pid_2_for_1_2"
result=$($DCLD_BIN_NEW query compliance provisional-model --vid=$vid_for_1_2 --pid=$pid_2_for_1_2 --softwareVersion=$software_version_for_1_2 --certificationType=$certification_type_for_1_2)
check_response "$result" "\"value\": false"
check_response "$result" "\"vid\": $vid_for_1_2"
check_response "$result" "\"pid\": $pid_2_for_1_2"

echo "Get compliance-info model with VID: $vid_for_1_2 PID: $pid_1_for_1_2"
result=$($DCLD_BIN_NEW query compliance compliance-info --vid=$vid_for_1_2 --pid=$pid_1_for_1_2 --softwareVersion=$software_version_for_1_2 --certificationType=$certification_type_for_1_2)
check_response "$result" "\"vid\": $vid_for_1_2"
check_response "$result" "\"pid\": $pid_1_for_1_2"
check_response "$result" "\"softwareVersion\": $software_version_for_1_2"
check_response "$result" "\"certificationType\": \"$certification_type_for_1_2\""

echo "Get compliance-info model with VID: $vid_for_1_4_3 PID: $pid_1_for_1_4_3"
result=$($DCLD_BIN_NEW query compliance compliance-info --vid=$vid_for_1_4_3 --pid=$pid_1_for_1_4_3 --softwareVersion=$software_version_for_1_4_3 --certificationType=$certification_type_for_1_4_3)
check_response "$result" "\"vid\": $vid_for_1_4_3"
check_response "$result" "\"pid\": $pid_1_for_1_4_3"
check_response "$result" "\"softwareVersion\": $software_version_for_1_4_3"
check_response "$result" "\"certificationType\": \"$certification_type_for_1_4_3\""

echo "Get compliance-info model with VID: $vid PID: $pid_1"
result=$($DCLD_BIN_NEW query compliance compliance-info --vid=$vid --pid=$pid_1 --softwareVersion=$software_version --certificationType=$certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_1"
check_response "$result" "\"softwareVersion\": $software_version"
check_response "$result" "\"certificationType\": \"$certification_type\""

echo "Get compliance-info model with VID: $vid_for_1_2 PID: $pid_2_for_1_2"
result=$($DCLD_BIN_NEW query compliance compliance-info --vid=$vid_for_1_2 --pid=$pid_2_for_1_2 --softwareVersion=$software_version_for_1_2 --certificationType=$certification_type_for_1_2)
check_response "$result" "\"vid\": $vid_for_1_2"
check_response "$result" "\"pid\": $pid_2_for_1_2"
check_response "$result" "\"softwareVersion\": $software_version_for_1_2"
check_response "$result" "\"certificationType\": \"$certification_type_for_1_2\""

echo "Get compliance-info model with VID: $vid_for_1_4_3 PID: $pid_2_for_1_4_3"
result=$($DCLD_BIN_NEW query compliance compliance-info --vid=$vid_for_1_4_3 --pid=$pid_2_for_1_4_3 --softwareVersion=$software_version_for_1_4_3 --certificationType=$certification_type_for_1_4_3)
check_response "$result" "\"vid\": $vid_for_1_4_3"
check_response "$result" "\"pid\": $pid_2_for_1_4_3"
check_response "$result" "\"softwareVersion\": $software_version_for_1_4_3"
check_response "$result" "\"certificationType\": \"$certification_type_for_1_4_3\""

echo "Get device software compliance cDCertificateId=$cd_certificate_id"
result=$($DCLD_BIN_NEW query compliance device-software-compliance --cdCertificateId=$cd_certificate_id)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_1"

echo "Get device software compliance cDCertificateId=$cd_certificate_id_for_1_2"
result=$($DCLD_BIN_NEW query compliance device-software-compliance --cdCertificateId=$cd_certificate_id_for_1_2)
check_response "$result" "\"vid\": $vid_for_1_2"
check_response "$result" "\"pid\": $pid_1_for_1_2"

echo "Get device software compliance cDCertificateId=$cd_certificate_id_for_1_4_3"
result=$($DCLD_BIN_NEW query compliance device-software-compliance --cdCertificateId=$cd_certificate_id_for_1_4_3)
check_response "$result" "\"vid\": $vid_for_1_4_3"
check_response "$result" "\"pid\": $pid_1_for_1_4_3"

echo "Get all certified models"
result=$($DCLD_BIN_NEW query compliance all-certified-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_1"
check_response "$result" "\"vid\": $vid_for_1_2"
check_response "$result" "\"pid\": $pid_1_for_1_2"
check_response "$result" "\"vid\": $vid_for_1_4_3"
check_response "$result" "\"pid\": $pid_1_for_1_4_3"

echo "Get all provisional models"
result=$($DCLD_BIN_NEW query compliance all-provisional-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_3"

echo "Get all revoked models"
result=$($DCLD_BIN_NEW query compliance all-revoked-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_2"
check_response "$result" "\"vid\": $vid_for_1_2"
check_response "$result" "\"pid\": $pid_2_for_1_2"
check_response "$result" "\"vid\": $vid_for_1_4_3"
check_response "$result" "\"pid\": $pid_2_for_1_4_3"

echo "Get all compliance infos"
result=$($DCLD_BIN_NEW query compliance all-compliance-info)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_1"
check_response "$result" "\"pid\": $pid_2"
check_response "$result" "\"vid\": $vid_for_1_2"
check_response "$result" "\"pid\": $pid_1_for_1_2"
check_response "$result" "\"pid\": $pid_2_for_1_2"
check_response "$result" "\"vid\": $vid_for_1_4_3"
check_response "$result" "\"pid\": $pid_1_for_1_4_3"
check_response "$result" "\"pid\": $pid_2_for_1_4_3"

echo "Get all device software compliances"
result=$($DCLD_BIN_NEW query compliance all-device-software-compliance)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_1"
check_response "$result" "\"cDCertificateId\": \"$cd_certificate_id\""
check_response "$result" "\"vid\": $vid_for_1_2"
check_response "$result" "\"pid\": $pid_1_for_1_2"
check_response "$result" "\"cDCertificateId\": \"$cd_certificate_id_for_1_2\""
check_response "$result" "\"vid\": $vid_for_1_4_3"
check_response "$result" "\"pid\": $pid_1_for_1_4_3"
check_response "$result" "\"cDCertificateId\": \"$cd_certificate_id_for_1_4_3\""

test_divider

# PKI

echo "Get all x509 certificates"

echo "Get all x509 certificates (GLOBAL)"
result=$($DCLD_BIN_NEW query pki all-certs)
check_response "$result" "\"subjectKeyId\": \"$root_cert_with_vid_subject_key_id_for_1_4_3\""
check_response "$result" "\"subjectKeyId\": \"$test_root_cert_subject_key_id_for_1_2\""
check_response "$result" "\"subjectKeyId\": \"$test_root_cert_subject_key_id\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_for_1_4_3\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_ica_cert_1_subject_key_id_for_1_4_3\""

echo "Get all x509 certificates (DA)"
result=$($DCLD_BIN_NEW query pki all-x509-certs)
check_response "$result" "\"subjectKeyId\": \"$root_cert_with_vid_subject_key_id_for_1_4_3\""
check_response "$result" "\"subjectKeyId\": \"$test_root_cert_subject_key_id_for_1_2\""
check_response "$result" "\"subjectKeyId\": \"$test_root_cert_subject_key_id\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_for_1_4_3\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_ica_cert_1_subject_key_id_for_1_4_3\""

echo "Get all x509 certificates (NOC)"
result=$($DCLD_BIN_NEW query pki all-noc-x509-certs)
check_response "$result" "\[\]"
response_does_not_contain "$result" "\"subjectKeyId\": \"$root_cert_with_vid_subject_key_id_for_1_4_3\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$test_root_cert_subject_key_id_for_1_2\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$test_root_cert_subject_key_id\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_for_1_4_3\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_ica_cert_1_subject_key_id_for_1_4_3\""

echo "Get subject certificates (GLOBAL)"
result=$($DCLD_BIN_NEW query pki all-subject-certs --subject=$root_cert_with_vid_subject_for_1_4_3)
check_response "$result" "$root_cert_with_vid_subject_key_id_for_1_4_3"

result=$($DCLD_BIN_NEW query pki all-subject-certs --subject=$test_root_cert_subject_for_1_2)
check_response "$result" "$test_root_cert_subject_key_id_for_1_2"

result=$($DCLD_BIN_NEW query pki all-subject-certs --subject=$test_root_cert_subject)
check_response "$result" "$test_root_cert_subject_key_id"

result=$($DCLD_BIN_NEW query pki all-subject-certs --subject=$noc_root_cert_1_subject_for_1_4_3)
check_response "$result" "Not Found"

echo "Get subject certificates (DA)"
result=$($DCLD_BIN_NEW query pki all-subject-x509-certs --subject=$root_cert_with_vid_subject_for_1_4_3)
check_response "$result" "$root_cert_with_vid_subject_key_id_for_1_4_3"

result=$($DCLD_BIN_NEW query pki all-subject-x509-certs --subject=$test_root_cert_subject_for_1_2)
check_response "$result" "$test_root_cert_subject_key_id_for_1_2"

result=$($DCLD_BIN_NEW query pki all-subject-x509-certs --subject=$test_root_cert_subject)
check_response "$result" "$test_root_cert_subject_key_id"

result=$($DCLD_BIN_NEW query pki all-subject-x509-certs --subject=$noc_root_cert_1_subject_for_1_4_3)
check_response "$result" "Not Found"

echo "Get subject certificates (NOC)"
result=$($DCLD_BIN_NEW query pki all-noc-subject-x509-certs --subject=$root_cert_with_vid_subject_for_1_4_3)
check_response "$result" "Not Found"

result=$($DCLD_BIN_NEW query pki all-noc-subject-x509-certs --subject=$test_root_cert_subject_for_1_2)
check_response "$result" "Not Found"

result=$($DCLD_BIN_NEW query pki all-noc-subject-x509-certs --subject=$test_root_cert_subject)
check_response "$result" "Not Found"

result=$($DCLD_BIN_NEW query pki all-noc-subject-x509-certs --subject=$noc_root_cert_1_subject_for_1_4_3)
check_response "$result" "Not Found"

echo "Get x509 certificates"

echo "Get x509 certificate (GLOBAL)"
result=$($DCLD_BIN_NEW query pki cert --subject=$root_cert_with_vid_subject_for_1_4_3 --subject-key-id=$root_cert_with_vid_subject_key_id_for_1_4_3)
check_response "$result" "\"subject\": \"$root_cert_with_vid_subject_for_1_4_3\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_with_vid_subject_key_id_for_1_4_3\""
check_response "$result" "\"vid\": $root_cert_vid_for_1_4_3"

result=$($DCLD_BIN_NEW query pki cert --subject=$test_root_cert_subject_for_1_2 --subject-key-id=$test_root_cert_subject_key_id_for_1_2)
check_response "$result" "\"subject\": \"$test_root_cert_subject_for_1_2\""
check_response "$result" "\"subjectKeyId\": \"$test_root_cert_subject_key_id_for_1_2\""
check_response "$result" "\"vid\": $test_root_cert_vid_for_1_2"

result=$($DCLD_BIN_NEW query pki cert --subject=$test_root_cert_subject --subject-key-id=$test_root_cert_subject_key_id)
check_response "$result" "\"subject\": \"$test_root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$test_root_cert_subject_key_id\""
check_response "$result" "\"vid\": $test_root_cert_vid"

result=$($DCLD_BIN_NEW query pki cert --subject=$noc_root_cert_1_subject_for_1_4_3 --subject-key-id=$noc_root_cert_1_subject_key_id_for_1_4_3)
check_response "$result" "Not Found"

echo "Get x509 certificate (DA)"
result=$($DCLD_BIN_NEW query pki x509-cert --subject=$root_cert_with_vid_subject_for_1_4_3 --subject-key-id=$root_cert_with_vid_subject_key_id_for_1_4_3)
check_response "$result" "\"subject\": \"$root_cert_with_vid_subject_for_1_4_3\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_with_vid_subject_key_id_for_1_4_3\""
check_response "$result" "\"vid\": $root_cert_vid_for_1_4_3"

result=$($DCLD_BIN_NEW query pki x509-cert --subject=$test_root_cert_subject_for_1_2 --subject-key-id=$test_root_cert_subject_key_id_for_1_2)
check_response "$result" "\"subject\": \"$test_root_cert_subject_for_1_2\""
check_response "$result" "\"subjectKeyId\": \"$test_root_cert_subject_key_id_for_1_2\""
check_response "$result" "\"vid\": $test_root_cert_vid_for_1_2"

result=$($DCLD_BIN_NEW query pki x509-cert --subject=$test_root_cert_subject --subject-key-id=$test_root_cert_subject_key_id)
check_response "$result" "\"subject\": \"$test_root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$test_root_cert_subject_key_id\""
check_response "$result" "\"vid\": $test_root_cert_vid"

result=$($DCLD_BIN_NEW query pki x509-cert --subject=$noc_root_cert_1_subject_for_1_4_3 --subject-key-id=$noc_root_cert_1_subject_key_id_for_1_4_3)
check_response "$result" "Not Found"

echo "Get x509 certificate (NOC)"
result=$($DCLD_BIN_NEW query pki noc-x509-cert --subject="$root_cert_with_vid_subject_for_1_4_3" --subject-key-id="$root_cert_with_vid_subject_key_id_for_1_4_3")
check_response "$result" "Not Found"

result=$($DCLD_BIN_NEW query pki noc-x509-cert --subject="$test_root_cert_subject_for_1_2" --subject-key-id="$test_root_cert_subject_key_id_for_1_2")
check_response "$result" "Not Found"

result=$($DCLD_BIN_NEW query pki noc-x509-cert --subject="$test_root_cert_subject" --subject-key-id="$test_root_cert_subject_key_id")
check_response "$result" "Not Found"

result=$($DCLD_BIN_NEW query pki noc-x509-cert --subject="$noc_root_cert_1_subject_for_1_4_3" --subject-key-id="$noc_root_cert_1_subject_key_id_for_1_4_3")
check_response "$result" "Not Found"

echo "Get all x509 certificates by subjectKeyId"
result=$($DCLD_BIN_NEW query pki cert --subject-key-id="$root_cert_with_vid_subject_key_id_for_1_4_3")
check_response "$result" "$root_cert_with_vid_subject_for_1_4_3"
check_response "$result" "\"subjectKeyId\": \"$root_cert_with_vid_subject_key_id_for_1_4_3\""

result=$($DCLD_BIN_NEW query pki cert --subject-key-id="$test_root_cert_subject_key_id_for_1_2")
check_response "$result" "$test_root_cert_subject_for_1_2"
check_response "$result" "\"subjectKeyId\": \"$test_root_cert_subject_key_id_for_1_2\""

result=$($DCLD_BIN_NEW query pki cert --subject-key-id="$test_root_cert_subject_key_id")
check_response "$result" "$test_root_cert_subject"
check_response "$result" "\"subjectKeyId\": \"$test_root_cert_subject_key_id\""

result=$($DCLD_BIN_NEW query pki cert --subject-key-id="$noc_root_cert_1_subject_key_id_for_1_4_3")
check_response "$result" "Not Found"

result=$($DCLD_BIN_NEW query pki x509-cert --subject-key-id="$root_cert_with_vid_subject_key_id_for_1_4_3")
check_response "$result" "$root_cert_with_vid_subject_for_1_4_3"
check_response "$result" "\"subjectKeyId\": \"$root_cert_with_vid_subject_key_id_for_1_4_3\""

result=$($DCLD_BIN_NEW query pki x509-cert --subject-key-id="$test_root_cert_subject_key_id_for_1_2")
check_response "$result" "$test_root_cert_subject_for_1_2"
check_response "$result" "\"subjectKeyId\": \"$test_root_cert_subject_key_id_for_1_2\""

result=$($DCLD_BIN_NEW query pki x509-cert --subject-key-id="$test_root_cert_subject_key_id")
check_response "$result" "$test_root_cert_subject"
check_response "$result" "\"subjectKeyId\": \"$test_root_cert_subject_key_id\""

result=$($DCLD_BIN_NEW query pki x509-cert --subject-key-id="$noc_root_cert_1_subject_key_id_for_1_4_3")
check_response "$result" "Not Found"

result=$($DCLD_BIN_NEW query pki proposed-x509-root-cert --subject=$google_root_cert_subject_for_1_2 --subject-key-id=$google_root_cert_subject_key_id_for_1_2)
check_response "$result" "\"subject\": \"$google_root_cert_subject_for_1_2\""
check_response "$result" "\"subjectKeyId\": \"$google_root_cert_subject_key_id_for_1_2\""
check_response "$result" "\"vid\": $google_root_cert_path_random_vid_for_1_2"

result=$($DCLD_BIN_NEW query pki proposed-x509-root-cert --subject=$google_root_cert_subject --subject-key-id=$google_root_cert_subject_key_id)
check_response "$result" "\"subject\": \"$google_root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$google_root_cert_subject_key_id\""

echo "Get revoked x509 certificate"
result=$($DCLD_BIN_NEW query pki revoked-x509-cert --subject=$intermediate_cert_with_vid_subject_for_1_4_3 --subject-key-id=$intermediate_cert_with_vid_subject_key_id_for_1_4_3)
check_response "$result" "\"subject\": \"$intermediate_cert_with_vid_subject_for_1_4_3\""
check_response "$result" "\"subjectKeyId\": \"$intermediate_cert_with_vid_subject_key_id_for_1_4_3\""
check_response "$result" "\"vid\": $intermediate_cert_with_vid_65521_vid_for_1_4_3"

result=$($DCLD_BIN_NEW query pki revoked-x509-cert --subject=$intermediate_cert_subject_for_1_2 --subject-key-id=$intermediate_cert_subject_key_id_for_1_2)
check_response "$result" "\"subject\": \"$intermediate_cert_subject_for_1_2\""
check_response "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id_for_1_2\""

result=$($DCLD_BIN_NEW query pki revoked-x509-cert --subject=$intermediate_cert_subject --subject-key-id=$intermediate_cert_subject_key_id)
check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""

result=$($DCLD_BIN_NEW query pki revoked-x509-cert --subject=$noc_root_cert_1_subject_for_1_4_3 --subject-key-id=$noc_root_cert_1_subject_key_id_for_1_4_3)
check_response "$result" "Not Found"

echo "Get proposed x509 root certificate to revoke"
result=$($DCLD_BIN_NEW query pki proposed-x509-root-cert-to-revoke --subject=$root_cert_with_vid_subject_for_1_4_3 --subject-key-id=$root_cert_with_vid_subject_key_id_for_1_4_3)
check_response "$result" "\"subject\": \"$root_cert_with_vid_subject_for_1_4_3\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_with_vid_subject_key_id_for_1_4_3\""

result=$($DCLD_BIN_NEW query pki proposed-x509-root-cert-to-revoke --subject=$test_root_cert_subject_for_1_2 --subject-key-id=$test_root_cert_subject_key_id_for_1_2)
check_response "$result" "\"subject\": \"$test_root_cert_subject_for_1_2\""
check_response "$result" "\"subjectKeyId\": \"$test_root_cert_subject_key_id_for_1_2\""

result=$($DCLD_BIN_NEW query pki proposed-x509-root-cert-to-revoke --subject=$test_root_cert_subject --subject-key-id=$test_root_cert_subject_key_id)
check_response "$result" "\"subject\": \"$test_root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$test_root_cert_subject_key_id\""

echo "Get revocation point"
result=$($DCLD_BIN_NEW query pki revocation-point --vid=$vid_for_1_2 --label=$product_label_for_1_2 --issuer-subject-key-id=$issuer_subject_key_id)
check_response "$result" "\"vid\": $vid_for_1_2"
check_response "$result" "\"issuerSubjectKeyID\": \"$issuer_subject_key_id\""
check_response "$result" "\"label\": \"$product_label_for_1_2\""
check_response "$result" "\"dataURL\": \"$test_data_url\""

echo "Get revocation points by issuer subject key id"
result=$($DCLD_BIN_NEW query pki revocation-points --issuer-subject-key-id=$issuer_subject_key_id)
check_response "$result" "\"vid\": $vid_for_1_2"
check_response "$result" "\"issuerSubjectKeyID\": \"$issuer_subject_key_id\""
check_response "$result" "\"label\": \"$product_label_for_1_2\""
check_response "$result" "\"dataURL\": \"$test_data_url\""

echo "Get all proposed x509 root certificates"
result=$($DCLD_BIN_NEW query pki all-proposed-x509-root-certs)
check_response "$result" "\"subject\": \"$google_root_cert_subject_for_1_2\""
check_response "$result" "\"subjectKeyId\": \"$google_root_cert_subject_key_id_for_1_2\""
check_response "$result" "\"subject\": \"$google_root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$google_root_cert_subject_key_id\""

echo "Get all revoked x509 root certificates"
result=$($DCLD_BIN_NEW query pki all-revoked-x509-root-certs)
check_response "$result" "\"subject\": \"$root_cert_subject_for_1_2\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id_for_1_2\""
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""

echo "Get all proposed x509 root certificates to revoke"
result=$($DCLD_BIN_NEW query pki all-proposed-x509-root-certs-to-revoke)
check_response "$result" "\"subject\": \"$test_root_cert_subject_for_1_2\""
check_response "$result" "\"subjectKeyId\": \"$test_root_cert_subject_key_id_for_1_2\""
check_response "$result" "\"subject\": \"$test_root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$test_root_cert_subject_key_id\""

echo "Get all revocation points"
result=$($DCLD_BIN_NEW query pki all-revocation-points)
check_response "$result" "\"vid\": $vid_for_1_2"
check_response "$result" "\"issuerSubjectKeyID\": \"$issuer_subject_key_id\""
check_response "$result" "\"label\": \"$product_label_for_1_2\""
check_response "$result" "\"dataURL\": \"$test_data_url\""
check_response "$result" "\"vid\": $vid_for_1_4_3"
check_response "$result" "\"label\": \"$product_label_for_1_4_3\""
check_response "$result" "\"dataURL\": \"$test_data_url_for_1_4_3\""

echo "Get all noc certificates"
result=$($DCLD_BIN_NEW query pki all-noc-x509-certs)
check_response "$result" "\[\]"
response_does_not_contain "$result" "$noc_root_cert_1_subject_key_id_for_1_4_3"
response_does_not_contain "$result" "$noc_ica_cert_1_subject_key_id_for_1_4_3"

echo "Get all noc x509 root certificates"
result=$($DCLD_BIN_NEW query pki noc-x509-root-certs --vid=$vid_for_1_4_3)
check_response "$result" "Not Found"
response_does_not_contain "$result" "$noc_root_cert_1_subject_key_id_for_1_4_3"

echo "Get all noc x509 root certificates by vid=$vid_for_1_4_3 and skid=$noc_root_cert_1_subject_key_id_for_1_4_3 (must be empty)"
result=$($DCLD_BIN_NEW query pki noc-x509-cert --vid=$vid_for_1_4_3 --subject-key-id=$noc_root_cert_1_subject_key_id_for_1_4_3)
check_response "$result" "Not Found"

echo "Get noc x509 root certificate by subject and subject key id"
result=$($DCLD_BIN_NEW query pki noc-x509-cert --subject="$noc_root_cert_1_subject_for_1_4_3" --subject-key-id="$noc_root_cert_1_subject_key_id_for_1_4_3")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id_for_1_4_3\""

echo "Get noc x509 ica certificate  by subject and subject key id"
result=$($DCLD_BIN_NEW query pki noc-x509-cert --subject="$noc_ica_cert_1_subject_for_1_4_3" --subject-key-id="$noc_ica_cert_1_subject_key_id_for_1_4_3")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_ica_cert_1_subject_key_id_for_1_4_3\""

test_divider

# AUTH

echo "Get all accounts"
result=$($DCLD_BIN_NEW query auth all-accounts)
check_response "$result" "\"address\": \"$user_5_address\""
check_response "$result" "\"address\": \"$user_2_address\""

echo "Get account"
result=$($DCLD_BIN_NEW query auth account --address=$user_5_address)
check_response "$result" "\"address\": \"$user_5_address\""

result=$($DCLD_BIN_NEW query auth account --address=$user_2_address)
check_response "$result" "\"address\": \"$user_2_address\""

echo "Get all proposed accounts"
result=$($DCLD_BIN_NEW query auth all-proposed-accounts)
check_response "$result" "\"address\": \"$user_6_address\""
check_response "$result" "\"address\": \"$user_3_address\""

echo "Get proposed account"
result=$($DCLD_BIN_NEW query auth proposed-account --address=$user_6_address)
check_response "$result" "\"address\": \"$user_6_address\""

result=$($DCLD_BIN_NEW query auth proposed-account --address=$user_3_address)
check_response "$result" "\"address\": \"$user_3_address\""

echo "Get all proposed accounts to revoke"
result=$($DCLD_BIN_NEW query auth all-proposed-accounts-to-revoke)
check_response "$result" "\"address\": \"$user_5_address\""
check_response "$result" "\"address\": \"$user_2_address\""

echo "Get proposed account to revoke"
result=$($DCLD_BIN_NEW query auth proposed-account-to-revoke --address=$user_5_address)
check_response "$result" "\"address\": \"$user_5_address\""

result=$($DCLD_BIN_NEW query auth proposed-account-to-revoke --address=$user_2_address)
check_response "$result" "\"address\": \"$user_2_address\""

echo "Get all revoked accounts"
result=$($DCLD_BIN_NEW query auth all-revoked-accounts)
check_response "$result" "\"address\": \"$user_4_address\""
check_response "$result" "\"address\": \"$user_1_address\""

echo "Get revoked account"
result=$($DCLD_BIN_NEW query auth revoked-account --address=$user_4_address)
check_response "$result" "\"address\": \"$user_4_address\""

result=$($DCLD_BIN_NEW query auth revoked-account --address=$user_1_address)
check_response "$result" "\"address\": \"$user_1_address\""

test_divider

# Validator

echo "Get node"
# FIXME: use proper binary (not dcld but $DCLD_BIN_OLD)
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator all-nodes")
check_response "$result" "\"owner\": \"$validator_address\""

########################################################################################

# after upgrade constants

vid_for_1_4_4=65522
pid_1_for_1_4_4=77
pid_2_for_1_4_4=88
pid_3_for_1_4_4=99
device_type_id_for_1_4_4=4433
product_name_for_1_4_4="ProductName1.4.4"
product_label_for_1_4_4="ProductLabel1.4.4"
part_number_for_1_4_4="RCU2245B"
software_version_for_1_4_4=2
software_version_string_for_1_4_4="4.0"
cd_version_number_for_1_4_4=513
min_applicable_software_version_for_1_4_4=4
max_applicable_software_version_for_1_4_4=4000

certification_type_for_1_4_4="matter"
certification_date_for_1_4_4="2023-01-01T00:00:00Z"
provisional_date_for_1_4_4="2014-12-12T00:00:00Z"
cd_certificate_id_for_1_4_4="20DEXC"

da_root_cert_1_path_for_1_4_4="integration_tests/constants/upgrade_1_4_4_da_root_cert"
da_root_cert_1_subject_for_1_4_4="MEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQ="
da_root_cert_1_subject_key_id_for_1_4_4="A8:A0:95:18:9B:9F:81:4D:C7:9F:5E:B5:82:09:27:95:13:0C:9F:87"

da_intermediate_cert_1_path_for_1_4_4="integration_tests/constants/upgrade_1_4_4_da_intermediate_cert"
da_intermediate_cert_1_subject_for_1_4_4="MIGCMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTETMBEGA1UEBwwKU29tZSBTdGF0ZTEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRQwEgYDVQQDDAtOT0MtY2hpbGQtMw=="
da_intermediate_cert_1_subject_key_id_for_1_4_4="A8:A0:95:18:9B:9F:81:4D:C7:9F:5E:B5:82:09:27:95:13:0C:9F:87"
da_intermediate_cert_1_serial_number_for_1_4_4="3"

da_root_cert_2_path_for_1_4_4="integration_tests/constants/upgrade_1_4_4_da_root_cert_2"
da_root_cert_2_subject_for_1_4_4="MDsxCzAJBgNVBAYTAlRFMRMwEQYDVQQIDApTb21lLVN0YXRlMRcwFQYDVQQKDA5VcGdyYWRlMS40LjRfMQ=="
da_root_cert_2_subject_key_id_for_1_4_4="A8:A0:95:18:9B:9F:81:4D:C7:9F:5E:B5:82:09:27:95:13:0C:9F:87"

da_intermediate_cert_2_path_for_1_4_4="integration_tests/constants/upgrade_1_4_4_da_intermediate_cert_2"
da_intermediate_cert_2_subject_for_1_4_4="MIGBMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTETMBEGA1UEBwwKU29tZSBTdGF0ZTEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRMwEQYDVQQDDApEQS1jaGlsZC0z"
da_intermediate_cert_2_subject_key_id_for_1_4_4="A8:A0:95:18:9B:9F:81:4D:C7:9F:5E:B5:82:09:27:95:13:0C:9F:87"
da_intermediate_cert_2_serial_number_for_1_4_4="3"
da_intermediate_cert_2_vid_for_1_4_4=65521

noc_root_cert_1_path_for_1_4_4="integration_tests/constants/noc_root_cert_2"
noc_root_cert_1_subject_for_1_4_4="MHoxCzAJBgNVBAYTAlVaMRMwEQYDVQQIDApTb21lIFN0YXRlMREwDwYDVQQHDAhUYXNoa2VudDEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMQ4wDAYDVQQDDAVOT0MtMg=="
noc_root_cert_1_subject_key_id_for_1_4_4="CF:E6:DD:37:2B:4C:B2:B9:A9:F2:75:30:1C:AA:B1:37:1B:11:7F:1B"

noc_ica_cert_1_path_for_1_4_4="integration_tests/constants/noc_cert_2"
noc_ica_cert_1_subject_for_1_4_4="MIGCMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTETMBEGA1UEBwwKU29tZSBTdGF0ZTEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRQwEgYDVQQDDAtOT0MtY2hpbGQtMg=="
noc_ica_cert_1_subject_key_id_for_1_4_4="87:48:A2:33:12:1F:51:5C:93:E6:90:40:4A:2C:AB:9E:D6:19:E5:AD"

noc_root_cert_2_path_for_1_4_4="integration_tests/constants/noc_root_cert_3"
noc_root_cert_2_subject_for_1_4_4="MFUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQxDjAMBgNVBAMMBU5PQy0z"
noc_root_cert_2_subject_key_id_for_1_4_4="88:0D:06:D9:64:22:29:34:78:7F:8C:3B:AE:F5:08:93:86:8F:0D:20"

noc_ica_cert_2_path_for_1_4_4="integration_tests/constants/noc_cert_3"
noc_ica_cert_2_subject_for_1_4_4="MIGCMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTETMBEGA1UEBwwKU29tZSBTdGF0ZTEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRQwEgYDVQQDDAtOT0MtY2hpbGQtMw=="
noc_ica_cert_2_subject_key_id_for_1_4_4="DE:F9:1D:90:D5:A1:0F:23:59:5C:3F:5C:C7:2D:31:58:2F:A8:79:33"

test_data_url_for_1_4_4="https://url.data.dclmodel-1.4.4"

vendor_name_for_1_4_4="Vendor65522"
company_legal_name_for_1_4_4="LegalCompanyName65522"
company_preferred_name_for_1_4_4="CompanyPreferredName65522"
vendor_landing_page_url_for_1_4_4="https://www.new65522example.com"

vendor_account_for_1_4_4="vendor_account_65522"

echo "Create Vendor account $vendor_account_for_1_4_4"

result="$(echo $passphrase | $DCLD_BIN_NEW keys add "$vendor_account_for_1_4_4")"
echo "keys add $result"
_address=$(echo $passphrase | $DCLD_BIN_NEW keys show $vendor_account_for_1_4_4 -a)
_pubkey=$(echo $passphrase | $DCLD_BIN_NEW keys show $vendor_account_for_1_4_4 -p)
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth propose-add-account --address="$_address" --pubkey="$_pubkey" --vid="$vid_for_1_4_4" --roles="Vendor" --from "$trustee_account_1" --yes)
echo "propose-add-account $result"
result=$(get_txn_result "$result")
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth approve-add-account --address="$_address" --from "$trustee_account_2" --yes)
result=$(get_txn_result "$result")
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth approve-add-account --address="$_address" --from "$trustee_account_3" --yes)
result=$(get_txn_result "$result")
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth approve-add-account --address="$_address" --from "$trustee_account_4" --yes)
result=$(get_txn_result "$result")

random_string user_10
echo "$user_10 generates keys"
cmd="(echo $passphrase; echo $passphrase) | $DCLD_BIN_NEW keys add $user_10"
result="$(bash -c "$cmd")"
user_10_address=$(echo $passphrase | $DCLD_BIN_NEW keys show $user_10 -a)
user_10_pubkey=$(echo $passphrase | $DCLD_BIN_NEW keys show $user_10 -p)

random_string user_11
echo "$user_11 generates keys"
cmd="(echo $passphrase; echo $passphrase) | $DCLD_BIN_NEW keys add $user_11"
result="$(bash -c "$cmd")"
user_11_address=$(echo $passphrase | $DCLD_BIN_NEW keys show $user_11 -a)
user_11_pubkey=$(echo $passphrase | $DCLD_BIN_NEW keys show $user_11 -p)

random_string user_12
echo "$user_12 generates keys"
cmd="(echo $passphrase; echo $passphrase) | $DCLD_BIN_NEW keys add $user_12"
result="$(bash -c "$cmd")"
user_12_address=$(echo $passphrase | $DCLD_BIN_NEW keys show $user_12 -a)
user_12_pubkey=$(echo $passphrase | $DCLD_BIN_NEW keys show $user_12 -p)

# send all ledger update transactions after upgrade

# VENDOR_INFO
echo "Add vendor $vendor_name_for_1_4_4"
result=$(echo $passphrase | $DCLD_BIN_NEW tx vendorinfo add-vendor --vid=$vid_for_1_4_4 --vendorName=$vendor_name_for_1_4_4 --companyLegalName=$company_legal_name_for_1_4_4 --companyPreferredName=$company_preferred_name_for_1_4_4 --vendorLandingPageURL=$vendor_landing_page_url_for_1_4_4 --from=$vendor_account_for_1_4_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Update vendor $vendor_name_for_1_2"
result=$(echo $passphrase | $DCLD_BIN_NEW tx vendorinfo update-vendor --vid=$vid_for_1_2 --vendorName=$vendor_name_for_1_2 --companyLegalName=$company_legal_name_for_1_2 --companyPreferredName=$company_preferred_name_for_1_4_4 --vendorLandingPageURL=$vendor_landing_page_url_for_1_4_4 --from=$vendor_account_for_1_2 --yes)
result=$(get_txn_result "$result")
echo $result
check_response "$result" "\"code\": 0"

test_divider

# MODEL and MODEL_VERSION

echo "Add model vid=$vid_for_1_4_4 pid=$pid_1_for_1_4_4"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model --vid=$vid_for_1_4_4 --pid=$pid_1_for_1_4_4 --deviceTypeID=$device_type_id_for_1_4_4 --productName=$product_name_for_1_4_4 --productLabel=$product_label_for_1_4_4 --partNumber=$part_number_for_1_4_4 --from=$vendor_account_for_1_4_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Add model version vid=$vid_for_1_4_4 pid=$pid_1_for_1_4_4"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model-version --vid=$vid_for_1_4_4 --pid=$pid_1_for_1_4_4 --softwareVersion=$software_version_for_1_4_4 --softwareVersionString=$software_version_string_for_1_4_4 --cdVersionNumber=$cd_version_number_for_1_4_4 --minApplicableSoftwareVersion=$min_applicable_software_version_for_1_4_4 --maxApplicableSoftwareVersion=$max_applicable_software_version_for_1_4_4 --from=$vendor_account_for_1_4_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Add model vid=$vid_for_1_4_4 pid=$pid_2_for_1_4_4"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model --vid=$vid_for_1_4_4 --pid=$pid_2_for_1_4_4 --deviceTypeID=$device_type_id_for_1_4_4 --productName=$product_name_for_1_4_4 --productLabel=$product_label_for_1_4_4 --partNumber=$part_number_for_1_4_4 --from=$vendor_account_for_1_4_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Add model version vid=$vid_for_1_4_4 pid=$pid_2_for_1_4_4"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model-version --vid=$vid_for_1_4_4 --pid=$pid_2_for_1_4_4 --softwareVersion=$software_version_for_1_4_4 --softwareVersionString=$software_version_string_for_1_4_4 --cdVersionNumber=$cd_version_number_for_1_4_4 --minApplicableSoftwareVersion=$min_applicable_software_version_for_1_4_4 --maxApplicableSoftwareVersion=$max_applicable_software_version_for_1_4_4 --from=$vendor_account_for_1_4_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Add model vid=$vid_for_1_4_4 pid=$pid_3_for_1_4_4"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model --vid=$vid_for_1_4_4 --pid=$pid_3_for_1_4_4 --deviceTypeID=$device_type_id_for_1_4_4 --productName=$product_name_for_1_4_4 --productLabel=$product_label_for_1_4_4 --partNumber=$part_number_for_1_4_4 --from=$vendor_account_for_1_4_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Add model version vid=$vid_for_1_4_4 pid=$pid_3_for_1_4_4"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model-version --vid=$vid_for_1_4_4 --pid=$pid_3_for_1_4_4 --softwareVersion=$software_version_for_1_4_4 --softwareVersionString=$software_version_string_for_1_4_4 --cdVersionNumber=$cd_version_number_for_1_4_4 --minApplicableSoftwareVersion=$min_applicable_software_version_for_1_4_4 --maxApplicableSoftwareVersion=$max_applicable_software_version_for_1_4_4 --from=$vendor_account_for_1_4_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Delete model vid=$vid_for_1_4_4 pid=$pid_3_for_1_4_4"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model delete-model --vid=$vid_for_1_4_4 --pid=$pid_3_for_1_4_4 --from=$vendor_account_for_1_4_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Update model vid=$vid pid=$pid_2"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model update-model --vid=$vid --pid=$pid_2 --productName=$product_name --productLabel=$product_label_for_1_4_4 --partNumber=$part_number_for_1_4_4 --from=$vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Update model version vid=$vid pid=$pid_2"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model update-model-version --vid=$vid --pid=$pid_2 --softwareVersion=$software_version --minApplicableSoftwareVersion=$min_applicable_software_version_for_1_4_4 --maxApplicableSoftwareVersion=$max_applicable_software_version_for_1_4_4 --from=$vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

# CERTIFY_DEVICE_COMPLIANCE

echo "Certify model vid=$vid_for_1_4_4 pid=$pid_1_for_1_4_4"
result=$(echo $passphrase | $DCLD_BIN_NEW tx compliance certify-model --vid=$vid_for_1_4_4 --pid=$pid_1_for_1_4_4 --softwareVersion=$software_version_for_1_4_4 --softwareVersionString=$software_version_string_for_1_4_4  --certificationType=$certification_type_for_1_4_4 --certificationDate=$certification_date_for_1_4_4 --cdCertificateId=$cd_certificate_id_for_1_4_4 --from=$certification_center_account --cdVersionNumber=$cd_version_number_for_1_4_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Provision model vid=$vid_for_1_4_4 pid=$pid_2_for_1_4_4"
result=$(echo $passphrase | $DCLD_BIN_NEW tx compliance provision-model --vid=$vid_for_1_4_4 --pid=$pid_2_for_1_4_4 --softwareVersion=$software_version_for_1_4_4 --softwareVersionString=$software_version_string_for_1_4_4 --certificationType=$certification_type_for_1_4_4 --provisionalDate=$provisional_date_for_1_4_4 --cdCertificateId=$cd_certificate_id_for_1_4_4 --from=$certification_center_account --cdVersionNumber=$cd_version_number_for_1_4_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Certify model vid=$vid_for_1_4_4 pid=$pid_2_for_1_4_4"
result=$(echo $passphrase | $DCLD_BIN_NEW tx compliance certify-model --vid=$vid_for_1_4_4 --pid=$pid_2_for_1_4_4 --softwareVersion=$software_version_for_1_4_4 --softwareVersionString=$software_version_string_for_1_4_4  --certificationType=$certification_type_for_1_4_4 --certificationDate=$certification_date_for_1_4_4 --cdCertificateId=$cd_certificate_id_for_1_4_4 --from=$certification_center_account --cdVersionNumber=$cd_version_number_for_1_4_4  --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Revoke model certification vid=$vid_for_1_4_4 pid=$pid_2_for_1_4_4"
result=$(echo $passphrase | $DCLD_BIN_NEW tx compliance revoke-model --vid=$vid_for_1_4_4 --pid=$pid_2_for_1_4_4 --softwareVersion=$software_version_for_1_4_4 --softwareVersionString=$software_version_string_for_1_4_4 --certificationType=$certification_type_for_1_4_4 --revocationDate=$certification_date_for_1_4_4 --from=$certification_center_account --cdVersionNumber=$cd_version_number_for_1_4_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

# X509 PKI

echo "Propose add da_root_cert_1_path_for_1_4_4"
result=$(echo $passphrase | $DCLD_BIN_NEW tx pki propose-add-x509-root-cert --certificate="$da_root_cert_1_path_for_1_4_4" --vid="$vid_for_1_4_4" --from=$trustee_account_1 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add da_root_cert_1"
result=$(echo $passphrase | $DCLD_BIN_NEW tx pki approve-add-x509-root-cert --subject="$da_root_cert_1_subject_for_1_4_4" --subject-key-id=$da_root_cert_1_subject_key_id_for_1_4_4 --from=$trustee_account_2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "reject add da_root_cert_1"
result=$(echo $passphrase | $DCLD_BIN_NEW tx pki reject-add-x509-root-cert --subject="$da_root_cert_1_subject_for_1_4_4" --subject-key-id=$da_root_cert_1_subject_key_id_for_1_4_4 --from=$trustee_account_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add da_root_cert_1"
result=$(echo $passphrase | $DCLD_BIN_NEW tx pki approve-add-x509-root-cert --subject="$da_root_cert_1_subject_for_1_4_4" --subject-key-id=$da_root_cert_1_subject_key_id_for_1_4_4 --from=$trustee_account_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add da_root_cert_1"
result=$(echo $passphrase | $DCLD_BIN_NEW tx pki approve-add-x509-root-cert --subject="$da_root_cert_1_subject_for_1_4_4" --subject-key-id=$da_root_cert_1_subject_key_id_for_1_4_4 --from=$trustee_account_5 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Propose add da_root_cert_2"
result=$(echo $passphrase | $DCLD_BIN_NEW tx pki propose-add-x509-root-cert --certificate="$da_root_cert_2_path_for_1_4_4" --vid="$vid_for_1_4_4" --from=$trustee_account_5 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add da_root_cert_2"
result=$(echo $passphrase | $DCLD_BIN_NEW tx pki approve-add-x509-root-cert --subject="$da_root_cert_2_subject_for_1_4_4" --subject-key-id=$da_root_cert_2_subject_key_id_for_1_4_4 --from=$trustee_account_1 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

result=$(echo $passphrase | $DCLD_BIN_NEW tx pki approve-add-x509-root-cert --subject="$da_root_cert_2_subject_for_1_4_4" --subject-key-id=$da_root_cert_2_subject_key_id_for_1_4_4 --from=$trustee_account_2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

result=$(echo $passphrase | $DCLD_BIN_NEW tx pki approve-add-x509-root-cert --subject="$da_root_cert_2_subject_for_1_4_4" --subject-key-id=$da_root_cert_2_subject_key_id_for_1_4_4 --from=$trustee_account_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Add da_intermediate_cert_1"
result=$(echo $passphrase | $DCLD_BIN_NEW tx pki add-x509-cert --certificate="$da_intermediate_cert_1_path_for_1_4_4" --from=$vendor_account_for_1_4_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Add da_intermediate_cert_2"
result=$(echo $passphrase | $DCLD_BIN_NEW tx pki add-x509-cert --certificate="$da_intermediate_cert_2_path_for_1_4_4" --from=$vendor_account_for_1_4_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Propose revoke paa_cert_no_vid"
result=$(echo "$passphrase" | $DCLD_BIN_NEW tx pki propose-revoke-x509-root-cert --subject="$da_root_cert_1_subject_for_1_4_4" --subject-key-id="$da_root_cert_1_subject_key_id_for_1_4_4" --from="$trustee_account_1" --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve revoke paa_cert_no_vid"
result=$(echo "$passphrase" | $DCLD_BIN_NEW tx pki approve-revoke-x509-root-cert --subject="$da_root_cert_1_subject_for_1_4_4" --subject-key-id="$da_root_cert_1_subject_key_id_for_1_4_4" --from="$trustee_account_2" --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve revoke $da_intermediate_cert_1_path_for_1_4_4"
result=$(echo "$passphrase" | $DCLD_BIN_NEW tx pki approve-revoke-x509-root-cert --subject="$da_root_cert_1_subject_for_1_4_4" --subject-key-id="$da_root_cert_1_subject_key_id_for_1_4_4" --from="$trustee_account_3" --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve revoke $da_intermediate_cert_1_path_for_1_4_4"
result=$(echo "$passphrase" | $DCLD_BIN_NEW tx pki approve-revoke-x509-root-cert --subject="$da_root_cert_1_subject_for_1_4_4" --subject-key-id="$da_root_cert_1_subject_key_id_for_1_4_4" --from="$trustee_account_4" --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Propose revoke da_root_cert_2"
result=$(echo $passphrase | $DCLD_BIN_NEW tx pki propose-revoke-x509-root-cert --subject="$da_root_cert_2_subject_for_1_4_4" --subject-key-id="$da_root_cert_2_subject_key_id_for_1_4_4" --from $trustee_account_1 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Revoke da_intermediate_cert_1"
result=$(echo $passphrase | $DCLD_BIN_NEW tx pki revoke-x509-cert --subject="$da_intermediate_cert_1_subject_for_1_4_4" --subject-key-id="$da_intermediate_cert_1_subject_key_id_for_1_4_4" --from=$vendor_account_for_1_4_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Add NOC Root certificate by vendor with VID = $vid_for_1_4_4"
result=$(echo "$passphrase" | $DCLD_BIN_NEW tx pki add-noc-x509-root-cert --certificate="$noc_root_cert_1_path_for_1_4_4" --from $vendor_account_for_1_4_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Add NOC ICA certificate by vendor with VID = $vid_for_1_4_4"
result=$(echo "$passphrase" | $DCLD_BIN_NEW tx pki add-noc-x509-ica-cert --certificate="$noc_ica_cert_1_path_for_1_4_4" --from $vendor_account_for_1_4_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Add NOC Root certificate by vendor with VID = $vid_for_1_4_4"
result=$(echo "$passphrase" | $DCLD_BIN_NEW tx pki add-noc-x509-root-cert --certificate="$noc_root_cert_2_path_for_1_4_4" --from $vendor_account_for_1_4_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Add NOC ICA certificate by vendor with VID = $vid_for_1_4_4"
result=$(echo "$passphrase" | $DCLD_BIN_NEW tx pki add-noc-x509-ica-cert --certificate="$noc_ica_cert_2_path_for_1_4_4" --from $vendor_account_for_1_4_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Revoke NOC root certificate by vendor with VID = $vid_for_1_4_4"
result=$(echo "$passphrase" | $DCLD_BIN_NEW tx pki revoke-noc-x509-root-cert --subject="$noc_root_cert_1_subject_for_1_4_4" --subject-key-id="$noc_root_cert_1_subject_key_id_for_1_4_4" --from=$vendor_account_for_1_4_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Revoke NOC ICA certificate by vendor with VID = $vid_for_1_4_4"
result=$(echo "$passphrase" | $DCLD_BIN_NEW tx pki revoke-noc-x509-ica-cert --subject="$noc_ica_cert_1_subject_for_1_4_4" --subject-key-id="$noc_ica_cert_1_subject_key_id_for_1_4_4" --from=$vendor_account_for_1_4_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

# PKI Revocation point

echo "Add new revocation point for vid_for_1_4_4"
result=$(echo $passphrase | $DCLD_BIN_NEW tx pki add-revocation-point --vid=$vid_for_1_4_4 --revocation-type=1 --is-paa="true" --certificate="$da_root_cert_2_path_for_1_4_4" --label="$product_label_for_1_4_4" --data-url="$test_data_url_for_1_4_4" --issuer-subject-key-id=$issuer_subject_key_id --from=$vendor_account_for_1_4_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Update revocation point"
result=$(echo $passphrase | $DCLD_BIN_NEW tx pki update-revocation-point --vid=$vid_for_1_4_4 --certificate="$da_root_cert_2_path_for_1_4_4" --label="$product_label_for_1_4_4" --data-url="$test_data_url_for_1_4_4/new" --issuer-subject-key-id=$issuer_subject_key_id --from=$vendor_account_for_1_4_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Delete revocation point"
result=$(echo $passphrase | $DCLD_BIN_NEW tx pki delete-revocation-point --vid=$vid_for_1_4_4 --label="$product_label_for_1_4_4" --issuer-subject-key-id=$issuer_subject_key_id --from=$vendor_account_for_1_4_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Add new revocation point"
result=$(echo $passphrase | $DCLD_BIN_NEW tx pki add-revocation-point --vid=$vid_for_1_4_4 --revocation-type=1 --is-paa="true" --certificate="$da_root_cert_2_path_for_1_4_4" --label="$product_label_for_1_4_4" --data-url="$test_data_url_for_1_4_4" --issuer-subject-key-id=$issuer_subject_key_id --from=$vendor_account_for_1_4_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

# AUTH

echo "Propose add account $user_10_address"
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth propose-add-account --address="$user_10_address" --pubkey="$user_10_pubkey" --roles="CertificationCenter" --from="$trustee_account_1" --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_10_address"
result=$($DCLD_BIN_NEW tx auth approve-add-account --address="$user_10_address" --from=$trustee_account_2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_10_address"
result=$($DCLD_BIN_NEW tx auth approve-add-account --address="$user_10_address" --from=$trustee_account_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_10_address"
result=$($DCLD_BIN_NEW tx auth approve-add-account --address="$user_10_address" --from=$trustee_account_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Propose add account $user_11_address"
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth propose-add-account --address="$user_11_address" --pubkey=$user_11_pubkey --roles=CertificationCenter --from=$trustee_account_1 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_11_address"
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth approve-add-account --address="$user_11_address" --from=$trustee_account_2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_11_address"
result=$($DCLD_BIN_NEW tx auth approve-add-account --address="$user_11_address" --from=$trustee_account_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_11_address"
result=$($DCLD_BIN_NEW tx auth approve-add-account --address="$user_11_address" --from=$trustee_account_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Propose add account $user_12_address"
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth propose-add-account --address="$user_12_address" --pubkey=$user_12_pubkey --roles=CertificationCenter --from=$trustee_account_1 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Propose revoke account $user_10_address"
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth propose-revoke-account --address="$user_10_address" --from=$trustee_account_1 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve revoke account $user_10_address"
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth approve-revoke-account --address="$user_10_address" --from=$trustee_account_2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve revoke account $user_10_address"
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth approve-revoke-account --address="$user_10_address" --from=$trustee_account_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve revoke account $user_10_address"
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth approve-revoke-account --address="$user_10_address" --from=$trustee_account_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Propose revoke account $user_11_address"
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth propose-revoke-account --address="$user_11_address" --from=$trustee_account_1 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

# VALIDATOR_NODE
result=$(docker exec "$container" /bin/sh -c "echo test1234  | dcld config broadcast-mode sync")

echo "Disable node"
result=$(docker exec "$container" /bin/sh -c "echo test1234  | dcld tx validator disable-node --from=$account --yes")
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Enable node"
result=$(docker exec "$container" /bin/sh -c "echo test1234  | dcld tx validator enable-node --from=$account --yes")
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve disable node"
result=$(echo $passphrase | $DCLD_BIN_NEW tx validator approve-disable-node --address=$validator_address --from=$trustee_account_2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve disable node"
result=$(echo $passphrase | $DCLD_BIN_NEW tx validator approve-disable-node --address=$validator_address --from=$trustee_account_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve disable node"
result=$(echo $passphrase | $DCLD_BIN_NEW tx validator approve-disable-node --address=$validator_address --from=$trustee_account_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Enable node"
result=$(docker exec "$container" /bin/sh -c "echo test1234  | dcld tx validator enable-node --from=$account --yes")
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Propose disable node"
result=$(echo $passphrase | $DCLD_BIN_NEW tx validator propose-disable-node --address=$validator_address --from=$trustee_account_1 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Verify that new data is not corrupted"

test_divider

# VENDORINFO

echo "Verify if VendorInfo Record for VID: $vid_for_1_4_4 is present or not"
result=$($DCLD_BIN_NEW query vendorinfo vendor --vid=$vid_for_1_4_4)
check_response "$result" "\"vendorID\": $vid_for_1_4_4"
check_response "$result" "\"companyLegalName\": \"$company_legal_name_for_1_4_4\""

echo "Verify if VendorInfo Record for VID: $vid_for_1_2 updated or not"
result=$($DCLD_BIN_NEW query vendorinfo vendor --vid=$vid_for_1_2)
check_response "$result" "\"vendorID\": $vid_for_1_2"
check_response "$result" "\"vendorName\": \"$vendor_name_for_1_2\""
check_response "$result" "\"companyPreferredName\": \"$company_preferred_name_for_1_4_4\""
check_response "$result" "\"vendorLandingPageURL\": \"$vendor_landing_page_url_for_1_4_4\""

echo "Request all vendor infos"
result=$($DCLD_BIN_NEW query vendorinfo all-vendors)
check_response "$result" "\"vendorID\": $vid_for_1_4_4"
check_response "$result" "\"companyLegalName\": \"$company_legal_name_for_1_4_4\""
check_response "$result" "\"vendorName\": \"$vendor_name_for_1_4_4\""

test_divider

# MODEL

echo "Get Model with VID: $vid_for_1_4_4 PID: $pid_1_for_1_4_4"
result=$($DCLD_BIN_NEW query model get-model --vid=$vid_for_1_4_4 --pid=$pid_1_for_1_4_4)
check_response "$result" "\"vid\": $vid_for_1_4_4"
check_response "$result" "\"pid\": $pid_1_for_1_4_4"
check_response "$result" "\"productLabel\": \"$product_label_for_1_4_4\""

echo "Get Model with VID: $vid_for_1_4_4 PID: $pid_2_for_1_4_4"
result=$($DCLD_BIN_NEW query model get-model --vid=$vid_for_1_4_4 --pid=$pid_2_for_1_4_4)
check_response "$result" "\"vid\": $vid_for_1_4_4"
check_response "$result" "\"pid\": $pid_2_for_1_4_4"
check_response "$result" "\"productLabel\": \"$product_label_for_1_4_4\""

echo "Check Model with VID: $vid_for_1_4_4 PID: $pid_2_for_1_4_4 updated"
result=$($DCLD_BIN_NEW query model get-model --vid=$vid --pid=$pid_2)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_2"
check_response "$result" "\"productLabel\": \"$product_label_for_1_4_4\""
check_response "$result" "\"partNumber\": \"$part_number_for_1_4_4\""

echo "Check Model version with VID: $vid_for_1_4_4 PID: $pid_2_for_1_4_4 updated"
result=$($DCLD_BIN_NEW query model model-version --vid=$vid --pid=$pid_2  --softwareVersion=$software_version)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_2"
check_response "$result" "\"minApplicableSoftwareVersion\": $min_applicable_software_version_for_1_4_4"
check_response "$result" "\"maxApplicableSoftwareVersion\": $max_applicable_software_version_for_1_4_4"

echo "Get all models"
result=$($DCLD_BIN_NEW query model all-models)
check_response "$result" "\"vid\": $vid_for_1_4_4"
check_response "$result" "\"pid\": $pid_1_for_1_4_4"
check_response "$result" "\"pid\": $pid_2_for_1_4_4"

echo "Get all model versions"
result=$($DCLD_BIN_NEW query model all-model-versions --vid=$vid_for_1_4_4 --pid=$pid_1_for_1_4_4)
check_response "$result" "\"vid\": $vid_for_1_4_4"
check_response "$result" "\"pid\": $pid_1_for_1_4_4"

echo "Get Vendor Models with VID: ${vid_for_1_4_4}"
result=$($DCLD_BIN_NEW query model vendor-models --vid=$vid_for_1_4_4)
check_response "$result" "\"pid\": $pid_1_for_1_4_4"
check_response "$result" "\"pid\": $pid_2_for_1_4_4"

echo "Get model version VID: $vid_for_1_4_4 PID: $pid_1_for_1_4_4"
result=$($DCLD_BIN_NEW query model model-version --vid=$vid_for_1_4_4 --pid=$pid_1_for_1_4_4 --softwareVersion=$software_version_for_1_4_4)
check_response "$result" "\"vid\": $vid_for_1_4_4"
check_response "$result" "\"pid\": $pid_1_for_1_4_4"
check_response "$result" "\"softwareVersion\": $software_version_for_1_4_4"

echo "Get model version VID: $vid_for_1_4_4 PID: $pid_2_for_1_4_4"
result=$($DCLD_BIN_NEW query model model-version --vid=$vid_for_1_4_4 --pid=$pid_2_for_1_4_4 --softwareVersion=$software_version_for_1_4_4)
check_response "$result" "\"vid\": $vid_for_1_4_4"
check_response "$result" "\"pid\": $pid_2_for_1_4_4"
check_response "$result" "\"softwareVersion\": $software_version_for_1_4_4"

test_divider

# COMPLIANCE

echo "Get certified model vid=$vid_for_1_4_4 pid=$pid_1_for_1_4_4"
result=$($DCLD_BIN_NEW query compliance certified-model --vid=$vid_for_1_4_4 --pid=$pid_1_for_1_4_4 --softwareVersion=$software_version_for_1_4_4 --certificationType=$certification_type_for_1_4_4)
check_response "$result" "\"value\": true"
check_response "$result" "\"vid\": $vid_for_1_4_4"
check_response "$result" "\"pid\": $pid_1_for_1_4_4"
check_response "$result" "\"softwareVersion\": $software_version_for_1_4_4"
check_response "$result" "\"certificationType\": \"$certification_type_for_1_4_4\""

echo "Get revoked Model with VID: $vid_for_1_4_4 PID: $pid_2_for_1_4_4"
result=$($DCLD_BIN_NEW query compliance revoked-model --vid=$vid_for_1_4_4 --pid=$pid_2_for_1_4_4 --softwareVersion=$software_version_for_1_4_4 --certificationType=$certification_type_for_1_4_4)
check_response "$result" "\"vid\": $vid_for_1_4_4"
check_response "$result" "\"pid\": $pid_2_for_1_4_4"

echo "Get certified model with VID: $vid_for_1_4_4 PID: $pid_1_for_1_4_4"
result=$($DCLD_BIN_NEW query compliance certified-model --vid=$vid_for_1_4_4 --pid=$pid_1_for_1_4_4 --softwareVersion=$software_version_for_1_4_4 --certificationType=$certification_type_for_1_4_4)
check_response "$result" "\"value\": true"
check_response "$result" "\"vid\": $vid_for_1_4_4"
check_response "$result" "\"pid\": $pid_1_for_1_4_4"

echo "Get provisional model with VID: $vid_for_1_4_4 PID: $pid_2_for_1_4_4"
result=$($DCLD_BIN_NEW query compliance provisional-model --vid=$vid_for_1_4_4 --pid=$pid_2_for_1_4_4 --softwareVersion=$software_version_for_1_4_4 --certificationType=$certification_type_for_1_4_4)
check_response "$result" "\"value\": false"
check_response "$result" "\"vid\": $vid_for_1_4_4"
check_response "$result" "\"pid\": $pid_2_for_1_4_4"

echo "Get compliance-info model with VID: $vid_for_1_4_4 PID: $pid_1_for_1_4_4"
result=$($DCLD_BIN_NEW query compliance compliance-info --vid=$vid_for_1_4_4 --pid=$pid_1_for_1_4_4 --softwareVersion=$software_version_for_1_4_4 --certificationType=$certification_type_for_1_4_4)
check_response "$result" "\"vid\": $vid_for_1_4_4"
check_response "$result" "\"pid\": $pid_1_for_1_4_4"
check_response "$result" "\"softwareVersion\": $software_version_for_1_4_4"
check_response "$result" "\"certificationType\": \"$certification_type_for_1_4_4\""

echo "Get compliance-info model with VID: $vid_for_1_4_4 PID: $pid_2_for_1_4_4"
result=$($DCLD_BIN_NEW query compliance compliance-info --vid=$vid_for_1_4_4 --pid=$pid_2_for_1_4_4 --softwareVersion=$software_version_for_1_4_4 --certificationType=$certification_type_for_1_4_4)
check_response "$result" "\"vid\": $vid_for_1_4_4"
check_response "$result" "\"pid\": $pid_2_for_1_4_4"
check_response "$result" "\"softwareVersion\": $software_version_for_1_4_4"
check_response "$result" "\"certificationType\": \"$certification_type_for_1_4_4\""

echo "Get device software compliance cDCertificateId=$cd_certificate_id_for_1_4_4"
result=$($DCLD_BIN_NEW query compliance device-software-compliance --cdCertificateId=$cd_certificate_id_for_1_4_4)
check_response "$result" "\"vid\": $vid_for_1_4_4"
check_response "$result" "\"pid\": $pid_1_for_1_4_4"

echo "Get all certified models"
result=$($DCLD_BIN_NEW query compliance all-certified-models)
check_response "$result" "\"vid\": $vid_for_1_4_4"
check_response "$result" "\"pid\": $pid_1_for_1_4_4"

echo "Get all provisional models"
result=$($DCLD_BIN_NEW query compliance all-provisional-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_3"

echo "Get all revoked models"
result=$($DCLD_BIN_NEW query compliance all-revoked-models)
check_response "$result" "\"vid\": $vid_for_1_4_4"
check_response "$result" "\"pid\": $pid_2_for_1_4_4"

echo "Get all compliance infos"
result=$($DCLD_BIN_NEW query compliance all-compliance-info)
check_response "$result" "\"vid\": $vid_for_1_4_4"
check_response "$result" "\"pid\": $pid_1_for_1_4_4"
check_response "$result" "\"pid\": $pid_2_for_1_4_4"

echo "Get all device software compliances"
result=$($DCLD_BIN_NEW query compliance all-device-software-compliance)
check_response "$result" "\"vid\": $vid_for_1_4_4"
check_response "$result" "\"pid\": $pid_1_for_1_4_4"
check_response "$result" "\"cDCertificateId\": \"$cd_certificate_id_for_1_4_4\""

test_divider

# PKI

echo "Get certificates"

echo "Get certificates (ALL)"
result=$($DCLD_BIN_NEW query pki all-certs)
echo $result | jq
check_response "$result" "\"subjectKeyId\": \"$da_root_cert_2_subject_key_id_for_1_4_4\""
check_response "$result" "\"subjectKeyId\": \"$da_intermediate_cert_2_subject_key_id_for_1_4_4\""
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_2_subject_key_id_for_1_4_4\""
check_response "$result" "\"subjectKeyId\": \"$noc_ica_cert_2_subject_key_id_for_1_4_4\""

echo "Get certificates (DA)"
result=$($DCLD_BIN_NEW query pki all-x509-certs)
check_response "$result" "\"subjectKeyId\": \"$da_root_cert_2_subject_key_id_for_1_4_4\""
check_response "$result" "\"subjectKeyId\": \"$da_intermediate_cert_2_subject_key_id_for_1_4_4\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_2_subject_key_id_for_1_4_4\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_ica_cert_2_subject_key_id_for_1_4_4\""

echo "Get certificates (NOC)"
result=$($DCLD_BIN_NEW query pki all-noc-x509-certs)
check_response "$result" "$noc_root_cert_2_subject_key_id_for_1_4_4"
check_response "$result" "$noc_ica_cert_2_subject_key_id_for_1_4_4"
response_does_not_contain "$result" "$da_root_cert_1_subject_key_id_for_1_4_4"
response_does_not_contain "$result" "$noc_root_cert_1_subject_key_id_for_1_4_4"
response_does_not_contain "$result" "$noc_ica_cert_1_subject_key_id_for_1_4_4"

echo "Get certificate"

echo "Get certificate (ALL)"
result=$($DCLD_BIN_NEW query pki cert --subject=$da_root_cert_2_subject_for_1_4_4 --subject-key-id=$da_root_cert_2_subject_key_id_for_1_4_4)
check_response "$result" "\"subject\": \"$da_root_cert_2_subject_for_1_4_4\""
check_response "$result" "\"subjectKeyId\": \"$da_root_cert_2_subject_key_id_for_1_4_4\""

result=$($DCLD_BIN_NEW query pki cert --subject=$da_intermediate_cert_2_subject_for_1_4_4 --subject-key-id=$da_intermediate_cert_2_subject_key_id_for_1_4_4)
check_response "$result" "\"subject\": \"$da_intermediate_cert_2_subject_for_1_4_4\""
check_response "$result" "\"subjectKeyId\": \"$da_intermediate_cert_2_subject_key_id_for_1_4_4\""

result=$($DCLD_BIN_NEW query pki cert --subject=$noc_root_cert_2_subject_for_1_4_4 --subject-key-id=$noc_root_cert_2_subject_key_id_for_1_4_4)
check_response "$result" "\"subject\": \"$noc_root_cert_2_subject_for_1_4_4\""
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_2_subject_key_id_for_1_4_4\""

result=$($DCLD_BIN_NEW query pki cert --subject=$noc_ica_cert_2_subject_for_1_4_4 --subject-key-id=$noc_ica_cert_2_subject_key_id_for_1_4_4)
check_response "$result" "\"subject\": \"$noc_ica_cert_2_subject_for_1_4_4\""
check_response "$result" "\"subjectKeyId\": \"$noc_ica_cert_2_subject_key_id_for_1_4_4\""

echo "Get certificate (DA)"
result=$($DCLD_BIN_NEW query pki x509-cert --subject=$da_root_cert_2_subject_for_1_4_4 --subject-key-id=$da_root_cert_2_subject_key_id_for_1_4_4)
check_response "$result" "\"subject\": \"$da_root_cert_2_subject_for_1_4_4\""
check_response "$result" "\"subjectKeyId\": \"$da_root_cert_2_subject_key_id_for_1_4_4\""

result=$($DCLD_BIN_NEW query pki x509-cert --subject=$da_intermediate_cert_2_subject_for_1_4_4 --subject-key-id=$da_intermediate_cert_2_subject_key_id_for_1_4_4)
check_response "$result" "\"subject\": \"$da_intermediate_cert_2_subject_for_1_4_4\""
check_response "$result" "\"subjectKeyId\": \"$da_intermediate_cert_2_subject_key_id_for_1_4_4\""

result=$($DCLD_BIN_NEW query pki x509-cert --subject=$noc_root_cert_2_subject_for_1_4_4 --subject-key-id=$noc_root_cert_2_subject_key_id_for_1_4_4)
check_response "$result" "Not Found"

result=$($DCLD_BIN_NEW query pki x509-cert --subject=$noc_ica_cert_2_subject_for_1_4_4 --subject-key-id=$noc_ica_cert_2_subject_key_id_for_1_4_4)
check_response "$result" "Not Found"

echo "Get certificate (NOC)"
result=$($DCLD_BIN_NEW query pki noc-x509-cert --subject=$noc_root_cert_2_subject_for_1_4_4 --subject-key-id=$noc_root_cert_2_subject_key_id_for_1_4_4)
check_response "$result" "\"subject\": \"$noc_root_cert_2_subject_for_1_4_4\""
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_2_subject_key_id_for_1_4_4\""

result=$($DCLD_BIN_NEW query pki noc-x509-cert --subject=$noc_ica_cert_2_subject_for_1_4_4 --subject-key-id=$noc_ica_cert_2_subject_key_id_for_1_4_4)
check_response "$result" "\"subject\": \"$noc_ica_cert_2_subject_for_1_4_4\""
check_response "$result" "\"subjectKeyId\": \"$noc_ica_cert_2_subject_key_id_for_1_4_4\""

result=$($DCLD_BIN_NEW query pki noc-x509-cert --subject=$da_root_cert_2_subject_for_1_4_4 --subject-key-id=$da_root_cert_2_subject_key_id_for_1_4_4)
check_response "$result" "Not Found"

result=$($DCLD_BIN_NEW query pki noc-x509-cert --subject=$da_intermediate_cert_2_subject_for_1_4_4 --subject-key-id=$da_intermediate_cert_2_subject_key_id_for_1_4_4)
check_response "$result" "Not Found"

echo "Get all subject certificates"

echo "Get all subject certificates (Global)"
result=$($DCLD_BIN_NEW query pki all-subject-certs --subject=$da_root_cert_2_subject_for_1_4_4)
check_response "$result" "$da_root_cert_2_subject_key_id_for_1_4_4"

result=$($DCLD_BIN_NEW query pki all-subject-certs --subject=$noc_root_cert_2_subject_for_1_4_4)
check_response "$result" "$noc_root_cert_2_subject_for_1_4_4"

echo "Get all subject certificates (DA)"
result=$($DCLD_BIN_NEW query pki all-subject-x509-certs --subject=$da_root_cert_2_subject_for_1_4_4)
check_response "$result" "$da_root_cert_2_subject_key_id_for_1_4_4"

result=$($DCLD_BIN_NEW query pki all-subject-x509-certs --subject=$noc_root_cert_2_subject_for_1_4_4)
check_response "$result" "Not Found"

echo "Get all subject certificates (NOC)"
result=$($DCLD_BIN_NEW query pki all-noc-subject-x509-certs --subject=$noc_root_cert_2_subject_for_1_4_4)
check_response "$result" "$noc_root_cert_2_subject_for_1_4_4"

result=$($DCLD_BIN_NEW query pki all-noc-subject-x509-certs --subject=$da_root_cert_2_subject_for_1_4_4)
check_response "$result" "Not Found"

echo "Get all certificates by SKID"

echo "Get all certificates by SKID (Global)"
result=$($DCLD_BIN_NEW query pki cert --subject-key-id=$da_root_cert_2_subject_key_id_for_1_4_4)
check_response "$result" "\"subjectKeyId\": \"$da_root_cert_2_subject_key_id_for_1_4_4\""

result=$($DCLD_BIN_NEW query pki cert --subject-key-id=$noc_root_cert_2_subject_key_id_for_1_4_4)
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_2_subject_key_id_for_1_4_4\""

echo "Get all certificates by SKID (DA)"
result=$($DCLD_BIN_NEW query pki x509-cert --subject-key-id=$da_root_cert_2_subject_key_id_for_1_4_4)
check_response "$result" "\"subjectKeyId\": \"$da_root_cert_2_subject_key_id_for_1_4_4\""

result=$($DCLD_BIN_NEW query pki x509-cert --subject-key-id=$noc_root_cert_2_subject_key_id_for_1_4_4)
check_response "$result" "Not Found"

echo "Get all certificates by SKID (NOC)"
result=$($DCLD_BIN_NEW query pki noc-x509-cert --subject-key-id=$noc_root_cert_2_subject_key_id_for_1_4_4)
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_2_subject_key_id_for_1_4_4\""

result=$($DCLD_BIN_NEW query pki noc-x509-cert --subject-key-id=$da_root_cert_2_subject_key_id_for_1_4_4)
check_response "$result" "Not Found"

echo "Get all revoked x509 root certificates"

echo "Get all revoked x509 certificates (DA)"
result=$($DCLD_BIN_NEW query pki all-revoked-x509-certs)
check_response "$result" "\"subjectKeyId\": \"$da_root_cert_1_subject_key_id_for_1_4_4\""
check_response "$result" "\"subjectKeyId\": \"$da_intermediate_cert_1_subject_key_id_for_1_4_4\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id_for_1_4_4\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_ica_cert_1_subject_key_id_for_1_4_4\""

echo "Get all revoked x509 root certificates (DA)"
result=$($DCLD_BIN_NEW query pki all-revoked-x509-root-certs)
check_response "$result" "\"subjectKeyId\": \"$da_root_cert_1_subject_key_id_for_1_4_4\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$da_root_cert_2_subject_for_1_4_4\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id_for_1_4_4\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id_for_1_4_4\""

echo "Get all revoked x509 root certificates (NOC)"
result=$($DCLD_BIN_NEW query pki all-revoked-noc-x509-root-certs)
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id_for_1_4_4\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$da_root_cert_1_subject_key_id_for_1_4_4\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_ica_cert_1_subject_key_id_for_1_4_4\""

echo "Get all revoked x509 ica certificates (NOC)"
result=$($DCLD_BIN_NEW query pki all-revoked-noc-x509-ica-certs)
check_response "$result" "\"subjectKeyId\": \"$noc_ica_cert_1_subject_key_id_for_1_4_4\""

echo "Get revoked x509 certificate"

echo "Get revoked x509 certificate (DA)"
result=$($DCLD_BIN_NEW query pki revoked-x509-cert --subject=$da_root_cert_1_subject_for_1_4_4 --subject-key-id=$da_root_cert_1_subject_key_id_for_1_4_4)
check_response "$result" "\"subject\": \"$da_root_cert_1_subject_for_1_4_4\""
check_response "$result" "\"subjectKeyId\": \"$da_root_cert_1_subject_key_id_for_1_4_4\""

result=$($DCLD_BIN_NEW query pki revoked-x509-cert --subject=$noc_root_cert_1_subject_for_1_4_4 --subject-key-id=$noc_root_cert_1_subject_key_id_for_1_4_4)
check_response "$result" "Not Found"

echo "Get revoked x509 certificate (NOC)"
result=$($DCLD_BIN_NEW query pki revoked-noc-x509-root-cert --subject=$noc_root_cert_1_subject_for_1_4_4 --subject-key-id=$noc_root_cert_1_subject_key_id_for_1_4_4)
check_response "$result" "\"subject\": \"$noc_root_cert_1_subject_for_1_4_4\""
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id_for_1_4_4\""

result=$($DCLD_BIN_NEW query pki revoked-noc-x509-root-cert --subject=$da_root_cert_1_subject_for_1_4_4 --subject-key-id=$da_root_cert_1_subject_key_id_for_1_4_4)
check_response "$result" "Not Found"

echo "Get revocation point"
result=$($DCLD_BIN_NEW query pki revocation-point --vid=$vid_for_1_4_4 --label=$product_label_for_1_4_4 --issuer-subject-key-id=$issuer_subject_key_id)
check_response "$result" "\"vid\": $vid_for_1_4_4"
check_response "$result" "\"issuerSubjectKeyID\": \"$issuer_subject_key_id\""
check_response "$result" "\"label\": \"$product_label_for_1_4_4\""
check_response "$result" "\"dataURL\": \"$test_data_url_for_1_4_4\""

echo "Get revocation points by issuer subject key id"
result=$($DCLD_BIN_NEW query pki revocation-points --issuer-subject-key-id=$issuer_subject_key_id)
check_response "$result" "\"vid\": $vid_for_1_4_4"
check_response "$result" "\"issuerSubjectKeyID\": \"$issuer_subject_key_id\""
check_response "$result" "\"label\": \"$product_label_for_1_4_4\""
check_response "$result" "\"dataURL\": \"$test_data_url_for_1_4_4\""

echo "Get all revocation points"
result=$($DCLD_BIN_NEW query pki all-revocation-points)
check_response "$result" "\"vid\": $vid_for_1_4_4"
check_response "$result" "\"issuerSubjectKeyID\": \"$issuer_subject_key_id\""
check_response "$result" "\"label\": \"$product_label_for_1_4_4\""
check_response "$result" "\"dataURL\": \"$test_data_url_for_1_4_4\""

echo "Get all noc x509 root certificates by vid=$vid_for_1_4_4 and skid=$noc_root_cert_2_subject_key_id_for_1_4_4"
result=$($DCLD_BIN_NEW query pki noc-x509-cert --vid=$vid_for_1_4_4 --subject-key-id=$noc_root_cert_2_subject_key_id_for_1_4_4)
check_response "$result" "\"subject\": \"$noc_root_cert_2_subject_for_1_4_4\""
check_response "$result" "$noc_root_cert_2_subject_key_id_for_1_4_4"

echo "Get all noc x509 root certificates by vid $vid_for_1_4_4 and skid=$noc_ica_cert_2_subject_key_id_for_1_4_4"
result=$($DCLD_BIN_NEW query pki noc-x509-cert --vid=$vid_for_1_4_4 --subject-key-id=$noc_ica_cert_2_subject_key_id_for_1_4_4)
check_response "$result" "\"subject\": \"$noc_ica_cert_2_subject_for_1_4_4\""
check_response "$result" "$noc_ica_cert_2_subject_key_id_for_1_4_4"

test_divider

# AUTH

echo "Get all accounts"
result=$($DCLD_BIN_NEW query auth all-accounts)
check_response "$result" "\"address\": \"$user_11_address\""

echo "Get account"
result=$($DCLD_BIN_NEW query auth account --address=$user_11_address)
check_response "$result" "\"address\": \"$user_11_address\""

echo "Get all proposed accounts"
result=$($DCLD_BIN_NEW query auth all-proposed-accounts)
check_response "$result" "\"address\": \"$user_12_address\""

echo "Get proposed account"
result=$($DCLD_BIN_NEW query auth proposed-account --address=$user_12_address)
check_response "$result" "\"address\": \"$user_12_address\""

echo "Get all proposed accounts to revoke"
result=$($DCLD_BIN_NEW query auth all-proposed-accounts-to-revoke)
check_response "$result" "\"address\": \"$user_11_address\""

echo "Get proposed account to revoke"
result=$($DCLD_BIN_NEW query auth proposed-account-to-revoke --address=$user_11_address)
check_response "$result" "\"address\": \"$user_11_address\""

echo "Get all revoked accounts"
result=$($DCLD_BIN_NEW query auth all-revoked-accounts)
check_response "$result" "\"address\": \"$user_10_address\""

echo "Get revoked account"
result=$($DCLD_BIN_NEW query auth revoked-account --address=$user_10_address)
check_response "$result" "\"address\": \"$user_10_address\""

test_divider

# Validator

echo "Get node"
# FIXME: use proper binary (not dcld but $DCLD_BIN_OLD)
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator all-nodes")
check_response "$result" "\"owner\": \"$validator_address\""

test_divider

echo "Upgrade from 1.4.3 to 1.4.4 passed"

rm -f $DCLD_BIN_OLD
rm -f $DCLD_BIN_NEW
