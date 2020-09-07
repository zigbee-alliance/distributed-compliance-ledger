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

set -e
source integration_tests/cli/common.sh

root_cert_subject="CN=DST Root CA X3,O=Digital Signature Trust Co."
root_cert_subject_key_id="C4:A7:B1:A4:7B:2C:71:FA:DB:E1:4B:90:75:FF:C4:15:60:85:89:10"
root_cert_serial_number="91299735575339953335919266965803778155"

intermediate_cert_subject="CN=Let's Encrypt Authority X3,O=Let's Encrypt,C=US"
intermediate_cert_subject_key_id="A8:4A:6A:63:4:7D:DD:BA:E6:D1:39:B7:A6:45:65:EF:F3:A8:EC:A1"
intermediate_cert_serial_number="13298795840390663119752826058995181320"

leaf_cert_subject="CN=dsr-corporation.com"
leaf_cert_subject_key_id="8A:E9:AC:D4:16:81:2F:87:66:8E:61:BE:A9:C5:1C:0:1B:F7:BB:AE"
leaf_cert_serial_number="393904870890265262371394210372104514174397"

# Preparation of Actors

trustee_account="jack"
second_trustee_account="alice"

echo "Create regular account"
create_new_account user_account ""

# Body

echo "$user_account (Not Trustee) propose Root certificate"
root_path="integration_tests/constants/root_cert"
result=$(echo "test1234" | zblcli tx pki propose-add-x509-root-cert --certificate="$root_path" --from $user_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Request proposed Root certificate"
result=$(zblcli query pki proposed-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$root_cert_subject_key_id\""
check_response "$result" "\"serial_number\": \"$root_cert_serial_number\""
echo "$result"

echo "Request all proposed Root certificates"
result=$(zblcli query pki all-proposed-x509-root-certs)
check_response "$result" "\"total\": \"1\""
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$root_cert_subject_key_id\""
check_response "$result" "\"serial_number\": \"$root_cert_serial_number\""
echo "$result"

echo "Request all active certificates must be empty"
result=$(zblcli query pki all-x509-certs)
check_response "$result" "\"total\": \"0\""
echo "$result"

echo "Request all active root certificates must be empty"
result=$(zblcli query pki all-x509-root-certs)
check_response "$result" "\"total\": \"0\""
echo "$result"

echo "$trustee_account (Trustee) approve Root certificate"
result=$(echo "test1234" | zblcli tx pki approve-add-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from $trustee_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Certificate must be still in Proposed state. Request proposed Root certificate"
result=$(zblcli query pki proposed-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$root_cert_subject_key_id\""
check_response "$result" "\"serial_number\": \"$root_cert_serial_number\""
check_response "$result" "[\"$(zblcli keys show jack -a)\"]"
echo "$result"

echo "Request all active certificates must be empty"
result=$(zblcli query pki all-x509-certs)
check_response "$result" "\"total\": \"0\""
echo "$result"

echo "$second_trustee_account (Trustee) approve Root certificate"
result=$(echo "test1234" | zblcli tx pki approve-add-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from $second_trustee_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Certificate must be Approved. Request Root certificate"
result=$(zblcli query pki x509-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$root_cert_subject_key_id\""
check_response "$result" "\"serial_number\": \"$root_cert_serial_number\""
echo "$result"

echo "Request certificate chain for Root certificate"
result=$(zblcli query pki x509-cert-chain --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$root_cert_subject_key_id\""
check_response "$result" "\"serial_number\": \"$root_cert_serial_number\""
echo "$result"

echo "Request all proposed Root certificates must be empty"
result=$(zblcli query pki all-proposed-x509-root-certs)
check_response "$result" "\"total\": \"0\""
echo "$result"

echo "Request all active certificates"
result=$(zblcli query pki all-x509-certs)
check_response "$result" "\"total\": \"1\""
check_response "$result" "\"subject_key_id\": \"$root_cert_subject_key_id\""
echo "$result"

echo "$user_account (Not Trustee) add Intermediate certificate"
intermediate_path="integration_tests/constants/intermediate_cert"
result=$(echo "test1234" | zblcli tx pki add-x509-cert --certificate="$intermediate_path" --from $user_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Request Intermediate certificate"
result=$(zblcli query pki x509-cert --subject="$intermediate_cert_subject" --subject-key-id="$intermediate_cert_subject_key_id")
check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$intermediate_cert_subject_key_id\""
check_response "$result" "\"serial_number\": \"$intermediate_cert_serial_number\""
echo "$result"

echo "Request certificate chain for Intermediate certificate"
result=$(zblcli query pki x509-cert-chain --subject="$intermediate_cert_subject" --subject-key-id="$intermediate_cert_subject_key_id")
check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$intermediate_cert_subject_key_id\""
check_response "$result" "\"serial_number\": \"$intermediate_cert_serial_number\""
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$root_cert_subject_key_id\""
check_response "$result" "\"serial_number\": \"$root_cert_serial_number\""
echo "$result"

echo "Request all proposed Root certificates must be empty"
result=$(zblcli query pki all-proposed-x509-root-certs)
check_response "$result" "\"total\": \"0\""

echo "Request all active certificates"
result=$(zblcli query pki all-x509-certs)
check_response "$result" "\"total\": \"2\""
check_response "$result" "\"subject_key_id\": \"$root_cert_subject_key_id\""
check_response "$result" "\"subject_key_id\": \"$intermediate_cert_subject_key_id\""
echo "$result"

echo "Request all active root certificates"
result=$(zblcli query pki all-x509-root-certs)
check_response "$result" "\"total\": \"1\""
check_response "$result" "\"subject_key_id\": \"$root_cert_subject_key_id\""
echo "$result"

echo "$trustee_account (Trustee) add Leaf certificate"
leaf_path="integration_tests/constants/leaf_cert"
result=$(echo "test1234" | zblcli tx pki add-x509-cert --certificate="$leaf_path" --from $trustee_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Request Leaf certificate"
result=$(zblcli query pki x509-cert --subject="$leaf_cert_subject" --subject-key-id="$leaf_cert_subject_key_id")
check_response "$result" "\"subject\": \"$leaf_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$leaf_cert_subject_key_id\""
check_response "$result" "\"serial_number\": \"$leaf_cert_serial_number\""
echo "$result"

echo "Request certificate chain for Leaf certificate"
result=$(zblcli query pki x509-cert-chain --subject="$leaf_cert_subject" --subject-key-id="$leaf_cert_subject_key_id")
check_response "$result" "\"subject\": \"$leaf_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$leaf_cert_subject_key_id\""
check_response "$result" "\"serial_number\": \"$leaf_cert_serial_number\""
check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$intermediate_cert_subject_key_id\""
check_response "$result" "\"serial_number\": \"$intermediate_cert_serial_number\""
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$root_cert_subject_key_id\""
check_response "$result" "\"serial_number\": \"$root_cert_serial_number\""
echo "$result"

echo "Request all proposed Root certificates must be empty"
result=$(zblcli query pki all-proposed-x509-root-certs)
check_response "$result" "\"total\": \"0\""

echo "Request all active certificates"
result=$(zblcli query pki all-x509-certs)
check_response "$result" "\"total\": \"3\""
check_response "$result" "\"subject_key_id\": \"$root_cert_subject_key_id\""
check_response "$result" "\"subject_key_id\": \"$intermediate_cert_subject_key_id\""
check_response "$result" "\"subject_key_id\": \"$leaf_cert_subject_key_id\""
echo "$result"

echo "Request all active root certificates"
result=$(zblcli query pki all-x509-root-certs)
check_response "$result" "\"total\": \"1\""
check_response "$result" "\"subject_key_id\": \"$root_cert_subject_key_id\""
echo "$result"

echo "Request all subject certificates"
result=$(zblcli query pki all-subject-x509-certs --subject="$leaf_cert_subject")
check_response "$result" "\"total\": \"1\""
check_response "$result" "\"subject_key_id\": \"$leaf_cert_subject_key_id\""
echo "$result"
