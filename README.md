# Distributed Compliance Ledger

If you are interested in how to build and run the project locally, please look at [README-DEV](README-DEV.md)

Please note, that the only officially supported platform now is Linux.
It's recommended to develop and deploy the App on Ubuntu 18.04 or Ubuntu 20.04.

## Overview
DC Ledger is a public permissioned Ledger which can be used for two main use cases:
 - ZB/Matter compliance certification of device models
 - Public key infrastructure (PKI)
 
More information about use cases can be found in [DC Ledger Overview](docs/ZB_Ledger_overview.pdf) and [Use Case Diagrams](docs/use_cases).

DC Ledger is based on [Tendermint](https://tendermint.com/) and [Cosmos SDK](https://cosmos.network/sdk).

### Main Components

 - **Pool of Nodes**
   - A network of Tendermint-based validator nodes (Validators and Observers) maintaining the ledger.
   - Every validator node (`dcld` binary) runs DC Ledger application code (based on Cosmos SDK) implementing the use cases.
   - See the proposed deployment in [deployment](docs/deployment.png).
 - **Clients** for interactions with the pool of nodes (sending write and read requests).
     - **CLI** 
       - The same `dcld` binary as a Node
       - The CLI is based on the Cosmos SDK
       - See [CLI Usage](#cli-usage) section for details.
     - **REST**
       - Exposed by every running node.
       - See [transactions](docs/transactions.md) for a full list of endpoints.
     - **gRPC**.
       - A client code can be generated for all popular languages from the proto files [proto](proto), see https://grpc.io/docs/languages/.   
     - **Light Client**
       - Tendermint's Light Client can be used for a direct communication on API level.
       - There are currently no DC Ledger specific API libraries for various platforms and languages, 
    but they may be provided in future.
       - These libraries can be based on the following Light Client implementations: 
            - [Golang Light Client implementation](https://pkg.go.dev/github.com/tendermint/tendermint/lite2)
            - [Rust Light Client implementation](https://docs.rs/tendermint-light-client/0.23.3/tendermint_light_client/)  
 - **Public UI** (Outdated)
    - https://dcl.dev.dsr-corporation.com
    - based on the REST API
    - can be used to browse the ledger
        - please note that it doesn't show all the accounts on the ledger
        - it shows only the default (demo) accounts created on the UI server
    - **for demo purposes only**: can be used for sending write requests from the default (demo) accounts     
    - source code and documentation are located in [dcl-ui](/dcl-ui) directory

### Public Permissioned Ledger
DC Ledger is a public permissioned ledger in the following sense:
 - Anyone can read from the ledger (that's why it's public). See [How to read from the Ledger](docs/transactions.md#how-to-read-from-the-ledger).
 - Writes to the ledger are permissioned. See [How to write to the Ledger](docs/transactions.md#how-to-write-to-the-ledger) for details.
 - PoA (proof of authority) approach is used for adding new validator nodes to the network 
 (see [Add New Node Use Case](docs/use_cases/use_cases_add_validator_node.png)) and
  [Running Node Instructions](docs/running-validator-node.md).

In order to send write transactions to the ledger you need: 
   - Have a private/public key pair
   - Have an Account created on the ledger via `ACCOUNT` transaction (see [Use Case Txn Auth](use_cases/use_cases_txn_auth.puml)).
       - The Account stores the public part of the key
       - The Account has an associated role. The role is used for authorization policies.
   - Sign every transaction by the private key.


## How To

### CLI Usage
A full list of all CLI commands can be found there: [transactions.md](docs/transactions.md).

Please configure the CLI before using (see [how-to.md](docs/how-to.md#cli-configuration)).

If write requests to the Ledger needs to be sent, please make sure that you have
an Account created on the Ledger with an appropriate role (see [how-to.md](docs/how-to.md#getting-account)).

Sending read requests to the Ledger doesn't require an Account (Ledger is public for reads).

### REST Usage
Any running node exposes a REST API at port `26640`. 

A list of all REST API calls can be found in [transactions.md](docs/transactions.md).

Details on how a REST API can be used for write and read requests can be found in
[How to write to the Ledger](docs/transactions.md#how-to-write-to-the-ledger)
and [How to read from the Ledger](docs/transactions.md#how-to-read-from-the-ledger).

If write requests to the Ledger needs to be sent, please make sure that you have
an Account created on the Ledger with an appropriate role (see [how-to.md](docs/how-to.md#getting-account)).

Sending read requests to the Ledger doesn't require an Account (Ledger is public for reads).

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
- [Test House](docs/how-to.md#test-house-instructions) 
    - publish compliance test results
    - publish X509 certificates
    - revoke X509 root certificates    
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
A recommended way for deployment and client connection: [diagram](docs/deployment.png)

One can either deploy its own network of validator nodes or join one of the persistent DC Ledger Networks. 

- If you want to deploy your own network for debug purposes,
you can use the provided Ansible Playbook: [ansible/readme.md](deployment/ansible/README.md).
- If you want to join an existing network (either a custom or persistent) as a validator node,
please follow the [Running Node](docs/running-node.md) or a more detailed [Running a Validator Node](docs/advanced/running-validator-node.md) instructions.
- If you want to join an existing network  as an observer node,
please follow the [Running Node](docs/running-node.md) or a more detailed [Running an Observer Node](docs/advanced/running-observer-node.md) instructions.
- If you want to deploy your own persistent network,
you will need to create a genesis node and a genesis file first as described in
[Running Node](docs/running-node.md) or a more detailed [Running a Genesis Validator Node](docs/advanced/running-genesis-node.md).
Please note, that these instructions describe the case when the genesis block consists of a single node only. 
This is done just for simplicity, and nothing prevents you from adding more nodes to the genesis file by adapting the instructions accordingly. 

## Useful Links 
- [Quick Start](docs/quickStartGuide.adoc)
- [List of Transactions, Queries, CLI command, REST API](docs/transactions.md)
- [How To Guide](docs/how-to.md)
- [Use Case Diagrams](docs/use_cases)
    - [PKI](docs/use_cases/use_cases_pki.png)
    - [Device on-ledger certification](docs/use_cases/use_cases_device_on_ledger_certification.png)
    - [Device off-ledger certification](docs/use_cases/use_cases_device_off_ledger_certification.png)
    - [Auth](docs/use_cases/use_cases_txn_auth.png)
    - [Validators](docs/use_cases/use_cases_add_validator_node.png)
- [DC Ledger Overview](docs/design/DCL-Overview.pdf)
- [DC Ledger Architecture Details](docs/design/DCL-arch-overview.pdf)
- [Deployment Pattern](docs/deployment.png)
- [Running a Node](docs/running-node.md)
  - [Running a Genesis Validator Node](docs/advanced/running-genesis-node.md)
  - [Running a Validator Node](docs/advanced/running-validator-node.md)
  - [Running an Observer Node](docs/advanced/running-observer-node.md)
- [Tendermint](https://tendermint.com/)
- [Cosmos SDK](https://cosmos.network/sdk)
     



