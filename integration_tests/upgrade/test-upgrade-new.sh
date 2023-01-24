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

# Preparation

# constants
trustee_account_1="jack"
trustee_account_2="alice"
vendor_account="van"

plan_name="test-upgrade"

original_model_module_version=1
updated_model_module_version=2

vid=1
pid_1=1
pid_2=2
device_type_id=12345
product_name="ProductName"
product_label="ProductLabel"
part_number="RCU2205A"
software_version=1
software_version_string="1.0"
cd_version_number=312
min_applicable_software_version=1
max_applicable_software_version=1000

certification_type="zigbee"
certification_date="2020-01-01T00:00:00Z"
provisional_date="2019-12-12T00:00:00Z"
cd_certificate_id="15DEXF"

root_cert_path="integration_tests/constants/root_cert"
root_cert_subject="MDQxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMRAwDgYDVQQKDAdyb290LWNh"
root_cert_subject_key_id="5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB"

test_root_cert_path="integration_tests/constants/test_root_cert"
test_root_cert_subject="MDAxGDAWBgNVBAMMD01hdHRlciBUZXN0IFBBQTEUMBIGCisGAQQBgqJ8AgEMBDEyNUQ="
test_root_cert_subject_key_id="E2:90:8D:36:9C:3C:A3:C1:13:BB:09:E2:4D:C1:CC:C5:A6:66:91:D4"

google_root_cert_path="integration_tests/constants/google_root_cert"
google_root_cert_subject="MEsxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZHb29nbGUxFTATBgNVBAMMDE1hdHRlciBQQUEgMTEUMBIGCisGAQQBgqJ8AgEMBDYwMDY="
google_root_cert_subject_key_id="B0:00:56:81:B8:88:62:89:62:80:E1:21:18:A1:A8:BE:09:DE:93:21"

intermediate_cert_path="integration_tests/constants/intermediate_cert"

vendor_name="VendorName"
company_legal_name="LegalCompanyName"
company_preferred_name="CompanyPreferredName"
vendor_landing_page_url="https://www.example.com"

random_string user_1
echo "$user_1 generates keys"
cmd="(echo $passphrase; echo $passphrase) | dcld keys add $user_1"
result="$(bash -c "$cmd")"
user_1_address=$(echo $passphrase | dcld keys show $user_1 -a)
user_1_pubkey=$(echo $passphrase | dcld keys show $user_1 -p)

random_string user_2
echo "$user_2 generates keys"
cmd="(echo $passphrase; echo $passphrase) | dcld keys add $user_2"
result="$(bash -c "$cmd")"
user_2_address=$(echo $passphrase | dcld keys show $user_2 -a)
user_2_pubkey=$(echo $passphrase | dcld keys show $user_2 -p)

random_string user_3
echo "$user_3 generates keys"
cmd="(echo $passphrase; echo $passphrase) | dcld keys add $user_3"
result="$(bash -c "$cmd")"
user_3_address=$(echo $passphrase | dcld keys show $user_3 -a)
user_3_pubkey=$(echo $passphrase | dcld keys show $user_3 -p)

echo "Create Vendor account $vendor_account"
create_new_vendor_account $vendor_account $vid

echo "Create Vendor account $vendor_account"
create_new_account certification_center_account "CertificationCenter"

test_divider

# Body

get_height current_height
echo "Current height is $current_height"

plan_height=$(expr $current_height \* 20)

test_divider

echo "Propose upgrade $plan_name at height $plan_height"
result=$(echo $passphrase | dcld tx dclupgrade propose-upgrade --name $plan_name --upgrade-height $plan_height --from $trustee_account_1 --yes)
echo "$result"
check_response "$result" "\"code\": 0"

test_divider

echo "Approve upgrade $plan_name"
result=$(echo $passphrase | dcld tx dclupgrade approve-upgrade --name $plan_name --from $trustee_account_2 --yes)
echo "$result"
check_response "$result" "\"code\": 0"

test_divider

# send all ledger update transactions

