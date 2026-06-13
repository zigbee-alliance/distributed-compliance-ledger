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

noc_root_cert_1_path="integration_tests/constants/noc_root_cert_1"
noc_root_cert_1_subject="MHoxCzAJBgNVBAYTAlVaMRMwEQYDVQQIEwpTb21lIFN0YXRlMREwDwYDVQQHEwhUYXNoa2VudDEYMBYGA1UEChMPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLExBUZXN0aW5nIERpdmlzaW9uMQ4wDAYDVQQDEwVOT0MtMQ=="
noc_root_cert_1_subject_key_id="0E:10:B8:5D:96:7A:08:33:C7:C5:44:49:0E:28:0F:C1:6E:D5:D4:7C"
noc_root_cert_1_serial_number="313831573505791137291636389937677533381171619492"

noc_root_cert_2_path="integration_tests/constants/noc_root_cert_2"
noc_root_cert_2_subject="MHoxCzAJBgNVBAYTAlVaMRMwEQYDVQQIEwpTb21lIFN0YXRlMREwDwYDVQQHEwhUYXNoa2VudDEYMBYGA1UEChMPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLExBUZXN0aW5nIERpdmlzaW9uMQ4wDAYDVQQDEwVOT0MtMg=="
noc_root_cert_2_subject_key_id="46:C0:B0:74:0C:63:C8:9E:E0:5C:14:C2:71:62:F8:67:24:5C:8E:29"
noc_root_cert_2_serial_number="727423814323052015089749828769570958840545369270"

noc_root_cert_1_copy_path="integration_tests/constants/noc_root_cert_1_copy"
noc_root_cert_1_copy_serial_number="12722088350714347345576486793058060481880825999"

noc_cert_1_path="integration_tests/constants/noc_cert_1"
noc_cert_1_subject="MIGCMQswCQYDVQQGEwJVWjETMBEGA1UECBMKU29tZSBTdGF0ZTETMBEGA1UEBxMKU29tZSBTdGF0ZTEYMBYGA1UEChMPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLExBUZXN0aW5nIERpdmlzaW9uMRQwEgYDVQQDEwtOT0MtY2hpbGQtMQ=="
noc_cert_1_subject_key_id="06:9F:5A:E0:1F:23:3E:9F:C7:4F:B6:F9:A2:33:47:33:62:7A:07:C5"
noc_cert_1_serial_number="577430346509479530103103319788179390906984119670"

# Matter §6.5.12 §6.4.10 chain: VvscRoot1 → VvscIca1 → noc_leaf_cert_1 (
# repurposed as the VVSC leaf so downstream assertions keep their names).
# noc_leaf_cert_1 was originally a NOC end-entity chained under noc_cert_1;
# the strict §6.5.12 ICA handler no longer accepts that profile. The VVSC
# chain runs in parallel with the OperationalPKI chain (noc_root_cert_1 →
# noc_cert_1) so the existing assertions about each chain still hold.
vvsc_root_path="integration_tests/constants/vvsc_root_cert_1"
vvsc_root_subject="MIGWMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTERMA8GA1UEBwwIVGFzaGtlbnQxGDAWBgNVBAoMD0V4YW1wbGUgQ29tcGFueTEZMBcGA1UECwwQVGVzdGluZyBEaXZpc2lvbjEUMBIGA1UEAwwLVlZTQy1Sb290LTExFDASBgorBgEEAYKifAIBDAQwMDAx"
vvsc_root_subject_key_id="21:B9:21:60:2D:53:8B:86:DA:A4:16:5C:AA:40:90:25:EB:FE:7E:28"
vvsc_root_serial_number="5068329979261235249"
vvsc_ica_1_path="integration_tests/constants/vvsc_ica_cert_1"
vvsc_ica_1_subject="MIGXMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTETMBEGA1UEBwwKU29tZSBTdGF0ZTEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRMwEQYDVQQDDApWVlNDLUlDQS0xMRQwEgYKKwYBBAGConwCAQwEMDAwMQ=="
vvsc_ica_1_subject_key_id="98:4B:EE:D7:40:A2:FE:29:CB:AF:C0:0A:67:B7:AE:FF:12:A5:DA:DD"
vvsc_ica_1_serial_number="5068329979109130545"
noc_leaf_cert_1_path="integration_tests/constants/vvsc_leaf_cert_1"
noc_leaf_cert_1_subject="MIGYMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTETMBEGA1UEBwwKU29tZSBTdGF0ZTEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRQwEgYDVQQDDAtWVlNDLUxlYWYtMTEUMBIGCisGAQQBgqJ8AgEMBDAwMDE="
noc_leaf_cert_1_subject_key_id="42:24:A6:34:C8:C1:2F:88:9D:9C:7F:BE:8A:7A:6E:40:DB:C8:2B:F1"
noc_leaf_cert_1_serial_number="5068329979159654449"

