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
- certificate_id: `string`
- certified_date: `rfc3339 encoded date`
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
- vid: `int16`
- pid: `int16`
- test_result: `string`
- owner: `bech32 encoded address`
- created_at: `datetime`

Permissions:
- All the transactions below must be signed. Use `--from` flag.
- Signer must have `TestHouse` role. See `Authorization` module for details.

Transactions:
- ` zblcli tx compliancetest add-test-result [vid] [pid] [test-result]` - Add new Testing Result.
  - Signature is required. Use `--from` flag.

  Example: `zblcli tx compliancetest add-test-result 1 1 "Test Document" --from jack`
  
Queries:
- `zblcli query compliancetest test-result [vid] [pid]` - Query Testing Results associated with VID/PID.

  Example: `zblcli query compliancetest test-result 1 1`

### Compliance

Certified Models type:
- vid: `int16`
- pid: `int16`
- certification_date: `rfc3339 encoded date`
- certification_type:(optional) `string`  - zb is the default and the only supported value now
- owner: `bech32 encoded address`

Permissions:
- All the transactions below must be signed. Use `--from` flag.
- Signer must have `ZBCertificationCenter` role. See `Authorization` module for details.

Transactions:
- ` zblcli tx compliance certify-model [vid] [pid] [certification-date]` - Certify model.
  - Signature is required. Use `--from` flag.
  - Optional flags: 
    - `--certification-type` string

  Example: `zblcli tx compliance certify-model 1 1 "2020-04-16T06:04:57.05Z" --from jack`
  
  Example: `zblcli tx compliance certify-model 1 1 "2020-04-16T06:04:57.05Z" --certification-type "zb" --from jack`
  
Queries:
- `zblcli query compliance certified-model [vid] [pid]` - Query certification data for model associated with VID/PID.

  Example: `zblcli query compliance certified-model 1 1`
  
- `zblcli query compliance all-certified-models` - Query all certified models.

  Example: `zblcli query compliance all-certified-models`

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
 