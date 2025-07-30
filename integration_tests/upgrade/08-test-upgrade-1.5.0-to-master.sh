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

# Upgrade constants
binary_version_old="v1.5.0"

DCLD_BIN_OLD="./dcld_old"
DCLD_BIN_NEW="./dcld_new"
DCL_DIR="/var/lib/dcl/.dcl"

wget -O dcld_old "https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/$binary_version_old/dcld"
chmod ugo+x dcld_old

MASTER_UPGRADE_DOCKERFILE="./integration_tests/upgrade/Dockerfile-build-master"

cleanup_container $MASTER_UPGRADE_CONTAINER_NAME
docker build -f "$MASTER_UPGRADE_DOCKERFILE" -t "$MASTER_UPGRADE_IMAGE" .
docker container create --name "$MASTER_UPGRADE_CONTAINER_NAME" "$MASTER_UPGRADE_IMAGE"
docker cp "$MASTER_UPGRADE_CONTAINER_NAME:/go/bin/dcld" "$DCLD_BIN_NEW"
MASTER_UPGRADE_PLAN_NAME="$(docker run "$MASTER_UPGRADE_IMAGE" /bin/sh -c "cd /go/src/distributed-compliance-ledger && git rev-parse --short HEAD")"

for node_name in node0 node1 node2 node3 observer0 lightclient0; do
    if [[ -d "$LOCALNET_DIR/$node_name" ]]; then
        docker cp "$MASTER_UPGRADE_CONTAINER_NAME:/go/bin/dcld" "$LOCALNET_DIR/$node_name/"
        docker exec "$node_name" /bin/sh -c "cosmovisor add-upgrade "$MASTER_UPGRADE_PLAN_NAME" "$DCL_DIR"/dcld"
    fi
done

$DCLD_BIN_NEW config broadcast-mode sync
########################################################################################

# Upgrade to master version

get_height current_height
echo "Current height is $current_height"

plan_height=$(expr $current_height \+ 20)

test_divider

echo "Propose upgrade $MASTER_UPGRADE_PLAN_NAME at height $plan_height"
result=$(echo $passphrase | $DCLD_BIN_OLD tx dclupgrade propose-upgrade --name=$MASTER_UPGRADE_PLAN_NAME --upgrade-height=$plan_height --from $trustee_account_1 --yes)
echo "$result"
check_response "$result" "\"code\": 0"

test_divider

echo "Approve upgrade $MASTER_UPGRADE_PLAN_NAME"
result=$(echo $passphrase | $DCLD_BIN_OLD tx dclupgrade approve-upgrade --name $MASTER_UPGRADE_PLAN_NAME --from $trustee_account_2 --yes)
echo "$result"
check_response "$result" "\"code\": 0"

echo "Approve upgrade $MASTER_UPGRADE_PLAN_NAME"
result=$(echo $passphrase | $DCLD_BIN_OLD tx dclupgrade approve-upgrade --name $MASTER_UPGRADE_PLAN_NAME --from $trustee_account_3 --yes)
echo "$result"
check_response "$result" "\"code\": 0"

echo "Approve upgrade $MASTER_UPGRADE_PLAN_NAME"
result=$(echo $passphrase | $DCLD_BIN_OLD tx dclupgrade approve-upgrade --name $MASTER_UPGRADE_PLAN_NAME --from $trustee_account_4 --yes)
echo "$result"
check_response "$result" "\"code\": 0"

test_divider

echo "Wait for block height to become greater than upgrade $MASTER_UPGRADE_PLAN_NAME plan height"
wait_for_height $(expr $plan_height + 1) 300 outage-safe

test_divider

echo "Verify that no upgrade has been scheduled anymore"
result=$($DCLD_BIN_NEW query upgrade plan 2>&1) || true
check_response_and_report "$result" "no upgrade scheduled" raw

test_divider

echo "Verify that upgrade is applied"
result=$($DCLD_BIN_NEW query upgrade applied $MASTER_UPGRADE_PLAN_NAME)
echo "$result"

test_divider

########################################################################################

echo "Verify that new data is not corrupted"

test_divider

# VENDORINFO

echo "Verify if VendorInfo Record for VID: $vid_for_1_5_0 is present or not"
result=$($DCLD_BIN_NEW query vendorinfo vendor --vid=$vid_for_1_5_0)
check_response "$result" "\"vendorID\": $vid_for_1_5_0"
check_response "$result" "\"companyLegalName\": \"$company_legal_name_for_1_5_0\""

echo "Verify if VendorInfo Record for VID: $vid_for_1_2 updated or not"
result=$($DCLD_BIN_NEW query vendorinfo vendor --vid=$vid_for_1_2)
check_response "$result" "\"vendorID\": $vid_for_1_2"
check_response "$result" "\"vendorName\": \"$vendor_name_for_1_2\""
check_response "$result" "\"companyPreferredName\": \"$company_preferred_name_for_1_5_0\""
check_response "$result" "\"vendorLandingPageURL\": \"$vendor_landing_page_url_for_1_5_0\""

echo "Request all vendor infos"
result=$($DCLD_BIN_NEW query vendorinfo all-vendors)
check_response "$result" "\"vendorID\": $vid_for_1_5_0"
check_response "$result" "\"companyLegalName\": \"$company_legal_name_for_1_5_0\""
check_response "$result" "\"vendorName\": \"$vendor_name_for_1_5_0\""

test_divider

# MODEL

echo "Get Model with VID: $vid_for_1_5_0 PID: $pid_1_for_1_5_0"
result=$($DCLD_BIN_NEW query model get-model --vid=$vid_for_1_5_0 --pid=$pid_1_for_1_5_0)
check_response "$result" "\"vid\": $vid_for_1_5_0"
check_response "$result" "\"pid\": $pid_1_for_1_5_0"
check_response "$result" "\"productLabel\": \"$product_label_for_1_5_0\""

echo "Get Model with VID: $vid_for_1_5_0 PID: $pid_2_for_1_5_0"
result=$($DCLD_BIN_NEW query model get-model --vid=$vid_for_1_5_0 --pid=$pid_2_for_1_5_0)
check_response "$result" "\"vid\": $vid_for_1_5_0"
check_response "$result" "\"pid\": $pid_2_for_1_5_0"
check_response "$result" "\"productLabel\": \"$product_label_for_1_5_0\""

echo "Check Model with VID: $vid_for_1_5_0 PID: $pid_2_for_1_5_0 updated"
result=$($DCLD_BIN_NEW query model get-model --vid=$vid --pid=$pid_2)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_2"
check_response "$result" "\"productLabel\": \"$product_label_for_1_5_0\""
check_response "$result" "\"partNumber\": \"$part_number_for_1_5_0\""

