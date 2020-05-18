## Running Node

This document describes in details how to configure a validator node, and add it to the existing pool.

### Hardware requirements

Minimal:
- 1GB RAM
- 25GB of disk space
- 1.4 GHz CPU

Recommended (for highload applications):
- 2GB RAM
- 100GB SSD
- x64 2.0 GHz 2v CPU

### Operating System

Current delivery is compiled and tested under `Ubuntu 18.04.3 LTS` so we recommend using this distribution for now. In future, it will be possible to compile the application for a wide range of operating systems thanks to Go language.

## Components

The delivery consists of the following components:

* Binary artifacts:
    * zbld: The binary used for running a node.
    * zblcli: The binary that allow users to interact with pool ledger.
* Genesis transactions file: `genesis.json`
* The list of alive peers: `persistent_peers.txt`. It has the following format: `<node id>@<node ip>,<node2 id>@<node2 ip>,...`.
* The service configuration file `zbld.service`

### Deployment steps

1. Put `zbld` and `zblcli` binaries to `/usr/bin/` and configure permissions.

2. Configure zblcli:
    * `zblcli config chain-id zblchain`
    * `zblcli config output json` - Output format (text/json).
    * `zblcli config indent true` - Add indent to JSON response.
    * `zblcli config trust-node false` - Verify proofs for node responses.
    * `zblcli config node <ip address>` - Address of a node to connect. 
    Choose one of the listed in `persistent_peers.txt` file. 
    Example: `tcp://18.157.114.34:26657`.

3. Prepare keys and account:
    * zblcli keys add <name> - Derive a new private key and encrypt to disk.
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
    * Copy generated `address` and `pubkey` and share them to any `Trustee`. 
    * `Trustee` will register the account on the ledger and assign `NodeAdmin` role.
    * In order to ensure that account is created and has assigned role you can use the command: 
    `zblcli query authnext account --address=<address>`.
    Expected output format: 
        ```json
        {
          "result": {
            "address": string, // bench32 encoded address
            "public_key": "string, // bench32 encoded public key
            "roles": [
              "NodeAdmin"
            ],
            "coins": [],
            "account_number": string,
            "sequence": string
          },
          "height": string
        }
        ```

4. Initialize the node and create the necessary config files:
    * Init Node: `zbld init <node name> --chain-id zblchain`.
    * Put `genesis.json` into zbld's config directory (usually `$HOME/.zbld/config/`).
    * Open `$HOME/.zbld/config/config.toml` file in your favorite text editor and 
    set the value for the `persistent_peers` field as the content of `persistent_peers.txt` file.
    * Open `26656` (p2p) and `26657` (RPC) ports. 
        * `sudo ufw allow 26656/tcp`
        * `sudo ufw allow 26657/tcp`
    * Edit `zbld.service`
        * Replace `ubuntu` with a user name you want to start service on behalf
    * Copy service configuration.
        * `cp zbld.service /etc/systemd/system/`
    * Optionally, edit `$HOME/.zbld/config/config.toml` in order to set different setting (like listen address).

5. Add validator node to the network:
   * Get this node's tendermint validator address: `zbld tendermint show-address`.
       Expected output format: 
           ```
           cosmosvalcons1yrs697lxpwugy7h465wskwu2a5w9dgklx608f0
           ```
   * Get this node's tendermint validator pubkey: `zbld tendermint show-validator`.
       Expected output format: 
           ```
           cosmosvalconspub1zcjduepqcwg4eenpcxgs0269xuup5jlzj3pdquxlvj494cjxtqtcathsq7esfrsapa
           ```
   * Note that *validator address* and *validator pubkey* are not the same as `address` and `pubkey` were used for account creation.
   
   * Add validator node: `zblcli tx validator add-node --validator-address=<validator address> --validator-pubkey=<validator pubkey> --name=<node name> --from=<key name>`.
   If the transaction has been successfully written you would find `"success": true` in the output JSON. 
   
   * Enable the service: `sudo systemctl enable zbld`
   * Start node: `sudo systemctl start zbld`
   
   * For testing purpose the node can be started in CLI mode: `zbld start` (instead of two previous `systemctl` commands).
   Service mode is recommended for demo and production environment.

6. Check the node is running and participates in consensus:
    * Get the list of all nodes: `zblcli query validator all-nodes`. 
    The node must present in the list and has the following params: `power:10` and `jailed:false`.

    * Get the node status: `zblcli status --node <node ip>`. 
    The value of `node ip` matches to `[rpc] laddr` field in `$HOME/.zbld/config/config.toml`
    (TCP or UNIX socket address for the RPC server to listen on).  
    Make sure that `result.sync_info.latest_block_height` is increasing over the time (once in about 5 sec).
       Expected output format: 
        ```json
        {
          "node_info": {
            "protocol_version": {
              "p2p": "7",
              "block": "10",
              "app": "0"
            },
            "id": string, // matches to prefix <ID> of the file: $HOME/.zbld/config/gentx/gentx-<ID>.json
            "listen_addr": "tcp://0.0.0.0:26656", // Address to listen for incoming connections. Matches to $HOME/.zbld/config/config.toml [p2p] `laddr` filed.
            "network": "zblchain",
            "version": "0.32.8",
            "channels": string,
            "moniker": string,
            "other": {
              "tx_index": "on",
              "rpc_address": "tcp://127.0.0.1:26657" // TCP or UNIX socket address for the RPC server to listen on. Matches to $HOME/.zbld/config/config.toml [rpc] `laddr` filed. 
            }
          },
          "sync_info": {
            "latest_block_hash": string,
            "latest_app_hash": "string,
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
    
    * Get the list of nodes participating in the consensus for the last block: `zblcli tendermint-validator-set`.
        * You can pass the additional value to get the result for a specific height: `zblcli tendermint-validator-set 100`  .
      
7. Congrats! You are an owner of the validator node.
