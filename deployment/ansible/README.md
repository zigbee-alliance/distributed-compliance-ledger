# Ansible playbooks

This playbook can be used to deploy your own network of validator nodes for debug purposes.

If you like to join one of the persistent networks instead, please follow the [Running Node](../../docs/running-node.md) instructions. 

## Setting up nodes:

1. Generate genesis and configuration files for nodes by running `genlocalnetconfig.sh`
2. Make sure you have hosts with names `node0`, `node1`, `node2` and so on in host group `nodes` in your hosts file
3. Execute playbooks in the following order: `binary`, `config`, `service`, `start`
4. Don't forget to open `26656` (p2p) and `26657` (RPC) ports
5. Make sure nodes are running and writing blocks by running `status` playbook

PS. Don't forget to define node ips in their configurations so they will be able to find each other

## Setting up client:

Ansible is used to quickly deploy `dclcli` running in rest-server mode.

1. Add target machine to the `clients` group in your inventory file:
    - Hosts file is usually located in: `/etc/ansible/hosts`
    - It should be looking like this: 
    ```
    [clients]
    54.248.129.25 ansible_connection=ssh ansible_user=ubuntu
    ```
2. Execute playbooks in the following order: `binary`, `config`, `service`, `start`
    - Commands should be looking like: `ansible-playbook binary.yml`
4. Make sure clients are running by running `status` playbook
