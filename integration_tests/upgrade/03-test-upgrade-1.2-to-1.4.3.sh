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

plan_name="v1.4"
upgrade_checksum="sha256:a007f58d61632af107a09c89b7392eedd05d8127d0df67ace50f318948c62001"
binary_version_new="v1.4.3"

DCLD_BIN_OLD="/tmp/dcld_bins/dcld_v1.2.2"
DCLD_BIN_NEW="/tmp/dcld_bins/dcld_v1.4.3"
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

echo "Request all vendor infos"
result=$($DCLD_BIN_NEW query vendorinfo all-vendors)
check_response "$result" "\"vendorID\": $vid"
check_response "$result" "\"vendorID\": $vid_for_1_2"
check_response "$result" "\"companyLegalName\": \"$company_legal_name\""
check_response "$result" "\"companyLegalName\": \"$company_legal_name_for_1_2\""
check_response "$result" "\"vendorName\": \"$vendor_name\""
check_response "$result" "\"vendorName\": \"$vendor_name_for_1_2\""

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
check_response "$result" "\"productLabel\": \"$product_label_for_1_2\""
check_response "$result" "\"partNumber\": \"$part_number_for_1_2\""

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

echo "Get all models"
result=$($DCLD_BIN_NEW query model all-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_1"
check_response "$result" "\"pid\": $pid_2"
check_response "$result" "\"vid\": $vid_for_1_2"
check_response "$result" "\"pid\": $pid_1_for_1_2"
check_response "$result" "\"pid\": $pid_2_for_1_2"

echo "Get Vendor Models with VID: ${vid}"
result=$($DCLD_BIN_NEW query model vendor-models --vid=$vid)
check_response "$result" "\"pid\": $pid_1"
check_response "$result" "\"pid\": $pid_2"

echo "Get Vendor Models with VID: ${vid_for_1_2}"
result=$($DCLD_BIN_NEW query model vendor-models --vid=$vid_for_1_2)
check_response "$result" "\"pid\": $pid_1_for_1_2"
check_response "$result" "\"pid\": $pid_2_for_1_2"

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
check_response "$result" "\"minApplicableSoftwareVersion\": $min_applicable_software_version_for_1_2"
check_response "$result" "\"maxApplicableSoftwareVersion\": $max_applicable_software_version_for_1_2"

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

echo "Get revoked Model with VID: $vid PID: $pid_2"
result=$($DCLD_BIN_NEW query compliance revoked-model --vid=$vid --pid=$pid_2 --softwareVersion=$software_version --certificationType=$certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_2"

echo "Get revoked Model with VID: $vid_for_1_2 PID: $pid_2_for_1_2"
result=$($DCLD_BIN_NEW query compliance revoked-model --vid=$vid_for_1_2 --pid=$pid_2_for_1_2 --softwareVersion=$software_version_for_1_2 --certificationType=$certification_type_for_1_2)
check_response "$result" "\"vid\": $vid_for_1_2"
check_response "$result" "\"pid\": $pid_2_for_1_2"

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

echo "Get device software compliance cDCertificateId=$cd_certificate_id"
result=$($DCLD_BIN_NEW query compliance device-software-compliance --cdCertificateId=$cd_certificate_id)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_1"

echo "Get device software compliance cDCertificateId=$cd_certificate_id_for_1_2"
result=$($DCLD_BIN_NEW query compliance device-software-compliance --cdCertificateId=$cd_certificate_id_for_1_2)
check_response "$result" "\"vid\": $vid_for_1_2"
check_response "$result" "\"pid\": $pid_1_for_1_2"

echo "Get all certified models"
result=$($DCLD_BIN_NEW query compliance all-certified-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_1"
check_response "$result" "\"vid\": $vid_for_1_2"
check_response "$result" "\"pid\": $pid_1_for_1_2"

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

echo "Get all compliance infos"
result=$($DCLD_BIN_NEW query compliance all-compliance-info)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_1"
check_response "$result" "\"pid\": $pid_2"
check_response "$result" "\"vid\": $vid_for_1_2"
check_response "$result" "\"pid\": $pid_1_for_1_2"
check_response "$result" "\"pid\": $pid_2_for_1_2"

echo "Get all device software compliances"
result=$($DCLD_BIN_NEW query compliance all-device-software-compliance)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_1"
check_response "$result" "\"cDCertificateId\": \"$cd_certificate_id\""
check_response "$result" "\"vid\": $vid_for_1_2"
check_response "$result" "\"pid\": $pid_1_for_1_2"
check_response "$result" "\"cDCertificateId\": \"$cd_certificate_id_for_1_2\""

test_divider

# PKI

echo "Get x509 root certificate"
result=$($DCLD_BIN_NEW query pki x509-cert --subject=$test_root_cert_subject_for_1_2 --subject-key-id=$test_root_cert_subject_key_id_for_1_2)
check_response "$result" "\"subject\": \"$test_root_cert_subject_for_1_2\""
check_response "$result" "\"subjectKeyId\": \"$test_root_cert_subject_key_id_for_1_2\""
check_response "$result" "\"vid\": $test_root_cert_vid_for_1_2"

result=$($DCLD_BIN_NEW query pki x509-cert --subject=$test_root_cert_subject --subject-key-id=$test_root_cert_subject_key_id)
check_response "$result" "\"subject\": \"$test_root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$test_root_cert_subject_key_id\""
check_response "$result" "\"vid\": $test_root_cert_vid"

echo "Get all x509 certificates by subjectKeyId $test_root_cert_subject_key_id"
result=$($DCLD_BIN_NEW query pki x509-cert --subject-key-id="$test_root_cert_subject_key_id_for_1_2")
check_response "$result" "$test_root_cert_subject_for_1_2"
check_response "$result" "\"subjectKeyId\": \"$test_root_cert_subject_key_id_for_1_2\""

result=$($DCLD_BIN_NEW query pki x509-cert --subject-key-id="$test_root_cert_subject_key_id")
check_response "$result" "$test_root_cert_subject"
check_response "$result" "\"subjectKeyId\": \"$test_root_cert_subject_key_id\""

echo "Get all subject x509 root certificates"
result=$($DCLD_BIN_NEW query pki all-subject-x509-certs --subject=$test_root_cert_subject_for_1_2)
check_response "$result" "\"subject\": \"$test_root_cert_subject_for_1_2\""
check_response "$result" "$test_root_cert_subject_key_id_for_1_2"

result=$($DCLD_BIN_NEW query pki all-subject-x509-certs --subject=$test_root_cert_subject)
check_response "$result" "\"subject\": \"$test_root_cert_subject\""
check_response "$result" "$test_root_cert_subject_key_id"

echo "Get proposed x509 root certificate"
result=$($DCLD_BIN_NEW query pki proposed-x509-root-cert --subject=$google_root_cert_subject_for_1_2 --subject-key-id=$google_root_cert_subject_key_id_for_1_2)
check_response "$result" "\"subject\": \"$google_root_cert_subject_for_1_2\""
check_response "$result" "\"subjectKeyId\": \"$google_root_cert_subject_key_id_for_1_2\""
check_response "$result" "\"vid\": $google_root_cert_path_random_vid_for_1_2"

result=$($DCLD_BIN_NEW query pki proposed-x509-root-cert --subject=$google_root_cert_subject --subject-key-id=$google_root_cert_subject_key_id)
check_response "$result" "\"subject\": \"$google_root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$google_root_cert_subject_key_id\""

echo "Get revoked x509 certificate"
result=$($DCLD_BIN_NEW query pki revoked-x509-cert --subject=$intermediate_cert_subject_for_1_2 --subject-key-id=$intermediate_cert_subject_key_id_for_1_2)
check_response "$result" "\"subject\": \"$intermediate_cert_subject_for_1_2\""
check_response "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id_for_1_2\""

result=$($DCLD_BIN_NEW query pki revoked-x509-cert --subject=$intermediate_cert_subject --subject-key-id=$intermediate_cert_subject_key_id)
check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""

echo "Get proposed x509 root certificate to revoke"
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

echo "Get all x509 certificates"
result=$($DCLD_BIN_NEW query pki all-x509-certs)
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

# after upgrade constatnts

vid_for_1_4_3=65521
pid_1_for_1_4_3=44
pid_2_for_1_4_3=55
pid_3_for_1_4_3=66
device_type_id_for_1_4_3=4321
product_name_for_1_4_3="ProductName13"
product_label_for_1_4_3="ProductLabel13"
part_number_for_1_4_3="RCU2225B"
software_version_for_1_4_3=2
software_version_string_for_1_4_3="3.0"
cd_version_number_for_1_4_3=413
min_applicable_software_version_for_1_4_3=3
max_applicable_software_version_for_1_4_3=3000

certification_type_for_1_4_3="matter"
certification_date_for_1_4_3="2022-01-01T00:00:00Z"
provisional_date_for_1_4_3="2012-12-12T00:00:00Z"
cd_certificate_id_for_1_4_3="18DEXC"

root_cert_with_vid_subject_for_1_4_3="MIGYMQswCQYDVQQGEwJVUzERMA8GA1UECAwITmV3IFlvcmsxETAPBgNVBAcMCE5ldyBZb3JrMRgwFgYDVQQKDA9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsMEFRlc3RpbmcgRGl2aXNpb24xGDAWBgNVBAMMD3d3dy5leGFtcGxlLmNvbTEUMBIGCisGAQQBgqJ8AgEMBEZGRjE="
root_cert_with_vid_subject_key_id_for_1_4_3="CE:A8:92:66:EA:E0:80:BD:2B:B5:68:E4:0B:07:C4:FA:2C:34:6D:31"
root_cert_with_vid_path_for_1_4_3="integration_tests/constants/root_cert_with_vid"
root_cert_vid_for_1_4_3=65521

paa_cert_no_vid_path_for_1_4_3="integration_tests/constants/paa_cert_no_vid"
paa_cert_no_vid_subject_for_1_4_3="MBoxGDAWBgNVBAMMD01hdHRlciBUZXN0IFBBQQ=="
paa_cert_no_vid_subject_key_id_for_1_4_3="78:5C:E7:05:B8:6B:8F:4E:6F:C7:93:AA:60:CB:43:EA:69:68:82:D5"

root_cert_subject_for_1_4_3="MIGCMQswCQYDVQQGEwJVUzERMA8GA1UECAwITmV3IFlvcmsxETAPBgNVBAcMCE5ldyBZb3JrMRgwFgYDVQQKDA9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsMEFRlc3RpbmcgRGl2aXNpb24xGDAWBgNVBAMMD3d3dy5leGFtcGxlLmNvbQ=="
root_cert_subject_key_id_for_1_4_3="33:5E:0C:07:44:F8:B5:9C:CD:55:01:9B:6D:71:23:83:6F:D0:D4:BE"
root_cert_path_for_1_4_3="integration_tests/constants/root_with_same_subject_and_skid_1"

intermediate_cert_with_vid_subject_for_1_4_3="MIGuMQswCQYDVQQGEwJVUzERMA8GA1UECAwITmV3IFlvcmsxETAPBgNVBAcMCE5ldyBZb3JrMRgwFgYDVQQKDA9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsMEFRlc3RpbmcgRGl2aXNpb24xGDAWBgNVBAMMD3d3dy5leGFtcGxlLmNvbTEUMBIGCisGAQQBgqJ8AgEMBEZGRjExFDASBgorBgEEAYKifAICDARGRkYx"
intermediate_cert_with_vid_subject_key_id_for_1_4_3="0E:8C:E8:C8:B8:AA:50:BC:25:85:56:B9:B1:9C:C2:C7:D9:C5:2F:17"
intermediate_cert_with_vid_path_for_1_4_3="integration_tests/constants/intermediate_cert_with_vid_1"
intermediate_cert_with_vid_serial_number_for_1_4_3="3"
intermediate_cert_with_vid_65521_vid_for_1_4_3=65521

noc_root_cert_1_path_for_1_4_3="integration_tests/constants/noc_root_cert_1"
noc_root_cert_1_subject_for_1_4_3="MHoxCzAJBgNVBAYTAlVaMRMwEQYDVQQIDApTb21lIFN0YXRlMREwDwYDVQQHDAhUYXNoa2VudDEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMQ4wDAYDVQQDDAVOT0MtMQ=="
noc_root_cert_1_subject_key_id_for_1_4_3="44:EB:4C:62:6B:25:48:CD:A2:B3:1C:87:41:5A:08:E7:2B:B9:83:26"

noc_ica_cert_1_path_for_1_4_3="integration_tests/constants/noc_cert_1"
noc_ica_cert_1_subject_for_1_4_3="MIGCMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTETMBEGA1UEBwwKU29tZSBTdGF0ZTEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRQwEgYDVQQDDAtOT0MtY2hpbGQtMQ=="
noc_ica_cert_1_subject_key_id_for_1_4_3="02:72:6E:BC:BB:EF:D6:BD:8D:9B:42:AE:D4:3C:C0:55:5F:66:3A:B3"

crl_signer_delegated_by_pai_1="integration_tests/constants/leaf_cert_with_vid_65521"
delegator_cert_with_vid_65521_path="integration_tests/constants/intermediate_cert_with_vid_1"
delegator_cert_with_vid_subject_key_id="0E8CE8C8B8AA50BC258556B9B19CC2C7D9C52F17"

test_data_url_for_1_4_3="https://url.data.dclmodel-1.4"

vendor_name_for_1_4_3="Vendor65521"
company_legal_name_for_1_4_3="LegalCompanyName65521"
company_preferred_name_for_1_4_3="CompanyPreferredName65521"
vendor_landing_page_url_for_1_4_3="https://www.new65521example.com"

vendor_account_for_1_4_3="vendor_account_65521"

echo "Create Vendor account $vendor_account_for_1_4_3"

result="$(echo $passphrase | $DCLD_BIN_NEW keys add "$vendor_account_for_1_4_3")"
echo "keys add $result"
_address=$(echo $passphrase | $DCLD_BIN_NEW keys show $vendor_account_for_1_4_3 -a)
_pubkey=$(echo $passphrase | $DCLD_BIN_NEW keys show $vendor_account_for_1_4_3 -p)
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth propose-add-account --address="$_address" --pubkey="$_pubkey" --vid="$vid_for_1_4_3" --roles="Vendor" --from "$trustee_account_1" --yes)
echo "propose-add-account $result"
result=$(get_txn_result "$result")
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth approve-add-account --address="$_address" --from "$trustee_account_2" --yes)
result=$(get_txn_result "$result")
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth approve-add-account --address="$_address" --from "$trustee_account_3" --yes)
result=$(get_txn_result "$result")
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth approve-add-account --address="$_address" --from "$trustee_account_4" --yes)
result=$(get_txn_result "$result")

