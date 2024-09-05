# Transactions and Queries
<!-- markdownlint-disable MD036 -->

See use case sequence diagrams for the examples of how transaction can be used.

1. [General](#general)
2. [How to write to the Ledger](#how-to-write-to-the-ledger)
3. [How to read from the Ledger](#how-to-read-from-the-ledger)
4. [Vendor Info](#vendor-info)
5. [Model and Model Version](#model-and-model_version)
6. [Compliance](#certify_device_compliance)
7. [X509 PKI](#x509-pki)
    - [Device Attestation (DA): PAA, PAI](#device-attestation-certificates-da-paa-pai)
    - [E2E (NOC): RCAC, ICAC](#e2e-noc-rcac-icac)
8. [Auth](#auth)
9. [Validator Node](#validator_node)
10. [Upgrade](#upgrade)
11. [Extensions](#extensions)

## General

- Every writer to the Ledger must  
  - Have a private/public key pair.
  - Have an Account created on the ledger via `ACCOUNT` transaction (see [Use Case Txn Auth](use_cases/use_cases_txn_auth.puml)).
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
      See `CLI Configuration` section in [how-to.md](how-to.md#cli-configuration).
  - Generate and store a private key for the Account to be used for sending.
      See `Getting Account` section in [how-to.md](how-to.md#getting-account).
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
  - Generate a client code from the proto files [proto](../proto) for the client language (see <https://grpc.io/docs/languages/>)
  - Build, sign, and broadcast the message (transaction).
      See [grpc/rest integration tests](../integration_tests/grpc_rest) as an example.
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
  - If there are no trusted Observer or Validator nodes to connect a CLI, then a [Light Client Proxy](running-light-client-proxy.md) can be used.
- REST API
  - OpenAPI specification: <https://zigbee-alliance.github.io/distributed-compliance-ledger/>.
  - Any running node exposes a REST API at port `1317`. See <https://docs.cosmos.network/v0.45/core/grpc_rest.html>.
  - See `REST API` section for every read request.
  - See [grpc/rest integration tests](../integration_tests/grpc_rest) as an example.
  - There are no state proofs in REST, so REST queries should be sent to trusted Validator or Observer nodes only.
- gRPC
  - Any running node exposes a REST API at port `9090`. See <https://docs.cosmos.network/v0.45/core/grpc_rest.html>.
  - Generate a client code from the proto files [proto](../proto) for the client language (see <https://grpc.io/docs/languages/>).
  - See [grpc/rest integration tests](../integration_tests/grpc_rest) as an example.
  - There are no state proofs in gRPC, so gRPC queries should be sent to trusted Validator or Observer nodes only.
- Tendermint RPC
  - Tendermint RPC OpenAPI specification can be found in <https://zigbee-alliance.github.io/distributed-compliance-ledger/>.
  - Tendermint RPC is exposed by every running node  at port `26657`. See <https://docs.cosmos.network/v0.45/core/grpc_rest.html#tendermint-rpc>.
  - Tendermint RPC supports state proofs. Tendermint's Light Client library can be used to verify the state proofs.
    So, if Light Client API is used, then it's possible to communicate with non-trusted nodes.
  - Please note, that multi-value queries don't have state proofs support and should be sent to trusted nodes only.
  - Refer to [this doc](./cometbft-rpc.md) to see how to [subscribe](./cometbft-rpc.md#subscribe) to a Tendermint WebSocket based events and/or [query](./cometbft-rpc.md#querying-application-components) an application components. 

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

## VENDOR INFO

### ADD_VENDOR_INFO

**Status: Implemented**

Adds a record about a Vendor.

- Parameters:
  - vid: `uint16` -  Vendor ID (positive non-zero)
  - vendorName: `string` -  Vendor name
  - companyLegalName: `string` -  Legal name of the vendor company
  - companyPreferredName: `optional(string)` -  Preferred name of the vendor company
  - vendorLandingPageURL: `optional(string)` -  URL of the vendor's landing page
  - schemaVersion: `optional(uint16)` - Schema version to support backward/forward compatability. Should be equal to 0 (default 0)
- In State: `vendorinfo/VendorInfo/value/<vid>`
- Who can send:
  - Account with a vendor role who has the matching Vendor ID
  - Account with a vendor admin role
- CLI command:
  - `dcld tx vendorinfo add-vendor --vid=<uint16> --vendorName=<string> --companyLegalName=<string> --companyPreferredName=<string> --vendorLandingPageURL=<string> --from=<account>`

### UPDATE_VENDOR_INFO

**Status: Implemented**

Updates a record about a Vendor.

- Parameters:
  - vid: `uint16` -  Vendor ID (positive non-zero)
  - vendorName: `optional(string)` -  Vendor name
  - companyLegalName: `optional(string)` -  Legal name of the vendor company
  - companyPreferredName: `optional(string)` -  Preferred name of the vendor company
  - vendorLandingPageURL: `optional(string)` -  URL of the vendor's landing page
  - schemaVersion: `optional(uint16)` - Schema version to support backward/forward compatability. Should be equal to 0 (default 0)
- In State: `vendorinfo/VendorInfo/value/<vid>`
- Who can send:
  - Account with a vendor role who has the matching Vendor ID
  - Account with a vendor admin role
- CLI command:
  - `dcld tx vendorinfo update-vendor --vid=<uint16> ... --from=<account>`

### GET_VENDOR_INFO

**Status: Implemented**

Gets a Vendor Info for the given `vid` (vendor ID).

- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
- CLI command:
  - `dcld query vendorinfo vendor --vid=<uint16>`
- REST API:
  - GET `/dcl/vendorinfo/vendors/{vid}`

### GET_ALL_VENDOR_INFO

**Status: Implemented**

Gets information about all vendors for all VIDs.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query vendorinfo all-vendors`
- REST API:
  - GET `/dcl/vendorinfo/vendors`

## MODEL and MODEL_VERSION

### ADD_MODEL

**Status: Implemented**

Adds a new Model identified by a unique combination of `vid` (vendor ID) and `pid` (product ID). If `account` was created with product ID ranges then the `pid` must fall within that specified range.

Not all fields can be edited (see `EDIT_MODEL`).

- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - deviceTypeID: `uint16` -  DeviceTypeID is the device type identifier. For example, DeviceTypeID 10 (0x000a), is the device type identifier for a Door Lock.
  - productName: `string` -  model name
  - productLabel: `optional(string)` -  model description (string or path to file containing data)
  - partNumber: `optional(string)` -  stock keeping unit
  - commissioningCustomFlow: `optional(uint8)` - A value of 1 indicates that user interaction with the device (pressing a button, for example) is required before commissioning can take place. When CommissioningCustomflow is set to a value of 2, the commissioner SHOULD attempt to obtain a URL which MAY be used to provide an end user with the necessary details for how to configure the product for initial commissioning
  - commissioningCustomFlowURL: `optional(string)` - commissioningCustomFlowURL SHALL identify a vendor specific commissioning URL for the device model when the commissioningCustomFlow field is set to '2'
  - commissioningModeInitialStepsHint: `optional(uint32)` - commissioningModeInitialStepsHint SHALL identify a hint for the steps that can be used to put into commissioning mode a device that has not yet been commissioned. This field is a bitmap with values defined in the Pairing Hint Table. For example, a value of 1 (bit 0 is set) indicates that a device that has not yet been commissioned will enter Commissioning Mode upon a power cycle (default 1).
  - commissioningModeInitialStepsInstruction: `optional(string)` - commissioningModeInitialStepsInstruction SHALL contain text which relates to specific values of CommissioningModeInitialStepsHint. Certain values of CommissioningModeInitialStepsHint, as defined in the Pairing Hint Table, indicate a Pairing Instruction (PI) dependency, and for these values the commissioningModeInitialStepsInstruction SHALL be set
  - commissioningModeSecondaryStepsHint: `optional(uint32)` - commissioningModeSecondaryStepsHint SHALL identify a hint for steps that can be used to put into commissioning mode a device that has already been commissioned. This field is a bitmap with values defined in the Pairing Hint Table. For example, a value of 4 (bit 2 is set) indicates that a device that has already been commissioned will require the user to visit a current CHIP Administrator to put the device into commissioning mode.
  - commissioningModeSecondaryStepInstruction: `optional(string)` - commissioningModeSecondaryStepInstruction SHALL contain text which relates to specific values of commissioningModeSecondaryStepsHint. Certain values of commissioningModeSecondaryStepsHint, as defined in the Pairing Hint Table, indicate a Pairing Instruction (PI) dependency, and for these values the commissioningModeSecondaryStepInstruction SHALL be set
  - userManualURL: `optional(string)` - URL that contains product specific web page that contains user manual for the device model.
  - supportURL: `optional(string)` - URL that contains product specific web page that contains support details for the device model.
  - productURL: `optional(string)` - URL that contains product specific web page that contains details for the device model.
  - lsfURL: `optional(string)` - URL to the Localized String File of this product.
  - enhancedSetupFlowOptions: `optional(uint16)` - enhancedSetupFlowOptions SHALL identify the configuration options for the Enhanced Setup Flow.
  - enhancedSetupFlowTCUrl: `optional(string)` - enhancedSetupFlowTCUrl SHALL identify a link to the Enhanced Setup Flow Terms and Condition File for this product. This field SHALL be present if and only if the EnhancedSetupFlowOptions field has bit 0 set.
  - enhancedSetupFlowTCRevision: `optional(uint16)` - enhancedSetupFlowTCRevision is an increasing positive integer indicating the latest available version of the Enhanced Setup Flow Terms and Conditions file. This field SHALL be present if and only if the EnhancedSetupFlowOptions field has bit 0 set.
  - enhancedSetupFlowTCDigest: `optional(string)` - enhancedSetupFlowTCDigest SHALL contain the digest of the entire contents of the associated file downloaded from the EnhancedSetupFlowTCUrl field, encoded in base64 string representation and SHALL be used to ensure the contents of the downloaded file are authentic. This field SHALL be present if and only if the EnhancedSetupFlowOptions field has bit 0 set.
  - enhancedSetupFlowTCFileSize: `optional(uint32)` - enhancedSetupFlowTCFileSize SHALL indicate the total size of the Enhanced Setup Flow Terms and Conditions file in bytes, and SHALL be used to ensure the downloaded file size is within the bounds of EnhancedSetupFlowTCFileSize. This field SHALL be present if and only if the EnhancedSetupFlowOptions field has bit 0 set.
  - maintenanceUrl: `optional(string)` - maintenanceUrl SHALL identify a link to a vendor-specific URL which SHALL provide a manufacturer specific means to resolve any functionality limitations indicated by the TERMS_AND_CONDITIONS_CHANGED status code. This field SHALL be present if and only if the EnhancedSetupFlowOptions field has bit 0 set.
  - schemaVersion: `optional(uint16)` - Schema version to support backward/forward compatability. Should be equal to 0 (default 0)
  - discoveryCapabilitiesBitmask: `optional(uint16)` - Identifies the device's available technologies for device discovery (default 0)
  - commissioningFallbackURL: `optional(string)` - This field SHALL identify a vendor-specific commissioning-fallback URL for the device model, which can be used by a Commissioner in case commissioning fails to direct the user to a manufacturer-provided mechanism to provide resolution to commissioning issues.
- In State:
  - `model/Model/value/<vid>/<pid>`
  - `model/VendorProducts/value/<vid>`
- Who can send:
  - Vendor account who is associated with the given vid
- CLI command minimal:

```bash
dcld tx model add-model --vid=<uint16> --pid=<uint16> --deviceTypeID=<uint16> --productName=<string> --productLabel=<string or path> --partNumber=<string> --from=<account>
```

- CLI command full:

```bash
dcld tx model add-model --vid=<uint16> --pid=<uint16> --deviceTypeID=<uint16> --productName=<string> --productLabel=<string or path> --partNumber=<string> 
    --commissioningCustomFlow=<uint8> --commissioningCustomFlowUrl=<string> --commissioningModeInitialStepsHint=<uint32> --commissioningModeInitialStepsInstruction=<string>
    --commissioningModeSecondaryStepsHint=<uint32> --commissioningModeSecondaryStepsInstruction=<string> --userManualURL=<string> --supportURL=<string> --productURL=<string> --lsfURL=<string> --discoveryCapabilitiesBitmask=<uint16> --commissioningFallbackURL<string>
    --from=<account>
```

### EDIT_MODEL

**Status: Implemented**

Edits an existing Model identified by a unique combination of `vid` (vendor ID) and `pid` (product ID)
by the vendor account. If `account` was created with product ID ranges then the `pid` must fall within that specified range.

Only the fields listed below (except `vid` and `pid`) can be edited. If other fields need to be edited -
a new model info with a new `vid` or `pid` can be created.

All non-edited fields remain the same.

If one of EnhancedSetupFlow or MaintenanceUrl fields needs to be updated, ALL EnhancedSetupFlow fields MUST be specified, and EnhancedSetupFlowOptions field must have bit 0 set.

- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - productName: `optional(string)` -  model name
  - productLabel: `optional(string)` -  model description (string or path to file containing data)
  - partNumber: `optional(string)` -  stock keeping unit
  - commissioningCustomFlowURL: `optional(string)` - commissioningCustomFlowURL SHALL identify a vendor specific commissioning URL for the device model when the commissioningCustomFlow field is set to '2'
  - commissioningModeInitialStepsInstruction: `optional(string)` - commissioningModeInitialStepsInstruction SHALL contain text which relates to specific values of CommissioningModeInitialStepsHint. Certain values of CommissioningModeInitialStepsHint, as defined in the Pairing Hint Table, indicate a Pairing Instruction (PI) dependency, and for these values the commissioningModeInitialStepsInstruction SHALL be set
  - commissioningModeSecondaryStepInstruction: `optional(string)` - commissioningModeSecondaryStepInstruction SHALL contain text which relates to specific values of commissioningModeSecondaryStepsHint. Certain values of commissioningModeSecondaryStepsHint, as defined in the Pairing Hint Table, indicate a Pairing Instruction (PI) dependency, and for these values the commissioningModeSecondaryStepInstruction SHALL be set
  - userManualURL: `optional(string)` - URL that contains product specific web page that contains user manual for the device model.
  - supportURL: `optional(string)` - URL that contains product specific web page that contains support details for the device model.
  - productURL: `optional(string)` - URL that contains product specific web page that contains details for the device model.  
  - lsfURL: `optional(string)` - URL to the Localized String File of this product.
  - lsfRevision: `optional(uint32)` - LsfRevision is a monotonically increasing positive integer indicating the latest available version of Localized String File.
  - commissioningModeInitialStepsHint: `optional(uint32)` - commissioningModeInitialStepsHint SHALL identify a hint for the steps that can be used to put into commissioning mode a device that has not yet been commissioned. This field is a bitmap with values defined in the Pairing Hint Table. For example, a value of 1 (bit 0 is set) indicates that a device that has not yet been commissioned will enter Commissioning Mode upon a power cycle. Note that this value cannot be updated to 0.
  - enhancedSetupFlowOptions: `optional(uint16)` - enhancedSetupFlowOptions SHALL identify the configuration options for the Enhanced Setup Flow.
  - enhancedSetupFlowTCUrl: `optional(string)` - enhancedSetupFlowTCUrl SHALL identify a link to the Enhanced Setup Flow Terms and Condition File for this product. This field SHALL be present if and only if the EnhancedSetupFlowOptions field has bit 0 set.
  - enhancedSetupFlowTCRevision: `optional(uint16)` - enhancedSetupFlowTCRevision is an increasing positive integer indicating the latest available version of the Enhanced Setup Flow Terms and Conditions file. This field SHALL be present if and only if the EnhancedSetupFlowOptions field has bit 0 set.
  - enhancedSetupFlowTCDigest: `optional(string)` - enhancedSetupFlowTCDigest SHALL contain the digest of the entire contents of the associated file downloaded from the EnhancedSetupFlowTCUrl field, encoded in base64 string representation and SHALL be used to ensure the contents of the downloaded file are authentic. This field SHALL be present if and only if the EnhancedSetupFlowOptions field has bit 0 set.
  - enhancedSetupFlowTCFileSize: `optional(uint32)` - enhancedSetupFlowTCFileSize SHALL indicate the total size of the Enhanced Setup Flow Terms and Conditions file in bytes, and SHALL be used to ensure the downloaded file size is within the bounds of EnhancedSetupFlowTCFileSize. This field SHALL be present if and only if the EnhancedSetupFlowOptions field has bit 0 set.
  - maintenanceUrl: `optional(string)` - maintenanceUrl SHALL identify a link to a vendor-specific URL which SHALL provide a manufacturer specific means to resolve any functionality limitations indicated by the TERMS_AND_CONDITIONS_CHANGED status code. This field SHALL be present if and only if the EnhancedSetupFlowOptions field has bit 0 set.
  - schemaVersion: `optional(uint16)` - Schema version to support backward/forward compatability. Should be equal to 0 (default 0)
  - commissioningFallbackURL: `optional(string)` - This field SHALL identify a vendor-specific commissioning-fallback URL for the device model, which can be used by a Commissioner in case commissioning fails to direct the user to a manufacturer-provided mechanism to provide resolution to commissioning issues.
- In State: `model/Model/value/<vid>/<pid>`
- Who can send:
  - Vendor account associated with the same vid who has created the model
- CLI command:
  - `dcld tx model update-model --vid=<uint16> --pid=<uint16> ... --from=<account>`

### DELETE_MODEL

**Status: Implemented**

Deletes an existing Model identified by a unique combination of `vid` (vendor ID) and `pid` (product ID)
by the vendor account. If `account` was created with product ID ranges then the `pid` must fall within that specified range.

If one of Model Versions associated with the Model is certified then Model can not be deleted. When Model is deleted, all associated Model Versions will be deleted as well.

- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
- In State: `model/Model/value/<vid>/<pid>`
- Who can send:
  - Vendor account associated with the same vid who has created the model
- CLI command:
  - `dcld tx model delete-model --vid=<uint16> --pid=<uint16> --from=<account>`

### ADD_MODEL_VERSION

**Status: Implemented**

Adds a new Model Software Version identified by a unique combination of `vid` (vendor ID), `pid` (product ID) and `softwareVersion`.
If `account` was created with product ID ranges then the `pid` must fall within that specified range

Not all Model Software Version fields can be edited (see `EDIT_MODEL_VERSION`).

If one of `OTA_URl`, `OTA_checksum` or `OTA_checksum_type` fields is set, then the other two must also be set.

- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - softwareVersion: `uint32` - model software version
  - softwareVersionSting: `string` - model software version string
  - cdVersionNumber `uint16` - CD Version Number of the certification
  - minApplicableSoftwareVersion `uint32` - MinApplicableSoftwareVersion should specify the lowest SoftwareVersion for which this image can be applied
  - maxApplicableSoftwareVersion `uint32` - MaxApplicableSoftwareVersion should specify the highest SoftwareVersion for which this image can be applied
  - firmwareInformation `optional(string)` - FirmwareInformation field included in the Device Attestation response when this Software Image boots on the device
  - softwareVersionValid `optional(bool)` - Flag to indicate whether the software version is valid or not (default true)
  - otaURL `optional(string)` - URL where to obtain the OTA image
  - otaFileSize `optional(string)`  - OtaFileSize is the total size of the OTA software image in bytes
  - otaChecksum `optional(string)` - Digest of the entire contents of the associated OTA Software Update Image under the OtaUrl attribute, encoded in base64 string representation. The digest SHALL have been computed using the algorithm specified in OtaChecksumType
  - otaChecksumType `optional(string)` - Numeric identifier as defined in IANA Named Information Hash Algorithm Registry for the type of otaChecksum. For example, a value of 1 would match the sha-256 identifier, which maps to the SHA-256 digest algorithm
  - releaseNotesURL `optional(string)` - URL that contains product specific web page that contains release notes for the device model.
  - schemaVersion: `optional(uint16)` - Schema version to support backward/forward compatability. Should be equal to 0 (default 0)
- In State:
  - `model/ModelVersion/value/<vid>/<pid>/<softwareVersion>`
  - `model/ModelVersions/value/<vid>/<pid>`
- Who can send:
  - Vendor with same vid who created the Model
- CLI command minimal:

```bash
dcld tx model add-model-version --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --softwareVersionString=<string> --cdVersionNumber=<uint32>
--minApplicableSoftwareVersion=<uint32> --maxApplicableSoftwareVersion=<uint32> --from=<account>
```

- CLI command full:

```bash
dcld tx model add-model-version --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --softwareVersionString=<string> --cdVersionNumber=<uint32>
--minApplicableSoftwareVersion=<uint32> --maxApplicableSoftwareVersion=<uint32>
--firmwareInformation=<string> --softwareVersionValid=<bool> --otaURL=<string> --otaFileSize=<string> --otaChecksum=<string> --otaChecksumType=<string> --releaseNotesURL=<string> 
--from=<account>
```

### EDIT_MODEL_VERSION

**Status: Implemented**

Edits an existing Model Software Version identified by a unique combination of `vid` (vendor ID) `pid` (product ID) and `softwareVersion`
by the vendor. If `account` was created with product ID ranges then the `pid` must fall within that specified range.

Only the fields listed below (except `vid` `pid` and `softwareVersion`)  can be edited.

All non-edited fields remain the same.

`otaURL` can be edited only if  `otaFileSize`, `otaChecksum` and `otaChecksumType` are already set.

- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - softwareVersion: `uint32` - model software version (positive non-zero)
  - softwareVersionValid `optional(bool)` - Flag to indicate whether the software version is valid or not (default true)
  - otaURL `optional(string)` - URL where to obtain the OTA image
  - maxApplicableSoftwareVersion `optional(uint32)` - MaxApplicableSoftwareVersion should specify the highest SoftwareVersion for which this image can be applied
  - minApplicableSoftwareVersion `optional(uint32)` - MinApplicableSoftwareVersion should specify the lowest SoftwareVersion for which this image can be applied
  - releaseNotesURL `optional(string)` - URL that contains product specific web page that contains release notes for the device model.
  - otaURL `optional(string)` - URL where to obtain the OTA image
  - otaFileSize `optional(string)`  - OtaFileSize is the total size of the OTA software image in bytes
  - otaChecksum `optional(string)` - Digest of the entire contents of the associated OTA Software Update Image under the OtaUrl attribute, encoded in base64 string representation. The digest SHALL have been computed using the algorithm specified in OtaChecksumType
  - schemaVersion: `optional(uint16)` - Schema version to support backward/forward compatability. Should be equal to 0 (default 0)

- In State: `model/ModelVersion/value/<vid>/<pid>/<softwareVersion>`
- Who can send:
  - Vendor associated with the same vid who created the Model
- CLI command:
  - `dcld tx model update-model-version --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> ... --from=<account>`

### DELETE_MODEL_VERSION

**Status: Implemented**

Deletes an existing Model Version identified by a unique combination of `vid` (vendor ID), `pid` (product ID) and `softwareVersion`
by the vendor account. If `account` was created with product ID ranges then the `pid` must fall within that specified range.

Model Version can be deleted only before it is certified.

- Parameters:
  - vid: `uint16` -  model version vendor ID (positive non-zero)
  - pid: `uint16` -  model version product ID (positive non-zero)
  - softwareVersion: `uint32` - model version software version (positive non-zero)
- In State: `model/ModelVersion/value/<vid>/<pid>/<softwareVersion>`
- Who can send:
  - Vendor account associated with the same vid who has created the model version
- CLI command:
  - `dcld tx model delete-model-version --vid=< uint16 > --pid=< uint16 > --softwareVersion=<uint32> --from=<account>`

### GET_MODEL

**Status: Implemented**

Gets a Model Info with the given `vid` (vendor ID) and `pid` (product ID).

- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
- CLI command:
  - `dcld query model get-model --vid=<uint16> --pid=<uint16>`
- REST API:
  - GET `/dcl/model/models/{vid}/{pid}`

### GET_MODEL_VERSION

**Status: Implemented**

Gets a Model Software Versions for the given `vid`, `pid` and `softwareVersion`.

- Parameters
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - softwareVersion: `uint32` - model software version (positive non-zero)
- CLI command:
  - `dcld query model model-version --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32>`
- REST API:
  - GET `/dcl/model/versions/{vid}/{pid}/{softwareVersion}`

### GET_ALL_MODELS

**Status: Implemented**

Gets all Model Infos for all vendors.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query model all-models`
- REST API:
  - GET `/dcl/model/models`

### GET_ALL_VENDOR_MODELS

**Status: Implemented**

Gets all Model Infos by the given Vendor (`vid`).

- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
- CLI command:
  - `dcld query model vendor-models --vid=<uint16>`
- REST API:
  - GET `/dcl/model/models/{vid}`

### GET_ALL_MODEL_VERSIONS

**Status: Implemented**

Gets all Model Software Versions for the given `vid` and `pid` combination.

- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
- CLI command:
  - `dcld query model all-model-versions --vid=<uint16> --pid=<uint16>`
- REST API:
  - GET `/dcl/model/versions/{vid}/{pid}`

## CERTIFY_DEVICE_COMPLIANCE

### CERTIFY_MODEL

**Status: Implemented**

Attests compliance of the Model Version to the ZB or Matter standard.

`REVOKE_MODEL_CERTIFICATION` should be used for revoking (disabling) the compliance.
It's possible to call `CERTIFY_MODEL` for revoked model versions to enable them back.

The corresponding Model and Model Version must be present on the ledger.

It must be called for every compliant device for use cases where compliance
is tracked on ledger.

It can be used for use cases where only revocation is tracked on the ledger to remove a Model Version
from the revocation list.

- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - softwareVersion: `uint32` - model software version
  - softwareVersionSting: `string` - model software version string
  - certificationDate: `string` - The date of model certification (rfc3339 encoded), for example 2019-10-12T07:20:50.52Z
  - certificationType: `string` - Certification type - Currently 'zigbee', 'matter', 'access control', 'product security' types are supported
  - cdCertificateId: `string` - CD Certificate ID 
  - reason `optional(string)` - optional comment describing the reason of the certification
  - cDVersionNumber `optional(uint32)` - optional field describing the CD version number
  - familyId `optional(string)` - optional field describing the family ID
  - supportedClusters `optional(string)` - optional field describing the supported clusters
  - compliantPlatformUsed `optional(string)` - optional field describing the compliant platform used
  - compliantPlatformVersion `optional(string)` - optional field describing the compliant platform version
  - OSVersion `optional(string)` - optional field describing the OS version
  - certificationRoute `optional(string)` - optional field describing the certification route
  - programType `optional(string)` - optional field describing the program type
  - programTypeVersion `optional(string)` - optional field describing the program type version
  - transport `optional(string)` - optional field describing the transport
  - parentChild `optional(string)` - optional field describing the parent/child - Currently 'parent' and 'child' types are supported
  - certificationIDOfSoftwareComponent `optional(string)` - optional field describing the certification ID of software component
  - schemaVersion: `optional(uint16)` - Schema version to support backward/forward compatability. Should be equal to 0 (default 0)
- In State:
  - `compliance/ComplianceInfo/value/<vid>/<pid>/<softwareVersion>/<certificationType>`
  - `compliance/CertifiedModel/value/<vid>/<pid>/<softwareVersion>/<certificationType>`
- Who can send:
  - CertificationCenter
- CLI command:
  - `dcld tx compliance certify-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --softwareVersionString=<string>  --certificationType=<matter|zigbee|access control|product security> --certificationDate=<rfc3339 encoded date> --cdCertificateId=<string> --from=<account>`
- CLI command full:
  - `dcld tx compliance certify-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --softwareVersionString=<string>  --certificationType=<matter|zigbee|access control|product security> --certificationDate=<rfc3339 encoded date> --cdCertificateId=<string> --reason=<string> --cDVersionNumber=<uint32> --familyId=<string> --supportedClusters=<string> --compliantPlatformUsed=<string> --compliantPlatformVersion=<string> --OSVersion=<string> --certificationRoute=<string> --programType=<string> --programTypeVersion=<string> --transport=<string> --parentChild=<string> --certificationIDOfSoftwareComponent=<string> --from=<account>`

### UPDATE_COMPLIANCE_INFO

**Status: Implemented**

Updates a compliance info by VID, PID, Software Version and Certification Type.


- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - softwareVersion: `uint32` - model software version
  - certificationType: `string` - Certification type - Currently 'zigbee', 'matter', 'access control', 'product security' types are supported
  - certificationDate: `optional(string)` - The date of model certification (rfc3339 encoded), for example 2019-10-12T07:20:50.52Z
  - cdCertificateId: `optional(string)` - CD Certificate ID
  - reason `optional(string)` - optional comment describing the reason of the certification
  - cDVersionNumber `optional(string)` - optional field (a uint32-parsable string) describing the CD version number, must be the same with the associated model version
  - familyId `optional(string)` - optional field describing the family ID
  - supportedClusters `optional(string)` - optional field describing the supported clusters
  - compliantPlatformUsed `optional(string)` - optional field describing the compliant platform used
  - compliantPlatformVersion `optional(string)` - optional field describing the compliant platform version
  - OSVersion `optional(string)` - optional field describing the OS version
  - certificationRoute `optional(string)` - optional field describing the certification route
  - programType `optional(string)` - optional field describing the program type
  - programTypeVersion `optional(string)` - optional field describing the program type version
  - transport `optional(string)` - optional field describing the transport
  - parentChild `optional(string)` - optional field describing the parent/child - Currently 'parent' and 'child' types are supported
  - certificationIDOfSoftwareComponent `optional(string)` - optional field describing the certification ID of software component
  - schemaVersion: `optional(uint16)` - Schema version to support backward/forward compatability. Should be equal to 0 (default 0)
- Who can send:
  - CertificationCenter
- CLI command:
  - `dcld tx compliance update-compliance-info`
- CLI command full:
  - `dcld tx compliance update-compliance-info --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<string> --cdVersionNumber=<string> --certificationDate=$upd_certification_date --reason=$upd_reason --cdCertificateId=$upd_cd_certificate_id --certificationRoute=$upd_certification_route --programType=$upd_program_type --programTypeVersion=$upd_program_type_version --compliantPlatformUsed=$upd_compliant_platform_used --compliantPlatformVersion=$upd_compliant_platform_version --transport=$upd_transport --familyId=$upd_familyID --supportedClusters=$upd_supported_clusters --OSVersion=$upd_os_version --parentChild=$upd_parent_child --certificationIDOfSoftwareComponent=$upd_certification_id_of_software_component --from=$zb_account`
- REST API:
  - `/dcl/compliance/update-compliance-info`

### DELETE_COMPLIANCE_INFO

**Status: Implemented**

Delete compliance of the Model Version to the ZB or Matter standard.

The corresponding Compliance Info is required to be present on the ledger

- Parameters:
  - vid: `uint16` - model vendor ID (positive non-zero)
  - pid: `uint16` - model product ID (positive non-zero)
  - softwareVersion: `uint32` - model software version
  - certificationType: `string` - Certification type - Currently 'zigbee' and 'matter', 'access control', 'product security' types are supported
- Who can send:
  - CertificationCenter
- CLI command:
  - `dcld tx compliance delete-compliance-info --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<matter|zigbee|access control|product security> --from=<account>`

### REVOKE_MODEL_CERTIFICATION

**Status: Implemented**

Revoke compliance of the Model Version to the ZB or Matter standard.

The corresponding Model and Model Version are not required to be present on the ledger.

It can be used in cases where every compliance result
is written on the ledger (`CERTIFY_MODEL` was called), or
 cases where only revocation list is stored on the ledger.

- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - softwareVersion: `uint32` - model software version
  - softwareVersionSting: `string` - model software version string
  - revocationDate: `string` - The date of model revocation (rfc3339 encoded), for example 2019-10-12T07:20:50.52Z
  - certificationType: `string`  - Certification type - Currently 'zigbee' and 'matter', 'access control', 'product security' types are supported
  - reason `optional(string)`  - optional comment describing the reason of revocation
  - schemaVersion: `optional(uint16)` - Schema version to support backward/forward compatability. Should be equal to 0 (default 0)
- In State:
  - `compliance/ComplianceInfo/value/<vid>/<pid>/<softwareVersion>/<certificationType>`
  - `compliance/RevokedModel/value/<vid>/<pid>/<softwareVersion>/<certificationType>`
- Who can send:
  - CertificationCenter
- CLI command:
  - `dcld tx compliance revoke-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<matter|zigbee|access control|product security> --revocationDate=<rfc3339 encoded date> --reason=<string> --from=<account>`

### PROVISION_MODEL

**Status: Implemented**

Sets provisional state for the Model Version.

The corresponding Model and Model Version is required to be present in the ledger.

Can not be set if there is already a certification record on the ledger (certified or revoked).

- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - softwareVersion: `uint32` - model software version
  - softwareVersionSting: `string` - model software version string
  - provisionalDate: `string` - The date of model provisioning (rfc3339 encoded), for example 2019-10-12T07:20:50.52Z
  - certificationType: `string`  - Certification type - Currently 'zigbee' and 'matter', 'access control', 'product security' types are supported
  - cdCertificateId: `string` - CD Certificate ID 
  - reason `optional(string)`  - optional comment describing the reason of revocation
  - cDVersionNumber `optional(uint32)` - optional field describing the CD version number
  - familyId `optional(string)` - optional field describing the family ID
  - supportedClusters `optional(string)` - optional field describing the supported clusters
  - compliantPlatformUsed `optional(string)` - optional field describing the compliant platform used
  - compliantPlatformVersion `optional(string)` - optional field describing the compliant platform version
  - OSVersion `optional(string)` - optional field describing the OS version
  - certificationRoute `optional(string)` - optional field describing the certification route
  - programType `optional(string)` - optional field describing the program type
  - programTypeVersion `optional(string)` - optional field describing the program type version
  - transport `optional(string)` - optional field describing the transport
  - parentChild `optional(string)` - optional field describing the parent/child - Currently 'parent' and 'child' types are supported
  - certificationIDOfSoftwareComponent `optional(string)` - optional field describing the certification ID of software component
  - schemaVersion: `optional(uint16)` - Schema version to support backward/forward compatability. Should be equal to 0 (default 0)
- In State:
  - `compliance/ComplianceInfo/value/<vid>/<pid>/<softwareVersion>/<certificationType>`
  - `compliance/ProvisionalModel/value/<vid>/<pid>/<softwareVersion>/<certificationType>`
- Who can send:
  - CertificationCenter
- CLI command:
  - `dcld tx compliance provision-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<matter|zigbee|access control|product security> --provisionalDate=<rfc3339 encoded date> --from=<account>`
- CLI command full:
  - `dcld tx compliance provision-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<matter|zigbee|access control|product security> --provisionalDate=<rfc3339 encoded date> --cdCertificateId=<string> --reason=<string> --cDVersionNumber=<uint32> --familyId=<string> --supportedClusters=<string> --compliantPlatformUsed=<string> --compliantPlatformVersion=<string> --OSVersion=<string> --certificationRoute=<string> --programType=<string> --programTypeVersion=<string> --transport=<string> --parentChild=<string> --certificationIDOfSoftwareComponent=<string> --from=<account>`

### GET_CERTIFIED_MODEL

**Status: Implemented**

Gets a structure containing the Model Version / Certification Type key (`vid`, `pid`, `softwareVersion`, `certificationType`) and a flag (`value`) indicating whether the given Model Version is compliant to `certificationType` standard.

This is the aggregation of compliance and
revocation information for every vid/pid/softwareVersion/certificationType. It should be used in cases where compliance
is tracked on the ledger.

This function responds with `NotFound` (404 code) if Model Version was never certified earlier.

This function returns `true` if compliance information is found on ledger and it's in `certified` state.

You can use `GET_COMPLICE_INFO` method to get the whole compliance information.

- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - softwareVersion: `uint32` - model software version
  - certification_type: `string`  - Certification type - Currently 'zigbee' and 'matter' types are supported
- CLI command:
  - `dcld query compliance certified-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<zigbee|matter|access control|product security>`
- REST API:
  - GET `/dcl/compliance/certified-models/{vid}/{pid}/{software_version}/{certification_type}`

### GET_REVOKED_MODEL

**Status: Implemented**

Gets a structure containing the Model Version / Certification Type key (`vid`, `pid`, `softwareVersion`, `certificationType`) and a flag (`value`) indicating whether the given Model Version is revoked for `certificationType` standard.

It contains information about revocation only, so it should be used in cases
 where only revocation is tracked on the ledger.

This function responds with `NotFound` (404 code) if Model Version was never certified or revoked earlier.

This function returns `true` if compliance information is found on ledger and it's in `revoked` state.

You can use `GET_COMPLICE_INFO` method to get the whole compliance information.

- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - softwareVersion: `uint32` - model software version
  - certification_type: `string`  - Certification type - Currently 'zigbee' and 'matter' types are supported
- CLI command:
  - `dcld query compliance revoked-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<zigbee|matter|access control|product security>`
- REST API:
  - GET `/dcl/compliance/revoked-models/{vid}/{pid}/{software_version}/{certification_type}`

### GET_PROVISIONAL_MODEL

**Status: Implemented**

Gets a structure containing the Model Version / Certification Type key (`vid`, `pid`, `softwareVersion`, `certificationType`) and a flag (`value`) indicating whether the given Model Version is in provisional state for `certificationType` standard.

This function responds with `NotFound` (404 code) if Model Version was never provisioned or certified earlier.

This function returns `true` if compliance information is found on the ledger and it's in `provisional` state.

You can use `GET_COMPLICE_INFO` method to get the whole compliance information.

- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - softwareVersion: `uint32` - model software version
  - certification_type: `string`  - Certification type - Currently 'zigbee' and 'matter' types are supported
- CLI command:
  - `dcld query compliance provisional-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<zigbee|matter|access control|product security>`
- REST API:
  - GET `/dcl/compliance/provisional-models/{vid}/{pid}/{software_version}/{certification_type}`

### GET_COMPLIANCE_INFO

**Status: Implemented**

Gets compliance information associated with the Model Version and Certification Type (identified by the `vid`, `pid`, `softwareVersion` and `certification_type`).

It can be used instead of GET_CERTIFIED_MODEL / GET_REVOKED_MODEL / GET_PROVISIONAL_MODEL methods
to get the whole compliance information without additional state check.

This function responds with `NotFound` (404 code) if compliance information is not found in store.

- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - softwareVersion: `uint32` - model software version
  - certification_type: `string`  - Certification type - Currently 'zigbee' and 'matter' types are supported
- CLI command:
  - `dcld query compliance compliance-info --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<zigbee|matter|access control|product security>`
- REST API:
  - GET `/dcl/compliance/compliance-info/{vid}/{pid}/{software_version}/{certification_type}`

### GET_DEVICE_SOFTWARE_COMPLIANCE

**Status: Implemented**

Gets device software compliance associated with the `cDCertificateId`.

This function responds with `NotFound` (404 code) if device software compliance is not found in store.

- Parameters:
  - cDCertificateId: `string` - CD Certificate ID
- CLI command:
  - `dcld query compliance device-software-compliance --cDCertificateId=<string>`
- REST API:
  - GET `/dcl/compliance/device-software-compliance/{cDCertificateId}`

### GET_ALL_CERTIFIED_MODELS

**Status: Implemented**

Gets all compliant Model Versions for all vendors (`vid`s).

This is the aggregation of compliance and
revocation information for every vid/pid. It should be used in cases where compliance is tracked on ledger.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query compliance all-certified-models`
- REST API:
  - GET `/dcl/compliance/certified-models`

### GET_ALL_REVOKED_MODELS

**Status: Implemented**

Gets all revoked Model Versions for all vendors (`vid`s).

It contains information about revocation only, so it should be used in cases
 where only revocation is tracked on the ledger.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query compliance all-revoked-models`
- REST API:
  - GET `/dcl/compliance/revoked-models`

### GET_ALL_PROVISIONAL_MODELS

**Status: Implemented**

Gets all Model Versions in provisional state for all vendors (`vid`s).

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query compliance all-provisional-models`
- REST API:
  - GET `/dcl/compliance/provisional-models`

### GET_ALL_COMPLIANCE_INFO_RECORDS

**Status: Implemented**

Gets all stored compliance information records for all vendors (`vid`s).

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query compliance all-compliance-info`
- REST API:
  - GET `/dcl/compliance/compliance-info`

### GET_ALL_DEVICE_SOFTWARE_COMPLIANCES

**Status: Implemented**

Gets all stored device software compliances.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query compliance all-device-software-compliance`
- REST API:
  - `/dcl/compliance/device-software-compliance`

## X509 PKI

**NOTE**: X.509 v3 certificates are only supported (all certificates MUST contain `Subject Key ID` field).
All PKI related methods are based on this restriction.

### Device Attestation Certificates (DA): PAA, PAI

#### PROPOSE_ADD_PAA

**Status: Implemented**

Proposes a new PAA (self-signed root certificate).

If more than 1 Trustee signature is required to add the PAA certificate, the PAA certificate
will be in a pending state until sufficient number of approvals is received.

The PAA certificate is immutable. It can only be revoked by either the owner or a quorum of Trustees.

- Who can send:
  - Trustee
- Parameters:
  - cert: `string` - PEM encoded certificate. The corresponding CLI parameter can contain either a PEM string or a path to a file containing the data.
  - info: `optional(string)` - information/notes for the proposal. Can contain up to 4096 characters.
  - time: `optional(int64)` - proposal time (number of nanoseconds elapsed since January 1, 1970 UTC). This field cannot be specified using a CLI command and will use the current time by default.
  - vid: `uint16` -  Vendor ID (positive non-zero). Must be equal to the Certificate's `vid` field for VID-scoped PAA.
  - schemaVersion: `optional(uint16)` - Certificate's schema version to support backward/forward compatability. Should be equal to 0 (default 0)
- In State: `pki/ProposedCertificate/value/<Certificate's Subject>/<Certificate's Subject Key ID>`
- CLI command:
  - `dcld tx pki propose-add-x509-root-cert --certificate=<string-or-path> --from=<account>`
- Validation:
  - provided certificate must be root:
    - `Issuer` == `Subject`
    - `Authority Key Identifier` == `Subject Key Identifier`
  - no existing `Proposed` certificate with the same `<Certificate's Subject>:<Certificate's Subject Key ID>` combination.
  - no existing certificate with the same `<Certificate's Issuer>:<Certificate's Serial Number>` combination.
  - if approved certificates with the same `<Certificate's Subject>:<Certificate's Subject Key ID>` combination already exists:
    - the existing certificate must not be NOC certificate
    - sender must match to the owner of the existing certificates.
  - the signature (self-signature) and expiration date are valid.

#### APPROVE_ADD_PAA

**Status: Implemented**

Approves the proposed PAA (self-signed root certificate). It also can be used for revote (i.e. change vote from reject to approve)

The PAA certificate is not active until sufficient number of Trustees approve it.

- Who can send:
  - Trustee
- Parameters:
  - subject: `string`  - proposed certificates's `Subject` is base64 encoded subject DER sequence bytes.
  - subject_key_id: `string`  - proposed certificates's `Subject Key Id` in hex string format, e.g: `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`.
  - info: `optional(string)` - information/notes for the approval. Can contain up to 4096 characters.
  - time: `optional(int64)` - proposal time (number of nanoseconds elapsed since January 1, 1970 UTC). This field cannot be specified using a CLI command and will use the current time by default.
- In State: `pki/ApprovedCertificates/value/<Certificate's Subject>/<Certificate's Subject Key ID>`.
- Number of required approvals:
  - greater than or equal 2/3 of Trustees (proposal by a Trustee is also counted as an approval)
- CLI command:
  - `dcld tx pki approve-add-x509-root-cert --subject=<base64 string> --subject-key-id=<hex string> --from=<account>`
- Validation:
  - the proposal to add a root certificate with the provided subject and subject_key_id, must be submitted first.
  - the proposed certificate hasn't been approved by the signer yet.

#### REJECT_ADD_PAA

**Status: Implemented**

Rejects the proposed PAA (self-signed root certificate). It also can be used for revote (i.e. change vote from approve to reject)

If proposed PAA certificate has only proposer's approval and no rejects then proposer can send this transaction to remove the proposal

The certificate is not reject until sufficient number of Trustees reject it.

- Who can send:
  - Trustee
- Parameters:
  - subject: `string` - proposed certificates's `Subject` is base64 encoded subject DER sequence bytes
  - subject_key_id: `string` - proposed certificates's `Subject Key Id` in hex string format, e.g:
  `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`
  - info: `optional(string)` - information/notes for the reject. Can contain up to 4096 characters.
  - time: `optional(int64)` - reject time (number of nanoseconds elapsed since January 1, 1970 UTC). This field cannot be specified using a CLI command and will use the current time by default.
- In State: `pki/RejectedCertificates/value/<Certificate's Subject>/<Certificate's Subject Key ID>`
- Number of required rejects:
  - more than 1/3 of Trustees
- CLI command:
  - `dcld tx pki reject-add-x509-root-cert --subject=<base64 string> --subject-key-id=<hex string> --from=<account>`
- Validation:
  - the proposal to add a root certificate with the provided subject and subject_key_id, must be submitted first.
  - the proposed certificate hasn't been rejected by the signer yet

#### PROPOSE_REVOKE_PAA

**Status: Implemented**

Proposes revocation of the given PAA (self-signed root certificate) by a Trustee.

Revocation works as a soft-delete, meaning that the certificates are not entirely removed but moved from the approved list to the revoked list.
Revoked certificates can be retrieved by using the [GET_REVOKED_CERT](#get_revoked_cert) query.

If a Revocation Distribution Point needs to be published (such as RFC5280 Certificate Revocation List), please use [ADD_REVOCATION_DISTRIBUTION_POINT](#add_revocation_distribution_point).

If `revoke-child` flag is set to `true` then all the certificates in the chain signed by the revoked certificate will be revoked as well.

If more than 1 Trustee signature is required to revoke a PAA certificate,
then the certificate will be in a pending state until sufficient number of other Trustee's approvals is received.

- Who can send:
  - Trustee
- Parameters:
  - subject: `string`  - certificates's `Subject` is base64 encoded subject DER sequence bytes.
  - subject_key_id: `string`  - certificates's `Subject Key Id` in hex string format, e.g: `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`.
  - serial-number: `optional(string)`  - certificate's serial number.
  - revoke-child: `optional(bool)`  - to revoke child certificates in the chain - default is false.
  - info: `optional(string)` - information/notes for the revocation proposal. Can contain up to 4096 characters.
  - time: `optional(int64)` - revocation proposal time (number of nanoseconds elapsed since January 1, 1970 UTC). This field cannot be specified using a CLI command and will use the current time by default.
- In State: `pki/ProposedCertificateRevocation/value/<Certificate's Subject>/<Certificate's Subject Key ID>`
- CLI command:
  - `dcld tx pki propose-revoke-x509-root-cert --subject=<base64 string> --subject-key-id=<hex string> --from=<account>`
- Validation:
  - revoked certificate must be root:
    - `Issuer` == `Subject`
    - `Authority Key Identifier` == `Subject Key Identifier`
  - no existing `Proposed` certificate with the same `<Certificate's Subject>:<Certificate's Subject Key ID>` combination.

#### APPROVE_REVOKE_PAA

**Status: Implemented**

Approves the revocation of the given PAA (self-signed root certificate) by a Trustee.

Revocation works as a soft-delete, meaning that the certificates are not entirely removed but moved from the approved list to the revoked list.
Revoked certificates can be retrieved by using the [GET_REVOKED_CERT](#get_revoked_cert) query.

If a Revocation Distribution Point needs to be published (such as RFC5280 Certificate Revocation List), please use [ADD_REVOCATION_DISTRIBUTION_POINT](#add_revocation_distribution_point).

The revocation is not applied until sufficient number of Trustees approve it.

- Who can send:
  - Trustee
- Parameters:
  - subject: `string`  - certificates's `Subject` is base64 encoded subject DER sequence bytes.
  - subject_key_id: `string`  - certificates's `Subject Key Id` in hex string format, e.g: `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`.
  - serial-number: `optional(string)` - certificate's serial number.
  - info: `optional(string)` - information/notes for the revocation approval. Can contain up to 4096 characters.
  - time: `optional(int64)` - revocation approval time (number of nanoseconds elapsed since January 1, 1970 UTC). This field cannot be specified using a CLI command and will use the current time by default.
- In State: `pki/RevokedCertificates/value/<Certificate's Subject>/<Certificate's Subject Key ID>`
- Number of required approvals:
  - greater than or equal 2/3 of Trustees (proposal by a Trustee is also counted as an approval)
- CLI command:
  - `dcld tx pki approve-revoke-x509-root-cert --subject=<base64 string> --subject-key-id=<hex string> --from=<account>`
- Validation:
  - the proposal to revoke a root certificate with the provided subject and subject_key_id, must be submitted first.
  - the proposed certificate revocation hasn't been approved by the signer yet.

#### ASSIGN_VID_TO_PAA

**Status: Implemented**

Assigns a Vendor ID (VID) to non-VID scoped PAAs (self-signed root certificate) already present on the ledger.

- Who can send:
  - Vendor Admin
- Parameters:
  - subject: `string` - certificates's `Subject` is base64 encoded subject DER sequence bytes.
  - subject_key_id: `string` - certificates's `Subject Key Id` in hex string format, e.g: `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`.
  - vid: `uint16` - Vendor ID (positive non-zero). Must be the same as `vid` field in the VID-scoped PAA certificate.
- CLI command:
  - `dcld tx pki assign-vid --subject=<base64 string> --subject-key-id=<hex string> --vid=<uint16> --from=<account>`
- Validation:
  - PAA Certificate with the provided `subject` and `subject_key_id` must exist in the ledger.
  - If the PAA is a VID scoped one, then the `vid` field must be equal to the VID value in the PAA's subject.

#### ADD_REVOCATION_DISTRIBUTION_POINT

**Status: Implemented**

Publishes a PKI Revocation distribution endpoint (such as RFC5280 Certificate Revocation List) owned by the Vendor.

If `crlSignerCertificate` is a PAA (root certificate), then it must be present on DCL.

If `crlSignerCertificate` is a PAI (intermediate certificate) or delegated by PAA, then it must be chained back to a valid PAA (root certificate) present on DCL.
In this case `crlSignerCertificate` is not required to be present on DCL, and will not be added to DCL as a result of this transaction.
If PAI needs to be added to DCL, it should be done via [ADD_PAI](#add_pai) transaction.

Publishing the revocation distribution endpoint doesn't automatically remove PAI (Intermediate certificates)
and DACs (leaf certificates) added to DCL if they are revoked in the CRL identified by this distribution point.
[REVOKE_PAI](#revoke_pai) needs to be called to remove an intermediate or leaf certificate from the ledger.∂

- Who can send: Vendor account
  - `vid` field in the transaction (`VendorID`) must be equal to the Vendor account's VID
  - VID-scoped PAAs (Root certs) and PAIs (Intermediate certs): `vid` field in the `CRLSignerCertificate`'s subject must be equal to the Vendor account's VID
  - Non-VID scoped PAAs (Root certs): `vid` field associated with the corresponding PAA on the ledger must be equal to the Vendor account's VID
- Parameters:
  - vid: `uint16` -  Vendor ID (positive non-zero). Must be the same as Vendor account's VID and `vid` field in the VID-scoped `CRLSignerCertificate`. Must be the same as a `vid` associated with non-VID scoped `CRLSignerCertificate` on the ledger.
  - pid: `optional(uint16)` -  Product ID (positive non-zero). Must be empty if `IsPAA` is true. Must be equal to a `pid` field in `CRLSignerCertificate`.
  - isPAA: `bool` -  True if the revocation information distribution point relates to a PAA
  - label: `string` -  A label to disambiguate multiple revocation information partitions of a particular issuer.
  - crlSignerCertificate: `string` - The issuer certificate whose revocation information is provided in the distribution point entry, encoded in X.509v3 PEM format. The corresponding CLI parameter can contain either a PEM string or a path to a file containing the data. Please note that if crlSignerCertificate is a delegated certificate by a PAI, the delegator certificate must be provided using the `crlSignerDelegator` field.
  - crlSignerDelegator: `optional(string)` - If crlSignerCertificate is a delegated certificate by a PAI, then crlSignerDelegator must contain the delegator PAI certificate which must be chained back to an approved certificate in the ledger, encoded in X.509v3 PEM format. Otherwise this field can be omitted. The corresponding CLI parameter can contain either a PEM string or a path to a file containing the data.
  - issuerSubjectKeyID: `string` - Uniquely identifies the PAA or PAI for which this revocation distribution point is provided. Must consist of even number of uppercase hexadecimal characters ([0-9A-F]), with no whitespace and no non-hexadecimal characters., e.g: `5A880E6C3653D07FB08971A3F473790930E62BDB`.
  - dataUrl: `string` -  The URL where to obtain the information in the format indicated by the RevocationType field. Must start with either `http` or `https`. Must be unique for all pairs of VendorID and IssuerSubjectKeyID.
  - dataFileSize: `optional(uint64)` -  Total size in bytes of the file found at the DataUrl. Must be omitted if RevocationType is 1.
  - dataDigest: `optional(string)` -  Digest of the entire contents of the associated file downloaded from the DataUrl. Must be omitted if RevocationType is 1. Must be provided if and only if the `DataFileSize` field is present.
  - dataDigestType: `optional(uint32)` - The type of digest used in the DataDigest field from the list of [1, 7, 8, 10, 11, 12] (IANA Named Information Hash Algorithm Registry). Must be provided if and only if the `DataDigest` field is present.
  - revocationType: `uint32` - The type of file found at the DataUrl for this entry. Supported types: 1 - RFC5280 Certificate Revocation List (CRL).
  - schemaVersion: `optional(uint16)` - Schema version to support backward/forward compatibility. Should be equal to 0 (default 0)
- In State:
  - `pki/RevocationDistributionPoint/value/<IssuerSubjectKeyID>` -> list of Revocation Distribution Points
  - `pki/RevocationDistributionPoint/value/<IssuerSubjectKeyID>/<vid>/<label>`-> Revocation Distribution Point
- CLI command:
  - `dcld tx pki add-revocation-point --vid=<uint16> --pid=<uint16> --issuer-subject-key-id=<string> --is-paa=<bool> --label=<string>
    --certificate=<string-or-path> --certificate-delegator=<string-or-path> --data-url=<string> --revocation-type=1 --from=<account>`

#### UPDATE_REVOCATION_DISTRIBUTION_POINT

**Status: Implemented**

Updates an existing PKI Revocation distribution endpoint (such as RFC5280 Certificate Revocation List) owned by the Vendor.

- Who can send: Vendor account
  - `vid` field in the transaction (`VendorID`) must be equal to the Vendor account's VID
  - VID-scoped PAAs (Root certs) and PAIs (Intermediate certs): `vid` field in the `CRLSignerCertificate`'s subject must be equal to the Vendor account's VID
  - Non-VID scoped PAAs (Root certs): `vid` field associated with the corresponding PAA on the ledger must be equal to the Vendor account's VID
- Parameters:
  - vid: `uint16` -  Vendor ID (positive non-zero). Must be the same as Vendor account's VID and `vid` field in the VID-scoped `CRLSignerCertificate`. Must be the same as a `vid` associated with non-VID scoped `CRLSignerCertificate` on the ledger.
  - label: `string` -  A label to disambiguate multiple revocation information partitions of a particular issuer.
  - issuerSubjectKeyID: `string` - Uniquely identifies the PAA or PAI for which this revocation distribution point is provided. Must consist of even number of uppercase hexadecimal characters ([0-9A-F]), with no whitespace and no non-hexadecimal characters., e.g: `5A880E6C3653D07FB08971A3F473790930E62BDB`.
  - crlSignerCertificate: `optional(string)` - The issuer certificate whose revocation information is provided in the distribution point entry, encoded in X.509v3 PEM format. The corresponding CLI parameter can contain either a PEM string or a path to a file containing the data. Please note that if crlSignerCertificate is a delegated certificate by a PAI, the delegator certificate must be provided using the `crlSignerDelegator` field.
  - crlSignerDelegator: `optional(string)` - If crlSignerCertificate is a delegated certificate by a PAI, then crlSignerDelegator must contain the delegator PAI certificate which must be chained back to an approved certificate in the ledger, encoded in X.509v3 PEM format. Otherwise this field can be omitted. The corresponding CLI parameter can contain either a PEM string or a path to a file containing the data.
  - dataUrl: `optional(string)` -  The URL where to obtain the information in the format indicated by the RevocationType field. Must start with either `http` or `https`. Must be unique for all pairs of VendorID and IssuerSubjectKeyID.
  - dataFileSize: `optional(uint64)` -  Total size in bytes of the file found at the DataUrl. Must be omitted if RevocationType is 1.
  - dataDigest: `optional(string)` -  Digest of the entire contents of the associated file downloaded from the DataUrl. Must be omitted if RevocationType is 1. Must be provided if and only if the `DataFileSize` field is present.
  - dataDigestType: `optional(uint32)` - The type of digest used in the DataDigest field from the list of [1, 7, 8, 10, 11, 12] (IANA Named Information Hash Algorithm Registry). Must be provided if and only if the `DataDigest` field is present.
  - schemaVersion: `optional(uint16)` - Schema version to support backward/forward compatibility. Should be equal to 0 (default 0)
- In State:
  - `pki/RevocationDistributionPoint/value/<IssuerSubjectKeyID>` -> list of Revocation Distribution Points
  - `pki/RevocationDistributionPoint/value/<IssuerSubjectKeyID>/<vid>/<label>` -> Revocation Distribution Point
- CLI command:
  - `dcld tx pki update-revocation-point --vid=<uint16> --issuer-subject-key-id=<string> --label=<string>
    --data-url=<string> --certificate=<string-or-path> --certificate-delegator=<string-or-path> --from=<account>`

#### DELETE_REVOCATION_DISTRIBUTION_POINT

**Status: Implemented**

Deletes a PKI Revocation distribution endpoint (such as RFC5280 Certificate Revocation List)  owned by the Vendor.

- Who can send: Vendor account
  - `vid` field in the transaction (`VendorID`) must be equal to the Vendor account's VID
  - VID-scoped PAAs (Root certs) and PAIs (Intermediate certs): `vid` field in the `CRLSignerCertificate`'s subject must be equal to the Vendor account's VID
  - Non-VID scoped PAAs (Root certs): `vid` field associated with the corresponding PAA on the ledger must be equal to the Vendor account's VID
- Parameters:
  - vid: `uint16` -  Vendor ID (positive non-zero). Must be the same as Vendor account's VID and `vid` field in the VID-scoped `CRLSignerCertificate`. Must be the same as a `vid` associated with non-VID scoped `CRLSignerCertificate` on the ledger.
  - label: `string` -  A label to disambiguate multiple revocation information partitions of a particular issuer.
  - issuerSubjectKeyID: `string` - Uniquely identifies the PAA or PAI for which this revocation distribution point is provided. Must consist of even number of uppercase hexadecimal characters ([0-9A-F]), with no whitespace and no non-hexadecimal characters., e.g: `5A880E6C3653D07FB08971A3F473790930E62BDB`.
- In State:
  - `pki/RevocationDistributionPoint/value/<IssuerSubjectKeyID>` -> list of Revocation Distribution Points
  - `pki/RevocationDistributionPoint/value/<IssuerSubjectKeyID>/<vid>/<label>` -> Revocation Distribution Point
- CLI command:
  - `dcld tx pki delete-revocation-point --vid=<uint16> --issuer-subject-key-id=<string> --label=<string> --from=<account>`

#### ADD_PAI

**Status: Implemented**

Adds a PAI (intermediate certificate) signed by a chain of certificates which must be already present on the ledger.

- Who can send:
  - Vendor Account
- Parameters:
  - cert: `string` - PEM encoded certificate. The corresponding CLI parameter can contain either a PEM string or a path to a file containing the data.
  - certificate-schema-version: `optional(uint16)` - Certificate's schema version to support backward/forward compatability(default 0)
- In State:
  - `pki/ApprovedCertificates/value/<Certificate's Subject>/<Certificate's Subject Key ID>`
  - `pki/ChildCertificates/value/<Certificate's Subject>/<Certificate's Subject Key ID>`
- CLI command:
  - `dcld tx pki add-x509-cert --certificate=<string-or-path> --from=<account>`
- Validation:
  - provided certificate must not be root:
    - `Issuer` != `Subject`
    - `Authority Key Identifier` != `Subject Key Identifier`
  - no existing certificate with the same `<Certificate's Issuer>:<Certificate's Serial Number>` combination.
  - if certificates with the same `<Certificate's Subject>:<Certificate's Subject Key ID>` combination already exist:
    - the existing certificate must not be NOC certificate.
    - the sender's VID must match the VID of the existing certificate's owner.
  - the signature and expiration date are valid.
  - parent certificate must be already stored on the ledger and a valid chain to some root certificate can be built.
  - if the parent root certificate is VID scoped:
    - the provided certificate must also be VID scoped.
    - the `vid` in the subject of the root certificate must be equal to the `vid` in the subject of the provided certificate.
    - the `vid` in the subjects of both certificates must be equal to the sender Vendor account's VID.
  - if the parent root certificate is not VID scoped but has an associated VID:
    - the provided certificate can be either VID scoped or non-VID scoped.
    - if the provided certificate is VID scoped, the `vid` in the subject of the certificate must be equal to the VID associated with the root certificate and to the sender Vendor account's VID.
  - if the parent root certificate is non-VID scoped and does not have an associated VID, an error will occur.

> **_Note:_**  Multiple certificates can refer to the same `<Certificate's Subject>:<Certificate's Subject Key ID>` combination.

#### REVOKE_PAI

**Status: Implemented**

Revokes the given PAI (intermediate certificate).

Revocation works as a soft-delete, meaning that the certificates are not entirely removed but moved from the approved list to the revoked list.
Revoked certificates can be retrieved by using the [GET_REVOKED_CERT](#get_revoked_cert) query.
To entirely remove a PAI certificate, please use [REMOVE_PAI](#remove_pai).

If a Revocation Distribution Point needs to be published (such as RFC5280 Certificate Revocation List), please use [ADD_REVOCATION_DISTRIBUTION_POINT](#add_revocation_distribution_point).

If `revoke-child` flag is set to `true` then all the certificates in the chain signed by the revoked certificate will be revoked as well.

Root certificates can not be revoked this way, use  [PROPOSE_REVOKE_PAA](#propose_revoke_paa) and [APPROVE_REVOKE_PAA](#approve_revoke_paa) instead.  

- Who can send: Vendor account
  - the sender's VID must match the VID of the revoking certificate's owner.
- Parameters:
  - subject: `string`  - certificates's `Subject` is base64 encoded subject DER sequence bytes.
  - subject_key_id: `string`  - certificates's `Subject Key Id` in hex string format, e.g: `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`.
  - serial-number: `optional(string)` - certificate's serial number.
  - revoke-child: `optional(bool)` - to revoke child certificates in the chain - default is false.
  - info: `optional(string)` - information/notes for the revocation. Can contain up to 4096 characters.
  - time: `optional(int64)` - revocation time (number of nanoseconds elapsed since January 1, 1970 UTC). This field cannot be specified using a CLI command and will use the current time by default.
- In State: `pki/RevokedCertificates/value/<Certificate's Subject>/<Certificate's Subject Key ID>`
- CLI command:
  - `dcld tx pki revoke-x509-cert --subject=<base64 string> --subject-key-id=<hex string> --from=<account>`
- Validation:
  - a PAI Certificate with the provided `subject` and `subject_key_id` must exist in the ledger.

#### REMOVE_PAI

**Status: Implemented**

This transaction completely removes the given PAI (intermediate certificate) from both the approved and revoked certificates list.

PAA (self-signed root certificate) can not be removed this way.  

- Who can send: Vendor account
  - the sender's VID must match the VID of the removing certificate's owner.
- Parameters:
  - subject: `string`  - certificates's `Subject` is base64 encoded subject DER sequence bytes
  - subject_key_id: `string`  - certificates's `Subject Key Id` in hex string format, e.g: `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`
  - serial-number: `optional(string)`  - certificate's serial number.
- CLI command:
  - `dcld tx pki remove-x509-cert --subject=<base64 string> --subject-key-id=<hex string> --from=<account>`
- Validation:
  - a PAI Certificate with the provided `subject` and `subject_key_id` must exist in the ledger.

#### GET_PKI_REVOCATION_DISTRIBUTION_POINT

**Status: Implemented**

Gets a revocation distribution point (such as RFC5280 Certificate Revocation List) identified by (VendorID, Label, IssuerSubjectKeyID) unique combination.
Use [GET_ALL_PKI_REVOCATION_DISTRIBUTION_POINT](#get_all_pki_revocation_distribution_point) to get a list of all revocation distribution points.

- Parameters:
  - vid: `uint16` -  Vendor ID (positive non-zero)
  - label: `string` -  A label to disambiguate multiple revocation information partitions of a particular issuer.
  - issuerSubjectKeyID: `string` - Uniquely identifies the PAA or PAI for which this revocation distribution point is provided. Must consist of even number of uppercase hexadecimal characters ([0-9A-F]), with no whitespace and no non-hexadecimal characters., e.g: `5A880E6C3653D07FB08971A3F473790930E62BDB`.
- CLI command:
  - `dcld query pki revocation-point --vid=<uint16> --label=<string> --issuer-subject-key-id=<string>`
- REST API:
  - GET `/dcl/pki/revocation-points/{issuerSubjectKeyID}/{vid}/{label}`

#### GET_PKI_REVOCATION_DISTRIBUTION_POINTS_BY_SUBJECT_KEY_ID

**Status: Implemented**

Gets a list of revocation distribution point (such as RFC5280 Certificate Revocation List) identified by IssuerSubjectKeyID.

- Parameters:
  - issuerSubjectKeyID: `string` - Uniquely identifies the PAA or PAI for which this revocation distribution point is provided. Must consist of even number of uppercase hexadecimal characters ([0-9A-F]), with no whitespace and no non-hexadecimal characters., e.g: `5A880E6C3653D07FB08971A3F473790930E62BDB`.
- CLI command:
  - `dcld query pki revocation-points --issuer-subject-key-id=<string>`
- REST API:
  - GET `/dcl/pki/revocation-points/{issuerSubjectKeyID}`

#### GET_ALL_PKI_REVOCATION_DISTRIBUTION_POINT

**Status: Implemented**

Gets a list of all revocation distribution points (such as RFC5280 Certificate Revocation List).

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters
- CLI command:
  - `dcld query pki all-revocation-points`
- REST API:
  - GET `/dcl/pki/revocation-points`

#### GET_PROPOSED_PAA

**Status: Implemented**

Gets a proposed but not approved PAA certificate with the given subject and subject key ID attributes.

- Parameters:
  - subject: `string`  - certificates's `Subject` is base64 encoded subject DER sequence bytes
  - subject_key_id: `string`  - certificates's `Subject Key Id` in hex string format, e.g: `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`
- CLI command:
  - `dcld query pki proposed-x509-root-cert --subject=<base64 string> --subject-key-id=<hex string>`
- REST API:
  - GET `/dcl/pki/proposed-certificates/{subject}/{subject_key_id}`

#### GET_REJECTED_PAA

**Status: Implemented**

Get a rejected PAA certificate with the given subject and subject key ID attributes.

- Parameters:
  - subject: `string` - certificates's `Subject` is base64 encoded subject DER sequence bytes
  - subject_key_id: `string`  - certificates's `Subject Key Id` in hex string format, e.g: `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`
- CLI command:
  - `dcld query pki rejected-x509-root-cert --subject=<base64 string> --subject-key-id=<hex string>`
- REST API:
  - GET `/dcl/pki/rejected-certificates/{subject}/{subject_key_id}`

#### GET_PROPOSED_PAA_TO_REVOKE

**Status: Implemented**

Gets a proposed but not approved PAA certificate to be revoked.

Revocation here just means removing it from the ledger.
If a Revocation Distribution Point (such as RFC5280 Certificate Revocation List) published to the ledger needs to be queried, please use [GET_PKI_REVOCATION_DISTRIBUTION_POINT](#get_pki_revocation_distribution_point).

- Parameters:
  - subject: `string`  - certificates's `Subject` is base64 encoded subject DER sequence bytes
  - subject_key_id: `string`  - certificates's `Subject Key Id` in hex string format, e.g: `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`
  - serial-number: `optional(string)`  - certificate's serial number
- CLI command:
  - `dcld query pki proposed-x509-root-cert-to-revoke --subject=<base64 string> --subject-key-id=<hex string>`
- REST API:
  - GET `/dcl/pki/proposed-revocation-certificates/{subject}/{subject_key_id}?serialnumber={serialnumber}`

#### GET_ALL_PAA

**Status: Implemented**

Gets all approved PAA certificates. Revoked certificates are not returned.
Use [GET_ALL_REVOKED_PAA](#get_all_revoked_paa) to get a list of all revoked PAA certificates.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query pki all-x509-root-certs`
- REST API:
  - GET `/dcl/pki/root-certificates`

#### GET_ALL_REVOKED_PAA

**Status: Implemented**

Gets all revoked PAA certificates.

Revocation here just means removing it from the ledger.
If a Revocation Distribution Point (such as RFC5280 Certificate Revocation List) published to the ledger needs to be queried, please use [GET_PKI_REVOCATION_DISTRIBUTION_POINT](#get_pki_revocation_distribution_point).

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query pki all-revoked-x509-root-certs`
- REST API:
  - GET `/dcl/pki/revoked-root-certificates`

#### GET_ALL_PROPOSED_PAA

**Status: Implemented**

Gets all proposed but not approved root certificates.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query pki all-proposed-x509-root-certs`
- REST API:
  - GET `dcl/pki/proposed-certificates`

#### GET_ALL_REJECTED_PAA

 **Status: Implemented**

Gets all rejected root certificates.

Shoudl be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query pki all-rejected-x509-root-certs`
- REST API:
  - GET `dcl/pki/rejected-certificates`

#### GET_ALL_PROPOSED_PAA_TO_REVOKE

**Status: Implemented**

Gets all proposed but not approved root certificates to be revoked.

Revocation here just means removing it from the ledger.
If a Revocation Distribution Point (such as RFC5280 Certificate Revocation List) published to the ledger needs to be queried, please use [GET_PKI_REVOCATION_DISTRIBUTION_POINT](#get_pki_revocation_distribution_point).

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query pki all-proposed-x509-root-certs-to-revoke`
- REST API:
  - GET `/dcl/pki/proposed-revocation-certificates`

### E2E (NOC): RCAC, ICAC

#### ADD_NOC_ROOT (RCAC)

**Status: Implemented**

This transaction adds a NOC root certificate (RCAC) owned by the Vendor.

- Who can send
  - Vendor account
- Parameters:
  - cert: `string` - The NOC Root Certificate (RCAC), encoded in X.509v3 PEM format. Can be a PEM string or a file path.
  - schemaVersion: `optional(uint16)` - Certificate's schema version to support backward/forward compatability. Should be equal to 0 (default 0)
- In State:
  - `pki/ApprovedCertificates/value/<Subject>/<SubjectKeyID>`
  - `pki/ApprovedCertificatesBySubject/value/<Subject>`
  - `pki/ApprovedCertificatesBySubjectKeyID/value/<SubjectKeyID>`
  - `pki/NocRootCertificates/value/<VID>`
- CLI Command:
  - `dcld tx pki add-noc-x509-root-cert --certificate=<string-or-path> --from=<account>`
- Validation:
  - the provided certificate must be a root certificate (RCAC):
    - `Issuer` == `Subject`
    - `Authority Key Identifier` == `Subject Key Identifier`
  - no existing certificate with the same `<Certificate's Issuer>:<Certificate's Serial Number>` combination.
  - if certificates with the same `<Certificate's Subject>:<Certificate's Subject Key ID>` combination already exist:
    - the existing certificate must be NOC root certificate (RCAC)
    - the sender's VID must match the `vid` field of the existing certificates.
  - the signature (self-signature) and expiration date must be valid.

#### REVOKE_NOC_ROOT (RCAC)

**Status: Implemented**

This transaction revokes a NOC root certificate (RCAC) owned by the Vendor.
Revoked NOC root certificates (RCACs) can be re-added using the [ADD_NOC_ROOT](#add_noc_root-(rcac)) transaction.

Revocation works as a soft-delete, meaning that the certificates are not entirely removed but moved from the approved list to the revoked list.
Revoked certificates can be retrieved by using the [GET_REVOKED_CERT](#get_revoked_cert) query.

- Who can send: Vendor account
  - Vid field associated with the corresponding NOC root certificate (RCAC) on the ledger must be equal to the Vendor account's VID.
- Parameters:
  - subject: `string` - base64 encoded subject DER sequence bytes of the certificate.
  - subject_key_id: `string` - certificate's `Subject Key Id` in hex string format, e.g., `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`.
  - serial_number: `optional(string)` - certificate's serial number. If not provided, the transaction will revoke all certificates that match the given `subject` and `subject_key_id` combination.
  - revoke-child: `optional(bool)` - if true, then all certificates in the chain signed by the revoked certificate (intermediate, leaf) are revoked as well. If false, only the current root cert is revoked (default: false).
  - info: `optional(string)` - information/notes for the revocation. Can contain up to 4096 characters.
  - time: `optional(int64)` - revocation time (number of nanoseconds elapsed since January 1, 1970 UTC). This field cannot be specified using a CLI command and will use the current time by default.
- In State:
  - `pki/RevokedCertificates/value/<Certificate's Subject>/<Certificate's Subject Key ID>`
  - `pki/RevokedNocRootCertificates/value/<Certificate's Subject>/<Certificate's Subject Key ID>`
- CLI command:
  - `dcld tx pki revoke-noc-x509-root-cert --subject=<base64 string> --subject-key-id=<hex string> --serial-number=<string> --info=<string> --time=<int64> --revoke-child=<bool> --from=<account>`
- Validation:
  - a NOC Root Certificate (RCAC) with the provided `subject` and `subject_key_id` must exist in the ledger.

#### REMOVE_NOC_ROOT (RCAC)

**Status: Implemented**

This transaction completely removes the given NOC root certificate (RCAC) owned by the Vendor from the ledger.
Removed NOC root certificates (RCACs) can be re-added using the [ADD_NOC_ROOT](#add_noc_root-(rcac)) transaction.

- Who can send: Vendor account
  - Vid field associated with the corresponding NOC certificate on the ledger must be equal to the Vendor account's VID.
- Validation:
  - a NOC Root Certificate (RCAC) with the provided `subject` and `subject_key_id` must exist in the ledger.
- Parameters:
  - subject: `string` - base64 encoded subject DER sequence bytes of the certificate.
  - subject_key_id: `string` - certificate's `Subject Key Id` in hex string format, e.g., `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`.
  - serial_number: `optional(string)` - certificate's serial number. If not provided, the transaction will remove all certificates that match the given `subject` and `subject_key_id` combination.
- CLI command:
  - `dcld tx pki remove-noc-x509-root-cert --subject=<base64 string> --subject-key-id=<hex string> --from=<account>`



#### ADD_NOC_ICA (ICAC)

**Status: Implemented**

This transaction adds a NOC ICA certificate (ICAC) owned by the Vendor signed by a chain of certificates which must be
already present on the ledger.

- Who can send: Vendor account
- Validation:
  - the provided certificate must be a non-root certificate:
    - `Issuer` != `Subject`
    - `Authority Key Identifier` != `Subject Key Identifier`
  - the root certificate must be a NOC certificate and added by the same vendor
    - `isNoc` field of the root certificate must be set to true
    - `VID of root certificate` == `VID of account`
  - no existing certificate with the same `<Certificate's Issuer>:<Certificate's Serial Number>` combination.
  - if certificates with the same `<Certificate's Subject>:<Certificate's Subject Key ID>` combination already exist:
    - the existing certificate must be NOC non-root certificate
    - the sender's VID must match the vid field of the existing certificates.
  - the signature and expiration date must be valid.
- Parameters:
  - cert: `string` - The NOC non-root Certificate, encoded in X.509v3 PEM format. Can be a PEM string or a file path.
  - certificate-schema-version: `optional(uint16)` - Certificate's schema version to support backward/forward compatability(default 0)
- In State:
  - `pki/ApprovedCertificates/value/<Subject>/<SubjectKeyID>`
  - `pki/ApprovedCertificatesBySubject/value/<Subject>`
  - `pki/ApprovedCertificatesBySubjectKeyID/value/<SubjectKeyID>`
  - `pki/NocIcaCertificates/value/<VID>`
  - `pki/ChildCertificates/value/<Certificate's Subject>/<Certificate's Subject Key ID>`
- CLI Command:
  - `dcld tx pki add-noc-x509-ica-cert --certificate=<string-or-path> --from=<account>`

#### REVOKE_NOC_ICA (ICAC)

**Status: Implemented**

This transaction revokes a NOC ICA certificate (ICAC) owned by the Vendor.
Revoked NOC ICA certificates (ICACs) can be re-added using the [ADD_NOC_ICA](#add_noc_ica-(icac)) transaction.

Revocation works as a soft-delete, meaning that the certificates are not entirely removed but moved from the approved list to the revoked list.
Revoked certificates can be retrieved by using the [GET_REVOKED_CERT](#get_revoked_cert) query.

- Who can send: Vendor account
  - Vid field associated with the corresponding NOC certificate on the ledger must be equal to the Vendor account's VID.
- Validation:
  - a NOC Certificate with the provided `subject` and `subject_key_id` must exist in the ledger.
- Parameters:
  - subject: `string` - base64 encoded subject DER sequence bytes of the certificate.
  - subject_key_id: `string` - certificate's `Subject Key Id` in hex string format, e.g., `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`.
  - serial_number: `optional(string)` - certificate's serial number. If not provided, the transaction will revoke all certificates that match the given `subject` and `subject_key_id` combination.
  - revoke-child: `optional(bool)` - if true, then all certificates in the chain signed by the revoked certificate (leaf) are revoked as well. If false, only the current cert is revoked (default: false).
  - info: `optional(string)` - information/notes for the revocation. Can contain up to 4096 characters.
  - time: `optional(int64)` - revocation time (number of nanoseconds elapsed since January 1, 1970 UTC). This field cannot be specified using a CLI command and will use the current time by default.
- In State:
  - `pki/RevokedCertificates/value/<Certificate's Subject>/<Certificate's Subject Key ID>`
- CLI command:
  - `dcld tx pki revoke-noc-x509-ica-cert --subject=<base64 string> --subject-key-id=<hex string> --serial-number=<string> --info=<string> --time=<int64> --revoke-child=<bool> --from=<account>`

#### REMOVE_NOC_ICA (ICAC)

**Status: Implemented**

This transaction completely removes the given NOC ICA (ICAC) owned by the Vendor from the ledger.
Removed NOC ICA certificates (ICACs) can be re-added using the [ADD_NOC_ICA](#add_noc_ica-(icac)) transaction.

- Who can send: Vendor account
  - Vid field associated with the corresponding NOC certificate on the ledger must be equal to the Vendor account's VID.
- Validation:
  - a NOC ICA Certificate (ICAC) with the provided `subject` and `subject_key_id` must exist in the ledger.
- Parameters:
  - subject: `string` - base64 encoded subject DER sequence bytes of the certificate.
  - subject_key_id: `string` - certificate's `Subject Key Id` in hex string format, e.g., `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`.
  - serial_number: `optional(string)` - certificate's serial number. If not provided, the transaction will remove all certificates that match the given `subject` and `subject_key_id` combination.
- CLI command:
  - `dcld tx pki remove-noc-x509-ica-cert --subject=<base64 string> --subject-key-id=<hex string> --from=<account>`


#### GET_NOC_ROOT_BY_VID (RCACs)

**Status: Implemented**

Retrieve NOC root certificates (RCACs) associated with a specific VID.

Revoked NOC root certificates (RCACs) are not returned.
Use [GET_ALL_REVOKED_NOC_ROOT](#get_all_revoked_noc_root-(rcacs)) to get a list of all revoked NOC root certificates (RCACs).

- Who can send: Any account
- Parameters:
  - vid: `uint16` - Vendor ID (positive non-zero)
- CLI Command:
  - `dcld query pki noc-x509-root-certs --vid=<uint16>`
- REST API:
  - GET `/dcl/pki/noc-root-certificates/{vid}`

#### GET_NOC_BY_VID_AND_SKID (RCACs/ICACs)

**Status: Implemented**

Retrieve NOC (Root/ICA) certificates (RCACs/ICACs) associated with a specific VID and subject key ID.
This request also returns the Trust Quotient (TQ) value of the certificate

Revoked NOC certificates are not returned.
Use [GET_ALL_REVOKED_NOC_ROOT](#get_all_revoked_noc_root-(rcacs)) to get a list of all revoked NOC root certificates.
Use [GET_ALL_REVOKED_CERT](#get_all_revoked_certs) to get a list of all revoked certificates (including ICACs).

- Who can send: Any account
- Parameters:
  - vid: `uint16` - Vendor ID (positive non-zero)
  - subject_key_id: `string` - Certificate's `Subject Key Id` in hex string format, e.g., `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`
- CLI Command:
  - `dcld query pki noc-x509-certs --vid=<uint16> --subject-key-id=<hex string>`
- REST API:
  - GET `/dcl/pki/noc-certificates/{vid}/{subject_key_id}`

#### GET_NOC_ICA_BY_VID (ICACs)

**Status: Implemented**

Retrieve NOC ICA certificates (ICACs) associated with a specific VID.

Revoked certificates are not returned.
Use [GET_ALL_REVOKED_CERT](#get_all_revoked_certs) to get a list of all revoked certificates.

- Who can send: Any account
- Parameters:
  - vid: `uint16` - Vendor ID (positive non-zero)
- CLI Command:
  - `dcld query pki noc-x509-ica-certs --vid=<uint16>`
- REST API:
  - GET `/dcl/pki/noc-ica-certificates/{vid}`

#### GET_REVOKED_NOC_ROOT (RCAC)

**Status: Implemented**

Gets a revoked NOC root certificate (RCAC) by the given subject and subject key ID attributes.

Revocation works as a soft-delete, meaning that the certificates are not entirely removed but moved from the approved list to the revoked list.

- Parameters:
  - subject: `string` - Base64 encoded subject DER sequence bytes of the certificate.
  - subject_key_id: `string` - Certificate's `Subject Key Id` in hex string format, e.g., `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`.
- CLI command:
  - `dcld query pki revoked-noc-x509-root-cert --subject=<base64 string> --subject-key-id=<hex string>`
- REST API:
  - GET `/dcl/pki/revoked-noc-root-certificates/{subject}/{subject_key_id}`

#### GET_ALL_NOC_ROOT (RCACs)

**Status: Implemented**

Retrieve a list of all of NOC root certificates (RCACs).

Revoked NOC root certificates (RCACs) are not returned.
Use [GET_ALL_REVOKED_NOC_ROOT](#get_all_revoked_noc_root-(rcacs)) to get a list of all revoked NOC root certificates (RCACs).

- Who can send: Any account
- Parameters:
  - Common pagination parameters
- CLI Command:
  - `dcld query pki all-noc-x509-root-certs`
- REST API:
  - GET `/dcl/pki/noc-root-certificates`

#### GET_ALL_NOC_ICA (ICACs)

**Status: Implemented**

Retrieve a list of all of NOC ICA certificates (ICACs).

Revoked certificates are not returned.
Use [GET_ALL_REVOKED_CERT](#get_all_revoked_certs) to get a list of all revoked certificates.

- Who can send: Any account
- Parameters:
  - Common pagination parameters
- CLI Command:
  - `dcld query pki all-noc-x509-ica-certs`
- REST API:
  - GET `/dcl/pki/noc-ica-certificates`

#### GET_ALL_REVOKED_NOC_ROOT (RCACs)

Gets all revoked NOC root certificates (RCACs).

Revocation works as a soft-delete, meaning that the certificates are not entirely removed but moved from the approved list to the revoked list.

- Who can send: Any account
- Parameters:
  - Common pagination parameters
- CLI command:
  - `dcld query pki all-revoked-noc-x509-root-certs`
- REST API:
  - GET `/dcl/pki/revoked-noc-root-certificates`

### Common

#### GET_CERT

**Status: Implemented**

Gets a certificate by the given subject and subject key ID attributes. This query works for all types of certificates (PAA, PAI, NOC_ROOT, NOC_ICA).
Revoked certificates are not returned.
Use [GET_REVOKED_CERT](#get_revoked_cert) to get a revoked certificate.

- Parameters:
  - subject: `string`  - certificates's `Subject` is base64 encoded subject DER sequence bytes
  - subject_key_id: `string`  - certificates's `Subject Key Id` in hex string format, e.g: `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`
- CLI command:
  - `dcld query pki x509-cert --subject=<base64 string> --subject-key-id=<hex string>`
- REST API:
  - GET `/dcl/pki/certificates/{subject}/{subject_key_id}`

#### GET_REVOKED_CERT

**Status: Implemented**

Gets a revoked certificate by the given subject and subject key ID attributes. This query works for all types of certificates (PAA, PAI, NOC_ROOT, NOC_ICA).

Revocation works as a soft-delete, meaning that the certificates are not entirely removed but moved from the approved list to the revoked list.
If a Revocation Distribution Point (such as RFC5280 Certificate Revocation List) published to the ledger needs to be queried, please use [GET_PKI_REVOCATION_DISTRIBUTION_POINT](#get_pki_revocation_distribution_point).

- Parameters:
  - subject: `string`  - certificates's `Subject` is base64 encoded subject DER sequence bytes
  - subject_key_id: `string`  - certificates's `Subject Key Id` in hex string format, e.g: `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`
- CLI command:
  - `dcld query pki revoked-x509-cert --subject=<base64 string> --subject-key-id=<hex string>`
- REST API:
  - GET `/dcl/pki/revoked-certificates/{subject}/{subject_key_id}`

#### GET_CERTS_BY_SKID

**Status: Implemented**

Gets all certificates by the given subject key ID attribute. This query works for all types of certificates (PAA, PAI, NOC_ROOT, NOC_ICA).

Revoked certificates are not returned.
Use `GET_ALL_REVOKED_CERTS` to get a list of all revoked certificates.

- Parameters:
  - subject_key_id: `string`  - certificates's `Subject Key Id` in hex string format, e.g: `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`
- CLI command:
  - `dcld query pki x509-cert --subject-key-id=<hex string>`
- REST API:
  - GET `/dcl/pki/certificates?subjectKeyId={subjectKeyId}`

#### GET_CERTS_BY_SUBJECT

**Status: Implemented**

Gets all certificates associated with a subject. This query works for all types of certificates (PAA, PAI, NOC_ROOT, NOC_ICA).

Revoked certificates are not returned.
Use [GET_ALL_REVOKED_CERTS](#get_all_revoked_certs) to get a list of all revoked certificates.

- Parameters:
  - subject: `string`  - certificates's `Subject` is base64 encoded subject DER sequence bytes
- CLI command:
  - `dcld query pki all-subject-x509-certs --subject=<base64 string>`
- REST API:
  - GET `/dcl/pki/certificates/{subject}`

#### GET_CHILD_CERTS

**Status: Implemented**

Gets all child certificates for the given certificate. This query works for both PAI and NOC_ICA.
Revoked certificates are not returned.

- Parameters:
  - subject: `string`  - certificates's `Subject` is base64 encoded subject DER sequence bytes
  - subject_key_id: `string`  - certificates's `Subject Key Id` in hex string format, e.g: `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`
- CLI command:
  - `dcld query pki all-child-x509-certs (--subject=<base64 string> --subject-key-id=<hex string>`
- REST API:
  - GET `/dcl/pki/child-certificates/{subject}/{subject_key_id}`

#### GET_ALL_CERTS

**Status: Implemented**

Gets all certificates. This query works for all types of certificates (PAA, PAI, NOC_ROOT, NOC_ICA).

Revoked certificates are not returned.
Use [GET_ALL_REVOKED_CERTS](#get_all_revoked_certs) to get a list of all revoked certificates.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query pki all-x509-certs`
- REST API:
  - GET `/dcl/pki/certificates`

#### GET_ALL_REVOKED_CERTS

**Status: Implemented**

Gets all revoked certificates. This query works for all types of certificates (PAA, PAI, NOC_ROOT, NOC_ICA).

Revocation works as a soft-delete, meaning that the certificates are not entirely removed but moved from the approved list to the revoked list.
If a Revocation Distribution Point (such as RFC5280 Certificate Revocation List) published to the ledger needs to be queried, please use [GET_PKI_REVOCATION_DISTRIBUTION_POINT](#get_pki_revocation_distribution_point).

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query pki all-revoked-x509-certs`
- REST API:
  - GET `/dcl/pki/revoked-certificates`

## AUTH

### PROPOSE_ADD_ACCOUNT

**Status: Implemented**

Proposes a new Account with the given address, public key and role.

If more than 1 Trustee signature is required to add the account, the account
will be in a pending state until sufficient number of approvals is received.

- Parameters:
  - address: `string` - account address; Bech32 encoded
  - pub_key: `string` - account's Protobuf JSON encoded public key
  - vid: `optional(uint16)` - vendor ID (only needed for vendor role)
  - pid_ranges: `optional(array<uint16 range>)` - the list of product-id ranges (range item separated with "-"), comma-separated, in increasing order, associated with this account: `1-100,201-300...`
  - roles: `array<string>` - the list of roles, comma-separated, assigning to the account. Supported roles: `Vendor`, `TestHouse`, `CertificationCenter`, `Trustee`, `NodeAdmin`, `VendorAdmin`.
  - info: `optional(string)` - information/notes for the proposal
  - time: `optional(int64)` - proposal time (number of nanoseconds elapsed since January 1, 1970 UTC). This field cannot be specified using a CLI command and will use the current time by default.
- In State: `dclauth/PendingAccount/value/<address>`
- Who can send:
  - Trustee
- CLI command:
  - `dcld tx auth propose-add-account --address=<bench32 encoded string> --pubkey='<protobuf JSON encoded>' --roles=<role1,role2,...> --vid=<uint16> --pid_ranges=<uint16-range,uint16-range,...> --from=<account>`

### APPROVE_ADD_ACCOUNT

**Status: Implemented**

Approves the proposed account. It also can be used for revote (i.e. change vote from reject to approve)

The account is not active until sufficient number of Trustees approve it.

- Parameters:
  - address: `string` - account address; Bech32 encoded
  - info: `optional(string)` - information/notes for the approval
  - time: `optional(int64)` - approval time (number of nanoseconds elapsed since January 1, 1970 UTC). This field cannot be specified using a CLI command and will use the current time by default.
- In State: `dclauth/Account/value/<address>`
- Who can send:
  - Trustee
- Number of required approvals:
  - greater than or equal 2/3 of Trustees for account roles: `TestHouse`, `CertificationCenter`, `Trustee`, `NodeAdmin`, `VendorAdmin` (proposal by a Trustee is also counted as an approval)
  - greater than 1/3 of Trustees for account role: `Vendor` (proposal by a Trustee is also counted as an approval)
- CLI command:
  - `dcld tx auth approve-add-account --address=<bench32 encoded string> --from=<account>`

> **_Note:_**  If we are approving an account with role `Vendor`, then we need more than 1/3 of Trustees approvals.

### REJECT_ADD_ACCOUNT

**Status: Implemented**

Rejects the proposed account. It also can be used for revote (i.e. change vote from approve to reject)

If proposed account has only proposer's approval and no rejects then proposer can send this transaction to remove the proposal

The account is not reject until sufficient number of Trustees reject it.

- Parameters:
  - address: `string` - account address; Bech32 encoded
  - info: `optional(string)` - information/notes for the reject
  - time: `optional(int64)` - reject time (number of nanoseconds elapsed since January 1, 1970 UTC). This field cannot be specified using a CLI command and will use the current time by default.
- In State: `dclauth/RejectedAccount/value/<address>`
- Who can send:
  - Trustee
- Number of required rejects:
  - greater than 1/3 of Trustees for account roles: `TestHouse`, `CertificationCenter`, `Trustee`, `NodeAdmin`, `VendorAdmin` (proposal by a Trustee is also counted as an approval)
  - greater than or equal 2/3 of Trustees for account role: `Vendor` (proposal by a Trustee is also counted as an approval)
- CLI command:
  - `dcld tx auth reject-add-account --address=<bench32 encoded string> --from=<account>`

### PROPOSE_REVOKE_ACCOUNT

**Status: Implemented**

Proposes revocation of the Account with the given address.

If more than 1 Trustee signature is required to revoke the account, the revocation
will be in a pending state until sufficient number of approvals is received.

- Parameters:
  - address: `string` - account address; Bech32 encoded
  - info: `optional(string)` - information/notes for the revocation proposal
  - time: `optional(int64)` - revocation proposal time (number of nanoseconds elapsed since January 1, 1970 UTC). This field cannot be specified using a CLI command and will use the current time by default.
- In State: `dclauth/Account/value/<address>`
- Who can send:
  - Trustee
- CLI command:
  - `dcld tx auth propose-revoke-account --address=<bench32 encoded string> --from=<account>`

### APPROVE_REVOKE_ACCOUNT

**Status: Implemented**

Approves the proposed revocation of the account.

The account is not revoked until sufficient number of Trustees approve it.

- Parameters:
  - address: `string` - account address; Bech32 encoded
  - info: `optional(string)` - information/notes for the revocation approval
  - time: `optional(int64)` - revocation approval time (number of nanoseconds elapsed since January 1, 1970 UTC). This field cannot be specified using a CLI command and will use the current time by default.
- In State: `dclauth/Account/value/<address>`
- Who can send:
  - Trustee
- Number of required approvals:
  - greater than or equal 2/3 of Trustees (proposal by a Trustee is also counted as an approval)
- CLI command:
  - `dcld tx auth approve-revoke-account --address=<bench32 encoded string> --from=<account>`

> **_Note:_**  If revoking an account has sufficient number of Trustees approve it then this account is placed in Revoked Account.

### GET_ACCOUNT

**Status: Implemented**

Gets an accounts by the address. Revoked accounts are not returned.

- Parameters:
  - address: `string` - account address; Bech32 encoded
- CLI command:
  - `dcld query auth account --addres <bench32 encoded string>`
- REST API:
  - GET `/dcl/auth/accounts/{address}`

### GET_PROPOSED_ACCOUNT

**Status: Implemented**

Gets a proposed but not approved accounts by its address

- Parameters:
  - address: `string` - account address; Bech32 encoded
- CLI command:
  - `dcld query auth proposed-account --address <bench32 encoded string>`
- REST API:
  - GET `/dcl/auth/proposed-accounts/{address}`

### GET_REJECTED_ACCOUNT

**Status: Implemented**

Get a rejected accounts by its address

- Parameters:
  - address: `string` - account address; Bech32 encoded
- CLI command:
  - `dcld query auth rejected-account --address <bech32 encoded string>`
- REST API:
  - GET `/dcl/auth/rejected-accounts/{address}`

### GET_PROPOSED_ACCOUNT_TO_REVOKE

**Status: Implemented**

Gets a proposed but not approved accounts to be revoked by its address.

- Parameters:
  - address: `string` - account address; Bech32 encoded
- CLI command:
  - `dcld query auth proposed-account-to-revoke --address <bench32 encoded string>`
- REST API:
  - GET `/dcl/auth/proposed-revocation-accounts/{address}`

### GET_REVOKED_ACCOUNT

**Status: Implemented**

Gets a revoked account by its address.

- Parameters:
  - address: `string` - account address; Bech32 encoded
- CLI command:
  - `dcld query auth revoked-account --address <bench32 encoded string>`
- REST API:
  - GET `/dcl/auth/revoked-accounts/{address}`

### GET_ALL_ACCOUNTS

**Status: Implemented**

Gets all accounts. Revoked accounts are not returned.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query auth all-accounts`
- REST API:
  - GET `/dcl/auth/accounts`

### GET_ALL_PROPOSED_ACCOUNTS

**Status: Implemented**

Gets all proposed but not approved accounts.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query auth all-proposed-accounts`
- REST API:
  - GET `/dcl/auth/proposed-accounts`

### GET_ALL_REJECTED_ACCOUNTS

**Status: Implemented**

Get all rejected accounts.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params]
   (#common-pagination-parameters))
- CLI command:
  - `dcld query auth all-rejected-accounts`
- REST API:
  - GET `/dcl/auth/rejected-accounts`

### GET_ALL_PROPOSED_ACCOUNTS_TO_REVOKE

**Status: Implemented**

Gets all proposed but not approved accounts to be revoked.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query auth all-proposed-accounts-to-revoke`
- REST API:
  - GET `/dcl/auth/proposed-revocation-accounts`

### GET_ALL_REVOKED_ACCOUNTS

**Status: Implemented**

Gets all revoked accounts.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query auth all-revoked-accounts`
- REST API:
  - GET `/dcl/auth/revoked-accounts`

### ROTATE_KEY

**Status: Not Implemented**

Rotate's the Account's public key by the owner.

- Who can send:
  - Any role; owner

## VALIDATOR_NODE

### ADD_VALIDATOR_NODE

**Status: Implemented**

Adds a new Validator node.

- Parameters:
  - pubkey: `string` - The validator's Protobuf JSON encoded public key
  - moniker: `string` - The validator's human-readable name
  - identity: `optional(string)` - identity signature (ex. UPort or Keybase)
  - website: `optional(string)` - The validator's site link
  - details: `optional(string)` - The validator's details
  - ip: `optional(string)` - The node's public IP
  - node-id: `optional(string)` - The node's ID
- In State: `validator/Validator/value/<owner-address>`
- Who can send:
  - NodeAdmin
- CLI command:
  - `dcld tx validator add-node --pubkey='<protobuf JSON encoded>' --moniker=<string> --from=<account>`

### DISABLE_VALIDATOR_NODE

**Status: Implemented**

Disables the Validator node (removes from the validator set) by the owner.

- Who can send:
  - NodeAdmin; owner
- Parameters: No
- CLI command:
  - `dcld tx validator disable-node --from=<account>`

### PROPOSE_DISABLE_VALIDATOR_NODE

**Status: Implemented**

Proposes disabling of the Validator node from the validator set by a Trustee.

If more than 1 Trustee signature is required to disable a node, the disable
will be in a pending state until sufficient number of approvals is received.

- Parameters:
  - address: `string` - Bech32 encoded validator address or owner account
  - info: `optional(string)` - information/notes for the proposal
- Who can send:
  - Trustee
- CLI command:
  - `dcld tx validator propose-disable-node --address=<validator address> --from=<account>`
   e.g.:

    ```bash
    dcld query validator propose-disable-node --address=cosmos1nlt926tzc280ntkdmqvqumgrnvym8xc5wqwg3q --from alice
    ```

> **_Note:_** You can get Validator's address or owner address using query [GET_VALIDATOR](#get_validator)

### APPROVE_DISABLE_VALIDATOR_NODE

**Status: Implemented**

Approves disabling of the Validator node by a Trustee. It also can be used for revote (i.e. change vote from reject to approve)

The validator node is not disabled until sufficient number of Trustees approve it.

- Parameters:
  - address: `string` - Bech32 encoded validator address or owner account
  - info: `optional(string)` - information/notes for the approval
- Who can send:
  - Trustee
- Number of required approvals:
  - greater than 2/3 of Trustees (proposal by a Trustee is also counted as an approval)
- CLI command:
  - `dcld tx validator approve-disable-node --address=<validator address> --from=<account>`
   e.g.:

    ```bash
    dcld tx validator approve-disable-node --address=cosmos1nlt926tzc280ntkdmqvqumgrnvym8xc5wqwg3q from alice
    ```

> **_Note:_** You can get Validator's address or owner address using query [GET_VALIDATOR](#get_validator)

### REJECT_DISABLE_VALIDATOR_NODE

**Status: Implemented**

Rejects disabling of the Validator node by a Trustee. It also can be used for revote (i.e. change vote from approve to reject)

If disable validator proposal has only proposer's approval and no rejects then proposer can send this transaction to remove the proposal

The validator node is not reject until sufficient number of Trustees rejects it.

- Parameters:
  - address: `string` - Bech32 encoded validator address
  - info: `optional(string)` - information/notes for the reject
- Who can send:
  - Trustee
- Number of required rejects:
  - more than 1/3 of Trustees
- CLI command:
  - `dcld tx validator reject-disable-node --address=<validator address> --from=<account>`
   e.g.:

  ```bash
  dcld tx validator reject-disable-node --address=cosmos1nlt926tzc280ntkdmqvqumgrnvym8xc5wqwg3q --from alice
  ```

> **_Note:_** You can get Validator's address or owner address using query [GET_VALIDATOR](#get_validator)

### ENABLE_VALIDATOR_NODE

**Status: Implemented**

Enables the Validator node (returns to the validator set) by the owner.

the node will be enabled and returned to the active validator set.

- Who can send:
  - NodeAdmin; owner
- Parameters: No
- CLI command:
  - `dcld tx validator enable-node --from=<account>`

### GET_VALIDATOR

**Status: Implemented**

Gets a validator node.

- Parameters:
  - address: `string` - Bech32 encoded validator address or owner account
- CLI command:
  - `dcld query validator node --address=<validator address|account>`  e.g.:

    ```bash
    dcld query validator node --address=cosmosvaloper1qse069r3w0d82dul4xluqapxfg62qlndsdw9ms
    ```

    or

    ```bash
    dcld query validator node --address=cosmos1nlt926tzc280ntkdmqvqumgrnvym8xc5wqwg3q
    ```

- REST API:
  - GET `/dcl/validator/nodes/{owner}`

### GET_ALL_VALIDATORS

**Status: Implemented**

Gets the list of all validator nodes from the store.

> **_Note:_**  All stored validator nodes (`active` and `jailed`) will be returned by default.
In order to get an active validator set use specific command [validator set](#validator-set).

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query validator all-nodes`
- REST API:
  - GET `/dcl/validator/nodes`

### GET_PROPOSED_DISABLE_VALIDATOR

**Status: Implemented**

Gets a proposed validator node.

- Parameters:
  - address: `string` - Bech32 encoded validator address or owner account
- CLI command:
  - `dcld query validator proposed-disable-node --address=<validator address|account>`  e.g.:

    ```bash
    dcld query validator proposed-disable-node --address=cosmosvaloper1qse069r3w0d82dul4xluqapxfg62qlndsdw9ms
    ```

    or

    ```bash
    dcld query validator proposed-disable-node --address=cosmos1nlt926tzc280ntkdmqvqumgrnvym8xc5wqwg3q
    ```

- REST API:
  - GET `/dcl/validator/proposed-disable-nodes/{address}`

### GET_ALL_PROPOSED_DISABLE_VALIDATORS

**Status: Implemented**

Gets the list of all proposed disable validator nodes from the store.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query validator all-proposed-disable-nodes`
- REST API:
  - GET `/dcl/validator/proposed-disable-nodes`

### GET_REJECTED_DISABLE_VALIDATOR

**Status: Implemented**

Gets a rejected validator node.

- Parameters:
  - address: `string` - Bech32 encoded validator address or owner account
- CLI command:
  - `dcld query validator rejected-disable-node --address=<validator address|account>`  e.g.:

    ```bash
    dcld query validator rejected-disable-node --address=cosmosvaloper1qse069r3w0d82dul4xluqapxfg62qlndsdw9ms
    ```

    or

    ```bash
    dcld query validator rejected-disable-node --address=cosmos1nlt926tzc280ntkdmqvqumgrnvym8xc5wqwg3q
    ```

- REST API:
  - GET `/dcl/validator/rejected-disable-nodes/{address}`

### GET_ALL_REJECTED_DISABLE_VALIDATORS

**Status: Implemented**

Gets the list of all rejected disable validator nodes from the store.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query validator all-rejected-disable-nodes`
- REST API:
  - GET `/dcl/validator/rejected-disable-nodes`

### GET_DISABLED_VALIDATOR

**Status: Implemented**

Gets a disabled validator node.

- Parameters:
  - address: `string` - Bech32 encoded validator address or owner account
- CLI command:
  - `dcld query validator disabled-node --address=<validator address|account>`
   e.g.:

    ```bash
    dcld query validator disabled-node --address=cosmosvaloper1qse069r3w0d82dul4xluqapxfg62qlndsdw9ms
    ```

    or

    ```bash
    dcld query validator disabled-node --address=cosmos1nlt926tzc280ntkdmqvqumgrnvym8xc5wqwg3q
    ```

- REST API:
  - GET `/dcl/validator/disabled-nodes/{address}`

### GET_ALL_DISABLED_VALIDATORS

**Status: Implemented**

Gets the list of all disabled validator nodes from the store.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query validator all-disabled-nodes`
- REST API:
  - GET `/dcl/validator/disabled-nodes`

### GET_LAST_VALIDATOR_POWER

**Status: Implemented**

Gets a last validator node power.

- Parameters:
  - address: `string` - Bech32 encoded validator address or owner account
- CLI command:
  - `dcld query validator last-power --address=<validator address|account>`
   e.g.:

    ```bash
    dcld query validator last-power --address=cosmosvaloper1qse069r3w0d82dul4xluqapxfg62qlndsdw9ms
    ```

    or

    ```bash
    dcld query validator last-power --address=cosmos1nlt926tzc280ntkdmqvqumgrnvym8xc5wqwg3q
    ```

- REST API:
  - GET `/dcl/validator/last-powers/{owner}`

### GET_ALL_LAST_VALIDATORS_POWER

**Status: Implemented**

Gets the list of all last validator nodes power from the store.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query validator all-last-powers`
- REST API:
  - GET `/dcl/validator/last-powers`

### UPDATE_VALIDATOR_NODE

**Status: Not Implemented**

Updates the Validator node by the owner.  
`address` is used to reference the node, but can not be changed.

- Parameters:
  - address: `string` - Bech32 encoded validator address or owner account
  - moniker: `string` - The validator's human-readable name
  - identity: `optional(string)` - identity signature (ex. UPort or Keybase)
  - website: `optional(string)` - The validator's site link
  - details: `optional(string)` - The validator's details
  - ip: `optional(string)` - The node's public IP
  - node-id: `optional(string)` - The node's ID
- Who can send:
  - NodeAdmin; owner

## UPGRADE

### PROPOSE_UPGRADE

**Status: Implemented**

Proposes an upgrade plan with the given name at the given height.

- Parameters:
  - name: `string` - upgrade plan name
  - upgrade-height: `int64` -  upgrade plan height (positive non-zero)
  - upgrade-info: `optional(string)` - upgrade plan info (for node admins to
      read). Recommended format is an os/architecture -> application binary URL
      map as a JSON under `binaries` key where each URL should include the
      corresponding checksum as `checksum` query parameter with the value in the
      format `type:value` where `type` is `sha256` or `sha512` and `value` is
      the actual checksum value. For example:

```json
{
  "binaries": {
    "linux/amd64":"https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v0.7.0/dcld?checksum=sha256:aec070645fe53ee3b3763059376134f058cc337247c978add178b6ccdfb0019f"
  }
}
```

- In State: `dclupgrade/ProposedUpgrade/value/<name>`
- Who can send:
  - Trustee
- Number of required approvals:
  - greater than 2/3 of Trustees (proposal by a Trustee is also counted as an approval)
- CLI command minimal:

```bash
dcld tx dclupgrade propose-upgrade --name=<string> --upgrade-height=<int64> --from=<account>
```

- CLI command full:

```bash
dcld tx dclupgrade propose-upgrade --name=<string> --upgrade-height=<int64> --upgrade-info=<string> --from=<account>
```

> **_Note:_**  If the current upgrade proposal is out of date(when the current network height is greater than the proposed upgrade height), we can resubmit the upgrade proposal with the same name.

### APPROVE_UPGRADE

**Status: Implemented**

Approves the proposed upgrade plan with the given name. It also can be used for revote (i.e. change vote from reject to approve)

- Parameters:
  - name: `string` - upgrade plan name
- In State: `upgrade/0x0`
- Who can send:
  - Trustee
- Number of required approvals:
  - greater than 2/3 of Trustees (proposal by a Trustee is also counted as an approval)
- CLI command:

```bash
dcld tx dclupgrade approve-upgrade --name=<string> --from=<account>
```

### REJECT_UPGRADE

**Status: Implemented**

Rejects the proposed upgrade plan with the given name. It also can be used for revote (i.e. change vote from approve to reject)

If proposed upgrade has only proposer's approval and no rejects then proposer can send this transaction to remove the proposal

- Paramaters:
  - name: `string` - upgrade plan name
- In State: `RejectUpgrade/value/<name>`
- Who can send:
  - Trustee
- Number of required rejects:
  - more than 1/3 of Trustees
- CLI command:

```bash
dcld tx dclupgrade reject-upgrade --name=<string> --from=<account>
```

### GET_PROPOSED_UPGRADE

**Status: Implemented**

Gets the proposed upgrade plan with the given name.

- Parameters:
  - name: `string` - upgrade plan name
- CLI command:

```bash
dcld query dclupgrade proposed-upgrade --name=<string>
```

- REST API:
  - GET `/dcl/dclupgrade/proposed-upgrades/{name}`

### GET_APPROVED_UPGRADE

**Status: Implemented**

Gets the approved upgrade plan with the given name.

- Parameters:
  - name: `string` - upgrade plan name
- CLI command:

```bash
dcld query dclupgrade approved-upgrade --name=<string>
```

- REST API:
  - GET `/dcl/dclupgrade/approved-upgrades/{name}`

### GET_REJECTED_UPGRADE

**Status: Implemented**

Gets the rejected upgrade plan with the given name.

- Parameters:
  - name: `string` - upgrade plan name
- CLI command:

```bash
dcld query dclupgrade rejected-upgrade --name=<string>
```

- REST API:
  - GET `/dcl/dclupgrade/rejected-upgrades/{name}`

### GET_ALL_PROPOSED_UPGRADES

**Status: Implemented**

Gets all the proposed upgrade plans.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:

```bash
dcld query dclupgrade all-proposed-upgrades
```

- REST API:
  - GET `/dcl/dclupgrade/proposed-upgrades`

### GET_ALL_APPROVED_UPGRADES

**Status: Implemented**

Gets all the approved upgrade plans.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:

```bash
dcld query dclupgrade all-approved-upgrades
```

- REST API:
  - GET `/dcl/dclupgrade/approved-upgrades`

### GET_ALL_REJECTED_UPGRADES

**Status: Implemented**

Gets all the rejected upgrade plans.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:

```bash
dcld query dclupgrade all-rejected-upgrades
```

- REST API:
  - GET `/dcl/dclupgrade/rejected-upgrades`

### GET_UPGRADE_PLAN

**Status: Implemented**

Gets the currently scheduled upgrade plan, if it exists.

- CLI command:

```bash
dcld query upgrade plan
```

- REST API:
  - GET `/cosmos/upgrade/v1beta1/current_plan`

### GET_APPLIED_UPGRADE

**Status: Implemented**

Returns the header for the block at which the upgrade with the given name was
applied, if it was previously executed on the chain. This helps a client
determine which binary was valid over a given range of blocks, as well as gives
more context to understand past migrations.

- Parameters:
  - `string` - upgrade name
- CLI command:

```bash
dcld query upgrade applied <string>
```

- REST API:
  - GET `/cosmos/upgrade/v1beta1/applied_plan/{name}`

### GET_MODULE_VERSIONS

**Status: Implemented**

Gets a list of module names and their respective consensus versions. Following
the command with a specific module name will return only that module's
information.

- Parameters:
  - `optional(string)` - module name
- CLI command minimal:

```bash
dcld query upgrade module_versions
```

- CLI command full:

```bash
dcld query upgrade module_versions <string>
```

- REST API:
  - GET `/cosmos/upgrade/v1beta1/module_versions`

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
