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
noc_root_cert_1_subject_as_text="CN=NOC-1,OU=Testing Division,O=Example Company,L=Tashkent,ST=Some State,C=UZ"

noc_root_cert_1_copy_path="integration_tests/constants/noc_root_cert_1_copy"
noc_root_cert_1_copy_serial_number="12722088350714347345576486793058060481880825999"

noc_root_cert_2_path="integration_tests/constants/noc_root_cert_2"
noc_root_cert_2_subject="MHoxCzAJBgNVBAYTAlVaMRMwEQYDVQQIEwpTb21lIFN0YXRlMREwDwYDVQQHEwhUYXNoa2VudDEYMBYGA1UEChMPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLExBUZXN0aW5nIERpdmlzaW9uMQ4wDAYDVQQDEwVOT0MtMg=="
noc_root_cert_2_subject_key_id="46:C0:B0:74:0C:63:C8:9E:E0:5C:14:C2:71:62:F8:67:24:5C:8E:29"
noc_root_cert_2_serial_number="727423814323052015089749828769570958840545369270"
noc_root_cert_2_subject_as_text="CN=NOC-2,OU=Testing Division,O=Example Company,L=Tashkent,ST=Some State,C=UZ"

noc_root_cert_3_path="integration_tests/constants/noc_root_cert_3"
noc_root_cert_3_subject="MFUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpTb21lLVN0YXRlMSEwHwYDVQQKExhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQxDjAMBgNVBAMTBU5PQy0z"
noc_root_cert_3_subject_key_id="0F:D2:F8:12:06:F1:38:2D:D2:19:2F:29:52:42:AA:FB:E7:2F:7B:A3"
noc_root_cert_3_serial_number="620481712672111766723531823383547399894194653186"
noc_root_cert_3_subject_as_text="CN=NOC-3,O=Internet Widgits Pty Ltd,ST=Some-State,C=AU"

noc_cert_1_path="integration_tests/constants/noc_cert_1"
noc_cert_1_subject="MIGCMQswCQYDVQQGEwJVWjETMBEGA1UECBMKU29tZSBTdGF0ZTETMBEGA1UEBxMKU29tZSBTdGF0ZTEYMBYGA1UEChMPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLExBUZXN0aW5nIERpdmlzaW9uMRQwEgYDVQQDEwtOT0MtY2hpbGQtMQ=="
noc_cert_1_subject_key_id="06:9F:5A:E0:1F:23:3E:9F:C7:4F:B6:F9:A2:33:47:33:62:7A:07:C5"
noc_cert_1_serial_number="577430346509479530103103319788179390906984119670"

noc_cert_1_copy_path="integration_tests/constants/noc_cert_1_copy"
noc_cert_1_copy_serial_number="617357865778805507017637943649984133152592305888"

noc_cert_2_path="integration_tests/constants/noc_cert_2"
noc_cert_2_subject="MIGCMQswCQYDVQQGEwJVWjETMBEGA1UECBMKU29tZSBTdGF0ZTETMBEGA1UEBxMKU29tZSBTdGF0ZTEYMBYGA1UEChMPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLExBUZXN0aW5nIERpdmlzaW9uMRQwEgYDVQQDEwtOT0MtY2hpbGQtMg=="
noc_cert_2_subject_key_id="17:E2:72:19:E1:7F:19:D7:0D:02:1A:B0:40:7B:04:26:CC:D4:2B:F5"
noc_cert_2_serial_number="634591262660314610068979921875981241084684028375"

noc_leaf_cert_1_path="integration_tests/constants/noc_leaf_cert_1"
noc_leaf_cert_1_subject="MIGBMQswCQYDVQQGEwJVWjETMBEGA1UECBMKU29tZSBTdGF0ZTETMBEGA1UEBxMKU29tZSBTdGF0ZTEYMBYGA1UEChMPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLExBUZXN0aW5nIERpdmlzaW9uMRMwEQYDVQQDEwpOT0MtbGVhZi0x"
noc_leaf_cert_1_subject_key_id="F0:3A:A5:96:8F:60:63:F0:15:21:24:0C:DA:0A:E6:2B:CC:A0:58:F9"
noc_leaf_cert_1_serial_number="201294310322324358101935163941973786245732555938"

