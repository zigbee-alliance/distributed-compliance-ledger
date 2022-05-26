# DCL Deployment on AWS using Terraform and Ansible (Prerequisites)

## Overview

This document describes all necessary steps to deploy a new DCL network on AWS cloud in accordance with this [design document](../deployment-design-aws.md).

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