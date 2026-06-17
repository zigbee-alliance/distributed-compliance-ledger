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

echo "ISSUE #593: Delete model version vid=$vid_for_1_6_0 pid=$pid_3_for_1_6_0 sv=$software_version_2_for_1_6_0"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model delete-model-version --vid=$vid_for_1_6_0 --pid=$pid_3_for_1_6_0 --softwareVersion=$software_version_2_for_1_6_0 --from=$vendor_account_for_1_2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "ISSUE #593: Get all model versions(ghost Model Version should be returned)"
result=$($DCLD_BIN_NEW query model all-model-versions --vid=$vid_for_1_6_0 --pid=$pid_3_for_1_6_0)
check_response "$result" "\"vid\": $vid_for_1_6_0"
check_response "$result" "\"pid\": $pid_3_for_1_6_0"
check_response "$result" "$software_version_1_for_1_6_0"
response_does_not_contain "$result" "$software_version_2_for_1_6_0"

test_divider

# Upgrade constants

plan_name="v1.6.0"
upgrade_checksum="sha256:ab07c87f6bdd1ee903ca4a7c26c8a503b2f1d14c63acf6ebfa6968b41efb236f"
# TODO it must be v1.6.0 before actual release
binary_version_new="v1.6.0-0.dev.4"

DCLD_BIN_OLD="/tmp/dcld_bins/dcld_v1.5.2"
DCLD_BIN_NEW="/tmp/dcld_bins/dcld_v1.6.0-0.dev.4" # TODO it must be v1.6.0 before actual release
$DCLD_BIN_NEW config broadcast-mode sync
########################################################################################

# Upgrade to 1.6 version

get_height current_height
echo "Current height is $current_height"

plan_height=$(expr $current_height \+ 20)

test_divider

echo "Propose upgrade $plan_name at height $plan_height"
echo "https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/$binary_version_new/dcld?checksum=$upgrade_checksum"
result=$(echo $passphrase | $DCLD_BIN_OLD tx dclupgrade propose-upgrade --name=$plan_name --upgrade-height=$plan_height --upgrade-info="{\"binaries\":{\"linux/amd64\":\"https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/$binary_version_new/dcld?checksum=$upgrade_checksum\"}}" --from $trustee_account_1 --yes)
echo "$result"
check_response "$result" "\"code\": 0"

test_divider

echo "Approve upgrade $plan_name"
result=$(echo $passphrase | $DCLD_BIN_OLD tx dclupgrade approve-upgrade --name $plan_name --from $trustee_account_2 --yes)
echo "$result"
check_response "$result" "\"code\": 0"

echo "Approve upgrade $plan_name"
result=$(echo $passphrase | $DCLD_BIN_OLD tx dclupgrade approve-upgrade --name $plan_name --from $trustee_account_3 --yes)
echo "$result"
check_response "$result" "\"code\": 0"

echo "Approve upgrade $plan_name"
result=$(echo $passphrase | $DCLD_BIN_OLD tx dclupgrade approve-upgrade --name $plan_name --from $trustee_account_4 --yes)
echo "$result"
check_response "$result" "\"code\": 0"

test_divider

echo "Wait for block height to become greater than upgrade $plan_name plan height"
wait_for_height $(expr $plan_height + 1) 300 outage-safe

test_divider

echo "Verify that no upgrade has been scheduled anymore"
result=$($DCLD_BIN_NEW query upgrade plan 2>&1) || true
check_response_and_report "$result" "no upgrade scheduled" raw

test_divider

echo "Verify that upgrade is applied"
result=$($DCLD_BIN_NEW query upgrade applied $plan_name)
echo "$result"

test_divider

########################################################################################

echo "Verify that new data is not corrupted"

test_divider

# MODEL

echo "ISSUE #593: After migration, ghost Model Versions must be removed: Get all model versions for VID: $vid_for_1_6_0 PID: $pid_3_for_1_6_0"
result=$($DCLD_BIN_NEW query model all-model-versions --vid=$vid_for_1_6_0 --pid=$pid_3_for_1_6_0)
echo "$result"
check_response "$result" "Not Found"

echo "ISSUE #593: Now we can remove Model. Delete model vid=$vid_for_1_6_0 pid=$pid_3_for_1_6_0"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model delete-model --vid=$vid_for_1_6_0 --pid=$pid_3_for_1_6_0 --from=$vendor_account_for_1_2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Get Model with VID: $vid_for_1_5_2 PID: $pid_1_for_1_5_2"
result=$($DCLD_BIN_NEW query model get-model --vid=$vid_for_1_5_2 --pid=$pid_1_for_1_5_2)
check_response "$result" "\"vid\": $vid_for_1_5_2"
check_response "$result" "\"pid\": $pid_1_for_1_5_2"
check_response "$result" "\"productLabel\": \"$product_label_for_1_5_2\""

