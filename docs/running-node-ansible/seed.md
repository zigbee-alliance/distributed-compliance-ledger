# Running Seed Node using Ansible
<!-- markdownlint-disable MD033 -->

## Prerequisites

Make sure you have all [prerequisites](./prerequisites.md) set up

## Configure DCL network parameters (local machine)

### 1. Set network chain ID

[`deployment/ansible/inventory/hosts.yml`]

```yaml
all:
  vars:
    chain_id: <chain-id>
  ...
```

Every network must have a unique chain ID (e.g. `test-net`, `main-net` etc.)

<details>
<summary>Example for Testnet 2.0 (clickable) </summary>

```yaml
all:
  vars:
    chain_id: testnet-2.0
  ...
```

</details>

<details>
<summary>Example for Mainnet (clickable) </summary>

```yaml
all:
  vars:
    chain_id: main-net
  ...
```
</details>

### 2. Put `genesis.json` file under specific directory

- Get or download `genesis.json` file of a network your node will be joining and put it under the following path:

  ```text
  deployment/persistent_chains/<chain-id>/genesis.json
  ```

  where `<chain-id>` is the chain ID of a network spefied in the previous step.

  <details>
  <summary>Example for Testnet 2.0 (clickable) </summary>

  For `testnet-2.0` the genesis file is already in place. So you don't need to do anything!

  ```text
  deployment/persistent_chains/testnet-2.0/genesis.json
  ```

  </details>

  <details>
  <summary>Example for Mainnet (clickable) </summary>

  For `main-net` the genesis file is already in place. So you don't need to do anything!

  ```text
  deployment/persistent_chains/main-net/genesis.json
  ```

  </details>

## Configure node type specific parameters (local machine)

### 1. Specify target instance address in the inventory file

[`deployment/ansible/inventory/hosts.yml`]

```yaml
all:
    ...
    children:
      ...
      seeds:
        hosts:
        <seed node IP address or hostname>
      ...
```

### 2. Set persistent peers string in seed configuration

[`deployment/ansible/roles/configure/vars/seed.yml`]

```yaml
config:
  p2p:
    persistent_peers: "<node1-ID>@<node1-IP>:26656,..."
...
```

`persistent_peers` values:
  - `Public Sentry` nodes with public IP
  - Use the following command to get `node-ID` of a node: `./dcld tendermint show-validator`.

### 3. (Optional) If you are joining a long-running network, enable `statesync` or use one of the options in [running-node-in-existing-network.md](../advanced/running-node-in-existing-network.md)

[`deployment/ansible/roles/configure/vars/seed.yml`]

```yaml
config:
...
  statesync:
    enable: true
    rpc_servers: "http(s):<node1-IP>:26657, ..."
    trust_height: <trust-height>
    trust_hash: "<trust-hash>"
...
```

<details>
<summary>Example for Testnet 2.0 (clickable) </summary>

```yaml
config:
  statesync:
    enable: true
    rpc_servers: "https://on.test-net.dcl.csa-iot.org:26657,https://on.test-net.dcl.csa-iot.org:26657"
```

</details>

<details>
<summary>Example for Mainnet (clickable) </summary>

```yaml
config:
  statesync:
    enable: true
    rpc_servers: "https://on.dcl.csa-iot.org:26657,https://on.dcl.csa-iot.org:26657"
```

</details>

> **_NOTE:_**  You should provide at least 2 addresses for `rpc_servers`. It can be 2 identical addresses

You can use the following command to obtain `<trust-height>` and `<trust-hash>` of your network

```bash
curl -s http(s)://<host>:<port>/commit | jq "{height: .result.signed_header.header.height, hash: .result.signed_header.commit.block_id.hash}"
```

<details>
<summary>Example for Testnet 2.0 (clickable) </summary>

```bash
curl -s https://on.test-net.dcl.csa-iot.org:26657/commit | jq "{height: .result.signed_header.header.height, hash: .result.signed_header.commit.block_id.hash}"
```

</details>

<details>
<summary>Example for Mainnet (clickable) </summary>

```bash
curl -s https://on.dcl.csa-iot.org:26657/commit | jq "{height: .result.signed_header.header.height, hash: .result.signed_header.commit.block_id.hash}"
```

</details>

- `<host>` - RPC endpoint host of the network being joined
- `<port>` - RPC endpoint port of the network being joined

> **_NOTE:_** State sync is not attempted if the node has any local state (LastBlockHeight > 0)

## Run ansible (local machine)

### 1. Verify that all the configuration parameters from the previous section are correct

### 2. Run ansible

```bash
ansible-playbook -i ./deployment/ansible/inventory  -u <target-host-ssh-user> ./deployment/ansible/deploy.yml
```

- `<target-host-ssh-username>` - target host ssh user
- Ansible provisioning can take several minutes depending on number of nodes being provisioned

## Deployment Verification (target machine)

### 1. Switch to cosmovisor user

```bash
sudo su -s /bin/bash cosmovisor
```

### 2. Query status

```bash
dcld status
```
