## How To

This document contains tutorials demonstrating how to accomplish common tasks using CLI.

- Please configure the CLI before using (see [CLI Configuration](#cli-configuration) section).
- If write requests to the Ledger needs to be sent, please make sure that you have
an Account created on the Ledger with an appropriate role (see [Getting Account](#getting-account) section in [how-to.md](how-to.md)).
- Sending read requests to the Ledger doesn't require an Account (Ledger is public for reads).
- A full list of all CLI commands can be found there: [cli-help](cli-help.md).
- After the CLI is configured and Account with an appropriate role is created,
the following instructions can be used for every actor (see [Use Case Diagrams](use_cases)):
    - [Trustee](#trustee-instructions) 
        - create and approve new accounts
        - approve X509 root certificates
        - publish X509 certificates
    - [CA](#ca-instructions)
        - propose X509 root certificates
    - [Vendor](#vendor-instructions) 
        - publish device model info
        - publish X509 certificates
    - [Test House](#test-house-instructions) 
        - publish compliance test results
        - publish X509 certificates
    - [ZB Certification Center](#certification-center-instructions)
        - certify or revoke certification of device models
        - publish X509 certificates
    - [Node Admin](#node-admin-instructions-setting-up-a-new-validator-node) 
        - add a new Validator node
        - publish X509 certificates

### CLI Configuration

CLI configuration file can be created or updated by executing of the command: `zblcli config <key> [value]`.
Here is the list of supported settings:
* chain-id <chain id> - Chain ID of pool node
* output <type> - Output format (text/json)
* indent <bool> - Add indent to JSON response
* trust-node <bool> - Trust connected full node (don't verify proofs for responses). The `false` value is recommended.
* node <node-ip> - Address `<host>:<port>` of the node to connect. 
* trace <bool> - Print out full stack trace on errors.
* broadcast-mode <mode> - Write transaction broadcast mode to use (one of: `sync`, `async`, `block`. `block` is default).

In order to connect the CLI to the ZB Ledger Demo Pool, the following parameters should be used:

* `zblcli config chain-id zblchain`
* `zblcli config output json` - Output format (text/json).
* `zblcli config indent true` - Add indent to JSON response.
* `zblcli config trust-node false` - Verify proofs for node responses.
* `zblcli config node <ip address>` - Address of a node to connect. 
Choose one of the listed in `persistent_peers.txt` file. 
Example: `tcp://18.157.114.34:26657`.


### Getting Account
Ledger is public for read which means that anyone can read from the Ledger without a need to have 
an Account but it is private for write. 
In order to send write transactions to the ledger you need: 

  - Have a private/public key pair.
  - Have an Account created on the ledger via `ACCOUNT` transaction (see [Use Case Txn Auth](use_cases/use_cases_txn_auth.puml)).
      - The Account stores the public part of the key
      - The Account has an associated role. The role is used for authorization policies.
  - Sign every transaction by the private key.

Here is steps for getting an account:
* Generate keys and local account: `zblcli keys add <name>`.
* Share generated `address` and `pubkey` to any `Trustee`. 
* `Trustee` registers the account on the ledger: `zblcli tx authnext create-account --address=<account address> --pubkey=<account pubkey> --from <trustee>`
* Optionally, `Trustee` can assign some role to the account: `zblcli tx authz assign-role --address=<account address> --role=<role> --from <trustee>`
* Check account is created: `zblcli query authnext account --address=<account address>`

Example:
* `zblcli keys add jack`
* `zblcli tx authnext create-account --address=cosmos1sug8cquqnn5jddkqt4ud6hcr290sn4wh96x5tv --pubkey=cosmospub1addwnpepqvnfd2f99vew4t7phe3mqprmceq3jgavm0rguef3gkv8z8jd6lg25egq6d5 --from trustee`
* `zblcli tx authz assign-role --address=cosmos1sug8cquqnn5jddkqt4ud6hcr290sn4wh96x5tv --role="NodeAdmin" --from trustee`
* `zblcli query authnext account --address=cosmos1sug8cquqnn5jddkqt4ud6hcr290sn4wh96x5tv`

### Trustee Instructions

### CA instructions

### Vendor Instructions

### Test House Instructions

### Certification Center Instructions



### Node Admin instructions (Setting up a new Validator Node)

Validators are responsible for committing of new blocks to the ledger.
Here are steps for setting up a new validator node.

A more detailed instruction on how to add a validator node to the current ZB Ledger pool
can be found here: [running-node.md](running-node.md).

* Firstly you have to posses an account with `NodeAdmin` role on the ledger. See [Getting Account](#getting-account):
    
* Initialize the node and create the necessary config files:
    * Init Node: `zbld init <node name> --chain-id <chain id>`.
    * Fetch the network `genesis.json` file and put it into zbld's config directory (usually `$HOME/.zbld/config/`).
    * In order to join network your node needs to know how to find alive peers. 
    Update `persistent_peers` field of `$HOME/.zbld/config/config.toml` file to contain peers info in the format:
    `<node1 id>@<node1 listen_addr>,<node2 id>@<node2 listen_addr>,.....`
    * Open `26656` (p2p) and `26657` (RPC) ports.

* Add validator node to the network:
    * Get this node's tendermint validator *consensus address*: `zbld tendermint show-address`
    * Get this node's tendermint validator *consensus pubkey*: `zbld tendermint show-validator`
    * Note that *consensus address* and *consensus pubkey* are not the same as `address` and `pubkey` were used for account creation.
    * Start node: `zbld start`
    * In the output, you can notice that `height` increases quickly over time. 
    This means that the node is in updating to the latest network state (it takes some time).
    
        You can also check node status by connecting CLI to your local node `zblcli config node tcp://0.0.0.0:26657`
        and executing the command `zblcli status` to get the current status.
        The `true` value for `catching_up` field means that the node is in the updating process.
        The value of `latest_block_height` reflects the current node height.

    * Wait until the value of `catching_up` field gets to `false` value.
    * Add validator node: `zblcli tx validator add-node --validator-address=<validator address> --validator-pubkey=<validator pubkey> --name=<node name> --from=<name>`

* Congrats! You are an owner of the validator node.

* Check node is alive and participate in consensus:
    * Get the list of all nodes: `zblcli query validator all-nodes`. 
    The node must present in the list and has the following params: `power:10` and `jailed:false`.
    * Get the list of nodes participating in the consensus for the last block: `zblcli tendermint-validator-set`
        * You can pass the additional value to get the result for a specific height: `zblcli tendermint-validator-set 100`  
    * Get the node status: `zblcli status --node <node ip>`

Example:
* `zbld init node-name --chain-id zblchain`
* `cp /source/genesis.json $HOME/.zbld/config/`
* `sed -i "s/persistent_peers = \"\"/<node id>@<node ip>,<node2 id>@<node2 ip>/g" $HOME/.zbld/config/config.toml`
* `sudo ufw allow 26656/tcp`
* `sudo ufw allow 26657/tcp`
* `zbld start`
* `zbld status`
* `zblcli tx validator add-node --validator-address=$(zbld tendermint show-address) --validator-pubkey=$(zbld tendermint show-validator) --name=node-name --from=node_admin`
* `zblcli query validator all-nodes`

##### Policy
- Maximum number of nodes (`MaxNodes`): 100 
- Size (number of blocks) of the sliding window used to track validator liveness (`SignedBlocksWindow`): 100
- Minimal number of blocks must have been signed per window (`MinSignedPerWindow`): 50

Node will be jailed and removed from the active validator set in the following conditions are met:
- Node passed minimum height: `node start height + SignedBlocksWindow`
- Node exceeded the maximum number of unsigned blocks withing the window: `SignedBlocksWindow - MinSignedPerWindow`

Note that jailed node will not be removed from the main index to track validator nodes. 
So it is not possible to create a new node with the same `validator address`.
In order to unjail the node and return it to the active tendermint validator set the sufficient number of Trustee's approvals is needed 
(see authorization rules).
