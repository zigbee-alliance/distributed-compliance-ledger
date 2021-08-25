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

# client

dclcli config chain-id dclchain
dclcli config output json
dclcli config indent true
dclcli config trust-node false

echo 'test1234' | dclcli keys add jack
echo 'test1234' | dclcli keys add alice
echo 'test1234' | dclcli keys add bob

# node

dcld init node0 --chain-id dclchain

dcld add-genesis-account --address=$(dclcli keys show jack -a) --pubkey=$(dclcli keys show jack -p) --roles="Trustee,NodeAdmin"
dcld add-genesis-account --address=$(dclcli keys show alice -a) --pubkey=$(dclcli keys show alice -p) --roles="Trustee,NodeAdmin,Vendor"
dcld add-genesis-account --address=$(dclcli keys show bob -a) --pubkey=$(dclcli keys show bob -p) --roles="Trustee,NodeAdmin"

echo 'test1234' | dcld gentx --from jack

dcld collect-gentxs
dcld validate-genesis
