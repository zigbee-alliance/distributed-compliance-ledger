# Add Observer Node to Testnet 2.0

Here is the step-by-step guide for adding an Observer Node to Testnet 2.0 using `state sync` option
as described in  [Runing Node in Existing Network](../../../docs/advanced/running-node-in-existing-network.md) document:

## 1. Configure an instance

### 1.1 Ubuntu 20.04 LTS is recommended

### 1.2 (Optional) Clean up the system

> **_Note:_** Following steps are needed if you earlier version of DCL installed on the same computer

```bash
sudo systemctl stop dcld
sudo rm -f "$(which dcld)"
rm -rf "$HOME/.dcl"
```

### 1.3 Configure the firewall

* p2p and RPC (by default: `26656` and `26657` respectively) should be available for TCP connections.
  For Ubuntu:

```bash
# P2P
sudo ufw allow 26656/tcp

# RPC
sudo ufw allow 26657/tcp

# gRPC
sudo ufw allow 9090/tcp

# REST
sudo ufw allow 1317/tcp
```

* In case of IP filtering rules ensure they allow incoming and outcoming connections from/to other peers.

## 2. Download required artifacts

### 2.1 Get the release artifacts

> **_Note:_** Downloading latest release artifacts is recommended. Cosmovisor service is available only starting from DCL `v0.8.0`

```bash
curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/<version>/dcld
curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/<version>/cosmovisor
curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/<version>/cosmovisor.service

# helper script to automate running a node
curl -L -O https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/<release>/run_dcl_node
```

`<version>` - DCL release version (ex: `v0.9.0`)

### 2.2 Download genesis file for Testnet 2.0

```bash
curl -L -O https://raw.githubusercontent.com/zigbee-alliance/distributed-compliance-ledger/master/deployment/persistent_chains/testnet-2.0/genesis.json
```

### 2.3. Put `cosmovisor` binary in a folder listed in `$PATH` (e.g. `/usr/bin/`)

```bash
sudo cp ./cosmovisor /usr/bin/
sudo chmod a+x /usr/bin/cosmovisor
```

## 3. Run Observer Node

### 3.1 Initialize node

```bash
chmod u+x ./dcld
./dcld init "<node-name>" --chain-id "testnet-2.0"
```

### 3.2 Set the following `state sync` parameters under `[statesync]` section in `$HOME/.dcl/config/config.toml`

```ini
[statesync]
enable = true

rpc_servers = "https://on.test-net.dcl.csa-iot.org:26657,https://on.test-net.dcl.csa-iot.org:26657"

trust_height = <trust-height>
trust_hash = "<trust-hash>"
```

You can get `<trust-height>` and `<trust-hash>` parameters using the following command:

```bash
curl -s https://on.test-net.dcl.csa-iot.org:26657/commit | jq "{height: .result.signed_header.header.height, hash: .result.signed_header.commit.block_id.hash}"
```

### 3.3 Set `seeds` parameter under `[p2p]` section in `$HOME/.dcl/config/config.toml`

```ini
[p2p]
seeds = "8190bf7a220892165727896ddac6e71e735babe5@100.25.175.140:26656"
```

### 3.4 Create an empty `persistent_peers.txt` in the current path because this file is required by `run_dcl_node` script

Your observer node will use `Seed Node` to discover peers (see 3.3)

```bash
touch persistent_peers.txt
```

### 3.5 Run Observer Node

```bash
chmod u+x ./run_dcl_node
./run_dcl_node -t observer -c "testnet-2.0" "<node-name>"
```

### 3.6 Verify node is running

Execute `source $HOME/.profile` to take the updated `$PATH` into effect, now
it includes the directory containing the current version of `dcld` binary (if
you have not modified or commented out the line doing the corresponding
`$PATH` assignment in `run_dcl_node` script):

```bash
source $HOME/.profile
```

Check whether `catching_up` field is set to `false` after a while using the following query:

```bash
dcld status
```

It may take couple of minutes to catch up using `state-sync` depending on how far the `statesync snapshot` was from the current state of the network
