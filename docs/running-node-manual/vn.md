# Running a Validator Node manually

## Prerequisites
Make sure you have all [prerequisites](./prerequisites.md) set up
## Deployment steps

#### 1. Put `cosmovisor` binary to `/usr/bin/`, set proper owner and execution permissions.

#### 2. Locate the genesis app version to genesis application version directory:
- Create `$HOME/.dcl/cosmovisor/genesis/bin` directory.
- Copy `dcld` binary to it, set proper owner and execution permissions.
    Please note that execution permissions on `dcld` should be granted to all (i.e. User, Group and Others classes)
    because cosmovisor requires execution permission on the application binary to be granted to Others class.

3. Configure CLI:
    - `./dcld config chain-id <chain-id>`
      - Use `testnet` if you want to connect to the persistent Test Net
    - `./dcld config output json` - Output format (text/json).
    - `./dcld config node tcp://<host>:<port>` - Address of a node to connect.
    Choose one of the listed in `persistent_peers.txt` file.
    Example: `tcp://18.157.114.34:26657`.

#### 5. Create keys for a node admin and a trustee (optional) accounts

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
- Use `testnet` if you want to connect to the persistent Test Net

#### 7. (Optional) Consider enabling `state sync` in the configuration if you are joining long-running network
- For more information refer to [running-node-in-existing-network.md](../advanced/running-node-in-existing-network.md)

#### 8. Configure p2p and consensus parameters in `[~/.dcl/config.toml]` file:
  ```toml
  [p2p]
  pex = false
  addr_book_strict = false

  [consensus]
  create_empty_blocks = false
  create_empty_blocks_interval = "600s" # 10 mins
  ```

#### 9. (Optional) Enable `state sync` snapshots in`[~/.dcl/app.toml]` file:

  ```toml
  [state-sync]
  snapshot-interval = "snapshot-interval"
  snapshot-keep-recent = "snapshot-keep-recent"
  ```

#### *** Step 10 can be automated using `run_dcl_node` script
Run node:

```bash
./run_dcl_node -c "<chain-id>" "<node-name>"
```

> Notes:
>
> * the script assumes that:
>   * current user is going to be used for `cosmovisor` service to run as
>   * current user is in sudoers list

This command:

* properly locates `genesis.json`
* configures and starts the node

#### 10. Run node:
- Put `genesis.json` into dcld's config directory (usually `$HOME/.dcl/config/`).
    - Use `deployment/persistent_chains/testnet/genesis.json` if you want to connect to the persistent Test Net
- Open `$HOME/.dcl/config/config.toml` file in your favorite text editor:
    - Tell node how to connect to the network:
        - Set the value for the `persistent_peers` field as the content of `persistent_peers.txt` file.
        - Use `deployment/persistent_chains/testnet/persistent_peers.txt` if you want to connect to the persistent Test Net.
    - Make your node public:
        - Open `$HOME/.dcl/config/config.toml`
        - Find the line under `# TCP or UNIX socket address for the RPC server to listen on`
        - Change it to: `laddr = "tcp://0.0.0.0:26657"`
    - Optionally change other setting.
- Open `26656` (p2p) and `26657` (RPC) ports.
    - `sudo ufw allow 26656/tcp`
    - `sudo ufw allow 26657/tcp`
- Edit `cosmovisor.service`
    - Replace `ubuntu` with a username you want to start service on behalf
- Copy service configuration.
    - `cp cosmovisor.service /etc/systemd/system/`
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

#### 11. Add validator node to the network:
- Get this node's tendermint validator address: `./dcld tendermint show-address`.
    Expected output format:

    ```text
    cosmosvalcons1yrs697lxpwugy7h465wskwu2a5w9dgklx608f0
    ```

- Get this node's tendermint validator pubkey: `./dcld tendermint show-validator`.
    Expected output format:

    ```text
    cosmosvalconspub1zcjduepqcwg4eenpcxgs0269xuup5jlzj3pdquxlvj494cjxtqtcathsq7esfrsapa
    ```

- Note that *validator address* and *validator pubkey* are not the same as `address` and `pubkey` were used for account creation.

- Wait until the value of `catching_up` field gets to `false` value.

- Add validator node: `dcld tx validator add-node --pubkey=<validator pubkey> --moniker=<node name> --from=<key name>`.
If the transaction has been successfully written you would find `"code": 0` in the output JSON.

#### 12. Check the node is running and participates in consensus:
- Get the list of all nodes: `dcld query validator all-nodes`.
The node must present in the list and has the following params: `power:10` and `jailed:false`.

- Get the node status: `dcld status --node tcp://<node_ip>:26657`.
The value of `node ip` matches to `[rpc] laddr` field in `$HOME/.dcl/config/config.toml`
(TCP or UNIX socket address for the RPC server to listen on).  
Make sure that `result.sync_info.latest_block_height` is increasing over the time (once in about 10 mins).
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
    "network": "dclchain",
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
    - You can pass the additional value to get the result for a specific height: `dcld query tendermint-validator-set 100`  .

#### 13. Congrats! You are an owner of the validator node.