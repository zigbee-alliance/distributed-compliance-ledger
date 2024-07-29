set -euo pipefail
source integration_tests/cli/common.sh

root_cert_1_path="integration_tests/constants/noc_root_cert_1"
root_cert_subject="MHoxCzAJBgNVBAYTAlVaMRMwEQYDVQQIDApTb21lIFN0YXRlMREwDwYDVQQHDAhUYXNoa2VudDEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMQ4wDAYDVQQDDAVOT0MtMQ=="
root_cert_subject_key_id="44:EB:4C:62:6B:25:48:CD:A2:B3:1C:87:41:5A:08:E7:2B:B9:83:26"
root_cert_1_serial_number="47211865327720222621302679792296833381734533449"

root_cert_1_copy_path="integration_tests/constants/noc_root_cert_1_copy"
root_cert_1_copy_serial_number="460647353168152946606945669687905527879095841977"

root_cert_vid=65521

intermediate_cert_1_path="integration_tests/constants/noc_cert_1"
intermediate_cert_2_path="integration_tests/constants/noc_cert_1_copy"
intermediate_cert_subject="MIGCMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTETMBEGA1UEBwwKU29tZSBTdGF0ZTEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRQwEgYDVQQDDAtOT0MtY2hpbGQtMQ=="
intermediate_cert_subject_key_id="02:72:6E:BC:BB:EF:D6:BD:8D:9B:42:AE:D4:3C:C0:55:5F:66:3A:B3"
intermediate_cert_1_serial_number="631388393741945881054190991612463928825155142122"
intermediate_cert_2_serial_number="169445068204646961882009388640343665944683778293"

leaf_cert_path="integration_tests/constants/noc_leaf_cert_1"
leaf_cert_subject="MIGBMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTETMBEGA1UEBwwKU29tZSBTdGF0ZTEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRMwEQYDVQQDDApOT0MtbGVhZi0x"
leaf_cert_subject_key_id="77:1F:DB:C4:4C:B1:29:7E:3C:EB:3E:D8:2A:38:0B:63:06:07:00:01"
leaf_cert_serial_number="281347277961838999749763518155363401757954575313"

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
result=$(dcld query pki all-x509-certs)
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
result=$(dcld query pki all-revoked-x509-certs)
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
result=$(dcld query pki all-x509-certs)
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
result=$(dcld query pki x509-cert --subject="$intermediate_cert_subject" --subject-key-id="$intermediate_cert_subject_key_id")
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
result=$(dcld query pki x509-cert --subject="$intermediate_cert_subject" --subject-key-id="$intermediate_cert_subject_key_id")
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
result=$(dcld query pki all-revoked-x509-certs)
echo $result | jq
check_response "$result" "\[\]"
response_does_not_contain "$result" "\"subject\": \"$intermediate_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$intermediate_cert_1_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$intermediate_cert_2_serial_number\""

echo "Request all certificates should contain only root and leaf certificates"
result=$(dcld query pki all-x509-certs)
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
result=$(dcld query pki x509-cert --subject="$leaf_cert_subject" --subject-key-id="$leaf_cert_subject_key_id")
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
result=$(dcld query pki all-x509-certs)
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
result=$(dcld query pki all-x509-certs)
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
result=$(dcld query pki all-revoked-x509-certs)
echo $result | jq
check_response "$result" "\"serialNumber\": \"$root_cert_1_serial_number\""

echo "Remove NOC root certificate with serialNumber $root_cert_1_serial_number"
result=$(echo "$passphrase" | dcld tx pki remove-noc-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --serial-number="$root_cert_1_serial_number" --from=$vendor_account_65521 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Request all certificates should contain only one NOC root certificate"
result=$(dcld query pki all-x509-certs)
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
result=$(dcld query pki x509-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"serialNumber\": \"$root_cert_1_copy_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$root_cert_1_serial_number\""

echo "Re-add NOC root certificate"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-root-cert --certificate="$root_cert_1_path" --from $vendor_account_65521 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Check that root cert is added. Request all approved certificates."
result=$(dcld query pki all-x509-certs)
echo $result | jq
check_response "$result" "\"serialNumber\": \"$root_cert_1_serial_number\""
check_response "$result" "\"serialNumber\": \"$root_cert_1_copy_serial_number\""

echo "Remove NOC root certificates"
result=$(echo "$passphrase" | dcld tx pki remove-noc-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from=$vendor_account_65521 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Request approved NOC root certificates should be empty"
result=$(dcld query pki x509-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
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
result=$(dcld query pki all-x509-certs)
echo $result | jq
check_response "$result" "\"serialNumber\": \"$intermediate_cert_1_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$root_cert_1_serial_number"
response_does_not_contain "$result" "\"serialNumber\": \"$root_cert_1_copy_serial_number"

echo "Request NOC certificate by VID = $vid and SKID = $intermediate_cert_subject_key_id should not be empty"
result=$(dcld query pki noc-x509-certs --vid="$vid" --skid="$intermediate_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"serialNumber\": \"$intermediate_cert_1_serial_number\""

echo "Request NOC certificate by VID = $vid and SKID = $root_cert_subject_key_id should be empty"
result=$(dcld query pki noc-x509-certs --vid="$vid" --skid="$root_cert_subject_key_id")
echo $result | jq
check_response "$result" "Not Found"

test_divider
