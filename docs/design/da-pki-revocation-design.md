# DA PKI Revocation

## User Stories

1. Revocation of a VID-scoped PAI

   A Vendor with DCL write privilege can submit a transaction to publish the location of the CRL distribution point for a PAA associated with the Vendor ID. This information can be updated or deleted by the Vendor with the same Vendor ID.

2. Revocation of a DAC 

   A Vendor with DCL write privilege can submit a transaction to publish the location of the CRL distribution point for a PAI associated with the Vendor ID. This information can be updated or deleted by the Vendor with the same Vendor ID.

3. Revocation of a non-VID-scoped PAI

   A Vendor Admin with DCL write privilege can submit a transaction to publish the location of the CRL distribution point for a non-VID scoped PAA. This information can be updated or deleted by any Vendor Admin account.


## Revocation Distribution Point Schema

| Name                 | Type   | Constraint          | Mutable | Conformance |
|:---------------------|:-------|:--------------------|:--------|-------------|
| VendorID             | uint16 | all                 | No      | M           |
| ProductID            | uint16 | all                 | No      | O (desc)    |
| IsPAA                | bool   | all                 | No      | M           |
| Label                | string | max 64              | No      | M           |
| CRLSignerCertificate | string | max 2048            | Yes     | M           |
| IssuerSubjectKeyID   | string | max 64              | No      | M           |
| DataUrl              | string | max 256             | Yes     | M           |
| DataFileSize         | uint64 | desc                | Yes     | O           |
| DataDigest           | string | max 128             | Yes     | O (desc)    |
| DataDigestType       | uint32 | all                 | Yes     | O (desc)    |
| RevocationType       | uint32 | all                 | No      | M           |


## A need for Proof-of-possession of `CRLSignerCertificate` key
No additional proof-of-possession is required because 
- The transaction is already signed by the Vendor account key, so only a trusted Vendor can send it.
- DCL checks that `CRLSignerCertificate` is either a valid PAI (signed by a trusted PAA from DCL) or a PAA present on the ledger.
- It can be complicated to create an interactive PoP of `CRLSignerCertificate` for some PAIs.
- Indirect CRLs are out of scope for now

## Transactions

### 1. ADD_PKI_REVOCATION_DISTRIBUTION_POINT
Publishes a PKI Revocation distribution endpoint owned by the Vendor. 

If `crlSignerCertificate` is a PAA (root certificate), then it must be present on DCL.

If `crlSignerCertificate` is a PAI (intermediate certificate), then it must be chained back to a valid PAA (root certificate) present on DCL.
In this case `crlSignerCertificate` is not required to be present on DCL, and will not be added to DCL as a result of this transaction.
If PAI needs to be added to DCL, it should be done via `ADD_X509_CERT` transaction.

Publishing the revocation distribution endpoint doesn't automatically remove PAI (Intermediate certificates)
and DACs (leaf certificates) added to DCL if they are revoked in the CRL identified by this distribution point.


- Who can send: Vendor account
    - `vid` field in the transaction (`VendorID`) must be equal to the Vendor account's VID
    - `vid` field in the `CRLSignerCertificate` must be equal to the Vendor account's VID
