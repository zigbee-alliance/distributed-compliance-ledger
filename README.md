# ZB Ledger

## Build and run the app

To build and run, follow the tutorial: https://github.com/cosmos/sdk-tutorials/blob/master/nameservice/tutorial/21-build-run.md

Use __zbld__, __zblcli__ instead of __nsd__, __nscli__.

## Modules

### Compliance

ModelInfo type:
- `ID: string`
- `Family: string`
- `Cert: string`
- `Owner: sdk.AccAddress`

Transactions:
- `zblcli tx compliance add-model-info [id:string] [family:string] [certificate:string] [owner:Bech32Addr]` - Add new
 ModelInfo.
  - Signature is required. Use `--from` flag.
- `zblcli tx compliance update-model-info [id:string] [new-family:string] [new-certificate:string] [new-owner:Bech32Addr
]` - Update existing ModelInfo.
  - Signature is required. Use `--from` flag.
- `zblcli tx compliance delete-model-info [id:string]` - Delete existing ModelInfo.
  - Signature is required. Use `--from` flag.

All the transactions above require the signer to be either `administrator` or `manufacturer` and record's owner.

Queries:
- `zblcli query compliance model-info [id]` - Query single ModelInfo.
- `zblcli query compliance model-info-with-proof [id]` - Query single ModelInfo with proof.
- `zblcli query compliance model-info-headers --skip [x] --take [y]` - Query list of ModelInfo headers. Flags are
 optional.

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

## Localnet

To start localnet:

- make
- make localnet_init
- make localnet_up