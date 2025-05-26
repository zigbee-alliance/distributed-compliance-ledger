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

plan_name="wrong_plan_name_2"
upgrade_checksum="sha256:a007f58d61632af107a09c89b7392eedd05d8127d0df67ace50f318948c62001"
binary_version_old="v1.2.2"
binary_version_new="v1.4.3"

wget -O dcld_old "https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/$binary_version_old/dcld"
chmod ugo+x dcld_old

DCLD_BIN_OLD="./dcld_old"
DCLD_BIN_NEW="./dcld_old"
$DCLD_BIN_NEW config broadcast-mode sync
########################################################################################

# Upgrade to version wrong_plan_name_2

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

echo "Verify that upgrade is not applied"
result=$(! $DCLD_BIN_NEW query upgrade applied $plan_name)
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
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator all-nodes")
check_response "$result" "\"owner\": \"$validator_address\""

########################################################################################

# after upgrade constants

vid_for_1_2_r2=4703
pid_1_for_1_2_r2=16
pid_2_for_1_2_r2=27
pid_3_for_1_2_r2=38
device_type_id_for_1_2_r2=1239
product_name_for_1_2_r2="ProductName1.2_r2"
product_label_for_1_2_r2="ProductLabe1.2_r2"
part_number_for_1_2_r2="RCU2205F"
software_version_for_1_2_r2=2
software_version_string_for_1_2_r2="2.0"
cd_version_number_for_1_2_r2=313
min_applicable_software_version_for_1_2_r2=2
max_applicable_software_version_for_1_2_r2=2000

certification_type_for_1_2_r2="matter"
certification_date_for_1_2_r2="2021-01-03T00:00:00Z"
provisional_date_for_1_2_r2="2010-12-11T00:00:00Z"
cd_certificate_id_for_1_2_r2="15DEXZ"

# root_cert_path_for_1_2_r2="integration_tests/constants/google_root_cert_gsr4"
# root_cert_subject_for_1_2_r2="MFAxJDAiBgNVBAsTG0dsb2JhbFNpZ24gRUNDIFJvb3QgQ0EgLSBSNDETMBEGA1UEChMKR2xvYmFsU2lnbjETMBEGA1UEAxMKR2xvYmFsU2lnbg=="
# root_cert_subject_key_id_for_1_2_r2="54:B0:7B:AD:45:B8:E2:40:7F:FB:0A:6E:FB:BE:33:C9:3C:A3:84:D5"
# root_cert_path_for_1_2_r2_random_vid="1234"

# test_root_cert_path_for_1_2_r2="integration_tests/constants/paa_cert_numeric_vid"
# test_root_cert_subject_for_1_2_r2="MDAxGDAWBgNVBAMMD01hdHRlciBUZXN0IFBBQTEUMBIGCisGAQQBgqJ8AgEMBEZGRjE="
# test_root_cert_subject_key_id_for_1_2_r2="6A:FD:22:77:1F:51:1F:EC:BF:16:41:97:67:10:DC:DC:31:A1:71:7E"
# test_root_cert_vid_for_1_2_r2="65521"

# google_root_cert_path_for_1_2_r2="integration_tests/constants/google_root_cert_r2"
# google_root_cert_subject_for_1_2_r2="MEcxCzAJBgNVBAYTAlVTMSIwIAYDVQQKExlHb29nbGUgVHJ1c3QgU2VydmljZXMgTExDMRQwEgYDVQQDEwtHVFMgUm9vdCBSMg=="
# google_root_cert_subject_key_id_for_1_2_r2="BB:FF:CA:8E:23:9F:4F:99:CA:DB:E2:68:A6:A5:15:27:17:1E:D9:0E"
# google_root_cert_path_random_vid_for_1_2_r2="1234"

# intermediate_cert_path_for_1_2_r2="integration_tests/constants/intermediate_cert_gsr4"
# intermediate_cert_subject_for_1_2_r2="MEYxCzAJBgNVBAYTAlVTMSIwIAYDVQQKExlHb29nbGUgVHJ1c3QgU2VydmljZXMgTExDMRMwEQYDVQQDEwpHVFMgQ0EgMkQ0"
# intermediate_cert_subject_key_id_for_1_2_r2="A8:88:D9:8A:39:AC:65:D5:82:4B:37:A8:95:6C:65:43:CD:44:01:E0"

