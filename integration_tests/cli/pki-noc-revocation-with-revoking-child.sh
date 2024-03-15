set -euo pipefail
source integration_tests/cli/common.sh

noc_root_cert_1_path="integration_tests/constants/noc_root_cert_1"
noc_root_cert_1_subject="MHoxCzAJBgNVBAYTAlVaMRMwEQYDVQQIDApTb21lIFN0YXRlMREwDwYDVQQHDAhUYXNoa2VudDEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMQ4wDAYDVQQDDAVOT0MtMQ=="
noc_root_cert_1_subject_key_id="44:EB:4C:62:6B:25:48:CD:A2:B3:1C:87:41:5A:08:E7:2B:B9:83:26"
noc_root_cert_1_serial_number="47211865327720222621302679792296833381734533449"

noc_root_cert_1_copy_path="integration_tests/constants/noc_root_cert_1_copy"
noc_root_cert_1_copy_serial_number="460647353168152946606945669687905527879095841977"

noc_cert_1_path="integration_tests/constants/noc_cert_1"
noc_cert_1_subject="MIGCMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTETMBEGA1UEBwwKU29tZSBTdGF0ZTEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRQwEgYDVQQDDAtOT0MtY2hpbGQtMQ=="
noc_cert_1_subject_key_id="02:72:6E:BC:BB:EF:D6:BD:8D:9B:42:AE:D4:3C:C0:55:5F:66:3A:B3"
noc_cert_1_serial_number="631388393741945881054190991612463928825155142122"

noc_leaf_cert_1_path="integration_tests/constants/noc_leaf_cert_1"
noc_leaf_cert_1_subject="MIGBMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTETMBEGA1UEBwwKU29tZSBTdGF0ZTEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRMwEQYDVQQDDApOT0MtbGVhZi0x"
noc_leaf_cert_1_subject_key_id="77:1F:DB:C4:4C:B1:29:7E:3C:EB:3E:D8:2A:38:0B:63:06:07:00:01"
noc_leaf_cert_1_serial_number="281347277961838999749763518155363401757954575313"

vid_in_hex_format=0x6006
vid=24582

vendor_account=vendor_account_$vid_in_hex_format
echo "Create Vendor account - $vendor_account"
create_new_vendor_account $vendor_account $vid_in_hex_format

test_divider

echo "Add first NOC root certificate by vendor with VID = $vid"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-root-cert --certificate="$noc_root_cert_1_path" --from $vendor_account --yes)
check_response "$result" "\"code\": 0"

echo "Add second NOC root certificate by vendor with VID = $vid"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-root-cert --certificate="$noc_root_cert_1_copy_path" --from $vendor_account --yes)
check_response "$result" "\"code\": 0"

echo "Add first NOC certificate by vendor with VID = $vid"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-cert --certificate="$noc_cert_1_path" --from $vendor_account --yes)
check_response "$result" "\"code\": 0"

echo "Add NOC leaf certificate by vendor with VID = $vid"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-cert --certificate="$noc_leaf_cert_1_path" --from $vendor_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Request All NOC root certificate"
result=$(dcld query pki all-noc-x509-root-certs)
echo $result | jq
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_copy_serial_number\""

echo "Request all NOC certificates"
result=$(dcld query pki all-noc-x509-certs)
echo $result | jq
check_response "$result" "\"serialNumber\": \"$noc_cert_1_serial_number\""
check_response "$result" "\"serialNumber\": \"$noc_leaf_cert_1_serial_number\""

echo "$vendor_account Vendor revokes root NOC certificate by setting \"revoke-child\" flag to true, it should revoke child certificates too"
result=$(echo "$passphrase" | dcld tx pki revoke-noc-x509-root-cert --subject="$noc_root_cert_1_subject" --subject-key-id="$noc_root_cert_1_subject_key_id" --revoke-child=true --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"

echo "Request all revoked certificates should contain two root, one intermediate and one leaf certificates"
result=$(dcld query pki all-revoked-x509-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$noc_root_cert_1_subject"
check_response "$result" "\"subject\": \"$noc_cert_1_subject\""
check_response "$result" "\"subject\": \"$noc_leaf_cert_1_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
check_response "$result" "\"subjectKeyId\": \"$noc_cert_1_subject_key_id\""
check_response "$result" "\"subjectKeyId\": \"$noc_leaf_cert_1_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_copy_serial_number\""
check_response "$result" "\"serialNumber\": \"$noc_cert_1_serial_number\""
check_response "$result" "\"serialNumber\": \"$noc_leaf_cert_1_serial_number\""

echo "Request all revoked NOC root certificates should contain two root certificates"
result=$(dcld query pki all-revoked-noc-x509-root-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$noc_root_cert_1_subject"
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_copy_serial_number\""

echo "Request revoked NOC root certificate by subject and subjectKeyId should contain two root certificates"
result=$(dcld query pki revoked-noc-x509-root-cert --subject="$noc_root_cert_1_subject" --subject-key-id="$noc_root_cert_1_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$noc_root_cert_1_subject"
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_copy_serial_number\""

echo "Request all x509 root revoked certificates should not contain revoked NOC root certificates"
result=$(dcld query pki all-revoked-x509-root-certs)
response_does_not_contain "$result" "\"subject\": \"$noc_root_cert_1_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_copy_serial_number\""
echo $result | jq

echo "Request NOC root certificate by VID = $vid should be empty"
result=$(dcld query pki noc-x509-root-certs --vid="$vid")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$noc_root_cert_1_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_copy_serial_number\""
echo $result | jq

echo "Request all certificates by NOC root certificate's subject should be empty"
result=$(dcld query pki all-subject-x509-certs --subject="$noc_root_cert_1_subject")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"$noc_root_cert_1_subject\""
response_does_not_contain "$result" "\"$noc_root_cert_1_subject_key_id\""
echo $result | jq

echo "Request all certificates by NOC root certificate's subjectKeyId should be empty"
result=$(dcld query pki x509-cert --subject-key-id="$noc_root_cert_1_subject_key_id")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$noc_root_cert_1_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_copy_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
echo $result | jq

echo "Request NOC certificate by VID = $vid should be empty"
result=$(dcld query pki noc-x509-certs --vid="$vid")
echo $result | jq
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$noc_cert_1_subject\""
response_does_not_contain "$result" "\"subject\": \"$noc_cert_1_subject\""
response_does_not_contain "$result" "\"subject\": \"$noc_leaf_cert_1_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_leaf_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_cert_1_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_leaf_cert_1_serial_number\""

echo "Request all approved certificates should be empty"
result=$(dcld query pki all-x509-certs)
check_response "$result" "\[\]"
response_does_not_contain "$result" "\"subject\": \"$noc_root_cert_1_subject\""
response_does_not_contain "$result" "\"subject\": \"$noc_cert_1_subject\""
response_does_not_contain "$result" "\"subject\": \"$noc_leaf_cert_1_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_leaf_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_copy_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_cert_1_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_leaf_cert_1_serial_number\""
echo $result | jq

test_divider