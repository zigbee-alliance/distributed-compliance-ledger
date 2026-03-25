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

plan_name="v1.5.2"
upgrade_checksum="sha256:746e4d24969f45f55b7d4a2f143ffe9609cf4f7a60c1472e38ecfe781b2327dc"
# TODO it must be v1.5.2 before actual 1.5.2 release
binary_version_new="v1.5.2"

DCLD_BIN_OLD="/tmp/dcld_bins/dcld_v1.5.1"
DCLD_BIN_NEW="/tmp/dcld_bins/dcld_v1.5.2"
$DCLD_BIN_NEW config broadcast-mode sync
########################################################################################

# Upgrade to 1.5 version

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

echo "Get Model with VID: $vid_for_1_5_1 PID: $pid_1_for_1_5_1"
result=$($DCLD_BIN_NEW query model get-model --vid=$vid_for_1_5_1 --pid=$pid_1_for_1_5_1)
check_response "$result" "\"vid\": $vid_for_1_5_1"
check_response "$result" "\"pid\": $pid_1_for_1_5_1"
check_response "$result" "\"productLabel\": \"$product_label_for_1_5_1\""
# check migration
check_response "$result" "\"commissioningModeSecondaryStepsHint\": $commissioning_mode_sec_hint_for_1_5_1"

echo "Get Model with VID: $vid_for_1_5_1 PID: $pid_2_for_1_5_1"
result=$($DCLD_BIN_NEW query model get-model --vid=$vid_for_1_5_1 --pid=$pid_2_for_1_5_1)
check_response "$result" "\"vid\": $vid_for_1_5_1"
check_response "$result" "\"pid\": $pid_2_for_1_5_1"
check_response "$result" "\"productLabel\": \"$product_label_for_1_5_1\""
# check migration
check_response "$result" "\"commissioningModeSecondaryStepsHint\": 4"

echo "Check Model with VID: $vid_for_1_5_1 PID: $pid_2_for_1_5_1 updated"
result=$($DCLD_BIN_NEW query model get-model --vid=$vid --pid=$pid_2)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_2"
check_response "$result" "\"productLabel\": \"$product_label_for_1_5_1\""
check_response "$result" "\"partNumber\": \"$part_number_for_1_5_1\""

echo "Check Model version with VID: $vid_for_1_5_1 PID: $pid_2_for_1_5_1 updated"
result=$($DCLD_BIN_NEW query model model-version --vid=$vid --pid=$pid_2  --softwareVersion=$software_version)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_2"
check_response "$result" "\"minApplicableSoftwareVersion\": $min_applicable_software_version_for_1_5_1"
check_response "$result" "\"maxApplicableSoftwareVersion\": $max_applicable_software_version_for_1_5_1"

echo "Get all models"
result=$($DCLD_BIN_NEW query model all-models)
check_response "$result" "\"vid\": $vid_for_1_5_1"
check_response "$result" "\"pid\": $pid_1_for_1_5_1"
check_response "$result" "\"pid\": $pid_2_for_1_5_1"

echo "Get all model versions"
result=$($DCLD_BIN_NEW query model all-model-versions --vid=$vid_for_1_5_1 --pid=$pid_1_for_1_5_1)
check_response "$result" "\"vid\": $vid_for_1_5_1"
check_response "$result" "\"pid\": $pid_1_for_1_5_1"

echo "Get Vendor Models with VID: ${vid_for_1_5_1}"
result=$($DCLD_BIN_NEW query model vendor-models --vid=$vid_for_1_5_1)
check_response "$result" "\"pid\": $pid_1_for_1_5_1"
check_response "$result" "\"pid\": $pid_2_for_1_5_1"

echo "Get model version VID: $vid_for_1_5_1 PID: $pid_1_for_1_5_1"
result=$($DCLD_BIN_NEW query model model-version --vid=$vid_for_1_5_1 --pid=$pid_1_for_1_5_1 --softwareVersion=$software_version_for_1_5_1)
check_response "$result" "\"vid\": $vid_for_1_5_1"
check_response "$result" "\"pid\": $pid_1_for_1_5_1"
check_response "$result" "\"softwareVersion\": $software_version_for_1_5_1"

