set -euo pipefail
source integration_tests/cli/common.sh

noc_root_cert_1_path="integration_tests/constants/noc_root_cert_1"
noc_root_cert_1_subject="MIGCMQswCQYDVQQGEwJVUzERMA8GA1UECAwITmV3IFlvcmsxETAPBgNVBAcMCE5ldyBZb3JrMRgwFgYDVQQKDA9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsMEFRlc3RpbmcgRGl2aXNpb24xGDAWBgNVBAMMD3d3dy5leGFtcGxlLmNvbQ=="
noc_root_cert_1_subject_key_id="44:EB:4C:62:6B:25:48:CD:A2:B3:1C:87:41:5A:08:E7:2B:B9:83:26"
noc_root_cert_1_serial_number="1"
noc_root_cert_1_subject_as_text="CN=www.example.com,OU=Testing Division,O=Example Company,L=New York,ST=New York,C=US"

noc_root_cert_2_path="integration_tests/constants/noc_root_cert_2"
noc_root_cert_2_subject="MIGCMQswCQYDVQQGEwJVUzERMA8GA1UECAwITmV3IFlvcmsxETAPBgNVBAcMCE5ldyBZb3JrMRgwFgYDVQQKDA9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsMEFRlc3RpbmcgRGl2aXNpb24xGDAWBgNVBAMMD3d3dy5leGFtcGxlLmNvbQ=="
noc_root_cert_2_subject_key_id="CF:E6:DD:37:2B:4C:B2:B9:A9:F2:75:30:1C:AA:B1:37:1B:11:7F:1B"
noc_root_cert_2_serial_number="2"
noc_root_cert_2_subject_as_text="CN=www.example.com,OU=Testing Division,O=Example Company,L=New York,ST=New York,C=US"

noc_root_cert_3_path="integration_tests/constants/noc_root_cert_3"
noc_root_cert_3_subject="MFUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQxDjAMBgNVBAMMBU5PQy0z"
noc_root_cert_3_subject_key_id="88:0D:06:D9:64:22:29:34:78:7F:8C:3B:AE:F5:08:93:86:8F:0D:20"
noc_root_cert_3_serial_number="38457288443253426021793906708335409501754677187"
noc_root_cert_3_subject_as_text="CN=NOC-3,O=Internet Widgits Pty Ltd,ST=Some-State,C=AU"

noc_cert_1_path="integration_tests/constants/noc_cert_1"
noc_cert_1_subject="MIGCMQswCQYDVQQGEwJVUzERMA8GA1UECAwITmV3IFlvcmsxETAPBgNVBAcMCE5ldyBZb3JrMRgwFgYDVQQKDA9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsMEFRlc3RpbmcgRGl2aXNpb24xGDAWBgNVBAMMD3d3dy5leGFtcGxlLmNvbQ=="
noc_cert_1_subject_key_id="02:72:6E:BC:BB:EF:D6:BD:8D:9B:42:AE:D4:3C:C0:55:5F:66:3A:B3"
noc_cert_1_serial_number="3"

noc_cert_2_path="integration_tests/constants/noc_cert_2"
noc_cert_2_subject="MIGCMQswCQYDVQQGEwJVUzERMA8GA1UECAwITmV3IFlvcmsxETAPBgNVBAcMCE5ldyBZb3JrMRgwFgYDVQQKDA9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsMEFRlc3RpbmcgRGl2aXNpb24xGDAWBgNVBAMMD3d3dy5leGFtcGxlLmNvbQ=="
noc_cert_2_subject_key_id="87:48:A2:33:12:1F:51:5C:93:E6:90:40:4A:2C:AB:9E:D6:19:E5:AD"
noc_cert_2_serial_number="4"

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