trustee_account="jack"
second_trustee_account="alice"

vid_in_hex_format=0x6006
vid=24582

vendor_account=vendor_account_$vid_in_hex_format
echo "Create Vendor account - $vendor_account"
create_new_vendor_account $vendor_account $vid_in_hex_format

vid_2_in_hex_format=0x125D
vid_2=4701

vendor_account_2=vendor_account_$vid_2_in_hex_format
echo "Create Vendor account - $vendor_account_2"
create_new_vendor_account $vendor_account_2 $vid_2_in_hex_format


test_divider

echo "Request NOC root certificate by VID must be empty"
result=$(dcld query pki noc-x509-root-certs --vid="$vid")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$noc_root_cert_1_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$noc_root_cert_1_subject_as_text\""
echo $result | jq

test_divider

echo "Request NOC root certificate by VID = $vid and SKID = $noc_root_cert_1_subject_key_id must be empty"
result=$(dcld query pki noc-x509-cert --vid="$vid" --subject-key-id="$noc_root_cert_1_subject_key_id")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$noc_root_cert_1_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$noc_root_cert_1_subject_as_text\""
echo $result | jq

test_divider

echo "Request NOC root certificate by VID = $vid and SKID = $noc_root_cert_2_subject_key_id must be empty"
result=$(dcld query pki noc-x509-cert --vid="$vid" --subject-key-id="$noc_root_cert_2_subject_key_id")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$noc_root_cert_2_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_2_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_2_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$noc_root_cert_2_subject_as_text\""
echo $result | jq

test_divider

echo "Request NOC root certificate by VID = $vid and SKID = $noc_root_cert_3_subject_key_id must be empty"
result=$(dcld query pki noc-x509-cert --vid="$vid" --subject-key-id="$noc_root_cert_3_subject_key_id")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$noc_root_cert_3_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_3_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_3_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$noc_root_cert_3_subject_as_text\""
echo $result | jq

test_divider

echo "Request all NOC root certificates must be empty"
result=$(dcld query pki all-noc-x509-root-certs)
check_response "$result" "\[\]"
response_does_not_contain "$result" "\"subject\": \"$noc_root_cert_1_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$noc_root_cert_1_subject_as_text\""
echo $result | jq

test_divider

echo "Request approved certificate must be empty"
result=$(dcld query pki noc-x509-cert --subject="$noc_root_cert_1_subject" --subject-key-id="$noc_root_cert_1_subject_key_id")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$noc_root_cert_1_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$noc_root_cert_1_subject_as_text\""
echo $result | jq

test_divider

echo "Request all certificates by subject must be empty"
result=$(dcld query pki all-noc-subject-x509-certs --subject="$noc_root_cert_1_subject")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$noc_root_cert_1_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
echo $result | jq

test_divider

echo "Request all certificates by subjectKeyId must be empty"
result=$(dcld query pki noc-x509-cert --subject-key-id="$noc_root_cert_1_subject_key_id")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$noc_root_cert_1_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$noc_root_cert_1_subject_as_text\""
echo $result | jq

test_divider

echo "Try to add intermediate cert using add-noc-x509-root-cert command"
intermediate_path="integration_tests/constants/intermediate_cert"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-root-cert --certificate="$intermediate_path" --from $vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 414"

echo "Add first NOC root certificate by vendor with VID = $vid"
cert_schema_version_0=0
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-root-cert --certificate="$noc_root_cert_1_path" --schemaVersion=$cert_schema_version_0 --from $vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

