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
create_new_account certification_house "CertificationCenter"

vid=$RANDOM
pid=$RANDOM
softwareVersionString=$RANDOM
vendor_account=vendor_account_$vid
echo "Create Vendor account - $vendor_account"
create_new_vendor_account $vendor_account $vid

test_divider

echo "Create CertificationCenter account"
create_new_account zb_account "CertificationCenter"

((vid_with_pids=vid + 1))
pid_ranges="1-100"
vendor_account_with_pids=vendor_account_$vid_with_pids
echo "Create Vendor account - $vid_with_pids with ProductIDs - $pid_ranges"
create_new_vendor_account $vendor_account_with_pids $vid_with_pids $pid_ranges

# Body

# Ledger side errors

echo "Add Model with VID: $vid PID: $pid: Not Vendor"
result=$(echo "test1234" | dcld tx model add-model --vid=$vid --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0 --from=$certification_house --yes)
result=$(get_txn_result "$result")
check_response_and_report "$result" "\"code\": 4"
echo "$result"

echo "Add Model with VID: $vid_with_pids PID: 101 :Vendor with non-associated PID"
result=$(echo "test1234" | dcld tx model add-model --vid=$vid_with_pids --pid=101 --deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0 --from=$vendor_account_with_pids --yes)
result=$(get_txn_result "$result")
check_response_and_report "$result" "\"code\": 4"
echo "$result"

test_divider

vid1=$RANDOM
echo "Add Model with VID: $vid1 PID: $pid: Vendor ID does not belong to vendor"
result=$(echo "test1234" | dcld tx model add-model --vid=$vid1 --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0 --from=$vendor_account --yes)
result=$(get_txn_result "$result")
check_response_and_report "$result" "\"code\": 4"
echo "$result"

test_divider

echo "Add Model with VID: $vid PID: $pid: Twice"
result=$(echo "test1234" | dcld tx model add-model --vid=$vid --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0 --from=$vendor_account --yes)
result=$(echo "test1234" | dcld tx model add-model --vid=$vid --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0 --from=$vendor_account --yes)
result=$(get_txn_result "$result")
check_response_and_report "$result" "\"code\": 501"
echo "$result"

test_divider

sv=$RANDOM
svs=$RANDOM
echo "Create a Device Model Version with VID: $vid PID: $pid SV: $sv"
result=$(echo 'test1234' | dcld tx model add-model-version --cdVersionNumber=1 --maxApplicableSoftwareVersion=10 --minApplicableSoftwareVersion=1 --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$softwareVersionString --from=$vendor_account --yes)
result=$(get_txn_result "$result")
echo "$result"
check_response "$result" "\"code\": 0"

certification_date="2020-01-01T00:00:01Z"
zigbee_certification_type="zigbee"
matter_certification_type="matter"
cd_certificate_id="123"
result=$(echo 'test1234' | dcld tx compliance certify-model --vid=$vid --pid=$pid --softwareVersion=$sv --cdVersionNumber=1 --certificationType="$zigbee_certification_type" --certificationDate="$certification_date" --softwareVersionString=$softwareVersionString --cdCertificateId="$cd_certificate_id" --from $zb_account --yes)
result=$(get_txn_result "$result")
echo "$result"

echo "Delete Model with VID: ${vid} PID: ${pid}"
result=$(dcld tx model delete-model --vid=$vid --pid=$pid --from=$vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 525" # code for model certified error


# CLI side errors

echo "Add Model with VID: $vid PID: $pid: Unknown account"
result=$(echo "test1234" | dcld tx model add-model --vid=$vid --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0 --from=$vendor_account --yes)
result=$(dcld tx model add-model --vid=$vid --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0 --from "Unknown"  2>&1) || true
result=$(get_txn_result "$result")
check_response_and_report "$result" "key not found" raw

test_divider

echo "Add model with invalid VID/PID"
i="-1" 
result=$(echo "test1234" | dcld tx model add-model --vid=$i --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0 --from $vendor_account --yes 2>&1) || true
result=$(get_txn_result "$result")
check_response_and_report "$result" "Vid must not be less than 1" raw

result=$(echo "test1234" | dcld tx model add-model --vid=$vid --pid=$i --deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0 --from $vendor_account --yes 2>&1) || true
result=$(get_txn_result "$result")
heck_response_and_report "$result" "Pid must not be less than 1" raw

i="0" 
result=$(echo "test1234" | dcld tx model add-model --vid=$i --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0 --from $vendor_account --yes 2>&1) || true
result=$(get_txn_result "$result")
check_response_and_report "$result" "Vid must not be less than 1" raw

result=$(echo "test1234" | dcld tx model add-model --vid=$vid --pid=$i --deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0 --from $vendor_account --yes 2>&1) || true
result=$(get_txn_result "$result")
check_response_and_report "$result" "Pid must not be less than 1" raw

i="65536"
result=$(echo "test1234" | dcld tx model add-model --vid=$i --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0 --from $vendor_account --yes 2>&1) || true
result=$(get_txn_result "$result")
check_response_and_report "$result" "Vid must not be greater than 65535" raw

result=$(echo "test1234" | dcld tx model add-model --vid=$vid --pid=$i --deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0 --from $vendor_account --yes 2>&1) || true
result=$(get_txn_result "$result")
check_response_and_report "$result" "Pid must not be greater than 65535" raw

i="string" 
result=$(echo "test1234" | dcld tx model add-model --vid=$i --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0 --from $vendor_account --yes 2>&1) || true
result=$(get_txn_result "$result")
check_response_and_report "$result" "invalid syntax" raw

result=$(echo "test1234" | dcld tx model add-model --vid=$vid --pid=$i --deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0 --from $vendor_account --yes 2>&1) || true
result=$(get_txn_result "$result")
check_response_and_report "$result" "invalid syntax" raw

test_divider

echo "Add model with empty name"
result=$(echo "test1234" | dcld tx model add-model --vid=$vid --pid=$pid --deviceTypeID=1 --productName="" --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0  --from $vendor_account --yes 2>&1) || true
result=$(get_txn_result "$result")
check_response_and_report "$result" "ProductName is a required field" raw

test_divider

echo "Add model with empty --from flag"
result=$(dcld tx model add-model --vid=$vid --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0  --from "" --yes 2>&1) || true
result=$(get_txn_result "$result")
check_response_and_report "$result" "invalid creator address (empty address string is not allowed)" raw

test_divider

echo "Add model without --from flag"
result=$(dcld tx model add-model --vid=$vid --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0  --yes 2>&1) || true
result=$(get_txn_result "$result")
check_response_and_report "$result" "required flag(s) \"from\" not set" raw

test_divider

echo "Update model with Non Mutable fields" 
pid=$RANDOM
sv=$RANDOM
svs=$RANDOM
create_model_and_version $vid $pid $sv $svs $vendor_account 
result=$(dcld query model get-model --vid=$vid --pid=$pid)
echo "$result"
