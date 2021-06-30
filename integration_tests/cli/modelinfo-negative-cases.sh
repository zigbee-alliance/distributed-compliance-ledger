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

echo "Create regular account"
create_new_account user_account ""

echo "Create Vendor account"
create_new_account vendor_account "Vendor"

# Body

# Ledger side errors

vid=$RANDOM
pid=$RANDOM

echo "Add Model with VID: $vid PID: $pid: Not Vendor"

result=$(echo "test1234" | dclcli tx modelinfo add-model --vid=$vid --pid=$pid --productName="Product Name" --description="Device Description" --sku="SKU12FS" --softwareVersion="10123" --softwareVersionString="1.0b123"  --hardwareVersion="5123" --hardwareVersionString="5.1.23"  --cdVersionNumber="32" --from $user_account --yes)
check_response_and_report "$result" "\"success\": false"
check_response_and_report "$result" "\"code\": 4"
echo "$result"

echo "Add Model with VID: $vid PID: $pid: Twice"
result=$(echo "test1234" | dclcli tx modelinfo add-model --vid=$vid --pid=$pid --productName="Product Name" --description="Device Description" --sku="SKU12FS" --softwareVersion="10123" --softwareVersionString="1.0b123"  --hardwareVersion="5123" --hardwareVersionString="5.1.23"  --cdVersionNumber="32" --from $vendor_account --yes)
result=$(echo "test1234" | dclcli tx modelinfo add-model --vid=$vid --pid=$pid --productName="Product Name" --description="Device Description" --sku="SKU12FS" --softwareVersion="10123" --softwareVersionString="1.0b123"  --hardwareVersion="5123" --hardwareVersionString="5.1.23"  --cdVersionNumber="32" --from $vendor_account --yes)
check_response_and_report "$result" "\"success\": false"
check_response_and_report "$result" "\"code\": 501"
echo "$result"

# CLI side errors

echo "Add Model with VID: $vid PID: $pid: Unknown account"
result=$(dclcli tx modelinfo add-model --vid=$vid --pid=$pid --productName="Product Name" --description="Device Description" --sku="SKU12FS" --softwareVersion="10123" --softwareVersionString="1.0b123"  --hardwareVersion="5123" --hardwareVersionString="5.1.23"  --cdVersionNumber="32" --from "Unknown"  2>&1) || true
check_response_and_report "$result" "Key Unknown not found"

echo "Add model with invalid VID/PID"
for i in "0" "string"
do
  result=$(echo "test1234" | dclcli tx modelinfo add-model --vid=$i --pid=$pid --productName="Product Name" --description="Device Description" --sku="SKU12FS" --softwareVersion="10123" --softwareVersionString="1.0b123"  --hardwareVersion="5123" --hardwareVersionString="5.1.23"  --cdVersionNumber="32" --from $vendor_account --yes 2>&1) || true
  check_response_and_report "$result" "Invalid VID"

  result=$(echo "test1234" | dclcli tx modelinfo add-model --vid=$vid --pid=$i --productName="Product " --description="Device Description" --sku="SKU12FS" --softwareVersion="10123" --softwareVersionString="1.0b123"  --hardwareVersion="5123" --hardwareVersionString="5.1.23"  --cdVersionNumber="32" --from $vendor_account --yes 2>&1) || true
  check_response_and_report "$result" "Invalid PID"
done

echo "Add model with empty name"
result=$(echo "test1234" | dclcli tx modelinfo add-model --vid=$vid --pid=$pid --productName="" --description="Device Description" --sku="SKU12FS" --softwareVersion="10123" --softwareVersionString="1.0b123"  --hardwareVersion="5123" --hardwareVersionString="5.1.23"  --cdVersionNumber="32" --from $vendor_account --yes 2>&1) || true
check_response_and_report "$result" "Code: 6"
check_response_and_report "$result" "Invalid ProductName"

echo "Add model with empty description"
result=$(echo "test1234" | dclcli tx modelinfo add-model --vid=$vid --pid=$pid --productName="Product Name" --description="" --sku="SKU12FS" --softwareVersion="10123" --softwareVersionString="1.0b123"  --hardwareVersion="5123" --hardwareVersionString="5.1.23"  --cdVersionNumber="32" --from $vendor_account --yes 2>&1) || true
check_response_and_report "$result" "Code: 6"
check_response_and_report "$result" "Invalid Description"

