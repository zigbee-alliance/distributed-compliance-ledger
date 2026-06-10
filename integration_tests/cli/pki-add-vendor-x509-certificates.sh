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

root_cert_with_vid_65521_subject="MIGYMQswCQYDVQQGEwJVUzERMA8GA1UECBMITmV3IFlvcmsxETAPBgNVBAcTCE5ldyBZb3JrMRgwFgYDVQQKEw9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsTEFRlc3RpbmcgRGl2aXNpb24xGDAWBgNVBAMTD3d3dy5leGFtcGxlLmNvbTEUMBIGCisGAQQBgqJ8AgETBEZGRjE="
root_cert_with_vid_65521_subject_key_id="01:A9:FC:05:0D:83:72:97:53:65:E8:24:EB:33:CB:90:E6:D7:23:DB"
root_cert_with_vid_65521_path="integration_tests/constants/root_cert_with_vid"
root_cert_with_vid_65521_vid=65521

intermediate_cert_with_vid_65521_subject="MIGuMQswCQYDVQQGEwJVUzERMA8GA1UECBMITmV3IFlvcmsxETAPBgNVBAcTCE5ldyBZb3JrMRgwFgYDVQQKEw9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsTEFRlc3RpbmcgRGl2aXNpb24xGDAWBgNVBAMTD3d3dy5leGFtcGxlLmNvbTEUMBIGCisGAQQBgqJ8AgETBEZGRjExFDASBgorBgEEAYKifAICEwRGRkYx"
intermediate_cert_with_vid_65521_subject_key_id="01:EA:C7:E5:68:AF:D6:F3:35:DB:FC:D5:9D:7C:59:4A:81:63:33:65"
intermediate_cert_with_vid_65521_path="integration_tests/constants/intermediate_cert_with_vid_1"
intermediate_cert_with_vid_65521_serial_number="3"
intermediate_cert_with_vid_65521_vid=65521

intermediate_cert_with_vid_65522_path="integration_tests/constants/intermediate_cert_with_vid_2"
intermediate_cert_with_vid_65522_serial_number="4"
intermediate_cert_with_vid_65522_vid=65522

trustee_account="jack"
second_trustee_account="alice"

test_divider

echo "ADD VID SCOPED X509 CERTIFICATES"

vendor_vid_65521=$root_cert_with_vid_65521_vid
vendor_account_65521=vendor_account_$vendor_vid_65521
echo "Create Vendor account - $vendor_account_65521"
create_new_vendor_account $vendor_account_65521 $vendor_vid_65521

