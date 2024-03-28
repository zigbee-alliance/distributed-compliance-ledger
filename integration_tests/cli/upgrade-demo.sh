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
random_string upgrade_info

echo "propose and approve upgrade"
echo "Create Trustee account"
create_new_account trustee_account "Trustee"
random_string upgrade_name
propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_name --upgrade-height=$upgrade_height --upgrade-info=$upgrade_info --from $trustee_account --yes)
propose=$(get_txn_result "$propose")
echo "propose upgrade response: $propose"
check_response "$propose" "\"code\": 0"

proposed_dclupgrade_query=$(dcld query dclupgrade proposed-upgrade --name=$upgrade_name)
echo "dclupgrade proposed upgrade query: $proposed_dclupgrade_query"
check_response_and_report "$proposed_dclupgrade_query" "\"name\": \"$upgrade_name\""
check_response_and_report "$proposed_dclupgrade_query" "\"height\": \"$upgrade_height\""
check_response_and_report "$proposed_dclupgrade_query" "\"info\": \"$upgrade_info\""

approve=$(dcld tx dclupgrade approve-upgrade --name=$upgrade_name --from alice --yes)
approve=$(get_txn_result "$approve")
echo "approve upgrade response: $approve"
check_response "$approve" "\"code\": 0"

reject=$(dcld tx dclupgrade reject-upgrade --name=$upgrade_name --from alice --yes)
reject=$(get_txn_result "$reject")
echo "reject upgrade response: $reject"
check_response "$reject" "\"code\": 0"

approve=$(dcld tx dclupgrade approve-upgrade --name=$upgrade_name --from alice --yes)
approve=$(get_txn_result "$approve")
echo "approve upgrade response: $approve"
check_response "$approve" "\"code\": 0"

proposed_dclupgrade_query=$(dcld query dclupgrade proposed-upgrade --name=$upgrade_name)
echo "dclupgrade proposed upgrade query: $proposed_dclupgrade_query"
check_response_and_report "$proposed_dclupgrade_query" "\"name\": \"$upgrade_name\""
check_response_and_report "$proposed_dclupgrade_query" "\"height\": \"$upgrade_height\""
check_response_and_report "$proposed_dclupgrade_query" "\"info\": \"$upgrade_info\""

approved_dclupgrade_query=$(dcld query dclupgrade approved-upgrade --name=$upgrade_name)
echo "dclupgrade approved upgrade query: $approved_dclupgrade_query"
check_response "$approved_dclupgrade_query" "Not Found"

plan_query=$(dcld query upgrade plan 2>&1) || true
check_response "$plan_query" "no upgrade scheduled" raw

approve=$(dcld tx dclupgrade approve-upgrade --name=$upgrade_name --from bob --yes)
approve=$(get_txn_result "$approve")
echo "approve upgrade response: $approve"
check_response "$approve" "\"code\": 0"

plan_query=$(dcld query upgrade plan)
echo "plan query: $plan_query"
check_response_and_report "$plan_query" "\"name\": \"$upgrade_name\""
check_response_and_report "$plan_query" "\"height\": \"$upgrade_height\""
check_response_and_report "$plan_query" "\"info\": \"$upgrade_info\""

approved_dclupgrade_query=$(dcld query dclupgrade approved-upgrade --name=$upgrade_name)
echo "dclupgrade approved upgrade query: $approved_dclupgrade_query"
check_response "$approved_dclupgrade_query" "\"name\": \"$upgrade_name\""
check_response "$approved_dclupgrade_query" "\"height\": \"$upgrade_height\""
check_response "$approved_dclupgrade_query" "\"info\": \"$upgrade_info\""

proposed_dclupgrade_query=$(dcld query dclupgrade proposed-upgrade --name=$upgrade_name)
echo "dclupgrade proposed upgrade query: $proposed_dclupgrade_query"
check_response "$proposed_dclupgrade_query" "Not Found" raw


test_divider


echo "proposer cannot approve upgrade"
random_string upgrade_name
propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_name --upgrade-height=$upgrade_height --upgrade-info=$upgrade_info --from alice --yes)
propose=$(get_txn_result "$propose")
echo "propose upgrade response: $propose"
check_response "$propose" "\"code\": 0"