test_data_url="https://url.data.dclmodel"
issuer_subject_key_id="5A880E6C3653D07FB08971A3F473790930E62BDB"

vendor_name_for_1_2_r2="VendorName4705"
company_legal_name_for_1_2_r2="LegalCompanyName4705"
company_preferred_name_for_1_2_r2="CompanyPreferredName4705"
vendor_landing_page_url_for_1_2_r2="https://www.newexample_R2.com"

vendor_account_for_1_2_r2="vendor_account_4705"
vendor_admin_account="vendor_admin_account"
certification_center_account="certification_center_account"

# send all ledger update transactions after upgrade

echo "Create Vendor account $vendor_account_for_1_2_r2"

result="$(echo $passphrase | $DCLD_BIN_NEW keys add "$vendor_account_for_1_2_r2")"
result=$(get_txn_result "$result")
_address=$(echo $passphrase | $DCLD_BIN_NEW keys show $vendor_account_for_1_2_r2 -a)
result=$(get_txn_result "$result")
_pubkey=$(echo $passphrase | $DCLD_BIN_NEW keys show $vendor_account_for_1_2_r2 -p)
result=$(get_txn_result "$result")
result="$(echo $passphrase | $DCLD_BIN_NEW tx auth propose-add-account --address="$_address" --pubkey="$_pubkey" --vid="$vid_for_1_2_r2" --roles="Vendor" --from "$trustee_account_1" --yes)"
echo "propose-add-account $result"
result=$(get_txn_result "$result")
result="$(echo $passphrase | $DCLD_BIN_NEW tx auth approve-add-account --address="$_address" --from "$trustee_account_2" --yes)"
result=$(get_txn_result "$result")
result="$(echo $passphrase | $DCLD_BIN_NEW tx auth approve-add-account --address="$_address" --from "$trustee_account_3" --yes)"
result=$(get_txn_result "$result")
result="$(echo $passphrase | $DCLD_BIN_NEW tx auth approve-add-account --address="$_address" --from "$trustee_account_4" --yes)"

random_string user_4
echo "$user_4 generates keys"
cmd="(echo $passphrase; echo $passphrase) | $DCLD_BIN_NEW keys add $user_4"
result="$(bash -c "$cmd")"
user_4_address=$(echo $passphrase | $DCLD_BIN_NEW keys show $user_4 -a)
user_4_pubkey=$(echo $passphrase | $DCLD_BIN_NEW keys show $user_4 -p)

random_string user_5
echo "$user_5 generates keys"
cmd="(echo $passphrase; echo $passphrase) | $DCLD_BIN_NEW keys add $user_5"
result="$(bash -c "$cmd")"
user_5_address=$(echo $passphrase | $DCLD_BIN_NEW keys show $user_5 -a)
user_5_pubkey=$(echo $passphrase | $DCLD_BIN_NEW keys show $user_5 -p)

random_string user_6
echo "$user_6 generates keys"
cmd="(echo $passphrase; echo $passphrase) | $DCLD_BIN_NEW keys add $user_6"
result="$(bash -c "$cmd")"
user_6_address=$(echo $passphrase | $DCLD_BIN_NEW keys show $user_6 -a)
user_6_pubkey=$(echo $passphrase | $DCLD_BIN_NEW keys show $user_6 -p)

