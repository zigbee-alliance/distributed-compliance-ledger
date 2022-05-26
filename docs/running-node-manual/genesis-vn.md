# Running a Genesis Validator Node manually

## Prerequisites
Make sure you have all [prerequisites](./prerequisites.md) set up
## Deployment steps

### 1. Put `cosmovisor` binary to `/usr/bin/`, set proper owner and execution permissions.

### 2. Locate the genesis app version to genesis application version directory:
- Create `$HOME/.dcl/cosmovisor/genesis/bin` directory.
- Copy `dcld` binary to it, set proper owner and execution permissions.
    Please note that execution permissions on `dcld` should be granted to all (i.e. User, Group and Others classes)
    because cosmovisor requires execution permission on the application binary to be granted to Others class.

### 3. Choose the chain ID. Every network (for example, test-net, main-net, etc.) must have a unique chain ID.

### 4. Configure CLI:
- `./dcld config chain-id <chain-id>` - the chosen unique chain ID.
- `./dcld config output json` - Output format (text/json).

### 5. Create keys for a node admin and a trustee genesis accounts

```bash
./dcld keys add "<key-name>" 2>&1 | tee "<key-name>.dclkey.data"
```
Expected output format:

```json
{
    "name": <name>, // key name. can be used for signing transactions
    "type": "local",
    "address": string, // bench32 encoded address
    "pubkey": string, // bench32 encoded public key
    "mnemonic": string // seed that can be used to generate the same private/public key pair
}
```
- Remember generated `address` and `pubkey` they will be used later.
You can retrieve `address` and `pubkey` values anytime using `./dcld keys show <name>`.
Of course, only on the machine where the keypair was generated.

> Notes: It's important to keep the generated data (especially a mnemonic that allows to recover a key) in a safe place

### 6. Initilize the node:

```bash
./dcld init "<node-name>" --chain-id "<chain-id>"
```

### *** Steps (7-8) can be automated using `run_dcl_node` script
Run node:

```bash
./run_dcl_node -t genesis -c "<chain-id>" --gen-key-name "<node-admin-key>" [--gen-key-name-trustee "<trustee-key>"] "<node-name>"
```

This command:

* generates `genesis.json` file with the following entries:
  * a genesis account with `NodeAdmin` role
  * (if a trustee key is provided) a genesis account with `Trustee` role
  * a genesis txn that makes the local node a validator
* configures and starts the node

* the script assumes that:
  * current user is going to be used for `cosmovisor` service to run as
  * current user is in sudoers list
* you may likely want to note the summary that this script prints, in particular: node's address, public key and ID.

### 7. Prepare genesis node configuration:
- Add genesis account with the generated key and `Trustee`, `NodeAdmin` roles:
`./dcld add-genesis-account --address=<address> --pubkey=<pubkey> --roles="Trustee,NodeAdmin"`
- Optionally, add other genesis accounts using the same command.
- Create genesis transaction: `./dcld gentx --from <name>`, where `<name>` is the keys' name specified at Step 5.
- Collect genesis transactions: `./dcld collect-gentxs`.
- Validate genesis file: `./dcld validate-genesis`.
- Genesis file is located in `$HOME/.dcl/config/genesis.json`. Give this file to each new node admin.

### 8. Run node:
- Open `26656` (p2p) and `26657` (RPC) ports.
    - `sudo ufw allow 26656/tcp`
    - `sudo ufw allow 26657/tcp`
- Edit `cosmovisor.service`
    - Replace `ubuntu` with a username you want to start service on behalf
- Copy service configuration.
    - `cp cosmovisor.service /etc/systemd/system/`
- Make your node public:
    - Open `$HOME/.dcl/config/config.toml`
    - Find the line under `# TCP or UNIX socket address for the RPC server to listen on`
    - Change it to: `laddr = "tcp://0.0.0.0:26657"`
- Optionally, edit `$HOME/.dcl/config/config.toml` in order to set different setting (like listen address).
- Enable the service: `sudo systemctl enable cosmovisor`
- Start node: `sudo systemctl start cosmovisor`
- For testing purpose the node process can be started directly: `./dcld start` (instead of two previous `systemctl` commands using `cosmovisor` service).
Service mode is recommended for demo and production environment.

- Use `systemctl status cosmovisor` to get the node service status.
- Use `journalctl -u cosmovisor.service -f` to see node logs.
- You can also check node status by executing the command `./dcld status` to get the current status.
    The value of `latest_block_height` reflects the current node height.

