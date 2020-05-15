## Running Node

This document describes in details how to configure and run validator node to a from scratch.

1. Assume we have the following artifacts:
    * Binary artifacts:
        * zbld: The binary used for running a node.
        * zblcli: The binary that allow users to interact with pool ledger.
    * Genesis transactions file: `genesis.json`
    * The list of alive peers: `persistent_peers.txt`
    * The service configuration file `zblcli.service`

2. Put `zbld` and `zblcli` binaries to `/usr/bin/` and configure permissions.

3. Configure zblcli:
    * `zblcli config chain-id zblchain`
    * `zblcli config output json` - Output format (text/json)
    * `zblcli config indent true` - Add indent to JSON response
    * `zblcli config trust-node false` - Verify proofs for node responses.
    * `zblcli config node <ip address>` - Address of a node to connect.

4. Prepare keys and account:
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
    * In order to ensure that account is created and has assigned role you can use the command: `zblcli query authnext account --address=<address>`
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

5. Initialize the node and create the necessary config files:
    * Init Node: `zbld init <node name> --chain-id zblchain`.
    * Put `genesis.json` into zbld's config directory (usually `$HOME/.zbld/config/`).
    * Edit `persistent_peers` field of `$HOME/.zbld/config/config.toml` file.
    Set it value into the body of `peers.txt`
    * Open `26656` (p2p) and `26657` (RPC) ports. 
        * `sudo ufw allow 26656/tcp`
        * `sudo ufw allow 26657/tcp`
    * Copy service configuration
        * `cp zbld.service /etc/systemd/system/`
    * Optionally, edit `$HOME/.zbld/config/config.toml` in order to set different setting (like listen address).

6. Add validator node to the network:
   * Get this node's tendermint validator address: `zbld tendermint show-address`
       Expected output format: 
           ```
           cosmosvalcons1yrs697lxpwugy7h465wskwu2a5w9dgklx608f0
           ```
   * Get this node's tendermint validator pubkey: `zbld tendermint show-validator`
       Expected output format: 
           ```
           cosmosvalconspub1zcjduepqcwg4eenpcxgs0269xuup5jlzj3pdquxlvj494cjxtqtcathsq7esfrsapa
           ```
   * Note that *validator address* and *validator pubkey* are not the same as `address` and `pubkey` were used for account creation.
   
   * Add validator node: `zblcli tx validator add-node --validator-address=<validator address> --validator-pubkey=<validator pubkey> --name=<node name> --from=<key name>`.
   If the transaction has been successfully written you would find `"success": true` in the output JSON. 
   
   * Start node: `sudo service zbld start`

7. Check node is running and participate in consensus:
    * Get the list of all nodes: `zblcli query validator all-nodes`. 
    The node must present in the list and has the following params: `power:10` and `jailed:false`.

    * Get the node status: `zblcli status --node <node ip>`
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
    
    * Get the list of nodes participating in the consensus for the last block: `zblcli tendermint-validator-set`
        * You can pass the additional value to get the result for a specific height: `zblcli tendermint-validator-set 100`  
      
8. Congrats! You are an owner of the validator node.