echo "Get model version VID: $vid_for_1_5_1 PID: $pid_2_for_1_5_1"
result=$($DCLD_BIN_NEW query model model-version --vid=$vid_for_1_5_1 --pid=$pid_2_for_1_5_1 --softwareVersion=$software_version_for_1_5_1)
check_response "$result" "\"vid\": $vid_for_1_5_1"
check_response "$result" "\"pid\": $pid_2_for_1_5_1"
check_response "$result" "\"softwareVersion\": $software_version_for_1_5_1"

test_divider

########################################################################################

# after upgrade constants

vid_for_1_5_2=65519
pid_1_for_1_5_2=59
pid_2_for_1_5_2=69
pid_3_for_1_5_2=57
device_type_id_for_1_5_2=4433
product_name_for_1_5_2="ProductName_1_5_2"
product_label_for_1_5_2="ProductLabel_1_5_2"
icd_user_active_mode_trigger_hint_for_1_5_2=4
icd_user_active_mode_trigger_instruction_for_1_5_2="icd_user_active_mode_trigger_hint_for_1_5_2"
factory_reset_steps_hint_for_1_5_2=3
factory_reset_steps_instruction_for_1_5_2="factory_reset_steps_instruction_for_1_5_2"
commissioning_mode_sec_hint_for_1_5_2=7
specification_version_for_1_5_2=2
part_number_for_1_5_2="RCU2245M"
software_version_for_1_5_2=4
software_version_string_for_1_5_2="4.3"
cd_version_number_for_1_5_2=513
min_applicable_software_version_for_1_5_2=8
max_applicable_software_version_for_1_5_2=8000

certification_type_for_1_5_2="matter"
certification_date_for_1_5_2="2024-01-01T00:00:00Z"
provisional_date_for_1_5_2="2016-12-12T00:00:00Z"
cd_certificate_id_for_1_5_2="20DEXZ"
cd_certificate_id_for_1_5_2="20DEXZ"
certification_id_of_software_component_1_5_2="some_component"

test_data_url_for_1_5_2="https://url.data.dclmodel-1.5"

vendor_name_for_1_5_2="Vendor_1_5_2"
company_legal_name_for_1_5_2="LegalCompanyName_1_5_2"
company_preferred_name_for_1_5_2="CompanyPreferredName_1_5_2"
vendor_landing_page_url_for_1_5_2="https://www.new_1_5_2_example.com"

vendor_account_for_1_5_2="vendor_account_1_5_2"

echo "Create Vendor account $vendor_account_for_1_5_2"

result="$(echo $passphrase | $DCLD_BIN_NEW keys add "$vendor_account_for_1_5_2")"
echo "keys add $result"
_address=$(echo $passphrase | $DCLD_BIN_NEW keys show $vendor_account_for_1_5_2 -a)
_pubkey=$(echo $passphrase | $DCLD_BIN_NEW keys show $vendor_account_for_1_5_2 -p)
result=$(echo $passphrase | $DCLD_BIN_NEW tx auth propose-add-account --address="$_address" --pubkey="$_pubkey" --vid="$vid_for_1_5_2" --roles="Vendor" --from "$trustee_account_1" --yes)
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

# MODEL and MODEL_VERSION

