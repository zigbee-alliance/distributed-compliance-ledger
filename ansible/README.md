# Ansible playbooks

## Setting up nodes:

1. Generate genesis and configuration files for nodes by running `genlocalnetconfig.sh`
2. Make sure you have hosts with names `node0`, `node1`, `node2` and so on in your hosts file
3. Execute playbooks in the following order: `binary`, `config`, `service`, `start`
4. Make sure nodes are running and writing blocks by running `status` playbook

PS. Don't forget to define node ips in their configurations so they will be able to find each other

## Setting up client: