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
upgrade_plan_name_v1_4_0="v1.4.0"
upgrade_plan_info_v1_4_0="{\"binaries\":{\"linux/amd64\":\"https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v1.4.0/dcld?checksum=sha256:e4031c6a77aa8e58add391be671a334613271bcf6e7f11d23b04a0881ece6958\"}}"
upgrade_plan_name_v1_4_1="v1.4.1"
upgrade_plan_info_v1_4_1="{\"binaries\":{\"linux/amd64\":\"https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v1.4.1/dcld?checksum=sha256:e4031c6a77aa8e58add391be671a334613271bcf6e7f11d23b04a0881ece6958\"}}"

echo "propose and approve upgrade"
echo "Create Trustee account"
create_new_account trustee_account "Trustee"

propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_plan_name_v1_2_0 --upgrade-height=$upgrade_height --upgrade-info=$upgrade_plan_info_v1_2_0 --from $trustee_account --yes)
propose=$(get_txn_result "$propose")
echo "propose upgrade response: $propose"
check_response "$propose" "\"code\": 0"

proposed_dclupgrade_query=$(dcld query dclupgrade proposed-upgrade --name=$upgrade_plan_name_v1_2_0)
echo "dclupgrade proposed upgrade query: $proposed_dclupgrade_query"
check_response_and_report "$proposed_dclupgrade_query" "\"name\": \"$upgrade_plan_name_v1_2_0\""
check_response_and_report "$proposed_dclupgrade_query" "\"height\": \"$upgrade_height\""
check_upgrade_plan_info "$proposed_dclupgrade_query" "$upgrade_plan_info_v1_2_0"

approve=$(dcld tx dclupgrade approve-upgrade --name=$upgrade_plan_name_v1_2_0 --from alice --yes)
approve=$(get_txn_result "$approve")
echo "approve upgrade response: $approve"
check_response "$approve" "\"code\": 0"

reject=$(dcld tx dclupgrade reject-upgrade --name=$upgrade_plan_name_v1_2_0 --from alice --yes)
reject=$(get_txn_result "$reject")
echo "reject upgrade response: $reject"
check_response "$reject" "\"code\": 0"

approve=$(dcld tx dclupgrade approve-upgrade --name=$upgrade_plan_name_v1_2_0 --from alice --yes)
approve=$(get_txn_result "$approve")
echo "approve upgrade response: $approve"
check_response "$approve" "\"code\": 0"

proposed_dclupgrade_query=$(dcld query dclupgrade proposed-upgrade --name=$upgrade_plan_name_v1_2_0)
echo "dclupgrade proposed upgrade query: $proposed_dclupgrade_query"
check_response_and_report "$proposed_dclupgrade_query" "\"name\": \"$upgrade_plan_name_v1_2_0\""
check_response_and_report "$proposed_dclupgrade_query" "\"height\": \"$upgrade_height\""
check_upgrade_plan_info "$proposed_dclupgrade_query" "$upgrade_plan_info_v1_2_0"

approved_dclupgrade_query=$(dcld query dclupgrade approved-upgrade --name=$upgrade_plan_name_v1_2_0)
echo "dclupgrade approved upgrade query: $approved_dclupgrade_query"
check_response "$approved_dclupgrade_query" "Not Found"

plan_query=$(dcld query upgrade plan 2>&1) || true
check_response "$plan_query" "no upgrade scheduled" raw

approve=$(dcld tx dclupgrade approve-upgrade --name=$upgrade_plan_name_v1_2_0 --from bob --yes)
approve=$(get_txn_result "$approve")
echo "approve upgrade response: $approve"
check_response "$approve" "\"code\": 0"

plan_query=$(dcld query upgrade plan)
echo "plan query: $plan_query"
check_response_and_report "$plan_query" "\"name\": \"$upgrade_plan_name_v1_2_0\""
check_response_and_report "$plan_query" "\"height\": \"$upgrade_height\""
check_upgrade_plan_info "$plan_query" "$upgrade_plan_info_v1_2_0"

approved_dclupgrade_query=$(dcld query dclupgrade approved-upgrade --name=$upgrade_plan_name_v1_2_0)
echo "dclupgrade approved upgrade query: $approved_dclupgrade_query"
check_response "$approved_dclupgrade_query" "\"name\": \"$upgrade_plan_name_v1_2_0\""
check_response "$approved_dclupgrade_query" "\"height\": \"$upgrade_height\""
check_upgrade_plan_info "$approved_dclupgrade_query" "$upgrade_plan_info_v1_2_0"