echo "Add model vid=$vid_for_1_5_2 pid=$pid_1_for_1_5_2"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model --vid=$vid_for_1_5_2 --pid=$pid_1_for_1_5_2 \
  --deviceTypeID=$device_type_id_for_1_5_2 --productName=$product_name_for_1_5_2 --productLabel=$product_label_for_1_5_2 --partNumber=$part_number_for_1_5_2 \
  --icdUserActiveModeTriggerHint="$icd_user_active_mode_trigger_hint_for_1_5_2" --icdUserActiveModeTriggerInstruction="$icd_user_active_mode_trigger_instruction_for_1_5_2" \
  --factoryResetStepsHint="$factory_reset_steps_hint_for_1_5_2" --factoryResetStepsInstruction="$factory_reset_steps_instruction_for_1_5_2" \
  --commissioningModeSecondaryStepsHint="$commissioning_mode_sec_hint_for_1_5_2" \
  --from=$vendor_account_for_1_5_2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Add model version vid=$vid_for_1_5_2 pid=$pid_1_for_1_5_2"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model-version --vid=$vid_for_1_5_2 --pid=$pid_1_for_1_5_2 --softwareVersion=$software_version_for_1_5_2 --softwareVersionString=$software_version_string_for_1_5_2 --cdVersionNumber=$cd_version_number_for_1_5_2 --minApplicableSoftwareVersion=$min_applicable_software_version_for_1_5_2 --maxApplicableSoftwareVersion=$max_applicable_software_version_for_1_5_2 --specificationVersion="$specification_version_for_1_5_2" --from=$vendor_account_for_1_5_2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Add model vid=$vid_for_1_5_2 pid=$pid_2_for_1_5_2"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model --vid=$vid_for_1_5_2 --pid=$pid_2_for_1_5_2 --deviceTypeID=$device_type_id_for_1_5_2 \
  --productName=$product_name_for_1_5_2 --productLabel=$product_label_for_1_5_2 --partNumber=$part_number_for_1_5_2 --from=$vendor_account_for_1_5_2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Add model version vid=$vid_for_1_5_2 pid=$pid_2_for_1_5_2"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model-version --vid=$vid_for_1_5_2 --pid=$pid_2_for_1_5_2 --softwareVersion=$software_version_for_1_5_2 --softwareVersionString=$software_version_string_for_1_5_2 --cdVersionNumber=$cd_version_number_for_1_5_2 --minApplicableSoftwareVersion=$min_applicable_software_version_for_1_5_2 --maxApplicableSoftwareVersion=$max_applicable_software_version_for_1_5_2 --from=$vendor_account_for_1_5_2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Add model vid=$vid_for_1_5_2 pid=$pid_3_for_1_5_2"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model --vid=$vid_for_1_5_2 --pid=$pid_3_for_1_5_2 --deviceTypeID=$device_type_id_for_1_5_2 --productName=$product_name_for_1_5_2 --productLabel=$product_label_for_1_5_2 --partNumber=$part_number_for_1_5_2 --from=$vendor_account_for_1_5_2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Add model version vid=$vid_for_1_5_2 pid=$pid_3_for_1_5_2"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model add-model-version --vid=$vid_for_1_5_2 --pid=$pid_3_for_1_5_2 --softwareVersion=$software_version_for_1_5_2 --softwareVersionString=$software_version_string_for_1_5_2 --cdVersionNumber=$cd_version_number_for_1_5_2 --minApplicableSoftwareVersion=$min_applicable_software_version_for_1_5_2 --maxApplicableSoftwareVersion=$max_applicable_software_version_for_1_5_2 --from=$vendor_account_for_1_5_2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Delete model vid=$vid_for_1_5_2 pid=$pid_3_for_1_5_2"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model delete-model --vid=$vid_for_1_5_2 --pid=$pid_3_for_1_5_2 --from=$vendor_account_for_1_5_2 --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Update model vid=$vid pid=$pid_2"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model update-model --vid=$vid --pid=$pid_2 --productName=$product_name --productLabel=$product_label_for_1_5_2 --partNumber=$part_number_for_1_5_2 --from=$vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Update model version vid=$vid pid=$pid_2"
