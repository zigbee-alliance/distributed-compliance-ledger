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
- `Vendor` - Is able to add models that belong to the vendor ID associated with the vendor account.
- `TestHouse` - Is able to add testing results for a model.
- `CertificationCenter` - Is able to certify and revoke models.
- `NodeAdmin` - Is able to add validator nodes to the network.

##### Transactions

- Propose a new account.

  Role: `Trustee`

  Command: `dclcli tx auth propose-add-account --address=<string> --pubkey=<string> --roles=<roles> --from=<account>`

  Flags:
  - address: `string` - bench32 encoded account address
  - pubkey: `string` - bench32 encoded public key
  - roles: `optional(string)` - comma-separated list of roles (supported roles: Vendor, TestHouse, CertificationCenter, Trustee, NodeAdmin)
  - from: `string` - name or address of private key with which to sign

  Example: `dclcli tx auth propose-add-account --address=cosmos15ljvz60tfekhstz8lcyy0c9l8dys5qa2nnx4d7 --pubkey=cosmospub1addwnpepqtrnrp93hswlsrzvltc3n8z7hjg9dxuh3n4rkp2w2verwfr8yg27c95l4k3 --roles=Trustee,NodeAdmin --from=jack`

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

- Propose a new account (Vendor Role). A vendor role is tied to a Vendor ID, hence while proposing a Vendor Role vid is a required field.

  Role: `Trustee`

  Command: `dclcli tx auth propose-add-account --address=<string> --pubkey=<string> --roles=<roles> --vid=<vendorID> --from=<account>`

  Flags:
  - address: `string` - bench32 encoded account address
  - pubkey: `string` - bench32 encoded public key
  - roles: `optional(string)` - comma-separated list of roles (supported roles: Vendor, TestHouse, CertificationCenter, Trustee, NodeAdmin)
  - vid: `string` - Vendor ID associated with this account. Required only for Vendor Roles
  - from: `string` - name or address of private key with which to sign

  Example: `dclcli tx auth propose-add-account --address=cosmos15ljvz60tfekhstz8lcyy0c9l8dys5qa2nnx4d7 --pubkey=cosmospub1addwnpepqtrnrp93hswlsrzvltc3n8z7hjg9dxuh3n4rkp2w2verwfr8yg27c95l4k3 --roles=Vendor,NodeAdmin --vid=123 --from=jack`
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


### Vendor Info

The set of commands that allows the Vendor to update contact information.
##### Transactions
- Add a new vendor contact info.

  Role: `Vendor`
  
  Command: `dclcli tx vendorinfo add-vendor --companyLegalName=<string> --companyPreferredName=<string> --vendorName=<string> --vid=<uint16> ----vendorLandingPageURL=<url> --from=<account>`

  Flags:
  - companyLegalName `string (max64)` - Company Legal Name
  - companyPreferredName `optional string (max64)` - Company Preferred Name  
  - vendorLandingPageURL `optional string (max256)` - Landing Page URL for the Vendor
  - vendorName `string (max32)` - Vendor Name
  - vid `uint16` - Vendor ID      
  

  Example: `dclcli tx vendorinfo add-vendor --companyLegalName="XYZ Technology Solutions" --vid=123 --vendorName="XYZ Inc" --from="jack"`

- Update vendor contact info.

  Role: `Vendor`

  Command: `dclcli tx vendorinfo update-vendor --companyLegalName=<string> --companyPreferredName=<string> --vendorName=<string> --vid=<uint16> ----vendorLandingPageURL=<url> --from=<account>`

    
  Flags:
  - companyLegalName `string (max64)` - Company Legal Name
  - companyPreferredName `optional string (max64)` - Company Preferred Name  
  - vendorLandingPageURL `optional string (max256)` - Landing Page URL for the Vendor
  - vendorName `string (max32)` - Vendor Name
  - vid `uint16` - Vendor ID     
    
  Example: `dclcli tx vendorinfo update-vendor --vendorLandingPageURL="https://producturl.vendor.info" --vid=123 --from="jack"`
  
##### Queries
- Query a single vendor info.

  Command: `dclcli query vendorinfo vendor --vid=<uint16>`

  Flags:
  - vid: `uint16` -  vendor ID

  Example: `dclcli query vendorinfo vendor --vid=1`
  