echo "Request NOC certificate by VID must be empty"
result=$(dcld query pki noc-x509-root-certs --vid="$vid")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$noc_root_cert_1_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$noc_root_cert_1_subject_as_text\""
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
result=$(dcld query pki x509-cert --subject="$noc_root_cert_1_subject" --subject-key-id="$noc_root_cert_1_subject_key_id")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$noc_root_cert_1_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$noc_root_cert_1_subject_as_text\""
echo $result | jq

test_divider

echo "Request all certificates by subject must be empty"
result=$(dcld query pki all-subject-x509-certs --subject="$noc_root_cert_1_subject")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$noc_root_cert_1_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
echo $result | jq

test_divider

echo "Request all certificates by subjectKeyId must be empty"
result=$(dcld query pki x509-cert --subject-key-id="$noc_root_cert_1_subject_key_id")
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
check_response "$result" "\"code\": 414"

echo "Add first NOC root certificate by vendor with VID = $vid"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-root-cert --certificate="$noc_root_cert_1_path" --from $vendor_account --yes)
check_response "$result" "\"code\": 0"

echo "Add second NOC root certificate by vendor with VID = $vid"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-root-cert --certificate="$noc_root_cert_2_path" --from $vendor_account --yes)
check_response "$result" "\"code\": 0"

echo "Add third NOC root certificate by vendor with VID = $vid_2"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-root-cert --certificate="$noc_root_cert_3_path" --from $vendor_account_2 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Request NOC root certificate by VID"
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
check_response "$result" "\"vid\": $vid"

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

echo "Request NOC root certificate by Subject and SubjectKeyID"
result=$(dcld query pki x509-cert --subject="$noc_root_cert_1_subject" --subject-key-id="$noc_root_cert_1_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$noc_root_cert_1_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
check_response "$result" "\"subjectAsText\": \"$noc_root_cert_1_subject_as_text\""
check_response "$result" "\"approvals\": \\[\\]"

test_divider

echo "Request NOC root certificate by Subject"
result=$(dcld query pki all-subject-x509-certs --subject="$noc_root_cert_1_subject")
echo $result | jq
check_response "$result" "\"$noc_root_cert_1_subject\""
check_response "$result" "\"$noc_root_cert_1_subject_key_id\""

test_divider

echo "Request NOC root certificate by SubjectKeyID"
result=$(dcld query pki x509-cert --subject-key-id="$noc_root_cert_1_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$noc_root_cert_1_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
check_response "$result" "\"subjectAsText\": \"$noc_root_cert_1_subject_as_text\""

echo "Add first NOC certificate by vendor with VID = $vid"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-cert --certificate="$noc_cert_1_path" --from $vendor_account --yes)
check_response "$result" "\"code\": 0"

echo "Request NOC certificate by VID = $vid"
result=$(dcld query pki noc-x509-certs --vid="$vid")
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
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-cert --certificate="$noc_cert_2_path" --from $vendor_account_2 --yes)
check_response "$result" "\"code\": 439"

test_divider

echo "Add second NOC certificate by vendor with VID = $vid"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-cert --certificate="$noc_cert_2_path" --from $vendor_account --yes)
check_response "$result" "\"code\": 0"

echo "Request all NOC certificates"
result=$(dcld query pki all-noc-x509-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$noc_cert_1_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_cert_1_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_cert_1_serial_number\""
check_response "$result" "\"subject\": \"$noc_cert_2_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_cert_2_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_cert_2_serial_number\""
check_response "$result" "\"vid\": $vid"

echo "Request all approved certificates"
result=$(dcld query pki all-x509-certs)
check_response "$result" "\"subject\": \"$noc_root_cert_1_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_root_cert_1_serial_number\""
check_response "$result" "\"subject\": \"$noc_cert_1_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_cert_1_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_cert_1_serial_number\""
check_response "$result" "\"subject\": \"$noc_cert_2_subject\""
check_response "$result" "\"subjectKeyId\": \"$noc_cert_2_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$noc_cert_2_serial_number\""
echo $result | jq

test_divider