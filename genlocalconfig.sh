rm -rf ~/.zblcli
rm -rf ~/.zbld

# client

zblcli config chain-id zblchain
zblcli config output json
zblcli config indent true
zblcli config trust-node false

echo 'test1234' | zblcli keys add jack

# node

zbld init node0 --chain-id zblchain

zbld add-genesis-account --address=$(zblcli keys show jack -a) --pubkey=$(zblcli keys show jack -p) --roles="Trustee,NodeAdmin"

echo 'test1234' | zbld gentx --from jack

zbld collect-gentxs
zbld validate-genesis
