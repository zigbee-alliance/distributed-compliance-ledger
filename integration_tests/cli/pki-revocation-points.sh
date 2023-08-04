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

pai_cert_vid_path="integration_tests/constants/pai_cert_vid"
pai_cert_with_numeric_vid_path="integration_tests/constants/pai_cert_numeric_vid"
pai_cert_with_numeric_vid_pid_path="integration_tests/constants/pai_cert_numeric_vid_pid"

root_cert_path="integration_tests/constants/root_cert"
root_cert_subject="MDQxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMRAwDgYDVQQKDAdyb290LWNh"
root_cert_subject_key_id="5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB"

test_root_cert_path="integration_tests/constants/test_root_cert"
test_root_cert_subject="MDAxGDAWBgNVBAMMD01hdHRlciBUZXN0IFBBQTEUMBIGCisGAQQBgqJ8AgEMBDEyNUQ="
test_root_cert_subject_key_id="E2:90:8D:36:9C:3C:A3:C1:13:BB:09:E2:4D:C1:CC:C5:A6:66:91:D4"

trustee_account="jack"
second_trustee_account="alice"
third_trustee_account="bob"

trustee_account_address=$(echo $passphrase | dcld keys show jack -a)
second_trustee_account_address=$(echo $passphrase | dcld keys show alice -a)
third_trustee_account_address=$(echo $passphrase | dcld keys show bob -a)

label="label"
label_pai="label_pai"
vid=65521
vid_65522=65522
vid_non_vid_scoped=4701
label_non_vid_scoped="label2"
pid=8
data_url="https://url.data.dclmodel"
data_url_non_vid_scoped="https://url.data.dclmodel2"
issuer_subject_key_id="5A880E6C3653D07FB08971A3F473790930E62BDB"

vendor_account=vendor_account_$vid
echo "Create Vendor account - $vendor_account"
create_new_vendor_account $vendor_account $vid
test_divider

vendor_account_65522=vendor_account_$vid_65522
echo "Create Vendor account - $vendor_account_65522"
create_new_vendor_account $vendor_account_65522 $vid_65522
test_divider

vendor_account_non_vid_scoped=vendor_account_$vid_non_vid_scoped
echo "Create Vendor account - $vendor_account_non_vid_scoped"
create_new_vendor_account $vendor_account_non_vid_scoped $vid_non_vid_scoped
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

echo "Trustees add PAA cert with numeric vid to the ledger"
result=$(echo "$passphrase" | dcld tx pki propose-add-x509-root-cert --certificate="$paa_cert_with_numeric_vid_path" --vid $vid --from $trustee_account --yes)
check_response "$result" "\"code\": 0"
result=$(echo "$passphrase" | dcld tx pki approve-add-x509-root-cert --subject="$paa_cert_with_numeric_vid_subject" --subject-key-id="$paa_cert_with_numeric_vid_subject_key_id" --from $second_trustee_account --yes)
check_response "$result" "\"code\": 0"

echo "Trustees add PAA no VID"
result=$(echo "$passphrase" | dcld tx pki propose-add-x509-root-cert --certificate="$paa_cert_no_vid_path" --vid $vid_non_vid_scoped --from $trustee_account --yes)
check_response "$result" "\"code\": 0"
result=$(echo "$passphrase" | dcld tx pki approve-add-x509-root-cert --subject="$paa_cert_no_vid_subject" --subject-key-id="$paa_cert_no_vid_subject_key_id" --from $second_trustee_account --yes)
check_response "$result" "\"code\": 0"

echo "Trustees add root cert"
result=$(echo "$passphrase" | dcld tx pki propose-add-x509-root-cert --certificate="$root_cert_path" --vid $vid --from $trustee_account --yes)
check_response "$result" "\"code\": 0"
result=$(echo "$passphrase" | dcld tx pki approve-add-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from $second_trustee_account --yes)
check_response "$result" "\"code\": 0"

