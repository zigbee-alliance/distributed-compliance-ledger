# Transactions and Queries
See use case sequence diagrams for the examples of how transaction can be used.

## General
- Every writer to the Ledger must  
    - Have a private/public key pair.
    - Have an Account created on the ledger via `ACCOUNT` transaction (see [Use Case Txn Auth](use_cases/use_cases_txn_auth.puml)).
        - The Account stores the public part of the key
        - The Account has an associated role. The role is used for authorization policies.
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
    - Example
        ```json
        dclcli tx modelinfo add-model 1 1 "Device #1" "Device Description" "SKU12FS" "1.0" "2.0" true --from cosmos1ar04n6hxwk8ny54s2kzkpyqjcsnqm7jzv5y62y
        ```
- CLI (keys at the edge)
    - There are two CLIs are started in a CLI mode.
        - CLI 1: Stores private key. Does not have a connection to the network of nodes.
        - CLI 2: Is connected to the network of nodes. Doesn't has access to private key.
    - CLI 1: A private key is generated and stored off-server (in the user's private wallet).
    - CLI 2: Register account containing generated `Address` and `PubKey` on the ledger.
    - CLI 2: Build transaction using the account (`--from`) and `--generate-only` flag.
    - CLI 2: Fetch `account number` and `sequence`
    - CLI 1: Sign the transaction manually. `dclcli tx sign [path-to-txn-file] --from [address] --account-number [value] --sequence [value] --gas "auto" --offline`
    - CLI 2: Broadcast signed transaction using CLI (`broadcast command)
    - Example
        ```json
        CLI 2: dclcli tx modelinfo add-model 1 1 "Device #1" "Device Description" "SKU12FS" "1.0" "2.0" true --from cosmos1ar04n6hxwk8ny54s2kzkpyqjcsnqm7jzv5y62y --generate-only
        CLI 2: dclcli query auth accounts
        CLI 1: dclcli tx sign /home/artem/dc-ledger/txn.json --from cosmos1ar04n6hxwk8ny54s2kzkpyqjcsnqm7jzv5y62y --account-number 0 --sequence 24 --gas "auto" --offline --output-document txn.json
        CLI 2: dclcli tx broadcast /home/artem/dc-ledger/txn.json
        ```
- Non-trusted REST API (keys at the edge):
    - CLI is started in a server mode.
    - A private key is generated and stored off-server (in the user's private wallet).
    - Build transaction one of the following ways
        - The user builds the transaction manually (see `Extensions` section to find transactions format):.
        - The user does a `POST` to the server specifying the transaction parameters. The server builds the transaction.
        - The user uses corresponding CLI commands with specifying of `--generate-only` flag.
    - Sign the transaction manually using CLI or `tx/sign` endpoint.
    - The user does a `POST` of the signed request to the CLI-based server for broadcasting using `tx/broadcast`.     
    - Example
        ```json
        POST /modelinfo/models
        POST tx/sign
        POST tx/broadcast
        ```
- Trusted REST API (keys at the server):
    - CLI is started in a server mode.
    - A private key is generated and stored on the server (assuming it's within the user's domain and the user trusts it).
    - The user sends a POST to the server specifying the transaction parameters and account details (passphrase etc.) using `Authorization` header`.
    The server builds the request, signs it by the specified account's private key and broadcasts to the ledger
     in a way similar to the local CLI case.
    - See `REST API` section for every write request (transaction).
    - Example
        ```json
        POST /modelinfo/models with setting Authorization header 
        ```

## How to read from the Ledger
- Local CLI
    - CLI is started in a CLI mode.
    - No keys/account is needed as the ledger is public for reads
    - See `CLI` section for every read request.
- REST API
    - CLI is started in a server mode.
    - No keys/account is needed as the ledger is public for reads
    - See `REST API` section for every read request.   
    
##### Query types     
- Query single value:
    - By default, all read requests are performed for the last ledger state (height) and 
        then state proof verification is performed for result values.
        In order to verify state proof, app must have the next ledger state (height + 1) which will be created after some time.
        Because of this expectation queries may take significant time. 
    
        In order to avoid this delay CLI and REST exposes optional flag/parameter `prev-height`.
        When this flag is specified read requests will be performed for previous state (height-1). 
        The next state (height) is already present on the ledger so state proof verification is performed immediately.
    
        Note:     
        * In case simple reading use `prev-height` flag to get quick response.
        * In case of sequent add/read requests flag `prev-height` can be used. In case of failure for height-1 one more request for current height will be sent.
        * In case of sequent update/read requests flag `prev-height` must not be used because data before modification can be returned.
             
- Query list of values:
    - At the current moment, there is no state proof verification for list queries so there are no delays for those queries.
        

## KV Store
A summary of KV store and paths used:
- KV store name: `pki`
    - Proposed but not approved root certificates:
        - `1:<Certificate's Subject>:<Certificate's Subject Key ID>` : `<Certificate> + <List of approved trustee account IDs>`
    - Approved root and Non-root certificates:    
        - `2:<Certificate's Subject>:<Certificate's Subject Key ID>` : list[`<Certificate>`]
    - Certificate Chain:
        - `3:<Certificate's Subject>:<Certificate's Subject Key ID>` : `<a list of child certificates (subject/subjectKeyId pairs)>`
    - Proposed but not approved revocation of certificates:
        - `4:<Certificate's Subject>:<Certificate's Subject Key ID>` : `<List of approved trustee account IDs>`
    - CRL (Certificate Revocation List):
        - `5` : `CRL (Certificate Revocation List)`
    - Certificate uniqueness:
        - `6:<Certificate's Subject>:<Certificate's Subject Key ID>` : bool
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
       - `1:<certification_type>:<vid>:<pid>` : `<compliance info>`
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
    - Main index to store validators (active/jailed)
      - `1:<Validator Address>` : `<Validator>`
    - Helper index to track the last active validator set
      - `2:<Validator Address>` : `<Validator Last Power>`
    - Helper index to track that each validator owner has only one node
      - `5:<Account Address>` : `<Validator Address>`
    - Helper index to track validator signatures
      - `6:<Validator Address>` : `<Signing Info>`
    - Helper index to track whether validator signed a block
      - `7:<Validator Address:index>` : `<bool>`
   
## X509 PKI

**NOTE**: X.509 v3 certificates are only supported (all certificates MUST contain `Subject Key ID` field).
All PKI related methods are based on this restriction.

#### PROPOSE_ADD_X509_ROOT_CERT
**Status: Implemented**

Proposes a new self-signed root certificate.

If it's sent by a non-Trustee account, or more than 1 Trustee signature is required to add a root certificate, 
then the certificate
will be in a pending state until sufficient number of other Trustee's approvals is received.

The certificate is immutable. It can only be revoked by either the owner or a quorum of Trustees.

- Parameters:
  - `cert`: PEM-encoded certificate
- In State:
  - `pki` store  
  - `1:<Certificate's Subject>:<Certificate's Subject Key ID>` : `<Certificate> + <List of approved trustee account IDs>`
  - `2:<Certificate's Subject>:<Certificate's Subject Key ID>` : `List[<Certificate>]` (if just 1 Trustee is required)  
  - `6:<Certificate's Subject>:<Certificate's Subject Key ID>` : bool
- Who can send: 
    - Any role
- The current number of required approvals: 
    - 2
- CLI command: 
    -   `dclcli tx pki propose-add-x509-root-cert --certificate=<string-or-path> --from=<account>`
- REST API: 
    -   POST `/pki/certs/proposed/root`
- Validation:
    - provided certificate must be root: 
        - `Issuer` == `Subject` 
        - `Authority Key Identifier` == `Subject Key Identifier`
    - no existing `Proposed` certificate with the same `<Certificate's Subject>:<Certificate's Subject Key ID>` combination.
    - no existing certificate with the same `<Certificate's Issuer>:<Certificate's Serial Number>` combination.
    - if approved certificates with the same `<Certificate's Subject>:<Certificate's Subject Key ID>` combination already exists:
        - sender must match to the owner of the existing certificates.
    - the signature (self-signature) and expiration date are valid.

#### APPROVE_ADD_X509_ROOT_CERT
**Status: Implemented**

Approves the proposed root certificate.

The certificate is not active until sufficient number of Trustees approve it. 

- Parameters:
  - `subject`: string  - proposed certificates's `Subject`
  - `subject_key_id`: string  - proposed certificates's `Subject Key Id`
- In State:
  - `pki` store  
  - `1:<Certificate's Subject>:<Certificate's Subject Key ID>` : `<Certificate> + <List of approved trustee account IDs>`
  - `2:<Certificate's Subject>:<Certificate's Subject Key ID>` : `List[<Certificate>]`
- Who can send: 
    - Trustee
- The current number of required approvals: 
    - 2
- CLI command: 
    -   `dclcli tx pki approve-add-x509-root-cert --subject=<string> --subject-key-id=<hex string> --from=<account>`
- REST API: 
    -   PATCH `/pki/certs/proposed/root/<subject>/<subject_key_id>`
- Validation:
    - the proposed certificate hasn't been approved by the signer yet
        
#### ADD_X509_CERT
**Status: Implemented**

Adds an intermediate or leaf X509 certificate signed by a chain of certificates which must be
already present on the ledger.

The certificate is immutable. It can only be revoked by either the owner or a quorum of Trustees.

- Parameters:
  - `cert`: PEM-encoded certificate
- In State:
  - `pki` store  
  - `2:<Certificate's Subject>:<Certificate's Subject Key ID>` : `List[<Certificate>]`
  - `3:<Parent Certificate's Subject>:<Parent Certificate's Subject Key ID>` : `<a list of child certificates (subject/subjectKeyId pairs)>`
  - `6:<Certificate's Subject>:<Certificate's Subject Key ID>` : bool
- Who can send: 
    - Any role
- CLI command: 
    -   `dclcli tx pki add-x509-cert --certificate=<string-or-path> --from=<account>`
- REST API: 
    -   POST `/pki/certs`
- Validation:
    - provided certificate must not be root: 
        - `Issuer` != `Subject` 
        - `Authority Key Identifier` != `Subject Key Identifier`
    - no existing certificate with the same `<Certificate's Issuer>:<Certificate's Serial Number>` combination.
    - if certificates with the same `<Certificate's Subject>:<Certificate's Subject Key ID>` combination already exist:
        - sender must match to the owner of the existing certificates.
    - the signature (self-signature) and expiration date are valid.
    - parent certificate must be already stored on the ledger and a valid chain to some root certificate can be built. 

Note: Multiple certificates can refer to the same `<Certificate's Subject>:<Certificate's Subject Key ID>` combination.
    
#### REVOKE_X509_CERT
**Status: Implemented**

Revokes the given X509 certificate (either intermediate or leaf).
All the certificates in the chain signed by the revoked certificate will be revoked as well.

Only the owner (sender) can revoke the certificate.
Root certificates can not be revoked this way, use  `PROPOSE_X509_CERT_REVOC` and `APPROVE_X509_ROOT_CERT_REVOC` instead.  

- Parameters:
  - `subject`: string  - certificates's `Subject`
  - `subject_key_id`: string  - certificates's `Subject Key Id`
- In State:
  - `pki` store  
  - `2:<Certificate's Subject>:<Certificate's Subject Key ID>` : `List[<Certificate>]`  
  - `3:<Parent Certificate's Subject>:<Parent Certificate's Subject Key ID>` : `<a list of child certificates (subject/subjectKeyId pairs)>`
  - `5` : `CRL (Certificate Revocation List)`
- Who can send: 
    - Any role; owner
- CLI command: 
    -   `dclcli tx pki revoke-x509-cert --subject=<string> --subject-key-id=<hex string> --from=<account>`
- REST API: 
    -   DELETE `/pki/certs/<subject>/<subject_key_id>`

#### PROPOSE_REVOKE_X509_ROOT_CERT
**Status: Implemented**

Proposes revocation of the given X509 root certificate by a Trustee.

All the certificates in the chain signed by the revoked certificate will be revoked as well.

If more than 1 Trustee signature is required to revoke a root certificate, 
then the certificate will be in a pending state until sufficient number of other Trustee's approvals is received.

- Parameters:
  - `subject`: string  - certificates's `Subject`
  - `subject_key_id`: string  - certificates's `Subject Key Id`
- In State:
  - `pki` store  
  - `4:<Certificate's Subject>:<Certificate's Subject Key ID>` : `<List of approved trustee account IDs>`
- Who can send: 
    - Trustee
- CLI command: 
    -   `dclcli tx pki propose-revoke-x509-root-cert --subject=<string> --subject-key-id=<hex string> --from=<account>`
- REST API: 
    -   PUT `/pki/certs/proposed/revoked/root/<subject>/<subject_key_id>`
    


#### APPROVE_REVOKE_X509_ROOT_CERT
**Status: Implemented**

Approves the revocation of the given X509 root certificate by a Trustee.
All the certificates in the chain signed by the revoked certificate will be revoked as well.

The revocation is not applied until sufficient number of Trustees approve it. 

- Parameters:
  - `subject`: string  - certificates's `Subject`
  - `subject_key_id`: string  - certificates's `Subject Key Id`
- In State:
  - `pki` store  
  - `4:<Certificate's Subject>:<Certificate's Subject Key ID>` : `<List of approved trustee account IDs>`
  - `2:<Certificate's Subject>:<Certificate's Subject Key ID>` : `List[<Certificate>]`  
  - `3:<Parent Certificate's Subject>:<Parent Certificate's Subject Key ID>` : `<a list of child certificates (subject/subjectKeyId pairs)>`
  - `5` : `CRL (Certificate Revocation List)`
- Who can send: 
    - Trustee
- CLI command: 
    -   `dclcli tx pki approve-revoke-x509-root-cert --subject=<string> --subject-key-id=<hex string> --from=<account>`
- REST API: 
    -   PATCH `/pki/certs/proposed/revoked/root/<subject>/<subject_key_id>`
        
#### GET_ALL_PROPOSED_X509_ROOT_CERTS
**Status: Implemented**

Gets all proposed but not approved root certificates.

- Parameters:
  - `skip`: optional(int)  - number records to skip (`0` by default)
  - `take`: optional(int)  - number records to take (all records are returned by default)
- CLI command: 
    -   `dclcli query pki all-proposed-x509-root-certs ... `
- REST API: 
    -   GET `/pki/certs/proposed/root`
- Result:
```json
{
  "result": {
    "total": string,
    "items": [
      {
        "pem_cert": string, //pem encoded certificate
        "subject": string,
        "subject_key_id": string,
        "serial_number": string,
        "approvals": optional([string]),
        "owner": string
      }
    ]
  },
  "height": string
}
```

#### GET_PROPOSED_X509_ROOT_CERT
**Status: Implemented**

Gets a proposed but not approved root certificate with the given subject and subject key id attributes.

- Parameters:
  - `subject`: string  - certificates's `Subject`
  - `subject_key_id`: string  - certificates's `Subject Key Id`
  - `prev-height`: optional(bool) - query data from previous height to avoid delay linked to state proof verification
- CLI command: 
    -   `dclcli query pki proposed-x509-root-cert --subject=<string> --subject-key-id=<hex string> ... `
- REST API: 
    -   GET `/pki/certs/proposed/root/<subject>/<subject_key_id>`
```json
{
  "result": {
    "pem_cert": string, //pem encoded certificate
    "subject": string,
    "subject_key_id": string,
    "serial_number": string,
    "approvals": optional([string]),
    "owner": string
  },
  "height": string
}
```

#### GET_ALL_X509_ROOT_CERTS
**Status: Implemented**

Gets all approved root certificates. Revoked certificates are not returned. 
Use `GET_ALL_REVOKED_X509_CERTS_ROOT` to get a list of all revoked root certificates. 

- Parameters:
  - `skip`: optional(int)  - number records to skip (`0` by default)
  - `take`: optional(int)  - number records to take (all records are returned by default)
- CLI command: 
    -   `dclcli query pki all-x509-root-certs .... `
- REST API: 
    -   GET `/pki/certs/root`
```json
{
  "result": {
    "total": string,
    "items": [
      {
        "pem_cert": string, //pem encoded certificate
        "subject": string,
        "subject_key_id": string,
        "serial_number": string,
        "issuer": string,
        "authority_key_id": string,
        "root_subject": string,
        "root_subject_key_id": string,      
        "is_root": boolean, 
        "owner": string,
      }
    ]
  },
  "height": string
}
```

#### GET_X509_CERT
**Status: Implemented**

Gets a certificate (either root, intermediate or leaf) by the given subject and subject key id attributes.
Revoked certificates are not returned. 
Use `GET_ALL_REVOKED_X509_CERTS` to get a list of all revoked certificates. 

- Parameters:
  - `subject`: string  - certificates's `Subject`
  - `subject_key_id`: string  - certificates's `Subject Key Id`
  - `prev-height`: optional(bool) - query data from previous height to avoid delay linked to state proof verification
- CLI command: 
    -   `dclcli query pki x509-cert --subject=<string> --subject-key-id=<hex string> ... `
- REST API: 
    -   GET `/pki/certs/<subject>/<subject_key_id>`
```json
{
  "result": {
    "items": [
      {
        "pem_cert": string, //pem encoded certificate
        "subject": string,
        "subject_key_id": string,
        "serial_number": string,
        "issuer": string, // omitted for root certificates
        "authority_key_id": string, // omitted for root certificates
        "root_subject": string, // omitted for root certificates
        "root_subject_key_id": string, // omitted for root certificates
        "is_root": boolean, 
        "owner": string,
      }
    ]
  },
  "height": string
}
```

#### GET_X509_CERT_CHAIN
**Status: Implemented**

Gets the complete chain for a certificate with the given combination of subject and subject-key-id.
Revoked certificates are not returned. 
Use `GET_ALL_REVOKED_X509_CERTS` to get a list of all revoked certificates. 

- Parameters:
  - `subject`: string  - certificates's `Subject`
  - `subject_key_id`: string  - certificates's `Subject Key Id`
  - `prev-height`: optional(bool) - query data from previous height to avoid delay linked to state proof verification
- CLI command: 
    -   `dclcli query pki x509-cert-chain --subject=<string> --subject-key-id=<hex string> ... `
- REST API: 
    -   GET `/pki/certs/chain/<subject>/<subject_key_id>`
```json
{
  "result": {
    "items": [
      {
        "pem_cert": string, //pem encoded certificate
        "subject": string,
        "subject_key_id": string,
        "serial_number": string,
        "issuer": string, // omitted for root certificate (last in the chain)
        "authority_key_id": string, // omitted for root certificate (last in the chain)
        "root_subject": string, // omitted for root certificate (last in the chain)
        "root_subject_key_id": string, // omitted for root certificate (last in the chain)
        "is_root": boolean, 
        "owner": string,
      }
    ]
  },
  "height": string
}
```

#### GET_ALL_X509_CERTS
**Status: Implemented**

Gets all certificates (root, intermediate and leaf).

Revoked certificates are not returned. 
Use `GET_ALL_REVOKED_X509_CERTS` to get a list of all revoked certificates. 

Can optionally be filtered by the root certificate's subject or subject key id so that 
only the certificate chains started with the given root certificate are returned.   

`GET_ALL_X509_CERTS_SINCE` can be used to incrementally update the list stored locally. 

- Parameters:
  - `skip`: optional(int)  - number records to skip (`0` by default)
  - `take`: optional(int)  - number records to take (all records are returned by default)
  - `root_subject`: string (optional) - root certificates's `Subject`
  - `root_subject_key_id`: string (optional) - root certificates's `Subject Key Id`
- CLI command: 
    -   `dclcli query pki all-x509-certs .... `
- REST API: 
    -   GET `/pki/certs`
    -   GET `/pki/certs?root_subject=<>`
    -   GET `/pki/certs?root_subject_key_id=<>`
    -   GET `/pki/certs?root_subject=<>;root_subject_key_id=<>`
```json
{
  "result": {
    "total": string,
    "items": [
      {
        "pem_cert": string, //pem encoded certificate
        "subject": string,
        "subject_key_id": string,
        "serial_number": string,
        "issuer": string, // omitted for root certificates
        "authority_key_id": string, // omitted for root certificates
        "root_subject": string, // omitted for root certificates
        "root_subject_key_id": string, // omitted for root certificates
        "is_root": boolean, 
        "owner": string,
      }
    ]
  },
  "height": string
}
```

#### GET_ALL_SUBJECT_X509_CERTS
**Status: Implemented**

Gets all certificates (root, intermediate and leaf) associated with a subject.

Revoked certificates are not returned. 
Use `GET_ALL_REVOKED_X509_CERTS` to get a list of all revoked certificates. 

Can optionally be filtered by the root certificate's subject or subject key id so that 
only the certificate chains started with the given root certificate are returned. 

- Parameters:
  - `skip`: optional(int)  - number records to skip (`0` by default)
  - `take`: optional(int)  - number records to take (all records are returned by default)
  - `root_subject`: string (optional) - root certificates's `Subject`
  - `root_subject_key_id`: string (optional) - root certificates's `Subject Key Id`
- CLI command: 
    -   `dclcli query pki all-subject-x509-certs --subject=<string> .... `
- REST API: 
    -   GET `/pki/certs/<subject>`
    -   GET `/pki/certs/<subject>?root_subject=<>`
    -   GET `/pki/certs/<subject>?root_subject_key_id=<>`
    -   GET `/pki/certs/<subject>?root_subject=<>;root_subject_key_id=<>`
```json
{
  "result": {
    "total": string,
    "items": [
      {
        "pem_cert": string, //pem encoded certificate
        "subject": string,
        "subject_key_id": string,
        "serial_number": string,
        "issuer": string, // omitted for root certificates
        "authority_key_id": string, // omitted for root certificates
        "root_subject": string, // omitted for root certificates
        "root_subject_key_id": string, // omitted for root certificates
        "is_root": boolean, 
        "owner": string,
      }
    ]
  },
  "height": string
}
```

#### GET_ALL_PROPOSED_X509_ROOT_CERTS_TO_REVOKE
**Status: Implemented**

Gets all proposed but not approved root certificates to be revoked.

- Parameters:
  - `skip`: optional(int)  - number records to skip (`0` by default)
  - `take`: optional(int)  - number records to take (all records are returned by default)
- CLI command: 
    -   `dclcli query pki all-proposed-x509-root-certs-to-revoke .... `
- REST API: 
    -   GET `/pki/certs/proposed/revoked/root`

#### GET_PROPOSED_X509_ROOT_CERT_TO_REVOKE
**Status: Implemented**

Gets a proposed but not approved root certificate to be revoked.

- Parameters:
  - `subject`: string  - certificates's `Subject`
  - `subject_key_id`: string  - certificates's `Subject Key Id`
- CLI command: 
    -   `dclcli query pki proposed-x509-root-cert-to-revoke --subject=<string> --subject-key-id=<hex string>`
- REST API: 
    -   GET `/pki/certs/proposed/revoked/root/<subject>/<subject_key_id>`

#### GET_ALL_REVOKED_X509_CERTS
**Status: Implemented**

Gets all revoked certificates (both root and non-root).
   
- Parameters:
  - `skip`: optional(int)  - number records to skip (`0` by default)
  - `take`: optional(int)  - number records to take (all records are returned by default)
- CLI command: 
    -   `dclcli query pki all-revoked-x509-certs .... `
- REST API: 
    -   GET `/pki/certs/revoked`

#### GET_ALL_REVOKED_X509_ROOT_CERTS
**Status: Implemented**

Gets all revoked root certificates.
   
- Parameters:
  - `skip`: optional(int)  - number records to skip (`0` by default)
  - `take`: optional(int)  - number records to take (all records are returned by default)
- CLI command: 
    -   `dclcli query pki all-revoked-x509-root-certs .... `
- REST API: 
    -   GET `/pki/certs/revoked/root`    

#### GET_ALL_X509_CERTS_SINCE
**Status: Not Implemented**

Gets all certificates (root, intermediate and leaf) which has been added since 
the given ledger's `height`.

Can optionally be filtered by the root certificate's subject and subject key id so that 
only the certificate chains started with the given root certificate are returned.   

- Parameters:
  - `since`: integer - the last ledger's height the user has locally.
  - `root_subject`: string (optional) - root certificates's `Subject`
  - `root_subject_key_id`: string (optional) - root certificates's `Subject Key Id` 
- CLI command: 
    -   `dclcli query pki all-x509-certs-delta --since=<integer>.... `
- REST API: 
    -   GET `/pki/certs?since=<>`
    -   GET `/pki/certs?since=<>;root_subject=<>;root_subject_key_id={}`
    
## MODEL INFO

#### ADD_MODEL_INFO
**Status: Implemented**

Adds a new Model Info identified by a unique combination of `vid` (vendor ID) and `pid` (product ID).

Only some of Model Info fields can be edited (see `EDIT_MODEL_INFO`). If other fields need to be edited - 
a new model info with a new `vid` or `pid` can be created.

If one of `OTA_URl`, `OTA_checksum` and `OTA_checksum_type` fields is set, then the other two must also be set.

- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - name: `string` -  model name
  - description: `string` -  model description (string or path to file containing data)
  - sku: `string` -  stock keeping unit
  - software-version: `uint32` -  Software Version of model (uint32)
  - software-version-string: `string` - Software Version String of model
  - hardware-version: `uint32` -  version of model hardware
  - hardware-version-string: `string` - Hardware Version String of model
  - cd-version-number: `uint32` -  CD Version Number of the Certification
  - from: `string` - Name or address of private key with which to sign
  - cid: `optional(uint16)` - model category ID (positive non-zero)
  - revoked: `optional(bool)` - boolean flag to revoke the model
  - ota-url: `optional(string)` - the URL of the OTA
  - ota-checksum: `optional(string)` - the checksum of the OTA 
  - ota-checksum-type: `optional(string)` - the type of the OTA checksum 
  - ota-blob: `optional(string)` - metadata about OTA 
  - commisioning-custom-flow: `optional(uint8)` - A value of 1 indicates that user interaction with the device (pressing a button, for example) is required before commissioning can take place. When CommissioningCustomflow is set to a value of 2, the commissioner SHOULD attempt to obtain a URL which MAY be used to provide an end-user with the necessary details for how to configure the product for initial commissioning
  - commisioning-custom-flow-url: `optional(string)` - commisioning-custom-flow-url SHALL identify a vendor specific commissioning URL for the device model when the commisioning-custom-flow field is set to '2'
  - commissioning-mode-initial-steps-hint: `optional(uint32)` - commissioning-mode-initial-steps-hint SHALL identify a hint for the steps that can be used to put into commissioning mode a device that has not yet been commissioned. This field is a bitmap with values defined in the Pairing Hint Table. For example, a value of 1 (bit 0 is set) indicates that a device that has not yet been commissioned will enter Commissioning Mode upon a power cycle.
  - commissioning-mode-initial-steps-instruction: `optional(string)` - commissioning-mode-initial-steps-instruction SHALL contain text which relates to specific values of CommissioningModeInitialStepsHint. Certain values of CommissioningModeInitialStepsHint, as defined in the Pairing Hint Table, indicate a Pairing Instruction (PI) dependency, and for these values the commissioning-mode-initial-steps-instruction SHALL be set
  - commissioning-mode-secondary-steps-hint: `optional(uint32)` - commissioning-mode-secondary-steps-hint SHALL identify a hint for steps that can be used to put into commissioning mode a device that has already been commissioned. This field is a bitmap with values defined in the Pairing Hint Table. For example, a value of 4 (bit 2 is set) indicates that a device that has already been commissioned will require the user to visit a current CHIP Administrator to put the device into commissioning mode.
  - commissioning-mode-secondary-step-instruction: `optional(string)` - commissioning-mode-secondary-step-instruction SHALL contain text which relates to specific values of commissioning-mode-secondary-steps-hint. Certain values of commissioning-mode-secondary-steps-hint, as defined in the Pairing Hint Table, indicate a Pairing Instruction (PI) dependency, and for these values the commissioning-mode-secondary-step-instruction SHALL be set
  - release_notes_url: `optional(string)` - URL that contains product specific web page that contains release notes for the device model.
  - user-manual-url: `optional(string)` - URL that contains product specific web page that contains user manual for the device model.
  - support-url: `optional(string)` - URL that contains product specific web page that contains support details for the device model.
  - product-url: `optional(string)` - URL that contains product specific web page that contains details for the device model.  
  - chip-blob: `optional(string)` - chip-blob SHALL identify CHIP specific configurations
  - vendor-blob: `optional(string)` - field for vendors to provide any additional metadata about the device model using a string, blob, or URL.  

- In State:
  - `modelinfo` store  
  - `1:<vid>:<pid>` : `<model info>`
  - `2:<vid>` : `<list of pids + metadata>`
- Who can send: 
    - Vendor
- CLI command: 
    -   `dclcli tx modelinfo add-model --vid=<uint16> --pid=<uint16> --name=<string> --description=<string or path> --sku=<string> 
--software-version=<uint32> --software-version-string=<string> --hardware-version=<uint32> --hardware-version-string=<string> --cd-version-number=<uint16> 
--from=<account> .... `
- REST API: 
    -   POST `/modelinfo/models`

#### EDIT_MODEL_INFO
**Status: Implemented**

Edits an existing Model Info identified by a unique combination of `vid` (vendor ID) and `pid` (product ID)
by the owner.

Only the fields listed below (besides `vid` and `pid`) can be edited. If other fields need to be edited - 
a new model info with a new `vid` or `pid` can be created.

All non-edited fields remain the same.

`OTA_URL` can be edited only if  `OTA_checksum` and `OTA_checksum_type` are already set.

- Parameters:
    - `vid`: 16 bits int
    - `pid`: 16 bits int
    - `cid`: 16 bits int (optional)
    - `OTA_URL`: string (optional)
    - `description`: string (optional)
    - `tis_or_trp_testing_completed`: bool
    - `custom`: string (optional)
- In State:
  - `modelinfo` store  
  - `1:<vid>:<pid>` : `<model info>`
  - `2:<vid>` : `<list of pids + metadata>`
- Who can send: 
    - Vendor; owner
- CLI command: 
    -   `dclcli tx modelinfo update-model --vid=<uint16> --pid=<uint16> --tis-or-trp-testing-completed=<bool> --from=<account> .... `
- REST API: 
    -   PUT `/modelinfo/models/vid/pid`


#### GET_ALL_MODEL_INFO
**Status: Implemented**

Gets all Model Infos for all vendors.

- Parameters:
  - `skip`: optional(int)  - number records to skip (`0` by default)
  - `take`: optional(int)  - number records to take (all records are returned by default)
- CLI command: 
    -   `dclcli query modelinfo all-models ...`
- REST API: 
    -   GET `/modelinfo/models`
- Result
```json
{
  "height": string,
  "result": {
    "total": string,
    "items": [
      {
        "vid": 16 bits int,
        "pid": 16 bits int,
        "name": string,
        "owner": string,
        "sku": string
      }
    ]
  }
}
```

#### GET_VENDOR_MODEL_INFO
**Status: Implemented**

Gets all Model Info by the given Vendor (`vid`).

- Parameters:
    - `vid`: 16 bits int
- CLI command: 
    -   `dclcli query modelinfo vendor-models --vid=<uint16>`
- REST API: 
    -   GET `/modelinfo/models/vid`
- Result
```json
{
  "height": string,
  "result": {
    "vid": 16 bits int,
    "products": [
      {
        "pid": 16 bits int,
        "name": string,
        "owner": string,
        "sku": string
      }
    ]
  }
}
```

#### GET_MODEL_INFO
**Status: Implemented**

Gets a Model Info with the given `vid` (vendor ID) and `pid` (product ID).

- Parameters:
    - `vid`: 16 bits int
    - `pid`: 16 bits int
    - `prev-height`: optional(bool) - query data from previous height to avoid delay linked to state proof verification
- CLI command: 
    -   `dclcli query modelinfo model --vid=<uint16> --pid=<uint16> .... `
- REST API: 
    -   GET `/modelinfo/models/vid/pid`
- Result
```json
{
  "height": string,
  "result": {
    "vid": 16 bits int,
    "pid": 16 bits int,
    "cid": (optional) 16 bits int,
    "name": string,
    "owner": string,
    "description": string,
    "sku": string,
    "firmware_version": string,
    "hardware_version": string,
    "OTA_URL": string (optional),
    "OTA_checksum" string (optional),
    "OTA_checksum_type": string (optional),
    "custom": (optional) string,
    "tis_or_trp_testing_completed": bool
  }
}
```

#### GET_VENDORS    
**Status: Implemented**

Get a list of all Vendors (`vid`s). 

- Parameters:
  - `skip`: optional(int)  - number records to skip (`0` by default)
  - `take`: optional(int)  - number records to take (all records are returned by default)
- CLI command: 
    -   `dclcli query modelinfo vendors .... `
- REST API: 
    -   GET `/modelinfo/vendors`
- Result
```json
{
  "height": string,
  "result": {
    "total": string,
    "items": [
      {
        "vid": 16 bits int
      }
    ]
  }
}
```

## TEST_DEVICE_COMPLIANCE

#### ADD_TEST_RESULT
**Status: Implemented**

Submits result of a compliance testing for the given device (`vid` and `pid`).
The test result can be a blob of data or a reference (URL) to an external storage.

Multiple test results (potentially from different test houses) can be added for the same device type. 

The test result is immutable and can not be deleted or removed after submitting. 
Another test result can be submitted instead.

- Parameters:
    - `vid`: 16 bits int
    - `pid`: 16 bits int
    - `test_result`: string
    - `test_date`: rfc3339 encoded date
- In State:
  - `compliancetest` store  
  - `1:<vid>:<pid>` : `<list of test results>`
- Who can send: 
    - TestHouse
- CLI command: 
    -   `dclcli tx compliancetest add-test-result --vid=<uint16> --pid=<uint16> --test-result=<string> --test-date=<rfc3339 encoded date> --from=<account>`
- REST API: 
    -   POST `/compliancetest/testresults`

#### GET_TEST_RESULT
**Status: Implemented**

Gets a test result for the given `vid` (vendor ID) and `pid` (product ID).

- Parameters:
    - `vid`: 16 bits int
    - `pid`: 16 bits int
    - `prev-height`: optional(bool) - query data from previous height to avoid delay linked to state proof verification
- CLI command: 
    -   `dclcli query compliancetest test-result --vid=<uint16> --pid=<uint16> .... `
- REST API: 
    -   GET `/compliancetest/testresults/vid/pid`
- Result:
```json
{
  "height": string,
  "result": {
    "vid": 16 bits int,
    "pid": 16 bits int,
    "results": [
      {
        "test_result": string,
        "test_date": datetime,
        "owner": string
      }
    ]
  }
}
```

## CERTIFY_DEVICE_COMPLIANCE

#### CERTIFY_MODEL
**Status: Implemented**

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
    - `certification_date`: rfc3339 encoded date - date of certification
    - `certification_type`: string  - `zb` is the default and the only supported value now
    - `reason` (optional): string  - optional comment describing the reason of the certification
- In State:
  - `compliance` store  
  - `1:<certification_type>:<vid>:<pid>` : `<compliance info>`
  - `2:<vid>` : `<compliance pids>`  
  - `3:<vid>` : `<revoked pids>`  
- Who can send: 
    - ZBCertificationCenter
- CLI command: 
    -   `dclcli tx compliance certify-model --vid=<uint16> --pid=<uint16> --certification-type=<zb> --certification-date=<rfc3339 encoded date> --from=<account> .... `
- REST API: 
    -   PUT `/compliance/certified/vid/pid/certification_type`
    
#### REVOKE_MODEL_CERTIFICATION
**Status: Implemented**

Revoke compliance of the Model to the ZB standard.

The corresponding Model Info and test results are not required to be on the ledger 
to be used in cases where revocation only is tracked on the ledger.

It can be used in cases where every compliance result 
is written on the ledger (`CERTIFY_MODEL` was called), or
 cases where only revocation list is stored on the ledger.
 
- Parameters:
    - `vid`: 16 bits int
    - `pid`: 16 bits int
    - `revocation_date`: rfc3339 encoded date - date of revocation
    - `certification_type` string - `zb` is the default and the only supported value now
    - `reason` (optional): string - optional comment describing the reason of the revocation
- In State:
  - `compliance` store  
  - `1:<certification_type>:<vid>:<pid>` : `<compliance info>`
  - `2:<vid>` : `<compliance pids>`  
  - `3:<vid>` : `<revocation pids>`  
- Who can send: 
    - ZBCertificationCenter
- CLI command: 
    -   `dclcli tx compliance revoke-model --vid=<uint16> --pid=<uint16> --certification-type=<zb> --revocation-date=<rfc3339 encoded date> --from=<account> .... `
- REST API: 
    -   PUT `/compliance/revoked/vid/pid/certification_type`    
    
#### GET_CERTIFIED_MODEL
**Status: Implemented**

Gets a boolean if the given Model (identified by the `vid`, `pid` and `certification_type`) is compliant to ZB standards. 

This is the aggregation of compliance and
revocation information for every vid/pid. It should be used in cases where compliance 
is tracked on the ledger.

Note: This function returns `false` in two cases:
- compliance information not found in the store.
- compliance information is found but it is in `revoked` state. 

You can use `GET_COMPLICE_INFO` method to get the whole compliance information.
 
- Parameters:
    - `vid`: 16 bits int
    - `pid`: 16 bits int
    - `certification_type`: string - `zb` is the default and the only supported value now
    - `prev-height`: optional(bool) - query data from previous height to avoid delay linked to state proof verification
- CLI command: 
    -   `dclcli query compliance certified-model --vid=<uint16> --pid=<uint16> --certification-type=<zb> .... `
- REST API: 
    -   GET `/compliance/certified/vid/pid/certification_type`
- Result
```json
{
  "result": {
    "value": bool
  },
  "height": string
}
```

#### GET_REVOKED_MODEL
**Status: Implemented**

Gets a boolean if the given Model (identified by the `vid`, `pid` and `certification_type`) is revoked. 

It contains information about revocation only, so it should be used in cases
 where only revocation is tracked on the ledger.

Note: This function returns `false` in two cases:
- compliance information not found in the store.
- compliance information is found but it is in `certified` state. 
 
You can use `GET_COMPLICE_INFO` method to get the whole compliance information.

- Parameters:
    - `vid`: 16 bits int
    - `pid`: 16 bits int
    - `certification_type`: string - `zb` is the default and the only supported value now
    - `prev-height`: optional(bool) - query data from previous height to avoid delay linked to state proof verification
- CLI command: 
    -   `dclcli query compliance revoked-model --vid=<uint16> --pid=<uint16> --certification-type=<zb> .... `
- REST API: 
    -   GET `/compliance/revoked/vid/pid/certification_type`
- Result:
```json
{
  "result": {
    "value": bool
  },
  "height": string
}
```

#### GET_COMPLIANCE_INFO
**Status: Implemented**

Gets compliance information associated with the Model (identified by the `vid` and `pid` and `certification_type`).

It can be used instead of GET_CERTIFIED_MODEL / GET_REVOKED_MODEL methods 
 to get the whole compliance information without additional state check.
This function responds with `NotFoundError` (404 code) if compliance information (identified by the `vid` and `pid`) not found in store.
 
- Parameters:
    - `vid`: 16 bits int
    - `pid`: 16 bits int
    - `certification_type`: string - `zb` is the default and the only supported value now
    - `prev-height`: optional(bool) - query data from previous height to avoid delay linked to state proof verification
- CLI command: 
    -   `dclcli query compliance compliance-info --vid=<uint16> --pid=<uint16> --certification-type=<zb> .... `
- REST API: 
    -   GET `/compliance/vid/pid/certification_type`
- Result:
```json
{
  "result": {
    "vid": 16 bits int,
    "pid": 16 bits int,
    "state": string, // certified or revoked
    "date": rfc3339 encoded date,
    "certification_type": string,
    "reason": optional(string),
    "owner": string
  },
  "height": string
}
```
- Result in case the sate of the model was changed before:
```json
{
  "result": {
    "vid": 16 bits int,
    "pid": 16 bits int,
    "state": string, // certified or revoked
    "date": rfc3339 encoded date,
    "certification_type": string,
    "reason": optional(string),
    "owner": string,
    "history": [
      {
        "state": string, // certified or revoked
        "date": rfc3339 encoded date,
        "reason": optional(string)
      }
    ]
  },
  "height": string
}
```

#### GET_ALL_REVOKED_MODELS
**Status: Implemented**

Gets all revoked Models (`pid`s) for all vendors (`vid`s).

It contains information about revocation only, so it should be used in cases
 where only revocation is tracked on the ledger.
 
`GET_ALL_REVOKED_MODELS_SINCE` can be used to incrementally update the list stored locally. 

- Parameters:
  - `skip`: optional(int)  - number records to skip (`0` by default)
  - `take`: optional(int)  - number records to take (all records are returned by default)
- CLI command: 
    -   `dclcli query compliance all-revoked-models ... `
- REST API: 
    -   GET `/compliance/revoked`
        - optional query parameter `certification_type` can be passed to filter by certification type.
 - Result
 ```json
{
  "result": {
    "total": string,
    "items": [
      {
        "vid": 16 bits int,
        "pid": 16 bits int,
        "certification_type": string,
      }
    ]
  },
  "height": string
}
 ```

#### GET_ALL_CERTIFIED_MODELS
**Status: Implemented**

Gets all compliant Models (`pid`s) for all the vendors (`vid`s).

This is the aggregation of compliance and
revocation information for every vid/pid. It should be used in cases where compliance is tracked on ledger.
 
`GET_ALL_CERTIFIED_MODELS_SINCE` can be used to incrementally update the list stored locally. 
 
- Parameters:
  - `skip`: optional(int)  - number records to skip (`0` by default)
  - `take`: optional(int)  - number records to take (all records are returned by default)
- CLI command: 
    -   `dclcli query compliance all-certified-models `
- REST API: 
    -   GET `/compliance/certified`
        - optional query parameter `certification_type` can be passed to filter by certification type.
 - Result
 ```json
{
  "result": {
    "total": string,
    "items": [
      {
        "vid": 16 bits int,
        "pid": 16 bits int,
        "certification_type": string,
      }
    ]
  },
  "height": string
}
 ```
    
#### GET_ALL_COMPLIANCE_INFO_RECORDS
**Status: Implemented**

Gets all stored compliance information records.

`GET_ALL_COMPLIANCE_INFO_RECORDS_SINCE` can be used to incrementally update the list stored locally.
 
- Parameters:
  - `skip`: optional(int)  - number records to skip (`0` by default)
  - `take`: optional(int)  - number records to take (all records are returned by default)
- CLI command: 
    -   `dclcli query compliance all-compliance-info-records`
- REST API: 
    -   GET `/compliance`
        - optional query parameter `certification_type` can be passed
 - Result
 ```json
{
  "result": {
    "total": string,
    "items": [
      {
        "vid": 16 bits int,
        "pid": 16 bits int,
        "state": string, // certified or revoked
        "date": rfc3339 encoded date,
        "certification_type": string,
        "reason": optional(string),
        "owner": string
        "history":  array // (as for `GET_COMPLIANCE_INFO`) if not empty
      }
    ]
  },
  "height": string
}
 ```

#### GET_VENDOR_CERTIFIED_MODELS
**Status: Not Implemented**

Gets all the Models (`pid`s) issued by the given Vendor (`vid`) which are complaint to ZB standards.

This is the aggregation of compliance and
revocation information for every vid/pid. It should be used in cases where compliance is tracked on ledger.

- Parameters:
    - `vid`: 16 bits int
- CLI command: 
    -   `dclcli query compliance certified-vendor-models --vid=<uint16> .... `
- REST API: 
    -   GET `/compliance/certified/vid`
    
#### GET_VENDOR_REVOKED_MODELS
**Status: Not Implemented**

Gets all the Models (`pid`s) which certification is revoked for the given Vendor (`vid`).

It contains information about revocation only, so  it should be used in cases
 where only revocation is tracked on the ledger.


- Parameters:
    - `vid`: 16 bits int
- CLI command: 
    -   `dclcli query compliance revoked-vendor-models --vid=<uint16>  .... `
- REST API: 
    -   GET `/compliance/revoked/vid`

#### GET_ALL_CERTIFIED_MODELS_SINCE
**Status: Not Implemented**

Gets a delta of all compliant Models (`pid`s) for every vendor (`vid`s) which has been added or revoked since 
the given ledger's `height`.

This is the aggregation of compliance and
revocation information for every vid/pid. It should be used in cases where compliance is tracked on ledger.
 
- Parameters: 
  - `since`: integer - the last ledger's height the user has locally.
- CLI command: 
    -   `dclcli query compliance all-certified-models-delta `
- REST API: 
    -   GET `/compliance/certified?since=<>`
    
#### GET_ALL_REVOKED_MODELS_SINCE
**Status: Not Implemented**

Gets a delta of all revoked Models (`pid`s) for every vendor (`vid`s) which has been added or revoked since 
the given ledger's `height`.

This is the aggregation of compliance and
revocation information for every vid/pid. It should be used in cases where compliance is tracked on ledger.
 
- Parameters: 
  - `since`: integer - the last ledger's height the user has locally.
- CLI command: 
    -   `dclcli query compliance all-revoked-models-delta `
- REST API: 
    -   GET `/compliance/revoked?since=<>`
    
#### GET_ALL_COMPLIANCE_INFO_RECORDS_SINCE
**Status: Not Implemented**

Gets a delta of all compliance info records which has been added or revoked since the given ledger's `height`.

- Parameters: 
  - `since`: integer - the last ledger's height the user has locally.
- CLI command: 
    -   `dclcli query compliance all-compliance-info-records-delta `
- REST API: 
    -   GET `/compliance?since=<>`
    
## AUTH

#### PROPOSE_ADD_ACCOUNT
**Status: Implemented**

Proposes a new Account with the given address, public key and role.

If more than 1 Trustee signature is required to add the account, the account
will be in a pending state until sufficient number of approvals is received.

- Parameters:
    - `address`: string // account address; bech32 encoded
    - `pub_key`: string // account public key; bech32 encoded
    - `roles`: array<string> // the list of roles to assign to account 
- In State:
  - `auth` store  
  - `1:<address>` : `<account info> + <list of approvers>`
  - `2:<address>` : `<account info>` (if just 1 Trustee is required)  
- Who can send: 
    - Trustee
- CLI command: 
    -   `dclcli tx auth propose-add-account --address=<account address> --pubkey=<account pubkey> --roles=<role1,role2,...> --from=<trustee name>`
- REST API: 
    -   POST `/auth/accounts/proposed`
    
#### APPROVE_ADD_ACCOUNT
**Status: Implemented**

Approves the proposed account.    

The account is not active until sufficient number of Trustees approve it. 

- Parameters:
    - `address`: string // account address; bech32 encoded
- In State:
  - `auth` store  
  - `1:<address>` : `<account info> + <list of approvers>`  
  - `2:<address>` : `<account info>`
- Who can send: 
    - Trustee
- CLI command: 
    -   `dclcli tx auth approve-add-account --address=<account address> --from=<trustee name>`
- REST API: 
    -   PATCH `/auth/accounts/proposed/<address>`
    
    
  
#### PROPOSE_REVOKE_ACCOUNT
**Status: Implemented**

Proposes revocation of the Account with the given address.

If more than 1 Trustee signature is required to revoke the account, the revocation
will be in a pending state until sufficient number of approvals is received.

- Parameters:
    - `address`: string // account address; bech32 encoded
- In State:
  - `auth` store  
  - `3:<address>` : `<account info> + <list of approvers>`
  - `2:<address>` : `<account info>` (if just 1 Trustee is required)  
- Who can send: 
    - Trustee
- CLI command: 
    -   `dclcli tx auth propose-revoke-account --address=<account address> --from=<trustee name>`
- REST API: 
    -   POST `/auth/accounts/proposed/revoked`
    
#### APPROVE_REVOKE_ACCOUNT
**Status: Implemented**

Approves the proposed revocation of the account.    

The account is not revoked until sufficient number of Trustees approve it. 

- Parameters:
    - `address`: string // account address; bech32 encoded
- In State:
  - `auth` store  
  - `3:<address>` : `<account info> + <list of approvers>`  
  - `2:<address>` : `<account info>`
- Who can send: 
    - Trustee
- CLI command: 
    -   `dclcli tx auth approve-revoke-account --address=<account address> --from=<trustee name>`
- REST API: 
    -   PATCH `/auth/accounts/proposed/revoked/<address>`
    
#### GET_ALL_PROPOSED_ACCOUNTS
**Status: Implemented**

Gets all proposed but not approved accounts.

- Parameters: No
- CLI command: 
    -   `dclcli query auth all-proposed-accounts .... `
- REST API: 
    -   GET `/auth/accounts/proposed`
    
#### GET_ALL_ACCOUNTS
**Status: Implemented**

Gets all accounts. Revoked accounts are not returned.

- Parameters: No
- CLI command: 
    -   `dclcli query auth all-accounts .... `
- REST API: 
    -   GET `/auth/accounts`           

#### GET_ACCOUNT
**Status: Implemented**

Gets an accounts by the address. Revoked accounts are not returned.

- Parameters:
    - `address`
- CLI command: 
    -   `dclcli query auth account .... `
- REST API: 
    -   GET `/auth/accounts/<address>`         
    
#### GET_ALL_PROPOSED_ACCOUNTS_TO_REVOKE
**Status: Implemented**

Gets all proposed but not approved accounts to be revoked.

- Parameters: No
- CLI command: 
    -   `dclcli query auth all-proposed-accounts-to-revoke`
- REST API: 
    -   GET `/auth/accounts/proposed/revoked`
    

#### ROTATE_KEY
**Status: Not Implemented**

Rotate's the Account's public key by the owner.

- Parameters:
    - `pub_key`
    - `pub_key_type`
- In State:
  - `auth` store  
  - `1:<address>` : `<account info>`
- Who can send: 
    - Any role; owner
- CLI command: 
    -   `dclcli tx auth rotate-key .... `
- REST API: 
    -   PUT `/auth/accounts/<address>`
    
## VALIDATOR_NODE                      

#### ADD_VALIDATOR_NODE
**Status: Implemented**

Adds a new Validator node.

- Parameters:
    - `validator_address`: string // the tendermint validator address; bech32 encoded
    - `validator_pubkey`: string // the tendermint validator public key; bech32 encoded
    - `description`: json
        - `name`: string // validator name
        - `identity`: string (optional) // identity signature (ex. UPort or Keybase)
        - `website`: string (optional) // website link
        - `details`: string (optional) // details
- In State:
  - `validator` store  
  - `1:<Validator Address>` : `<Validator>` - main index to store validators (there are two state of validator: active/jailed)
  - `2:<Validator Address>` : `<Validator Last Power>` - helper index to track the last active validator set
  - `5:<Account Address>` : `<Validator Address>` - helper index to track that each validator owner has only one node
  - `6:<Validator Address>` : `<Signing Info>` - helper index to track validator signatures
  - `7:<Validator Address>:<index>` : `<bool>` - helper index to track validator signatures over blocks window
- Who can send: 
    - NodeAdmin
- CLI command: 
    -   `dclcli tx validator add-node --validator-address=<validator address> --validator-pubkey=<validator pubkey> --name=<node name> --from=<name> .... `
- REST API: 
    -   POST `/validators`

#### GET_ALL_VALIDATORS
**Status: Implemented**

Gets the list of all validator nodes from the store.

Note: All stored validator nodes (`active` and `jailed`) will be returned by default.
In order to get an active validator set use `state` query parameter or specific command [validator set](#validator-set).

- Parameters:
  - `skip`: optional(int)  - number records to skip (`0` by default)
  - `take`: optional(int)  - number records to take (all records are returned by default)
  - `state`: string (optional) - state of the validator (active/jailed)
- CLI command: 
    -   `dclcli query validator all-nodes .... `
- REST API: 
    -   GET `/validators`   
- Result:
    ```json
    {
      "height": string,
      "result": {
        "total": string,
        "items": [
          {
            "description": {
              "name": string // validator name
              "identity": optional(string) // identity signature (ex. UPort or Keybase)
              "website": optional(string) // website link
              "details": optional(string) // additional details
            },
            "validator_address": string, // the tendermint validator address
            "validator_pubkey": string, // the tendermint validator public key
            "power": string, // validator consensus power
            "jailed": bool, // has the validator been removed from validator set because of cheating
            "jailed_reason": optional(string), // the reason of validator jailing
            "owner": string // the account address of validator owner (original sender of transaction)
          },
          ...
        ]
      }
    }
    ```
     
#### GET_VALIDATOR
**Status: Implemented**

Gets a validator node.

- Parameters:
    - `validator_address`: string // the tendermint validator address; bech32 encoded
- CLI command: 
    -   `dclcli query validator node --validator-address=<validator address>.... `
- REST API: 
    -   GET `/validators/<validator_address>`   
- Result:
    ```json
    {
      "height": string,
      "result": {
        "description": {
          "name": string // validator name
          "identity": optional(string) // identity signature (ex. UPort or Keybase)
          "website": optional(string) // website link
          "details": optional(string) // additional details
        },
        "validator_address": string, // the tendermint validator address
        "validator_pubkey": string, // the tendermint validator public key
        "power": string, // validator consensus power
        "jailed": bool, // has the validator been removed from validator set because of cheating
        "jailed_reason": optional(string), // the reason of validator jailing
        "owner": string // the account address of validator owner (original sender of transaction)
      }
    }
    ```
    
#### UPDATE_VALIDATOR_NODE
**Status: Not Implemented**

Updates the Validator node by the owner. Only `description` can be changed. 
`validator_address` is used to reference the node, but can not be changed. 

- Parameters: 
    - `validator_address`: string // the tendermint validator address; bech32 encoded
    - `description`: json
        - `name`: string // validator name
        - `identity`: string (optional) // identity signature (ex. UPort or Keybase)
        - `website`: string (optional) // website link
        - `details`: string (optional) // details
- In State: 
  - `validator` store  
  - `1:<Validator Address>` : `<Validator>` - main index to store validators (there are two state of validator: active/jailed)
- Who can send: 
    - NodeAdmin; owner
- CLI command: 
    -   `dclcli tx validator update-node --validator-address=<validator address> --from=<owner>.... `
- REST API: 
    -   PUT `/validators/<validator_address>`    

#### REMOVE_VALIDATOR_NODE
**Status: Not Implemented**

Deletes the Validator node (removes from the validator set) by the owner.

- Parameters:
    - `validator_address`: string // the tendermint validator address; bech32 encoded
- In State: 
  - `validator` store  
  - `1:<Validator Address>` : `<Validator>` - main index to store validators (there are two state of validator: active/jailed)
  - `2:<Validator Address>` : `<Validator Last Power>` - helper index to track the last active validator set
  - `5:<Account Address>` : `<Validator Address>` - helper index to track that each validator owner has only one node
  - `6:<Validator Address>` : `<Signing Info>` - helper index to track validator signatures
  - `7:<Validator Address>:<index>` : `<bool>` - helper index to track validator signatures over blocks window
- Who can send: 
    - NodeAdmin; owner
- CLI command: 
    -   `dclcli tx validator remove-node --validator-address=<validator address> --from=<owner>.... `
- REST API: 
    -   DELETE `/validators/<validator_address>`

#### PROPOSE_REMOVE_VALIDATOR_NODE
**Status: Not Implemented**

Proposes removing the Validator node from the validator set by a Trustee. 

If more than 1 Trustee signature is required to remove a node, the removal
will be in a pending state until sufficient number of approvals is received.

- Parameters:
    - `validator_address`: string // the tendermint validator address; bech32 encoded
- In State: Same as in Tendermint/Cosmos-sdk
- Who can send: 
    - Trustee
- CLI command: 
    -   `dclcli tx validator propose-remove-node .... `
- REST API: 
    -   POST `/validators/proposed/removed`

#### APPROVE_REMOVE_VALIDATOR_NODE
**Status: Not Implemented**

Approves removing of the Validator node by a Trustee. 

The account is not removed until sufficient number of Trustees approve it. 

- Parameters:
    - `validator_address`: string // the tendermint validator address; bech32 encoded
- In State: Same as in Tendermint/Cosmos-sdk
- Who can send: 
    - Trustee
- CLI command: 
    -   `dclcli tx validator approve-remove-node .... `
- REST API: 
    -   PATCH `/validators/proposed/removed/<validator_address>`
                      
#### UNJAIL_VALIDATOR_NODE
**Status: Not Implemented**

Approves unjail of the Validator node from jailed state and returning to the active validator state. 

If more than 1 Trustee approval is required to unjail a node, the node still
will be in a jailed state until sufficient number of approvals is received.

If 1 Trustee approval is required to unjail a nod or sufficient number of approvals is received,
the node will be unjailed and returned to the active validator set.

- Parameters:
    - `validator_address`: string // the tendermint validator address; bech32 encoded
- In State:
  - `validator` store  
  - `1:<Validator Address>` : `<Validator + List of Approvals>`
- Who can send: 
    - Trustee
- CLI command: 
    -   `dclcli tx validator unjail-node --validator-address=<validator address> --from=<trustee>.... `
- REST API: 
    -   PATCH `/validators/<validator_address>`
            
   
#### GET_ALL_PROPOSED_VALIDATORS_TO_REMOVE
**Status: Not Implemented**

Gets all proposed but not approved validator nodes to be removed.

- Parameters: No
- CLI command: 
    -   `dclcli query validator all-proposed-nodes-to-remove .... `
- REST API: 
    -   GET `/validators/proposed/removed`   
    
    
## Extensions    

#### Sign
Sign transaction by the given key.

- Parameters:
    - `txn` - transaction to sign.
    - `from` -  name or address of private key to use to sign.
    - `account-number` - (optional) the account number of the signing account.
    - `sequence` - (optional) the sequence number of the signing account.
    - `chain-id` - (optional) chain id.
- CLI command: 
    -   `dclcli tx sign [path-to-txn-file] --from [address]`
    - Transaction:
        ```
        {
            "type":"cosmos-sdk/StdTx",
            "value":{
                "msg":[
                    {msg to sign}
                ],
                "fee":{
                    "amount":[],
                    "gas":string
                },
                "signatures":null,
                "memo":""
            }
        }
        ```
- REST API: 
    - POST `/tx/sign`  
    - Request
    ```
        base_req: {
            "from": string,
            "chain_id": string,
            "account_number": optional(string),
            "sequence": optional(string),
        },
        txn: {
            "type":"cosmos-sdk/StdTx",
            "value":{
                "msg":[
                    {msg to sign}
                ],
                "fee":{
                    "amount":[],
                    "gas":"200000"
                },
                "signatures":null,
                "memo":""
            }
        }
    ```
Note: if `account_number` and `sequence`  are not specified they will be fetched from the ledger automatically.  
   
#### Broadcast
Broadcast transaction to the ledger.

- Parameters:
    - `txn` - transaction to broadcast
- CLI command: 
    -   `dclcli tx broadcast [path-to-txn-file]`
- REST API: 
    - POST `/tx/broadcast`  
- Transaction:
    ```
    {
        "type":"cosmos-sdk/StdTx",
        "value":{
            "msg":[
                {msg to sign}
            ],
            "fee":{
                "amount":[],
                "gas":"200000"
            },
            "signatures": [
                {
                    "pub_key": {
                        "type": "tendermint/PubKeySecp256k1",
                        "value": "AiXgamO0AZKu38IZE/j9Tt54mQq4yza/5zilm6rlHwrb"
                    },
                    "signature": "SS+52lsXPWtQuEUFJv6Tl8U+vfdyatWK3piYjNlZWb1pg2tYMZlH4IEIIYs+Cfg6/F3lPEOb/SJeuCsh/Zl2/w=="
                }
            ],
            "memo":""
        }
    }
    ```
  
#### Status
Query status of a node.

- Parameters:
    - `node`: optional(string) - node physical address to query (by default queries the node specified in CLI config file or else "tcp://localhost:26657")
- CLI command: 
    -   `dclcli status [--node=<node ip>]`
        ```
- REST API: 
    - GET `/status?node=<node ip>`  
- Result:
    ```json
    {
      "node_info": {
        "protocol_version": {
          "p2p": string
          "block": string,
          "app": string
        },
        "id": string,
        "listen_addr": string,
        "network": string,
        "version": string,
        "channels": string,
        "moniker": string,
        "other": {
          "tx_index": string,
          "rpc_address": string
        }
      },
      "sync_info": {
        "latest_block_hash": string,
        "latest_app_hash": string,
        "latest_block_height": string,
        "latest_block_time": string,
        "catching_up": bool
      },
      "validator_info": {
        "address": string,
        "pub_key": {
          "type": string,
          "value": string
        },
        "voting_power": string
      }
    }
    ```

#### Validator set
Get the list of tendermint validators participating in the consensus at given height.

- Parameters:
    - `height`: optional(uint) - height to query (the latest by default)
- CLI command: 
    -   `dclcli tendermint-validator-set [height]`
        ```
- REST API: 
    - GET `/validator-set?height=<height>`
- Result:
    ```json
    {
      "block_height": string,
      "validators": [
        {
          "address": string,
          "pub_key": string,
          "proposer_priority": string,
          "voting_power": string
        },
        ...
      ]
    }
    ```
