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

Transactions:   // TODO
- `zblcli tx compliance add-model-info` - Add new ModelInfo.
  - Owner's signature is required. Use `--from` flag.
- `zblcli tx compliance update-model-info` - Update existing ModelInfo.
- `zblcli tx compliance delete-model-info` - Delete existing ModelInfo.

Queries
- `zblcli query compliance model-info`
- `zblcli query compliance model-info-with-proof`
- `zblcli query compliance model-info-headers`

### Authorization

Commands:
- `zblcli query account-headers --skip x --take y` - The command to query list of account headers. Flags are optional.

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

- `zblcli query authnext account-headers --skip x --take y` - The command to query list of account headers. Flags are
 optional.
