## Running a Validator Node

This document describes in details how to configure a validator node, and add it to the existing network.

The existing network can be either a custom one, or one of the persistent networks (such as a Test Net).
Configuration of all persistent networks can be found in [Persistent Chains](../../deployment/persistent_chains)
where each subfolder represents a `<chain-id>`.

If a new network needs to be initialized, please follow the [Running Genesis Node](running-genesis-node.md)
instructions first. After this more validator nodes can be added by following the instructions from this doc. 
 

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

The latest release can be found at [DCL Releases](https://github.com/zigbee-alliance/distributed-compliance-ledger/releases).

The following components will be needed:

* dcld (part of the release): The binary used for running a node and CLI.
* The service configuration file `dcld.service` 
(either part of the release or [deployment](../../deployment) folder).    
* Genesis transactions file: `genesis.json`
* The list of alive peers: `persistent_peers.txt`. It has the following format: `<node id>@<node ip>,<node2 id>@<node2 ip>,...`.

If you want to join an existing persistent network (such as Test Net), then look at the [Persistent Chains](../../deployment/persistent_chains)
folder for a list of available networks. Each subfolder there represents a `<chain-id>` 
and contains the genesis and persistent_peers files. 

### Deployment steps

1. Put `dcld` binary to `/usr/bin/` and configure permissions.

2. Configure CLI:
    * `dcld config chain-id <chain-id>`
      * Use `testnet` if you want to connect to the persistent Test Net
    * `dcld config output json` - Output format (text/json).
    * `dcld config node tcp://<host>:<port>` - Address of a node to connect. 
    Choose one of the listed in `persistent_peers.txt` file. 
    Example: `tcp://18.157.114.34:26657`.

3. Prepare keys and account:
    * dcld keys add <name> - Derive a new private key and encrypt to disk.
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
    `dcld query auth account --address=<address>`.
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
    * Init Node: `dcld init <node name> --chain-id <chain-id>`.
        * Use `testnet` if you want to connect to the persistent Test Net
    * Put `genesis.json` into dcld's config directory (usually `$HOME/.dcl/config/`).
        * Use `deployment/persistent_chains/testnet/genesis.json` if you want to connect to the persistent Test Net
    * Open `$HOME/.dcl/config/config.toml` file in your favorite text editor:
        * Tell node how to connect to the network:
            * Set the value for the `persistent_peers` field as the content of `persistent_peers.txt` file.
            * Use `deployment/persistent_chains/testnet/persistent_peers.txt` if you want to connect to the persistent Test Net.
        * Make your node public:
            * Open `$HOME/.dcl/config/config.toml`
            * Find the line under `# TCP or UNIX socket address for the RPC server to listen on`
            * Change it to: `laddr = "tcp://0.0.0.0:26657"`
        * Optionally change other setting.
    * Open `26656` (p2p) and `26657` (RPC) ports. 
        * `sudo ufw allow 26656/tcp`
        * `sudo ufw allow 26657/tcp`
    * Edit `dcld.service`
        * Replace `ubuntu` with a username you want to start service on behalf
    * Copy service configuration.
        * `cp dcld.service /etc/systemd/system/`

5. Add validator node to the network:
   * Get this node's tendermint validator address: `dcld tendermint show-address`.
       Expected output format: 
           ```
           cosmosvalcons1yrs697lxpwugy7h465wskwu2a5w9dgklx608f0
           ```
   * Get this node's tendermint validator pubkey: `dcld tendermint show-validator`.
       Expected output format: 
           ```
           cosmosvalconspub1zcjduepqcwg4eenpcxgs0269xuup5jlzj3pdquxlvj494cjxtqtcathsq7esfrsapa
           ```
   * Note that *validator address* and *validator pubkey* are not the same as `address` and `pubkey` were used for account creation.

   * Enable the service: `sudo systemctl enable dcld`
   * Start node: `sudo systemctl start dcld`
    
   * For testing purpose the node can be started in CLI mode: `dcld start` (instead of two previous `systemctl` commands).
   Service mode is recommended for demo and production environment.
   
   * Use `systemctl status dcld` to get the node service status. 
    In the output, you can notice that `height` increases quickly over time. 
    This means that the node in updating to the latest network state (it takes some time).
        
        You can also check node status by connecting CLI to your local node `dcld config node tcp://0.0.0.0:26657`
        and executing the command `dcld status` to get the current status.
        The `true` value for `catching_up` field means that the node is in the updating process.
        The value of `latest_block_height` reflects the current node height.
       
   * Wait until the value of `catching_up` field gets to `false` value.
      
   * Add validator node: `dcld tx validator add-node --pubkey=<validator pubkey> --moniker=<node name> --from=<key name>`.
   If the transaction has been successfully written you would find `"code": 0` in the output JSON. 

6. Check the node is running and participates in consensus:
    * Get the list of all nodes: `dcld query validator all-nodes`. 
    The node must present in the list and has the following params: `power:10` and `jailed:false`.

    * Get the node status: `dcld status --node tcp://<node_ip>:26657`. 
    The value of `node ip` matches to `[rpc] laddr` field in `$HOME/.dcl/config/config.toml`
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
    
    * Get the list of nodes participating in the consensus for the last block: `dcld query tendermint-validator-set`.
        * You can pass the additional value to get the result for a specific height: `dcld query tendermint-validator-set 100`  .
      
7. Congrats! You are an owner of the validator node.
