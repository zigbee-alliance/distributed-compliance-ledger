set -euo pipefail
source integration_tests/cli/common.sh

# 1. check non-existent values via light client when no entry added
echo "check non-existent values via light client when no entry added"

test_divider

# connect to light client proxy
dcld config node tcp://localhost:26620
sleep 10

vid=$RANDOM

echo "Query non existent vendorinfo"
result=$(execute_with_retry "dcld query vendorinfo vendor --vid=$vid")
echo "$result"
check_response "$result" "Not Found"

test_divider

# 2. list queries should return a warning via light client

echo "list queries should return a warning via light client"

test_divider

echo "Request all vendorinfos"
result=$(execute_with_retry "dcld query vendorinfo all-vendors")
echo "$result"
check_response "$result" "List queries don't work with a Light Client Proxy"

test_divider

# 3. write entries

echo "write entries"

test_divider

# write requests can be sent via Full Node only
dcld config node tcp://localhost:26657 

vid=$RANDOM

vendor_account=vendor_account_$vid
echo "Create Vendor account - $vendor_account"
create_new_vendor_account $vendor_account $vid

test_divider

echo "Create VendorInfo Record for VID: $vid"
companyLegalName="XYZ IOT Devices Inc"
vendorName="XYZ Devices"
result=$(echo "test1234" | dcld tx vendorinfo add-vendor --vid=$vid --companyLegalName="$companyLegalName" --vendorName="$vendorName" --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"
echo "$result"

test_divider

# 4. check existent values via light client

echo "check existent values via light client"

test_divider

# connect to light client proxy
dcld config node tcp://localhost:26620
sleep 5

echo "Get vendorinfo"
result=$(execute_with_retry "dcld query vendorinfo vendor --vid=$vid")
echo "$result"
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"company_legal_name\": \"$companyLegalName\""
check_response "$result" "\"vendor_name\": \"$vendorName\""

test_divider


# 5. check non-existent values via light client when entry added

echo "check non-existent values via light client when entry added"

test_divider

vid=$RANDOM

echo "Query non existent vendorinfo"
result=$(execute_with_retry "dcld query vendorinfo vendor --vid=$vid")
echo "$result"
check_response "$result" "Not Found"

test_divider

# 6. try to write via light client proxy

echo "try to write via light client proxy"

test_divider

echo "Add vendorinfo"
result=$(echo "test1234" | dcld tx vendorinfo add-vendor --vid=$vid --companyLegalName="$companyLegalName" --vendorName="$vendorName" --from=$vendor_account --yes)
echo "$result"
check_response "$result" "Write requests don't work with a Light Client Proxy"


test_divider
