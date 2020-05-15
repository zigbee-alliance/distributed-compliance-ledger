# ZB Ledger

## Build and run the app

To build and run, follow the tutorial: https://github.com/cosmos/sdk-tutorials/blob/master/nameservice/tutorial/21-build-run.md

Use __zbld__, __zblcli__ instead of __nsd__, __nscli__.

## Localnet

- To start localnet using docker-compose run `make install && make localnet_init && make localnet_start`
  - 4 nodes will be started and will expose their RPC enpoints on ports `26657`, `26659`, `26661`, `26662`
- To stop localnet run `make localnet_stop`

## Deployment

- Read more about deployment in `ansible/readme.md`.

## Docs
- Requests and transactions: [transactions.md](docs/transactions.md)
- Use cases:
    - [PKI](docs/use_cases_pki.png)
    - [Device on-ledger certification](docs/use_cases_device_on_ledger_certification.png)
    - [Device off-ledger certification](docs/use_cases_device_off_ledger_certification.png)
    - [Auth](docs/use_cases_txn_auth.png)
    - [Validators](docs/use_cases_add_validator_node.png)

## Modules

Some of the modules are being refactored against [transactions.md](docs/transactions.md) and may look
a bit different than specified below.

### PKI

Proposed Certificate type:
- pem_cert: `string` - pem encoded certificate
- subject: `string` - certificates's `Subject`
- subject_key_id: `string` - certificates's `Subject Key Id`
- serial_number: `string` - certificates's `Serial Number`
- approvals: `string` - Trustees addresses who approved root certificate
- owner: `bech32 encoded address` - the address used for sending the original message

Certificate type:
- pem_cert: `string` - pem encoded certificate
- subject: `string` - certificates's `Subject`
- subject_key_id: `string` - certificates's `Subject Key Id`
- serial_number: `string` - certificates's `Serial Number`
- root_subject: `string` - root certificates's `Subject`
- root_subject_key_id: `string` - root certificates's `Subject Key Id`
- type: `string` - certificate type: either root or intermediate
- owner: `bech32 encoded address` - the address used for sending the original message

Permissions:
- All the transactions below must be signed. Use `--from` flag.
- Propose Root Certificate: Any Role.
- Approve Root Certificate: Signer must have `Trustee` role. See `Authorization` module for details.
- Add Leaf Certificate: Any Role.

Transactions:
- ` zblcli tx pki propose-add-x509-root-cert [certificate-path-or-pem-string]` - Proposes a new self-signed root certificate.
  - Signature is required. Use `--from` flag.

  Example: `zblcli tx pki propose-add-x509-root-cert "/path/to/certificate/file" --from jack`
  
  Example: `zblcli tx pki propose-add-x509-root-cert "----BEGIN CERTIFICATE----- ......" --from jack`

- ` zblcli tx pki approve-add-x509-root-cert [subject] [subject-key-id]` - Approves the proposed root certificate.
  - Signature is required. Use `--from` flag.

  Example: `zblcli tx pki approve-add-x509-root-cert "CN=dsr-corporation.com" "8A:34:B:5C:D8:42:18:F2:C1:2A:AC:7A:B3:8F:6E:90:66:F4:4E:5C" --from jack`
  
- ` zblcli tx pki add-x509-cert [certificate-path-or-pem-string]` - Adds an intermediate or leaf X509 certificate signed by a chain of certificates which must be already present on the ledger.
  - Signature is required. Use `--from` flag.

  Example: `zblcli tx pki add-x509-cert "/path/to/certificate/file" --from jack`
  
  Example: `zblcli tx pki add-x509-cert "----BEGIN CERTIFICATE----- ......" --from jack`  
    
Queries:
- `zblcli query pki all-proposed-x509-root-certs` - Gets all proposed but not approved root certificates.
  - Optional flags: 
    - `--skip` int
    - `--take` int
    
  Example: `zblcli query pki all-proposed-x509-root-certs`
  
- `zblcli query pki proposed-x509-root-cert [subject] [subject-key-id]` - Gets a proposed but not approved root certificate.

  Example: `zblcli query pki proposed-x509-root-cert "CN=dsr-corporation.com" "8A:34:B:5C:D8:42:18:F2:C1:2A:AC:7A:B3:8F:6E:90:66:F4:4E:5C"`

