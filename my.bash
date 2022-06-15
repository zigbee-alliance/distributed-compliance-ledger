set -euo pipefail
source integration_tests/cli/common.sh

# Preparation of Actors
vid=$RANDOM
vendor_account=vendor_account_$vid
echo "Create Vendor account - $vendor_account"
create_new_vendor_account $vendor_account $vid

test_divider

echo "Create CertificationCenter account"
create_new_account zb_account "CertificationCenter"

test_divider

# Body
pid=$RANDOM
sv=$RANDOM
svs=$RANDOM
certification_date="2020-01-01T00:00:01Z"
zigbee_certification_type="zigbee"
matter_certification_type="matter"
full_certification_type="Full"
CbS_certification_type="CbS"
CTP_certification_type="CTP"
PFC_certification_type="PFC"

test_divider

echo "Add Model with VID: $vid PID: $pid"
result=$(echo "$passphrase" | dcld tx model add-model --vid=$vid --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0  --from $vendor_account --yes)
echo $result
check_response "$result" "\"code\": 0"

test_divider

echo "Add Model Version with VID: $vid PID: $pid SV: $sv SoftwareVersionString:$svs"
result=$(echo '$passphrase' | dcld tx model add-model-version --cdVersionNumber=1 --maxApplicableSoftwareVersion=10 --minApplicableSoftwareVersion=1 --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --from=$vendor_account --yes)
echo $result
check_response "$result" "\"code\": 0"

test_divider

echo "Certify Model with VID: $vid PID: $pid  SV: ${sv} with zigbee certification"
result=$(echo "$passphrase" | dcld tx compliance certify-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$zigbee_certification_type" --certificationDate="$certification_date" --programTypeVersion="1.0"  --CDCertificationID="someID" --familyID="someFID" --supportedClusters="someClusters" --compliancePlatformUsed="WIFI" --compliancePlatformVersion="V1" --OSVersion="someV" --certificationRoute="Full" --programType="pType" --transport="someTransport" --parentChild="parent" --from $zb_account --yes)
echo "$result"
check_response "$result" "\"code\": 0"

test_divider

echo "Certify Model with VID: $vid PID: $pid  SV: ${sv} with matter certification"
echo "dcld tx compliance certify-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$matter_certification_type" --certificationDate="$certification_date" --from $zb_account --yes"
result=$(echo "$passphrase" | dcld tx compliance certify-model --vid=$vid --pid=$pid --softwareVersion=$sv --softwareVersionString=$svs --certificationType="$matter_certification_type" --certificationDate="$certification_date" --programTypeVersion="1.0" --CDCertificationID="someID" --familyID="someFID" --supportedClusters="someClusters" --compliancePlatformUsed="WIFI" --compliancePlatformVersion="V1" --OSVersion="someV" --certificationRoute="Full" --programType="pType" --transport="someTransport" --parentChild="parent" --from $zb_account --yes)
echo "$result"
check_response "$result" "\"code\": 0"

test_divider

echo "Get Certified Model with VID: ${vid} PID: ${pid} SV: ${sv} for matter certification"
result=$(dcld query compliance certified-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$matter_certification_type)
check_response "$result" "\"value\": true"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"vid\": $vid"
echo "$result" | jq

test_divider

echo "Get Certified Model with VID: ${vid} PID: ${pid} SV: ${sv} for zigbee certification"
result=$(dcld query compliance certified-model --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$zigbee_certification_type)
check_response "$result" "\"value\": true"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"vid\": $vid"
echo "$result" | jq

test_divider

echo "Get Compliance Info for Model with VID: ${vid} PID: ${pid} SV: ${sv} for $zigbee_certification_type"
result=$(dcld query compliance compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$zigbee_certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"softwareVersionCertificationStatus\": 2"
check_response "$result" "\"date\": \"$certification_date\""
check_response "$result" "\"certificationType\": \"$zigbee_certification_type\""
check_response "$result" "\"programTypeVersion\": \"1.0\""
check_response "$result" "\"CDCertificationID\": \"someID\""
check_response "$result" "\"familyID\": \"someFID\""
check_response "$result" "\"supportedClusters\": \"someClusters\""
check_response "$result" "\"compliancePlatformUsed\": \"WIFI\""
check_response "$result" "\"compliancePlatformVersion\": \"V1\""
check_response "$result" "\"OSVersion\": \"someV\""
check_response "$result" "\"certificationRoute\": \"Full\""
check_response "$result" "\"programType\": \"pType\""
check_response "$result" "\"transport\": \"someTransport\""
check_response "$result" "\"parentChild\": \"parent\""
echo "$result" | jq

test_divider

echo "Get Compliance Info for Model with VID: ${vid} PID: ${pid} SV: ${sv} for $matter_certification_type"
result=$(dcld query compliance compliance-info --vid=$vid --pid=$pid --softwareVersion=$sv --certificationType=$matter_certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"softwareVersionCertificationStatus\": 2"
check_response "$result" "\"date\": \"$certification_date\""
check_response "$result" "\"certificationType\": \"$matter_certification_type\""
check_response "$result" "\"programTypeVersion\": \"1.0\""
check_response "$result" "\"CDCertificationID\": \"someID\""
check_response "$result" "\"familyID\": \"someFID\""
check_response "$result" "\"supportedClusters\": \"someClusters\""
check_response "$result" "\"compliancePlatformUsed\": \"WIFI\""
check_response "$result" "\"compliancePlatformVersion\": \"V1\""
check_response "$result" "\"certificationRoute\": \"Full\""
check_response "$result" "\"programType\": \"pType\""
check_response "$result" "\"transport\": \"someTransport\""
check_response "$result" "\"parentChild\": \"parent\""
echo "$result" | jq

test_divider

echo "PASSED"