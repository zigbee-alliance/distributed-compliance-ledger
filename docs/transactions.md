# Transactions and Queries
See use case sequence diagrams for the examples of how transaction can be used.

## General
- Every writer to the Ledger must  
    - Have a private/public key pair.
    - Have an Account created on the ledger via `ACCOUNT` transaction (see [Use Case Txn Auth](use_cases_txn_auth.puml)).
        - The Account stores the public part of the key
        - The Account has an associated role. The role is used for authorization policies (see [Auth Map](auth_map.md)).
    - Sign every transaction by the private key.
- Ledger is public for read which means that anyone can read from the Ledger without a need to have 
an Account or sign the request.  
- The following roles are supported:
    - Trustee
    - Vendor
    - TestHouse
    - ZBCertificationCenter
    - NodeAdmin   
- All read (get) requests return the current `height` of the ledger in addition to the
requested data. The `height` can be used to get a delta (changes) from the last state that the user has.
This is useful to avoid correlation by the sender's IP address.        

## How to write to the Ledger
- Local CLI
    - CLI is started in a CLI mode.
    - Generate and store a private key for the Account to be used for sending.
    - Send transactions to the ledger from the Account (`--from`).
        - it will automatically build a request, sign it by the account's key, and broadcast to the ledger.
    - See `CLI` section for every write request (transaction).
