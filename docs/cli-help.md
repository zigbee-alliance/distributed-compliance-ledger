# CLI Help

This document contains descriptions and examples of CLI commands.

Please configure the CLI before using (see `CLI Configuration` section in [how-to.md](how-to.md#cli-configuration)).

If write requests to the Ledger needs to be sent, please make sure that you have
an Account created on the Ledger with an appropriate role (see `Getting Account` section in [how-to.md](how-to.md#getting-account)).

All the transactions (write requests) must be signed. Use `--from` flag.

### Keys

The set of commands that allows you to manage your local keystore.

Commands:
- Derive a new private key and encrypt to disk. 
You will be prompted to enter encryption passphrase. 
This passphrase will be requested each time you send write transactions on the ledger using this key.

  Command: `dclcli keys add <key name>`

  Example: `dclcli keys add jack`

- Recover existing key using seed instead of creating a new one.
You will be prompted to enter encryption passphrase and seed.
This passphrase will be requested each time you send write transactions on the ledger using this key.

  Command: `dclcli keys add <key name> --recover`

  Example: `dclcli keys add jack --recover`

- Get a list of all stored public keys. 

  Command: `dclcli keys list`

  Example: `dclcli keys list`

- Get details for a key.

  Command: `dclcli keys show <key name>`

  Example: `dclcli keys show jack`

  
### Authorization

The set of commands that allows you to manage accounts and assigned roles.

Roles:
- `Trustee` - Is able to create accounts, assign roles, approve root certificates.
- `Vendor` - Is able to add models.
- `TestHouse` - Is able to add testing results for a model.
- `ZBCertificationCenter` - Is able to certify and revoke models.
- `NodeAdmin` - Is able to add nodes to the network of validator nodes.

##### Transactions

- Propose a new account.

  Role: `Trustee`

  Command: `dclcli tx auth propose-add-account --address=<string> --pubkey=<string> --roles=<roles> --from=<account>`

  Flags:
  - address: `string` - bench32 encoded account address
  - pubkey: `string` - bench32 encoded public key
  - roles: `optional(string)` - comma-separated list of roles (supported roles: Vendor, TestHouse, ZBCertificationCenter, Trustee, NodeAdmin)
  - from: `string` - name or address of private key with which to sign

  Example: `dclcli tx auth propose-add-account --address=cosmos15ljvz60tfekhstz8lcyy0c9l8dys5qa2nnx4d7 --pubkey=cosmospub1addwnpepqtrnrp93hswlsrzvltc3n8z7hjg9dxuh3n4rkp2w2verwfr8yg27c95l4k3 --roles=Vendor,NodeAdmin --from=jack`

- Approve a proposed account.

  Role: `Trustee`

  Command: `dclcli tx auth approve-add-account --address=<string> --from=<account>`

  Flags:
  - address: `string` - bench32 encoded account address to approve
  - from: `string` - name or address of private key with which to sign

  Example: `dclcli tx auth approve-add-account --address=cosmos15ljvz60tfekhstz8lcyy0c9l8dys5qa2nnx4d7 --from=jack`

- Propose revocation of an account.

  Role: `Trustee`

  Command: `dclcli tx auth propose-revoke-account --address=<string> --from=<account>`

  Flags:
  - address: `string` - bench32 encoded account address to approve
  - from: `string` - name or address of private key with which to sign

  Example: `dclcli tx auth propose-revoke-account --address=cosmos15ljvz60tfekhstz8lcyy0c9l8dys5qa2nnx4d7 --from=jack`

- Approve revocation of an account.

  Role: `Trustee`

  Command: `dclcli tx auth approve-revoke-account --address=<string> --from=<account>`

  Flags:
  - address: `string` - bench32 encoded account address to approve
  - from: `string` - name or address of private key with which to sign

  Example: `dclcli tx auth approve-revoke-account --address=cosmos15ljvz60tfekhstz8lcyy0c9l8dys5qa2nnx4d7 --from=jack`


##### Queries

- Get a single account. Revoked accounts are not returned.

  Command: `dclcli query auth account --address=<string>`

  Flags:
  - address: `string` - bench32 encoded account address

  Example: `dclcli query auth account --address=cosmos15ljvz60tfekhstz8lcyy0c9l8dys5qa2nnx4d7`

- Get all accounts. Revoked accounts are not returned.

  Command: `dclcli query auth all-accounts`
 
  Flags:
  - skip: `optional(int)` - number records to skip (`0` by default)
  - take: `optional(int)` - number records to take (all records are returned by default) 
 
  Example: `dclcli query auth accounts`

- Get all proposed accounts.

  Command: `dclcli query auth all-proposed-accounts`
 
  Flags:
  - skip: `optional(int)` - number records to skip (`0` by default)
  - take: `optional(int)` - number records to take (all records are returned by default) 
 
  Example: `dclcli query auth all-proposed-accounts`

- Get all proposed accounts to revoke.

  Command: `dclcli query auth all-proposed-accounts-to-revoke`
 
  Flags:
  - skip: `optional(int)` - number records to skip (`0` by default)
  - take: `optional(int)` - number records to take (all records are returned by default) 
 
  Example: `dclcli query auth all-proposed-accounts-to-revoke`

### PKI

The set of commands that allows you to manage X.509 certificates.

##### Transactions
- Propose a new self-signed root certificate.

  Role: Any  

  Command: `dclcli tx pki propose-add-x509-root-cert --certificate=<string-or-path> --from=<account>`

  Flags:
    - certificate: `string` - PEM encoded certificate (string or path to file containing data).
    - from: `string` - Name or address of private key with which to sign.

  Example: `dclcli tx pki propose-add-x509-root-cert --certificate="/path/to/certificate/file" --from=jack`
  
  Example: `dclcli tx pki propose-add-x509-root-cert --certificate="----BEGIN CERTIFICATE----- ......" --from=jack`

- Approves the proposed root certificate

  Role: `Trustee`

  Command: `dclcli tx pki approve-add-x509-root-cert --subject=<string> --subject-key-id=<hex string> --from=<account>`

  Flags:
    - subject: `string` - certificates's `Subject`.
    - subject-key-id: `string` - certificates's `Subject Key ID` (hex-encoded uppercase string).
    - from: `string` - Name or address of private key with which to sign.

  Example: `dclcli tx pki approve-add-x509-root-cert --subject="CN=dsr-corporation.com" --subject-key-id="8A:E9:AC:D4:16:81:2F:87:66:8E:61:BE:A9:C5:1C:0:1B:F7:BB:AE" --from=jack`
  
- Add an intermediate or leaf X509 certificate signed by a chain of certificates which must be already present on the ledger.

  Role: Any

  Command: `dclcli tx pki add-x509-cert --certificate=<string-or-path> --from=<account>`

  Flags:
    - certificate: `string` - PEM encoded certificate (string or path to file containing data).
    - from: `string` - Name or address of private key with which to sign.

  Example: `dclcli tx pki add-x509-cert --certificate="/path/to/certificate/file" --from=jack`
  
  Example: `dclcli tx pki add-x509-cert --certificate="----BEGIN CERTIFICATE----- ......" --from=jack`  

- Revoke the given intermediate or leaf X509 certificate. Can be done by the certificate's issuer only.

  Role: Any 

  Command: `dclcli tx pki revoke-x509-cert --subject=<string> --subject-key-id=<hex string> --from=<account>`

  Flags:
    - subject: `string` - certificates's `Subject`.
    - subject-key-id: `string` - certificates's `Subject Key ID` (hex-encoded uppercase string).
    - from: `string` - Name or address of private key with which to sign.

  Example: `dclcli tx pki revoke-x509-cert --subject="CN=dsr-corporation.com" --subject-key-id="8A:E9:AC:D4:16:81:2F:87:66:8E:61:BE:A9:C5:1C:0:1B:F7:BB:AE" --from=jack`

- Propose revocation of a X509 root certificate.  

  Role: `Trustee`
  
  Command: `dclcli tx pki propose-revoke-x509-root-cert --subject=<string> --subject-key-id=<hex string> --from=<account>`

  Flags:
    - subject: `string` - certificates's `Subject`.
    - subject-key-id: `string` - certificates's `Subject Key ID` (hex-encoded uppercase string).
    - from: `string` - Name or address of private key with which to sign.

  Example: `dclcli tx pki propose-revoke-x509-root-cert --subject="CN=dsr-corporation.com" --subject-key-id="8A:E9:AC:D4:16:81:2F:87:66:8E:61:BE:A9:C5:1C:0:1B:F7:BB:AE" --from=jack`

- Approve revocation of a X509 root certificate.  

  Role: `Trustee`
  
  Command: `dclcli tx pki approve-revoke-x509-root-cert --subject=<string> --subject-key-id=<hex string> --from=<account>`

  Flags:
    - subject: `string` - certificates's `Subject`.
    - subject-key-id: `string` - certificates's `Subject Key ID` (hex-encoded uppercase string).
    - from: `string` - Name or address of private key with which to sign.

  Example: `dclcli tx pki approve-revoke-x509-root-cert --subject="CN=dsr-corporation.com" --subject-key-id="8A:E9:AC:D4:16:81:2F:87:66:8E:61:BE:A9:C5:1C:0:1B:F7:BB:AE" --from=jack`

    
##### Queries
- Get all proposed but not approved root certificates.

  Command: `dclcli query pki all-proposed-x509-root-certs`

  Flags:
  - skip: `optional(int)` - number records to skip (`0` by default)
  - take: `optional(int)` - number records to take (all records are returned by default)
    
  Example: `dclcli query pki all-proposed-x509-root-certs`
  
- Get a proposed but not approved root certificate.

  Command: `dclcli query pki proposed-x509-root-cert --subject=<string> --subject-key-id=<hex string>`

  Flags:
    - subject: `string` - certificates's `Subject`.
    - subject-key-id: `string` - certificates's `Subject Key ID` (hex-encoded uppercase string).

  Example: `dclcli query pki proposed-x509-root-cert --subject="CN=dsr-corporation.com" --subject-key-id="8A:E9:AC:D4:16:81:2F:87:66:8E:61:BE:A9:C5:1C:0:1B:F7:BB:AE"`

- Get all approved root certificates. Revoked certificates are not returned.

  Command: `dclcli query pki all-x509-root-certs`

  Flags:
  - skip: `optional(int)` - number records to skip (`0` by default)
  - take: `optional(int)` - number records to take (all records are returned by default)
  
  Example: `dclcli query pki all-x509-root-certs`
  
- Get a certificate (either root, intermediate or leaf). Revoked certificates are not returned.

  Command: `dclcli query pki x509-cert --subject=<string> --subject-key-id=<hex string>`

  Flags:
    - subject: `string` - certificates's `Subject`.
    - subject-key-id: `string` - certificates's `Subject Key ID` (hex-encoded uppercase string).

  Example: `dclcli query pki x509-certs --subject="CN=dsr-corporation.com" --subject-key-id="8A:E9:AC:D4:16:81:2F:87:66:8E:61:BE:A9:C5:1C:0:1B:F7:BB:AE"`
  
- Get all certificates (root, intermediate and leaf). Revoked certificates are not returned.

  Command: `dclcli query pki all-x509-certs`

  Flags:
  - skip: `optional(int)` - number records to skip (`0` by default)
  - take: `optional(int)` - number records to take (all records are returned by default)
  - root-subject: `optional(string)` - filter certificates by Subject of root certificate (only the certificates started with the given root certificate are returned)
  - root-subject-key-id: `optional(string)` -   - root-subject-key-id: `optional(string)` - filter certificates by Subject of root certificate (only the certificates started with the given root certificate are returned)

  Example: `dclcli query pki x509-certs`
  
  Example: `dclcli query pki x509-certs --root-subject-key-id="8A:E9:AC:D4:16:81:2F:87:66:8E:61:BE:A9:C5:1C:0:1B:F7:BB:AE"`
  
  Example: `dclcli query pki x509-certs --root-subject="CN=dsr-corporation.com"`
  
- Get all certificates (root, intermediate and leaf) associated with subject. Revoked certificates are not returned.

  Command: `dclcli query pki all-subject-x509-certs --subject=<subject>`

  Flags:
  - subject: `string` - certificates's `Subject`
  - skip: `optional(int)` - number records to skip (`0` by default)
  - take: `optional(int)` - number records to take (all records are returned by default)
  - root-subject: `optional(string)` - filter certificates by Subject of root certificate (only the certificates started with the given root certificate are returned)
  - root-subject-key-id: `optional(string)` -   - root-subject-key-id: `optional(string)` - filter certificates by Subject of root certificate (only the certificates started with the given root certificate are returned)

  Example: `dclcli query pki all-subject-x509-certs --subject="CN=dsr"`

  Example: `dclcli query pki all-subject-x509-certs --subject="CN=dsr" --root-subject-key-id="8A:E9:AC:D4:16:81:2F:87:66:8E:61:BE:A9:C5:1C:0:1B:F7:BB:AE"`
  
  Example: `dclcli query pki all-subject-x509-certs --subject="CN=dsr" --root-subject="CN=dsr-corporation.com"`

- Get a complete chain for a certificate. Revoked certificates are not returned.

  Command: `dclcli query pki x509-cert-chain --subject=<string> --subject-key-id=<hex string>`

  Flags:
    - subject: `string` - certificates's `Subject`.
    - subject-key-id: `string` - certificates's `Subject Key ID` (hex-encoded uppercase string).

  Example: `dclcli query pki x509-cert-chain --subject="CN=dsr-corporation.com" --subject-key-id="8A:E9:AC:D4:16:81:2F:87:66:8E:61:BE:A9:C5:1C:0:1B:F7:BB:AE"`

- Get all proposed but not approved root certificates to be revoked.

  Command: `dclcli query pki all-proposed-x509-root-certs-to-revoke`

  Flags:
  - skip: `optional(int)` - number records to skip (`0` by default)
  - take: `optional(int)` - number records to take (all records are returned by default)

  Example: `dclcli query pki all-proposed-x509-root-certs-to-revoke`

- Get a proposed but not approved root certificate to be revoked.

  Command: `dclcli query pki proposed-x509-root-cert-to-revoke --subject=<string> --subject-key-id=<hex string>`

  Flags:
    - subject: `string` - certificates's `Subject`.
    - subject-key-id: `string` - certificates's `Subject Key ID` (hex-encoded uppercase string).

  Example: `dclcli query pki proposed-x509-root-cert-to-revoke --subject="CN=dsr-corporation.com" --subject-key-id="8A:E9:AC:D4:16:81:2F:87:66:8E:61:BE:A9:C5:1C:0:1B:F7:BB:AE"`

- Get all revoked certificates (both root and non-root).

  Command: `dclcli query pki all-revoked-x509-certs`

  Flags:
  - skip: `optional(int)` - number records to skip (`0` by default)
  - take: `optional(int)` - number records to take (all records are returned by default)

  Example: `dclcli query pki all-revoked-x509-certs`

- Get all revoked root certificates.

  Command: `dclcli query pki all-revoked-x509-root-certs`

  Flags:
  - skip: `optional(int)` - number records to skip (`0` by default)
  - take: `optional(int)` - number records to take (all records are returned by default)

  Example: `dclcli query pki all-revoked-x509-root-certs`

### Model Info

The set of commands that allows you to manage model infos.

##### Transactions
- Add a new model info.

  Role: `Vendor`
  
  Command: `dclcli tx modelinfo add-model --vid=<uint16> --pid=<uint16> --name=<string> --description=<string or path> --sku=<string> 
--firmware-version=<string> --hardware-version=<string> --tis-or-trp-testing-completed=<bool> --from=<account>`

  Flags:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - name: `string` -  model name
  - description: `string` -  model description (string or path to file containing data)
  - sku: `string` -  stock keeping unit
  - firmware-version: `string` -  version of model firmware
  - hardware-version: `string` -  version of model hardware
  - tis-or-trp-testing-completed: `bool` -  whether model has successfully completed TIS/TRP testing
  - from: `string` - Name or address of private key with which to sign
  - cid: `optional(uint16)` - model category ID (positive non-zero)
  - custom: `optional(string)` - custom information (string or path to file containing data)
  - ota-url: `optional(string)` - the URL of the OTA
  - ota-checksum: `optional(string)` - the checksum of the OTA 
  - ota-checksum-type: `optional(string)` - the type of the OTA checksum 

  Example: `dclcli tx modelinfo add-model --vid=1 --pid=1 --name="Device #1" --description="Device Description" --sku="SKU12FS" --firmware-version="1.0" --hardware-version="2.0" --tis-or-trp-testing-completed=true "--from=jack`
  
  Example: `dclcli tx modelinfo add-model --vid=1 --pid=1 --name="Device #1" --description="Device Description" --sku="SKU12FS" --firmware-version="1.0" --hardware-version="2.0" --tis-or-trp-testing-completed=true --from=jack --cid=1 --custom="Some Custom information" --ota-url="http://my-ota.com" --ota-checksum="df56hf" --ota-checksum-type="SHA-256"`

- Update an existing model info. Only the owner can edit a Model Info.

  Role: `Vendor`

  Command: `dclcli tx modelinfo update-model --vid=<uint16> --pid=<uint16> --tis-or-trp-testing-completed=<bool> --from=<account>`
  existing ModelInfo.
    
  Flags:
  - vid: `uint16` -  model vendor ID
  - pid: `uint16` -  model product ID
  - tis-or-trp-testing-completed: `bool` -  whether model has successfully completed TIS/TRP testing
  - from: `string` - Name or address of private key with which to sign
  - description: `optional(string)` -  model description (string or path to file containing data)
  - cid: `optional(uint16)` - model category ID
  - custom: `optional(string)` - custom information (string or path to file containing data)
  - ota-url: `optional(string)` - a new OTA URL. Can be edited only if `OTA_checksum` and `OTA_checksum_type` are already set.
    
  Example: `dclcli tx modelinfo update-model --vid=1 --pid=1 --tis-or-trp-testing-completed=true --from=jack --description="New Description"`
  
  Example: `dclcli tx modelinfo update-model --vid=1 --pid=1 --tis-or-trp-testing-completed=true --from=jack --custom="Custom Data" --ota-url="http://new-ota.com"`

##### Queries
- Query a single model info.

  Command: `dclcli query modelinfo model --vid=<uint16> --pid=<uint16>`

  Flags:
  - vid: `uint16` -  model vendor ID
  - pid: `uint16` -  model product ID

  Example: `dclcli query modelinfo model --vid=1 --pid=1`
  
- Query a list of all model infos. 

  Command: `dclcli query modelinfo all-models`

  Flags:
  - skip: `optional(int)` - number records to skip (`0` by default)
  - take: `optional(int)` - number records to take (all records are returned by default)
    
  Example: `dclcli query modelinfo all-models`

- Query a list of all vendors.

  Command: `dclcli query modelinfo vendors`
  
  Flags:
  - skip: `optional(int)` - number records to skip (`0` by default)
  - take: `optional(int)` - number records to take (all records are returned by default)
    
  Example: `dclcli query modelinfo vendors`
  
- Query a list of all model infos for the given vendor.

  Command: `dclcli query modelinfo vendor-models --vid=<uint16>`

  Flags:
  - vid: `uint16` -  model vendor ID
  - skip: `optional(int)` - number records to skip (`0` by default)
  - take: `optional(int)` - number records to take (all records are returned by default)

  Example: `dclcli query modelinfo vendor-models --vid=1`

### Compliance Test

The set of commands that allows you to manage testing results associated with a model.

##### Transactions
- Add new testing result for model associated with the given VID/PID. Note that the corresponding model must present on the ledger. 

  Role: `TestHouse`

  Command: ` dclcli tx compliancetest add-test-result --vid=<uint16> --pid=<uint16> --test-result=<string> --test-date=<rfc3339 encoded date> --from=<account>`

  Flags:
  - vid: `uint16` -  model vendor ID
  - pid: `uint16` -  model product ID
  - test-result: `string` -  test result (string or path to file containing data)
  - test-date: `string` -  date of test result (rfc3339 encoded)
  - from: `string` - Name or address of private key with which to sign

  Example: `dclcli tx compliancetest add-test-result --vid=1 --pid=1 --test-result="Test Document" --test-date="2020-04-16T06:04:57.05Z" --from=jack`
  
  Example: `dclcli tx compliancetest add-test-result --vid=1 --pid=1 --test-result="path/to/document" --test-date="2020-04-16T06:04:57.05Z" --from=jack`
  
##### Queries
- Query testing results for model associated with the given VID/PID.

  Command: `dclcli query compliancetest test-result --vid=<uint16> --pid=<uint16>`

  Flags:
  - vid: `uint16` -  model vendor ID
  - pid: `uint16` -  model product ID

  Example: `dclcli query compliancetest test-result --vid=1 --pid=1`

### Compliance

The set of commands that allows you to manage model certification information.

##### Transactions
- Certify a model associated with the given VID/PID. Note that the corresponding model and the test results must present on the ledger.
Only the owner can update an existing record. 

  Role: `ZBCertificationCenter`

  Command: `dclcli tx compliance certify-model --vid=<uint16> --pid=<uint16> --certification-type=<zb> --certification-date=<rfc3339 encoded date> --from=<account>`
  
  Flags:
  - vid: `uint16` -  model vendor ID
  - pid: `uint16` -  model product ID
  - certification-type: `string` -  certification type (zb` is the only supported value now)
  - certification-date: `string` -  the date of model certification (rfc3339 encoded)
  - from: `string` - name or address of private key with which to sign
  - reason: `optional(string)` -  an optional comment describing the reason of certification

  Example: `dclcli tx compliance certify-model --vid=1 --pid=1 --certification-type="zb" --certification-date="2020-04-16T06:04:57.05Z" --from=jack`
 
- Revoke certification for a model associated with the given VID/PID. Only the owner can update an existing record. 

  Role: `ZBCertificationCenter`

  Command: ` dclcli tx compliance revoke-model --vid=<uint16> --pid=<uint16> --certification-type=<zb> --revocation-date=<rfc3339 encoded date> --from=<account>`

  Flags:
  - vid: `uint16` -  model vendor ID
  - pid: `uint16` -  model product ID
  - certification-type: `string` -  certification type (zb` is the only supported value now)
  - revocation-date: `string` -  the date of model revocation (rfc3339 encoded)
  - from: `string` - name or address of private key with which to sign
  - reason: `optional(string)` -  an optional comment describing the reason of revocation

  Example: `dclcli tx compliance revoke-model --vid=1 --pid=1 --certification-type="zb" --revocation-date="2020-04-16T06:04:57.05Z" --from=jack`
  
  Example: `dclcli tx compliance revoke-model --vid=1 --pid=1 --certification-type="zb" --revocation-date="2020-04-16T06:04:57.05Z" --reason "Some Reason" --from=jack`
  
##### Queries
- Check if the model associated with the given VID/PID is certified.

  Command: `dclcli query compliance certified-model --vid=<uint16> --pid=<uint16> --certification-type=<zb>`

  Flags:
  - vid: `uint16` -  model vendor ID
  - pid: `uint16` -  model product ID
  - certification-type: `string` -  certification type (zb` is the only supported value now)

  Example: `dclcli query compliance certified-model --vid=1 --pid=1 --certification-type="zb"`
  
- Query all certified models.

  Command: `dclcli query compliance all-certified-models`

  Flags:
  - skip: `optional(int)` - number records to skip (`0` by default)
  - take: `optional(int)` - number records to take (all records are returned by default)

  Example: `dclcli query compliance all-certified-models`
  
- Check if the model associated with the given VID/PID is revoked.

  Command: `dclcli query compliance revoked-model --vid=<uint16> --pid=<uint16> --certification-type=<zb>`

  Flags:
  - vid: `uint16` -  model vendor ID
  - pid: `uint16` -  model product ID
  - certification-type: `string` -  certification type (zb` is the only supported value now)

  Example: `dclcli query compliance revoked-model --vid=1 --pid=1 --certification-type="zb"`
  
- Query all revoked models.

  Command: `dclcli query compliance all-revoked-models`

  Flags:
  - skip: `optional(int)` - number records to skip (`0` by default)
  - take: `optional(int)` - number records to take (all records are returned by default)

  Example: `dclcli query compliance all-revoked-models`
  
- Query compliance info for model associated with VID/PID.

  Command: `dclcli query compliance compliance-info --vid=<uint16> --pid=<uint16> --certification-type=<zb>`

  Flags:
  - vid: `uint16` -  model vendor ID
  - pid: `uint16` -  model product ID
  - certification-type: `string` -  certification type (zb` is the only supported value now)

  Example: `dclcli query compliance compliance-info --vid=1 --pid=1 --certification-type="zb"`
  
- Query all compliance infos. 

  Command: `dclcli query compliance all-compliance-info-records`

  Flags:
  - skip: `optional(int)` - number records to skip (`0` by default)
  - take: `optional(int)` - number records to take (all records are returned by default)

  Example: `dclcli query compliance all-compliance-info-records`

### Validator

The set of commands that allows you to manage the set of validator nodes in the network.

##### Transactions
- Add a new validator node.

  Role: `NodeAdmin`

  Command: `dclcli tx validator add-node --validator-address=<address> --validator-pubkey=<pubkey> --name=<node name> --from=<account>`
  
  Flags:
  - validator-address: `string` - the tendermint validator address
  - validator-pubkey: `string` - the tendermint validator public key
  - name: `string` -  validator name
  - from: `string` - name or address of private key with which to sign
  - website: `optional(string)` - optional validator's website
  - identity: `optional(string)` - optional identity signature
  - details: `optional(string)` - optional validator's details
  
  Example: `dclcli tx validator add-node --validator-address=cosmosvalcons1tl46nm39xtuutvw2wqaeyyd6csknfe0a7xqnrw --validator-pubkey=cosmosvalconspub1zcjduepqn5nz4c8n5jwfmgd6tqfqzu8arpne3au4g7tfsz33g8y6dcvhkf4sw054j8 --name=node1 --from=jack`
 
##### Queries
- Query a validator node by the given validator address.

  Command: `dclcli query validator node --validator-address=<address>`

  Flags:
  - validator-address: `string` - the tendermint validator address

  Example: `dclcli query validator node --validator-address=cosmosvaloper1jnr3hcrvcpvqm5fdcafg70azkg0awf96tvu79z`
  
- Query all validator nodes.

  Command: `dclcli query validator all-nodes`

  Flags:
  - state: `string` - state of a validator (active/jailed)

  Example: `dclcli query validator all-nodes`
