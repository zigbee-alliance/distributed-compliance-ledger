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

# DA
da_root_path="integration_tests/constants/root_cert"
da_root_subject="MDQxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMRAwDgYDVQQKDAdyb290LWNh"
da_root_subject_key_id="5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB"

da_intermediate_path="integration_tests/constants/intermediate_cert"
da_intermediate_subject="MDwxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMRgwFgYDVQQKDA9pbnRlcm1lZGlhdGUtY2E="
da_intermediate_subject_key_id="4E:3B:73:F4:70:4D:C2:98:0D:DB:C8:5A:5F:02:3B:BF:86:25:56:2B"

da_leaf_path="integration_tests/constants/leaf_cert"
da_leaf_subject="MDExCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMQ0wCwYDVQQKDARsZWFm"
da_leaf_subject_key_id="30:F4:65:75:14:20:B2:AF:3D:14:71:17:AC:49:90:93:3E:24:A0:1F"

# NOC
noc_root_path="integration_tests/constants/noc_root_cert_1"
noc_root_subject="MHoxCzAJBgNVBAYTAlVaMRMwEQYDVQQIDApTb21lIFN0YXRlMREwDwYDVQQHDAhUYXNoa2VudDEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMQ4wDAYDVQQDDAVOT0MtMQ=="
noc_root_subject_key_id="44:EB:4C:62:6B:25:48:CD:A2:B3:1C:87:41:5A:08:E7:2B:B9:83:26"

noc_intermediate_path="integration_tests/constants/noc_cert_1"
noc_intermediate_subject="MIGCMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTETMBEGA1UEBwwKU29tZSBTdGF0ZTEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRQwEgYDVQQDDAtOT0MtY2hpbGQtMQ=="
noc_intermediate_subject_key_id="02:72:6E:BC:BB:EF:D6:BD:8D:9B:42:AE:D4:3C:C0:55:5F:66:3A:B3"

noc_leaf_path="integration_tests/constants/noc_leaf_cert_1"
noc_leaf_subject="MIGBMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTETMBEGA1UEBwwKU29tZSBTdGF0ZTEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRMwEQYDVQQDDApOT0MtbGVhZi0x"
noc_leaf_subject_key_id="77:1F:DB:C4:4C:B1:29:7E:3C:EB:3E:D8:2A:38:0B:63:06:07:00:01"

# Accounts
trustee_account="jack"
second_trustee_account="alice"

vid_in_hex_format=0x6006
vid=24582

vendor_account=vendor_account_$vid_in_hex_format
echo "Create Vendor account - $vendor_account"
create_new_vendor_account $vendor_account $vid_in_hex_format

vid_2_in_hex_format=0x125D
vid_2=4701

test_divider

# Body

echo "1. QUERY Certificates (EMPTY)"

test_divider

echo "Request all approved DA certificates must be empty"
result=$(dcld query pki all-x509-certs)
echo $result | jq
check_response "$result" "\[\]"
response_does_not_contain "$result" "\"subjectKeyId\": \"$da_root_subject_key_id\""

echo "Request all NOC certificates must be empty"
result=$(dcld query pki all-noc-x509-certs)
echo $result | jq
check_response "$result" "\[\]"
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_subject_key_id\""

test_divider

echo "Request approved DA certificate must be empty"
result=$(dcld query pki x509-cert --subject="$da_root_subject" --subject-key-id="$da_root_subject_key_id")
echo $result | jq
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subjectKeyId\": \"$da_root_subject_key_id\""

echo "Request NOC Root certificate must be empty"
result=$(dcld query pki noc-x509-cert --subject="$noc_root_subject" --subject-key-id="$noc_root_subject_key_id")
echo $result | jq
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_subject_key_id\""

test_divider

echo "Request proposed DA Root certificate must be empty"
result=$(dcld query pki proposed-x509-root-cert --subject="$da_root_subject" --subject-key-id="$da_root_subject_key_id")
echo $result | jq
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subjectKeyId\": \"$da_root_subject_key_id\""

