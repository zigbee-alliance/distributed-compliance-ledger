# Running a DCLedger Node

This document describes in details how to configure different types of DCLedger nodes.

## Components

*   Common release artifacts:
    * Binary artifacts (part of the release):
        * dcld: The binary used for running a node.
        * dclcli: The binary that allows users to interact with the network of nodes.
    * The service configuration file `dcld.service` 
    (either part of the release or [deployment](https://github.com/zigbee-alliance/distributed-compliance-ledger/deployment) folder).    
    * where to get:
        * The latest release can be found at [DCL Releases](https://github.com/zigbee-alliance/distributed-compliance-ledger/releases).
* Additional generated data (for validators and observers):
    * Genesis transactions file: `genesis.json`
    * The list of alive peers: `persistent_peers.txt` with the following format: `<node id>@<node ip>,<node2 id>@<node2 ip>,...`.
    * where to get:
        * If you want to join an existing persistent network (such as Test Net), then look at the [Persistent Chains](../deployment/persistent_chains)
          folder for a list of available networks. Each subfolder there represents a `<chain-id>` and contains the genesis and persistent_peers files. 

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

Notes:
    * the deployment commands below will try to setup and run `dcld` systemd service on Ubuntu
    * that will require `sudo` for a user
    * in case non-Ubuntu system these steps will be scipped so you would need to take care about

## Deployment

Pre-requisites:
1. `dcld` and `dclcli` binaries are located in `/usr/bin`
2. (for systemd systems) `dcld.service` is available in the current directory
3. (for validator and observer) `genesis.json` and `persistent_peers.txt` are available in the current directory

### Genesis Validator Node

This part describes how to configure a genesis node - a starting point of any new network.

The following steps automates a set of instructions that you can find in [Running Genesis Node](running-genesis-node.md) document.

1. Choose the chain ID. Every network (for example, test-net, main-net, etc.)
must have a unique chain ID.

2. run

```bash
$ run_dcl_node -t genesis -c <chain-id> node0
```

This command:

    - generates a new key entry for a node admin account
    - generates `genesis.json` file with the following entries:
        - a genenesis account for the key entry above with `Trustee` and `NodeAdmin` roles
        - a genesis txn that makes the local node a validator
    - configures and starts the node

Outputs:

    - `*.dclkey.json` file in the current directory with node admint key data (address, public key, mnemonic)
    - standard output:
        - genesis file `$HOME/.dcld/config/genesis.json`

## Validator Node

This part describes how to configure a validator node and add it to the existing network.

The following steps automates a set of instructions that you can find in [Running Validator Node](running-validator-node.md) document

1. run

```bash
$ run_dcl_node -c <chain-id> <node-name>
```

This command:

    - generates a new key entry for a node admin account
        - by default `<node_name>admin` key name is used
        - can be configured using `-k/--key-name` option
    - properly locates `genesis.json`
    - configures and starts the node

Outputs:

    - `*.dclkey.json` file in the current directory with node admint key data (address, public key, mnemonic)
    - standard output:
        - node admin key data: `address` and `pubkey`
        - validator data: `address` and `pubkey`

2. Provide generated node admin key `address` and `pubkey` to any `Trustee`(s). So they may create
   an account with `NodeAdmin` role. And **wait** until:

   * Account is created
   * The node completed a catch-up:
        * `dclcli status --node <node ip>` returns `false` for `catching_up` field

3. Make the node a validator:

```bash
$ dclcli tx validator add-node \
    --validator-address=<validator address> --validator-pubkey=<validator pubkey> \
    --name=<node name> --from=<key name>
```

If the transaction has been successfully written you would find `"success": true` in the output JSON. 


### Observer Node

This part describes how to configure an observer node and add it to the existing network.

The following command automates a set of instructions that you can find in [Running Observer Node](running-observer-node.md) document

So to launch an observer node you only need to run:

```bash
$ run_dcl_node -t observer -c <chain-id> <node-name>
```

## Deployment Verification

* Check the account:
    * `dclcli query auth account --address=<address>`
* Check the node is running and participates in consensus: 
    * `dclcli status --node <node ip>`
    * The value of `node ip` matches to `[rpc] laddr` field in `$HOME/.dcld/config/config.toml`
    * Make sure that `result.sync_info.latest_block_height` is increasing over the time (once in about 5 sec).
* Get the list of nodes participating in the consensus for the last block:
    * `dclcli tendermint-validator-set`.
    * You can pass the additional value to get the result for a specific height: `dclcli tendermint-validator-set 100`.