echo "Check Model version with VID: $vid_for_1_5_0 PID: $pid_2_for_1_5_0 updated"
result=$($DCLD_BIN_NEW query model model-version --vid=$vid --pid=$pid_2  --softwareVersion=$software_version)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_2"
check_response "$result" "\"minApplicableSoftwareVersion\": $min_applicable_software_version_for_1_5_0"
check_response "$result" "\"maxApplicableSoftwareVersion\": $max_applicable_software_version_for_1_5_0"

echo "Get all models"
result=$($DCLD_BIN_NEW query model all-models)
check_response "$result" "\"vid\": $vid_for_1_5_0"
check_response "$result" "\"pid\": $pid_1_for_1_5_0"
check_response "$result" "\"pid\": $pid_2_for_1_5_0"

echo "Get all model versions"
result=$($DCLD_BIN_NEW query model all-model-versions --vid=$vid_for_1_5_0 --pid=$pid_1_for_1_5_0)
check_response "$result" "\"vid\": $vid_for_1_5_0"
check_response "$result" "\"pid\": $pid_1_for_1_5_0"

echo "Get Vendor Models with VID: ${vid_for_1_5_0}"
result=$($DCLD_BIN_NEW query model vendor-models --vid=$vid_for_1_5_0)
check_response "$result" "\"pid\": $pid_1_for_1_5_0"
check_response "$result" "\"pid\": $pid_2_for_1_5_0"

echo "Get model version VID: $vid_for_1_5_0 PID: $pid_1_for_1_5_0"
result=$($DCLD_BIN_NEW query model model-version --vid=$vid_for_1_5_0 --pid=$pid_1_for_1_5_0 --softwareVersion=$software_version_for_1_5_0)
check_response "$result" "\"vid\": $vid_for_1_5_0"
check_response "$result" "\"pid\": $pid_1_for_1_5_0"
check_response "$result" "\"softwareVersion\": $software_version_for_1_5_0"

echo "Get model version VID: $vid_for_1_5_0 PID: $pid_2_for_1_5_0"
result=$($DCLD_BIN_NEW query model model-version --vid=$vid_for_1_5_0 --pid=$pid_2_for_1_5_0 --softwareVersion=$software_version_for_1_5_0)
check_response "$result" "\"vid\": $vid_for_1_5_0"
check_response "$result" "\"pid\": $pid_2_for_1_5_0"
check_response "$result" "\"softwareVersion\": $software_version_for_1_5_0"

test_divider

# COMPLIANCE

echo "Get certified model vid=$vid_for_1_5_0 pid=$pid_1_for_1_5_0"
result=$($DCLD_BIN_NEW query compliance certified-model --vid=$vid_for_1_5_0 --pid=$pid_1_for_1_5_0 --softwareVersion=$software_version_for_1_5_0 --certificationType=$certification_type_for_1_5_0)
check_response "$result" "\"value\": true"
check_response "$result" "\"vid\": $vid_for_1_5_0"
check_response "$result" "\"pid\": $pid_1_for_1_5_0"
check_response "$result" "\"softwareVersion\": $software_version_for_1_5_0"
check_response "$result" "\"certificationType\": \"$certification_type_for_1_5_0\""

echo "Get revoked Model with VID: $vid_for_1_5_0 PID: $pid_2_for_1_5_0"
result=$($DCLD_BIN_NEW query compliance revoked-model --vid=$vid_for_1_5_0 --pid=$pid_2_for_1_5_0 --softwareVersion=$software_version_for_1_5_0 --certificationType=$certification_type_for_1_5_0)
check_response "$result" "\"vid\": $vid_for_1_5_0"
check_response "$result" "\"pid\": $pid_2_for_1_5_0"

echo "Get certified model with VID: $vid_for_1_5_0 PID: $pid_1_for_1_5_0"
result=$($DCLD_BIN_NEW query compliance certified-model --vid=$vid_for_1_5_0 --pid=$pid_1_for_1_5_0 --softwareVersion=$software_version_for_1_5_0 --certificationType=$certification_type_for_1_5_0)
check_response "$result" "\"value\": true"
check_response "$result" "\"vid\": $vid_for_1_5_0"
check_response "$result" "\"pid\": $pid_1_for_1_5_0"

echo "Get provisional model with VID: $vid_for_1_5_0 PID: $pid_2_for_1_5_0"
result=$($DCLD_BIN_NEW query compliance provisional-model --vid=$vid_for_1_5_0 --pid=$pid_2_for_1_5_0 --softwareVersion=$software_version_for_1_5_0 --certificationType=$certification_type_for_1_5_0)
check_response "$result" "\"value\": false"
check_response "$result" "\"vid\": $vid_for_1_5_0"
check_response "$result" "\"pid\": $pid_2_for_1_5_0"

echo "Get compliance-info model with VID: $vid_for_1_5_0 PID: $pid_1_for_1_5_0"
result=$($DCLD_BIN_NEW query compliance compliance-info --vid=$vid_for_1_5_0 --pid=$pid_1_for_1_5_0 --softwareVersion=$software_version_for_1_5_0 --certificationType=$certification_type_for_1_5_0)
check_response "$result" "\"vid\": $vid_for_1_5_0"
check_response "$result" "\"pid\": $pid_1_for_1_5_0"
check_response "$result" "\"softwareVersion\": $software_version_for_1_5_0"
check_response "$result" "\"certificationType\": \"$certification_type_for_1_5_0\""

echo "Get compliance-info model with VID: $vid_for_1_5_0 PID: $pid_2_for_1_5_0"
result=$($DCLD_BIN_NEW query compliance compliance-info --vid=$vid_for_1_5_0 --pid=$pid_2_for_1_5_0 --softwareVersion=$software_version_for_1_5_0 --certificationType=$certification_type_for_1_5_0)
check_response "$result" "\"vid\": $vid_for_1_5_0"
check_response "$result" "\"pid\": $pid_2_for_1_5_0"
check_response "$result" "\"softwareVersion\": $software_version_for_1_5_0"
check_response "$result" "\"certificationType\": \"$certification_type_for_1_5_0\""

echo "Get device software compliance cDCertificateId=$cd_certificate_id_for_1_5_0"
result=$($DCLD_BIN_NEW query compliance device-software-compliance --cdCertificateId=$cd_certificate_id_for_1_5_0)
check_response "$result" "\"vid\": $vid_for_1_5_0"
check_response "$result" "\"pid\": $pid_1_for_1_5_0"

echo "Get all certified models"
result=$($DCLD_BIN_NEW query compliance all-certified-models)
check_response "$result" "\"vid\": $vid_for_1_5_0"
check_response "$result" "\"pid\": $pid_1_for_1_5_0"

