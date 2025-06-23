#!/bin/bash
# Copyright 2025 DSR Corporation
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

# TestNet constants

plan_name="v1.4.4"
chain_id=testnet-2.0
node_endpoint=https://on.test-net.dcl.csa-iot.org:26657
upgrade_checksum="sha256:e4031c6a77aa8e58add391be671a334613271bcf6e7f11d23b04a0881ece6958"

vid="1"
vid_1="4107"
vid_2="4161"
vid_3="4251" # certified model
vid_revoked="12288"
vid_complience="5218"
pid="1111"
pid_1="1"
pid_for_vid_3="4096"
pid_revoked="1000"
vendor_name="Panasonic"
company_legal_name="Panasonic Corporation"
software_version=1
software_version_for_vid_2=103
software_version_revoked=4
certification_type="matter"
cd_certificate_id="ZIG20142ZB330003-24"

dcld config broadcast-mode sync #TODO

echo "Configure CLI"
dcld config output json
dcld config chain-id $chain_id
dcld config node $node_endpoint

# UPGRADE

echo "Verify that upgrade is applied"
result=$(dcld query upgrade applied $plan_name)
echo "$result"

# test_divider

########################################################################################

echo "Verify that old data is not corrupted"

# VENDORINFO

echo "Verify if VendorInfo Record for VID: $vid is present or not"
result=$(dcld query vendorinfo vendor --vid=$vid)
check_response "$result" "\"vendorID\": $vid"
check_response "$result" "\"companyLegalName\": \"$company_legal_name\""
check_response "$result" "\"vendorName\": \"$vendor_name\""

echo "Request all vendor infos"
result=$(dcld query vendorinfo all-vendors)
check_response "$result" "\"vendorID\": $vid"
check_response "$result" "\"vendorID\": $vid_1"
check_response "$result" "\"vendorID\": $vid_2"

test_divider

# MODEL

echo "Get Model with VID: $vid PID: $pid"
result=$(dcld query model get-model --vid=$vid --pid=$pid)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"

echo "Get model version VID: $vid PID: $pid"
result=$(dcld query model model-version --vid=$vid --pid=$pid --softwareVersion=$software_version)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"softwareVersion\": $software_version"

