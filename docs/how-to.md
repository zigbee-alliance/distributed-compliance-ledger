## How To

This document contains tutorials demonstrating how to accomplish common tasks.

### Setting up a Validator Node

Validators are responsible for committing of new blocks to the ledger.
Here are steps for setting up a new validator node.

* Firstly you have to posses an account with `NodeAdmin` role on the ledger. Here are steps to get an appropriate account:
    * Generate keys and local account: `zblcli keys add <name>`
    * Register account on the ledger by sharing generated `address` and `pubkey` to any `Trustee`. 
    `Trustee` registers the account on the ledger: `zblcli tx authnext create-account <address> <pubkey> --from <trustee>`  
    * `Trustee` assigns `NodeAdmin` role to the account: `zblcli tx authz assign-role <address> "NodeAdmin" --from <trustee>`
    
* Initialize the node and create the necessary config files:
    * Init Node: `zbld init <node name> --chain-id <chain id>`.
    * Fetch the network `genesis.json` file and put it into zbld's config directory (usually `$HOME/.zbld/config/`).
    * In order to join network your node needs to know how to find alive peers. 
    Update `persistent_peers` field of `$HOME/.zbld/config/config.toml` file to contain peers info in the format:
    `<node1 id>@<node1 listen_addr>,<node2 id>@<node2 listen_addr>,.....`

* Add validator node to the network:
    * Get your `pubkey` that can be used to create a new validator: `zbld tendermint show-validator`
    * Add validator node: `zblcli tx validator add-node --pubkey=<pubkey> --name=<node name> --from=<name>`
    * Start node: `zbld start`

* Congrats! You are an owner of the validator node.