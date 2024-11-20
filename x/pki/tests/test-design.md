### [Add DA Root](./handler_add_paa_cert_test.go)

* Propose adding of DA root certificate: 
  * Indexes to check: 
    * Present:
      * `ProposedCertificate`
      * `UniqueCertificate`
    * Missing:
      * `RejectedCertificate`
      * `All Certificates`: Subject+SKID, SKID, Subject
      * `DA Certificates`: Subject+SKID (approved), Subject+SKID (root), SKID, Subject
  * Tests:
    * `TestHandler_ProposeAddX509RootCert_ByTrustee`
* Propose add approve adding of DA root certificate:
  * Indexes:
    * Present:
      * `UniqueCertificate`
      * `All Certificates`: Subject+SKID, SKID, Subject
      * `DA Certificates`: Subject+SKID (approved), Subject+SKID (root), SKID, Subject
    * Missing:
      * `ProposedCertificate`
  * Tests:
    * `TestHandler_ApproveAddX509RootCert_ForEnoughApprovals`
    * `TestHandler_TwoThirdApprovalsNeededForAddingRootCertification`
    * `TestHandler_ApproveX509RootCert_FourApprovalsAreNeeded_FiveTrustees`

### [Add DA Intermediate](./handler_add_pai_cert_test.go)

* Add DA intermediate certificate:
  * Indexes to check:
    * Present:
      * `UniqueCertificate`
      * `All Certificates`: Subject+SKID, SKID, Subject
      * `DA Certificates`: Subject+SKID (approved), SKID, Subject
      * `ChildCertificates`: for parent
    * Missing:
      * `ProposedCertificate`
  * Tests:
    * `TestHandler_AddX509Cert`

### [Revoke DA Root](./handler_revoke_paa_cert_test.go)

* Propose revocation of DA root certificate:
  * Indexes to check:
    * Present:
      * `ProposedCertificateRevocation`
      * `UniqueCertificate`
      * `All Certificates`: Subject+SKID, SKID, Subject
      * `DA Certificates`: Subject+SKID (approved), Subject+SKID (root), SKID, Subject
    * Missing:
      * `RevokedCertificates`
  * Tests:
    * `TestHandler_ProposeRevokeX509RootCert_ByTrusteeOwner`
* Propose and approve revocation of DA root certificate:
  * Indexes:
    * Present:
      * `RevokedCertificates`
      * `UniqueCertificate`
    * Missing:
      * `ProposedCertificateRevocation`
      * `All Certificates`: Subject+SKID, SKID, Subject
      * `DA Certificates`: Subject+SKID (approved), Subject+SKID (root), SKID, Subject
  * Tests:
    * `TestHandler_TwoThirdApprovalsNeededForRevokingRootCertification`

### [Revoke DA Intermediate](./handler_revoke_pai_cert_test.go)

* Revoke DA intermediate certificate:
  * Indexes to check:
    * Present:
      * `RevokedCertificates`
      * `UniqueCertificate`
      * Root - stays approved
    * Missing:
      * `ProposedCertificateRevocation`
      * `All Certificates`: Subject+SKID, SKID, Subject
      * `DA Certificates`: Subject+SKID (approved), SKID, Subject
      * `ChildCertificates`: for parent
  * Tests:
    * `TestHandler_RevokeX509Cert`

### [Remove DA Intermediate](./handler_remove_pai_cert_test.go)

* Remove DA intermediate certificate:
  * Indexes to check:
    * Present:
      * -
    * Missing:
      * `RevokedCertificates`
      * `UniqueCertificate`
      * `All Certificates`: Subject+SKID, SKID, Subject
      * `DA Certificates`: Subject+SKID (approved), SKID, Subject
      * `ChildCertificates`: for parent
  * Tests:
    * `TestHandler_RemoveX509Cert_BySubjectAndSKID`

### [Add Noc Root](./handler_add_noc_root_cert_test.go)

* Add Noc root certificate:
  * Indexes to check:
    * Present:
      * `UniqueCertificate`
      * `All Certificates`: Subject+SKID, SKID, Subject
      * `Noc Certificates`: Subject+SKID, SKID, Subject, VID (root), VID+SKID
    * Missing:
      * -
  * Tests:
    * `TestHandler_AddNocX509Cert_AddNewRoot`

### [Add Noc Intermediate](./handler_add_noc_ica_cert_test.go)

* Add Noc intermediate certificate:
  * Indexes to check:
    * Present:
      * `UniqueCertificate`
      * `All Certificates`: Subject+SKID, SKID, Subject
      * `Noc Certificates`: Subject+SKID, SKID, Subject, VID (ica), VID+SKID
      * `ChildCertificates`: for parent
    * Missing:
      * -
  * Tests:
    * `TestHandler_AddNocX509Cert_AddNewIca`

### [Revoke Noc Root](./handler_revoke_noc_root_cert_test.go)

* Revoke Noc root certificate:
  * Indexes:
    * Present:
      * `RevokedCertificates` (root)
      * `UniqueCertificate`
    * Missing:
      * `RevokedCertificates` (ica)
      * `All Certificates`: Subject+SKID, SKID, Subject
      * `Noc Certificates`: Subject+SKID, SKID, Subject, VID (root), VID+SKID
  * Tests:
    * `TestHandler_RevokeNocX509RootCert`

### [Revoke Noc Ica](./handler_revoke_noc_ica_cert_test.go)

* Revoke Noc ica certificate:
  * Indexes:
    * Present:
      * `RevokedCertificates` (ica)
      * `UniqueCertificate`
    * Missing:
      * `RevokedCertificates` (root)
      * `All Certificates`: Subject+SKID, SKID, Subject
      * `Noc Certificates`: Subject+SKID, SKID, Subject, VID (ica), VID+SKID
      * `ChildCertificates`: for parent
  * Tests:
    * `TestHandler_RevokeNocX509Cert`

### [Remove Noc Root](./handler_remove_noc_root_cert_test.go)

* Remove Noc root certificate by Subject/SKID:
  * Indexes to check:
    * Present:
      * -
    * Missing:
      * `RevokedCertificates` (root)
      * `UniqueCertificate`
      * `All Certificates`: Subject+SKID, SKID, Subject
      * `Noc Certificates`: Subject+SKID, SKID, Subject, VID (root), VID+SKID
  * Tests:
    * `TestHandler_RemoveNocX509RootCert`

### [Remove Noc Root](./handler_remove_noc_ica_cert_test.go)

* Remove Noc ica certificate by Subject/SKID:
  * Indexes to check:
    * Present:
      * -
    * Missing:
      * `RevokedCertificates` (ica)
      * `UniqueCertificate`
      * `All Certificates`: Subject+SKID, SKID, Subject
      * `Noc Certificates`: Subject+SKID, SKID, Subject, VID (ica), VID+SKID
      * `ChildCertificates`: for parent
  * Tests:
    * `TestHandler_RemoveNocX509IcaCert`
