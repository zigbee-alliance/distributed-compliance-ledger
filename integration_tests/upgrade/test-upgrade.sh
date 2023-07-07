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

TESTS_TO_RUN=${1:-all}

DETAILED_OUTPUT=true

LOCALNET_DIR=".localnet"

LOG_PREFIX="[run all] "
SED_EXT=
if [ "$(uname)" == "Darwin" ]; then
    # Mac OS X sed needs the file extension when -i flag is used. Keeping it empty as we don't need backupfile
    SED_EXT="''"
fi

if ${DETAILED_OUTPUT}; then
  DETAILED_OUTPUT_TARGET=/dev/stdout
else
  DETAILED_OUTPUT_TARGET=/dev/null
fi

source integration_tests/cli/common.sh

log() {
  echo "${LOG_PREFIX}$1"
}

  # patch configs properly by having all values >= 1 sec, otherwise headers may start having time from the future and light client verification will fail
  # if we patch config to have new blocks created in less than 1 sec, the min time in a time header is still 1 sec.
  # So, new blocks started to be from the future.
patch_consensus_config() {
  local NODE_CONFIGS="$(find "$LOCALNET_DIR" -type f -name "config.toml" -wholename "*node*")"

  for NODE_CONFIG in ${NODE_CONFIGS}; do
    sed -i $SED_EXT 's/timeout_propose = "3s"/timeout_propose = "1s"/g' "${NODE_CONFIG}"
    #sed -i $SED_EXT 's/timeout_prevote = "1s"/timeout_prevote = "1s"/g' "${NODE_CONFIG}"
    #sed -i $SED_EXT 's/timeout_precommit = "1s"/timeout_precommit = "1s"/g' "${NODE_CONFIG}"
    sed -i $SED_EXT 's/timeout_commit = "5s"/timeout_commit = "1s"/g' "${NODE_CONFIG}"
  done
}

init_pool() {
  local _patch_config="${1:-yes}";
  local _localnet_init_target=${2:-localnet_init}

  log "Setting up pool"

  log "-> Generating network configuration" >${DETAILED_OUTPUT_TARGET}
  make ${_localnet_init_target} &>${DETAILED_OUTPUT_TARGET}

  if [ "$_patch_config" = "yes" ];
  then
    patch_consensus_config
  fi;

  log "-> Running pool" >${DETAILED_OUTPUT_TARGET}
  make localnet_start &>${DETAILED_OUTPUT_TARGET}

  log "-> Waiting for the second block (needed to request proofs)" >${DETAILED_OUTPUT_TARGET}
  wait_for_height 2 20
}

cleanup_pool() {
  log "Cleaning up pool"
  log "-> Stopping pool & Removing configurations" >${DETAILED_OUTPUT_TARGET}
  make localnet_clean
}

stop_rest_server() {
  log "Stopping cli in rest-server mode"
  killall dcld
}

container="validator-demo"
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

  docker build -f Dockerfile-build -t dcld-build .
  docker container create --name dcld-build-inst dcld-build
  docker cp dcld-build-inst:/go/bin/dcld ./
  docker rm dcld-build-inst

  docker run -d --name $container --ip $ip -p "$node_p2p_port-$node_client_port:26656-26657" --network $docker_network -i dcledger

  docker cp ./dcld "$container":"$DCL_USER_HOME"/
  rm -f ./dcld

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
  alice_address="$(dcld keys show alice -a)"
  bob_address="$(dcld keys show bob -a)"
  jack_address="$(dcld keys show jack -a)"
  echo "Create account for $account and Assign NodeAdmin role"
  echo $passphrase | dcld tx auth propose-add-account --address="$address" --pubkey="$pubkey" --roles="NodeAdmin" --from jack --yes
  echo $passphrase | dcld tx auth approve-add-account --address="$address" --from alice --yes
 echo $passphrase | dcld tx auth approve-add-account --address="$address" --from bob --yes

  test_divider
  vaddress=$(docker exec $container ./dcld tendermint show-address)
  vpubkey=$(docker exec $container ./dcld tendermint show-validator)

  echo "Check pool response for yet unknown node \"$node_name\""
  result=$(dcld query validator node --address "$address")
  check_response "$result" "Not Found"
  echo "$result"
  result=$(dcld query validator last-power --address "$address")
  check_response "$result" "Not Found"
  echo "$result"

  echo "$account Add Node \"$node_name\" to validator set"

  ! read -r -d '' _script << EOF
      set -eu; echo test1234 | dcld tx validator add-node --pubkey='$vpubkey' --moniker="$node_name" --from="$account" --yes
EOF
  result="$(docker exec "$container" /bin/sh -c "echo test1234 | ./dcld tx validator add-node --pubkey='$vpubkey' --moniker="$node_name" --from="$account" --yes")"
  check_response "$result" "\"code\": 0"
  echo "$result"


  test_divider


  echo "Locating the app to $DCL_DIR/cosmovisor/genesis/bin directory"
  docker exec $container mkdir -p "$DCL_DIR"/cosmovisor/genesis/bin
  docker exec $container cp -f ./dcld "$DCL_DIR"/cosmovisor/genesis/bin/

  echo "$account Start Node \"$node_name\""
  docker exec -d $container cosmovisor start
  sleep 10

  result=$(dcld query validator node --address "$address")
  validator_address=$(echo "$result" | jq -r '.owner')
  echo "$result"
}

cleanup() {
  if docker container ls -a | grep -q $container; then
    if docker container inspect $container | grep -q '"Status": "running"'; then
      echo "Stopping container"
      docker container kill $container
    fi

    echo "Removing container"
    docker container rm -f "$container"
  fi
}
trap cleanup EXIT

cleanup

# Preparation

# constants
trustee_account_1="jack"
trustee_account_2="alice"
trustee_account_3="bob"
vendor_account="vendor_account"

plan_name="v0.13.0-pre"
upgrade_checksum="sha256:2e1df8fcd97d0f35f214fcc0fc9bc9bb28827f8815bba4b6ab797b19d8d0d4ac"

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

test_root_cert_path="integration_tests/constants/test_root_cert"
test_root_cert_subject="MDAxGDAWBgNVBAMMD01hdHRlciBUZXN0IFBBQTEUMBIGCisGAQQBgqJ8AgEMBDEyNUQ="
test_root_cert_subject_key_id="E2:90:8D:36:9C:3C:A3:C1:13:BB:09:E2:4D:C1:CC:C5:A6:66:91:D4"

google_root_cert_path="integration_tests/constants/google_root_cert"
google_root_cert_subject="MEsxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZHb29nbGUxFTATBgNVBAMMDE1hdHRlciBQQUEgMTEUMBIGCisGAQQBgqJ8AgEMBDYwMDY="
google_root_cert_subject_key_id="B0:00:56:81:B8:88:62:89:62:80:E1:21:18:A1:A8:BE:09:DE:93:21"

intermediate_cert_path="integration_tests/constants/intermediate_cert"
intermediate_cert_subject="MDwxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMRgwFgYDVQQKDA9pbnRlcm1lZGlhdGUtY2E="
intermediate_cert_subject_key_id="4E:3B:73:F4:70:4D:C2:98:0D:DB:C8:5A:5F:02:3B:BF:86:25:56:2B"

vendor_name="VendorName"
company_legal_name="LegalCompanyName"
company_preferred_name="CompanyPreferredName"
vendor_landing_page_url="https://www.example.com"

random_string user_1
echo "$user_1 generates keys"
cmd="(echo $passphrase; echo $passphrase) | dcld keys add $user_1"
result="$(bash -c "$cmd")"
user_1_address=$(echo $passphrase | dcld keys show $user_1 -a)
user_1_pubkey=$(echo $passphrase | dcld keys show $user_1 -p)