echo "Get Model with VID: $vid_for_1_5_2 PID: $pid_2_for_1_5_2"
result=$($DCLD_BIN_NEW query model get-model --vid=$vid_for_1_5_2 --pid=$pid_2_for_1_5_2)
check_response "$result" "\"vid\": $vid_for_1_5_2"
check_response "$result" "\"pid\": $pid_2_for_1_5_2"
check_response "$result" "\"productLabel\": \"$product_label_for_1_5_2\""

echo "Check Model with VID: $vid_for_1_5_2 PID: $pid_2_for_1_5_2 updated"
result=$($DCLD_BIN_NEW query model get-model --vid=$vid --pid=$pid_2)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_2"
check_response "$result" "\"productLabel\": \"$product_label_for_1_5_2\""
check_response "$result" "\"partNumber\": \"$part_number_for_1_5_2\""

echo "Check Model version with VID: $vid PID: $pid_2 updated"
result=$($DCLD_BIN_NEW query model model-version --vid=$vid --pid=$pid_2  --softwareVersion=$software_version)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_2"
check_response "$result" "\"minApplicableSoftwareVersion\": $min_applicable_software_version_for_1_5_2"
check_response "$result" "\"maxApplicableSoftwareVersion\": $max_applicable_software_version_for_1_5_2"

echo "Get all models"
result=$($DCLD_BIN_NEW query model all-models)
check_response "$result" "\"vid\": $vid_for_1_5_2"
check_response "$result" "\"pid\": $pid_1_for_1_5_2"
check_response "$result" "\"pid\": $pid_2_for_1_5_2"

echo "Get all model versions"
result=$($DCLD_BIN_NEW query model all-model-versions --vid=$vid_for_1_5_2 --pid=$pid_1_for_1_5_2)
check_response "$result" "\"vid\": $vid_for_1_5_2"
check_response "$result" "\"pid\": $pid_1_for_1_5_2"

echo "Get Vendor Models with VID: ${vid_for_1_5_2}"
result=$($DCLD_BIN_NEW query model vendor-models --vid=$vid_for_1_5_2)
check_response "$result" "\"pid\": $pid_1_for_1_5_2"
check_response "$result" "\"pid\": $pid_2_for_1_5_2"

echo "Get model version VID: $vid_for_1_5_2 PID: $pid_1_for_1_5_2"
result=$($DCLD_BIN_NEW query model model-version --vid=$vid_for_1_5_2 --pid=$pid_1_for_1_5_2 --softwareVersion=$software_version_for_1_5_2)
check_response "$result" "\"vid\": $vid_for_1_5_2"
check_response "$result" "\"pid\": $pid_1_for_1_5_2"
check_response "$result" "\"softwareVersion\": $software_version_for_1_5_2"

echo "Get model version VID: $vid_for_1_5_2 PID: $pid_2_for_1_5_2"
result=$($DCLD_BIN_NEW query model model-version --vid=$vid_for_1_5_2 --pid=$pid_2_for_1_5_2 --softwareVersion=$software_version_for_1_5_2)
check_response "$result" "\"vid\": $vid_for_1_5_2"
check_response "$result" "\"pid\": $pid_2_for_1_5_2"
check_response "$result" "\"softwareVersion\": $software_version_for_1_5_2"

test_divider

# COMPLIANCE
# After upgrade to v1.6.0, the compliance record certified in 08 must remain queryable.
# Schema-v1 bump (#730) tightens write-path constraints but pre-existing stored records keep their values.

echo "Get compliance-info created in 1.5.2 for VID: $vid_for_1_5_2 PID: $pid_1_for_1_5_2"
result=$($DCLD_BIN_NEW query compliance compliance-info --vid=$vid_for_1_5_2 --pid=$pid_1_for_1_5_2 --softwareVersion=$software_version_for_1_5_2 --certificationType=$certification_type_for_1_5_2)
check_response "$result" "\"vid\": $vid_for_1_5_2"
check_response "$result" "\"pid\": $pid_1_for_1_5_2"
check_response "$result" "\"softwareVersion\": $software_version_for_1_5_2"
check_response "$result" "\"certificationType\": \"$certification_type_for_1_5_2\""
check_response "$result" "\"cDCertificateId\": \"$cd_certificate_id_for_1_5_2\""

echo "Get certified-model created in 1.5.2 for VID: $vid_for_1_5_2 PID: $pid_1_for_1_5_2"
result=$($DCLD_BIN_NEW query compliance certified-model --vid=$vid_for_1_5_2 --pid=$pid_1_for_1_5_2 --softwareVersion=$software_version_for_1_5_2 --certificationType=$certification_type_for_1_5_2)
check_response "$result" "\"value\": true"
check_response "$result" "\"vid\": $vid_for_1_5_2"
check_response "$result" "\"pid\": $pid_1_for_1_5_2"

