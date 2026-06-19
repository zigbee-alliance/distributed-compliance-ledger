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

# See pki-noc-revocation-with-revoking-child.sh for the VVSC chain rationale.
# noc_leaf_cert_1 now points at the Matter §6.5.12 VVSC leaf chained under
# VvscRoot1 → VvscIca1 (path length 3, the §6.4.10 step 12.a.iii cap),
# submitted with --is-vid-verification-signer=true.
vvsc_root_path="integration_tests/constants/vvsc_root_cert_1"
vvsc_root_subject="MIGWMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTERMA8GA1UEBwwIVGFzaGtlbnQxGDAWBgNVBAoMD0V4YW1wbGUgQ29tcGFueTEZMBcGA1UECwwQVGVzdGluZyBEaXZpc2lvbjEUMBIGA1UEAwwLVlZTQy1Sb290LTExFDASBgorBgEEAYKifAIBDAQwMDAx"
vvsc_root_subject_key_id="21:B9:21:60:2D:53:8B:86:DA:A4:16:5C:AA:40:90:25:EB:FE:7E:28"
# VvscRootCert1Copy reuses VvscRootCert1's key (so the same Subject + SKID), but
# with a different serial number — used to re-establish an active VVSC root after
# section 1 soft-deletes VvscRootCert1 (the UniqueCertificate record keyed by
# Issuer+SerialNumber survives revocation, so re-adding the same serial fails
# with "certificate already exists").
vvsc_root_copy_path="integration_tests/constants/vvsc_root_cert_1_copy"
vvsc_ica_1_path="integration_tests/constants/vvsc_ica_cert_1"
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

echo "Add intermidiate NOC certificate by vendor with VID = $vid"
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

echo "Request NOC root certificate by VID = $vid and SKID=$noc_root_cert_1_subject_key_id"
result=$(dcld query pki noc-x509-cert --vid="$vid" --subject-key-id="$noc_root_cert_1_subject_key_id")
check_response "$result" "\"subject\": \"$noc_root_cert_1_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_copy_serial_number\""
echo $result | jq

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

