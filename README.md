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
- VID: `int16`
- PID: `int16`
- CID: `int16` (optional)
- Name: `string`
- Owner: `bech32 encoded address`
- Description: `string`
- SKU: `string`
- FirmwareVersion: `string`
- HardwareVersion: `string`
- Custom: `string` (optional)
- CertificateID: `string`
- CertifiedDate: `rfc3339 encoded date`
- TisOrTrpTestingCompleted: `bool`

Permissions:
- All the transactions below must be signed. Use `--from` flag.
- Signer must have `vendor` role. See `Authorization` module for details.

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

### Authorization

Roles:
- `administrator` - Is able to assign or revoke roles.
- `vendor`

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
          "administrator"
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
 