echo "Get device-software-compliance created in 1.5.2 by cDCertificateId=$cd_certificate_id_for_1_5_2"
result=$($DCLD_BIN_NEW query compliance device-software-compliance --cdCertificateId=$cd_certificate_id_for_1_5_2)
check_response "$result" "\"cDCertificateId\": \"$cd_certificate_id_for_1_5_2\""
check_response "$result" "\"vid\": $vid_for_1_5_2"
check_response "$result" "\"pid\": $pid_1_for_1_5_2"

test_divider

########################################################################################

# after upgrade constants

vid_for_1_6_0=65520
pid_1_for_1_6_0=60
pid_2_for_1_6_0=70
pid_3_for_1_6_0=58
device_type_id_for_1_6_0=4434
product_name_for_1_6_0="ProductName_1_6_0"
product_label_for_1_6_0="ProductLabel_1_6_0"
icd_user_active_mode_trigger_hint_for_1_6_0=5
icd_user_active_mode_trigger_instruction_for_1_6_0="icd_user_active_mode_trigger_hint_for_1_6_0"
factory_reset_steps_hint_for_1_6_0=4
factory_reset_steps_instruction_for_1_6_0="factory_reset_steps_instruction_for_1_6_0"
commissioning_mode_sec_hint_for_1_6_0=8
commissioning_custom_flow_for_1_6_0=0
specification_version_for_1_6_0=3
part_number_for_1_6_0="RCU2246M"
software_version_for_1_6_0=5
software_version_string_for_1_6_0="5.0"
cd_version_number_for_1_6_0=514
min_applicable_software_version_for_1_6_0=9
max_applicable_software_version_for_1_6_0=9000

certification_type_for_1_6_0="matter"
certification_date_for_1_6_0="2024-02-01T00:00:00Z"
provisional_date_for_1_6_0="2017-01-01T00:00:00Z"
cd_certificate_id_for_1_6_0="21DEXZ"

test_data_url_for_1_6_0="https://url.data.dclmodel-1.6"

vendor_name_for_1_6_0="Vendor_1_6_0"
company_legal_name_for_1_6_0="LegalCompanyName_1_6_0"
company_preferred_name_for_1_6_0="CompanyPreferredName_1_6_0"
vendor_landing_page_url_for_1_6_0="https://www.new_1_6_0_example.com"

vendor_account_for_1_6_0="vendor_account_1_6_0"

echo "Create Vendor account $vendor_account_for_1_6_0"

result="$(echo $passphrase | $DCLD_BIN_NEW keys add "$vendor_account_for_1_6_0")"
echo "keys add $result"
_address=$(echo $passphrase | $DCLD_BIN_NEW keys show $vendor_account_for_1_6_0 -a)
_pubkey=$(echo $passphrase | $DCLD_BIN_NEW keys show $vendor_account_for_1_6_0 -p)
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth propose-add-account --address="$_address" --pubkey="$_pubkey" --vid="$vid_for_1_6_0" --roles="Vendor" --from "$trustee_account_1" --yes)
echo "propose-add-account $result"
result=$(get_txn_result "$result")
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth approve-add-account --address="$_address" --from "$trustee_account_2" --yes)
result=$(get_txn_result "$result")
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth approve-add-account --address="$_address" --from "$trustee_account_3" --yes)
result=$(get_txn_result "$result")
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth approve-add-account --address="$_address" --from "$trustee_account_4" --yes)
result=$(get_txn_result "$result")

random_string user_16
echo "$user_16 generates keys"
cmd="(echo $passphrase; echo $passphrase) | $DCLD_BIN_NEW keys add $user_16"
result="$(bash -c "$cmd")"
user_16_address=$(echo $passphrase | $DCLD_BIN_NEW keys show $user_16 -a)
user_16_pubkey=$(echo $passphrase | $DCLD_BIN_NEW keys show $user_16 -p)

# send all ledger update transactions after upgrade

# MODEL and MODEL_VERSION

