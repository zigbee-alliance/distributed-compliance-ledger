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

root_cert_subject="MDQxCzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpzb21lLXN0YXRlMRAwDgYDVQQKEwdyb290LWNh"
root_cert_subject_key_id="DF:4E:AF:B0:8C:9C:37:78:1A:E7:53:12:CA:E4:78:6B:48:1E:AF:B0"
root_cert_serial_number="81311506302208030248766861785118937702312370677"
vid=1

intermediate_cert_subject="MDwxCzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpzb21lLXN0YXRlMRgwFgYDVQQKEw9pbnRlcm1lZGlhdGUtY2E="
intermediate_cert_subject_key_id="1B:73:2A:91:34:46:8A:90:2A:87:19:91:E4:BD:8F:69:3A:F9:04:77"
intermediate_cert_serial_number="486736128900935106101503663840421220667833341899"

leaf_cert_subject="MDExCzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpzb21lLXN0YXRlMQ0wCwYDVQQKEwRsZWFm"
leaf_cert_subject_key_id="2A:31:8D:39:6E:50:DA:96:DF:95:C5:98:83:68:F0:58:B2:15:B3:3A"
leaf_cert_serial_number="409691117370409054634487600348183880852961428328"

# Preparation of Actors

first_trustee_account="jack"
second_trustee_account="alice"
third_trustee_account="bob"

first_trustee_address=$(echo $passphrase | dcld keys show jack -a)
second_trustee_address=$(echo $passphrase | dcld keys show alice -a)
third_trustee_address=$(echo $passphrase | dcld keys show bob -a)

echo "Create regular account"
create_new_account user_account "CertificationCenter"

test_divider

echo "Add 3 new Trustee accounts, this will result in a total of 6 Trustees and 4 approvals needed for 2/3 quorum"

test_divider

echo "Add Fourth Trustee Account"
random_string fourth_trustee_account
echo "$fourth_trustee_account generates keys"
result=$(dcld keys add $fourth_trustee_account)
fourth_trustee_address=$(echo $passphrase | dcld keys show $fourth_trustee_account -a)
fourth_trustee_pubkey=$(echo $passphrase | dcld keys show $fourth_trustee_account -p)

