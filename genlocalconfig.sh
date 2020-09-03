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


rm -rf ~/.zblcli
rm -rf ~/.zbld

# client

zblcli config chain-id zblchain
zblcli config output json
zblcli config indent true
zblcli config trust-node false

echo 'test1234' | zblcli keys add jack
echo 'test1234' | zblcli keys add alice
echo 'test1234' | zblcli keys add bob

# node

zbld init node0 --chain-id zblchain

zbld add-genesis-account --address=$(zblcli keys show jack -a) --pubkey=$(zblcli keys show jack -p) --roles="Trustee,NodeAdmin"
zbld add-genesis-account --address=$(zblcli keys show alice -a) --pubkey=$(zblcli keys show alice -p) --roles="Trustee,NodeAdmin"
zbld add-genesis-account --address=$(zblcli keys show bob -a) --pubkey=$(zblcli keys show bob -p) --roles="Trustee,NodeAdmin"

echo 'test1234' | zbld gentx --from jack

zbld collect-gentxs
zbld validate-genesis
