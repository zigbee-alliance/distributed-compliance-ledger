#!/bin/bash
set -e

passphrase="test1234"

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

create_new_account(){
  local  __resultvar=$1
  local name=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 6 | head -n 1)
  eval $__resultvar="'$name'"

  roles=$2

  echo "Account name: $name"

  echo "Generate key for $name"
  echo $passphrase | zblcli keys add "$name"

  address=$(zblcli keys show $name -a)
  pubkey=$(zblcli keys show $name -p)

  echo "Jack prupose account for \"$name\" with roles: \"$roles\""
  result=$(echo $passphrase | zblcli tx auth propose-add-account --address="$address" --pubkey="$pubkey" --roles=$roles --from jack --yes)
  check_response "$result" "\"success\": true"
  echo "$result"

  echo "Alice approve account for \"$name\" with roles: \"$roles\""
  result=$(echo $passphrase | zblcli tx auth approve-add-account --address="$address" --from alice --yes)
  check_response "$result" "\"success\": true"
  echo "$result"
}

