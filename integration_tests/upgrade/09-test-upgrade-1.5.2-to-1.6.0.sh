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
upgrade_checksum="sha256:47d91b6be0b0a15e7edde7b78e3013d4eedbbb3c2c1b164de78409198548a2de"
# TODO it must be v1.6.0 before actual release
binary_version_new="v1.6.0-0.dev.3"

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

test_divider

echo "Upgrade from 1.5.2 to 1.6.0 PASSED"