- Add the following line to the end of `$HOME/.profile` file:
    - `export PATH=$PATH:$HOME/.dcl/cosmovisor/current/bin`
- Execute the following command to apply the updated `$PATH` immediately:
    - `source $HOME/.profile`

### 9. Check that genesis account is created:

- In order to ensure that account is created and has assigned role you can use the command:
`dcld query auth account --address=<address>`.
Expected output format:
```json
{
    "result": {
    "address": string, // bench32 encoded address
    "public_key": string, // bench32 encoded public key
    "roles": [
        "NodeAdmin", "Trustee"
    ],
    "coins": [],
    "account_number": string,
    "sequence": string
    },
    "height": string
}
```

### 10. Check the node is running and participates in consensus:
- Get the list of all nodes: `dcld query validator all-nodes`.
The node must present in the list and has the following params: `power:10` and `jailed:false`.

- Get the node status: `dcld status --node tcp://<node_ip>:26657`.
The value of `node ip` matches to `[rpc] laddr` field in `$HOME/.dcl/config/config.toml`
(TCP or UNIX socket address for the RPC server to listen on).  
Make sure that `result.sync_info.latest_block_height` is increasing over the time (once in about 5 sec).
<br>Expected output format:

```json
{
    "node_info": {
    "protocol_version": {
        "p2p": "7",
        "block": "10",
        "app": "0"
    },
    "id": string, // matches to prefix <ID> of the file: $HOME/.dcl/config/gentx/gentx-<ID>.json
    "listen_addr": "tcp://0.0.0.0:26656", // Address to listen for incoming connections. Matches to $HOME/.dcl/config/config.toml [p2p] `laddr` filed.
    "network": "<chain-id>",
    "version": "0.32.8",
    "channels": string,
    "moniker": string,
    "other": {
        "tx_index": "on",
        "rpc_address": "tcp://127.0.0.1:26657" // TCP or UNIX socket address for the RPC server to listen on. Matches to $HOME/.dcl/config/config.toml [rpc] `laddr` filed. 
    }
    },
    "sync_info": {
    "latest_block_hash": string,
    "latest_app_hash": string,
    "latest_block_height": string,
    "latest_block_time": string,
    "catching_up": bool
    },
    "validator_info": {
    "address": string,
    "pub_key": {
        "type": string,
        "value": string
    },
    "voting_power": string
    }
}
```

- Get the list of nodes participating in the consensus for the last block: `dcld query tendermint-validator-set`.
    - You can pass the additional value to get the result for a specific height: `dcld query tendermint-validator-set 100`.
    - As for the genesis state we have just 1 node, the command should return only our node at this phase.

### 11. Add more initial trusted validator nodes to the network

- Let's add more nodes that will be considered as trusted persistent peers by other nodes.
- Create a sub-folder in [deployment](../../deployment/persistent_chains)
with the chosen chain ID (Step 3) as a name.
- Persist the genesis file in the created chain-id subfolder.
The genesis file can be found in `$HOME/.dcl/config/genesis.json` (Step 6).
- Create the `persistent_peers.txt` file from the template located in [deployment](../../deployment)
and persist it in the chain-id subfolder.
    - If there is only one genesis node (like in the current tutorial), it will contain
    a single entry `<node1_id>@<node1_IP>:26656`.
    - `<node1_id>` - Node ID. Can be found as an `<ID>` prefix of the file `$HOME/.dcl/config/gentx/gentx-<ID>.json`.
    - `<node1-IP>` - public IP address of the genesis node.

- Every node that needs to join the network should follow the [Running Node](../running-node.md) instructions
using the chosen chain-id and the genesis and persistent_peers files above (or from the chain-id subfolder).
- Update the `persistent_peers.txt` file by including the entries for every added initial node.
- Update the `persistent_peers` field in `$HOME/.dcl/config/config.toml`
for every initial node (including the genesis one) to match the `persistent_peers.txt` content.
    - See Step 4 from [Running Node](../running-node.md) for details.

### 12. Adding more validator nodes to the network

- Just follow the [Running Node](../running-node.md) instructions
using the chosen chain-id and the genesis and persistent_peers files above (from the chain-id subfolder).
- Please make sure, that the persistent_peers contain all the nodes (including the genesis node)
added at Step 10.

### 13. Congrats! You are an owner of the genesis validator node.