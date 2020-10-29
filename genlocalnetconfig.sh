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

rm -rf ~/.dclcli
rm -rf ~/.dcld

rm -rf localnet
mkdir localnet localnet/client localnet/node0 localnet/node1 localnet/node2 localnet/node3

# client

dclcli config chain-id dclchain
dclcli config output json
dclcli config indent true
dclcli config trust-node false

echo 'test1234' | dclcli keys add jack
echo 'test1234' | dclcli keys add alice
echo 'test1234' | dclcli keys add bob
echo 'test1234' | dclcli keys add anna

cp -r ~/.dclcli/* localnet/client

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

mv ~/.dcld/* localnet/node0

# node 1

dcld init node1 --chain-id dclchain

dcld add-genesis-account --address=$jack_address --pubkey=$jack_pubkey --roles="Trustee,NodeAdmin"
dcld add-genesis-account --address=$alice_address --pubkey=$alice_pubkey --roles="Trustee,NodeAdmin"
dcld add-genesis-account --address=$bob_address --pubkey=$bob_pubkey --roles="Trustee,NodeAdmin"
dcld add-genesis-account --address=$anna_address --pubkey=$anna_pubkey --roles="NodeAdmin"

echo 'test1234' | dcld gentx --from alice

mv ~/.dcld/* localnet/node1

# node 2

dcld init node2 --chain-id dclchain

dcld add-genesis-account --address=$jack_address --pubkey=$jack_pubkey --roles="Trustee,NodeAdmin"
dcld add-genesis-account --address=$alice_address --pubkey=$alice_pubkey --roles="Trustee,NodeAdmin"
dcld add-genesis-account --address=$bob_address --pubkey=$bob_pubkey --roles="Trustee,NodeAdmin"
dcld add-genesis-account --address=$anna_address --pubkey=$anna_pubkey --roles="NodeAdmin"

echo 'test1234' | dcld gentx --from bob

mv ~/.dcld/* localnet/node2

# node 3

dcld init node3 --chain-id dclchain

dcld add-genesis-account --address=$jack_address --pubkey=$jack_pubkey --roles="Trustee,NodeAdmin"
dcld add-genesis-account --address=$alice_address --pubkey=$alice_pubkey --roles="Trustee,NodeAdmin"
dcld add-genesis-account --address=$bob_address --pubkey=$bob_pubkey --roles="Trustee,NodeAdmin"
dcld add-genesis-account --address=$anna_address --pubkey=$anna_pubkey --roles="NodeAdmin"

echo 'test1234' | dcld gentx --from anna

cp -r ~/.dcld/* localnet/node3

# Collect all validator creation transactions

cp localnet/node0/config/gentx/* ~/.dcld/config/gentx
cp localnet/node1/config/gentx/* ~/.dcld/config/gentx
cp localnet/node2/config/gentx/* ~/.dcld/config/gentx
cp localnet/node3/config/gentx/* ~/.dcld/config/gentx

# Embed them into genesis

dcld collect-gentxs
dcld validate-genesis

# Update genesis for all nodes

cp ~/.dcld/config/genesis.json localnet/node0/config/
cp ~/.dcld/config/genesis.json localnet/node1/config/
cp ~/.dcld/config/genesis.json localnet/node2/config/
cp ~/.dcld/config/genesis.json localnet/node3/config/

# Find out node ids

id0=$(ls localnet/node0/config/gentx | sed 's/gentx-\(.*\).json/\1/')
id1=$(ls localnet/node1/config/gentx | sed 's/gentx-\(.*\).json/\1/')
id2=$(ls localnet/node2/config/gentx | sed 's/gentx-\(.*\).json/\1/')
id3=$(ls localnet/node3/config/gentx | sed 's/gentx-\(.*\).json/\1/')

# Update address book of the first node
peers="$id0@192.167.10.2:26656,$id1@192.167.10.3:26656,$id2@192.167.10.4:26656,$id3@192.167.10.5:26656"
sed -i "s/persistent_peers = \"\"/persistent_peers = \"$peers\"/g" localnet/node0/config/config.toml

# Make RPC enpoint available externally

sed -i 's/laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/g' localnet/node0/config/config.toml
sed -i 's/laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/g' localnet/node1/config/config.toml
sed -i 's/laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/g' localnet/node2/config/config.toml
sed -i 's/laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/g' localnet/node3/config/config.toml
