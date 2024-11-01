# Transactions and Queries
<!-- markdownlint-disable MD036 -->

See use case sequence diagrams for the examples of how transaction can be used.

1. [General](#general)
2. [How to write to the Ledger](#how-to-write-to-the-ledger)
3. [How to read from the Ledger](#how-to-read-from-the-ledger)
4. [Module](#modules)
   * [Vendor Info](#vendor-info)
   * [Model and Model Version](#model-and-model_version)
   * [Compliance](#certify_device_compliance)
   * [X509 PKI](#x509-pki)
       - [Device Attestation (DA): PAA, PAI](#device-attestation-certificates-da-paa-pai)
       - [E2E (NOC): RCAC, ICAC](#e2e-noc-rcac-icac)
   * [Auth](#auth)
   * [Validator Node](#validator_node)
   * [Upgrade](#upgrade)
5. [Extensions](#extensions)

## General

- Every writer to the Ledger must  
  - Have a private/public key pair.
  - Have an Account created on the ledger via `ACCOUNT` transaction (see [Use Case Txn Auth](../use_cases/use_cases_txn_auth.puml)).
    - The Account stores the public part of the key
    - The Account has an associated role. The role is used for authorization policies.
  - Sign every transaction by the private key.
- Ledger is public for read which means that anyone can read from the Ledger without a need to have
an Account or sign the request.  
- The following roles are supported:
  - `Trustee` - can create and approve accounts, approve root certificates.
  - `Vendor` - can add models that belong to the vendor ID associated with the vendor account.
  - `VendorAdmin` - can add vendor info records and update any vendor info.
  - `CertificationCenter` - can certify and revoke models.
  - `NodeAdmin` - can add validator nodes to the network.

## How to write to the Ledger

- Local CLI
  - Configure the CLI before using.
      See `CLI Configuration` section in [how-to.md](../how-to.md#cli-configuration).
  - Generate and store a private key for the Account to be used for sending.
      See `Getting Account` section in [how-to.md](../how-to.md#getting-account).
  - Send transactions to the ledger from the Account (`--from`).
    - it will automatically build a request, sign it by the account's key, and broadcast to the ledger.
  - See `CLI` sub-sections for every write request (transaction).
  - It's possible to build, sign and broadcast a transaction separately (possibly by diferent CLIs):
    - Let's assume we have two CLIs:
      - CLI 1: Is connected to the network of nodes. Doesn't have access to private keys.
      - CLI 2: Stores private key. Does not have a connection to the network of nodes.
    - Build transaction by CLI 1: `dcld tx ... --generate-only`
    - Fetch `account-number` and `sequence` by CLI 1:  `dcld query auth account --address <address>`
    - Sign transaction by CLI 2: `dcld tx sign txn.json --from <from> --account-number <int> --sequence <int> --gas "auto" --offline --output-document txn.json`
    - Broadcast transaction by CLI 1: `dcld tx broadcast txn.json`
    - To get the actual result of transaction, `dcld query tx=txHash` call must be executed, where `txHash` is the hash of previously executed transaction.
- gRPC:
  - Generate a client code from the proto files [proto](../../proto) for the client language (see <https://grpc.io/docs/languages/>)
  - Build, sign, and broadcast the message (transaction).
      See [grpc/rest integration tests](../../integration_tests/grpc_rest) as an example.
- REST API
  - Build and sign a transaction by one of the following ways
    - In code via gRPC (see above)
    - Via CLI commands specifying `--generate-only` flag and using `dcld tx sign` (see above)
  - The user does a `POST` of the signed request to `http://<node-ip>:1317/cosmos/tx/v1beta1/txs` endpoint.
  - Example

```bash
dcld tx ... --generate-only
dcld query auth account --address <address>
dcld tx sign txn.json --from <from> --account-number <int> --sequence <int> --gas "auto" --offline --output-document txn.json
POST http://<node-ip>:1317/cosmos/tx/v1beta1/txs
```

## How to read from the Ledger

No keys/account is needed as the ledger is public for reads.

Please note, that multi-value queries don't have state proofs support and should be sent to trusted nodes only.

Please make sure that TLS is enabled in gRPC, REST or Light Client Proxy for secure communication with a Node.

- Local CLI
  - See `CLI` section for every read request.
  - If there are no trusted Observer or Validator nodes to connect a CLI, then a [Light Client Proxy](../running-light-client-proxy.md) can be used.
- REST API
  - OpenAPI specification: <https://zigbee-alliance.github.io/distributed-compliance-ledger/>.
  - Any running node exposes a REST API at port `1317`. See <https://docs.cosmos.network/v0.45/core/grpc_rest.html>.
  - See `REST API` section for every read request.
  - See [grpc/rest integration tests](../../integration_tests/grpc_rest) as an example.
  - There are no state proofs in REST, so REST queries should be sent to trusted Validator or Observer nodes only.
- gRPC
  - Any running node exposes a REST API at port `9090`. See <https://docs.cosmos.network/v0.45/core/grpc_rest.html>.
  - Generate a client code from the proto files [proto](../../proto) for the client language (see <https://grpc.io/docs/languages/>).
  - See [grpc/rest integration tests](../../integration_tests/grpc_rest) as an example.
  - There are no state proofs in gRPC, so gRPC queries should be sent to trusted Validator or Observer nodes only.
- Tendermint RPC
  - Tendermint RPC OpenAPI specification can be found in <https://zigbee-alliance.github.io/distributed-compliance-ledger/>.
  - Tendermint RPC is exposed by every running node  at port `26657`. See <https://docs.cosmos.network/v0.45/core/grpc_rest.html#tendermint-rpc>.
  - Tendermint RPC supports state proofs. Tendermint's Light Client library can be used to verify the state proofs.
    So, if Light Client API is used, then it's possible to communicate with non-trusted nodes.
  - Please note, that multi-value queries don't have state proofs support and should be sent to trusted nodes only.
  - Refer to [this doc](../cometbft-rpc.md) to see how to [subscribe](../cometbft-rpc.md#subscribe) to a Tendermint WebSocket based events and/or [query](../cometbft-rpc.md#querying-application-components) an application components. 

`NotFound` (404 code) is returned if an entry is not found on the ledger.

### Query types

- Query single value
- Query list of values with pagination support (should be sent to trusted nodes only)

### Common pagination parameters

- count-total `optional(bool)`:  count total number of records
- limit `optional(uint)`:        pagination limit (default 100)
- offset `optional(uint)`:       pagination offset
- page `optional(uint)`:         pagination page. This sets offset to a multiple of limit (default 1).
- page-key `optional(string)`:   pagination page-key
- reverse `optional(bool)`:       results are sorted in descending order

## Modules

### [VENDOR INFO](./vendor-info.md)

### [MODEL and MODEL_VERSION](./model.md)

### [CERTIFY_DEVICE_COMPLIANCE](./compliance.md)

### [X509 PKI](./pki.md)

### [AUTH](./auth.md)

### [VALIDATOR_NODE](./validator-node.md)

### [UPGRADE](./upgrade.md)

## Extensions

### Sign

Sign transaction by the given key.

- Parameters:
  - `txn` - transaction to sign.
  - `from` -  name or address of private key to use to sign.
  - `account-number` - (optional) the account number of the signing account.
  - `sequence` - (optional) the sequence number of the signing account.
  - `chain-id` - (optional) chain ID.
- CLI command:
  - `dcld tx sign [path-to-txn-file] --from [address]`

> **_Note:_**  if `account_number` and `sequence`  are not specified they will be fetched from the ledger automatically.  

### Broadcast

Broadcast transaction to the ledger.

- Parameters:
  - `txn` - transaction to broadcast
- CLI command:
  - `dcld tx broadcast [path-to-txn-file]`
- REST API:
  - POST `/cosmos/tx/v1beta1/txs`
  
### Status

Query status of a node.

- Parameters:
  - `node`: optional(string) - node physical address to query (by default queries the node specified in CLI config file or else "tcp://localhost:26657")
- CLI command:
  - `dcld status [--node=<node ip>]`
- REST API:
  - GET `/cosmos/base/tendermint/v1beta1/node_info`

### Validator set

Get the list of tendermint validators participating in the consensus at given height.

- Parameters:
  - `height`: optional(uint) - height to query (the latest by default)
- CLI command:
  - `dcld query tendermint-validator-set [height]`
- REST API:
  - GET `/cosmos/base/tendermint/v1beta1/validatorsets/latest`
  - GET `/cosmos/base/tendermint/v1beta1/validatorsets/{height}`

### Keys

The set of CLI commands that allows you to manage your local keystore.

Commands:

- Derive a new private key and encrypt to disk.

  You will be prompted to create an encryption passphrase.
  This passphrase will be requested each time you send write transactions on the ledger using this key.
  You can remember and securely save the mnemonic phrase shown after the key is created
  to be able to recover the key later.

  Command: `dcld keys add <key name>`

  Example: `dcld keys add jack`

- Recover existing key instead of creating a new one.

  The key can be recovered from a seed obtained from the mnemonic passphrase (see the previous command).
  You will be prompted to create an encryption passphrase and enter the seed's mnemonic.
  This passphrase will be requested each time you send write transactions on the ledger using this key.

  Command: `dcld keys add <key name> --recover`

  Example: `dcld keys add jack --recover`

- Get a list of all stored public keys.

  Command: `dcld keys list`

  Example: `dcld keys list`

- Get details for a key.

  Command: `dcld keys show <key name>`

  Example: `dcld keys show jack`

- Export a key.

  A private key from the local keystore can be exported in ASCII-armored encrypted format.
  You will be prompted to enter the decryption passphrase for the key and  
  to create an encryption passphrase for the exported key.
  The exported key can be stored to a file for import.

  Command: `dcld keys export <key name>`
  
  Example: `dcld keys export jack`
  
- Import a key.

  A key can be imported from the ASCII-armored encrypted format
  obtained by the export key command.
  You will be prompted to enter the decryption passphrase for the exported key
  which was used during the export process.

  Command: `dcld keys import <key name> <key file>`
  
  Example: `dcld keys import jack jack_exported_priv_key_file`
<!-- markdownlint-enable MD036 -->
