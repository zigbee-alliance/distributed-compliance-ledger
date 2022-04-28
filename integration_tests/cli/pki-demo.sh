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

root_cert_subject="Tz1yb290LWNhLFNUPXNvbWUtc3RhdGUsQz1BVQ=="
root_cert_subject_key_id="5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB"
root_cert_serial_number="442314047376310867378175982234956458728610743315"
root_cert_subject_as_text="O=root-ca,ST=some-state,C=AU"

intermediate_cert_subject="Tz1pbnRlcm1lZGlhdGUtY2EsU1Q9c29tZS1zdGF0ZSxDPUFV"
intermediate_cert_subject_key_id="4E:3B:73:F4:70:4D:C2:98:0D:DB:C8:5A:5F:02:3B:BF:86:25:56:2B"
intermediate_cert_serial_number="169917617234879872371588777545667947720450185023"
intermediate_cert_subject_as_text="O=intermediate-ca,ST=some-state,C=AU"

leaf_cert_subject="Tz1sZWFmLFNUPXNvbWUtc3RhdGUsQz1BVQ=="
leaf_cert_subject_key_id="30:F4:65:75:14:20:B2:AF:3D:14:71:17:AC:49:90:93:3E:24:A0:1F"
leaf_cert_serial_number="143290473708569835418599774898811724528308722063"
leaf_cert_subject_as_text="O=leaf,ST=some-state,C=AU"

# Preparation of Actors

trustee_account="jack"
second_trustee_account="alice"

trustee_account_address=$(echo $passphrase | dcld keys show jack -a)
second_trustee_account_address=$(echo $passphrase | dcld keys show alice -a)

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
result=$(echo "$passphrase" | dcld tx pki propose-add-x509-root-cert --certificate="$root_path" --from $user_account --yes)
check_response "$result" "\"code\": 0"


test_divider

