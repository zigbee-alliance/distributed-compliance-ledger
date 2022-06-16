# Running a DCLedger Node

## DCL Node types

- `Genesis Validator` - Validator Node created at the beginning of a network
- `Non-genesis Validator` - Validator Node joined a network after a significant time period
- `Private Sentry` - Full Node to connect other (external) Validator Nodes ([Sentry Node Architecture](https://forum.cosmos.network/t/sentry-node-architecture-overview/454))
- `Public Sentry` - Full Node to connect other (external) Full Nodes
- `Observer` - Full Node for serving gRPC / REST / RPC clients
- `Seed` - Full Node for sharing IP addresses of `Public Sentry` Nodes ([Seed Node](https://docs.tendermint.com/master/nodes/#seed-nodes))

## DCL network architecture overview

DCL network architecture can logically have the following variations based on different use cases

### 1. Genesis Validator + (Optional) Private Sentry + (Optional) Public Sentry + (Optional) Observer + (Optional) Seed

- Running a full DCL network from scratch with Genesis Validator Node
- Refer to [deployment-design-aws.md](./deployment-design-aws.md) for more info

### 2. Validator + (Optional) Private Sentry + (Optional) Public Sentry + (Optional) Observer + (Optional) Seed

- Running a full DCL infrastructure with Validator Node to join existing DCL network
- Refer to [deployment-design-aws.md](./deployment-design-aws.md) for more info

### 3. Observer only

- Running an Observer Node only to join existing DCL network

## Running a node

Depending on your use cases you can choose one of the following options to run your nodes

### 1. Manual (advanced)

- [Genesis Validator](./running-node-manual/genesis-vn.md)
- [Non-genesis Validator](./running-node-manual/vn.md)
- [Private Sentry](./running-node-manual/private-sentry.md)
- [Public Sentry](./running-node-manual/public-sentry.md)
- [Observer](./running-node-manual/on.md)
- [Seed](./running-node-manual/seed.md)

### 2. Using ansible (semi automated)

- [Genesis Validator](./running-node-ansible/genesis-vn.md)
- [Non-genesis Validator](./running-node-ansible/vn.md)
- [Private Sentry](./running-node-ansible/private-sentry.md)
- [Public Sentry](./running-node-ansible/public-sentry.md)
- [Observer](./running-node-ansible/on.md)
- [Seed](./running-node-ansible/seed.md)

### 3. Using Terraform and Ansible on AWS cloud (fully automated)

- [AWS deployment](./running-node-aws-terraform-ansible/deployment.md) - this option includes configurable scripts to run DCL network according to [deployment-design-aws.md](./deployment-design-aws.md)

## Security and DDoS mitigation

- To protect your node against DDoS attacks you can consider one of the following options:
  - No VPN, just whitelist/blacklist via firewall rule
  - IPSec site-to-site VPN (Cloud providers)
  - [WireGuard](https://www.wireguard.com) P2P VPN
- Consider enabling TLS for public endpoints (RPC/gRPC/REST) (TLS 1.3 is recommended)
  - Most cloud providers including AWS provide TLS encryption integrated into Load Balancers
  - Offload TLS encryption using a reverse proxy (i.e [Nginx][4])
- See [deployment wiki][1] for more info

## Health and Monitoring

Health and monitoring can be configured various ways depending on a cloud provider or user needs.

Some general recommendations:

- [Prometheus][2] - for monitoring application and server performance metrics
- [ELK][3] - for collecting application and system logs
- See [deployment wiki][1] for more info

[AWS deployment](./running-node-aws-terraform-ansible/deployment.md) - contains automation scripts for setting up health and monitoring on AWS

[1]: https://github.com/zigbee-alliance/distributed-compliance-ledger/wiki/DCL-MainNet-Deployment#4-health-and-monitoring
[2]: https://prometheus.io
[3]: https://github.com/elastic
[4]: https://docs.nginx.com/nginx/admin-guide/web-server/reverse-proxy/
