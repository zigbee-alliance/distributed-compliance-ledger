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

