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

unknown_cert_subj="Tz11bmtub3duLWNhLFNUPXNvbWUtc3RhdGUsQz1BVQ=="
unknown_cert_subject_key_id="68:99:0E:76:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB"

# 1. check non-existent values via light client when no entry added
echo "check non-existent values via light client when no entry added"

test_divider

# connect to light client proxy
dcld config node tcp://localhost:26620
sleep 10


echo "Query non existent cert"
result=$(execute_with_retry "dcld query pki x509-cert --subject=$root_cert_subject --subject-key-id=$root_cert_subject_key_id")
echo "$result"
check_response "$result" "Not Found"

test_divider

echo "Query non existent revoked cert"
result=$(execute_with_retry "dcld query pki revoked-x509-cert --subject=$root_cert_subject --subject-key-id=$root_cert_subject_key_id")
echo "$result"
check_response "$result" "Not Found"

test_divider


echo "Query non existent proposed root cert"
result=$(execute_with_retry "dcld query pki proposed-x509-root-cert --subject=$root_cert_subject --subject-key-id=$root_cert_subject_key_id")
echo "$result"
check_response "$result" "Not Found"

test_divider

echo "Query non existent cert by subject"
result=$(execute_with_retry "dcld query pki all-subject-x509-certs --subject=$root_cert_subject")
echo "$result"
check_response "$result" "Not Found"

test_divider

echo "Query non existent all root certs"
result=$(execute_with_retry "dcld query pki all-x509-root-certs")
echo "$result"
check_response "$result" "Not Found"

test_divider

echo "Query non existent all revoked root certs"
result=$(execute_with_retry "dcld query pki all-revoked-x509-root-certs")
echo "$result"
check_response "$result" "Not Found"

test_divider

echo "Query non existent proposed revoked cert"
result=$(execute_with_retry "dcld query pki proposed-x509-root-cert-to-revoke --subject=$root_cert_subject --subject-key-id=$root_cert_subject_key_id")
echo "$result"
check_response "$result" "Not Found"

test_divider

echo "Query non existent child certs"
result=$(execute_with_retry "dcld query pki all-child-x509-certs --subject=$root_cert_subject --subject-key-id=$root_cert_subject_key_id")
echo "$result"
check_response "$result" "Not Found"

test_divider

# 2. list queries should return a warning via light client

echo "list queries should return a warning via light client"

test_divider

echo "Request all approved certs"
result=$(execute_with_retry "dcld query pki all-x509-certs")
echo "$result"
check_response "$result" "List queries don't work with a Light Client Proxy"

test_divider

echo "Request all revoked certs"
result=$(execute_with_retry "dcld query pki all-revoked-x509-certs")
echo "$result"
check_response "$result" "List queries don't work with a Light Client Proxy"

test_divider

echo "Request all proposed certs"
result=$(execute_with_retry "dcld query pki all-proposed-x509-root-certs")
echo "$result"
check_response "$result" "List queries don't work with a Light Client Proxy"

test_divider

echo "Request all proposed revoked certs"
result=$(execute_with_retry "dcld query pki all-proposed-x509-root-certs-to-revoke")
echo "$result"
check_response "$result" "List queries don't work with a Light Client Proxy"

test_divider

# 3. write entries

echo "write entries"

test_divider

# write requests can be sent via Full Node only
dcld config node tcp://localhost:26657 


trustee_account="jack"
second_trustee_account="alice"

vendor_account=vendor_account_$vid
echo "Create Vendor account - $vendor_account"
create_new_vendor_account $vendor_account $vid

test_divider

echo "$vendor_account (Not Trustee) propose Root certificate"
root_path="integration_tests/constants/root_cert"
result=$(echo "$passphrase" | dcld tx pki propose-add-x509-root-cert --certificate="$root_path" --vid $vid --from $vendor_account --yes)
result=$(get_txn_result "$result")
response_does_not_contain "$result" "\"code\": 0"
echo "$result"

test_divider

