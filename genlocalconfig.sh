#!/bin/bash

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
