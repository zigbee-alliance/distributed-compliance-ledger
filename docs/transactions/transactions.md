# Transactions and Queries
<!-- markdownlint-disable MD036 -->

See use case sequence diagrams for the examples of how transaction can be used.

1. [General](#general)
2. [How to write to the Ledger](#how-to-write-to-the-ledger)
3. [How to read from the Ledger](#how-to-read-from-the-ledger)
4. [Modules](#modules)
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
    - Build transaction by CLI 1: `dcld tx --generate-only`
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
dcld tx --generate-only
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

| Method                                                      | Description                                        | CLI command                         |
|-------------------------------------------------------------|----------------------------------------------------|-------------------------------------|
| [ADD_VENDOR_INFO](./vendor-info.md#add_vendor_info)         | Adds a record about a Vendor                       | `dcld tx vendorinfo add-vendor`     |
| [UPDATE_VENDOR_INFO](./vendor-info.md#update_vendor_info)   | Updates a record about a Vendor                    | `dcld tx vendorinfo update-vendor`  |
| [GET_VENDOR_INFO](./vendor-info.md#get_vendor_info)         | Gets a Vendor Info for the given `vid` (vendor ID) | `dcld query vendorinfo vendor`      |
| [GET_ALL_VENDOR_INFO](./vendor-info.md#get_all_vendor_info) | Gets information about all vendors for all VIDs    | `dcld query vendorinfo all-vendors` |

### [MODEL and MODEL_VERSION](./model.md)

| Method                                                      | Description                                      | CLI command                           |
|-------------------------------------------------------------|--------------------------------------------------|---------------------------------------|
| [ADD_MODEL](./model.md#add_model)                           | Adds a new Model                                 | `dcld tx model add-model`             |
| [EDIT_MODEL](./model.md#edit_model)                         | Edits an existing Model                          | `dcld tx model update-model`          |
| [DELETE_MODEL](./model.md#delete_model)                     | Deletes an existing Model                        | `dcld tx model delete-model`          |
| [ADD_MODEL_VERSION](./model.md#add_model_version)           | Adds a new Model Software Version                | `dcld tx model add-model-version`     |
| [EDIT_MODEL_VERSION](./model.md#edit_model_version)         | Edits an existing Model Software Version         | `dcld tx model update-model-version`  |
| [DELETE_MODEL_VERSION](./model.md#delete_model_version)     | Deletes an existing Model Version                | `dcld tx model delete-model-version`  |
| [GET_MODEL](./model.md#get_model)                           | Gets a Model Info                                | `dcld query model get-model`          |
| [GET_MODEL_VERSION](./model.md#get_model_version)           | Gets a Model Software Versions                   | `dcld query model model-version`      |
| [GET_ALL_MODELS](./model.md#get_all_models)                 | Gets all Model Infos for all vendors             | `dcld query model all-models`         |
| [GET_ALL_VENDOR_MODELS](./model.md#get_all_vendor_models)   | Gets all Model Infos by the given Vendor         | `dcld query model vendor-models`      |
| [GET_ALL_MODEL_VERSIONS](./model.md#get_all_model_versions) | Gets all Model Software Versions for vid and pid | `dcld query model all-model-versions` |

### [CERTIFY_DEVICE_COMPLIANCE](./compliance.md)

| Method                                                                                     | Description                                    | CLI command                                            |
|--------------------------------------------------------------------------------------------|------------------------------------------------|--------------------------------------------------------|
| [CERTIFY_MODEL](./compliance.md#certify_model)                                             | Attests compliance of the Model Version        | `dcld tx compliance certify-model`                     |
| [UPDATE_COMPLIANCE_INFO](./compliance.md#update_compliance_info)                           | Updates a compliance info                      | `dcld tx compliance update-compliance-info`            |
| [DELETE_COMPLIANCE_INFO](./compliance.md#delete_compliance_info)                           | Delete compliance of the Model Version         | `dcld tx compliance delete-compliance-info`            |
| [REVOKE_MODEL_CERTIFICATION](./compliance.md#revoke_model_certification)                   | Revoke compliance of the Model Version         | `dcld tx compliance revoke-model`                      |
| [PROVISION_MODEL](./compliance.md#provision_model)                                         | Sets provisional state for the Model Version   | `dcld tx compliance provision-model`                   |
| [GET_CERTIFIED_MODEL](./compliance.md#get_certified_model)                                 | Gets Model compliance information              | `dcld query compliance certified-model`                |
| [GET_REVOKED_MODEL](./compliance.md#get_revoked_model)                                     | Gets Model revocation information              | `dcld query compliance revoked-model`                  |
| [GET_PROVISIONAL_MODEL](./compliance.md#get_provisional_model)                             | Gets Model Version provisional information     | `dcld query compliance provisional-model`              |
| [GET_COMPLIANCE_INFO](./compliance.md#get_compliance_info)                                 | Gets Model Version compliance information      | `dcld query compliance compliance-info`                |
| [GET_DEVICE_SOFTWARE_COMPLIANCE](./compliance.md#get_device_software_compliance)           | Gets device software compliance                | `dcld query compliance device-software-compliance`     |
| [GET_ALL_CERTIFIED_MODELS](./compliance.md#get_all_certified_models)                       | Gets all compliant Model Versions              | `dcld query compliance all-certified-models`           |
| [GET_ALL_REVOKED_MODELS](./compliance.md#get_all_revoked_models)                           | Gets all revoked Model Versions                | `dcld query compliance all-revoked-models`             |
| [GET_ALL_PROVISIONAL_MODELS](./compliance.md#get_all_provisional_models)                   | Gets all Model Versions in provisional state   | `dcld query compliance all-provisional-models`         |
| [GET_ALL_COMPLIANCE_INFO_RECORDS](./compliance.md#get_all_compliance_info_records)         | Gets all stored compliance information records | `dcld query compliance all-compliance-info`            |
| [GET_ALL_DEVICE_SOFTWARE_COMPLIANCES](./compliance.md#get_all_device_software_compliances) | Gets all stored device software compliance's   | `dcld query compliance all-device-software-compliance` |

### [X509 PKI](./pki.md)

| Method                                                                                                                        | Description                                                                     | CLI command                                             |
|-------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------|---------------------------------------------------------|
| Global                                                                                                                        | Work for all certificate types (DA, NOC)                                        |                                                         |
| [GET_CERT](./pki.md#get_cert)                                                                                                 | Gets a certificate (PAA, PAI, RCAC, ICAC)                                       | `dcld query pki cert`                                   |
| [GET_ALL_CERTS](./pki.md#get_all_certs)                                                                                       | Gets all certificates (PAA, PAI, RCAC, ICAC)                                    | `dcld query pki all-certs`                              |
| [GET_CHILD_CERTS](./pki.md#get_child_certs)                                                                                   | ets all child certificates for the given certificate                            | `dcld query pki all-child-x509-certs`                   |
| DA                                                                                                                            | Work for DA certificate types (PAA, PAI)                                        |                                                         |
| [PROPOSE_ADD_PAA](./pki.md#propose_add_paa)                                                                                   | Proposes a new PAA (self-signed root certificate)                               | `dcld tx pki propose-add-x509-root-cert`                |
| [APPROVE_ADD_PAA](./pki.md#approve_add_paa)                                                                                   | Approves the proposed PAA / Re-vote                                             | `dcld tx pki approve-add-x509-root-cert`                |
| [REJECT_ADD_PAA](./pki.md#reject_add_paa)                                                                                     | Rejects the proposed PAA / Re-vote                                              | `dcld tx pki reject-add-x509-root-cert`                 |
| [PROPOSE_REVOKE_PAA](./pki.md#propose_revoke_paa)                                                                             | Proposes revocation of the given PAA                                            | `dcld tx pki propose-revoke-x509-root-cert`             |
| [APPROVE_REVOKE_PAA](./pki.md#approve_revoke_paa)                                                                             | Approves the revocation of the given PAA                                        | `dcld tx pki approve-revoke-x509-root-cert`             |
| [ASSIGN_VID_TO_PAA](./pki.md#assign_vid_to_paa)                                                                               | Assigns a Vendor ID to non-VID scoped PAAs                                      | `dcld tx pki assign-vid`                                |
| [ADD_REVOCATION_DISTRIBUTION_POINT](./pki.md#add_revocation_distribution_point)                                               | Publishes a PKI Revocation distribution endpoint                                | `dcld tx pki add-revocation-point`                      |
| [UPDATE_REVOCATION_DISTRIBUTION_POINT](./pki.md#update_revocation_distribution_point)                                         | Updates an existing PKI Revocation distribution endpoint                        | `dcld tx pki update-revocation-point`                   |
| [DELETE_REVOCATION_DISTRIBUTION_POINT](./pki.md#delete_revocation_distribution_point)                                         | Deletes a PKI Revocation distribution endpoint                                  | `dcld tx pki delete-revocation-point`                   |
| [ADD_PAI](./pki.md#add_pai)                                                                                                   | Adds a PAI (intermediate certificate)                                           | `dcld tx pki add-x509-cert`                             |
| [REVOKE_PAI](./pki.md#revoke_pai)                                                                                             | Revokes the given PAI (intermediate certificate)                                | `dcld tx pki revoke-x509-cert`                          |
| [REMOVE_PAI](./pki.md#remove_pai)                                                                                             | Removes the given PAI from approved and revoked lists                           | `dcld tx pki remove-x509-cert`                          |
| [GET_DA_CERT](./pki.md#get_da_cert)                                                                                           | Gets a DA certificate (PAA, PAI)                                                | `dcld query pki x509-cert`                              |
| [GET_REVOKED_DA_CERT](./pki.md#get_revoked_da_cert)                                                                           | Gets a revoked DA certificate (PAA, PAI)                                        | `dcld query pki revoked-x509-cert`                      |
| [GET_DA_CERTS_BY_SKID](./pki.md#get_da_certs_by_skid)                                                                         | Gets all DA certificates by the given subject key ID                            | `dcld query pki x509-cert`                              |
| [GET_DA_CERTS_BY_SUBJECT](./pki.md#get_da_certs_by_subject)                                                                   | Gets all DA certificates associated with a subject                              | `dcld query pki all-subject-x509-certs`                 |
| [GET_ALL_DA_CERTS](./pki.md#get_all_da_certs)                                                                                 | Gets all DA certificates                                                        | `dcld query pki all-x509-certs`                         |
| [GET_ALL_REVOKED_DA_CERTS](./pki.md#get_all_revoked_da_certs)                                                                 | Gets all revoked DA certificates                                                | `dcld query pki all-revoked-x509-certs`                 |
| [GET_PKI_REVOCATION_DISTRIBUTION_POINT](./pki.md#get_pki_revocation_distribution_point)                                       | Gets a revocation distribution point                                            | `dcld query pki revocation-point`                       |
| [GET_PKI_REVOCATION_DISTRIBUTION_POINTS_BY_SUBJECT_KEY_ID](./pki.md#get_pki_revocation_distribution_points_by_subject_key_id) | Gets a list of revocation distribution point by subject key id                  | `dcld query pki revocation-points`                      |
| [GET_ALL_PKI_REVOCATION_DISTRIBUTION_POINT](./pki.md#get_all_pki_revocation_distribution_point)                               | Gets a list of all revocation distribution points                               | `dcld query pki all-revocation-points`                  |
| [GET_PROPOSED_PAA](./pki.md#get_proposed_paa)                                                                                 | Gets a proposed but not approved PAA certificate                                | `dcld query pki proposed-x509-root-cert`                |
| [GET_REJECTED_PAA](./pki.md#get_rejected_paa)                                                                                 | Get a rejected PAA certificate                                                  | `dcld query pki rejected-x509-root-cert`                |
| [GET_PROPOSED_PAA_TO_REVOKE](./pki.md#get_proposed_paa_to_revoke)                                                             | Gets a proposed but not approved PAA certificate to be revoked                  | `dcld query pki proposed-x509-root-cert-to-revoke`      |
| [GET_ALL_PAA](./pki.md#get_all_paa)                                                                                           | Gets all approved PAA certificates                                              | `dcld query pki all-x509-root-certs`                    |
| [GET_ALL_REVOKED_PAA](./pki.md#get_all_revoked_paa)                                                                           | Gets all revoked PAA certificates                                               | `dcld query pki all-revoked-x509-root-certs`            |
| [GET_ALL_PROPOSED_PAA](./pki.md#get_all_proposed_paa)                                                                         | Gets all proposed but not approved root certificates                            | `dcld query pki all-proposed-x509-root-certs`           |
| [GET_ALL_REJECTED_PAA](./pki.md#get_all_rejected_paa)                                                                         | Gets all rejected root certificates                                             | `dcld query pki all-rejected-x509-root-certs`           |
| [GET_ALL_PROPOSED_PAA_TO_REVOKE](./pki.md#get_all_proposed_paa_to_revoke)                                                     | Gets all proposed but not approved root certificates to be revoked              | `dcld query pki all-proposed-x509-root-certs-to-revoke` |
| NOC                                                                                                                           | Work for NOC certificate types (RCAC, ICAC)                                     |                                                         |
| [ADD_NOC_ROOT](./pki.md#add_noc_root-rcac)                                                                                    | Adds a NOC root certificate (RCAC)                                              | `dcld tx pki add-noc-x509-root-cert`                    |
| [REVOKE_NOC_ROOT](./pki.md#revoke_noc_root-rcac)                                                                              | Revokes a NOC root certificate (RCAC)                                           | `dcld tx pki revoke-noc-x509-root-cert`                 |
| [ADD_NOC_ICA](./pki.md#add_noc_ica-icac)                                                                                      | Adds a NOC ica certificate (ICAC)                                               | `dcld tx pki add-noc-x509-ica-cert`                     |
| [REVOKE_NOC_ICA](./pki.md#revoke_noc_ica-icac)                                                                                | Revokes a NOC ica certificate (ICAC)                                            | `dcld tx pki revoke-noc-x509-ica-cert`                  |
| [REMOVE_NOC_ICA](./pki.md#remove_noc_ica-icac)                                                                                | Remove a NOC ica certificate (ICAC)                                             | `dcld tx pki remove-noc-x509-ica-cert`                  |
| [GET_NOC_CERT](./pki.md#get_noc_cert)                                                                                         | Gets a NOC certificate (RCAC, ICAC)                                             | `dcld query pki noc-x509-cert`                          |
| [GET_NOC_ROOT_BY_VID](./pki.md#get_noc_root_by_vid-rcacs)                                                                     | Retrieve NOC root certificates (RCACs) associated with a specific VID           | `dcld query pki noc-x509-root-certs`                    |
| [GET_NOC_BY_VID_AND_SKID](./pki.md#get_noc_by_vid_and_skid-rcacsicacs)                                                        | Retrieve NOC certificates (RCACs/ICACs) associated with a specific VID and SKID | `dcld query pki noc-x509-certs`                         |
| [GET_NOC_ICA_BY_VID](./pki.md#get_noc_ica_by_vid-icacs)                                                                       | Retrieve NOC ICA certificates (ICACs) associated with a specific VID            | `dcld query pki noc-x509-ica-certs`                     |
| [GET_NOC_CERTS_BY_SUBJECT](./pki.md#get_noc_certs_by_subject)                                                                 | Gets all NOC certificates associated with a subject                             | `dcld query pki all-noc-subject-x509-certs`             |
| [GET_REVOKED_NOC_ROOT](./pki.md#get_revoked_noc_root-rcac)                                                                    | Gets a revoked NOC root certificate (RCAC) by the given subject and SKID        | `dcld query pki revoked-noc-x509-root-cert`             |
| [GET_REVOKED_NOC_ICA](./pki.md#get_revoked_noc_ica-icac)                                                                      | Gets a revoked NOC ica certificate (ICAC) by the given subject and SKID         | `dcld query pki revoked-noc-x509-ica-cert`              |
| [GET_ALL_NOC](./pki.md#get_all_noc-rcacsicacs)                                                                                | Retrieve a list of all of NOC certificates (RCACs of ICACs)                     | `dcld query pki all-noc-x509-certs`                     |
| [GET_ALL_NOC_ROOT](./pki.md#get_all_noc_root-rcacs)                                                                           | Retrieve a list of all of NOC root certificates (RCACs)                         | `dcld query pki all-noc-x509-root-certs`                |
| [GET_ALL_NOC_ICA](./pki.md#get_all_noc_ica-icacs)                                                                             | Retrieve a list of all of NOC ICA certificates (ICACs)                          | `dcld query pki all-noc-x509-ica-certs`                 |
| [GET_ALL_REVOKED_NOC_ROOT](./pki.md#get_all_revoked_noc_root-rcacs)                                                           | Gets all revoked NOC root certificates (RCACs)                                  | `dcld query pki all-revoked-noc-x509-root-certs`        |
| [GET_ALL_REVOKED_NOC_ICA](./pki.md#get_all_revoked_noc_ica-icacs)                                                             | Gets all revoked NOC ica certificates (ICACs)                                   | `dcld query pki all-revoked-noc-x509-ica-certs`         |

### [AUTH](./auth.md)

| Method                                                                               | Description                                               | CLI command                                       |
|--------------------------------------------------------------------------------------|-----------------------------------------------------------|---------------------------------------------------|
| [PROPOSE_ADD_ACCOUNT](./auth.md#propose_add_account)                                 | Proposes a new Account                                    | `dcld tx auth propose-add-account`                |
| [APPROVE_ADD_ACCOUNT](./auth.md#approve_add_account)                                 | Approves the proposed account / Re-vote                   | `dcld tx auth approve-add-account`                |
| [REJECT_ADD_ACCOUNT](./auth.md#reject_add_account)                                   | Rejects the proposed account / Re-vote                    | `dcld tx auth reject-add-account`                 |
| [PROPOSE_REVOKE_ACCOUNT](./auth.md#propose_revoke_account)                           | Proposes revocation of the Account                        | `dcld tx auth propose-revoke-account`             |
| [APPROVE_REVOKE_ACCOUNT](./auth.md#approve_revoke_account)                           | Approves the proposed revocation of the account           | `dcld tx auth approve-revoke-account`             |
| [GET_ACCOUNT](./auth.md#get_account)                                                 | Gets an accounts                                          | `dcld query auth account`                         |
| [GET_PROPOSED_ACCOUNT](./auth.md#get_proposed_account)                               | Gets a proposed but not approved accounts                 | `dcld query auth proposed-account`                |
| [GET_REJECTED_ACCOUNT](./auth.md#get_rejected_account)                               | Get a rejected accounts                                   | `dcld query auth rejected-account`                |
| [GET_PROPOSED_ACCOUNT_TO_REVOKE](./auth.md#get_proposed_account_to_revoke)           | Gets a proposed but not approved accounts to be revoked   | `dcld query auth proposed-account-to-revoke`      |
| [GET_REVOKED_ACCOUNT](./auth.md#get_revoked_account)                                 | Gets a revoked account by its address                     | `dcld query auth revoked-account`                 |
| [GET_ALL_ACCOUNTS](./auth.md#get_all_accounts)                                       | Gets all accounts                                         | `dcld query auth all-accounts`                    |
| [GET_ALL_PROPOSED_ACCOUNTS](./auth.md#get_all_proposed_accounts)                     | Gets all proposed but not approved accounts               | `dcld query auth all-proposed-accounts`           |
| [GET_ALL_REJECTED_ACCOUNTS](./auth.md#get_all_rejected_accounts)                     | Get all rejected accounts                                 | `dcld query auth all-rejected-accounts`           |
| [GET_ALL_PROPOSED_ACCOUNTS_TO_REVOKE](./auth.md#get_all_proposed_accounts_to_revoke) | Gets all proposed but not approved accounts to be revoked | `dcld query auth all-proposed-accounts-to-revoke` |
| [GET_ALL_REVOKED_ACCOUNTS](./auth.md#get_all_revoked_accounts)                       | Gets all revoked accounts                                 | `dcld query auth all-revoked-accounts`            |

### [VALIDATOR_NODE](./validator-node.md)

| Method                                                                                         | Description                                                     | CLI command                                       |
|------------------------------------------------------------------------------------------------|-----------------------------------------------------------------|---------------------------------------------------|
| [ADD_VALIDATOR_NODE](./validator-node.md#add_validator_node)                                   | Adds a new Validator node                                       | `dcld tx validator add-node`                      |
| [DISABLE_VALIDATOR_NODE](./validator-node.md#disable_validator_node)                           | Disables the Validator node by owner                            | `dcld tx validator disable-node`                  |
| [PROPOSE_DISABLE_VALIDATOR_NODE](./validator-node.md#propose_disable_validator_node)           | Proposes disabling of the Validator node by a Trustee           | `dcld tx validator propose-disable-node`          |
| [APPROVE_DISABLE_VALIDATOR_NODE](./validator-node.md#approve_disable_validator_node)           | Approves disabling of the Validator node by a Trustee / Re-vote | `dcld tx validator approve-disable-node`          |
| [REJECT_DISABLE_VALIDATOR_NODE](./validator-node.md#reject_disable_validator_node)             | Rejects disabling of the Validator node by a Trustee / Re-vote  | `dcld tx validator reject-disable-node`           |
| [ENABLE_VALIDATOR_NODE](./validator-node.md#enable_validator_node)                             | Enables the Validator node by the owner                         | `dcld tx validator enable-node`                   |
| [GET_VALIDATOR](./validator-node.md#get_validator)                                             | Gets a validator node                                           | `dcld query validator node`                       |
| [GET_ALL_VALIDATORS](./validator-node.md#get_all_validators)                                   | Gets the list of all validator nodes                            | `dcld query validator all-nodes`                  |
| [GET_PROPOSED_DISABLE_VALIDATOR](./validator-node.md#get_proposed_disable_validator)           | Gets a proposed validator node                                  | `dcld query validator proposed-disable-node`      |
| [GET_ALL_PROPOSED_DISABLE_VALIDATORS](./validator-node.md#get_all_proposed_disable_validators) | Gets the list of all proposed disable validator nodes           | `dcld query validator all-proposed-disable-nodes` |
| [GET_REJECTED_DISABLE_VALIDATOR](./validator-node.md#get_rejected_disable_validator)           | Gets a rejected validator node                                  | `dcld query validator rejected-disable-node`      |
| [GET_ALL_REJECTED_DISABLE_VALIDATORS](./validator-node.md#get_all_rejected_disable_validators) | Gets the list of all rejected disable validator nodes           | `dcld query validator all-rejected-disable-nodes` |
| [GET_DISABLED_VALIDATOR](./validator-node.md#get_disabled_validator)                           | Gets a disabled validator node                                  | `dcld query validator disabled-node`              |
| [GET_ALL_DISABLED_VALIDATORS](./validator-node.md#get_all_disabled_validators)                 | Gets the list of all disabled validator nodes                   | `dcld query validator all-disabled-nodes`         |
| [GET_LAST_VALIDATOR_POWER](./validator-node.md#get_last_validator_power)                       | Gets a last validator node power                                | `dcld query validator last-power`                 |
| [GET_ALL_LAST_VALIDATORS_POWER](./validator-node.md#get_all_last_validators_power)             | Gets the list of all last validator nodes power                 | `dcld query validator all-last-powers`            |
| [UPDATE_VALIDATOR_NODE](./validator-node.md#update_validator_node)                             | Updates the Validator node by the owner                         |                                                   |

### [UPGRADE](./upgrade.md)

| Method                                                              | Description                                                          | CLI command                                   |
|---------------------------------------------------------------------|----------------------------------------------------------------------|-----------------------------------------------|
| [PROPOSE_UPGRADE](./upgrade.md#propose_upgrade)                     | Proposes an upgrade plan with the given name at the given height     | `dcld tx dclupgrade propose-upgrade`          |
| [APPROVE_UPGRADE](./upgrade.md#approve_upgrade)                     | Approves the proposed upgrade plan with the given name / Re-vote     | `dcld tx dclupgrade approve-upgrade`          |
| [REJECT_UPGRADE](./upgrade.md#reject_upgrade)                       | Rejects the proposed upgrade plan with the given name / Re-vote      | `dcld tx dclupgrade reject-upgrade`           |
| [GET_PROPOSED_UPGRADE](./upgrade.md#get_proposed_upgrade)           | Gets the proposed upgrade plan with the given name                   | `dcld query dclupgrade proposed-upgrade`      |
| [GET_APPROVED_UPGRADE](./upgrade.md#get_approved_upgrade)           | Gets the approved upgrade plan with the given name                   | `dcld query dclupgrade approved-upgrade`      |
| [GET_REJECTED_UPGRADE](./upgrade.md#get_rejected_upgrade)           | Gets the rejected upgrade plan with the given name                   | `dcld query dclupgrade rejected-upgrade`      |
| [GET_ALL_PROPOSED_UPGRADES](./upgrade.md#get_all_proposed_upgrades) | Gets all the proposed upgrade plans                                  | `dcld query dclupgrade all-proposed-upgrades` |
| [GET_ALL_APPROVED_UPGRADES](./upgrade.md#get_all_approved_upgrades) | Gets all the approved upgrade plans                                  | `dcld query dclupgrade all-approved-upgrades` |
| [GET_ALL_REJECTED_UPGRADES](./upgrade.md#get_all_rejected_upgrades) | Gets all the rejected upgrade plans                                  | `dcld query dclupgrade all-rejected-upgrades` |
| [GET_UPGRADE_PLAN](./upgrade.md#get_upgrade_plan)                   | Gets the currently scheduled upgrade plan, if it exists              | `dcld query upgrade plan`                     |
| [GET_APPLIED_UPGRADE](./upgrade.md#get_applied_upgrade)             | Gets header block at which the upgrade was applied                   | `dcld query upgrade applied`                  |
| [GET_MODULE_VERSIONS](./upgrade.md#get_module_versions)             | Gets a list of module names and their respective consensus versions  | `dcld query upgrade module_versions`          |

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