proposed_dclupgrade_query=$(dcld query dclupgrade proposed-upgrade --name=$upgrade_plan_name_v1_2_0)
echo "dclupgrade proposed upgrade query: $proposed_dclupgrade_query"
check_response "$proposed_dclupgrade_query" "Not Found" raw


test_divider

random_string upgrade_name

echo "proposer cannot approve upgrade"
propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_name --upgrade-height=$upgrade_height --from alice --yes)
propose=$(get_txn_result "$propose")
echo "propose upgrade response: $propose"
check_response "$propose" "\"code\": 0"

approve=$(dcld tx dclupgrade approve-upgrade --name=$upgrade_name --from alice --yes)
approve=$(get_txn_result "$approve")
echo "approve upgrade response: $approve"
check_response_and_report "$approve" "unauthorized" raw


test_divider

random_string upgrade_name
echo "cannot approve upgrade twice"
propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_name --upgrade-height=$upgrade_height --from alice --yes)
propose=$(get_txn_result "$propose")
echo "propose upgrade response: $propose"
check_response "$propose" "\"code\": 0"

approve=$(dcld tx dclupgrade approve-upgrade --name=$upgrade_name --from bob --yes)
approve=$(get_txn_result "$approve")
echo "approve upgrade response: $approve"
check_response "$approve" "\"code\": 0"

second_approve=$(dcld tx dclupgrade approve-upgrade --name=$upgrade_name --from bob --yes)
second_approve=$(get_txn_result "$second_approve")
echo "second approve upgrade response: $second_approve"
check_response_and_report "$second_approve" "unauthorized" raw

test_divider

random_string upgrade_name
echo "cannot propose upgrade twice"
propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_name --upgrade-height=$upgrade_height --from alice --yes)
propose=$(get_txn_result "$propose")
echo "propose upgrade response: $propose"
check_response "$propose" "\"code\": 0"

second_propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_name --upgrade-height=$upgrade_height --from alice --yes)
second_propose=$(get_txn_result "$second_propose")
echo "second propose upgrade response: $second_propose"
check_response_and_report "$second_propose" "proposed upgrade already exists" raw


test_divider

random_string upgrade_name
echo "upgrade height < current height"
height=1
propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_name --upgrade-height=$height --from alice --yes)
propose=$(get_txn_result "$propose")
echo "propose upgrade response: $propose"
check_response_and_report "$propose" "upgrade cannot be scheduled in the past" raw


test_divider

echo "propose and reject upgrade"
upgrade_height=$(($RANDOM + 10000000))
propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_plan_name_v1_2_1 --upgrade-height=$upgrade_height --upgrade-info=$upgrade_plan_info_v1_2_1 --from $trustee_account --yes)
propose=$(get_txn_result "$propose")
echo "propose upgrade response: $propose"
check_response "$propose" "\"code\": 0"

approve=$(dcld tx dclupgrade approve-upgrade --name=$upgrade_plan_name_v1_2_1 --from alice --yes)
approve=$(get_txn_result "$approve")
echo "approve upgrade response: $approve"
check_response "$approve" "\"code\": 0"

proposed_dclupgrade_query=$(dcld query dclupgrade proposed-upgrade --name=$upgrade_plan_name_v1_2_1)
echo "dclupgrade proposed upgrade query: $proposed_dclupgrade_query"
check_response_and_report "$proposed_dclupgrade_query" "\"name\": \"$upgrade_plan_name_v1_2_1\""
check_response_and_report "$proposed_dclupgrade_query" "\"height\": \"$upgrade_height\""
check_upgrade_plan_info "$proposed_dclupgrade_query" "$upgrade_plan_info_v1_2_1"

reject=$(dcld tx dclupgrade reject-upgrade --name=$upgrade_plan_name_v1_2_1 --from $trustee_account --yes)
reject=$(get_txn_result "$reject")
echo "reject upgrade response: $reject"
check_response "$reject" "\"code\": 0"

approve=$(dcld tx dclupgrade approve-upgrade --name=$upgrade_plan_name_v1_2_1 --from $trustee_account --yes)
approve=$(get_txn_result "$approve")
echo "approve upgrade response: $approve"
check_response "$approve" "\"code\": 0"