echo "Get all models"
result=$(dcld query model all-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"

echo "Get Vendor Models with VID: ${vid}"
result=$(dcld query model vendor-models --vid=$vid)
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"pid\": $pid_1"

echo "Get all model versions with VID: $vid PID: $pid"
result=$(dcld query model all-model-versions --vid=$vid --pid=$pid)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
# check_response "$result" "\"softwareVersions\": [$software_version]" TODO: fix

test_divider

# COMPLIANCE

echo "Get certified model vid=$vid_2 pid=$pid_1"
result=$(dcld query compliance certified-model --vid=$vid_2 --pid=$pid_1 --softwareVersion=$software_version_for_vid_2 --certificationType=$certification_type)
check_response "$result" "\"value\": true"
check_response "$result" "\"vid\": $vid_2"
check_response "$result" "\"pid\": $pid_1"
check_response "$result" "\"softwareVersion\": $software_version_for_vid_2"
check_response "$result" "\"certificationType\": \"$certification_type\""

echo "Get all certified models"
result=$(dcld query compliance all-certified-models)
check_response "$result" "\"vid\": $vid_2"
check_response "$result" "\"pid\": $pid_1"
check_response "$result" "\"vid\": $vid_3"
check_response "$result" "\"pid\": $pid_for_vid_3"

echo "Get revoked Model with VID: $vid_revoked PID: $pid_revoked"
result=$(dcld query compliance revoked-model --vid=$vid_revoked --pid=$pid_revoked --softwareVersion=$software_version_revoked --certificationType=$certification_type)
check_response "$result" "\"vid\": $vid_revoked"
check_response "$result" "\"pid\": $pid_revoked"
check_response "$result" "\"softwareVersion\": $software_version_revoked"
check_response "$result" "\"certificationType\": \"$certification_type\""

echo "Get all revoked models"
result=$(dcld query compliance all-revoked-models)
check_response "$result" "\"vid\": $vid_revoked"
check_response "$result" "\"pid\": $pid_revoked"

echo "Get compliance-info model with VID: $vid_2 PID: $pid_1"
result=$(dcld query compliance compliance-info --vid=$vid_2 --pid=$pid_1 --softwareVersion=$software_version_for_vid_2 --certificationType=$certification_type)
check_response "$result" "\"vid\": $vid_2"
check_response "$result" "\"pid\": $pid_1"
check_response "$result" "\"softwareVersion\": $software_version_for_vid_2"
check_response "$result" "\"certificationType\": \"$certification_type\""

echo "Get device software compliance cDCertificateId=$cd_certificate_id"
result=$(dcld query compliance device-software-compliance --cdCertificateId=$cd_certificate_id)
check_response "$result" "\"vid\": $vid_complience"
check_response "$result" "\"pid\": $pid_1"
check_response "$result" "\"cDCertificateId\": \"$cd_certificate_id\""

echo "Get all device software compliances"
result=$(dcld query compliance all-device-software-compliance)
check_response "$result" "\"vid\": $vid_complience"
check_response "$result" "\"pid\": $pid_1"
check_response "$result" "\"cDCertificateId\": \"$cd_certificate_id\""

test_divider

# PKI

echo "Get certificates (ALL)"
result=$(dcld query pki all-certs)
echo $result | jq
check_response "$result" "\"subjectKeyId\": \"$da_root_cert_2_subject_key_id_for_1_4_4\""

echo "Get all certificates by subject (Global)"
result=$(dcld query pki all-subject-certs --subject=$da_root_cert_2_subject_for_1_4_4)
check_response "$result" "$da_root_cert_2_subject_key_id_for_1_4_4"

echo "Get all certificates by SKID (Global)"
result=$(dcld query pki cert --subject-key-id=$da_root_cert_2_subject_key_id_for_1_4_4)
check_response "$result" "\"subjectKeyId\": \"$da_root_cert_2_subject_key_id_for_1_4_4\""
check_response "$result" "$da_root_cert_2_subject_key_id_for_1_4_4"

echo "Get certificate (ALL)"
result=$(dcld query pki cert --subject=$da_root_cert_2_subject_for_1_4_4 --subject-key-id=$da_root_cert_2_subject_key_id_for_1_4_4)
check_response "$result" "\"subject\": \"$da_root_cert_2_subject_for_1_4_4\""
check_response "$result" "\"subjectKeyId\": \"$da_root_cert_2_subject_key_id_for_1_4_4\""

echo "Get certificates (NOC)"
result=$(dcld query pki all-noc-x509-certs)
check_response "$result" "$noc_root_cert_2_subject_key_id_for_1_4_4"

echo "Get certificate (NOC)"
result=$(dcld query pki noc-x509-cert --subject=$noc_root_cert_2_subject_for_1_4_4 --subject-key-id=$noc_root_cert_2_subject_key_id_for_1_4_4)
check_response "$result" "\"subject\": \"$noc_root_cert_2_subject_for_1_4_4\""
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_2_subject_key_id_for_1_4_4\""

echo "Get certificates (DA)"
result=$(dcld query pki all-x509-certs)
check_response "$result" "\"subjectKeyId\": \"$da_root_cert_2_subject_key_id_for_1_4_4\""

echo "Get certificate (DA)"
result=$(dcld query pki x509-cert --subject=$da_root_cert_2_subject_for_1_4_4 --subject-key-id=$da_root_cert_2_subject_key_id_for_1_4_4)
check_response "$result" "\"subject\": \"$da_root_cert_2_subject_for_1_4_4\""
check_response "$result" "\"subjectKeyId\": \"$da_root_cert_2_subject_key_id_for_1_4_4\""

test_divider

# AUTH

echo "Get account"
result=$(dcld query auth account --address=$user_11_address)
check_response "$result" "\"address\": \"$user_11_address\""

echo "Get all accounts"
result=$(dcld query auth all-accounts)
check_response "$result" "\"address\": \"$user_11_address\""
check_response "$result" "\"address\": \"$user_2_address\""

echo "Get rejected account"
result=$(dcld query auth rejected-account --address=$user_rejected_address)
check_response "$result" "\"address\": \"$user_rejected_address\""

echo "Get all rejected accounts"
result=$(dcld query auth all-rejected-accounts)
check_response "$result" "\"address\": \"$user_rejected_address\""
check_response "$result" "\"address\": \"$user_rejected_1_address\""

test_divider

# Validator Node

echo "Get node"
# FIXME: use proper binary (not dcld but $DCLD_BIN_OLD)
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator all-nodes")
check_response "$result" "\"owner\": \"$validator_address\""

########################################################################################

# Write commands for vendor test account

# MODEL and MODEL_VERSION

echo "Add model vid=$vid_testnet pid=$pid_random"
result=$(echo $passphrase | dcld tx model add-model --vid=$vid_testnet --pid=$pid_random --deviceTypeID=$device_type_id_for_1_4_4 --productName=$product_name_for_1_4_4 --productLabel=$product_label_for_1_4_4 --partNumber=$part_number_for_1_4_4 --from=$vendor_account_testnet --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Add model version vid=$vid_for_1_4_4 pid=$pid_1_for_1_4_4"
result=$(echo $passphrase | dcld tx model add-model-version --vid=$vid_testnet --pid=$pid_1_for_1_4_4 --softwareVersion=$software_version_for_1_4_4 --softwareVersionString=$software_version_string_for_1_4_4 --cdVersionNumber=$cd_version_number_for_1_4_4 --minApplicableSoftwareVersion=$min_applicable_software_version_for_1_4_4 --maxApplicableSoftwareVersion=$max_applicable_software_version_for_1_4_4 --from=$vendor_account_testnet --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

# X509 PKI

echo "Add NOC Root certificate by vendor with VID = $vid_testnet"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-root-cert --certificate="$noc_root_cert_1_path_for_1_4_4" --from $vid_testnet --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Upgrade of TestNet from 1.4.3 to 1.4.4 passed"