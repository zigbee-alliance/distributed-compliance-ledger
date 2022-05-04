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

# get_height current_height
# echo "Current height is $current_height"

# plan_height=$(expr $current_height + 2)
# random_string upgrade_info
# random_string upgrade_name

# echo "approve upgrade plan height less block height and we can re-propose upgrade and approve"

# test_divider

# propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_name --upgrade-height=$plan_height --upgrade-info=$upgrade_info --from jack --yes)
# echo "propose upgrade response: $propose"
# check_response "$propose" "\"code\": 0"

# test_divider

# proposed_dclupgrade_query=$(dcld query dclupgrade proposed-upgrade --name=$upgrade_name)
# echo "dclupgrade proposed upgrade query: $proposed_dclupgrade_query"
# check_response_and_report "$proposed_dclupgrade_query" "\"name\": \"$upgrade_name\""
# check_response_and_report "$proposed_dclupgrade_query" "\"height\": \"$plan_height\""
# check_response_and_report "$proposed_dclupgrade_query" "\"info\": \"$upgrade_info\""

# test_divider

# sleep 25

# approve=$(dcld tx dclupgrade approve-upgrade --name=$upgrade_name --from alice --yes 2>&1 || true)
# check_response_and_report "$approve" "\"code\": 0"

# test_divider

# proposed_dclupgrade_query=$(dcld query dclupgrade proposed-upgrade --name=$upgrade_name)
# check_response_and_report "$proposed_dclupgrade_query" "Not Found"

# test_divider

# get_height current_height
# echo "Current height is $current_height"

# plan_height=$(expr $current_height + 5)

# test_divider

# propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_name --upgrade-height=$plan_height --upgrade-info=$upgrade_info --from jack --yes)
# echo "propose upgrade response: $propose"
# check_response "$propose" "\"code\": 0"

# test_divider

# proposed_dclupgrade_query=$(dcld query dclupgrade proposed-upgrade --name=$upgrade_name)
# echo "dclupgrade proposed upgrade query: $proposed_dclupgrade_query"
# check_response_and_report "$proposed_dclupgrade_query" "\"name\": \"$upgrade_name\""
# check_response_and_report "$proposed_dclupgrade_query" "\"height\": \"$plan_height\""
# check_response_and_report "$proposed_dclupgrade_query" "\"info\": \"$upgrade_info\""

# test_divider

# approve=$(dcld tx dclupgrade approve-upgrade --name=$upgrade_name --from alice --yes)
# echo "approve upgrade response: $approve"
# check_response "$approve" "\"code\": 0"

# test_divider

upgrade_height=$(($RANDOM + 10000000))
random_string upgrade_info

echo "propose and approve upgrade"
echo "Create Trustee account"
create_new_account trustee_account "Trustee"
random_string upgrade_name
propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_name --upgrade-height=$upgrade_height --upgrade-info=$upgrade_info --from $trustee_account --yes)
echo "propose upgrade response: $propose"
check_response "$propose" "\"code\": 0"

proposed_dclupgrade_query=$(dcld query dclupgrade proposed-upgrade --name=$upgrade_name)
echo "dclupgrade proposed upgrade query: $proposed_dclupgrade_query"
check_response_and_report "$proposed_dclupgrade_query" "\"name\": \"$upgrade_name\""
check_response_and_report "$proposed_dclupgrade_query" "\"height\": \"$upgrade_height\""
check_response_and_report "$proposed_dclupgrade_query" "\"info\": \"$upgrade_info\""

approve=$(dcld tx dclupgrade approve-upgrade --name=$upgrade_name --from alice --yes)
echo "approve upgrade response: $approve"
check_response "$approve" "\"code\": 0"

reject=$(dcld tx dclupgrade reject-upgrade --name=$upgrade_name --from alice --yes)
echo "reject upgrade response: $reject"
response_does_not_contain "$reject" "\"code\": 0"

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
echo "propose upgrade response: $propose"
check_response "$propose" "\"code\": 0"

approve=$(dcld tx dclupgrade approve-upgrade --name=$upgrade_name --from alice --yes)
echo "approve upgrade response: $approve"
check_response_and_report "$approve" "unauthorized" raw


test_divider


echo "cannot approve upgrade twice"
random_string upgrade_name
propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_name --upgrade-height=$upgrade_height --upgrade-info=$upgrade_info --from alice --yes)
echo "propose upgrade response: $propose"
check_response "$propose" "\"code\": 0"

approve=$(dcld tx dclupgrade approve-upgrade --name=$upgrade_name --from bob --yes)
echo "approve upgrade response: $approve"
check_response "$approve" "\"code\": 0"

second_approve=$(dcld tx dclupgrade approve-upgrade --name=$upgrade_name --from bob --yes)
echo "second approve upgrade response: $second_approve"
check_response_and_report "$second_approve" "unauthorized" raw

test_divider


echo "cannot propose upgrade twice"
random_string upgrade_name
propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_name --upgrade-height=$upgrade_height --upgrade-info=$upgrade_info --from alice --yes)
echo "propose upgrade response: $propose"
check_response "$propose" "\"code\": 0"

second_propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_name --upgrade-height=$upgrade_height --upgrade-info=$upgrade_info --from alice --yes)
echo "second propose upgrade response: $second_propose"
check_response_and_report "$second_propose" "proposed upgrade already exists" raw


test_divider


echo "upgrade height < current height"
random_string upgrade_name
height=1
propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_name --upgrade-height=$height --upgrade-info=$upgrade_info --from alice --yes)
echo "propose upgrade response: $propose"
check_response_and_report "$propose" "upgrade cannot be scheduled in the past" raw


test_divider

upgrade_height=$(($RANDOM + 10000000))
random_string upgrade_info

echo "propose and reject upgrade"
random_string upgrade_name
propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_name --upgrade-height=$upgrade_height --upgrade-info=$upgrade_info --from $trustee_account --yes)
echo "propose upgrade response: $propose"
check_response "$propose" "\"code\": 0"

test_divider

proposed_dclupgrade_query=$(dcld query dclupgrade proposed-upgrade --name=$upgrade_name)
echo "dclupgrade proposed upgrade query: $proposed_dclupgrade_query"
check_response_and_report "$proposed_dclupgrade_query" "\"name\": \"$upgrade_name\""
check_response_and_report "$proposed_dclupgrade_query" "\"height\": \"$upgrade_height\""
check_response_and_report "$proposed_dclupgrade_query" "\"info\": \"$upgrade_info\""

test_divider

reject=$(dcld tx dclupgrade reject-upgrade --name=$upgrade_name --from alice --yes)
echo "reject upgrade response: $reject"
check_response "$reject" "\"code\": 0"

test_divider

second_reject=$(dcld tx dclupgrade reject-upgrade --name=$upgrade_name --from alice --yes)
echo "second_reject upgrade response: $reject"
response_does_not_contain "$second_reject" "\"code\": 0"

approve=$(dcld tx dclupgrade approve-upgrade --name=$upgrade_name --from alice --yes)
echo "approve upgrade response: $approve"
response_does_not_contain "$approve" "\"code\": 0"

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

echo "PASSED"