reject=$(dcld tx dclupgrade reject-upgrade --name=$upgrade_plan_name_v1_2_1 --from alice --yes)
reject=$(get_txn_result "$reject")
echo "reject upgrade response: $reject"
check_response "$reject" "\"code\": 0"

# second_reject=$(dcld tx dclupgrade reject-upgrade ---name=$upgrade_plan_name_v1_2_1 --from alice --yes)
# second_reject=$(get_txn_result "$second_reject")
# echo "second_reject upgrade response: $reject"
# response_does_not_contain "$second_reject" "\"code\": 0"

proposed_dclupgrade_query=$(dcld query dclupgrade proposed-upgrade --name=$upgrade_plan_name_v1_2_1)
echo "dclupgrade proposed upgrade query: $proposed_dclupgrade_query"
check_response_and_report "$proposed_dclupgrade_query" "\"name\": \"$upgrade_plan_name_v1_2_1\""
check_response_and_report "$proposed_dclupgrade_query" "\"height\": \"$upgrade_height\""
check_upgrade_plan_info "$proposed_dclupgrade_query" "$upgrade_plan_info_v1_2_1"

rejected_dclupgrade_query=$(dcld query dclupgrade rejected-upgrade --name=$upgrade_plan_name_v1_2_1)
echo "dclupgrade rejected upgrade query: $rejected_dclupgrade_query"
check_response "$rejected_dclupgrade_query" "Not Found"

approved_dclupgrade_query=$(dcld query dclupgrade approved-upgrade --name=$upgrade_plan_name_v1_2_1)
echo "dclupgrade approved upgrade query: $approved_dclupgrade_query"
check_response "$approved_dclupgrade_query" "Not Found"

reject=$(dcld tx dclupgrade reject-upgrade --name=$upgrade_plan_name_v1_2_1 --from bob --yes)
reject=$(get_txn_result "$reject")
echo "reject upgrade response: $reject"
check_response "$reject" "\"code\": 0"

rejected_dclupgrade_query=$(dcld query dclupgrade rejected-upgrade --name=$upgrade_plan_name_v1_2_1)
echo "dclupgrade rejected upgrade query: $rejected_dclupgrade_query"
check_response_and_report "$rejected_dclupgrade_query" "\"name\": \"$upgrade_plan_name_v1_2_1\""
check_response_and_report "$rejected_dclupgrade_query" "\"height\": \"$upgrade_height\""
check_upgrade_plan_info "$rejected_dclupgrade_query" "$upgrade_plan_info_v1_2_1"

proposed_dclupgrade_query=$(dcld query dclupgrade proposed-upgrade --name=$upgrade_plan_name_v1_2_1)
echo "dclupgrade rejected upgrade query: $proposed_dclupgrade_query"
check_response "$proposed_dclupgrade_query" "Not Found"

approved_dclupgrade_query=$(dcld query dclupgrade approved-upgrade --name=$upgrade_plan_name_v1_2_1)
echo "dclupgrade rejected upgrade query: $approved_dclupgrade_query"
check_response "$approved_dclupgrade_query" "Not Found"

# APPROVE UPGRADE'S PLAN HEIGHT LESS THAN BLOCK HEIGHT AND RE-PROPOSE UPGRADE PLAN HEIGHT BIGGER THAN BLOCK HEIGHT
###########################################################################################################################################
get_height current_height
echo "Current height is $current_height"
plan_height=$(expr $current_height + 3)
echo "$plan_height"

echo "propose upgrade's plan height bigger than block height"
propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_plan_name_v1_2_2 --upgrade-height=$plan_height --upgrade-info=$upgrade_plan_info_v1_2_2 --from jack --yes)
propose=$(get_txn_result "$propose")
echo "propose upgrade response: $propose"
check_response "$propose" "\"code\": 0"

proposed_dclupgrade_query=$(dcld query dclupgrade proposed-upgrade --name=$upgrade_plan_name_v1_2_2)
echo "dclupgrade proposed upgrade query: $proposed_dclupgrade_query"
check_response_and_report "$proposed_dclupgrade_query" "\"name\": \"$upgrade_plan_name_v1_2_2\""
check_response_and_report "$proposed_dclupgrade_query" "\"height\": \"$plan_height\""
check_upgrade_plan_info "$proposed_dclupgrade_query" "$upgrade_plan_info_v1_2_2"