noc_cert_2_path="integration_tests/constants/noc_cert_2"
noc_cert_2_subject="MIGCMQswCQYDVQQGEwJVWjETMBEGA1UECBMKU29tZSBTdGF0ZTETMBEGA1UEBxMKU29tZSBTdGF0ZTEYMBYGA1UEChMPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLExBUZXN0aW5nIERpdmlzaW9uMRQwEgYDVQQDEwtOT0MtY2hpbGQtMg=="
noc_cert_2_subject_key_id="17:E2:72:19:E1:7F:19:D7:0D:02:1A:B0:40:7B:04:26:CC:D4:2B:F5"
noc_cert_2_serial_number="634591262660314610068979921875981241084684028375"

noc_cert_2_copy_path="integration_tests/constants/noc_cert_2_copy"
noc_cert_2_copy_serial_number="252687488758567844896720928536709119387931444024"

# Second VVSC chain (VvscRoot1 → VvscIca2 → VvscLeaf2) repurposed as
# noc_leaf_cert_2 so the second cascade-revoke section keeps its assertions.
vvsc_ica_2_path="integration_tests/constants/vvsc_ica_cert_2"
vvsc_ica_2_subject="MIGXMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTETMBEGA1UEBwwKU29tZSBTdGF0ZTEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRMwEQYDVQQDDApWVlNDLUlDQS0yMRQwEgYKKwYBBAGConwCAQwEMDAwMQ=="
vvsc_ica_2_subject_key_id="ED:8C:5B:36:E7:3C:E4:54:09:A2:59:D4:E8:0A:D6:6C:99:C6:A2:CC"
vvsc_ica_2_serial_number="5068329979109130546"
noc_leaf_cert_2_path="integration_tests/constants/vvsc_leaf_cert_2"
noc_leaf_cert_2_subject="MIGYMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTETMBEGA1UEBwwKU29tZSBTdGF0ZTEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRQwEgYDVQQDDAtWVlNDLUxlYWYtMjEUMBIGCisGAQQBgqJ8AgEMBDAwMDE="
noc_leaf_cert_2_subject_key_id="8D:F6:2A:9C:24:D0:92:36:83:32:38:47:35:3A:0B:E9:19:CD:90:B3"
noc_leaf_cert_2_serial_number="5068329979159654450"

vid_in_hex_format=0x6006
vid=24582

echo "REVOCATION OF NOC ROOT CERTIFICATES"

vendor_account=vendor_account_$vid_in_hex_format
echo "Create Vendor account - $vendor_account"
create_new_vendor_account $vendor_account $vid_in_hex_format

test_divider

echo "Add first NOC root certificate by vendor with VID = $vid"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-root-cert --certificate="$noc_root_cert_1_path" --from $vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Add second NOC root certificate by vendor with VID = $vid"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-root-cert --certificate="$noc_root_cert_1_copy_path" --from $vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Add first NOC certificate by vendor with VID = $vid"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-ica-cert --certificate="$noc_cert_1_path" --from $vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Pre-seed the VVSC chain (Matter §6.4.5.4) so the leaf below has a"
echo "§6.4.10 step 12.a.iii path-length-3 chain to validate against."
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-root-cert --certificate="$vvsc_root_path" --is-vid-verification-signer=true --from $vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

