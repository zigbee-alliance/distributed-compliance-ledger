set -euo pipefail
source integration_tests/cli/common.sh

root_cert_subject="MIGCMQswCQYDVQQGEwJVUzERMA8GA1UECAwITmV3IFlvcmsxETAPBgNVBAcMCE5ldyBZb3JrMRgwFgYDVQQKDA9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsMEFRlc3RpbmcgRGl2aXNpb24xGDAWBgNVBAMMD3d3dy5leGFtcGxlLmNvbQ=="
root_cert_subject_key_id="33:5E:0C:07:44:F8:B5:9C:CD:55:01:9B:6D:71:23:83:6F:D0:D4:BE"
root_cert_1_path="integration_tests/constants/root_with_same_subject_and_skid_1"
root_cert_1_serial_number="1"
root_cert_vid=65521
intermediate_cert_subject="MEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQ="
intermediate_cert_subject_key_id="2E:13:3B:44:52:2C:30:E9:EC:FB:45:FA:5D:E5:04:0A:C1:C6:E6:B9"
intermediate_cert_1_path="integration_tests/constants/intermediate_with_same_subject_and_skid_1"
intermediate_cert_2_path="integration_tests/constants/intermediate_with_same_subject_and_skid_2"
intermediate_cert_1_serial_number="3"
intermediate_cert_2_serial_number="4"
leaf_cert_subject="MEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQ="
leaf_cert_subject_key_id="12:16:55:8E:5E:2A:DF:04:D7:E6:FE:D1:53:69:61:98:EF:17:2F:03"
leaf_cert_path="integration_tests/constants/leaf_with_same_subject_and_skid"
leaf_cert_serial_number="5"

trustee_account="jack"
second_trustee_account="alice"

test_divider

echo "REMOVE X509 CERTIFICATES"

vendor_account_65521=vendor_account_$root_cert_vid
echo "Create Vendor account - $vendor_account_65521"
create_new_vendor_account $vendor_account_65521 $root_cert_vid

vendor_account_65522=vendor_account_65522
echo "Create Vendor account - $vendor_account_65522"
create_new_vendor_account $vendor_account_65522 65522

echo "Propose and approve root certificate 1"
result=$(echo "$passphrase" | dcld tx pki propose-add-x509-root-cert --certificate="$root_cert_1_path" --vid "$root_cert_vid" --from $trustee_account --yes)
check_response "$result" "\"code\": 0"
result=$(echo "$passphrase" | dcld tx pki approve-add-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from $second_trustee_account --yes)
check_response "$result" "\"code\": 0"

echo "Add an intermediate certificate with serialNumber 3"
result=$(echo "$passphrase" | dcld tx pki add-x509-cert --certificate="$intermediate_cert_1_path" --from $vendor_account_65521 --yes)
check_response "$result" "\"code\": 0"

echo "Add an intermediate certificate with serialNumber 4"
result=$(echo "$passphrase" | dcld tx pki add-x509-cert --certificate="$intermediate_cert_2_path" --from $vendor_account_65521 --yes)
check_response "$result" "\"code\": 0"

echo "Add a leaf certificate with serialNumber 5"
result=$(echo "$passphrase" | dcld tx pki add-x509-cert --certificate="$leaf_cert_path" --from $vendor_account_65521 --yes)
check_response "$result" "\"code\": 0"

echo "Request all approved root certificates."
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

echo "Revoke an intermediate certificate with serialNumber 3"
result=$(echo "$passphrase" | dcld tx pki revoke-x509-cert --subject="$intermediate_cert_subject" --subject-key-id="$intermediate_cert_subject_key_id" --serial-number="$intermediate_cert_1_serial_number" --from=$vendor_account_65521 --yes)
check_response "$result" "\"code\": 0"

echo "Request all revoked certificates should contain only one intermediate certificate with serialNumber 3"
result=$(dcld query pki all-revoked-x509-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$intermediate_cert_1_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$intermediate_cert_2_serial_number\""

echo "Remove intermediate certificate with invalid serialNumber"
result=$(echo "$passphrase" | dcld tx pki remove-x509-cert --subject="$intermediate_cert_subject" --subject-key-id="$intermediate_cert_subject_key_id" --serial-number="invalid" --from=$vendor_account_65521 --yes)
check_response "$result" "\"code\": 404"

echo "Try to remove the intermediate certificate using an account that does not have vendor role"
result=$(echo "$passphrase" | dcld tx pki remove-x509-cert --subject="$intermediate_cert_subject" --subject-key-id="$intermediate_cert_subject_key_id" --serial-number="$intermediate_cert_1_serial_number" --from=$trustee_account --yes)
check_response "$result" "\"code\": 4"

echo "Try to remove the intermediate certificate using a vendor account with other VID"
result=$(echo "$passphrase" | dcld tx pki remove-x509-cert --subject="$intermediate_cert_subject" --subject-key-id="$intermediate_cert_subject_key_id" --serial-number="$intermediate_cert_1_serial_number" --from=$vendor_account_65522 --yes)
check_response "$result" "\"code\": 4"

echo "Remove intermediate certificate with serialNumber 3"
result=$(echo "$passphrase" | dcld tx pki remove-x509-cert --subject="$intermediate_cert_subject" --subject-key-id="$intermediate_cert_subject_key_id" --serial-number="$intermediate_cert_1_serial_number" --from=$vendor_account_65521 --yes)
check_response "$result" "\"code\": 0"

echo "Request all certificates should not contain intermediate certificate with serialNumber 3"
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

echo "Request approved certificates by an intermediate certificate's subject and subjectKeyId should contain only one certificate with serialNumber 4"
result=$(dcld query pki x509-cert --subject="$intermediate_cert_subject" --subject-key-id="$intermediate_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$intermediate_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$intermediate_cert_2_serial_number\""
response_does_not_contain "$result" "\"serialNumber\": \"$intermediate_cert_1_serial_number\""

echo "Remove an intermediate certificate with subject and subjectKeyId"
result=$(echo "$passphrase" | dcld tx pki remove-x509-cert --subject="$intermediate_cert_subject" --subject-key-id="$intermediate_cert_subject_key_id" --from=$vendor_account_65521 --yes)
check_response "$result" "\"code\": 0"

echo "Request approved certificates by an intermediate certificate's subject and subjectKeyId should be empty"
result=$(dcld query pki x509-cert --subject="$intermediate_cert_subject" --subject-key-id="$intermediate_cert_subject_key_id")
echo $result | jq
check_response "$result" "Not Found"
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
result=$(echo "$passphrase" | dcld tx pki remove-x509-cert --subject="$leaf_cert_subject" --subject-key-id="$leaf_cert_subject_key_id" --from=$vendor_account_65521 --yes)
check_response "$result" "\"code\": 0"

echo "Request approved leaf certificates should be empty"
result=$(dcld query pki x509-cert --subject="$leaf_cert_subject" --subject-key-id="$leaf_cert_subject_key_id")
echo $result | jq
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$leaf_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$leaf_cert_subject_key_id\""
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

test_divider
