set -euo pipefail
source integration_tests/cli/common.sh

root_cert_path="integration_tests/constants/root_cert"
root_cert_subject="MDQxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMRAwDgYDVQQKDAdyb290LWNh"
root_cert_subject_key_id="5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB"
root_cert_serial_number="442314047376310867378175982234956458728610743315"
root_cert_subject_as_text="O=root-ca,ST=some-state,C=AU"

google_root_path="integration_tests/constants/google_root_cert"
google_root_subject="MEsxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZHb29nbGUxFTATBgNVBAMMDE1hdHRlciBQQUEgMTEUMBIGCisGAQQBgqJ8AgEMBDYwMDY="
google_root_subject_key_id="B0:00:56:81:B8:88:62:89:62:80:E1:21:18:A1:A8:BE:09:DE:93:21"
google_root_serial_number="1"
google_root_subject_as_text="CN=Matter PAA 1,O=Google,C=US,vid=0x6006"

test_root_path="integration_tests/constants/test_root_cert"
test_root_subject="MDAxGDAWBgNVBAMMD01hdHRlciBUZXN0IFBBQTEUMBIGCisGAQQBgqJ8AgEMBDEyNUQ="
test_root_subject_key_id="E2:90:8D:36:9C:3C:A3:C1:13:BB:09:E2:4D:C1:CC:C5:A6:66:91:D4"
test_root_serial_number="1647312298631"
test_root_subject_as_text="CN=Matter Test PAA,vid=0x125D"

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
response_does_not_contain "$result" "\"subject\": \"$root_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$root_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$root_cert_subject_as_text\""
echo $result | jq

test_divider

echo "Request all NOC root certificates must be empty"
result=$(dcld query pki all-noc-x509-root-certs)
check_response "$result" "\[\]"
response_does_not_contain "$result" "\"subject\": \"$root_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$root_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$root_cert_subject_as_text\""
echo $result | jq

test_divider

echo "Request approved certificate must be empty"
result=$(dcld query pki x509-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$root_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$root_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$root_cert_subject_as_text\""
echo $result | jq

test_divider

echo "Request all certificates by subject must be empty"
result=$(dcld query pki all-subject-x509-certs --subject="$root_cert_subject")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$root_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
echo $result | jq

test_divider

echo "Request all certificates by subjectKeyId must be empty"
result=$(dcld query pki x509-cert --subject-key-id="$root_cert_subject_key_id")
check_response "$result" "Not Found"
response_does_not_contain "$result" "\"subject\": \"$root_cert_subject\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
response_does_not_contain "$result" "\"serialNumber\": \"$root_cert_serial_number\""
response_does_not_contain "$result" "\"subjectAsText\": \"$root_cert_subject_as_text\""
echo $result | jq

test_divider

echo "Add first NOC root certificate by vendor with VID = $vid"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-root-cert --certificate="$root_cert_path" --from $vendor_account --yes)
check_response "$result" "\"code\": 0"

echo "Add second NOC root certificate by vendor with VID = $vid"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-root-cert --certificate="$google_root_path" --from $vendor_account --yes)
check_response "$result" "\"code\": 0"

echo "Add NOC root certificate by vendor with VID = $vid_2"
result=$(echo "$passphrase" | dcld tx pki add-noc-x509-root-cert --certificate="$test_root_path" --from $vendor_account_2 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Request NOC root certificate by VID"
result=$(dcld query pki noc-x509-root-certs --vid="$vid")
echo $result | jq
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$root_cert_serial_number\""
check_response "$result" "\"subjectAsText\": \"$root_cert_subject_as_text\""
check_response "$result" "\"subject\": \"$google_root_subject\""
check_response "$result" "\"subjectKeyId\": \"$google_root_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$google_root_serial_number\""
check_response "$result" "\"subjectAsText\": \"$google_root_subject_as_text\""
check_response "$result" "\"vid\": $vid"

test_divider

echo "Request All NOC root certificate"
result=$(dcld query pki all-noc-x509-root-certs)
echo $result | jq
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$root_cert_serial_number\""
check_response "$result" "\"subjectAsText\": \"$root_cert_subject_as_text\""
check_response "$result" "\"subject\": \"$google_root_subject\""
check_response "$result" "\"subjectKeyId\": \"$google_root_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$google_root_serial_number\""
check_response "$result" "\"subjectAsText\": \"$google_root_subject_as_text\""
check_response "$result" "\"subject\": \"$test_root_subject\""
check_response "$result" "\"subjectKeyId\": \"$test_root_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$test_root_serial_number\""
check_response "$result" "\"subjectAsText\": \"$test_root_subject_as_text\""
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"vid\": $vid_2"

test_divider

echo "Request NOC root certificate by Subject and SubjectKeyID"
result=$(dcld query pki x509-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$root_cert_serial_number\""
check_response "$result" "\"subjectAsText\": \"$root_cert_subject_as_text\""
check_response "$result" "\"approvals\": \\[\\]"

test_divider

echo "Request NOC root certificate by Subject"
result=$(dcld query pki all-subject-x509-certs --subject="$root_cert_subject")
echo $result | jq
check_response "$result" "\"$root_cert_subject\""
check_response "$result" "\"$root_cert_subject_key_id\""

test_divider

echo "Request NOC root certificate by SubjectKeyID"
result=$(dcld query pki x509-cert --subject-key-id="$root_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$root_cert_serial_number\""
check_response "$result" "\"subjectAsText\": \"$root_cert_subject_as_text\""

test_divider