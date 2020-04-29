#!/bin/bash
set -e

check_response() {
  result=$1
  expected_string=$2
  if [[ $result != *$expected_string* ]]; then
    echo "ERROR: command filed. The expected string: $expected_string not found in the result: $result"
    exit 1
  fi
}

response_does_not_contain() {
  result=$1
  unexpected_string=$2
  if [[ ${testmystring} == *$expected_string* ]];then
    echo "ERROR: command filed. The unexpected string: $unexpected_string found in the result: $result"
    exit 1
  fi
}

create_account_with_name(){
  name=$1

  echo "Generate key for $name"
  echo 'test1234' | zblcli keys add "$name"

  address=$(zblcli keys show $name -a)
  pubkey=$(zblcli keys show $name -p)

  echo "Jack creates account for $name"
  result=$(echo "test1234" | zblcli tx authnext create-account "$address" "$pubkey" --from jack --yes)
  check_response "$result" "\"success\": true"
  echo "$result"
}