random_string user_7
echo "$user_7 generates keys"
cmd="(echo $passphrase; echo $passphrase) | $DCLD_BIN_NEW keys add $user_7"
result="$(bash -c "$cmd")"
user_7_address=$(echo $passphrase | $DCLD_BIN_NEW keys show $user_7 -a)
user_7_pubkey=$(echo $passphrase | $DCLD_BIN_NEW keys show $user_7 -p)

random_string user_8
echo "$user_8 generates keys"
cmd="(echo $passphrase; echo $passphrase) | $DCLD_BIN_NEW keys add $user_8"
result="$(bash -c "$cmd")"
user_8_address=$(echo $passphrase | $DCLD_BIN_NEW keys show $user_8 -a)
user_8_pubkey=$(echo $passphrase | $DCLD_BIN_NEW keys show $user_8 -p)

random_string user_9
echo "$user_9 generates keys"
cmd="(echo $passphrase; echo $passphrase) | $DCLD_BIN_NEW keys add $user_9"
result="$(bash -c "$cmd")"
user_9_address=$(echo $passphrase | $DCLD_BIN_NEW keys show $user_9 -a)
user_9_pubkey=$(echo $passphrase | $DCLD_BIN_NEW keys show $user_9 -p)

# send all ledger update transactions after upgrade

