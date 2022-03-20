# Genesis

An Ansible role that initializes and configures the DCL Genesis node.

## Requirements

None

## Role Variables

```yaml
chain_id: test-net
```

The unique chain ID to identify the network.

```yaml
dcld:
  user: dcl
  group: dcl
  home: /var/lib/dcl
```

The *user* and *group* to be used by OS to run the DCL service. The *home* var
specifies the path to store DCL config information.

## Dependencies

1. `bootstrap` ansible role.

## Example Playbook

example inventory.yaml

```yaml
all:
  vars:
    chain_id: dev-net
  hosts:
    node0:
      moniker: dev-net-node0
      accounts:
        - name: jack
          passphrase: test1234
          roles:
            - NodeAdmin
            - Trustee
```

in your playbook:

```yaml
- name: init DCL genesis node
  hosts: genesis
  serial: 1
  roles:
    - genesis
```
