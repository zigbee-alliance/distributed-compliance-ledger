# Running a DCLedger Node using Ansible

This document describes in how to:

* configure different types of DCLedger nodes: genesis, validator and observer
* create a node administrator (or admin) account as a necessary part of validator and genesis node configuration

## Requirements (local machine)

### 1. Clone [DCL][5] project
```
git clone https://github.com/zigbee-alliance/distributed-compliance-ledger.git
```

### 2. Install `Python` and `Pip`
```bash
sudo apt-get update
sudo apt-get install -y --no-install-recommends python3
sudo apt install python3-pip
sudo apt install python3-testresources
```

### 3. Install [Ansible][4] and its dependencies
    
Run the following commands from the [DCL][5] project home
```bash
sudo pip install -r deployment/requirements.txt
ansible-galaxy install -r deployment/galaxy-requirements.yml
```

## Requirements (target machine)

### Hardware requirements

Minimal:

* 1GB RAM
* 25GB of disk space
* 1.4 GHz CPU

Recommended (for highload applications):

* 2GB RAM
* 100GB SSD
* x64 2.0 GHz 2v CPU

### Operating System

Current delivery is compiled and tested under `Ubuntu 20.04 LTS` so we recommend using this distribution for now.
In future, it will be possible to compile the application for a wide range of operating systems (thanks to Go language).

### Install `Python`
Python3 needs to be installed on target machine to run ansible playbooks
```bash
sudo apt-get update
sudo apt-get install -y --no-install-recommends python3
```

## Configure Nodes

### 1. Set network chain id in [`deployment/ansible/inventory/hosts.yml`]
```
all:
  vars:
    chain_id: <chain-id>
  ...
```
Every network must have a unique chain ID (e.g. `test-net`, `main-net` etc.)

### 2. Put `genesis.json` file under specific directory (if you are not running a genesis node)
Get or download `genesis.json` file of a network your node will be joining and put it under the following path:
```
deployment/persistent_chains/<chain-id>/genesis.json
```
where `<chain-id>` is the chain id of a network spefied in the previous step

### 3. Set node type specific config parameters

#### Genesis Node:
[`deployment/ansible/inventory/hosts.yml`]
  - Set Genesis node instance IP address or hostname
    ```
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

[`deployment/ansible/roles/configure/vars/validator.yml`]
  - Set `persisten_peers` value in 
      ```
      config:
        p2p:
          persistent_peers: "<node1-ID>@<node1-IP>:26656,..."
        ...
      ```
[`deployment/ansible/inventory/group_vars/genesis.yaml`]
  - Set genesis accounts:
      ```
      accounts:
      - name: validator-key
        passphrase: password123
        roles:
          - NodeAdmin
          - Trustee
      ```
    Genesis nodes should be created with at least one validator and trustee accounts

#### Validator Node:
[`deployment/ansible/inventory/hosts.yml`]
  - Set Validator node instance IP address or hostname
    ```
    all:
      ...
      children:
        ...
        validators:
          hosts:
            <validator node IP address or hostname>
        ...
    ```

[`deployment/ansible/roles/configure/vars/validator.yml`]
  - Set `persisten_peers` value in 
      ```
      config:
        p2p:
          persistent_peers: "<node1-ID>@<node1-IP>:26656,..."
        ...
      ```
  - (Optional) If you are joining a long-running network, enable `statesync`
      ```
      config:
        ...
        statesync:
          enable: true
          rpc_servers: "http(s):<node1-IP>:26657, ..."
          trust_height: 0
          trust_hash: ""
        ...
      ```
      refer to this [document](./running-node-in-existing-network.md) for detailed info

#### Observer Nodes:
[`deployment/ansible/inventory/hosts.yml`]
  - Set Observer node instances IP address or hostnames
    ```
    all:
      ...
      children:
        ...
        observers:
          hosts:
            <observer IP address or hostname>
            ...
        ...
    ```
  You can specify multiple observer node addresses to configure them all at the same time. 

[`deployment/ansible/roles/configure/vars/observer.yml`]
  - Set `persisten_peers` value in 
      ```
      config:
        p2p:
          persistent_peers: "<node1-ID>@<node1-IP>:26656,..."
        ...
      ```
  - (Optional) If you are joining a long-running network, enable `statesync`
      ```
      config:
        ...
        statesync:
          enable: true
          rpc_servers: "http(s):<node1-IP>:26657, ..."
          trust_height: 0
          trust_hash: ""
        ...
      ```
      refer to this [document](./running-node-in-existing-network.md) for detailed info
  

## Run nodes
### 1. Verify that all the configuration parameters from the previous section are correctly
### 2. Run ansible
```bash
ansible-playbook -i ./deployment/ansible/inventory  -u <target-host-ssh-user> ./deployment/ansible/deploy.yml
```
- `<target-host-ssh-username>` - target host ssh user
- Ansible provisioning can take several minutes depending on number of nodes being provisioned

[1]: https://www.terraform.io/
[2]: https://learn.hashicorp.com/tutorials/terraform/install-cli
[3]: https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html
[4]: https://www.ansible.com
[5]: https://github.com/zigbee-alliance/distributed-compliance-ledger.git
[6]: https://github.com/TomWright/dasel