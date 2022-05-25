# Running Private Sentry Node using Ansible
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
    private_sentries:
      hosts:
        <private sentry node IP address or hostname>
    ...
```

### 2. Set persistent peers string in validator configuration
[`deployment/ansible/roles/configure/vars/private-sentry.yml`]

```yaml
config:
  p2p:
    persistent_peers: "<node1-ID>@<node1-IP>:26656,..."
...
```
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