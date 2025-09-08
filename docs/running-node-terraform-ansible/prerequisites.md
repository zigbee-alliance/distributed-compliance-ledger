# DCL Deployment on cloud using Terraform and Ansible (Prerequisites)

## Environment

Officially supported OS for development is `Ubuntu 20.04 LTS` and all the following intructions are tested on it.
But you are free to use any other environment that supports [Terraform][1] and [Ansible][4].

## Requirements

### 1. Clone [DCL][5] project

```bash
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

Run the following commands from the [DCL][5] project home (uses [python virtual environment][6]):

**option 1 rootless (Recommended)**

```bash
python3 -m venv dcld-venv
source dcld-venv/bin/activate
pip install -r deployment/requirements.txt
ansible-galaxy install -r deployment/galaxy-requirements.yml 
```

**option 2**

```bash
sudo pip3 install -r deployment/requirements.txt
ansible-galaxy install -r deployment/galaxy-requirements.yml 
```

### 5. Preliminary cloud configuration

<details>
<summary> AWS </summary>

In case you choose [`s3`](https://developer.hashicorp.com/terraform/language/v1.5.x/settings/backends/s3) as a terraform backend:

*   create S3 bucket
*   (optional but recommended) create DynamoDB table to support [remote state locking](https://developer.hashicorp.com/terraform/language/v1.5.x/state/locking)
    *   **Note** The table must have a partition key named `LockID` with a type of `String`.

</details>


<details>
<summary> GCP </summary>

In case you choose [`gcs`](https://developer.hashicorp.com/terraform/language/v1.5.x/settings/backends/gcs) as a terraform backend:

*   create a Cloud Storage bucket (e.g. like decribed [here](ihttps://cloud.google.com/storage/docs/creating-buckets)), recommended:
    *   enable the versioning (with limited number of concurrent versions and expiration days)
    *   prevent the public access

</details>

<details>
<summary> Azure </summary>

The Azure deployment automation logic considers the following:

*   resource group exists
*   subscription resource providers should be registered
    *   or the client should have permissions to do that (e.g. via `Contributor` role)
    *   please see more details [here](https://learn.microsoft.com/en-us/azure/azure-resource-manager/management/resource-providers-and-types)
*   in case [`azurerm`](https://developer.hashicorp.com/terraform/language/v1.5.x/settings/backends/azurerm) backend is used the following resources are needed:
    *   a storage account (see [here](https://learn.microsoft.com/en-us/azure/storage/common/storage-account-create?toc=%2Fazure%2Fstorage%2Fblobs%2Ftoc.json&bc=%2Fazure%2Fstorage%2Fblobs%2Fbreadcrumb%2Ftoc.json&tabs=azure-portal) for the details)
        *   (recommended) versioning is [enabled](https://learn.microsoft.com/en-us/azure/storage/blobs/versioning-enable?tabs=portal)
    *   a container in the storage account (see [here](https://learn.microsoft.com/en-us/azure/storage/blobs/storage-quickstart-blobs-portal#create-a-container) for the details)

</details>



[1]: https://www.terraform.io/
[2]: https://learn.hashicorp.com/tutorials/terraform/install-cli
[3]: https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html
[4]: https://www.ansible.com
[5]: https://github.com/zigbee-alliance/distributed-compliance-ledger.git
[6]: https://docs.python.org/3/library/venv.html
