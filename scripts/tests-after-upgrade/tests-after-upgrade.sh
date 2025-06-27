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

ENVIRONMENT=${1:-local}
ENV_FILE="scripts/tests-after-upgrade/${ENVIRONMENT}.env"

set -euo pipefail
source integration_tests/cli/common.sh
  source "$ENV_FILE"

if [ "$ENVIRONMENT" != "mainnet" ]; then
  if ! grep -qE '^[[:space:]]*passphrase=' "$ENV_FILE"; then
    echo "Error: 'passphrase' is not defined in $ENV_FILE"
    exit 1
  fi
  if [ -z "${passphrase:-}" ]; then
    echo "Error: 'passphrase' is empty in $ENV_FILE"
    exit 1
  fi
fi

dcld config broadcast-mode sync #TODO

echo "Configure CLI"
dcld config output json
dcld config chain-id $chain_id
dcld config node $node_endpoint

random_four_digit_int() {
  echo $(( RANDOM % 9000 + 1000 ))
}

# UPGRADE

echo "Verify that upgrade is applied"
result=$(dcld query upgrade applied $plan_name)
echo "$result"

test_divider

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

test_divider

# MODEL

echo "Get Model with VID: $vid PID: $pid_for_vid"
result=$(dcld query model get-model --vid=$vid --pid=$pid_for_vid)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_for_vid"

echo "Get model version VID: $vid PID: $pid_for_vid"
result=$(dcld query model model-version --vid=$vid --pid=$pid_for_vid --softwareVersion=$software_version)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_for_vid"
check_response "$result" "\"softwareVersion\": $software_version"

