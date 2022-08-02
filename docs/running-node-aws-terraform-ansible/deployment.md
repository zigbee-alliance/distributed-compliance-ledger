# DCL Deployment on AWS using Terraform and Ansible
<!-- markdownlint-disable MD033 -->

## Overview

This document describes all necessary steps to deploy a new DCL network on AWS cloud in accordance with this [design document](../deployment-design-aws.md).

## Prerequisites

Make sure you have all [prerequisites](./prerequisites.md) set up

## Terraform and Ansible Configuration

### 1. Set up an AWS user for use with Terraform

Create credentials file [`~/.aws/credentials`] with the following content:

```text
[default]
aws_access_key_id = <access_key_id_here>
aws_secret_access_key = <secret_access_key_here>
```

> **_Note:_** Your account must have enough privileges to manage all AWS resources required by terraform

### 2. Disable host key checking for Ansible (to avoid host key checking when Ansible connects to AWS instances using ssh)

Create Ansible configuration file [`~/.ansible.cfg`] with the following content:

```text
[defaults]
HOST_KEY_CHECKING=False
```

## Deployment Configuration

### 1. Configure AWS infrastructure parameters

[`deployment/terraform/aws/terraform.tfvars`]

#### AWS Regions

```hcl
region_1 = "us-west-1"
region_2 = "us-east-2"
```

- Selects two regions where nodes will be created

#### (Genesis) Validator

```hcl
validator_config = {
    instance_type = "t3.medium"
    is_genesis    = true
}
```