echo "Add model vid=$vid_for_1_6_0 pid=$pid_1_for_1_6_0"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model --vid=$vid_for_1_6_0 --pid=$pid_1_for_1_6_0 \
  --deviceTypeID=$device_type_id_for_1_6_0 --productName=$product_name_for_1_6_0 --productLabel=$product_label_for_1_6_0 --partNumber=$part_number_for_1_6_0 \
  --icdUserActiveModeTriggerHint="$icd_user_active_mode_trigger_hint_for_1_6_0" --icdUserActiveModeTriggerInstruction="$icd_user_active_mode_trigger_instruction_for_1_6_0" \
  --factoryResetStepsHint="$factory_reset_steps_hint_for_1_6_0" --factoryResetStepsInstruction="$factory_reset_steps_instruction_for_1_6_0" \
  --commissioningCustomFlow=$commissioning_custom_flow_for_1_6_0 --commissioningModeSecondaryStepsHint="$commissioning_mode_sec_hint_for_1_6_0" \
  --from=$vendor_account_for_1_6_0 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Add model version vid=$vid_for_1_6_0 pid=$pid_1_for_1_6_0"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model-version --vid=$vid_for_1_6_0 --pid=$pid_1_for_1_6_0 --softwareVersion=$software_version_for_1_6_0 --softwareVersionString=$software_version_string_for_1_6_0 --cdVersionNumber=$cd_version_number_for_1_6_0 --minApplicableSoftwareVersion=$min_applicable_software_version_for_1_6_0 --maxApplicableSoftwareVersion=$max_applicable_software_version_for_1_6_0 --specificationVersion="$specification_version_for_1_6_0" --from=$vendor_account_for_1_6_0 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Add model vid=$vid_for_1_6_0 pid=$pid_2_for_1_6_0"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model --vid=$vid_for_1_6_0 --pid=$pid_2_for_1_6_0 --deviceTypeID=$device_type_id_for_1_6_0 \
  --productName=$product_name_for_1_6_0 --productLabel=$product_label_for_1_6_0 --partNumber=$part_number_for_1_6_0 --commissioningCustomFlow=$commissioning_custom_flow_for_1_6_0 --from=$vendor_account_for_1_6_0 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Add model version vid=$vid_for_1_6_0 pid=$pid_2_for_1_6_0"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model-version --vid=$vid_for_1_6_0 --pid=$pid_2_for_1_6_0 --softwareVersion=$software_version_for_1_6_0 --softwareVersionString=$software_version_string_for_1_6_0 --cdVersionNumber=$cd_version_number_for_1_6_0 --minApplicableSoftwareVersion=$min_applicable_software_version_for_1_6_0 --maxApplicableSoftwareVersion=$max_applicable_software_version_for_1_6_0 --from=$vendor_account_for_1_6_0 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Update model vid=$vid_for_1_5_2 pid=$pid_2_for_1_5_2"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model update-model --vid=$vid_for_1_5_2 --pid=$pid_2_for_1_5_2 --productName=$product_name_for_1_6_0 --productLabel=$product_label_for_1_6_0 --partNumber=$part_number_for_1_6_0 --from=$vendor_account_for_1_5_2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

# Model with discoveryCapabilitiesBitmask=20 — allowed range widened from 0-14 to 0-30 in v1.6.0.
pid_widened_bitmask_for_1_6_0=$((pid_3_for_1_6_0 + 100))
echo "Add model with widened discoveryCapabilitiesBitmask=20 vid=$vid_for_1_6_0 pid=$pid_widened_bitmask_for_1_6_0"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model --vid=$vid_for_1_6_0 --pid=$pid_widened_bitmask_for_1_6_0 \
  --deviceTypeID=$device_type_id_for_1_6_0 --productName=$product_name_for_1_6_0 --productLabel=$product_label_for_1_6_0 --partNumber=$part_number_for_1_6_0 \
  --commissioningCustomFlow=$commissioning_custom_flow_for_1_6_0 --discoveryCapabilitiesBitmask=20 \
  --from=$vendor_account_for_1_6_0 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

