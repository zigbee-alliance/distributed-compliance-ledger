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

unknown_cert_subj="O=unknown-ca,ST=some-state,C=AU"
unknown_cert_subject_key_id="68:99:E:76:36:53:D0:7F:B0:89:71:A3:F4:73:79:9:30:E6:2B:DB"

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

echo "Create regular account"
create_new_account user_account "CertificationCenter"

test_divider

echo "$user_account (Not Trustee) propose Root certificate"
root_path="integration_tests/constants/root_cert"
result=$(echo "$passphrase" | dcld tx pki propose-add-x509-root-cert --certificate="$root_path" --from $user_account --yes)
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "$trustee_account (Trustee) approve Root certificate"
result=$(echo $passphrase | dcld tx pki approve-add-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from $trustee_account --yes)
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "$second_trustee_account (Trustee) approve Root certificate"
result=$(echo "$passphrase" | dcld tx pki approve-add-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from $second_trustee_account --yes)
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "$user_account (Not Trustee) add Intermediate certificate"
intermediate_path="integration_tests/constants/intermediate_cert"
result=$(echo "$passphrase" | dcld tx pki add-x509-cert --certificate="$intermediate_path" --from $user_account --yes)
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "$trustee_account (Trustee) add Leaf certificate"
leaf_path="integration_tests/constants/leaf_cert"
result=$(echo "$passphrase" | dcld tx pki add-x509-cert --certificate="$leaf_path" --from $trustee_account --yes)
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "$trustee_account (Trustee) revokes Leaf certificate."
result=$(echo "$passphrase" | dcld tx pki revoke-x509-cert --subject="$leaf_cert_subject" --subject-key-id="$leaf_cert_subject_key_id" --from=$trustee_account --yes)
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

echo "$trustee_account (Trustee) proposes to revoke Root certificate"
result=$(echo "$passphrase" | dcld tx pki propose-revoke-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from $trustee_account --yes)
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
check_response "$result" "\"subject_key_id\": \"$root_cert_subject_key_id\""
check_response "$result" "\"serial_number\": \"$root_cert_serial_number\""

test_divider

echo "Query existent revoked cert"
result=$(execute_with_retry "dcld query pki revoked-x509-cert --subject=$leaf_cert_subject --subject-key-id=$leaf_cert_subject_key_id")
echo "$result"
check_response "$result" "\"subject\": \"$leaf_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$leaf_cert_subject_key_id\""
check_response "$result" "\"serial_number\": \"$leaf_cert_serial_number\""
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
check_response "$result" "\"subject_key_id\": \"$root_cert_subject_key_id\""

test_divider

echo "Query existent child certs"
result=$(execute_with_retry "dcld query pki all-child-x509-certs --subject=$root_cert_subject --subject-key-id=$root_cert_subject_key_id")
echo "$result"
check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$intermediate_cert_subject_key_id\""

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
result=$(echo "$passphrase" | dcld tx pki add-x509-cert --certificate="$intermediate_path" --from $user_account --yes)
echo "$result"
check_response "$result" "Write requests don't work with a Light Client Proxy"


test_divider
