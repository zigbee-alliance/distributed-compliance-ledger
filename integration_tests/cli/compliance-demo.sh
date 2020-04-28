#!/bin/bash
set -e
source integration_tests/cli/common.sh

echo "Assign Vendor role to Jack"
result=$(echo "test1234" | zblcli tx authz assign-role $(zblcli keys show jack -a) "Vendor" --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

vid=$RANDOM
pid=$RANDOM
echo "Add Model with VID: $vid PID: $pid"
result=$(echo "test1234" | zblcli tx modelinfo add-model $vid $pid "Device #1" "Device Description" "SKU12FS" "1.0" "2.0" true --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Assign TestHouse role to Jack"
result=$(echo "test1234" | zblcli tx authz assign-role $(zblcli keys show jack -a) "TestHouse" --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Add Testing Result for Model VID: $vid PID: $pid"
testing_result="http://first.place.com"
test_date="2020-11-24T10:00:00Z"
result=$(echo "test1234" | zblcli tx compliancetest add-test-result $vid $pid "$testing_result" "$test_date" --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get Certified Model with VID: ${vid} PID: ${pid} before compliance record was created"
result=$(zblcli query compliance certified-model $vid $pid "zb")
check_response "$result" "\"value\": false"
echo "$result"

echo "Get Revoked Model with VID: ${vid} PID: ${pid} before compliance record was created"
result=$(zblcli query compliance revoked-model $vid $pid "zb")
check_response "$result" "\"value\": false"
echo "$result"

echo "Assign ZBCertificationCenter role to Jack"
result=$(echo "test1234" | zblcli tx authz assign-role $(zblcli keys show jack -a) "ZBCertificationCenter" --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Certify Model with VID: $vid PID: $pid"
certification_date="2020-01-01T00:00:00Z"
certification_type="zb"
result=$(echo "test1234" | zblcli tx compliance certify-model $vid $pid "$certification_type" "$certification_date" --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get Certified Model with VID: ${vid} PID: ${pid}"
result=$(zblcli query compliance certified-model $vid $pid $certification_type)
check_response "$result" "\"value\": true"
echo "$result"

echo "Get Compliance Info for Model with VID: ${vid} PID: ${pid}"
result=$(zblcli query compliance compliance-info $vid $pid $certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"state\": \"certified\""
check_response "$result" "\"date\": \"$certification_date\""
check_response "$result" "\"certification_type\": \"$certification_type\""
echo "$result"

echo "Get All Certified Models"
result=$(zblcli query compliance all-certified-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"certification_type\": \"$certification_type\""
echo "$result"

echo "Get All Compliance Info Recordss"
result=$(zblcli query compliance all-compliance-info-records)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"date\": \"$certification_date\""
echo "$result"

echo "Revoke Certification for Model with VID: $vid PID: $pid"
revocation_date="2020-02-02T02:20:20Z"
revocation_reason="some reason"
result=$(echo "test1234" | zblcli tx compliance revoke-model $vid $pid "$certification_type" "$revocation_date" --reason "$revocation_reason" --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get Compliance Info for Model with VID: ${vid} PID: ${pid}"
result=$(zblcli query compliance compliance-info $vid $pid $certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"state\": \"revoked\""
check_response "$result" "\"date\": \"$revocation_date\""
check_response "$result" "\"reason\": \"$revocation_reason\""
check_response "$result" "\"certification_type\": \"$certification_type\""
check_response "$result" "\"history\""
echo "$result"

echo "Get Revoked Model with VID: ${vid} PID: ${pid}"
result=$(zblcli query compliance revoked-model $vid $pid $certification_type)
check_response "$result" "\"value\": true"
echo "$result"

echo "Get Certified Model with VID: ${vid} PID: ${pid}"
result=$(zblcli query compliance certified-model $vid $pid $certification_type)
check_response "$result" "\"value\": false"
echo "$result"

echo "Get All Revoked Models"
result=$(zblcli query compliance all-revoked-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"certification_type\": \"$certification_type\""
echo "$result"

echo "Again Certify Model with VID: $vid PID: $pid"
certification_date="2020-03-03T00:00:00Z"
result=$(echo "test1234" | zblcli tx compliance certify-model $vid $pid "$certification_type" "$certification_date" --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get Compliance Info for Model with VID: ${vid} PID: ${pid}"
result=$(zblcli query compliance compliance-info $vid $pid $certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"state\": \"certified\""
check_response "$result" "\"date\": \"$certification_date\""
check_response "$result" "\"certification_type\": \"zb\""
echo "$result"

echo "Get Certified Model with VID: ${vid} PID: ${pid}"
result=$(zblcli query compliance certified-model $vid $pid $certification_type)
check_response "$result" "\"value\": true"
echo "$result"

echo "Get Revoked Model with VID: ${vid} PID: ${pid}"
result=$(zblcli query compliance revoked-model $vid $pid $certification_type)
check_response "$result" "\"value\": false"
echo "$result"

echo "Get All Compliance Infos"
result=$(zblcli query compliance all-compliance-info-records)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"state\": \"certified\""
echo "$result"