echo "Add model with empty SKU"
result=$(echo "test1234" | dclcli tx modelinfo add-model --vid=$vid --pid=$pid --productName="Product Name" --description="Description" --sku="" --softwareVersion=12 --softwareVersionString="1.0b123"  --hardwareVersion="5123" --hardwareVersionString="5.1.23"  --cdVersionNumber="32" --from $vendor_account --yes 2>&1) || true
check_response_and_report "$result" "Code: 6"
check_response_and_report "$result" "Invalid SKU"

echo "Add model with empty Software Version"
result=$(echo "test1234" | dclcli tx modelinfo add-model --vid=$vid --pid=$pid --productName="Product Name" --description="Description" --sku="SKU12FS" --softwareVersion="" --softwareVersionString="1.0b123"  --hardwareVersion="5123" --hardwareVersionString="5.1.23"  --cdVersionNumber="32" --from $vendor_account --yes 2>&1) || true
check_response_and_report "$result" "Code: 6"
check_response_and_report "$result" "Invalid SoftwareVersion"

echo "Add model with empty Software Version String"
result=$(echo "test1234" | dclcli tx modelinfo add-model --vid=$vid --pid=$pid --productName="Product Name" --description="Description" --sku="SKU12FS" --softwareVersion="1" --softwareVersionString=""  --hardwareVersion="5123" --hardwareVersionString="5.1.23"  --cdVersionNumber="32" --from $vendor_account --yes 2>&1) || true
check_response_and_report "$result" "Code: 6"
check_response_and_report "$result" "Invalid SoftwareVersionString"

echo "Add model with empty Hardware Version"
result=$(echo "test1234" | dclcli tx modelinfo add-model --vid=$vid --pid=$pid --productName="Product Name" --description="Description" --sku="SKU12FS" --softwareVersion="1" --softwareVersionString="1.0b123"  --hardwareVersion="" --hardwareVersionString="5.1.23"  --cdVersionNumber="32" --from $vendor_account --yes 2>&1) || true
check_response_and_report "$result" "Code: 6"
check_response_and_report "$result" "Invalid HardwareVersion"

echo "Add model with empty Hardware Version String"
result=$(echo "test1234" | dclcli tx modelinfo add-model --vid=$vid --pid=$pid --productName="Product Name" --description="Description" --sku="SKU12FS" --softwareVersion="1" --softwareVersionString="1.0b123"  --hardwareVersion="5123" --hardwareVersionString=""  --cdVersionNumber="32" --from $vendor_account --yes 2>&1) || true
check_response_and_report "$result" "Code: 6"
check_response_and_report "$result" "Invalid HardwareVersionString"

echo "Add model with empty CD Version Number"
result=$(echo "test1234" | dclcli tx modelinfo add-model --vid=$vid --pid=$pid --productName="Product Name" --description="Description" --sku="SKU12FS" --softwareVersion="1" --softwareVersionString="1.0b123"  --hardwareVersion="5123" --hardwareVersionString="51.1235.3"  --cdVersionNumber="" --from $vendor_account --yes 2>&1) || true
check_response_and_report "$result" "Code: 6"
check_response_and_report "$result" "Invalid CDVersionNumber"

echo "Add model without --from flag"
result=$(echo "test1234" | dclcli tx modelinfo add-model --vid=$vid --pid=$pid --productName="Product Name" --description="Device Description" --sku="SKU12FS" --softwareVersion="10123" --softwareVersionString="1.0b123"  --hardwareVersion="5123" --hardwareVersionString="5.1.23"  --cdVersionNumber="32"  --yes 2>&1) || true
check_response_and_report "$result" "required flag(s) \"from\" not set"

echo "Add model without enough parameters"
result=$(echo "test1234" | dclcli tx modelinfo add-model --vid=$vid --pid=$pid --productName="Product Name"  --sku="SKU12FS" --softwareVersion="10123" --softwareVersionString="1.0b123"  --hardwareVersion="5123" --hardwareVersionString="5.1.23"  --cdVersionNumber="32" --from $vendor_account --yes 2>&1) || true
check_response_and_report "$result" "required flag(s) \"description\" not set"