echo "Get all provisional models"
result=$($DCLD_BIN_NEW query compliance all-provisional-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_3"

echo "Get all revoked models"
result=$($DCLD_BIN_NEW query compliance all-revoked-models)
check_response "$result" "\"vid\": $vid_for_1_5_0"
check_response "$result" "\"pid\": $pid_2_for_1_5_0"

echo "Get all compliance infos"
result=$($DCLD_BIN_NEW query compliance all-compliance-info)
check_response "$result" "\"vid\": $vid_for_1_5_0"
check_response "$result" "\"pid\": $pid_1_for_1_5_0"
check_response "$result" "\"pid\": $pid_2_for_1_5_0"

echo "Get all device software compliances"
result=$($DCLD_BIN_NEW query compliance all-device-software-compliance)
check_response "$result" "\"vid\": $vid_for_1_5_0"
check_response "$result" "\"pid\": $pid_1_for_1_5_0"
check_response "$result" "\"cDCertificateId\": \"$cd_certificate_id_for_1_5_0\""

test_divider

# PKI

echo "Get certificates"

echo "Get certificates (ALL)"
result=$($DCLD_BIN_NEW query pki all-certs)
echo $result | jq
check_response "$result" "\"subjectKeyId\": \"$da_root_cert_2_subject_key_id_for_1_5_0\""
check_response "$result" "\"subjectKeyId\": \"$da_intermediate_cert_2_subject_key_id_for_1_5_0\""
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_2_subject_key_id_for_1_5_0\""
check_response "$result" "\"subjectKeyId\": \"$noc_ica_cert_2_subject_key_id_for_1_5_0\""

echo "Get certificates (DA)"
result=$($DCLD_BIN_NEW query pki all-x509-certs)
check_response "$result" "\"subjectKeyId\": \"$da_root_cert_2_subject_key_id_for_1_5_0\""
check_response "$result" "\"subjectKeyId\": \"$da_intermediate_cert_2_subject_key_id_for_1_5_0\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_2_subject_key_id_for_1_5_0\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_ica_cert_2_subject_key_id_for_1_5_0\""

echo "Get certificates (NOC)"
result=$($DCLD_BIN_NEW query pki all-noc-x509-certs)
check_response "$result" "$noc_root_cert_2_subject_key_id_for_1_5_0"
check_response "$result" "$noc_ica_cert_2_subject_key_id_for_1_5_0"
response_does_not_contain "$result" "$da_root_cert_1_subject_key_id_for_1_5_0"
response_does_not_contain "$result" "$noc_root_cert_1_subject_key_id_for_1_5_0"
response_does_not_contain "$result" "$noc_ica_cert_1_subject_key_id_for_1_5_0"

echo "Get certificate"

echo "Get certificate (ALL)"
result=$($DCLD_BIN_NEW query pki cert --subject=$da_root_cert_2_subject_for_1_5_0 --subject-key-id=$da_root_cert_2_subject_key_id_for_1_5_0)
check_response "$result" "\"subject\": \"$da_root_cert_2_subject_for_1_5_0\""
check_response "$result" "\"subjectKeyId\": \"$da_root_cert_2_subject_key_id_for_1_5_0\""

result=$($DCLD_BIN_NEW query pki cert --subject=$da_intermediate_cert_2_subject_for_1_5_0 --subject-key-id=$da_intermediate_cert_2_subject_key_id_for_1_5_0)
check_response "$result" "\"subject\": \"$da_intermediate_cert_2_subject_for_1_5_0\""
check_response "$result" "\"subjectKeyId\": \"$da_intermediate_cert_2_subject_key_id_for_1_5_0\""

result=$($DCLD_BIN_NEW query pki cert --subject=$noc_root_cert_2_subject_for_1_5_0 --subject-key-id=$noc_root_cert_2_subject_key_id_for_1_5_0)
check_response "$result" "\"subject\": \"$noc_root_cert_2_subject_for_1_5_0\""
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_2_subject_key_id_for_1_5_0\""

result=$($DCLD_BIN_NEW query pki cert --subject=$noc_ica_cert_2_subject_for_1_5_0 --subject-key-id=$noc_ica_cert_2_subject_key_id_for_1_5_0)
check_response "$result" "\"subject\": \"$noc_ica_cert_2_subject_for_1_5_0\""
check_response "$result" "\"subjectKeyId\": \"$noc_ica_cert_2_subject_key_id_for_1_5_0\""

echo "Get certificate (DA)"
result=$($DCLD_BIN_NEW query pki x509-cert --subject=$da_root_cert_2_subject_for_1_5_0 --subject-key-id=$da_root_cert_2_subject_key_id_for_1_5_0)
check_response "$result" "\"subject\": \"$da_root_cert_2_subject_for_1_5_0\""
check_response "$result" "\"subjectKeyId\": \"$da_root_cert_2_subject_key_id_for_1_5_0\""

result=$($DCLD_BIN_NEW query pki x509-cert --subject=$da_intermediate_cert_2_subject_for_1_5_0 --subject-key-id=$da_intermediate_cert_2_subject_key_id_for_1_5_0)
check_response "$result" "\"subject\": \"$da_intermediate_cert_2_subject_for_1_5_0\""
check_response "$result" "\"subjectKeyId\": \"$da_intermediate_cert_2_subject_key_id_for_1_5_0\""

result=$($DCLD_BIN_NEW query pki x509-cert --subject=$noc_root_cert_2_subject_for_1_5_0 --subject-key-id=$noc_root_cert_2_subject_key_id_for_1_5_0)
check_response "$result" "Not Found"

result=$($DCLD_BIN_NEW query pki x509-cert --subject=$noc_ica_cert_2_subject_for_1_5_0 --subject-key-id=$noc_ica_cert_2_subject_key_id_for_1_5_0)
check_response "$result" "Not Found"

echo "Get certificate (NOC)"
result=$($DCLD_BIN_NEW query pki noc-x509-cert --subject=$noc_root_cert_2_subject_for_1_5_0 --subject-key-id=$noc_root_cert_2_subject_key_id_for_1_5_0)
check_response "$result" "\"subject\": \"$noc_root_cert_2_subject_for_1_5_0\""
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_2_subject_key_id_for_1_5_0\""

result=$($DCLD_BIN_NEW query pki noc-x509-cert --subject=$noc_ica_cert_2_subject_for_1_5_0 --subject-key-id=$noc_ica_cert_2_subject_key_id_for_1_5_0)
check_response "$result" "\"subject\": \"$noc_ica_cert_2_subject_for_1_5_0\""
check_response "$result" "\"subjectKeyId\": \"$noc_ica_cert_2_subject_key_id_for_1_5_0\""

result=$($DCLD_BIN_NEW query pki noc-x509-cert --subject=$da_root_cert_2_subject_for_1_5_0 --subject-key-id=$da_root_cert_2_subject_key_id_for_1_5_0)
check_response "$result" "Not Found"

result=$($DCLD_BIN_NEW query pki noc-x509-cert --subject=$da_intermediate_cert_2_subject_for_1_5_0 --subject-key-id=$da_intermediate_cert_2_subject_key_id_for_1_5_0)
check_response "$result" "Not Found"

echo "Get all subject certificates"

echo "Get all subject certificates (Global)"
result=$($DCLD_BIN_NEW query pki all-subject-certs --subject=$da_root_cert_2_subject_for_1_5_0)
check_response "$result" "$da_root_cert_2_subject_key_id_for_1_5_0"

result=$($DCLD_BIN_NEW query pki all-subject-certs --subject=$noc_root_cert_2_subject_for_1_5_0)
check_response "$result" "$noc_root_cert_2_subject_for_1_5_0"

echo "Get all subject certificates (DA)"
result=$($DCLD_BIN_NEW query pki all-subject-x509-certs --subject=$da_root_cert_2_subject_for_1_5_0)
check_response "$result" "$da_root_cert_2_subject_key_id_for_1_5_0"

result=$($DCLD_BIN_NEW query pki all-subject-x509-certs --subject=$noc_root_cert_2_subject_for_1_5_0)
check_response "$result" "Not Found"

echo "Get all subject certificates (NOC)"
result=$($DCLD_BIN_NEW query pki all-noc-subject-x509-certs --subject=$noc_root_cert_2_subject_for_1_5_0)
check_response "$result" "$noc_root_cert_2_subject_for_1_5_0"

result=$($DCLD_BIN_NEW query pki all-noc-subject-x509-certs --subject=$da_root_cert_2_subject_for_1_5_0)
check_response "$result" "Not Found"

echo "Get all certificates by SKID"

echo "Get all certificates by SKID (Global)"
result=$($DCLD_BIN_NEW query pki cert --subject-key-id=$da_root_cert_2_subject_key_id_for_1_5_0)
check_response "$result" "\"subjectKeyId\": \"$da_root_cert_2_subject_key_id_for_1_5_0\""

result=$($DCLD_BIN_NEW query pki cert --subject-key-id=$noc_root_cert_2_subject_key_id_for_1_5_0)
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_2_subject_key_id_for_1_5_0\""

echo "Get all certificates by SKID (DA)"
result=$($DCLD_BIN_NEW query pki x509-cert --subject-key-id=$da_root_cert_2_subject_key_id_for_1_5_0)
check_response "$result" "\"subjectKeyId\": \"$da_root_cert_2_subject_key_id_for_1_5_0\""

result=$($DCLD_BIN_NEW query pki x509-cert --subject-key-id=$noc_root_cert_2_subject_key_id_for_1_5_0)
check_response "$result" "Not Found"

echo "Get all certificates by SKID (NOC)"
result=$($DCLD_BIN_NEW query pki noc-x509-cert --subject-key-id=$noc_root_cert_2_subject_key_id_for_1_5_0)
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_2_subject_key_id_for_1_5_0\""

result=$($DCLD_BIN_NEW query pki noc-x509-cert --subject-key-id=$da_root_cert_2_subject_key_id_for_1_5_0)
check_response "$result" "Not Found"

echo "Get all revoked x509 root certificates"

echo "Get all revoked x509 certificates (DA)"
result=$($DCLD_BIN_NEW query pki all-revoked-x509-certs)
check_response "$result" "\"subjectKeyId\": \"$da_root_cert_1_subject_key_id_for_1_5_0\""
check_response "$result" "\"subjectKeyId\": \"$da_intermediate_cert_1_subject_key_id_for_1_5_0\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id_for_1_5_0\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_ica_cert_1_subject_key_id_for_1_5_0\""

echo "Get all revoked x509 root certificates (DA)"
result=$($DCLD_BIN_NEW query pki all-revoked-x509-root-certs)
check_response "$result" "\"subjectKeyId\": \"$da_root_cert_1_subject_key_id_for_1_5_0\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$da_root_cert_2_subject_for_1_5_0\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id_for_1_5_0\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id_for_1_5_0\""

echo "Get all revoked x509 root certificates (NOC)"
result=$($DCLD_BIN_NEW query pki all-revoked-noc-x509-root-certs)
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id_for_1_5_0\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$da_root_cert_1_subject_key_id_for_1_5_0\""
response_does_not_contain "$result" "\"subjectKeyId\": \"$noc_ica_cert_1_subject_key_id_for_1_5_0\""

echo "Get all revoked x509 ica certificates (NOC)"
result=$($DCLD_BIN_NEW query pki all-revoked-noc-x509-ica-certs)
check_response "$result" "\"subjectKeyId\": \"$noc_ica_cert_1_subject_key_id_for_1_5_0\""

echo "Get revoked x509 certificate"

echo "Get revoked x509 certificate (DA)"
result=$($DCLD_BIN_NEW query pki revoked-x509-cert --subject=$da_root_cert_1_subject_for_1_5_0 --subject-key-id=$da_root_cert_1_subject_key_id_for_1_5_0)
check_response "$result" "\"subject\": \"$da_root_cert_1_subject_for_1_5_0\""
check_response "$result" "\"subjectKeyId\": \"$da_root_cert_1_subject_key_id_for_1_5_0\""

result=$($DCLD_BIN_NEW query pki revoked-x509-cert --subject=$noc_root_cert_1_subject_for_1_5_0 --subject-key-id=$noc_root_cert_1_subject_key_id_for_1_5_0)
check_response "$result" "Not Found"

echo "Get revoked x509 certificate (NOC)"
result=$($DCLD_BIN_NEW query pki revoked-noc-x509-root-cert --subject=$noc_root_cert_1_subject_for_1_5_0 --subject-key-id=$noc_root_cert_1_subject_key_id_for_1_5_0)
check_response "$result" "\"subject\": \"$noc_root_cert_1_subject_for_1_5_0\""
check_response "$result" "\"subjectKeyId\": \"$noc_root_cert_1_subject_key_id_for_1_5_0\""

result=$($DCLD_BIN_NEW query pki revoked-noc-x509-root-cert --subject=$da_root_cert_1_subject_for_1_5_0 --subject-key-id=$da_root_cert_1_subject_key_id_for_1_5_0)
check_response "$result" "Not Found"

echo "Get revocation point"
result=$($DCLD_BIN_NEW query pki revocation-point --vid=$vid_for_1_5_0 --label=$product_label_for_1_5_0 --issuer-subject-key-id=$issuer_subject_key_id)
check_response "$result" "\"vid\": $vid_for_1_5_0"
check_response "$result" "\"issuerSubjectKeyID\": \"$issuer_subject_key_id\""
check_response "$result" "\"label\": \"$product_label_for_1_5_0\""
check_response "$result" "\"dataURL\": \"$test_data_url_for_1_5_0\""

echo "Get revocation points by issuer subject key id"
result=$($DCLD_BIN_NEW query pki revocation-points --issuer-subject-key-id=$issuer_subject_key_id)
check_response "$result" "\"vid\": $vid_for_1_5_0"
check_response "$result" "\"issuerSubjectKeyID\": \"$issuer_subject_key_id\""
check_response "$result" "\"label\": \"$product_label_for_1_5_0\""
check_response "$result" "\"dataURL\": \"$test_data_url_for_1_5_0\""

echo "Get all revocation points"
result=$($DCLD_BIN_NEW query pki all-revocation-points)
check_response "$result" "\"vid\": $vid_for_1_5_0"
check_response "$result" "\"issuerSubjectKeyID\": \"$issuer_subject_key_id\""
check_response "$result" "\"label\": \"$product_label_for_1_5_0\""
check_response "$result" "\"dataURL\": \"$test_data_url_for_1_5_0\""

echo "Get all noc x509 root certificates by vid=$vid_for_1_5_0 and skid=$noc_root_cert_2_subject_key_id_for_1_5_0"
result=$($DCLD_BIN_NEW query pki noc-x509-cert --vid=$vid_for_1_5_0 --subject-key-id=$noc_root_cert_2_subject_key_id_for_1_5_0)
check_response "$result" "\"subject\": \"$noc_root_cert_2_subject_for_1_5_0\""
check_response "$result" "$noc_root_cert_2_subject_key_id_for_1_5_0"

echo "Get all noc x509 root certificates by vid $vid_for_1_5_0 and skid=$noc_ica_cert_2_subject_key_id_for_1_5_0"
result=$($DCLD_BIN_NEW query pki noc-x509-cert --vid=$vid_for_1_5_0 --subject-key-id=$noc_ica_cert_2_subject_key_id_for_1_5_0)
check_response "$result" "\"subject\": \"$noc_ica_cert_2_subject_for_1_5_0\""
check_response "$result" "$noc_ica_cert_2_subject_key_id_for_1_5_0"

test_divider

# AUTH

echo "Get all accounts"
result=$($DCLD_BIN_NEW query auth all-accounts)
check_response "$result" "\"address\": \"$user_11_address\""

echo "Get account"
result=$($DCLD_BIN_NEW query auth account --address=$user_11_address)
check_response "$result" "\"address\": \"$user_11_address\""

echo "Get all proposed accounts"
result=$($DCLD_BIN_NEW query auth all-proposed-accounts)
check_response "$result" "\"address\": \"$user_12_address\""

echo "Get proposed account"
result=$($DCLD_BIN_NEW query auth proposed-account --address=$user_12_address)
check_response "$result" "\"address\": \"$user_12_address\""

echo "Get all proposed accounts to revoke"
result=$($DCLD_BIN_NEW query auth all-proposed-accounts-to-revoke)
check_response "$result" "\"address\": \"$user_11_address\""

echo "Get proposed account to revoke"
result=$($DCLD_BIN_NEW query auth proposed-account-to-revoke --address=$user_11_address)
check_response "$result" "\"address\": \"$user_11_address\""

echo "Get all revoked accounts"
result=$($DCLD_BIN_NEW query auth all-revoked-accounts)
check_response "$result" "\"address\": \"$user_10_address\""

echo "Get revoked account"
result=$($DCLD_BIN_NEW query auth revoked-account --address=$user_10_address)
check_response "$result" "\"address\": \"$user_10_address\""

test_divider

# Validator

echo "Get node"
# FIXME: use proper binary (not dcld but $DCLD_BIN_OLD)
result=$(docker exec "$VALIDATOR_DEMO_CONTAINER_NAME" /bin/sh -c "echo test1234 | dcld query validator all-nodes")
check_response "$result" "\"owner\": \"$validator_address\""

########################################################################################

# after upgrade constants

vid_for_master=62529
pid_1_for_master=89
pid_2_for_master=99
pid_3_for_master=77
device_type_id_for_master=3433
product_name_for_master="ProductName_master"
product_label_for_master="ProductLabel_master"
part_number_for_master="ZCU2245M"
software_version_for_master=4
software_version_string_for_master="5.3"
cd_version_number_for_master=743
min_applicable_software_version_for_master=4
max_applicable_software_version_for_master=4000

certification_type_for_master="matter"
certification_date_for_master="2024-02-01T00:00:00Z"
provisional_date_for_master="2016-10-12T00:00:00Z"
cd_certificate_id_for_master="10DEXZ"

test_data_url_for_master="https://url.data.dclmodel-master"

vendor_name_for_master="Vendor_master"
company_legal_name_for_master="LegalCompanyName_master"
company_preferred_name_for_master="CompanyPreferredName_master"
vendor_landing_page_url_for_master="https://www.new_master_example.com"

vendor_account_for_master="vendor_account_master"

echo "Create Vendor account $vendor_account_for_master"

result="$(echo $passphrase | $DCLD_BIN_NEW keys add "$vendor_account_for_master")"
echo "keys add $result"
_address=$(echo $passphrase | $DCLD_BIN_NEW keys show $vendor_account_for_master -a)
_pubkey=$(echo $passphrase | $DCLD_BIN_NEW keys show $vendor_account_for_master -p)
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth propose-add-account --address="$_address" --pubkey="$_pubkey" --vid="$vid_for_master" --roles="Vendor" --from "$trustee_account_1" --yes)
echo "propose-add-account $result"
result=$(get_txn_result "$result")
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth approve-add-account --address="$_address" --from "$trustee_account_2" --yes)
result=$(get_txn_result "$result")
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth approve-add-account --address="$_address" --from "$trustee_account_3" --yes)
result=$(get_txn_result "$result")
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth approve-add-account --address="$_address" --from "$trustee_account_4" --yes)
result=$(get_txn_result "$result")

