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

root_cert_1_path="integration_tests/constants/noc_root_cert_1"
root_cert_subject="MHoxCzAJBgNVBAYTAlVaMRMwEQYDVQQIEwpTb21lIFN0YXRlMREwDwYDVQQHEwhUYXNoa2VudDEYMBYGA1UEChMPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLExBUZXN0aW5nIERpdmlzaW9uMQ4wDAYDVQQDEwVOT0MtMQ=="
root_cert_subject_key_id="0E:10:B8:5D:96:7A:08:33:C7:C5:44:49:0E:28:0F:C1:6E:D5:D4:7C"
root_cert_1_serial_number="313831573505791137291636389937677533381171619492"

root_cert_1_copy_path="integration_tests/constants/noc_root_cert_1_copy"
root_cert_1_copy_serial_number="12722088350714347345576486793058060481880825999"

root_cert_vid=65521

intermediate_cert_1_path="integration_tests/constants/noc_cert_1"
intermediate_cert_2_path="integration_tests/constants/noc_cert_1_copy"
intermediate_cert_subject="MIGCMQswCQYDVQQGEwJVWjETMBEGA1UECBMKU29tZSBTdGF0ZTETMBEGA1UEBxMKU29tZSBTdGF0ZTEYMBYGA1UEChMPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLExBUZXN0aW5nIERpdmlzaW9uMRQwEgYDVQQDEwtOT0MtY2hpbGQtMQ=="
intermediate_cert_subject_key_id="06:9F:5A:E0:1F:23:3E:9F:C7:4F:B6:F9:A2:33:47:33:62:7A:07:C5"
intermediate_cert_1_serial_number="577430346509479530103103319788179390906984119670"
intermediate_cert_2_serial_number="617357865778805507017637943649984133152592305888"

leaf_cert_path="integration_tests/constants/noc_leaf_cert_1"
leaf_cert_subject="MIGBMQswCQYDVQQGEwJVWjETMBEGA1UECBMKU29tZSBTdGF0ZTETMBEGA1UEBxMKU29tZSBTdGF0ZTEYMBYGA1UEChMPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLExBUZXN0aW5nIERpdmlzaW9uMRMwEQYDVQQDEwpOT0MtbGVhZi0x"
leaf_cert_subject_key_id="F0:3A:A5:96:8F:60:63:F0:15:21:24:0C:DA:0A:E6:2B:CC:A0:58:F9"
leaf_cert_serial_number="201294310322324358101935163941973786245732555938"

trustee_account="jack"

test_divider

echo "REMOVING the NOC and ICA CERTIFICATES"

vendor_account_65521=vendor_account_$root_cert_vid
echo "Create Vendor account - $vendor_account_65521"
create_new_vendor_account $vendor_account_65521 $root_cert_vid

vendor_account_65522=vendor_account_65522
echo "Create Vendor account - $vendor_account_65522"
create_new_vendor_account $vendor_account_65522 65522

