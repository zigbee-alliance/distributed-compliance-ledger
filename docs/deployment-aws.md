# DCL Deployment on AWS using Terraform and Ansible

## Overview

This document describes all necessary steps to deploy a new DCL network on AWS cloud in accordance with this [design document](./deployment-design-aws.md).

## Environment
Officially supported OS for development is `Ubuntu 20.04 LTS` and all the following intructions are tested on it.
But you are free to use any other environment that supports [Terraform][1] and [Ansible][4].
## Requirements
### 1. Clone [DCL][5] project
```
git clone https://github.com/zigbee-alliance/distributed-compliance-ledger.git
```
### 2. Install [Terraform][2] CLI
```bash
sudo apt-get update && sudo apt-get install -y gnupg software-properties-common curl
curl -fsSL https://apt.releases.hashicorp.com/gpg | sudo apt-key add -
sudo apt-add-repository "deb [arch=amd64] https://apt.releases.hashicorp.com $(lsb_release -cs) main"
sudo apt-get update && sudo apt-get install terraform
```

### 3. Install `Python` and `pip`
```bash
sudo apt-get update
sudo apt-get install -y --no-install-recommends python3
sudo apt install python3-pip
sudo apt install python3-testresources
```

### 4. Install [Ansible][4] and its dependencies
    
Run the following commands from the [DCL][5] project home
```bash
sudo pip install -r deployment/requirements.txt
ansible-galaxy install -r deployment/galaxy-requirements.yml 
```

### 5. Install [Dasel][6] to make conversions between `json` and `yaml`
```bash
sudo wget -qO /usr/local/bin/dasel https://github.com/TomWright/dasel/releases/latest/download/dasel_linux_amd64
sudo chmod a+x /usr/local/bin/dasel
```

## Terraform and Ansible Configuration
### 1. Set up an AWS user for use with Terraform
    
Create credentials file [`~/.aws/credentials`] with the following content:
```
[default]
aws_access_key_id = <access_key_id_here>
aws_secret_access_key = <secret_access_key_here>
```
> **_Note:_** Your account must have enough privileges to manage all AWS resources required by terraform

### 2. Disable host key checking for Ansible (to avoid host key checking when Ansible connects to AWS instances using ssh)

Create Ansible configuration file [`~/.ansible.cfg`] with the following content:
```
ANSIBLE_HOST_KEY_CHECKING=False
```

## Deployment Configuration

### 1. Configure AWS infrastructure parameters in [`deployment/terraform/aws/terraform.tfvars`]
    
### AWS Regions:
```
region_1 = "us-west-1"
region_2 = "us-east-2"
```
- Selects two regions where nodes will be created

### Validator:
```
validator_config = {
    instance_type = "t3.medium"
}
```
- Validator node is created in `region_1` by default


### Private Sentries:

```
private_sentries_config = {
    enable        = true
    nodes_count   = 2
    instance_type = "t3.medium"
}
```
- Private sentry nodes are created in the region as Validator by default
- Can be disabled by setting `enable = false`
- Only one instance of private sentry is created with static ip address

### Public Sentries:
```
public_sentries_config = {
    enable        = true
    enable_ipv6   = false
    nodes_count   = 2
    instance_type = "t3.medium"

    regions = [
        1,
        2
    ]
}
```
- Requires `Private Sentries` being enabled
- Can be configured to have IPv6 addresses using `enable_ipv6 = true`
    > **_Note:_** Number of available IPv4 static addresses is restricted to 5 per region on AWS by default
- Can be configured to run on multiple regions

### Observers:
```
observers_config = {
    enable           = true
    nodes_count      = 3
    instance_type    = "t3.medium"
    root_domain_name = "matterprotocol.com"
    enable_tls       = true

    regions = [
        1,
        2
    ]
}
```

- Requires `Private Sentries` being enabled
- When `root_domain_name` parameter is set, will be serving requests under `on.<root_domain_name>` domain
    - Name servers for `root_domain_name` servers must be pointed to AWS and managed by the AWS account
- TLS can be enabled or disabled
- Can be configured to run on multiple regions

### 2. Configure Ansible inventory variables in [`deployment/ansible/inventory/aws/group_vars/all.yaml`]
```
chain_id: test-net2

dcl_home: /var/lib/dcl/.dcl
dcl_version: 0.9.0
dcld:
  version: "{{ dcl_version }}"
  path: "{{ dcl_home }}/cosmovisor/genesis/bin/dcld"
cosmovisor:
  version: "{{ dcl_version }}"
  user: cosmovisor
  group: dcl
  path: /usr/bin/cosmovisor
  home: "{{ dcl_home | dirname }}"

dcld_checksums:
  0.9.0: c333d828a124e527dd7a9c0170f77d61ad07091d9f6cd61dd0175a36b55aadce
cosmovisor_checksums:
  0.9.0: c05705efe5369b9d83e65ef7b252bd7c610eec414ae3f6c08681bcf49dc38e6d

dcld_download_url: "https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v{{ dcld.version }}/dcld"
dcld_binary_checksum: "sha256:{{ dcld_checksums[dcld.version] }}"
cosmovisor_download_url: "https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v{{ cosmovisor.version }}/cosmovisor"
cosmovisor_binary_checksum: "sha256:{{ cosmovisor_checksums[cosmovisor.version] }}"
```
- Specify DCL network `chain_id`
- Specify `dcld` and `cosmovisor` versions and checksums

## Deployment
### 1. Run terraform from [`deployment/terraform/aws`]
```
terraform apply
```

> **_Note:_** Terraform asks a confirmation before applying changes

### 2. Generate ansible inventory from terraform output
Once terraform completes successfully, run the following command to genarate ansible inventory file:
```bash
terraform output -json ansible_inventory | dasel -r json -w yaml . > ../../ansible/inventory/aws/aws_all.yaml
```

### 3. Run Ansible
Run the following command from the project home:
```bash
ansible-playbook -i ./deployment/ansible/inventory/aws  -u ubuntu ./deployment/ansible/deploy.yml
```
- Ansible provisioning can take several minutes depending on number of nodes being provisioned

## Deployment Verification
### 1. Verify [`deployment/persistent_chains/<chain_id>/genesis.json`] is created
- `<chain_id>` - chain id of the network specified in Ansible inventory variables
### 2. Verify `Observers` REST endpoint is available under `http(s)://on.<root_domain_name>` using your browser
### 3. Verify `Observers` RPC endpoint is available under `http(s)://on.<root_domain_name>:26657` using your browser
### 4. Verify `Observers` gRPC endpoint is available under `http(s)://on.<root_domain_name>:8443` using postman (or similar tool)

- `<root_domain_name>` - domain name specified in terraform `Observers` config


[1]: https://www.terraform.io/
[2]: https://learn.hashicorp.com/tutorials/terraform/install-cli
[3]: https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html
[4]: https://www.ansible.com
[5]: https://github.com/zigbee-alliance/distributed-compliance-ledger.git
[6]: https://github.com/TomWright/dasel