- Set `is_genesis = false` to deploy just a validator node (not genesis). This option requires:
  - Putting `genesis.json` file of an existing network to the following path before the step [run-ansible](#4-run-ansible)

    ```text
    deployment/persistent_chains/<chain-id>/genesis.json
    ```

    where `<chain-id>` is the chain ID of a network being joined.
    - For `testnet-2.0` the genesis file is already in place

        ```text
        deployment/persistent_chains/testnet-2.0/genesis.json
        ```

  - Manually adding the validator to the network (see [making node a validator](../running-node-ansible/vn.md#make-your-node-a-validator-target-machine)) after the step [run-ansible](#4-run-ansible)

- Manually set `persistent_peers` string in validator config (only if `Private Sentries` are disabled)
  [`deployment/ansible/roles/configure/vars/validator.yml`]

  ```yaml
  config:
    p2p:
      persistent_peers: "<node1-ID>@<node1-IP>:26656,..."
    ...
  ```

  - For `testnet-2.0` get the latest `persistent_peers` string from the CSA slack channel

- Validator/Genesis node is created in `region_1` by default

#### Private Sentries (optional)

```hcl
private_sentries_config = {
    enable        = true
    nodes_count   = 2
    instance_type = "t3.medium"
}
```

- Private sentry nodes are created in the region as Validator by default
- Can be disabled by setting `enable = false`
- Only one instance of private sentry is created with static ip address
- Manually set `persistent_peers` string in private sentry config
  [`deployment/ansible/roles/configure/vars/private-sentry.yml`]

  ```yaml
  config:
    p2p:
      persistent_peers: "<node1-ID>@<node1-IP>:26656,..."
    ...
  ```

  - For `testnet-2.0` get the latest `persistent_peers` string from the CSA slack channel

#### Public Sentries (optional)

```hcl
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

#### Observers (optional)

```hcl
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

#### Prometheus (optional)

```hcl
prometheus_config = {
  enable        = true
  instance_type = "t3.small"
}
```

- Requires `Private Sentries` being enabled
- When enabled runs a dedicated Prometheus server on Private Sentries VPC to collect Tendermint metrics from all DCL nodes
- Collected metrics are written to AWS [AMP workspace](https://aws.amazon.com/prometheus/)

### 2. Set DCL network chain ID in ansible inventory

[`deployment/ansible/inventory/aws/group_vars/all.yaml`]

```yaml
chain_id: test-net
...
```

<details>
<summary>Example for Testnet 2.0 (clickable) </summary>

```yaml
chain_id: testnet-2.0
...
```

</details>

## Deployment

### 1. Run terraform

```bash
cd deployment/terraform/aws
terraform apply
```

> **_Note:_** Terraform asks a confirmation before applying changes

### 2. Generate ansible inventory from terraform output

```bash
terraform output -json ansible_inventory | dasel -r json -w yaml . > ../../ansible/inventory/aws/aws_all.yaml
```

### 3. (Optional) Consider enabling state sync

When joining an existing pool, you may want to enable state sync for all the nodes.
To do so, you should set state sync parameters:

```yaml
config:
...
  statesync:
    enable: true
    rpc_servers: "http(s):<node1-IP>:26657,..."
    trust_height: <trust-height>
    trust_hash: "<trust-hash>"
...
```

in the following ansible config files

```text
deployment/ansible/roles/configure/vars/
  validator.yml
  private-sentry.yml
  observer.yml
  public-senrty.yml
  seed.yml
```

<details>
<summary>Example for Testnet 2.0 (clickable) </summary>

```yaml
config:
  statesync:
    enable: true
    rpc_servers: "https://on.test-net.dcl.csa-iot.org:26657,https://on.test-net.dcl.csa-iot.org:26657"
    
```

</details>

> **_NOTE:_**  You should provide at least 2 addresses for `rpc_servers`. It can be 2 identical addresses

You can use the following command to obtain `<trust-height>` and `<trust-hash>` of your network

```bash
curl -s http(s)://<host>:<port>/commit | jq "{height: .result.signed_header.header.height, hash: .result.signed_header.commit.block_id.hash}"
```

<details>
<summary>Example for Testnet 2.0 (clickable) </summary>

```bash
curl -s https://on.test-net.dcl.csa-iot.org:26657/commit | jq "{height: .result.signed_header.header.height, hash: .result.signed_header.commit.block_id.hash}"
```

</details>

- `<host>` - RPC endpoint host of the network being joined
- `<port>` - RPC endpoint port of the network being joined

### 4. Run Ansible

Run the following command from the project home:

```bash
ansible-playbook -i ./deployment/ansible/inventory/aws  -u ubuntu ./deployment/ansible/deploy.yml
```

- Ansible provisioning can take several minutes depending on number of nodes being provisioned

### 5. (For non-genesis validator nodes) Add Validator to the network

- Manually add the validator to the network (see [making node a validator](../running-node-ansible/vn.md#make-your-node-a-validator-target-machine))

## Deployment Verification

1. Verify [`deployment/persistent_chains/<chain_id>/genesis.json`] is created
    - `<chain_id>` - chain ID of the network specified in Ansible inventory variables
2. Verify `Observers` REST endpoint is available under `http(s)://on.<root_domain_name>` using your browser
3. Verify `Observers` RPC endpoint is available under `http(s)://on.<root_domain_name>:26657` using your browser
4. Verify `Observers` gRPC endpoint is available under `http(s)://on.<root_domain_name>:8443` using postman (or similar tool)

- `<root_domain_name>` - domain name specified in terraform `Observers` config

## Health and Monitoring

### Logs

- Logs from all DCL nodes are collected by AWS cloudwatch agent and available at [AWS Cloudwatch](https://aws.amazon.com/cloudwatch/)
- Logs are collected per DCL node and AWS region

### Monitoring

- Metrics of AWS EC2 instances where the DCL nodes run on are available at [AWS Cloudwatch](https://aws.amazon.com/cloudwatch/)
- Metrics of underlying blockchain engine (Tendermint) are pushed to [AWS AMP Service](https://aws.amazon.com/prometheus/) when prometheus is enabled
- Use

    ```bash
    terraform output prometheus_endpoint
    ```

    to get prometheus endpoint which can be used by Grafana to visualize metrics. See detailed instructions [here][1]

[1]: https://aws.amazon.com/blogs/opensource/using-amazon-managed-service-for-prometheus-to-monitor-ec2-environments/