approve=$(dcld tx dclupgrade approve-upgrade --name=$upgrade_name --from alice --yes)
approve=$(get_txn_result "$approve")
echo "approve upgrade response: $approve"
check_response_and_report "$approve" "unauthorized" raw


test_divider


echo "cannot approve upgrade twice"
random_string upgrade_name
propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_name --upgrade-height=$upgrade_height --upgrade-info=$upgrade_info --from alice --yes)
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


echo "cannot propose upgrade twice"
random_string upgrade_name
propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_name --upgrade-height=$upgrade_height --upgrade-info=$upgrade_info --from alice --yes)
propose=$(get_txn_result "$propose")
echo "propose upgrade response: $propose"
check_response "$propose" "\"code\": 0"

second_propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_name --upgrade-height=$upgrade_height --upgrade-info=$upgrade_info --from alice --yes)
second_propose=$(get_txn_result "$second_propose")
echo "second propose upgrade response: $second_propose"
check_response_and_report "$second_propose" "proposed upgrade already exists" raw


test_divider


echo "upgrade height < current height"
random_string upgrade_name
height=1
propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_name --upgrade-height=$height --upgrade-info=$upgrade_info --from alice --yes)
propose=$(get_txn_result "$propose")
echo "propose upgrade response: $propose"
check_response_and_report "$propose" "upgrade cannot be scheduled in the past" raw


test_divider

upgrade_height=$(($RANDOM + 10000000))
random_string upgrade_info

echo "propose and reject upgrade"
random_string upgrade_name
propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_name --upgrade-height=$upgrade_height --upgrade-info=$upgrade_info --from $trustee_account --yes)
propose=$(get_txn_result "$propose")
echo "propose upgrade response: $propose"
check_response "$propose" "\"code\": 0"

test_divider

approve=$(dcld tx dclupgrade approve-upgrade --name=$upgrade_name --from alice --yes)
approve=$(get_txn_result "$approve")
echo "approve upgrade response: $approve"
check_response "$approve" "\"code\": 0"

test_divider

proposed_dclupgrade_query=$(dcld query dclupgrade proposed-upgrade --name=$upgrade_name)
echo "dclupgrade proposed upgrade query: $proposed_dclupgrade_query"
check_response_and_report "$proposed_dclupgrade_query" "\"name\": \"$upgrade_name\""
check_response_and_report "$proposed_dclupgrade_query" "\"height\": \"$upgrade_height\""
check_response_and_report "$proposed_dclupgrade_query" "\"info\": \"$upgrade_info\""

test_divider

reject=$(dcld tx dclupgrade reject-upgrade --name=$upgrade_name --from $trustee_account --yes)
reject=$(get_txn_result "$reject")
echo "reject upgrade response: $reject"
check_response "$reject" "\"code\": 0"

test_divider

approve=$(dcld tx dclupgrade approve-upgrade --name=$upgrade_name --from $trustee_account --yes)
approve=$(get_txn_result "$approve")
echo "approve upgrade response: $approve"
check_response "$approve" "\"code\": 0"

test_divider

reject=$(dcld tx dclupgrade reject-upgrade --name=$upgrade_name --from alice --yes)
reject=$(get_txn_result "$reject")
echo "reject upgrade response: $reject"
check_response "$reject" "\"code\": 0"

test_divider

second_reject=$(dcld tx dclupgrade reject-upgrade --name=$upgrade_name --from alice --yes)
second_reject=$(get_txn_result "$second_reject")
echo "second_reject upgrade response: $reject"
response_does_not_contain "$second_reject" "\"code\": 0"

test_divider

proposed_dclupgrade_query=$(dcld query dclupgrade proposed-upgrade --name=$upgrade_name)
echo "dclupgrade proposed upgrade query: $proposed_dclupgrade_query"
check_response_and_report "$proposed_dclupgrade_query" "\"name\": \"$upgrade_name\""
check_response_and_report "$proposed_dclupgrade_query" "\"height\": \"$upgrade_height\""
check_response_and_report "$proposed_dclupgrade_query" "\"info\": \"$upgrade_info\""

