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

