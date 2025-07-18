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

upgrade_height=$(($RANDOM + 10000000))

upgrade_plan_name_v1_2_0="v1.2.0"
upgrade_plan_info_v1_2_0="{\"binaries\":{\"linux/amd64\":\"https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v1.2.0/dcld?checksum=sha256:e4031c6a77aa8e58add391be671a334613271bcf6e7f11d23b04a0881ece6958\"}}"
upgrade_plan_name_v1_2_1="v1.2.1"
upgrade_plan_info_v1_2_1="{\"binaries\":{\"linux/amd64\":\"https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v1.2.1/dcld?checksum=sha256:e4031c6a77aa8e58add391be671a334613271bcf6e7f11d23b04a0881ece6958\"}}"
upgrade_plan_name_v1_2_2="v1.2.2"
upgrade_plan_info_v1_2_2="{\"binaries\":{\"linux/amd64\":\"https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v1.2.2/dcld?checksum=sha256:e4031c6a77aa8e58add391be671a334613271bcf6e7f11d23b04a0881ece6958\"}}"

# APPROVE UPGRADE'S PLAN HEIGHT LESS THAN BLOCK HEIGHT AND RE-PROPOSE UPGRADE PLAN HEIGHT BIGGER THAN BLOCK HEIGHT
###########################################################################################################################################
get_height current_height
echo "Current height is $current_height"
plan_height=$(expr $current_height + 3)
echo "$plan_height"

echo "propose upgrade's plan height bigger than block height"
propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_plan_name_v1_2_0 --upgrade-height=$plan_height --upgrade-info=$upgrade_plan_info_v1_2_0 --from jack --yes)
propose=$(get_txn_result "$propose")
echo "propose upgrade response: $propose"
check_response "$propose" "\"code\": 0"

proposed_dclupgrade_query=$(dcld query dclupgrade proposed-upgrade --name=$upgrade_plan_name_v1_2_0)
echo "dclupgrade proposed upgrade query: $proposed_dclupgrade_query"
check_response_and_report "$proposed_dclupgrade_query" "\"name\": \"$upgrade_plan_name_v1_2_0\""
check_response_and_report "$proposed_dclupgrade_query" "\"height\": \"$plan_height\""
check_upgrade_plan_info "$proposed_dclupgrade_query" "$upgrade_plan_info_v1_2_0"

wait_for_height $(expr $plan_height + 5) 300 outage-safe
get_height current_height
echo "Current height is $current_height"

echo "approve upgrade's plan height less than block height"
approve=$(dcld tx dclupgrade approve-upgrade --name=$upgrade_plan_name_v1_2_0 --from alice --yes)
approve=$(get_txn_result "$approve")
echo "approve upgrade response: $approve"
check_response_and_report "$approve" "upgrade cannot be scheduled in the past" raw

get_height current_height
echo "Current height is $current_height"
plan_height=$(expr $current_height + 3)

echo "re-propose upgrade's plan height bigger than block height"
propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_plan_name_v1_2_0 --upgrade-height=$plan_height --upgrade-info=$upgrade_plan_info_v1_2_0 --from jack --yes)
propose=$(get_txn_result "$propose")
echo "propose upgrade response: $propose"
check_response "$propose" "\"code\": 0"

proposed_dclupgrade_query=$(dcld query dclupgrade proposed-upgrade --name=$upgrade_plan_name_v1_2_0)
echo "dclupgrade proposed upgrade query: $proposed_dclupgrade_query"
check_response_and_report "$proposed_dclupgrade_query" "\"name\": \"$upgrade_plan_name_v1_2_0\""
check_response_and_report "$proposed_dclupgrade_query" "\"height\": \"$plan_height\""
check_upgrade_plan_info "$proposed_dclupgrade_query" "$upgrade_plan_info_v1_2_0"
###########################################################################################################################################

test_divider

# APPROVE UPGRADE'S PLAN HEIGHT LESS THAN BLOCK HEIGHT AND RE-PROPOSE UPGRADE PLAN HEIGHT LESS THAN BLOCK HEIGHT
###########################################################################################################################################
get_height current_height
echo "Current height is $current_height"
plan_height=$(expr $current_height + 3)

echo "propose upgrade's plan height bigger than block height"
propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_plan_name_v1_2_1 --upgrade-height=$plan_height --upgrade-info=$upgrade_plan_info_v1_2_1 --from jack --yes)
propose=$(get_txn_result "$propose")
echo "propose upgrade response: $propose"
check_response "$propose" "\"code\": 0"

proposed_dclupgrade_query=$(dcld query dclupgrade proposed-upgrade --name=$upgrade_plan_name_v1_2_1)
echo "dclupgrade proposed upgrade query: $proposed_dclupgrade_query"
check_response_and_report "$proposed_dclupgrade_query" "\"name\": \"$upgrade_plan_name_v1_2_1\""
check_response_and_report "$proposed_dclupgrade_query" "\"height\": \"$plan_height\""
check_upgrade_plan_info "$proposed_dclupgrade_query" "$upgrade_plan_info_v1_2_1"

wait_for_height $(expr $plan_height + 5) 300 outage-safe
get_height current_height
echo "Current height is $current_height"

echo "approve upgrade's plan height less than block height"
approve=$(dcld tx dclupgrade approve-upgrade --name=$upgrade_plan_name_v1_2_1 --from alice --yes)
approve=$(get_txn_result "$approve")
echo "approve upgrade response: $approve"
check_response_and_report "$approve" "upgrade cannot be scheduled in the past" raw

echo "re-propose upgrade's plan height bigger than block height"
propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_plan_name_v1_2_1 --upgrade-height=$plan_height --upgrade-info=$upgrade_plan_info_v1_2_1 --from jack --yes)
propose=$(get_txn_result "$propose")
echo "propose upgrade response: $propose"
check_response_and_report "$propose" "upgrade cannot be scheduled in the past" raw
###########################################################################################################################################

test_divider
get_height current_height
echo "Current height is $current_height"
plan_height=$(expr $current_height + 3)

# 18. TEST PROPOSE AND REJECT UPGRADE
echo "18. TEST PROPOSE AND REJECT UPGRADE"
test_divider

echo "jack (Trustee) propose upgrade"
result=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_plan_name_v1_2_2 --upgrade-height=$plan_height --upgrade-info=$upgrade_plan_info_v1_2_2 --from jack --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "jack (Trustee) rejects upgrade"
result=$(dcld tx dclupgrade reject-upgrade --name=$upgrade_plan_name_v1_2_2 --from jack --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Upgrade not found in proposed upgrade"
result=$(dcld query dclupgrade proposed-upgrade --name=$upgrade_plan_name_v1_2_2)
echo $result | jq
check_response "$result" "Not Found"

test_divider

echo "Upgrade not found in rejected upgrade"
result=$(dcld query dclupgrade rejected-upgrade --name=$upgrade_plan_name_v1_2_2)
echo $result | jq
check_response "$result" "Not Found"

test_divider

echo "Upgrade not found in approved upgrade query"
result=$(dcld query dclupgrade approved-upgrade --name=$upgrade_plan_name_v1_2_2)
echo $result | jq
check_response "$result" "Not Found"

echo "PASSED"