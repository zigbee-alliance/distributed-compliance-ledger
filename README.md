# ZB Ledger

## Build and run the app

To build and run, follow the tutorial: https://github.com/cosmos/sdk-tutorials/blob/master/nameservice/tutorial/21-build-run.md

Use __zbld__, __zblcli__ instead of __nsd__, __nscli__.

## Modules

### Compliance

ModelInfo type:
- ID: `string`
- Name: `string`
- Owner: `bech32 encoded address`
- Description: `string`
- SKU: `string`
- FirmwareVersion: `string`
- HardwareVersion: `string`
- CertificateID: `string`
- CertifiedDate: `rfc3339 encoded date`
- TisOrTrpTestingCompleted: `bool`

Permissions:
- All the transactions below must be signed. Use `--from` flag.
- Signer must be either `administrator` or `manufacturer` and record's owner. See `Authorization` module for details

Transactions:
- `zblcli tx compliance add-model-info [id:string] [name:string] [owner:bech32 encoded address] [description:string
] [sku:string] [firmware-version:string] [hardware-version:string][certificate-id:string] [certified-date:rfc3339
 encoded date] [tis-or-trp-testing-completed:bool]` - Add new ModelInfo.
  - Signature is required. Use `--from` flag.
- `zblcli tx compliance update-model-info update-model-info [id:string] [new-name:string] [new-owner:bech32 encoded
 address] [new-description:string] [new-sku:string] [new-firmware-version:string] [new-hardware-version:string] [new
 -certificate-id:string] [new-certified-date:rfc3339 encoded date] [new-tis-or-trp-testing-completed:bool]` - Update
  existing ModelInfo.
  - Signature is required. Use `--from` flag.
- `zblcli tx compliance delete-model-info [id:string]` - Delete existing ModelInfo.
  - Signature is required. Use `--from` flag.

Queries:
- `zblcli query compliance model-info [id]` - Query single ModelInfo.
- `zblcli query compliance model-info-with-proof [id]` - Query single ModelInfo with proof.
- `zblcli query compliance model-info-headers --skip [x] --take [y]` - Query list of ModelInfo headers. Flags are
 optional.

Examples:
- `zblcli tx compliance add-model-info "b4a3b939-c5ab-42b5-a163-928b3b147f9f" "TestName
" "cosmos1g4936hdq8mr5p6vs0qevdvxuvgpfpesh86cvc7" "Test description" "id1, id2" "1.2.3" "3.2.1" "cert34" "2020-01
-24T14:04:21+03:00" true --from jack `

Genesis:

- Use `nsd add-genesis-account` to add users to genesis.

### Authorization

Roles:
- `administrator` - Is able to assign or revoke roles.
- `manufacturer`

Commands:
- `zblcli tx assign-role [address] [role]` - Assign role to specified account.
  - Administrator's signature is required. Use `--from` flag.
- `zblcli tx revoke-role [address] [role]` - Revoke role from specified account.
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