result=$(echo "$passphrase" | dcld tx pki add-noc-x509-ica-cert --certificate="$vvsc_ica_1_path" --is-vid-verification-signer=true --from $vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Add VVSC leaf certificate (replaces the legacy NocLeafCert1) by vendor with VID = $vid"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-ica-cert --certificate="$noc_leaf_cert_1_path" --is-vid-verification-signer=true --from $vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Request All NOC root certificate"
result=$(dcld query pki all-noc-x509-root-certs)
echo $result | jq
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_copy_serial_number\""

echo "Request all NOC certificates"
result=$(dcld query pki all-noc-x509-ica-certs)
echo $result | jq
check_response "$result" "\"serialNumber\": \"$noc_cert_1_serial_number\""
check_response "$result" "\"serialNumber\": \"$noc_leaf_cert_1_serial_number\""

echo "$vendor_account Vendor revokes the OperationalPKI root NOC certificate with revoke-child=true."
echo "Cascade hits NocCert1 only — VVSC chain is structurally disjoint (Matter §6.5.12)."
result=$(echo "$passphrase" | dcld tx pki revoke-noc-x509-root-cert --subject="$noc_root_cert_1_subject" --subject-key-id="$noc_root_cert_1_subject_key_id" --revoke-child=true --from=$vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Also revoke the VVSC root with revoke-child=true so the cascade also picks up VvscIca1 + the VVSC leaf."
result=$(echo "$passphrase" | dcld tx pki revoke-noc-x509-root-cert --subject="$vvsc_root_subject" --subject-key-id="$vvsc_root_subject_key_id" --revoke-child=true --from=$vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Request all revoked NOC root certificates should contain two root certificates"
result=$(dcld query pki all-revoked-noc-x509-root-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$noc_root_cert_1_subject"
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_copy_serial_number\""

echo "Request all revoked NOC ica certificates should contain one intermediate and one leaf certificates"
result=$(dcld query pki all-revoked-noc-x509-ica-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$noc_cert_1_subject\""
check_response "$result" "\"subject\": \"$noc_leaf_cert_1_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_cert_1_subject_key_id\""
check_response "$result" "\"subjectKeyId\": \"$noc_leaf_cert_1_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_cert_1_serial_number\""
check_response "$result" "\"serialNumber\": \"$noc_leaf_cert_1_serial_number\""

echo "Request revoked NOC root certificate by subject and subjectKeyId should contain two root certificates"
result=$(dcld query pki revoked-noc-x509-root-cert --subject="$noc_root_cert_1_subject" --subject-key-id="$noc_root_cert_1_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$noc_root_cert_1_subject"
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_copy_serial_number\""

echo "Request all x509 root revoked certificates should not contain revoked NOC root certificates"
result=$(dcld query pki all-revoked-x509-root-certs)
response_does_not_contain "$result" "\"subject\": \"$noc_root_cert_1_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_copy_serial_number\""
echo $result | jq

echo "Request NOC root certificate by VID = $vid should be empty"
result=$(dcld query pki noc-x509-root-certs --vid="$vid")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$noc_root_cert_1_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_copy_serial_number\""
echo $result | jq

echo "Request all certificates by NOC root certificate's subject should be empty"
result=$(dcld query pki all-noc-subject-x509-certs --subject="$noc_root_cert_1_subject")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"$noc_root_cert_1_subject\""
response_does_not_contain "$result" "\"$noc_root_cert_1_subject_key_id\""
echo $result | jq

echo "Request all certificates by NOC root certificate's subjectKeyId should be empty"
result=$(dcld query pki noc-x509-cert --subject-key-id="$noc_root_cert_1_subject_key_id")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$noc_root_cert_1_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_copy_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
echo $result | jq

echo "Request NOC certificate by VID = $vid should be empty"
result=$(dcld query pki noc-x509-ica-certs --vid="$vid")
echo $result | jq
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$noc_cert_1_subject\""
response_does_not_contain "$result" "\"subject\": \"$noc_leaf_cert_1_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_leaf_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_cert_1_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_leaf_cert_1_serial_number\""

echo "Request all approved certificates should be empty"
result=$(dcld query pki all-x509-certs)
check_response "$result" "\[\]"
response_does_not_contain "$result" "\"subject\": \"$noc_root_cert_1_subject\""
response_does_not_contain "$result" "\"subject\": \"$noc_cert_1_subject\""
response_does_not_contain "$result" "\"subject\": \"$noc_leaf_cert_1_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_leaf_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_copy_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_cert_1_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_leaf_cert_1_serial_number\""
echo $result | jq

echo "Request NOC certificate by VID = $vid and SKID = $noc_cert_1_subject_key_id should be empty"
result=$(dcld query pki noc-x509-cert --vid="$vid" --subject-key-id="$noc_cert_1_subject_key_id")
echo $result | jq
check_response "$result" "Not Found"

echo "Request NOC certificate by VID = $vid and SKID = $noc_leaf_cert_1_subject_key_id should be empty"
result=$(dcld query pki noc-x509-cert --vid="$vid" --subject-key-id="$noc_leaf_cert_1_subject_key_id")
echo $result | jq
check_response "$result" "Not Found"

test_divider

echo "REVOCATION OF NOC NON-ROOT CERTIFICATES"

echo "Add NOC root certificate by vendor with VID = $vid"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-root-cert --certificate="$noc_root_cert_2_path" --from $vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Add intermidiate NOC certificate by vendor with VID = $vid"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-ica-cert --certificate="$noc_cert_2_path" --from $vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Add second intermidiate NOC certificate by vendor with VID = $vid"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-ica-cert --certificate="$noc_cert_2_copy_path" --from $vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Re-add the self-issued VVSC root. Revocation is a soft delete (cert moves"
echo "to the revoked list), so the active store no longer contains it and"
echo "verifyVVSCCertificate would fail the chain walk for VvscIca2."
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-root-cert --certificate="$vvsc_root_path" --is-vid-verification-signer=true --from $vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Pre-seed the second VVSC intermediate (VvscIca2 under VvscRoot1) so the"
echo "leaf-2 chain VvscRoot1 → VvscIca2 → noc_leaf_cert_2 resolves through"
echo "verifyVVSCCertificate."
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-ica-cert --certificate="$vvsc_ica_2_path" --is-vid-verification-signer=true --from $vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Add VVSC leaf certificate 2 (replaces the legacy NocLeafCert2) by vendor with VID = $vid"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-ica-cert --certificate="$noc_leaf_cert_2_path" --is-vid-verification-signer=true --from $vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Request All NOC root certificate"
result=$(dcld query pki all-noc-x509-root-certs)
echo $result | jq
check_response "$result" "\"serialNumber\": \"$noc_root_cert_2_serial_number\""

echo "Request all intermidiate NOC certificates"
result=$(dcld query pki all-noc-x509-ica-certs)
echo $result | jq
check_response "$result" "\"serialNumber\": \"$noc_cert_2_serial_number\""
check_response "$result" "\"serialNumber\": \"$noc_cert_2_copy_serial_number\""
check_response "$result" "\"serialNumber\": \"$noc_leaf_cert_2_serial_number\""

echo "$vendor_account Vendor revokes the OperationalPKI ICA with revoke-child=true."
echo "Cascade hits NocCert2 / NocCert2Copy but, per Matter §6.5.12, does not reach the VVSC leaf."
result=$(echo "$passphrase" | dcld tx pki revoke-noc-x509-ica-cert --subject="$noc_cert_2_subject" --subject-key-id="$noc_cert_2_subject_key_id" --revoke-child=true --from=$vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Also revoke VvscIca2 with revoke-child=true so the cascade also picks up the VVSC leaf 2."
result=$(echo "$passphrase" | dcld tx pki revoke-noc-x509-ica-cert --subject="$vvsc_ica_2_subject" --subject-key-id="$vvsc_ica_2_subject_key_id" --revoke-child=true --from=$vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Request all revoked certificates should two intermediate and one leaf certificates"
result=$(dcld query pki all-revoked-noc-x509-ica-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$noc_cert_2_subject\""
check_response "$result" "\"subject\": \"$noc_leaf_cert_2_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_cert_2_subject_key_id\""
check_response "$result" "\"subjectKeyId\": \"$noc_leaf_cert_2_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_cert_2_serial_number\""
check_response "$result" "\"serialNumber\": \"$noc_cert_2_copy_serial_number\""
check_response "$result" "\"serialNumber\": \"$noc_leaf_cert_2_serial_number\""
response_does_not_contain "$result" "\"subject\": \"$noc_root_cert_2_subject"
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_2_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_2_serial_number\""

echo "Request all certificates by NOC certificate's subject should be empty"
result=$(dcld query pki all-noc-subject-x509-certs --subject="$noc_cert_2_subject")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"$noc_cert_1_subject\""
response_does_not_contain "$result" "\"$noc_cert_1_subject_key_id\""
echo $result | jq

echo "Request all certificates by NOC certificate's subjectKeyId should be empty"
result=$(dcld query pki noc-x509-cert --subject-key-id="$noc_cert_2_subject_key_id")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$noc_cert_2_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_cert_2_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_cert_2_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_cert_2_copy_serial_number\""
echo $result | jq

echo "Request NOC certificate by VID = $vid should not contain intermediate and leaf certificates"
result=$(dcld query pki noc-x509-ica-certs --vid="$vid")
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$noc_cert_2_subject\""
response_does_not_contain "$result" "\"subject\": \"$noc_leaf_cert_2_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_cert_2_subject_key_id\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_leaf_cert_2_subject_key_id\""

echo "Request all noc certificates should not contain intermediate and leaf certificates"
result=$(dcld query pki all-noc-x509-certs)
check_response "$result" "\"subject\": \"$noc_root_cert_2_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_2_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_2_serial_number\""
response_does_not_contain "$result" "\"subject\": \"$noc_cert_2_subject\""
response_does_not_contain "$result" "\"subject\": \"$noc_leaf_cert_2_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_cert_2_subject_key_id\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_leaf_cert_2_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_cert_2_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_cert_2_copy_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_leaf_cert_2_serial_number\""
echo $result | jq

echo "Request NOC certificate by VID = $vid and SKID = $noc_root_cert_2_subject_key_id should not be empty"
result=$(dcld query pki noc-x509-cert --vid="$vid" --subject-key-id="$noc_root_cert_2_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$noc_root_cert_2_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_2_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_2_serial_number\""

echo "Request NOC certificate by VID = $vid and SKID = $noc_cert_2_subject_key_id should be empty"
result=$(dcld query pki noc-x509-cert --vid="$vid" --subject-key-id="$noc_cert_2_subject_key_id")
echo $result | jq
check_response "$result" "Not Found"

echo "Request NOC certificate by VID = $vid and SKID = $noc_leaf_cert_2_subject_key_id should be empty"
result=$(dcld query pki noc-x509-cert --vid="$vid" --subject-key-id="$noc_leaf_cert_2_subject_key_id")
echo $result | jq
check_response "$result" "Not Found"

test_divider