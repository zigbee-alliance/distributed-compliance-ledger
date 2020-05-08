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

zbld add-genesis-account $(zblcli keys show jack -a) 1000nametoken,100000000stake

jack_address=$(zblcli keys show jack -a)
sed -i 's/"account_roles": \[\]/"account_roles": \[{"address":'\""$jack_address\""',"roles":\[\"Trustee"\,\"NodeAdmin\"]}\]/' ~/.zbld/config/genesis.json

echo 'test1234' | zbld gentx --name jack

zbld collect-gentxs
zbld validate-genesis