echo "Request all proposed Root certificates - There should be no approvals"
result=$(dcld query pki all-proposed-x509-root-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""


test_divider


echo "Request proposed Root certificate - there should be no Approval"
result=$(dcld query pki proposed-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$root_cert_serial_number\""
check_response "$result" "\"subjectAsText\": \"$root_cert_subject_as_text\""
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

echo "$trustee_account (Trustee) approve Root certificate"
result=$(echo $passphrase | dcld tx pki approve-add-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from $trustee_account --yes)
check_response "$result" "\"code\": 0"


test_divider

echo "Certificate must be still in Proposed state but with Approval from $trustee_account. Request proposed Root certificate"
result=$(dcld query pki proposed-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$root_cert_serial_number\""
check_response "$result" "\"address\": \"$trustee_account_address\""
check_response "$result" "\"subjectAsText\": \"$root_cert_subject_as_text\""
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

echo "Certificate must be Approved and contain 2 approvals. Request Root certificate"
result=$(dcld query pki x509-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$root_cert_subject\""
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


echo "$user_account (Not Trustee) adds Intermediate certificate"
intermediate_path="integration_tests/constants/intermediate_cert"
result=$(echo "$passphrase" | dcld tx pki add-x509-cert --certificate="$intermediate_path" --from $user_account --yes)
check_response "$result" "\"code\": 0"


test_divider

echo "Request Intermediate certificate - There are no approvals for Intermidiate Certificates"
result=$(dcld query pki x509-cert --subject="$intermediate_cert_subject" --subject-key-id="$intermediate_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
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

echo "$trustee_account (Trustee) add Leaf certificate"
leaf_path="integration_tests/constants/leaf_cert"
result=$(echo "$passphrase" | dcld tx pki add-x509-cert --certificate="$leaf_path" --from $trustee_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Request Leaf certificate - There is no approvals on leaf certificate"
result=$(dcld query pki x509-cert --subject="$leaf_cert_subject" --subject-key-id="$leaf_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$leaf_cert_subject\""
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

echo "$user_account (Not Trustee) revokes Intermediate certificate. This must also revoke its child - Leaf certificate."
result=$(echo "$passphrase" | dcld tx pki revoke-x509-cert --subject="$intermediate_cert_subject" --subject-key-id="$intermediate_cert_subject_key_id" --from=$user_account --yes)
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
check_response "$result" "\"subject\": \"$leaf_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""
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
check_response "$result" "\"subject\": \"$leaf_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$leaf_cert_serial_number\""

test_divider

echo "Request all approved certificates"
result=$(dcld query pki all-x509-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"subject\": \"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
response_does_not_contain "$result" "\"subject\": \"$leaf_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""

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
response_does_not_contain "$result" "\"$root_cert_subject\""
response_does_not_contain "$result" "\"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"$intermediate_cert_subject_key_id\""
response_does_not_contain "$result" "\"$leaf_cert_subject\""
response_does_not_contain "$result" "\"$leaf_cert_subject_key_id\""

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

echo "Approved Leaf certificate must be empty"
result=$(dcld query pki x509-cert --subject="$leaf_cert_subject" --subject-key-id="$leaf_cert_subject_key_id")
echo $result | jq
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$leaf_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$leaf_cert_serial_number\""

test_divider

# 7. PROPOSE REVOCATION OF ROOT CERT

echo "7. PROPOSE REVOCATION OF ROOT CERT"
test_divider

echo "$trustee_account (Trustee) proposes to revoke Root certificate"
result=$(echo "$passphrase" | dcld tx pki propose-revoke-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from $trustee_account --yes)
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
check_response "$result" "\"subject\": \"$leaf_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""
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
response_does_not_contain "$result" "\"subject\": \"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
response_does_not_contain "$result" "\"subject\": \"$leaf_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""


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
response_does_not_contain "$result" "\"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"$intermediate_cert_subject_key_id\""
response_does_not_contain "$result" "\"$leaf_cert_subject\""
response_does_not_contain "$result" "\"$leaf_cert_subject_key_id\""

test_divider

# 8. APPROVE REVOCATION OF ROOT CERT

echo "8. APPROVE REVOCATION OF ROOT CERT"
test_divider


echo "$second_trustee_account (Second Trustee) approves to revoke Root certificate"
result=$(echo "$passphrase" | dcld tx pki approve-revoke-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from $second_trustee_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Request all root certificates proposed to revoke. Nothing left in list as the certficate is revoked"
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
check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
check_response "$result" "\"subject\": \"$leaf_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
check_response "$result" "\"address\": \"$trustee_account_address\""
check_response "$result" "\"address\": \"$second_trustee_account_address\""


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

echo "Request all approved certificates must be empty"
result=$(dcld query pki all-x509-certs)
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$root_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"subject\": \"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
response_does_not_contain "$result" "\"subject\": \"$leaf_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""


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
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$intermediate_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$intermediate_cert_subject_as_text\""


test_divider

echo "Approved Leaf certificate must be empty"
result=$(dcld query pki x509-cert --subject="$leaf_cert_subject" --subject-key-id="$leaf_cert_subject_key_id")
echo $result | jq
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$leaf_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$leaf_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$leaf_cert_subject_as_text\""


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

google_cert_subject="Q049TWF0dGVyIFBBQSAxLE89R29vZ2xlLEM9VVMsMS4zLjYuMS40LjEuMzcyNDQuMi4xPSMxMzA0MzYzMDMwMzY="
google_cert_subject_key_id="B0:00:56:81:B8:88:62:89:62:80:E1:21:18:A1:A8:BE:09:DE:93:21"
google_cert_serial_number="1"
google_cert_subject_as_text="CN=Matter PAA 1,O=Google,C=US,vid=0x6006"

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

echo "Request all approved certificates must be empty"
result=$(dcld query pki all-x509-certs)
check_response "$result" "\[\]"
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

echo "Request all revoked certificates must be empty"
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

echo "$user_account (Not Trustee) propose Root certificate"
google_root_path="integration_tests/constants/google_root_cert"
result=$(echo "$passphrase" | dcld tx pki propose-add-x509-root-cert --certificate="$google_root_path" --from $user_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Request proposed Root certificate - there should be no Approval"
result=$(dcld query pki proposed-x509-root-cert --subject="$google_cert_subject" --subject-key-id="$google_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$google_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$google_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$google_cert_serial_number\""
check_response "$result" "\"subjectAsText\": \"$google_cert_subject_as_text\""
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

echo "$trustee_account (Trustee) approve Root certificate"
result=$(echo $passphrase | dcld tx pki approve-add-x509-root-cert --subject="$google_cert_subject" --subject-key-id="$google_cert_subject_key_id" --from $trustee_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Certificate must be still in Proposed state but with Approval from $trustee_account. Request proposed Root certificate"
result=$(dcld query pki proposed-x509-root-cert --subject="$google_cert_subject" --subject-key-id="$google_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$google_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$google_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$google_cert_serial_number\""
check_response "$result" "\"subjectAsText\": \"$google_cert_subject_as_text\""
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

echo "Certificate must be Approved and contain 2 approvals. Request Root certificate"
result=$(dcld query pki x509-cert --subject="$google_cert_subject" --subject-key-id="$google_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$google_cert_subject\""
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
test_cert_subject="Q049TWF0dGVyIFRlc3QgUEFBLDEuMy42LjEuNC4xLjM3MjQ0LjIuMT0jMTMwNDMxMzIzNTQ0"
test_cert_subject_key_id="E2:90:8D:36:9C:3C:A3:C1:13:BB:09:E2:4D:C1:CC:C5:A6:66:91:D4"
test_cert_serial_number="1647312298631"
test_cert_subject_as_text="CN=Matter Test PAA,vid=0x125D"

# 14. PROPOSE TEST ROOT CERT
echo "14. PROPOSE TEST ROOT CERT"
test_divider

echo "$user_account (Not Trustee) propose Root certificate"
test_root_path="integration_tests/constants/test_root_cert"
result=$(echo "$passphrase" | dcld tx pki propose-add-x509-root-cert --certificate="$test_root_path" --from $user_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Request proposed Root certificate - there should be no Approval"
result=$(dcld query pki proposed-x509-root-cert --subject="$test_cert_subject" --subject-key-id="$test_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$test_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$test_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$test_cert_serial_number\""
check_response "$result" "\"subjectAsText\": \"$test_cert_subject_as_text\""
echo $result | jq

# 15. REJECT TEST ROOT CERT
echo "11. REJECT TEST ROOT CERT"
test_divider

echo "$trustee_account (Trustee) reject Root certificate"
result=$(echo $passphrase | dcld tx pki reject-add-x509-root-cert --subject="$test_cert_subject" --subject-key-id="$test_cert_subject_key_id" --from $trustee_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "$trustee_account (Trustee) doesn't reject Root certificate at the second time"
result=$(echo $passphrase | dcld tx pki reject-add-x509-root-cert --subject="$test_cert_subject" --subject-key-id="$test_cert_subject_key_id" --from $trustee_account --yes 2>&1 || true)
response_does_not_contain "$result" "\"code\": 0"

test_divider

echo "$trustee_account (Trustee) doesn't approve Root certificate, because already has rejected"
result=$(echo $passphrase | dcld tx pki approve-add-x509-root-cert --subject="$test_cert_subject" --subject-key-id="$test_cert_subject_key_id" --from $trustee_account --yes 2>&1 || true)
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

echo "$second_trustee_account (Second Trustee) rejects Root certificate"
result=$(echo "$passphrase" | dcld tx pki reject-add-x509-root-cert --subject="$test_cert_subject" --subject-key-id="$test_cert_subject_key_id" --from $second_trustee_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Certificate must be Rejected and contain 2 rejects. Request Root certificate"
result=$(dcld query pki rejected-x509-root-cert --subject="$test_cert_subject" --subject-key-id="$test_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$test_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$test_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$test_cert_serial_number\""
check_response "$result" "\"subjectAsText\": \"$test_cert_subject_as_text\""
check_response "$result" "\"address\": \"$trustee_account_address\""
check_response "$result" "\"address\": \"$second_trustee_account_address\""

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

test_divider

echo "PASS"