echo "Add first NOC root certificate"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-root-cert --certificate="$root_cert_1_path" --from $vendor_account_65521 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Add first an ICA certificate"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-ica-cert --certificate="$intermediate_cert_1_path" --from $vendor_account_65521 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Add second an ICA certificate"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-ica-cert --certificate="$intermediate_cert_2_path" --from $vendor_account_65521 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Add a leaf ICA certificate"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-ica-cert --certificate="$leaf_cert_path" --from $vendor_account_65521 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Request all approved certificates."
result=$(dcld query pki all-noc-x509-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
check_response "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
check_response "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$root_cert_1_serial_number\""
check_response "$result" "\"serialNumber\": \"$intermediate_cert_1_serial_number\""
check_response "$result" "\"serialNumber\": \"$intermediate_cert_2_serial_number\""
check_response "$result" "\"serialNumber\": \"$leaf_cert_serial_number\""

echo "Revoke an ICA certificate with serialNumber $intermediate_cert_1_serial_number"
result=$(echo "$passphrase" | dcld tx pki revoke-noc-x509-ica-cert --subject="$intermediate_cert_subject" --subject-key-id="$intermediate_cert_subject_key_id" --serial-number="$intermediate_cert_1_serial_number" --from=$vendor_account_65521 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Request all revoked certificates should contain only one intermediate ICA certificate with serialNumber $intermediate_cert_1_serial_number"
result=$(dcld query pki all-revoked-noc-x509-ica-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$intermediate_cert_1_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$intermediate_cert_2_serial_number\""

echo "Remove intermediate ICA certificate with invalid serialNumber"
result=$(echo "$passphrase" | dcld tx pki remove-noc-x509-ica-cert --subject="$intermediate_cert_subject" --subject-key-id="$intermediate_cert_subject_key_id" --serial-number="invalid" --from=$vendor_account_65521 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 404"

echo "Try to remove the intermediate ICA certificate when sender is not Vendor account"
result=$(echo "$passphrase" | dcld tx pki remove-noc-x509-ica-cert --subject="$intermediate_cert_subject" --subject-key-id="$intermediate_cert_subject_key_id" --serial-number="$intermediate_cert_1_serial_number" --from=$trustee_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 4"

echo "Try to remove the intermediate ICA certificate using a vendor account with other VID"
result=$(echo "$passphrase" | dcld tx pki remove-noc-x509-ica-cert --subject="$intermediate_cert_subject" --subject-key-id="$intermediate_cert_subject_key_id" --serial-number="$intermediate_cert_1_serial_number" --from=$vendor_account_65522 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 4"

echo "Remove revoked intermediate ICA certificate with serialNumber $intermediate_cert_1_serial_number"
result=$(echo "$passphrase" | dcld tx pki remove-noc-x509-ica-cert --subject="$intermediate_cert_subject" --subject-key-id="$intermediate_cert_subject_key_id" --serial-number="$intermediate_cert_1_serial_number" --from=$vendor_account_65521 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Request all certificates should not contain intermediate ICA certificate with serialNumber $intermediate_cert_1_serial_number"
result=$(dcld query pki all-noc-x509-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
check_response "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
check_response "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$root_cert_1_serial_number\""
check_response "$result" "\"serialNumber\": \"$intermediate_cert_2_serial_number\""
check_response "$result" "\"serialNumber\": \"$leaf_cert_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$intermediate_cert_1_serial_number\""

echo "Request ICA certificates by VID should contain one ICA and leaf certificates"
result=$(dcld query pki noc-x509-ica-certs --vid="$root_cert_vid")
echo $result | jq
check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
check_response "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$intermediate_cert_2_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$intermediate_cert_1_serial_number\""

echo "Request approved certificates by an intermediate certificate's subject and subjectKeyId should contain only one certificate with serialNumber $intermediate_cert_2_serial_number"
result=$(dcld query pki noc-x509-cert --subject="$intermediate_cert_subject" --subject-key-id="$intermediate_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$intermediate_cert_2_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$intermediate_cert_1_serial_number\""

echo "Remove an intermediate certificate with subject and subjectKeyId"
result=$(echo "$passphrase" | dcld tx pki remove-noc-x509-ica-cert --subject="$intermediate_cert_subject" --subject-key-id="$intermediate_cert_subject_key_id" --from=$vendor_account_65521 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Request approved certificates by an intermediate certificate's subject and subjectKeyId should be empty"
result=$(dcld query pki noc-x509-cert --subject="$intermediate_cert_subject" --subject-key-id="$intermediate_cert_subject_key_id")
echo $result | jq
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$intermediate_cert_2_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$intermediate_cert_1_serial_number\""

echo "Request ICA certificates by VID should contain only one leaf certificate"
result=$(dcld query pki noc-x509-ica-certs --vid="$root_cert_vid")
echo $result | jq
check_response "$result" "\"subject\": \"$leaf_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""
response_does_not_contain "$result" "\"subject\": \"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$intermediate_cert_2_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$intermediate_cert_1_serial_number\""

echo "Request all revoked certificates should be empty"
result=$(dcld query pki all-revoked-noc-x509-ica-certs)
echo $result | jq
check_response "$result" "\[\]"
response_does_not_contain "$result" "\"subject\": \"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$intermediate_cert_1_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$intermediate_cert_2_serial_number\""

echo "Request all certificates should contain only root and leaf certificates"
result=$(dcld query pki all-noc-x509-certs)
echo $result | jq
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
check_response "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$root_cert_1_serial_number\""
check_response "$result" "\"serialNumber\": \"$leaf_cert_serial_number\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$intermediate_cert_2_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$intermediate_cert_1_serial_number\""

echo "Remove leaf certificate"
result=$(echo "$passphrase" | dcld tx pki remove-noc-x509-ica-cert --subject="$leaf_cert_subject" --subject-key-id="$leaf_cert_subject_key_id" --from=$vendor_account_65521 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Request approved leaf certificates should be empty"
result=$(dcld query pki noc-x509-cert --subject="$leaf_cert_subject" --subject-key-id="$leaf_cert_subject_key_id")
echo $result | jq
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$leaf_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$leaf_cert_serial_number"

echo "Request ICA certificates by VID should be empty"
result=$(dcld query pki noc-x509-ica-certs --vid="$root_cert_vid")
echo $result | jq
response_does_not_contain "$result" "\"subject\": \"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"subject\": \"$leaf_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$intermediate_cert_1_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$intermediate_cert_2_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$leaf_cert_serial_number"

echo "Request all certificates should contain only root certificate"
result=$(dcld query pki all-noc-x509-certs)
echo $result | jq
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$root_cert_1_serial_number\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id"
response_does_not_contain "$result" "\"serialNumber\": \"$intermediate_cert_1_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$intermediate_cert_2_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$leaf_cert_serial_number"

echo "Add second NOC root certificate"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-root-cert --certificate="$root_cert_1_copy_path" --from $vendor_account_65521 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Re-add an ICA certificate"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-ica-cert --certificate="$intermediate_cert_1_path" --from $vendor_account_65521 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Check that root cert is added. Request all approved certificates."
result=$(dcld query pki all-noc-x509-certs)
echo $result | jq
check_response "$result" "\"serialNumber\": \"$root_cert_1_serial_number\""
check_response "$result" "\"serialNumber\": \"$root_cert_1_copy_serial_number\""

echo "Try to remove NOC root certificate with invalid serialNumber"
result=$(echo "$passphrase" | dcld tx pki remove-noc-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --serial-number="invalid" --from=$vendor_account_65521 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 404"

echo "Try to remove NOC root certificate when sender is not Vendor account"
result=$(echo "$passphrase" | dcld tx pki remove-noc-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --serial-number="$root_cert_1_serial_number" --from=$trustee_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 4"

echo "Try to remove NOC root certificate using a vendor account with other VID"
result=$(echo "$passphrase" | dcld tx pki remove-noc-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --serial-number="$root_cert_1_serial_number" --from=$vendor_account_65522 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 4"

echo "Revoke NOC root certificate with serialNumber $root_cert_1_serial_number"
result=$(echo "$passphrase" | dcld tx pki revoke-noc-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --serial-number="$root_cert_1_serial_number" --from=$vendor_account_65521 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Request all revoked certificates should contain NOC root certificate with serialNumber $root_cert_1_serial_number"
result=$(dcld query pki all-revoked-noc-x509-root-certs)
echo $result | jq
check_response "$result" "\"serialNumber\": \"$root_cert_1_serial_number\""

echo "Remove NOC root certificate with serialNumber $root_cert_1_serial_number"
result=$(echo "$passphrase" | dcld tx pki remove-noc-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --serial-number="$root_cert_1_serial_number" --from=$vendor_account_65521 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Request all certificates should contain only one NOC root certificate"
result=$(dcld query pki all-noc-x509-certs)
echo $result | jq
check_response "$result" "\"serialNumber\": \"$root_cert_1_copy_serial_number\""
check_response "$result" "\"serialNumber\": \"$intermediate_cert_1_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$root_cert_1_serial_number\""

echo "Request NOC certificates by VID should contain one NOC root certificate"
result=$(dcld query pki noc-x509-root-certs --vid="$root_cert_vid")
echo $result | jq
check_response "$result" "\"serialNumber\": \"$root_cert_1_copy_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$root_cert_1_serial_number\""

echo "Request approved certificates by NOC root's subject and subjectKeyId should contain only one root certificate"
result=$(dcld query pki noc-x509-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"serialNumber\": \"$root_cert_1_copy_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$root_cert_1_serial_number\""

echo "Re-add NOC root certificate"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-root-cert --certificate="$root_cert_1_path" --from $vendor_account_65521 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Check that root cert is added. Request all approved certificates."
result=$(dcld query pki all-noc-x509-certs)
echo $result | jq
check_response "$result" "\"serialNumber\": \"$root_cert_1_serial_number\""
check_response "$result" "\"serialNumber\": \"$root_cert_1_copy_serial_number\""

echo "Remove NOC root certificates"
result=$(echo "$passphrase" | dcld tx pki remove-noc-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from=$vendor_account_65521 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Request approved NOC root certificates should be empty"
result=$(dcld query pki noc-x509-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
echo $result | jq
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"serialNumber\": \"$root_cert_1_serial_number"
response_does_not_contain "$result" "\"serialNumber\": \"$root_cert_1_copy_serial_number"

echo "Request NOC root certificates by VID should be empty"
result=$(dcld query pki noc-x509-root-certs --vid="$root_cert_vid")
echo $result | jq
response_does_not_contain "$result" "\"serialNumber\": \"$root_cert_1_serial_number"
response_does_not_contain "$result" "\"serialNumber\": \"$root_cert_1_copy_serial_number"

echo "Request all certificates should contain only ICA certificate"
result=$(dcld query pki all-noc-x509-certs)
echo $result | jq
check_response "$result" "\"serialNumber\": \"$intermediate_cert_1_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$root_cert_1_serial_number"
response_does_not_contain "$result" "\"serialNumber\": \"$root_cert_1_copy_serial_number"

echo "Request NOC certificate by VID = $root_cert_vid and SKID = $intermediate_cert_subject_key_id should not be empty"
result=$(dcld query pki noc-x509-cert --vid="$root_cert_vid" --subject-key-id="$intermediate_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"serialNumber\": \"$intermediate_cert_1_serial_number\""

echo "Request NOC certificate by VID = $root_cert_vid and SKID = $root_cert_subject_key_id should be empty"
result=$(dcld query pki noc-x509-cert --vid="$root_cert_vid" --subject-key-id="$root_cert_subject_key_id")
echo $result | jq
check_response "$result" "Not Found"

test_divider
