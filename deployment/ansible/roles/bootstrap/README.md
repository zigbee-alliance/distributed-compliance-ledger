# Bootstrap

An Ansible role that fetches DCL binary and configures the service. This role
creates keys defined on `accounts` list variable and set up the command-line
interface.

## Requirements
None

## Role Variables

```yaml
chain_id: test-net
```

The unique chain ID to identify the network.

```yaml
accounts:
  - name: user1
    passphrase: password123
    roles:
      - NodeAdmin
      - Trustee
```

A list of DCL accounts to be created on a specific target node.

```yaml
dcld:
  version: 0.7.0
```

The DCL binary version to be used for the deployment.

```yaml
dcld:
  user: dcl
  group: dcl
  home: /var/lib/dcl
```

The *user* and *group* to be used by OS to run the DCL service. The *home* var
specifies the path to store DCL config information.

## Dependencies
None

## Example Playbook

example inventory.yaml

```yaml
all:
  vars:
    chain_id: dev-net 
  hosts:
    node0:
      accounts:
        - name: jack
          passphrase: test1234
          roles:
            - NodeAdmin
            - Trustee
    node1:
      accounts:
        - name: alice
          passphrase: s3cr3t123
          roles:
            - NodeAdmin
            - Trustee
    node2:
      accounts:
        - name: bob
          passphrase: admin1234
          roles:
            - NodeAdmin
            - Trustee
    node3:
      accounts:
        - name: anna
          passphrase: test1234
          roles:
            - NodeAdmin
```

in your playbook:

```yaml
- name: bootstrap DCL nodes
  hosts: all
  become: true
  roles:
    - bootstrap
  tasks:
    - name: start DCL service
      service:
        name: dcld
        state: started
```