echo "$trustee_account (Trustee) propose Root certificate"
root_path="integration_tests/constants/root_cert"
result=$(echo "$passphrase" | dcld tx pki propose-add-x509-root-cert --certificate="$root_path" --vid $vid --from $trustee_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "$second_trustee_account (Trustee) approve Root certificate"
result=$(echo "$passphrase" | dcld tx pki approve-add-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from $second_trustee_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "$vendor_account (Vendor) add Intermediate certificate"
intermediate_path="integration_tests/constants/intermediate_cert"
result=$(echo "$passphrase" | dcld tx pki add-x509-cert --certificate="$intermediate_path" --from $vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "$vendor_account (Vendor) add Leaf certificate"
leaf_path="integration_tests/constants/leaf_cert"
result=$(echo "$passphrase" | dcld tx pki add-x509-cert --certificate="$leaf_path" --from $vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "$vendor_account (Vendor) revokes Leaf certificate."
result=$(echo "$passphrase" | dcld tx pki revoke-x509-cert --subject="$leaf_cert_subject" --subject-key-id="$leaf_cert_subject_key_id" --from=$vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "$trustee_account (Trustee) proposes to revoke Root certificate"
result=$(echo "$passphrase" | dcld tx pki propose-revoke-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from $trustee_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

# 4. check existent values via light client

echo "check existent values via light client"

test_divider

# connect to light client proxy
dcld config node tcp://localhost:26620
sleep 5

echo "Query existent cert"
result=$(execute_with_retry "dcld query pki x509-cert --subject=$root_cert_subject --subject-key-id=$root_cert_subject_key_id")
echo "$result"
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$root_cert_serial_number\""

test_divider

echo "Query existent revoked cert"
result=$(execute_with_retry "dcld query pki revoked-x509-cert --subject=$leaf_cert_subject --subject-key-id=$leaf_cert_subject_key_id")
echo "$result"
check_response "$result" "\"subject\": \"$leaf_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$leaf_cert_serial_number\""
test_divider


echo "Query existent cert by subject"
result=$(execute_with_retry "dcld query pki all-subject-x509-certs --subject=$root_cert_subject")
echo "$result"
check_response "$result" "\"$root_cert_subject\""
check_response "$result" "\"$root_cert_subject_key_id\""

test_divider

echo "Query existent all root certs"
result=$(execute_with_retry "dcld query pki all-x509-root-certs")
echo "$result"
check_response "$result" "\"$root_cert_subject\""
check_response "$result" "\"$root_cert_subject_key_id\""
test_divider

echo "Query existent proposed revoked cert"
result=$(execute_with_retry "dcld query pki proposed-x509-root-cert-to-revoke --subject=$root_cert_subject --subject-key-id=$root_cert_subject_key_id")
echo "$result"
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""

test_divider

echo "Query existent child certs"
result=$(execute_with_retry "dcld query pki all-child-x509-certs --subject=$root_cert_subject --subject-key-id=$root_cert_subject_key_id")
echo "$result"
check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""

test_divider


# 5. check non-existent values via light client when entry added

echo "check non-existent values via light client when entry added"

test_divider

echo "Query non existent cert"
result=$(execute_with_retry "dcld query pki x509-cert --subject=$unknown_cert_subj --subject-key-id=$unknown_cert_subject_key_id")
echo "$result"
check_response "$result" "Not Found"

test_divider

echo "Query non existent proposed cert"
result=$(execute_with_retry "dcld query pki proposed-x509-root-cert --subject=$unknown_cert_subj --subject-key-id=$unknown_cert_subject_key_id")
echo "$result"
check_response "$result" "Not Found"

test_divider

echo "Query non existent revoked cert"
result=$(execute_with_retry "dcld query pki revoked-x509-cert --subject=$unknown_cert_subj --subject-key-id=$unknown_cert_subject_key_id")
echo "$result"
check_response "$result" "Not Found"

test_divider

echo "Query non existent cert by subject"
result=$(execute_with_retry "dcld query pki all-subject-x509-certs --subject=$unknown_cert_subj")
echo "$result"
check_response "$result" "Not Found"

test_divider

echo "Query non existent all revoked root certs"
result=$(execute_with_retry "dcld query pki all-revoked-x509-root-certs")
echo "$result"
check_response "$result" "\[\]"

test_divider

echo "Query non existent proposed revoked cert"
result=$(execute_with_retry "dcld query pki proposed-x509-root-cert-to-revoke --subject=$unknown_cert_subj --subject-key-id=$unknown_cert_subject_key_id")
echo "$result"
check_response "$result" "Not Found"

test_divider

echo "Query non existent child certs"
result=$(execute_with_retry "dcld query pki all-child-x509-certs --subject=$unknown_cert_subj --subject-key-id=$unknown_cert_subject_key_id")
echo "$result"
check_response "$result" "Not Found"

test_divider


# 6. try to write via light client proxy

echo "try to write via light client proxy"

test_divider

echo "Add cert"
result=$(echo "$passphrase" | dcld tx pki add-x509-cert --certificate="$intermediate_path" --from $vendor_account --yes)
echo "$result"
check_response "$result" "Write requests don't work with a Light Client Proxy"


test_divider
