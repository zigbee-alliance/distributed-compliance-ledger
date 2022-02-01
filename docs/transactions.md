# Transactions and Queries
See use case sequence diagrams for the examples of how transaction can be used.

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
  - `TestHouse` - can add testing results for a model.
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
       - Fetch `account number` and `sequence` by CLI 1:  `dcld query auth account --address <address>`
       - Sign transaction by CLI 2: `dcld tx sign txn.json --from <from> --account-number <int> --sequence <int> --gas "auto" --offline --output-document txn.json`
       - Broadcat transaction by CLI 1: `dcld tx broadcast txn.json`
- gRPC:
    - Generate a client code from the proto files [proto](../proto) for the client language (see https://grpc.io/docs/languages/)   
    - Build, sign, and broadcast the message (transaction). 
      See [grpc/rest integration tests](../integration_tests/grpc_rest) as an example.   
- REST API
    - Build and sign a transaction by one of the following ways
        - In code via gRPC (see above)
        - Via CLI commands specifying `--generate-only` flag and using `dcld tx sign` (see above)
    - The user does a `POST` of the signed request to `http://<node-ip>:1317/cosmos/tx/v1beta1/txs` endpoint.     
    - Example
        ```
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
  - OpenAPI specification: https://zigbee-alliance.github.io/distributed-compliance-ledger/.
  - Any running node exposes a REST API at port `1317`. See https://docs.cosmos.network/master/core/grpc_rest.html.
  - See `REST API` section for every read request.
  - See [grpc/rest integration tests](../integration_tests/grpc_rest) as an example.
  - There are no state proofs in REST, so REST queries should be sent to trusted Validator or Observer nodes only.
- gRPC
  - Any running node exposes a REST API at port `9090`. See https://docs.cosmos.network/master/core/grpc_rest.html.
  - Generate a client code from the proto files [proto](../proto) for the client language (see https://grpc.io/docs/languages/).
  - See [grpc/rest integration tests](../integration_tests/grpc_rest) as an example.
  - There are no state proofs in gRPC, so gRPC queries should be sent to trusted Validator or Observer nodes only.
- Tendermint RPC
  - Tendermint RPC OpenAPI specification can be found in https://zigbee-alliance.github.io/distributed-compliance-ledger/.
  - Tendermint RPC is exposed by every running node  at port `26657`. See https://docs.cosmos.network/master/core/grpc_rest.html#tendermint-rpc.
  - Tendermint RPC supports state proofs. Tendermint's Light Client library can be used to verify the state proofs.
    So, if Light Client API is used, then it's possible to communicate with non-trusted nodes.
  - Please note, that multi-value queries don't have state proofs support and should be sent to trusted nodes only.


`NotFound` (404 code) is returned if an entry is not found on the ledger.
    
##### Query types     
- Query single value
- Query list of values with pagination support (should be sent to trusted nodes only)

##### Common pagination parameters         
- count-total `optional(bool)`:  count total number of records 
- limit `optional(uint)`:        pagination limit (default 100)
- offset `optional(uint)`:       pagination offset 
- page `optional(uint)`:         pagination page. This sets offset to a multiple of limit (default 1).
- page-key `optional(string)`:   pagination page-key
- reverse `optional(bool)`:       results are sorted in descending order


## VENDOR INFO


#### ADD_VENDOR_INFO
**Status: Implemented**

Adds a record about a Vendor.

- Parameters:
  - vid: `uint16` -  Vendor ID (positive non-zero)
  - vendorName: `string` -  Vendor name
  - companyLegalName: `string` -  Legal name of the vendor company
  - companyPreferredName: `optional(string)` -  Preferred name of the vendor company
  - vendorLandingPageURL: `optional(string)` -  URL of the vendor's landing page
- In State: `vendorinfo/VendorInfo/value/<vid>`
- Who can send: 
    - Vendor account
- CLI command:
  - `dcld tx vendorinfo add-vendor --vid=<uint16> --vendorName=<string> --companyLegalName=<string> --companyPreferredName=<string> --vendorLandingPageURL=<string> --from=<account>`


#### UPDATE_VENDOR_INFO
**Status: Implemented**

Updates a record about a Vendor.

- Parameters:
  - vid: `uint16` -  Vendor ID (positive non-zero)
  - vendorName: `optional(string)` -  Vendor name
  - companyLegalName: `optional(string)` -  Legal name of the vendor company
  - companyPreferredName: `optional(string)` -  Preferred name of the vendor company
  - vendorLandingPageURL: `optional(string)` -  URL of the vendor's landing page
- In State: `vendorinfo/VendorInfo/value/<vid>`
- Who can send: 
    - Account with a vendor role who has the matching Vendor ID
- CLI command:
  - `dcld tx vendorinfo update-vendor --vid=<uint16> ... --from=<account>`


#### GET_VENDOR_INFO
**Status: Implemented**

Gets a Vendor Info for the given `vid` (vendor ID).

- Parameters:
    - vid: `uint16` -  model vendor ID (positive non-zero)
- CLI command: 
    -   `dcld query vendorinfo vendor --vid=<uint16>`
- REST API: 
    - GET `/dcl/vendorinfo/vendors/{vid}`

#### GET_ALL_VENDOR_INFO
**Status: Implemented**

Gets information about all vendors for all VIDs.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command: 
    -   `dcld query vendorinfo all-vendors`
- REST API: 
    - GET `/dcl/vendorinfo/vendors`

## MODEL and MODEL_VERSION

#### ADD_MODEL
**Status: Implemented**

Adds a new Model identified by a unique combination of `vid` (vendor ID) and `pid` (product ID).

Not all fields can be edited (see `EDIT_MODEL`).


- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - deviceTypeID: `uint16` -  DeviceTypeID is the device type identifier. For example, DeviceTypeID 10 (0x000a), is the device type identifier for a Door Lock.
  - productName: `string` -  model name
  - productLabel: `string` -  model description (string or path to file containing data)
  - partNumber: `string` -  stock keeping unit
  - commissioningCustomFlow: `optional(uint8)` - A value of 1 indicates that user interaction with the device (pressing a button, for example) is required before commissioning can take place. When CommissioningCustomflow is set to a value of 2, the commissioner SHOULD attempt to obtain a URL which MAY be used to provide an end-user with the necessary details for how to configure the product for initial commissioning
  - commissioningCustomFlowURL: `optional(string)` - commissioningCustomFlowURL SHALL identify a vendor specific commissioning URL for the device model when the commissioningCustomFlow field is set to '2'
  - commissioningModeInitialStepsHint: `optional(uint32)` - commissioningModeInitialStepsHint SHALL identify a hint for the steps that can be used to put into commissioning mode a device that has not yet been commissioned. This field is a bitmap with values defined in the Pairing Hint Table. For example, a value of 1 (bit 0 is set) indicates that a device that has not yet been commissioned will enter Commissioning Mode upon a power cycle.
  - commissioningModeInitialStepsInstruction: `optional(string)` - commissioningModeInitialStepsInstruction SHALL contain text which relates to specific values of CommissioningModeInitialStepsHint. Certain values of CommissioningModeInitialStepsHint, as defined in the Pairing Hint Table, indicate a Pairing Instruction (PI) dependency, and for these values the commissioningModeInitialStepsInstruction SHALL be set
  - commissioningModeSecondaryStepsHint: `optional(uint32)` - commissioningModeSecondaryStepsHint SHALL identify a hint for steps that can be used to put into commissioning mode a device that has already been commissioned. This field is a bitmap with values defined in the Pairing Hint Table. For example, a value of 4 (bit 2 is set) indicates that a device that has already been commissioned will require the user to visit a current CHIP Administrator to put the device into commissioning mode.
  - commissioningModeSecondaryStepInstruction: `optional(string)` - commissioningModeSecondaryStepInstruction SHALL contain text which relates to specific values of commissioningModeSecondaryStepsHint. Certain values of commissioningModeSecondaryStepsHint, as defined in the Pairing Hint Table, indicate a Pairing Instruction (PI) dependency, and for these values the commissioningModeSecondaryStepInstruction SHALL be set
  - userManualURL: `optional(string)` - URL that contains product specific web page that contains user manual for the device model.
  - supportURL: `optional(string)` - URL that contains product specific web page that contains support details for the device model.
  - productURL: `optional(string)` - URL that contains product specific web page that contains details for the device model.
  - lsfURL: `optional(string)` - URL to the Localized String File of this product.
  - lsfRevision: `optional(uint32)` - LsfRevision is a monotonically increasing positive integer indicating the latest available.version of Localized String File
- In State:
  - `model/Model/value/<vid>/<pid>`
  - `model/VendorProducts/value/<vid>`
- Who can send: 
    - Vendor account who is associated with the given vid
- CLI command minimal:
```
 dcld tx model add-model --vid=<uint16> --pid=<uint16> --deviceTypeID=<uint16> --productName=<string> --productLabel=<string or path> --partNumber=<string> 
 --from=<account>
```
- CLI command full:
```
 dcld tx model add-model --vid=<uint16> --pid=<uint16> --deviceTypeID=<uint16> --productName=<string> --productLabel=<string or path> --partNumber=<string> 
    --commissioningCustomFlow=<uint8> --commissioningCustomFlowUrl=<string> --commissioningModeInitialStepsHint=<uint32> --commissioningModeInitialStepsInstruction=<string>
    --commissioningModeSecondaryStepsHint=<uint32> --commissioningModeSecondaryStepsInstruction=<string> --userManualURL=<string> --supportURL=<string> --productURL=<string>
    --from=<account>
```

#### EDIT_MODEL
**Status: Implemented**

Edits an existing Model identified by a unique combination of `vid` (vendor ID) and `pid` (product ID)
by the vendor account.

Only the fields listed below (except `vid` and `pid`) can be edited. If other fields need to be edited -
a new model info with a new `vid` or `pid` can be created.

All non-edited fields remain the same.


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
  - lsfRevision: `optional(uint32)` - LsfRevision is a monotonically increasing positive integer indicating the latest available.  
- In State: `model/Model/value/<vid>/<pid>`
- Who can send: 
    - Vendor account associated with the same vid who has created the model
- CLI command: 
    -   `dcld tx model update-model --vid=<uint16> --pid=<uint16> ... --from=<account>`

#### ADD_MODEL_VERSION
**Status: Implemented**

Adds a new Model Software Version identified by a unique combination of `vid` (vendor ID), `pid` (product ID) and `softwareVersion`. 

Not all Model Software Version fields can be edited (see `EDIT_MODEL_VERSION`). 

If one of `OTA_URl`, `OTA_checksum` or `OTA_checksum_type` fields is set, then the other two must also be set.

- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - softwareVersion: `uint32` - model software version
  - softwareVersionSting: `string` - model software version string 
  - cdVersionNumber `uint32` - CD Version Number of the certification
  - minApplicableSoftwareVersion `uint32` - MinApplicableSoftwareVersion should specify the lowest SoftwareVersion for which this image can be applied 
  - maxApplicableSoftwareVersion `uint32` - MaxApplicableSoftwareVersion should specify the highest SoftwareVersion for which this image can be applied
  - firmwareDigests `optional(string)` - FirmwareDigests field included in the Device Attestation response when this Software Image boots on the device
  - softwareVersionValid `optional(bool)` - Flag to indicate whether the software version is valid or not (default true)
  - otaURL `optional(string)` - URL where to obtain the OTA image
  - otaFileSize `optional(string)`  - OtaFileSize is the total size of the OTA software image in bytes
  - otaChecksum `optional(string)` - Digest of the entire contents of the associated OTA Software Update Image under the OtaUrl attribute, encoded in base64 string representation. The digest SHALL have been computed using the algorithm specified in OtaChecksumType
  - otaChecksumType `optional(string)` - Numeric identifier as defined in IANA Named Information Hash Algorithm Registry for the type of otaChecksum. For example, a value of 1 would match the sha-256 identifier, which maps to the SHA-256 digest algorithm
  - releaseNotesURL `optional(string)` - URL that contains product specific web page that contains release notes for the device model.
- In State:
  - `model/ModelVersion/value/<vid>/<pid>/<softwareVersion>`
  - `model/ModelVersions/value/<vid>/<pid>`
- Who can send: 
    - Vendor with same vid who created the Model
- CLI command minimal:
```
dcld tx model add-model-version --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --softwareVersionString=<string> --cdVersionNumber=<uint32>
--minApplicableSoftwareVersion=<uint32> --maxApplicableSoftwareVersion=<uint32> --from=<account>
```
- CLI command full:
```
dcld tx model add-model-version --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --softwareVersionString=<string> --cdVersionNumber=<uint32>
--minApplicableSoftwareVersion=<uint32> --maxApplicableSoftwareVersion=<uint32>
--firmwareDigests=<string> --softwareVersionValid=<bool> --otaURL=<string> --otaFileSize=<string> --otaChecksum=<string> --otaChecksumType=<string> --releaseNotesURL=<string> 
--from=<account>
```

#### EDIT_MODEL_VERSION
**Status: Implemented**

Edits an existing Model Software Version identified by a unique combination of `vid` (vendor ID) `pid` (product ID) and `softwareVersion`
by the vendor.

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

- In State: `model/ModelVersion/value/<vid>/<pid>/<softwareVersion>`
- Who can send: 
    - Vendor associated with the same vid who created the Model
- CLI command: 
    -   `dcld tx model update-model-version --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> ... --from=<account>`



#### GET_MODEL
**Status: Implemented**

Gets a Model Info with the given `vid` (vendor ID) and `pid` (product ID).

- Parameters:
    - vid: `uint16` -  model vendor ID (positive non-zero)
    - pid: `uint16` -  model product ID (positive non-zero)
- CLI command: 
    -   `dcld query model get-model --vid=<uint16> --pid=<uint16>`
- REST API: 
    -   GET `/dcl/model/models/{vid}/{pid}`
 
#### GET_MODEL_VERSION
**Status: Implemented**

Gets a Model Software Versions for the given `vid`, `pid` and `softwareVersion`.

- Parameters
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - softwareVersion: `uint32` - model software version (positive non-zero)
- CLI command: 
    -   `dcld query model model-version --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32>`
- REST API: 
    -   GET `/dcl/model/versions/{vid}/{pid}/{softwareVersion}`

#### GET_ALL_MODELS
**Status: Implemented**

Gets all Model Infos for all vendors.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command: 
    -   `dcld query model all-models`
- REST API: 
    -   GET `/dcl/model/models`

#### GET_ALL_VENDOR_MODELS
**Status: Implemented**

Gets all Model Infos by the given Vendor (`vid`).

- Parameters:
    - vid: `uint16` -  model vendor ID (positive non-zero)
- CLI command: 
    -   `dcld query model vendor-models --vid=<uint16>`
- REST API: 
    - GET `/dcl/model/models/{vid}`

#### GET_ALL_MODEL_VERSIONS
**Status: Implemented**

Gets all Model Software Versions for the given `vid` and `pid` combination.

- Parameters:
    - vid: `uint16` -  model vendor ID (positive non-zero)
    - pid: `uint16` -  model product ID (positive non-zero)
- CLI command: 
    -   `dcld query model all-model-versions --vid=<uint16> --pid=<uint16>`
- REST API: 
    - GET `/dcl/model/versions/{vid}/{pid}`

## TEST_DEVICE_COMPLIANCE

#### ADD_TEST_RESULT
**Status: Implemented**

Submits result of a compliance testing for the given Model Version (`vid`, `pid`, `softwareVersion` and `softwareVersionString`).
The test result can be a blob of data or a reference (URL) to an external storage.

The corresponding Model Version must be present on ledger.

Multiple test results (potentially from different test houses) can be added for the same Model Version.

The test result is immutable and can not be deleted or removed after submitting. 
Another test result can be submitted instead.

- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - softwareVersion: `uint32` - model software version
  - softwareVersionSting: `string` - model software version string 
  - test_result: `string` - Test result (string or path to file containing data)
  - test_date: `string` - Date of test result (rfc3339 encoded), for example 2019-10-12T07:20:50.52Z
- In State: `compliancetest/TestingResults/value/<vid>/<pid>/<softwareVersion>`
- Who can send: 
    - TestHouse
- CLI command: 
    -   `dcld tx compliancetest add-test-result --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --softwareVersionString=<string> --test-result=<string-or-path> --test-date=<rfc3339 encoded date> --from=<account>`

#### GET_TEST_RESULT
**Status: Implemented**

Gets a test result for the given `vid` (vendor ID), `pid` (product ID) and `softwareVersion`.

- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - softwareVersion: `uint32` - model software version
- CLI command: 
    -   `dcld query compliancetest test-result --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32>`
- REST API: 
    -   GET `/dcl/compliancetest/testing-results/{vid}/{pid}/{softwareVersion}`

#### GET_ALL_TEST_RESULTS
**Status: Implemented**

Gets all test results.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command: 
    -   `dcld query compliancetest all-test-results`
- REST API: 
    -   GET `/dcl/compliancetest/testing-results`

## CERTIFY_DEVICE_COMPLIANCE

#### CERTIFY_MODEL
**Status: Implemented**

Attests compliance of the Model Version to the ZB or Matter standard.

`REVOKE_MODEL_CERTIFICATION` should be used for revoking (disabling) the compliance.
It's possible to call `CERTIFY_MODEL` for revoked model versions to enable them back. 

The corresponding Model Version must be present on the ledger.

It must be called for every compliant device for use cases where compliance 
is tracked on ledger.
For such use cases the corresponding test results must be present on the ledger.

It can be used for use cases where only revocation is tracked on the ledger to remove a Model Version
from the revocation list.
For such use cases the corresponding test results are not required to be on the ledger.
 
- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - softwareVersion: `uint32` - model software version
  - softwareVersionSting: `string` - model software version string
  - certification_date: `string` - The date of model certification (rfc3339 encoded), for example 2019-10-12T07:20:50.52Z
  - certification_type: `string`  - Certification type - Currently 'zigbee' and 'matter' types are supported
  - reason `optional(string)`  - optional comment describing the reason of the certification
- In State:
  - `compliance/ComplianceInfo/value/<vid>/<pid>/<softwareVersion>/<certificationType>`
  - `compliance/CertifiedModel/value/<vid>/<pid>/<softwareVersion>/<certificationType>`
- Who can send: 
    - CertificationCenter
- CLI command: 
    -   `dcld tx compliance certify-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --softwareVersionString=<string>  --certificationType=<matter|zigbee> --certificationDate=<rfc3339 encoded date> --reason=<string> --from=<account>`
    
#### REVOKE_MODEL_CERTIFICATION
**Status: Implemented**

Revoke compliance of the Model Version to the ZB or Matter standard.

The corresponding Model Version must be present on the ledger. The corresponding test results are not required to be on the ledger.

It can be used in cases where every compliance result 
is written on the ledger (`CERTIFY_MODEL` was called), or
 cases where only revocation list is stored on the ledger.
 
- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - softwareVersion: `uint32` - model software version
  - softwareVersionSting: `string` - model software version string
  - revocation_date: `string` - The date of model revocation (rfc3339 encoded), for example 2019-10-12T07:20:50.52Z
  - certification_type: `string`  - Certification type - Currently 'zigbee' and 'matter' types are supported
  - reason `optional(string)`  - optional comment describing the reason of revocation
- In State: 
  - `compliance/ComplianceInfo/value/<vid>/<pid>/<softwareVersion>/<certificationType>`
  - `compliance/RevokedModel/value/<vid>/<pid>/<softwareVersion>/<certificationType>`
- Who can send: 
    - CertificationCenter
- CLI command: 
    -   `dcld tx compliance revoke-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<matter|zigbee> --revocationDate=<rfc3339 encoded date> --reason=<string> --from=<account>`
 
#### PROVISION_MODEL
**Status: Implemented**

Sets provisional state for the Model Version.

The corresponding Model Version and test results are not required to be on the ledger.

Can not be set if there is already a certification record on the ledger (certified or revoked).
 
- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - softwareVersion: `uint32` - model software version
  - softwareVersionSting: `string` - model software version string
  - provisional_date: `string` - The date of model provisioning (rfc3339 encoded), for example 2019-10-12T07:20:50.52Z
  - certification_type: `string`  - Certification type - Currently 'zigbee' and 'matter' types are supported
  - reason `optional(string)`  - optional comment describing the reason of revocation
- In State: 
  - `compliance/ComplianceInfo/value/<vid>/<pid>/<softwareVersion>/<certificationType>`
  - `compliance/ProvisionalModel/value/<vid>/<pid>/<softwareVersion>/<certificationType>`
- Who can send: 
    - CertificationCenter
- CLI command: 
    -   `dcld tx compliance provision-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<matter|zigbee> --provisionalDate=<rfc3339 encoded date> --reason=<string> --from=<account>`
 
    
    
#### GET_CERTIFIED_MODEL
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
    -   `dcld query compliance certified-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<zigbee|matter>`
- REST API: 
    -   GET `/dcl/compliance/certified-models/{vid}/{pid}/{software_version}/{certification_type}`


#### GET_REVOKED_MODEL
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
    -   `dcld query compliance revoked-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<zigbee|matter>`
- REST API: 
    -   GET `/dcl/compliance/revoked-models/{vid}/{pid}/{software_version}/{certification_type}`

#### GET_PROVISIONAL_MODEL
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
    -   `dcld query compliance provisional-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<zigbee|matter>`
- REST API: 
    -   GET `/dcl/compliance/provisional-models/{vid}/{pid}/{software_version}/{certification_type}`

#### GET_COMPLIANCE_INFO
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
    -   `dcld query compliance compliance-info --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<zigbee|matter>`
- REST API: 
    -   GET `/dcl/compliance/compliance-info/{vid}/{pid}/{software_version}/{certification_type}`

#### GET_ALL_CERTIFIED_MODELS
**Status: Implemented**

Gets all compliant Model Versions for all vendors (`vid`s).

This is the aggregation of compliance and
revocation information for every vid/pid. It should be used in cases where compliance is tracked on ledger.

Should be sent to trusted nodes only.
 
- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command: 
    -   `dcld query compliance all-certified-models`
- REST API: 
    - GET `/dcl/compliance/certified-models`
     
#### GET_ALL_REVOKED_MODELS
**Status: Implemented**

Gets all revoked Model Versions for all vendors (`vid`s).

It contains information about revocation only, so it should be used in cases
 where only revocation is tracked on the ledger.
 
Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command: 
    -   `dcld query compliance all-revoked-models`
- REST API: 
    -   GET `/dcl/compliance/revoked-models`

#### GET_ALL_PROVISIONAL_MODELS
**Status: Implemented**

Gets all Model Versions in provisional state for all vendors (`vid`s).

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command: 
    -   `dcld query compliance all-provisional-models`
- REST API: 
    -   GET `/dcl/compliance/provisional-models`

    
#### GET_ALL_COMPLIANCE_INFO_RECORDS
**Status: Implemented**

Gets all stored compliance information records for all vendors (`vid`s).

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command: 
    -   `dcld query compliance all-compliance-info`
- REST API: 
    -   GET `/dcl/compliance/compliance-info`

## X509 PKI

**NOTE**: X.509 v3 certificates are only supported (all certificates MUST contain `Subject Key ID` field).
All PKI related methods are based on this restriction.

#### PROPOSE_ADD_X509_ROOT_CERT
**Status: Implemented**

Proposes a new self-signed root certificate.

If it's sent by a non-Trustee account, or more than 1 Trustee signature is required to add a root certificate, 
then the certificate
will be in a pending state until sufficient number of other Trustee's approvals is received.

The certificate is immutable. It can only be revoked by either the owner or a quorum of Trustees.

- Parameters:
  - cert: `string` - PEM encoded certificate (string or path to file containing data)
- In State: `pki/ProposedCertificate/value/<Certificate's Subject>/<Certificate's Subject Key ID>`
- Who can send: 
    - Any role
- CLI command: 
    -   `dcld tx pki propose-add-x509-root-cert --certificate=<string-or-path> --from=<account>`
- Validation:
    - provided certificate must be root: 
        - `Issuer` == `Subject` 
        - `Authority Key Identifier` == `Subject Key Identifier`
    - no existing `Proposed` certificate with the same `<Certificate's Subject>:<Certificate's Subject Key ID>` combination.
    - no existing certificate with the same `<Certificate's Issuer>:<Certificate's Serial Number>` combination.
    - if approved certificates with the same `<Certificate's Subject>:<Certificate's Subject Key ID>` combination already exists:
        - sender must match to the owner of the existing certificates.
    - the signature (self-signature) and expiration date are valid.

#### APPROVE_ADD_X509_ROOT_CERT
**Status: Implemented**

Approves the proposed root certificate.

The certificate is not active until sufficient number of Trustees approve it. 

- Parameters:
  - subject: `string`  - proposed certificates's `Subject`
  - subject_key_id: `string`  - proposed certificates's `Subject Key Id`
- In State: `pki/ApprovedCertificates/value/<Certificate's Subject>/<Certificate's Subject Key ID>`
- Who can send: 
    - Trustee
- Number of required approvals: 
    - 2/3 of Trustees
- CLI command: 
    -   `dcld tx pki approve-add-x509-root-cert --subject=<string> --subject-key-id=<hex string> --from=<account>`
- Validation:
    - the proposed certificate hasn't been approved by the signer yet
        
#### ADD_X509_CERT
**Status: Implemented**

Adds an intermediate or leaf X509 certificate signed by a chain of certificates which must be
already present on the ledger.

The certificate is immutable. It can only be revoked by either the owner or a quorum of Trustees.

- Parameters:
  - cert: `string` - PEM encoded certificate (string or path to file containing data)
- In State:
  - `pki/ApprovedCertificates/value/<Certificate's Subject>/<Certificate's Subject Key ID>`
  - `pki/ChildCertificates/value/<Certificate's Subject>/<Certificate's Subject Key ID>`
- Who can send: 
    - Any role
- CLI command: 
    -   `dcld tx pki add-x509-cert --certificate=<string-or-path> --from=<account>`
- Validation:
    - provided certificate must not be root: 
        - `Issuer` != `Subject` 
        - `Authority Key Identifier` != `Subject Key Identifier`
    - no existing certificate with the same `<Certificate's Issuer>:<Certificate's Serial Number>` combination.
    - if certificates with the same `<Certificate's Subject>:<Certificate's Subject Key ID>` combination already exist:
        - sender must match to the owner of the existing certificates.
    - the signature (self-signature) and expiration date are valid.
    - parent certificate must be already stored on the ledger and a valid chain to some root certificate can be built. 

Note: Multiple certificates can refer to the same `<Certificate's Subject>:<Certificate's Subject Key ID>` combination.
    
#### REVOKE_X509_CERT
**Status: Implemented**

Revokes the given X509 certificate (either intermediate or leaf).
All the certificates in the chain signed by the revoked certificate will be revoked as well.

Only the owner (sender) can revoke the certificate.
Root certificates can not be revoked this way, use  `PROPOSE_X509_CERT_REVOC` and `APPROVE_X509_ROOT_CERT_REVOC` instead.  

- Parameters:
  - subject: `string`  - certificates's `Subject`
  - subject_key_id: `string`  - certificates's `Subject Key Id`
- In State: `pki/RevokedCertificates/value/<Certificate's Subject>/<Certificate's Subject Key ID>`
- Who can send: 
    - Any role; owner
- CLI command: 
    -   `dcld tx pki revoke-x509-cert --subject=<string> --subject-key-id=<hex string> --from=<account>`

#### PROPOSE_REVOKE_X509_ROOT_CERT
**Status: Implemented**

Proposes revocation of the given X509 root certificate by a Trustee.

All the certificates in the chain signed by the revoked certificate will be revoked as well.

If more than 1 Trustee signature is required to revoke a root certificate, 
then the certificate will be in a pending state until sufficient number of other Trustee's approvals is received.

- Parameters:
  - subject: `string`  - certificates's `Subject`
  - subject_key_id: `string`  - certificates's `Subject Key Id`
- In State: `pki/ProposedCertificateRevocation/value/<Certificate's Subject>/<Certificate's Subject Key ID>`
- Who can send: 
    - Trustee
- CLI command: 
    -   `dcld tx pki propose-revoke-x509-root-cert --subject=<string> --subject-key-id=<hex string> --from=<account>`
    


#### APPROVE_REVOKE_X509_ROOT_CERT
**Status: Implemented**

Approves the revocation of the given X509 root certificate by a Trustee.
All the certificates in the chain signed by the revoked certificate will be revoked as well.

The revocation is not applied until sufficient number of Trustees approve it. 

- Parameters:
  - subject: `string`  - certificates's `Subject`
  - subject_key_id: `string`  - certificates's `Subject Key Id`
- In State: `pki/RevokedCertificates/value/<Certificate's Subject>/<Certificate's Subject Key ID>`
- Who can send: 
    - Trustee
- Number of required approvals: 
    - 2/3 of Trustees
- CLI command: 
    -   `dcld tx pki approve-revoke-x509-root-cert --subject=<string> --subject-key-id=<hex string> --from=<account>`

#### GET_X509_CERT
**Status: Implemented**

Gets a certificate (either root, intermediate or leaf) by the given subject and subject key id attributes.
Revoked certificates are not returned. 
Use `GET_ALL_REVOKED_X509_CERTS` to get a list of all revoked certificates. 

- Parameters:
  - subject: `string`  - certificates's `Subject`
  - subject_key_id: `string`  - certificates's `Subject Key Id`
- CLI command: 
    -   `dcld query pki x509-cert --subject=<string> --subject-key-id=<hex string>`
- REST API: 
    -   GET `/dcl/pki/certificates/{subject}/{subject_key_id}`

#### GET_ALL_SUBJECT_X509_CERTS
**Status: Implemented**

Gets all certificates (root, intermediate and leaf) associated with a subject.

Revoked certificates are not returned. 
Use `GET_ALL_REVOKED_X509_CERTS` to get a list of all revoked certificates. 

- Parameters:
  - subject: `string`  - certificates's `Subject`
- CLI command: 
    -   `dcld query pki all-subject-x509-certs --subject=<string>`
- REST API: 
    - GET `/dcl/pki/certificates/{subject}`

#### GET_ALL_CHILD_X509_CERTS
**Status: Implemented**

Gets all child certificates for the given certificate.
Revoked certificates are not returned. 

- Parameters:
  - subject: `string`  - certificates's `Subject`
  - subject_key_id: `string`  - certificates's `Subject Key Id`
- CLI command: 
    -   `dcld query pki all-child-x509-certs --subject=<string> --subject-key-id=<hex string>`
- REST API: 
    -   GET `/dcl/pki/child-certificates/{subject}/{subject_key_id}`

#### GET_PROPOSED_X509_ROOT_CERT
**Status: Implemented**

Gets a proposed but not approved root certificate with the given subject and subject key id attributes.

- Parameters:
  - subject: `string`  - certificates's `Subject`
  - subject_key_id: `string`  - certificates's `Subject Key Id`
- CLI command: 
    -   `dcld query pki proposed-x509-root-cert --subject=<string> --subject-key-id=<hex string>`
- REST API: 
    -   GET `/dcl/pki/proposed-certificates/{subject}/{subject_key_id}`

#### GET_REVOKED_CERT
**Status: Implemented**

Gets a revoked certificate (either root, intermediate or leaf) by the given subject and subject key id attributes.

- Parameters:
  - subject: `string`  - certificates's `Subject`
  - subject_key_id: `string`  - certificates's `Subject Key Id`
- CLI command: 
    -   `dcld query pki revoked-x509-cert --subject=<string> --subject-key-id=<hex string>`
- REST API: 
    -   GET `/dcl/pki/revoked-certificates/{subject}/{subject_key_id}`

#### GET_PROPOSED_X509_ROOT_CERT_TO_REVOKE
**Status: Implemented**

Gets a proposed but not approved root certificate to be revoked.

- Parameters:
  - subject: `string`  - certificates's `Subject`
  - subject_key_id: `string`  - certificates's `Subject Key Id`
- CLI command: 
    -   `dcld query pki proposed-x509-root-cert-to-revoke --subject=<string> --subject-key-id=<hex string>`
- REST API: 
    -   GET `/dcl/pki/proposed-revocation-certificates/{subject}/{subject_key_id}`

#### GET_ALL_X509_ROOT_CERTS
**Status: Implemented**

Gets all approved root certificates. Revoked certificates are not returned. 
Use `GET_ALL_REVOKED_X509_CERTS_ROOT` to get a list of all revoked root certificates. 

- Parameters: No
- CLI command: 
    -   `dcld query pki all-x509-root-certs`
- REST API: 
    -   GET `/dcl/pki/root-certificates`

#### GET_ALL_REVOKED_X509_ROOT_CERTS
**Status: Implemented**

Gets all revoked root certificates.

- Parameters: No   
- CLI command: 
    -   `dcld query pki all-revoked-x509-root-certs`
- REST API: 
    -   GET `/dcl/pki/revoked-root-certificates`    

#### GET_ALL_X509_CERTS
**Status: Implemented**

Gets all certificates (root, intermediate and leaf).

Revoked certificates are not returned. 
Use `GET_ALL_REVOKED_X509_CERTS` to get a list of all revoked certificates. 

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command: 
    -   `dcld query pki all-x509-certs`
- REST API: 
    - GET `/dcl/pki/certificates`


#### GET_ALL_REVOKED_X509_CERTS
**Status: Implemented**

Gets all revoked certificates (both root and non-root).
   
Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command: 
    -   `dcld query pki all-revoked-x509-certs`
- REST API: 
    -   GET `/dcl/pki/revoked-certificates`

#### GET_ALL_PROPOSED_X509_ROOT_CERTS
**Status: Implemented**

Gets all proposed but not approved root certificates.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command: 
    -   `dcld query pki all-proposed-x509-root-certs`
- REST API: 
    -   GET `dcl/pki/proposed-certificates`

#### GET_ALL_PROPOSED_X509_ROOT_CERTS_TO_REVOKE
**Status: Implemented**

Gets all proposed but not approved root certificates to be revoked.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command: 
    -   `dcld query pki all-proposed-x509-root-certs-to-revoke`
- REST API:
    -   GET `/dcl/pki/proposed-revocation-certificates`




    
## AUTH

#### PROPOSE_ADD_ACCOUNT
**Status: Implemented**

Proposes a new Account with the given address, public key and role.

If more than 1 Trustee signature is required to add the account, the account
will be in a pending state until sufficient number of approvals is received.

- Parameters:
    - address: `string` - account address; Bench32 encoded
    - pub_key: `string` - account's Protobuf JSON encoded public key
    - vid: `optional(uint16)` - vendor id (only needed for vendor role)
    - roles: `array<string>` - the list of roles, comma-separated, assigning to the account. Supported roles: `Vendor`, `TestHouse`, `CertificationCenter`, `Trustee`, `NodeAdmin`. 
- In State: `dclauth/PendingAccount/value/<address>`
- Who can send: 
    - Trustee
- CLI command: 
    -   `dcld tx auth propose-add-account --address=<bench32 encoded string> --pubkey=<protobuf JSON encoded> --roles=<role1,role2,...> --vid=<uint16> --from=<account>`
    
#### APPROVE_ADD_ACCOUNT
**Status: Implemented**

Approves the proposed account.    

The account is not active until sufficient number of Trustees approve it. 

- Parameters:
    - address: `string` - account address; Bench32 encoded
- In State: `dclauth/Account/value/<address>`
- Who can send: 
    - Trustee
- Number of required approvals: 
    - 2/3 of Trustees
- CLI command: 
    -   `dcld tx auth approve-add-account --address=<bench32 encoded string> --from=<account>`
    
  
#### PROPOSE_REVOKE_ACCOUNT
**Status: Implemented**

Proposes revocation of the Account with the given address.

If more than 1 Trustee signature is required to revoke the account, the revocation
will be in a pending state until sufficient number of approvals is received.

- Parameters:
    - address: `string` - account address; Bench32 encoded
- In State: `dclauth/Account/value/<address>`
- Who can send: 
    - Trustee
- CLI command: 
    -   `dcld tx auth propose-revoke-account --address=<bench32 encoded string> --from=<account>`
    
#### APPROVE_REVOKE_ACCOUNT
**Status: Implemented**

Approves the proposed revocation of the account.    

The account is not revoked until sufficient number of Trustees approve it. 

- Parameters:
    - address: `string` - account address; Bench32 encoded
- In State: `dclauth/Account/value/<address>`
- Who can send: 
    - Trustee
- Number of required approvals: 
    - 2/3 of Trustees
- CLI command: 
    -   `dcld tx auth approve-revoke-account --address=<bench32 encoded string> --from=<account>`

#### GET_ACCOUNT
**Status: Implemented**

Gets an accounts by the address. Revoked accounts are not returned.

- Parameters:
    - address: `string` - account address; Bench32 encoded
- CLI command: 
    -   `dcld query auth account --addres <bench32 encoded string>`
- REST API: 
    -   GET `/dcl/auth/accounts/{address}`         

#### GET_PROPOSED_ACCOUNT
**Status: Implemented**

Gets a proposed but not approved accounts by its address

- Parameters:
    - address: `string` - account address; Bench32 encoded
- CLI command: 
    -   `dcld query auth proposed-account --address <bench32 encoded string>`
- REST API: 
    -   GET `/dcl/auth/proposed-accounts/{address}`       

#### GET_PROPOSED_ACCOUNT_TO_REVOKE
**Status: Implemented**

Gets a proposed but not approved accounts to be revoked by its address.

- Parameters:
    - address: `string` - account address; Bench32 encoded
- CLI command: 
    -   `dcld query auth proposed-account-to-revoke --address <bench32 encoded string>`
- REST API: 
    -   GET `/dcl/auth/proposed-revocation-accounts/{address}`      

#### GET_ALL_ACCOUNTS
**Status: Implemented**

Gets all accounts. Revoked accounts are not returned.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command: 
    -   `dcld query auth all-accounts`
- REST API: 
    -   GET `/dcl/auth/accounts`           
    
#### GET_ALL_PROPOSED_ACCOUNTS
**Status: Implemented**

Gets all proposed but not approved accounts.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command: 
    -   `dcld query auth all-proposed-accounts`
- REST API: 
    -   GET `/dcl/auth/proposed-accounts`
    
   
#### GET_ALL_PROPOSED_ACCOUNTS_TO_REVOKE
**Status: Implemented**

Gets all proposed but not approved accounts to be revoked.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command: 
    -   `dcld query auth all-proposed-accounts-to-revoke`
- REST API: 
    -   GET `/dcl/auth/proposed-revocation-accounts`
    

#### ROTATE_KEY
**Status: Not Implemented**

Rotate's the Account's public key by the owner.

- Who can send: 
    - Any role; owner

    
## VALIDATOR_NODE                      

#### ADD_VALIDATOR_NODE
**Status: Implemented**

Adds a new Validator node.

- Parameters:
    - pubkey: `string` - The validator's Protobuf JSON encoded public key
    - moniker: `string` - The validator's human-readable name
    - identity: `optional(string)` - identity signature (ex. UPort or Keybase)
    - website: `optional(string)` - The validator's website link
    - details: `optional(string)` - The validator's details
    - ip: `optional(string)` - The node's public IP
    - node-id: `optional(string)` - The node's ID
- In State: `validator/Validator/value/<owner-address>`
- Who can send: 
    - NodeAdmin
- CLI command: 
    -   `dcld tx validator add-node --pubkey=<protobuf JSON encoded> --moniker=<string> --from=<account>`

#### GET_VALIDATOR
**Status: Implemented**

Gets a validator node.

- Parameters:
    - address: `string` - Bench32 encoded validator address or owner account
- CLI command: 
    -   `dcld query validator node --address=<validator address|account>`
- REST API: 
    -   GET `/dcl/validator/nodes/{owner}`   

#### GET_ALL_VALIDATORS
**Status: Implemented**

Gets the list of all validator nodes from the store.

Note: All stored validator nodes (`active` and `jailed`) will be returned by default.
In order to get an active validator set use specific command [validator set](#validator-set).

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command: 
    -   `dcld query validator all-nodes`
- REST API: 
    -   GET `/dcl/validator/nodes`   
     

  
#### UPDATE_VALIDATOR_NODE
**Status: Not Implemented**

Updates the Validator node by the owner.  
`address` is used to reference the node, but can not be changed. 

- Parameters: 
    - address: `string` - Bench32 encoded validator address or owner account
    - moniker: `string` - The validator's human-readable name
    - identity: `optional(string)` - identity signature (ex. UPort or Keybase)
    - website: `optional(string)` - The validator's website link
    - details: `optional(string)` - The validator's details
    - ip: `optional(string)` - The node's public IP
    - node-id: `optional(string)` - The node's ID
- Who can send: 
    - NodeAdmin; owner

#### REMOVE_VALIDATOR_NODE
**Status: Not Implemented**

Deletes the Validator node (removes from the validator set) by the owner.

- Parameters:
    - address: `string` - Bench32 encoded validator address or owner account
- Who can send: 
    - NodeAdmin; owner

#### PROPOSE_REMOVE_VALIDATOR_NODE
**Status: Not Implemented**

Proposes removing the Validator node from the validator set by a Trustee. 

If more than 1 Trustee signature is required to remove a node, the removal
will be in a pending state until sufficient number of approvals is received.

- Parameters:
    - address: `string` - Bench32 encoded validator address or owner account
- Who can send: 
    - Trustee

#### APPROVE_REMOVE_VALIDATOR_NODE
**Status: Not Implemented**

Approves removing of the Validator node by a Trustee. 

The account is not removed until sufficient number of Trustees approve it. 

- Parameters:
    - address: `string` - Bench32 encoded validator address or owner account
- Who can send: 
    - Trustee
- Number of required approvals: 
    - 2/3 of Trustees
                      
#### UNJAIL_VALIDATOR_NODE
**Status: Not Implemented**

Approves unjail of the Validator node from jailed state and returning to the active validator state. 

If more than 1 Trustee approval is required to unjail a node, the node still
will be in a jailed state until sufficient number of approvals is received.

If 1 Trustee approval is required to unjail a nod or sufficient number of approvals is received,
the node will be unjailed and returned to the active validator set.

- Parameters:
    - address: `string` - Bench32 encoded validator address or owner account
- Who can send: 
    - Trustee
- Number of required approvals: 
    - 2/3 of Trustees
            
   
## Extensions    

#### Sign
Sign transaction by the given key.

- Parameters:
    - `txn` - transaction to sign.
    - `from` -  name or address of private key to use to sign.
    - `account-number` - (optional) the account number of the signing account.
    - `sequence` - (optional) the sequence number of the signing account.
    - `chain-id` - (optional) chain id.
- CLI command: 
    -   `dcld tx sign [path-to-txn-file] --from [address]`
Note: if `account_number` and `sequence`  are not specified they will be fetched from the ledger automatically.  
   
#### Broadcast
Broadcast transaction to the ledger.

- Parameters:
    - `txn` - transaction to broadcast
- CLI command: 
    -   `dcld tx broadcast [path-to-txn-file]`
- REST API: 
    - POST `/cosmos/tx/v1beta1/txs`  
  
#### Status
Query status of a node.

- Parameters:
    - `node`: optional(string) - node physical address to query (by default queries the node specified in CLI config file or else "tcp://localhost:26657")
- CLI command: 
    -   `dcld status [--node=<node ip>]`
- REST API: 
    - GET `/cosmos/base/tendermint/v1beta1/node_info`  

#### Validator set
Get the list of tendermint validators participating in the consensus at given height.

- Parameters:
    - `height`: optional(uint) - height to query (the latest by default)
- CLI command: 
    -   `dcld query tendermint-validator-set [height]`
- REST API: 
    - GET `/cosmos/base/tendermint/v1beta1/validatorsets/latest`
    - GET `/cosmos/base/tendermint/v1beta1/validatorsets/{height}`

#### Keys

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