random_string user_13
echo "$user_13 generates keys"
cmd="(echo $passphrase; echo $passphrase) | $DCLD_BIN_NEW keys add $user_13"
result="$(bash -c "$cmd")"
user_13_address=$(echo $passphrase | $DCLD_BIN_NEW keys show $user_13 -a)
user_13_pubkey=$(echo $passphrase | $DCLD_BIN_NEW keys show $user_13 -p)

random_string user_14
echo "$user_14 generates keys"
cmd="(echo $passphrase; echo $passphrase) | $DCLD_BIN_NEW keys add $user_14"
result="$(bash -c "$cmd")"
user_14_address=$(echo $passphrase | $DCLD_BIN_NEW keys show $user_14 -a)
user_14_pubkey=$(echo $passphrase | $DCLD_BIN_NEW keys show $user_14 -p)

random_string user_15
echo "$user_15 generates keys"
cmd="(echo $passphrase; echo $passphrase) | $DCLD_BIN_NEW keys add $user_15"
result="$(bash -c "$cmd")"
user_15_address=$(echo $passphrase | $DCLD_BIN_NEW keys show $user_15 -a)
user_15_pubkey=$(echo $passphrase | $DCLD_BIN_NEW keys show $user_15 -p)

# send all ledger update transactions after upgrade

