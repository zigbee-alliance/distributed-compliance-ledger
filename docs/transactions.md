# Transactions and Queries

## General

- Every transaction (write request) must be signed
- The signer's public key must be stored on the ledger via `ACCOUNT` transaction. 
- Depending on the transaction type, the signer   

## MODEL_INFO

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

##### Unique key:
 `vid:pid`

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

## X509_CERT

#### PROPOSE_X509_ROOT_CERT

#### APPROVE_X509_ROOT_CERT

#### ADD_X509_CERT

#### GET_PROPOSED_X509_ROOT_CERTS

#### GET_X509_ROOT_CERTS

#### GET_X509_CERT

####  

 
