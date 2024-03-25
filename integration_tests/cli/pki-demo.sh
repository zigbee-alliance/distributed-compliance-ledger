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

root_cert_subject="MDQxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMRAwDgYDVQQKDAdyb290LWNh"
root_cert_subject_key_id="5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB"
root_cert_serial_number="442314047376310867378175982234956458728610743315"
root_cert_subject_as_text="O=root-ca,ST=some-state,C=AU"
vid=1

intermediate_cert_subject="MDwxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMRgwFgYDVQQKDA9pbnRlcm1lZGlhdGUtY2E="
intermediate_cert_subject_key_id="4E:3B:73:F4:70:4D:C2:98:0D:DB:C8:5A:5F:02:3B:BF:86:25:56:2B"
intermediate_cert_serial_number="169917617234879872371588777545667947720450185023"
intermediate_cert_subject_as_text="O=intermediate-ca,ST=some-state,C=AU"

leaf_cert_subject="MDExCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMQ0wCwYDVQQKDARsZWFm"
leaf_cert_subject_key_id="30:F4:65:75:14:20:B2:AF:3D:14:71:17:AC:49:90:93:3E:24:A0:1F"
leaf_cert_serial_number="143290473708569835418599774898811724528308722063"
leaf_cert_subject_as_text="O=leaf,ST=some-state,C=AU"

# Preparation of Actors

trustee_account="jack"
second_trustee_account="alice"
third_trustee_account="bob"

trustee_account_address=$(echo $passphrase | dcld keys show jack -a)
second_trustee_account_address=$(echo $passphrase | dcld keys show alice -a)
third_trustee_account_address=$(echo $passphrase | dcld keys show bob -a)

vendor_account=vendor_account_$vid
echo "Create Vendor account - $vendor_account"
create_new_vendor_account $vendor_account $vid

vendor_account_65522=vendor_account_65522
echo "Create Vendor account - $vendor_account_65522"
create_new_vendor_account $vendor_account_65522 65522

echo "Create regular account"
create_new_account user_account "CertificationCenter"
test_divider

# Body

# 1. QUERY ALL (EMPTY)

echo "1. QUERY ALL EMPTY"
test_divider