random_string user_2
echo "$user_2 generates keys"
cmd="(echo $passphrase; echo $passphrase) | dcld keys add $user_2"
result="$(bash -c "$cmd")"
user_2_address=$(echo $passphrase | dcld keys show $user_2 -a)
user_2_pubkey=$(echo $passphrase | dcld keys show $user_2 -p)

random_string user_3
echo "$user_3 generates keys"
cmd="(echo $passphrase; echo $passphrase) | dcld keys add $user_3"
result="$(bash -c "$cmd")"
user_3_address=$(echo $passphrase | dcld keys show $user_3 -a)
user_3_pubkey=$(echo $passphrase | dcld keys show $user_3 -p)

echo "Create Vendor account $vendor_account"
create_new_vendor_account $vendor_account $vid

echo "Create CertificationCenter account"
create_new_account certification_center_account "CertificationCenter"

random_string trustee_account_4
random_string trustee_account_5

echo "Generate key for $trustee_account_4"
(echo $passphrase; echo $passphrase) | dcld keys add "$trustee_account_4"

echo "Generate key for $trustee_account_5"
(echo $passphrase; echo $passphrase) | dcld keys add "$trustee_account_5"

trustee_4_address=$(echo $passphrase | dcld keys show $trustee_account_4 -a)
trustee_4_pubkey=$(echo $passphrase | dcld keys show $trustee_account_4 -p)

trustee_5_address=$(echo $passphrase | dcld keys show $trustee_account_5 -a)
trustee_5_pubkey=$(echo $passphrase | dcld keys show $trustee_account_5 -p)
  
echo "Jack proposes account for trustee \"$trustee_account_4\""
result=$(echo $passphrase | dcld tx auth propose-add-account --address="$trustee_4_address" --pubkey="$trustee_4_pubkey" --roles=Trustee --from jack --yes)
check_response "$result" "\"code\": 0"

echo "Alice approves account for trustee \"$trustee_account_4\""
result=$(echo $passphrase | dcld tx auth approve-add-account --address="$trustee_4_address" --from alice --yes)
check_response "$result" "\"code\": 0"

echo "Jack proposes account for trustee \"$trustee_account_5\""
result=$(echo $passphrase | dcld tx auth propose-add-account --address="$trustee_5_address" --pubkey="$trustee_5_pubkey" --roles=Trustee --from jack --yes)
check_response "$result" "\"code\": 0"

echo "Alice approves account for trustee \"$trustee_account_5\""
result=$(echo $passphrase | dcld tx auth approve-add-account --address="$trustee_5_address" --from alice --yes)
check_response "$result" "\"code\": 0"

echo "$trustee_account_4 approves account for trustee \"$trustee_account_5\""
result=$(echo $passphrase | dcld tx auth approve-add-account --address="$trustee_5_address" --from $trustee_account_4 --yes)
check_response "$result" "\"code\": 0"

test_divider

# Body
echo "send all ledger update transactions before upgrade"

