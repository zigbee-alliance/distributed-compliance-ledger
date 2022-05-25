# Distributed Compliance Ledger

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://www.apache.org/licenses/LICENSE-2.0)
[![Unit Tests](https://github.com/zigbee-alliance/distributed-compliance-ledger/workflows/Check%20Tests%20and%20Documentation/badge.svg)](https://github.com/zigbee-alliance/distributed-compliance-ledger/actions/workflows/verify.yml)

If you are interested in how to build and run the project locally, please look at [README-DEV](README-DEV.md)

Please note, that the only officially supported platform now is Linux.
It's recommended to develop and deploy the App on Ubuntu 18.04 or Ubuntu 20.04.

**Please note, that there were breaking changes in DCL 0.6 (migration to the latest Cosmos SDK), so
the current master and DCL releases 0.6+ are not compatible with pools and Test Nets running DCL 0.5.**

## Overview

DC Ledger is a public permissioned Ledger which can be used for two main use cases:

- ZB/Matter compliance certification of device models
- Public key infrastructure (PKI)

More information about use cases can be found in [DC Ledger Overview](./docs/design/DCL-Overview.pdf) and [Use Case Diagrams](docs/use_cases).

DC Ledger is based on [Tendermint](https://tendermint.com/) and [Cosmos SDK](https://cosmos.network/sdk).

DC Ledger is a public permissioned ledger in the following sense:

- Anyone can read from the ledger (that's why it's public). See [How to read from the Ledger](docs/transactions.md#how-to-read-from-the-ledger).
- Writes to the ledger are permissioned. See [How to write to the Ledger](docs/transactions.md#how-to-write-to-the-ledger) for details.
- PoA (proof of authority) approach is used for adding new validator nodes to the network
 (see [Add New Node Use Case](docs/use_cases/use_cases_add_validator_node.png)) and
  [Running Node Instructions](docs/advanced/running-validator-node.md).

In order to send write transactions to the ledger you need:

- Have a private/public key pair
- Have an Account created on the ledger via `ACCOUNT` transaction (see [Use Case Txn Auth](docs/use_cases/use_cases_txn_auth.puml)).
  - The Account stores the public part of the key
  - The Account has an associated role. The role is used for authorization policies.
- Sign every transaction by the private key.

## Main Components

### Pool of Nodes

- A network of Tendermint-based validator nodes (Validators and Observers) maintaining the ledger.
- Every validator node (`dcld` binary) runs DC Ledger application code (based on Cosmos SDK) implementing the use cases.
- See the proposed deployment in [deployment](docs/deployment.png) and [deployment-detailed](docs/deployment-detailed.png).
- See recommended design for DCL MainNet deployment on AWS in [aws deployment](./docs/deployment-design-aws.md)

### Node Types

- **Full Node**: contains a full replication of data (ledger, state, etc.):
  - **Validator Node (VN)**: a full node participating in consensus protocol (ordering transactions).
  - **Sentry Node:** a full nodes that doesn't participate in consensus and wraps the validator node representing it for the rest of the network
    as one of the ways for DDoS protection.
    - **Private Sentry Node:** a full node to connect other Validator or Sentry nodes only; should not be accessed by clients.
    - **Public Sentry Node:** a full node to connect other external full nodes (possibly observer nodes).
  - **Observer Node (ON):** a full node that doesn't participate in consensus. Should be used to receive read/write requests from the clients.
- **Light Client Proxy Node**: doesn't contain a full replication of data. Can be used as a proxy to untrusted Full nodes for single-value query requests sent via CLI or Tendermint RPC.
  It will verify all state proofs automatically.
- **Seed Node**: provides a list of peers which a node can connect to.

See

- [Deployment](docs/deployment.png)
- [Deployment-detailed](docs/deployment-detailed.png).
- [Deployment Recommendations](https://github.com/zigbee-alliance/distributed-compliance-ledger/wiki/DCL-MainNet-Deployment)
- [Deployment Recommendations for AWS](./docs/deployment-design-aws.md)
- <https://docs.tendermint.com/master/nodes/validators.html>
- [Run Light Client Proxy](docs/running-light-client-proxy.md)

### Clients

For interactions with the pool of nodes (sending write and read requests).

Every client must be connected to a Node (either Observer or Validator).

If there is no trusted node for connection, a [Light Client Proxy](#light-client-proxy) can be used.
A Light Client Proxy can be connected to multiple nodes and will verify the state proofs for every single value query request.

- [CLI](#cli)  
- [REST](#rest)
- [gRPC](#grpc)
- [Tendermint RPC and Light Client](#tendermint-rpc-and-light-client)

**Please note, that multi-value queries don't have state proofs support and should be sent to trusted nodes only.**

**Please make sure that TLS is enabled in gRPC, REST or Light Client Proxy for secure communication with a Node.**

### Public UI (Outdated, doesn't work with DCL 0.6+)

- based on the REST API
- can be used to browse the ledger
  - please note that it doesn't show all the accounts on the ledger
  - it shows only the default (demo) accounts created on the UI server
- **for demo purposes only**: can be used for sending write requests from the default (demo) accounts
- source code and documentation are located in [dcl-ui](dcl-ui) directory

## How To

### CLI

- The same `dcld` binary as a Node
- A full list of all CLI commands can be found there: [transactions.md](docs/transactions.md).
- CLI can be used for write and read requests.
- Please configure the CLI before using (see [how-to.md](docs/how-to.md#cli-configuration)).
- **If there are no trusted Observer or Validator nodes to connect a CLI, then a [Light Client Proxy](#light-client-proxy) can be used.**

### Light Client Proxy

Should be used if there are no trusted Observer or Validator nodes to connect.

It can be a proxy for CLI or direct requests from code done via Tendermint RPC.

Please note, that CLI can use a Light Client proxy only for single-value query requests.
A Full Node (Validator or Observer) should be used for multi-value query requests and write requests.

Please note, that multi-value queries don't have state proofs support and should be sent to trusted node only.

See [Run Light Client Proxy](docs/running-light-client-proxy.md) for details how to run it.

### REST

- **There are no state proofs in REST, so REST queries should be sent to trusted Validator or Observer nodes only.**
- OpenAPI specification: <https://zigbee-alliance.github.io/distributed-compliance-ledger/>.
- Any running node exposes a REST API at port `1317`. See <https://docs.cosmos.network/v0.44/core/grpc_rest.html>.
- See [transactions](docs/transactions.md) for a full list of endpoints.
- REST HTTP(S) queries can be directly used for read requests.
  See [How to read from the Ledger](docs/transactions.md#how-to-read-from-the-ledger).
- REST HTTP(S) queries can be directly used to broadcast generated and signed transaction.
- Generation and signing of transactions need to be done in code or via CLI.
  See [How to write to the Ledger](docs/transactions.md#how-to-write-to-the-ledger).

### gRPC

- **There are no state proofs in gRPC, so gRPC queries should be sent to trusted Validator or Observer nodes only.**
- Any running node exposes a REST API at port `9090`. See <https://docs.cosmos.network/v0.44/core/grpc_rest.html>.
- A client code can be generated for all popular languages from the proto files [proto](proto), see <https://grpc.io/docs/languages/>.
- The generated client code can be used for read and write requests, i.e. generation and signing of transactions
  See [How to read from the Ledger](docs/transactions.md#how-to-read-from-the-ledger) and [How to write to the Ledger](docs/transactions.md#how-to-write-to-the-ledger) for details.

### Tendermint RPC and Light Client

- Tendermint RPC is exposed by every running node  at port `26657`. See <https://docs.cosmos.network/v0.44/core/grpc_rest.html#tendermint-rpc>.
- Tendermint RPC supports state proofs. Tendermint's Light Client library can be used to verify the state proofs.
    So, if Light Client API is used, then it's possible to communicate with non-trusted nodes.
- Please note, that multi-value queries don't have state proofs support and should be sent to trusted nodes only.
- There are currently no DC Ledger specific API libraries for various platforms and languages,
but they may be provided in the future.
- The following libraries can be used as light clients:
  - [Golang Light Client implementation](https://pkg.go.dev/github.com/tendermint/tendermint/lite2)
  - [Rust Light Client implementation](https://docs.rs/tendermint-light-client/0.23.3/tendermint_light_client/)

### Instructions

After the CLI or REST API is configured and Account with an appropriate role is created,
the following instructions from [how-to.md](docs/how-to.md) can be used for every role
(see [Use Case Diagrams](docs/use_cases)):

- [Trustee](docs/how-to.md#trustee-instructions)
  - propose new accounts
  - approve new accounts
  - propose revocation of accounts
  - approve revocation of accounts
  - propose X509 root certificates
  - approve X509 root certificates
  - propose revocation of X509 root certificates
  - approve revocation of X509 root certificates
  - publish X509 certificates
  - revoke X509 certificates
  - propose pool upgrade
  - approve pool upgrade

- [CA](docs/how-to.md#ca-instructions)
  - propose X509 root certificates
  - publish X509 certificates
  - revoke X509 certificates
- [Vendor](docs/how-to.md#vendor-instructions)
  - publish vendor info
  - publish device model info
  - publish device model version
  - publish X509 certificates
  - revoke X509 certificates
- [Certification Center](docs/how-to.md#certification-center-instructions)
  - certify or revoke certification of device models
  - publish X509 certificates
  - revoke X509 certificates
- [Node Admin](docs/how-to.md#node-admin-instructions-setting-up-a-new-validator-node)
  - add a new Validator node
  - publish X509 certificates
  - revoke X509 certificates

### Run a local pool of nodes in Docker

See [Run local pool](README-DEV.md#run-local-pool) section in [README-DEV.md](README-DEV.md).

### Deploy a persistent pool of nodes

A recommended way for deployment and client connection: [diagram](docs/deployment.png), [diagram-detailed](docs/deployment-detailed.png) and [diagram-aws](docs/deployment-design-aws-diagram.png).

One can either deploy its own network of validator nodes or join one of the persistent DC Ledger Networks.

- If you want to deploy your own network for debug purposes,
you can use the provided Ansible Playbook: [ansible/readme.md](deployment/ansible/README.md).
- If you want to join an existing network (either a custom or persistent) as a validator node,
please follow the [Running Node](docs/running-node.md) or a more detailed [Running a Validator Node](docs/advanced/running-validator-node.md) instructions.
  > **_NOTE:_** If you are joining a `long-running` network,
  consider the following additional instructions [Running Node in existing network](docs/running-node-in-existing-network.md)
- If you want to join an existing new network as an observer node,
please follow the [Running Node](docs/running-node.md) or a more detailed [Running an Observer Node](docs/advanced/running-observer-node.md) instructions.
  > **_NOTE:_** If you are joining a `long-running` network,
  consider the following additional instructions [Running Node in existing network](docs/running-node-in-existing-network.md)
- If you want to deploy your own persistent network,
you will need to create a genesis node and a genesis file first as described in
[Running Node](docs/running-node.md) or a more detailed [Running a Genesis Validator Node](docs/advanced/running-genesis-node.md).
Please note, that these instructions describe the case when the genesis block consists of a single node only.
This is done just for simplicity, and nothing prevents you from adding more nodes to the genesis file by adapting the instructions accordingly.

### Upgrade all nodes in a pool to a new version of DCL application

DCL application can be simultaneously updated on all nodes in the pool without breaking consensus.
See [Pool Upgrade](docs/pool-upgrade.md) and [Pool Upgrade How To](docs/pool-upgrade-how-to.md) for details.

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
- [Running Node in a new network](docs/running-node.md)
  - [Running Genesis Validator Node](docs/advanced/running-genesis-node.md)
  - [Running Validator Node](docs/advanced/running-validator-node.md)
  - [Running Observer Node](docs/advanced/running-observer-node.md)
- [Running Node in an existing network](docs/running-node-in-existing-network.md)
- [Pool Upgrade](docs/pool-upgrade.md)
- [Pool Upgrade How To Guide](docs/pool-upgrade-how-to.md)
- [Tendermint](https://tendermint.com/)
- [Cosmos SDK](https://cosmos.network/sdk)
