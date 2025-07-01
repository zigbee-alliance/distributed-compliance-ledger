# DCL Deployment on AWS using Terraform and Ansible (Prerequisites)

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


[1]: https://www.terraform.io/
[2]: https://learn.hashicorp.com/tutorials/terraform/install-cli
[3]: https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html
[4]: https://www.ansible.com
[5]: https://github.com/zigbee-alliance/distributed-compliance-ledger.git
[6]: https://docs.python.org/3/library/venv.html
