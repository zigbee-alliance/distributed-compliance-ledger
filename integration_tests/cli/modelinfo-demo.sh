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

# Preparation of Actors

echo "Create Vendor account"
create_new_account vendor_account "Vendor"

# Body

vid=$RANDOM
pid=$RANDOM
sv=$RANDOM
hv=$RANDOM
productName="Device #1"
echo "Add Model with VID: $vid PID: $pid"
echo "dclcli tx modelinfo add-model --vid=$vid --pid=$pid --softwareVersion=$sv --hardwareVersion=$hv --supportURL="https://originalsupporturl.test" --productName="$productName" --description="Device Description" --sku="SKU12FS"  --softwareVersionString="1.0b123"   --hardwareVersionString="5.1.23"  --cdVersionNumber="32" --from $vendor_account --yes)"

result=$(echo "test1234" | dclcli tx modelinfo add-model --vid=$vid --pid=$pid --softwareVersion=$sv --hardwareVersion=$hv --supportURL="https://originalsupporturl.test" --productName="$productName" --description="Device Description" --sku="SKU12FS"  --softwareVersionString="1.0b123"   --hardwareVersionString="5.1.23"  --cdVersionNumber="32" --from $vendor_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get Model with VID: $vid PID: $pid"
echo "dclcli query modelinfo model --vid=$vid --pid=$pid --softwareVersion=$sv --hardwareVersion=$hv"
result=$(dclcli query modelinfo model --vid=$vid --pid=$pid --softwareVersion=$sv --hardwareVersion=$hv)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"softwareVersion\": $sv"
check_response "$result" "\"hardwareVersion\": $hv"
check_response "$result" "\"productName\": \"$productName\""
echo "$result"

echo "Get all model infos"
result=$(dclcli query modelinfo all-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"softwareVersion\": $sv"
check_response "$result" "\"hardwareVersion\": $hv"
echo "$result"

echo "Get all vendors"
result=$(dclcli query modelinfo vendors)
check_response "$result" "\"vid\": $vid"
echo "$result"

echo "Get Vendor Models with VID: ${vid}"
result=$(dclcli query modelinfo vendor-models --vid=$vid)
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"softwareVersion\": $sv"
check_response "$result" "\"hardwareVersion\": $hv"
echo "$result"

echo "Update Model with VID: ${vid} PID: ${pid} with new description"
description="New Device Description"
result=$(echo "test1234" | dclcli tx modelinfo update-model --vid=$vid --pid=$pid --softwareVersion=$sv --hardwareVersion=$hv --cdVersionNumber="32" --from $vendor_account --yes --description "$description")
check_response "$result" "\"success\": true"
echo "$result"

echo "Get Model with VID: ${vid} PID: ${pid}"
result=$(dclcli query modelinfo model --vid=$vid --pid=$pid --softwareVersion=$sv --hardwareVersion=$hv)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"softwareVersion\": $sv"
check_response "$result" "\"hardwareVersion\": $hv"
check_response "$result" "\"description\": \"$description\""
echo "$result"

echo "Update Model with VID: ${vid} PID: ${pid} modifying supportURL"
supportURL="https://newsupporturl.test"
echo dclcli tx modelinfo update-model --vid=$vid --pid=$pid --softwareVersion=$sv --hardwareVersion=$hv --cdVersionNumber="32" --from $vendor_account --yes --supportURL "$supportURL"
result=$(echo "test1234" | dclcli tx modelinfo update-model --vid=$vid --pid=$pid --softwareVersion=$sv --hardwareVersion=$hv --cdVersionNumber="33" --from $vendor_account --yes --supportURL "$supportURL")
check_response "$result" "\"success\": true"
echo "$result"

echo "Get Model with VID: ${vid} PID: ${pid}"
result=$(dclcli query modelinfo model --vid=$vid --pid=$pid --softwareVersion=$sv --hardwareVersion=$hv)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"softwareVersion\": $sv"
check_response "$result" "\"hardwareVersion\": $hv"
check_response "$result" "\"supportURL\": \"$supportURL\""
echo "$result"
