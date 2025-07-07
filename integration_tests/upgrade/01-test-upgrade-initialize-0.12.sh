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

binary_version="v0.12.0"

wget -O dcld-initial "https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/$binary_version/dcld"
chmod ugo+x dcld-initial

DCLD_BIN="./dcld-initial"
$DCLD_BIN config broadcast-mode block

container="validator-demo"

cleanup_validator_node() {
  if docker container ls -a | grep -q $container; then
    if docker container inspect $container | grep -q '"Status": "running"'; then
      echo "Stopping container"
      docker container kill $container
    fi

    echo "Removing container"
    docker container rm -f "$container"
  fi
}

add_validator_node() {
  random_string account
  address=""
  LOCALNET_DIR=".localnet"
  DCL_USER_HOME="/var/lib/dcl"
  DCL_DIR="$DCL_USER_HOME/.dcl"

  node_name="node-demo"
  node_p2p_port=26670
  node_client_port=26671
  chain_id="dclchain"
  ip="192.167.10.6"
  node0conn="tcp://192.167.10.2:26657"
  passphrase="test1234"
  docker_network="distributed-compliance-ledger_localnet"

  docker run -d --name $container --ip $ip -p "$node_p2p_port-$node_client_port:26656-26657" --network $docker_network -i dcledger

  docker cp $DCLD_BIN "$container":"$DCL_USER_HOME"/dcld

  test_divider

  echo "$account Configure CLI"
  docker exec $container /bin/sh -c "
    ./dcld config chain-id dclchain &&
    ./dcld config output json &&
    ./dcld config node $node0conn &&
    ./dcld config keyring-backend test &&
    ./dcld config broadcast-mode block"

  test_divider

  echo "$account Prepare Node configuration files"
  docker exec $container ./dcld init $node_name --chain-id $chain_id
  docker cp "$LOCALNET_DIR/node0/config/genesis.json" $container:$DCL_DIR/config
  peers="$(cat "$LOCALNET_DIR/node0/config/config.toml" | grep -o -E "persistent_peers = \".*\"")"
  docker exec $container sed -i "s/persistent_peers = \"\"/$peers/g" $DCL_DIR/config/config.toml
  docker exec $container sed -i 's/laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/g' $DCL_DIR/config/config.toml

  test_divider

  echo "Generate keys for $account"
  cmd="(echo $passphrase; echo $passphrase) | ./dcld keys add $account"
  docker exec $container /bin/sh -c "$cmd"

  address="$(docker exec $container /bin/sh -c "echo $passphrase | ./dcld keys show $account -a")"
  pubkey="$(docker exec $container /bin/sh -c "echo $passphrase | ./dcld keys show $account -p")"
  alice_address="$($DCLD_BIN keys show alice -a)"
  bob_address="$($DCLD_BIN keys show bob -a)"
  jack_address="$($DCLD_BIN keys show jack -a)"
  echo "Create account for $account and Assign NodeAdmin role"
  echo $passphrase | $DCLD_BIN tx auth propose-add-account --address="$address" --pubkey="$pubkey" --roles="NodeAdmin" --from jack --yes
  echo $passphrase | $DCLD_BIN tx auth approve-add-account --address="$address" --from alice --yes
  echo $passphrase | $DCLD_BIN tx auth approve-add-account --address="$address" --from bob --yes

  test_divider
  vaddress=$(docker exec $container ./dcld tendermint show-address)
  vpubkey=$(docker exec $container ./dcld tendermint show-validator)

  echo "Check pool response for yet unknown node \"$node_name\""
  result=$($DCLD_BIN query validator node --address "$address")
  check_response "$result" "Not Found"
  echo "$result"
  result=$($DCLD_BIN query validator last-power --address "$address")
  check_response "$result" "Not Found"
  echo "$result"

  echo "$account Add Node \"$node_name\" to validator set"

  ! read -r -d '' _script << EOF
      set -eu; echo test1234 | $DCLD_BIN tx validator add-node --pubkey='$vpubkey' --moniker="$node_name" --from="$account" --yes