echo "Get all models"
result=$(dcld query model all-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_for_vid"

echo "Get Vendor Models with VID: ${vid}"
result=$(dcld query model vendor-models --vid=$vid)
check_response "$result" "\"pid\": $pid_for_vid"
check_response "$result" "\"pid\": $pid_1_for_vid"

echo "Get all model versions with VID: $vid PID: $pid_for_vid"
result=$(dcld query model all-model-versions --vid=$vid --pid=$pid_for_vid)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_for_vid"
# # check_response "$result" "\"softwareVersions\": [$software_version]" TODO: fix

test_divider

# COMPLIANCE

echo "Get certified model vid=$vid_2 pid=$pid_for_vid_2"
result=$(dcld query compliance certified-model --vid=$vid_2 --pid=$pid_for_vid_2 --softwareVersion=$software_version_for_vid_2 --certificationType=$certification_type)
check_response "$result" "\"value\": true"
check_response "$result" "\"vid\": $vid_2"
check_response "$result" "\"pid\": $pid_for_vid_2"
check_response "$result" "\"softwareVersion\": $software_version_for_vid_2"
check_response "$result" "\"certificationType\": \"$certification_type\""

echo "Get all certified models"
result=$(dcld query compliance all-certified-models)
check_response "$result" "\"vid\": $vid_2"
check_response "$result" "\"pid\": $pid_for_vid_2"

if [ "$ENVIRONMENT" != "mainnet" ]; then
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
fi

echo "Get compliance-info model with VID: $vid_2 PID: $pid_for_vid_2"
result=$(dcld query compliance compliance-info --vid=$vid_2 --pid=$pid_for_vid_2 --softwareVersion=$software_version_for_vid_2 --certificationType=$certification_type)
check_response "$result" "\"vid\": $vid_2"
check_response "$result" "\"pid\": $pid_for_vid_2"
check_response "$result" "\"softwareVersion\": $software_version_for_vid_2"
check_response "$result" "\"certificationType\": \"$certification_type\""

echo "Get device software compliance cDCertificateId=$cd_certificate_id"
result=$(dcld query compliance device-software-compliance --cdCertificateId=$cd_certificate_id)
check_response "$result" "\"vid\": $vid_complience"
check_response "$result" "\"pid\": $pid_complience"
check_response "$result" "\"cDCertificateId\": \"$cd_certificate_id\""

echo "Get all device software compliances"
result=$(dcld query compliance all-device-software-compliance)
check_response "$result" "\"vid\": $vid_complience"
check_response "$result" "\"pid\": $pid_complience"
check_response "$result" "\"cDCertificateId\": \"$cd_certificate_id\""

test_divider

# AUTH

echo "Get account"
result=$(dcld query auth account --address=$user_address)
check_response "$result" "\"address\": \"$user_address\""

echo "Get all accounts"
result=$(dcld query auth all-accounts)
check_response "$result" "\"address\": \"$user_address\""
check_response "$result" "\"address\": \"$user_address_1\""

if [ "$ENVIRONMENT" != "mainnet" ]; then

echo "Get rejected account"
result=$(dcld query auth rejected-account --address=$user_rejected_address)
check_response "$result" "\"address\": \"$user_rejected_address\""

echo "Get all rejected accounts"
result=$(dcld query auth all-rejected-accounts)
check_response "$result" "\"address\": \"$user_rejected_address\""
  
fi

test_divider

# PKI

echo "Get certificates (ALL)"
result=$(dcld query pki all-certs)
check_response "$result" "\"subject\": \"$da_root_cert_subject\""

echo "Get all certificates by subject (Global)"
result=$(dcld query pki all-subject-certs --subject=$da_root_cert_subject)
check_response "$result" "\"subject\": \"$da_root_cert_subject\""

echo "Get all certificates by SKID (Global)"
result=$(dcld query pki cert --subject-key-id=$da_root_cert_subject_key_id)
check_response "$result" "\"subject\": \"$da_root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$da_root_cert_subject_key_id\""

echo "Get certificate (ALL)"
result=$(dcld query pki cert --subject=$da_root_cert_subject --subject-key-id=$da_root_cert_subject_key_id)
check_response "$result" "\"subject\": \"$da_root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$da_root_cert_subject_key_id\""

if [ "$ENVIRONMENT" != "mainnet" ]; then

echo "Get certificates (NOC)"
result=$(dcld query pki all-noc-x509-certs)
check_response "$result" "\"subject\": \"$noc_root_cert_subject\""

echo "Get certificate (NOC)"
result=$(dcld query pki noc-x509-cert --subject=$noc_root_cert_subject --subject-key-id=$noc_root_cert_subject_key_id)
check_response "$result" "\"subject\": \"$noc_root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_subject_key_id\""

fi

echo "Get certificates (DA)"
result=$(dcld query pki all-x509-certs)
check_response "$result" "\"subject\": \"$da_root_cert_subject\""

echo "Get certificate (DA)"
result=$(dcld query pki x509-cert --subject=$da_root_cert_subject --subject-key-id=$da_root_cert_subject_key_id)
check_response "$result" "\"subject\": \"$da_root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$da_root_cert_subject_key_id\""

test_divider

# Validator Node

echo "Get validator"
result=$(dcld query validator node --address=$validator_address)
check_response "$result" "\"owner\": \"$validator_address\""

echo "Get all validators"
result=$(dcld query validator all-nodes)
check_response "$result" "\"owner\": \"$validator_address\""

test_divider

########################################################################################

# Write commands for vendor test account

if [ "$ENVIRONMENT" = "local" ]; then
  :
  # create_new_vendor_account $vendor_account $vid_vendor
fi

if [ "$ENVIRONMENT" != "mainnet" ]; then

echo "Use keys for a $vendor_account"
result="$(echo "$passphrase" | dcld keys add "$vendor_account" --recover)"
  # echo "$(dcld keys list)" # for debug

# MODEL and MODEL_VERSION

pid_random=$(random_four_digit_int)
echo "Add model vid=$vid_vendor pid=$pid_random"
result=$(dcld tx model add-model --vid=$vid_vendor --pid=$pid_random --deviceTypeID=$device_type_id --productName=$product_name --from=$vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Add model version vid=$vid_vendor pid=$pid_random"
result=$(dcld tx model add-model-version --vid=$vid_vendor --pid=$pid_random --softwareVersion=$software_version --softwareVersionString="$software_version_string" --cdVersionNumber=$cd_version_number --minApplicableSoftwareVersion=$min_applicable_software_version --maxApplicableSoftwareVersion=$max_applicable_software_version --from=$vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

# X509 PKI

echo "Add NOC Root certificate by vendor with VID = $vendor_account"
result=$(dcld tx pki add-noc-x509-root-cert --certificate="$noc_root_cert" --from $vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"
echo "Teardown: delete the added NOC Root certificate"
result=$(dcld tx pki remove-noc-x509-root-cert --subject="$subject" --subject-key-id="$subject_key_id" --from $vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider
fi

echo "Upgrade of TestNet from 1.4.3 to 1.4.4 passed"