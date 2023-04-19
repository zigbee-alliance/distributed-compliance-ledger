# Running Observer Node using Ansible
<!-- markdownlint-disable MD033 -->
 

## Prerequisites
Follow all steps from [common.md](./common.md).

## Configure node type specific parameters (local machine)

### 1. Set node type

[`deployment/ansible/inventory/group_vars/all.yml`]

```yaml
  type_name: "ON"
  ...
```

### 2. Specify target instance address

[`deployment/ansible/inventory/hosts.yml`]

```yaml
all:
    ...
    children:
      ...
      observers:
        hosts:
        <observer node IP address or hostname>
      ...
```

### 3. Specify the nodes for connection

#### Option 1: an Observer node is connected to another organization's Public Sentries via Seed nodes

This is the main option if you want to connect an ON to CSA public nodes.

[`deployment/ansible/roles/configure/vars/observer.yml`]

```yaml
...
config:
  p2p:
    seeds: "<seed-node-ID>@<seed-node-public-IP>:26656"
...
```

<details>
<summary>CSA `seeds` Example for Testnet 2.0 (clickable) </summary>

```bash
seeds: "8190bf7a220892165727896ddac6e71e735babe5@100.25.175.140:26656"
```

</details>

<details>
<summary>CSA `seeds` Example for MainNet (clickable) </summary>

  ```bash
seeds: "ba1f547b83040904568f181a39ebe6d7e29dd438@54.183.6.67:26656"
```

</details>

#### Option 2: an Observer node is connected to another organization's public nodes

This option can be used if you have a trusted relationship with some organization and that organization
provided you with access to its nodes.   

[`deployment/ansible/roles/configure/vars/observer.yml`]

```yaml
...
config:
  p2p:
    persistent_peers: "<node1-ID>@<node1-IP>:26656,..."
...
```

`persistent_peers` value:
  - another organization nodes with public IPs that this organization shared with you. 

#### Option 3: an Observer node is connected to my organization nodes

This is the option when you have a VN and want to create an ON connected to it.

[`deployment/ansible/roles/configure/vars/observer.yml`]

```yaml
...
config:
  p2p:
    persistent_peers: "<node1-ID>@<node1-IP>:26656,..."
...
```

`persistent_peers` value:
  - If your VN doesn't use any Private Sentry nodes, then it must point to the `Validator` node with private IP.
  - Otherwise, it must point to the Private Sentry nodes with private IPs.
  - Use the following command to get `node-ID` of a node: `./dcld tendermint show-node-id`.

 

### 4. If you are joining a long-running network, enable `statesync` or use one of the options in [running-node-in-existing-network.md](../advanced/running-node-in-existing-network.md)

[`deployment/ansible/roles/configure/vars/observer.yml`]

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
