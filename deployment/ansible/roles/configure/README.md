# Configure

An Ansible role that configures DCL client, app and config files

## Requirements

None

## Role Variables

```yaml
chain_id: test-net
```

The unique chain ID to identify the network.

## Dependencies

None

## Example Playbook

example inventory.yaml

```yaml
all:
  vars:
    chain_id: dev-net
```

in your playbook:

```yaml
- name: setup validators
  hosts: validators
  roles:
    - bootstrap
    - role: configure
      config:
        p2p:
          pex: false
          persistent_peers:
          addr_book_strict: false
        statesync:
          rpc_servers:
      app:
        state-sync:
          snapshot-interval: snapshot-interval
          snapshot-keep-recent: snapshot-keep-recent
```