- Query a list of all vendor infos. 

  Command: `dclcli query vendorinfo all-vendors`

  Flags:
  - skip: `optional(int)` - number records to skip (`0` by default)
  - take: `optional(int)` - number records to take (all records are returned by default)
    
  Example: `dclcli query vendorinfo all-vendors`

### Device Model 

The set of commands that allows you to manage model and model versions.

##### Transactions
- Add a new model. Only vendor role with associated vendor ID can add the model for given Vendor ID

  Role: `Vendor`
  
  Command: `dclcli tx model add-model --vid=<uint16> --pid=<uint16> --productName=<string> --productLabel=<string or path> --sku=<string> 
--softwareVersion=<uint32> --softwareVersionString=<string> --hardwareVersion=<uint32> --hardwareVersionString=<string> --cdVersionNumber=<uint16> 
--from=<account>`

  Flags:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - deviceTypeID: `uint16` -  DeviceTypeID is the device type identifier. For example, DeviceTypeID 10 (0x000a), is the device type identifier for a Door Lock.
  - productName: `string` -  model name
  - productLabel: `string` -  model description (string or path to file containing data)
  - partNumber: `string` -  stock keeping unit
   
  - commissioningCustomFlow: `optional(uint8)` - A value of 1 indicates that user interaction with the device (pressing a button, for example) is required before commissioning can take place. When CommissioningCustomflow is set to a value of 2, the commissioner SHOULD attempt to obtain a URL which MAY be used to provide an end-user with the necessary details for how to configure the product for initial commissioning
  - commissioningCustomFlowURL: `optional(string)` - commissioningCustomFlowURL SHALL identify a vendor specific commissioning URL for the device model when the commissioningCustomFlow field is set to '2'
  - commissioningModeInitialStepsHint: `optional(uint32)` - commissioningModeInitialStepsHint SHALL identify a hint for the steps that can be used to put into commissioning mode a device that has not yet been commissioned. This field is a bitmap with values defined in the Pairing Hint Table. For example, a value of 1 (bit 0 is set) indicates that a device that has not yet been commissioned will enter Commissioning Mode upon a power cycle.
  - commissioningModeInitialStepsInstruction: `optional(string)` - commissioningModeInitialStepsInstruction SHALL contain text which relates to specific values of CommissioningModeInitialStepsHint. Certain values of CommissioningModeInitialStepsHint, as defined in the Pairing Hint Table, indicate a Pairing Instruction (PI) dependency, and for these values the commissioningModeInitialStepsInstruction SHALL be set
  - commissioningModeSecondaryStepsHint: `optional(uint32)` - commissioningModeSecondaryStepsHint SHALL identify a hint for steps that can be used to put into commissioning mode a device that has already been commissioned. This field is a bitmap with values defined in the Pairing Hint Table. For example, a value of 4 (bit 2 is set) indicates that a device that has already been commissioned will require the user to visit a current CHIP Administrator to put the device into commissioning mode.
  - commissioningModeSecondaryStepInstruction: `optional(string)` - commissioningModeSecondaryStepInstruction SHALL contain text which relates to specific values of commissioningModeSecondaryStepsHint. Certain values of commissioningModeSecondaryStepsHint, as defined in the Pairing Hint Table, indicate a Pairing Instruction (PI) dependency, and for these values the commissioningModeSecondaryStepInstruction SHALL be set
  - userManualURL: `optional(string)` - URL that contains product specific web page that contains user manual for the device model.
  - supportURL: `optional(string)` - URL that contains product specific web page that contains support details for the device model.
  - productURL: `optional(string)` - URL that contains product specific web page that contains details for the device model.  
  
  

  Example: `dclcli tx model add-model --vid=1 --pid=1 --productName="Device #1" --productLabel="Device Description" --partNumber="SKU12FS"  --from="jack"`
  
