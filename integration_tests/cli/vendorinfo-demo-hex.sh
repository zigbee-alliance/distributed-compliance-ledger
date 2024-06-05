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

# Preparation of Actors
vid_in_hex_format=0xA13
vid2_in_hex_format=0xA14
vid=2579
vid2=2580
vendor_account=vendor_account_$vid_in_hex_format
second_vendor_account=vendor_account_$vid2_in_hex_format
echo "Create First Vendor Account - $vendor_account"
create_new_vendor_account $vendor_account $vid_in_hex_format
echo "Create Second Vendor Account - $second_vendor_account"
create_new_vendor_account $second_vendor_account $vid2_in_hex_format

test_divider

# Create a vendor info record
echo "Create VendorInfo Record for VID: $vid_in_hex_format"
companyLegalName="XYZ IOT Devices Inc"
vendorName="XYZ Devices"
result=$(echo "test1234" | dcld tx vendorinfo add-vendor --vid=$vid_in_hex_format --companyLegalName="$companyLegalName" --vendorName="$vendorName" --from=$vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

# Query vendor info record
echo "Verify if VendorInfo Record for VID: $vid_in_hex_format is present or not"
result=$(dcld query vendorinfo vendor --vid=$vid_in_hex_format)
check_response "$result" "\"vendorID\": $vid"
check_response "$result" "\"companyLegalName\": \"$companyLegalName\""
check_response "$result" "\"vendorName\": \"$vendorName\""
echo "$result"

test_divider

# Update vendor info record
echo "Update vendor info record for VID: $vid_in_hex_format"
companyLegalName="ABC Subsidiary Corporation"
vendorLandingPageURL="https://www.w3.org/"
result=$(echo "test1234" | dcld tx vendorinfo update-vendor --vid=$vid_in_hex_format --companyLegalName="$companyLegalName" --vendorLandingPageURL=$vendorLandingPageURL --vendorName="$vendorName" --from=$vendor_account --yes)
result=$(get_txn_result "$result")
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

# Query updated vendor info record
echo "Verify if VendorInfo Record for VID: $vid_in_hex_format is updated or not"
result=$(dcld query vendorinfo vendor --vid=$vid_in_hex_format)
check_response "$result" "\"vendorID\": $vid"
check_response "$result" "\"companyLegalName\": \"$companyLegalName\""
check_response "$result" "\"vendorName\": \"$vendorName\""
check_response "$result" "\"vendorLandingPageURL\": \"$vendorLandingPageURL\""
echo "$result"

test_divider

# Create a vendor info record from a vendor account belonging to another vendor_account
vid1_in_hex_format=0xA15
vid1=2581
result=$(echo "test1234" | dcld tx vendorinfo add-vendor --vid=$vid1_in_hex_format --companyLegalName="$companyLegalName" --vendorName="$vendorName" --from=$vendor_account --yes 2>&1) || true
result=$(get_txn_result "$result")
echo "$result"
check_response_and_report "$result" "transaction should be signed by a vendor account associated with the vendorID $vid1"

test_divider

# Update a vendor info record from a vendor account belonging to another vendor_account
result=$(echo "test1234" | dcld tx vendorinfo update-vendor --vid=$vid_in_hex_format --companyLegalName="$companyLegalName" --vendorName="$vendorName" --from=$second_vendor_account --yes 2>&1) || true
result=$(get_txn_result "$result")
echo "$result"
check_response_and_report "$result" "transaction should be signed by a vendor account associated with the vendorID $vid"

test_divider
