# DCL Deployment on AWS using Terraform and Ansible
<!-- markdownlint-disable MD033 -->

## Overview

This document describes all necessary steps to deploy a new DCL network on AWS cloud in accordance with this [design document](../deployment-design-aws.md).

Please note, that the current version of the automation scripts can not be used to run just an Observer node without a VN.
If you need to run an ON connected to another organization nodes, please follow the manual steps from [running-node-manual/on.md](../running-node-manual/on.md).

The deployment consits of AWS infrastructure and application deployment steps automated by Terraform and Ansible tools respectively.

## Prerequisites

Make sure you have all [prerequisites](./prerequisites.md) set up

## 1. Infrastructure Deployment

### 1.1 Configuration

#### 1.1.1 Cloud authentication

Terraform scirpts appliance requires cloud authentication configuration.

The details below provide the options how do that for each cloud.

<details>
<summary> AWS </summary>

Create credentials file [`~/.aws/credentials`] with the following content:

```text
[default]
aws_access_key_id = <access_key_id_here>
aws_secret_access_key = <secret_access_key_here>
```

> **_Note:_** Your account must have enough privileges to manage all AWS resources required by terraform

Please see other authentication configuration options [here](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#authentication-and-configuration).

</details>

<details>
<summary> GCP </summary>

*   install `gcloud` CLI tool following [this](https://cloud.google.com/sdk/docs/install) docs
*   run `gcloud auth application-default login`

Please see the additional details [here](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs#authenticating-to-azur://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/getting_started).

</details>


<details>
<summary> Azure </summary>

The following steps are based on Azure CLI (see the details [here](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/guides/azure_cli)):

*   install `az` CLI tool using [this](https://learn.microsoft.com/en-us/cli/azure/install-azure-cli?view=azure-cli-latest) docs
*   run `az login`
    *   once logged in you may run `az account list` to list the availble subscriptions
    *   run `az account set --subscription="SUBSCRIPTION_ID"` in case multiple subscriptions are listed 
*  run `export ARM_SUBSCRIPTION_ID="SUBSCRIPTION_ID"`

Please see other authentication configuration options [here](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs#authenticating-to-azure).

</details>



#### 1.1.2 Terraform backend

The general recommendation is:

*   use [`local`](https://developer.hashicorp.com/terraform/language/v1.5.x/settings/backends/local) backend only for development
*   use one of the remote [backends](https://developer.hashicorp.com/terraform/language/v1.5.x/settings/backends/configuration#available-backends)
available for your cloud provider

The default backends:

*   AWS: [`s3`](https://developer.hashicorp.com/terraform/language/v1.5.x/settings/backends/s3)
*   GCP: [`gcs`](https://developer.hashicorp.com/terraform/language/v1.5.x/settings/backends/gcs)
*   Azure: [`azurerm`](https://developer.hashicorp.com/terraform/language/v1.5.x/settings/backends/azurerm)

> **_Note_** In case you need to switch a backend for already deployed infrastracture:
> * configure new backend
> * use `-migrate-state` option for `terraform init` command (see the details [here](https://developer.hashicorp.com/terraform/cli/v1.5.x/commands/init#backend-initialization))

Please see below more backend configuration details for each supported cloud.

<details>
<summary> AWS </summary>

By default AWS infrastructure backend is set as `s3` (see [`deployment/terraform/aws/backend.tf`](../../deployment/terraform/aws/backend.tf)).

S3 backend configuration implies:

*   existent S3 bucket
*   (optional but recommended) DynamoDB table to support [remote state locking](https://developer.hashicorp.com/terraform/language/v1.5.x/state/locking)
    *   **_Note_** The table must have a partition key named `LockID` with a type of `String`.

To complete the configuration please specify:

*   S3 bucket name
*   S3 object key
*   region
*   (optional) DynamoDB table name

using one of the following ways:

*   as parameters in [`deployment/terraform/aws/backend.tf`](../../deployment/terraform/aws/backend.tf)
*   as a separate configuration file:
    *   use `-backend-config=<config.s3.tfbackend>` (more [details](https://developer.hashicorp.com/terraform/language/v1.5.x/settings/backends/configuration#partial-configuration))
    *   example: [`deployment/terraform/aws/config.s3.tfbackend.example`](../../deployment/terraform/aws/config.s3.tfbackend.example)
*   as command line arguments
*   interactively during terraform initialization

Please see also Terraform [docs](https://developer.hashicorp.com/terraform/language/v1.5.x/settings/backends/s3) for the details.
</details>


<details>
<summary> GCP </summary>

By default GCP infrastructure backend is set as `gcs` (see [`deployment/terraform/gcp/backend.tf`](../../deployment/terraform/gcp/backend.tf)).

S3 backend configuration implies:

*   existent S3 bucket
*   (optional but recommended) DynamoDB table to support [remote state locking](https://developer.hashicorp.com/terraform/language/v1.5.x/state/locking)
    *   **_Note_** The table must have a partition key named `LockID` with a type of `String`.

To complete the configuration please specify:

*   Cloud Storage bucket name
*   (optional) prefix inside the bucket

using one of the following ways:

*   as parameters in [`deployment/terraform/gcp/backend.tf`](../../deployment/terraform/gcp/backend.tf)
*   as a separate configuration file:
    *   use `-backend-config=<config.gcs.tfbackend>` (more [details](https://developer.hashicorp.com/terraform/language/v1.5.x/settings/backends/configuration#partial-configuration))
    *   example: [`deployment/terraform/gcp/config.gcs.tfbackend.example`](../../deployment/terraform/gcp/config.gcs.tfbackend.example)
*   as command line arguments
*   interactively during terraform initialization

Please see also Terraform [docs](https://developer.hashicorp.com/terraform/language/v1.5.x/settings/backends/gcs) for the details.

</details>

<details>
<summary> Azure </summary>

By default AWS infrastructure backend is set as `azurerm` (see [`deployment/terraform/azure/backend.tf`](../../deployment/terraform/azure/backend.tf)).

`azurerm` backend configuration implies:

*   existent storage account along with storage account container (see [prerequisites](./prerequisites.md#5-preliminary-cloud-configuration) for the details)

To complete the configuration please specify:

*   resource group name
*   storage account name
*   container name
*   key name

using one of the following ways:

*   as parameters in [`deployment/terraform/azure/backend.tf`](../../deployment/terraform/azure/backend.tf)
*   as a separate configuration file
    *   use `-backend-config=<config.azurerm.tfbackend>` option (more [details](https://developer.hashicorp.com/terraform/language/v1.5.x/settings/backends/configuration#partial-configuration))
    *   example: [`deployment/terraform/azure/config.azurerm.tfbackend.example`](../../deployment/terraform/azure/config.azurerm.tfbackend.example)
*   as command line arguments
*   interactively during terraform initialization

Please see also Terraform [docs](https://developer.hashicorp.com/terraform/language/v1.5.x/settings/backends/azurerm) for the details.
</details>


#### 1.1.3 Configure infrastructure parameters

Deployment configuration for each of the supported clouds includes common set of parameters:

*   common parameters (regions / locations, tags, other cloud specific ones)
*   validator parameteres
*   private sentries parameteres
*   public sentries parameteres
*   observers parameteres
*   prometheus (**only for AWS for now**)

And should be defined in the floowing files:

*   AWS [`deployment/terraform/aws/terraform.tfvars`](../../deployment/terraform/aws/terraform.tfvars)
*   GCP [`deployment/terraform/aws/terraform.tfvars`](../../deployment/terraform/gcp/terraform.tfvars)
*   Azure [`deployment/terraform/aws/terraform.tfvars`](../../deployment/terraform/azure/terraform.tfvars)

##### 1.1.3.1  Common parameters

**Regions / Locations**

Selects two regions (locations in Azure) where nodes will be created. Examples:

<details>
<summary> AWS </summary>

```hcl
region_1 = "us-west-1"
region_2 = "us-east-2"
```
</details>


<details>
<summary> GCP </summary>

```hcl
region_1 = "us-east1"
region_2 = "us-west1"
```
</details>


<details>
<summary> Azure </summary>

```hcl
location_1 = "eastus"
location_2 = "westus2"
```

>**_Note:_** ensure you have enough quota for the planned virtual machine sizes
in the required locations and request the new quotes if needed (more details in
[docs](https://learn.microsoft.com/en-us/azure/quotas/per-vm-quota-requests))

</details>

**Commmon tags**

These tags will be applied to all the resourcers that support them

<details>
<summary> AWS & Azure </summary>

```hcl
common_tags = {
  project		   = "DCL"  # (optional, default - "DCL")
  environment      = "issue-123" (optional, default - workspace name)
  purpose          = "some context details" (optional)
  created-by       = "user@domain.com" (optional)
}
```
</details>


<details>
<summary> GCP </summary>

```hcl
common_labels= {
  project		   = "DCL"  # (optional, default - "DCL")
  environment      = "issue-123" (optional, default - workspace name)
  purpose          = "some-context-details" (optional)
  created-by       = "username" (optional)
}
```
</details>

**Cloud specific parameters**

<details>
<summary> GCP </summary>

```hcl
project_id = "<your-project-id>"
```
</details>

<details>
<summary> Azure </summary>

```hcl
resource_group_name = "<your-resource-group-name>"
```
</details>

##### 1.1.3.2 (Genesis) Validator

*   Set `is_genesis = false` to deploy just a validator node (not genesis).
*   Validator/Genesis node is created in `region_1` by default

<details>
<summary> AWS </summary>

```hcl
validator_config = {
    instance_type = "t3.medium"
    is_genesis    = true
}
```
</details>


<details>
<summary> GCP </summary>

```hcl
validator_config = {
  instance_type = "e2-standard-2"
  is_genesis    = true
}
```
</details>


<details>
<summary> Azure </summary>

```hcl
validator_config = {
  instance_size = "Standard_B1s"
  is_genesis    = true
  enable_encryption_at_host = true
}
```

>**_Note:_** choose `instance_size` taking into account the following:
>  *    limits & quotas for you subscription
>  *    SKUs [availbility](https://learn.microsoft.com/en-us/azure/azure-resource-manager/troubleshooting/error-sku-not-available?tabs=azure-cli#solution)

>**_Note:_** if disk encryption is requied it's necessary to [enable](https://learn.microsoft.com/en-us/azure/virtual-machines/linux/disks-enable-host-based-encryption-cli)
that for the subscription

</details>

##### 1.1.3.3 Private Sentries (optional)

*   Private sentry nodes are created in the same region as Validator by default
*   Can be disabled by setting `enable = false`

<details>
<summary> AWS </summary>

```hcl
private_sentries_config = {
    enable        = true
    nodes_count   = 2
    instance_type = "t3.medium"
}
```
</details>


<details>
<summary> GCP </summary>

```hcl
private_sentries_config = {
  enable        = true
  nodes_count   = 2
  instance_type = "e2-standard-2"
}
```
</details>

<details>
<summary> Azure </summary>

```hcl
private_sentries_config = {
  enable        = true
  nodes_count   = 2
  instance_size = "Standard_B2s"
}
```
</details>


##### 1.1.3.4 Public Sentries (optional)

*   **Requires** `Private Sentries` being enabled
*   Can be configured to have IPv6 addresses using `enable_ipv6 = true` (**only for AWS for now**)
*   Can be configured to run on multiple regions

<details>
<summary> AWS </summary>

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

> **_Note:_** Number of available IPv4 static addresses is restricted to 5 per region on AWS by default
```
</details>

<details>
<summary> GCP </summary>

```hcl
public_sentries_config = {
  enable        = true
  region1_nodes_count   = 2
  region2_nodes_count   = 2
  instance_type = "e2-standard-2"
}
```

> **_Note:_** The properties block differs from other clouds:
> * you set number of nodes for each region separately
> * the reason: in GCP a single virtual network span both regions

</details>

<details>
<summary> Azure </summary>

```hcl
public_sentries_config = {
  enable        = true
  nodes_count   = 2
  instance_size = "Standard_B2s"

  locations = [
    1,
    2
  ]
}
```
</details>

##### 1.1.3.5 Observers (optional)

*    **Requires** `Private Sentries` being enabled
*    Can be configured to run on multiple regions


<details>
<summary> AWS </summary>

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

*    When `root_domain_name` parameter is set, will be serving requests under `on.<root_domain_name>` domain
    *   Name servers for `root_domain_name` servers must be pointed to AWS and managed by the AWS account
*    TLS can be enabled or disabled
</details>


<details>
<summary> GCP </summary>

```hcl
observers_config = {
  enable           = true
  region1_nodes_count      = 3
  region2_nodes_count      = 3
  instance_type    = "e2-standard-2"
}
```

> **_Note:_** The properties block differs from other clouds:
> * you set number of nodes for each region separately
> * the reason: in GCP a single virtual network span both regions
</details>


<details>
<summary> Azure </summary>

```hcl
observers_config = {
  enable           = true
  nodes_count      = 3
  instance_size    = "Standard_B2s"

  locations = [
    1,
    2
  ]
  azs              = [[2], [2]]
}
```

> **_Note:_** The properties block differs from other clouds:
> * additional property `azs` allows to specify observer nodes distribution across availbility zones as follows:
>   *   `azs` is a list of number sets of zones logical ids to use
>   *   only first 2 items are expected (equal to number of locations)
>   *   `null` or empty set are allowed and treated as "use all the available zones for the location"
>       *   **_Note:_** this might be not the best option since not all the zones in a location have quotes for the particular VM sizes
>       *   please use `az vm list-skus --location <location> --resource-type virtualMachines --zone --all --output table --size <vn-size>` to explore
>   *   examples:
>       *   `[[2], [2]]` zones with logical ids `2` in each location
>       *   `[[1, 2], null]` or `[[1, 2], []]` or `[[1, 2]]` - zones `1` and `2` for location 1, all available zones for location 2
>       *   `null` or `[[], []]` or omitted - all availble zones for each location

</details>

##### 1.1.3.6 Prometheus (optional)

> **_Note:_** Only for AWS cloud for now

*   **Requires** `Private Sentries` being enabled

<details>
<summary> AWS </summary>

```hcl
prometheus_config = {
  enable        = true
  instance_type = "t3.small"
}
```

*   When enabled runs a dedicated Prometheus server on Private Sentries VPC to collect Tendermint metrics from all DCL nodes
*   Collected metrics are written to AWS [AMP workspace](https://aws.amazon.com/prometheus/)
</details>

### 1.2 Deployment

#### 1.2.1 Initialize terraform

```bash
cd deployment/terraform/aws

terraform init -backend-config=<backend-config-file> # in case backend configuration is in a file
```

where `<backend-config-file>` is the backend configuration file ([AWS S3 backend example](../../deployment/terraform/aws/config.s3.tfbackend.example)).

(optional) Create/Activate the deployment workspace:

```bash
terraform workspace select -or-create=true <workspace-name>
```

where `<workspace-name>` is the name of the Terraform workspace (e.g. `prod` or `issue-123`).

#### 1.2.2 Run terraform

Before applying the configuration it is recommended to make the checks:

```bash
terraform workspace show

terraform plan
```

Apply the configuration:

```bash
terraform apply
```

> **_Note:_** Terraform asks a confirmation before applying changes

#### 1.2.3 Generate ansible inventory from terraform output

```bash
terraform output -raw ansible_inventory_yaml > ../../ansible/inventory/cloud/all.yaml
```

## 2. Application Deployment

### 2.1 Configuration

#### 2.1.1 SSH connection configuration

##### 2.1.1.1 Disable host key checking for Ansible

This is done to avoid host key checking when Ansible connects to AWS instances using ssh.

Create Ansible configuration file [`~/.ansible.cfg`] with the following content:

```text
[defaults]
HOST_KEY_CHECKING=False
```

##### 2.1.1.2 Configure SSH authentication

The recommended way is to use `ssh-agent`, e.g.

```bash

# considering the agent is already running

ssh-add <path-to-ssh-private-key>
```

Please read more about Ansible ssh connection configuration [here](https://docs.ansible.com/ansible/latest/collections/ansible/builtin/ssh_connection.html)

#### 2.1.2 Set base DCL network parameters

[`deployment/ansible/inventory/cloud/group_vars/all.yaml`](../../deployment/ansible/inventory/cloud/group_vars/all.yaml)

```yaml
chain_id: <chain-id>
company_name: <company>
dcl_version: <version>
...
```

where:
*   `<chain-id>`: an unique chain ID every network must have
*   `<company>`: the company name that owns the node, will be used (along with node type) to generate a human-readable node username
*   `<version>`: one of the available DCL [releases](https://github.com/zigbee-alliance/distributed-compliance-ledger/releases) **without forwarding `v`** (e.g. `1.4.4`)

<details>
<summary>Example for Testnet 2.0 (clickable) </summary>

```yaml
chain_id: testnet-2.0
company_name: YourCompany
dcl_version: 0.12.0
...
```

</details>

<details>
<summary>Example for Mainnet (clickable) </summary>

```yaml
chain_id: main-net
company_name: YourCompany
dcl_version: 0.12.0
...
```

</details>


#### 2.1.3 (For non-genesis validator nodes) Configure genesis data

Put `genesis.json` file of an existing network as `deployment/persistent_chains/<chain-id>/genesis.json`,
where `<chain-id>` is the chain ID of a network being joined.

**Note** For `testnet-2.0` and `main-net` the genesis files are already in place

  - [`deployment/persistent_chains/testnet-2.0/genesis.json`](../../deployment/persistent_chains/testnet-2.0/genesis.json)
  - [`deployment/persistent_chains/main-net/genesis.json`](../../deployment/persistent_chains/main-net/genesis.json)

#### 2.1.4 Persistent peers configuration

[Persistent peers](https://docs.tendermint.com/v0.34/tendermint-core/using-tendermint.html#persistent-peer) are the nodes
this node is constantly connected with.

The parameter is set as `config.p2p.persistent_peers` for each node type separately 
and has the following format:

```yaml
config:
  p2p:
    persistent_peers: "<node1-ID>@<node1-IP>:26656,<node2-ID>@<node2-IP>:26656,..."
  ...
```

*   node ID can be found using `dcld tendermint show-node-id` command
*   node IP can be either from LAN IP or WAN IP depending on the node type and the deployment architecture (see below)

In particular:

*   validator (**required if private sentries are disabled, otherwise by default resolved from private sentries**) [`deployment/ansible/roles/configure/vars/validator.yml`](../../deployment/ansible/roles/configure/vars/validator.yml)
*   private-sentry (**required if private sentries are enabled, otherwise by default resolved from validators**) [`deployment/ansible/roles/configure/vars/private-sentry.yml`](../../deployment/ansible/roles/configure/vars/private-sentry.yml)
*   observer (optional, by default resolved from private sentries) [`deployment/ansible/roles/configure/vars/observer.yml`](../../deployment/ansible/roles/configure/vars/observer.yml)
*   public-sentry (optional, by default resolved from private sentries) [`deployment/ansible/roles/configure/vars/public-sentry.yml`](../../deployment/ansible/roles/configure/vars/public-sentry.yml)
*   seed (optional, by default resolved from public sentries) [`deployment/ansible/roles/configure/vars/seed.yml`](../../deployment/ansible/roles/configure/vars/seed.yml)

**Note** For `testnet-2.0` or `main-net` get the latest `persistent_peers` string from the CSA slack channel

#### 2.1.5 Consider enabling state sync

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

<details>
<summary>Example for Mainnet (clickable) </summary>

```yaml
config:
  statesync:
    enable: true
    rpc_servers: "https://on.dcl.csa-iot.org:26657,https://on.dcl.csa-iot.org:26657"

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

<details>
<summary>Example for Mainnet (clickable) </summary>

```bash
curl -s https://on.dcl.csa-iot.org:26657/commit | jq "{height: .result.signed_header.header.height, hash: .result.signed_header.commit.block_id.hash}"
```

</details>

*   `<host>` - RPC endpoint host of the network being joined
*   `<port>` - RPC endpoint port of the network being joined

### 2.2 Deployment

#### 2.2.1 Run Ansible

Switch the directory to ansible:

```bash
cd ./deployment/ansible
```

Run the following command:

```bash
ansible-playbook -i inventory/aws  -u ubuntu deploy.yml
```

*   Ansible provisioning can take several minutes depending on number of nodes being provisioned

#### 2.2.2 (For non-genesis validator nodes) Add Validator to the network

*   Manually add the validator to the network (see [making node a validator](../running-node-ansible/vn.md#make-your-node-a-validator-target-machine))

## 3. Deployment Verification

1. Verify `deployment/persistent_chains/<chain_id>/genesis.json` is created
    - `<chain_id>` - chain ID of the network specified in Ansible inventory variables
2. Verify `Observers` REST endpoint is available under `http(s)://on.<root_domain_name>` using your browser
3. Verify `Observers` RPC endpoint is available under `http(s)://on.<root_domain_name>:26657` using your browser
4. Verify `Observers` gRPC endpoint is available under `http(s)://on.<root_domain_name>:8443` using Postman (or similar tool)

*   `<root_domain_name>` - domain name specified in terraform `Observers` config

## 4. DCL Web UI integration

### Using AWS Amplify Service

1. Make a fork from [DCL UI](https://github.com/Comcast/dcl-ui) repo to your own github account
2. Access your AWS Amplify console and follow the [instructions][1]
    - skip instructions for setting up a backend since you are going to deploy only static Vue.js app
    - use the following environment variables:
      ```.env
      VUE_APP_DCL_API_NODE=https://on.<root_domain_name>
      VUE_APP_DCL_RPC_NODE=https://on.<root_domain_name>:26657
      VUE_APP_DCL_WEBSOCKET_NODE=wss://on.<root_domain_name>:26657/websocket
      VUE_APP_DCL_CHAIN_ID=<chain-id>
      VUE_APP_DCL_ADDR_PREFIX=cosmos
      VUE_APP_DCL_SDK_VERSION=Stargate
      VUE_APP_DCL_TX_API=/rpc/tx?hash=0x
      VUE_APP_DCL_REFRESH=500000
      ```
    - add your `<root_domain_name>` with a free SSL certificate
3. Your DCL UI should be available under `https://<root_domain_name>`

## 5. Health and Monitoring

### Logs

*   Logs from all DCL nodes are collected by AWS cloudwatch agent and available at [AWS Cloudwatch](https://aws.amazon.com/cloudwatch/)
*   Logs are collected per DCL node and AWS region

### Monitoring

*   Metrics of AWS EC2 instances where the DCL nodes run on are available at [AWS Cloudwatch](https://aws.amazon.com/cloudwatch/)
*   Metrics of underlying blockchain engine (Tendermint) are pushed to [AWS AMP Service](https://aws.amazon.com/prometheus/) when prometheus is enabled
*   Use

    ```bash
    terraform output prometheus_endpoint
    ```

    to get prometheus endpoint which can be used by Grafana to visualize metrics. See detailed instructions [here][1]

[1]: https://aws.amazon.com/blogs/opensource/using-amazon-managed-service-for-prometheus-to-monitor-ec2-environments/
[2]: https://docs.aws.amazon.com/amplify/latest/userguide/getting-started.html

## 6. Destruction

**TODO**:

*   double-check no actions needed neither on business level logic nor on infra conifguration one (Ansible)

### 1. Disable validator node termination protection

```bash
cd deployment/terraform/aws
terraform apply -var="disable_validator_protection=true"
```

### 2. Destroy the infrastructure

```bash
terraform destroy
```