# Compliance writes after upgrade — schemaVersion=1 (default), specificationVersion required by #730.
echo "Certify-model on v1.6.0 with v1 schema vid=$vid_for_1_6_0 pid=$pid_1_for_1_6_0"
result=$(echo $passphrase | $DCLD_BIN_NEW tx compliance certify-model \
  --vid=$vid_for_1_6_0 --pid=$pid_1_for_1_6_0 \
  --softwareVersion=$software_version_for_1_6_0 --softwareVersionString=$software_version_string_for_1_6_0 \
  --cdVersionNumber=$cd_version_number_for_1_6_0 \
  --certificationType=$certification_type_for_1_6_0 --certificationDate=$certification_date_for_1_6_0 \
  --specificationVersion=$specification_version_for_1_6_0 \
  --cdCertificateId=$cd_certificate_id_for_1_5_2 \
  --schemaVersion=1 \
  --from=$certification_center_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Provision-model on v1.6.0 with v1 schema vid=$vid_for_1_6_0 pid=$pid_2_for_1_6_0"
result=$(echo $passphrase | $DCLD_BIN_NEW tx compliance provision-model \
  --vid=$vid_for_1_6_0 --pid=$pid_2_for_1_6_0 \
  --softwareVersion=$software_version_for_1_6_0 --softwareVersionString=$software_version_string_for_1_6_0 \
  --cdVersionNumber=$cd_version_number_for_1_6_0 \
  --certificationType=$certification_type_for_1_6_0 --provisionalDate=$provisional_date_for_1_6_0 \
  --specificationVersion=$specification_version_for_1_6_0 \
  --cdCertificateId=$cd_certificate_id_for_1_5_2 \
  --schemaVersion=1 \
  --from=$certification_center_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Verify that new data is not corrupted"

test_divider

# MODEL

echo "Get Model with VID: $vid_for_1_6_0 PID: $pid_1_for_1_6_0"
result=$($DCLD_BIN_NEW query model get-model --vid=$vid_for_1_6_0 --pid=$pid_1_for_1_6_0)
check_response "$result" "\"vid\": $vid_for_1_6_0"
check_response "$result" "\"pid\": $pid_1_for_1_6_0"
check_response "$result" "\"productLabel\": \"$product_label_for_1_6_0\""

echo "Get Model with VID: $vid_for_1_6_0 PID: $pid_2_for_1_6_0"
result=$($DCLD_BIN_NEW query model get-model --vid=$vid_for_1_6_0 --pid=$pid_2_for_1_6_0)
check_response "$result" "\"vid\": $vid_for_1_6_0"
check_response "$result" "\"pid\": $pid_2_for_1_6_0"
check_response "$result" "\"productLabel\": \"$product_label_for_1_6_0\""

echo "Check Model with VID: $vid_for_1_5_2 PID: $pid_2_for_1_5_2 updated"
result=$($DCLD_BIN_NEW query model get-model --vid=$vid_for_1_5_2 --pid=$pid_2_for_1_5_2)
check_response "$result" "\"vid\": $vid_for_1_5_2"
check_response "$result" "\"pid\": $pid_2_for_1_5_2"
check_response "$result" "\"productLabel\": \"$product_label_for_1_6_0\""
check_response "$result" "\"partNumber\": \"$part_number_for_1_6_0\""

echo "Get all models"
result=$($DCLD_BIN_NEW query model all-models)
check_response "$result" "\"vid\": $vid_for_1_6_0"
check_response "$result" "\"pid\": $pid_1_for_1_6_0"
check_response "$result" "\"pid\": $pid_2_for_1_6_0"

echo "Get model version VID: $vid_for_1_6_0 PID: $pid_1_for_1_6_0"
result=$($DCLD_BIN_NEW query model model-version --vid=$vid_for_1_6_0 --pid=$pid_1_for_1_6_0 --softwareVersion=$software_version_for_1_6_0)
check_response "$result" "\"vid\": $vid_for_1_6_0"
check_response "$result" "\"pid\": $pid_1_for_1_6_0"
check_response "$result" "\"softwareVersion\": $software_version_for_1_6_0"
check_response "$result" "\"specificationVersion\": $specification_version_for_1_6_0"

echo "Get Model with widened discoveryCapabilitiesBitmask VID: $vid_for_1_6_0 PID: $pid_widened_bitmask_for_1_6_0"
result=$($DCLD_BIN_NEW query model get-model --vid=$vid_for_1_6_0 --pid=$pid_widened_bitmask_for_1_6_0)
check_response "$result" "\"vid\": $vid_for_1_6_0"
check_response "$result" "\"pid\": $pid_widened_bitmask_for_1_6_0"
check_response "$result" "\"discoveryCapabilitiesBitmask\": 20"

# Verify compliance writes made after the upgrade landed under schema v1 (specificationVersion now stored on ComplianceInfo).

echo "Get certified-model written after upgrade for VID: $vid_for_1_6_0 PID: $pid_1_for_1_6_0"
result=$($DCLD_BIN_NEW query compliance certified-model --vid=$vid_for_1_6_0 --pid=$pid_1_for_1_6_0 --softwareVersion=$software_version_for_1_6_0 --certificationType=$certification_type_for_1_6_0)
check_response "$result" "\"value\": true"
check_response "$result" "\"vid\": $vid_for_1_6_0"

echo "Get compliance-info written after upgrade for VID: $vid_for_1_6_0 PID: $pid_1_for_1_6_0"
result=$($DCLD_BIN_NEW query compliance compliance-info --vid=$vid_for_1_6_0 --pid=$pid_1_for_1_6_0 --softwareVersion=$software_version_for_1_6_0 --certificationType=$certification_type_for_1_6_0)
check_response "$result" "\"schemaVersion\": 1"
check_response "$result" "\"specificationVersion\": $specification_version_for_1_6_0"

echo "Get provisional-model written after upgrade for VID: $vid_for_1_6_0 PID: $pid_2_for_1_6_0"
result=$($DCLD_BIN_NEW query compliance provisional-model --vid=$vid_for_1_6_0 --pid=$pid_2_for_1_6_0 --softwareVersion=$software_version_for_1_6_0 --certificationType=$certification_type_for_1_6_0)
check_response "$result" "\"value\": true"
check_response "$result" "\"vid\": $vid_for_1_6_0"

test_divider

########################################################################################
# v1.6.0 enforcement smoke tests — confirm that the upgraded binary applies the new
# Matter-spec-conformance constraints (#703, #713, #718, #730, #727, #726). Each
# negative case is exercised against $DCLD_BIN_NEW only; the same checks run as
# unit / negative-cases tests, but here we verify they are enabled post-upgrade.
########################################################################################

# Helper PID values for v1.6.0 enforcement tests (distinct from previously-used PIDs)
pid_neg1_for_1_6_0=$((pid_2_for_1_6_0 + 100))
pid_neg2_for_1_6_0=$((pid_2_for_1_6_0 + 101))
pid_neg3_for_1_6_0=$((pid_2_for_1_6_0 + 102))
pid_neg4_for_1_6_0=$((pid_2_for_1_6_0 + 103))
pid_neg5_for_1_6_0=$((pid_2_for_1_6_0 + 104))
pid_vendor_admin_for_1_6_0=$((pid_2_for_1_6_0 + 105))
pid_ota_for_1_6_0=$((pid_2_for_1_6_0 + 106))
sv_neg_for_1_6_0=$((software_version_for_1_6_0 + 100))
svs_neg_for_1_6_0="6.0"

# --- Compliance schema-v1 negative cases (#730) -------------------------------------

echo "Reject certify-model with schemaVersion=0 (must be 1 after #730)"
result=$(echo $passphrase | $DCLD_BIN_NEW tx compliance certify-model \
  --vid=$vid_for_1_6_0 --pid=$pid_neg1_for_1_6_0 \
  --softwareVersion=$sv_neg_for_1_6_0 --softwareVersionString=$svs_neg_for_1_6_0 \
  --cdVersionNumber=$cd_version_number_for_1_6_0 \
  --certificationType=$certification_type_for_1_6_0 --certificationDate=$certification_date_for_1_6_0 \
  --specificationVersion=$specification_version_for_1_6_0 \
  --cdCertificateId=$cd_certificate_id_for_1_5_2 \
  --schemaVersion=0 \
  --from=$certification_center_account --yes 2>&1) || true
check_response_and_report "$result" "SchemaVersion must be equal 1" raw

echo "Reject certify-model with cdCertificateId shorter than 19 chars"
short_cd_id_for_1_6_0="1234567890abcdefgh" # 18 chars
result=$(echo $passphrase | $DCLD_BIN_NEW tx compliance certify-model \
  --vid=$vid_for_1_6_0 --pid=$pid_neg1_for_1_6_0 \
  --softwareVersion=$sv_neg_for_1_6_0 --softwareVersionString=$svs_neg_for_1_6_0 \
  --cdVersionNumber=$cd_version_number_for_1_6_0 \
  --certificationType=$certification_type_for_1_6_0 --certificationDate=$certification_date_for_1_6_0 \
  --specificationVersion=$specification_version_for_1_6_0 \
  --cdCertificateId="$short_cd_id_for_1_6_0" \
  --schemaVersion=1 \
  --from=$certification_center_account --yes 2>&1) || true
check_response_and_report "$result" "minimum length for CDCertificateId allowed is 19" raw

echo "Reject certify-model with cdCertificateId longer than 19 chars"
long_cd_id_for_1_6_0="12345678910abcdefghX" # 20 chars
result=$(echo $passphrase | $DCLD_BIN_NEW tx compliance certify-model \
  --vid=$vid_for_1_6_0 --pid=$pid_neg1_for_1_6_0 \
  --softwareVersion=$sv_neg_for_1_6_0 --softwareVersionString=$svs_neg_for_1_6_0 \
  --cdVersionNumber=$cd_version_number_for_1_6_0 \
  --certificationType=$certification_type_for_1_6_0 --certificationDate=$certification_date_for_1_6_0 \
  --specificationVersion=$specification_version_for_1_6_0 \
  --cdCertificateId="$long_cd_id_for_1_6_0" \
  --schemaVersion=1 \
  --from=$certification_center_account --yes 2>&1) || true
check_response_and_report "$result" "maximum length for CDCertificateId allowed is 19" raw

echo "Reject certify-model with specificationVersion=0 (now required by #730)"
result=$(echo $passphrase | $DCLD_BIN_NEW tx compliance certify-model \
  --vid=$vid_for_1_6_0 --pid=$pid_neg1_for_1_6_0 \
  --softwareVersion=$sv_neg_for_1_6_0 --softwareVersionString=$svs_neg_for_1_6_0 \
  --cdVersionNumber=$cd_version_number_for_1_6_0 \
  --certificationType=$certification_type_for_1_6_0 --certificationDate=$certification_date_for_1_6_0 \
  --specificationVersion=0 \
  --cdCertificateId=$cd_certificate_id_for_1_5_2 \
  --schemaVersion=1 \
  --from=$certification_center_account --yes 2>&1) || true
check_response_and_report "$result" "SpecificationVersion is a required field" raw

echo "Reject certify-model with certificationType longer than 20 chars (#713)"
long_certification_type="this_certification_type_is_way_too_long"
result=$(echo $passphrase | $DCLD_BIN_NEW tx compliance certify-model \
  --vid=$vid_for_1_6_0 --pid=$pid_neg1_for_1_6_0 \
  --softwareVersion=$sv_neg_for_1_6_0 --softwareVersionString=$svs_neg_for_1_6_0 \
  --cdVersionNumber=$cd_version_number_for_1_6_0 \
  --certificationType="$long_certification_type" --certificationDate=$certification_date_for_1_6_0 \
  --specificationVersion=$specification_version_for_1_6_0 \
  --cdCertificateId=$cd_certificate_id_for_1_5_2 \
  --schemaVersion=1 \
  --from=$certification_center_account --yes 2>&1) || true
check_response_and_report "$result" "maximum length for CertificationType allowed is 20" raw

echo "Reject provision-model with schemaVersion=0"
result=$(echo $passphrase | $DCLD_BIN_NEW tx compliance provision-model \
  --vid=$vid_for_1_6_0 --pid=$pid_neg2_for_1_6_0 \
  --softwareVersion=$sv_neg_for_1_6_0 --softwareVersionString=$svs_neg_for_1_6_0 \
  --cdVersionNumber=$cd_version_number_for_1_6_0 \
  --certificationType=$certification_type_for_1_6_0 --provisionalDate=$provisional_date_for_1_6_0 \
  --specificationVersion=$specification_version_for_1_6_0 \
  --cdCertificateId=$cd_certificate_id_for_1_5_2 \
  --schemaVersion=0 \
  --from=$certification_center_account --yes 2>&1) || true
check_response_and_report "$result" "SchemaVersion must be equal 1" raw

echo "Reject revoke-model with schemaVersion=0"
result=$(echo $passphrase | $DCLD_BIN_NEW tx compliance revoke-model \
  --vid=$vid_for_1_6_0 --pid=$pid_neg2_for_1_6_0 \
  --softwareVersion=$sv_neg_for_1_6_0 --softwareVersionString=$svs_neg_for_1_6_0 \
  --certificationType=$certification_type_for_1_6_0 --revocationDate=$certification_date_for_1_6_0 \
  --cdVersionNumber=$cd_version_number_for_1_6_0 \
  --schemaVersion=0 \
  --from=$certification_center_account --yes 2>&1) || true
check_response_and_report "$result" "SchemaVersion must be equal 1" raw

test_divider

# --- Model required-fields negative cases (#718) ------------------------------------

echo "Reject add-model without productLabel (now required by #718)"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model \
  --vid=$vid_for_1_6_0 --pid=$pid_neg3_for_1_6_0 \
  --deviceTypeID=$device_type_id_for_1_6_0 --productName=$product_name_for_1_6_0 \
  --partNumber=$part_number_for_1_6_0 \
  --commissioningCustomFlow=$commissioning_custom_flow_for_1_6_0 \
  --from=$vendor_account_for_1_6_0 --yes 2>&1) || true
check_response_and_report "$result" 'required flag(s) "productLabel" not set' raw

echo "Reject add-model without partNumber (now required by #718)"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model \
  --vid=$vid_for_1_6_0 --pid=$pid_neg3_for_1_6_0 \
  --deviceTypeID=$device_type_id_for_1_6_0 --productName=$product_name_for_1_6_0 \
  --productLabel=$product_label_for_1_6_0 \
  --commissioningCustomFlow=$commissioning_custom_flow_for_1_6_0 \
  --from=$vendor_account_for_1_6_0 --yes 2>&1) || true
