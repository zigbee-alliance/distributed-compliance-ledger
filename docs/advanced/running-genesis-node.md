# Running a Genesis Validator Node

This document describes in details how to configure a genesis (first) validator node.

Here we assume (for simplicity) that the genesis block consists of a single node only.
Please note that nothing prevents you from adding more nodes to the genesis file by adapting the instructions accordingly.

## Hardware requirements

Minimal:

- 1GB RAM
- 25GB of disk space
- 1.4 GHz CPU

Recommended (for highload applications):

- 2GB RAM
- 100GB SSD
- x64 2.0 GHz 2v CPU

## Operating System

Current delivery is compiled and tested under `Ubuntu 18.04.3 LTS` so we recommend using this distribution for now. In future, it will be possible to compile the application for a wide range of operating systems thanks to Go language.

## Components

The latest release can be found at [DCL Releases](https://github.com/zigbee-alliance/distributed-compliance-ledger/releases).

The following components will be needed:

- `dcld` (part of the release): The application binary which can be run in node or CLI mode.
- `cosmovisor` (part of the release): A small process manager that supports an automated process of Cosmos SDK based application upgrade (`dcld` upgrade in our case).
- The service configuration file `cosmovisor.service`
(either part of the release or [deployment](../../deployment) folder).

### Deployment steps

1. Put `cosmovisor` binary to `/usr/bin/`, set proper owner and execution permissions.

2. Locate the genesis app version to genesis application version directory:
  * Create `$HOME/.dcl/cosmovisor/genesis/bin` directory.
  * Copy `dcld` binary to it, set proper owner and execution permissions.
      Please note that execution permissions on `dcld` should be granted to all (i.e. User, Group and Others classes)
      because cosmovisor requires execution permission on the application binary to be granted to Others class.

3. Choose the chain ID. Every network (for example, test-net, main-net, etc.)
must have a unique chain ID.

4. Configure CLI:
    - `./dcld config chain-id <chain-id>` - the chosen unique chain ID.
    - `./dcld config output json` - Output format (text/json).

5. Prepare keys:
    - Derive a new private key and encrypt to disk: `./dcld keys add <name>`.
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

6. Prepare genesis node configuration:

    - Initialize new configuration: `./dcld init <node-name> --chain-id <chain-id>`.
    - Add genesis account with the generated key and `Trustee`, `NodeAdmin` roles:
    `./dcld add-genesis-account --address=<address> --pubkey=<pubkey> --roles="Trustee,NodeAdmin"`
    - Optionally, add other genesis accounts using the same command.
    - Create genesis transaction: `./dcld gentx --from <name>`, where `<name>` is the keys' name specified at Step 5.
    - Collect genesis transactions: `./dcld collect-gentxs`.
    - Validate genesis file: `./dcld validate-genesis`.
    - Genesis file is located in `$HOME/.dcl/config/genesis.json`. Give this file to each new node admin.

7. Run node:
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
    - Execute the following command to apply the updated PATH immediately:
        - `source $HOME/.profile`

8. Check that genesis account is created:

    - In order to ensure that account is created and has assigned role you can use the command:
    `dcld query auth account --address=<address>`.
    Expected output format:

        ```json
        {
          "result": {
            "address": string, // bench32 encoded address
            "public_key": "string, // bench32 encoded public key
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

9. Check the node is running and participates in consensus:
    - Get the list of all nodes: `dcld query validator all-nodes`.
    The node must present in the list and has the following params: `power:10` and `jailed:false`.

    - Get the node status: `dcld status --node tcp://<node_ip>:26657`.
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

    - Get the list of nodes participating in the consensus for the last block: `dcld query tendermint-validator-set`.
        - You can pass the additional value to get the result for a specific height: `dcld query tendermint-validator-set 100`.
        - As for the genesis state we have just 1 node, the command should return only our node at this phase.

10. Add more initial trusted validator nodes to the network

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

11. Adding more validator nodes to the network

    - Just follow the [Running Node](../running-node.md) instructions
    using the chosen chain-id and the genesis and persistent_peers files above (from the chain-id subfolder).
    - Please make sure, that the persistent_peers contain all the nodes (including the genesis node)
    added at Step 10.

12. Congrats! You are an owner of the genesis validator node.
