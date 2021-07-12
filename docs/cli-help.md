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

  You will be prompted to create an encryption passphrase. 
  This passphrase will be requested each time you send write transactions on the ledger using this key.
  You can remember and securely save the mnemonic phrase shown after the key is created
  to be able to recover the key later.

  Command: `dclcli keys add <key name>`

  Example: `dclcli keys add jack`

- Recover existing key instead of creating a new one.

  The key can be recovered from a seed obtained from the mnemonic passphrase (see the previous command).
  You will be prompted to create an encryption passphrase and enter the seed's mnemonic.
  This passphrase will be requested each time you send write transactions on the ledger using this key.

  Command: `dclcli keys add <key name> --recover`

  Example: `dclcli keys add jack --recover`

- Get a list of all stored public keys. 

  Command: `dclcli keys list`

  Example: `dclcli keys list`

- Get details for a key.

  Command: `dclcli keys show <key name>`

  Example: `dclcli keys show jack`

- Export a key.

  A private key from the local keystore can be exported in ASCII-armored encrypted format.
  You will be prompted to enter the decryption passphrase for the key and  
  to create an encryption passphrase for the exported key.
  The exported key can be stored to a file for import.
 
  Command: `dclcli keys export <key name>`
  
  Example: `dclcli keys export jack`
  
- Import a key.

  A key can be imported from the ASCII-armored encrypted format 
  obtained by the export key command.
  You will be prompted to enter the decryption passphrase for the exported key
  which was used during the export process.

  Command: `dclcli keys import <key name> <key file>`
  
  Example: `dclcli keys import jack jack_exported_priv_key_file`

### Authorization

The set of commands that allows you to manage accounts and assigned roles.

Roles:
- `Trustee` - Is able to create accounts, assign roles, approve root certificates.
- `Vendor` - Is able to add models.
- `TestHouse` - Is able to add testing results for a model.
- `ZBCertificationCenter` - Is able to certify and revoke models.
- `NodeAdmin` - Is able to add validator nodes to the network.

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
  
  Command: `dclcli tx modelinfo add-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --hardwareVersion=<uint32> --productName=<string> --description=<string or path> --sku=<string> --softwareVersionString=<string>  --hardwareVersionString=<string> --cdVersionNumber=<uint16> --from=<account>`

  Flags:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - softwareVersion: `uint32` -  Software Version of model (uint32)
  - hardwareVersion: `uint32` -  version of model hardware
  - name: `string` -  model name
  - description: `string` -  model description (string or path to file containing data)
  - sku: `string` -  stock keeping unit
  - softwareVersionString: `string` - Software Version String of model
  - hardwareVersionString: `string` - Hardware Version String of model
  - cdVersionNumber: `uint32` -  CD Version Number of the Certification
  - from: `string` - Name or address of private key with which to sign
  - cid: `optional(uint16)` - model category ID (positive non-zero)
  - revoked: `optional(bool)` - boolean flag to revoke the model
  - otaURL: `optional(string)` - the URL of the OTA
  - otaChecksum: `optional(string)` - the checksum of the OTA 
  - otaChecksumType: `optional(string)` - the type of the OTA checksum 
  - otaBlob: `optional(string)` - metadata about OTA 
  - commissioningCustomFlow: `optional(uint8)` - A value of 1 indicates that user interaction with the device (pressing a button, for example) is required before commissioning can take place. When CommissioningCustomflow is set to a value of 2, the commissioner SHOULD attempt to obtain a URL which MAY be used to provide an end-user with the necessary details for how to configure the product for initial commissioning
  - commissioningCustomFlowURL: `optional(string)` - commissioningCustomFlowURL SHALL identify a vendor specific commissioning URL for the device model when the commissioningCustomFlow field is set to '2'
  - commissioningModeInitialStepsHint: `optional(uint32)` - commissioningModeInitialStepsHint SHALL identify a hint for the steps that can be used to put into commissioning mode a device that has not yet been commissioned. This field is a bitmap with values defined in the Pairing Hint Table. For example, a value of 1 (bit 0 is set) indicates that a device that has not yet been commissioned will enter Commissioning Mode upon a power cycle.
  - commissioningModeInitialStepsInstruction: `optional(string)` - commissioningModeInitialStepsInstruction SHALL contain text which relates to specific values of CommissioningModeInitialStepsHint. Certain values of CommissioningModeInitialStepsHint, as defined in the Pairing Hint Table, indicate a Pairing Instruction (PI) dependency, and for these values the commissioningModeInitialStepsInstruction SHALL be set
  - commissioningModeSecondaryStepsHint: `optional(uint32)` - commissioningModeSecondaryStepsHint SHALL identify a hint for steps that can be used to put into commissioning mode a device that has already been commissioned. This field is a bitmap with values defined in the Pairing Hint Table. For example, a value of 4 (bit 2 is set) indicates that a device that has already been commissioned will require the user to visit a current CHIP Administrator to put the device into commissioning mode.
  - commissioningModeSecondaryStepInstruction: `optional(string)` - commissioningModeSecondaryStepInstruction SHALL contain text which relates to specific values of commissioningModeSecondaryStepsHint. Certain values of commissioningModeSecondaryStepsHint, as defined in the Pairing Hint Table, indicate a Pairing Instruction (PI) dependency, and for these values the commissioningModeSecondaryStepInstruction SHALL be set
  - releaseNotesURL: `optional(string)` - URL that contains product specific web page that contains release notes for the device model.
  - userManualURL: `optional(string)` - URL that contains product specific web page that contains user manual for the device model.
  - supportURL: `optional(string)` - URL that contains product specific web page that contains support details for the device model.
  - productURL: `optional(string)` - URL that contains product specific web page that contains details for the device model.  
  - chipBlob: `optional(string)` - chipBlob SHALL identify CHIP specific configurations
  - vendorBlob: `optional(string)` - field for vendors to provide any additional metadata about the device model using a string, blob, or URL.  
  

  Example: `dclcli tx modelinfo add-model --vid=1 --pid=1 --softwareVersion=1 --hardwareVersion=1 --productName="Device #1" --description="Device Description" --sku="SKU12FS" --softwareVersionString="1.0b123"  --hardwareVersionString="5.1.23"  --cdVersionNumber="32" --from="jack"`
  
  Example: `dclcli tx modelinfo add-model --vid=1 --pid=1 --softwareVersion=1 --hardwareVersion=1 --productName="Device #1" --description="Device Description" --sku="SKU12FS" --softwareVersionString="1.0b123"  --hardwareVersion="5123" --cdVersionNumber="32"  --cid=1 --custom="Some Custom information" --otaURL="http://my-ota.com" --otaChecksum="df56hf" --otaChecksumType="SHA-256" --from=jack `

