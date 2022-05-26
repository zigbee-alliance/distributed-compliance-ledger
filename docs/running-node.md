# Running a DCLedger Node

## DCL network architecture overview
DCL network architecture can logically have the following variations based on different use cases
#### 1. Genesis Validator + (Optional) Private Sentry + (Optional) Public Sentry + (Optional) Observer + (Optional) Seed
- Running a full DCL network from scratch with Genesis Validator Node
- Refer to [deployment-design-aws.md](./deployment-design-aws.md) for more info

#### 2. Validator + (Optional) Private Sentry + (Optional) Public Sentry + (Optional) Observer + (Optional) Seed
- Running a full DCL infrastructure with Validator Node to join existing DCL network
- Refer to [deployment-design-aws.md](./deployment-design-aws.md) for more info

#### 3. Observer only
- Running an Observer Node only to join existing DCL network


## Running a node
Depending on your use cases you can choose one of the following options to run your nodes 
### 1. Manual (advanced)
- [Genesis validator](./running-node-manual/genesis-vn.md)
- [Validator](./running-node-manual/vn.md) 
- [Private Sentry](./running-node-manual/private-sentry.md)
- [Public Sentry](./running-node-manual/public-sentry.md) 
- [Observer](./running-node-manual/on.md)
- [Seed](./running-node-manual/seed.md)

### 2. Using ansible (semi automated)
- [Genesis validator](./running-node-ansible/genesis-vn.md)
- [Validator](./running-node-ansible/vn.md) 
- [Private Sentry](./running-node-ansible/private-sentry.md) 
- [Public Sentry](./running-node-ansible/public-sentry.md) 
- [Observer](./running-node-ansible/on.md)
- [Seed](./running-node-ansible/seed.md) 

### 3. Using Terraform and Ansible on AWS cloud (fully automated)
- [AWS deployment](./running-node-aws-terraform-ansible/deployment.md)