echo "Try to revoke intermediate with invalid serialNumber"
result=$(echo "$passphrase" | dcld tx pki revoke-noc-x509-root-cert --subject="$noc_root_cert_1_subject" --subject-key-id="$noc_root_cert_1_subject_key_id" --serial-number="invalid" --from $vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 404"

echo "$vendor_account Vendor revokes root NOC certificate with serialNumber=$noc_root_cert_1_serial_number only, it should not revoke child certificates"
result=$(echo "$passphrase" | dcld tx pki revoke-noc-x509-root-cert --subject="$noc_root_cert_1_subject" --subject-key-id="$noc_root_cert_1_subject_key_id" --serial-number="$noc_root_cert_1_serial_number" --from=$vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Request all revoked NOC root certificates should contain root certificate with serialNumber=$noc_root_cert_1_serial_number"
result=$(dcld query pki all-revoked-noc-x509-root-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$noc_root_cert_1_subject"
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_copy_serial_number\""
response_does_not_contain "$result" "\"subject\": \"$noc_cert_1_subject\""
response_does_not_contain "$result" "\"subject\": \"$noc_leaf_cert_1_subject\""

echo "Request revoked NOC root certificate by subject and subjectKeyId should contain root certificate with serialNumber=$noc_root_cert_1_serial_number"
result=$(dcld query pki revoked-noc-x509-root-cert --subject="$noc_root_cert_1_subject" --subject-key-id="$noc_root_cert_1_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$noc_root_cert_1_subject"
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_copy_serial_number\""

echo "Request all x509 root revoked certificates should not contain revoked NOC root certificates"
result=$(dcld query pki all-revoked-x509-root-certs)
response_does_not_contain "$result" "\"subject\": \"$noc_root_cert_1_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
echo $result | jq

echo "Request NOC root certificate by VID = $vid should contain only one root certificate with serialNumber=$noc_root_cert_1_copy_serial_number"
result=$(dcld query pki noc-x509-root-certs --vid="$vid")
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_copy_serial_number\""
check_response "$result" "\"subject\": \"$noc_root_cert_1_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
echo $result | jq

echo "Request NOC root certificate by VID = $vid and SKID=$noc_root_cert_1_subject_key_id should contain only one root certificate with serialNumber=$noc_root_cert_1_copy_serial_number"
result=$(dcld query pki noc-x509-cert --vid="$vid" --subject-key-id="$noc_root_cert_1_subject_key_id")
check_response "$result" "\"subject\": \"$noc_root_cert_1_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_copy_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
echo $result | jq

echo "Request all certificates by subject should not be empty"
result=$(dcld query pki all-noc-subject-x509-certs --subject="$noc_root_cert_1_subject")
check_response "$result" "\"$noc_root_cert_1_subject\""
check_response "$result" "\"$noc_root_cert_1_subject_key_id\""
echo $result | jq

echo "Request all certificates by subjectKeyId should contain only one root certificate with serialNumber=$noc_root_cert_1_copy_serial_number"
result=$(dcld query pki noc-x509-cert --subject-key-id="$noc_root_cert_1_subject_key_id")
check_response "$result" "\"subject\": \"$noc_root_cert_1_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_copy_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
echo $result | jq

echo "Request NOC certificate by VID = $vid should contain intermediate and leaf certificates"
result=$(dcld query pki noc-x509-ica-certs --vid="$vid")
echo $result | jq
check_response "$result" "\"subject\": \"$noc_cert_1_subject\""
check_response "$result" "\"subject\": \"$noc_cert_1_subject\""
check_response "$result" "\"subject\": \"$noc_leaf_cert_1_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_cert_1_subject_key_id\""
check_response "$result" "\"subjectKeyId\": \"$noc_leaf_cert_1_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_cert_1_serial_number\""
check_response "$result" "\"serialNumber\": \"$noc_leaf_cert_1_serial_number\""

echo "Request all approved certificates should not contain revoked NOC root certificate"
result=$(dcld query pki all-noc-x509-certs)
check_response "$result" "\"subject\": \"$noc_root_cert_1_subject\""
check_response "$result" "\"subject\": \"$noc_cert_1_subject\""
check_response "$result" "\"subject\": \"$noc_leaf_cert_1_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
check_response "$result" "\"subjectKeyId\": \"$noc_cert_1_subject_key_id\""
check_response "$result" "\"subjectKeyId\": \"$noc_leaf_cert_1_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_copy_serial_number\""
check_response "$result" "\"serialNumber\": \"$noc_cert_1_serial_number\""
check_response "$result" "\"serialNumber\": \"$noc_leaf_cert_1_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
echo $result | jq

test_divider

echo "$vendor_account Vendor revokes second root NOC certificate by serialNumber with \"revoke-child\" flag set to true, it should remove child certificates too"
result=$(echo "$passphrase" | dcld tx pki revoke-noc-x509-root-cert --subject="$noc_root_cert_1_subject" --subject-key-id="$noc_root_cert_1_subject_key_id" --serial-number="$noc_root_cert_1_copy_serial_number" --revoke-child=true --from=$vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Also revoke the VVSC root with revoke-child=true. The VVSC chain"
echo "VvscRoot1 → VvscIca1 → VvscLeaf1 is structurally disjoint from the"
echo "OperationalPKI cascade (Matter §6.5.12 / §6.4.10) — without an explicit"
echo "VVSC root revocation the leaf would remain active and the revoked-ICA"
echo "assertion below would fail."
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

echo "Request all revoked NOC root certificates should contain one intermediate and one leaf certificates"
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

echo "Request NOC root certificate by VID = $vid should be empty"
result=$(dcld query pki noc-x509-root-certs --vid="$vid")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$noc_root_cert_1_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_copy_serial_number\""
echo $result | jq

echo "Request all certificates by subject should be empty"
result=$(dcld query pki all-noc-subject-x509-certs --subject="$noc_root_cert_1_subject")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"$noc_root_cert_1_subject\""
response_does_not_contain "$result" "\"$noc_root_cert_1_subject_key_id\""
echo $result | jq

echo "Request all certificates by subjectKeyId should be empty"
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
response_does_not_contain "$result" "\"subject\": \"$noc_cert_1_subject\""
response_does_not_contain "$result" "\"subject\": \"$noc_leaf_cert_1_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_leaf_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_cert_1_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_leaf_cert_1_serial_number\""

echo "Request all approved certificates should be empty"
result=$(dcld query pki all-noc-x509-certs)
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

test_divider

echo "REVOCATION OF NOC NON-ROOT CERTIFICATES"

echo "Add NOC root certificate by vendor with VID = $vid"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-root-cert --certificate="$noc_root_cert_2_path" --from $vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Add first intermidiate NOC certificate by vendor with VID = $vid"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-ica-cert --certificate="$noc_cert_2_path" --from $vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Add second intermidiate NOC certificate by vendor with VID = $vid"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-ica-cert --certificate="$noc_cert_2_copy_path" --from $vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Re-establish an active VVSC root via VvscRootCert1Copy. Revocation is a"
echo "soft delete (cert moves to the revoked list) but the (Issuer, SerialNumber)"
echo "UniqueCertificate record survives — re-adding the same PEM would fail with"
echo "ErrCertificateAlreadyExists. The Copy shares VvscRoot1's key (same Subject"
echo "and SubjectKeyID) so VvscIca2's AuthorityKeyID still resolves to a present"
echo "VIDSignerPKI entry during verifyVVSCCertificate's chain walk."
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-root-cert --certificate="$vvsc_root_copy_path" --is-vid-verification-signer=true --from $vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Pre-seed VvscIca2 under the VVSC root so the leaf-2 chain"
echo "VvscRoot1 → VvscIca2 → noc_leaf_cert_2 resolves through verifyVVSCCertificate."
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

echo "Request all NOC certificates"
result=$(dcld query pki all-noc-x509-ica-certs)
echo $result | jq
check_response "$result" "\"serialNumber\": \"$noc_cert_2_serial_number\""
check_response "$result" "\"serialNumber\": \"$noc_cert_2_copy_serial_number\""
check_response "$result" "\"serialNumber\": \"$noc_leaf_cert_2_serial_number\""

echo "Try to revoke intermediate with invalid serialNumber"
result=$(echo "$passphrase" | dcld tx pki revoke-noc-x509-ica-cert --subject="$noc_cert_2_subject" --subject-key-id="$noc_cert_2_subject_key_id" --serial-number="invalid" --from $vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 404"

echo "$vendor_account Vendor revokes NOC certificate with serialNumber=$noc_cert_2_serial_number only, it should not revoke child certificates"
result=$(echo "$passphrase" | dcld tx pki revoke-noc-x509-ica-cert --subject="$noc_cert_2_subject" --subject-key-id="$noc_cert_2_subject_key_id" --serial-number="$noc_cert_2_serial_number" --from=$vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Request all revoked noc ica certificates should contain one intermediate certificate only"
result=$(dcld query pki all-revoked-noc-x509-ica-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$noc_cert_2_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_cert_2_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_cert_2_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_cert_2_copy_serial_number\""
response_does_not_contain "$result" "\"subject\": \"$noc_leaf_cert_2_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_leaf_cert_2_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_leaf_cert_2_serial_number\""

echo "Request all certificates by NOC certificate's subject should not be empty"
result=$(dcld query pki all-noc-subject-x509-certs --subject="$noc_cert_2_subject")
check_response "$result" "\"$noc_cert_2_subject\""
check_response "$result" "\"$noc_cert_2_subject_key_id\""
echo $result | jq

echo "Request all certificates by NOC certificate's subjectKeyId should not be empty"
result=$(dcld query pki noc-x509-cert --subject-key-id="$noc_cert_2_subject_key_id")
check_response "$result" "\"subject\": \"$noc_cert_2_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_cert_2_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_cert_2_copy_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_cert_2_serial_number\""
echo $result | jq

echo "Request NOC certificate by VID = $vid should contain one intermediate and leaf certificates"
result=$(dcld query pki noc-x509-ica-certs --vid="$vid")
echo $result | jq
check_response "$result" "\"subject\": \"$noc_cert_2_subject\""
check_response "$result" "\"subject\": \"$noc_leaf_cert_2_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_cert_2_subject_key_id\""
check_response "$result" "\"subjectKeyId\": \"$noc_leaf_cert_2_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_cert_2_copy_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_cert_2_serial_number\""

echo "Request all approved certificates should contain one intermediate and leaf certificates"
result=$(dcld query pki all-noc-x509-certs)
check_response "$result" "\"subject\": \"$noc_root_cert_2_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_2_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_2_serial_number\""
check_response "$result" "\"subject\": \"$noc_cert_2_subject\""
check_response "$result" "\"subject\": \"$noc_leaf_cert_2_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_cert_2_subject_key_id\""
check_response "$result" "\"subjectKeyId\": \"$noc_leaf_cert_2_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_leaf_cert_2_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_cert_2_serial_number\""
echo $result | jq

echo "$vendor_account Vendor revokes NOC certificate with serialNumber=$noc_cert_2_serial_number with \"revoke-child\" flag set to true, it should revoke child certificates too"
result=$(echo "$passphrase" | dcld tx pki revoke-noc-x509-ica-cert --subject="$noc_cert_2_subject" --subject-key-id="$noc_cert_2_subject_key_id" --serial-number="$noc_cert_2_copy_serial_number" --revoke-child=true --from=$vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Also revoke VvscIca2 with revoke-child=true so the VVSC leaf is cascaded."
echo "noc_leaf_cert_2 now lives in the VVSC chain (VvscRoot1 → VvscIca2 →"
echo "VvscLeaf2), which is disjoint from the OperationalPKI chain that the"
echo "previous revoke just walked."
result=$(echo "$passphrase" | dcld tx pki revoke-noc-x509-ica-cert --subject="$vvsc_ica_2_subject" --subject-key-id="$vvsc_ica_2_subject_key_id" --revoke-child=true --from=$vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Request all revoked certificates should contain two intermediate and one leaf certificates"
result=$(dcld query pki all-revoked-noc-x509-ica-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$noc_cert_2_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_cert_2_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_cert_2_serial_number\""
check_response "$result" "\"serialNumber\": \"$noc_cert_2_copy_serial_number\""
check_response "$result" "\"subject\": \"$noc_leaf_cert_2_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_leaf_cert_2_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_leaf_cert_2_serial_number\""

echo "Request all certificates by NOC certificate's subject should be empty"
result=$(dcld query pki all-noc-subject-x509-certs --subject="$noc_cert_2_subject")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"$noc_cert_2_subject\""
response_does_not_contain "$result" "\"$noc_cert_2_subject_key_id\""
echo $result | jq

echo "Request all certificates by NOC certificate's subjectKeyId should be empty"
result=$(dcld query pki noc-x509-cert --subject-key-id="$noc_cert_2_subject_key_id")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$noc_cert_2_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_cert_2_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_cert_2_copy_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_cert_2_serial_number\""
echo $result | jq

echo "Request NOC certificate by VID = $vid should be empty"
result=$(dcld query pki noc-x509-ica-certs --vid="$vid")
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$noc_cert_2_subject\""
response_does_not_contain "$result" "\"subject\": \"$noc_leaf_cert_2_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_cert_2_subject_key_id\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_leaf_cert_2_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_cert_2_copy_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_cert_2_serial_number\""

echo "Request all approved certificates should contain only one root certificate"
result=$(dcld query pki all-noc-x509-certs)
check_response "$result" "\"subject\": \"$noc_root_cert_2_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_2_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_2_serial_number\""
response_does_not_contain "$result" "\"subject\": \"$noc_cert_2_subject\""
response_does_not_contain "$result" "\"subject\": \"$noc_leaf_cert_2_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_cert_2_subject_key_id\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_leaf_cert_2_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_cert_2_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_leaf_cert_2_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_cert_2_serial_number\""
echo $result | jq

test_divider