- Update an existing model. Only the vendor role with associated vendorID can edit a Model.

  Role: `Vendor`

  Command: `dclcli tx model update-model --vid=<uint16> --pid=<uint16>  --from=<account>`
  
    
  Flags:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - productName: `string` -  model name
  - productLabel: `string` -  model description (string or path to file containing data)
  - partNumber: `string` -  stock keeping unit
  - commissioningCustomFlowURL: `optional(string)` - commissioningCustomFlowURL SHALL identify a vendor specific commissioning URL for the device model when the commissioningCustomFlow field is set to '2'
  - commissioningModeInitialStepsInstruction: `optional(string)` - commissioningModeInitialStepsInstruction SHALL contain text which relates to specific values of CommissioningModeInitialStepsHint. Certain values of CommissioningModeInitialStepsHint, as defined in the Pairing Hint Table, indicate a Pairing Instruction (PI) dependency, and for these values the commissioningModeInitialStepsInstruction SHALL be set
  - commissioningModeSecondaryStepInstruction: `optional(string)` - commissioningModeSecondaryStepInstruction SHALL contain text which relates to specific values of commissioningModeSecondaryStepsHint. Certain values of commissioningModeSecondaryStepsHint, as defined in the Pairing Hint Table, indicate a Pairing Instruction (PI) dependency, and for these values the commissioningModeSecondaryStepInstruction SHALL be set
  - userManualURL: `optional(string)` - URL that contains product specific web page that contains user manual for the device model.
  - supportURL: `optional(string)` - URL that contains product specific web page that contains support details for the device model.
  - productURL: `optional(string)` - URL that contains product specific web page that contains details for the device model.  

  Example: `dclcli tx model update-model --vid=1 --pid=1 --productLabel="New Description" --from=jack `
  
  Example: `dclcli tx model update-model --vid=1 --pid=1 --supportURL="https://product-support.url.test"  --from=jack `


- Add a new model version. Only vendor role with associated vendor ID can add the model versions for given Vendor ID. Also the Model with a given pid should be present on the ledger before model versions can be added.

  Role: `Vendor`
  
  Command: `dclcli tx model add-model-version --cdVersionNumber=1 --maxApplicableSoftwareVersion=10 --minApplicableSoftwareVersion=1 --vid=$vid --pid=$pid1 --softwareVersion=$software_version --softwareVersionString=1 --from=jack `

  Flags:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - softwareVersion: `uint32` - model software version
  - softwareVersionSting: `string` - model software version string 
  - cdVersionNumber `uint32` - CD Version Number of the certification
  - firmwareDigests `string` - FirmwareDigests field included in the Device Attestation response when this Software Image boots on the device
  - softwareVersionValid `bool` - Flag to indicate whether the software version is valid or not (default true)
  - otaURL `string` - URL where to obtain the OTA image
  - otaFileSize `string`  - OtaFileSize is the total size of the OTA software image in bytes
  - otaChecksum `string` - Digest of the entire contents of the associated OTA Software Update Image under the OtaUrl attribute, encoded in base64 string representation. The digest SHALL have been computed using the algorithm specified in OtaChecksumType
  - otaChecksumType `string` - Numeric identifier as defined in IANA Named Information Hash Algorithm Registry for the type of otaChecksum. For example, a value of 1 would match the sha-256 identifier, which maps to the SHA-256 digest algorithm
  - maxApplicableSoftwareVersion `uint32` - MaxApplicableSoftwareVersion should specify the highest SoftwareVersion for which this image can be applied
  - minApplicableSoftwareVersion `uint32` - MinApplicableSoftwareVersion should specify the lowest SoftwareVersion for which this image can be applied
  - releaseNotesURL `string` - URL that contains product specific web page that contains release notes for the device model.

  Example: `dclcli tx model add-model-version --vid=1 --pid=1 --softwareVersion=20 --softwareVersionString="1.0" --cdVersionNumber=1 --minApplicableSoftwareVersion=1 --maxApplicableSoftwareVersion=10  --from="jack"`
  
- Update an existing model version. Only the vendor role with associated vendorID can edit a Model.

  Role: `Vendor`

  Command: `dclcli tx model update-model-version --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --from=<account>`
 
    
  Flags:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - softwareVersion: `uint32` - model software version
  - softwareVersionValid `bool` - Flag to indicate whether the software version is valid or not (default true)
  - otaURL `string` - URL where to obtain the OTA image
  
  - maxApplicableSoftwareVersion `uint32` - MaxApplicableSoftwareVersion should specify the highest SoftwareVersion for which this image can be applied
  - minApplicableSoftwareVersion `uint32` - MinApplicableSoftwareVersion should specify the lowest SoftwareVersion for which this image can be applied
  - releaseNotesURL `string` - URL that contains product specific web page that contains release notes for the device model.

 
    
  Example: `dclcli tx model update-model-version --vid=1 --pid=1 --softwareVersion=1 --releaseNotesURL="https://release.notes.url.info" --from=jack `
  
