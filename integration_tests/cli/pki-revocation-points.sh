set -euo pipefail
source integration_tests/cli/common.sh

paa_cert_with_numeric_vid_path="integration_tests/constants/paa_cert_numeric_vid"
paa_cert_with_numeric_vid_subject="MDAxGDAWBgNVBAMMD01hdHRlciBUZXN0IFBBQTEUMBIGCisGAQQBgqJ8AgEMBEZGRjE="
paa_cert_with_numeric_vid_subject_key_id="6A:FD:22:77:1F:51:1F:EC:BF:16:41:97:67:10:DC:DC:31:A1:71:7E"

paa_cert_no_vid_path="integration_tests/constants/paa_cert_no_vid"
paa_cert_no_vid_subject="MBoxGDAWBgNVBAMMD01hdHRlciBUZXN0IFBBQQ=="
paa_cert_no_vid_subject_key_id="78:5C:E7:05:B8:6B:8F:4E:6F:C7:93:AA:60:CB:43:EA:69:68:82:D5"
    
paa_cert_with_numeric_vid1_path="integration_tests/constants/paa_cert_numeric_vid_1"
paa_cert_with_numeric_vid1_subject="MDAxGDAWBgNVBAMMD01hdHRlciBUZXN0IFBBQTEUMBIGCisGAQQBgqJ8AgEMBEZGRjI="
paa_cert_with_numeric_vid1_subject_key_id="7F:1D:AA:F2:44:98:B9:86:68:0E:A0:8F:C1:89:21:E8:48:48:9D:17"

pai_cert_with_numeric_vid_pid_path="integration_tests/constants/pai_cert_numeric_vid_pid"

trustee_account="jack"
second_trustee_account="alice"
third_trustee_account="bob"

trustee_account_address=$(echo $passphrase | dcld keys show jack -a)
second_trustee_account_address=$(echo $passphrase | dcld keys show alice -a)
third_trustee_account_address=$(echo $passphrase | dcld keys show bob -a)

label="label"
vid=65521
vid_65522=65522
pid=8
data_url="https://url.data.dclmodel"
issuer_subject_key_id="5A880E6C3653D07FB08971A3F473790930E62BDB"

vendor_account=vendor_account_$vid
echo "Create Vendor account - $vendor_account"
create_new_vendor_account $vendor_account $vid
test_divider

vendor_account_65522=vendor_account_$vid_65522
echo "Create Vendor account - $vendor_account_65522"
create_new_vendor_account $vendor_account_65522 $vid_65522
test_divider

echo "1. QUERY ALL REVOCATION POINTS EMPTY"

result=$(dcld query pki all-revocation-points)
check_response "$result" "\[\]"

test_divider

echo "2. QUERY REVOCATION POINT EMPTY"

result=$(dcld query pki revocation-point --vid=$vid --label=$label --issuer-subject-key-id=AB)

response_does_not_contain "$result" "\"vid\": 123"
response_does_not_contain "$result" "\"label\": \"test\""
response_does_not_contain "$result" "\"issuerSubjectKeyID\": \"AB\""
check_response "$result" "Not Found"

test_divider

echo "3. QUERY REVOCATION BY ISSUER SUBJECT KEY ID"

result=$(dcld query pki revocation-points --issuer-subject-key-id=AB)

response_does_not_contain "$result" "\"issuerSubjectKeyID\": \"AB\""
check_response "$result" "Not Found"

test_divider

echo "4. ADD REVOCATION POINT NOT BY VENDOR"

result=$(dcld tx pki add-revocation-point --vid=$vid --is-paa="true" --certificate="$paa_cert_with_numeric_vid_path" --label="$label" --data-url="$data_url" --issuer-subject-key-id=$issuer_subject_key_id --revocation-type=1 --from=$trustee_account --yes)

response_does_not_contain "$result" "\"code\": 0"
echo $result

test_divider

echo "5. ADD REVOCATION POINT FOR PAA WHEN SENDER VID IS NOT EQUAL VID FIELD"

result=$(dcld tx pki add-revocation-point --vid=$vid --is-paa="true" --certificate="$paa_cert_with_numeric_vid_path" --label="$label" --data-url="$data_url" --issuer-subject-key-id=$issuer_subject_key_id --revocation-type=1 --from=$vendor_account_65522 --yes)

response_does_not_contain "$result" "\"code\": 0"
echo $result

test_divider

echo "6. ADD REVOCATION POINT FOR PAA WHEN CERTIFICATE DOES NOT EXIST"

result=$(dcld tx pki add-revocation-point --vid=$vid --is-paa="true" --certificate="$paa_cert_with_numeric_vid_path" --label="$label" --data-url="$data_url" --issuer-subject-key-id=$issuer_subject_key_id --revocation-type=1 --from=$vendor_account --yes)

response_does_not_contain "$result" "\"code\": 0"
echo $result

test_divider

echo "7. ADD REVOCATION POINT FOR PAA WHEN CRL SIGNER CERTIFICATE PEM VALUE IS NOT EQUAL TO STORED CERTIFICATE PEM VALUE"

echo "Trustees add PAA cert with numeric vid to the ledger"
result=$(echo "$passphrase" | dcld tx pki propose-add-x509-root-cert --certificate="$paa_cert_with_numeric_vid_path" --from $trustee_account --yes)
check_response "$result" "\"code\": 0"
result=$(echo "$passphrase" | dcld tx pki approve-add-x509-root-cert --subject="$paa_cert_with_numeric_vid_subject" --subject-key-id="$paa_cert_with_numeric_vid_subject_key_id" --from $second_trustee_account --yes)
check_response "$result" "\"code\": 0"

