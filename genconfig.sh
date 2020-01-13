rm -rf ~/.zblcli
rm -rf ~/.zbld

rm -rf testnet
mkdir testnet testnet/client testnet/node0 testnet/node1 testnet/node2 testnet/node3

# client

zblcli config chain-id zblchain
zblcli config output json
zblcli config indent true
zblcli config trust-node false

echo 'test1234' | zblcli keys add jack
echo 'test1234' | zblcli keys add alice
echo 'test1234' | zblcli keys add bob
echo 'test1234' | zblcli keys add anna

cp -r ~/.zblcli/* testnet/client

# node 0

zbld init node0 --chain-id zblchain

zbld add-genesis-account $(zblcli keys show jack -a) 1000nametoken,100000000stake
zbld add-genesis-account $(zblcli keys show alice -a) 1000nametoken,100000000stake
zbld add-genesis-account $(zblcli keys show bob -a) 1000nametoken,100000000stake
zbld add-genesis-account $(zblcli keys show anna -a) 1000nametoken,100000000stake

echo 'test1234' | zbld gentx --name jack

mv ~/.zbld/* testnet/node0

# node 1

zbld init node1 --chain-id zblchain

zbld add-genesis-account $(zblcli keys show jack -a) 1000nametoken,100000000stake
zbld add-genesis-account $(zblcli keys show alice -a) 1000nametoken,100000000stake
zbld add-genesis-account $(zblcli keys show bob -a) 1000nametoken,100000000stake
zbld add-genesis-account $(zblcli keys show anna -a) 1000nametoken,100000000stake

echo 'test1234' | zbld gentx --name alice

mv ~/.zbld/* testnet/node1

# node 2

zbld init node2 --chain-id zblchain

zbld add-genesis-account $(zblcli keys show jack -a) 1000nametoken,100000000stake
zbld add-genesis-account $(zblcli keys show alice -a) 1000nametoken,100000000stake
zbld add-genesis-account $(zblcli keys show bob -a) 1000nametoken,100000000stake
zbld add-genesis-account $(zblcli keys show anna -a) 1000nametoken,100000000stake

echo 'test1234' | zbld gentx --name bob

mv ~/.zbld/* testnet/node2

# node 3

zbld init node3 --chain-id zblchain

zbld add-genesis-account $(zblcli keys show jack -a) 1000nametoken,100000000stake
zbld add-genesis-account $(zblcli keys show alice -a) 1000nametoken,100000000stake
zbld add-genesis-account $(zblcli keys show bob -a) 1000nametoken,100000000stake
zbld add-genesis-account $(zblcli keys show anna -a) 1000nametoken,100000000stake

echo 'test1234' | zbld gentx --name anna

cp -r ~/.zbld/* testnet/node3

# collect all genesis transactions

cp testnet/node0/config/gentx/* ~/.zbld/config/gentx
cp testnet/node1/config/gentx/* ~/.zbld/config/gentx
cp testnet/node2/config/gentx/* ~/.zbld/config/gentx
cp testnet/node3/config/gentx/* ~/.zbld/config/gentx

zbld collect-gentxs
zbld validate-genesis

cp ~/.zbld/config/genesis.json testnet/node0/config/
cp ~/.zbld/config/genesis.json testnet/node1/config/
cp ~/.zbld/config/genesis.json testnet/node2/config/
cp ~/.zbld/config/genesis.json testnet/node3/config/