# VENDOR_INFO
echo "Add vendor $vendor_name"
result=$(echo $passphrase | dcld tx vendorinfo add-vendor --vid=$vid --vendorName=$vendor_name --companyLegalName=$company_legal_name --companyPreferredName=$company_preferred_name --vendorLandingPageURL=$vendor_landing_page_url --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"

test_divider

# MODEL and MODEL_VERSION

echo "Add model vid=$vid pid=$pid_1"
result=$(echo $passphrase | dcld tx model add-model --vid=$vid --pid=$pid_1 --deviceTypeID=$device_type_id --productName=$product_name --productLabel=$product_label --partNumber=$part_number --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Add model version vid=$vid pid=$pid_1"
result=$(echo $passphrase | dcld tx model add-model-version --vid=$vid --pid=$pid_1 --softwareVersion=$software_version --softwareVersionString=$software_version_string --cdVersionNumber=$cd_version_number --minApplicableSoftwareVersion=$min_applicable_software_version --maxApplicableSoftwareVersion=$max_applicable_software_version --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Add model vid=$vid pid=$pid_2"
result=$(echo $passphrase | dcld tx model add-model --vid=$vid --pid=$pid_2 --deviceTypeID=$device_type_id --productName=$product_name --productLabel=$product_label --partNumber=$part_number --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Add model version vid=$vid pid=$pid_2"
result=$(echo $passphrase | dcld tx model add-model-version --vid=$vid --pid=$pid_2 --softwareVersion=$software_version --softwareVersionString=$software_version_string --cdVersionNumber=$cd_version_number --minApplicableSoftwareVersion=$min_applicable_software_version --maxApplicableSoftwareVersion=$max_applicable_software_version --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"

test_divider

# CERTIFY_DEVICE_COMPLIANCE

echo "Certify model vid=$vid pid=$pid_1"
result=$(echo $passphrase | dcld tx compliance certify-model --vid=$vid --pid=$pid_1 --softwareVersion=$software_version --softwareVersionString=$software_version_string  --certificationType=$certification_type --certificationDate=$certification_date --cdCertificateId=$cd_certificate_id --from=$certification_center_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Provision model vid=$vid pid=$pid_2"
result=$(echo $passphrase | dcld tx compliance provision-model --vid=$vid --pid=$pid_2 --softwareVersion=$software_version --softwareVersionString=$software_version_string --certificationType=$certification_type --provisionalDate=$provisional_date --cdCertificateId=$cd_certificate_id --from=$certification_center_account --yes)
check_response "$result" "\"code\": 0"

test_divider

# X509 PKI

echo "Propose add root_certificate"
result=$(echo $passphrase | dcld tx pki propose-add-x509-root-cert --certificate="$root_cert_path" --from=$trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add root_certificate"
result=$(echo $passphrase | dcld tx pki approve-add-x509-root-cert --subject="$root_cert_subject" --subject-key-id=$root_cert_subject_key_id --from=$trustee_account_2 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Propose add test_root_certificate"
result=$(echo $passphrase | dcld tx pki propose-add-x509-root-cert --certificate="$test_root_cert_path" --from=$trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add test_root_certificate"
result=$(echo $passphrase | dcld tx pki approve-add-x509-root-cert --subject="$test_root_cert_subject" --subject-key-id=$test_root_cert_subject_key_id --from=$trustee_account_2 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Propose add google_root_certificate"
result=$(echo $passphrase | dcld tx pki propose-add-x509-root-cert --certificate="$google_root_cert_path" --from=$trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Reject add google_root_certificate"
result=$(echo $passphrase | dcld tx pki reject-add-x509-root-cert --subject="$google_root_cert_subject" --subject-key-id=$google_root_cert_subject_key_id --from=$trustee_account_2 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Add intermediate_cert"
result=$(echo $passphrase | dcld tx pki add-x509-cert --certificate="$intermediate_cert_path" --from=$trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Propose revoke root_certificate"
result=$(echo "$passphrase" | dcld tx pki propose-revoke-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from="$trustee_account_1" --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve revoke root_certificate"
result=$(echo "$passphrase" | dcld tx pki approve-revoke-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from="$trustee_account_2" --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Propose revoke test_root_certificate"
result=$(echo $passphrase | dcld tx pki propose-revoke-x509-root-cert --subject="$test_root_cert_subject" --subject-key-id="$test_root_cert_subject_key_id" --from $trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

# AUTH

echo "Propose add account $user_1_address"
result=$(echo $passphrase | dcld tx auth propose-add-account --address="$user_1_address" --pubkey="$user_1_pubkey" --roles="CertificationCenter" --from="$trustee_account_1" --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_1_address"
result=$(dcld tx auth approve-add-account --address="$user_1_address" --from=$trustee_account_2 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Propose add account $user_2_address"
result=$(echo $passphrase | dcld tx auth propose-add-account --address="$user_2_address" --pubkey=$user_2_pubkey --roles=CertificationCenter --from=$trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_2_address"
result=$(echo $passphrase | dcld tx auth approve-add-account --address="$user_2_address" --from=$trustee_account_2 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Propose add account $user_3_address"
result=$(echo $passphrase | dcld tx auth propose-add-account --address="$user_3_address" --pubkey=$user_3_pubkey --roles=CertificationCenter --from=$trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_3_address"
result=$(echo $passphrase | dcld tx auth reject-add-account --address="$user_3_address" --from=$trustee_account_2 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Propose revoke account $user_1_address"
result=$(echo $passphrase | dcld tx auth propose-revoke-account --address="$user_1_address" --from=$trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve revoke account $user_1_address"
result=$(echo $passphrase | dcld tx auth approve-revoke-account --address="$user_1_address" --from=$trustee_account_2 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Propose revoke account $user_2_address"
result=$(echo $passphrase | dcld tx auth propose-revoke-account --address="$user_2_address" --from=$trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

# VALIDATOR_NODE

# dcld tx validator add-node --pubkey=<protobuf JSON encoded> --moniker=<string> --from=<account>
# dcld tx validator disable-node --from=<account>

# dcld tx validator add-node --pubkey=<protobuf JSON encoded> --moniker=<string> --from=<account>
# dcld tx validator propose-disable-node --address=<validator address> --from=<account>

test_divider

echo "Wait for block height to become greater than upgrade $plan_name plan height"
wait_for_height $(expr $plan_height + 1) 300 outage-safe

test_divider

echo "Verify that no upgrade has been scheduled anymore"
result=$(dcld query upgrade plan 2>&1) || true
check_response_and_report "$result" "no upgrade scheduled" raw

test_divider

echo "PASSED"
