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
echo 'test1234' | zblcli keys add anna

# node

zbld init node0 --chain-id zblchain

zbld add-genesis-account $(zblcli keys show jack -a) 1000nametoken,100000000stake
zbld add-genesis-account $(zblcli keys show alice -a) 1000nametoken,100000000stake
zbld add-genesis-account $(zblcli keys show bob -a) 1000nametoken,100000000stake
zbld add-genesis-account $(zblcli keys show anna -a) 1000nametoken,100000000stake

echo 'test1234' | zbld gentx --name jack

zbld collect-gentxs
zbld validate-genesis