##### Queries
- Query a single model.

  Command: `dclcli query model get-model --vid=<uint16> --pid=<uint16>`

  Flags:
  - vid: `uint16` -  model vendor ID
  - pid: `uint16` -  model product ID

  Example: `dclcli query model get-model --vid=1 --pid=1`

- Query a single model version.

  Command: `dclcli query model get-model-version --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32>`

  Flags:
  - vid: `uint16` -  model vendor ID
  - pid: `uint16` -  model product ID
  - softwareVersion: `uint32` - model software version
  

  Example: `dclcli query model get-model-version --vid=1 --pid=1 --softwareVersion=1`  
  
- Query a list of all models. 

  Command: `dclcli query model all-models`

  Flags:
  - skip: `optional(int)` - number records to skip (`0` by default)
  - take: `optional(int)` - number records to take (all records are returned by default)
    
  Example: `dclcli query model all-models`

- Query a list of all model versions for a given model. 

  Command: `dclcli query model all-model-versions --vid=<uint16> --pid=<uint16>`

  Flags:
  - vid: `uint16` -  model vendor ID
  - pid: `uint16` -  model product ID
  - skip: `optional(int)` - number records to skip (`0` by default)
  - take: `optional(int)` - number records to take (all records are returned by default)
    
  Example: `dclcli query model all-model-versions --vid=1 --pid=1`  

- Query a list of all vendors.

  Command: `dclcli query model vendors`
  
  Flags:
  - skip: `optional(int)` - number records to skip (`0` by default)
  - take: `optional(int)` - number records to take (all records are returned by default)
    
  Example: `dclcli query model vendors`
  
- Query a list of all models for the given vendor.

  Command: `dclcli query model vendor-models --vid=<uint16>`

  Flags:
  - vid: `uint16` -  model vendor ID
  - skip: `optional(int)` - number records to skip (`0` by default)
  - take: `optional(int)` - number records to take (all records are returned by default)

  Example: `dclcli query model vendor-models --vid=1`

### Compliance Test

The set of commands that allows you to manage testing results associated with a model.

##### Transactions
- Add new testing result for model associated with the given VID/PID. Note that the corresponding model must present on the ledger. 

  Role: `TestHouse`

  Command: ` dclcli tx compliancetest add-test-result --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --softwareVersionString=<string> --test-result=<string> --test-date=<rfc3339 encoded date> --from=<account>`

  Flags:
  - vid: `uint16` -  model vendor ID
  - pid: `uint16` -  model product ID
  - softwareVersion: `uint32` - model software version
  - softwareVersionSting: `string` - model software version string 
  - test-result: `string` -  test result (string or path to file containing data)
  - test-date: `string` -  date of test result (rfc3339 encoded)
  - from: `string` - Name or address of private key with which to sign

  Example: `dclcli tx compliancetest add-test-result --vid=1 --pid=1 --softwareVersion=1 --softwareVersionString="1.0"" --test-result="Test Document" --test-date="2020-04-16T06:04:57.05Z" --from=jack`
  
  Example: `dclcli tx compliancetest add-test-result --vid=1 --pid=1 --softwareVersion=1 --softwareVersionString="1.0"" --test-result="path/to/document" --test-date="2020-04-16T06:04:57.05Z" --from=jack`
  
##### Queries
- Query testing results for model associated with the given VID/PID/SoftwareVersion.

  Command: `dclcli query compliancetest test-result --vid=<uint16> --pid=<uint16> --softwareVersion=<uint16>"`

  Flags:
  - vid: `uint16` -  model vendor ID
  - pid: `uint16` -  model product ID
  - softwareVersion: `uint32` - model software version

  Example: `dclcli query compliancetest test-result --vid=1 --pid=1 --softwareVersion=1`

### Compliance

The set of commands that allows you to manage model certification information.