echo "Propose and approve root certificate with vid=$root_cert_with_vid_65521_vid"
result=$(echo "$passphrase" | dcld tx pki propose-add-x509-root-cert --certificate="$root_cert_with_vid_65521_path" --vid "$root_cert_with_vid_65521_vid" --from $trustee_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"
result=$(echo "$passphrase" | dcld tx pki approve-add-x509-root-cert --subject="$root_cert_with_vid_65521_subject" --subject-key-id="$root_cert_with_vid_65521_subject_key_id" --from $second_trustee_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Try to add the intermediate certificate using an account that does not have vendor role"
result=$(echo "$passphrase" | dcld tx pki add-x509-cert --certificate="$intermediate_cert_with_vid_65521_path" --from $trustee_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 4"

echo "Add an intermediate certificate with vid=$intermediate_cert_with_vid_65521_vid by $vendor_account_65521 with vid=$vendor_vid_65521"
result=$(echo "$passphrase" | dcld tx pki add-x509-cert --certificate="$intermediate_cert_with_vid_65521_path" --from $vendor_account_65521 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Request all approved root certificates."
result=$(dcld query pki all-x509-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$root_cert_with_vid_65521_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_with_vid_65521_subject_key_id\""
check_response "$result" "\"subject\": \"$intermediate_cert_with_vid_65521_subject\""
check_response "$result" "\"subjectKeyId\": \"$intermediate_cert_with_vid_65521_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$intermediate_cert_with_vid_65521_serial_number\""

echo "Request intermediate certificate by subject=$intermediate_cert_with_vid_65521_subject and skid=$intermediate_cert_with_vid_65521_subject_key_id."
result=$(dcld query pki x509-cert --subject="$intermediate_cert_with_vid_65521_subject" --subject-key-id="$intermediate_cert_with_vid_65521_subject_key_id")
echo $result | jq
check_response "$result" "\"serialNumber\": \"$intermediate_cert_with_vid_65521_serial_number\""
check_response "$result" "\"vid\": $intermediate_cert_with_vid_65521_vid"

echo "Try to add an intermediate certificate with vid=$intermediate_cert_with_vid_65522_vid by $vendor_account_65521 with vid=$vendor_vid_65521"
result=$(echo "$passphrase" | dcld tx pki add-x509-cert --certificate="$intermediate_cert_with_vid_65522_path" --from $vendor_account_65521 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 440"

echo "Request all approved root certificates should not contain intermediate cert with serialNumber=$intermediate_cert_with_vid_65522_serial_number"
result=$(dcld query pki all-x509-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$root_cert_with_vid_65521_subject\""
check_response "$result" "\"subject\": \"$intermediate_cert_with_vid_65521_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_with_vid_65521_subject_key_id\""
check_response "$result" "\"subjectKeyId\": \"$intermediate_cert_with_vid_65521_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$intermediate_cert_with_vid_65521_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$intermediate_cert_with_vid_65522_serial_number\""

root_cert_with_no_vid_path="integration_tests/constants/paa_cert_no_vid"
root_cert_with_no_vid_subject="MBoxGDAWBgNVBAMMD01hdHRlciBUZXN0IFBBQQ=="
root_cert_with_no_vid_subject_key_id="78:5C:E7:05:B8:6B:8F:4E:6F:C7:93:AA:60:CB:43:EA:69:68:82:D5"

intermediate_cert_with_vid_65522_path="integration_tests/constants/pai_cert_numeric_vid"
intermediate_cert_with_vid_65522_subject="MDAxGDAWBgNVBAMMD01hdHRlciBUZXN0IFBBSTEUMBIGCisGAQQBgqJ8AgEMBEZGRjI="
intermediate_cert_with_vid_65522_subject_key_id="61:3D:D0:87:35:5E:F0:8B:AE:01:E4:C6:9A:8F:C7:3D:AC:8C:7D:FD"
intermediate_cert_with_vid_65522_vid=65522
intermediate_cert_with_vid_65522_serial_number="4428370313154203676"

echo "Propose and approve non-vid root certificate"
result=$(echo "$passphrase" | dcld tx pki propose-add-x509-root-cert --certificate="$root_cert_with_no_vid_path" --vid "65522" --from $trustee_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"
result=$(echo "$passphrase" | dcld tx pki approve-add-x509-root-cert --subject="$root_cert_with_no_vid_subject" --subject-key-id="$root_cert_with_no_vid_subject_key_id" --from $second_trustee_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

vendor_vid_65523=65523
vendor_account_65523=vendor_account_$vendor_vid_65523
echo "Create Vendor account - $vendor_account_65523"
create_new_vendor_account $vendor_account_65523 $vendor_vid_65523

echo "Try to add an intermediate certificate with vid=$intermediate_cert_with_vid_65522_vid by $vendor_account_65523 with vid=$vendor_vid_65523"
result=$(echo "$passphrase" | dcld tx pki add-x509-cert --certificate="$intermediate_cert_with_vid_65522_path" --from $vendor_account_65523 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 439"

echo "Request all approved root certificates should not contain intermediate cert with serialNumber=$intermediate_cert_with_vid_65522_serial_number"
result=$(dcld query pki all-x509-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$root_cert_with_no_vid_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_with_no_vid_subject_key_id\""
response_does_not_contain "$result" "\"subject\": \"$intermediate_cert_with_vid_65522_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$intermediate_cert_with_vid_65522_subject_key_id"
response_does_not_contain "$result" "\"serialNumber\": \"$intermediate_cert_with_vid_65522_serial_number\""

vendor_vid_65522=65522
vendor_account_65522=vendor_account_$vendor_vid_65522
echo "Create Vendor account - $vendor_account_65522"
create_new_vendor_account $vendor_account_65522 $vendor_vid_65522

echo "Add an intermediate certificate with vid=$intermediate_cert_with_vid_65522_vid by $vendor_account_65522 with vid=$vendor_vid_65522"
result=$(echo "$passphrase" | dcld tx pki add-x509-cert --certificate="$intermediate_cert_with_vid_65522_path" --from $vendor_account_65522 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Request all approved root certificates should contain intermediate cert with serialNumber=$intermediate_cert_with_vid_65522_serial_number"
result=$(dcld query pki all-x509-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$root_cert_with_no_vid_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_with_no_vid_subject_key_id\""
check_response "$result" "\"subject\": \"$intermediate_cert_with_vid_65522_subject\""
check_response "$result" "\"subjectKeyId\": \"$intermediate_cert_with_vid_65522_subject_key_id"
check_response "$result" "\"serialNumber\": \"$intermediate_cert_with_vid_65522_serial_number\""

test_divider

root_cert_subject="MIGCMQswCQYDVQQGEwJVUzERMA8GA1UECBMITmV3IFlvcmsxETAPBgNVBAcTCE5ldyBZb3JrMRgwFgYDVQQKEw9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsTEFRlc3RpbmcgRGl2aXNpb24xGDAWBgNVBAMTD3d3dy5leGFtcGxlLmNvbQ=="
root_cert_subject_key_id="76:16:85:A1:B2:F9:C0:5A:03:A4:5B:70:CE:18:BC:78:9E:3A:9B:68"
root_cert_path="integration_tests/constants/root_with_same_subject_and_skid_1"
root_cert_serial_number="1"

intermediate_cert_subject="MEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpTb21lLVN0YXRlMSEwHwYDVQQKExhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQ="
intermediate_cert_subject_key_id="4C:C8:CA:45:F6:0F:FA:C8:9B:9B:80:89:23:A9:CA:6A:79:B1:E7:EC"
intermediate_cert_1_path="integration_tests/constants/intermediate_with_same_subject_and_skid_1"
intermediate_cert_1_serial_number="3"
intermediate_cert_2_path="integration_tests/constants/intermediate_with_same_subject_and_skid_2"
intermediate_cert_2_serial_number="4"

echo "Propose and approve root certificate"
result=$(echo "$passphrase" | dcld tx pki propose-add-x509-root-cert --certificate="$root_cert_path" --vid "$vendor_vid_65521" --from $trustee_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"
result=$(echo "$passphrase" | dcld tx pki approve-add-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from $second_trustee_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Add first intermediate certificate by $vendor_account_65521"
result=$(echo "$passphrase" | dcld tx pki add-x509-cert --certificate="$intermediate_cert_1_path" --from $vendor_account_65521 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Request all approved root certificates."
result=$(dcld query pki all-x509-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$intermediate_cert_1_serial_number\""

echo "Try to add second intermediate certificate with same subject and SKID by $vendor_account_65523"
result=$(echo "$passphrase" | dcld tx pki add-x509-cert --certificate="$intermediate_cert_2_path" --from $vendor_account_65523 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 4"

echo "Request all approved root certificates should not contain intermediate cert with serialNumber=$intermediate_cert_2_serial_number"
result=$(dcld query pki all-x509-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$intermediate_cert_1_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$intermediate_cert_2_serial_number\""

second_vendor_account_65521=second_vendor_account_$vendor_vid_65521
echo "Create Vendor account - $second_vendor_account_65521"
create_new_vendor_account $second_vendor_account_65521 $vendor_vid_65521

echo "Add second intermediate certificate with same subject and SKID by $second_vendor_account_65521"
result=$(echo "$passphrase" | dcld tx pki add-x509-cert --certificate="$intermediate_cert_2_path" --from $vendor_account_65521 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Request all approved root certificates should contain intermediate cert with serialNumber=$intermediate_cert_2_serial_number"
result=$(dcld query pki all-x509-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$intermediate_cert_1_serial_number\""
check_response "$result" "\"serialNumber\": \"$intermediate_cert_2_serial_number\""