# VENDOR_INFO
echo "Add vendor $vendor_name_for_1_2_r2"
result=$(echo $passphrase | $DCLD_BIN_NEW tx vendorinfo add-vendor --vid=$vid_for_1_2_r2 --vendorName=$vendor_name_for_1_2_r2 --companyLegalName=$company_legal_name_for_1_2_r2 --companyPreferredName=$company_preferred_name_for_1_2_r2 --vendorLandingPageURL=$vendor_landing_page_url_for_1_2_r2 --from=$vendor_account_for_1_2_r2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

# MODEL and MODEL_VERSION

echo "Add model vid=$vid_for_1_2_r2 pid=$pid_1_for_1_2_r2"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model --vid=$vid_for_1_2_r2 --pid=$pid_1_for_1_2_r2 --deviceTypeID=$device_type_id_for_1_2_r2 --productName=$product_name_for_1_2_r2 --productLabel=$product_label_for_1_2_r2 --partNumber=$part_number_for_1_2_r2 --from=$vendor_account_for_1_2_r2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Add model version vid=$vid_for_1_2_r2 pid=$pid_1_for_1_2_r2"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model-version --vid=$vid_for_1_2_r2 --pid=$pid_1_for_1_2_r2 --softwareVersion=$software_version_for_1_2_r2 --softwareVersionString=$software_version_string_for_1_2_r2 --cdVersionNumber=$cd_version_number_for_1_2_r2 --minApplicableSoftwareVersion=$min_applicable_software_version_for_1_2_r2 --maxApplicableSoftwareVersion=$max_applicable_software_version_for_1_2_r2 --from=$vendor_account_for_1_2_r2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Add model vid=$vid_for_1_2_r2 pid=$pid_2_for_1_2_r2"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model --vid=$vid_for_1_2_r2 --pid=$pid_2_for_1_2_r2 --deviceTypeID=$device_type_id_for_1_2_r2 --productName=$product_name_for_1_2_r2 --productLabel=$product_label_for_1_2_r2 --partNumber=$part_number_for_1_2_r2 --from=$vendor_account_for_1_2_r2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Add model version vid=$vid_for_1_2_r2 pid=$pid_2_for_1_2_r2"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model-version --vid=$vid_for_1_2_r2 --pid=$pid_2_for_1_2_r2 --softwareVersion=$software_version_for_1_2_r2 --softwareVersionString=$software_version_string_for_1_2_r2 --cdVersionNumber=$cd_version_number_for_1_2_r2 --minApplicableSoftwareVersion=$min_applicable_software_version_for_1_2_r2 --maxApplicableSoftwareVersion=$max_applicable_software_version_for_1_2_r2 --from=$vendor_account_for_1_2_r2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Add model vid=$vid_for_1_2_r2 pid=$pid_3_for_1_2_r2"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model --vid=$vid_for_1_2_r2 --pid=$pid_3_for_1_2_r2 --deviceTypeID=$device_type_id_for_1_2_r2 --productName=$product_name_for_1_2_r2 --productLabel=$product_label_for_1_2_r2 --partNumber=$part_number_for_1_2_r2 --from=$vendor_account_for_1_2_r2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Add model version vid=$vid_for_1_2_r2 pid=$pid_3_for_1_2_r2"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model-version --vid=$vid_for_1_2_r2 --pid=$pid_3_for_1_2_r2 --softwareVersion=$software_version_for_1_2_r2 --softwareVersionString=$software_version_string_for_1_2_r2 --cdVersionNumber=$cd_version_number_for_1_2_r2 --minApplicableSoftwareVersion=$min_applicable_software_version_for_1_2_r2 --maxApplicableSoftwareVersion=$max_applicable_software_version_for_1_2_r2 --from=$vendor_account_for_1_2_r2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Delete model vid=$vid_for_1_2_r2 pid=$pid_3_for_1_2_r2"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model delete-model --vid=$vid_for_1_2_r2 --pid=$pid_3_for_1_2_r2 --from=$vendor_account_for_1_2_r2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

# CERTIFY_DEVICE_COMPLIANCE

echo "Certify model vid=$vid_for_1_2_r2 pid=$pid_1_for_1_2_r2"
result=$(echo $passphrase | $DCLD_BIN_NEW tx compliance certify-model --vid=$vid_for_1_2_r2 --pid=$pid_1_for_1_2_r2 --softwareVersion=$software_version_for_1_2_r2 --softwareVersionString=$software_version_string_for_1_2_r2  --certificationType=$certification_type_for_1_2_r2 --certificationDate=$certification_date_for_1_2_r2 --cdCertificateId=$cd_certificate_id_for_1_2_r2 --from=$certification_center_account --cdVersionNumber=$cd_version_number_for_1_2_r2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Provision model vid=$vid_for_1_2_r2 pid=$pid_2_for_1_2_r2"
result=$(echo $passphrase | $DCLD_BIN_NEW tx compliance provision-model --vid=$vid_for_1_2_r2 --pid=$pid_2_for_1_2_r2 --softwareVersion=$software_version_for_1_2_r2 --softwareVersionString=$software_version_string_for_1_2_r2 --certificationType=$certification_type_for_1_2_r2 --provisionalDate=$provisional_date_for_1_2_r2 --cdCertificateId=$cd_certificate_id_for_1_2_r2 --from=$certification_center_account --cdVersionNumber=$cd_version_number_for_1_2_r2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Certify model vid=$vid_for_1_2_r2 pid=$pid_2_for_1_2_r2"
result=$(echo $passphrase | $DCLD_BIN_NEW tx compliance certify-model --vid=$vid_for_1_2_r2 --pid=$pid_2_for_1_2_r2 --softwareVersion=$software_version_for_1_2_r2 --softwareVersionString=$software_version_string_for_1_2_r2  --certificationType=$certification_type_for_1_2_r2 --certificationDate=$certification_date_for_1_2_r2 --cdCertificateId=$cd_certificate_id_for_1_2_r2 --from=$certification_center_account --cdVersionNumber=$cd_version_number_for_1_2_r2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

# echo "Revoke model certification vid=$vid_for_1_2_r2 pid=$pid_2_for_1_2_r2"
# result=$(echo $passphrase | $DCLD_BIN_NEW tx compliance revoke-model --vid=$vid_for_1_2_r2 --pid=$pid_2_for_1_2_r2 --softwareVersion=$software_version_for_1_2_r2 --softwareVersionString=$software_version_string_for_1_2_r2 --certificationType=$certification_type_for_1_2_r2 --revocationDate=$certification_date_for_1_2_r2 --from=$certification_center_account --cdVersionNumber=$cd_version_number_for_1_2_r2 --yes)
# result=$(get_txn_result "$result")
# check_response "$result" "\"code\": 0"

# test_divider

# X509 PKI

# echo "Assign VID to test_root_certificate"
# result=$(echo $passphrase | $DCLD_BIN_NEW tx pki assign-vid --subject="$test_root_cert_subject" --subject-key-id="$test_root_cert_subject_key_id" --vid="$test_root_cert_vid" --from $vendor_admin_account --yes)
# result=$(get_txn_result "$result")
# check_response "$result" "\"code\": 0"

# test_divider

# echo "Verify that vid is assigned to test_root_certificate"
# result=$($DCLD_BIN_NEW query pki x509-cert --subject="$test_root_cert_subject" --subject-key-id="$test_root_cert_subject_key_id")
# result=$(get_txn_result "$result")
# echo $result | jq
# check_response "$result" "\"subject\": \"$test_root_cert_subject\""
# check_response "$result" "\"subjectKeyId\": \"$test_root_cert_subject_key_id\""
# check_response "$result" "\"vid\": $test_root_cert_vid"

# test_divider

# echo "Propose add root_certificate"
# result=$(echo $passphrase | $DCLD_BIN_NEW tx pki propose-add-x509-root-cert --certificate="$root_cert_path_for_1_2_r2" --vid="$root_cert_path_for_1_2_r2_random_vid" --from=$trustee_account_1 --yes)
# result=$(get_txn_result "$result")
# check_response "$result" "\"code\": 0"

# test_divider

# echo "Approve add root_certificate"
# result=$(echo $passphrase | $DCLD_BIN_NEW tx pki approve-add-x509-root-cert --subject="$root_cert_subject_for_1_2_r2" --subject-key-id=$root_cert_subject_key_id_for_1_2_r2 --from=$trustee_account_2 --yes)
# result=$(get_txn_result "$result")
# check_response "$result" "\"code\": 0"

# test_divider

# echo "reject add root_certificate"
# result=$(echo $passphrase | $DCLD_BIN_NEW tx pki reject-add-x509-root-cert --subject="$root_cert_subject_for_1_2_r2" --subject-key-id=$root_cert_subject_key_id_for_1_2_r2 --from=$trustee_account_2 --yes)
# result=$(get_txn_result "$result")
# check_response "$result" "\"code\": 0"

# test_divider

# echo "Approve add root_certificate"
# result=$(echo $passphrase | $DCLD_BIN_NEW tx pki approve-add-x509-root-cert --subject="$root_cert_subject_for_1_2_r2" --subject-key-id=$root_cert_subject_key_id_for_1_2_r2 --from=$trustee_account_3 --yes)
# result=$(get_txn_result "$result")
# check_response "$result" "\"code\": 0"

# test_divider

# echo "Approve add root_certificate"
# result=$(echo $passphrase | $DCLD_BIN_NEW tx pki approve-add-x509-root-cert --subject="$root_cert_subject_for_1_2_r2" --subject-key-id=$root_cert_subject_key_id_for_1_2_r2 --from=$trustee_account_4 --yes)
# result=$(get_txn_result "$result")
# check_response "$result" "\"code\": 0"

# test_divider

# echo "Approve add root_certificate"
# result=$(echo $passphrase | $DCLD_BIN_NEW tx pki approve-add-x509-root-cert --subject="$root_cert_subject_for_1_2_r2" --subject-key-id=$root_cert_subject_key_id_for_1_2_r2 --from=$trustee_account_5 --yes)
# result=$(get_txn_result "$result")
# check_response "$result" "\"code\": 0"

# test_divider

# echo "Propose add test_root_certificate"
# result=$(echo $passphrase | $DCLD_BIN_NEW tx pki propose-add-x509-root-cert --certificate="$test_root_cert_path_for_1_2_r2" --vid="$test_root_cert_vid_for_1_2_r2" --from=$trustee_account_1 --yes)
# result=$(get_txn_result "$result")
# check_response "$result" "\"code\": 0"

# test_divider

# echo "Approve add test_root_certificate"
# result=$(echo $passphrase | $DCLD_BIN_NEW tx pki approve-add-x509-root-cert --subject="$test_root_cert_subject_for_1_2_r2" --subject-key-id=$test_root_cert_subject_key_id_for_1_2_r2 --from=$trustee_account_2 --yes)
# result=$(get_txn_result "$result")
# check_response "$result" "\"code\": 0"

# test_divider

# echo "Approve add test_root_certificate"
# result=$(echo $passphrase | $DCLD_BIN_NEW tx pki approve-add-x509-root-cert --subject="$test_root_cert_subject_for_1_2_r2" --subject-key-id=$test_root_cert_subject_key_id_for_1_2_r2 --from=$trustee_account_3 --yes)
# result=$(get_txn_result "$result")
# check_response "$result" "\"code\": 0"

# test_divider

# echo "Approve add test_root_certificate"
# result=$(echo $passphrase | $DCLD_BIN_NEW tx pki approve-add-x509-root-cert --subject="$test_root_cert_subject_for_1_2_r2" --subject-key-id=$test_root_cert_subject_key_id_for_1_2_r2 --from=$trustee_account_4 --yes)
# result=$(get_txn_result "$result")
# check_response "$result" "\"code\": 0"

# test_divider

# echo "Propose add google_root_certificate"
# result=$(echo $passphrase | $DCLD_BIN_NEW tx pki propose-add-x509-root-cert --certificate="$google_root_cert_path_for_1_2_r2" --vid="$google_root_cert_path_random_vid_for_1_2_r2" --from=$trustee_account_1 --yes)
# result=$(get_txn_result "$result")
# check_response "$result" "\"code\": 0"

# test_divider

# echo "Add intermediate_cert"
# result=$(echo $passphrase | $DCLD_BIN_NEW tx pki add-x509-cert --certificate="$intermediate_cert_path_for_1_2_r2" --from=$vendor_account --yes)
# result=$(get_txn_result "$result")
# check_response "$result" "\"code\": 0"

# test_divider

# echo "Revoke intermediate_cert"
# result=$(echo $passphrase | $DCLD_BIN_NEW tx pki revoke-x509-cert --subject="$intermediate_cert_subject_for_1_2_r2" --subject-key-id="$intermediate_cert_subject_key_id_for_1_2_r2" --from=$vendor_account --yes)
# result=$(get_txn_result "$result")
# check_response "$result" "\"code\": 0"

# test_divider

# echo "Propose revoke root_certificate"
# result=$(echo "$passphrase" | $DCLD_BIN_NEW tx pki propose-revoke-x509-root-cert --subject="$root_cert_subject_for_1_2_r2" --subject-key-id="$root_cert_subject_key_id_for_1_2_r2" --from="$trustee_account_1" --yes)
# result=$(get_txn_result "$result")
# check_response "$result" "\"code\": 0"

# test_divider

# echo "Approve revoke root_certificate"
# result=$(echo "$passphrase" | $DCLD_BIN_NEW tx pki approve-revoke-x509-root-cert --subject="$root_cert_subject_for_1_2_r2" --subject-key-id="$root_cert_subject_key_id_for_1_2_r2" --from="$trustee_account_2" --yes)
# result=$(get_txn_result "$result")
# check_response "$result" "\"code\": 0"

# test_divider

# echo "Approve revoke root_certificate"
# result=$(echo "$passphrase" | $DCLD_BIN_NEW tx pki approve-revoke-x509-root-cert --subject="$root_cert_subject_for_1_2_r2" --subject-key-id="$root_cert_subject_key_id_for_1_2_r2" --from="$trustee_account_3" --yes)
# result=$(get_txn_result "$result")
# check_response "$result" "\"code\": 0"

# test_divider

# echo "Approve revoke root_certificate"
# result=$(echo "$passphrase" | $DCLD_BIN_NEW tx pki approve-revoke-x509-root-cert --subject="$root_cert_subject_for_1_2_r2" --subject-key-id="$root_cert_subject_key_id_for_1_2_r2" --from="$trustee_account_4" --yes)
# result=$(get_txn_result "$result")
# check_response "$result" "\"code\": 0"

# test_divider

# echo "Propose revoke test_root_certificate"
# result=$(echo $passphrase | $DCLD_BIN_NEW tx pki propose-revoke-x509-root-cert --subject="$test_root_cert_subject_for_1_2_r2" --subject-key-id="$test_root_cert_subject_key_id_for_1_2_r2" --from $trustee_account_1 --yes)
# result=$(get_txn_result "$result")
# check_response "$result" "\"code\": 0"

# test_divider

# AUTH

echo "Propose add account $user_4_address"
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth propose-add-account --address="$user_4_address" --pubkey="$user_4_pubkey" --roles="CertificationCenter" --from="$trustee_account_1" --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_4_address"
result=$($DCLD_BIN_NEW tx auth approve-add-account --address="$user_4_address" --from=$trustee_account_2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_4_address"
result=$($DCLD_BIN_NEW tx auth approve-add-account --address="$user_4_address" --from=$trustee_account_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_4_address"
result=$($DCLD_BIN_NEW tx auth approve-add-account --address="$user_4_address" --from=$trustee_account_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Propose add account $user_5_address"
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth propose-add-account --address="$user_5_address" --pubkey=$user_5_pubkey --roles=CertificationCenter --from=$trustee_account_1 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_5_address"
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth approve-add-account --address="$user_5_address" --from=$trustee_account_2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_5_address"
result=$($DCLD_BIN_NEW tx auth approve-add-account --address="$user_5_address" --from=$trustee_account_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_5_address"
result=$($DCLD_BIN_NEW tx auth approve-add-account --address="$user_5_address" --from=$trustee_account_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Propose add account $user_6_address"
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth propose-add-account --address="$user_6_address" --pubkey=$user_6_pubkey --roles=CertificationCenter --from=$trustee_account_1 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Propose revoke account $user_4_address"
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth propose-revoke-account --address="$user_4_address" --from=$trustee_account_1 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve revoke account $user_4_address"
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth approve-revoke-account --address="$user_4_address" --from=$trustee_account_2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve revoke account $user_4_address"
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth approve-revoke-account --address="$user_4_address" --from=$trustee_account_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve revoke account $user_4_address"
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth approve-revoke-account --address="$user_4_address" --from=$trustee_account_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Propose revoke account $user_5_address"
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth propose-revoke-account --address="$user_5_address" --from=$trustee_account_1 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

# VALIDATOR_NODE
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
result=$(echo $passphrase | $DCLD_BIN_OLD tx validator propose-disable-node --address=$validator_address --from=$trustee_account_1 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

# Validator

echo "Get node"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator all-nodes")
check_response "$result" "\"owner\": \"$validator_address\""

test_divider

echo "Rollback when update to wrong_plan_name_2 PASSED"

rm -f $DCLD_BIN_OLD
