set -euo pipefail
source integration_tests/cli/common.sh

root_cert_subject_path="integration_tests/constants/root_cert"
root_cert_subject="MDQxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMRAwDgYDVQQKDAdyb290LWNh"
root_cert_subject_key_id="5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB"
root_cert_vid=65521

trustee_account="jack"
second_trustee_account="alice"

vendor_account=vendor_account_$root_cert_vid
echo "Create Vendor account - $vendor_account"
create_new_vendor_account $vendor_account $root_cert_vid

test_divider

# ASSIGN VID TO ROOT CERTIFICATE THAT ALREADY HAS VID
echo "ASSIGN VID TO ROOT CERTIFICATE THAT ALREADY HAS VID"

echo "Propose and approve root certificate"
result=$(echo "$passphrase" | dcld tx pki propose-add-x509-root-cert --certificate="$root_cert_subject_path"  --vid "$root_cert_vid" --from $trustee_account --yes)
check_response "$result" "\"code\": 0"
result=$(echo "$passphrase" | dcld tx pki approve-add-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from $second_trustee_account --yes)
check_response "$result" "\"code\": 0"

result=$(dcld tx pki assign-vid --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --vid="$root_cert_vid" --from=$vendor_account --yes)
check_response "$result" "vid is not empty"

test_divider