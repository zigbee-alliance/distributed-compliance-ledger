# X509 PKI

<!-- markdownlint-disable MD036 -->

**NOTE**: X.509 v3 certificates are only supported (all certificates MUST contain `Subject Key ID` field).
All PKI related methods are based on this restriction.


### All Certificates (DA, NOC)

#### GET_CERT

**Status: Implemented**

Gets a certificate by the given subject and subject key ID attributes. This query works for all types of certificates (PAA, PAI, RCAC, ICAC).
Revoked certificates are not returned.
Use [GET_REVOKED_DA_CERT](#get_revoked_da_cert) to get a revoked DA certificate.
Use [GET_REVOKED_NOC_ROOT_CERT](#get_revoked_noc_root-rcac) to get a revoked Noc Root certificate.
Use [GET_REVOKED_NOC_ICA_CERT](#get_revoked_noc_ica-icac) to get a revoked Noc ICA certificate.

- Parameters:
  - subject: `string`  - certificates's `Subject` is base64 encoded subject DER sequence bytes
  - subject_key_id: `string`  - certificates's `Subject Key Id` in hex string format, e.g: `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`
- CLI command:
  - `dcld query pki cert --subject=<base64 string> --subject-key-id=<hex string>`
- REST API:
  - GET `/dcl/pki/all-certificates/{subject}/{subject_key_id}`

#### GET_ALL_CERTS

**Status: Implemented**

Gets all certificates. This query works for all types of certificates (PAA, PAI, RCAC, ICAC).

Revoked certificates are not returned.
Use [GET_ALL_REVOKED_DA_CERTS](#get_all_revoked_da_certs) to get a list of all revoked DA certificates.
Use [GET_ALL_REVOKED_NOC_ROOT_CERTS](#get_all_revoked_noc_root-rcacs) to get a list of all revoked Noc Root certificates.
Use [GET_ALL_REVOKED_NOC_ICA_CERTS](#get_all_revoked_noc_ica-icacs) to get a list of all revoked Noc ICA certificates.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query pki all-certs`
- REST API:
  - GET `/dcl/pki/all-certificates`

#### GET_CHILD_CERTS

**Status: Implemented**

Gets all child certificates for the given certificate. This query works for both PAI and NOC_ICA.
Revoked certificates are not returned.

- Parameters:
  - subject: `string`  - certificates's `Subject` is base64 encoded subject DER sequence bytes
  - subject_key_id: `string`  - certificates's `Subject Key Id` in hex string format, e.g: `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`
- CLI command:
  - `dcld query pki all-child-x509-certs (--subject=<base64 string> --subject-key-id=<hex string>`
- REST API:
  - GET `/dcl/pki/child-certificates/{subject}/{subject_key_id}`

### Device Attestation Certificates (DA): PAA, PAI

#### PROPOSE_ADD_PAA

**Status: Implemented**

Proposes a new PAA (self-signed root certificate).

If more than 1 Trustee signature is required to add the PAA certificate, the PAA certificate
will be in a pending state until sufficient number of approvals is received.

The PAA certificate is immutable. It can only be revoked by either the owner or a quorum of Trustees.

- Who can send:
  - Trustee
- Parameters:
  - cert: `string` - PEM encoded certificate. The corresponding CLI parameter can contain either a PEM string or a path to a file containing the data.
  - info: `optional(string)` - information/notes for the proposal. Can contain up to 4096 characters.
  - time: `optional(int64)` - proposal time (number of nanoseconds elapsed since January 1, 1970 UTC). This field cannot be specified using a CLI command and will use the current time by default.
  - vid: `uint16` -  Vendor ID (positive non-zero). Must be equal to the Certificate's `vid` field for VID-scoped PAA.
  - schemaVersion: `optional(uint16)` - Certificate's schema version to support backward/forward compatability. Should be equal to 0 (default 0)
- In State: `pki/ProposedCertificate/value/<Certificate's Subject>/<Certificate's Subject Key ID>`
- CLI command:
  - `dcld tx pki propose-add-x509-root-cert --certificate=<string-or-path> --from=<account>`
- Validation:
  - provided certificate must be root:
    - `Issuer` == `Subject`
    - `Authority Key Identifier` == `Subject Key Identifier`
  - no existing `Proposed` certificate with the same `<Certificate's Subject>:<Certificate's Subject Key ID>` combination.
  - no existing certificate with the same `<Certificate's Issuer>:<Certificate's Serial Number>` combination.
  - if approved certificates with the same `<Certificate's Subject>:<Certificate's Subject Key ID>` combination already exists:
    - the existing certificate must not be NOC certificate
    - sender must match to the owner of the existing certificates.
  - the signature (self-signature) and expiration date are valid.

#### APPROVE_ADD_PAA

**Status: Implemented**

Approves the proposed PAA (self-signed root certificate). It also can be used for revote (i.e. change vote from reject to approve)

The PAA certificate is not active until sufficient number of Trustees approve it.

- Who can send:
  - Trustee
- Parameters:
  - subject: `string`  - proposed certificates's `Subject` is base64 encoded subject DER sequence bytes.
  - subject_key_id: `string`  - proposed certificates's `Subject Key Id` in hex string format, e.g: `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`.
  - info: `optional(string)` - information/notes for the approval. Can contain up to 4096 characters.
  - time: `optional(int64)` - proposal time (number of nanoseconds elapsed since January 1, 1970 UTC). This field cannot be specified using a CLI command and will use the current time by default.
- In State:
  - `pki/AllCertificates/value/<Certificate's Subject>/<Certificate's Subject Key ID>`.
  - `pki/ApprovedCertificates/value/<Certificate's Subject>/<Certificate's Subject Key ID>`.
  - `pki/ApprovedCertificatesBySubject/value/<Certificate's Subject>`
  - `pki/ApprovedCertificatesBySubjectKeyId/value/<Certificate's Subject Key ID>`.
- Number of required approvals:
  - greater than or equal 2/3 of Trustees (proposal by a Trustee is also counted as an approval)
- CLI command:
  - `dcld tx pki approve-add-x509-root-cert --subject=<base64 string> --subject-key-id=<hex string> --from=<account>`
- Validation:
  - the proposal to add a root certificate with the provided subject and subject_key_id, must be submitted first.
  - the proposed certificate hasn't been approved by the signer yet.

#### REJECT_ADD_PAA

**Status: Implemented**

Rejects the proposed PAA (self-signed root certificate). It also can be used for revote (i.e. change vote from approve to reject)

If proposed PAA certificate has only proposer's approval and no rejects then proposer can send this transaction to remove the proposal

The certificate is not reject until sufficient number of Trustees reject it.

- Who can send:
  - Trustee
- Parameters:
  - subject: `string` - proposed certificates's `Subject` is base64 encoded subject DER sequence bytes
  - subject_key_id: `string` - proposed certificates's `Subject Key Id` in hex string format, e.g:
  `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`
  - info: `optional(string)` - information/notes for the reject. Can contain up to 4096 characters.
  - time: `optional(int64)` - reject time (number of nanoseconds elapsed since January 1, 1970 UTC). This field cannot be specified using a CLI command and will use the current time by default.
- In State: `pki/RejectedCertificates/value/<Certificate's Subject>/<Certificate's Subject Key ID>`
- Number of required rejects:
  - more than 1/3 of Trustees
- CLI command:
  - `dcld tx pki reject-add-x509-root-cert --subject=<base64 string> --subject-key-id=<hex string> --from=<account>`
- Validation:
  - the proposal to add a root certificate with the provided subject and subject_key_id, must be submitted first.
  - the proposed certificate hasn't been rejected by the signer yet

#### PROPOSE_REVOKE_PAA

**Status: Implemented**

Proposes revocation of the given PAA (self-signed root certificate) by a Trustee.

Revocation works as a soft-delete, meaning that the certificates are not entirely removed but moved from the approved list to the revoked list.
Revoked certificates can be retrieved by using the [GET_REVOKED_DA_CERT](#get_revoked_da_cert) query.

If a Revocation Distribution Point needs to be published (such as RFC5280 Certificate Revocation List), please use [ADD_REVOCATION_DISTRIBUTION_POINT](#add_revocation_distribution_point).

If `revoke-child` flag is set to `true` then all the certificates in the chain signed by the revoked certificate will be revoked as well.

If more than 1 Trustee signature is required to revoke a PAA certificate,
then the certificate will be in a pending state until sufficient number of other Trustee's approvals is received.

- Who can send:
  - Trustee
- Parameters:
  - subject: `string`  - certificates's `Subject` is base64 encoded subject DER sequence bytes.
  - subject_key_id: `string`  - certificates's `Subject Key Id` in hex string format, e.g: `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`.
  - serial-number: `optional(string)`  - certificate's serial number.
  - revoke-child: `optional(bool)`  - to revoke child certificates in the chain - default is false.
  - info: `optional(string)` - information/notes for the revocation proposal. Can contain up to 4096 characters.
  - time: `optional(int64)` - revocation proposal time (number of nanoseconds elapsed since January 1, 1970 UTC). This field cannot be specified using a CLI command and will use the current time by default.
- In State: `pki/ProposedCertificateRevocation/value/<Certificate's Subject>/<Certificate's Subject Key ID>`
- CLI command:
  - `dcld tx pki propose-revoke-x509-root-cert --subject=<base64 string> --subject-key-id=<hex string> --from=<account>`
- Validation:
  - revoked certificate must be root:
    - `Issuer` == `Subject`
    - `Authority Key Identifier` == `Subject Key Identifier`
  - no existing `Proposed` certificate with the same `<Certificate's Subject>:<Certificate's Subject Key ID>` combination.

#### APPROVE_REVOKE_PAA

**Status: Implemented**

Approves the revocation of the given PAA (self-signed root certificate) by a Trustee.

Revocation works as a soft-delete, meaning that the certificates are not entirely removed but moved from the approved list to the revoked list.
Revoked certificates can be retrieved by using the [GET_REVOKED_DA_CERT](#get_revoked_da_cert) query.

If a Revocation Distribution Point needs to be published (such as RFC5280 Certificate Revocation List), please use [ADD_REVOCATION_DISTRIBUTION_POINT](#add_revocation_distribution_point).

The revocation is not applied until sufficient number of Trustees approve it.

- Who can send:
  - Trustee
- Parameters:
  - subject: `string`  - certificates's `Subject` is base64 encoded subject DER sequence bytes.
  - subject_key_id: `string`  - certificates's `Subject Key Id` in hex string format, e.g: `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`.
  - serial-number: `optional(string)` - certificate's serial number.
  - info: `optional(string)` - information/notes for the revocation approval. Can contain up to 4096 characters.
  - time: `optional(int64)` - revocation approval time (number of nanoseconds elapsed since January 1, 1970 UTC). This field cannot be specified using a CLI command and will use the current time by default.
- In State: `pki/RevokedCertificates/value/<Certificate's Subject>/<Certificate's Subject Key ID>`
- Number of required approvals:
  - greater than or equal 2/3 of Trustees (proposal by a Trustee is also counted as an approval)
- CLI command:
  - `dcld tx pki approve-revoke-x509-root-cert --subject=<base64 string> --subject-key-id=<hex string> --from=<account>`
- Validation:
  - the proposal to revoke a root certificate with the provided subject and subject_key_id, must be submitted first.
  - the proposed certificate revocation hasn't been approved by the signer yet.

#### ASSIGN_VID_TO_PAA

**Status: Implemented**

Assigns a Vendor ID (VID) to non-VID scoped PAAs (self-signed root certificate) already present on the ledger.

- Who can send:
  - Vendor Admin
- Parameters:
  - subject: `string` - certificates's `Subject` is base64 encoded subject DER sequence bytes.
  - subject_key_id: `string` - certificates's `Subject Key Id` in hex string format, e.g: `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`.
  - vid: `uint16` - Vendor ID (positive non-zero). Must be the same as `vid` field in the VID-scoped PAA certificate.
- CLI command:
  - `dcld tx pki assign-vid --subject=<base64 string> --subject-key-id=<hex string> --vid=<uint16> --from=<account>`
- Validation:
  - PAA Certificate with the provided `subject` and `subject_key_id` must exist in the ledger.
  - If the PAA is a VID scoped one, then the `vid` field must be equal to the VID value in the PAA's subject.

#### ADD_REVOCATION_DISTRIBUTION_POINT

**Status: Implemented**

Publishes a PKI Revocation distribution endpoint (such as RFC5280 Certificate Revocation List) owned by the Vendor.

If `crlSignerCertificate` is a PAA (root certificate), then it must be present on DCL.

If `crlSignerCertificate` is a PAI (intermediate certificate) or delegated by PAA, then it must be chained back to a valid PAA (root certificate) present on DCL.
In this case `crlSignerCertificate` is not required to be present on DCL, and will not be added to DCL as a result of this transaction.
If PAI needs to be added to DCL, it should be done via [ADD_PAI](#add_pai) transaction.

Publishing the revocation distribution endpoint doesn't automatically remove PAI (Intermediate certificates)
and DACs (leaf certificates) added to DCL if they are revoked in the CRL identified by this distribution point.
[REVOKE_PAI](#revoke_pai) needs to be called to remove an intermediate or leaf certificate from the ledger.∂

- Who can send: Vendor account
  - `vid` field in the transaction (`VendorID`) must be equal to the Vendor account's VID
  - VID-scoped PAAs (Root certs) and PAIs (Intermediate certs): `vid` field in the `CRLSignerCertificate`'s subject must be equal to the Vendor account's VID
  - Non-VID scoped PAAs (Root certs): `vid` field associated with the corresponding PAA on the ledger must be equal to the Vendor account's VID
- Parameters:
  - vid: `uint16` -  Vendor ID (positive non-zero). Must be the same as Vendor account's VID and `vid` field in the VID-scoped `CRLSignerCertificate`. Must be the same as a `vid` associated with non-VID scoped `CRLSignerCertificate` on the ledger.
  - pid: `optional(uint16)` -  Product ID (positive non-zero). Must be empty if `IsPAA` is true. Must be equal to a `pid` field in `CRLSignerCertificate`.
  - isPAA: `bool` -  True if the revocation information distribution point relates to a PAA
  - label: `string` -  A label to disambiguate multiple revocation information partitions of a particular issuer.
  - crlSignerCertificate: `string` - The issuer certificate whose revocation information is provided in the distribution point entry, encoded in X.509v3 PEM format. The corresponding CLI parameter can contain either a PEM string or a path to a file containing the data. Please note that if crlSignerCertificate is a delegated certificate by a PAI, the delegator certificate must be provided using the `crlSignerDelegator` field.
  - crlSignerDelegator: `optional(string)` - If crlSignerCertificate is a delegated certificate by a PAI, then crlSignerDelegator must contain the delegator PAI certificate which must be chained back to an approved certificate in the ledger, encoded in X.509v3 PEM format. Otherwise this field can be omitted. The corresponding CLI parameter can contain either a PEM string or a path to a file containing the data.
  - issuerSubjectKeyID: `string` - Uniquely identifies the PAA or PAI for which this revocation distribution point is provided. Must consist of even number of uppercase hexadecimal characters ([0-9A-F]), with no whitespace and no non-hexadecimal characters., e.g: `5A880E6C3653D07FB08971A3F473790930E62BDB`.
  - dataUrl: `string` -  The URL where to obtain the information in the format indicated by the RevocationType field. Must start with either `http` or `https`. Must be unique for all pairs of VendorID and IssuerSubjectKeyID.
  - dataFileSize: `optional(uint64)` -  Total size in bytes of the file found at the DataUrl. Must be omitted if RevocationType is 1.
  - dataDigest: `optional(string)` -  Digest of the entire contents of the associated file downloaded from the DataUrl. Must be omitted if RevocationType is 1. Must be provided if and only if the `DataFileSize` field is present.
  - dataDigestType: `optional(uint32)` - The type of digest used in the DataDigest field from the list of [1, 7, 8, 10, 11, 12] (IANA Named Information Hash Algorithm Registry). Must be provided if and only if the `DataDigest` field is present.
  - revocationType: `uint32` - The type of file found at the DataUrl for this entry. Supported types: 1 - RFC5280 Certificate Revocation List (CRL).
  - schemaVersion: `optional(uint16)` - Schema version to support backward/forward compatibility. Should be equal to 0 (default 0)
- In State:
  - `pki/RevocationDistributionPoint/value/<IssuerSubjectKeyID>` -> list of Revocation Distribution Points
  - `pki/RevocationDistributionPoint/value/<IssuerSubjectKeyID>/<vid>/<label>`-> Revocation Distribution Point
- CLI command:
  - `dcld tx pki add-revocation-point --vid=<uint16> --pid=<uint16> --issuer-subject-key-id=<string> --is-paa=<bool> --label=<string>
    --certificate=<string-or-path> --certificate-delegator=<string-or-path> --data-url=<string> --revocation-type=1 --from=<account>`

#### UPDATE_REVOCATION_DISTRIBUTION_POINT

**Status: Implemented**

Updates an existing PKI Revocation distribution endpoint (such as RFC5280 Certificate Revocation List) owned by the Vendor.

- Who can send: Vendor account
  - `vid` field in the transaction (`VendorID`) must be equal to the Vendor account's VID
  - VID-scoped PAAs (Root certs) and PAIs (Intermediate certs): `vid` field in the `CRLSignerCertificate`'s subject must be equal to the Vendor account's VID
  - Non-VID scoped PAAs (Root certs): `vid` field associated with the corresponding PAA on the ledger must be equal to the Vendor account's VID
- Parameters:
  - vid: `uint16` -  Vendor ID (positive non-zero). Must be the same as Vendor account's VID and `vid` field in the VID-scoped `CRLSignerCertificate`. Must be the same as a `vid` associated with non-VID scoped `CRLSignerCertificate` on the ledger.
  - label: `string` -  A label to disambiguate multiple revocation information partitions of a particular issuer.
  - issuerSubjectKeyID: `string` - Uniquely identifies the PAA or PAI for which this revocation distribution point is provided. Must consist of even number of uppercase hexadecimal characters ([0-9A-F]), with no whitespace and no non-hexadecimal characters., e.g: `5A880E6C3653D07FB08971A3F473790930E62BDB`.
  - crlSignerCertificate: `optional(string)` - The issuer certificate whose revocation information is provided in the distribution point entry, encoded in X.509v3 PEM format. The corresponding CLI parameter can contain either a PEM string or a path to a file containing the data. Please note that if crlSignerCertificate is a delegated certificate by a PAI, the delegator certificate must be provided using the `crlSignerDelegator` field.
  - crlSignerDelegator: `optional(string)` - If crlSignerCertificate is a delegated certificate by a PAI, then crlSignerDelegator must contain the delegator PAI certificate which must be chained back to an approved certificate in the ledger, encoded in X.509v3 PEM format. Otherwise this field can be omitted. The corresponding CLI parameter can contain either a PEM string or a path to a file containing the data.
  - dataUrl: `optional(string)` -  The URL where to obtain the information in the format indicated by the RevocationType field. Must start with either `http` or `https`. Must be unique for all pairs of VendorID and IssuerSubjectKeyID.
  - dataFileSize: `optional(uint64)` -  Total size in bytes of the file found at the DataUrl. Must be omitted if RevocationType is 1.
  - dataDigest: `optional(string)` -  Digest of the entire contents of the associated file downloaded from the DataUrl. Must be omitted if RevocationType is 1. Must be provided if and only if the `DataFileSize` field is present.
  - dataDigestType: `optional(uint32)` - The type of digest used in the DataDigest field from the list of [1, 7, 8, 10, 11, 12] (IANA Named Information Hash Algorithm Registry). Must be provided if and only if the `DataDigest` field is present.
  - schemaVersion: `optional(uint16)` - Schema version to support backward/forward compatibility. Should be equal to 0 (default 0)
- In State:
  - `pki/RevocationDistributionPoint/value/<IssuerSubjectKeyID>` -> list of Revocation Distribution Points
  - `pki/RevocationDistributionPoint/value/<IssuerSubjectKeyID>/<vid>/<label>` -> Revocation Distribution Point
- CLI command:
  - `dcld tx pki update-revocation-point --vid=<uint16> --issuer-subject-key-id=<string> --label=<string>
    --data-url=<string> --certificate=<string-or-path> --certificate-delegator=<string-or-path> --from=<account>`

#### DELETE_REVOCATION_DISTRIBUTION_POINT

**Status: Implemented**

Deletes a PKI Revocation distribution endpoint (such as RFC5280 Certificate Revocation List)  owned by the Vendor.

- Who can send: Vendor account
  - `vid` field in the transaction (`VendorID`) must be equal to the Vendor account's VID
  - VID-scoped PAAs (Root certs) and PAIs (Intermediate certs): `vid` field in the `CRLSignerCertificate`'s subject must be equal to the Vendor account's VID
  - Non-VID scoped PAAs (Root certs): `vid` field associated with the corresponding PAA on the ledger must be equal to the Vendor account's VID
- Parameters:
  - vid: `uint16` -  Vendor ID (positive non-zero). Must be the same as Vendor account's VID and `vid` field in the VID-scoped `CRLSignerCertificate`. Must be the same as a `vid` associated with non-VID scoped `CRLSignerCertificate` on the ledger.
  - label: `string` -  A label to disambiguate multiple revocation information partitions of a particular issuer.
  - issuerSubjectKeyID: `string` - Uniquely identifies the PAA or PAI for which this revocation distribution point is provided. Must consist of even number of uppercase hexadecimal characters ([0-9A-F]), with no whitespace and no non-hexadecimal characters., e.g: `5A880E6C3653D07FB08971A3F473790930E62BDB`.
- In State:
  - `pki/RevocationDistributionPoint/value/<IssuerSubjectKeyID>` -> list of Revocation Distribution Points
  - `pki/RevocationDistributionPoint/value/<IssuerSubjectKeyID>/<vid>/<label>` -> Revocation Distribution Point
- CLI command:
  - `dcld tx pki delete-revocation-point --vid=<uint16> --issuer-subject-key-id=<string> --label=<string> --from=<account>`

#### ADD_PAI

**Status: Implemented**

Adds a PAI (intermediate certificate) signed by a chain of certificates which must be already present on the ledger.

- Who can send:
  - Vendor Account
- Parameters:
  - cert: `string` - PEM encoded certificate. The corresponding CLI parameter can contain either a PEM string or a path to a file containing the data.
  - certificate-schema-version: `optional(uint16)` - Certificate's schema version to support backward/forward compatability(default 0)
- In State:
  - `pki/ApprovedCertificates/value/<Certificate's Subject>/<Certificate's Subject Key ID>`
  - `pki/ChildCertificates/value/<Certificate's Subject>/<Certificate's Subject Key ID>`
- CLI command:
  - `dcld tx pki add-x509-cert --certificate=<string-or-path> --from=<account>`
- Validation:
  - provided certificate must not be root:
    - `Issuer` != `Subject`
    - `Authority Key Identifier` != `Subject Key Identifier`
  - no existing certificate with the same `<Certificate's Issuer>:<Certificate's Serial Number>` combination.
  - if certificates with the same `<Certificate's Subject>:<Certificate's Subject Key ID>` combination already exist:
    - the existing certificate must not be NOC certificate.
    - the sender's VID must match the VID of the existing certificate's owner.
  - the signature and expiration date are valid.
  - parent certificate must be already stored on the ledger and a valid chain to some root certificate can be built.
  - if the parent root certificate is VID scoped:
    - the provided certificate must also be VID scoped.
    - the `vid` in the subject of the root certificate must be equal to the `vid` in the subject of the provided certificate.
    - the `vid` in the subjects of both certificates must be equal to the sender Vendor account's VID.
  - if the parent root certificate is not VID scoped but has an associated VID:
    - the provided certificate can be either VID scoped or non-VID scoped.
    - if the provided certificate is VID scoped, the `vid` in the subject of the certificate must be equal to the VID associated with the root certificate and to the sender Vendor account's VID.
  - if the parent root certificate is non-VID scoped and does not have an associated VID, an error will occur.

> **_Note:_**  Multiple certificates can refer to the same `<Certificate's Subject>:<Certificate's Subject Key ID>` combination.

#### REVOKE_PAI

**Status: Implemented**

Revokes the given PAI (intermediate certificate).

Revocation works as a soft-delete, meaning that the certificates are not entirely removed but moved from the approved list to the revoked list.
Revoked certificates can be retrieved by using the [GET_REVOKED_DA_CERT](#get_revoked_da_cert) query.
To entirely remove a PAI certificate, please use [REMOVE_PAI](#remove_pai).

If a Revocation Distribution Point needs to be published (such as RFC5280 Certificate Revocation List), please use [ADD_REVOCATION_DISTRIBUTION_POINT](#add_revocation_distribution_point).

If `revoke-child` flag is set to `true` then all the certificates in the chain signed by the revoked certificate will be revoked as well.

Root certificates can not be revoked this way, use  [PROPOSE_REVOKE_PAA](#propose_revoke_paa) and [APPROVE_REVOKE_PAA](#approve_revoke_paa) instead.  

- Who can send: Vendor account
  - the sender's VID must match the VID of the revoking certificate's owner.
- Parameters:
  - subject: `string`  - certificates's `Subject` is base64 encoded subject DER sequence bytes.
  - subject_key_id: `string`  - certificates's `Subject Key Id` in hex string format, e.g: `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`.
  - serial-number: `optional(string)` - certificate's serial number.
  - revoke-child: `optional(bool)` - to revoke child certificates in the chain - default is false.
  - info: `optional(string)` - information/notes for the revocation. Can contain up to 4096 characters.
  - time: `optional(int64)` - revocation time (number of nanoseconds elapsed since January 1, 1970 UTC). This field cannot be specified using a CLI command and will use the current time by default.
- In State: `pki/RevokedCertificates/value/<Certificate's Subject>/<Certificate's Subject Key ID>`
- CLI command:
  - `dcld tx pki revoke-x509-cert --subject=<base64 string> --subject-key-id=<hex string> --from=<account>`
- Validation:
  - a PAI Certificate with the provided `subject` and `subject_key_id` must exist in the ledger.

#### REMOVE_PAI

**Status: Implemented**

This transaction completely removes the given PAI (intermediate certificate) from both the approved and revoked certificates list.

PAA (self-signed root certificate) can not be removed this way.  

- Who can send: Vendor account
  - the sender's VID must match the VID of the removing certificate's owner.
- Parameters:
  - subject: `string`  - certificates's `Subject` is base64 encoded subject DER sequence bytes
  - subject_key_id: `string`  - certificates's `Subject Key Id` in hex string format, e.g: `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`
  - serial-number: `optional(string)`  - certificate's serial number.
- CLI command:
  - `dcld tx pki remove-x509-cert --subject=<base64 string> --subject-key-id=<hex string> --from=<account>`
- Validation:
  - a PAI Certificate with the provided `subject` and `subject_key_id` must exist in the ledger.

#### GET_DA_CERT

**Status: Implemented**

Gets a DA certificate by the given subject and subject key ID attributes. This query works for all types of DA certificates (PAA, PAI).
Revoked certificates are not returned.
Use [GET_REVOKED_DA_CERT](#get_revoked_da_cert) to get a revoked certificate.

- Parameters:
  - subject: `string`  - certificates's `Subject` is base64 encoded subject DER sequence bytes
  - subject_key_id: `string`  - certificates's `Subject Key Id` in hex string format, e.g: `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`
- CLI command:
  - `dcld query pki x509-cert --subject=<base64 string> --subject-key-id=<hex string>`
- REST API:
  - GET `/dcl/pki/certificates/{subject}/{subject_key_id}`

#### GET_REVOKED_DA_CERT

**Status: Implemented**

Gets a revoked DA certificate by the given subject and subject key ID attributes. This query works for all types of DA certificates (PAA, PAI).

Revocation works as a soft-delete, meaning that the certificates are not entirely removed but moved from the approved list to the revoked list.
If a Revocation Distribution Point (such as RFC5280 Certificate Revocation List) published to the ledger needs to be queried, please use [GET_PKI_REVOCATION_DISTRIBUTION_POINT](#get_pki_revocation_distribution_point).

- Parameters:
  - subject: `string`  - certificates's `Subject` is base64 encoded subject DER sequence bytes
  - subject_key_id: `string`  - certificates's `Subject Key Id` in hex string format, e.g: `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`
- CLI command:
  - `dcld query pki revoked-x509-cert --subject=<base64 string> --subject-key-id=<hex string>`
- REST API:
  - GET `/dcl/pki/revoked-certificates/{subject}/{subject_key_id}`

#### GET_DA_CERTS_BY_SKID

**Status: Implemented**

Gets all DA certificates by the given subject key ID attribute. This query works for all types of DA certificates (PAA, PAI).

Revoked certificates are not returned.
Use [GET_ALL_REVOKED_DA_CERTS](#get_all_revoked_da_cert) to get a revoked certificate.

- Parameters:
  - subject_key_id: `string`  - certificates's `Subject Key Id` in hex string format, e.g: `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`
- CLI command:
  - `dcld query pki x509-cert --subject-key-id=<hex string>`
- REST API:
  - GET `/dcl/pki/certificates?subjectKeyId={subjectKeyId}`

#### GET_DA_CERTS_BY_SUBJECT

**Status: Implemented**

Gets all DA certificates associated with a subject. This query works for all types of DA certificates (PAA, PAI).

Revoked certificates are not returned.
Use [GET_ALL_REVOKED_DA_CERTS](#get_all_revoked_da_certs) to get a list of all revoked certificates.

- Parameters:
  - subject: `string`  - certificates's `Subject` is base64 encoded subject DER sequence bytes
- CLI command:
  - `dcld query pki all-subject-x509-certs --subject=<base64 string>`
- REST API:
  - GET `/dcl/pki/certificates/{subject}`

#### GET_ALL_DA_CERTS

**Status: Implemented**

Gets all DA certificates. This query works for all types of DA certificates (PAA, PAI).

Revoked certificates are not returned.
Use [GET_ALL_REVOKED_DA_CERTS](#get_all_revoked_da_certs) to get a list of all revoked certificates.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query pki all-x509-certs`
- REST API:
  - GET `/dcl/pki/certificates`

#### GET_ALL_REVOKED_DA_CERTS

**Status: Implemented**

Gets all revoked DA certificates. This query works for all types of DA certificates (PAA, PAI).

Revocation works as a soft-delete, meaning that the certificates are not entirely removed but moved from the approved list to the revoked list.
If a Revocation Distribution Point (such as RFC5280 Certificate Revocation List) published to the ledger needs to be queried, please use [GET_PKI_REVOCATION_DISTRIBUTION_POINT](#get_pki_revocation_distribution_point).

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query pki all-revoked-x509-certs`
- REST API:
  - GET `/dcl/pki/revoked-certificates` 

#### GET_PKI_REVOCATION_DISTRIBUTION_POINT

**Status: Implemented**

Gets a revocation distribution point (such as RFC5280 Certificate Revocation List) identified by (VendorID, Label, IssuerSubjectKeyID) unique combination.
Use [GET_ALL_PKI_REVOCATION_DISTRIBUTION_POINT](#get_all_pki_revocation_distribution_point) to get a list of all revocation distribution points.

- Parameters:
  - vid: `uint16` -  Vendor ID (positive non-zero)
  - label: `string` -  A label to disambiguate multiple revocation information partitions of a particular issuer.
  - issuerSubjectKeyID: `string` - Uniquely identifies the PAA or PAI for which this revocation distribution point is provided. Must consist of even number of uppercase hexadecimal characters ([0-9A-F]), with no whitespace and no non-hexadecimal characters., e.g: `5A880E6C3653D07FB08971A3F473790930E62BDB`.
- CLI command:
  - `dcld query pki revocation-point --vid=<uint16> --label=<string> --issuer-subject-key-id=<string>`
- REST API:
  - GET `/dcl/pki/revocation-points/{issuerSubjectKeyID}/{vid}/{label}`

#### GET_PKI_REVOCATION_DISTRIBUTION_POINTS_BY_SUBJECT_KEY_ID

**Status: Implemented**

Gets a list of revocation distribution point (such as RFC5280 Certificate Revocation List) identified by IssuerSubjectKeyID.

- Parameters:
  - issuerSubjectKeyID: `string` - Uniquely identifies the PAA or PAI for which this revocation distribution point is provided. Must consist of even number of uppercase hexadecimal characters ([0-9A-F]), with no whitespace and no non-hexadecimal characters., e.g: `5A880E6C3653D07FB08971A3F473790930E62BDB`.
- CLI command:
  - `dcld query pki revocation-points --issuer-subject-key-id=<string>`
- REST API:
  - GET `/dcl/pki/revocation-points/{issuerSubjectKeyID}`

#### GET_ALL_PKI_REVOCATION_DISTRIBUTION_POINT

**Status: Implemented**

Gets a list of all revocation distribution points (such as RFC5280 Certificate Revocation List).

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters
- CLI command:
  - `dcld query pki all-revocation-points`
- REST API:
  - GET `/dcl/pki/revocation-points`

#### GET_PROPOSED_PAA

**Status: Implemented**

Gets a proposed but not approved PAA certificate with the given subject and subject key ID attributes.

- Parameters:
  - subject: `string`  - certificates's `Subject` is base64 encoded subject DER sequence bytes
  - subject_key_id: `string`  - certificates's `Subject Key Id` in hex string format, e.g: `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`
- CLI command:
  - `dcld query pki proposed-x509-root-cert --subject=<base64 string> --subject-key-id=<hex string>`
- REST API:
  - GET `/dcl/pki/proposed-certificates/{subject}/{subject_key_id}`

#### GET_REJECTED_PAA

**Status: Implemented**

Get a rejected PAA certificate with the given subject and subject key ID attributes.

- Parameters:
  - subject: `string` - certificates's `Subject` is base64 encoded subject DER sequence bytes
  - subject_key_id: `string`  - certificates's `Subject Key Id` in hex string format, e.g: `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`
- CLI command:
  - `dcld query pki rejected-x509-root-cert --subject=<base64 string> --subject-key-id=<hex string>`
- REST API:
  - GET `/dcl/pki/rejected-certificates/{subject}/{subject_key_id}`

#### GET_PROPOSED_PAA_TO_REVOKE

**Status: Implemented**

Gets a proposed but not approved PAA certificate to be revoked.

Revocation here just means removing it from the ledger.
If a Revocation Distribution Point (such as RFC5280 Certificate Revocation List) published to the ledger needs to be queried, please use [GET_PKI_REVOCATION_DISTRIBUTION_POINT](#get_pki_revocation_distribution_point).

- Parameters:
  - subject: `string`  - certificates's `Subject` is base64 encoded subject DER sequence bytes
  - subject_key_id: `string`  - certificates's `Subject Key Id` in hex string format, e.g: `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`
  - serial-number: `optional(string)`  - certificate's serial number
- CLI command:
  - `dcld query pki proposed-x509-root-cert-to-revoke --subject=<base64 string> --subject-key-id=<hex string>`
- REST API:
  - GET `/dcl/pki/proposed-revocation-certificates/{subject}/{subject_key_id}?serialnumber={serialnumber}`

#### GET_ALL_PAA

**Status: Implemented**

Gets all approved PAA certificates. Revoked certificates are not returned.
Use [GET_ALL_REVOKED_PAA](#get_all_revoked_paa) to get a list of all revoked PAA certificates.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query pki all-x509-root-certs`
- REST API:
  - GET `/dcl/pki/root-certificates`

#### GET_ALL_REVOKED_PAA

**Status: Implemented**

Gets all revoked PAA certificates.

Revocation here just means removing it from the ledger.
If a Revocation Distribution Point (such as RFC5280 Certificate Revocation List) published to the ledger needs to be queried, please use [GET_PKI_REVOCATION_DISTRIBUTION_POINT](#get_pki_revocation_distribution_point).

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query pki all-revoked-x509-root-certs`
- REST API:
  - GET `/dcl/pki/revoked-root-certificates`

#### GET_ALL_PROPOSED_PAA

**Status: Implemented**

Gets all proposed but not approved root certificates.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query pki all-proposed-x509-root-certs`
- REST API:
  - GET `dcl/pki/proposed-certificates`

#### GET_ALL_REJECTED_PAA

 **Status: Implemented**

Gets all rejected root certificates.

Shoudl be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query pki all-rejected-x509-root-certs`
- REST API:
  - GET `dcl/pki/rejected-certificates`

#### GET_ALL_PROPOSED_PAA_TO_REVOKE

**Status: Implemented**

Gets all proposed but not approved root certificates to be revoked.

Revocation here just means removing it from the ledger.
If a Revocation Distribution Point (such as RFC5280 Certificate Revocation List) published to the ledger needs to be queried, please use [GET_PKI_REVOCATION_DISTRIBUTION_POINT](#get_pki_revocation_distribution_point).

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query pki all-proposed-x509-root-certs-to-revoke`
- REST API:
  - GET `/dcl/pki/proposed-revocation-certificates`

### E2E (NOC): RCAC, ICAC

#### ADD_NOC_ROOT (RCAC)

**Status: Implemented**

This transaction adds a NOC root certificate (RCAC) owned by the Vendor.

- Who can send
  - Vendor account
- Parameters:
  - cert: `string` - The NOC Root Certificate (RCAC), encoded in X.509v3 PEM format. Can be a PEM string or a file path.
  - schemaVersion: `optional(uint16)` - Certificate's schema version to support backward/forward compatability. Should be equal to 0 (default 0)
- In State:
  - `pki/AllCertificates/value/<Subject>/<SubjectKeyID>`
  - `pki/NocCertificates/value/<Subject>/<SubjectKeyID>`
  - `pki/NocRootCertificates/value/<VID>`
  - `pki/NocCertificatesBySubject/value/<Subject>`
  - `pki/NocCertificatesBySubjectKeyId/value/<SubjectKeyID>`
  - `pki/NocCertificatesByVidAndSkid/value/<VID>/<SubjectKeyID>`
- CLI Command:
  - `dcld tx pki add-noc-x509-root-cert --certificate=<string-or-path> --from=<account>`
- Validation:
  - the provided certificate must be a root certificate (RCAC):
    - `Issuer` == `Subject`
    - `Authority Key Identifier` == `Subject Key Identifier`
  - no existing certificate with the same `<Certificate's Issuer>:<Certificate's Serial Number>` combination.
  - if certificates with the same `<Certificate's Subject>:<Certificate's Subject Key ID>` combination already exist:
    - the existing certificate must be NOC root certificate (RCAC)
    - the sender's VID must match the `vid` field of the existing certificates.
  - the signature (self-signature) and expiration date must be valid.

#### REVOKE_NOC_ROOT (RCAC)

**Status: Implemented**

This transaction revokes a NOC root certificate (RCAC) owned by the Vendor.
Revoked NOC root certificates (RCACs) can be re-added using the [ADD_NOC_ROOT](#add_noc_root-rcac) transaction.

Revocation works as a soft-delete, meaning that the certificates are not entirely removed but moved from the approved list to the revoked list.
Revoked certificates can be retrieved by using the [GET_REVOKED_CERT](#get_revoked_noc_root-rcac) query.

- Who can send: Vendor account
  - Vid field associated with the corresponding NOC root certificate (RCAC) on the ledger must be equal to the Vendor account's VID.
- Parameters:
  - subject: `string` - base64 encoded subject DER sequence bytes of the certificate.
  - subject_key_id: `string` - certificate's `Subject Key Id` in hex string format, e.g., `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`.
  - serial_number: `optional(string)` - certificate's serial number. If not provided, the transaction will revoke all certificates that match the given `subject` and `subject_key_id` combination.
  - revoke-child: `optional(bool)` - if true, then all certificates in the chain signed by the revoked certificate (intermediate, leaf) are revoked as well. If false, only the current root cert is revoked (default: false).
  - info: `optional(string)` - information/notes for the revocation. Can contain up to 4096 characters.
  - time: `optional(int64)` - revocation time (number of nanoseconds elapsed since January 1, 1970 UTC). This field cannot be specified using a CLI command and will use the current time by default.
- In State:
  - `pki/RevokedNocRootCertificates/value/<Certificate's Subject>/<Certificate's Subject Key ID>`
- CLI command:
  - `dcld tx pki revoke-noc-x509-root-cert --subject=<base64 string> --subject-key-id=<hex string> --serial-number=<string> --info=<string> --time=<int64> --revoke-child=<bool> --from=<account>`
- Validation:
  - a NOC Root Certificate (RCAC) with the provided `subject` and `subject_key_id` must exist in the ledger.

#### REMOVE_NOC_ROOT (RCAC)

**Status: Implemented**

This transaction completely removes the given NOC root certificate (RCAC) owned by the Vendor from the ledger.
Removed NOC root certificates (RCACs) can be re-added using the [ADD_NOC_ROOT](#add_noc_root-rcac) transaction.

- Who can send: Vendor account
  - Vid field associated with the corresponding NOC certificate on the ledger must be equal to the Vendor account's VID.
- Validation:
  - a NOC Root Certificate (RCAC) with the provided `subject` and `subject_key_id` must exist in the ledger.
- Parameters:
  - subject: `string` - base64 encoded subject DER sequence bytes of the certificate.
  - subject_key_id: `string` - certificate's `Subject Key Id` in hex string format, e.g., `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`.
  - serial_number: `optional(string)` - certificate's serial number. If not provided, the transaction will remove all certificates that match the given `subject` and `subject_key_id` combination.
- CLI command:
  - `dcld tx pki remove-noc-x509-root-cert --subject=<base64 string> --subject-key-id=<hex string> --from=<account>`

#### ADD_NOC_ICA (ICAC)

**Status: Implemented**

This transaction adds a NOC ICA certificate (ICAC) owned by the Vendor signed by a chain of certificates which must be
already present on the ledger.

- Who can send: Vendor account
- Validation:
  - the provided certificate must be a non-root certificate:
    - `Issuer` != `Subject`
    - `Authority Key Identifier` != `Subject Key Identifier`
  - the root certificate must be a NOC certificate and added by the same vendor
    - `isNoc` field of the root certificate must be set to true
    - `VID of root certificate` == `VID of account`
  - no existing certificate with the same `<Certificate's Issuer>:<Certificate's Serial Number>` combination.
  - if certificates with the same `<Certificate's Subject>:<Certificate's Subject Key ID>` combination already exist:
    - the existing certificate must be NOC non-root certificate
    - the sender's VID must match the vid field of the existing certificates.
  - the signature and expiration date must be valid.
- Parameters:
  - cert: `string` - The NOC non-root Certificate, encoded in X.509v3 PEM format. Can be a PEM string or a file path.
  - certificate-schema-version: `optional(uint16)` - Certificate's schema version to support backward/forward compatability(default 0)
- In State:
  - `pki/AllCertificates/value/<Subject>/<SubjectKeyID>`
  - `pki/NocCertificates/value/<Subject>/<SubjectKeyID>`
  - `pki/NocIcaCertificates/value/<VID>`
  - `pki/NocCertificatesBySubject/value/<Subject>`
  - `pki/NocCertificatesBySubjectKeyID/value/<SubjectKeyID>`
  - `pki/NocCertificatesByVidAndSkid/value/<VID>/<SubjectKeyID>`
  - `pki/ChildCertificates/value/<Certificate's Subject>/<Certificate's Subject Key ID>`
- CLI Command:
  - `dcld tx pki add-noc-x509-ica-cert --certificate=<string-or-path> --from=<account>`

#### REVOKE_NOC_ICA (ICAC)

**Status: Implemented**

This transaction revokes a NOC ICA certificate (ICAC) owned by the Vendor.
Revoked NOC ICA certificates (ICACs) can be re-added using the [ADD_NOC_ICA](#add_noc_ica-icac) transaction.

Revocation works as a soft-delete, meaning that the certificates are not entirely removed but moved from the approved list to the revoked list.
Revoked certificates can be retrieved by using the [GET_REVOKED_CERT](#get_revoked_noc_ica-icac) query.

- Who can send: Vendor account
  - Vid field associated with the corresponding NOC certificate on the ledger must be equal to the Vendor account's VID.
- Validation:
  - a NOC Certificate with the provided `subject` and `subject_key_id` must exist in the ledger.
- Parameters:
  - subject: `string` - base64 encoded subject DER sequence bytes of the certificate.
  - subject_key_id: `string` - certificate's `Subject Key Id` in hex string format, e.g., `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`.
  - serial_number: `optional(string)` - certificate's serial number. If not provided, the transaction will revoke all certificates that match the given `subject` and `subject_key_id` combination.
  - revoke-child: `optional(bool)` - if true, then all certificates in the chain signed by the revoked certificate (leaf) are revoked as well. If false, only the current cert is revoked (default: false).
  - info: `optional(string)` - information/notes for the revocation. Can contain up to 4096 characters.
  - time: `optional(int64)` - revocation time (number of nanoseconds elapsed since January 1, 1970 UTC). This field cannot be specified using a CLI command and will use the current time by default.
- In State:
  - `pki/RevokedCertificates/value/<Certificate's Subject>/<Certificate's Subject Key ID>`
- CLI command:
  - `dcld tx pki revoke-noc-x509-ica-cert --subject=<base64 string> --subject-key-id=<hex string> --serial-number=<string> --info=<string> --time=<int64> --revoke-child=<bool> --from=<account>`

#### REMOVE_NOC_ICA (ICAC)

**Status: Implemented**

This transaction completely removes the given NOC ICA (ICAC) owned by the Vendor from the ledger.
Removed NOC ICA certificates (ICACs) can be re-added using the [ADD_NOC_ICA](#add_noc_ica-icac) transaction.

- Who can send: Vendor account
  - Vid field associated with the corresponding NOC certificate on the ledger must be equal to the Vendor account's VID.
- Validation:
  - a NOC ICA Certificate (ICAC) with the provided `subject` and `subject_key_id` must exist in the ledger.
- Parameters:
  - subject: `string` - base64 encoded subject DER sequence bytes of the certificate.
  - subject_key_id: `string` - certificate's `Subject Key Id` in hex string format, e.g., `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`.
  - serial_number: `optional(string)` - certificate's serial number. If not provided, the transaction will remove all certificates that match the given `subject` and `subject_key_id` combination.
- CLI command:
  - `dcld tx pki remove-noc-x509-ica-cert --subject=<base64 string> --subject-key-id=<hex string> --from=<account>`

#### GET_NOC_CERT

**Status: Implemented**

Gets a NOC certificate by the given subject and subject key ID attributes. This query works for all types of Noc certificates (NOC_ROOT, NOC_ICA).
Revoked certificates are not returned.
Use [GET_REVOKED_ROOT_ICA](#get_revoked_noc_root-rcac) to get a revoked root certificate.
Use [GET_REVOKED_NOC_ICA](#get_revoked_noc_ica-icac) to get a revoked ica certificate.

- Parameters:
  - subject: `string`  - certificate's `Subject` is base64 encoded subject DER sequence bytes
  - subject_key_id: `string`  - certificate's `Subject Key Id` in hex string format, e.g: `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`
- CLI command:
  - `dcld query pki noc-x509-cert --subject=<base64 string> --subject-key-id=<hex string>`
- REST API:
  - GET `/dcl/pki/noc-certificates/{subject}/{subject_key_id}`

#### GET_NOC_ROOT_BY_VID (RCACs)

**Status: Implemented**

Retrieve NOC root certificates (RCACs) associated with a specific VID.

Revoked NOC root certificates (RCACs) are not returned.
Use [GET_ALL_REVOKED_NOC_ROOT](#get_all_revoked_noc_root-rcacs) to get a list of all revoked NOC root certificates (RCACs).

- Who can send: Any account
- Parameters:
  - vid: `uint16` - Vendor ID (positive non-zero)
- CLI Command:
  - `dcld query pki noc-x509-root-certs --vid=<uint16>`
- REST API:
  - GET `/dcl/pki/noc-root-certificates/{vid}`

#### GET_NOC_BY_VID_AND_SKID (RCACs/ICACs)

**Status: Implemented**

Retrieve NOC (Root/ICA) certificates (RCACs/ICACs) associated with a specific VID and subject key ID.
This request also returns the Trust Quotient (TQ) value of the certificate

Revoked NOC certificates are not returned.
Use [GET_ALL_REVOKED_NOC_ROOT](#get_all_revoked_noc_root-rcacs) to get a list of all revoked NOC root certificates.
Use [GET_ALL_REVOKED_NOC_ICA](#get_all_revoked_noc_ica-icacs) to get a list of all revoked NOC ica certificates.

- Who can send: Any account
- Parameters:
  - vid: `uint16` - Vendor ID (positive non-zero)
  - subject_key_id: `string` - Certificate's `Subject Key Id` in hex string format, e.g., `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`
- CLI Command:
  - `dcld query pki noc-x509-certs --vid=<uint16> --subject-key-id=<hex string>`
- REST API:
  - GET `/dcl/pki/noc-certificates/{vid}/{subject_key_id}`

#### GET_NOC_ICA_BY_VID (ICACs)

**Status: Implemented**

Retrieve NOC ICA certificates (ICACs) associated with a specific VID.

Revoked certificates are not returned.
Use [GET_ALL_REVOKED_CERT](#get_all_revoked_certs) to get a list of all revoked certificates.

- Who can send: Any account
- Parameters:
  - vid: `uint16` - Vendor ID (positive non-zero)
- CLI Command:
  - `dcld query pki noc-x509-ica-certs --vid=<uint16>`
- REST API:
  - GET `/dcl/pki/noc-ica-certificates/{vid}`

#### GET_NOC_CERTS_BY_SUBJECT

**Status: Implemented**

Gets all NOC certificates associated with a subject. This query works for both types of certificates (NOC_ROOT, NOC_ICA).

Revoked certificates are not returned.
Use [GET_ALL_REVOKED_NOC_ROOT](#get_all_revoked_noc_root-rcacs) to get a list of all revoked NOC root certificates.
Use [GET_ALL_REVOKED_NOC_ICA](#get_all_revoked_noc_ica-icacs) to get a list of all revoked NOC ica certificates.

- Parameters:
  - subject: `string`  - certificate's `Subject` is base64 encoded subject DER sequence bytes
- CLI command:
  - `dcld query pki all-noc-subject-x509-certs --subject=<base64 string>`
- REST API:
  - GET `/dcl/pki/noc-certificates/{subject}`

#### GET_REVOKED_NOC_ROOT (RCAC)

**Status: Implemented**

Gets a revoked NOC root certificate (RCAC) by the given subject and subject key ID attributes.

Revocation works as a soft-delete, meaning that the certificates are not entirely removed but moved from the approved list to the revoked list.

- Parameters:
  - subject: `string` - Base64 encoded subject DER sequence bytes of the certificate.
  - subject_key_id: `string` - Certificate's `Subject Key Id` in hex string format, e.g., `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`.
- CLI command:
  - `dcld query pki revoked-noc-x509-root-cert --subject=<base64 string> --subject-key-id=<hex string>`
- REST API:
  - GET `/dcl/pki/revoked-noc-root-certificates/{subject}/{subject_key_id}`

#### GET_REVOKED_NOC_ICA (ICAC)

**Status: Implemented**

Gets a revoked NOC ica certificate (ICAC) by the given subject and subject key ID attributes.

Revocation works as a soft-delete, meaning that the certificates are not entirely removed but moved from the approved list to the revoked list.

- Parameters:
  - subject: `string` - Base64 encoded subject DER sequence bytes of the certificate.
  - subject_key_id: `string` - Certificate's `Subject Key Id` in hex string format, e.g., `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`.
- CLI command:
  - `dcld query pki revoked-noc-x509-ica-cert --subject=<base64 string> --subject-key-id=<hex string>`
- REST API:
  - GET `/dcl/pki/revoked-noc-ica-certificates/{subject}/{subject_key_id}`

#### GET_ALL_NOC (RCACs/ICACs)

**Status: Implemented**

Retrieve a list of all of NOC certificates (RCACs of ICACs).

Revoked NOC certificates (RCACs and ICACs) are not returned.
Use [GET_ALL_REVOKED_NOC_ROOT](#get_all_revoked_noc_ica-icacs) to get a list of all revoked NOC root certificates (RCACs).
Use [GET_ALL_REVOKED_NOC_ICA](#get_all_revoked_noc_ica-icacs) to get a list of all revoked NOC ica certificates (ICACs).

- Who can send: Any account
- Parameters:
  - Common pagination parameters
- CLI Command:
  - `dcld query pki all-noc-x509-certs`
- REST API:
  - GET `/dcl/pki/noc-certificates`

#### GET_ALL_NOC_ROOT (RCACs)

**Status: Implemented**

Retrieve a list of all of NOC root certificates (RCACs).

Revoked NOC root certificates (RCACs) are not returned.
Use [GET_ALL_REVOKED_NOC_ROOT](#get_all_revoked_noc_root-rcacs) to get a list of all revoked NOC root certificates (RCACs).

- Who can send: Any account
- Parameters:
  - Common pagination parameters
- CLI Command:
  - `dcld query pki all-noc-x509-root-certs`
- REST API:
  - GET `/dcl/pki/noc-root-certificates`

#### GET_ALL_NOC_ICA (ICACs)

**Status: Implemented**

Retrieve a list of all of NOC ICA certificates (ICACs).

Revoked certificates are not returned.
Use [GET_ALL_REVOKED_NOC_ICA](#get_all_revoked_noc_ica-icacs) to get a list of all revoked certificates.

- Who can send: Any account
- Parameters:
  - Common pagination parameters
- CLI Command:
  - `dcld query pki all-noc-x509-ica-certs`
- REST API:
  - GET `/dcl/pki/noc-ica-certificates`

#### GET_ALL_REVOKED_NOC_ROOT (RCACs)

Gets all revoked NOC root certificates (RCACs).

Revocation works as a soft-delete, meaning that the certificates are not entirely removed but moved from the approved list to the revoked list.

- Who can send: Any account
- Parameters:
  - Common pagination parameters
- CLI command:
  - `dcld query pki all-revoked-noc-x509-root-certs`
- REST API:
  - GET `/dcl/pki/revoked-noc-root-certificates`

#### GET_ALL_REVOKED_NOC_ICA (ICACs)

Gets all revoked NOC ica certificates (ICACs).

Revocation works as a soft-delete, meaning that the certificates are not entirely removed but moved from the approved list to the revoked list.

- Who can send: Any account
- Parameters:
  - Common pagination parameters
- CLI command:
  - `dcld query pki all-revoked-noc-x509-ica-certs`
- REST API:
  - GET `/dcl/pki/revoked-noc-ica-certificates`

<!-- markdownlint-enable MD036 -->