result=$(dcld tx pki add-revocation-point --vid=$vid_65522 --is-paa="true" --certificate="$paa_cert_with_numeric_vid1_path" --label="$label" --data-url="$data_url" --issuer-subject-key-id=$issuer_subject_key_id --revocation-type=1 --from=$vendor_account_65522 --yes)

response_does_not_contain "$result" "\"code\": 0"
echo $result

test_divider

echo "8. ADD REVOCATION POINT FOR PAI NOT BY VENDOR"

result=$(dcld tx pki add-revocation-point --vid=$vid --pid=$pid --is-paa="false" --certificate="$pai_cert_with_numeric_vid_pid_path" --label="$label" --data-url="$data_url" --issuer-subject-key-id=$issuer_subject_key_id --revocation-type=1 --from=$trustee_account --yes)

response_does_not_contain "$result" "\"code\": 0"
echo $result

test_divider

# echo "9. ADD REVOCATION POINT FOR PAI WHEN CERT IS NOT CHAINED BACK TO DCL CERTS"

# result=$(dcld tx pki add-revocation-point --vid=$vid --pid=$pid --is-paa="false" --certificate="$pai_cert_with_numeric_vid_pid_path" --label="$label" --data-url="$data_url" --issuer-subject-key-id=$issuer_subject_key_id --revocation-type=1 --from=$vendor_account --yes)

# response_does_not_contain "$result" "\"code\": 0"
# echo $result

echo "10. ADD REVOCATION POINT FOR PAA WHEN REVOCATION POINT ALREADY EXISTS"

result=$(dcld tx pki add-revocation-point --vid=$vid --is-paa="true" --certificate="$paa_cert_with_numeric_vid_path" --label="$label" --data-url="$data_url" --issuer-subject-key-id=$issuer_subject_key_id --revocation-type=1 --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"
echo $result

result=$(dcld query pki all-revocation-points)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"label\": \"$label\""
check_response "$result" "\"issuerSubjectKeyID\": \"$issuer_subject_key_id\""

result=$(dcld query pki revocation-points --issuer-subject-key-id=$issuer_subject_key_id)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"label\": \"$label\""
check_response "$result" "\"issuerSubjectKeyID\": \"$issuer_subject_key_id\""

result=$(dcld query pki revocation-point --vid=$vid --label=$label --issuer-subject-key-id=$issuer_subject_key_id)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"label\": \"$label\""
check_response "$result" "\"issuerSubjectKeyID\": \"$issuer_subject_key_id\""

result=$(dcld tx pki add-revocation-point --vid=$vid --is-paa="true" --certificate="$paa_cert_with_numeric_vid_path" --label="$label" --data-url="$data_url" --issuer-subject-key-id=$issuer_subject_key_id --revocation-type=1 --from=$vendor_account --yes)
response_does_not_contain "$result" "\"code\": 0"
echo $result

echo "11. UPDATE REVOCATION POINT WHEN POINT NOT FOUND"

result=$(dcld tx pki update-revocation-point --vid=$vid_65522 --certificate="$pai_cert_with_numeric_vid_pid_path" --label="$label" --data-url="$data_url" --issuer-subject-key-id=$issuer_subject_key_id --from=$vendor_account_65522 --yes)
response_does_not_contain "$result" "\"code\": 0"
echo $result

echo "12. UPDATE REVOCATION POINT FOR PAA WHEN NEW CERT IS NOT PAA"

result=$(dcld tx pki update-revocation-point --vid=$vid --certificate="$pai_cert_with_numeric_vid_pid_path" --label="$label" --data-url="$data_url" --issuer-subject-key-id=$issuer_subject_key_id --from=$vendor_account --yes)
response_does_not_contain "$result" "\"code\": 0"
echo $result

test_divider

echo "13. UPDATE REVOCATION POINT WHEN SENDER IS NOT VENDOR"

result=$(dcld tx pki update-revocation-point --vid=$vid --certificate="$paa_cert_with_numeric_vid_path" --label="$label" --data-url="$data_url" --issuer-subject-key-id=$issuer_subject_key_id --from=$trustee_account --yes)
response_does_not_contain "$result" "\"code\": 0"
echo $result

test_divider

echo "14. UPDATE REVOCATION POINT FOR PAA WHEN SENDER VID IS NOT EQUAL TO CERT VID"

result=$(dcld tx pki update-revocation-point --vid=$vid --certificate="$paa_cert_with_numeric_vid_path" --label="$label" --data-url="$data_url" --issuer-subject-key-id=$issuer_subject_key_id --from=$vendor_account_65522 --yes)
response_does_not_contain "$result" "\"code\": 0"
echo $result

test_divider

echo "15. UPDATE REVOCATION POINT FOR PAA WHEN MSG VID IS NOT EQUAL TO CERT VID"

result=$(dcld tx pki update-revocation-point --vid=$vid_65522 --certificate="$paa_cert_with_numeric_vid_path" --label="$label" --data-url="$data_url" --issuer-subject-key-id=$issuer_subject_key_id --from=$vendor_account --yes)
response_does_not_contain "$result" "\"code\": 0"
echo $result

test_divider

echo "16. DELETE REVOCATION"

result=$(dcld tx pki delete-revocation-point --vid=$vid --label="$label" --issuer-subject-key-id=$issuer_subject_key_id --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"
echo $result

test_divider