test_divider

proposed_dclupgrade_query=$(dcld query dclupgrade proposed-upgrade --name=$upgrade_name)
echo "dclupgrade proposed upgrade query: $proposed_dclupgrade_query"
check_response_and_report "$proposed_dclupgrade_query" "\"name\": \"$upgrade_name\""
check_response_and_report "$proposed_dclupgrade_query" "\"height\": \"$upgrade_height\""
check_response_and_report "$proposed_dclupgrade_query" "\"info\": \"$upgrade_info\""

test_divider

rejected_dclupgrade_query=$(dcld query dclupgrade rejected-upgrade --name=$upgrade_name)
echo "dclupgrade rejected upgrade query: $rejected_dclupgrade_query"
check_response "$rejected_dclupgrade_query" "Not Found"

test_divider

approved_dclupgrade_query=$(dcld query dclupgrade approved-upgrade --name=$upgrade_name)
echo "dclupgrade approved upgrade query: $approved_dclupgrade_query"
check_response "$approved_dclupgrade_query" "Not Found"

test_divider

reject=$(dcld tx dclupgrade reject-upgrade --name=$upgrade_name --from bob --yes)
reject=$(get_txn_result "$reject")
echo "reject upgrade response: $reject"
check_response "$reject" "\"code\": 0"

test_divider

rejected_dclupgrade_query=$(dcld query dclupgrade rejected-upgrade --name=$upgrade_name)
echo "dclupgrade rejected upgrade query: $rejected_dclupgrade_query"
check_response_and_report "$rejected_dclupgrade_query" "\"name\": \"$upgrade_name\""
check_response_and_report "$rejected_dclupgrade_query" "\"height\": \"$upgrade_height\""
check_response_and_report "$rejected_dclupgrade_query" "\"info\": \"$upgrade_info\""

test_divider

proposed_dclupgrade_query=$(dcld query dclupgrade proposed-upgrade --name=$upgrade_name)
echo "dclupgrade rejected upgrade query: $proposed_dclupgrade_query"
check_response "$proposed_dclupgrade_query" "Not Found"

test_divider

approved_dclupgrade_query=$(dcld query dclupgrade approved-upgrade --name=$upgrade_name)
echo "dclupgrade rejected upgrade query: $approved_dclupgrade_query"
check_response "$approved_dclupgrade_query" "Not Found"

test_divider

# APPROVE UPGRADE'S PLAN HEIGHT LESS THAN BLOCK HEIGHT AND RE-PROPOSE UPGRADE PLAN HEIGHT BIGGER THAN BLOCK HEIGHT
###########################################################################################################################################
get_height current_height
echo "Current height is $current_height"
plan_height=$(expr $current_height + 3)
echo "$plan_height"
random_string upgrade_name
random_string upgrade_info

echo "propose upgrade's plan height bigger than block height"
propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_name --upgrade-height=$plan_height --upgrade-info=$upgrade_info --from jack --yes)
propose=$(get_txn_result "$propose")
echo "propose upgrade response: $propose"
check_response "$propose" "\"code\": 0"

proposed_dclupgrade_query=$(dcld query dclupgrade proposed-upgrade --name=$upgrade_name)
echo "dclupgrade proposed upgrade query: $proposed_dclupgrade_query"
check_response_and_report "$proposed_dclupgrade_query" "\"name\": \"$upgrade_name\""
check_response_and_report "$proposed_dclupgrade_query" "\"height\": \"$plan_height\""
check_response_and_report "$proposed_dclupgrade_query" "\"info\": \"$upgrade_info\""

wait_for_height $(expr $plan_height + 5) 300 outage-safe
get_height current_height
echo "Current height is $current_height"

echo "approve upgrade's plan height less than block height"
approve=$(dcld tx dclupgrade approve-upgrade --name=$upgrade_name --from alice --yes)
approve=$(get_txn_result "$approve")
echo "approve upgrade response: $approve"
check_response_and_report "$approve" "upgrade cannot be scheduled in the past" raw