EOF
  result="$(docker exec "$container" /bin/sh -c "echo test1234 | ./dcld tx validator add-node --pubkey='$vpubkey' --moniker="$node_name" --from="$account" --yes")"
  check_response "$result" "\"code\": 0"
  echo "$result"


  test_divider


  echo "Locating the app to $DCL_DIR/cosmovisor/genesis/bin directory"
  docker exec $container mkdir -p "$DCL_DIR"/cosmovisor/genesis/bin
  docker exec $container cp -f ./dcld "$DCL_DIR"/cosmovisor/genesis/bin/

  echo "$account Start Node \"$node_name\""
  docker exec -d $container /var/lib/dcl/./node_helper.sh
  sleep 10

  result=$($DCLD_BIN query validator node --address "$address")
  validator_address=$(echo "$result" | jq -r '.owner')
  echo "$result"
}

# constants
trustee_account_1="jack"
trustee_account_2="alice"
trustee_account_3="bob"
vendor_account="vendor_account"

vid=1
pid_1=1
pid_2=2
pid_3=3
device_type_id=12345
product_name="ProductName"
product_label="ProductLabel"
part_number="RCU2205A"
software_version=1
software_version_string="1.0"
cd_version_number=312
min_applicable_software_version=1
max_applicable_software_version=1000

certification_type="zigbee"
certification_date="2020-01-01T00:00:00Z"
provisional_date="2019-12-12T00:00:00Z"
cd_certificate_id="15DEXF"

root_cert_path="integration_tests/constants/root_cert"
root_cert_subject="MDQxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMRAwDgYDVQQKDAdyb290LWNh"
root_cert_subject_key_id="5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB"
root_cert_serial_number="442314047376310867378175982234956458728610743315"
root_cert_subject_as_text="O=root-ca,ST=some-state,C=AU"

test_root_cert_path="integration_tests/constants/test_root_cert"
test_root_cert_subject="MDAxGDAWBgNVBAMMD01hdHRlciBUZXN0IFBBQTEUMBIGCisGAQQBgqJ8AgEMBDEyNUQ="
test_root_cert_subject_key_id="E2:90:8D:36:9C:3C:A3:C1:13:BB:09:E2:4D:C1:CC:C5:A6:66:91:D4"
test_root_cert_serial_number="1647312298631"
test_root_cert_subject_as_text="CN=Matter Test PAA,vid=0x125D"
test_root_cert_vid=4701

google_root_cert_path="integration_tests/constants/google_root_cert"
google_root_cert_subject="MEsxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZHb29nbGUxFTATBgNVBAMMDE1hdHRlciBQQUEgMTEUMBIGCisGAQQBgqJ8AgEMBDYwMDY="
google_root_cert_subject_key_id="B0:00:56:81:B8:88:62:89:62:80:E1:21:18:A1:A8:BE:09:DE:93:21"
google_cert_serial_number="1"
google_cert_subject_as_text="CN=Matter PAA 1,O=Google,C=US,vid=0x6006"

intermediate_cert_path="integration_tests/constants/intermediate_cert"
intermediate_cert_subject="MDwxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMRgwFgYDVQQKDA9pbnRlcm1lZGlhdGUtY2E="
intermediate_cert_subject_key_id="4E:3B:73:F4:70:4D:C2:98:0D:DB:C8:5A:5F:02:3B:BF:86:25:56:2B"

vendor_name="VendorName"
company_legal_name="LegalCompanyName"
company_preferred_name="CompanyPreferredName"
vendor_landing_page_url="https://www.example.com"

random_string user_1
echo "$user_1 generates keys"
cmd="(echo $passphrase; echo $passphrase) | $DCLD_BIN keys add $user_1"
result="$(bash -c "$cmd")"
user_1_address=$(echo $passphrase | $DCLD_BIN keys show $user_1 -a)
user_1_pubkey=$(echo $passphrase | $DCLD_BIN keys show $user_1 -p)