# VENDOR_INFO
echo "Add vendor $vendor_name"
result=$(echo $passphrase | dcld tx vendorinfo add-vendor --vid=$vid --vendorName=$vendor_name --companyLegalName=$company_legal_name --companyPreferredName=$company_preferred_name --vendorLandingPageURL=$vendor_landing_page_url --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"

test_divider

# MODEL and MODEL_VERSION

echo "Add model vid=$vid pid=$pid_1"
result=$(echo $passphrase | dcld tx model add-model --vid=$vid --pid=$pid_1 --deviceTypeID=$device_type_id --productName=$product_name --productLabel=$product_label --partNumber=$part_number --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Add model version vid=$vid pid=$pid_1"
result=$(echo $passphrase | dcld tx model add-model-version --vid=$vid --pid=$pid_1 --softwareVersion=$software_version --softwareVersionString=$software_version_string --cdVersionNumber=$cd_version_number --minApplicableSoftwareVersion=$min_applicable_software_version --maxApplicableSoftwareVersion=$max_applicable_software_version --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Add model vid=$vid pid=$pid_2"
result=$(echo $passphrase | dcld tx model add-model --vid=$vid --pid=$pid_2 --deviceTypeID=$device_type_id --productName=$product_name --productLabel=$product_label --partNumber=$part_number --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Add model version vid=$vid pid=$pid_2"
result=$(echo $passphrase | dcld tx model add-model-version --vid=$vid --pid=$pid_2 --softwareVersion=$software_version --softwareVersionString=$software_version_string --cdVersionNumber=$cd_version_number --minApplicableSoftwareVersion=$min_applicable_software_version --maxApplicableSoftwareVersion=$max_applicable_software_version --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Add model vid=$vid pid=$pid_3"
result=$(echo $passphrase | dcld tx model add-model --vid=$vid --pid=$pid_3 --deviceTypeID=$device_type_id --productName=$product_name --productLabel=$product_label --partNumber=$part_number --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Add model version vid=$vid pid=$pid_3"
result=$(echo $passphrase | dcld tx model add-model-version --vid=$vid --pid=$pid_3 --softwareVersion=$software_version --softwareVersionString=$software_version_string --cdVersionNumber=$cd_version_number --minApplicableSoftwareVersion=$min_applicable_software_version --maxApplicableSoftwareVersion=$max_applicable_software_version --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Delete model vid=$vid pid=$pid_3"
result=$(echo $passphrase | dcld tx model delete-model --vid=$vid --pid=$pid_3 --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"

test_divider

# CERTIFY_DEVICE_COMPLIANCE

echo "Certify model vid=$vid pid=$pid_1"
result=$(echo $passphrase | dcld tx compliance certify-model --vid=$vid --pid=$pid_1 --softwareVersion=$software_version --softwareVersionString=$software_version_string  --certificationType=$certification_type --certificationDate=$certification_date --cdCertificateId=$cd_certificate_id --from=$certification_center_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Certify model vid=$vid pid=$pid_2"
result=$(echo $passphrase | dcld tx compliance certify-model --vid=$vid --pid=$pid_2 --softwareVersion=$software_version --softwareVersionString=$software_version_string  --certificationType=$certification_type --certificationDate=$certification_date --cdCertificateId=$cd_certificate_id --from=$certification_center_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Revoke model certification vid=$vid pid=$pid_2"
result=$(echo $passphrase | dcld tx compliance revoke-model --vid=$vid --pid=$pid_2 --softwareVersion=$software_version --softwareVersionString=$software_version_string --certificationType=$certification_type --revocationDate=$certification_date --from=$certification_center_account --yes)
check_response "$result" "\"code\": 0"

test_divider

test_divider

echo "Provision model vid=$vid pid=$pid_3"
result=$(echo $passphrase | dcld tx compliance provision-model --vid=$vid --pid=$pid_3 --softwareVersion=$software_version --softwareVersionString=$software_version_string --certificationType=$certification_type --provisionalDate=$provisional_date --cdCertificateId=$cd_certificate_id --from=$certification_center_account --yes)
check_response "$result" "\"code\": 0"

test_divider

# X509 PKI

echo "Propose add root_certificate"
result=$(echo $passphrase | dcld tx pki propose-add-x509-root-cert --certificate="$root_cert_path" --from=$trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add root_certificate"
result=$(echo $passphrase | dcld tx pki approve-add-x509-root-cert --subject="$root_cert_subject" --subject-key-id=$root_cert_subject_key_id --from=$trustee_account_2 --yes)
check_response "$result" "\"code\": 0"

echo "Approve add root_certificate"
result=$(echo $passphrase | dcld tx pki approve-add-x509-root-cert --subject="$root_cert_subject" --subject-key-id=$root_cert_subject_key_id --from=$trustee_account_3 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Propose add test_root_certificate"
result=$(echo $passphrase | dcld tx pki propose-add-x509-root-cert --certificate="$test_root_cert_path" --from=$trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add test_root_certificate"
result=$(echo $passphrase | dcld tx pki approve-add-x509-root-cert --subject="$test_root_cert_subject" --subject-key-id=$test_root_cert_subject_key_id --from=$trustee_account_2 --yes)
check_response "$result" "\"code\": 0"

echo "Approve add test_root_certificate"
result=$(echo $passphrase | dcld tx pki approve-add-x509-root-cert --subject="$test_root_cert_subject" --subject-key-id=$test_root_cert_subject_key_id --from=$trustee_account_3 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Propose add google_root_certificate"
result=$(echo $passphrase | dcld tx pki propose-add-x509-root-cert --certificate="$google_root_cert_path" --from=$trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Reject add google_root_certificate"
result=$(echo $passphrase | dcld tx pki reject-add-x509-root-cert --subject="$google_root_cert_subject" --subject-key-id=$google_root_cert_subject_key_id --from=$trustee_account_2 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Add intermediate_cert"
result=$(echo $passphrase | dcld tx pki add-x509-cert --certificate="$intermediate_cert_path" --from=$trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Propose revoke root_certificate"
result=$(echo "$passphrase" | dcld tx pki propose-revoke-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from="$trustee_account_1" --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve revoke root_certificate"
result=$(echo "$passphrase" | dcld tx pki approve-revoke-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from="$trustee_account_2" --yes)
check_response "$result" "\"code\": 0"

echo "Approve revoke root_certificate"
result=$(echo "$passphrase" | dcld tx pki approve-revoke-x509-root-cert --subject="$root_cert_subject" --subject-key-id="$root_cert_subject_key_id" --from="$trustee_account_3" --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Propose revoke test_root_certificate"
result=$(echo $passphrase | dcld tx pki propose-revoke-x509-root-cert --subject="$test_root_cert_subject" --subject-key-id="$test_root_cert_subject_key_id" --from $trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

# AUTH

echo "Propose add account $user_1_address"
result=$(echo $passphrase | dcld tx auth propose-add-account --address="$user_1_address" --pubkey="$user_1_pubkey" --roles="CertificationCenter" --from="$trustee_account_1" --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_1_address"
result=$(dcld tx auth approve-add-account --address="$user_1_address" --from=$trustee_account_2 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_1_address"
result=$(dcld tx auth approve-add-account --address="$user_1_address" --from=$trustee_account_3 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Propose add account $user_2_address"
result=$(echo $passphrase | dcld tx auth propose-add-account --address="$user_2_address" --pubkey=$user_2_pubkey --roles=CertificationCenter --from=$trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_2_address"
result=$(echo $passphrase | dcld tx auth approve-add-account --address="$user_2_address" --from=$trustee_account_2 --yes)
check_response "$result" "\"code\": 0"

echo "Approve add account $user_2_address"
result=$(echo $passphrase | dcld tx auth approve-add-account --address="$user_2_address" --from=$trustee_account_3 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Propose add account $user_3_address"
result=$(echo $passphrase | dcld tx auth propose-add-account --address="$user_3_address" --pubkey=$user_3_pubkey --roles=CertificationCenter --from=$trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Propose revoke account $user_1_address"
result=$(echo $passphrase | dcld tx auth propose-revoke-account --address="$user_1_address" --from=$trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve revoke account $user_1_address"
result=$(echo $passphrase | dcld tx auth approve-revoke-account --address="$user_1_address" --from=$trustee_account_2 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve revoke account $user_1_address"
result=$(echo $passphrase | dcld tx auth approve-revoke-account --address="$user_1_address" --from=$trustee_account_3 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Propose revoke account $user_2_address"
result=$(echo $passphrase | dcld tx auth propose-revoke-account --address="$user_2_address" --from=$trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

# VALIDATOR_NODE

echo "Add new validator node"
add_validator_node

test_divider

echo "Disable node"
result=$(docker exec "$container" /bin/sh -c "echo test1234  | dcld tx validator disable-node --from=$account --yes")
check_response "$result" "\"code\": 0"

test_divider

echo "Enable node"
result=$(docker exec "$container" /bin/sh -c "echo test1234  | dcld tx validator enable-node --from=$account --yes")
check_response "$result" "\"code\": 0"

test_divider

echo "Propose disable node"
result=$(echo $passphrase | dcld tx validator propose-disable-node --address=$validator_address --from=$trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve disable node"
result=$(echo $passphrase | dcld tx validator approve-disable-node --address=$validator_address --from=$trustee_account_2 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve disable node"
result=$(echo $passphrase | dcld tx validator approve-disable-node --address=$validator_address --from=$trustee_account_3 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Enable node"
result=$(docker exec "$container" /bin/sh -c "echo test1234  | dcld tx validator enable-node --from=$account --yes")
check_response "$result" "\"code\": 0"

test_divider

echo "Propose disable node"
result=$(echo $passphrase | dcld tx validator propose-disable-node --address=$validator_address --from=$trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider


get_height current_height
echo "Current height is $current_height"

plan_height=$(expr $current_height \+ 20)

test_divider

echo "Propose upgrade $plan_name at height $plan_height"
echo "https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/$plan_name/dcld.ubuntu.tar.gz?checksum=$upgrade_checksum"
result=$(echo $passphrase | dcld tx dclupgrade propose-upgrade --name=$plan_name --upgrade-height=$plan_height --upgrade-info="{\"binaries\":{\"linux/amd64\":\"https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/$plan_name/dcld.ubuntu.tar.gz?checksum=$upgrade_checksum\"}}" --from $trustee_account_1 --yes)
echo "$result"
check_response "$result" "\"code\": 0"

test_divider

echo "Approve upgrade $plan_name"
result=$(echo $passphrase | dcld tx dclupgrade approve-upgrade --name $plan_name --from $trustee_account_2 --yes)
echo "$result"
check_response "$result" "\"code\": 0"

test_divider

echo "Approve upgrade $plan_name"
result=$(echo $passphrase | dcld tx dclupgrade approve-upgrade --name $plan_name --from $trustee_account_3 --yes)
echo "$result"
check_response "$result" "\"code\": 0"

test_divider

echo "Wait for block height to become greater than upgrade $plan_name plan height"
wait_for_height $(expr $plan_height + 1) 300 outage-safe

test_divider

echo "Verify that no upgrade has been scheduled anymore"
result=$(dcld query upgrade plan 2>&1) || true
check_response_and_report "$result" "no upgrade scheduled" raw

test_divider

echo "Verify that upgrade is applied"
result=$(dcld query upgrade applied $plan_name)
echo "$result"

test_divider

echo "Verify that old data is not corrupted"

# VENDORINFO

echo "Verify if VendorInfo Record for VID: $vid is present or not"
result=$(dcld query vendorinfo vendor --vid=$vid)
check_response "$result" "\"vendorID\": $vid"
check_response "$result" "\"companyLegalName\": \"$company_legal_name\""
check_response "$result" "\"vendorName\": \"$vendor_name\""

echo "Request all vendor infos"
result=$(dcld query vendorinfo all-vendors)
check_response "$result" "\"vendorID\": $vid"
check_response "$result" "\"companyLegalName\": \"$company_legal_name\""
check_response "$result" "\"vendorName\": \"$vendor_name\""

test_divider

# MODEL

echo "Get Model with VID: $vid PID: $pid_1"
result=$(dcld query model get-model --vid=$vid --pid=$pid_1)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_1"
check_response "$result" "\"productLabel\": \"$product_label\""

echo "Get Model with VID: $vid PID: $pid_2"
result=$(dcld query model get-model --vid=$vid --pid=$pid_2)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_2"
check_response "$result" "\"productLabel\": \"$product_label\""

echo "Get all models"
result=$(dcld query model all-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_1"
check_response "$result" "\"pid\": $pid_2"

echo "Get Vendor Models with VID: ${vid}"
result=$(dcld query model vendor-models --vid=$vid)
check_response "$result" "\"pid\": $pid_1"
check_response "$result" "\"pid\": $pid_2"

echo "Get model version VID: $vid PID: $pid_1"
result=$(dcld query model model-version --vid=$vid --pid=$pid_1 --softwareVersion=$software_version)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_1"
check_response "$result" "\"softwareVersion\": $software_version"

echo "Get model version VID: $vid PID: $pid_2"
result=$(dcld query model model-version --vid=$vid --pid=$pid_2 --softwareVersion=$software_version)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_2"
check_response "$result" "\"softwareVersion\": $software_version"

test_divider

# COMPLIANCE

echo "Get certified model vid=$vid pid=$pid_1"
result=$(dcld query compliance certified-model --vid=$vid --pid=$pid_1 --softwareVersion=$software_version --certificationType=$certification_type)
check_response "$result" "\"value\": true"
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_1"
check_response "$result" "\"softwareVersion\": $software_version"
check_response "$result" "\"certificationType\": \"$certification_type\""

echo "Get revoked Model with VID: $vid PID: $pid_2"
result=$(dcld query compliance revoked-model --vid=$vid --pid=$pid_2 --softwareVersion=$software_version --certificationType=$certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_2"

echo "Get provisional model with VID: $vid PID: $pid_3"
result=$(dcld query compliance provisional-model --vid=$vid --pid=$pid_3 --softwareVersion=$software_version --certificationType=$certification_type)
check_response "$result" "\"value\": true"
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_3"

echo "Get compliance-info model with VID: $vid PID: $pid_1"
result=$(dcld query compliance compliance-info --vid=$vid --pid=$pid_1 --softwareVersion=$software_version --certificationType=$certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_1"
check_response "$result" "\"softwareVersion\": $software_version"
check_response "$result" "\"certificationType\": \"$certification_type\""

echo "Get compliance-info model with VID: $vid PID: $pid_2"
result=$(dcld query compliance compliance-info --vid=$vid --pid=$pid_2 --softwareVersion=$software_version --certificationType=$certification_type)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_2"
check_response "$result" "\"softwareVersion\": $software_version"
check_response "$result" "\"certificationType\": \"$certification_type\""

echo "Get device software compliance cDCertificateId=$cd_certificate_id"
result=$(dcld query compliance device-software-compliance --cdCertificateId=$cd_certificate_id)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_1"

echo "Get all certified models"
result=$(dcld query compliance all-certified-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_1"

echo "Get all provisional models"
result=$(dcld query compliance all-provisional-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_3"

echo "Get all revoked models"
result=$(dcld query compliance all-revoked-models)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_2"

echo "Get all compliance infos"
result=$(dcld query compliance all-compliance-info)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_1"
check_response "$result" "\"pid\": $pid_2"

echo "Get all device software compliances"
result=$(dcld query compliance all-device-software-compliance)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_1"
check_response "$result" "\"cDCertificateId\": \"$cd_certificate_id\""

test_divider

# PKI

echo "Get all x509 root certificates"
result=$(dcld query pki all-x509-root-certs)
check_response "$result" "\"subject\": \"$test_root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$test_root_cert_subject_key_id\""

echo "Get all revoked x509 root certificates"
result=$(dcld query pki all-revoked-x509-root-certs)
check_response "$result" "\"subject\": \"$root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id\""

echo "Get all proposed x509 root certificates"
result=$(dcld query pki all-proposed-x509-root-certs)
check_response "$result" "\"subject\": \"$google_root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$google_root_cert_subject_key_id\""

echo "Get all proposed x509 root certificates"
result=$(dcld query pki all-proposed-x509-root-certs-to-revoke)
check_response "$result" "\"subject\": \"$test_root_cert_subject\""
check_response "$result" "\"subjectKeyId\": \"$test_root_cert_subject_key_id\""

test_divider

# AUTH

echo "Get all accounts"
result=$(dcld query auth all-accounts)
check_response "$result" "\"address\": \"$user_2_address\""

echo "Get all proposed accounts"
result=$(dcld query auth all-proposed-accounts)
check_response "$result" "\"address\": \"$user_3_address\""

echo "Get all proposed accounts to revoke"
result=$(dcld query auth all-proposed-accounts-to-revoke)
check_response "$result" "\"address\": \"$user_2_address\""

echo "Get all revoked accounts"
result=$(dcld query auth all-revoked-accounts)
check_response "$result" "\"address\": \"$user_1_address\""

test_divider

# Validator

echo "Get proposed node to disable"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator proposed-disable-node --address="$address"")
check_response "$result" "\"address\": \"$validator_address\""

test_divider

########################################################################################

# after upgrade constatnts

vid_new=2
pid_1_new=11
pid_2_new=22
pid_3_new=33
device_type_id_new=1234
product_name_new="ProductNameNew"
product_label_new="ProductLabelNew"
part_number_new="RCU2205B"
software_version_new=2
software_version_string_new="2.0"
cd_version_number_new=313
min_applicable_software_version_new=2
max_applicable_software_version_new=2000

certification_type_new="matter"
certification_date_new="2021-01-01T00:00:00Z"
provisional_date_new="2010-12-12T00:00:00Z"
cd_certificate_id_new="15DEXC"

root_cert_path_new="integration_tests/constants/google_root_cert_gsr4"
root_cert_subject_new="MFAxJDAiBgNVBAsTG0dsb2JhbFNpZ24gRUNDIFJvb3QgQ0EgLSBSNDETMBEGA1UEChMKR2xvYmFsU2lnbjETMBEGA1UEAxMKR2xvYmFsU2lnbg=="
root_cert_subject_key_id_new="54:B0:7B:AD:45:B8:E2:40:7F:FB:0A:6E:FB:BE:33:C9:3C:A3:84:D5"

test_root_cert_path_new="integration_tests/constants/google_root_cert_r1"
test_root_cert_subject_new="MEcxCzAJBgNVBAYTAlVTMSIwIAYDVQQKExlHb29nbGUgVHJ1c3QgU2VydmljZXMgTExDMRQwEgYDVQQDEwtHVFMgUm9vdCBSMQ=="
test_root_cert_subject_key_id_new="E4:AF:2B:26:71:1A:2B:48:27:85:2F:52:66:2C:EF:F0:89:13:71:3E"

google_root_cert_path_new="integration_tests/constants/google_root_cert_r2"
google_root_cert_subject_new="MEcxCzAJBgNVBAYTAlVTMSIwIAYDVQQKExlHb29nbGUgVHJ1c3QgU2VydmljZXMgTExDMRQwEgYDVQQDEwtHVFMgUm9vdCBSMg=="
google_root_cert_subject_key_id_new="BB:FF:CA:8E:23:9F:4F:99:CA:DB:E2:68:A6:A5:15:27:17:1E:D9:0E"

intermediate_cert_path_new="integration_tests/constants/intermediate_cert_gsr4"
intermediate_cert_subject_new="MEYxCzAJBgNVBAYTAlVTMSIwIAYDVQQKExlHb29nbGUgVHJ1c3QgU2VydmljZXMgTExDMRMwEQYDVQQDEwpHVFMgQ0EgMkQ0"
intermediate_cert_subject_key_id_new="A8:88:D9:8A:39:AC:65:D5:82:4B:37:A8:95:6C:65:43:CD:44:01:E0"

test_data_url="https://www.newexample.com"

vendor_name_new="VendorNameNew"
company_legal_name_new="LegalCompanyNameNew"
company_preferred_name_new="CompanyPreferredNameNew"
vendor_landing_page_url_new="https://www.newexample.com"

vendor_account_new="vendor_account_new"
certification_center_account_new="certification_center_account_new"

echo "Create Vendor account $vendor_account_new"

result="$(echo $passphrase | dcld keys add "$vendor_account_new")"
_address=$(echo $passphrase | dcld keys show $vendor_account_new -a)
_pubkey=$(echo $passphrase | dcld keys show $vendor_account_new -p)
result="$(echo $passphrase | dcld tx auth propose-add-account --address="$_address" --pubkey="$_pubkey" --vid="$vid_new" --roles="Vendor" --from "$trustee_account_1" --yes)"
result="$(echo $passphrase | dcld tx auth approve-add-account --address="$_address" --from "$trustee_account_2" --yes)"
result="$(echo $passphrase | dcld tx auth approve-add-account --address="$_address" --from "$trustee_account_3" --yes)"
result="$(echo $passphrase | dcld tx auth approve-add-account --address="$_address" --from "$trustee_account_4" --yes)"

echo "Create CertificationCenter account"

result="$(echo $passphrase | dcld keys add "$certification_center_account_new")"
_address=$(echo $passphrase | dcld keys show $certification_center_account_new -a)
_pubkey=$(echo $passphrase | dcld keys show $certification_center_account_new -p)
result="$(echo $passphrase | dcld tx auth propose-add-account --address="$_address" --pubkey="$_pubkey" --roles="CertificationCenter" --from "$trustee_account_1" --yes)"
result="$(echo $passphrase | dcld tx auth approve-add-account --address="$_address" --from "$trustee_account_2" --yes)"
result="$(echo $passphrase | dcld tx auth approve-add-account --address="$_address" --from "$trustee_account_3" --yes)"
result="$(echo $passphrase | dcld tx auth approve-add-account --address="$_address" --from "$trustee_account_4" --yes)"

random_string user_4
echo "$user_4 generates keys"
cmd="(echo $passphrase; echo $passphrase) | dcld keys add $user_4"
result="$(bash -c "$cmd")"
user_4_address=$(echo $passphrase | dcld keys show $user_4 -a)
user_4_pubkey=$(echo $passphrase | dcld keys show $user_4 -p)

random_string user_5
echo "$user_5 generates keys"
cmd="(echo $passphrase; echo $passphrase) | dcld keys add $user_5"
result="$(bash -c "$cmd")"
user_5_address=$(echo $passphrase | dcld keys show $user_5 -a)
user_5_pubkey=$(echo $passphrase | dcld keys show $user_5 -p)

random_string user_6
echo "$user_6 generates keys"
cmd="(echo $passphrase; echo $passphrase) | dcld keys add $user_6"
result="$(bash -c "$cmd")"
user_6_address=$(echo $passphrase | dcld keys show $user_6 -a)
user_6_pubkey=$(echo $passphrase | dcld keys show $user_6 -p)

# send all ledger update transactions after upgrade

# VENDOR_INFO
echo "Add vendor $vendor_name_new"
result=$(echo $passphrase | dcld tx vendorinfo add-vendor --vid=$vid_new --vendorName=$vendor_name_new --companyLegalName=$company_legal_name_new --companyPreferredName=$company_preferred_name_new --vendorLandingPageURL=$vendor_landing_page_url_new --from=$vendor_account_new --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Update vendor $vendor_name"
result=$(echo $passphrase | dcld tx vendorinfo update-vendor --vid=$vid --vendorName=$vendor_name --companyLegalName=$company_legal_name --companyPreferredName=$company_preferred_name_new --vendorLandingPageURL=$vendor_landing_page_url_new --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"

test_divider

# MODEL and MODEL_VERSION

echo "Add model vid=$vid_new pid=$pid_1_new"
result=$(echo $passphrase | dcld tx model add-model --vid=$vid_new --pid=$pid_1_new --deviceTypeID=$device_type_id_new --productName=$product_name_new --productLabel=$product_label_new --partNumber=$part_number_new --from=$vendor_account_new --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Add model version vid=$vid_new pid=$pid_1_new"
result=$(echo $passphrase | dcld tx model add-model-version --vid=$vid_new --pid=$pid_1_new --softwareVersion=$software_version_new --softwareVersionString=$software_version_string_new --cdVersionNumber=$cd_version_number_new --minApplicableSoftwareVersion=$min_applicable_software_version_new --maxApplicableSoftwareVersion=$max_applicable_software_version_new --from=$vendor_account_new --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Add model vid=$vid_new pid=$pid_2_new"
result=$(echo $passphrase | dcld tx model add-model --vid=$vid_new --pid=$pid_2_new --deviceTypeID=$device_type_id_new --productName=$product_name_new --productLabel=$product_label_new --partNumber=$part_number_new --from=$vendor_account_new --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Add model version vid=$vid_new pid=$pid_2_new"
result=$(echo $passphrase | dcld tx model add-model-version --vid=$vid_new --pid=$pid_2_new --softwareVersion=$software_version_new --softwareVersionString=$software_version_string_new --cdVersionNumber=$cd_version_number_new --minApplicableSoftwareVersion=$min_applicable_software_version_new --maxApplicableSoftwareVersion=$max_applicable_software_version_new --from=$vendor_account_new --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Add model vid=$vid_new pid=$pid_3_new"
result=$(echo $passphrase | dcld tx model add-model --vid=$vid_new --pid=$pid_3_new --deviceTypeID=$device_type_id_new --productName=$product_name_new --productLabel=$product_label_new --partNumber=$part_number_new --from=$vendor_account_new --yes)
check_response "$result" "\"code\": 0"

echo "Add model version vid=$vid_new pid=$pid_3_new"
result=$(echo $passphrase | dcld tx model add-model-version --vid=$vid_new --pid=$pid_3_new --softwareVersion=$software_version_new --softwareVersionString=$software_version_string_new --cdVersionNumber=$cd_version_number_new --minApplicableSoftwareVersion=$min_applicable_software_version_new --maxApplicableSoftwareVersion=$max_applicable_software_version_new --from=$vendor_account_new --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Delete model vid=$vid_new pid=$pid_3_new"
result=$(echo $passphrase | dcld tx model delete-model --vid=$vid_new --pid=$pid_3_new --from=$vendor_account_new --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Update model vid=$vid pid=$pid_2"
result=$(echo $passphrase | dcld tx model update-model --vid=$vid --pid=$pid_2 --productName=$product_name --productLabel=$product_label_new --partNumber=$part_number_new --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Update model version vid=$vid pid=$pid_2"
result=$(echo $passphrase | dcld tx model update-model-version --vid=$vid --pid=$pid_2 --softwareVersion=$software_version --minApplicableSoftwareVersion=$min_applicable_software_version_new --maxApplicableSoftwareVersion=$max_applicable_software_version_new --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"

test_divider

# CERTIFY_DEVICE_COMPLIANCE

echo "Certify model vid=$vid_new pid=$pid_1_new"
result=$(echo $passphrase | dcld tx compliance certify-model --vid=$vid_new --pid=$pid_1_new --softwareVersion=$software_version_new --softwareVersionString=$software_version_string_new  --certificationType=$certification_type_new --certificationDate=$certification_date_new --cdCertificateId=$cd_certificate_id_new --from=$certification_center_account_new --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Certify model vid=$vid_new pid=$pid_2_new"
result=$(echo $passphrase | dcld tx compliance certify-model --vid=$vid_new --pid=$pid_2_new --softwareVersion=$software_version_new --softwareVersionString=$software_version_string_new  --certificationType=$certification_type_new --certificationDate=$certification_date_new --cdCertificateId=$cd_certificate_id_new --from=$certification_center_account_new --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Revoke model certification vid=$vid_new pid=$pid_2_new"
result=$(echo $passphrase | dcld tx compliance revoke-model --vid=$vid_new --pid=$pid_2_new --softwareVersion=$software_version_new --softwareVersionString=$software_version_string_new --certificationType=$certification_type_new --revocationDate=$certification_date_new --from=$certification_center_account_new --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Provision model vid=$vid_new pid=$pid_3_new"
result=$(echo $passphrase | dcld tx compliance provision-model --vid=$vid_new --pid=$pid_3_new --softwareVersion=$software_version_new --softwareVersionString=$software_version_string_new --certificationType=$certification_type_new --provisionalDate=$provisional_date_new --cdCertificateId=$cd_certificate_id_new --from=$certification_center_account_new --yes)
check_response "$result" "\"code\": 0"

test_divider

# X509 PKI

echo "Propose add root_certificate"
result=$(echo $passphrase | dcld tx pki propose-add-x509-root-cert --certificate="$root_cert_path_new" --from=$trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add root_certificate"
result=$(echo $passphrase | dcld tx pki approve-add-x509-root-cert --subject="$root_cert_subject_new" --subject-key-id=$root_cert_subject_key_id_new --from=$trustee_account_2 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "reject add root_certificate"
result=$(echo $passphrase | dcld tx pki reject-add-x509-root-cert --subject="$root_cert_subject_new" --subject-key-id=$root_cert_subject_key_id_new --from=$trustee_account_2 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add root_certificate"
result=$(echo $passphrase | dcld tx pki approve-add-x509-root-cert --subject="$root_cert_subject_new" --subject-key-id=$root_cert_subject_key_id_new --from=$trustee_account_3 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add root_certificate"
result=$(echo $passphrase | dcld tx pki approve-add-x509-root-cert --subject="$root_cert_subject_new" --subject-key-id=$root_cert_subject_key_id_new --from=$trustee_account_4 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add root_certificate"
result=$(echo $passphrase | dcld tx pki approve-add-x509-root-cert --subject="$root_cert_subject_new" --subject-key-id=$root_cert_subject_key_id_new --from=$trustee_account_5 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Propose add test_root_certificate"
result=$(echo $passphrase | dcld tx pki propose-add-x509-root-cert --certificate="$test_root_cert_path_new" --from=$trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add test_root_certificate"
result=$(echo $passphrase | dcld tx pki approve-add-x509-root-cert --subject="$test_root_cert_subject_new" --subject-key-id=$test_root_cert_subject_key_id_new --from=$trustee_account_2 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add test_root_certificate"
result=$(echo $passphrase | dcld tx pki approve-add-x509-root-cert --subject="$test_root_cert_subject_new" --subject-key-id=$test_root_cert_subject_key_id_new --from=$trustee_account_3 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add test_root_certificate"
result=$(echo $passphrase | dcld tx pki approve-add-x509-root-cert --subject="$test_root_cert_subject_new" --subject-key-id=$test_root_cert_subject_key_id_new --from=$trustee_account_4 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Propose add google_root_certificate"
result=$(echo $passphrase | dcld tx pki propose-add-x509-root-cert --certificate="$google_root_cert_path_new" --from=$trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Add intermediate_cert"
result=$(echo $passphrase | dcld tx pki add-x509-cert --certificate="$intermediate_cert_path_new" --from=$trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Revoke intermediate_cert"
result=$(echo $passphrase | dcld tx pki revoke-x509-cert --subject="$intermediate_cert_subject_new" --subject-key-id="$intermediate_cert_subject_key_id_new" --from=$trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Propose revoke root_certificate"
result=$(echo "$passphrase" | dcld tx pki propose-revoke-x509-root-cert --subject="$root_cert_subject_new" --subject-key-id="$root_cert_subject_key_id_new" --from="$trustee_account_1" --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve revoke root_certificate"
result=$(echo "$passphrase" | dcld tx pki approve-revoke-x509-root-cert --subject="$root_cert_subject_new" --subject-key-id="$root_cert_subject_key_id_new" --from="$trustee_account_2" --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve revoke root_certificate"
result=$(echo "$passphrase" | dcld tx pki approve-revoke-x509-root-cert --subject="$root_cert_subject_new" --subject-key-id="$root_cert_subject_key_id_new" --from="$trustee_account_3" --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve revoke root_certificate"
result=$(echo "$passphrase" | dcld tx pki approve-revoke-x509-root-cert --subject="$root_cert_subject_new" --subject-key-id="$root_cert_subject_key_id_new" --from="$trustee_account_4" --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Propose revoke test_root_certificate"
result=$(echo $passphrase | dcld tx pki propose-revoke-x509-root-cert --subject="$test_root_cert_subject_new" --subject-key-id="$test_root_cert_subject_key_id_new" --from $trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

# PKI Revocation point

echo "Add new revocaton point for a old vid"
result=$(echo $passphrase | tx pki add-revocation-point --vid=$vid --is-paa="true" --certificate="$test_root_cert_path" --label="$product_label" --data-url="$test_data_url" --issuer-subject-key-id=$test_root_cert_subject_key_id --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Add new revocaton point for a new vid"
result=$(echo $passphrase | tx pki add-revocation-point --vid=$vid_new --is-paa="true" --certificate="$test_root_cert_path_new" --label="$product_label_new" --data-url="$test_data_url" --issuer-subject-key-id=$test_root_cert_subject_key_id_new --from=$vendor_account_new --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Update revocaton point for an old vid"
result=$(echo $passphrase | tx pki update-revocation-point --vid=$vid --certificate="$test_root_cert_path" --label="$product_label" --data-url="$test_data_url" --issuer-subject-key-id=$test_root_cert_subject_key_id --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Update revocaton point for a new vid"
result=$(echo $passphrase | tx pki update-revocation-point --vid=$vid_new --certificate="$test_root_cert_path_new" --label="$product_label_new" --data-url="$test_data_url" --issuer-subject-key-id=$test_root_cert_subject_key_id_new --from=$vendor_account_new --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Delete revocaton point for the old vid"
result=$(echo $passphrase | tx pki delete-revocation-point --vid=$vid --label="$product_label" --issuer-subject-key-id=$test_root_cert_subject_key_id --from=$vendor_account --yes)
check_response "$result" "\"code\": 0"

test_divider

# AUTH

echo "Propose add account $user_4_address"
result=$(echo $passphrase | dcld tx auth propose-add-account --address="$user_4_address" --pubkey="$user_4_pubkey" --roles="CertificationCenter" --from="$trustee_account_1" --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_4_address"
result=$(dcld tx auth approve-add-account --address="$user_4_address" --from=$trustee_account_2 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_4_address"
result=$(dcld tx auth approve-add-account --address="$user_4_address" --from=$trustee_account_3 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_4_address"
result=$(dcld tx auth approve-add-account --address="$user_4_address" --from=$trustee_account_4 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Propose add account $user_5_address"
result=$(echo $passphrase | dcld tx auth propose-add-account --address="$user_5_address" --pubkey=$user_5_pubkey --roles=CertificationCenter --from=$trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_5_address"
result=$(echo $passphrase | dcld tx auth approve-add-account --address="$user_5_address" --from=$trustee_account_2 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_5_address"
result=$(dcld tx auth approve-add-account --address="$user_5_address" --from=$trustee_account_3 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve add account $user_5_address"
result=$(dcld tx auth approve-add-account --address="$user_5_address" --from=$trustee_account_4 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Propose add account $user_6_address"
result=$(echo $passphrase | dcld tx auth propose-add-account --address="$user_6_address" --pubkey=$user_6_pubkey --roles=CertificationCenter --from=$trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Propose revoke account $user_4_address"
result=$(echo $passphrase | dcld tx auth propose-revoke-account --address="$user_4_address" --from=$trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve revoke account $user_4_address"
result=$(echo $passphrase | dcld tx auth approve-revoke-account --address="$user_4_address" --from=$trustee_account_2 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve revoke account $user_4_address"
result=$(echo $passphrase | dcld tx auth approve-revoke-account --address="$user_4_address" --from=$trustee_account_3 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve revoke account $user_4_address"
result=$(echo $passphrase | dcld tx auth approve-revoke-account --address="$user_4_address" --from=$trustee_account_4 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Propose revoke account $user_5_address"
result=$(echo $passphrase | dcld tx auth propose-revoke-account --address="$user_5_address" --from=$trustee_account_1 --yes)
check_response "$result" "\"code\": 0"

test_divider

# VALIDATOR_NODE
echo "Disable node"
result=$(docker exec "$container" /bin/sh -c "echo test1234  | dcld tx validator disable-node --from=$account --yes")
check_response "$result" "\"code\": 0"

test_divider

echo "Enable node"
result=$(docker exec "$container" /bin/sh -c "echo test1234  | dcld tx validator enable-node --from=$account --yes")
check_response "$result" "\"code\": 0"

test_divider

echo "Approve disable node"
result=$(echo $passphrase | dcld tx validator approve-disable-node --address=$validator_address --from=$trustee_account_2 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve disable node"
result=$(echo $passphrase | dcld tx validator approve-disable-node --address=$validator_address --from=$trustee_account_3 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Approve disable node"
result=$(echo $passphrase | dcld tx validator approve-disable-node --address=$validator_address --from=$trustee_account_4 --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Enable node"
result=$(docker exec "$container" /bin/sh -c "echo test1234  | dcld tx validator enable-node --from=$account --yes")
check_response "$result" "\"code\": 0"

test_divider

echo "Verify that new data is not corrupted"

test_divider

# VENDORINFO

echo "Verify if VendorInfo Record for VID: $vid_new is present or not"
result=$(dcld query vendorinfo vendor --vid=$vid_new)
check_response "$result" "\"vendorID\": $vid_new"
check_response "$result" "\"companyLegalName\": \"$company_legal_name_new\""

echo "Verify if VendorInfo Record for VID: $vid updated or not"
result=$(dcld query vendorinfo vendor --vid=$vid)
check_response "$result" "\"vendorID\": $vid"
check_response "$result" "\"vendorName\": \"$vendor_name\""
check_response "$result" "\"companyPreferredName\": \"$company_preferred_name_new\""
check_response "$result" "\"vendorLandingPageURL\": \"$vendor_landing_page_url_new\""

echo "Request all vendor infos"
result=$(dcld query vendorinfo all-vendors)
check_response "$result" "\"vendorID\": $vid_new"
check_response "$result" "\"companyLegalName\": \"$company_legal_name_new\""
check_response "$result" "\"vendorName\": \"$vendor_name_new\""

test_divider

# MODEL

echo "Get Model with VID: $vid_new PID: $pid_1_new"
result=$(dcld query model get-model --vid=$vid_new --pid=$pid_1_new)
check_response "$result" "\"vid\": $vid_new"
check_response "$result" "\"pid\": $pid_1_new"
check_response "$result" "\"productLabel\": \"$product_label_new\""

echo "Get Model with VID: $vid_new PID: $pid_2_new"
result=$(dcld query model get-model --vid=$vid_new --pid=$pid_2_new)
check_response "$result" "\"vid\": $vid_new"
check_response "$result" "\"pid\": $pid_2_new"
check_response "$result" "\"productLabel\": \"$product_label_new\""

echo "Check Model with VID: $vid_new PID: $pid_2_new updated"
result=$(dcld query model get-model --vid=$vid --pid=$pid_2)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_2"
check_response "$result" "\"productLabel\": \"$product_label_new\""
check_response "$result" "\"partNumber\": \"$part_number_new\""

echo "Check Model version with VID: $vid_new PID: $pid_2_new updated"
result=$(dcld query model model-version --vid=$vid --pid=$pid_2  --softwareVersion=$software_version)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid_2"
check_response "$result" "\"minApplicableSoftwareVersion\": $min_applicable_software_version_new"
check_response "$result" "\"maxApplicableSoftwareVersion\": $max_applicable_software_version_new"

echo "Get all models"
result=$(dcld query model all-models)
check_response "$result" "\"vid\": $vid_new"
check_response "$result" "\"pid\": $pid_1_new"
check_response "$result" "\"pid\": $pid_2_new"

echo "Get all model versions"
result=$(dcld query model all-model-versions --vid=$vid_new --pid=$pid_1_new)
check_response "$result" "\"vid\": $vid_new"
check_response "$result" "\"pid\": $pid_1_new"

echo "Get Vendor Models with VID: ${vid_new}"
result=$(dcld query model vendor-models --vid=$vid_new)
check_response "$result" "\"pid\": $pid_1_new"
check_response "$result" "\"pid\": $pid_2_new"

echo "Get model version VID: $vid_new PID: $pid_1_new"
result=$(dcld query model model-version --vid=$vid_new --pid=$pid_1_new --softwareVersion=$software_version_new)
check_response "$result" "\"vid\": $vid_new"
check_response "$result" "\"pid\": $pid_1_new"
check_response "$result" "\"softwareVersion\": $software_version_new"

echo "Get model version VID: $vid_new PID: $pid_2_new"
result=$(dcld query model model-version --vid=$vid_new --pid=$pid_2_new --softwareVersion=$software_version_new)
check_response "$result" "\"vid\": $vid_new"
check_response "$result" "\"pid\": $pid_2_new"
check_response "$result" "\"softwareVersion\": $software_version_new"

test_divider

# COMPLIANCE

echo "Get certified model vid=$vid_new pid=$pid_1_new"
result=$(dcld query compliance certified-model --vid=$vid_new --pid=$pid_1_new --softwareVersion=$software_version_new --certificationType=$certification_type_new)
check_response "$result" "\"value\": true"
check_response "$result" "\"vid\": $vid_new"
check_response "$result" "\"pid\": $pid_1_new"
check_response "$result" "\"softwareVersion\": $software_version_new"
check_response "$result" "\"certificationType\": \"$certification_type_new\""

echo "Get revoked Model with VID: $vid_new PID: $pid_2_new"
result=$(dcld query compliance revoked-model --vid=$vid_new --pid=$pid_2_new --softwareVersion=$software_version_new --certificationType=$certification_type_new)
check_response "$result" "\"vid\": $vid_new"
check_response "$result" "\"pid\": $pid_2_new"

echo "Get certified model with VID: $vid_new PID: $pid_1_new"
result=$(dcld query compliance certified-model --vid=$vid_new --pid=$pid_1_new --softwareVersion=$software_version_new --certificationType=$certification_type_new)
check_response "$result" "\"value\": true"
check_response "$result" "\"vid\": $vid_new"
check_response "$result" "\"pid\": $pid_1_new"

echo "Get provisional model with VID: $vid_new PID: $pid_3_new"
result=$(dcld query compliance provisional-model --vid=$vid_new --pid=$pid_3_new --softwareVersion=$software_version_new --certificationType=$certification_type_new)
check_response "$result" "\"value\": true"
check_response "$result" "\"vid\": $vid_new"
check_response "$result" "\"pid\": $pid_3_new"

echo "Get compliance-info model with VID: $vid_new PID: $pid_1_new"
result=$(dcld query compliance compliance-info --vid=$vid_new --pid=$pid_1_new --softwareVersion=$software_version_new --certificationType=$certification_type_new)
check_response "$result" "\"vid\": $vid_new"
check_response "$result" "\"pid\": $pid_1_new"
check_response "$result" "\"softwareVersion\": $software_version_new"
check_response "$result" "\"certificationType\": \"$certification_type_new\""

echo "Get compliance-info model with VID: $vid_new PID: $pid_2_new"
result=$(dcld query compliance compliance-info --vid=$vid_new --pid=$pid_2_new --softwareVersion=$software_version_new --certificationType=$certification_type_new)
check_response "$result" "\"vid\": $vid_new"
check_response "$result" "\"pid\": $pid_2_new"
check_response "$result" "\"softwareVersion\": $software_version_new"
check_response "$result" "\"certificationType\": \"$certification_type_new\""

echo "Get device software compliance cDCertificateId=$cd_certificate_id_new"
result=$(dcld query compliance device-software-compliance --cdCertificateId=$cd_certificate_id_new)
check_response "$result" "\"vid\": $vid_new"
check_response "$result" "\"pid\": $pid_1_new"

echo "Get all certified models"
result=$(dcld query compliance all-certified-models)
check_response "$result" "\"vid\": $vid_new"
check_response "$result" "\"pid\": $pid_1_new"

echo "Get all provisional models"
result=$(dcld query compliance all-provisional-models)
check_response "$result" "\"vid\": $vid_new"
check_response "$result" "\"pid\": $pid_3_new"

echo "Get all revoked models"
result=$(dcld query compliance all-revoked-models)
check_response "$result" "\"vid\": $vid_new"
check_response "$result" "\"pid\": $pid_2_new"

echo "Get all compliance infos"
result=$(dcld query compliance all-compliance-info)
check_response "$result" "\"vid\": $vid_new"
check_response "$result" "\"pid\": $pid_1_new"
check_response "$result" "\"pid\": $pid_2_new"

echo "Get all device software compliances"
result=$(dcld query compliance all-device-software-compliance)
check_response "$result" "\"vid\": $vid_new"
check_response "$result" "\"pid\": $pid_1_new"
check_response "$result" "\"cDCertificateId\": \"$cd_certificate_id_new\""

test_divider

# PKI

echo "Get x509 root certificate"
result=$(dcld query pki x509-cert --subject=$test_root_cert_subject_new --subject-key-id=$test_root_cert_subject_key_id_new)
check_response "$result" "\"subject\": \"$test_root_cert_subject_new\""
check_response "$result" "\"subjectKeyId\": \"$test_root_cert_subject_key_id_new\""

echo "Get all subject x509 root certificates"
result=$(dcld query pki all-subject-x509-certs --subject=$test_root_cert_subject_new)
check_response "$result" "\"subject\": \"$test_root_cert_subject_new\""
check_response "$result" "$test_root_cert_subject_key_id_new"

echo "Get proposed x509 root certificate"
result=$(dcld query pki proposed-x509-root-cert --subject=$google_root_cert_subject_new --subject-key-id=$google_root_cert_subject_key_id_new)
check_response "$result" "\"subject\": \"$google_root_cert_subject_new\""
check_response "$result" "\"subjectKeyId\": \"$google_root_cert_subject_key_id_new\""

echo "Get revoked x509 certificate"
result=$(dcld query pki revoked-x509-cert --subject=$intermediate_cert_subject_new --subject-key-id=$intermediate_cert_subject_key_id_new)
check_response "$result" "\"subject\": \"$intermediate_cert_subject_new\""
check_response "$result" "\"subjectKeyId\": \"$intermediate_cert_subject_key_id_new\""

echo "Get proposed x509 root certificate to revoke"
result=$(dcld query pki proposed-x509-root-cert-to-revoke --subject=$test_root_cert_subject_new --subject-key-id=$test_root_cert_subject_key_id_new)
check_response "$result" "\"subject\": \"$test_root_cert_subject_new\""
check_response "$result" "\"subjectKeyId\": \"$test_root_cert_subject_key_id_new\""

echo "Get revocation point"
result=$(dcld query pki revocation-point --vid=$vid_new --label=$product_label_new --issuer-subject-key-id=$test_root_cert_subject_key_id_new)
check_response "$result" "\"vid\": \"$vid_new\""
check_response "$result" "\"issuerSubjectKeyID\": \"$test_root_cert_subject_key_id_new\""
check_response "$result" "\"label\": \"$product_label_new\""
check_response "$result" "\"dataUrl\": \"$test_data_url\""

echo "Get revocation points by issuer subject key id"
result=$(dcld query pki revocation-point --issuer-subject-key-id=$test_root_cert_subject_key_id_new)
check_response "$result" "\"vid\": \"$vid_new\""
check_response "$result" "\"issuerSubjectKeyID\": \"$test_root_cert_subject_key_id_new\""
check_response "$result" "\"label\": \"$product_label_new\""
check_response "$result" "\"dataUrl\": \"$test_data_url\""

echo "Get all proposed x509 root certificates"
result=$(dcld query pki all-proposed-x509-root-certs)
check_response "$result" "\"subject\": \"$google_root_cert_subject_new\""
check_response "$result" "\"subjectKeyId\": \"$google_root_cert_subject_key_id_new\""

echo "Get all revoked x509 root certificates"
result=$(dcld query pki all-revoked-x509-root-certs)
check_response "$result" "\"subject\": \"$root_cert_subject_new\""
check_response "$result" "\"subjectKeyId\": \"$root_cert_subject_key_id_new\""

echo "Get all proposed x509 root certificates to revoke"
result=$(dcld query pki all-proposed-x509-root-certs-to-revoke)
check_response "$result" "\"subject\": \"$test_root_cert_subject_new\""
check_response "$result" "\"subjectKeyId\": \"$test_root_cert_subject_key_id_new\""

echo "Get all x509 certificates"
result=$(dcld query pki all-x509-certs)
check_response "$result" "\"subject\": \"$test_root_cert_subject_new\""
check_response "$result" "\"subjectKeyId\": \"$test_root_cert_subject_key_id_new\""

echo "Get all revocation points"
result=$(dcld query pki all-revocation-points)
check_response "$result" "\"vid\": \"$vid_new\""
check_response "$result" "\"issuerSubjectKeyID\": \"$test_root_cert_subject_key_id_new\""
check_response "$result" "\"label\": \"$product_label_new\""
check_response "$result" "\"dataUrl\": \"$test_data_url\""

test_divider

# AUTH

echo "Get all accounts"
result=$(dcld query auth all-accounts)
check_response "$result" "\"address\": \"$user_5_address\""

echo "Get account"
result=$(dcld query auth account --address=$user_5_address)
check_response "$result" "\"address\": \"$user_5_address\""

echo "Get all proposed accounts"
result=$(dcld query auth all-proposed-accounts)
check_response "$result" "\"address\": \"$user_6_address\""

echo "Get proposed account"
result=$(dcld query auth proposed-account --address=$user_6_address)
check_response "$result" "\"address\": \"$user_6_address\""

echo "Get all proposed accounts to revoke"
result=$(dcld query auth all-proposed-accounts-to-revoke)
check_response "$result" "\"address\": \"$user_5_address\""

echo "Get proposed account to revoke"
result=$(dcld query auth proposed-account-to-revoke --address=$user_5_address)
check_response "$result" "\"address\": \"$user_5_address\""

echo "Get all revoked accounts"
result=$(dcld query auth all-revoked-accounts)
check_response "$result" "\"address\": \"$user_4_address\""

echo "Get revoked account"
result=$(dcld query auth revoked-account --address=$user_4_address)
check_response "$result" "\"address\": \"$user_4_address\""

test_divider

# Validator

echo "Get node"
result=$(docker exec "$container" /bin/sh -c "echo test1234 | dcld query validator all-nodes")
check_response "$result" "\"owner\": \"$validator_address\""

echo "PASSED"
