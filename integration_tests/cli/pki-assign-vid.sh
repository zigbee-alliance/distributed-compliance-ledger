#!/bin/bash
# Copyright 2020 DSR Corporation
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -euo pipefail
source integration_tests/cli/common.sh

root_cert_subject_path="integration_tests/constants/root_cert"
root_cert_subject="MDQxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMRAwDgYDVQQKDAdyb290LWNh"
root_cert_subject_key_id="5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB"
root_cert_vid=65521

trustee_account="jack"
second_trustee_account="alice"

echo "Create a VendorAdmin Account"
create_new_account vendor_admin_account "VendorAdmin"

test_divider

# ASSIGN VID TO ROOT CERTIFICATE THAT ALREADY HAS VID
echo "ASSIGN VID TO ROOT CERTIFICATE THAT ALREADY HAS VID"

echo "Propose and approve root certificate"
result=$(echo "$passphrase" | dcld tx pki propose-add-x509-root-cert --certificate="$root_cert_subject_path"  --vid "$root_cert_vid" --from $trustee_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"
result=$(echo "$passphrase" | dcld tx pki approve-add-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from $second_trustee_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Assing VID"
result=$(dcld tx pki assign-vid --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --vid="$root_cert_vid" --from $vendor_admin_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "vid is not empty"

test_divider