check_response_and_report "$result" 'required flag(s) "partNumber" not set' raw

test_divider

# --- OTA field constraint tests (#726, #727) ----------------------------------------

# First create the parent model that the OTA-related model-versions will hang off.
echo "Add model vid=$vid_for_1_6_0 pid=$pid_ota_for_1_6_0 (parent for OTA-tests)"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model \
  --vid=$vid_for_1_6_0 --pid=$pid_ota_for_1_6_0 \
  --deviceTypeID=$device_type_id_for_1_6_0 --productName=$product_name_for_1_6_0 \
  --productLabel=$product_label_for_1_6_0 --partNumber=$part_number_for_1_6_0 \
  --commissioningCustomFlow=$commissioning_custom_flow_for_1_6_0 \
  --from=$vendor_account_for_1_6_0 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Reject add-model-version with otaChecksum longer than 88 chars (#726)"
long_ota_checksum=$(printf 'A%.0s' {1..89})
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model-version \
  --vid=$vid_for_1_6_0 --pid=$pid_ota_for_1_6_0 \
  --softwareVersion=$sv_neg_for_1_6_0 --softwareVersionString=$svs_neg_for_1_6_0 \
  --cdVersionNumber=$cd_version_number_for_1_6_0 \
  --minApplicableSoftwareVersion=$min_applicable_software_version_for_1_6_0 \
  --maxApplicableSoftwareVersion=$max_applicable_software_version_for_1_6_0 \
  --otaURL="https://example.org" --otaFileSize=123 \
  --otaChecksum="$long_ota_checksum" \
  --otaChecksumType=1 \
  --from=$vendor_account_for_1_6_0 --yes 2>&1) || true
