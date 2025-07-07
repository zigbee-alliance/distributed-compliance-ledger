# Running a DCLedger Node manually (prerequisites)

## Components

* Common release artifacts:
  * Binary artifacts (part of the release):
    * `dcld`: The binary used for both running a node and interacting with it.
    * `cosmovisor`: A small process manager that supports an automated process of Cosmos SDK based application upgrade (`dcld` upgrade in our case).
  * The scripts files (part of the release):
    * `cosmovisor_preupgrade.sh`: The script that cosmovisor uses before updating dcld.
    * `cosmovisor_start.sh`: The script that start cosmovisor.
  * The service configuration files `cosmovisor.service` and `cosmovisor.conf`
        (either part of the release or [deployment](../../deployment) folder).
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

_NOTE: Some of the commands above may fail depending on whether or not `cosmovisor` was used in the previous setup._

</p>
</details>
<!-- markdownlint-enable MD033 -->

### Get the artifacts

* download `dcld`, `cosmovisor`, `cosmovisor.service`, `cosmovisor.conf`, `cosmovisor_preupgrade.sh` and `cosmovisor_start.sh` from GitHub [release page](https://github.com/zigbee-alliance/distributed-compliance-ledger/releases)
* Get setup scripts either from [release page](https://github.com/zigbee-alliance/distributed-compliance-ledger/releases) or
    from [repository](../../deployment/scripts) if you need latest development version.
* (for validator and observer) Get the running DCLedegr network data:
  * `genesis.json` can be found in a `<chain-id>` sub-directory of the [persistent_chains](../../deployment/persistent_chains) folder
  * `persistent_peers.txt`: that file may be published there as well or can be requested from the DCLedger network admins otherwise

See [this](../advanced/running-node-in-existing-network.md) document for running node in existing network. Also, note that if the [3<sup>rd</sup> option](../advanced/running-node-in-existing-network.md#3-catchup-from-genesis) is used, then a version at the time of genesis needs to be utilized.

<!-- markdownlint-disable MD033 -->
<details>
<summary>Example (click to expand)</summary>
<p>

```bash
# release artifacts
curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/<release>/dcld
curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/<release>/cosmovisor
curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/<release>/cosmovisor.service
curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/<release>/cosmovisor.conf
curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/<release>/cosmovisor_start.sh
curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/<release>/cosmovisor_preupgrade.sh

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

> Note:
>
> * `run_dcl_node` script adds the cosmovisor-controlled directory containing
the current version of `dcld` binary to `$PATH` of current user. To do this the
script adds a line doing the corresponding `$PATH` assignment to
`$HOME/.profile` file. If for some reason it is not effective for your
environment, please modify the corresponding line in the script in the way you
need or just comment out the corresponding line and manually add
`$HOME/.dcl/cosmovisor/current/bin` to `$PATH` of current user after
`run_dcl_node` script is executed (see below).

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