echo "Trustees add test root cert"
result=$(echo "$passphrase" | dcld tx pki propose-add-x509-root-cert --certificate="$test_root_cert_path" --vid $vid_non_vid_scoped --from $trustee_account --yes)
check_response "$result" "\"code\": 0"
result=$(echo "$passphrase" | dcld tx pki approve-add-x509-root-cert --subject="$test_root_cert_subject" --subject-key-id="$test_root_cert_subject_key_id" --from $second_trustee_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "7. ADD REVOCATION POINT FOR PAA WHEN CRL SIGNER CERTIFICATE PEM VALUE IS NOT EQUAL TO STORED CERTIFICATE PEM VALUE"

result=$(dcld tx pki add-revocation-point --vid=$vid_65522 --is-paa="true" --certificate="$paa_cert_with_numeric_vid1_path" --label="$label" --data-url="$data_url" --issuer-subject-key-id=$issuer_subject_key_id --revocation-type=1 --from=$vendor_account_65522 --yes)
response_does_not_contain "$result" "\"code\": 0"
echo $result

test_divider

echo "8. ADD REVOCATION POINT FOR PAI NOT BY VENDOR"

result=$(dcld tx pki add-revocation-point --vid=$vid --pid=$pid --is-paa="false" --certificate="$pai_cert_with_numeric_vid_pid_path" --label="$label" --data-url="$data_url" --issuer-subject-key-id=$issuer_subject_key_id --revocation-type=1 --from=$trustee_account --yes)

response_does_not_contain "$result" "\"code\": 0"
echo $result

test_divider

echo "9. ADD REVOCATION POINT FOR VID-SCOPED PAA"

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


echo "Can not add the same REVOCATION POINT second time"
result=$(dcld tx pki add-revocation-point --vid=$vid --is-paa="true" --certificate="$paa_cert_with_numeric_vid_path" --label="$label" --data-url="$data_url" --issuer-subject-key-id=$issuer_subject_key_id --revocation-type=1 --from=$vendor_account --yes)
response_does_not_contain "$result" "\"code\": 0"
echo $result

test_divider

echo "10. ADD REVOCATION POINT FOR NON-VID-SCOPED PAA"

result=$(dcld tx pki add-revocation-point --vid=$vid_non_vid_scoped --is-paa="true" --certificate="$paa_cert_no_vid_path" --label="$label_non_vid_scoped" --data-url="$data_url_non_vid_scoped" --issuer-subject-key-id=$issuer_subject_key_id --revocation-type=1 --from=$vendor_account_non_vid_scoped --yes)
check_response "$result" "\"code\": 0"
echo $result

result=$(dcld query pki all-revocation-points)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"label\": \"$label\""
check_response "$result" "\"issuerSubjectKeyID\": \"$issuer_subject_key_id\""
check_response "$result" "\"vid\": $vid_non_vid_scoped"
check_response "$result" "\"label\": \"$label_non_vid_scoped\""

result=$(dcld query pki revocation-points --issuer-subject-key-id=$issuer_subject_key_id)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"label\": \"$label\""
check_response "$result" "\"issuerSubjectKeyID\": \"$issuer_subject_key_id\""
check_response "$result" "\"vid\": $vid_non_vid_scoped"
check_response "$result" "\"label\": \"$label_non_vid_scoped\""

result=$(dcld query pki revocation-point --vid=$vid_non_vid_scoped --label=$label_non_vid_scoped --issuer-subject-key-id=$issuer_subject_key_id)
check_response "$result" "\"vid\": $vid_non_vid_scoped"
check_response "$result" "\"label\": \"$label_non_vid_scoped\""
check_response "$result" "\"issuerSubjectKeyID\": \"$issuer_subject_key_id\""
response_does_not_contain "$result" "\"label\": \"$label\""
response_does_not_contain  "$result" "\"vid\": $vid"

test_divider

echo "11. ADD REVOCATION POINT FOR PAI"

result=$(dcld tx pki add-revocation-point --vid=$vid_65522 --is-paa="false" --certificate="$pai_cert_with_numeric_vid_path" --label="$label_pai" --data-url="$data_url" --issuer-subject-key-id=$issuer_subject_key_id --revocation-type=1 --from=$vendor_account_65522 --yes)
check_response "$result" "\"code\": 0"
echo $result

result=$(dcld query pki all-revocation-points)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"label\": \"$label\""
check_response "$result" "\"vid\": $vid_non_vid_scoped"
check_response "$result" "\"label\": \"$label_non_vid_scoped\""
check_response "$result" "\"vid\": $vid_65522"
check_response "$result" "\"label\": \"$label_pai\""
check_response "$result" "\"issuerSubjectKeyID\": \"$issuer_subject_key_id\""

result=$(dcld query pki revocation-points --issuer-subject-key-id=$issuer_subject_key_id)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"label\": \"$label\""
check_response "$result" "\"vid\": $vid_non_vid_scoped"
check_response "$result" "\"label\": \"$label_non_vid_scoped\""
check_response "$result" "\"vid\": $vid_65522"
check_response "$result" "\"label\": \"$label_pai\""
check_response "$result" "\"issuerSubjectKeyID\": \"$issuer_subject_key_id\""

result=$(dcld query pki revocation-point --vid=$vid_65522 --label=$label_pai --issuer-subject-key-id=$issuer_subject_key_id)
check_response "$result" "\"vid\": $vid_65522"
check_response "$result" "\"label\": \"$label_pai\""
check_response "$result" "\"issuerSubjectKeyID\": \"$issuer_subject_key_id\""
response_does_not_contain "$result" "\"vid\": $vid"
response_does_not_contain "$result" "\"label\": \"$label\""
response_does_not_contain "$result" "\"vid\": \"$label_non_vid_scoped\""
response_does_not_contain "$result" "\"label\": \"$vid_non_vid_scoped\""

test_divider

echo "12. UPDATE REVOCATION POINT WHEN POINT NOT FOUND"

result=$(dcld tx pki update-revocation-point --vid=$vid_65522 --certificate="$pai_cert_with_numeric_vid_pid_path" --label="$label" --data-url="$data_url" --issuer-subject-key-id=$issuer_subject_key_id --from=$vendor_account_65522 --yes)
response_does_not_contain "$result" "\"code\": 0"
echo $result

test_divider

echo "13. UPDATE REVOCATION POINT FOR PAA WHEN NEW CERT IS NOT PAA"

result=$(dcld tx pki update-revocation-point --vid=$vid --certificate="$pai_cert_with_numeric_vid_pid_path" --label="$label" --data-url="$data_url" --issuer-subject-key-id=$issuer_subject_key_id --from=$vendor_account --yes)
response_does_not_contain "$result" "\"code\": 0"
echo $result

test_divider

echo "14. UPDATE REVOCATION POINT WHEN SENDER IS NOT VENDOR"

result=$(dcld tx pki update-revocation-point --vid=$vid --certificate="$paa_cert_with_numeric_vid_path" --label="$label" --data-url="$data_url" --issuer-subject-key-id=$issuer_subject_key_id --from=$trustee_account --yes)
response_does_not_contain "$result" "\"code\": 0"
echo $result

test_divider

echo "15. UPDATE REVOCATION POINT FOR PAA WHEN SENDER VID IS NOT EQUAL TO CERT VID"

result=$(dcld tx pki update-revocation-point --vid=$vid --certificate="$paa_cert_with_numeric_vid_path" --label="$label" --data-url="$data_url" --issuer-subject-key-id=$issuer_subject_key_id --from=$vendor_account_65522 --yes)
response_does_not_contain "$result" "\"code\": 0"
echo $result

test_divider

echo "16. UPDATE REVOCATION POINT FOR PAA WHEN MSG VID IS NOT EQUAL TO CERT VID"

result=$(dcld tx pki update-revocation-point --vid=$vid_65522 --certificate="$paa_cert_with_numeric_vid_path" --label="$label" --data-url="$data_url" --issuer-subject-key-id=$issuer_subject_key_id --from=$vendor_account --yes)
response_does_not_contain "$result" "\"code\": 0"
echo $result

test_divider

echo "17. UPDATE REVOCATION POINT FOR VID-SCOPED PAA"

result=$(dcld tx pki update-revocation-point --vid=$vid --certificate="$root_cert_path" --label="$label" --data-url="$data_url" --issuer-subject-key-id=$issuer_subject_key_id --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"
echo $result

result=$(dcld query pki revocation-point --vid=$vid --label=$label --issuer-subject-key-id=$issuer_subject_key_id)
check_response "$result" "\"CrlSignerCertificate\": $(<$root_cert_path)"
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"label\": \"$label\""
check_response "$result" "\"issuerSubjectKeyID\": \"$issuer_subject_key_id\""

test_divider

echo "18. UPDATE REVOCATION POINT FOR NON-VID SCOPED PAA"

result=$(dcld tx pki update-revocation-point --vid=$vid_non_vid_scoped --certificate="$test_root_cert_path" --label="$label_non_vid_scoped" --data-url="$data_url_non_vid_scoped" --issuer-subject-key-id=$issuer_subject_key_id --from=$vendor_account_non_vid_scoped --yes)
check_response "$result" "\"code\": 0"
echo $result

result=$(dcld query pki revocation-point --vid=$vid_non_vid_scoped --label=$label_non_vid_scoped --issuer-subject-key-id=$issuer_subject_key_id)
check_response "$result" "\"CrlSignerCertificate\": $(<$test_root_cert_path)"
check_response "$result" "\"vid\": $vid_non_vid_scoped"
check_response "$result" "\"label\": \"$label_non_vid_scoped\""
check_response "$result" "\"issuerSubjectKeyID\": \"$issuer_subject_key_id\""

test_divider

echo "19. UPDATE REVOCATION POINT FOR PAI"

result=$(dcld tx pki update-revocation-point --vid=$vid_65522 --certificate="$pai_cert_vid_path" --label="$label_pai" --data-url="$data_url" --issuer-subject-key-id=$issuer_subject_key_id --from=$vendor_account_65522 --yes)
check_response "$result" "\"code\": 0"
echo $result

result=$(dcld query pki revocation-point --vid=$vid_65522 --label=$label_pai --issuer-subject-key-id=$issuer_subject_key_id)
check_response "$result" "\"CrlSignerCertificate\": $(<$pai_cert_vid_path)"
check_response "$result" "\"vid\": $vid_65522"
check_response "$result" "\"label\": \"$label_pai\""
check_response "$result" "\"issuerSubjectKeyID\": \"$issuer_subject_key_id\""

test_divider

echo "20. DELETE REVOCATION PAA"

result=$(dcld tx pki delete-revocation-point --vid=$vid --label="$label" --issuer-subject-key-id=$issuer_subject_key_id --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"
echo $result

result=$(dcld query pki revocation-point --vid=$vid --label=$label --issuer-subject-key-id=$issuer_subject_key_id)
check_response "$result" "Not Found"

test_divider

echo "22. DELETE REVOCATION PAI"

result=$(dcld tx pki delete-revocation-point --vid=$vid_65522 --label="$label_pai" --issuer-subject-key-id=$issuer_subject_key_id --from=$vendor_account_65522 --yes)
check_response "$result" "\"code\": 0"
echo $result

result=$(dcld query pki revocation-point --vid=$vid_65522 --label=$label_pai --issuer-subject-key-id=$issuer_subject_key_id)
check_response "$result" "Not Found"

test_divider