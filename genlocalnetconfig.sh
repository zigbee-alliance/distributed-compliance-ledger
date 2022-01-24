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
KEYPASSWD=test1234  # NOTE not necessary actually since we yse 'test' keyring backend now
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
dcld config node "tcp://localhost:26657"
dcld config keyring-backend test
dcld config broadcast-mode block

(echo "$KEYPASSWD"; echo "$KEYPASSWD") | dcld keys add jack
(echo "$KEYPASSWD"; echo "$KEYPASSWD") | dcld keys add alice
(echo "$KEYPASSWD"; echo "$KEYPASSWD") | dcld keys add bob
(echo "$KEYPASSWD"; echo "$KEYPASSWD") | dcld keys add anna

# common keyring (client) data for all the nodes
# TODO issue 99: do we need all the keys on all the nodes
jack_address=$(echo "$KEYPASSWD" | dcld keys show jack -a)
jack_pubkey=$(echo "$KEYPASSWD" | dcld keys show jack -p)

alice_address=$(echo "$KEYPASSWD" | dcld keys show alice -a)
alice_pubkey=$(echo "$KEYPASSWD" | dcld keys show alice -p)

bob_address=$(echo "$KEYPASSWD" | dcld keys show bob -a)
bob_pubkey=$(echo "$KEYPASSWD" | dcld keys show bob -p)

anna_address=$(echo "$KEYPASSWD" | dcld keys show anna -a)
anna_pubkey=$(echo "$KEYPASSWD" | dcld keys show anna -p)

mv "$DCL_DIR"/* $LOCALNET_DIR/client


function add_genesis_accounts {
    dcld add-genesis-account --address="$jack_address" --pubkey="$jack_pubkey" --roles="Trustee,NodeAdmin"
    dcld add-genesis-account --address=$alice_address --pubkey=$alice_pubkey --roles="Trustee,NodeAdmin"
    dcld add-genesis-account --address=$bob_address --pubkey=$bob_pubkey --roles="Trustee,NodeAdmin"
    dcld add-genesis-account --address=$anna_address --pubkey=$anna_pubkey --roles="NodeAdmin"
}


function gentx {
    local _node_name="$1"
    local _key_name="$2"
    echo "$KEYPASSWD" | dcld gentx "$_key_name" --chain-id "$CHAIN_ID" --moniker "$_node_name"
}


function init_node {
    local _node_name="$1"
    local _key_name="${2:-}"
    local _copy_only="${3:-}"

    dcld init "$_node_name" --chain-id "$CHAIN_ID"
    cp -R "$LOCALNET_DIR"/client/* "$DCL_DIR"

    # we need to make them in an app state for each node
    add_genesis_accounts

    if [[ -n "$_key_name" ]]; then
        gentx "$_node_name" "$_key_name"
    fi

    if [[ -n "$_copy_only" ]]; then
        cp -r "$DCL_DIR"/* "$LOCALNET_DIR/$_node_name"
    else
        mv "$DCL_DIR"/* "$LOCALNET_DIR/$_node_name"
    fi
}


init_node node0 jack
init_node node1 alice
init_node node2 bob
init_node node3 anna yes


if [[ -d "$LOCALNET_DIR/observer0" ]]; then
    rm -rf "$DCL_DIR"/*
    init_node observer0 "" yes
fi

# Collect all validator creation transactions

mkdir -p "$DCL_DIR"/config/gentx
for node_name in node0 node1 node2 node3; do
    cp "$LOCALNET_DIR/$node_name"/config/gentx/* "$DCL_DIR"/config/gentx
done

# Embed them into genesis

dcld collect-gentxs
dcld validate-genesis

# Update genesis for all nodes

cp "$DCL_DIR"/config/genesis.json "$LOCALNET_DIR"

for node_name in node0 node1 node2 node3; do
    cp "$DCL_DIR"/config/genesis.json "$LOCALNET_DIR/$node_name/config/"
done

if [[ -d "$LOCALNET_DIR/observer0" ]]; then
    cp "$DCL_DIR"/config/genesis.json "$LOCALNET_DIR/observer0/config/"
fi

# Find out node ids

id0=$(ls "$LOCALNET_DIR/node0/config/gentx" | sed 's/gentx-\(.*\).json/\1/')
id1=$(ls "$LOCALNET_DIR/node1/config/gentx" | sed 's/gentx-\(.*\).json/\1/')
id2=$(ls "$LOCALNET_DIR/node2/config/gentx" | sed 's/gentx-\(.*\).json/\1/')
id3=$(ls "$LOCALNET_DIR/node3/config/gentx" | sed 's/gentx-\(.*\).json/\1/')

# Update address book of the first node
peers="$id0@192.167.10.2:26656,$id1@192.167.10.3:26656,$id2@192.167.10.4:26656,$id3@192.167.10.5:26656"

echo "$peers" >"$LOCALNET_DIR"/persistent_peers.txt

# Update address book of the first node 
sed -i $SED_EXT "s/persistent_peers = \"\"/persistent_peers = \"$peers\"/g" "$LOCALNET_DIR/node0/config/config.toml"
if [[ -d "$LOCALNET_DIR/observer0" ]]; then
    sed -i $SED_EXT "s/persistent_peers = \"\"/persistent_peers = \"$peers\"/g" "$LOCALNET_DIR/observer0/config/config.toml"
fi

for node_name in node0 node1 node2 node3 observer0; do
    if [[ -d "$LOCALNET_DIR/${node_name}" ]]; then
        # Make RPC endpoints available externally
        sed -i $SED_EXT 's/laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/g' "$LOCALNET_DIR/${node_name}/config/config.toml"

        # sets proper moniker
        sed -i $SED_EXT "s/moniker = .*/moniker = \"$node_name\"/g" "$LOCALNET_DIR/${node_name}/config/config.toml"

        # enables RPC and prometheus endpoints
        # FIXME issue 99: not good code
        sed -i $SED_EXT '0,/^enable = false/{s~enable = false~enable = true~}' "$LOCALNET_DIR/${node_name}/config/app.toml"

        # enables prometheus endpoints
        sed -i $SED_EXT 's/prometheus = false/prometheus = true/g' "$LOCALNET_DIR/${node_name}/config/config.toml"
    fi
done

# Init Light CLient Proxy if needed DCL_LIGHT_CLIENT_PROXY=1
# DCL_LIGHT_CLIENT_PROXY="${DCL_LIGHT_CLIENT_PROXY:-}"
DCL_LIGHT_CLIENT_PROXY=1

function init_light_client_proxy {
    local _node_name="$1"

    rm -rf "$DCL_DIR"

    dcld config chain-id "$CHAIN_ID"
    dcld config node "tcp://localhost:26657"
    dcld config broadcast-mode block
    dcld config keyring-backend test
    dcld config broadcast-mode block

    cp -R "$LOCALNET_DIR"/client/* "$DCL_DIR"

    cp -r "$DCL_DIR"/* "$LOCALNET_DIR/$_node_name"
}

if [[ -n "$DCL_LIGHT_CLIENT_PROXY" ]]; then
    mkdir "$LOCALNET_DIR/lightclient0"
    init_light_client_proxy lightclient0
fi
