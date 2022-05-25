# Running Observer Node using Ansible
## Prerequisites
Make sure you have all [prerequisites](./prerequisites.md) set up
## Set up ansible configuration (local machine)
### 1. Specify target instance address in the inventory file
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

### 2. Set persistent peers string in validator configuration
[`deployment/ansible/roles/configure/vars/observer.yml`]
```yaml
config:
  p2p:
    persistent_peers: "<node1-ID>@<node1-IP>:26656,..."
...
```

### 3. (Optional) If you are joining a long-running network, enable `statesync`
[`deployment/ansible/roles/configure/vars/observer.yml`]

    ```yaml
    config:
    ...
    statesync:
        enable: true
        rpc_servers: "http(s):<node1-IP>:26657, ..."
        trust_height: 0
        trust_hash: ""
    ...
    ```
> **_NOTE:_**  You should provide at least 2 addresses for `rpc_servers`. It can be 2 identical addresses

You can use the following command to obtain `<trust-height>` and `<trust-hash>` of your network

```bash
curl -s http(s)://<host>:<port>/commit | jq "{height: .result.signed_header.header.height, hash: .result.signed_header.commit.block_id.hash}"
```

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
```
sudo su -s /bin/bash cosmovisor
```

### 2. Query status
```
dcld status
```