# Transactions and Queries

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
    - CertificationCenter
    - NodeAdmin   
- All read (get) requests return the current `height` of the ledger in addition to the
requested data. The `height` can be used to get a delta (changes) from the last state that the user has.
This is useful to avoid correlation by the sender's IP address.        

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
        - `4:<Child Certificate's Issuer>:<Child Certificate's Serial Number>` : `<Parent certificates issuer/serialNumber pairs>`
    - Proposed but not approved revocation of certificates:
        - `5:<Certificate's Issuer>:<Certificate's Serial Number>` : `<List of approved trustee account IDs>`
    - CRL (Certificate Revocation List):
        - `6` : `CRL (Certificate Revocation List)`
- KV store name: `modeilinfo`
    - Model Infos 
        - `1:<vid>:<pid>` : `<model info>`
    - Vendor to Products (Models) index:
        - `2:<vid>` : `<list of pids>`
- KV store name: `testresult`
    - Test results for every model
        - `1:<vid>:<pid>` : `<list of test results>`
- KV store name: `compliance`
    - Compliance results for every model       
        - `1:<vid>:<pid>` : `<compliance bool>` 
   
## X509 PKI

#### PROPOSE_X509_ROOT_CERT
Proposes a new self-signed root certificate.
The certificate is not applied until sufficient number of Trustees approve it.

The certificate is immutable. It can only be revoked by either the owner or a quorum of Trustees.

- Parameters:
  - `cert`: PEM-encoded certificate
- In State:
  - `pki` store  
  - `1:<Certificate's Issuer>:<Certificate's Serial Number>` : `<Certificate in PEM format> + <List of approved trustee account IDs>`
- Who can send: 
    - Trustee
- CLI command: 
    -   `zblcli tx pki propose-x509-root-cert .... `
- REST API: 
    -   POST `/pki/certs/proposed/root`
    


#### APPROVE_X509_ROOT_CERT
Approves the proposed root certificate.
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
    -   `zblcli tx pki approve-x509-root-cert .... `
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
  - `4:<Certificate's Issuer>:<Certificate's Serial Number>` : `<Parent certificate Chain's issuer/serialNumber pairs>`
- Who can send: 
    - Any role
- CLI command: 
    -   `zblcli tx pki add-x509-cert .... `
- REST API: 
    -   POST `/pki/certs`

#### PROPOSE_X509_CERT_REVOC
Proposes revocation of the given X509 certificate (either root, intermediate or leaf).
All the certificates in the chain signed by the revoked certificate will be revoked as well.

The revocation is not applied until sufficient number of Trustees approve it. 

- Parameters:
  - `issuer`: string  - revoked certificates's Issuer
  - `serialNumber`: string  - revoked certificates's Serial Number
- In State:
  - `pki` store  
  - `5:<Certificate's Issuer>:<Certificate's Serial Number>` : `<List of approved trustee account IDs>`
- Who can send: 
    - Trustee
- CLI command: 
    -   `zblcli tx pki propose-x509-cert-revok .... `
- REST API: 
    -   POST `/pki/certs/revoked`
    


#### APPROVE_X509_ROOT_CERT_REVOC
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
  - `4:<Certificate's Issuer>:<Certificate's Serial Number>` : `<Certificate Chain's issuer/serialNumber pairs>`
  - `6` : `CRL (Certificate Revocation List)`
- Who can send: 
    - Trustee
- CLI command: 
    -   `zblcli tx pki approve-x509-cert-revoc .... `
- REST API: 
    -   PATCH `/pki/certs/revoked/<issuer>/<serialNumber>`
        
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

#### GET_ALL_PROPOSED_X509_REVOKED_CERTS
Gets all proposed but not approved certificates to be revoked.

- Parameters: No
- CLI command: 
    -   `zblcli query pki all-proposed-x509-revoked-certs .... `
- REST API: 
    -   GET `/pki/certs/revoked`

#### GET_PROPOSED_X509_REVOKED_CERT
Gets a proposed but not approved certificate to be revoked.

- Parameters:
  - `issuer`: string - certificates's Issuer
  - `serialNumber`: string - certificates's Serial Number
- CLI command: 
    -   `zblcli query pki proposed-x509-revoked-cert .... `