check_response_and_report "$result" "maximum length for OtaChecksum allowed is 88" raw

echo "Reject add-model-version with otaChecksumType=2 (must be in {1,7,8,10,11,12} after #727)"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model-version \
  --vid=$vid_for_1_6_0 --pid=$pid_ota_for_1_6_0 \
  --softwareVersion=$sv_neg_for_1_6_0 --softwareVersionString=$svs_neg_for_1_6_0 \
  --cdVersionNumber=$cd_version_number_for_1_6_0 \
  --minApplicableSoftwareVersion=$min_applicable_software_version_for_1_6_0 \
  --maxApplicableSoftwareVersion=$max_applicable_software_version_for_1_6_0 \
  --otaURL="https://example.org" --otaFileSize=123 \
  --otaChecksum="MjFiZmYxN2YyMTRlMGJiMGMwNzhlNzIzOGIxZWE1ODk=" \
  --otaChecksumType=2 \
  --from=$vendor_account_for_1_6_0 --yes 2>&1) || true
check_response_and_report "$result" "OtaChecksumType 2 is not supported" raw

echo "Accept add-model-version with valid OTA metadata (otaChecksumType=1, all OTA fields together)"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model-version \
  --vid=$vid_for_1_6_0 --pid=$pid_ota_for_1_6_0 \
  --softwareVersion=$sv_neg_for_1_6_0 --softwareVersionString=$svs_neg_for_1_6_0 \
  --cdVersionNumber=$cd_version_number_for_1_6_0 \
  --minApplicableSoftwareVersion=$min_applicable_software_version_for_1_6_0 \
  --maxApplicableSoftwareVersion=$max_applicable_software_version_for_1_6_0 \
  --otaURL="https://example.org" --otaFileSize=123 \
  --otaChecksum="MjFiZmYxN2YyMTRlMGJiMGMwNzhlNzIzOGIxZWE1ODk=" \
  --otaChecksumType=1 \
  --from=$vendor_account_for_1_6_0 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Verify stored model-version has the OTA fields"
