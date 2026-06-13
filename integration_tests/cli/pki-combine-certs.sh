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
da_root_subject="MDQxCzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpzb21lLXN0YXRlMRAwDgYDVQQKEwdyb290LWNh"
da_root_subject_key_id="DF:4E:AF:B0:8C:9C:37:78:1A:E7:53:12:CA:E4:78:6B:48:1E:AF:B0"

da_intermediate_path="integration_tests/constants/intermediate_cert"
da_intermediate_subject="MDwxCzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpzb21lLXN0YXRlMRgwFgYDVQQKEw9pbnRlcm1lZGlhdGUtY2E="
da_intermediate_subject_key_id="1B:73:2A:91:34:46:8A:90:2A:87:19:91:E4:BD:8F:69:3A:F9:04:77"

da_leaf_path="integration_tests/constants/leaf_cert"
da_leaf_subject="MDExCzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpzb21lLXN0YXRlMQ0wCwYDVQQKEwRsZWFm"
da_leaf_subject_key_id="2A:31:8D:39:6E:50:DA:96:DF:95:C5:98:83:68:F0:58:B2:15:B3:3A"

# NOC
noc_root_path="integration_tests/constants/noc_root_cert_1"
noc_root_subject="MHoxCzAJBgNVBAYTAlVaMRMwEQYDVQQIEwpTb21lIFN0YXRlMREwDwYDVQQHEwhUYXNoa2VudDEYMBYGA1UEChMPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLExBUZXN0aW5nIERpdmlzaW9uMQ4wDAYDVQQDEwVOT0MtMQ=="
noc_root_subject_key_id="0E:10:B8:5D:96:7A:08:33:C7:C5:44:49:0E:28:0F:C1:6E:D5:D4:7C"

noc_intermediate_path="integration_tests/constants/noc_cert_1"
noc_intermediate_subject="MIGCMQswCQYDVQQGEwJVWjETMBEGA1UECBMKU29tZSBTdGF0ZTETMBEGA1UEBxMKU29tZSBTdGF0ZTEYMBYGA1UEChMPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLExBUZXN0aW5nIERpdmlzaW9uMRQwEgYDVQQDEwtOT0MtY2hpbGQtMQ=="
noc_intermediate_subject_key_id="06:9F:5A:E0:1F:23:3E:9F:C7:4F:B6:F9:A2:33:47:33:62:7A:07:C5"

# NocLeafCert1 (a §6.5.12 NOC end-entity, cA=FALSE / KU=digitalSignature) is no
# longer accepted by the strict add-noc-x509-ica-cert handler — that path now
# requires the §6.5.12 ICAC profile (cA=TRUE) or the §6.5.12 VVSC profile
# (cA=FALSE / KU=digitalSignature, only with --is-vid-verification-signer=true).
# Use the VVSC leaf fixture instead, chained under a VVSC root + ICA registered
# below as VIDSignerPKI.
vvsc_root_path="integration_tests/constants/vvsc_root_cert_1"
vvsc_root_subject="MIGWMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTERMA8GA1UEBwwIVGFzaGtlbnQxGDAWBgNVBAoMD0V4YW1wbGUgQ29tcGFueTEZMBcGA1UECwwQVGVzdGluZyBEaXZpc2lvbjEUMBIGA1UEAwwLVlZTQy1Sb290LTExFDASBgorBgEEAYKifAIBDAQwMDAx"
vvsc_root_subject_key_id="21:B9:21:60:2D:53:8B:86:DA:A4:16:5C:AA:40:90:25:EB:FE:7E:28"

vvsc_ica_path="integration_tests/constants/vvsc_ica_cert_1"
vvsc_ica_subject="MIGXMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTETMBEGA1UEBwwKU29tZSBTdGF0ZTEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRMwEQYDVQQDDApWVlNDLUlDQS0xMRQwEgYKKwYBBAGConwCAQwEMDAwMQ=="
vvsc_ica_subject_key_id="98:4B:EE:D7:40:A2:FE:29:CB:AF:C0:0A:67:B7:AE:FF:12:A5:DA:DD"

vvsc_leaf_path="integration_tests/constants/vvsc_leaf_cert_1"
vvsc_leaf_subject="MIGYMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTETMBEGA1UEBwwKU29tZSBTdGF0ZTEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRQwEgYDVQQDDAtWVlNDLUxlYWYtMTEUMBIGCisGAQQBgqJ8AgEMBDAwMDE="
vvsc_leaf_subject_key_id="42:24:A6:34:C8:C1:2F:88:9D:9C:7F:BE:8A:7A:6E:40:DB:C8:2B:F1"

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

echo "$vendor_account adds a self-issued VVSC Root (Matter §6.4.5.4) as the"
echo "trust anchor for the VVSC leaf below — registered as VIDSignerPKI."
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-root-cert --certificate="$vvsc_root_path" --is-vid-verification-signer=true --from $vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "$vendor_account adds a VVSC ICA chained under the VVSC root"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-ica-cert --certificate="$vvsc_ica_path" --is-vid-verification-signer=true --from $vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "$vendor_account adds the VVSC Leaf (replaces the legacy NOC Leaf — full"
echo "chain VvscRoot1 → VvscIca1 → VvscLeaf1 is path length 3, the §6.4.10 cap)"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-ica-cert --certificate="$vvsc_leaf_path" --is-vid-verification-signer=true --from $vendor_account --yes)
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
check_response "$result" "\"subjectKeyId\": \"$vvsc_leaf_subject_key_id\""

echo "Request all approved DA certificates"
result=$(dcld query pki all-x509-certs)
echo $result | jq
check_response "$result" "\"subjectKeyId\": \"$da_root_subject_key_id\""
check_response "$result" "\"subjectKeyId\": \"$da_intermediate_subject_key_id\""
check_response "$result" "\"subjectKeyId\": \"$da_leaf_subject_key_id\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_subject_key_id\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_intermediate_subject_key_id\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$vvsc_leaf_subject_key_id\""

echo "Request all NOC certificates"
result=$(dcld query pki all-noc-x509-certs)
echo $result | jq
check_response "$result" "\"subjectKeyId\": \"$noc_root_subject_key_id\""
check_response "$result" "\"subjectKeyId\": \"$noc_intermediate_subject_key_id\""
check_response "$result" "\"subjectKeyId\": \"$vvsc_leaf_subject_key_id\""
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

echo "Request certificates by subject key id"
echo "Request DA certificate using global command"
result=$(dcld query pki cert --subject-key-id="$da_root_subject_key_id")
echo $result | jq
check_response "$result" "\"subjectKeyId\": \"$da_root_subject_key_id\""

echo "Request NOC certificate using global command"
result=$(dcld query pki cert --subject-key-id="$noc_root_subject_key_id")
echo $result | jq
check_response "$result" "\"subjectKeyId\": \"$noc_root_subject_key_id\""

echo "Request DA certificate"
result=$(dcld query pki x509-cert --subject-key-id="$da_root_subject_key_id")
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