echo "Request all proposed DA Root certificates must be empty"
result=$(dcld query pki all-proposed-x509-root-certs)
check_response "$result" "\[\]"
echo $result | jq
response_does_not_contain "$result" "\"subjectKeyId\": \"$da_root_subject_key_id\""

test_divider

echo "Request all revoked DA certificates must be empty"
result=$(dcld query pki all-revoked-x509-certs)
echo $result | jq
check_response "$result" "\[\]"
response_does_not_contain "$result" "\"subjectKeyId\": \"$da_root_subject_key_id\""

echo "Request all revoked NOC Root certificates must be empty"
result=$(dcld query pki all-revoked-noc-x509-root-certs)
echo $result | jq
check_response "$result" "\[\]"
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_subject_key_id\""

echo "Request all revoked NOC ICA certificates must be empty"
result=$(dcld query pki all-revoked-noc-x509-ica-certs)
echo $result | jq
check_response "$result" "\[\]"
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_intermediate_subject_key_id\""

echo "Request revoked DA certificate must be empty"
result=$(dcld query pki revoked-x509-cert --subject="$da_root_subject" --subject-key-id="$da_root_subject_key_id")
echo $result | jq
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subjectKeyId\": \"$da_root_subject_key_id\""

echo "Request revoked NOC Root certificate must be empty"
result=$(dcld query pki revoked-noc-x509-root-cert --subject="$noc_root_subject" --subject-key-id="$noc_root_subject_key_id")
echo $result | jq
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_subject_key_id\""

echo "Request revoked NOC ICA certificate must be empty"
result=$(dcld query pki revoked-noc-x509-ica-cert --subject="$noc_intermediate_subject" --subject-key-id="$noc_intermediate_subject_key_id")
echo $result | jq
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_intermediate_subject_key_id\""

test_divider

echo "Request all DA certificates by subject must be empty"
result=$(dcld query pki all-subject-x509-certs --subject="$da_root_subject")
echo $result | jq
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"$da_root_subject\""

echo "Request all NOC certificates by subject must be empty"
result=$(dcld query pki all-noc-subject-x509-certs --subject="$noc_root_subject")
echo $result | jq
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"$noc_root_subject\""

test_divider

echo "Request all child certificates for DA root must be empty"
result=$(dcld query pki all-child-x509-certs --subject="$da_root_subject" --subject-key-id="$da_root_subject_key_id")
echo $result | jq
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"$da_root_subject_key_id\""

echo "Request all child certificates for NOC root must be empty"
result=$(dcld query pki all-child-x509-certs --subject="$noc_root_subject" --subject-key-id="$noc_root_subject_key_id")
echo $result | jq
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"$noc_root_subject_key_id\""

test_divider

echo "2. ADD Certificates"

test_divider