random_string user_2
echo "$user_2 generates keys"
cmd="(echo $passphrase; echo $passphrase) | $DCLD_BIN keys add $user_2"
result="$(bash -c "$cmd")"
user_2_address=$(echo $passphrase | $DCLD_BIN keys show $user_2 -a)
user_2_pubkey=$(echo $passphrase | $DCLD_BIN keys show $user_2 -p)

random_string user_3
echo "$user_3 generates keys"
cmd="(echo $passphrase; echo $passphrase) | $DCLD_BIN keys add $user_3"
result="$(bash -c "$cmd")"
user_3_address=$(echo $passphrase | $DCLD_BIN keys show $user_3 -a)
user_3_pubkey=$(echo $passphrase | $DCLD_BIN keys show $user_3 -p)

echo "Create Vendor account $vendor_account"
result="$(echo $passphrase | $DCLD_BIN keys add "$vendor_account")"
_address=$(echo $passphrase | $DCLD_BIN keys show $vendor_account -a)
_pubkey=$(echo $passphrase | $DCLD_BIN keys show $vendor_account -p)
result="$(echo $passphrase | $DCLD_BIN tx auth propose-add-account --address="$_address" --pubkey="$_pubkey" --vid="$vid" --roles="Vendor" --from jack --yes)"

echo "Create CertificationCenter account"
certification_center_account="certification_center_account_"
result="$(echo $passphrase | $DCLD_BIN keys add $certification_center_account)"
_address=$(echo $passphrase | $DCLD_BIN keys show $certification_center_account -a)
_pubkey=$(echo $passphrase | $DCLD_BIN keys show $certification_center_account -p)
result="$(echo $passphrase | $DCLD_BIN tx auth propose-add-account --address="$_address" --pubkey="$_pubkey" --roles="CertificationCenter"  --from jack --yes)"
result="$(echo $passphrase | $DCLD_BIN tx auth approve-add-account --address="$_address" --from alice --yes)"

random_string trustee_account_4
random_string trustee_account_5

echo "Generate key for $trustee_account_4"
(echo $passphrase; echo $passphrase) | $DCLD_BIN keys add "$trustee_account_4"

echo "Generate key for $trustee_account_5"
(echo $passphrase; echo $passphrase) | $DCLD_BIN keys add "$trustee_account_5"

trustee_4_address=$(echo $passphrase | $DCLD_BIN keys show $trustee_account_4 -a)
trustee_4_pubkey=$(echo $passphrase | $DCLD_BIN keys show $trustee_account_4 -p)

trustee_5_address=$(echo $passphrase | $DCLD_BIN keys show $trustee_account_5 -a)
trustee_5_pubkey=$(echo $passphrase | $DCLD_BIN keys show $trustee_account_5 -p)

echo "Jack proposes account for trustee \"$trustee_account_4\""
result=$(echo $passphrase | $DCLD_BIN tx auth propose-add-account --address="$trustee_4_address" --pubkey="$trustee_4_pubkey" --roles=Trustee --from jack --yes)
check_response "$result" "\"code\": 0"

echo "Alice approves account for trustee \"$trustee_account_4\""
result=$(echo $passphrase | $DCLD_BIN tx auth approve-add-account --address="$trustee_4_address" --from alice --yes)
check_response "$result" "\"code\": 0"

echo "Jack proposes account for trustee \"$trustee_account_5\""
result=$(echo $passphrase | $DCLD_BIN tx auth propose-add-account --address="$trustee_5_address" --pubkey="$trustee_5_pubkey" --roles=Trustee --from jack --yes)
check_response "$result" "\"code\": 0"

echo "Alice approves account for trustee \"$trustee_account_5\""
result=$(echo $passphrase | $DCLD_BIN tx auth approve-add-account --address="$trustee_5_address" --from alice --yes)
check_response "$result" "\"code\": 0"

