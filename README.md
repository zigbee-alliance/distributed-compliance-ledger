# Distributed Compliance Ledger

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://www.apache.org/licenses/LICENSE-2.0)
[![Unit Tests](https://github.com/zigbee-alliance/distributed-compliance-ledger/workflows/Check%20Tests%20and%20Documentation/badge.svg)](https://github.com/zigbee-alliance/distributed-compliance-ledger/actions/workflows/verify.yml)

If you are interested in how to build and run the project locally, please look at [README-DEV](README-DEV.md)

Please note, that the only officially supported platform now is Linux.
It's recommended to develop and deploy the App on Ubuntu 18.04 or Ubuntu 20.04.


## Overview

DC Ledger is a public permissioned Ledger which can be used for two main use cases:

- ZB/Matter compliance certification of device models
- Public key infrastructure (PKI)

More information about use cases can be found in [DC Ledger Overview](./docs/design/DCL-Overview.pdf) and [Use Case Diagrams](docs/use_cases).

DC Ledger is based on [CometBFT](https://cometbft.com/) and [Cosmos SDK](https://cosmos.network/sdk).

DC Ledger is a public permissioned ledger in the following sense:

- Anyone can read from the ledger (that's why it's public). See [How to read from the Ledger](docs/transactions.md#how-to-read-from-the-ledger).
- Writes to the ledger are permissioned. See [How to write to the Ledger](docs/transactions.md#how-to-write-to-the-ledger) for details.
- PoA (proof of authority) approach is used for adding new validator nodes to the network
 (see [Add New Node Use Case](docs/use_cases/use_cases_add_validator_node.png)) and
  [Running Node](docs/running-node.md).

In order to send write transactions to the ledger you need:

- Have a private/public key pair
- Have an Account created on the ledger via `ACCOUNT` transaction (see [Use Case Txn Auth](docs/use_cases/use_cases_txn_auth.puml)).
  - The Account stores the public part of the key
  - The Account has an associated role. The role is used for authorization policies.
- Sign every transaction by the private key.

## Main Components

### Pool of Nodes

- A network of CometBFT-based validator nodes (Validators and Observers) maintaining the ledger.
- Every validator node (`dcld` binary) runs DC Ledger application code (based on Cosmos SDK) implementing the use cases.
- See the proposed deployment in [deployment](docs/deployment.png) and [deployment-detailed](docs/deployment-detailed.png).
- See recommended design for DCL MainNet deployment on AWS in [aws deployment](./docs/deployment-design-aws.md)

### Node Types

- **Full Node**: contains a full replication of data (ledger, state, etc.):
  - **Validator Node (VN)**: a full node participating in consensus protocol (ordering transactions).
  - **Sentry Node:** a full node that doesn't participate in consensus and wraps the validator node representing it for the rest of the network
    as one of the ways for DDoS protection.
    - **Private Sentry Node:** a full node to connect other Validator or Sentry nodes only; should not be accessed by clients.
    - **Public Sentry Node:** a full node to connect other external full nodes (possibly observer nodes).
  - **Observer Node (ON):** a full node that doesn't participate in consensus. Should be used to receive read/write requests from the clients.
- **Light Client Proxy Node**: doesn't contain a full replication of data. Can be used as a proxy to untrusted Full nodes for single-value query requests sent via CLI or CometBFT RPC.
  It will verify all state proofs automatically.
- **Seed Node**: provides a list of peers which a node can connect to.

See

- [Deployment](docs/deployment.png)
- [Deployment-detailed](docs/deployment-detailed.png).
- [Deployment Recommendations](https://github.com/zigbee-alliance/distributed-compliance-ledger/wiki/DCL-MainNet-Deployment)
- [Deployment Recommendations for AWS](./docs/deployment-design-aws.md)
- <https://docs.cometbft.com/v0.37/core/validators>
- [Run Light Client Proxy](docs/running-light-client-proxy.md)

### Clients

For interactions with the pool of nodes (sending write and read requests).

Every client must be connected to a Node (either Observer or Validator).

If there is no trusted node for connection, a [Light Client Proxy](#light-client-proxy) can be used.
A Light Client Proxy can be connected to multiple nodes and will verify the state proofs for every single value query request.

- [CLI](#cli)  
- [REST](#rest)
- [gRPC](#grpc)
- [CometBFT RPC and Light Client](#cometbft-rpc-and-light-client)

**Please note, that multi-value queries don't have state proofs support and should be sent to trusted nodes only.**

**Please make sure that TLS is enabled in gRPC, REST or Light Client Proxy for secure communication with a Node.**


## How To: Node Operators
### Add an Observer node to existing network
See [Running Node](docs/running-node.md). There are two options to add an Observer nodes:
- [Manually](docs/running-node-manual/on.md)
- [Ansible](docs/running-node-ansible/on.md)

Please take into account [running-node-in-existing-network.md](docs/advanced/running-node-in-existing-network.md).

### Add a Validator node to existing network
A recommended way for deployment and client connection: [diagram](docs/deployment.png), [diagram-detailed](docs/deployment-detailed.png) and [diagram-aws](docs/deployment-design-aws-diagram.png).

See [Running Node](docs/running-node.md) for possible patterns and instructions.

Please take into account [running-node-in-existing-network.md](docs/advanced/running-node-in-existing-network.md).

### Upgrade all nodes in a pool to a new version of DCL application

DCL application can be simultaneously updated on all nodes in the pool without breaking consensus.
See [Pool Upgrade](docs/pool-upgrade.md) and [Pool Upgrade How To](docs/pool-upgrade-how-to.md) for details.

### Run a local pool of nodes in Docker

This is for development purposes only. 

See [Run local pool](README-DEV.md#run-local-pool) section in [README-DEV.md](README-DEV.md).

## How To: Users

### CLI

- The same `dcld` binary as a Node
- A full list of all CLI commands can be found there: [transactions.md](docs/transactions.md).
- CLI can be used for write and read requests.
- Please configure the CLI before using (see [how-to.md](docs/how-to.md#cli-configuration)).
- **If there are no trusted Observer or Validator nodes to connect a CLI, then a [Light Client Proxy](#light-client-proxy) can be used.**

### Light Client Proxy

Should be used if there are no trusted Observer or Validator nodes to connect.

It can be a proxy for CLI or direct requests from code done via CometBFT RPC.

Please note, that CLI can use a Light Client proxy only for single-value query requests.
A Full Node (Validator or Observer) should be used for multi-value query requests and write requests.

Please note, that multi-value queries don't have state proofs support and should be sent to trusted node only.

See [Run Light Client Proxy](docs/running-light-client-proxy.md) for details how to run it.

### REST

- **There are no state proofs in REST, so REST queries should be sent to trusted Validator or Observer nodes only.**
- OpenAPI specification: <https://zigbee-alliance.github.io/distributed-compliance-ledger/>.
- Any running node exposes a REST API at port `1317`. See <https://docs.cosmos.network/v0.47/learn/advanced/grpc_rest>.
- See [transactions](docs/transactions.md) for a full list of endpoints.
- REST HTTP(S) queries can be directly used for read requests.
  See [How to read from the Ledger](docs/transactions.md#how-to-read-from-the-ledger).
- REST HTTP(S) queries can be directly used to broadcast generated and signed transaction.
- Generation and signing of transactions need to be done in code or via CLI.
  See [How to write to the Ledger](docs/transactions.md#how-to-write-to-the-ledger).

### gRPC

- **There are no state proofs in gRPC, so gRPC queries should be sent to trusted Validator or Observer nodes only.**
- Any running node exposes a REST API at port `9090`. See <https://docs.cosmos.network/v0.47/learn/advanced/grpc_rest>.
- A client code can be generated for all popular languages from the proto files [proto](proto), see <https://grpc.io/docs/languages/>.
- The generated client code can be used for read and write requests, i.e. generation and signing of transactions
  See [How to read from the Ledger](docs/transactions.md#how-to-read-from-the-ledger) and [How to write to the Ledger](docs/transactions.md#how-to-write-to-the-ledger) for details.

### CometBFT RPC and Light Client

- CometBFT RPC is exposed by every running node  at port `26657`. See <https://docs.cosmos.network/v0.47/learn/advanced/grpc_rest#cometbft-rpc>.
- CometBFT RPC supports state proofs. CometBFT's Light Client library can be used to verify the state proofs.
    So, if Light Client API is used, then it's possible to communicate with non-trusted nodes.
- Please note, that multi-value queries don't have state proofs support and should be sent to trusted nodes only.
- There are currently no DC Ledger specific API libraries for various platforms and languages,
but they may be provided in the future.
- The following libraries can be used as light clients:
  - [Golang Light Client implementation](https://pkg.go.dev/github.com/cometbft/cometbft/light)
  - [Rust Light Client implementation](https://docs.rs/cometbft-light-client/0.1.0-alpha.2/cometbft_light_client/)
- Refer to [this doc](./docs/cometbtf-rpc.md) to see how to [subscribe](./docs/cometbtf-rpc.md#subscribe) to a CometBFT WebSocket based events and/or [query](./docs/cometbtf-rpc.md#querying-application-components) an application components.


### Instructions

After the CLI or REST API is configured and Account with an appropriate role is created,
the following instructions from [how-to.md](docs/how-to.md) can be used for every role
(see [Use Case Diagrams](docs/use_cases)):

- [Trustee](docs/how-to.md#trustee-instructions)
  - propose new accounts
  - approve/reject new accounts
  - propose revocation of accounts
  - approve revocation of accounts
  - propose X509 root certificates
  - approve/reject X509 root certificates
  - propose revocation of X509 root certificates
  - approve revocation of X509 root certificates
  - propose pool upgrade
  - approve/reject pool upgrade
  - propose disable a validator node
  - approve/reject disable a validator node
- [Vendor](docs/how-to.md#vendor-instructions)
  - publish/update vendor info
  - publish/update/delete device model info
  - publish/update/delete device model version
  - publish/update/delete PKI Revocation Distribution Point
  - publish/remove X509 certificates
- [Certification Center](docs/how-to.md#certification-center-instructions)
  - certify or revoke certification of device models
  - update/delete compliance info
- [Vendor Admin](docs/how-to.md#vendor-admin-instructions)
  - publish/update vendor info for any vendor
- Node Admin
  - add a new Validator node
  - disable a Validator node
  - enable a Validator node


## Useful Links

- [OpenAPI specification](https://zigbee-alliance.github.io/distributed-compliance-ledger/)
- [Quick Start](docs/quickStartGuide.adoc)
- [List of Transactions, Queries, CLI command, REST API](docs/transactions.md)
- [How To Guide](docs/how-to.md)
- [Use Case Diagrams](docs/use_cases)
  - [PKI](docs/use_cases/use_cases_pki.png)
  - [Device on-ledger certification](docs/use_cases/use_cases_device_on_ledger_certification.png)
  - [Device off-ledger certification](docs/use_cases/use_cases_device_off_ledger_certification.png)
  - [Auth](docs/use_cases/use_cases_txn_auth.png)
  - [Validators](docs/use_cases/use_cases_add_validator_node.png)
  - [Pool Upgrade](docs/use_cases/use_cases_upgrade_pool.png)
- [DC Ledger Overview](docs/design/DCL-Overview.pdf)
- [DC Ledger Architecture Details](docs/design/DCL-arch-overview.pdf)
- [Deployment Pattern](docs/deployment.png)
- [Deployment Pattern Detailed](docs/deployment-detailed.png)
- [Deployment Recommendations](https://github.com/zigbee-alliance/distributed-compliance-ledger/wiki/DCL-MainNet-Deployment)
- [Deployment Recommendations for AWS](./docs/deployment-design-aws.md)
- [Running Node](docs/running-node.md)
- [Pool Upgrade](docs/pool-upgrade.md)
- [Pool Upgrade How To Guide](docs/pool-upgrade-how-to.md)
- [CometBFT](https://cometbft.com/)
- [Cosmos SDK](https://cosmos.network/sdk)
