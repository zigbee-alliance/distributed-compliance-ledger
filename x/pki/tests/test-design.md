## [Add DA Root](./handler_add_paa_cert_test.go)

### Propose adding of DA root certificate

Indexes to check:

* Present:
    * `ProposedCertificate`
    * `UniqueCertificate`
* Missing:
    * `RejectedCertificate`
    * `All Certificates`: Subject+SKID, SKID, Subject
    * `DA Certificates`: Subject+SKID (approved), Subject+SKID (root), SKID, Subject

Test cases:

* Positive:
    * Propose adding of DA root certificate: `TestHandler_ProposeAddDaRootCert`
    * Propose adding of previously rejected DA root certificate: ?
    * Propose adding of DA root certificate with same Subject/SKID as existing Approved certificate but different Serial
      Number: `TestHandler_ProposeAddX509RootCert_ForDifferentSerialNumber` (need to rewrite)
* Negative:
    * TBD

### Propose and approve adding of DA root certificate

Indexes:

* Present:
    * `UniqueCertificate`
    * `All Certificates`: Subject+SKID, SKID, Subject
    * `DA Certificates`: Subject+SKID (approved), Subject+SKID (root), SKID, Subject
* Missing:
    * `ProposedCertificate`

Test cases:

* Positive:
    * Add DA root certificate: `TestHandler_AddDaRootCert`,
      `TestHandler_AddDaRootCert_TwoThirdApprovalsNeeded`,
      `TestHandler_AddDaRootCert_FourApprovalsAreNeeded_FiveTrustees`
    * Add two DA root certificates with same SKID but different subject:
      `TestHandler_AddDaRootCerts_SameSubjectKeyIdButDifferentSubject`
* Negative:
    * TBD

### Propose and reject adding of DA root certificate

Indexes:

* Present:
    * `RejectedCertificate`
* Missing:
    * `ProposedCertificate`
    * `UniqueCertificate`
    * `All Certificates`: Subject+SKID, SKID, Subject
    * `DA Certificates`: Subject+SKID (approved), Subject+SKID (root), SKID, Subject

Test cases:

* Positive:
    * Propose add reject adding of DA root certificate: `TestHandler_RejectAddDaRootCert`,
* Negative:
    * TBD

## [Add DA Intermediate](./handler_add_pai_cert_test.go)

Indexes to check:

* Present:
    * `UniqueCertificate`
    * `All Certificates`: Subject+SKID, SKID, Subject
    * `DA Certificates`: Subject+SKID (approved), SKID, Subject
    * `ChildCertificates`: for parent
* Missing:
    * `ProposedCertificate`

Test cases:

* Positive:
    * Add DA intermediate certificate: `TestHandler_AddDaIntermediateCert`,
      `TestHandler_AddDaIntermediateCert_VidScoped`
    * Add two DA intermediate certificates with same Subject/SKID but different Serial Number:
      `TestHandler_AddDaIntermediateCert_SameSubjectAndSkid_DifferentSerialNumber`
    * Add tree of DA certificates (root, intermediate, leaf):
      `TestHandler_AddDaCert_ForTree`
    * Add intermediate DA certificate but other Vendor with the same VID:
      `TestHandler_AddDaIntermediateCert_ByNotOwnerButSameVendor`
* Negative:
    * TBD

## [Revoke DA Root](./handler_revoke_paa_cert_test.go)

### Propose revocation of DA root certificate

Indexes to check:

* Present:
    * `ProposedCertificateRevocation`
    * `UniqueCertificate`
    * `All Certificates`: Subject+SKID, SKID, Subject
    * `DA Certificates`: Subject+SKID (approved), Subject+SKID (root), SKID, Subject
* Missing:
    * `RevokedCertificates`
    * `RevokedRootCertificates`

Test cases:

* Positive:
    * Propose revocation of DA root certificate: `TestHandler_ProposeRevokeDaRootCert`
    * Propose revocation of DA root certificate by not owner: `TestHandler_ProposeRevokeDaRootCert_ByTrusteeNotOwner`
* Negative:
    * TBD

### Propose and approve revocation of DA root certificate

Indexes:

* Present:
    * `RevokedCertificates`
    * `RevokedRootCertificates`
    * `UniqueCertificate`
* Missing:
    * `ProposedCertificateRevocation`
    * `All Certificates`: Subject+SKID, SKID, Subject
    * `DA Certificates`: Subject+SKID (approved), Subject+SKID (root), SKID, Subject

Test cases:

* Positive:
    * Propose and approve revocation of DA root certificate by Subject/SKID: `TestHandler_RevokeDaRootCert`,
      `TestHandler_RevokeDaRootCert_TwoThirdApprovalsNeeded`
    * Revoke DA root certificate by Subject/SKID when two certs with the same SKID exist:
      `TestHandler_RevokeDaRootCert_BySubjectAndSkid_WhenTwoCertsWithSameSkidExist`