##### Transactions
- Certify a model associated with the given VID/PID/SoftwareVersion. Note that the corresponding model and the test results must present on the ledger.
Only the owner can update an existing record. 

  Role: `CertificationCenter`

  Command: `dclcli tx compliance certify-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint16> --softwareVersionString=<string> --certificationType=<zigbee|matter> --certificationDate=<rfc3339 encoded date> --from=<account>`
  
  Flags:
  - vid: `uint16` -  model vendor ID
  - pid: `uint16` -  model product ID
  - softwareVersion: `uint32` - model software version
  - softwareVersionSting: `string` - model software version string 
  - certificationType: `string` -  certification type (`zigbee` & `matter` are the only supported values for now)
  - certificationDate: `string` -  the date of model certification (rfc3339 encoded)
  - from: `string` - name or address of private key with which to sign
  - reason: `optional(string)` -  an optional comment describing the reason of certification

  Example: `dclcli tx compliance certify-model --vid=1 --pid=1 --softwareVersion=1 --certificationType="matter" --certificationDate="2020-04-16T06:04:57.05Z" --from=jack`
 
- Revoke certification for a model associated with the given VID/PID/SoftwareVersion. Only the owner can update an existing record. 

  Role: `CertificationCenter`

  Command: ` dclcli tx compliance revoke-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint16> --certificationType=<zigbee|matter> --revocationDate=<rfc3339 encoded date> --from=<account>`

  Flags:
  - vid: `uint16` -  model vendor ID
  - pid: `uint16` -  model product ID
  - softwareVersion: `uint32` - model software version
  - certificationType: `string` -  certification type (zigbee & matter` are the only supported values for now)
  - revocationDate: `string` -  the date of model revocation (rfc3339 encoded)
  - from: `string` - name or address of private key with which to sign
  - reason: `optional(string)` -  an optional comment describing the reason of revocation

  Example: `dclcli tx compliance revoke-model --vid=1 --pid=1 --softwareVersion=1 --certificationType="matter" --revocationDate="2020-04-16T06:04:57.05Z" --from=jack`
  
  Example: `dclcli tx compliance revoke-model --vid=1 --pid=1 --softwareVersion=1 --certificationType="matter" --revocationDate="2020-04-16T06:04:57.05Z" --reason "Some Reason" --from=jack`
  
##### Queries
- Check if the model associated with the given VID/PID/SoftwareVersion is certified.

  Command: `dclcli query compliance certified-model --vid=<uint16> --pid=<uint16> --certificationType=<zb>`

  Flags:
  - vid: `uint16` -  model vendor ID
  - pid: `uint16` -  model product ID
  - softwareVersion: `uint32` - model software version
  - certificationType: `string` -  certification type (zigbee & matter` are the only supported values for now)

  Example: `dclcli query compliance certified-model --vid=1 --pid=1 --softwareVersion=1 --certificationType="matter"`
  
- Query all certified models.

  Command: `dclcli query compliance all-certified-models`

  Flags:
  - skip: `optional(int)` - number records to skip (`0` by default)
  - take: `optional(int)` - number records to take (all records are returned by default)

  Example: `dclcli query compliance all-certified-models`
  
- Check if the model associated with the given VID/PID/SoftwareVersion is revoked.

  Command: `dclcli query compliance revoked-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint16> --certificationType=<zigbee|matter>`

  Flags:
  - vid: `uint16` -  model vendor ID
  - pid: `uint16` -  model product ID
  - softwareVersion: `uint32` - model software version
  - certificationType: `string` -  certification type (zigbee & matter` are the only supported values for now)

  Example: `dclcli query compliance revoked-model --vid=1 --pid=1 --softwareVersion=1 --certificationType="matter"`
  
- Query all revoked models.

  Command: `dclcli query compliance all-revoked-models`

  Flags:
  - skip: `optional(int)` - number records to skip (`0` by default)
  - take: `optional(int)` - number records to take (all records are returned by default)

  Example: `dclcli query compliance all-revoked-models`
  
- Query compliance info for model associated with VID/PID/SoftwareVersion.

  Command: `dclcli query compliance compliance-info --vid=<uint16> --pid=<uint16> --softwareVersion=<uint16> --certificationType=<zigbee|matter>`

  Flags:
  - vid: `uint16` -  model vendor ID
  - pid: `uint16` -  model product ID
  - certificationType: `string` -  certification type (zigbee & matter` are the only supported values for now)

  Example: `dclcli query compliance compliance-info --vid=1 --pid=1 --softwareVersion=1 --certificationType="matter"`
  
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
