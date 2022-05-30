# Running a Genesis Validator Node manually

## Prerequisites
Make sure you have all [prerequisites](./prerequisites.md) set up
## Deployment steps

#### 1. Put `cosmovisor` binary to `/usr/bin/`, set proper owner and execution permissions.

#### 2. Locate the genesis app version to genesis application version directory:
- Create `$HOME/.dcl/cosmovisor/genesis/bin` directory.
- Copy `dcld` binary to it, set proper owner and execution permissions.
    Please note that execution permissions on `dcld` should be granted to all (i.e. User, Group and Others classes)
    because cosmovisor requires execution permission on the application binary to be granted to Others class.

#### 3. Choose the chain ID. Every network (for example, test-net, main-net, etc.) must have a unique chain ID.

#### 4. Configure CLI:
- `./dcld config chain-id <chain-id>` - the chosen unique chain ID.
- `./dcld config output json` - Output format (text/json).

#### 5. Create keys for a node admin and a trustee genesis accounts

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

#### 6. Initilize the node:

```bash
./dcld init "<node-name>" --chain-id "<chain-id>"
```

#### 7. Configure p2p and consensus parameters in `[~/.dcl/config.toml]` file:
  ```toml
  [p2p]
  pex = false
  addr_book_strict = false

  [consensus]
  create_empty_blocks = false
  create_empty_blocks_interval = "600s" # 10 mins
  ```

#### 8. (Optional) Enable `state sync` snapshots in`[~/.dcl/app.toml]` file:

  ```toml
  [state-sync]
  snapshot-interval = "snapshot-interval"
  snapshot-keep-recent = "snapshot-keep-recent"
  ```

#### *** Steps (9-10) can be automated using `run_dcl_node` script
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

#### 9. Prepare genesis node configuration:
- Add genesis account with the generated key and `Trustee`, `NodeAdmin` roles:
`./dcld add-genesis-account --address=<address> --pubkey=<pubkey> --roles="Trustee,NodeAdmin"`
- Optionally, add other genesis accounts using the same command.
- Create genesis transaction: `./dcld gentx --from <name>`, where `<name>` is the keys' name specified at Step 5.
- Collect genesis transactions: `./dcld collect-gentxs`.
- Validate genesis file: `./dcld validate-genesis`.
- Genesis file is located in `$HOME/.dcl/config/genesis.json`. Give this file to each new node admin.

#### 10. Run node:
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

#### 11. Check that genesis account is created:

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

#### 12. Check the node is running and participates in consensus:
- Get the list of all nodes: `dcld query validator all-nodes`.
The node must present in the list and has the following params: `power:10` and `jailed:false`.

- Get the node status: `dcld status --node tcp://<node_ip>:26657`.
The value of `node ip` matches to `[rpc] laddr` field in `$HOME/.dcl/config/config.toml`
(TCP or UNIX socket address for the RPC server to listen on).  
Make sure that `result.sync_info.latest_block_height` is increasing over the time (once in about 10 mins).

- Get the list of nodes participating in the consensus for the last block: `dcld query tendermint-validator-set`.
    - You can pass the additional value to get the result for a specific height: `dcld query tendermint-validator-set 100`.
    - As for the genesis state we have just 1 node, the command should return only our node at this phase.

#### 13. Congrats! You are an owner of the genesis validator node.