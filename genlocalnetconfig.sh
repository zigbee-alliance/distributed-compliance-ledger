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

DCL_DIR="$HOME/.dcl"
KEYPASSWD=test1234
CHAIN_ID=dclchain

rm -rf "$DCL_DIR"

rm -rf "$LOCALNET_DIR"
mkdir "$LOCALNET_DIR" "$LOCALNET_DIR"/{client,node0,node1,node2,node3}

if [[ -n "$DCL_OBSERVERS" ]]; then
    mkdir "$LOCALNET_DIR/observer0"
fi

# client

dcld config chain-id "$CHAIN_ID"
dcld config output json
# TODO issue 99: check the replacement for the setting
# dcld config indent true
# TODO issue 99: check the replacement for the setting
# dcld config trust-node false
dcld config keyring-backend file


cp -r ~/.dclcli/* "$LOCALNET_DIR/client"

(echo "$KEYPASSWD"; echo "$KEYPASSWD") | dcld keys add jack
(echo "$KEYPASSWD"; echo "$KEYPASSWD") | dcld keys add alice
(echo "$KEYPASSWD"; echo "$KEYPASSWD") | dcld keys add bob
(echo "$KEYPASSWD"; echo "$KEYPASSWD") | dcld keys add anna

# common keyring (client) data for all the nodes
# TODO issue 99: do we need all the keys on all the nodes
mv ~/.dcl/* localnet/client

jack_address=$(echo "$KEYPASSWD" | dcld keys show jack -a)
jack_pubkey=$(echo "$KEYPASSWD" | dcld keys show jack -p)

alice_address=$(echo "$KEYPASSWD" | dcld keys show alice -a)
alice_pubkey=$(echo "$KEYPASSWD" | dcld keys show alice -p)

bob_address=$(echo "$KEYPASSWD" | dcld keys show bob -a)
bob_pubkey=$(echo "$KEYPASSWD" | dcld keys show bob -p)

anna_address=$(echo "$KEYPASSWD" | dcld keys show anna -a)
anna_pubkey=$(echo "$KEYPASSWD" | dcld keys show anna -p)

# node 0

dcld init node0 --chain-id "$CHAIN_ID"
cp -R localnet/client/* ~/.dcl

# TODO issue 99: do we really need to create multiple the same 4 accounts per each node
dcld add-genesis-account --address="$jack_address" --pubkey="$jack_pubkey" --roles="Trustee,NodeAdmin"
dcld add-genesis-account --address=$alice_address --pubkey=$alice_pubkey --roles="Trustee,NodeAdmin"
dcld add-genesis-account --address=$bob_address --pubkey=$bob_pubkey --roles="Trustee,NodeAdmin"
dcld add-genesis-account --address=$anna_address --pubkey=$anna_pubkey --roles="NodeAdmin"

echo "$KEYPASSWD" | dcld gentx jack --chain-id "$CHAIN_ID"

mv ~/.dcld/* "$LOCALNET_DIR/node0"

# node 1

dcld init node1 --chain-id "$CHAIN_ID"
cp -R localnet/client/* ~/.dcl

dcld add-genesis-account --address=$jack_address --pubkey=$jack_pubkey --roles="Trustee,NodeAdmin"
dcld add-genesis-account --address=$alice_address --pubkey=$alice_pubkey --roles="Trustee,NodeAdmin"
dcld add-genesis-account --address=$bob_address --pubkey=$bob_pubkey --roles="Trustee,NodeAdmin"
dcld add-genesis-account --address=$anna_address --pubkey=$anna_pubkey --roles="NodeAdmin"

echo "$KEYPASSWD" | dcld gentx alice --chain-id "$CHAIN_ID"

mv ~/.dcld/* "$LOCALNET_DIR/node1"

# node 2

dcld init node2 --chain-id "$CHAIN_ID"
cp -R localnet/client/* ~/.dcl

dcld add-genesis-account --address=$jack_address --pubkey=$jack_pubkey --roles="Trustee,NodeAdmin"
dcld add-genesis-account --address=$alice_address --pubkey=$alice_pubkey --roles="Trustee,NodeAdmin"
dcld add-genesis-account --address=$bob_address --pubkey=$bob_pubkey --roles="Trustee,NodeAdmin"
dcld add-genesis-account --address=$anna_address --pubkey=$anna_pubkey --roles="NodeAdmin"

echo "$KEYPASSWD" | dcld gentx bob --chain-id "$CHAIN_ID"

mv ~/.dcld/* "$LOCALNET_DIR/node2"

# node 3

dcld init node3 --chain-id "$CHAIN_ID"
cp -R localnet/client/* ~/.dcl

dcld add-genesis-account --address=$jack_address --pubkey=$jack_pubkey --roles="Trustee,NodeAdmin"
dcld add-genesis-account --address=$alice_address --pubkey=$alice_pubkey --roles="Trustee,NodeAdmin"
dcld add-genesis-account --address=$bob_address --pubkey=$bob_pubkey --roles="Trustee,NodeAdmin"
dcld add-genesis-account --address=$anna_address --pubkey=$anna_pubkey --roles="NodeAdmin"

echo "$KEYPASSWD" | dcld gentx anna --chain-id "$CHAIN_ID"

cp -r ~/.dcld/* "$LOCALNET_DIR/node3"

if [[ -d "$LOCALNET_DIR/observer0" ]]; then
    rm -rf ~/.dcld/*
    # observer0

    dcld init observer0 --chain-id "$CHAIN_ID"
    cp -R localnet/client/* ~/.dcl

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