- Non-trusted REST API (keys at the edge):
    - CLI is started in a server mode.
    - A private key is generated and stored off-server (in the user's private wallet).
    - The user builds and signs the request manually.
    - The user does a `POST` of the signed request to the CLI-based server for broadcasting using `tx/broadcast`. 
- Trusted REST API (keys at the server):
    - CLI is started in a server mode.
    - A private key is generated and stored on the server (assuming it's within the user's domain and the user trusts it).
    - The user sends a POST to the server specifying the transaction parameters and account details (passphrase etc.).
    The server builds the request, signs it by the specified account's private key and broadcasts to the ledger
     in a way similar to the local CLI case.
    - See `REST API` section for every write request (transaction).

## How to read from the Ledger
- Local CLI
    - CLI is started in a CLI mode.
    - No keys/account is needed as the ledger is public for reads
    - See `CLI` section for every read request.
- REST API
    - CLI is started in a server mode.
    - No keys/account is needed as the ledger is public for reads
    - See `REST API` section for every read request.    

        

## KV Store
A summary of KV store and paths used:
- KV store name: `pki`
    - Proposed but not approved root certificates:
        - `1:<Certificate's Issuer>:<Certificate's Serial Number>` : `<Certificate in PEM format> + <List of approved trustee account IDs>`
    - Approved root certificates:    
        - `2:<Certificate's Issuer>:<Certificate's Serial Number>` : `<Certificate in PEM format>`
    - Non-root certificates:
        - `3:<Certificate's Issuer>:<Certificate's Serial Number>` : `<Certificate in PEM format>`
    - Certificate Chain:
          - `4:<Parent Certificate's Issuer>:<Parent Certificate's Serial Number>` : `<a list of child certificates (issuer/serialNum pairs)>`
    - Proposed but not approved revocation of certificates:
        - `5:<Certificate's Issuer>:<Certificate's Serial Number>` : `<List of approved trustee account IDs>`
    - CRL (Certificate Revocation List):
        - `6` : `CRL (Certificate Revocation List)`
- KV store name: `modelinfo`
    - Model Infos 
        - `1:<vid>:<pid>` : `<model info>`
    - Vendor to products (models) index:
        - `2:<vid>` : `<list of pids + metadata>`
- KV store name: `compliancetest`
    - Test results for every model
        - `1:<vid>:<pid>` : `<list of test results>`
- KV store name: `compliance`
    - Compliance results for every model       
       - `1:<vid>:<pid>` : `<compliance info>`
    - A list of compliant models (`pid`s) for the given vendor. 
       - `2:<vid>` : `<compliance pids>`  
    - A list of revoked models (`pid`s) for the given vendor.       
       - `3:<vid>` : `<revocation pids>`     
- KV store name: `auth`
    - Proposed but not approved accounts
      - `1:<address>` : `<account info> + <list of approvers>`
    - Approved accounts
      - `2:<address>` : `<account info>` 
    - Proposed but not approved accounts to be revoked
      - `3:<address>` : `<account info> + <list of approvers>`
- KV store name: `validator`
    - Validators
    - Proposed but not approved accounts to be removed
   
## X509 PKI

#### PROPOSE_ADD_X509_ROOT_CERT
Proposes a new self-signed root certificate.

If it's sent by a non-Trustee account, or more than 1 Trustee signature is required to add a root certificate, 
then the certificate
will be in a pending state until sufficient number of other Trustee's approvals is received.

The certificate is immutable. It can only be revoked by either the owner or a quorum of Trustees.

- Parameters:
  - `cert`: PEM-encoded certificate
- In State:
  - `pki` store  
  - `1:<Certificate's Issuer>:<Certificate's Serial Number>` : `<Certificate in PEM format> + <List of approved trustee account IDs>`
  - `2:<Certificate's Issuer>:<Certificate's Serial Number>` : `<Certificate in PEM format>` (if just 1 Trustee is required)  
- Who can send: 
    - Any role
- CLI command: 
    -   `zblcli tx pki propose-add-x509-root-cert .... `
- REST API: 
    -   POST `/pki/certs/proposed/root`
    


#### APPROVE_ADD_X509_ROOT_CERT
Approves the proposed root certificate.

The certificate is not active until sufficient number of Trustees approve it. 

- Parameters:
  - `issuer`: string  - proposed certificates's Issuer
  - `serialNumber`: string  - proposed certificates's Serial Number
- In State:
  - `pki` store  
  - `1:<Certificate's Issuer>:<Certificate's Serial Number>` : `<Certificate in PEM format> + <List of approved trustee account IDs>`
  - `2:<Certificate's Issuer>:<Certificate's Serial Number>` : `<Certificate in PEM format>`
- Who can send: 
    - Trustee
- CLI command: 
    -   `zblcli tx pki approve-add-x509-root-cert .... `
- REST API: 
    -   PATCH `/pki/certs/proposed/root/<issuer>/<serialNumber>`

#### ADD_X509_CERT
Adds an intermediate or leaf X509 certificate signed by a chain of certificates which must be
already present on the ledger.

The certificate is immutable. It can only be revoked by either the owner or a quorum of Trustees.

- Parameters:
  - `cert`: PEM-encoded certificate
- In State:
  - `pki` store  
  - `3:<Certificate's Issuer>:<Certificate's Serial Number>` : `<Certificate in PEM format>`
  - `4:<Parent Certificate's Issuer>:<Parent Certificate's Serial Number>` : `<a list of child certificates (issuer/serialNum pairs)>`
- Who can send: 
    - Any role
- CLI command: 
    -   `zblcli tx pki add-x509-cert .... `
- REST API: 
    -   POST `/pki/certs`

#### REVOKE_X509_CERT
Revokes the given X509 certificate (either intermediate or leaf).
All the certificates in the chain signed by the revoked certificate will be revoked as well.

Only the owner (sender) can revoke the certificate.
Root certificates can not be revoked this way, use  `PROPOSE_X509_CERT_REVOC` and `APPROVE_X509_ROOT_CERT_REVOC` instead.  

- Parameters:
  - `issuer`: string  - revoked certificates's Issuer
  - `serialNumber`: string  - revoked certificates's Serial Number
- In State:
  - `pki` store  
  - `2:<Certificate's Issuer>:<Certificate's Serial Number>` : `<Certificate in PEM format>`  
  - `3:<Certificate's Issuer>:<Certificate's Serial Number>` : `<Certificate in PEM format>`
  - `4:<Parent Certificate's Issuer>:<Parent Certificate's Serial Number>` : `<a list of child certificates (issuer/serialNum pairs)>`
  - `6` : `CRL (Certificate Revocation List)`

- Who can send: 
    - Any role; owner
- CLI command: 
    -   `zblcli tx pki revoke-x509-cert .... `
- REST API: 
    -   DELETE `/pki/certs/<issuer>/<serialNumber>`

#### PROPOSE_REVOKE_X509_CERT
Proposes revocation of the given X509 certificate (either root, intermediate or leaf) by a Trustee.

All the certificates in the chain signed by the revoked certificate will be revoked as well.

If more than 1 Trustee signature is required to revoke a root certificate, 
then the certificate will be in a pending state until sufficient number of other Trustee's approvals is received.

- Parameters:
  - `issuer`: string  - revoked certificates's Issuer
  - `serialNumber`: string  - revoked certificates's Serial Number
- In State:
  - `pki` store  
  - `5:<Certificate's Issuer>:<Certificate's Serial Number>` : `<List of approved trustee account IDs>`
- Who can send: 
    - Trustee
- CLI command: 
    -   `zblcli tx pki propose-revoke-x509-cert .... `
- REST API: 
    -   PUT `/pki/certs/proposed/revoked/<issuer>/<serialNumber>`
    


#### APPROVE_REVOKE_X509_ROOT_CERT
Approves the revocation of the given X509 certificate (either root, intermediate or leaf).
All the certificates in the chain signed by the revoked certificate will be revoked as well.

The revocation is not applied until sufficient number of Trustees approve it. 

- Parameters:
  - `issuer`: string  - revoked certificates's Issuer
  - `serialNumber`: string  - revoked certificates's Serial Number
- In State:
  - `pki` store  
  - `5:<Certificate's Issuer>:<Certificate's Serial Number>` : `<List of approved trustee account IDs>`
  - `2:<Certificate's Issuer>:<Certificate's Serial Number>` : `<Certificate in PEM format>`  
  - `3:<Certificate's Issuer>:<Certificate's Serial Number>` : `<Certificate in PEM format>`
  - `4:<Parent Certificate's Issuer>:<Parent Certificate's Serial Number>` : `<a list of child certificates (issuer/serialNum pairs)>`
  - `6` : `CRL (Certificate Revocation List)`
- Who can send: 
    - Trustee
- CLI command: 
    -   `zblcli tx pki approve-revoke-x509-cert .... `
- REST API: 
    -   PATCH `/pki/certs/proposed/revoked/<issuer>/<serialNumber>`
        
#### GET_ALL_PROPOSED_X509_ROOT_CERTS
Gets all proposed but not approved root certificates.

- Parameters: No
- CLI command: 
    -   `zblcli query pki all-proposed-x509-root-certs .... `
- REST API: 
    -   GET `/pki/certs/proposed/root`

#### GET_PROPOSED_X509_ROOT_CERT
Gets a proposed but not approved root certificate with the given 
issuer and serial number attributes.

- Parameters:
  - `issuer`: string - certificates's Issuer
  - `serialNumber`: string - certificates's Serial Number
- CLI command: 
    -   `zblcli query pki proposed-x509-root-cert .... `
- REST API: 
    -   GET `/pki/certs/proposed/root/<issuer>/<serialNumber>`

#### GET_ALL_X509_ROOT_CERTS
Gets all approved root certificates.

- Parameters: No
- CLI command: 
    -   `zblcli query pki all-x509-root-certs .... `
- REST API: 
    -   GET `/pki/certs/root/`


#### GET_X509_CERT
Gets a certificate (either root, intermediate or leaf) by the given 
issuer and serial number attributes.

- Parameters:
  - `issuer`: string - certificates's Issuer
  - `serial Number`: string - certificates's Serial Number
- CLI command: 
    -   `zblcli query pki x509-cert .... `
- REST API: 
    -   GET `/pki/certs/<issuer>/<serialNumber>`

#### GET_ALL_X509_CERTS
Gets all certificates (root, intermediate and leaf).
Can optionally be filtered by the root certificate's issuer and serial number so that 
only the certificate chains started with the given root certificate are returned.   

`GET_ALL_X509_CERTS_SINCE` can be used to incrementally update the list stored locally. 

- Parameters:
  - `rootIssuer`: string (optional) - root certificates's Issuer
  - `rootSerialNumber`: string (optional) - root certificates's Serial Number
- CLI command: 
    -   `zblcli query pki all-x509-certs .... `
- REST API: 
    -   GET `/pki/certs`
    -   GET `/pki/certs?rootIssuer=<>;rootSerialNumber={}`

#### GET_ALL_X509_CERTS_SINCE
Gets all certificates (root, intermediate and leaf) which has been added since 
the given ledger's `height`.

Can optionally be filtered by the root certificate's issuer and serial number so that 
only the certificate chains started with the given root certificate are returned.   

- Parameters:
  - `since`: integer - the last ledger's height the user has locally.
  - `rootIssuer`: string (optional) - root certificates's Issuer
  - `rootSerialNumber`: string (optional) - root certificates's Serial Number  
- CLI command: 
    -   `zblcli query pki all-x509-certs-delta .... `
- REST API: 
    -   GET `/pki/certs?since=<>`
    -   GET `/pki/certs?since=<>;rootIssuer=<>;rootSerialNumber={}`
   

#### GET_ALL_PROPOSED_X509_CERTS_TO_REVOKE
Gets all proposed but not approved certificates to be revoked.

- Parameters: No
- CLI command: 
    -   `zblcli query pki all-proposed-x509-certs-to-revoke .... `
- REST API: 
    -   GET `/pki/certs/proposed/revoked`

#### GET_PROPOSED_X509_CERT_TO_REVOKE
Gets a proposed but not approved certificate to be revoked.

- Parameters:
  - `issuer`: string - certificates's Issuer
  - `serialNumber`: string - certificates's Serial Number
- CLI command: 
    -   `zblcli query pki proposed-x509-cert-to-revoke ....`
- REST API: 
    -   GET `/pki/certs/proposed/revoked/<issuer>/<serialNumber>`

#### GET_CRL
Gets all revoked certificates (CRL or certificate revocation list).
   
- Parameters: No
- CLI command: 
    -   `zblcli query pki all-revoked-certs .... `
- REST API: 
    -   GET `/pki/certs/revoked`
    
    
## MODEL INFO

#### ADD_MODEL_INFO
Adds a new Model Info identified by a unique combination of `vid` (vendor ID) and `pid` (product ID).

Only some of Model Info fields can be edited (see `EDIT_MODEL_INFO`). If other fields need to be edited - 
a new model info with a new `vid` or `pid` can be created.

- Parameters:
    - `vid`: 16 bits int
    - `pid`: 16 bits int
    - `cid`: 16 bits int (optional)
    - `Name`: string
    - `owner`: bech32 encoded address
    - `description`: string
    - `SKU`: string
    - `firmwareVersion`: string
    - `hardwareVersion`: string
    - `tisOrTrpTestingCompleted`: bool
    - `custom`: string (optional)
- In State:
  - `modelinfo` store  
  - `1:<vid>:<pid>` : `<model info>`
  - `2:<vid>` : `<list of pids + metadata>`
- Who can send: 
    - Vendor
- CLI command: 
    -   `zblcli tx modelinfo add-model .... `
- REST API: 
    -   POST `/modelinfo/models`

#### EDIT_MODEL_INFO
Edits an existing Model Info identified by a unique combination of `vid` (vendor ID) and `pid` (product ID)
by the owner.

Only the fields listed below (besides `vid` and `pid`) can be edited. If other fields need to be edited - 
a new model info with a new `vid` or `pid` can be created.

All non-edited fields remain the same. 

- Parameters:
    - `vid`: 16 bits int
    - `pid`: 16 bits int
    - `cid`: 16 bits int (optional)
    - `description`: string (optional)
    - `tisOrTrpTestingCompleted`: bool (optional)
    - `custom`: string (optional)
- In State:
  - `modelinfo` store  
  - `1:<vid>:<pid>` : `<model info>`
  - `2:<vid>` : `<list of pids + metadata>`
- Who can send: 
    - Vendor; owner
- CLI command: 
    -   `zblcli tx modelinfo edit-model .... `
- REST API: 
    -   PUT `/modelinfo/models/vid/pid`


#### GET_ALL_MODEL_INFO
Gets all Model Infos for all vendors.

- Parameters: No
- CLI command: 
    -   `zblcli query modelinfo all-models ...`
- REST API: 
    -   GET `/modelinfo/models`
    
#### GET_VENDOR_MODEL_INFO
Gets all Model Info by the given Vendor (`vid`).

- Parameters:
    - `vid`: 16 bits int
- CLI command: 
    -   `zblcli query modelinfo vendor-models .... `
- REST API: 
    -   GET `/modelinfo/models/vid`
    
#### GET_MODEL_INFO
Gets a Model Info with the given `vid` (vendor ID) and `pid` (product ID).

- Parameters:
    - `vid`: 16 bits int
    - `pid`: 16 bits int
- CLI command: 
    -   `zblcli query modelinfo model.... `
- REST API: 
    -   GET `/modelinfo/models/vid/pid`

#### GET_VENDORS    
Get a list of all Vendors (`vid`s). 

- Parameters: No
- CLI command: 
    -   `zblcli query modelinfo vendors .... `
- REST API: 
    -   GET `/modelinfo/vendors`

## TEST_DEVICE_COMPLIANCE

#### ADD_TEST_RESULT
Submits result of a compliance testing for the given device (`vid` and `pid`).
The test result can be a blob of data or a reference (URL) to an external storage.

Multiple test results (potentially from different test houses) can be added for the same device type. 

The test result is immutable and can not be deleted or removed after submitting. 
Another test result can be submitted instead.

- Parameters:
    - `vid`: 16 bits int
    - `pid`: 16 bits int
    - `testResult`: string
- In State:
  - `compliancetest` store  
  - `1:<vid>:<pid>` : `<list of test results>`
- Who can send: 
    - TestHouse
- CLI command: 
    -   `zblcli tx compliancetest add-test-result .... `
- REST API: 
    -   POST `/compliancetest/testresults`

#### GET_TEST_RESULT
Gets a test result for the given `vid` (vendor ID) and `pid` (product ID).

- Parameters:
    - `vid`: 16 bits int
    - `pid`: 16 bits int
- CLI command: 
    -   `zblcli query compliancetest .... `
- REST API: 
    -   GET `/compliancetest/testresults/vid/pid`

## CERTIFY_DEVICE_COMPLIANCE

#### CERTIFY_MODEL
Attests compliance of the Model to the ZB standard.

`REVOKE_MODEL_CERTIFICATION` should be used for revoking (disabling) the compliance.
It's possible to call it for revoked models to enable them back. 

The corresponding Model Info and test results must be present on ledger.

It must be called for every compliant device for use cases where compliance 
is tracked on ledger.
It can be used by use cases where only revocation is tracked on the ledger to remove a Model
from the revocation list.
 
- Parameters:
    - `vid`: 16 bits int
    - `pid`: 16 bits int
    - `certificationDate`: rfc3339 encoded date
    - `certificationType` (optional): string  - `zb` is the default and the only supported value now
- In State:
  - `compliance` store  
  - `1:<vid>:<pid>` : `<compliance info>`
  - `2:<vid>` : `<compliance pids>`  
  - `3:<vid>` : `<revoked pids>`  
- Who can send: 
    - ZBCertificationCenter
- CLI command: 
    -   `zblcli tx compliance certify-model .... `
- REST API: 
    -   PUT `/compliance/certified/vid/pid`
    
#### REVOKE_MODEL_CERTIFICATION
Revoke compliance of the Model to the ZB standard.

The corresponding Model Info and test results are not required to be on the ledger 
to be used in cases where revocation only is tracked on the ledger.

It can be used in cases where every compliance result 
is written on the ledger (`CERTIFY_MODEL` was called), or
 cases where only revocation list is stored on the ledger.
 
- Parameters:
    - `vid`: 16 bits int
    - `pid`: 16 bits int
    - `revocationDate`: rfc3339 encoded date    
- In State:
  - `compliance` store  
  - `1:<vid>:<pid>` : `<compliance info>`
  - `2:<vid>` : `<compliance pids>`  
  - `3:<vid>` : `<revocation pids>`  
- Who can send: 
    - ZBCertificationCenter
- CLI command: 
    -   `zblcli tx compliance revoke-model-certification .... `
- REST API: 
    -   PUT `/compliance/revoked/vid/pid`    
    
#### GET_CERTIFIED_MODEL
Gets a boolean if the given Model (identified by the `vid` and `pid`) is compliant to ZB standards. 

This is the aggregation of compliance and
revocation information for every vid/pid. It should be used in cases where compliance 
is tracked on the ledger.
 
- Parameters:
    - `vid`: 16 bits int
    - `pid`: 16 bits int
- CLI command: 
    -   `zblcli query compliance certified-model .... `
- REST API: 
    -   GET `/compliance/certified/vid/pid`

#### GET_REVOKED_MODEL
Gets a boolean if the given Model's compliance is revoked. 

It contains information about revocation only, so  it should be used in cases
 where only revocation is tracked on the ledger.
 
- Parameters:
    - `vid`: 16 bits int
    - `pid`: 16 bits int
- CLI command: 
    -   `zblcli query compliance revoked-model .... `
- REST API: 
    -   GET `/compliance/revoked/vid/pid`

#### GET_VENDOR_CERTIFIED_MODELS
Gets all the Models (`pid`s) issued by the given Vendor (`vid`) which are complaint to ZB standards.

This is the aggregation of compliance and
revocation information for every vid/pid. It should be used in cases where compliance is tracked on ledger.

- Parameters:
    - `vid`: 16 bits int
- CLI command: 
    -   `zblcli query compliance certified-vendor-models.... `
- REST API: 
    -   GET `/compliance/certified/vid`
    
#### GET_VENDOR_REVOKED_MODELS
Gets all the Models (`pid`s) which certification is revoked for the given Vendor (`vid`).

It contains information about revocation only, so  it should be used in cases
 where only revocation is tracked on the ledger.


- Parameters:
    - `vid`: 16 bits int
- CLI command: 
    -   `zblcli query compliance revoked-vendor-models .... `
- REST API: 
    -   GET `/compliance/revoked/vid`
    
#### GET_ALL_CERTIFIED_MODELS
Gets all compliant Models (`pid`s) for all the vendors (`vid`s).

This is the aggregation of compliance and
revocation information for every vid/pid. It should be used in cases where compliance is tracked on ledger.
 
`GET_ALL_CERTIFIED_MODELS_SINCE` can be used to incrementally update the list stored locally. 
 
- Parameters: No
- CLI command: 
    -   `zblcli query compliance all-certified-models `
- REST API: 
    -   GET `/compliance/certified`
 

#### GET_ALL_REVOKED_MODELS
Gets all revoked Models (`pid`s) for all vendors (`vid`s).

It contains information about revocation only, so  it should be used in cases
 where only revocation is tracked on the ledger.
 
`GET_ALL_REVOKED_MODELS_SINCE` can be used to incrementally update the list stored locally. 

- Parameters: No
- CLI command: 
    -   `zblcli query compliance all-revoked-models.... `
- REST API: 
    -   GET `/compliance/revoked`

#### GET_ALL_CERTIFIED_MODELS_SINCE
Gets a delta of all compliant Models (`pid`s) for every vendor (`vid`s) which has been added or revoked since 
the given ledger's `height`.

This is the aggregation of compliance and
revocation information for every vid/pid. It should be used in cases where compliance is tracked on ledger.
 
- Parameters: 
  - `since`: integer - the last ledger's height the user has locally.
- CLI command: 
    -   `zblcli query compliance all-certified-models-delta `
- REST API: 
    -   GET `/compliance/certified?since=<>`
    
#### GET_ALL_REVOKED_MODELS_SINCE
Gets a delta of all revoked Models (`pid`s) for every vendor (`vid`s) which has been added or revoked since 
the given ledger's `height`.

This is the aggregation of compliance and
revocation information for every vid/pid. It should be used in cases where compliance is tracked on ledger.
 
- Parameters: 
  - `since`: integer - the last ledger's height the user has locally.
- CLI command: 
    -   `zblcli query compliance revoked all-revoked-models-delta `
- REST API: 
    -   GET `/compliance/revoked?since=<>`
    
## AUTH

#### PROPOSE_ADD_ACCOUNT
Proposes a new Account with the given address, public key and role.

If more than 1 Trustee signature is required to add the account, the account
will be in a pending state until sufficient number of approvals is received.

- Parameters:
    - `address`
    - `pubKey`
    - `pubKeyType`
    - `role` 
- In State:
  - `auth` store  
  - `1:<address>` : `<account info> + <list of approvers>`
  - `2:<address>` : `<account info>` (if just 1 Trustee is required)  
- Who can send: 
    - Trustee
- CLI command: 
    -   `zblcli tx auth propose-add-account .... `
- REST API: 
    -   POST `/auth/accounts/proposed`
    
#### APPROVE_ADD_ACCOUNT
Approves the proposed account.    

The account is not active until sufficient number of Trustees approve it. 

- Parameters:
    - `address`
- In State:
  - `auth` store  
  - `1:<address>` : `<account info> + <list of approvers>`  
  - `2:<address>` : `<account info>`
- Who can send: 
    - Trustee
- CLI command: 
    -   `zblcli tx auth approve-add-account .... `
- REST API: 
    -   PATCH `/auth/accounts/proposed/<address>`
    
    
#### ROTATE_KEY
Rotate's the Account's public key by the owner.

- Parameters:
    - `pubKey`
    - `pubKeyType`
- In State:
  - `auth` store  
  - `1:<address>` : `<account info>`
- Who can send: 
    - Any role; owner
- CLI command: 
    -   `zblcli tx auth rotate-key .... `
- REST API: 
    -   PUT `/auth/accounts/<address>`
    
#### PROPOSE_REVOKE_ACCOUNT
Proposes revocation of the Account with the given address.

If more than 1 Trustee signature is required to revoke the account, the revocation
will be in a pending state until sufficient number of approvals is received.

- Parameters:
    - `address`
- In State:
  - `auth` store  
  - `3:<address>` : `<account info> + <list of approvers>`
  - `2:<address>` : `<account info>` (if just 1 Trustee is required)  
- Who can send: 
    - Trustee
- CLI command: 
    -   `zblcli tx auth propose-revoke-account .... `
- REST API: 
    -   POST `/auth/accounts/revoked`
    
#### APPROVE_REVOKE_ACCOUNT
Approves the proposed revocation of the account.    

The account is not revoked until sufficient number of Trustees approve it. 

- Parameters:
    - `address`
- In State:
  - `auth` store  
  - `3:<address>` : `<account info> + <list of approvers>`  
  - `2:<address>` : `<account info>`
- Who can send: 
    - Trustee
- CLI command: 
    -   `zblcli tx auth approve-revoke-account .... `
- REST API: 
    -   PATCH `/auth/accounts/revoked/<address>`
    
#### GET_ALL_PROPOSED_ACCOUNTS
Gets all proposed but not approved accounts.

- Parameters: No
- CLI command: 
    -   `zblcli query auth all-proposed-accounts .... `
- REST API: 
    -   GET `/auth/accounts/proposed`
    
#### GET_ALL_ACCOUNTS
Gets all accounts.

- Parameters: No
- CLI command: 
    -   `zblcli query auth all-accounts .... `
- REST API: 
    -   GET `/auth/accounts`           

#### GET_ACCOUNTS
Gets an accounts by the address.

- Parameters:
    - `address`
- CLI command: 
    -   `zblcli query auth account .... `
- REST API: 
    -   GET `/auth/accounts/<address>`         
    
#### GET_ALL_PROPOSED__ACCOUNTS_TO_REVOKE
Gets all proposed but not approved accounts to be revoked.

- Parameters: No
- CLI command: 
    -   `zblcli query auth all-proposed-accounts-to-revoke .... `
- REST API: 
    -   GET `/auth/accounts/proposed/revoked`
    
    
## VALIDATOR_NODE                      

#### ADD_VALIDATOR_NODE
Adds a new Validator node.

- Parameters: Same as in `cosmos-sdk/MsgCreateValidator`
- In State: Same as in Tendermint/Cosmos-sdk
- Who can send: 
    - NodeAdmin
- CLI command: 
    -   `zblcli tx validator add-node .... `
- REST API: 
    -   PUT `/validators/<address>`
    
#### UPDATE_VALIDATOR_NODE
Updates the Validator node by the owner.

- Parameters: Same as in `cosmos-sdk/MsgCreateValidator`
- In State: Same as in Tendermint/Cosmos-sdk
- Who can send: 
    - NodeAdmin; owner
- CLI command: 
    -   `zblcli tx validator update-node .... `
- REST API: 
    -   PUT `/validators/<address>`    

#### REMOVE_VALIDATOR_NODE
Deletes the Validator node (removes from the validator set) by the owner.

- Parameters:
    - `address`
- In State: Same as in Tendermint/Cosmos-sdk
- Who can send: 
    - NodeAdmin; owner
- CLI command: 
    -   `zblcli tx validator remove-node .... `
- REST API: 
    -   DELETE `/validators/<address>`

#### PROPOSE_REMOVE_VALIDATOR_NODE
Proposes removing the Validator node from the validator set by a Trustee. 

If more than 1 Trustee signature is required to remove a node, the removal
will be in a pending state until sufficient number of approvals is received.

- Parameters:
    - `address`
- In State: Same as in Tendermint/Cosmos-sdk
- Who can send: 
    - Trustee
- CLI command: 
    -   `zblcli tx validator propose-remove-node .... `
- REST API: 
    -   POST `/validators/proposed/removed`

#### APPROVE_REMOVE_VALIDATOR_NODE
Approves removing of the Validator node by a Trustee. 

The account is not removed until sufficient number of Trustees approve it. 

- Parameters:
    - `address`
- In State: Same as in Tendermint/Cosmos-sdk
- Who can send: 
    - Trustee
- CLI command: 
    -   `zblcli tx validator approve-remove-node .... `
- REST API: 
    -   PATCH `/validators/proposed/removed/<address>`
            
#### GET_ALL_VALIDATORS
Gets all validator nodes.

- Parameters: No
- CLI command: 
    -   `zblcli query validator all-nodes .... `
- REST API: 
    -   GET `/validators`   
     
#### GET_VALIDATOR
Gets a validator node.

- Parameters:
    - `address`
- CLI command: 
    -   `zblcli query validator node .... `
- REST API: 
    -   GET `/validators/<address>`   
    
#### GET_ALL_PROPOSED_VALIDATORS_TO_REMOVE
Gets all proposed but not approved validator nodes to be removed.

- Parameters: No
- CLI command: 
    -   `zblcli query validator all-proposed-nodes-to-remove .... `
- REST API: 
    -   GET `/validators/proposed/removed`   