- `zblcli query pki all-x509-root-certs` - Gets all approved root certificates.
  - Optional flags: 
    - `--skip` int
    - `--take` int

  Example: `zblcli query pki all-x509-root-certs`
  
- `zblcli query pki x509-cert [subject] [subject-key-id]` - Gets a certificates (either root, intermediate or leaf).

  Example: `zblcli query pki x509-certs "CN=dsr-corporation.com" "8A:34:B:5C:D8:42:18:F2:C1:2A:AC:7A:B3:8F:6E:90:66:F4:4E:5C"`
  
- `zblcli query pki all-x509-certs` - Gets all certificates (root, intermediate and leaf).
  - Optional flags: 
    - `--skip` int
    - `--take` int
    - `--root-subject` string
    - `--root-subject-key-id` string

  Example: `zblcli query pki x509-certs`
  
  Example: `zblcli query pki x509-certs --root-subject-key-id "8A:34:B:5C:D8:42:18:F2:C1:2A:AC:7A:B3:8F:6E:90:66:F4:4E:5C"`
  
  Example: `zblcli query pki x509-certs --root-subject "CN=dsr-corporation.com"`
  
- `zblcli query pki all-subject-x509-certs` - Gets all certificates (root, intermediate and leaf) associated with subject.
  - Optional flags: 
    - `--skip` int
    - `--take` int
    - `--root-subject` string
    - `--root-subject-key-id` string

  Example: `zblcli query pki all-subject-x509-certs`

  Example: `zblcli query pki all-subject-x509-certs --root-subject-key-id "8A:34:B:5C:D8:42:18:F2:C1:2A:AC:7A:B3:8F:6E:90:66:F4:4E:5C"`
  
  Example: `zblcli query pki all-subject-x509-certs --root-subject "CN=dsr-corporation.com"`

### Model Info

ModelInfo type:
- vid: `uint16`
- pid: `uint16`
- cid: `uint16` (optional)
- name: `string`
- owner: `bech32 encoded address`
- description: `string`
- sku: `string`
- firmware_version: `string`
- hardware_version: `string`
- custom: `string` (optional)
- tis_or_trp_testing_completed: `bool`

Permissions:
- All the transactions below must be signed. Use `--from` flag.
- Signer must have `Vendor` role. See `Authorization` module for details.

Transactions:
- `zblcli tx modelinfo add-model [vid:uint16] [pid:uint16] [name:string] [description:string] [sku:string] 
[firmware-version:string] [hardware-version:string] [tis-or-trp-testing-completed:bool]` - Add new ModelInfo.
  - Signature is required. Use `--from` flag.
  - Optional flags: 
    - `--cid` uint16
    - `--custom` string

  Example: `zblcli tx modelinfo add-model 1 1 "Device #1" "Device Description" "SKU12FS" "1.0" "2.0" true --from jack`
  
  Example: `zblcli tx modelinfo add-model 1 2 "Device #2" "Device Description" "SKU324S" "2.0" "2.0" true --from jack --cid 1 --custom "Some Custom information" --certificate-id "ID123" --certified-date "2020-01-01T00:00:00Z"`

- `zblcli tx modelinfo update-model [vid:uint16] [pid:uint16] [tis-or-trp-testing-completed:bool]` - Update
  existing ModelInfo.
  - Signature is required. Use `--from` flag.
  - Optional flags: 
    - `--cid` uint16
    - `--custom` string
    - `--description` string
    
  Example: `zblcli tx modelinfo update-model 1 1 true --from jack --description "New Description"`
  
  Example: `zblcli tx modelinfo update-model 1 1 true --from jack --custom "Custom Data"`

Queries:
- `zblcli query modelinfo model [vid] [pid]` - Query single ModelInfo.

  Example: `zblcli query modelinfo model 1 1`
  
- `zblcli query modelinfo all-models` - Query list of ModelInfos. Optional flags: 
    - `--skip` int
    - `--take` int
    
  Example: `zblcli query modelinfo all-models`

- `zblcli query modelinfo vendors` - Query list of Vendors. Optional flags: 
    - `--skip` int
    - `--take` int
    
  Example: `zblcli query modelinfo vendors`
  
