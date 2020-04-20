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

### Model Info

ModelInfo type:
- vid: `int16`
- pid: `int16`
- cid: `int16` (optional)
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
- `zblcli tx modelinfo add-model [vid:int16] [pid:int16] [name:string] [description:string] [sku:string] 
[firmware-version:string] [hardware-version:string] [tis-or-trp-testing-completed:bool]` - Add new ModelInfo.
  - Signature is required. Use `--from` flag.
  - Optional flags: 
    - `--cid` int16
    - `--custom` string

  Example: `zblcli tx modelinfo add-model 1 1 "Device #1" "Device Description" "SKU12FS" "1.0" "2.0" true --from jack`
  
  Example: `zblcli tx modelinfo add-model 1 2 "Device #2" "Device Description" "SKU324S" "2.0" "2.0" true --from jack --cid 1 --custom "Some Custom information" --certificate-id "ID123" --certified-date "2020-01-01T00:00:00Z"`

- `zblcli tx modelinfo update-model [vid:int16] [pid:int16] [tis-or-trp-testing-completed:bool]` - Update
  existing ModelInfo.
  - Signature is required. Use `--from` flag.
  - Optional flags: 
    - `--cid` int16
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
- vid: `int16` - vendor id
- pid: `int16` - product id
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
  
Queries:
- `zblcli query compliancetest test-result [vid] [pid]` - Query Testing Results associated with VID/PID.

  Example: `zblcli query compliancetest test-result 1 1`

### Compliance

Compliance Info type:
- vid: `int16` - vendor id
- pid: `int16` - product id
- state: `string` - current compliance state: either `certified` or `revoked`
- date:(optional) `rfc3339 encoded date` - depending on the state either certification date or revocation date 
- certification_type:(optional) `string`  - `zb` is the default and the only supported value now
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
- ` zblcli tx compliance certify-model [vid] [pid] [certification-date]` - Certify model.
  - Signature is required. Use `--from` flag.
  - Optional flags: 
    - `--certification-type` string -  `zb` is the default and the only supported value now
    - `--reason` string -  an optional comment describing the reason of certification

  Example: `zblcli tx compliance certify-model 1 1 "2020-04-16T06:04:57.05Z" --from jack`
  
  Example: `zblcli tx compliance certify-model 1 1 "2020-04-16T06:04:57.05Z" --certification-type "zb" --from jack`
 
- ` zblcli tx compliance revoke-model [vid] [pid] [revocation-date]` - Revoke certification for a model.
  - Signature is required. Use `--from` flag.
  - Optional flags: 
    - `--certification-type` string -  `zb` is the default and the only supported value now
    - `--reason` string -  an optional comment describing the reason of revocation

  Example: `zblcli tx compliance revoke-model 1 1 "2020-04-16T06:04:57.05Z" --from jack`
  
  Example: `zblcli tx compliance revoke-model 1 1 "2020-04-16T06:04:57.05Z" --reason "Some Reason" --from jack`
  
Queries:
- `zblcli query compliance certified-model [vid] [pid]` - Query certification data for model associated with VID/PID.

  Example: `zblcli query compliance certified-model 1 1`
  
- `zblcli query compliance all-certified-models` - Query all certified models.

  Example: `zblcli query compliance all-certified-models`
  
- `zblcli query compliance revoked-model [vid] [pid]` - Query revocation data for model associated with VID/PID.

  Example: `zblcli query compliance revoked-model 1 1`
  
- `zblcli query compliance all-certified-models` - Query all revoked models.

  Example: `zblcli query compliance all-revoked-models`
  
- `zblcli query compliance compliance-info [vid] [pid]` - Query compliance info for model associated with VID/PID.

  Example: `zblcli query compliance compliance-info 1 1`
  
- `zblcli query compliance all-compliance-info-records` - Query all compliance-infos.

  Example: `zblcli query compliance all-compliance-info-records`

### Authorization

Roles:
- `Administrator` - Is able to assign or revoke roles.
- `Vendor` - Is able to add models
- `TestHouse` - Is able to add testing results for an model
- `ZBCertificationCenter` - Is able to certify models

Commands:
- `zblcli tx authz assign-role [address] [role]` - Assign role to specified account.
  - Administrator's signature is required. Use `--from` flag.
- `zblcli tx authz revoke-role [address] [role]` - Revoke role from specified account.
  - Administrator's signature is required. Use `--from` flag.

Genesis template:
```json
{
  "app_state": {
    "authz": {
      "account_roles": [{
        "address": "cosmos1j8x9urmqs7p44va5p4cu29z6fc3g0cx2c2vxx2",
        "roles": [
          "Administrator"
        ]
      }]
    }
  }
}
```

### Authentication extensions

Queries:

- `zblcli query authnext account-headers --skip [x] --take [y]` - The command to list account headers with roles. Flags
 are optional.
 