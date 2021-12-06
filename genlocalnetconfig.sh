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

DCL_OBSERVERS="${DCL_OBSERVERS:-}"
LOCALNET_DIR=".localnet"

SED_EXT=
if [ "$(uname)" == "Darwin" ]; then
    # Mac OS X sed needs the file extension when -i flag is used. Keeping it empty as we don't need backupfile
    SED_EXT="''"
fi

rm -rf ~/.dclcli
rm -rf ~/.dcld

rm -rf "$LOCALNET_DIR"
mkdir "$LOCALNET_DIR" "$LOCALNET_DIR"/{client,node0,node1,node2,node3}

if [[ -n "$DCL_OBSERVERS" ]]; then
    mkdir "$LOCALNET_DIR/observer0"
fi

# client

dclcli config chain-id dclchain
dclcli config output json
dclcli config indent true
dclcli config trust-node false

echo 'test1234' | dclcli keys add jack
echo 'test1234' | dclcli keys add alice
echo 'test1234' | dclcli keys add bob
echo 'test1234' | dclcli keys add anna

cp -r ~/.dclcli/* "$LOCALNET_DIR/client"

# node 0

dcld init node0 --chain-id dclchain

jack_address=$(dclcli keys show jack -a)
jack_pubkey=$(dclcli keys show jack -p)

alice_address=$(dclcli keys show alice -a)
alice_pubkey=$(dclcli keys show alice -p)

bob_address=$(dclcli keys show bob -a)
bob_pubkey=$(dclcli keys show bob -p)

anna_address=$(dclcli keys show anna -a)
anna_pubkey=$(dclcli keys show anna -p)

dcld add-genesis-account --address=$jack_address --pubkey=$jack_pubkey --roles="Trustee,NodeAdmin"
dcld add-genesis-account --address=$alice_address --pubkey=$alice_pubkey --roles="Trustee,NodeAdmin"
dcld add-genesis-account --address=$bob_address --pubkey=$bob_pubkey --roles="Trustee,NodeAdmin"
dcld add-genesis-account --address=$anna_address --pubkey=$anna_pubkey --roles="NodeAdmin"

echo 'test1234' | dcld gentx --from jack

mv ~/.dcld/* "$LOCALNET_DIR/node0"

# node 1

dcld init node1 --chain-id dclchain

dcld add-genesis-account --address=$jack_address --pubkey=$jack_pubkey --roles="Trustee,NodeAdmin"
dcld add-genesis-account --address=$alice_address --pubkey=$alice_pubkey --roles="Trustee,NodeAdmin"
dcld add-genesis-account --address=$bob_address --pubkey=$bob_pubkey --roles="Trustee,NodeAdmin"
dcld add-genesis-account --address=$anna_address --pubkey=$anna_pubkey --roles="NodeAdmin"

echo 'test1234' | dcld gentx --from alice

mv ~/.dcld/* "$LOCALNET_DIR/node1"

# node 2

dcld init node2 --chain-id dclchain

dcld add-genesis-account --address=$jack_address --pubkey=$jack_pubkey --roles="Trustee,NodeAdmin"
dcld add-genesis-account --address=$alice_address --pubkey=$alice_pubkey --roles="Trustee,NodeAdmin"
dcld add-genesis-account --address=$bob_address --pubkey=$bob_pubkey --roles="Trustee,NodeAdmin"
dcld add-genesis-account --address=$anna_address --pubkey=$anna_pubkey --roles="NodeAdmin"

echo 'test1234' | dcld gentx --from bob

mv ~/.dcld/* "$LOCALNET_DIR/node2"

# node 3

dcld init node3 --chain-id dclchain

dcld add-genesis-account --address=$jack_address --pubkey=$jack_pubkey --roles="Trustee,NodeAdmin"
dcld add-genesis-account --address=$alice_address --pubkey=$alice_pubkey --roles="Trustee,NodeAdmin"
dcld add-genesis-account --address=$bob_address --pubkey=$bob_pubkey --roles="Trustee,NodeAdmin"
dcld add-genesis-account --address=$anna_address --pubkey=$anna_pubkey --roles="NodeAdmin"

echo 'test1234' | dcld gentx --from anna

cp -r ~/.dcld/* "$LOCALNET_DIR/node3"


if [[ -d "$LOCALNET_DIR/observer0" ]]; then
    rm -rf ~/.dcld/*
    # observer0

    dcld init observer0 --chain-id dclchain

    dcld add-genesis-account --address=$jack_address --pubkey=$jack_pubkey --roles="Trustee,NodeAdmin"
    dcld add-genesis-account --address=$alice_address --pubkey=$alice_pubkey --roles="Trustee,NodeAdmin"
    dcld add-genesis-account --address=$bob_address --pubkey=$bob_pubkey --roles="Trustee,NodeAdmin"
    dcld add-genesis-account --address=$anna_address --pubkey=$anna_pubkey --roles="NodeAdmin"

    cp -r ~/.dcld/* "$LOCALNET_DIR/observer0"
fi

# Collect all validator creation transactions

mkdir -p ~/.dcld/config/gentx
cp "$LOCALNET_DIR"/node0/config/gentx/* ~/.dcld/config/gentx
cp "$LOCALNET_DIR"/node1/config/gentx/* ~/.dcld/config/gentx
cp "$LOCALNET_DIR"/node2/config/gentx/* ~/.dcld/config/gentx
cp "$LOCALNET_DIR"/node3/config/gentx/* ~/.dcld/config/gentx

# Embed them into genesis

dcld collect-gentxs
dcld validate-genesis

# Update genesis for all nodes

cp ~/.dcld/config/genesis.json "$LOCALNET_DIR/node0/config/"
cp ~/.dcld/config/genesis.json "$LOCALNET_DIR/node1/config/"
cp ~/.dcld/config/genesis.json "$LOCALNET_DIR/node2/config/"
cp ~/.dcld/config/genesis.json "$LOCALNET_DIR/node3/config/"

if [[ -d "$LOCALNET_DIR/observer0" ]]; then
    cp ~/.dcld/config/genesis.json "$LOCALNET_DIR/observer0/config/"
fi

# Find out node ids

id0=$(ls "$LOCALNET_DIR/node0/config/gentx" | sed 's/gentx-\(.*\).json/\1/')
id1=$(ls "$LOCALNET_DIR/node1/config/gentx" | sed 's/gentx-\(.*\).json/\1/')
id2=$(ls "$LOCALNET_DIR/node2/config/gentx" | sed 's/gentx-\(.*\).json/\1/')
id3=$(ls "$LOCALNET_DIR/node3/config/gentx" | sed 's/gentx-\(.*\).json/\1/')

# Update address book of the first node
peers="$id0@192.167.10.2:26656,$id1@192.167.10.3:26656,$id2@192.167.10.4:26656,$id3@192.167.10.5:26656"

# Update address book of the first node 
sed -i $SED_EXT "s/persistent_peers = \"\"/persistent_peers = \"$peers\"/g" "$LOCALNET_DIR/node0/config/config.toml"
if [[ -d "$LOCALNET_DIR/observer0" ]]; then
    sed -i $SED_EXT "s/persistent_peers = \"\"/persistent_peers = \"$peers\"/g" "$LOCALNET_DIR/observer0/config/config.toml"
fi

# Make RPC endpoint available externally
for node_id in node0 node1 node2 node3 observer0; do
    if [[ -d "$LOCALNET_DIR/${node_id}" ]]; then
        sed -i $SED_EXT 's/laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/g' "$LOCALNET_DIR/${node_id}/config/config.toml"
        sed -i $SED_EXT 's/prometheus = false/prometheus = true/g' "$LOCALNET_DIR/${node_id}/config/config.toml"
    fi
done
