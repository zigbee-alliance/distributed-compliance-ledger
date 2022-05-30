# Running a DCLedger Node using Ansible (Prerequisites)

## Requirements (local machine)

### 1. Clone [DCL][2] project
```bash
git clone https://github.com/zigbee-alliance/distributed-compliance-ledger.git
```

### 2. Install `Python` and `Pip`
```bash
sudo apt-get update
sudo apt-get install -y --no-install-recommends python3
sudo apt install python3-pip python3-testresources
```

### 3. Install [Ansible][1] and its dependencies
    
Run the following commands from the [DCL][2] project home
```bash
sudo pip install -r deployment/requirements.txt
ansible-galaxy install -r deployment/galaxy-requirements.yml
```

## Requirements (target machine)

### Hardware requirements

Minimal:

* 1GB RAM
* 25GB of disk space
* 1.4 GHz CPU

Recommended (for highload applications):

* 2GB RAM
* 100GB SSD
* x64 2.0 GHz 2v CPU

### Operating System

Current delivery is compiled and tested under `Ubuntu 20.04 LTS` so we recommend using this distribution for now.
In future, it will be possible to compile the application for a wide range of operating systems (thanks to Go language).

### Install `Python`
Python3 needs to be installed on target machine to run ansible playbooks
```bash
sudo apt-get update
sudo apt-get install -y --no-install-recommends python3
```

[1]: https://www.ansible.com
[2]: https://github.com/zigbee-alliance/distributed-compliance-ledger.git