- REST API: 
    -   GET `/pki/certs/revoked/<issuer>/<serialNumber>`

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
Cn optionally bi filtered by the root certificate's issuer and serial number so that 
only the certificate chains started with the given root certificate are returned.   

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

- Parameters:
  - `since`: integer - the last ledger's height the user has locally.
- CLI command: 
    -   `zblcli query pki all-x509-certs-since .... `
- REST API: 
    -   GET `/pki/certs?since=<>`
   
## MODEL INFO

#### ADD_MODEL_INFO
Adds a new Model Info identified by a unique combination of `vid` (vendor ID) and `pid` (product ID).

The Model Info is immutable. If it needs to be edited - a new model info with a new `vid` or `pid` 
can be created.

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
    - `custom`: string
- In State:
  - `modelinfo` store  
  - `1:<vid>:<pid>` : `<model info>`
  - `2:<vid>` : `<list of pids>`
- Who can send: 
    - Vendor
- CLI command: 
    -   `zblcli tx modelinfo add .... `
- REST API: 
    -   POST `/modelinfo`

#### GET_ALL_MODEL_INFO
Gets all Model Info. 

- Parameters: No
- CLI command: 
    -   `zblcli query modelinfo .... `
- REST API: 
    -   GET `/modelinfo`
    
#### GET_VENDOR_MODEL_INFO
Gets all Model Info by the given `vid` (vendor ID).

- Parameters:
    - `vid`: 16 bits int
- CLI command: 
    -   `zblcli query modelinfo .... `
- REST API: 
    -   GET `/modelinfo/vid`
    
#### GET_MODEL_INFO
Gets a Model Info by the given `vid` (vendor ID) and `pid` (product ID).

- Parameters:
    - `vid`: 16 bits int
    - `pid`: 16 bits int
- CLI command: 
    -   `zblcli query modelinfo .... `
- REST API: 
    -   GET `/modelinfo/vid/pid`

## TEST_DEVICE_COMPLIANCE

#### ADD_TEST_RESULT
Submits result of compliance testing for the given device (`vid` and `pid`).
The test result can be a blob of data or a reference (URL) to an external storage.

Multiple test results (potentially from different test houses) can be added for the same device type. 

The test result is immutable and can not be deleted or removed after submitting. 
Another test result can be submitted instead.

- Parameters:
    - `vid`: 16 bits int
    - `pid`: 16 bits int
    - `testResult`: string
- In State:
  - `testresult` store  
  - `1:<vid>:<pid>` : `<list of test results>`
- Who can send: 
    - TestHouse
- CLI command: 
    -   `zblcli tx testresult add .... `
- REST API: 
    -   POST `/testresult`

#### GET_TEST_RESULT
Gets a test result for the given `vid` (vendor ID) and `pid` (product ID).

- Parameters:
    - `vid`: 16 bits int
    - `pid`: 16 bits int
- CLI command: 
    -   `zblcli query testresult .... `
- REST API: 
    -   GET `/testresult/vid/pid`

## ATTEST_DEVICE_COMPLIANCE

#### ATTEST_MODEL
Attests compliance of the Model to the ZB standard.
Can be used to enable or disable the compliance (that is it's mutable). 
 
- Parameters:
    - `vid`: 16 bits int
    - `pid`: 16 bits int
    - `compliance`: bool
- In State:
  - `compliance` store  
  - `1:<vid>:<pid>` : `<compliance bool>`
- Who can send: 
    - CertificationCenter
- CLI command: 
    -   `zblcli tx compliance add .... `
- REST API: 
    -   PUT `/compliance/vid/pid`
    
#### GET_MODEL_COMPLIANCE
Gets a boolean of the given Model (identified by the `vid` and `pid`) is compliant to ZB standards. 
 
- Parameters:
    - `vid`: 16 bits int
    - `pid`: 16 bits int
- CLI command: 
    -   `zblcli query compliance .... `
- REST API: 
    -   GET `/compliance/vid/pid`
    
#### GET_MODEL_COMPLIANCE_DELTA
Gets a boolean of the given Model (identified by the `vid` and `pid`) is compliant to ZB standards. 
 
- Parameters:
    - `vid`: 16 bits int
    - `pid`: 16 bits int
- CLI command: 
    -   `zblcli query compliance .... `
- REST API: 
    -   GET `/compliance/vid/pid`            