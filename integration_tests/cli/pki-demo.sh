#!/bin/bash
set -e
source integration_tests/cli/common.sh

root_cert_subject="CN=DST Root CA X3,O=Digital Signature Trust Co."
root_cert_subject_key_id="C4:A7:B1:A4:7B:2C:71:FA:DB:E1:4B:90:75:FF:C4:15:60:85:89:10"
root_cert_serial_number="91299735575339953335919266965803778155"

intermediate_cert_subject="CN=Let's Encrypt Authority X3,O=Let's Encrypt,C=US"
intermediate_cert_subject_key_id="A8:4A:6A:63:4:7D:DD:BA:E6:D1:39:B7:A6:45:65:EF:F3:A8:EC:A1"
intermediate_cert_serial_number="13298795840390663119752826058995181320"

leaf_cert_subject="CN=dsr-corporation.com"
leaf_cert_subject_key_id="8A:34:B:5C:D8:42:18:F2:C1:2A:AC:7A:B3:8F:6E:90:66:F4:4E:5C"
leaf_cert_serial_number="312128364102099997394566658874957944692446"

echo "Assign Trustee role to Jack"
result=$(echo "test1234" | zblcli tx auth assign-role --address=$(zblcli keys show jack -a) --role="Trustee" --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

trustee_account=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 6 | head -n 1)
echo "Create Trustee account with address: $trustee_account"
create_account_with_name $trustee_account
trustee_address=$(zblcli keys show "$trustee_account" -a)
result=$(echo "test1234" | zblcli tx auth assign-role --address=$trustee_address --role="Trustee" --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

user_account=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 6 | head -n 1)
echo "Create regular account with address: $user_account"
create_account_with_name $user_account
user_address=$(zblcli keys show "$user_account" -a)

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

echo "Request all proposed Root certificate"
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

echo "Jack (Trustee) approve Root certificate"
result=$(echo "test1234" | zblcli tx pki approve-add-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Certificate mut be still in Proposed state. Request proposed Root certificate"
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

echo "$trustee_account (Trustee) approve Root certificate"
result=$(echo "test1234" | zblcli tx pki approve-add-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from $trustee_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Certificate mut be Approved. Request Root certificate"
result=$(zblcli query pki x509-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
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

echo "$user_account (Not Trustee) add intermediate certificate"
intermediate_path="integration_tests/constants/intermediate_cert"
result=$(echo "test1234" | zblcli tx pki add-x509-cert --certificate="$intermediate_path" --from $user_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Request intermediate certificate"
result=$(zblcli query pki x509-cert --subject="$intermediate_cert_subject" --subject-key-id="$intermediate_cert_subject_key_id")
check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$intermediate_cert_subject_key_id\""
check_response "$result" "\"serial_number\": \"$intermediate_cert_serial_number\""
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

echo "$trustee_account (Trustee) add leaf certificate"
leaf_path="integration_tests/constants/leaf_cert"
result=$(echo "test1234" | zblcli tx pki add-x509-cert --certificate="$leaf_path" --from $trustee_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Request leaf certificate"
result=$(zblcli query pki x509-cert --subject="$leaf_cert_subject" --subject-key-id="$leaf_cert_subject_key_id")
check_response "$result" "\"subject\": \"$leaf_cert_subject\""
check_response "$result" "\"subject_key_id\": \"$leaf_cert_subject_key_id\""
check_response "$result" "\"serial_number\": \"$leaf_cert_serial_number\""
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
