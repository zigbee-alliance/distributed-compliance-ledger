# NOC Root Certificate Transactions Design

## User Stories

### 1. Add NOC Root Certificate
A Vendor with DCL write privilege can submit a transaction to add a NOC Root certificate associated with their Vendor ID.

### 2. Revoke NOC Root Certificate
A Vendor with DCL write privilege can submit a transaction to revoke a NOC Root certificate associated with their Vendor ID. 

## Transactions

### 1. ADD_NOC_X509_ROOT_CERTIFICATE
This transaction adds a NOC Root Certificate owned by the Vendor.

- Who can send: Vendor account
  - VID-scoped NOC Root Certificate: The `vid` field in the certificate's subject must be equal to the Vendor account's VID.
- Validation:
  - The provided certificate must be a root certificate:
    - `Issuer` == `Subject`
    - `Authority Key Identifier` == `Subject Key Identifier`
  - No existing certificate with the same `<Certificate's Issuer>:<Certificate's Serial Number>` combination.
  - If certificates with the same `<Certificate's Subject>:<Certificate's Subject Key ID>` combination already exist:
    - The sender must match the owner of the existing certificates.
  - The signature (self-signature) and expiration date must be valid.
- Parameters:
  - cert: `string` - The NOC Root Certificate, encoded in X.509v3 PEM format. Can be a PEM string or a file path.
- State Changes:
  - `pki/ApprovedCertificates/value/<Subject>/<SubjectKeyID>`
  - `pki/NOCRootCertificates/value/<VID>`
- CLI Command:
  - `dcld tx pki add-noc-x509-root-cert --certificate=<string-or-path> --from=<account>`

### 2. REVOKE_NOC_X509_ROOT_CERTIFICATE
This transaction revokes a NOC Root Certificate owned by the Vendor.

- Who can send: Vendor account
  - VID-scoped NOC Root Certificate: The `vid` field in the certificate's subject must be equal to the Vendor account's VID.
- Validation:
  - A NOC Root Certificate with the provided `subject` and `subject_key_id` must exist in the ledger.
- Parameters:
  - subject: `string` - Base64 encoded subject DER sequence bytes of the certificate.
  - subject_key_id: `string` - Certificate's `Subject Key Id` in hex string format, e.g., `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`.
- CLI Command:
  - `dcld tx pki revoke-noc-x509-root-cert --subject=<base64 string> --subject-key-id=<hex string> --from=<account>`

## Query

### GET_NOC_X509_ROOT_CERTS

Retrieve NOC Root Certificates associated with a specific VID. 
Use `GET_ALL_X509_ROOT_CERTS` to list all NOC certificates by Subject and Subject Key Identifier.

- Who can send: Any account
- Parameters:
  - vid: `uint16` - Vendor ID (positive non-zero)
- CLI Command:
  - `dcld query pki get_noc_x509_root_certs`
- REST API:
  - GET `/dcl/pki/NocRootCertificates/{vid}`

## Question
- Should the vendor add a revocation distribution point for NOC certificates?