result=$(echo $passphrase | $DCLD_BIN_NEW tx model update-model-version --vid=$vid --pid=$pid_2 --softwareVersion=$software_version --minApplicableSoftwareVersion=$min_applicable_software_version_for_1_5_2 --maxApplicableSoftwareVersion=$max_applicable_software_version_for_1_5_2 --from=$vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Verify that new data is not corrupted"

test_divider


# COMPLIANCE

echo "Certify model vid=$vid_for_1_5_2 pid=$pid_1_for_1_5_2"
result=$(echo $passphrase | $DCLD_BIN_NEW tx compliance certify-model --vid=$vid_for_1_5_2 --pid=$pid_1_for_1_5_2 --softwareVersion=$software_version_for_1_5_2 --softwareVersionString=$software_version_string_for_1_5_2  --certificationType=$certification_type_for_1_5_2 --certificationDate=$certification_date_for_1_5_2 --cdCertificateId=$cd_certificate_id_for_1_5_2 --certificationIDOfSoftwareComponent=$certification_id_of_software_component_1_5_2 --cdVersionNumber=$cd_version_number_for_1_5_2 --from=$certification_center_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "Get compliance-info model with VID: $vid_for_1_5_2 PID: $pid_2_for_1_5_2"
result=$($DCLD_BIN_NEW query compliance compliance-info --vid=$vid_for_1_5_2 --pid=$pid_2_for_1_5_2 --softwareVersion=$software_version_for_1_5_2 --certificationType=$certification_type_for_1_5_2)
check_response "$result" "\"vid\": $vid_for_1_5_2"
check_response "$result" "\"pid\": $pid_2_for_1_5_2"
check_response "$result" "\"softwareVersion\": $software_version_for_1_5_2"
check_response "$result" "\"certificationType\": \"$certification_type_for_1_5_2\""
check_response "$result" "\"certificationIDOfSoftwareComponent\": \"$certification_id_of_software_component_1_5_2\""

test_divider

# MODEL

echo "Get Model with VID: $vid_for_1_5_2 PID: $pid_1_for_1_5_2"
result=$($DCLD_BIN_NEW query model get-model --vid=$vid_for_1_5_2 --pid=$pid_1_for_1_5_2)
check_response "$result" "\"vid\": $vid_for_1_5_2"
check_response "$result" "\"pid\": $pid_1_for_1_5_2"
check_response "$result" "\"productLabel\": \"$product_label_for_1_5_2\""
check_response "$result" "\"icdUserActiveModeTriggerHint\": $icd_user_active_mode_trigger_hint_for_1_5_2"
check_response "$result" "\"icdUserActiveModeTriggerInstruction\": \"$icd_user_active_mode_trigger_instruction_for_1_5_2\""
check_response "$result" "\"factoryResetStepsHint\": $factory_reset_steps_hint_for_1_5_2"
check_response "$result" "\"factoryResetStepsInstruction\": \"$factory_reset_steps_instruction_for_1_5_2\""
check_response "$result" "\"commissioningModeSecondaryStepsHint\": $commissioning_mode_sec_hint_for_1_5_2"

echo "Get Model with VID: $vid_for_1_5_2 PID: $pid_2_for_1_5_2"
result=$($DCLD_BIN_NEW query model get-model --vid=$vid_for_1_5_2 --pid=$pid_2_for_1_5_2)
check_response "$result" "\"vid\": $vid_for_1_5_2"
check_response "$result" "\"pid\": $pid_2_for_1_5_2"
check_response "$result" "\"productLabel\": \"$product_label_for_1_5_2\""
check_response "$result" "\"commissioningModeSecondaryStepsHint\": 4"

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
check_response "$result" "\"specificationVersion\": $specification_version_for_1_5_2"

echo "Get model version VID: $vid_for_1_5_2 PID: $pid_2_for_1_5_2"
result=$($DCLD_BIN_NEW query model model-version --vid=$vid_for_1_5_2 --pid=$pid_2_for_1_5_2 --softwareVersion=$software_version_for_1_5_2)
check_response "$result" "\"vid\": $vid_for_1_5_2"
check_response "$result" "\"pid\": $pid_2_for_1_5_2"
check_response "$result" "\"softwareVersion\": $software_version_for_1_5_2"

test_divider

echo "Upgrade from 1.5.1 to 1.5.2 PASSED"
