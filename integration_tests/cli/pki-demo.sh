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

root_cert_subject="O=root-ca,ST=some-state,C=AU"
root_cert_subject_key_id="5A:88:E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:9:30:E6:2B:DB"
root_cert_serial_number="442314047376310867378175982234956458728610743315"

intermediate_cert_subject="O=intermediate-ca,ST=some-state,C=AU"
intermediate_cert_subject_key_id="4E:3B:73:F4:70:4D:C2:98:D:DB:C8:5A:5F:2:3B:BF:86:25:56:2B"
intermediate_cert_serial_number="169917617234879872371588777545667947720450185023"

leaf_cert_subject="O=leaf,ST=some-state,C=AU"
leaf_cert_subject_key_id="30:F4:65:75:14:20:B2:AF:3D:14:71:17:AC:49:90:93:3E:24:A0:1F"
leaf_cert_serial_number="143290473708569835418599774898811724528308722063"

# Preparation of Actors

trustee_account="jack"
second_trustee_account="alice"

echo "Create regular account"
create_new_account user_account ""

# Body

echo "$user_account (Not Trustee) propose Root certificate"
root_path="integration_tests/constants/root_cert"
result=$(echo "test1234" | dclcli tx pki propose-add-x509-root-cert --certificate="$root_path" --from $user_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

test_divider

echo "Request all proposed Root certificates"
result=$(dclcli query pki all-proposed-x509-root-certs)
check_response "$result" "\"total\": \"1\""
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$root_cert_subject_key_id\""
echo "$result"

test_divider

echo "Request proposed Root certificate"
result=$(dclcli query pki proposed-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$root_cert_subject_key_id\""
check_response "$result" "\"serial_number\": \"$root_cert_serial_number\""
echo "$result"

test_divider

echo "Request all approved certificates must be empty"
result=$(dclcli query pki all-x509-certs)
check_response "$result" "\"total\": \"0\""
echo "$result"

test_divider

echo "Request all approved root certificates must be empty"
result=$(dclcli query pki all-x509-root-certs)
check_response "$result" "\"total\": \"0\""
echo "$result"

test_divider

echo "$trustee_account (Trustee) approve Root certificate"
result=$(echo "test1234" | dclcli tx pki approve-add-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from $trustee_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

test_divider

echo "Certificate must be still in Proposed state. Request proposed Root certificate"
result=$(dclcli query pki proposed-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$root_cert_subject_key_id\""
check_response "$result" "\"serial_number\": \"$root_cert_serial_number\""
check_response "$result" "[\"$(dclcli keys show jack -a)\"]"
echo "$result"

test_divider

echo "Request all approved certificates must be empty"
result=$(dclcli query pki all-x509-certs)
check_response "$result" "\"total\": \"0\""
echo "$result"

test_divider

echo "$second_trustee_account (Trustee) approve Root certificate"
result=$(echo "test1234" | dclcli tx pki approve-add-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from $second_trustee_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

test_divider

echo "Certificate must be Approved. Request Root certificate"
result=$(dclcli query pki x509-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$root_cert_subject_key_id\""
check_response "$result" "\"serial_number\": \"$root_cert_serial_number\""
echo "$result"

test_divider

echo "Request certificate chain for Root certificate"
result=$(dclcli query pki x509-cert-chain --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$root_cert_subject_key_id\""
check_response "$result" "\"serial_number\": \"$root_cert_serial_number\""
echo "$result"

test_divider

echo "Request all proposed Root certificates must be empty"
result=$(dclcli query pki all-proposed-x509-root-certs)
check_response "$result" "\"total\": \"0\""
echo "$result"

test_divider

echo "Request all approved certificates"
result=$(dclcli query pki all-x509-certs)
check_response "$result" "\"total\": \"1\""
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$root_cert_subject_key_id\""
echo "$result"

test_divider

echo "$user_account (Not Trustee) add Intermediate certificate"
intermediate_path="integration_tests/constants/intermediate_cert"
result=$(echo "test1234" | dclcli tx pki add-x509-cert --certificate="$intermediate_path" --from $user_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

test_divider

echo "Request Intermediate certificate"
result=$(dclcli query pki x509-cert --subject="$intermediate_cert_subject" --subject-key-id="$intermediate_cert_subject_key_id")
check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$intermediate_cert_subject_key_id\""
check_response "$result" "\"serial_number\": \"$intermediate_cert_serial_number\""
echo "$result"

test_divider

echo "Request certificate chain for Intermediate certificate"
result=$(dclcli query pki x509-cert-chain --subject="$intermediate_cert_subject" --subject-key-id="$intermediate_cert_subject_key_id")
check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$intermediate_cert_subject_key_id\""
check_response "$result" "\"serial_number\": \"$intermediate_cert_serial_number\""
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$root_cert_subject_key_id\""
check_response "$result" "\"serial_number\": \"$root_cert_serial_number\""
echo "$result"

test_divider

echo "Request all proposed Root certificates must be empty"
result=$(dclcli query pki all-proposed-x509-root-certs)
check_response "$result" "\"total\": \"0\""

test_divider

echo "Request all approved certificates"
result=$(dclcli query pki all-x509-certs)
check_response "$result" "\"total\": \"2\""
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$root_cert_subject_key_id\""
check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$intermediate_cert_subject_key_id\""
echo "$result"

test_divider

echo "Request all approved root certificates"
result=$(dclcli query pki all-x509-root-certs)
check_response "$result" "\"total\": \"1\""
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$root_cert_subject_key_id\""
echo "$result"

test_divider

echo "$trustee_account (Trustee) add Leaf certificate"
leaf_path="integration_tests/constants/leaf_cert"
result=$(echo "test1234" | dclcli tx pki add-x509-cert --certificate="$leaf_path" --from $trustee_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

test_divider

echo "Request Leaf certificate"
result=$(dclcli query pki x509-cert --subject="$leaf_cert_subject" --subject-key-id="$leaf_cert_subject_key_id")
check_response "$result" "\"subject\": \"$leaf_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$leaf_cert_subject_key_id\""
check_response "$result" "\"serial_number\": \"$leaf_cert_serial_number\""
echo "$result"

test_divider

echo "Request certificate chain for Leaf certificate"
result=$(dclcli query pki x509-cert-chain --subject="$leaf_cert_subject" --subject-key-id="$leaf_cert_subject_key_id")
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

test_divider

echo "Request all proposed Root certificates must be empty"
result=$(dclcli query pki all-proposed-x509-root-certs)
check_response "$result" "\"total\": \"0\""

test_divider

echo "Request all approved certificates"
result=$(dclcli query pki all-x509-certs)
check_response "$result" "\"total\": \"3\""
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$root_cert_subject_key_id\""
check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$intermediate_cert_subject_key_id\""
check_response "$result" "\"subject\": \"$leaf_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$leaf_cert_subject_key_id\""
echo "$result"

test_divider

echo "Request all approved root certificates"
result=$(dclcli query pki all-x509-root-certs)
check_response "$result" "\"total\": \"1\""
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$root_cert_subject_key_id\""
echo "$result"

test_divider

echo "Request all subject certificates"
result=$(dclcli query pki all-subject-x509-certs --subject="$leaf_cert_subject")
check_response "$result" "\"total\": \"1\""
check_response "$result" "\"subject\": \"$leaf_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$leaf_cert_subject_key_id\""
echo "$result"

test_divider

echo "Request all root certificates proposed to revoke"
result=$(dclcli query pki all-proposed-x509-root-certs-to-revoke)
check_response "$result" "\"total\": \"0\""
echo "$result"

test_divider

echo "Request all revoked certificates"
result=$(dclcli query pki all-revoked-x509-certs)
check_response "$result" "\"total\": \"0\""
echo "$result"

test_divider

echo "$user_account (Not Trustee) revokes Intermediate certificate. This must also revoke its child - Leaf certificate."
result=$(echo "test1234" | dclcli tx pki revoke-x509-cert --subject="$intermediate_cert_subject" --subject-key-id="$intermediate_cert_subject_key_id" --from=$user_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

test_divider

echo "Request all root certificates proposed to revoke"
result=$(dclcli query pki all-proposed-x509-root-certs-to-revoke)
check_response "$result" "\"total\": \"0\""
echo "$result"

test_divider

echo "Request all revoked certificates"
result=$(dclcli query pki all-revoked-x509-certs)
check_response "$result" "\"total\": \"2\""
check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$intermediate_cert_subject_key_id\""
check_response "$result" "\"subject\": \"$leaf_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$leaf_cert_subject_key_id\""
echo "$result"

test_divider

echo "Request all revoked root certificates"
result=$(dclcli query pki all-revoked-x509-root-certs)
check_response "$result" "\"total\": \"0\""
echo "$result"

test_divider

echo "Request revoked Intermediate certificate"
result=$(dclcli query pki revoked-x509-cert --subject="$intermediate_cert_subject" --subject-key-id="$intermediate_cert_subject_key_id")
check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$intermediate_cert_subject_key_id\""
check_response "$result" "\"serial_number\": \"$intermediate_cert_serial_number\""
echo "$result"

test_divider

echo "Request revoked Leaf certificate"
result=$(dclcli query pki revoked-x509-cert --subject="$leaf_cert_subject" --subject-key-id="$leaf_cert_subject_key_id")
check_response "$result" "\"subject\": \"$leaf_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$leaf_cert_subject_key_id\""
check_response "$result" "\"serial_number\": \"$leaf_cert_serial_number\""
echo "$result"

test_divider

echo "Request all approved certificates"
result=$(dclcli query pki all-x509-certs)
check_response "$result" "\"total\": \"1\""
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$root_cert_subject_key_id\""
echo "$result"

test_divider

echo "$trustee_account (Trustee) proposes to revoke Root certificate"
result=$(echo "test1234" | dclcli tx pki propose-revoke-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from $trustee_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

test_divider

echo "Request all root certificates proposed to revoke"
result=$(dclcli query pki all-proposed-x509-root-certs-to-revoke)
check_response "$result" "\"total\": \"1\""
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$root_cert_subject_key_id\""
echo "$result"

test_divider

echo "Request all revoked certificates"
result=$(dclcli query pki all-revoked-x509-certs)
check_response "$result" "\"total\": \"2\""
check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$intermediate_cert_subject_key_id\""
check_response "$result" "\"subject\": \"$leaf_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$leaf_cert_subject_key_id\""
echo "$result"

test_divider

echo "Request all revoked root certificates"
result=$(dclcli query pki all-revoked-x509-root-certs)
check_response "$result" "\"total\": \"0\""
echo "$result"

test_divider

echo "Request Root certificate proposed to revoke"
result=$(dclcli query pki proposed-x509-root-cert-to-revoke --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$root_cert_subject_key_id\""
check_response "$result" "[\"$(dclcli keys show jack -a)\"]"
echo "$result"

test_divider

echo "Request all approved certificates"
result=$(dclcli query pki all-x509-certs)
check_response "$result" "\"total\": \"1\""
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$root_cert_subject_key_id\""
echo "$result"

test_divider

echo "$second_trustee_account (Trustee) approves to revoke Root certificate"
result=$(echo "test1234" | dclcli tx pki approve-revoke-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from $second_trustee_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

test_divider

echo "Request all root certificates proposed to revoke"
result=$(dclcli query pki all-proposed-x509-root-certs-to-revoke)
check_response "$result" "\"total\": \"0\""
echo "$result"

test_divider

echo "Request all revoked certificates"
result=$(dclcli query pki all-revoked-x509-certs)
check_response "$result" "\"total\": \"3\""
check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$intermediate_cert_subject_key_id\""
check_response "$result" "\"subject\": \"$leaf_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$leaf_cert_subject_key_id\""
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$root_cert_subject_key_id\""
echo "$result"

test_divider

echo "Request all revoked root certificates"
result=$(dclcli query pki all-revoked-x509-root-certs)
check_response "$result" "\"total\": \"1\""
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$root_cert_subject_key_id\""
echo "$result"

test_divider

echo "Request revoked Root certificate"
result=$(dclcli query pki revoked-x509-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$root_cert_subject_key_id\""
check_response "$result" "\"serial_number\": \"$root_cert_serial_number\""
echo "$result"

test_divider

echo "Request all approved certificates"
result=$(dclcli query pki all-x509-certs)
check_response "$result" "\"total\": \"0\""
echo "$result"
