# NOC Root Certificate Transactions Design

## User Stories

### 1. Add NOC Root Certificate
A Vendor with DCL write privilege can submit a transaction to add a NOC root certificate associated with their Vendor ID.

### 2. Revoke NOC Root Certificate
A Vendor with DCL write privilege can submit a transaction to revoke a NOC root certificate associated with their Vendor ID.

### 3. Remove NOC Root Certificate
A Vendor with DCL write privilege can submit a transaction to remove a NOC root certificate associated with their Vendor ID. So that the Vendor can remove certificates that were added by mistake.

## Certificate Schema

To distinguesh NOC root certificates from others, an `isNOC` boolean field will be added to the [certificates](https://github.com/zigbee-alliance/distributed-compliance-ledger/blob/master/proto/pki/certificate.proto) schema 

## Transactions

### 1. ADD_NOC_X509_ROOT_CERTIFICATE
This transaction adds a NOC root certificate owned by the Vendor.

- Who can send: Vendor account
   - `vid` field in the transaction (`VendorID`) must be equal to the Vendor account's VID
- Validation:
  - The provided certificate must be a root certificate:
    - `Issuer` == `Subject`
    - `Authority Key Identifier` == `Subject Key Identifier`
  - No existing certificate with the same `<Certificate's Issuer>:<Certificate's Serial Number>` combination.
  - If certificates with the same `<Certificate's Subject>:<Certificate's Subject Key ID>` combination already exist:
    - The sender's VID must match the vid field of the existing certificates.
  - No existing certificate with the same `<Certificate's Subject>:<Certificate's Subject Key ID>` combination already published by another vendor.
  - The signature (self-signature) and expiration date must be valid.
- Parameters:
  - cert: `string` - The NOC Root Certificate, encoded in X.509v3 PEM format. Can be a PEM string or a file path.
  - vid: `uint16` - Vendor ID (positive non-zero)
- In State:
  - `pki/ApprovedCertificates/value/<Subject>/<SubjectKeyID>`
  - `pki/ApprovedCertificatesBySubject/value/<Subject>`
  - `pki/NOCRootCertificates/value/<VID>`
- CLI Command:
  - `dcld tx pki add-noc-x509-root-cert --certificate=<string-or-path> --from=<account>`

### 2. REVOKE_NOC_X509_ROOT_CERTIFICATE
This transaction revokes a NOC root certificate owned by the Vendor.
Revoked NOC root certificates can be re-added using the `ADD_NOC_X509_ROOT_CERTIFICATE` transaction.

- Who can send: Vendor account
  - Vid field associated with the corresponding NOC root certificate on the ledger must be equal to the Vendor account's VID.
- Validation:
  - A NOC Root Certificate with the provided `subject` and `subject_key_id` must exist in the ledger.
- Parameters:
  - subject: `string` - Base64 encoded subject DER sequence bytes of the certificate.
  - subject_key_id: `string` - Certificate's `Subject Key Id` in hex string format, e.g., `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`.
  - serial_number: `optional(string)` - Certificate's serial number.
  - info: `optional(string)` - Information/notes for the revocation
  - time: `optional(int64)` - Revocation time (number of nanoseconds elapsed since January 1, 1970 UTC). CLI uses the current time for that field.
- In State:
  - `pki/RevokedCertificates/value/<subject>/<subject_key_id>`
  - `pki/RevokedNOCRootCertificates/value/<subject>/<subject_key_id>`
- CLI Command:
  - `dcld tx pki revoke-noc-x509-root-cert --subject=<base64 string> --subject-key-id=<hex string> --serial-number=<string> --from=<account>`

### 3. REMOVE_NOC_X509_ROOT_CERTIFICATE
This transaction completely removes a NOC root certificate owned by the Vendor. 
Removed NOC root certificates can be re-added using the `ADD_NOC_X509_ROOT_CERTIFICATE` transaction.

**Node:** Revoked certificates that match the specified parameters will also be removed.

- Who can send: Vendor account
  - Vid field associated with the corresponding NOC root certificate on the ledger must be equal to the Vendor account's VID.
- Validation:
  - A NOC root certificate with the provided `subject` and `subject_key_id` must exist in the ledger.
- Parameters:
  - subject: `string` - Base64 encoded subject DER sequence bytes of the certificate.
  - subject_key_id: `string` - Certificate's `Subject Key Id` in hex string format, e.g., `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`.
  - serial_number: `optional(string)` - Certificate's serial number.
- CLI Command:
  - `dcld tx pki remove-noc-x509-root-cert --subject=<base64 string> --subject-key-id=<hex string> --serial-number=<string> --from=<account>`

## Query

To retrieve NOC certificates by Subject and Subject Key Identifier, use the [GET_X509_CERT](https://github.com/zigbee-alliance/distributed-compliance-ledger/blob/master/docs/transactions.md#get_x509_cert) or [GET_ALL_SUBJECT_X509_CERTS](https://github.com/zigbee-alliance/distributed-compliance-ledger/blob/master/docs/transactions.md#get_all_subject_x509_certs:) query.

To retrieve a revoked NOC certificate by Subject and Subject Key Identifier, use the [GET_REVOKED_CERT](https://github.com/zigbee-alliance/distributed-compliance-ledger/blob/master/docs/transactions.md#get_revoked_cert)

### GET_NOC_X509_ROOT_CERTS_BY_VID

Retrieve NOC root certificates associated with a specific VID. 

- Who can send: Any account
- Parameters:
  - vid: `uint16` - Vendor ID (positive non-zero)
- CLI Command:
  - `dcld query pki get_noc_x509_root_certs --vid=<uint16>`
- REST API:
  - GET `/dcl/pki/noc-root-certificates/{vid}`

### GET_ALL_NOC_X509_ROOT_CERTS

Retrieve a list of all of NOC root certificates

- Who can send: Any account
- Parameters:
  - Common pagination parameters
- CLI Command:
  - `dcld query pki get_all_noc_x509_root_certs`
- REST API:
  - GET `/dcl/pki/noc-root-certificates`

### GET_ALL_REVOKED_NOC_X509_ROOT_CERTS

Gets all revoked NOC root certificates.

- Who can send: Any account
- Parameters:
  - Common pagination parameters
- CLI command:
  - `dcld query pki all-revoked-noc-x509-root-certs`
- REST API:
  - GET `/dcl/pki/revoked-noc-root-certificates`