echo "$trustee_account_4 approves account for trustee \"$trustee_account_5\""
result=$(echo $passphrase | $DCLD_BIN tx auth approve-add-account --address="$trustee_5_address" --from $trustee_account_4 --yes)
check_response "$result" "\"code\": 0"

test_divider

# Body
echo "send all ledger update transactions before upgrade"

# VENDOR_INFO
echo "Add vendor $vendor_name"
result=$(echo $passphrase | $DCLD_BIN tx vendorinfo add-vendor --vid=$vid --vendorName=$vendor_name --companyLegalName=$company_legal_name --companyPreferredName=$company_preferred_name --vendorLandingPageURL=$vendor_landing_page_url --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"

test_divider

# MODEL and MODEL_VERSION

echo "Add model vid=$vid pid=$pid_1"
result=$(echo $passphrase | $DCLD_BIN tx model add-model --vid=$vid --pid=$pid_1 --deviceTypeID=$device_type_id --productName=$product_name --productLabel=$product_label --partNumber=$part_number --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Add model version vid=$vid pid=$pid_1"
result=$(echo $passphrase | $DCLD_BIN tx model add-model-version --vid=$vid --pid=$pid_1 --softwareVersion=$software_version --softwareVersionString=$software_version_string --cdVersionNumber=$cd_version_number --minApplicableSoftwareVersion=$min_applicable_software_version --maxApplicableSoftwareVersion=$max_applicable_software_version --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Add model vid=$vid pid=$pid_2"
result=$(echo $passphrase | $DCLD_BIN tx model add-model --vid=$vid --pid=$pid_2 --deviceTypeID=$device_type_id --productName=$product_name --productLabel=$product_label --partNumber=$part_number --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Add model version vid=$vid pid=$pid_2"
result=$(echo $passphrase | $DCLD_BIN tx model add-model-version --vid=$vid --pid=$pid_2 --softwareVersion=$software_version --softwareVersionString=$software_version_string --cdVersionNumber=$cd_version_number --minApplicableSoftwareVersion=$min_applicable_software_version --maxApplicableSoftwareVersion=$max_applicable_software_version --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Add model vid=$vid pid=$pid_3"
result=$(echo $passphrase | $DCLD_BIN tx model add-model --vid=$vid --pid=$pid_3 --deviceTypeID=$device_type_id --productName=$product_name --productLabel=$product_label --partNumber=$part_number --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Add model version vid=$vid pid=$pid_3"
result=$(echo $passphrase | $DCLD_BIN tx model add-model-version --vid=$vid --pid=$pid_3 --softwareVersion=$software_version --softwareVersionString=$software_version_string --cdVersionNumber=$cd_version_number --minApplicableSoftwareVersion=$min_applicable_software_version --maxApplicableSoftwareVersion=$max_applicable_software_version --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Delete model vid=$vid pid=$pid_3"
result=$(echo $passphrase | $DCLD_BIN tx model delete-model --vid=$vid --pid=$pid_3 --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"

test_divider

# CERTIFY_DEVICE_COMPLIANCE

echo "Certify model vid=$vid pid=$pid_1"
result=$(echo $passphrase | $DCLD_BIN tx compliance certify-model --vid=$vid --pid=$pid_1 --softwareVersion=$software_version --softwareVersionString=$software_version_string  --certificationType=$certification_type --certificationDate=$certification_date --cdCertificateId=$cd_certificate_id --from=$certification_center_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Certify model vid=$vid pid=$pid_2"
result=$(echo $passphrase | $DCLD_BIN tx compliance certify-model --vid=$vid --pid=$pid_2 --softwareVersion=$software_version --softwareVersionString=$software_version_string  --certificationType=$certification_type --certificationDate=$certification_date --cdCertificateId=$cd_certificate_id --from=$certification_center_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Revoke model certification vid=$vid pid=$pid_2"
result=$(echo $passphrase | $DCLD_BIN tx compliance revoke-model --vid=$vid --pid=$pid_2 --softwareVersion=$software_version --softwareVersionString=$software_version_string --certificationType=$certification_type --revocationDate=$certification_date --from=$certification_center_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Provision model vid=$vid pid=$pid_3"
result=$(echo $passphrase | $DCLD_BIN tx compliance provision-model --vid=$vid --pid=$pid_3 --softwareVersion=$software_version --softwareVersionString=$software_version_string --certificationType=$certification_type --provisionalDate=$provisional_date --cdCertificateId=$cd_certificate_id --from=$certification_center_account --yes)
check_response "$result" "\"code\": 0"

test_divider

# X509 PKI

echo "Propose add root_certificate"
result=$(echo $passphrase | $DCLD_BIN tx pki propose-add-x509-root-cert --certificate="$root_cert_path" --from=$trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add root_certificate"
result=$(echo $passphrase | $DCLD_BIN tx pki approve-add-x509-root-cert --subject="$root_cert_subject" --subject-key-id=$root_cert_subject_key_id --from=$trustee_account_2 --yes)
check_response "$result" "\"code\": 0"

echo "Approve add root_certificate"
result=$(echo $passphrase | $DCLD_BIN tx pki approve-add-x509-root-cert --subject="$root_cert_subject" --subject-key-id=$root_cert_subject_key_id --from=$trustee_account_3 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Propose add test_root_certificate"
result=$(echo $passphrase | $DCLD_BIN tx pki propose-add-x509-root-cert --certificate="$test_root_cert_path" --from=$trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add test_root_certificate"
result=$(echo $passphrase | $DCLD_BIN tx pki approve-add-x509-root-cert --subject="$test_root_cert_subject" --subject-key-id=$test_root_cert_subject_key_id --from=$trustee_account_2 --yes)
check_response "$result" "\"code\": 0"

echo "Approve add test_root_certificate"
result=$(echo $passphrase | $DCLD_BIN tx pki approve-add-x509-root-cert --subject="$test_root_cert_subject" --subject-key-id=$test_root_cert_subject_key_id --from=$trustee_account_3 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Propose add google_root_certificate"
result=$(echo $passphrase | $DCLD_BIN tx pki propose-add-x509-root-cert --certificate="$google_root_cert_path" --from=$trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Reject add google_root_certificate"
result=$(echo $passphrase | $DCLD_BIN tx pki reject-add-x509-root-cert --subject="$google_root_cert_subject" --subject-key-id=$google_root_cert_subject_key_id --from=$trustee_account_2 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Add intermediate_cert"
result=$(echo $passphrase | $DCLD_BIN tx pki add-x509-cert --certificate="$intermediate_cert_path" --from=$trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Propose revoke root_certificate"
result=$(echo "$passphrase" | $DCLD_BIN tx pki propose-revoke-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from="$trustee_account_1" --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve revoke root_certificate"
result=$(echo "$passphrase" | $DCLD_BIN tx pki approve-revoke-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from="$trustee_account_2" --yes)
check_response "$result" "\"code\": 0"

echo "Approve revoke root_certificate"
result=$(echo "$passphrase" | $DCLD_BIN tx pki approve-revoke-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from="$trustee_account_3" --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Propose revoke test_root_certificate"
result=$(echo $passphrase | $DCLD_BIN tx pki propose-revoke-x509-root-cert --subject="$test_root_cert_subject" --subject-key-id="$test_root_cert_subject_key_id" --from $trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

# AUTH

echo "Propose add account $user_1_address"
result=$(echo $passphrase | $DCLD_BIN tx auth propose-add-account --address="$user_1_address" --pubkey="$user_1_pubkey" --roles="CertificationCenter" --from="$trustee_account_1" --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_1_address"
result=$($DCLD_BIN tx auth approve-add-account --address="$user_1_address" --from=$trustee_account_2 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_1_address"
result=$($DCLD_BIN tx auth approve-add-account --address="$user_1_address" --from=$trustee_account_3 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Propose add account $user_2_address"
result=$(echo $passphrase | $DCLD_BIN tx auth propose-add-account --address="$user_2_address" --pubkey=$user_2_pubkey --roles=CertificationCenter --from=$trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_2_address"
result=$(echo $passphrase | $DCLD_BIN tx auth approve-add-account --address="$user_2_address" --from=$trustee_account_2 --yes)
check_response "$result" "\"code\": 0"

echo "Approve add account $user_2_address"
result=$(echo $passphrase | $DCLD_BIN tx auth approve-add-account --address="$user_2_address" --from=$trustee_account_3 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Propose add account $user_3_address"
result=$(echo $passphrase | $DCLD_BIN tx auth propose-add-account --address="$user_3_address" --pubkey=$user_3_pubkey --roles=CertificationCenter --from=$trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Propose revoke account $user_1_address"
result=$(echo $passphrase | $DCLD_BIN tx auth propose-revoke-account --address="$user_1_address" --from=$trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve revoke account $user_1_address"
result=$(echo $passphrase | $DCLD_BIN tx auth approve-revoke-account --address="$user_1_address" --from=$trustee_account_2 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve revoke account $user_1_address"
result=$(echo $passphrase | $DCLD_BIN tx auth approve-revoke-account --address="$user_1_address" --from=$trustee_account_3 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Propose revoke account $user_2_address"
result=$(echo $passphrase | $DCLD_BIN tx auth propose-revoke-account --address="$user_2_address" --from=$trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

# VALIDATOR_NODE

echo "Add new validator node"

cleanup_validator_node
add_validator_node

test_divider

echo "Disable node"
# FIXME: use proper binary (not dcld but $DCLD_BIN)
result=$(docker exec "$container" /bin/sh -c "echo test1234  | dcld tx validator disable-node --from=$account --yes")
check_response "$result" "\"code\": 0"

test_divider

echo "Enable node"
# FIXME: use proper binary (not dcld but $DCLD_BIN)
result=$(docker exec "$container" /bin/sh -c "echo test1234  | dcld tx validator enable-node --from=$account --yes")
check_response "$result" "\"code\": 0"

test_divider

echo "Propose disable node"
result=$(echo $passphrase | $DCLD_BIN tx validator propose-disable-node --address=$validator_address --from=$trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve disable node"
result=$(echo $passphrase | $DCLD_BIN tx validator approve-disable-node --address=$validator_address --from=$trustee_account_2 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve disable node"
result=$(echo $passphrase | $DCLD_BIN tx validator approve-disable-node --address=$validator_address --from=$trustee_account_3 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Enable node"
# FIXME: use proper binary (not dcld but $DCLD_BIN)
result=$(docker exec "$container" /bin/sh -c "echo test1234  | dcld tx validator enable-node --from=$account --yes")
check_response "$result" "\"code\": 0"

test_divider

echo "Propose disable node"
result=$(echo $passphrase | $DCLD_BIN tx validator propose-disable-node --address=$validator_address --from=$trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Get x509 root certificates"
result=$($DCLD_BIN query pki x509-cert --subject="$test_root_cert_subject" --subject-key-id="$test_root_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$test_root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$test_root_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$test_root_cert_serial_number\""
check_response "$result" "\"subjectAsText\": \"$test_root_cert_subject_as_text\""
response_does_not_contain "$result" "\"vid\":"

echo "Get x509 proposed root certificates"
result=$($DCLD_BIN query pki proposed-x509-root-cert --subject="$google_root_cert_subject" --subject-key-id="$google_root_cert_subject_key_id")
echo $result | jq
check_response "$result" "\"subject\": \"$google_root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$google_root_cert_subject_key_id\""
check_response "$result" "\"serialNumber\": \"$google_cert_serial_number\""
check_response "$result" "\"subjectAsText\": \"$google_cert_subject_as_text\""
response_does_not_contain "$result" "\"vid\":"

echo "Initialize 0.12.0 passed"

test_divider

rm -f $DCLD_BIN