get_height current_height
echo "Current height is $current_height"
plan_height=$(expr $current_height + 3)

echo "re-propose upgrade's plan height bigger than block height"
propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_name --upgrade-height=$plan_height --upgrade-info=$upgrade_info --from jack --yes)
propose=$(get_txn_result "$propose")
echo "propose upgrade response: $propose"
check_response "$propose" "\"code\": 0"

proposed_dclupgrade_query=$(dcld query dclupgrade proposed-upgrade --name=$upgrade_name)
echo "dclupgrade proposed upgrade query: $proposed_dclupgrade_query"
check_response_and_report "$proposed_dclupgrade_query" "\"name\": \"$upgrade_name\""
check_response_and_report "$proposed_dclupgrade_query" "\"height\": \"$plan_height\""
check_response_and_report "$proposed_dclupgrade_query" "\"info\": \"$upgrade_info\""
###########################################################################################################################################

test_divider

# APPROVE UPGRADE'S PLAN HEIGHT LESS THAN BLOCK HEIGHT AND RE-PROPOSE UPGRADE PLAN HEIGHT LESS THAN BLOCK HEIGHT
###########################################################################################################################################
get_height current_height
echo "Current height is $current_height"
plan_height=$(expr $current_height + 3)
random_string upgrade_name
random_string upgrade_info

echo "propose upgrade's plan height bigger than block height"
propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_name --upgrade-height=$plan_height --upgrade-info=$upgrade_info --from jack --yes)
propose=$(get_txn_result "$propose")
echo "propose upgrade response: $propose"
check_response "$propose" "\"code\": 0"

proposed_dclupgrade_query=$(dcld query dclupgrade proposed-upgrade --name=$upgrade_name)
echo "dclupgrade proposed upgrade query: $proposed_dclupgrade_query"
check_response_and_report "$proposed_dclupgrade_query" "\"name\": \"$upgrade_name\""
check_response_and_report "$proposed_dclupgrade_query" "\"height\": \"$plan_height\""
check_response_and_report "$proposed_dclupgrade_query" "\"info\": \"$upgrade_info\""

wait_for_height $(expr $plan_height + 5) 300 outage-safe
get_height current_height
echo "Current height is $current_height"

echo "approve upgrade's plan height less than block height"
approve=$(dcld tx dclupgrade approve-upgrade --name=$upgrade_name --from alice --yes)
approve=$(get_txn_result "$approve")
echo "approve upgrade response: $approve"
check_response_and_report "$approve" "upgrade cannot be scheduled in the past" raw

echo "re-propose upgrade's plan height bigger than block height"
propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_name --upgrade-height=$plan_height --upgrade-info=$upgrade_info --from jack --yes)
propose=$(get_txn_result "$propose")
echo "propose upgrade response: $propose"
check_response_and_report "$propose" "upgrade cannot be scheduled in the past" raw
###########################################################################################################################################

test_divider
get_height current_height
echo "Current height is $current_height"
plan_height=$(expr $current_height + 3)
random_string upgrade_name
random_string upgrade_info

# 18. TEST PROPOSE AND REJECT UPGRADE
echo "18. TEST PROPOSE AND REJECT UPGRADE"
test_divider

echo "jack (Trustee) propose upgrade"
result=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_name --upgrade-height=$plan_height --upgrade-info=$upgrade_info --from jack --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

echo "jack (Trustee) rejects upgrade"
result=$(dcld tx dclupgrade reject-upgrade --name=$upgrade_name --from jack --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"

test_divider

echo "Upgrade not found in proposed upgrade"
result=$(dcld query dclupgrade proposed-upgrade --name=$upgrade_name)
echo $result | jq
check_response "$result" "Not Found"

test_divider

echo "Upgrade not found in rejected upgrade"
result=$(dcld query dclupgrade rejected-upgrade --name=$upgrade_name)
echo $result | jq
check_response "$result" "Not Found"

test_divider

echo "Upgrade not found in approved upgrade query"
result=$(dcld query dclupgrade approved-upgrade --name=$upgrade_name)
echo $result | jq
check_response "$result" "Not Found"

echo "PASSED"