* Negative:
    * TBD

## [Revoke DA Intermediate](./handler_revoke_pai_cert_test.go)

Indexes to check:

* Present:
    * `RevokedCertificates`
    * `UniqueCertificate`
    * Root - stays approved
* Missing:
    * `ProposedCertificateRevocation`
    * `All Certificates`: Subject+SKID, SKID, Subject
    * `DA Certificates`: Subject+SKID (approved), SKID, Subject
    * `ChildCertificates`: for parent

Test cases:

* Positive:
    * Revoke DA intermediate certificate: `TestHandler_RevokeDaIntermediateCert`
* Negative:
    * TBD

## [Remove DA Intermediate](./handler_remove_pai_cert_test.go)

Indexes to check:

* Present:
    * no
* Missing:
    * `RevokedCertificates`
    * `UniqueCertificate`
    * `All Certificates`: Subject+SKID, SKID, Subject
    * `DA Certificates`: Subject+SKID (approved), SKID, Subject
    * `ChildCertificates`: for parent

Test cases:

* Positive:
    * Remove DA intermediate certificate: `TestHandler_RemoveDaIntermediateCert_BySubjectAndSKID`
* Negative:
    * TBD

## [Add Noc Root](./handler_add_noc_root_cert_test.go)

Indexes to check:

* Present:
    * `UniqueCertificate`
    * `All Certificates`: Subject+SKID, SKID, Subject
    * `Noc Certificates`: Subject+SKID, SKID, Subject, VID (root), VID+SKID
* Missing:
    * no

Test cases:

* Positive:
    * Add Noc root certificate: `TestHandler_AddNocRootCert`
    * Add two Noc root certificates with same subject/skid but different serial number:
      `TestHandler_AddNocRootCert_SameSubjectAndSkid_DifferentSerialNumber`
* Negative:
    * TBD

## [Add Noc Intermediate](./handler_add_noc_ica_cert_test.go)

Indexes to check:

* Present:
    * `UniqueCertificate`
    * `All Certificates`: Subject+SKID, SKID, Subject
    * `Noc Certificates`: Subject+SKID, SKID, Subject, VID (ica), VID+SKID
    * `ChildCertificates`: for parent
* Missing:
    * no

Test cases:

* Positive:
    * Add Noc intermediate certificate: `TestHandler_AddNocIntermediateCert`
    * Add two Noc intermediate certificates with same subject/skid but different serial number:
      `TestHandler_AddNocIntermediateCert_SameSubjectAndSkid_DifferentSerialNumber`
* Negative:
    * TBD

## [Revoke Noc Root](./handler_revoke_noc_root_cert_test.go)

Indexes:

* Present:
    * `RevokedCertificates` (root)
    * `UniqueCertificate`
* Missing:
    * `RevokedCertificates` (ica)
    * `All Certificates`: Subject+SKID, SKID, Subject
        * `Noc Certificates`: Subject+SKID, SKID, Subject, VID (root), VID+SKID

* Positive:
    * Revoke Noc root certificate: `TestHandler_RevokeNoRootCert`
* Negative:
    * TBD

## [Revoke Noc Ica](./handler_revoke_noc_ica_cert_test.go)

Indexes:

* Present:
    * `RevokedCertificates` (ica)
    * `UniqueCertificate`
* Missing:
    * `RevokedCertificates` (root)
    * `All Certificates`: Subject+SKID, SKID, Subject
    * `Noc Certificates`: Subject+SKID, SKID, Subject, VID (ica), VID+SKID
    * `ChildCertificates`: for parent

Test cases:

* Positive:
    * Revoke Noc ica certificate: `TestHandler_RevokeNocIntermediateCert`
* Negative:
    * TBD

## [Remove Noc Root](./handler_remove_noc_root_cert_test.go)

Indexes to check:

* Present:
    * no
* Missing:
    * `RevokedCertificates` (root)
    * `UniqueCertificate`
    * `All Certificates`: Subject+SKID, SKID, Subject
    * `Noc Certificates`: Subject+SKID, SKID, Subject, VID (root), VID+SKID

Test cases:

* Positive:
    * Remove Noc root certificate by Subject/SKID: `TestHandler_RemoveNocRootCert`
* Negative:
    * TBD

## [Remove Noc Intermediate](./handler_remove_noc_ica_cert_test.go)

Indexes to check:

* Present:
    * no
* Missing:
    * `RevokedCertificates` (ica)
    * `UniqueCertificate`
    * `All Certificates`: Subject+SKID, SKID, Subject
    * `Noc Certificates`: Subject+SKID, SKID, Subject, VID (ica), VID+SKID
    * `ChildCertificates`: for parent

Test cases:

* Positive:
    * Remove Noc ica certificate by Subject/SKID: `TestHandler_RemoveNocIntermediateCert`
* Negative:
    * TBD