- `zblcli query modelinfo vendor-models [vid]` - Query list of ModelInfos for the given Vendor. Optional flags: 
    - `--skip` int
    - `--take` int

  Example: `zblcli query modelinfo vendor-models 1`

Genesis:

- Use `nsd add-genesis-account` to add users to genesis.

### Compliance Test

Testing Result type:
- vid: `uint16` - vendor id
- pid: `uint16` - product id
- test_result: `string` - test result report. It can contain url, blob, etc..
- test_date: `rfc3339 encoded date` - the date of testing
- owner: `bech32 encoded address` - the address used for sending the original message

Permissions:
- All the transactions below must be signed. Use `--from` flag.
- Signer must have `TestHouse` role. See `Authorization` module for details.

Transactions:
- ` zblcli tx compliancetest add-test-result [vid] [pid] [test-result] [test_date]` - Add new Testing Result.
  - Signature is required. Use `--from` flag.

  Example: `zblcli tx compliancetest add-test-result 1 1 "Test Document" "2020-04-16T06:04:57.05Z" --from jack`
  
  Example: `zblcli tx compliancetest add-test-result 1 1 "path/to/document" "2020-04-16T06:04:57.05Z" --from jack`
  
Queries:
- `zblcli query compliancetest test-result [vid] [pid]` - Query Testing Results associated with VID/PID.

  Example: `zblcli query compliancetest test-result 1 1`

### Compliance

Compliance Info type:
- vid: `uint16` - vendor id
- pid: `uint16` - product id
- state: `string` - current compliance state: either `certified` or `revoked`
- date:(optional) `rfc3339 encoded date` - depending on the state either certification date or revocation date 
- certification_type: `string`  - `zb` is the default and the only supported value now
- reason:(optional) `string` - an optional comment describing the reason of action
- owner: `bech32 encoded address` - the address used for sending the original message
- history: array of items - contains the history of all state changes
    - state: `string` - either `certified` or `revoked` - previous state
    - date:(optional) `rfc3339 encoded date` - previous date

Permissions:
- All the transactions below must be signed. Use `--from` flag.
- Signer must have `ZBCertificationCenter` role. See `Authorization` module for details.
- Only owner can update an existing record. 

Transactions:
- ` zblcli tx compliance certify-model [vid] [pid] [certification-type] [certification-date]` - Certify model.
  - Signature is required. Use `--from` flag.
  - Optional flags: 
    - `--reason` string -  an optional comment describing the reason of certification

  Example: `zblcli tx compliance certify-model 1 1 "zb" "2020-04-16T06:04:57.05Z" --from jack`
 
- ` zblcli tx compliance revoke-model [vid] [pid] [certification-type] [revocation-date]` - Revoke certification for a model.
  - Signature is required. Use `--from` flag.
  - Optional flags: 
    - `--reason` string -  an optional comment describing the reason of revocation

  Example: `zblcli tx compliance revoke-model 1 1 "zb" "2020-04-16T06:04:57.05Z" --from jack`
  
  Example: `zblcli tx compliance revoke-model 1 1 "zb" "2020-04-16T06:04:57.05Z" --reason "Some Reason" --from jack`
  
Queries:
- `zblcli query compliance certified-model [vid] [pid] [certification-type]` - Query certification data for model associated with VID/PID.

  Example: `zblcli query compliance certified-model 1 1 "zb"`
  
- `zblcli query compliance all-certified-models` - Query all certified models.

  Example: `zblcli query compliance all-certified-models`
  
- `zblcli query compliance revoked-model [vid] [pid] [certification-type]` - Query revocation data for model associated with VID/PID.

  Example: `zblcli query compliance revoked-model 1 1 "zb"`
  
- `zblcli query compliance all-revoked-models` - Query all revoked models.

  Example: `zblcli query compliance all-revoked-models`
  
- `zblcli query compliance compliance-info [vid] [pid] [certification-type]` - Query compliance info for model associated with VID/PID.

  Example: `zblcli query compliance compliance-info 1 1 "zb"`
  
- `zblcli query compliance all-compliance-info-records` - Query all compliance-infos.

  Example: `zblcli query compliance all-compliance-info-records`

### Validator

