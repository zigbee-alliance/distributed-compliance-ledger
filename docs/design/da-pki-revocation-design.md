# DA PKI Revocation

## User Stories

1. Revocation of PAI

   A Vendor with DCL write privilege can submit a transaction to publish the location of the CRL distribution point for a PAA associated with the Vendor ID. This information can be updated by the Vendor.

2. Revocation of a DAC 

   A Vendor with DCL write privilege can submit a transaction to publish the location of the CRL distribution point for a PAI associated with the Vendor ID. This information can be updated by the Vendor.


## Revocation Distribution Point Schema
The schema is the same as in `PR6360`.
However, the validation logic is proposed to be a bit different (see [Validation Logic](#validation-logic) and [Divergence](#divergence-from-pr6360)). 


| Name                 | Type    | Constraint          | Mutable | Conformance |
|:---------------------|:--------|:--------------------|:--------|-------------|
| VendorID             | uint16  | all                 | No      | M           |
| ProductID            | uint16* | all                 | No      | O (desc)    |
| IsPAA                | bool    | all                 | No      | M           |
| Label                | string  | max 64              | No      | M           |
| CRLSignerCertificate | string  | max 2048            | No      | M           |
| IssuerSubjectKeyID   | string  | max 64              | No      | M           |
| DataUrl              | string  | max 256             | Yes     | M           |
| DataFileSize         | uint64  | desc                | Yes     | O           |
| DataDigest           | string  | max 128             | Yes     | O (desc)    |
| DataDigestType       | uint32  | all                 | Yes     | O (desc)    |
| RevocationType       | uint32  | all                 | No      | M           |

*: `PR6360` defines PID as `uint32` which is incorrect. PID must be `uint16` as in other places.

## Divergence from PR6360
1. No validation of `DataUrl` content.
   If the content is invalid, it can be edited/removed by either owner Vendor or Trustees via Update and Delete commands.

   Such validation requires resolving an external URL, and this is a non-deterministic operation:
   some nodes may resolve the URL successfully, but some nodes may not resolve it or be redirected to a different content.
   Non-deterministic results of validation may cause a situation that on some nodes the transaction is considered as valid,
   and on some nodes as invalid. That will break the consensus protocol and DCL won't be able to process write requests anymore.
   An alternative option is to pass the whole content of `DataUrl` to the transaction, but it doesn't make a lot of sense because
    1. the content can be quite big (1 MB)
    2. there will be no guarantee that `DataUrl` actually points to the same content as passed

2. Validation of `CRLSignerCertificate`. If `CRLSignerCertificate` is a PAA, then it must be present on DCL (there must be a PAA on DCL with the same pem value).
   If `CRLSignerCertificate` is a PAI, then it must be chained back to a valid PAA present on the ledger.

3. If `IssuerSubjectKeyID` is not equal to the `Subject Key Identifier` of `CRLSignerCertificate`, the transaction is considered invalid and rejected.

4. If `RevocationType` is 1 (RFC5280 CRL), then `DataFileSize`, `DataDigest`, `DataDigestType` must be empty.

   Otherwise, it may cause confusion. Certificates to be revoked are added to the content of CRL URL without a need to send a transaction to the ledger.
   So, someone may look at the DCL entry, see that a file size is specified, then go to the URL, and see that the actual size there is different.
   It may cause a question if the data is valid, and why a different file size is in DCL.

5. `ProductID` is expected to be `unit16`, not `unit32` as in other places.


## A need for Proof-of-possession of `CRLSignerCertificate` key
No additional proof-of-possession is required because 
- The transaction is already signed by the Vendor account key, so only a trusted Vendor can send it.
- DCL checks that `CRLSignerCertificate` is a valid PAI (signed by a trusted PAA from DCL) or a PAA present on the ledger.
- It can be complicated to create an interactive PoP of `CRLSignerCertificate` for some PAIs.
- Indirect CRL are out of scope for now

## Transactions

### 1. ADD_PKI_REVOCATION_DISTRIBUTION_POINT
Publishes a PKI Revocation distribution endpoint owned by the Vendor. 

If `crlSignerCertificate` is a PAA (root certificate), then it must be present on DCL.

If `crlSignerCertificate` is a PAI (intermediate certificate), then it must be chained back to a valid PAA (root certificate) present on DCL.
In this case `crlSignerCertificate` is not required to be present on DCL, and will not be added to DCL as a result of this transaction.
If PAI needs to be added to DCL, it should be done via `ADD_X509_CERT` transaction.

Publishing the revocation distribution endpoint doesn't automatically remove PAI (Intermediate certificates)
and DACs (leaf certificates) added to DCL if they are revoked in the CRL identified by this distribution point.


- Who can send:
    - Vendor account
    - `vid` field in the transaction (`VendorID`) must be equal to the Vendor account's VID
    - `vid` field in the `CRLSignerCertificate` (for vendor-scoped PAAs and PAIs) must be equal to the Vendor account's VID
- Parameters:
    - vid: `uint16` -  Vendor ID (positive non-zero)
    - pid: `optional(uint16)` -  Product ID (positive non-zero)
    - isPAA: `bool` -  True if the revocation information distribution point relates to a PAA
    - label: `string` -  A label to disambiguate multiple revocation information partitions of a particular issuer.
    - crlSignerCertificate: `string` -  PEM encoded certificate (string or path to file containing data)
    - issuerSubjectKeyID: `string` -  crlSignerCertificate's `Subject Key Id` in hex string format, e.g: `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`
    - dataUrl: `string` -  The URL where to obtain the information in the format indicated by the RevocationType field
    - dataFileSize: `optional(uint64)` -  Total size in bytes of the file found at the DataUrl. Must be omitted if RevocationType is 1.
    - dataDigest: `optional(string)` -  Digest of the entire contents of the associated file downloaded from the DataUrl. Must be omitted if RevocationType is 1.
    - dataDigestType: `optional(uint32)` - The type of digest used in the DataDigest field from the list of [1, 7, 8, 10, 11, 12] (IANA Named Information Hash Algorithm Registry).
    - revocationType: `uint32` - The type of file found at the DataUrl for this entry. Supported types: 1 - RFC5280 Certificate Revocation List (CRL).
- In State:
    - `pki/RevocationDistributionPoint/value/<IssuerSubjectKeyID>` -> list of Revocation Distribution Points
    - `pki/RevocationDistributionPoint/value/<IssuerSubjectKeyID>/<vid>/<label>`-> Revocation Distribution Point
- CLI command:
    - `dcld tx pki add-revocation-point --vid=<uint16> --pid=<uint16> --is-paa=<bool> --label=<string>
      --certificate=<string-or-path> --data-url=<string> --revocation-type=1 --from=<account>`


### 2. UPDATE_PKI_REVOCATION_DISTRIBUTION_POINT
Updates an existing PKI Revocation distribution endpoint owned by the sender.

- Who can send:
    - Vendor account
    - `vid` field in the transaction (`VendorID`) must be equal to the Vendor account's VID
- Parameters:
    - vid: `uint16` -  Vendor ID (positive non-zero)
    - label: `string` -  A label to disambiguate multiple revocation information partitions of a particular issuer.
    - issuerSubjectKeyID: `string` -  issuer certificate's `Subject Key Id` in hex string format, e.g: `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`
    - dataUrl: `string` -  The URL where to obtain the information in the format indicated by the RevocationType field
    - dataFileSize: `optional(uint64)` -  Total size in bytes of the file found at the DataUrl. Must be omitted if RevocationType is 1.
    - dataDigest: `optional(string)` -  Digest of the entire contents of the associated file downloaded from the DataUrl. Must be omitted if RevocationType is 1.
    - dataDigestType: `optional(uint32)` - The type of digest used in the DataDigest field from the list of [1, 7, 8, 10, 11, 12] (IANA Named Information Hash Algorithm Registry).
- In State:
    - `pki/RevocationDistributionPoint/value/<IssuerSubjectKeyID>` -> list of Revocation Distribution Points
    - `pki/RevocationDistributionPoint/value/<IssuerSubjectKeyID>/<vid>/<label>` -> Revocation Distribution Point
- CLI command:
    - `dcld tx pki update-revocation-point --vid=<uint16> --label=<string>
      --subject-key-id=<string> --data-url=<string> --from=<account>`

### 3. DELETE_DELETE_PKI_REVOCATION_DISTRIBUTION_POINT
Deletes a PKI Revocation distribution endpoint owned by the sender.

- Who can send:
    - Vendor account
    - `vid` field in the transaction (`VendorID`) must be equal to the Vendor account's VID
- Parameters:
    - vid: `uint16` -  Vendor ID (positive non-zero)
    - label: `string` -  A label to disambiguate multiple revocation information partitions of a particular issuer.
    - issuerSubjectKeyID: `string` -  issuer certificate's `Subject Key Id` in hex string format, e.g: `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`
- In State:
    - `pki/RevocationDistributionPoint/value/<IssuerSubjectKeyID>` -> list of Revocation Distribution Points
    - `pki/RevocationDistributionPoint/value/<IssuerSubjectKeyID>/<vid>/<label>` -> Revocation Distribution Point
- CLI command:
    - `dcld tx pki delete-revocation-point --vid=<uint16> --label=<string> --subject-key-id=<string> --from=<account>`


## Query

### 1. GET_PKI_REVOCATION_DISTRIBUTION_POINT
Gets a revocation distribution point identified by (VendorID, Label, IssuerSubjectKeyID) unique combination.
Use `GET_ALL_PKI_REVOCATION_DISTRIBUTION_POINT` to get a list of all revocation distribution points.

- Parameters:
  - vid: `uint16` -  Vendor ID (positive non-zero)
  - label: `string` -  A label to disambiguate multiple revocation information partitions of a particular issuer.
  - issuerSubjectKeyID: `string` -  issuer certificate's `Subject Key Id` in hex string format, e.g: `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`
- CLI command:
  - `dcld query pki revocation-point --vid=<uint16> --label=<string> --subject-key-id=<string>`
- REST API:
  - GET `/dcl/pki/revocation-points/{subject_key_id}/{vid}/{label}`
  
### 2. GET_PKI_REVOCATION_DISTRIBUTION_POINTS_BY_SUBJECT_KEY_ID
Gets a list of revocation distribution point identified by IssuerSubjectKeyID.

- Parameters:
  - issuerSubjectKeyID: `string` -  issuer certificate's `Subject Key Id` in hex string format, e.g: `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`
- CLI command:
  - `dcld query pki revocation-points --subject-key-id=<string>`
- REST API:
  - GET `/dcl/pki/revocation-points/{subject_key_id}`
  - 
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
- `DataUrl` starts with either `http` or `https`.
- `DataDigest` is present if and only if the `DataFileSize` field is present.
- `DataDigestType` is provided if and only if the `DataDigest` field is present. 
- If `RevocationType` is 1 (RFC5280 CRL), then `DataFileSize`, `DataDigest`, `DataDigestType` must be empty.
- Check that `ProductID` field is provided if and only if `IsPAA` is false and `CRLSignerCertificate` has a PID in its subject.  
- Check that `ProductID` must be equal to the PID in `CRLSignerCertificate`'s subject.
- Check that `VendorID` is equal to the vid in `CRLSignerCertificate`'s subject.
- Check that `IssuerSubjectKeyID` is equal to the `CRLSignerCertificate`'s Subject Key Identifier.
- If `IsPAA` is true, then
  - check that the `CRLSignerCertificate` is a PAA (root certificate, self-signed).
  - If `CRLSignerCertificate` encodes a vid in its subject, then it must be equal to `VendorID` field.
- If `IsPAA` is false, then
  - check that `CRLSignerCertificate` is a non-root certificate (not self-signed).
  - `CRLSignerCertificate` must encode a vid in its subject equal to `VendorID` field.

### Dynamic validation (when adding to a block)
- `VendorID` field must be equal to the Vendor account's VID
- If `crlSignerCertificate` is a PAA (root certificate, self-signed):
   - Query a certificate by `crlSignerCertificate` Subject and Subject Key ID. If it's not found - error.
   - Check that pem value of the found certificate is equal to `crlSignerCertificate` value.
- If `crlSignerCertificate` is a PAI (intermediate certificate, not self-signed):
   - Check that `crlSignerCertificate` is chained back to certificates present on DCL (`verifyCertificate` method).
- Check that (VendorID, Label, IssuerSubjectKeyID) combination is unique when adding the distribution endpoint.