# VENDOR_INFO
echo "Add vendor $vendor_name_for_master"
result=$(echo $passphrase | $DCLD_BIN_NEW tx vendorinfo add-vendor --vid=$vid_for_master --vendorName=$vendor_name_for_master --companyLegalName=$company_legal_name_for_master --companyPreferredName=$company_preferred_name_for_master --vendorLandingPageURL=$vendor_landing_page_url_for_master --from=$vendor_account_for_master --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Update vendor $vendor_name_for_1_2"
result=$(echo $passphrase | $DCLD_BIN_NEW tx vendorinfo update-vendor --vid=$vid_for_1_2 --vendorName=$vendor_name_for_1_2 --companyLegalName=$company_legal_name_for_1_2 --companyPreferredName=$company_preferred_name_for_master --vendorLandingPageURL=$vendor_landing_page_url_for_master --from=$vendor_account_for_1_2 --yes)
result=$(get_txn_result "$result")
echo $result
check_response "$result" "\"code\": 0"

test_divider

# MODEL and MODEL_VERSION

echo "Add model vid=$vid_for_master pid=$pid_1_for_master"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model --vid=$vid_for_master --pid=$pid_1_for_master --deviceTypeID=$device_type_id_for_master --productName=$product_name_for_master --productLabel=$product_label_for_master --partNumber=$part_number_for_master --from=$vendor_account_for_master --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Add model version vid=$vid_for_master pid=$pid_1_for_master"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model-version --vid=$vid_for_master --pid=$pid_1_for_master --softwareVersion=$software_version_for_master --softwareVersionString=$software_version_string_for_master --cdVersionNumber=$cd_version_number_for_master --minApplicableSoftwareVersion=$min_applicable_software_version_for_master --maxApplicableSoftwareVersion=$max_applicable_software_version_for_master --from=$vendor_account_for_master --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Add model vid=$vid_for_master pid=$pid_2_for_master"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model --vid=$vid_for_master --pid=$pid_2_for_master --deviceTypeID=$device_type_id_for_master --productName=$product_name_for_master --productLabel=$product_label_for_master --partNumber=$part_number_for_master --from=$vendor_account_for_master --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Add model version vid=$vid_for_master pid=$pid_2_for_master"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model-version --vid=$vid_for_master --pid=$pid_2_for_master --softwareVersion=$software_version_for_master --softwareVersionString=$software_version_string_for_master --cdVersionNumber=$cd_version_number_for_master --minApplicableSoftwareVersion=$min_applicable_software_version_for_master --maxApplicableSoftwareVersion=$max_applicable_software_version_for_master --from=$vendor_account_for_master --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Add model vid=$vid_for_master pid=$pid_3_for_master"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model --vid=$vid_for_master --pid=$pid_3_for_master --deviceTypeID=$device_type_id_for_master --productName=$product_name_for_master --productLabel=$product_label_for_master --partNumber=$part_number_for_master --from=$vendor_account_for_master --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Add model version vid=$vid_for_master pid=$pid_3_for_master"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model-version --vid=$vid_for_master --pid=$pid_3_for_master --softwareVersion=$software_version_for_master --softwareVersionString=$software_version_string_for_master --cdVersionNumber=$cd_version_number_for_master --minApplicableSoftwareVersion=$min_applicable_software_version_for_master --maxApplicableSoftwareVersion=$max_applicable_software_version_for_master --from=$vendor_account_for_master --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Delete model vid=$vid_for_master pid=$pid_3_for_master"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model delete-model --vid=$vid_for_master --pid=$pid_3_for_master --from=$vendor_account_for_master --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Update model vid=$vid pid=$pid_2"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model update-model --vid=$vid --pid=$pid_2 --productName=$product_name --productLabel=$product_label_for_master --partNumber=$part_number_for_master --from=$vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Update model version vid=$vid pid=$pid_2"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model update-model-version --vid=$vid --pid=$pid_2 --softwareVersion=$software_version --minApplicableSoftwareVersion=$min_applicable_software_version_for_master --maxApplicableSoftwareVersion=$max_applicable_software_version_for_master --from=$vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

# CERTIFY_DEVICE_COMPLIANCE

echo "Certify model vid=$vid_for_master pid=$pid_1_for_master"
result=$(echo $passphrase | $DCLD_BIN_NEW tx compliance certify-model --vid=$vid_for_master --pid=$pid_1_for_master --softwareVersion=$software_version_for_master --softwareVersionString=$software_version_string_for_master  --certificationType=$certification_type_for_master --certificationDate=$certification_date_for_master --cdCertificateId=$cd_certificate_id_for_master --from=$certification_center_account --cdVersionNumber=$cd_version_number_for_master --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Provision model vid=$vid_for_master pid=$pid_2_for_master"
result=$(echo $passphrase | $DCLD_BIN_NEW tx compliance provision-model --vid=$vid_for_master --pid=$pid_2_for_master --softwareVersion=$software_version_for_master --softwareVersionString=$software_version_string_for_master --certificationType=$certification_type_for_master --provisionalDate=$provisional_date_for_master --cdCertificateId=$cd_certificate_id_for_master --from=$certification_center_account --cdVersionNumber=$cd_version_number_for_master --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Certify model vid=$vid_for_master pid=$pid_2_for_master"
result=$(echo $passphrase | $DCLD_BIN_NEW tx compliance certify-model --vid=$vid_for_master --pid=$pid_2_for_master --softwareVersion=$software_version_for_master --softwareVersionString=$software_version_string_for_master  --certificationType=$certification_type_for_master --certificationDate=$certification_date_for_master --cdCertificateId=$cd_certificate_id_for_master --from=$certification_center_account --cdVersionNumber=$cd_version_number_for_master  --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Revoke model certification vid=$vid_for_master pid=$pid_2_for_master"
result=$(echo $passphrase | $DCLD_BIN_NEW tx compliance revoke-model --vid=$vid_for_master --pid=$pid_2_for_master --softwareVersion=$software_version_for_master --softwareVersionString=$software_version_string_for_master --certificationType=$certification_type_for_master --revocationDate=$certification_date_for_master --from=$certification_center_account --cdVersionNumber=$cd_version_number_for_master --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

# AUTH

echo "Propose add account $user_13_address"
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth propose-add-account --address="$user_13_address" --pubkey="$user_13_pubkey" --roles="CertificationCenter" --from="$trustee_account_1" --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_13_address"
result=$($DCLD_BIN_NEW tx auth approve-add-account --address="$user_13_address" --from=$trustee_account_2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_13_address"
result=$($DCLD_BIN_NEW tx auth approve-add-account --address="$user_13_address" --from=$trustee_account_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_13_address"
result=$($DCLD_BIN_NEW tx auth approve-add-account --address="$user_13_address" --from=$trustee_account_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Propose add account $user_14_address"
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth propose-add-account --address="$user_14_address" --pubkey=$user_14_pubkey --roles=CertificationCenter --from=$trustee_account_1 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_14_address"
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth approve-add-account --address="$user_14_address" --from=$trustee_account_2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_14_address"
result=$($DCLD_BIN_NEW tx auth approve-add-account --address="$user_14_address" --from=$trustee_account_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_14_address"
result=$($DCLD_BIN_NEW tx auth approve-add-account --address="$user_14_address" --from=$trustee_account_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Propose add account $user_15_address"
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth propose-add-account --address="$user_15_address" --pubkey=$user_15_pubkey --roles=CertificationCenter --from=$trustee_account_1 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Propose revoke account $user_13_address"
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth propose-revoke-account --address="$user_13_address" --from=$trustee_account_1 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve revoke account $user_13_address"
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth approve-revoke-account --address="$user_13_address" --from=$trustee_account_2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve revoke account $user_13_address"
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth approve-revoke-account --address="$user_13_address" --from=$trustee_account_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve revoke account $user_13_address"
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth approve-revoke-account --address="$user_13_address" --from=$trustee_account_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Propose revoke account $user_14_address"
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth propose-revoke-account --address="$user_14_address" --from=$trustee_account_1 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

# VALIDATOR_NODE
result=$(docker exec "$VALIDATOR_DEMO_CONTAINER_NAME" /bin/sh -c "echo test1234  | dcld config broadcast-mode sync")

echo "Disable node"
result=$(docker exec "$VALIDATOR_DEMO_CONTAINER_NAME" /bin/sh -c "echo test1234  | dcld tx validator disable-node --from=$account --yes")
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Enable node"
result=$(docker exec "$VALIDATOR_DEMO_CONTAINER_NAME" /bin/sh -c "echo test1234  | dcld tx validator enable-node --from=$account --yes")
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve disable node"
result=$(echo $passphrase | $DCLD_BIN_NEW tx validator approve-disable-node --address=$validator_address --from=$trustee_account_2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve disable node"
result=$(echo $passphrase | $DCLD_BIN_NEW tx validator approve-disable-node --address=$validator_address --from=$trustee_account_3 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve disable node"
result=$(echo $passphrase | $DCLD_BIN_NEW tx validator approve-disable-node --address=$validator_address --from=$trustee_account_4 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Enable node"
result=$(docker exec "$VALIDATOR_DEMO_CONTAINER_NAME" /bin/sh -c "echo test1234  | dcld tx validator enable-node --from=$account --yes")
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Propose disable node"
result=$(echo $passphrase | $DCLD_BIN_NEW tx validator propose-disable-node --address=$validator_address --from=$trustee_account_1 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Verify that new data is not corrupted"

test_divider

# VENDORINFO

echo "Verify if VendorInfo Record for VID: $vid_for_master is present or not"
result=$($DCLD_BIN_NEW query vendorinfo vendor --vid=$vid_for_master)
check_response "$result" "\"vendorID\": $vid_for_master"
check_response "$result" "\"companyLegalName\": \"$company_legal_name_for_master\""

echo "Verify if VendorInfo Record for VID: $vid_for_1_2 updated or not"
result=$($DCLD_BIN_NEW query vendorinfo vendor --vid=$vid_for_1_2)
check_response "$result" "\"vendorID\": $vid_for_1_2"
check_response "$result" "\"vendorName\": \"$vendor_name_for_1_2\""
check_response "$result" "\"companyPreferredName\": \"$company_preferred_name_for_master\""
check_response "$result" "\"vendorLandingPageURL\": \"$vendor_landing_page_url_for_master\""

echo "Request all vendor infos"
result=$($DCLD_BIN_NEW query vendorinfo all-vendors)
check_response "$result" "\"vendorID\": $vid_for_master"
check_response "$result" "\"companyLegalName\": \"$company_legal_name_for_master\""
check_response "$result" "\"vendorName\": \"$vendor_name_for_master\""

test_divider

# MODEL

echo "Get Model with VID: $vid_for_master PID: $pid_1_for_master"
result=$($DCLD_BIN_NEW query model get-model --vid=$vid_for_master --pid=$pid_1_for_master)
check_response "$result" "\"vid\": $vid_for_master"
check_response "$result" "\"pid\": $pid_1_for_master"
check_response "$result" "\"productLabel\": \"$product_label_for_master\""

echo "Get Model with VID: $vid_for_master PID: $pid_2_for_master"
result=$($DCLD_BIN_NEW query model get-model --vid=$vid_for_master --pid=$pid_2_for_master)
check_response "$result" "\"vid\": $vid_for_master"
check_response "$result" "\"pid\": $pid_2_for_master"
check_response "$result" "\"productLabel\": \"$product_label_for_master\""

echo "Check Model with VID: $vid_for_master PID: $pid_2_for_master updated"
result=$($DCLD_BIN_NEW query model get-model --vid=$vid --pid=$pid_2)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_2"
check_response "$result" "\"productLabel\": \"$product_label_for_master\""
check_response "$result" "\"partNumber\": \"$part_number_for_master\""

echo "Check Model version with VID: $vid_for_master PID: $pid_2_for_master updated"
result=$($DCLD_BIN_NEW query model model-version --vid=$vid --pid=$pid_2  --softwareVersion=$software_version)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_2"
check_response "$result" "\"minApplicableSoftwareVersion\": $min_applicable_software_version_for_master"
check_response "$result" "\"maxApplicableSoftwareVersion\": $max_applicable_software_version_for_master"

echo "Get all models"
result=$($DCLD_BIN_NEW query model all-models)
check_response "$result" "\"vid\": $vid_for_master"
check_response "$result" "\"pid\": $pid_1_for_master"
check_response "$result" "\"pid\": $pid_2_for_master"

echo "Get all model versions"
result=$($DCLD_BIN_NEW query model all-model-versions --vid=$vid_for_master --pid=$pid_1_for_master)
check_response "$result" "\"vid\": $vid_for_master"
check_response "$result" "\"pid\": $pid_1_for_master"

echo "Get Vendor Models with VID: ${vid_for_master}"
result=$($DCLD_BIN_NEW query model vendor-models --vid=$vid_for_master)
check_response "$result" "\"pid\": $pid_1_for_master"
check_response "$result" "\"pid\": $pid_2_for_master"

echo "Get model version VID: $vid_for_master PID: $pid_1_for_master"
result=$($DCLD_BIN_NEW query model model-version --vid=$vid_for_master --pid=$pid_1_for_master --softwareVersion=$software_version_for_master)
check_response "$result" "\"vid\": $vid_for_master"
check_response "$result" "\"pid\": $pid_1_for_master"
check_response "$result" "\"softwareVersion\": $software_version_for_master"

echo "Get model version VID: $vid_for_master PID: $pid_2_for_master"
result=$($DCLD_BIN_NEW query model model-version --vid=$vid_for_master --pid=$pid_2_for_master --softwareVersion=$software_version_for_master)
check_response "$result" "\"vid\": $vid_for_master"
check_response "$result" "\"pid\": $pid_2_for_master"
check_response "$result" "\"softwareVersion\": $software_version_for_master"

test_divider

# COMPLIANCE

echo "Get certified model vid=$vid_for_master pid=$pid_1_for_master"
result=$($DCLD_BIN_NEW query compliance certified-model --vid=$vid_for_master --pid=$pid_1_for_master --softwareVersion=$software_version_for_master --certificationType=$certification_type_for_master)
check_response "$result" "\"value\": true"
check_response "$result" "\"vid\": $vid_for_master"
check_response "$result" "\"pid\": $pid_1_for_master"
check_response "$result" "\"softwareVersion\": $software_version_for_master"
check_response "$result" "\"certificationType\": \"$certification_type_for_master\""

echo "Get revoked Model with VID: $vid_for_master PID: $pid_2_for_master"
result=$($DCLD_BIN_NEW query compliance revoked-model --vid=$vid_for_master --pid=$pid_2_for_master --softwareVersion=$software_version_for_master --certificationType=$certification_type_for_master)
check_response "$result" "\"vid\": $vid_for_master"
check_response "$result" "\"pid\": $pid_2_for_master"

echo "Get certified model with VID: $vid_for_master PID: $pid_1_for_master"
result=$($DCLD_BIN_NEW query compliance certified-model --vid=$vid_for_master --pid=$pid_1_for_master --softwareVersion=$software_version_for_master --certificationType=$certification_type_for_master)
check_response "$result" "\"value\": true"
check_response "$result" "\"vid\": $vid_for_master"
check_response "$result" "\"pid\": $pid_1_for_master"

echo "Get provisional model with VID: $vid_for_master PID: $pid_2_for_master"
result=$($DCLD_BIN_NEW query compliance provisional-model --vid=$vid_for_master --pid=$pid_2_for_master --softwareVersion=$software_version_for_master --certificationType=$certification_type_for_master)
check_response "$result" "\"value\": false"
check_response "$result" "\"vid\": $vid_for_master"
check_response "$result" "\"pid\": $pid_2_for_master"

echo "Get compliance-info model with VID: $vid_for_master PID: $pid_1_for_master"
result=$($DCLD_BIN_NEW query compliance compliance-info --vid=$vid_for_master --pid=$pid_1_for_master --softwareVersion=$software_version_for_master --certificationType=$certification_type_for_master)
check_response "$result" "\"vid\": $vid_for_master"
check_response "$result" "\"pid\": $pid_1_for_master"
check_response "$result" "\"softwareVersion\": $software_version_for_master"
check_response "$result" "\"certificationType\": \"$certification_type_for_master\""

echo "Get compliance-info model with VID: $vid_for_master PID: $pid_2_for_master"
result=$($DCLD_BIN_NEW query compliance compliance-info --vid=$vid_for_master --pid=$pid_2_for_master --softwareVersion=$software_version_for_master --certificationType=$certification_type_for_master)
check_response "$result" "\"vid\": $vid_for_master"
check_response "$result" "\"pid\": $pid_2_for_master"
check_response "$result" "\"softwareVersion\": $software_version_for_master"
check_response "$result" "\"certificationType\": \"$certification_type_for_master\""

echo "Get device software compliance cDCertificateId=$cd_certificate_id_for_master"
result=$($DCLD_BIN_NEW query compliance device-software-compliance --cdCertificateId=$cd_certificate_id_for_master)
check_response "$result" "\"vid\": $vid_for_master"
check_response "$result" "\"pid\": $pid_1_for_master"

echo "Get all certified models"
result=$($DCLD_BIN_NEW query compliance all-certified-models)
check_response "$result" "\"vid\": $vid_for_master"
check_response "$result" "\"pid\": $pid_1_for_master"

echo "Get all provisional models"
result=$($DCLD_BIN_NEW query compliance all-provisional-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_3"

echo "Get all revoked models"
result=$($DCLD_BIN_NEW query compliance all-revoked-models)
check_response "$result" "\"vid\": $vid_for_master"
check_response "$result" "\"pid\": $pid_2_for_master"

echo "Get all compliance infos"
result=$($DCLD_BIN_NEW query compliance all-compliance-info)
check_response "$result" "\"vid\": $vid_for_master"
check_response "$result" "\"pid\": $pid_1_for_master"
check_response "$result" "\"pid\": $pid_2_for_master"

echo "Get all device software compliances"
result=$($DCLD_BIN_NEW query compliance all-device-software-compliance)
check_response "$result" "\"vid\": $vid_for_master"
check_response "$result" "\"pid\": $pid_1_for_master"
check_response "$result" "\"cDCertificateId\": \"$cd_certificate_id_for_master\""

test_divider

# Validator

echo "Get node"
result=$(docker exec "$VALIDATOR_DEMO_CONTAINER_NAME" /bin/sh -c "echo test1234 | dcld query validator all-nodes")
check_response "$result" "\"owner\": \"$validator_address\""

test_divider

echo "Upgrade from 1.4.4 to master passed"

rm -f $DCLD_BIN_OLD
rm -f $DCLD_BIN_NEW