Validator type:
- `validator_address`: `bech32 encoded address` - the tendermint validator address
- `validator_pubkey`: `bech32 encoded pubkey` - the tendermint validator public key
- `owner`: `bech32 encoded address` - the account address of validator owner (original sender of transaction)
- `power`: `int` - validator consensus power
- `jailed`: `bool` - has the validator been removed from validator set because of cheating
- `jailed_reason`: (optional) `string` - the reason of validator jailing
- `description`: 
    - `name`: `string` - validator name
    - `identity`:(optional) `string` - identity signature (ex. UPort or Keybase)
    - `website`:(optional) `string` - website link
    - `details`:(optional) `string` - additional details
Permissions:
- All the transactions below must be signed. Use `--from` flag.
- Signer must have `NodeAdmin` role. See `Authorization` module for details.

Transactions:
- ` zblcli tx validator add-node --validator-address=<address> --validator-pubkey=<pubkey> --name=<node name> --from=<account>` - Add a new Validator node.
  - Signature is required. Use `--from` flag.
  - Optional flags: 
    - `--website` string -  optional validator's website
    - `--identity` string -  optional identity signature
    - `--details` string -  optional validator's details
    - `--node-id` string -  optional node's ID
    - `--ip` string -  node's public IP. It takes effect only when used in combination with `--generate-only` flag

  Example: `zblcli tx validator add-node --validator-address=cosmosvalcons1tl46nm39xtuutvw2wqaeyyd6csknfe0a7xqnrw --validator-pubkey=cosmosvalconspub1zcjduepqn5nz4c8n5jwfmgd6tqfqzu8arpne3au4g7tfsz33g8y6dcvhkf4sw054j8 --name=node1 --from=jack`
 
Queries:
- `zblcli query validator node --validator-address=<address>` - Query validator node by given validator address.

  Example: `zblcli query validator node --validator-address=cosmosvaloper1jnr3hcrvcpvqm5fdcafg70azkg0awf96tvu79z`
  
- `zblcli query validator all-nodes` - Query all validator nodes.

  Example: `zblcli query validator all-nodes`
  
### Authorization

Roles:
- `Administrator` - Is able to assign or revoke roles
- `Vendor` - Is able to add models
- `TestHouse` - Is able to add testing results for an model
- `ZBCertificationCenter` - Is able to certify models
- `Trustee` - Is able to approve root certificates
- `NodeAdmin` - Is able to add nodes to validator pool

Transactions:
- `zblcli tx authz assign-role [address] [role]` - Assign role to specified account.
  - Trustee's signature is required. Use `--from` flag.

  Example: `zblcli tx authz assign-role cosmos15ljvz60tfekhstz8lcyy0c9l8dys5qa2nnx4d7 Vendor --from jack`
  
- `zblcli tx authz revoke-role [address] [role]` - Revoke role from specified account.
  - Trustee's signature is required. Use `--from` flag.

  Example: `zblcli tx authz revoke-role cosmos15ljvz60tfekhstz8lcyy0c9l8dys5qa2nnx4d7 Vendor --from jack`

Queries:

- `zblcli query authz account-roles [account]` - The command to query roles by account address.

  Example: `zblcli query authz account-roles cosmos15ljvz60tfekhstz8lcyy0c9l8dys5qa2nnx4d7`

Genesis template:
```json
{
  "app_state": {
    "authz": {
      "account_roles": [{
        "address": "cosmos1j8x9urmqs7p44va5p4cu29z6fc3g0cx2c2vxx2",
        "roles": [
          "Trustee"
        ]
      }]
    }
  }
}
```

### Authentication extensions

Transactions:

- `zblcli tx authnext create-account [account] [public-key]` - The command to creates a new account.
  - Signature is required. Use `--from` flag.

  Example: `zblcli tx authnext create-account cosmos15ljvz60tfekhstz8lcyy0c9l8dys5qa2nnx4d7 cosmospub1addwnpepqtrnrp93hswlsrzvltc3n8z7hjg9dxuh3n4rkp2w2verwfr8yg27c95l4k3  --from jack`

Queries:

- `zblcli query authnext account [account]` - The command to query single account.

  Example: `zblcli query authnext account cosmos15ljvz60tfekhstz8lcyy0c9l8dys5qa2nnx4d7`

- `zblcli query authnext accounts --skip [x] --take [y]` - The command to list account headers with roles. Flags
 are optional.
 
  Example: `zblcli query authnext accounts`