echo "Request approved certificate must be empty"
result=$(dcld query pki x509-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$root_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$root_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$root_cert_subject_as_text\""
echo $result | jq

echo "Request all approved certificates must be empty"
result=$(dcld query pki all-x509-certs)
check_response "$result" "\[\]"
response_does_not_contain "$result" "\"subject\": \"$root_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$root_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$root_cert_subject_as_text\""
echo $result | jq

test_divider

echo "Request proposed Root certificate must be empty"
result=$(dcld query pki proposed-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$root_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$root_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$root_cert_subject_as_text\""
echo $result | jq

test_divider

echo "Request all proposed Root certificates must be empty"
result=$(dcld query pki all-proposed-x509-root-certs)
check_response "$result" "\[\]"
response_does_not_contain "$result" "\"subject\": \"$root_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
echo $result | jq

test_divider

echo "Request revoked certificate must be empty"
result=$(dcld query pki revoked-x509-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$root_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
echo $result | jq

echo "Request all revoked certificates must be empty"
result=$(dcld query pki all-revoked-x509-certs)
check_response "$result" "\[\]"
response_does_not_contain "$result" "\"subject\": \"$root_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$root_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$root_cert_subject_as_text\""
echo $result | jq

echo "Request all certificates by subject must be empty"
result=$(dcld query pki all-subject-x509-certs --subject="$root_cert_subject")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"$root_cert_subject\""
response_does_not_contain "$result" "\"$root_cert_subject_key_id\""
echo $result | jq

echo "Request all certificates by subjectKeyId must be empty"
result=$(dcld query pki x509-cert --subject-key-id="$root_cert_subject_key_id")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"$root_cert_subject\""
response_does_not_contain "$result" "\"$root_cert_subject_key_id\""
echo $result | jq

test_divider

echo "Request all approved root certificates must be empty"
result=$(dcld query pki all-x509-root-certs)
check_response "$result" "\[\]"
response_does_not_contain "$result" "\"$root_cert_subject\""
response_does_not_contain "$result" "\"$root_cert_subject_key_id\""
echo $result | jq

test_divider

echo "Request all revoked root certificates must be empty"
result=$(dcld query pki all-revoked-x509-root-certs)
check_response "$result" "\[\]"
response_does_not_contain "$result" "\"$root_cert_subject\""
response_does_not_contain "$result" "\"$root_cert_subject_key_id\""
echo $result | jq

test_divider

echo "Request root certificate proposed to revoke must be empty"
result=$(dcld query pki proposed-x509-root-cert-to-revoke --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"$root_cert_subject\""
response_does_not_contain "$result" "\"$root_cert_subject_key_id\""
echo $result | jq

test_divider

echo "Request all root certificates proposed to revoke must be empty"
result=$(dcld query pki all-proposed-x509-root-certs-to-revoke)
check_response "$result" "\[\]"
response_does_not_contain "$result" "\"$root_cert_subject\""
response_does_not_contain "$result" "\"$root_cert_subject_key_id\""
echo $result | jq

test_divider

echo "Request all child certificates must be empty"
result=$(dcld query pki all-child-x509-certs --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"$root_cert_subject\""
response_does_not_contain "$result" "\"$root_cert_subject_key_id\""
echo $result | jq

test_divider


# 2. PROPOSE ROOT

echo "2. PROPOSE ROOT CERT"
test_divider

echo "$user_account (Not Trustee) propose Root certificate"
root_path="integration_tests/constants/root_cert"
cert_schema_version_1=1
schema_version_2=2
result=$(echo "$passphrase" | dcld tx pki propose-add-x509-root-cert --certificate="$root_path" --from $user_account --vid $vid --yes)
response_does_not_contain "$result" "\"code\": 0"

echo "$trustee_account (Trustee) propose Root certificate"
result=$(echo "$passphrase" | dcld tx pki propose-add-x509-root-cert --certificate="$root_path" --certificate-schema-version=$cert_schema_version_1 --schemaVersion=$schema_version_2 --from $trustee_account   --vid $vid --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Request all proposed Root certificates - There should be no approvals"
result=$(dcld query pki all-proposed-x509-root-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
check_response "$result" "\"certSchemaVersion\": $cert_schema_version_1"
check_response "$result" "\"schemaVersion\": $schema_version_2"

test_divider


echo "Request proposed Root certificate - there should be no Approval"
result=$(dcld query pki proposed-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$root_cert_serial_number\""
check_response "$result" "\"subjectAsText\": \"$root_cert_subject_as_text\""
check_response "$result" "\"vid\": $vid"
echo $result | jq

test_divider

echo "Request all approved certificates must be empty"
result=$(dcld query pki all-x509-certs)
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$root_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$root_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$root_cert_subject_as_text\""


test_divider

echo "Approved certificate must be empty"
result=$(dcld query pki x509-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
echo $result | jq
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$root_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$root_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$root_cert_subject_as_text\""


test_divider

echo "Request all revoked certificates must be empty"
result=$(dcld query pki all-revoked-x509-certs)
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$root_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$root_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$root_cert_subject_as_text\""




echo "Request all approved root certificates must be empty"
result=$(dcld query pki all-x509-root-certs)
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$root_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""


test_divider

echo "Request all revoked root certificates must be empty"
result=$(dcld query pki all-revoked-x509-root-certs)
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$root_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""


test_divider

echo "Request all certificates by subject must be empty"
result=$(dcld query pki all-subject-x509-certs --subject="$root_cert_subject")
echo $result | jq
response_does_not_contain "$result" "\"$root_cert_subject\""
response_does_not_contain "$result" "\"$root_cert_subject_key_id\""
echo $result | jq

test_divider

# 3. APPROVE ROOT CERT

echo "3. APPROVE ROOT CERT"
test_divider

echo "Certificate must be still in Proposed state but with Approval from $trustee_account. Request proposed Root certificate"
result=$(dcld query pki proposed-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$root_cert_serial_number\""
check_response "$result" "\"address\": \"$trustee_account_address\""
check_response "$result" "\"subjectAsText\": \"$root_cert_subject_as_text\""
check_response "$result" "\"vid\": $vid"
response_does_not_contain "$result" "\"address\": \"$second_trustee_account_address\""
check_response "$result" "[\"$(echo "$passphrase" | dcld keys show jack -a)\"]"


test_divider

echo "Request all approved certificates must be empty, only 1 Trustee has approved so far"
result=$(dcld query pki all-x509-certs)
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$root_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$root_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$root_cert_subject_as_text\""


test_divider

echo "$second_trustee_account (Second Trustee) approves Root certificate"
result=$(echo "$passphrase" | dcld tx pki approve-add-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from $second_trustee_account --yes)
check_response "$result" "\"code\": 0"


test_divider

echo "Get Certificates by subject and subjectKeyId must be Approved and contain 2 approvals. Request Root certificate"
result=$(dcld query pki x509-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$root_cert_serial_number\""
check_response "$result" "\"subjectAsText\": \"$root_cert_subject_as_text\""
check_response "$result" "\"address\": \"$trustee_account_address\""
check_response "$result" "\"address\": \"$second_trustee_account_address\""

echo "Get Certificates by subjectKeyId must be Approved and contain 2 approvals. Request Root certificate"
result=$(dcld query pki x509-cert --subject-key-id="$root_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$root_cert_serial_number\""
check_response "$result" "\"subjectAsText\": \"$root_cert_subject_as_text\""
check_response "$result" "\"address\": \"$trustee_account_address\""
check_response "$result" "\"address\": \"$second_trustee_account_address\""


test_divider

echo "Request all proposed Root certificates must be empty"
result=$(dcld query pki all-proposed-x509-root-certs)
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$root_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$root_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$root_cert_subject_as_text\""


test_divider

echo "Request all approved certificates. It should contain one certificate with 2 approvals"
result=$(dcld query pki all-x509-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
check_response "$result" "\"address\": \"$trustee_account_address\""
check_response "$result" "\"address\": \"$second_trustee_account_address\""


test_divider

echo "Request all approved root certificates."
result=$(dcld query pki all-x509-root-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""



echo "Request all revoked certificates must be empty"
result=$(dcld query pki all-revoked-x509-certs)
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$root_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$root_cert_serial_number\""

# 4. ADD INTERMEDIATE CERT

echo "4. ADD INTERMEDIATE CERT"
test_divider


echo "$vendor_account adds Intermediate certificate"
intermediate_path="integration_tests/constants/intermediate_cert"
result=$(echo "$passphrase" | dcld tx pki add-x509-cert --certificate="$intermediate_path" --certificate-schema-version=$cert_schema_version_1 --schemaVersion=$schema_version_2 --from $vendor_account --yes)
check_response "$result" "\"code\": 0"


test_divider

echo "Request Intermediate certificate by subject and subjectKeyId - There are no approvals for Intermidiate Certificates"
result=$(dcld query pki x509-cert --subject="$intermediate_cert_subject" --subject-key-id="$intermediate_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$intermediate_cert_serial_number\""
check_response "$result" "\"subjectAsText\": \"$intermediate_cert_subject_as_text\""
check_response "$result" "\"schemaVersion\": $cert_schema_version_1"
check_response "$result" "\"schemaVersion\": $schema_version_2"
check_response "$result" "\"approvals\": \\[\\]"

echo "Request Intermediate certificate by subjectKeyId - There are no approvals for Intermidiate Certificates"
result=$(dcld query pki x509-cert --subject-key-id="$intermediate_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$intermediate_cert_serial_number\""
check_response "$result" "\"subjectAsText\": \"$intermediate_cert_subject_as_text\""
check_response "$result" "\"approvals\": \\[\\]"

test_divider

echo "Request all proposed Root certificates must be empty"
result=$(dcld query pki all-proposed-x509-root-certs)
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$root_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$root_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$root_cert_subject_as_text\""
response_does_not_contain "$result" "\"subject\": \"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$intermediate_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$intermediate_cert_subject_as_text\""


test_divider

echo "Request all approved certificates"
result=$(dcld query pki all-x509-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""


test_divider

echo "Request all approved root certificates"
result=$(dcld query pki all-x509-root-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"subject\": \"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""

test_divider

# 5. ADD LEAF CERT

echo "5. ADD LEAF CERT"
test_divider

echo "$vendor_account add Leaf certificate"
leaf_path="integration_tests/constants/leaf_cert"
schema_version_0=0
result=$(echo "$passphrase" | dcld tx pki add-x509-cert --certificate="$leaf_path" --from $vendor_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Request Leaf certificate by subject and subjectKeyId - There is no approvals on leaf certificate"
result=$(dcld query pki x509-cert --subject="$leaf_cert_subject" --subject-key-id="$leaf_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$leaf_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$leaf_cert_serial_number\""
check_response "$result" "\"subjectAsText\": \"$leaf_cert_subject_as_text\""
check_response "$result" "\"schemaVersion\": $schema_version_0"
check_response "$result" "\"approvals\": \\[\\]"

echo "Request Leaf certificate by subjectKeyId - There is no approvals on leaf certificate"
result=$(dcld query pki x509-cert --subject-key-id="$leaf_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$leaf_cert_serial_number\""
check_response "$result" "\"subjectAsText\": \"$leaf_cert_subject_as_text\""
check_response "$result" "\"approvals\": \\[\\]"


test_divider

# TODO: there is no use case for x509-cert-chain, and it can be tricky, see Slack discussion

# echo "Request certificate chain for Intermediate certificate"
# result=$(dcld query pki x509-cert-chain --subject="$intermediate_cert_subject" --subject-key-id="$intermediate_cert_subject_key_id")
# check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
# check_response "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
# check_response "$result" "\"serialNumber\": \"$intermediate_cert_serial_number\""
# check_response "$result" "\"subject\": \"$root_cert_subject\""
# check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
# check_response "$result" "\"serialNumber\": \"$root_cert_serial_number\""
# echo "$result"

# test_divider


# echo "Request certificate chain for Leaf certificate"
# result=$(dcld query pki x509-cert-chain --subject="$leaf_cert_subject" --subject-key-id="$leaf_cert_subject_key_id")
# check_response "$result" "\"subject\": \"$leaf_cert_subject\""
# check_response "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""
# check_response "$result" "\"serialNumber\": \"$leaf_cert_serial_number\""
# check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
# check_response "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
# check_response "$result" "\"serialNumber\": \"$intermediate_cert_serial_number\""
# check_response "$result" "\"subject\": \"$root_cert_subject\""
# check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
# check_response "$result" "\"serialNumber\": \"$root_cert_serial_number\""
# echo "$result"

# test_divider

echo "Request all proposed Root certificates must be empty"
result=$(dcld query pki all-proposed-x509-root-certs)
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$root_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$root_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$root_cert_subject_as_text\""
response_does_not_contain "$result" "\"subject\": \"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$intermediate_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$intermediate_cert_subject_as_text\""
response_does_not_contain "$result" "\"subject\": \"$leaf_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$leaf_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$leaf_cert_subject_as_text\""

test_divider

echo "Request all approved certificates"
result=$(dcld query pki all-x509-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
check_response "$result" "\"subject\": \"$leaf_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""

test_divider

echo "Request all approved root certificates"
result=$(dcld query pki all-x509-root-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"subject\": \"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
response_does_not_contain "$result" "\"subject\": \"$leaf_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""

test_divider

echo "Request all subject: $root_cert_subject certificates"
result=$(dcld query pki all-subject-x509-certs --subject="$root_cert_subject")
echo $result | jq
check_response "$result" "\"$root_cert_subject\""
check_response "$result" "\"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"$intermediate_cert_subject_key_id\""
response_does_not_contain "$result" "\"$leaf_cert_subject\""
response_does_not_contain "$result" "\"$leaf_cert_subject_key_id\""

test_divider

echo "Request all subject:$leaf_cert_subject certificates"
result=$(dcld query pki all-subject-x509-certs --subject="$leaf_cert_subject")
echo $result | jq
response_does_not_contain "$result" "\"$root_cert_subject\""
response_does_not_contain "$result" "\"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"$intermediate_cert_subject_key_id\""
check_response "$result" "\"$leaf_cert_subject\""
check_response "$result" "\"$leaf_cert_subject_key_id\""

test_divider

echo "Request all intermediate_cert_subject:$intermediate_cert_subject certificates"
result=$(dcld query pki all-subject-x509-certs --subject="$intermediate_cert_subject")
echo $result | jq
response_does_not_contain "$result" "\"$root_cert_subject\""
response_does_not_contain "$result" "\"$root_cert_subject_key_id\""
check_response "$result" "\"$intermediate_cert_subject\""
check_response "$result" "\"$intermediate_cert_subject_key_id\""
response_does_not_contain "$result" "\"$leaf_cert_subject\""
response_does_not_contain "$result" "\"$leaf_cert_subject_key_id\""

test_divider

echo "Request all root certificates proposed to revoke. There should be nothing in the list"
result=$(dcld query pki all-proposed-x509-root-certs-to-revoke)
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$root_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$root_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$root_cert_subject_as_text\""
response_does_not_contain "$result" "\"subject\": \"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$intermediate_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$intermediate_cert_subject_as_text\""
response_does_not_contain "$result" "\"subject\": \"$leaf_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$leaf_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$leaf_cert_subject_as_text\""

test_divider

echo "Request all revoked certificates. There should be nothing in the list"
result=$(dcld query pki all-revoked-x509-certs)
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$root_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$root_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$root_cert_subject_as_text\""
response_does_not_contain "$result" "\"subject\": \"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$intermediate_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$intermediate_cert_subject_as_text\""
response_does_not_contain "$result" "\"subject\": \"$leaf_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$leaf_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$leaf_cert_subject_as_text\""

test_divider

echo "Request all child certificates for root"
result=$(dcld query pki all-child-x509-certs --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
response_does_not_contain "$result" "\"subject\": \"$leaf_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""

test_divider


echo "Request all child certificates for intermediate"
result=$(dcld query pki all-child-x509-certs --subject="$intermediate_cert_subject" --subject-key-id="$intermediate_cert_subject_key_id")
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$root_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
check_response "$result" "\"subject\": \"$leaf_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""

test_divider

echo "Request all child certificates for leaf. There should be no children"
result=$(dcld query pki all-child-x509-certs --subject="$leaf_cert_subject" --subject-key-id="$leaf_cert_subject_key_id")
echo $result | jq
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$root_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"subject\": \"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""

test_divider

# 6. REVOKE INTERMEDIATE (AND HENCE  LEAF) CERTS

echo "6. REVOKE INTERMEDIATE (AND HENCE  LEAF) CERTS - No Approvals needed"
test_divider

echo "Try to revoke the intermediate certificate when sender is not Vendor account"
result=$(echo "$passphrase" | dcld tx pki revoke-x509-cert --subject="$intermediate_cert_subject" --subject-key-id="$intermediate_cert_subject_key_id" --from=$user_account --yes)
check_response "$result" "\"code\": 4"

echo "Try to revoke the intermediate certificate using a vendor account with other VID"
result=$(echo "$passphrase" | dcld tx pki revoke-x509-cert --subject="$intermediate_cert_subject" --subject-key-id="$intermediate_cert_subject_key_id" --from=$vendor_account_65522 --yes)
check_response "$result" "\"code\": 4"

revoke_schema_version_3=3
echo "$vendor_account (Not Trustee) revokes only Intermediate certificate. This must not revoke its child - Leaf certificate."
result=$(echo "$passphrase" | dcld tx pki revoke-x509-cert --subject="$intermediate_cert_subject" --subject-key-id="$intermediate_cert_subject_key_id" --schemaVersion=$revoke_schema_version_3 --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Request all root certificates proposed to revoke"
result=$(dcld query pki all-proposed-x509-root-certs-to-revoke)
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$root_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$root_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$root_cert_subject_as_text\""
response_does_not_contain "$result" "\"subject\": \"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$intermediate_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$intermediate_cert_subject_as_text\""
response_does_not_contain "$result" "\"subject\": \"$leaf_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$leaf_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$leaf_cert_subject_as_text\""


test_divider

echo "Request all revoked certificates"
result=$(dcld query pki all-revoked-x509-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
check_response "$result" "\"schemaVersion\": $revoke_schema_version_3"
response_does_not_contain "$result" "\"subject\": \"$leaf_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""
response_does_not_contain "$result" "\"subject\": \"$root_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""

test_divider

echo "Request all revoked root certificates"
result=$(dcld query pki all-revoked-x509-root-certs)
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$root_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"subject\": \"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
response_does_not_contain "$result" "\"subject\": \"$leaf_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""

test_divider

echo "Request revoked Intermediate certificate"
result=$(dcld query pki revoked-x509-cert --subject="$intermediate_cert_subject" --subject-key-id="$intermediate_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$intermediate_cert_serial_number\""

test_divider

echo "Request revoked Leaf certificate"
result=$(dcld query pki revoked-x509-cert --subject="$leaf_cert_subject" --subject-key-id="$leaf_cert_subject_key_id")
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$leaf_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$leaf_cert_serial_number\""

test_divider

echo "Request all approved certificates"
result=$(dcld query pki all-x509-certs)
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
check_response "$result" "\"subject\": \"$leaf_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""

test_divider

echo "Request all approved root certificates"
result=$(dcld query pki all-x509-root-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"subject\": \"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
response_does_not_contain "$result" "\"subject\": \"$leaf_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""

test_divider

echo "Request all subject certificates"
result=$(dcld query pki all-subject-x509-certs --subject="$leaf_cert_subject")
echo $result | jq
check_response "$result" "\"$leaf_cert_subject\""
check_response "$result" "\"$leaf_cert_subject_key_id\""
response_does_not_contain "$result" "\"$root_cert_subject\""
response_does_not_contain "$result" "\"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"$intermediate_cert_subject_key_id\""

test_divider

echo "Request all subject certificates"
result=$(dcld query pki all-subject-x509-certs --subject="$intermediate_cert_subject")
echo $result | jq
response_does_not_contain "$result" "\"$root_cert_subject\""
response_does_not_contain "$result" "\"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"$intermediate_cert_subject_key_id\""
response_does_not_contain "$result" "\"$leaf_cert_subject\""
response_does_not_contain "$result" "\"$leaf_cert_subject_key_id\""

test_divider

echo "Approved Intermediate certificate must be empty"
result=$(dcld query pki x509-cert --subject="$intermediate_cert_subject" --subject-key-id="$intermediate_cert_subject_key_id")
echo $result | jq
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$intermediate_cert_serial_number\""

test_divider

echo "Approved Leaf certificate must not be empty"
result=$(dcld query pki x509-cert --subject="$leaf_cert_subject" --subject-key-id="$leaf_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$leaf_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$leaf_cert_serial_number\""

test_divider

# 7. PROPOSE REVOCATION OF ROOT CERT

echo "7. PROPOSE REVOCATION OF ROOT CERT"
test_divider

revoke_schema_version_4=4
echo "$trustee_account (Trustee) proposes to revoke only Root certificate(child certificates should not be revoked)"
result=$(echo "$passphrase" | dcld tx pki propose-revoke-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --schemaVersion=$revoke_schema_version_4 --from $trustee_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Request root certificate proposed to revoke and verify that it contains approval from $trustee_account_address"
result=$(dcld query pki proposed-x509-root-cert-to-revoke --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"$root_cert_subject\""
check_response "$result" "\"$root_cert_subject_key_id\""
check_response "$result" "\"address\": \"$trustee_account_address\""

echo "Request all root certificates proposed to revoke"
result=$(dcld query pki all-proposed-x509-root-certs-to-revoke)
echo $result | jq
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
check_response "$result" "\"schemaVersion\": $revoke_schema_version_4"
response_does_not_contain "$result" "\"subject\": \"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
response_does_not_contain "$result" "\"subject\": \"$leaf_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""

test_divider

echo "Request all revoked certificates"
result=$(dcld query pki all-revoked-x509-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
response_does_not_contain "$result" "\"subject\": \"$leaf_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""
response_does_not_contain "$result" "\"subject\": \"$root_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""


test_divider

echo "Request all revoked root certificates"
result=$(dcld query pki all-revoked-x509-root-certs)
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$root_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"subject\": \"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
response_does_not_contain "$result" "\"subject\": \"$leaf_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""


test_divider

echo "Request Root certificate proposed to revoke, it should have one approval from $trustee_account_address"
result=$(dcld query pki proposed-x509-root-cert-to-revoke --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
check_response "$result" "[\"$(echo "$passphrase" | dcld keys show jack -a)\"]"  
check_response "$result" "\"address\": \"$trustee_account_address\""


test_divider

echo "Request all approved certificates"
result=$(dcld query pki all-x509-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
check_response "$result" "\"subject\": \"$leaf_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""
response_does_not_contain "$result" "\"subject\": \"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""


test_divider

echo "Request all approved root certificates"
result=$(dcld query pki all-x509-root-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"subject\": \"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
response_does_not_contain "$result" "\"subject\": \"$leaf_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""


test_divider

echo "Request all subject certificates"
result=$(dcld query pki all-subject-x509-certs --subject="$root_cert_subject")
echo $result | jq
check_response "$result" "\"$root_cert_subject\""
check_response "$result" "\"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"$leaf_cert_subject\""
response_does_not_contain "$result" "\"$leaf_cert_subject_key_id\""
response_does_not_contain "$result" "\"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"$intermediate_cert_subject_key_id\""

test_divider

# 8. APPROVE REVOCATION OF ROOT CERT

echo "8. APPROVE REVOCATION OF ROOT CERT"
test_divider


echo "$second_trustee_account (Second Trustee) approves to revoke Root certificate"
result=$(echo "$passphrase" | dcld tx pki approve-revoke-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from $second_trustee_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Request all root certificates proposed to revoke. Nothing left in list as the certificates are revoked"
result=$(dcld query pki all-proposed-x509-root-certs-to-revoke)
response_does_not_contain "$result" "\"subject\": \"$root_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"subject\": \"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
response_does_not_contain "$result" "\"subject\": \"$leaf_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""
echo $result | jq

test_divider

echo "Request all revoked certificates should contain approvals from both trustees"
result=$(dcld query pki all-revoked-x509-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
check_response "$result" "\"address\": \"$trustee_account_address\""
check_response "$result" "\"address\": \"$second_trustee_account_address\""
check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
response_does_not_contain "$result" "\"subject\": \"$leaf_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""


test_divider

echo "Request all revoked root certificates"
result=$(dcld query pki all-revoked-x509-root-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"subject\": \"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
response_does_not_contain "$result" "\"subject\": \"$leaf_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""


test_divider

echo "Request revoked Root certificate and also check for approvals from both Trustees"
result=$(dcld query pki revoked-x509-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$root_cert_serial_number\""
response_does_not_contain "$result" "\"subject\": \"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
response_does_not_contain "$result" "\"subject\": \"$leaf_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""
check_response "$result" "\"address\": \"$trustee_account_address\""
check_response "$result" "\"address\": \"$second_trustee_account_address\""


test_divider

echo "Request all approved certificates must not contain root certificate"
result=$(dcld query pki all-x509-certs)
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$root_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"subject\": \"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
check_response "$result" "\"subject\": \"$leaf_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""


echo "Request all approved root certificates must be empty"
result=$(dcld query pki all-x509-root-certs)
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$root_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"subject\": \"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
response_does_not_contain "$result" "\"subject\": \"$leaf_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""


test_divider

echo "Approved Intermediate certificate must be empty"
result=$(dcld query pki x509-cert --subject="$intermediate_cert_subject" --subject-key-id="$intermediate_cert_subject_key_id")
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$intermediate_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$intermediate_cert_subject_as_text\""


test_divider

echo "Approved Leaf certificate must not be empty"
result=$(dcld query pki x509-cert --subject="$leaf_cert_subject" --subject-key-id="$leaf_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$leaf_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$leaf_cert_serial_number\""
check_response "$result" "\"subjectAsText\": \"$leaf_cert_subject_as_text\""


test_divider

echo "Approved Root certificate must be empty"
result=$(dcld query pki x509-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
echo $result | jq
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$root_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$root_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$root_cert_subject_as_text\""

test_divider

echo "Request all subject certificates must be empty"
result=$(dcld query pki all-subject-x509-certs --subject="$root_cert_subject")
echo $result | jq
response_does_not_contain "$result" "\"$root_cert_subject\""
response_does_not_contain "$result" "\"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"$intermediate_cert_subject_key_id\""
response_does_not_contain "$result" "\"$leaf_cert_subject\""
response_does_not_contain "$result" "\"$leaf_cert_subject_key_id\""

# CHECK GOOGLE ROOT CERTIFICATE WHICH INCLUDES VID

google_cert_subject="MEsxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZHb29nbGUxFTATBgNVBAMMDE1hdHRlciBQQUEgMTEUMBIGCisGAQQBgqJ8AgEMBDYwMDY="
google_cert_subject_key_id="B0:00:56:81:B8:88:62:89:62:80:E1:21:18:A1:A8:BE:09:DE:93:21"
google_cert_serial_number="1"
google_cert_subject_as_text="CN=Matter PAA 1,O=Google,C=US,vid=0x6006"
google_cert_vid=24582

# 9. QUERY ALL (EMPTY)

echo "9. QUERY ALL EMPTY"
test_divider

echo "Request approved certificate must be empty"
result=$(dcld query pki x509-cert --subject="$google_cert_subject" --subject-key-id="$google_cert_subject_key_id")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$google_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$google_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$google_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$google_cert_subject_as_text\""
echo $result | jq

echo "Request all approved certificates must not contain google certification"
result=$(dcld query pki all-x509-certs)
response_does_not_contain "$result" "\"subject\": \"$google_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$google_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$google_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$google_cert_subject_as_text\""
echo $result | jq

test_divider

echo "Request proposed Root certificate must be empty"
result=$(dcld query pki proposed-x509-root-cert --subject="$google_cert_subject" --subject-key-id="$google_cert_subject_key_id")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$google_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$google_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$google_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$google_cert_subject_as_text\""
echo $result | jq

test_divider

echo "Request all proposed Root certificates must be empty"
result=$(dcld query pki all-proposed-x509-root-certs)
check_response "$result" "\[\]"
response_does_not_contain "$result" "\"subject\": \"$google_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$google_cert_subject_key_id\""
echo $result | jq

test_divider

echo "Request revoked certificate must be empty"
result=$(dcld query pki revoked-x509-cert --subject="$google_cert_subject" --subject-key-id="$google_cert_subject_key_id")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$google_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$google_cert_subject_key_id\""
echo $result | jq

echo "Request all revoked certificates must not contain google certification"
result=$(dcld query pki all-revoked-x509-certs)
response_does_not_contain "$result" "\"$google_cert_subject\""
response_does_not_contain "$result" "\"$google_cert_subject_key_id\""
echo $result | jq

echo "Request all certificates by subject must not contain google certification"
result=$(dcld query pki all-subject-x509-certs --subject="$google_cert_subject")
response_does_not_contain "$result" "\"$google_cert_subject\""
response_does_not_contain "$result" "\"$google_cert_subject_key_id\""
echo $result | jq

echo "Request all approved root certificates must not contain google certification"
result=$(dcld query pki all-x509-root-certs)
response_does_not_contain "$result" "\"$google_cert_subject\""
response_does_not_contain "$result" "\"$google_cert_subject_key_id\""
echo $result | jq

test_divider

echo "Request all revoked root certificates must not contain google certification"
result=$(dcld query pki all-revoked-x509-root-certs)
response_does_not_contain "$result" "\"$google_cert_subject\""
response_does_not_contain "$result" "\"$google_cert_subject_key_id\""
echo $result | jq

test_divider

echo "Request google root certificate proposed to revoke must be empty"
result=$(dcld query pki proposed-x509-root-cert-to-revoke --subject="$google_cert_subject" --subject-key-id="$google_cert_subject_key_id")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"$google_cert_subject\""
response_does_not_contain "$result" "\"$google_cert_subject_key_id\""
echo $result | jq

test_divider

echo "Request all root certificates proposed to revoke must be empty"
result=$(dcld query pki all-proposed-x509-root-certs-to-revoke)
check_response "$result" "\[\]"
response_does_not_contain "$result" "\"$google_cert_subject\""
response_does_not_contain "$result" "\"$google_cert_subject_key_id\""
echo $result | jq

test_divider

echo "Request all child certificates must be empty"
result=$(dcld query pki all-child-x509-certs --subject="$google_cert_subject" --subject-key-id="$google_cert_subject_key_id")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"$google_cert_subject\""
response_does_not_contain "$result" "\"$google_cert_subject_key_id\""
echo $result | jq

test_divider

# 10. PROPOSE GOOGLE ROOT

echo "10. PROPOSE GOOGLE ROOT CERT"
test_divider

cert_schema_version_0=0
echo "$user_account (Not Trustee) propose Root certificate"
google_root_path="integration_tests/constants/google_root_cert"
result=$(echo "$passphrase" | dcld tx pki propose-add-x509-root-cert --certificate="$google_root_path" --from $user_account --vid=$google_cert_vid --yes)
response_does_not_contain "$result" "\"code\": 0"

echo "$trustee_account (Trustee) propose Root certificate"
google_root_path="integration_tests/constants/google_root_cert"
result=$(echo "$passphrase" | dcld tx pki propose-add-x509-root-cert --certificate="$google_root_path" --from $trustee_account --vid=$google_cert_vid --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Request proposed Root certificate - there should be Approval"
result=$(dcld query pki proposed-x509-root-cert --subject="$google_cert_subject" --subject-key-id="$google_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$google_cert_subject\""
check_response "$result" "\"address\": \"$trustee_account_address\""
check_response "$result" "\"subjectKeyId\": \"$google_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$google_cert_serial_number\""
check_response "$result" "\"subjectAsText\": \"$google_cert_subject_as_text\""
check_response "$result" "\"certSchemaVersion\": $cert_schema_version_0"
check_response "$result" "\"schemaVersion\": $schema_version_0"
check_response "$result" "\"vid\": $google_cert_vid"
echo $result | jq

test_divider

echo "Request all approved certificates must be empty"
result=$(dcld query pki all-x509-certs)
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$google_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$google_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$google_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$google_cert_subject_as_text\""

test_divider

echo "Approved certificate must be empty"
result=$(dcld query pki x509-cert --subject="$google_cert_subject" --subject-key-id="$google_cert_subject_key_id")
echo $result | jq
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$google_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$google_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$google_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$google_cert_subject_as_text\""

test_divider

echo "Request all revoked certificates must not contain google certification"
result=$(dcld query pki all-revoked-x509-certs)
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$google_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$google_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$google_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$google_cert_subject_as_text\""




echo "Request all approved root certificates must be empty"
result=$(dcld query pki all-x509-root-certs)
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$google_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$google_cert_subject_key_id\""

test_divider

echo "Request all revoked root certificates must not contain google certification"
result=$(dcld query pki all-revoked-x509-root-certs)
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$google_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$google_cert_subject_key_id\""


test_divider

echo "Request all certificates by subject must be empty"
result=$(dcld query pki all-subject-x509-certs --subject="$google_cert_subject")
echo $result | jq
response_does_not_contain "$result" "\"$google_cert_subject\""
response_does_not_contain "$result" "\"$google_cert_subject_key_id\""
echo $result | jq

test_divider

# 11. APPROVE GOOGLE ROOT CERT

echo "11. APPROVE GOOGLE ROOT CERT"
test_divider

echo "Certificate must be still in Proposed state but with Approval from $trustee_account. Request proposed Root certificate"
result=$(dcld query pki proposed-x509-root-cert --subject="$google_cert_subject" --subject-key-id="$google_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$google_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$google_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$google_cert_serial_number\""
check_response "$result" "\"subjectAsText\": \"$google_cert_subject_as_text\""
check_response "$result" "\"vid\": $google_cert_vid"
check_response "$result" "\"address\": \"$trustee_account_address\""
response_does_not_contain "$result" "\"address\": \"$second_trustee_account_address\""
check_response "$result" "[\"$(echo "$passphrase" | dcld keys show jack -a)\"]"

test_divider

echo "Request all approved certificates must be empty, only 1 Trustee has approved so far"
result=$(dcld query pki all-x509-certs)
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$google_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$google_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$google_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$google_cert_subject_as_text\""

test_divider

echo "$second_trustee_account (Second Trustee) approves Root certificate"
result=$(echo "$passphrase" | dcld tx pki approve-add-x509-root-cert --subject="$google_cert_subject" --subject-key-id="$google_cert_subject_key_id" --from $second_trustee_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Get Certificates by subject and subjectKeyId must be Approved and contain 2 approvals. Request Root certificate"
result=$(dcld query pki x509-cert --subject="$google_cert_subject" --subject-key-id="$google_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$google_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$google_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$google_cert_serial_number\""
check_response "$result" "\"subjectAsText\": \"$google_cert_subject_as_text\""
check_response "$result" "\"address\": \"$trustee_account_address\""
check_response "$result" "\"address\": \"$second_trustee_account_address\""

echo "Get Certificates by subjectKeyId must be Approved and contain 2 approvals. Request Root certificate"
result=$(dcld query pki x509-cert --subject-key-id="$google_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subjectKeyId\": \"$google_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$google_cert_serial_number\""
check_response "$result" "\"subjectAsText\": \"$google_cert_subject_as_text\""
check_response "$result" "\"address\": \"$trustee_account_address\""
check_response "$result" "\"address\": \"$second_trustee_account_address\""

test_divider

echo "Request all proposed Root certificates must be empty"
result=$(dcld query pki all-proposed-x509-root-certs)
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$google_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$google_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$google_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$google_cert_subject_as_text\""

test_divider

echo "Request all approved certificates. It should contain one certificate with 2 approvals"
result=$(dcld query pki all-x509-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$google_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$google_cert_subject_key_id\""
check_response "$result" "\"address\": \"$trustee_account_address\""
check_response "$result" "\"address\": \"$second_trustee_account_address\""

test_divider

echo "Request all approved root certificates."
result=$(dcld query pki all-x509-root-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$google_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$google_cert_subject_key_id\""


echo "Request all revoked certificates must not contain google certification"
result=$(dcld query pki all-revoked-x509-certs)
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$google_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$google_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$google_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$google_cert_subject_as_text\""

# 12. PROPOSE REVOCATION OF GOOGLE ROOT CERT

echo "12. PROPOSE REVOCATION OF GOOGLE ROOT CERT"
test_divider

echo "$trustee_account (Trustee) proposes to revoke Root certificate"
result=$(echo "$passphrase" | dcld tx pki propose-revoke-x509-root-cert --subject="$google_cert_subject" --subject-key-id="$google_cert_subject_key_id" --from $trustee_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Request root certificate proposed to revoke and verify that it contains approval from $trustee_account_address"
result=$(dcld query pki proposed-x509-root-cert-to-revoke --subject="$google_cert_subject" --subject-key-id="$google_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"$google_cert_subject\""
check_response "$result" "\"$google_cert_subject_key_id\""
check_response "$result" "\"address\": \"$trustee_account_address\""

echo "Request all root certificates proposed to revoke"
result=$(dcld query pki all-proposed-x509-root-certs-to-revoke)
echo $result | jq
check_response "$result" "\"subject\": \"$google_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$google_cert_subject_key_id\""

test_divider

echo "Request all revoked certificates"
result=$(dcld query pki all-revoked-x509-certs)
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$google_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$google_cert_subject_key_id\""

test_divider

echo "Request all revoked root certificates"
result=$(dcld query pki all-revoked-x509-root-certs)
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$google_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$google_cert_subject_key_id\""

test_divider

echo "Request Root certificate proposed to revoke, it should have one approval from $trustee_account_address"
result=$(dcld query pki proposed-x509-root-cert-to-revoke --subject="$google_cert_subject" --subject-key-id="$google_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$google_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$google_cert_subject_key_id\""
check_response "$result" "[\"$(echo "$passphrase" | dcld keys show jack -a)\"]"  
check_response "$result" "\"address\": \"$trustee_account_address\""

test_divider

echo "Request all approved certificates"
result=$(dcld query pki all-x509-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$google_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$google_cert_subject_key_id\""

test_divider

echo "Request all approved root certificates"
result=$(dcld query pki all-x509-root-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$google_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$google_cert_subject_key_id\""

test_divider

echo "Request all subject certificates"
result=$(dcld query pki all-subject-x509-certs --subject="$google_cert_subject")
echo $result | jq
check_response "$result" "\"$google_cert_subject\""
check_response "$result" "\"$google_cert_subject_key_id\""

test_divider

# 13. APPROVE REVOCATION OF GOOGLE ROOT CERT

echo "13. APPROVE REVOCATION OF GOOGLE ROOT CERT"
test_divider

echo "$second_trustee_account (Second Trustee) approves to revoke Root certificate"
result=$(echo "$passphrase" | dcld tx pki approve-revoke-x509-root-cert --subject="$google_cert_subject" --subject-key-id="$google_cert_subject_key_id" --from $second_trustee_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Request all root certificates proposed to revoke. Nothing left in list as the certficate is revoked"
result=$(dcld query pki all-proposed-x509-root-certs-to-revoke)
response_does_not_contain "$result" "\"subject\": \"$google_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$google_cert_subject_key_id\""
echo $result | jq

test_divider

echo "Request all revoked certificates should contain approvals from both trustees"
result=$(dcld query pki all-revoked-x509-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$google_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$google_cert_subject_key_id\""
check_response "$result" "\"address\": \"$trustee_account_address\""
check_response "$result" "\"address\": \"$second_trustee_account_address\""

test_divider

echo "Request all revoked root certificates"
result=$(dcld query pki all-revoked-x509-root-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$google_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$google_cert_subject_key_id\""

test_divider

echo "Request revoked Root certificate and also check for approvals from both Trustees"
result=$(dcld query pki revoked-x509-cert --subject="$google_cert_subject" --subject-key-id="$google_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$google_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$google_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$google_cert_serial_number\""
check_response "$result" "\"subjectAsText\": \"$google_cert_subject_as_text\""
check_response "$result" "\"address\": \"$trustee_account_address\""
check_response "$result" "\"address\": \"$second_trustee_account_address\""

test_divider

echo "Request all approved certificates must be empty"
result=$(dcld query pki all-x509-certs)
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$google_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$google_cert_subject_key_id\""

echo "Request all approved root certificates must be empty"
result=$(dcld query pki all-x509-root-certs)
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$google_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$google_cert_subject_key_id\""

test_divider

echo "Approved Root certificate must be empty"
result=$(dcld query pki x509-cert --subject="$google_cert_subject" --subject-key-id="$google_cert_subject_key_id")
echo $result | jq
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$google_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$google_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$google_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$google_cert_subject_as_text\""

test_divider

echo "Request all subject certificates must be empty"
result=$(dcld query pki all-subject-x509-certs --subject="$google_cert_subject")
echo $result | jq
response_does_not_contain "$result" "\"$google_cert_subject\""
response_does_not_contain "$result" "\"$google_cert_subject_key_id\""


# CHECK REJECT CERTIFICATE
test_cert_subject="MDAxGDAWBgNVBAMMD01hdHRlciBUZXN0IFBBQTEUMBIGCisGAQQBgqJ8AgEMBDEyNUQ="
test_cert_subject_key_id="E2:90:8D:36:9C:3C:A3:C1:13:BB:09:E2:4D:C1:CC:C5:A6:66:91:D4"
test_cert_serial_number="1647312298631"
test_cert_subject_as_text="CN=Matter Test PAA,vid=0x125D"
test_cert_vid=4701

# 14. TEST PROPOSE AND REJECT ROOT CERTIFICATE
echo "14. TEST PROPOSE AND REJECT ROOT CERTIFICATE"
test_divider

echo "$trustee_account (Trustee) propose Root certificate"
test_root_path="integration_tests/constants/test_root_cert"
result=$(echo "$passphrase" | dcld tx pki propose-add-x509-root-cert --certificate="$test_root_path" --vid=$test_cert_vid --from $trustee_account --yes)
check_response "$result" "\"code\": 0"

echo "$trustee_account (Trustee) rejects Root certificate"
result=$(echo $passphrase | dcld tx pki reject-add-x509-root-cert --subject="$test_cert_subject" --subject-key-id="$test_cert_subject_key_id" --from $trustee_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Certificate not found in proposed-x509-root-cert"
result=$(dcld query pki proposed-x509-root-cert --subject="$test_cert_subject" --subject-key-id="$test_cert_subject_key_id")
echo $result | jq
check_response "$result" "Not Found"

test_divider

echo "Certificate not found in rejected-x509-root-cert"
result=$(dcld query pki rejected-x509-root-cert --subject="$test_cert_subject" --subject-key-id="$test_cert_subject_key_id")
echo $result | jq
check_response "$result" "Not Found"

test_divider

echo "Get Certificate by subject and subjectKeyId not found in x509-cert"
result=$(dcld query pki x509-cert --subject="$test_cert_subject" --subject-key-id="$test_cert_subject_key_id")
echo $result | jq
check_response "$result" "Not Found"

echo "Get Certificate by subjectKeyId not found in x509-cert"
result=$(dcld query pki x509-cert --subject-key-id="$test_cert_subject_key_id")
echo $result | jq
check_response "$result" "Not Found"


# 15. PROPOSE TEST ROOT CERT
echo "15. PROPOSE TEST ROOT CERT"
test_divider

echo "$user_account (Not Trustee) propose Root certificate"
test_root_path="integration_tests/constants/test_root_cert"
result=$(echo "$passphrase" | dcld tx pki propose-add-x509-root-cert --certificate="$test_root_path" --vid=$test_cert_vid --from $user_account --yes)
response_does_not_contain "$result" "\"code\": 0"

echo "$trustee_account (Trustee) propose Root certificate"
test_root_path="integration_tests/constants/test_root_cert"
result=$(echo "$passphrase" | dcld tx pki propose-add-x509-root-cert --certificate="$test_root_path" --vid=$test_cert_vid --from $trustee_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Request proposed Root certificate - there should be Approval"
result=$(dcld query pki proposed-x509-root-cert --subject="$test_cert_subject" --subject-key-id="$test_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$test_cert_subject\""
check_response "$result" "\"address\": \"$trustee_account_address\""
check_response "$result" "\"subjectKeyId\": \"$test_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$test_cert_serial_number\""
check_response "$result" "\"subjectAsText\": \"$test_cert_subject_as_text\""
check_response "$result" "\"vid\": $test_cert_vid"
echo $result | jq

# 16. TEST REJECT ROOT CERT
echo "16.  TEST REJECT ROOT CERT"
test_divider

random_string new_trustee1
echo "$new_trustee1 generates keys"
cmd="(echo $passphrase; echo $passphrase) | dcld keys add $new_trustee1"
result="$(bash -c "$cmd")"

test_divider

echo "Get key info for $new_trustee1"
result=$(echo $passphrase | dcld keys show $new_trustee1)
check_response "$result" "\"name\": \"$new_trustee1\""

test_divider

new_trustee_address1=$(echo $passphrase | dcld keys show $new_trustee1 -a)
new_trustee_pubkey1=$(echo $passphrase | dcld keys show $new_trustee1 -p)

test_divider

echo "Jack proposes account for $new_trustee1"
result=$(echo $passphrase | dcld tx auth propose-add-account --info="Jack is proposing this account" --address="$new_trustee_address1" --pubkey="$new_trustee_pubkey1" --roles="Trustee" --from jack --yes)
check_response "$result" "\"code\": 0"

echo "Alice approves account for \"$new_trustee1\""
result=$(echo $passphrase | dcld tx auth approve-add-account --address="$new_trustee_address1" --info="Alice is approving this account" --from alice --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Bob (Trustee) approves Root certificate"
result=$(echo $passphrase | dcld tx pki approve-add-x509-root-cert --subject="$test_cert_subject" --subject-key-id="$test_cert_subject_key_id" --from bob --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "$trustee_account (Trustee) rejects Root certificate"
result=$(echo $passphrase | dcld tx pki reject-add-x509-root-cert --subject="$test_cert_subject" --subject-key-id="$test_cert_subject_key_id" --from $trustee_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "$trustee_account (Trustee) can approve Root certificate even if already has rejected"
result=$(echo $passphrase | dcld tx pki approve-add-x509-root-cert --subject="$test_cert_subject" --subject-key-id="$test_cert_subject_key_id" --from $trustee_account --yes 2>&1 || true)
check_response "$result" "\"code\": 0"

test_divider

echo "$trustee_account (Trustee) rejects Root certificate"
result=$(echo $passphrase | dcld tx pki reject-add-x509-root-cert --subject="$test_cert_subject" --subject-key-id="$test_cert_subject_key_id" --from $trustee_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "$trustee_account (Trustee) cannot reject Root certificate for the second time"
result=$(echo $passphrase | dcld tx pki reject-add-x509-root-cert --subject="$test_cert_subject" --subject-key-id="$test_cert_subject_key_id" --from $trustee_account --yes 2>&1 || true)
response_does_not_contain "$result" "\"code\": 0"

test_divider

echo "Certificate must be still in Proposed state but with Approval from $trustee_account. Request proposed Root certificate"
result=$(dcld query pki proposed-x509-root-cert --subject="$test_cert_subject" --subject-key-id="$test_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$test_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$test_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$test_cert_serial_number\""
check_response "$result" "\"subjectAsText\": \"$test_cert_subject_as_text\""
check_response "$result" "\"address\": \"$trustee_account_address\""
response_does_not_contain "$result" "\"address\": \"$second_trustee_account_address\""
check_response "$result" "[\"$(echo "$passphrase" | dcld keys show jack -a)\"]"

test_divider

echo "Request all rejected certificates must be empty, only 1 Trustee has rejected so far"
result=$(dcld query pki all-rejected-x509-root-certs)
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$test_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$test_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$test_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$test_cert_subject_as_text\""

test_divider

reject_schema_version_4=4
echo "$second_trustee_account (Second Trustee) rejects Root certificate"
result=$(echo "$passphrase" | dcld tx pki reject-add-x509-root-cert --subject="$test_cert_subject" --subject-key-id="$test_cert_subject_key_id" --schemaVersion=$reject_schema_version_4 --from $second_trustee_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Alice proposes to revoke account for $new_trustee1"
result=$(echo $passphrase | dcld tx auth propose-revoke-account --address="$new_trustee_address1" --from alice --yes)
check_response "$result" "\"code\": 0"

echo "Bob approves to revoke account for $new_trustee1"
result=$(echo $passphrase | dcld tx auth approve-revoke-account --address="$new_trustee_address1" --from bob --yes)
check_response "$result" "\"code\": 0"

echo "Jack approves to revoke account for $new_trustee1"
result=$(echo $passphrase | dcld tx auth approve-revoke-account --address="$new_trustee_address1" --from jack --yes)
check_response "$result" "\"code\": 0"


echo "Certificate must be Rejected and contains 2 rejects. Request Root certificate"
result=$(dcld query pki rejected-x509-root-cert --subject="$test_cert_subject" --subject-key-id="$test_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$test_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$test_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$test_cert_serial_number\""
check_response "$result" "\"subjectAsText\": \"$test_cert_subject_as_text\""
check_response "$result" "\"address\": \"$trustee_account_address\""
check_response "$result" "\"address\": \"$second_trustee_account_address\""
check_response "$result" "\"schemaVersion\": $reject_schema_version_4"

test_divider

echo "Request all proposed Root certificates must be empty"
result=$(dcld query pki all-proposed-x509-root-certs)
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$test_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$test_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$test_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$test_cert_subject_as_text\""

test_divider

echo "Request all approved root certificates."
result=$(dcld query pki all-x509-root-certs)
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$test_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$test_cert_subject_key_id\""

test_divider

echo "Request all rejected certificates. It should contain one certificate with 2 rejects"
result=$(dcld query pki all-rejected-x509-root-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$test_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$test_cert_subject_key_id\""
check_response "$result" "\"address\": \"$trustee_account_address\""
check_response "$result" "\"address\": \"$second_trustee_account_address\""


# 17. TEST PROPOSE ROOT CERTIFICATE
echo "17. PROPOSE TEST ROOT CERT"
test_divider

echo "$user_account (Not Trustee) propose Root certificate"
test_root_path="integration_tests/constants/test_root_cert"
result=$(echo "$passphrase" | dcld tx pki propose-add-x509-root-cert --certificate="$test_root_path" --vid $test_cert_vid --from $user_account --yes)
response_does_not_contain "$result" "\"code\": 0"

echo "$trustee_account (Trustee) propose Root certificate"
test_root_path="integration_tests/constants/test_root_cert"
result=$(echo "$passphrase" | dcld tx pki propose-add-x509-root-cert --certificate="$test_root_path" --vid $test_cert_vid --from $trustee_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Request proposed Root certificate - there should be Approval"
result=$(dcld query pki proposed-x509-root-cert --subject="$test_cert_subject" --subject-key-id="$test_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$test_cert_subject\""
check_response "$result" "\"address\": \"$trustee_account_address\""
check_response "$result" "\"subjectKeyId\": \"$test_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$test_cert_serial_number\""
check_response "$result" "\"subjectAsText\": \"$test_cert_subject_as_text\""
check_response "$result" "\"vid\": $test_cert_vid"
echo $result | jq

test_divider

echo "Certificate not found in rejected-x509-root-cert"
result=$(dcld query pki rejected-x509-root-cert --subject="$test_cert_subject" --subject-key-id="$test_cert_subject_key_id")
echo $result | jq
check_response "$result" "Not Found"

# 18. TEST APPROVE ROOT CERTIFICATE
echo "18. TEST APPROVE ROOT CERT"
test_divider

echo "$second_trustee_account (Trustee) approve Root certificate"
result=$(echo $passphrase | dcld tx pki approve-add-x509-root-cert --subject="$test_cert_subject" --subject-key-id="$test_cert_subject_key_id" --from $second_trustee_account_address --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Get Certificates by subject must be Approved and contains 2 approvals. Request Root certificate"
result=$(dcld query pki x509-cert --subject="$test_cert_subject" --subject-key-id="$test_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$test_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$test_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$test_cert_serial_number\""
check_response "$result" "\"subjectAsText\": \"$test_cert_subject_as_text\""
check_response "$result" "\"address\": \"$trustee_account_address\""
check_response "$result" "\"address\": \"$second_trustee_account_address\""
check_response "$result" "\"vid\": $test_cert_vid"

echo "Get Certificates by subjectKeyId must be Approved and contains 2 approvals. Request Root certificate"
result=$(dcld query pki x509-cert --subject-key-id="$test_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subjectKeyId\": \"$test_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$test_cert_serial_number\""
check_response "$result" "\"subjectAsText\": \"$test_cert_subject_as_text\""
check_response "$result" "\"address\": \"$trustee_account_address\""
check_response "$result" "\"address\": \"$second_trustee_account_address\""
check_response "$result" "\"vid\": $test_cert_vid"

test_divider

echo "PASS"