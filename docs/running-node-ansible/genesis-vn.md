# Running Genesis Validator Node using Ansible

## Prerequisites
Make sure you have all [prerequisites](./prerequisites.md) set up

## Configure DCL network parameters (local machine)
### 1. Set network chain id
[`deployment/ansible/inventory/hosts.yml`]
```yaml
all:
  vars:
    chain_id: <chain-id>
  ...
```
Every network must have a unique chain ID (e.g. `test-net`, `main-net` etc.)

## Configure node type specific parameters (local machine)
### 1. Specify target instance address in the inventory file
[`deployment/ansible/inventory/hosts.yml`]

```yaml
all:
  ...
  children:
    genesis:
      hosts:
        <genesis node IP address or hostname>
    ...
    validators:
      hosts:
        <genesis node IP address or hostname>
    ...
```
    You should set the same address for `validators` and `genesis` hosts because a genesis node is also a validator node

### 2. Set persistent peers string in validator configuration
[`deployment/ansible/roles/configure/vars/validator.yml`]

```yaml
config:
  p2p:
    persistent_peers: "<node1-ID>@<node1-IP>:26656,..."
...
```

### 3. Specify genesis accounts
[`deployment/ansible/inventory/group_vars/genesis.yaml`]
```yaml
accounts:
  - name: validator-key
    passphrase: password123
    roles:
        - NodeAdmin
        - Trustee
```
    Genesis nodes should be created with at least one validator and trustee accounts

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