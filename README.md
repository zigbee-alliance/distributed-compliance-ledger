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
- `zblcli tx pki propose-add-x509-root-cert --certificate=<string-or-path> --from=<account>` - Proposes a new self-signed root certificate.

  Example: `zblcli tx pki propose-add-x509-root-cert --certificate="/path/to/certificate/file" --from=jack`
  
  Example: `zblcli tx pki propose-add-x509-root-cert --certificate="----BEGIN CERTIFICATE----- ......" --from=jack`

- `zblcli tx pki approve-add-x509-root-cert --subject=<string> --subject-key-id=<hex string> --from=<account>` - Approves the proposed root certificate.

  Example: `zblcli tx pki approve-add-x509-root-cert --subject="CN=dsr-corporation.com" --subject-key-id="8A:34:B:5C:D8:42:18:F2:C1:2A:AC:7A:B3:8F:6E:90:66:F4:4E:5C" --from=jack`
  
- `zblcli tx pki add-x509-cert --certificate=<string-or-path> --from=<account>` - Adds an intermediate or leaf X509 certificate signed by a chain of certificates which must be already present on the ledger.

  Example: `zblcli tx pki add-x509-cert --certificate="/path/to/certificate/file" --from=jack`
  
  Example: `zblcli tx pki add-x509-cert --certificate="----BEGIN CERTIFICATE----- ......" --from=jack`  
    
Queries:
- `zblcli query pki all-proposed-x509-root-certs` - Gets all proposed but not approved root certificates.
  - Optional flags: 
    - `--skip` int
    - `--take` int
    
  Example: `zblcli query pki all-proposed-x509-root-certs`
  
- `zblcli query pki proposed-x509-root-cert --subject=<string> --subject-key-id=<hex string>` - Gets a proposed but not approved root certificate.

  Example: `zblcli query pki proposed-x509-root-cert --subject="CN=dsr-corporation.com" --subject-key-id="8A:34:B:5C:D8:42:18:F2:C1:2A:AC:7A:B3:8F:6E:90:66:F4:4E:5C"`

- `zblcli query pki all-x509-root-certs` - Gets all approved root certificates.
  - Optional flags: 
    - `--skip` int
    - `--take` int

  Example: `zblcli query pki all-x509-root-certs`
  
- `zblcli query pki x509-cert --subject=<string> --subject-key-id=<hex string>` - Gets a certificates (either root, intermediate or leaf).

  Example: `zblcli query pki x509-certs --subject="CN=dsr-corporation.com" --subject-key-id="8A:34:B:5C:D8:42:18:F2:C1:2A:AC:7A:B3:8F:6E:90:66:F4:4E:5C"`
  
- `zblcli query pki all-x509-certs` - Gets all certificates (root, intermediate and leaf).
  - Optional flags: 
    - `--skip` int
    - `--take` int
    - `--root-subject` string
    - `--root-subject-key-id` string

  Example: `zblcli query pki x509-certs`
  
  Example: `zblcli query pki x509-certs --root-subject-key-id="8A:34:B:5C:D8:42:18:F2:C1:2A:AC:7A:B3:8F:6E:90:66:F4:4E:5C"`
  
  Example: `zblcli query pki x509-certs --root-subject="CN=dsr-corporation.com"`
  
