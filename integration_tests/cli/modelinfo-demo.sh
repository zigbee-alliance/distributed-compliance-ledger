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
name="Device #1"
echo "Add Model with VID: $vid PID: $pid"
result=$(echo "test1234" | dclcli tx modelinfo add-model --vid=$vid --pid=$pid --support-url="https://originalsupporturl.test" --name="$name" --description="Device Description" --sku="SKU12FS" --software-version="10123" --software-version-string="1.0b123"  --hardware-version="5123" --hardware-version-string="5.1.23"  --cd-version-number="32" --from $vendor_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get Model with VID: $vid PID: $pid"
result=$(dclcli query modelinfo model --vid=$vid --pid=$pid)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"name\": \"$name\""
echo "$result"

echo "Get all model infos"
result=$(dclcli query modelinfo all-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
echo "$result"

echo "Get all vendors"
result=$(dclcli query modelinfo vendors)
check_response "$result" "\"vid\": $vid"
echo "$result"

echo "Get Vendor Models with VID: ${vid}"
result=$(dclcli query modelinfo vendor-models --vid=$vid)
check_response "$result" "\"pid\": $pid"
echo "$result"

echo "Update Model with VID: ${vid} PID: ${pid} with new description"
description="New Device Description"
result=$(echo "test1234" | dclcli tx modelinfo update-model --vid=$vid --pid=$pid --cd-version-number="32" --from $vendor_account --yes --description "$description")
check_response "$result" "\"success\": true"
echo "$result"

echo "Get Model with VID: ${vid} PID: ${pid}"
result=$(dclcli query modelinfo model --vid=$vid --pid=$pid)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"description\": \"$description\""
echo "$result"

echo "Update Model with VID: ${vid} PID: ${pid} modifying support-url"
support_url="https://newsupporturl.test"
echo dclcli tx modelinfo update-model --vid=$vid --pid=$pid --cd-version-number="32" --from $vendor_account --yes --support-url "$support_url"
result=$(echo "test1234" | dclcli tx modelinfo update-model --vid=$vid --pid=$pid --cd-version-number="33" --from $vendor_account --yes --support-url "$support_url")
check_response "$result" "\"success\": true"
echo "$result"

echo "Get Model with VID: ${vid} PID: ${pid}"
result=$(dclcli query modelinfo model --vid=$vid --pid=$pid)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"support_url\": \"$support_url\""
echo "$result"
