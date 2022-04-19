# Bootstrap

An Ansible role that fetches DCL binary and configures the service

## Requirements

None

## Role Variables

```yaml
dcl_version: 0.9.0
```

The DCL binary version to be used for the deployment.

```yaml
dcl_home: /var/lib/dcl/.dcl
cosmovisor:
  user: cosmovisor
  group: dcl
```

The *user* and *group* to be used by OS to run the cosmovisor service. The
*dcl_home* var specifies the path to store DCL config information.

## Dependencies

None

## Example Playbook

example inventory.yaml

```yaml
all:
  vars:
    dcl_version: 0.9.0
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