- `zblcli query pki all-subject-x509-certs --subject=<subject>` - Gets all certificates (root, intermediate and leaf) associated with subject.
  - Optional flags: 
    - `--skip` int
    - `--take` int
    - `--root-subject` string
    - `--root-subject-key-id` string

  Example: `zblcli query pki all-subject-x509-certs --subject="CN=dsr"`

  Example: `zblcli query pki all-subject-x509-certs --subject="CN=dsr" --root-subject-key-id="8A:34:B:5C:D8:42:18:F2:C1:2A:AC:7A:B3:8F:6E:90:66:F4:4E:5C"`
  
  Example: `zblcli query pki all-subject-x509-certs --subject="CN=dsr" --root-subject="CN=dsr-corporation.com"`

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
- `zblcli tx modelinfo add-model --vid=<uint16> --pid=<uint16> --name=<string> --description=<string or path> --sku=<string> 
--firmware-version=<string> --hardware-version=<string> --tis-or-trp-testing-completed=<bool> --from=<account>` - Add new ModelInfo.
  - Optional flags: 
    - `--cid` uint16
    - `--custom` string

  Example: `zblcli tx modelinfo add-model --vid=1 --pid=1 --name="Device #1" --description="Device Description" --sku="SKU12FS" --firmware-version="1.0" --hardware-version="2.0" --tis-or-trp-testing-completed=true --from=jack`
  
  Example: `--vid=1 --pid=1 --name="Device #1" --description="Device Description" --sku="SKU12FS" --firmware-version="1.0" --hardware-version="2.0" --tis-or-trp-testing-completed=true --from=jack --cid=1 --custom="Some Custom information"`

- `zblcli tx modelinfo update-model --vid=<uint16> --pid=<uint16> --tis-or-trp-testing-completed=<bool> --from=<account>` - Update
  existing ModelInfo.
  - Optional flags: 
    - `--cid` uint16
    - `--custom` string
    - `--description` string
    
  Example: `zblcli tx modelinfo update-model --vid=1 --pid=1 --tis-or-trp-testing-completed=true --from=jack --description="New Description"`
  
  Example: `zblcli tx modelinfo update-model --vid=1 --pid=1 --tis-or-trp-testing-completed=true --from=jack --custom="Custom Data"`

Queries:
- `zblcli query modelinfo model --vid=<uint16> --pid=<uint16>` - Query single ModelInfo.

  Example: `zblcli query modelinfo model --vid=1 --pid=1`
  
- `zblcli query modelinfo all-models` - Query list of ModelInfos. Optional flags: 
    - `--skip` int
    - `--take` int
    
  Example: `zblcli query modelinfo all-models`

- `zblcli query modelinfo vendors` - Query list of Vendors. Optional flags: 
    - `--skip` int
    - `--take` int
    
  Example: `zblcli query modelinfo vendors`
  
- `zblcli query modelinfo vendor-models --vid=<uint16>` - Query list of ModelInfos for the given Vendor. Optional flags: 
    - `--skip` int
    - `--take` int

  Example: `zblcli query modelinfo vendor-models --vid=1`

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
- ` zblcli tx compliancetest add-test-result --vid=<uint16> --pid=<uint16> --test-result=<string> --test-date=<rfc3339 encoded date> --from=<account>` - Add new Testing Result.

  Example: `zblcli tx compliancetest add-test-result --vid=1 --pid=1 --test-result="Test Document" --test-date="2020-04-16T06:04:57.05Z" --from=jack`
  
  Example: `zblcli tx compliancetest add-test-result --vid=1 --pid=1 --test-result="path/to/document" --test-date="2020-04-16T06:04:57.05Z" --from=jack`
  
Queries:
- `zblcli query compliancetest test-result --vid=<uint16> --pid=<uint16>` - Query Testing Results associated with VID/PID.

  Example: `zblcli query compliancetest test-result --vid=1 --pid=1`

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
- ` zblcli tx compliance certify-model --vid=<uint16> --pid=<uint16> --certification-type=<zb> --certification-date=<rfc3339 encoded date> --from=<account>` - Certify model.
    - Optional flags: 
    - `--reason` string -  an optional comment describing the reason of certification

  Example: `zblcli tx compliance certify-model --vid=1 --pid=1 --certification-type="zb" --certification-date=<"2020-04-16T06:04:57.05Z" --from=jack`
 
