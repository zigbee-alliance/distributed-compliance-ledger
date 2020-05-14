rm -rf ~/.zblcli
rm -rf ~/.zbld

rm -rf localnet
mkdir localnet localnet/client localnet/node0 localnet/node1 localnet/node2 localnet/node3

# client

zblcli config chain-id zblchain
zblcli config output json
zblcli config indent true
zblcli config trust-node false

echo 'test1234' | zblcli keys add jack
echo 'test1234' | zblcli keys add alice
echo 'test1234' | zblcli keys add bob
echo 'test1234' | zblcli keys add anna

cp -r ~/.zblcli/* localnet/client

# node 0

zbld init node0 --chain-id zblchain

jack_address=$(zblcli keys show jack -a)
jack_pubkey=$(zblcli keys show jack -p)

alice_address=$(zblcli keys show alice -a)
alice_pubkey=$(zblcli keys show alice -p)

bob_address=$(zblcli keys show bob -a)
bob_pubkey=$(zblcli keys show bob -p)

anna_address=$(zblcli keys show anna -a)
anna_pubkey=$(zblcli keys show anna -p)

zbld add-genesis-account --address=$jack_address --pubkey=$jack_pubkey --roles="Trustee,NodeAdmin"
zbld add-genesis-account --address=$alice_address --pubkey=$alice_pubkey --roles="NodeAdmin"
zbld add-genesis-account --address=$bob_address --pubkey=$bob_pubkey --roles="NodeAdmin"
zbld add-genesis-account --address=$anna_address --pubkey=$anna_pubkey --roles="NodeAdmin"

echo 'test1234' | zbld gentx --from jack

mv ~/.zbld/* localnet/node0

# node 1

zbld init node1 --chain-id zblchain

zbld add-genesis-account --address=$jack_address --pubkey=$jack_pubkey --roles="Trustee,NodeAdmin"
zbld add-genesis-account --address=$alice_address --pubkey=$alice_pubkey --roles="NodeAdmin"
zbld add-genesis-account --address=$bob_address --pubkey=$bob_pubkey --roles="NodeAdmin"
zbld add-genesis-account --address=$anna_address --pubkey=$anna_pubkey --roles="NodeAdmin"

echo 'test1234' | zbld gentx --from alice

mv ~/.zbld/* localnet/node1

# node 2

zbld init node2 --chain-id zblchain

zbld add-genesis-account --address=$jack_address --pubkey=$jack_pubkey --roles="Trustee,NodeAdmin"
zbld add-genesis-account --address=$alice_address --pubkey=$alice_pubkey --roles="NodeAdmin"
zbld add-genesis-account --address=$bob_address --pubkey=$bob_pubkey --roles="NodeAdmin"
zbld add-genesis-account --address=$anna_address --pubkey=$anna_pubkey --roles="NodeAdmin"

echo 'test1234' | zbld gentx --from bob

mv ~/.zbld/* localnet/node2

# node 3

zbld init node3 --chain-id zblchain

zbld add-genesis-account --address=$jack_address --pubkey=$jack_pubkey --roles="Trustee,NodeAdmin"
zbld add-genesis-account --address=$alice_address --pubkey=$alice_pubkey --roles="NodeAdmin"
zbld add-genesis-account --address=$bob_address --pubkey=$bob_pubkey --roles="NodeAdmin"
zbld add-genesis-account --address=$anna_address --pubkey=$anna_pubkey --roles="NodeAdmin"

echo 'test1234' | zbld gentx --from anna

cp -r ~/.zbld/* localnet/node3

# Collect all validator creation transactions

cp localnet/node0/config/gentx/* ~/.zbld/config/gentx
cp localnet/node1/config/gentx/* ~/.zbld/config/gentx
cp localnet/node2/config/gentx/* ~/.zbld/config/gentx
cp localnet/node3/config/gentx/* ~/.zbld/config/gentx

# Embed them into genesis

zbld collect-gentxs
zbld validate-genesis

# Update genesis for all nodes

cp ~/.zbld/config/genesis.json localnet/node0/config/
cp ~/.zbld/config/genesis.json localnet/node1/config/
cp ~/.zbld/config/genesis.json localnet/node2/config/
cp ~/.zbld/config/genesis.json localnet/node3/config/

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