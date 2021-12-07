# Running a DCLedger Node

This document describes in details how to configure different types of DCLedger nodes.

## Components

*   Common release artifacts:
    *   Binary artifacts (part of the release):
        *   dcld: The binary used for running a node.
        *   dclcli: The binary that allows users to interact with the network of nodes.
    *   The service configuration file `dcld.service`
        (either part of the release or [deployment](../deployment) folder).
    *   where to get:
        *   The latest release can be found at [DCL Releases](https://github.com/zigbee-alliance/distributed-compliance-ledger/releases).
*   Additional generated data (for validators and observers):
    *   Genesis transactions file: `genesis.json`
    *   The list of alive peers: `persistent_peers.txt` with the following format: `<node id>@<ip:port>,<node2 id>@<ip:port>,...`.
    *   where to get:
        *   If you want to join an existing persistent network (such as Test Net), then look at the [Persistent Chains](../deployment/persistent_chains)
            folder for a list of available networks. Each sub-directory there represents a `<chain-id>` and contains the genesis. Also it may contain `persistent_peers.txt` files.

## Hardware requirements

Minimal:

*   1GB RAM
*   25GB of disk space
*   1.4 GHz CPU

Recommended (for highload applications):

*   2GB RAM
*   100GB SSD
*   x64 2.0 GHz 2v CPU

## Operating System

Current delivery is compiled and tested under `Ubuntu 20.04 LTS` so we recommend using this distribution for now.
In future, it will be possible to compile the application for a wide range of operating systems (thanks to Go language).

**Notes**

*   A part of the deployment commands below will try to enable and run `dcld` as a systemd service, it means:
    *   that will require `sudo` for a user
    *   you may consider to use non-Ubuntu systemd systems but it's not officially supported for the moment
    *   in case non systemd system you would need to take care about `dlcd` service enablement and run as well

## Deployment

### Preparation

#### (Optional) System cleanup

Required if a host has been already used in another DCLedger setup.

<details>
<summary>Cleanup (click to expand)</summary>
<p>

```bash
$ sudo systemctl stop dcld 
$ rm -rf "$HOME/.dcld" "$HOME/.dclcli"
$ sudo rm -f "$(which dcld)" "$(which dclcli)"
```

</p>
</details>

#### Get the artifacts

*   download `dclcli`, `dcld` and `dcld.service` from GitHub [release page](https://github.com/zigbee-alliance/distributed-compliance-ledger/releases)
*   Get setup scripts either from [release page](https://github.com/zigbee-alliance/distributed-compliance-ledger/releases) or
    from [repository](../deployment/scripts) if you need latest development version.
*   (for validator and observer) Get the running DCLedegr network data:
    
    *   `persistent_peers.txt`: that file may be published there as well or can be requested from the DCLedger network administrators otherwise

<details>
<summary>Example (click to expand)</summary>
<p>

```bash
# release artifacts
curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/<release>/dclcli
curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/<release>/dcld
curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/<release>/dcld.service
curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/<release>/run_dcl_node

# OR latest dev version in case of some not yet released improvements
curl -L -O https://raw.githubusercontent.com/zigbee-alliance/distributed-compliance-ledger/master/deployment/scripts/run_dcl_node

curl -L -O https://raw.githubusercontent.com/zigbee-alliance/distributed-compliance-ledger/master/deployment/persistent_chains/<chain-id>/genesis.json
```

</p>
</details>

#### Setup DCL binaries

*   located `dlcd` and `dclcli` binaries in a folder listed in `$PATH` (e.g. `/usr/bin/`)
*   set a proper owner and executable permissions

<details>
<summary>Example for ubuntu user (click to expand)</summary>
<p>

```bash
$ sudo cp -f ./dclcli ./dcld -t /usr/bin
$ sudo chown ubuntu /usr/bin/dclcli /usr/bin/dcld
$ sudo chmod u+x /usr/bin/dclcli /usr/bin/dcld
```

</p>
</details>

#### Configure the firewall

Ports `26656` (p2p) and `26657` (RPC) should be available for TCP connections.

<details>
<summary>Example for Ubuntu (click to expand)</summary>
<p>

```bash
$ sudo ufw allow 26656/tcp
$ sudo ufw allow 26657/tcp
```

</p>
</details>

### Genesis Validator Node

This part describes how to configure a genesis node - a starting point of any new network.

The following steps automates a set of instructions that you can find in [Running Genesis Node](running-genesis-node.md) document.

**Note** This part is not requried for all validator owners: it si performed only once for the initial (genesis) node of a DCLedger network.
If you are not going to become a genesis node administrator you may jump to [Validator Node](#validator-node).

#### Choose the chain ID

Every network (for example, test-net, main-net, etc.) must have a unique chain ID.

#### Setup a node

Run

```bash
$ ./run_dcl_node -t genesis -c <chain-id> node0
```

This command:

*   generates a new key entry for a node admin account
*   generates `genesis.json` file with the following entries:
    *   a genesis account for the key entry above with `Trustee` and `NodeAdmin` roles
    *   a genesis txn that makes the local node a validator
*   configures and starts the node
*   if `-u` option is specified the command will set a custom user the `dcld` service is executed as

Outputs:

*   `*.dclkey.json` file in the current directory with node admint key data (address, public key, mnemonic)
*   standard output:
    *   genesis file `$HOME/.dcld/config/genesis.json`

**Notes**

    * by default the command will try to setup a systemd `dcld` service to start under `ubuntu` user
    * if needed you may consier to specify a different service user `-u user`

### Validator Node

This part describes how to configure a validator node and add it to the existing network.

The following steps automates a set of instructions that you can find in [Running Validator Node](running-validator-node.md) document

#### Setup a node

Run

```bash
$ ./run_dcl_node -c <chain-id> [-u <user>] <node-name>
```

This command:

*   generates a new key entry for a node admin account
    *   by default `<node_name>admin` key name is used
    *   can be configured using `-k/--key-name` option
*   properly locates `genesis.json`
*   configures and starts the node

Outputs:

*   `*.dclkey.json` file in the current directory with node admin key data (address, public key, mnemonic)
*   standard output:
    *   node admin key data: `address` and `pubkey`
    *   validator data: `address` and `pubkey`

#### Ask for a ledger account

Provide generated node admin key `address` and `pubkey` to any `Trustee`(s). So they may create
an account with `NodeAdmin` role. And **wait** until:

*   Account is created
*   The node completed a catch-up:
    *   `dclcli status --node <ip:port>` returns `false` for `catching_up` field

#### Make the node a validator

```bash
$ dclcli tx validator add-node \
    --validator-address=<validator address> --validator-pubkey=<validator pubkey> \
    --name=<node name> --from=<key name>
```

If the transaction has been successfully written you would find `"success": true` in the output JSON.

#### Notify other validator administrators

Provide the node's `id`, `ip` and peer port (by default `26656`) to other validator administrators

### Observer Node

This part describes how to configure an observer node and add it to the existing network.

The following command automates a set of instructions that you can find in [Running Observer Node](running-observer-node.md) document

Run

```bash
$ ./run_dcl_node -t observer -c <chain-id> [-u <user>] <node-name>
```

Notes:

    *   if `-u` option is specified the command will set a custom user the `dcld` service is executed as

## Deployment Verification

*   Check the account:
    *   `dclcli query auth account --address=<address>`
*   Check the node is running and participates in consensus:
    *   `dclcli status --node <ip:port>`
    *   The value of `<ip:port>` matches to `[rpc] laddr` field in `$HOME/.dcld/config/config.toml`
    *   Make sure that `result.sync_info.latest_block_height` is increasing over the time (once in about 5 sec).
*   Get the list of nodes participating in the consensus for the last block:
    *   `dclcli tendermint-validator-set`.
    *   You can pass the additional value to get the result for a specific height: `dclcli tendermint-validator-set 100`.

## Validator Node Maintenance

*   `persistent_peers` field in `$HOME/.dcld/config/config.toml` should include the latest version of the validators list
    *   **Note** `dcld` service should be restarted on any configuration changes
