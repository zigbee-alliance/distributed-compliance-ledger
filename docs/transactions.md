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
   
## X509 PKI

#### PROPOSE_X509_ROOT_CERT
Proposes a new self-signed root certificate.
The certificate is not applied until sufficient number of Trustees approve it.

- Parameters:
  - `Cert`: PEM-encoded certificate
- In State:
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
  - `Issuer`: string  - proposed certificates's Issuer
  - `SerialNumber`: string  - proposed certificates's Serial Number
- In State:
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

- Parameters:
  - `Cert`: PEM-encoded certificate
- In State:
  - `3:<Certificate's Issuer>:<Certificate's Serial Number>` : `<Certificate in PEM format>`
  - `4:<Certificate's Issuer>:<Certificate's Serial Number>` : `<Certificate Chain's issuer/serialNumber pairs>`
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
  - `Issuer`: string  - revoked certificates's Issuer
  - `SerialNumber`: string  - revoked certificates's Serial Number
- In State:
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
  - `Issuer`: string  - revoked certificates's Issuer
  - `SerialNumber`: string  - revoked certificates's Serial Number
- In State:
  - `5:<Certificate's Issuer>:<Certificate's Serial Number>` : `<List of approved trustee account IDs>`
  - `2:<Certificate's Issuer>:<Certificate's Serial Number>` : `<Certificate in PEM format>`  
  - `3:<Certificate's Issuer>:<Certificate's Serial Number>` : `<Certificate in PEM format>`
  - `4:<Certificate's Issuer>:<Certificate's Serial Number>` : `<Certificate Chain's issuer/serialNumber pairs>`
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
  - `Issuer`: string - certificates's Issuer
  - `SerialNumber`: string - certificates's Serial Number
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
  - `Issuer`: string - certificates's Issuer
  - `SerialNumber`: string - certificates's Serial Number
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
  - `Issuer`: string - certificates's Issuer
  - `Serial Number`: string - certificates's Serial Number
- CLI command: 
    -   `zblcli query pki x509-cert .... `
- REST API: 
    -   GET `/pki/certs/<issuer>/<serialNumber>`

#### GET_ALL_X509_CERTS
Gets all certificates (root, intermediate and leaf).
Cn optionally bi filtered by the root certificate's issuer and serial number so that 
only the certificate chains started with the given root certificate are returned.   

- Parameters:
  - `RootIssuer`: string (optional) - root certificates's Issuer
  - `RootSerialNumber`: string (optional) - root certificates's Serial Number
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

##### Parameters:
- Vid: 16 bits int
- Pid: 16 bits int
- Cid: 16 bits int (optional)
- Name: string
- Owner: bech32 encoded address
- Description: string
- SKU: string
- FirmwareVersion: string
- HardwareVersion: string
- CertificateID: string (optional)
- CertifiedDate: rfc3339 encoded date
- TisOrTrpTestingCompleted: bool
- Custom: string

The `owner` field must be equal to the signer's account ID.

Unique key: `vid:pid`

##### Signature:
-  Required

#### UPDATE_MODEL_INFO
Updates an existing Model Info identified by the given  `vid` (vendor ID) and `pid` (product ID).
Only the fields that need to be updated can be specified. Unspecified fields are not changed.

##### Parameters:

- Vid: 16 bits int
- Pid: 16 bits int
- Cid: 16 bits int (optional)
- Name: string (optional)
- Owner: bech32 encoded address (optional)
- Description: string (optional)
- SKU: string (optional)
- FirmwareVersion: string (optional)
- HardwareVersion: string (optional)
- CertificateID: string (optional)
- CertifiedDate: rfc3339 encoded date (optional)
- TisOrTrpTestingCompleted: bool (optional)
- Custom: string (optional)

##### Unique key:
`vid:pid`

##### Who can send:
-  Manufacturer

#### GET_MODEL_INFO
Gets a Model Info by the given  `vid` (vendor ID) and `pid` (product ID).

##### Parameters:
- Vid: 16 bits int
- Pid: 16 bits int

##### Who can send:
-  Anyone




 