- ` zblcli tx compliance revoke-model --vid=<uint16> --pid=<uint16> --certification-type=<zb> --revocation-date=<rfc3339 encoded date> --from=<account>` - Revoke certification for a model.
  - Signature is required. Use `--from` flag.
  - Optional flags: 
    - `--reason` string -  an optional comment describing the reason of revocation

  Example: `zblcli tx compliance revoke-model --vid=1 --pid=1 --certification-type="zb" --revocation-date="2020-04-16T06:04:57.05Z" --from=jack`
  
  Example: `zblcli tx compliance revoke-model --vid=1 --pid=1 --certification-type="zb" --revocation-date="2020-04-16T06:04:57.05Z" --reason "Some Reason" --from=jack`
  
Queries:
- `zblcli query compliance certified-model --vid=<uint16> --pid=<uint16> --certification-type=<zb>` - Query certification data for model associated with VID/PID.

  Example: `zblcli query compliance certified-model --vid=1 --pid=1 --certification-type="zb"`
  
- `zblcli query compliance all-certified-models` - Query all certified models.

  Example: `zblcli query compliance all-certified-models`
  
- `zblcli query compliance revoked-model --vid=<uint16> --pid=<uint16> --certification-type=<zb>` - Query revocation data for model associated with VID/PID.

  Example: `zblcli query compliance revoked-model --vid=1 --pid=1 --certification-type="zb"`
  
- `zblcli query compliance all-revoked-models` - Query all revoked models.

  Example: `zblcli query compliance all-revoked-models`
  
- `zblcli query compliance compliance-info --vid=<uint16> --pid=<uint16> --certification-type=<zb>` - Query compliance info for model associated with VID/PID.

  Example: `zblcli query compliance compliance-info --vid=1 --pid=1 --certification-type="zb"`
  
- `zblcli query compliance all-compliance-info-records` - Query all compliance-infos.

  Example: `zblcli query compliance all-compliance-info-records`

### Authorization

Permissions:
- All the transactions below must be signed. Use `--from` flag.
- Signer must have `Trustee` role. 

Roles:
- `Administrator` - Is able to assign or revoke roles.
- `Vendor` - Is able to add models
- `TestHouse` - Is able to add testing results for an model
- `ZBCertificationCenter` - Is able to certify models
- `Trustee` - Is able to approve root certificates

Transactions:
- `zblcli tx authz assign-role --address=<bench32 address> --role=<string> --from=<account>` - Assign role to specified account.

  Example: `zblcli tx authz assign-role --address=cosmos15ljvz60tfekhstz8lcyy0c9l8dys5qa2nnx4d7 --role=Vendor --from=jack`
  
- `zblcli tx authz revoke-role --address=<bench32 address> --role=<string> --from=<account>` - Revoke role from specified account.
  - Trustee's signature is required.

  Example: `zblcli tx authz revoke-role --address=cosmos15ljvz60tfekhstz8lcyy0c9l8dys5qa2nnx4d7 --role=Vendor --from=jack`

Queries:

- `zblcli query authz account-roles --address=<bench32 address>` - The command to query roles by account address.

  Example: `zblcli query authz account-roles --address=cosmos15ljvz60tfekhstz8lcyy0c9l8dys5qa2nnx4d7`

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

Permissions:
- All the transactions below must be signed. Use `--from` flag.
- Signer must have `Trustee` role. See `Authorization` module for details.

Transactions:

- `zblcli tx authnext create-account --address=<bench32 address> --pubkey=<bench32 pubkey> --from=<account>` - The command to creates a new account.

  Example: `zblcli tx authnext create-account --address=cosmos15ljvz60tfekhstz8lcyy0c9l8dys5qa2nnx4d7 --pubkey=cosmospub1addwnpepqtrnrp93hswlsrzvltc3n8z7hjg9dxuh3n4rkp2w2verwfr8yg27c95l4k3  --from=jack`

Queries:

- `zblcli query authnext account --address=<bench32 address>` - The command to query single account.

  Example: `zblcli query authnext account --address=cosmos15ljvz60tfekhstz8lcyy0c9l8dys5qa2nnx4d7`

- `zblcli query authnext accounts --skip [x] --take [y]` - The command to list account headers with roles. Flags
 are optional.
 
  Example: `zblcli query authnext accounts`

 