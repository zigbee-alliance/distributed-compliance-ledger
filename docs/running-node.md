# Running a DCLedger Node

This document describes in how to:

* configure different types of DCLedger nodes: genesis, validator and observer
* create a node administrator (or admin) account as a necessary part of validator and genesis node configuration

## Components

* Common release artifacts:
  * Binary artifacts (part of the release):
    * `dcld`: The binary used for both running a node and interacting with it.
    * `cosmovisor`: A small process manager that supports an automated process of Cosmos SDK based application upgrade (`dcld` upgrade in our case).
  * The service configuration file `cosmovisor.service`
        (either part of the release or [deployment](../deployment) folder).
* Additional generated data (for validators and observers):
  * Genesis transactions file: `genesis.json`
  * The list of alive peers: `persistent_peers.txt` with the following format: `<node id>@<ip:port>,<node2 id>@<ip:port>,...`.

Please check [Get the artifacts](#get-the-artifacts) for the details how to get them.

## Hardware requirements

Minimal:

* 1GB RAM
* 25GB of disk space
* 1.4 GHz CPU

Recommended (for highload applications):

* 2GB RAM
* 100GB SSD
* x64 2.0 GHz 2v CPU

## Operating System

Current delivery is compiled and tested under `Ubuntu 20.04 LTS` so we recommend using this distribution for now.
In future, it will be possible to compile the application for a wide range of operating systems (thanks to Go language).

> Notes:
>
> * A part of the deployment commands below will try to enable and run `cosmovisor` as a systemd service, it means:
>   * that will require `sudo` for a user
>   * you may consider to use non-Ubuntu systemd systems but it's not officially supported for the moment
>   * in case non systemd system you would need to take care about `cosmovisor` service enablement and run as well

## Deployment Preparation

### (Optional) System cleanup

Required if a host has been already used in another DCLedger setup.

<!-- markdownlint-disable MD033 -->
<details>
<summary>Cleanup (click to expand)</summary>
<p>

```bash
sudo systemctl stop cosmovisor
sudo rm -f "$(which cosmovisor)"
sudo systemctl stop dcld 
sudo rm -f "$(which dcld)"
rm -rf "$HOME/.dcl" 
```

*NOTE: Some of the commands above may fail depending on whether or not `cosmovisor` was used in the previous setup.*

</p>
</details>
<!-- markdownlint-enable MD033 -->

### Get the artifacts

* download `dcld`, `cosmovisor` and `cosmovisor.service` from GitHub [release page](https://github.com/zigbee-alliance/distributed-compliance-ledger/releases)
* Get setup scripts either from [release page](https://github.com/zigbee-alliance/distributed-compliance-ledger/releases) or
    from [repository](../deployment/scripts) if you need latest development version.
* (for validator and observer) Get the running DCLedegr network data:
  * `genesis.json` can be found in a `<chain-id>` sub-directory of the [persistent_chains](../deployment/persistent_chains) folder
  * `persistent_peers.txt`: that file may be published there as well or can be requested from the DCLedger network admins otherwise

<!-- markdownlint-disable MD033 -->
<details>
<summary>Example (click to expand)</summary>
<p>

```bash
# release artifacts
curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/<release>/dcld
curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/<release>/cosmovisor
curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/<release>/cosmovisor.service

# deployment scripts
    # from release (if available)
curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/<release>/run_dcl_node
    # OR latest dev version
curl -L -O https://raw.githubusercontent.com/zigbee-alliance/distributed-compliance-ledger/master/deployment/scripts/run_dcl_node

# genesis file
curl -L -O https://raw.githubusercontent.com/zigbee-alliance/distributed-compliance-ledger/master/deployment/persistent_chains/<chain-id>/genesis.json

# persistent peers file (if available)
curl -L -O https://raw.githubusercontent.com/zigbee-alliance/distributed-compliance-ledger/master/deployment/persistent_chains/<chain-id>/persistent_peers.txt
```

</p>
</details>
<!-- markdownlint-enable MD033 -->

### Setup DCL binaries

* put `cosmovisor` binary in a folder listed in `$PATH` (e.g. `/usr/bin/`)
* set owner of `cosmovisor` binary to the user who will be used for `cosmovisor` service to run as
* set executable permission on `cosmovisor` binary for owner 

<!-- markdownlint-disable MD033 -->
<details>
<summary>Example for ubuntu user (click to expand)</summary>
<p>

```bash
sudo cp -f ./cosmovisor -t /usr/bin
sudo chown ubuntu /usr/bin/cosmovisor
sudo chmod u+x /usr/bin/cosmovisor
```

</p>
</details>
<!-- markdownlint-enable MD033 -->

### Configure the firewall

* ports `26656` (p2p) and `26657` (RPC) should be available for TCP connections
* if you use IP filtering rules they should be in sync with the persistent peers list

<!-- markdownlint-disable MD033 -->
<details>
<summary>Example for Ubuntu (click to expand)</summary>
<p>

```bash
sudo ufw allow 26656/tcp
sudo ufw allow 26657/tcp
```

</p>
</details>
<!-- markdownlint-enable MD033 -->

## Genesis Validator Node

This part describes how to configure a genesis node - a starting point of any new network.

The following steps automates a set of instructions that you can find in [Running Genesis Node](advanced/running-genesis-node.md) document.

**Note** This part is not required for all validator owners: it is performed only once for the initial (genesis) node of a DCLedger network.
If you are not going to become a genesis node admin you may jump to [Validator Node](#validator-node).

### Choose the chain ID

Every network (e.g. `test-net`, `main-net` etc.) must have a unique chain ID.

### Create keys for a node admin and a trustee genesis accounts

```bash
./dcld keys add "<key-name>" 2>&1 | tee "<key-name>.dclkey.data"
```

> Notes: It's important to keep the generated data (especially a mnemonic that allows to recover a key) in a safe place

### Setup a genesis node

Init node:

```bash
./dcld init "<node-name>" --chain-id "<chain-id>"
```

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
* if it's not acceptable for your case please consult a less automated guide [Running Genesis Node](advanced/running-genesis-node.md)
* you may likely want to note the summary that this script prints, in particular: node's address, public key and ID.

## Validator Node

This part describes how to configure a validator node and add it to the existing network.

The following steps automates a set of instructions that you can find in [Running Validator Node](advanced/running-validator-node.md) document

### Create a NodeAdmin account

Run the following to create a key:

```bash
./dcld keys add "<key-name>" 2>&1 | tee -a "<key-name>.dclkey.data"
```

And provide the output address and a public key to the network trustees.

> Notes: it's important to keep the generated data (especially a mnemonic that allows to recover the key) in a safe place

### Setup a validator node

Init node:

```bash
./dcld init "<node-name>" --chain-id "<chain-id>"
```

Run node:

```bash
./run_dcl_node -c "<chain-id>" "<node-name>"
```

> Notes:
>
> * the script assumes that:
>   * current user is going to be used for `cosmovisor` service to run as
>   * current user is in sudoers list
> * if it's not acceptable for your case please consult a less automated guide [Running Validator Node](advanced/running-validator-node.md)

This command:

* properly locates `genesis.json`
* configures and starts the node

### Ask for a ledger account

Provide generated node admin key `address` and `pubkey` to any `Trustee`(s). So they may create
an account with `NodeAdmin` role. And **wait** until:

* Account is created
* The node completed a catch-up:
  * `./dcld status --node tcp://<ip:port>` returns `false` for `catching_up` field

### Make the node a validator

```bash
./dcld tx validator add-node --pubkey="<validator pubkey>" --moniker="<node name>" --from="<key name>"
```

If the transaction has been successfully written you would find `"code": 0` in the output JSON.

### Notify other validator admins

Provide the node's `ID`, `IP` and a peer port (by default `26656`) to other validator admins.

*Note* Node `ID` can be found either in the output of the `run_dcl_node` script or using `./dcld status` command.

### (Optional) Create a key for a new trustee

If necessary you may also create a key to be used for a new Trustee account.

The procedure is similar to [NodeAdmin account creation](#create-a-nodeadmin-account).

## Observer Node

This part describes how to configure an observer node and add it to the existing network.

The following command automates a set of instructions that you can find in [Running Observer Node](advanced/running-observer-node.md) document

Init node:

```bash
./dcld init "<node-name>" --chain-id "<chain-id>"
```

Run node:

```bash
./run_dcl_node -t observer -c "<chain-id>" "<node-name>"
```

> Notes:
>
> * the script assumes that:
>   * current user is going to be used for `cosmovisor` service to run as
>   * current user is in sudoers list
> * if it's not acceptable for your case please consult a less automated guide [Running Observer Node](advanced/running-observer-node.md)

### Observer Peers

The list of persistent peers for an observer is not required to match the one used by the validators.

As a general guidance you may consider to use only the peers you own and/or trust.

## Deployment Verification

* Execute `source $HOME/.profile` to take the updated PATH into effect (it now contains the path to the current version of `dcld` binary):
  * `source $HOME/.profile`
* Check the account:
  * `dcld query auth account --address="<address>"`
* Check the node is running properly:
  * `dcld status --node tcp://<ip:port>`
  * The value of `<ip:port>` matches to `[rpc] laddr` field in `$HOME/.dcl/config/config.toml`
  * Make sure that `result.sync_info.latest_block_height` is increasing over the time (once in about 5 sec).
* Get the list of nodes participating in the consensus for the last block:
  * `dcld query tendermint-validator-set`.
  * You can pass the additional value to get the result for a specific height: `dcld query tendermint-validator-set 100`.

## Validator Node Maintenance

* `persistent_peers` field in `$HOME/.dcl/config/config.toml` should include the latest version of the validators list

  * you can use [update_peers](../deployment/scripts/update_peers)

  ```bash
  # by default path to a file is './persistent_peers'
  ./update_peers [PATH-TO-PEERS-FILE]
  ```

  * *Notes*
    * `cosmovisor` service should be restarted on any configuration changes
    * in case of any IP filtering firewall rules they should be updated as well