- Validation of parameters:
    - See [Validation](#validation-logic) section for details. 
- Parameters:
    - vid: `uint16` -  Vendor ID (positive non-zero). Must be the same as Vendor account's VID and `vid` field in the VID-scoped `CRLSignerCertificate`.
    - pid: `optional(uint16)` -  Product ID (positive non-zero). Must be empty if `IsPAA` is true. Must be equal to a `pid` field in `CRLSignerCertificate`.
    - isPAA: `bool` -  True if the revocation information distribution point relates to a PAA
    - label: `string` -  A label to disambiguate multiple revocation information partitions of a particular issuer.
    - crlSignerCertificate: `string` - The issuer certificate whose revocation information is provided in the distribution point entry, encoded in X.509v3 PEM format. The corresponding CLI parameter can contain either a PEM string or a path to a file containing the data.
    - issuerSubjectKeyID: `string` - Uniquely identifies the PAA or PAI for which this revocation distribution point is provided. Must consist of even number of uppercase hexadecimal characters ([0-9A-F]), with no whitespace and no non-hexadecimal characters., e.g: `5A880E6C3653D07FB08971A3F473790930E62BDB`.
    - dataUrl: `string` -  The URL where to obtain the information in the format indicated by the RevocationType field. Must start with either `http` or `https`.
    - dataFileSize: `optional(uint64)` -  Total size in bytes of the file found at the DataUrl. Must be omitted if RevocationType is 1.
    - dataDigest: `optional(string)` -  Digest of the entire contents of the associated file downloaded from the DataUrl. Must be omitted if RevocationType is 1. Must be provided if and only if the `DataFileSize` field is present.
    - dataDigestType: `optional(uint32)` - The type of digest used in the DataDigest field from the list of [1, 7, 8, 10, 11, 12] (IANA Named Information Hash Algorithm Registry). Must be provided if and only if the `DataDigest` field is present.
    - revocationType: `uint32` - The type of file found at the DataUrl for this entry. Supported types: 1 - RFC5280 Certificate Revocation List (CRL).
- In State:
    - `pki/RevocationDistributionPoint/value/<IssuerSubjectKeyID>` -> list of Revocation Distribution Points
    - `pki/RevocationDistributionPoint/value/<IssuerSubjectKeyID>/<vid>/<label>`-> Revocation Distribution Point
- CLI command:
    - `dcld tx pki add-revocation-point --vid=<uint16> --pid=<uint16> --issuer-subject-key-id=<string> --is-paa=<bool> --label=<string>
      --certificate=<string-or-path> --data-url=<string> --revocation-type=1 --from=<account>`


### 2. UPDATE_PKI_REVOCATION_DISTRIBUTION_POINT
Updates an existing PKI Revocation distribution endpoint owned by the sender.

- Who can send: Vendor account
    - `vid` field in the transaction (`VendorID`) must be equal to the Vendor account's VID
    - `vid` field in the `CRLSignerCertificate` must be equal to the Vendor account's VID
- Parameters:
    - vid: `uint16` -  Vendor ID (positive non-zero). Must be the same as Vendor account's VID and `vid` field in the VID-scoped `CRLSignerCertificate`.
    - label: `string` -  A label to disambiguate multiple revocation information partitions of a particular issuer.
    - issuerSubjectKeyID: `string` - Uniquely identifies the PAA or PAI for which this revocation distribution point is provided. Must consist of even number of uppercase hexadecimal characters ([0-9A-F]), with no whitespace and no non-hexadecimal characters., e.g: `5A880E6C3653D07FB08971A3F473790930E62BDB`.
    - crlSignerCertificate: `optional(string)` - The issuer certificate whose revocation information is provided in the distribution point entry, encoded in X.509v3 PEM format. The corresponding CLI parameter can contain either a PEM string or a path to a file containing the data.
    - dataUrl: `optional(string)` -  The URL where to obtain the information in the format indicated by the RevocationType field. Must start with either `http` or `https`.
    - dataFileSize: `optional(uint64)` -  Total size in bytes of the file found at the DataUrl. Must be omitted if RevocationType is 1.
    - dataDigest: `optional(string)` -  Digest of the entire contents of the associated file downloaded from the DataUrl. Must be omitted if RevocationType is 1. Must be provided if and only if the `DataFileSize` field is present.
    - dataDigestType: `optional(uint32)` - The type of digest used in the DataDigest field from the list of [1, 7, 8, 10, 11, 12] (IANA Named Information Hash Algorithm Registry). Must be provided if and only if the `DataDigest` field is present.
- In State:
    - `pki/RevocationDistributionPoint/value/<IssuerSubjectKeyID>` -> list of Revocation Distribution Points
    - `pki/RevocationDistributionPoint/value/<IssuerSubjectKeyID>/<vid>/<label>` -> Revocation Distribution Point
- CLI command:
    - `dcld tx pki update-revocation-point --vid=<uint16> --issuer-subject-key-id=<string> --label=<string>
      --data-url=<string> --certificate=<string-or-path> --from=<account>`

### 3. DELETE_DELETE_PKI_REVOCATION_DISTRIBUTION_POINT
Deletes a PKI Revocation distribution endpoint owned by the sender.

- Who can send:
    - Vendor account for VID-scoped `crlSignerCertificate`
        - `vid` field in the transaction (`VendorID`) must be equal to the Vendor account's VID
        - `vid` field in the corresponding `CRLSignerCertificate` (for vendor-scoped PAAs and PAIs) must be equal to the Vendor account's VID
    - Vendor Admin account for non-VID scoped `crlSignerCertificate`
        - `vid` field must be absent in the corresponding`CRLSignerCertificate`
- Parameters:
    - vid: `uint16` -  Vendor ID (positive non-zero). Must be the same as Vendor account's VID and `vid` field in the VID-scoped `CRLSignerCertificate`.
    - label: `string` -  A label to disambiguate multiple revocation information partitions of a particular issuer.
  - issuerSubjectKeyID: `string` - Uniquely identifies the PAA or PAI for which this revocation distribution point is provided. Must consist of even number of uppercase hexadecimal characters ([0-9A-F]), with no whitespace and no non-hexadecimal characters., e.g: `5A880E6C3653D07FB08971A3F473790930E62BDB`.
- In State:
    - `pki/RevocationDistributionPoint/value/<IssuerSubjectKeyID>` -> list of Revocation Distribution Points
    - `pki/RevocationDistributionPoint/value/<IssuerSubjectKeyID>/<vid>/<label>` -> Revocation Distribution Point
- CLI command:
    - `dcld tx pki delete-revocation-point --vid=<uint16> --issuer-subject-key-id=<string> --label=<string> --from=<account>`


## Query

### 1. GET_PKI_REVOCATION_DISTRIBUTION_POINT
Gets a revocation distribution point identified by (VendorID, Label, IssuerSubjectKeyID) unique combination.
Use `GET_ALL_PKI_REVOCATION_DISTRIBUTION_POINT` to get a list of all revocation distribution points.

- Parameters:
  - vid: `uint16` -  Vendor ID (positive non-zero)
  - label: `string` -  A label to disambiguate multiple revocation information partitions of a particular issuer.
  - issuerSubjectKeyID: `string` - Uniquely identifies the PAA or PAI for which this revocation distribution point is provided. Must consist of even number of uppercase hexadecimal characters ([0-9A-F]), with no whitespace and no non-hexadecimal characters., e.g: `5A880E6C3653D07FB08971A3F473790930E62BDB`.
- CLI command:
  - `dcld query pki revocation-point --vid=<uint16> --label=<string> --issuer-subject-key-id=<string>`
- REST API:
  - GET `/dcl/pki/revocation-points/{issuerSubjectKeyID}/{vid}/{label}`
  
### 2. GET_PKI_REVOCATION_DISTRIBUTION_POINTS_BY_SUBJECT_KEY_ID
Gets a list of revocation distribution point identified by IssuerSubjectKeyID.

- Parameters:
    - issuerSubjectKeyID: `string` - Uniquely identifies the PAA or PAI for which this revocation distribution point is provided. Must consist of even number of uppercase hexadecimal characters ([0-9A-F]), with no whitespace and no non-hexadecimal characters., e.g: `5A880E6C3653D07FB08971A3F473790930E62BDB`.
- CLI command:
  - `dcld query pki revocation-points --issuer-subject-key-id=<string>`
- REST API:
  - GET `/dcl/pki/revocation-points/{issuerSubjectKeyID}`

### 3. GET_ALL_PKI_REVOCATION_DISTRIBUTION_POINT
Gets a list of all revocation distribution points.

- Parameters:
  - Common pagination parameters
- CLI command:
  - `dcld query pki all-revocation-points`
- REST API:
  - GET `/dcl/pki/revocation-points`
  
## Stored in State
  - `pki/RevocationDistributionPoint/value/<IssuerSubjectKeyID>` -> list of Revocation Distribution Points
  - `pki/RevocationDistributionPoint/value/<IssuerSubjectKeyID>/<vid>/<label>` -> Revocation Distribution Point

## Validation Logic

### Static validation (before transaction is propagated and proposed to a block)
- All types and constraints (see [Schema](#revocation-distribution-point-schema))
- `DataDigestType` is one of the following values: [1, 7, 8, 10, 11, 12].
- `RevocationType` is one of the following values: [1]
- `ProductID` must be empty if `IsPAA` is true. 
- `DataUrl` starts with either `http` or `https` and follow the syntax of [RFC3986](https://datatracker.ietf.org/doc/html/rfc3986).
- `DataDigest` is present if and only if the `DataFileSize` field is present.
- `DataDigestType` is provided if and only if the `DataDigest` field is present. 
- `IssuerSubjectKeyID` must consist of even number of uppercase hexadecimal characters ([0-9A-F]), with no whitespace and no non-hexadecimal characters., e.g: `5A880E6C3653D07FB08971A3F473790930E62BDB`.
- If `RevocationType` is 1 (RFC5280 CRL), then `DataFileSize`, `DataDigest`, `DataDigestType` must be empty. This is applied to both Add and Update transactions.
- Check that `ProductID` field is provided if and only if `IsPAA` is false and `CRLSignerCertificate` has a PID in its subject.  
  If `ProductID` is provided, it must be equal to the PID in `CRLSignerCertificate`'s subject.
  - both OID and text versions of PID in a certificate can be supported (either `Mpid` or `1.3.6.1.4.1.37244.2.2`, see `x509.ToSubjectAsText` method and "2.2. Encoding of Vendor ID and Product ID in subject and issuer fields" section in spec)
- If `IsPAA` is true, then
  - check that the `CRLSignerCertificate` is a PAA (root certificate, self-signed).
  - If `CRLSignerCertificate` encodes a vid in its subject, then it must be equal to `VendorID` field.
    - both OID and text versions of VID in a certificate can be supported (either `Mvid` or `1.3.6.1.4.1.37244.2.1`, see `x509.ToSubjectAsText` method and "2.2. Encoding of Vendor ID and Product ID in subject and issuer fields" section in spec)
- If `IsPAA` is false, then
  - check that `CRLSignerCertificate` is a non-root certificate (not self-signed).
  - `CRLSignerCertificate` must encode a vid in its subject equal to `VendorID` field.
    - both OID and text versions of VID in a certificate can be supported (either `Mvid` or `1.3.6.1.4.1.37244.2.1`, see `x509.ToSubjectAsText` method and "2.2. Encoding of Vendor ID and Product ID in subject and issuer fields" section in spec)

### Dynamic validation (when adding to a block)

#### ADD_PKI_REVOCATION_DISTRIBUTION_POINT
- The sender must be a Vendor account
- If `crlSignerCertificate` is a PAA (root certificate, self-signed):
   - If `CRLSignerCertificate` encodes a vid in its subject, then `VendorID` field must be equal to the Vendor account's VID. 
     Otherwise, Unsupported Error.
       - both OID and text versions of VID in a certificate can be supported (either `Mvid` or `1.3.6.1.4.1.37244.2.1`, see `x509.ToSubjectAsText` method and "2.2. Encoding of Vendor ID and Product ID in subject and issuer fields" section in spec) 
   - Query a certificate by `crlSignerCertificate` Subject and Subject Key ID. If it's not found - error.
   - Check that pem value of the found certificate is equal to `crlSignerCertificate` value.
- If `crlSignerCertificate` is a PAI (intermediate certificate, not self-signed):
   - Check that `VendorID` field must be equal to the Vendor account's VID.
   - Check that `crlSignerCertificate` is chained back to certificates present on DCL (`x509.verifyCertificate` method):
       - Query for a PAA where `Subject == CRLSignerCertificate.Issuer` and `SubjectKeyID == CRLSignerCertificate.AuthorityKeyId`
       - Build certification path with both elements, verify path
- Check that (VendorID, IssuerSubjectKeyID, Label) combination is unique when adding the distribution endpoint.
- Check that (VendorID, IssuerSubjectKeyID, DataUrl) combination is unique when adding the

#### UPDATE_PKI_REVOCATION_DISTRIBUTION_POINT

- Check that Revocation Distribution Point is found by (VendorID, Label, IssuerSubjectKeyID)
- The sender must be a Vendor account
- Check that Revocation Distribution Point's `VendorID` field must be equal to the Vendor account's VID.
- If `crlSignerCertificate` is provided:
   - If Revocation Distribution Point's Signer Certificate is a PAA (root certificate, self-signed):
      - `crlSignerCertificate` must be a PAA (root certificate, self-signed)
      - if `crlSignerCertificate` encodes a VID in its subject, it must be equal to Revocation Distribution Point's `VendorID` field
      - if `crlSignerCertificate` doesn't encode a VID in its subject, throw Unsupported Error 
   - If Revocation Distribution Point's Signer Certificate is a PAI (intermediate certificate, not self-signed):
     - `crlSignerCertificate` must be a PAI
     - `crlSignerCertificate` must encode a vid in its subject and this vid must be equal to `VendorID` field
     - if `crlSignerCertificate` encodes a pid in its subject, it must be equal to the `ProductID` field
- Check that (VendorID, IssuerSubjectKeyID, DataUrl) combination is unique when updating the distribution endpoint.

#### DELETE_PKI_REVOCATION_DISTRIBUTION_POINT

- Check that Revocation Distribution Point is found by (VendorID, Label, IssuerSubjectKeyID)
- Check that the sender is a Vendor account
- Check that Revocation Distribution Point's `VendorID` field must be equal to the Vendor account's VID.