echo "$trustee_account (Trustee) propose DA Root certificate"
cert_schema_version_0=0
result=$(echo "$passphrase" | dcld tx pki propose-add-x509-root-cert --certificate="$da_root_path" --schemaVersion=$cert_schema_version_0 --from $trustee_account   --vid $vid --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "$second_trustee_account (Second Trustee) approves DA Root certificate"
result=$(echo "$passphrase" | dcld tx pki approve-add-x509-root-cert --subject="$da_root_subject" --subject-key-id="$da_root_subject_key_id" --from $second_trustee_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "$vendor_account adds DA Intermediate certificate"
result=$(echo "$passphrase" | dcld tx pki add-x509-cert --certificate="$da_intermediate_path" --schemaVersion=$cert_schema_version_0 --from $vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "$vendor_account add DA Leaf certificate"
result=$(echo "$passphrase" | dcld tx pki add-x509-cert --certificate="$da_leaf_path" --from $vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "$vendor_account add NOC Root certificate"
cert_schema_version_0=0
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-root-cert --certificate="$noc_root_path" --schemaVersion=$cert_schema_version_0 --from $vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "$vendor_account adds NOC Intermediate certificate"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-ica-cert --certificate="$noc_intermediate_path" --from $vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "$vendor_account adds NOC Leaf certificate"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-ica-cert --certificate="$noc_leaf_path" --from $vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "3. QUERY Certificates (PRESENT)"

test_divider

echo "Request all certificates"
result=$(dcld query pki all-certs)
echo $result | jq
check_response "$result" "\"subjectKeyId\": \"$da_root_subject_key_id\""
check_response "$result" "\"subjectKeyId\": \"$da_intermediate_subject_key_id\""
check_response "$result" "\"subjectKeyId\": \"$da_leaf_subject_key_id\""
check_response "$result" "\"subjectKeyId\": \"$noc_root_subject_key_id\""
check_response "$result" "\"subjectKeyId\": \"$noc_intermediate_subject_key_id\""
check_response "$result" "\"subjectKeyId\": \"$noc_leaf_subject_key_id\""

echo "Request all approved DA certificates"
result=$(dcld query pki all-x509-certs)
echo $result | jq
check_response "$result" "\"subjectKeyId\": \"$da_root_subject_key_id\""
check_response "$result" "\"subjectKeyId\": \"$da_intermediate_subject_key_id\""
check_response "$result" "\"subjectKeyId\": \"$da_leaf_subject_key_id\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_subject_key_id\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_intermediate_subject_key_id\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_leaf_subject_key_id\""

echo "Request all NOC certificates"
result=$(dcld query pki all-noc-x509-certs)
echo $result | jq
check_response "$result" "\"subjectKeyId\": \"$noc_root_subject_key_id\""
check_response "$result" "\"subjectKeyId\": \"$noc_intermediate_subject_key_id\""
check_response "$result" "\"subjectKeyId\": \"$noc_leaf_subject_key_id\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$da_root_subject_key_id\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$da_intermediate_subject_key_id\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$da_leaf_subject_key_id\""

test_divider

echo "Request DA certificate using global command"
result=$(dcld query pki cert --subject="$da_root_subject" --subject-key-id="$da_root_subject_key_id")
echo $result | jq
check_response "$result" "\"subjectKeyId\": \"$da_root_subject_key_id\""

echo "Request NOC certificate using global command"
result=$(dcld query pki cert --subject="$noc_root_subject" --subject-key-id="$noc_root_subject_key_id")
echo $result | jq
check_response "$result" "\"subjectKeyId\": \"$noc_root_subject_key_id\""

echo "Request DA certificate"
result=$(dcld query pki x509-cert --subject="$da_root_subject" --subject-key-id="$da_root_subject_key_id")
echo $result | jq
check_response "$result" "\"subjectKeyId\": \"$da_root_subject_key_id\""

echo "Request NOC certificate using DA command (must be empty)"
result=$(dcld query pki x509-cert --subject="$noc_root_subject" --subject-key-id="$noc_root_subject_key_id")
echo $result | jq
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_subject_key_id\""

echo "Request NOC Root certificate"
result=$(dcld query pki noc-x509-cert --subject="$noc_root_subject" --subject-key-id="$noc_root_subject_key_id")
echo $result | jq
check_response "$result" "\"subjectKeyId\": \"$noc_root_subject_key_id\""

echo "Request DA certificate using NOC command (must be empty)"
result=$(dcld query pki noc-x509-cert --subject="$da_root_subject" --subject-key-id="$da_root_subject_key_id")
echo $result | jq
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subjectKeyId\": \"$da_root_subject_key_id\""

test_divider

echo "Request DA certificates by subject using global command"
result=$(dcld query pki all-subject-certs --subject=$da_root_subject)
echo $result | jq
check_response "$result" "\"$da_root_subject\""
check_response "$result" "\"$da_root_subject_key_id\""

echo "Request NOC certificates by subject using global command"
result=$(dcld query pki all-subject-certs --subject=$noc_root_subject)
echo $result | jq
check_response "$result" "\"$noc_root_subject\""
check_response "$result" "\"$noc_root_subject_key_id\""

echo "Request all DA certificates by subject must be empty"
result=$(dcld query pki all-subject-x509-certs --subject="$da_root_subject")
echo $result | jq
check_response "$result" "\"$da_root_subject\""

echo "Request all NOC certificates by subject using DA command (must be empty)"
result=$(dcld query pki all-subject-x509-certs --subject="$noc_root_subject")
echo $result | jq
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"$noc_root_subject\""

echo "Request all NOC certificates by subject"
result=$(dcld query pki all-noc-subject-x509-certs --subject="$noc_root_subject")
echo $result | jq
check_response "$result" "\"$noc_root_subject\""

echo "Request all DA certificates by subject using NOC command (must be empty)"
result=$(dcld query pki all-noc-subject-x509-certs --subject="$da_root_subject")
echo $result | jq
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"$da_root_subject\""

test_divider

echo "Request all child certificates for DA root"
result=$(dcld query pki all-child-x509-certs --subject="$da_root_subject" --subject-key-id="$da_root_subject_key_id")
echo $result | jq
check_response "$result" "\"$da_root_subject\""
check_response "$result" "\"$da_root_subject_key_id\""

echo "Request all child certificates for NOC root"
result=$(dcld query pki all-child-x509-certs --subject="$noc_root_subject" --subject-key-id="$noc_root_subject_key_id")
echo $result | jq
check_response "$result" "\"$noc_root_subject\""
check_response "$result" "\"$noc_root_subject_key_id\""

test_divider

echo "4. Revoke Certificates"

test_divider

echo "$vendor_account revoke an intermediate DA certificate"
result=$(echo "$passphrase" | dcld tx pki revoke-x509-cert --subject="$da_intermediate_subject" --subject-key-id="$da_intermediate_subject_key_id" --from=$vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "$vendor_account revoke an intermediate NOC certificate"
result=$(echo "$passphrase" | dcld tx pki revoke-noc-x509-ica-cert --subject="$noc_intermediate_subject" --subject-key-id="$noc_intermediate_subject_key_id" --from=$vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Request all revoked DA certificates"
result=$(dcld query pki all-revoked-x509-certs)
echo $result | jq
check_response "$result" "\"subjectKeyId\": \"$da_intermediate_subject_key_id\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_intermediate_subject_key_id\""

echo "Request all revoked NOC ICA certificates must be empty"
result=$(dcld query pki all-revoked-noc-x509-ica-certs)
echo $result | jq
check_response "$result" "\"subjectKeyId\": \"$noc_intermediate_subject_key_id\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$da_intermediate_subject_key_id\""

echo "Request revoked DA certificate "
result=$(dcld query pki revoked-x509-cert --subject="$da_intermediate_subject" --subject-key-id="$da_intermediate_subject_key_id")
echo $result | jq
check_response "$result" "\"subjectKeyId\": \"$da_intermediate_subject_key_id\""

echo "Request revoked NOC certificate using DA command (must be empty)"
result=$(dcld query pki revoked-x509-cert --subject="$noc_intermediate_subject" --subject-key-id="$noc_intermediate_subject_key_id")
echo $result | jq
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_intermediate_subject_key_id\""

echo "Request revoked NOC ICA certificate"
result=$(dcld query pki revoked-noc-x509-ica-cert --subject="$noc_intermediate_subject" --subject-key-id="$noc_intermediate_subject_key_id")
echo $result | jq
check_response "$result" "\"subjectKeyId\": \"$noc_intermediate_subject_key_id\""

echo "Request revoked DA certificate using NOC command (must be empty)"
result=$(dcld query pki revoked-noc-x509-ica-cert --subject="$da_intermediate_subject" --subject-key-id="$da_intermediate_subject_key_id")
echo $result | jq
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subjectKeyId\": \"$da_intermediate_subject_key_id\""

test_divider

echo "PASS"