wait_for_height $(expr $plan_height + 5) 300 outage-safe
get_height current_height
echo "Current height is $current_height"

echo "approve upgrade's plan height less than block height"
approve=$(dcld tx dclupgrade approve-upgrade --name=$upgrade_plan_name_v1_2_2 --from alice --yes)
approve=$(get_txn_result "$approve")
echo "approve upgrade response: $approve"
check_response_and_report "$approve" "upgrade cannot be scheduled in the past" raw

get_height current_height
echo "Current height is $current_height"
plan_height=$(expr $current_height + 3)

echo "re-propose upgrade's plan height bigger than block height"
propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_plan_name_v1_2_2 --upgrade-height=$plan_height --upgrade-info=$upgrade_plan_info_v1_2_2 --from jack --yes)
propose=$(get_txn_result "$propose")
echo "propose upgrade response: $propose"
check_response "$propose" "\"code\": 0"

proposed_dclupgrade_query=$(dcld query dclupgrade proposed-upgrade --name=$upgrade_plan_name_v1_2_2)
echo "dclupgrade proposed upgrade query: $proposed_dclupgrade_query"
check_response_and_report "$proposed_dclupgrade_query" "\"name\": \"$upgrade_plan_name_v1_2_2\""
check_response_and_report "$proposed_dclupgrade_query" "\"height\": \"$plan_height\""
check_upgrade_plan_info "$proposed_dclupgrade_query" "$upgrade_plan_info_v1_2_2"
###########################################################################################################################################

test_divider

# APPROVE UPGRADE'S PLAN HEIGHT LESS THAN BLOCK HEIGHT AND RE-PROPOSE UPGRADE PLAN HEIGHT LESS THAN BLOCK HEIGHT
###########################################################################################################################################
get_height current_height
echo "Current height is $current_height"
plan_height=$(expr $current_height + 3)

echo "propose upgrade's plan height bigger than block height"
propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_plan_name_v1_4_0 --upgrade-height=$plan_height --upgrade-info=$upgrade_plan_info_v1_4_0 --from jack --yes)
propose=$(get_txn_result "$propose")
echo "propose upgrade response: $propose"
check_response "$propose" "\"code\": 0"

proposed_dclupgrade_query=$(dcld query dclupgrade proposed-upgrade --name=$upgrade_plan_name_v1_4_0)
echo "dclupgrade proposed upgrade query: $proposed_dclupgrade_query"
check_response_and_report "$proposed_dclupgrade_query" "\"name\": \"$upgrade_plan_name_v1_4_0\""
check_response_and_report "$proposed_dclupgrade_query" "\"height\": \"$plan_height\""
check_upgrade_plan_info "$proposed_dclupgrade_query" "$upgrade_plan_info_v1_4_0"

wait_for_height $(expr $plan_height + 5) 300 outage-safe
get_height current_height
echo "Current height is $current_height"

echo "approve upgrade's plan height less than block height"
approve=$(dcld tx dclupgrade approve-upgrade --name=$upgrade_plan_name_v1_4_0 --from alice --yes)
approve=$(get_txn_result "$approve")
echo "approve upgrade response: $approve"
check_response_and_report "$approve" "upgrade cannot be scheduled in the past" raw

echo "re-propose upgrade's plan height bigger than block height"
propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_plan_name_v1_4_0 --upgrade-height=$plan_height --upgrade-info=$upgrade_plan_info_v1_4_0 --from jack --yes)
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
result=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_plan_name_v1_4_1 --upgrade-height=$plan_height --upgrade-info=$upgrade_plan_info_v1_4_1 --from jack --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "jack (Trustee) rejects upgrade"
result=$(dcld tx dclupgrade reject-upgrade --name=$upgrade_plan_name_v1_4_1 --from jack --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Upgrade not found in proposed upgrade"
result=$(dcld query dclupgrade proposed-upgrade --name=$upgrade_plan_name_v1_4_1)
echo $result | jq
check_response "$result" "Not Found"

test_divider

echo "Upgrade not found in rejected upgrade"
result=$(dcld query dclupgrade rejected-upgrade --name=$upgrade_plan_name_v1_4_1)
echo $result | jq
check_response "$result" "Not Found"

test_divider

echo "Upgrade not found in approved upgrade query"
result=$(dcld query dclupgrade approved-upgrade --name=$upgrade_plan_name_v1_4_1)
echo $result | jq
check_response "$result" "Not Found"

echo "PASSED"