schema_version_0=0
echo "Add second NOC root certificate by vendor with VID = $vid"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-root-cert --certificate="$noc_root_cert_2_path" --from $vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Add third NOC root certificate by vendor with VID = $vid_2"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-root-cert --certificate="$noc_root_cert_3_path" --from $vendor_account_2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Request NOC root certificate by VID = $vid"
result=$(dcld query pki noc-x509-root-certs --vid="$vid")
echo $result | jq
check_response "$result" "\"subject\": \"$noc_root_cert_1_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
check_response "$result" "\"subjectAsText\": \"$noc_root_cert_1_subject_as_text\""
check_response "$result" "\"subject\": \"$noc_root_cert_2_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_2_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_2_serial_number\""
check_response "$result" "\"subjectAsText\": \"$noc_root_cert_2_subject_as_text\""
check_response "$result" "\"schemaVersion\": $cert_schema_version_0"
check_response "$result" "\"schemaVersion\": $cert_schema_version_0"
check_response "$result" "\"schemaVersion\": $schema_version_0"
check_response "$result" "\"vid\": $vid"

test_divider

echo "Request NOC root certificate by VID = "$vid" and SKID = $noc_root_cert_1_subject_key_id"
result=$(dcld query pki noc-x509-cert --vid="$vid" --subject-key-id="$noc_root_cert_1_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$noc_root_cert_1_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
check_response "$result" "\"subjectAsText\": \"$noc_root_cert_1_subject_as_text\""
check_response "$result" "\"schemaVersion\": $cert_schema_version_0"
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"tq\": 1"

test_divider

echo "Request NOC root certificate by VID = "$vid" and SKID = $noc_root_cert_2_subject_key_id"
result=$(dcld query pki noc-x509-cert --vid="$vid" --subject-key-id="$noc_root_cert_2_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$noc_root_cert_2_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_2_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_2_serial_number\""
check_response "$result" "\"subjectAsText\": \"$noc_root_cert_2_subject_as_text\""
check_response "$result" "\"schemaVersion\": $cert_schema_version_0"
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"tq\": 1"

test_divider

echo "Request All NOC root certificate"
result=$(dcld query pki all-noc-x509-root-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$noc_root_cert_1_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
check_response "$result" "\"subjectAsText\": \"$noc_root_cert_1_subject_as_text\""
check_response "$result" "\"subject\": \"$noc_root_cert_2_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_2_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_2_serial_number\""
check_response "$result" "\"subjectAsText\": \"$noc_root_cert_2_subject_as_text\""
check_response "$result" "\"subject\": \"$noc_root_cert_3_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_3_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_3_serial_number\""
check_response "$result" "\"subjectAsText\": \"$noc_root_cert_3_subject_as_text\""
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"vid\": $vid_2"

test_divider

echo "Request NOC root certificate by Subject and SubjectKeyID (using noc-x509-cert command)"
result=$(dcld query pki noc-x509-cert --subject="$noc_root_cert_1_subject" --subject-key-id="$noc_root_cert_1_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$noc_root_cert_1_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
check_response "$result" "\"subjectAsText\": \"$noc_root_cert_1_subject_as_text\""
check_response "$result" "\"approvals\": \\[\\]"

echo "Request NOC root certificate by Subject and SubjectKeyID (using cert command)"
result=$(dcld query pki cert --subject="$noc_root_cert_1_subject" --subject-key-id="$noc_root_cert_1_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$noc_root_cert_1_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
check_response "$result" "\"subjectAsText\": \"$noc_root_cert_1_subject_as_text\""
check_response "$result" "\"approvals\": \\[\\]"

test_divider

echo "Request NOC root certificate by Subject "
result=$(dcld query pki all-noc-subject-x509-certs --subject="$noc_root_cert_1_subject")
echo $result | jq
check_response "$result" "\"$noc_root_cert_1_subject\""
check_response "$result" "\"$noc_root_cert_1_subject_key_id\""

test_divider

echo "Request NOC root certificate by SubjectKeyID"
result=$(dcld query pki noc-x509-cert --subject-key-id="$noc_root_cert_1_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$noc_root_cert_1_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
check_response "$result" "\"subjectAsText\": \"$noc_root_cert_1_subject_as_text\""

echo "Add first intermidiate NOC certificate by vendor with VID = $vid"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-ica-cert --certificate="$noc_cert_1_path" --from $vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Request intermidiate NOC certificate by VID = $vid"
result=$(dcld query pki noc-x509-ica-certs --vid="$vid")
echo $result | jq
check_response "$result" "\"subject\": \"$noc_cert_1_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_cert_1_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_cert_1_serial_number\""
check_response "$result" "\"vid\": $vid"

test_divider

echo "Request all child certificates by Subject and SubjectKeyID"
result=$(dcld query pki all-child-x509-certs --subject="$noc_root_cert_1_subject" --subject-key-id="$noc_root_cert_1_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$noc_cert_1_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_cert_1_subject_key_id\""

echo "Try to add intermediate with different VID = $vid_2"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-ica-cert --certificate="$noc_cert_2_path" --from $vendor_account_2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 439"

test_divider

echo "Add second NOC certificate by vendor with VID = $vid"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-ica-cert --certificate="$noc_cert_2_path" --schemaVersion=$cert_schema_version_0 --from $vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Add third NOC certificate by vendor with VID = $vid"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-ica-cert --certificate="$noc_cert_1_copy_path" --from $vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Request all NOC certificates"
result=$(dcld query pki all-noc-x509-ica-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$noc_cert_1_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_cert_1_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_cert_1_serial_number\""
check_response "$result" "\"serialNumber\": \"$noc_cert_1_copy_serial_number\""
check_response "$result" "\"subject\": \"$noc_cert_2_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_cert_2_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_cert_2_serial_number\""
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"schemaVersion\": $cert_schema_version_0"
check_response "$result" "\"schemaVersion\": $cert_schema_version_0"
check_response "$result" "\"schemaVersion\": $schema_version_0"

echo "Request all approved certificates (Must be empty)"
result=$(dcld query pki all-x509-certs)
check_response "$result" "\[\]"
response_does_not_contain "$result" "\"subject\": \"$noc_root_cert_1_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
response_does_not_contain "$result" "\"subject\": \"$noc_cert_1_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_cert_1_serial_number\""

echo "Request all Noc certificates"
result=$(dcld query pki all-noc-x509-certs)
check_response "$result" "\"subject\": \"$noc_root_cert_1_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
check_response "$result" "\"subject\": \"$noc_cert_1_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_cert_1_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_cert_1_serial_number\""
check_response "$result" "\"serialNumber\": \"$noc_cert_1_copy_serial_number\""
check_response "$result" "\"subject\": \"$noc_cert_2_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_cert_2_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_cert_2_serial_number\""
echo $result | jq

test_divider

echo "Add third NOC root certificate by vendor with VID = $vid"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-root-cert --certificate="$noc_root_cert_1_copy_path" --from $vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Add NOC leaf certificate by vendor with VID = $vid"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-ica-cert --certificate="$noc_leaf_cert_1_path" --from $vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Request All NOC root certificate"
result=$(dcld query pki all-noc-x509-root-certs)
echo $result | jq
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_copy_serial_number\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_2_serial_number\""

echo "Request all NOC certificates"
result=$(dcld query pki all-noc-x509-ica-certs)
echo $result | jq
check_response "$result" "\"serialNumber\": \"$noc_cert_1_serial_number\""
check_response "$result" "\"serialNumber\": \"$noc_cert_1_copy_serial_number\""
check_response "$result" "\"serialNumber\": \"$noc_cert_2_serial_number\""
check_response "$result" "\"serialNumber\": \"$noc_leaf_cert_1_serial_number\""

echo "Try to revoke NOC root certificate with different VID = $vid_2"
result=$(echo "$passphrase" | dcld tx pki revoke-noc-x509-root-cert --subject="$noc_root_cert_1_subject" --subject-key-id="$noc_root_cert_1_subject_key_id" --from $vendor_account_2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 439"

echo "$vendor_account Vendor revokes only root certificate, it should not revoke intermediate certificates"
result=$(echo "$passphrase" | dcld tx pki revoke-noc-x509-root-cert --subject="$noc_root_cert_1_subject" --subject-key-id="$noc_root_cert_1_subject_key_id" --from=$vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Request all revoked noc root certificates should contain two root certificates"
result=$(dcld query pki all-revoked-noc-x509-root-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$noc_root_cert_1_subject"
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_copy_serial_number\""
response_does_not_contain "$result" "\"subject\": \"$noc_cert_1_subject\""
response_does_not_contain "$result" "\"subject\": \"$noc_leaf_cert_1_subject\""

echo "Request revoked noc root certificate by subject and subjectKeyId should contain two root certificates"
result=$(dcld query pki revoked-noc-x509-root-cert --subject="$noc_root_cert_1_subject" --subject-key-id="$noc_root_cert_1_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$noc_root_cert_1_subject"
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_copy_serial_number\""
response_does_not_contain "$result" "\"subject\": \"$noc_root_cert_2_subject\""
response_does_not_contain "$result" "\"subject\": \"$noc_root_cert_3_subject\""

echo "Request all x509 root revoked certificates should not contain revoked NOC root certificates"
result=$(dcld query pki all-revoked-x509-root-certs)
response_does_not_contain "$result" "\"subject\": \"$noc_root_cert_1_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_copy_serial_number\""
echo $result | jq

echo "Request NOC root certificate by VID = $vid must not contain revoked root certificates"
result=$(dcld query pki noc-x509-root-certs --vid="$vid")
check_response "$result" "\"subject\": \"$noc_root_cert_2_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_2_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_2_serial_number\""
response_does_not_contain "$result" "\"subject\": \"$noc_root_cert_1_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_copy_serial_number\""
echo $result | jq

echo "Request NOC root certificate by VID = $vid and SKID = $noc_root_cert_1_subject must be empty"
result=$(dcld query pki noc-x509-cert --vid="$vid" --subject-key-id="$noc_root_cert_1_subject_key_id")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$noc_root_cert_1_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$noc_root_cert_1_subject_as_text\""
echo $result | jq

echo "Request NOC root certificate by VID = "$vid" and SKID = $noc_root_cert_2_subject_key_id"
result=$(dcld query pki noc-x509-cert --vid="$vid" --subject-key-id="$noc_root_cert_2_subject_key_id")
check_response "$result" "\"subject\": \"$noc_root_cert_2_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_2_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_2_serial_number\""
check_response "$result" "\"subjectAsText\": \"$noc_root_cert_2_subject_as_text\""
check_response "$result" "\"schemaVersion\": $cert_schema_version_0"
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"tq\": 1"
echo $result | jq

echo "Request all certificates by subject must be empty"
result=$(dcld query pki all-noc-subject-x509-certs --subject="$noc_root_cert_1_subject")
response_does_not_contain "$result" "\"subject\": \"$noc_root_cert_1_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
echo $result | jq

echo "Request all certificates by subjectKeyId must be empty"
result=$(dcld query pki noc-x509-cert --subject-key-id="$noc_root_cert_1_subject_key_id")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$noc_root_cert_1_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_copy_serial_number\""
echo $result | jq

echo "Request NOC certificate by VID = $vid should contain intermediate and leaf certificates"
result=$(dcld query pki noc-x509-ica-certs --vid="$vid")
echo $result | jq
check_response "$result" "\"subject\": \"$noc_cert_1_subject\""
check_response "$result" "\"subject\": \"$noc_leaf_cert_1_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_cert_1_subject_key_id\""
check_response "$result" "\"subjectKeyId\": \"$noc_leaf_cert_1_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_cert_1_serial_number\""
check_response "$result" "\"serialNumber\": \"$noc_cert_1_copy_serial_number\""
check_response "$result" "\"serialNumber\": \"$noc_leaf_cert_1_serial_number\""

echo "Request all noc certificates should not contain revoked NOC root certificates"
result=$(dcld query pki all-noc-x509-certs)
check_response "$result" "\"subject\": \"$noc_cert_1_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_cert_1_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_cert_1_serial_number\""
check_response "$result" "\"serialNumber\": \"$noc_cert_1_copy_serial_number\""
check_response "$result" "\"subject\": \"$noc_leaf_cert_1_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_leaf_cert_1_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_leaf_cert_1_serial_number\""
response_does_not_contain "$result" "\"subject\": \"$noc_root_cert_1_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_copy_serial_number\""
echo $result | jq

test_divider

echo "REVOCATION OF NON-ROOT NOC CERTIFICATES"

echo "Try to revoke NOC certificate with different VID = $vid_2"
result=$(echo "$passphrase" | dcld tx pki revoke-noc-x509-ica-cert --subject="$noc_cert_1_subject" --subject-key-id="$noc_cert_1_subject_key_id" --from $vendor_account_2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 439"

echo "$vendor_account Vendor revokes only NOC certificates, it should not revoke leaf certificates"
result=$(echo "$passphrase" | dcld tx pki revoke-noc-x509-ica-cert --subject="$noc_cert_1_subject" --subject-key-id="$noc_cert_1_subject_key_id" --from=$vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Request all revoked certificates should not contain leaf certificate"
result=$(dcld query pki all-revoked-noc-x509-ica-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$noc_cert_1_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_cert_1_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_cert_1_serial_number"
check_response "$result" "\"schemaVersion\": $schema_version_0"
response_does_not_contain "$result" "\"subject\": \"$noc_leaf_cert_1_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_leaf_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_leaf_cert_1_serial_number"

echo "Request all revoked noc root certificates should not contain non-root NOC certificates"
result=$(dcld query pki all-revoked-noc-x509-root-certs)
echo $result | jq
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_cert_1_subject_key_id"
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_leaf_cert_1_subject_key_id\""

echo "Request all certificates by subject must be empty"
result=$(dcld query pki all-noc-subject-x509-certs --subject="$noc_cert_1_subject")
response_does_not_contain "$result" "\"subject\": \"$noc_cert_1_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_cert_1_subject_key_id\""
echo $result | jq

echo "Request all certificates by subjectKeyId must be empty"
result=$(dcld query pki noc-x509-cert --subject-key-id="$noc_cert_1_subject_key_id")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$noc_cert_1_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_cert_1_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_cert_1_copy_serial_number\""
echo $result | jq

echo "Request NOC certificate by VID = $vid should contain one leaf certificate"
result=$(dcld query pki noc-x509-ica-certs --vid="$vid")
echo $result | jq
check_response "$result" "\"subject\": \"$noc_leaf_cert_1_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_leaf_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"subject\": \"$noc_cert_1_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_cert_1_subject_key_id\""

echo "Request NOC certificate by VID = $vid and SKID = $noc_cert_1_subject_key_id should be empty"
result=$(dcld query pki noc-x509-cert --vid="$vid" --subject-key-id="$noc_cert_1_subject_key_id")
echo $result | jq
check_response "$result" "Not Found"

echo "Request NOC certificate by VID = $vid and SKID = $noc_leaf_cert_1_subject_key_id should contain one leaf certificate"
result=$(dcld query pki noc-x509-cert --vid="$vid" --subject-key-id="$noc_leaf_cert_1_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$noc_leaf_cert_1_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_leaf_cert_1_subject_key_id\""

echo "Request all approved certificates should not contain revoked NOC certificates"
result=$(dcld query pki all-noc-x509-certs)
check_response "$result" "\"subject\": \"$noc_leaf_cert_1_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_leaf_cert_1_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_leaf_cert_1_serial_number\""
response_does_not_contain "$result" "\"subject\": \"$noc_cert_1_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_cert_1_serial_number\""
response_does_not_contain "$result" "\"subject\": \"$noc_root_cert_1_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_copy_serial_number\""
echo $result | jq

test_divider

echo "PASS"