result=$($DCLD_BIN_NEW query model model-version --vid=$vid_for_1_6_0 --pid=$pid_ota_for_1_6_0 --softwareVersion=$sv_neg_for_1_6_0)
check_response "$result" "\"otaChecksumType\": 1"
check_response "$result" "\"otaFileSize\": \"123\""

test_divider

# --- VendorAdmin model-write authorization (#703) -----------------------------------
# VendorAdmin can now add/edit/delete models for ANY vid (checkModelRights short-
# circuits on the VendorAdmin role). The vendor_admin_account was created back in
# 03-test-upgrade-0.12-to-1.2.sh with the VendorAdmin role and no associated VID.

vid_vendor_admin_for_1_6_0=$((vid_for_1_6_0 + 1))

echo "Add model as VendorAdmin for foreign vid=$vid_vendor_admin_for_1_6_0 pid=$pid_vendor_admin_for_1_6_0"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model \
  --vid=$vid_vendor_admin_for_1_6_0 --pid=$pid_vendor_admin_for_1_6_0 \
  --deviceTypeID=$device_type_id_for_1_6_0 --productName=$product_name_for_1_6_0 \
  --productLabel=$product_label_for_1_6_0 --partNumber=$part_number_for_1_6_0 \
  --commissioningCustomFlow=$commissioning_custom_flow_for_1_6_0 \
  --from=$vendor_admin_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Verify VendorAdmin-added model is on ledger for vid=$vid_vendor_admin_for_1_6_0 pid=$pid_vendor_admin_for_1_6_0"
result=$($DCLD_BIN_NEW query model get-model --vid=$vid_vendor_admin_for_1_6_0 --pid=$pid_vendor_admin_for_1_6_0)
check_response "$result" "\"vid\": $vid_vendor_admin_for_1_6_0"
check_response "$result" "\"pid\": $pid_vendor_admin_for_1_6_0"
check_response "$result" "\"productLabel\": \"$product_label_for_1_6_0\""

test_divider

echo "Upgrade from 1.5.2 to 1.6.0 PASSED"