- Update an existing model info. Only the owner can edit a Model Info.

  Role: `Vendor`

  Command: `dclcli tx modelinfo update-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --hardwareVersion=<uint32> --from=<account>`
  existing ModelInfo.
    
  Flags:
  - vid: `uint16` -  model vendor ID
  - pid: `uint16` -  model product ID
  - softwareVersion: `uint32` -  Software Version of model (uint32)
  - hardwareVersion: `uint32` -  version of model hardware
  - description: `string` -  model description (string or path to file containing data)
  - cdVersionNumber: `uint32` -  CD Version Number of the Certification
  - from: `string` - Name or address of private key with which to sign
  - revoked: `optional(bool)` - boolean flag to revoke the model
  - otaURL: `optional(string)` - the URL of the OTA
  - otaChecksum: `optional(string)` - the checksum of the OTA 
  - otaChecksumType: `optional(string)` - the type of the OTA checksum 
  - otaBlob: `optional(string)` - metadata about OTA 
  - commissioningCustomFlowURL: `optional(string)` - commissioningCustomFlowURL SHALL identify a vendor specific commissioning URL for the device model when the commissioningCustomFlow field is set to '2'
  - releaseNotesURL: `optional(string)` - URL that contains product specific web page that contains release notes for the device model.
  - userManualURL: `optional(string)` - URL that contains product specific web page that contains user manual for the device model.
  - supportURL: `optional(string)` - URL that contains product specific web page that contains support details for the device model.
  - productURL: `optional(string)` - URL that contains product specific web page that contains details for the device model.  
  - chipBlob: `optional(string)` - chipBlob SHALL identify CHIP specific configurations
  - vendorBlob: `optional(string)` - field for vendors to provide any additional metadata about the device model using a string, blob, or URL.  
    
  Example: `dclcli tx modelinfo update-model --vid=1 --pid=1 --softwareVersion=1 --hardwareVersion=1 --cdVersionNumber="32" --description="New Description" --from=jack `
  
  Example: `dclcli tx modelinfo update-model --vid=1 --pid=1 --softwareVersion=1 --hardwareVersion=1 --supportURL="https://product-support.url.test" --otaURL="http://new-ota.com" --from=jack `

##### Queries
- Query a single model info.

  Command: `dclcli query modelinfo model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --hardwareVersion=<uint32>`

  Flags:
  - vid: `uint16` -  model vendor ID
  - pid: `uint16` -  model product ID
  - softwareVersion: `uint32` -  Software Version of model (uint32)
  - hardwareVersion: `uint32` -  version of model hardware


  Example: `dclcli query modelinfo model --vid=1 --pid=1  --softwareVersion=1 --hardwareVersion=1`
  
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

- Query the list of all Model Versions by combination of Vendor ID and Product ID.

  Command: `dclcli query modelinfo model-versions --vid=<uint16> --pid=<uint16>`

  Flags:
  - vid: `uint16` -  model vendor ID
  - pid: `uint16` -  model product ID
  
  Example: `dclcli query modelinfo model-versions --vid=1 --pid=1`
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
