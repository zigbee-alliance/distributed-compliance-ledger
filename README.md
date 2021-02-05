# Distributed Compliance Ledger

If you are interested in how to build and run the project locally, please look at [README-DEV](README-DEV.ms)

Please note, that the only officially supported platform now is Linux.
It's recommended to develop and deploy the App on Ubuntu 18.04 or Ubuntu 20.04.

## Overview
DC Ledger is a public permissioned Ledger which can be used for two main use cases:
 - ZB compliance certification of device models
 - Public key infrastructure (PKI)
 
More information about use cases can be found in [DC Ledger Overview](docs/ZB_Ledger_overview.pdf) and [Use Case Diagrams](docs/use_cases).

DC Ledger is based on [Tendermint](https://tendermint.com/) and [Cosmos SDK](https://cosmos.network/sdk).

#### Main Components
The ledger consists of
 - A network of Tendermint-based validator nodes maintaining the ledger.
  Every validator node runs DC Ledger application code (based on Cosmos SDK) implementing the use cases.
 - The client to be used for interactions with the network of nodes (sending write and read requests).
 The following clients are supported: 
    - CLI to communicate with the network of nodes (as a [light client](https://pkg.go.dev/github.com/tendermint/tendermint/lite2?tab=doc)).
    The CLI is based on the Cosmos SDK. See [CLI Usage](#cli-usage) section for details.
    - REST API which can be deployed as a server. The server can communicate with the nodes
     (as a [light client](https://pkg.go.dev/github.com/tendermint/tendermint/lite2?tab=doc)). 
    The REST API is based on the Cosmos SDK. See [REST Usage](#rest-usage) section for details.
    - Tendermint's Light Client can be used for a direct communication on API level. 
    There are currently no DC Ledger specific API libraries for various platforms and languages, 
    but they may be provided in future.
    These libraries can be based on the following Light Client implementations: 
        - [Golang Light Client implementation](https://pkg.go.dev/github.com/tendermint/tendermint/lite2?tab=doc)
        - [Rust Light Client implementation](https://docs.rs/tendermint/0.13.0/tendermint/lite/index.html)  
 - Public UI
    - https://dcl.dev.dsr-corporation.com.
    - based onf the REST API
    - can be used to browse the ledger
        - please note that it doesn't show all the accounts on the ledger
        - it shows only the default (demo) accounts created on the UI server
    - **for demo purposes only**: can be used for sending write requests from the default (demo) accounts     

#### Public Permissioned Ledger
DC Ledger is a public permissioned ledger in the following sense:
 - Anyone can read from the ledger (that's why it's public). See [How to read from the Ledger](docs/transactions.md#how-to-read-from-the-ledger).
-  Writes to the ledger are permissioned. See [How to write to the Ledger](docs/transactions.md#how-to-write-to-the-ledger) for details.
In order to send write transactions to the ledger you need: 
      - Have a private/public key pair
      - Have an Account created on the ledger via `ACCOUNT` transaction (see [Use Case Txn Auth](use_cases/use_cases_txn_auth.puml)).
          - The Account stores the public part of the key
          - The Account has an associated role. The role is used for authorization policies.
      - Sign every transaction by the private key.
 - PoA (proof of authority) approach is used for adding new validator nodes to the network 
 (see [Add New Node Use Case](docs/use_cases/use_cases_add_validator_node.png)) and
  [Running Node Instructions](docs/running-node.md).



## How To

### CLI Usage
A full list of all CLI commands can be found there: [cli-help.md](docs/cli-help.md).

Please configure the CLI before using (see [how-to.md](docs/how-to.md#cli-configuration)).

If write requests to the Ledger needs to be sent, please make sure that you have
an Account created on the Ledger with an appropriate role (see [how-to.md](docs/how-to.md#getting-account)).

Sending read requests to the Ledger doesn't require an Account (Ledger is public for reads).

### REST Usage
A REST API server is a CLI run in a REST mode: 
`dclcli rest-server --chain-id <chain_id>`.
 
Please configure the CLI before using (see [how-to.md](docs/how-to.md#cli-configuration)).

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
    - publish device model info
    - publish X509 certificates
    - revoke X509 certificates    
- [Test House](docs/how-to.md#test-house-instructions) 
    - publish compliance test results
    - publish X509 certificates
    - revoke X509 root certificates    
- [ZB Certification Center](docs/how-to.md#certification-center-instructions)
    - certify or revoke certification of device models
    - publish X509 certificates
    - revoke X509 certificates    
- [Node Admin](docs/how-to.md#node-admin-instructions-setting-up-a-new-validator-node) 
    - add a new Validator node
    - publish X509 certificates
    - revoke X509 certificates    

#### Deploy a network of validator nodes 
One can either deploy its own network of validator nodes or join one of the persistent DC Ledger Networks. 

- If you want to deploy your own network for debug purposes,
you can use the provided Ansible Playbook: [ansible/readme.md](deployment/ansible/README.md).
- If you want to join an existing network (either a custom or persistent) as a validator node,
please follow the [Running a Validator Node](docs/running-node.md) instructions.
- If you want to deploy your own persistent network,
you will need to create a genesis node and a genesis file first as described in [Running a Genesis Validator Node](docs/running-genesis-node.md).
After this more nodes can be added by following the [Running a Validator Node](docs/running-node.md) instructions.
Please note, that [Running a Genesis Validator Node](docs/running-genesis-node.md) describes 
the case when the genesis block consist of a single node only. This is done just for simplicity, 
and nothing prevents you from adding more nodes to the genesis file by adapting the instructions accordingly. 

## Useful Links 
- [Use Case Diagrams](docs/use_cases)
    - [PKI](docs/use_cases/use_cases_pki.png)
    - [Device on-ledger certification](docs/use_cases/use_cases_device_on_ledger_certification.png)
    - [Device off-ledger certification](docs/use_cases/use_cases_device_off_ledger_certification.png)
    - [Auth](docs/use_cases/use_cases_txn_auth.png)
    - [Validators](docs/use_cases/use_cases_add_validator_node.png)
- [DC Ledger Overview](docs/DCL-Overview.pdf)
- [DC Ledger Architecture Details](docs/DCL-arch-overview.pdf)
- [List of Transactions](docs/transactions.md)
- [How To Guide](docs/how-to.md)
- [CLI Help](docs/cli-help.md)
- [Deployment](deployment/ansible/README.md)
- [Running a Genesis Validator Node](docs/running-genesis-node.md)
- [Running a Validator Node](docs/running-node.md)
- [Tendermint](https://tendermint.com/)
- [Cosmos SDK](https://cosmos.network/sdk)
     