echo "$first_trustee_account proposes account for $fourth_trustee_account"
result=$(echo $passphrase | dcld tx auth propose-add-account --info="Jack is proposing this account" --address="$fourth_trustee_address" --pubkey="$fourth_trustee_pubkey" --roles="Trustee" --from $first_trustee_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "$second_trustee_account approves account for $fourth_trustee_account"
result=$(echo $passphrase | dcld tx auth approve-add-account --address="$fourth_trustee_address" --from $second_trustee_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Verify that the account is now present"
result=$(echo $passphrase | dcld query auth account --address="$fourth_trustee_address")
check_response "$result" "\"address\": \"$fourth_trustee_address\""

test_divider

echo "Add Fifth Trustee Account"
random_string fifth_trustee_account
echo "$fifth_trustee_account generates keys"
result=$(dcld keys add $fifth_trustee_account)

fifth_trustee_address=$(echo $passphrase | dcld keys show $fifth_trustee_account -a)
fifth_trustee_pubkey=$(echo $passphrase | dcld keys show $fifth_trustee_account -p)

echo "$first_trustee_account proposes account for $fifth_trustee_account"
result=$(echo $passphrase | dcld tx auth propose-add-account --info="Jack is proposing this account" --address="$fifth_trustee_address" --pubkey="$fifth_trustee_pubkey" --roles="Trustee" --from $first_trustee_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "$second_trustee_account approves account for $fifth_trustee_account"
result=$(echo $passphrase | dcld tx auth approve-add-account --address="$fifth_trustee_address" --from $second_trustee_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "$third_trustee_account approves account for $fifth_trustee_account"
result=$(echo $passphrase | dcld tx auth approve-add-account --address="$fifth_trustee_address" --from $third_trustee_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Verify that fifth account is now present"
result=$(echo $passphrase | dcld query auth account --address="$fifth_trustee_address")
check_response "$result" "\"address\": \"$fifth_trustee_address\""

test_divider

echo "Add sixth Trustee Account"
random_string sixth_trustee_account
echo "$sixth_trustee_account generates keys"
result=$(dcld keys add $sixth_trustee_account)
sixth_trustee_address=$(echo $passphrase | dcld keys show $sixth_trustee_account -a)
sixth_trustee_pubkey=$(echo $passphrase | dcld keys show $sixth_trustee_account -p)

echo "$first_trustee_account proposes account for $sixth_trustee_account"
result=$(echo $passphrase | dcld tx auth propose-add-account --info="Jack is proposing this account" --address="$sixth_trustee_address" --pubkey="$sixth_trustee_pubkey" --roles="Trustee" --from $first_trustee_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "$second_trustee_account approves account for $sixth_trustee_account"
result=$(echo $passphrase | dcld tx auth approve-add-account --address="$sixth_trustee_address" --from $second_trustee_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "$third_trustee_account approves account for $sixth_trustee_account"
result=$(echo $passphrase | dcld tx auth approve-add-account --address="$sixth_trustee_address" --from $third_trustee_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "$fourth_trustee_account approves account for $sixth_trustee_account"
result=$(echo $passphrase | dcld tx auth approve-add-account --address="$sixth_trustee_address" --from $fourth_trustee_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Verify that sixth account is now present"
result=$(echo $passphrase | dcld query auth account --address="$sixth_trustee_address")
check_response "$result" "\"address\": \"$sixth_trustee_address\""

test_divider

echo "PROPOSE ROOT CERT"
echo "$user_account (Not Trustee) propose Root certificate"
root_path="integration_tests/constants/root_cert"
result=$(echo "$passphrase" | dcld tx pki propose-add-x509-root-cert --certificate="$root_path" --vid $vid --from $user_account --yes)
result=$(get_txn_result "$result")
response_does_not_contain "$result" "\"code\": 0"

echo "$fourth_trustee_account (Trustee) propose Root certificate"
result=$(echo "$passphrase" | dcld tx pki propose-add-x509-root-cert --certificate="$root_path" --vid $vid --from $fourth_trustee_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve Root certificate now 4 Approvals are needed as we have 6 trustees"

echo "$first_trustee_account (Trustee) approve Root certificate"
result=$(echo "$passphrase" | dcld tx pki approve-add-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from $first_trustee_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Verify Root certificate is not present in approved state"
result=$(echo "$passphrase" | dcld query pki x509-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")  
check_response "$result" "Not Found"

test_divider

echo "$second_trustee_account (Trustee) approve Root certificate"
result=$(echo "$passphrase" | dcld tx pki approve-add-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from $second_trustee_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Verify Root certificate is not present in approved state"
result=$(echo "$passphrase" | dcld query pki x509-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")  
check_response "$result" "Not Found"

test_divider

echo "$third_trustee_account (Trustee) approve Root certificate"
result=$(echo "$passphrase" | dcld tx pki approve-add-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from $third_trustee_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Verify Root certificate is in approved state"
result=$(echo "$passphrase" | dcld query pki x509-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")  
echo $result | jq
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"address\": \"$first_trustee_address\""
check_response "$result" "\"address\": \"$second_trustee_address\""
check_response "$result" "\"address\": \"$third_trustee_address\""

test_divider

echo "$sixth_trustee_account proposes revoke Root certificate"
result=$(echo "$passphrase" | dcld tx pki propose-revoke-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from $sixth_trustee_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Request root certificate proposed to revoke and verify that it contains approval from $sixth_trustee_account"
result=$(dcld query pki proposed-x509-root-cert-to-revoke --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"$root_cert_subject\""
check_response "$result" "\"$root_cert_subject_key_id\""
check_response "$result" "\"address\": \"$sixth_trustee_address\""

test_divider

echo "$fifth_trustee_account revokes Root certificate"
result=$(echo "$passphrase" | dcld tx pki approve-revoke-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from $fifth_trustee_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Request root certificate proposed to revoke and verify that it contains approval from $fifth_trustee_account"
result=$(dcld query pki proposed-x509-root-cert-to-revoke --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"$root_cert_subject\""
check_response "$result" "\"$root_cert_subject_key_id\""
check_response "$result" "\"address\": \"$sixth_trustee_address\""
check_response "$result" "\"address\": \"$fifth_trustee_address\""

echo "$fourth_trustee_account revokes Root certificate"
result=$(echo "$passphrase" | dcld tx pki approve-revoke-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from $fourth_trustee_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Request root certificate proposed to revoke and verify that it contains approval from $fourth_trustee_account"
result=$(dcld query pki proposed-x509-root-cert-to-revoke --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"$root_cert_subject\""
check_response "$result" "\"$root_cert_subject_key_id\""
check_response "$result" "\"address\": \"$sixth_trustee_address\""
check_response "$result" "\"address\": \"$fifth_trustee_address\""
check_response "$result" "\"address\": \"$fourth_trustee_address\""

echo "$third_trustee_account revokes Root certificate"
result=$(echo "$passphrase" | dcld tx pki approve-revoke-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from $third_trustee_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Verify Root certificate is now revoked"
result=$(echo "$passphrase" | dcld query pki revoked-x509-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"$root_cert_subject\""
check_response "$result" "\"$root_cert_subject_key_id\""
check_response "$result" "\"address\": \"$sixth_trustee_address\""
check_response "$result" "\"address\": \"$fifth_trustee_address\""
check_response "$result" "\"address\": \"$fourth_trustee_address\""
check_response "$result" "\"address\": \"$third_trustee_address\""
