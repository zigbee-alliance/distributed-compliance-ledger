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

set -e

passphrase="test1234"

check_response() {
  result=$1
  expected_string=$2
  if [[ $result != *$expected_string* ]]; then
    echo "ERROR: command failed. The expected string: $expected_string not found in the result: $result"
    exit 1
  fi
}

check_response_and_report() {
  result=$1
  expected_string=$2
  check_response "$result" "$expected_string"
  echo "INFO: Result contains expected substring: $expected_string"
}

response_does_not_contain() {
  result=$1
  unexpected_string=$2
  if [[ ${testmystring} == *$expected_string* ]];then
    echo "ERROR: command failed. The unexpected string: $unexpected_string found in the result: $result"
    exit 1
  fi
}

create_new_account(){
  echo 1
  local  __resultvar=$1
  echo 2

  cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 6 | head -n 1

  echo "2.1"
  local name=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 6 | head -n 1)
  echo 3
  eval $__resultvar="'$name'"
  echo 4

  roles=$2

  echo "Account name: $name"

  echo "Generate key for $name"
  echo $passphrase | dclcli keys add "$name"

  address=$(dclcli keys show $name -a)
  pubkey=$(dclcli keys show $name -p)

  echo "Jack prupose account for \"$name\" with roles: \"$roles\""
  result=$(echo $passphrase | dclcli tx auth propose-add-account --address="$address" --pubkey="$pubkey" --roles=$roles --from jack --yes)
  check_response "$result" "\"success\": true"
  echo "$result"

  echo "Alice approve account for \"$name\" with roles: \"$roles\""
  result=$(echo $passphrase | dclcli tx auth approve-add-account --address="$address" --from alice --yes)
  check_response "$result" "\"success\": true"
  echo "$result"
}

