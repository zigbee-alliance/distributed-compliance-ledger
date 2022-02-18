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

upgrade_height=$RANDOM
random_string upgrade_info

echo "propose and approve upgrade"
random_string upgrade_name
propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_name --upgrade-height=$upgrade_height --upgrade-info=$upgrade_info --from alice --yes)
approve=$(dcld tx dclupgrade approve-upgrade --name=$upgrade_name --from bob --yes)
proposed_dclupgrade_query=$(dcld query dclupgrade proposed-upgrade --name=$upgrade_name)
approved_dclupgrade_query=$(dcld query dclupgrade approved-upgrade --name=$upgrade_name)
plan_query=$(dcld query upgrade plan)
check_response "$proposed_dclupgrade_query" "Not Found" raw
check_response "$approved_dclupgrade_query" "\"name\":\"$upgrade_name\"" raw
check_response_and_report "$plan_query" "\"name\":\"$upgrade_name\"" raw
check_response_and_report "$plan_query" "\"height\":\"$upgrade_height\"" raw
check_response_and_report "$plan_query" "\"info\":\"$upgrade_info\"" raw

test_divider

echo "proposer cannot approve upgrade"
random_string upgrade_name
propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_name --upgrade-height=$upgrade_height --upgrade-info=$upgrade_info --from alice --yes)
approve=$(dcld tx dclupgrade approve-upgrade --name=$upgrade_name --from alice --yes)
check_response_and_report "$approve" "unauthorized" raw

test_divider

echo "upgrade more approvals needed"
echo "Create Trustee account"
create_new_account trustee_account "Trustee"
random_string upgrade_name
propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_name --upgrade-height=$upgrade_height --upgrade-info=$upgrade_info --from $trustee_account --yes)
approve=$(dcld tx dclupgrade approve-upgrade --name=$upgrade_name --from alice --yes)
proposed_dclupgrade_query=$(dcld query dclupgrade proposed-upgrade --name=$upgrade_name)
approved_dclupgrade_query=$(dcld query dclupgrade approved-upgrade --name=$upgrade_name)
result=$(dcld query dclupgrade approved-upgrade --name=$upgrade_name)
check_response "$result" "Not Found"
check_response "$approved_dclupgrade_query" "Not Found"
check_response_and_report "$proposed_dclupgrade_query" "\"name\":\"$upgrade_name\"" raw
check_response_and_report "$proposed_dclupgrade_query" "\"height\":\"$upgrade_height\"" raw
check_response_and_report "$proposed_dclupgrade_query" "\"info\":\"$upgrade_info\"" raw

test_divider

echo "cannot approve upgrade twice"
random_string upgrade_name
propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_name --upgrade-height=$upgrade_height --upgrade-info=$upgrade_info --from alice --yes)
approve=$(dcld tx dclupgrade approve-upgrade --name=$upgrade_name --from bob --yes)
result=$(dcld tx dclupgrade approve-upgrade --name=$upgrade_name --from bob --yes)
check_response_and_report "$result" "unauthorized" raw

test_divider

echo "cannot propose upgrade twice"
random_string upgrade_name
propose=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_name --upgrade-height=$upgrade_height --upgrade-info=$upgrade_info --from alice --yes)
result=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_name --upgrade-height=$upgrade_height --upgrade-info=$upgrade_info --from alice --yes)
check_response_and_report "$result" "proposed upgrade already exists" raw

test_divider

echo "upgrade height < current height"
random_string upgrade_name
height=1
result=$(dcld tx dclupgrade propose-upgrade --name=$upgrade_name --upgrade-height=$height --upgrade-info=$upgrade_info --from alice --yes)
check_response_and_report "$result" "upgrade cannot be scheduled in the past" raw

test_divider

echo "PASSED"