# VENDOR_INFO
echo "Add vendor $vendor_name_for_1_4_3"
result=$(echo $passphrase | $DCLD_BIN_NEW tx vendorinfo add-vendor --vid=$vid_for_1_4_3 --vendorName=$vendor_name_for_1_4_3 --companyLegalName=$company_legal_name_for_1_4_3 --companyPreferredName=$company_preferred_name_for_1_4_3 --vendorLandingPageURL=$vendor_landing_page_url_for_1_4_3 --from=$vendor_account_for_1_4_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Update vendor $vendor_name_for_1_2"
result=$(echo $passphrase | $DCLD_BIN_NEW tx vendorinfo update-vendor --vid=$vid_for_1_2 --vendorName=$vendor_name_for_1_2 --companyLegalName=$company_legal_name_for_1_2 --companyPreferredName=$company_preferred_name_for_1_4_3 --vendorLandingPageURL=$vendor_landing_page_url_for_1_4_3 --from=$vendor_account_for_1_2 --yes)
result=$(get_txn_result "$result")
echo $result
check_response "$result" "\"code\": 0"

test_divider

# MODEL and MODEL_VERSION

echo "Add model vid=$vid_for_1_4_3 pid=$pid_1_for_1_4_3"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model --vid=$vid_for_1_4_3 --pid=$pid_1_for_1_4_3 --deviceTypeID=$device_type_id_for_1_4_3 --productName=$product_name_for_1_4_3 --productLabel=$product_label_for_1_4_3 --partNumber=$part_number_for_1_4_3 --from=$vendor_account_for_1_4_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Add model version vid=$vid_for_1_4_3 pid=$pid_1_for_1_4_3"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model-version --vid=$vid_for_1_4_3 --pid=$pid_1_for_1_4_3 --softwareVersion=$software_version_for_1_4_3 --softwareVersionString=$software_version_string_for_1_4_3 --cdVersionNumber=$cd_version_number_for_1_4_3 --minApplicableSoftwareVersion=$min_applicable_software_version_for_1_4_3 --maxApplicableSoftwareVersion=$max_applicable_software_version_for_1_4_3 --from=$vendor_account_for_1_4_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Add model vid=$vid_for_1_4_3 pid=$pid_2_for_1_4_3"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model --vid=$vid_for_1_4_3 --pid=$pid_2_for_1_4_3 --deviceTypeID=$device_type_id_for_1_4_3 --productName=$product_name_for_1_4_3 --productLabel=$product_label_for_1_4_3 --partNumber=$part_number_for_1_4_3 --from=$vendor_account_for_1_4_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Add model version vid=$vid_for_1_4_3 pid=$pid_2_for_1_4_3"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model-version --vid=$vid_for_1_4_3 --pid=$pid_2_for_1_4_3 --softwareVersion=$software_version_for_1_4_3 --softwareVersionString=$software_version_string_for_1_4_3 --cdVersionNumber=$cd_version_number_for_1_4_3 --minApplicableSoftwareVersion=$min_applicable_software_version_for_1_4_3 --maxApplicableSoftwareVersion=$max_applicable_software_version_for_1_4_3 --from=$vendor_account_for_1_4_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Add model vid=$vid_for_1_4_3 pid=$pid_3_for_1_4_3"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model --vid=$vid_for_1_4_3 --pid=$pid_3_for_1_4_3 --deviceTypeID=$device_type_id_for_1_4_3 --productName=$product_name_for_1_4_3 --productLabel=$product_label_for_1_4_3 --partNumber=$part_number_for_1_4_3 --from=$vendor_account_for_1_4_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Add model version vid=$vid_for_1_4_3 pid=$pid_3_for_1_4_3"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model-version --vid=$vid_for_1_4_3 --pid=$pid_3_for_1_4_3 --softwareVersion=$software_version_for_1_4_3 --softwareVersionString=$software_version_string_for_1_4_3 --cdVersionNumber=$cd_version_number_for_1_4_3 --minApplicableSoftwareVersion=$min_applicable_software_version_for_1_4_3 --maxApplicableSoftwareVersion=$max_applicable_software_version_for_1_4_3 --from=$vendor_account_for_1_4_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Delete model vid=$vid_for_1_4_3 pid=$pid_3_for_1_4_3"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model delete-model --vid=$vid_for_1_4_3 --pid=$pid_3_for_1_4_3 --from=$vendor_account_for_1_4_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Update model vid=$vid pid=$pid_2"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model update-model --vid=$vid --pid=$pid_2 --productName=$product_name --productLabel=$product_label_for_1_4_3 --partNumber=$part_number_for_1_4_3 --from=$vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Update model version vid=$vid pid=$pid_2"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model update-model-version --vid=$vid --pid=$pid_2 --softwareVersion=$software_version --minApplicableSoftwareVersion=$min_applicable_software_version_for_1_4_3 --maxApplicableSoftwareVersion=$max_applicable_software_version_for_1_4_3 --from=$vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

# CERTIFY_DEVICE_COMPLIANCE

echo "Certify model vid=$vid_for_1_4_3 pid=$pid_1_for_1_4_3"
result=$(echo $passphrase | $DCLD_BIN_NEW tx compliance certify-model --vid=$vid_for_1_4_3 --pid=$pid_1_for_1_4_3 --softwareVersion=$software_version_for_1_4_3 --softwareVersionString=$software_version_string_for_1_4_3  --certificationType=$certification_type_for_1_4_3 --certificationDate=$certification_date_for_1_4_3 --cdCertificateId=$cd_certificate_id_for_1_4_3 --from=$certification_center_account --cdVersionNumber=$cd_version_number_for_1_4_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Provision model vid=$vid_for_1_4_3 pid=$pid_2_for_1_4_3"
result=$(echo $passphrase | $DCLD_BIN_NEW tx compliance provision-model --vid=$vid_for_1_4_3 --pid=$pid_2_for_1_4_3 --softwareVersion=$software_version_for_1_4_3 --softwareVersionString=$software_version_string_for_1_4_3 --certificationType=$certification_type_for_1_4_3 --provisionalDate=$provisional_date_for_1_4_3 --cdCertificateId=$cd_certificate_id_for_1_4_3 --from=$certification_center_account --cdVersionNumber=$cd_version_number_for_1_4_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Certify model vid=$vid_for_1_4_3 pid=$pid_2_for_1_4_3"
result=$(echo $passphrase | $DCLD_BIN_NEW tx compliance certify-model --vid=$vid_for_1_4_3 --pid=$pid_2_for_1_4_3 --softwareVersion=$software_version_for_1_4_3 --softwareVersionString=$software_version_string_for_1_4_3  --certificationType=$certification_type_for_1_4_3 --certificationDate=$certification_date_for_1_4_3 --cdCertificateId=$cd_certificate_id_for_1_4_3 --from=$certification_center_account --cdVersionNumber=$cd_version_number_for_1_4_3  --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Revoke model certification vid=$vid_for_1_4_3 pid=$pid_2_for_1_4_3"
result=$(echo $passphrase | $DCLD_BIN_NEW tx compliance revoke-model --vid=$vid_for_1_4_3 --pid=$pid_2_for_1_4_3 --softwareVersion=$software_version_for_1_4_3 --softwareVersionString=$software_version_string_for_1_4_3 --certificationType=$certification_type_for_1_4_3 --revocationDate=$certification_date_for_1_4_3 --from=$certification_center_account --cdVersionNumber=$cd_version_number_for_1_4_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

# X509 PKI

echo "Propose add root_cert_with_vid"
result=$(echo $passphrase | $DCLD_BIN_NEW tx pki propose-add-x509-root-cert --certificate="$root_cert_with_vid_path_for_1_4_3" --vid="$vid_for_1_4_3" --from=$trustee_account_1 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add root_cert_with_vid"
result=$(echo $passphrase | $DCLD_BIN_NEW tx pki approve-add-x509-root-cert --subject="$root_cert_with_vid_subject_for_1_4_3" --subject-key-id=$root_cert_with_vid_subject_key_id_for_1_4_3 --from=$trustee_account_2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "reject add root_cert_with_vid"
result=$(echo $passphrase | $DCLD_BIN_NEW tx pki reject-add-x509-root-cert --subject="$root_cert_with_vid_subject_for_1_4_3" --subject-key-id=$root_cert_with_vid_subject_key_id_for_1_4_3 --from=$trustee_account_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add root_cert_with_vid"
result=$(echo $passphrase | $DCLD_BIN_NEW tx pki approve-add-x509-root-cert --subject="$root_cert_with_vid_subject_for_1_4_3" --subject-key-id=$root_cert_with_vid_subject_key_id_for_1_4_3 --from=$trustee_account_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add root_cert_with_vid"
result=$(echo $passphrase | $DCLD_BIN_NEW tx pki approve-add-x509-root-cert --subject="$root_cert_with_vid_subject_for_1_4_3" --subject-key-id=$root_cert_with_vid_subject_key_id_for_1_4_3 --from=$trustee_account_5 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Propose add paa_cert_no_vid"
result=$(echo $passphrase | $DCLD_BIN_NEW tx pki propose-add-x509-root-cert --certificate="$paa_cert_no_vid_path_for_1_4_3" --vid="$vid_for_1_4_3" --from=$trustee_account_5 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add paa_cert_no_vid"
result=$(echo $passphrase | $DCLD_BIN_NEW tx pki approve-add-x509-root-cert --subject="$paa_cert_no_vid_subject_for_1_4_3" --subject-key-id=$paa_cert_no_vid_subject_key_id_for_1_4_3 --from=$trustee_account_1 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add paa_cert_no_vid"
result=$(echo $passphrase | $DCLD_BIN_NEW tx pki approve-add-x509-root-cert --subject="$paa_cert_no_vid_subject_for_1_4_3" --subject-key-id=$paa_cert_no_vid_subject_key_id_for_1_4_3 --from=$trustee_account_2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add paa_cert_no_vid"
result=$(echo $passphrase | $DCLD_BIN_NEW tx pki approve-add-x509-root-cert --subject="$paa_cert_no_vid_subject_for_1_4_3" --subject-key-id=$paa_cert_no_vid_subject_key_id_for_1_4_3 --from=$trustee_account_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Propose root_cert"
result=$(echo $passphrase | $DCLD_BIN_NEW tx pki propose-add-x509-root-cert --certificate="$root_cert_path_for_1_4_3" --vid="$vid_for_1_4_3" --from=$trustee_account_1 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Add intermediate_cert_with_vid"
result=$(echo $passphrase | $DCLD_BIN_NEW tx pki add-x509-cert --certificate="$intermediate_cert_with_vid_path_for_1_4_3" --from=$vendor_account_for_1_4_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Revoke intermediate_cert_with_vid"
result=$(echo $passphrase | $DCLD_BIN_NEW tx pki revoke-x509-cert --subject="$intermediate_cert_with_vid_subject_for_1_4_3" --subject-key-id="$intermediate_cert_with_vid_subject_key_id_for_1_4_3" --serial-number="$intermediate_cert_with_vid_serial_number_for_1_4_3" --from=$vendor_account_for_1_4_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Propose revoke paa_cert_no_vid"
result=$(echo "$passphrase" | $DCLD_BIN_NEW tx pki propose-revoke-x509-root-cert --subject="$paa_cert_no_vid_subject_for_1_4_3" --subject-key-id="$paa_cert_no_vid_subject_key_id_for_1_4_3" --from="$trustee_account_1" --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve revoke paa_cert_no_vid_path_for_1_4_3"
result=$(echo "$passphrase" | $DCLD_BIN_NEW tx pki approve-revoke-x509-root-cert --subject="$paa_cert_no_vid_subject_for_1_4_3" --subject-key-id="$paa_cert_no_vid_subject_key_id_for_1_4_3" --from="$trustee_account_2" --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve revoke paa_cert_no_vid_path_for_1_4_3"
result=$(echo "$passphrase" | $DCLD_BIN_NEW tx pki approve-revoke-x509-root-cert --subject="$paa_cert_no_vid_subject_for_1_4_3" --subject-key-id="$paa_cert_no_vid_subject_key_id_for_1_4_3" --from="$trustee_account_3" --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve revoke paa_cert_no_vid_path_for_1_4_3"
result=$(echo "$passphrase" | $DCLD_BIN_NEW tx pki approve-revoke-x509-root-cert --subject="$paa_cert_no_vid_subject_for_1_4_3" --subject-key-id="$paa_cert_no_vid_subject_key_id_for_1_4_3" --from="$trustee_account_4" --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Propose revoke root_cert_with_vid"
result=$(echo $passphrase | $DCLD_BIN_NEW tx pki propose-revoke-x509-root-cert --subject="$root_cert_with_vid_subject_for_1_4_3" --subject-key-id="$root_cert_with_vid_subject_key_id_for_1_4_3" --from $trustee_account_1 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Emulation of actions, which has been done with NOC certificates on the TestNet"

echo "Add NOC Root certificate by vendor with VID = $vid_for_1_4_3"
result=$(echo "$passphrase" | $DCLD_BIN_NEW tx pki add-noc-x509-root-cert --certificate="$noc_root_cert_1_path_for_1_4_3" --from $vendor_account_for_1_4_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Add NOC ICA certificate by vendor with VID = $vid_for_1_4_3"
result=$(echo "$passphrase" | $DCLD_BIN_NEW tx pki add-noc-x509-ica-cert --certificate="$noc_ica_cert_1_path_for_1_4_3" --from $vendor_account_for_1_4_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Remove NOC root certificate by vendor with VID = $vid_for_1_4_3"
result=$(echo "$passphrase" | $DCLD_BIN_NEW tx pki remove-noc-x509-root-cert --subject="$noc_root_cert_1_subject_for_1_4_3" --subject-key-id="$noc_root_cert_1_subject_key_id_for_1_4_3" --from=$vendor_account_for_1_4_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Remove NOC ICA certificate by vendor with VID = $vid_for_1_4_3"
result=$(echo "$passphrase" | $DCLD_BIN_NEW tx pki remove-noc-x509-ica-cert --subject="$noc_ica_cert_1_subject_for_1_4_3" --subject-key-id="$noc_ica_cert_1_subject_key_id_for_1_4_3" --from=$vendor_account_for_1_4_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

# PKI Revocation point

echo "Add new revocation point for"
result=$(echo $passphrase | $DCLD_BIN_NEW tx pki add-revocation-point --vid=$vid_for_1_4_3 --revocation-type=1 --is-paa="true" --certificate="$root_cert_with_vid_path_for_1_4_3" --label="$product_label_for_1_4_3" --data-url="$test_data_url_for_1_4_3" --issuer-subject-key-id=$issuer_subject_key_id --from=$vendor_account_for_1_4_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider


echo "Update revocation point"
result=$(echo $passphrase | $DCLD_BIN_NEW tx pki update-revocation-point --vid=$vid_for_1_4_3 --certificate="$root_cert_with_vid_path_for_1_4_3" --label="$product_label_for_1_4_3" --data-url="$test_data_url_for_1_4_3/new" --issuer-subject-key-id=$issuer_subject_key_id --from=$vendor_account_for_1_4_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Delete revocation point"
result=$(echo $passphrase | $DCLD_BIN_NEW tx pki delete-revocation-point --vid=$vid_for_1_4_3 --label="$product_label_for_1_4_3" --issuer-subject-key-id=$issuer_subject_key_id --from=$vendor_account_for_1_4_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Add new revocation point"
result=$(echo $passphrase | $DCLD_BIN_NEW tx pki add-revocation-point --vid=$vid_for_1_4_3 --revocation-type=1 --is-paa="true" --certificate="$root_cert_with_vid_path_for_1_4_3" --label="$product_label_for_1_4_3" --data-url="$test_data_url_for_1_4_3" --issuer-subject-key-id=$issuer_subject_key_id --from=$vendor_account_for_1_4_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Add revocation point for CRL SIGNER CERTIFICATE delegated by PAI"

result=$(echo $passphrase | $DCLD_BIN_NEW tx pki add-revocation-point --vid=$vid_for_1_4_3 --is-paa="false" --certificate="$crl_signer_delegated_by_pai_1" --label="$product_label_for_1_4_3" --data-url="$test_data_url_for_1_4_3" --issuer-subject-key-id=$delegator_cert_with_vid_subject_key_id --revocation-type=1 --certificate-delegator="$delegator_cert_with_vid_65521_path" --from=$vendor_account_for_1_4_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Update revocation point for CRL SIGNER CERTIFICATE delegated by PAI"
result=$(echo $passphrase | $DCLD_BIN_NEW tx pki update-revocation-point --vid=$vid_for_1_4_3 --certificate="$crl_signer_delegated_by_pai_1" --label="$product_label_for_1_4_3" --data-url="$test_data_url_for_1_4_3/new" --issuer-subject-key-id=$delegator_cert_with_vid_subject_key_id --certificate-delegator="$delegator_cert_with_vid_65521_path" --from=$vendor_account_for_1_4_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

# AUTH

echo "Propose add account $user_7_address"
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth propose-add-account --address="$user_7_address" --pubkey="$user_7_pubkey" --roles="CertificationCenter" --from="$trustee_account_1" --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_7_address"
result=$($DCLD_BIN_NEW tx auth approve-add-account --address="$user_7_address" --from=$trustee_account_2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_7_address"
result=$($DCLD_BIN_NEW tx auth approve-add-account --address="$user_7_address" --from=$trustee_account_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_7_address"
result=$($DCLD_BIN_NEW tx auth approve-add-account --address="$user_7_address" --from=$trustee_account_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Propose add account $user_8_address"
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth propose-add-account --address="$user_8_address" --pubkey=$user_8_pubkey --roles=CertificationCenter --from=$trustee_account_1 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_8_address"
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth approve-add-account --address="$user_8_address" --from=$trustee_account_2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_8_address"
result=$($DCLD_BIN_NEW tx auth approve-add-account --address="$user_8_address" --from=$trustee_account_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_8_address"
result=$($DCLD_BIN_NEW tx auth approve-add-account --address="$user_8_address" --from=$trustee_account_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Propose add account $user_9_address"
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth propose-add-account --address="$user_9_address" --pubkey=$user_9_pubkey --roles=CertificationCenter --from=$trustee_account_1 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Propose revoke account $user_7_address"
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth propose-revoke-account --address="$user_7_address" --from=$trustee_account_1 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve revoke account $user_7_address"
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth approve-revoke-account --address="$user_7_address" --from=$trustee_account_2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve revoke account $user_7_address"
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth approve-revoke-account --address="$user_7_address" --from=$trustee_account_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve revoke account $user_7_address"
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth approve-revoke-account --address="$user_7_address" --from=$trustee_account_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Propose revoke account $user_8_address"
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth propose-revoke-account --address="$user_8_address" --from=$trustee_account_1 --yes)
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

echo "Verify if VendorInfo Record for VID: $vid_for_1_4_3 is present or not"
result=$($DCLD_BIN_NEW query vendorinfo vendor --vid=$vid_for_1_4_3)
check_response "$result" "\"vendorID\": $vid_for_1_4_3"
check_response "$result" "\"companyLegalName\": \"$company_legal_name_for_1_4_3\""

echo "Verify if VendorInfo Record for VID: $vid_for_1_2 updated or not"
result=$($DCLD_BIN_NEW query vendorinfo vendor --vid=$vid_for_1_2)
check_response "$result" "\"vendorID\": $vid_for_1_2"
check_response "$result" "\"vendorName\": \"$vendor_name_for_1_2\""
check_response "$result" "\"companyPreferredName\": \"$company_preferred_name_for_1_4_3\""
check_response "$result" "\"vendorLandingPageURL\": \"$vendor_landing_page_url_for_1_4_3\""

echo "Request all vendor infos"
result=$($DCLD_BIN_NEW query vendorinfo all-vendors)
check_response "$result" "\"vendorID\": $vid_for_1_4_3"
check_response "$result" "\"companyLegalName\": \"$company_legal_name_for_1_4_3\""
check_response "$result" "\"vendorName\": \"$vendor_name_for_1_4_3\""

test_divider

# MODEL

echo "Get Model with VID: $vid_for_1_4_3 PID: $pid_1_for_1_4_3"
result=$($DCLD_BIN_NEW query model get-model --vid=$vid_for_1_4_3 --pid=$pid_1_for_1_4_3)
check_response "$result" "\"vid\": $vid_for_1_4_3"
check_response "$result" "\"pid\": $pid_1_for_1_4_3"
check_response "$result" "\"productLabel\": \"$product_label_for_1_4_3\""

echo "Get Model with VID: $vid_for_1_4_3 PID: $pid_2_for_1_4_3"
result=$($DCLD_BIN_NEW query model get-model --vid=$vid_for_1_4_3 --pid=$pid_2_for_1_4_3)
check_response "$result" "\"vid\": $vid_for_1_4_3"
check_response "$result" "\"pid\": $pid_2_for_1_4_3"
check_response "$result" "\"productLabel\": \"$product_label_for_1_4_3\""

echo "Check Model with VID: $vid_for_1_4_3 PID: $pid_2_for_1_4_3 updated"
result=$($DCLD_BIN_NEW query model get-model --vid=$vid --pid=$pid_2)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_2"
check_response "$result" "\"productLabel\": \"$product_label_for_1_4_3\""
check_response "$result" "\"partNumber\": \"$part_number_for_1_4_3\""

echo "Check Model version with VID: $vid_for_1_4_3 PID: $pid_2_for_1_4_3 updated"
result=$($DCLD_BIN_NEW query model model-version --vid=$vid --pid=$pid_2  --softwareVersion=$software_version)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_2"
check_response "$result" "\"minApplicableSoftwareVersion\": $min_applicable_software_version_for_1_4_3"
check_response "$result" "\"maxApplicableSoftwareVersion\": $max_applicable_software_version_for_1_4_3"

echo "Get all models"
result=$($DCLD_BIN_NEW query model all-models)
check_response "$result" "\"vid\": $vid_for_1_4_3"
check_response "$result" "\"pid\": $pid_1_for_1_4_3"
check_response "$result" "\"pid\": $pid_2_for_1_4_3"

echo "Get all model versions"
result=$($DCLD_BIN_NEW query model all-model-versions --vid=$vid_for_1_4_3 --pid=$pid_1_for_1_4_3)
check_response "$result" "\"vid\": $vid_for_1_4_3"
check_response "$result" "\"pid\": $pid_1_for_1_4_3"

echo "Get Vendor Models with VID: ${vid_for_1_4_3}"
result=$($DCLD_BIN_NEW query model vendor-models --vid=$vid_for_1_4_3)
check_response "$result" "\"pid\": $pid_1_for_1_4_3"
check_response "$result" "\"pid\": $pid_2_for_1_4_3"

echo "Get model version VID: $vid_for_1_4_3 PID: $pid_1_for_1_4_3"
result=$($DCLD_BIN_NEW query model model-version --vid=$vid_for_1_4_3 --pid=$pid_1_for_1_4_3 --softwareVersion=$software_version_for_1_4_3)
check_response "$result" "\"vid\": $vid_for_1_4_3"
check_response "$result" "\"pid\": $pid_1_for_1_4_3"
check_response "$result" "\"softwareVersion\": $software_version_for_1_4_3"

echo "Get model version VID: $vid_for_1_4_3 PID: $pid_2_for_1_4_3"
result=$($DCLD_BIN_NEW query model model-version --vid=$vid_for_1_4_3 --pid=$pid_2_for_1_4_3 --softwareVersion=$software_version_for_1_4_3)
check_response "$result" "\"vid\": $vid_for_1_4_3"
check_response "$result" "\"pid\": $pid_2_for_1_4_3"
check_response "$result" "\"softwareVersion\": $software_version_for_1_4_3"

test_divider

# COMPLIANCE

echo "Get certified model vid=$vid_for_1_4_3 pid=$pid_1_for_1_4_3"
result=$($DCLD_BIN_NEW query compliance certified-model --vid=$vid_for_1_4_3 --pid=$pid_1_for_1_4_3 --softwareVersion=$software_version_for_1_4_3 --certificationType=$certification_type_for_1_4_3)
check_response "$result" "\"value\": true"
check_response "$result" "\"vid\": $vid_for_1_4_3"
check_response "$result" "\"pid\": $pid_1_for_1_4_3"
check_response "$result" "\"softwareVersion\": $software_version_for_1_4_3"
check_response "$result" "\"certificationType\": \"$certification_type_for_1_4_3\""

echo "Get revoked Model with VID: $vid_for_1_4_3 PID: $pid_2_for_1_4_3"
result=$($DCLD_BIN_NEW query compliance revoked-model --vid=$vid_for_1_4_3 --pid=$pid_2_for_1_4_3 --softwareVersion=$software_version_for_1_4_3 --certificationType=$certification_type_for_1_4_3)
check_response "$result" "\"vid\": $vid_for_1_4_3"
check_response "$result" "\"pid\": $pid_2_for_1_4_3"

echo "Get certified model with VID: $vid_for_1_4_3 PID: $pid_1_for_1_4_3"
result=$($DCLD_BIN_NEW query compliance certified-model --vid=$vid_for_1_4_3 --pid=$pid_1_for_1_4_3 --softwareVersion=$software_version_for_1_4_3 --certificationType=$certification_type_for_1_4_3)
check_response "$result" "\"value\": true"
check_response "$result" "\"vid\": $vid_for_1_4_3"
check_response "$result" "\"pid\": $pid_1_for_1_4_3"

echo "Get provisional model with VID: $vid_for_1_4_3 PID: $pid_2_for_1_4_3"
result=$($DCLD_BIN_NEW query compliance provisional-model --vid=$vid_for_1_4_3 --pid=$pid_2_for_1_4_3 --softwareVersion=$software_version_for_1_4_3 --certificationType=$certification_type_for_1_4_3)
check_response "$result" "\"value\": false"
check_response "$result" "\"vid\": $vid_for_1_4_3"
check_response "$result" "\"pid\": $pid_2_for_1_4_3"

echo "Get compliance-info model with VID: $vid_for_1_4_3 PID: $pid_1_for_1_4_3"
result=$($DCLD_BIN_NEW query compliance compliance-info --vid=$vid_for_1_4_3 --pid=$pid_1_for_1_4_3 --softwareVersion=$software_version_for_1_4_3 --certificationType=$certification_type_for_1_4_3)
check_response "$result" "\"vid\": $vid_for_1_4_3"
check_response "$result" "\"pid\": $pid_1_for_1_4_3"
check_response "$result" "\"softwareVersion\": $software_version_for_1_4_3"
check_response "$result" "\"certificationType\": \"$certification_type_for_1_4_3\""

echo "Get compliance-info model with VID: $vid_for_1_4_3 PID: $pid_2_for_1_4_3"
result=$($DCLD_BIN_NEW query compliance compliance-info --vid=$vid_for_1_4_3 --pid=$pid_2_for_1_4_3 --softwareVersion=$software_version_for_1_4_3 --certificationType=$certification_type_for_1_4_3)
check_response "$result" "\"vid\": $vid_for_1_4_3"
check_response "$result" "\"pid\": $pid_2_for_1_4_3"
check_response "$result" "\"softwareVersion\": $software_version_for_1_4_3"
check_response "$result" "\"certificationType\": \"$certification_type_for_1_4_3\""

echo "Get device software compliance cDCertificateId=$cd_certificate_id_for_1_4_3"
result=$($DCLD_BIN_NEW query compliance device-software-compliance --cdCertificateId=$cd_certificate_id_for_1_4_3)
check_response "$result" "\"vid\": $vid_for_1_4_3"
check_response "$result" "\"pid\": $pid_1_for_1_4_3"

echo "Get all certified models"
result=$($DCLD_BIN_NEW query compliance all-certified-models)
check_response "$result" "\"vid\": $vid_for_1_4_3"
check_response "$result" "\"pid\": $pid_1_for_1_4_3"

echo "Get all provisional models"
result=$($DCLD_BIN_NEW query compliance all-provisional-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_3"

echo "Get all revoked models"
result=$($DCLD_BIN_NEW query compliance all-revoked-models)
check_response "$result" "\"vid\": $vid_for_1_4_3"
check_response "$result" "\"pid\": $pid_2_for_1_4_3"

echo "Get all compliance infos"
result=$($DCLD_BIN_NEW query compliance all-compliance-info)
check_response "$result" "\"vid\": $vid_for_1_4_3"
check_response "$result" "\"pid\": $pid_1_for_1_4_3"
check_response "$result" "\"pid\": $pid_2_for_1_4_3"

echo "Get all device software compliances"
result=$($DCLD_BIN_NEW query compliance all-device-software-compliance)
check_response "$result" "\"vid\": $vid_for_1_4_3"
check_response "$result" "\"pid\": $pid_1_for_1_4_3"
check_response "$result" "\"cDCertificateId\": \"$cd_certificate_id_for_1_4_3\""

test_divider

# PKI

echo "Get x509 root certificate"
result=$($DCLD_BIN_NEW query pki x509-cert --subject=$root_cert_with_vid_subject_for_1_4_3 --subject-key-id=$root_cert_with_vid_subject_key_id_for_1_4_3)
check_response "$result" "\"subject\": \"$root_cert_with_vid_subject_for_1_4_3\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_with_vid_subject_key_id_for_1_4_3\""
check_response "$result" "\"vid\": $vid_for_1_4_3"

echo "Get all subject x509 root certificates"
result=$($DCLD_BIN_NEW query pki all-subject-x509-certs --subject=$root_cert_with_vid_subject_for_1_4_3)
check_response "$result" "\"subject\": \"$root_cert_with_vid_subject_for_1_4_3\""
check_response "$result" "$root_cert_with_vid_subject_key_id_for_1_4_3"

echo "Get all x509 root certificates by SKID"
result=$($DCLD_BIN_NEW query pki x509-cert --subject-key-id=$root_cert_with_vid_subject_key_id_for_1_4_3)
check_response "$result" "\"subjectKeyId\": \"$root_cert_with_vid_subject_key_id_for_1_4_3\""
check_response "$result" "$root_cert_with_vid_subject_for_1_4_3"

echo "Get proposed x509 root certificate"
result=$($DCLD_BIN_NEW query pki proposed-x509-root-cert --subject=$root_cert_subject_for_1_4_3 --subject-key-id=$root_cert_subject_key_id_for_1_4_3)
check_response "$result" "\"subject\": \"$root_cert_subject_for_1_4_3\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id_for_1_4_3\""
check_response "$result" "\"vid\": $vid_for_1_4_3"

echo "Get revoked x509 certificate"
result=$($DCLD_BIN_NEW query pki revoked-x509-cert --subject=$intermediate_cert_with_vid_subject_for_1_4_3 --subject-key-id=$intermediate_cert_with_vid_subject_key_id_for_1_4_3)
check_response "$result" "\"subject\": \"$intermediate_cert_with_vid_subject_for_1_4_3\""
check_response "$result" "\"subjectKeyId\": \"$intermediate_cert_with_vid_subject_key_id_for_1_4_3\""
check_response "$result" "\"vid\": $intermediate_cert_with_vid_65521_vid_for_1_4_3"

echo "Get proposed x509 root certificate to revoke"
result=$($DCLD_BIN_NEW query pki proposed-x509-root-cert-to-revoke --subject=$root_cert_with_vid_subject_for_1_4_3 --subject-key-id=$root_cert_with_vid_subject_key_id_for_1_4_3)
check_response "$result" "\"subject\": \"$root_cert_with_vid_subject_for_1_4_3\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_with_vid_subject_key_id_for_1_4_3\""

echo "Get revocation point"
result=$($DCLD_BIN_NEW query pki revocation-point --vid=$vid_for_1_4_3 --label=$product_label_for_1_4_3 --issuer-subject-key-id=$issuer_subject_key_id)
check_response "$result" "\"vid\": $vid_for_1_4_3"
check_response "$result" "\"issuerSubjectKeyID\": \"$issuer_subject_key_id\""
check_response "$result" "\"label\": \"$product_label_for_1_4_3\""
check_response "$result" "\"dataURL\": \"$test_data_url_for_1_4_3\""

echo "Get revocation points by issuer subject key id"
result=$($DCLD_BIN_NEW query pki revocation-points --issuer-subject-key-id=$issuer_subject_key_id)
check_response "$result" "\"vid\": $vid_for_1_4_3"
check_response "$result" "\"issuerSubjectKeyID\": \"$issuer_subject_key_id\""
check_response "$result" "\"label\": \"$product_label_for_1_4_3\""
check_response "$result" "\"dataURL\": \"$test_data_url_for_1_4_3\""

echo "Get all revocation points"
result=$($DCLD_BIN_NEW query pki all-revocation-points)
check_response "$result" "\"vid\": $vid_for_1_4_3"
check_response "$result" "\"issuerSubjectKeyID\": \"$issuer_subject_key_id\""
check_response "$result" "\"label\": \"$product_label_for_1_4_3\""
check_response "$result" "\"dataURL\": \"$test_data_url_for_1_4_3\""

echo "Get all certificates"
result=$($DCLD_BIN_NEW query pki all-x509-certs)
check_response "$result" "\[\]"
response_does_not_contain "$result" "$noc_root_cert_1_subject_key_id_for_1_4_3"
response_does_not_contain "$result" "$noc_ica_cert_1_subject_key_id_for_1_4_3"

echo "Get all noc x509 root certificates"
result=$($DCLD_BIN_NEW query pki noc-x509-root-certs --vid=$vid_for_1_4_3)
check_response "$result" "Not Found"
response_does_not_contain "$result" "$noc_root_cert_1_subject_key_id_for_1_4_3"

echo "Get noc x509 root certificates by vid=$vid_for_1_4_3 and skid=$noc_root_cert_1_subject_key_id_for_1_4_3 (must be empty)"
result=$($DCLD_BIN_NEW query pki noc-x509-certs --vid=$vid_for_1_4_3 --subject-key-id=$noc_root_cert_1_subject_key_id_for_1_4_3)
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$noc_root_cert_1_subject_for_1_4_3\""
response_does_not_contain "$result" "$noc_root_cert_1_subject_key_id_for_1_4_3"

echo "Get noc x509 ica certificates by vid=$vid_for_1_4_3 and skid=$noc_ica_cert_1_subject_key_id_for_1_4_3 (must be empty)"
result=$($DCLD_BIN_NEW query pki noc-x509-certs --vid=$vid_for_1_4_3 --subject-key-id="$noc_ica_cert_1_subject_key_id_for_1_4_3")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$noc_ica_cert_1_subject_for_1_4_3\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_ica_cert_1_subject_key_id_for_1_4_3\""

test_divider

# AUTH

echo "Get all accounts"
result=$($DCLD_BIN_NEW query auth all-accounts)
check_response "$result" "\"address\": \"$user_8_address\""

echo "Get account"
result=$($DCLD_BIN_NEW query auth account --address=$user_8_address)
check_response "$result" "\"address\": \"$user_8_address\""

echo "Get all proposed accounts"
result=$($DCLD_BIN_NEW query auth all-proposed-accounts)
check_response "$result" "\"address\": \"$user_9_address\""

echo "Get proposed account"
result=$($DCLD_BIN_NEW query auth proposed-account --address=$user_9_address)
check_response "$result" "\"address\": \"$user_9_address\""

echo "Get all proposed accounts to revoke"
result=$($DCLD_BIN_NEW query auth all-proposed-accounts-to-revoke)
check_response "$result" "\"address\": \"$user_8_address\""

echo "Get proposed account to revoke"
result=$($DCLD_BIN_NEW query auth proposed-account-to-revoke --address=$user_8_address)
check_response "$result" "\"address\": \"$user_8_address\""

echo "Get all revoked accounts"
result=$($DCLD_BIN_NEW query auth all-revoked-accounts)
check_response "$result" "\"address\": \"$user_7_address\""

echo "Get revoked account"
result=$($DCLD_BIN_NEW query auth revoked-account --address=$user_7_address)
check_response "$result" "\"address\": \"$user_7_address\""

test_divider

# Validator

echo "Get node"
# FIXME: use proper binary (not dcld but $DCLD_BIN_OLD)
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator all-nodes")
check_response "$result" "\"owner\": \"$validator_address\""

test_divider

echo "Upgrade from 1.2 to 1.4.3 passed"
