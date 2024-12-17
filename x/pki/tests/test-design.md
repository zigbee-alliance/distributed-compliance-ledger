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
    * Propose single certificate: `TestHandler_ProposeAddDaRootCert`
    * Propose two certificates with same SKID but different Subject:
      `TestHandler_ProposeAddDaRootCert_SameSkidButDifferentSubject`
    * Propose certificate with Subject/SKID same as existing Approved certificate, but different SerialNumber:
      `TestHandler_ProposeAddDaRootCert_DifferentSerialNumber`
    * Propose adding of previously rejected certificate: `TestHandler_ProposeAddDaRootCert_PreviouslyRejected`
* Negative:
    * Propose by not Trustee: `TestHandler_ProposeAddDaRootCert_ByNotTrustee`
    * Propose invalid certificate: `TestHandler_ProposeAddDaRootCert_ForInvalidCertificate`
    * Propose with existing proposed certificate (Subject/SKID):
      `TestHandler_ProposeAddDaRootCert_ProposedCertificateAlreadyExists`
    * Propose with existing approved certificate (Subject/SKID/SerialNumber):
      `TestHandler_ProposeAddDaRootCert_CertificateAlreadyExists`
    * Propose not self-signed certificate: `TestHandler_ProposeAddDaRootCert_ForNonRootCertificate`
    * Propose not root certificate: `TestHandler_ProposeAddDaRootCert_ForNonRootCertificate`
    * Propose NOC root certificate: `TestHandler_ProposeAddDaRootCert_ForNocCertificate`
    * Propose with existing approved subject/SKID where signer is not owner of active:
      `TestHandler_ProposeAddDaRootCert_ForDifferentSigner`

### Approve adding of DA root certificate

Indexes:

* Present:
    * `UniqueCertificate`
    * `All Certificates`: Subject+SKID, SKID, Subject
    * `DA Certificates`: Subject+SKID (approved), Subject+SKID (root), SKID, Subject
* Missing:
    * `ProposedCertificate`

Test cases:

* Positive:
    * Add certificate: `TestHandler_AddDaRootCert`,
      `TestHandler_AddDaRootCert_TwoThirdApprovalsNeeded`,
      `TestHandler_AddDaRootCert_FourApprovalsAreNeeded_FiveTrustees`
    * Add two certificates with same SKID but different Subject:
      `TestHandler_AddDaRootCert_SameSkid_DifferentSubject`
    * Add two certificates with same Subject but different SKID:
    * Add two certificates with same Subject and SKID:
      `TestHandler_AddDaRootCert_SameSubjectAndSkid_DifferentSerialNumber`
    * Approve certificate for not enough approvals: `TestHandler_AddDaRootCert_TwoThirdApprovalsNeeded`
    * Approve certificate which was previously rejected by the current user:
      `TestHandler_ApproveAddDaRootCert_PreviouslyRejectedByCurrentTrustee`
* Negative:
    * Approve by not Trustee: `TestHandler_ApproveAddDaRootCert_ByNotTrustee`
    * Approve of non-existing proposed certificate: `TestHandler_ApproveAddDaRootCert_ForUnknownProposedCertificate`
    * Approve certificate already approved by the current user: `TestHandler_ApproveAddDaRootCert_Twice`

### Reject adding of DA root certificate

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
      `TestHandler_RejectX509RootCert_TwoRejectApprovalsAreNeeded_FiveTrustees`
    * Reject adding of DA root certificate for not enough rejects: `TestHandler_RejectAddDaRootCert`,
      `TestHandler_RejectX509RootCert_TwoRejectApprovalsAreNeeded_FiveTrustees`
    * Reject DA root certificate which was previously approved by the current user and certificate has other
      approval:
      `TestHandler_RejectAddDaRootCert_PreviouslyApprovedByCurrentTrustee_CertificateHasOtherApproval`
    * Reject DA root certificate which was previously approved by the current user and certificate has other
      rejects:
      `TestHandler_RejectAddDaRootCert_PreviouslyApprovedByCurrentTrustee_CertificateHasOtherReject`
    * Reject DA root certificate which was previously approved by the current user (and certificate does not have other
      rejects/approvals):
      `TestHandler_RejectAddDaRootCert_PreviouslyApprovedByCurrentTrustee_CertificateNotHasOtherApproval`
* Negative:
    * Reject by not Trustee: `TestHandler_RejectAddDaRootCert_ByNotTrustee`
    * Reject of non-existing proposed certificate: `TestHandler_RejectAddDaRootCert_ForUnknownProposedCertificate`
    * Reject certificate already rejected by the current user:
      `TestHandler_RejectX509RootCert_TwiceFromTheSameTrustee`

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
    * Add intermediate certificate: `TestHandler_AddDaIntermediateCert`,
      `TestHandler_AddDaIntermediateCert_VidScoped`
    * Add two certificates with same Subject/SKID but different SerialNumber:
      `TestHandler_AddDaIntermediateCert_SameSubjectAndSkid_DifferentSerialNumber`
    * Add two certificates with same Subject but different SKID: ?
    * Add two certificates with same SKID but different Subject: ?
    * Add tree of certificates (root, intermediate, leaf): `TestHandler_AddDaIntermediateCert_ForTree`
    * Add intermediate certificate but other Vendor with the same VID:
      `TestHandler_AddDaIntermediateCert_ByNotOwnerButSameVendor`
* Negative:
    * Add by not Vendor: `TestHandler_AddDaIntermediateCert_SenderNotVendor`
    * Add invalid certificate: `TestHandler_AddDaIntermediateCert_ForInvalidCertificate`
    * Add self-signed certificate: `TestHandler_AddDaIntermediateCert_ForRootCertificate`
    * Add with existing issuer/serial number: `TestHandler_AddDaIntermediateCert_ForDuplicate`
    * Add for root certificate: `TestHandler_AddDaIntermediateCert_ForRootCertificate`
    * Add for root NOC certificate: `TestHandler_AddDaIntermediateCert_RootIsNoc`
    * Add NOC certificate: TBD
    * Add with different VID: `TestHandler_AddDaIntermediateCert_ByOtherVendor`
    * Add with invalid chain: `TestHandler_AddDaIntermediateCert_ForAbsentDirectParentCert`

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
    * Propose revocation by Subject/SKID/SerialNumber - single certificate: `TestHandler_ProposeRevokeDaRootCert`
    * Propose revocation by Subject/SKID/SerialNumber - two certificates:
      `TestHandler_ProposeRevokeDaRootCert_TwoCertificates`
    * Propose revocation by Subject/SKID/SerialNumber - revoke child: `TestHandler_ProposeRevokeDaRootCert_RevokeChild`
    * Propose revocation by Subject/SKID/SerialNumber - keep child: `TestHandler_ProposeRevokeDaRootCert_KeepChild`
    * Propose revocation by other Vendor with the same VID: `TestHandler_ProposeRevokeDaRootCert_ByTrusteeNotOwner`
* Negative:
    * Propose revocation by not Trustee: `TestHandler_ProposeRevokeDaRootCert_ByNotTrustee`
    * Propose revocation of already proposed for revocation:
      `TestHandler_ProposeRevokeDaRootCert_ProposedRevocationAlreadyExists`
    * Propose revocation of not existing approved certificate (Subject/SKID):
      `TestHandler_ProposeRevokeDaRootCert_CertificateDoesNotExist`,
      `TestHandler_ProposeRevokeDaRootCert_ForProposedCertificate`
    * Propose revocation of not existing approved certificate (Subject/SKID + SerialNumber):
      `TestHandler_ProposeRevokeDaRootCert_CertificateDoesNotExistBySerialNumber`
    * Propose revocation of not root certificate: `TestHandler_ProposeRevokeDaRootCert_ForNonRootCertificate`

### Approve revocation of DA root certificate

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
    * Approve revocation DA root certificate when not enough approvals:
      `TestHandler_ApproveRevokeDaRootCert_NotEnoughApprovals`
    * Revoke by Subject/SKID: `TestHandler_RevokeDaRootCert_BySubjectAndSKID`,
      `TestHandler_RevokeDaRootCert_TwoThirdApprovalsNeeded`
    * Revoke by Subject/SKID/SerialNumber: `TestHandler_RevokeDaRootCert_BySerialNumber`
    * Revoke by Subject/SKID/SerialNumber - revoke child: `TestHandler_RevokeDaRootCert_RevokeChild`
    * Revoke by Subject/SKID/SerialNumber - keep child: `TestHandler_RevokeDaRootCert_KeepChild`
    * Revoke by Subject/SKID when two certs with the same SKID exist:
      `TestHandler_RevokeDaRootCert_BySubjectAndSkid_TwoCertificatesWithSameSkid`
    * Revoke by Subject/SKID when two certs with the same Subject exist: ?
* Negative:
    * Approve revocation by not Trustee: `TestHandler_ApproveRevokeDaRootCert_ByNotTrustee`
    * Approve revocation of not existing certificate (Subject/SKID):
      `TestHandler_ApproveRevokeDaRootCert_ProposedRevocationDoesNotExist`
    * Approve certificate revocation by not existing serial number (Subject/SKID + SerialNumber):
      `TestHandler_ApproveRevokeDaRootCert_BySerialNumber_ProposedRevocationDoesNotExist`
    * Approve certificate revocation twice by the same user: `TestHandler_ApproveRevokeDaRootCert_Twice`

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
    * Revoke by Subject/SKID: `TestHandler_RevokeDaIntermediateCert_BySubjectAndSKID`
    * Revoke by Subject/SKID/SerialNumber: `TestHandler_RevokeDaIntermediateCert_BySerialNumber`
    * Revoke by Subject/SKID - revoke child: `TestHandler_RevokeDaIntermediateCert_BySubjectAndSKID_RevokeChild`
    * Revoke by Subject/SKID/SerialNumber - revoke child:
      `TestHandler_RevokeDaIntermediateCert_BySerialNumber_RevokeChild`
    * Revoke by Subject/SKID - keep child: `TestHandler_RevokeDaIntermediateCert_BySubjectAndSKID_KeepChild`
    * Revoke by Subject/SKID/SerialNumber - keep child: `TestHandler_RevokeDaIntermediateCert_BySerialNumber_KeepChild`
    * Revoke by Subject/SKID - parent not affected: `TestHandler_RevokeDaIntermediateCert_BySubjectAndSKID_ParentExist`
    * Revoke by Subject/SKID/SerialNumber - parent not affected:
      `TestHandler_RevokeDaIntermediateCert_BySerialNumber_ParentExist`
    * Revoke by Subject/SKID - another certificate with same Subject exist: ?
    * Revoke by Subject/SKID - another certificate with same SKID exist: ?
    * Revoke by other Vendor with the same VID: `TestHandler_RevokeDaIntermediateCert_ByNotOwnerButSameVendor`
* Negative:
    * Revoke by not Vendor: `TestHandler_RevokeDaIntermediateCert_SenderNotVendor`
    * Revoke root certificate: `TestHandler_RevokeDaIntermediateCert_ForRootCertificate`
    * Revoke by Vendor with different VID: `TestHandler_RevokeDaIntermediateCert_ByVendorWithOtherVid`
    * Revoke not existing certificate (Subject/SKID): `TestHandler_RevokeDaIntermediateCert_CertificateDoesNotExist`
    * Revoke not existing certificate by SerialNumber (Subject/SKID + SerialNumber):
      `TestHandler_RevokeDaIntermediateCert_CertificateDoesNotExistBySerialNumber`

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
    * Remove by Subject/SKID: `TestHandler_RemoveDaIntermediateCert_BySubjectAndSKID`
    * Remove by Subject/SKID/SerialNumber: `TestHandler_RemoveDaIntermediateCert_BySerialNumber`
    * Remove by Subject/SKID - parent exist: `TestHandler_RemoveDaIntermediateCert_BySubjectAndSKID_ParentExist`
    * Remove by Subject/SKID/SerialNumber - parent exist:
      `TestHandler_RemoveDaIntermediateCert_BySerialNumber_ParentExist`
    * Remove by Subject/SKID - approved child exist:
      `TestHandler_RemoveDaIntermediateCert_BySubjectAndSKID_ApprovedChildExist`
    * Remove by Subject/SKID/SerialNumber - approved child exist:
      `TestHandler_RemoveDaIntermediateCert_BySerialNumber_ApprovedChildExist`
    * Remove by Subject/SKID - approved child exist:
      `TestHandler_RemoveDaIntermediateCert_BySubjectAndSKID_RevokedChildExist`
    * Remove by Subject/SKID/SerialNumber - approved child exist:
      `TestHandler_RemoveDaIntermediateCert_BySerialNumber_RevokedChildExist`
    * Remove by Subject/SKID - revoked certificate:
      `TestHandler_RemoveDaIntermediateCert_BySubjectAndSKID_RevokedCertificate`
    * Remove by Subject/SKID/SerialNumber - revoked certificate:
      `TestHandler_RemoveDaIntermediateCert_BySerialNumber_RevokedCertificate`
    * Remove by Subject/SKID - revoked and active certificates:
      `TestHandler_RemoveDaIntermediateCert_BySubjectAndSKID_RevokedAndActiveCertificate`
    * Remove by Subject/SKID - another certificate with same Subject exist: ?
    * Remove by Subject/SKID - another certificate with same SKID exist: ?
    * Remove by other Vendor with the same VID: `TestHandler_RemoveDaIntermediateCert_ByNotOwnerButSameVendor`
* Negative:
    * Remove by not Vendor: `TestHandler_RemoveDaIntermediateCert_SenderNotVendor`
    * Remove not existing certificated (Subject/SKID): `TestHandler_RemoveDaIntermediateCert_CertificateDoesNotExist`
    * Remove not existing certificated (Subject/SKID + SerialNumber):
      `TestHandler_RemoveDaIntermediateCert_InvalidSerialNumber`
    * Remove root certificate: `TestHandler_RemoveDaIntermediateCert_ForRootCertificate`
    * Remove NOC certificate: `TestHandler_RemoveDaIntermediateCert_ForNocIcaCertificate`
    * Remove by other Vendor with different VID: `TestHandler_RemoveDaIntermediateCert_ByOtherVendor`

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
    * Add certificate: `TestHandler_AddNocRootCert`
    * Add two certificates with same Subject/SKID but different SerialNumber:
      `TestHandler_AddNocRootCert_SameSubjectAndSkid_DifferentSerialNumber`
    * Add certificates with same Subject but different SKID: ?
    * Add two certificates with same SKID but different Subject: ?
    * Add two certificates but different Vendors with same VID: `TestHandler_AddNocRootCert_ByNotOwnerButSameVendor`
* Negative:
    * Add by not Vendor: `TestHandler_AddNocRootCert_SenderNotVendor`
    * Add invalid certificate: `TestHandler_AddNocRootCert_InvalidCertificate:NotValidPemCertificate`
    * Add not root: `TestHandler_AddNocRootCert_InvalidCertificate:NonRootCertificate`
    * Add with existing Issuer/SerialNumber: `TestHandler_AddNocRootCert_CertificateExist:Duplicate`
    * Add DA certificate: `TestHandler_AddNocRootCert_CertificateExist:ExistingNotNocCert`
    * Add by Vendor with different VID: `TestHandler_AddNocRootCert_CertificateExist:ExistingCertWithDifferentVid`

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
    * Add certificate: `TestHandler_AddNocIntermediateCert`
    * Add two certificates with same Subject/SKID but different SerialNumber:
      `TestHandler_AddNocIntermediateCert_SameSubjectAndSkid_DifferentSerialNumber`
    * Add two certificates with same Subject but different SKID: ?
    * Add two certificates with same SKID but different Subject: ?
    * Add two certificates but different Vendors with same VID:
      `TestHandler_AddNocIntermediateCert_ByNotOwnerButSameVendor`
* Negative:
    * Add by not Vendor: `TestHandler_AddNocIntermediateCert_SenderNotVendor`
    * Add invalid certificate: `TestHandler_AddNocIntermediateCert_ForInvalidCertificate`
    * Add NOC root: `TestHandler_AddNocIntermediateCert_ForNocRootCertificate`
    * Add with existing Issuer/SerialNumber: `TestHandler_AddNocIntermediateCert_CertificateExist`
    * Add for invalid chain of parent certificates: `TestHandler_AddNocIntermediateCert_WhenNocRootCertIsAbsent`
    * Add DA certificate: `TestHandler_AddNocIntermediateCert_ForRootNonNocCertificate`
    * Add by Vendor with different VID: `TestHandler_AddNocIntermediateCert_Root_VID_Does_Not_Equal_To_AccountVID`

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
    * Revoke by Subject/SKID: `TestHandler_RevokeNocRootCert_BySubjectAndSKID`
    * Revoke by Subject/SKID/SerialNumber: `TestHandler_RevokeNocRootCert_BySerialNumber`
    * Revoke by Subject/SKID - revoke child: `TestHandler_RevokeNocRootCert_BySubjectAndSKID_RevokeChild`
    * Revoke by Subject/SKID/SerialNumber - revoke child: `TestHandler_RevokeNocRootCert_BySerialNumber_RevokeChild`
    * Revoke by Subject/SKID - keep child: `TestHandler_RevokeNocRootCert_BySubjectAndSKID_KeepChild`
    * Revoke by Subject/SKID/SerialNumber - keep child: `TestHandler_RevokeNocRootCert_BySerialNumber_KeepChild`
    * Revoke by Subject/SKID - another certificate with same Subject exist: ?
    * Revoke by Subject/SKID - another certificate with same SKID exist: ?
    * Revoke by other Vendor with the same VID: `TestHandler_RevokeNocRootCert_OtherVendor`
* Negative:
    * Revoke by not Vendor: `TestHandler_RevokeNocRootCert_SenderNotVendor`
    * Revoke not existing certificate (Subject/SKID): `TestHandler_RevokeNocRootCert_CertificateDoesNotExist`
    * Revoke not existing certificate by SerialNumber (Subject/SKID + SerialNumber):
      `TestHandler_RevokeNocRootCert_CertificateExists`
    * Revoke not root certificate: `TestHandler_RevokeNocRootCert_CertificateExists`
    * Revoke not NOC certificate: `TestHandler_RevokeNocRootCert_CertificateExists`
    * Revoke by Vendor with different VID: `TestHandler_RevokeNocRootCert_CertificateExists`

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
    * Revoke by Subject/SKID: `TestHandler_RevokeNocIntermediateCert_BySubjectAndSKID`
    * Revoke by Subject/SKID/SerialNumber: `TestHandler_RevokeNocIntermediateCert_BySerialNumber`
    * Revoke by Subject/SKID - revoke child: `TestHandler_RevokeNocIntermediateCert_BySubjectAndSKID_RevokeChild`
    * Revoke by Subject/SKID/SerialNumber - revoke child:
      `TestHandler_RevokeNocIntermediateCert_BySerialNumber_RevokehChild`
    * Revoke by Subject/SKID - keep child: `TestHandler_RevokeNocIntermediateCert_BySubjectAndSKID_KeepChild`
    * Revoke by Subject/SKID/SerialNumber - keep child: `TestHandler_RevokeNocIntermediateCert_BySerialNumber_KeepChild`
    * Revoke by Subject/SKID - parent not affected: `TestHandler_RevokeNocIntermediateCert_BySubjectAndSKID_ParentExist`
    * Revoke by Subject/SKID/SerialNumber - parent not affected:
      `TestHandler_RevokeNocIntermediateCert_BySerialNumber_ParentExist`
    * Revoke by Subject/SKID - another certificate with same Subject exist: ?
    * Revoke by Subject/SKID - another certificate with same SKID exist: ?
    * Revoke by other Vendor with the same VID: `TestHandler_RevokeNocIntermediateCert_ByOtherVendor`
* Negative:
    * Revoke by not Vendor: `TestHandler_RevokeNocIntermediateCert_SenderNotVendor`
    * Revoke not existing certificate by Subject/SKID: `TestHandler_RevokeNocIntermediateCert_CertificateDoesNotExist`
    * Revoke not existing certificate by Subject/SKID/SerialNumber:
      `TestHandler_RevokeNocIntermediateCert_CertificateExists`
    * Revoke root certificate: `TestHandler_RevokeNocIntermediateCert_CertificateExists`
    * Revoke root DA certificate: `TestHandler_RevokeNocIntermediateCert_CertificateExists`
    * Revoke by Vendor with different VID: `TestHandler_RevokeNocIntermediateCert_CertificateExists`

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
    * Remove by Subject/SKID: `TestHandler_RemoveNocRootCert_BySubjectAndSKID`
    * Remove by Subject/SKID/SerialNumber: `TestHandler_RemoveNocRootCert_BySerialNumber`
    * Remove by Subject/SKID - child exist: `TestHandler_RemoveNocRootCert_BySubjectAndSKID_ChildExist`
    * Remove by Subject/SKID/SerialNumber - child exist: `TestHandler_RemoveNocRootCert_BySerialNumber_ChildExist`
    * Remove by Subject/SKID - revoked certificate:
      `TestHandler_RemoveNocRootCert_BySubjectAndSKID_RevokedCertificate`
    * Remove by Subject/SKID/SerialNumber - revoked certificate:
      `TestHandler_RemoveNocRootCert_BySerialNumber_RevokedCertificate`
    * Remove by Subject/SKID - revoked and active certificates:
      `TestHandler_RemoveNocRootCert_BySubjectAndSKID_RevokedAndActiveCertificate`
    * Remove by Subject/SKID - another certificate with same Subject exist: ?
    * Remove by Subject/SKID - another certificate with same SKID exist: ?
    * Remove by other Vendor with the same VID: `TestHandler_RemoveNocRootCert_ByNotOwnerButSameVendor`
* Negative:
    * Remove by not Vendor: `TestHandler_RemoveNocRootCert_SenderNotVendor`
    * Remove not existing certificated (Subject/SKID): `TestHandler_RemoveNocRootCert_CertificateDoesNotExist`
    * Remove not existing certificated (Subject/SKID + SerialNumber):
      `TestHandler_RemoveNocRootCert_InvalidSerialNumber`
    * Remove intermediate certificate: `TestHandler_RemoveNocRootCert_IntermediateCertificate`
    * Remove DA certificate: `TestHandler_RemoveNocRootCert_DaCertificate`
    * Remove by other Vendor with different VID: `TestHandler_RemoveNocRootCert_ByOtherVendor`

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
    * Remove by Subject/SKID: `TestHandler_RemoveNocIntermediateCert_BySubjectAndSKID`
    * Remove by Subject/SKID/SerialNumber: `TestHandler_RemoveNocIntermediateCert_BySerialNumber`
    * Remove by Subject/SKID - parent exist: `TestHandler_RemoveNocIntermediateCert_BySubjectAndSKID_ParentExist`
    * Remove by Subject/SKID/SerialNumber - parent exist:
      `TestHandler_RemoveNocIntermediateCert_BySerialNumber_ParentExist`
    * Remove by Subject/SKID - approved child exist:
      `TestHandler_RemoveNocIntermediateCert_BySubjectAndSKID_ApprovedChildExist`
    * Remove by Subject/SKID/SerialNumber - approved child exist:
      `TestHandler_RemoveNocIntermediateCert_BySerialNumber_ApprovedChildExist`
    * Remove by Subject/SKID - revoked child exist:
      `TestHandler_RemoveNocIntermediateCert_BySubjectAndSKID_RevokedChildExist`
    * Remove by Subject/SKID/SerialNumber - revoked child exist:
      `TestHandler_RemoveNocIntermediateCert_BySerialNumber_RevokedChildExist`
    * Remove by Subject/SKID - revoked certificate:
      `TestHandler_RemoveNocIntermediateCert_BySubjectAndSKID_RevokedCertificate`
    * Remove by Subject/SKID/SerialNumber - revoked certificate:
      `TestHandler_RemoveNocIntermediateCert_BySerialNumber_RevokedCertificate`
    * Remove by Subject/SKID - revoked and active certificates:
      `TestHandler_RemoveNocIntermediateCert_BySubjectAndSKID_RevokedAndActiveCertificate`
    * Remove by Subject/SKID - another certificate with same Subject exist: ?
    * Remove by Subject/SKID - another certificate with same SKID exist: ?
    * Remove by other Vendor with the same VID: `TestHandler_RemoveNocIntermediateCert_ByNotOwnerButSameVendor`
* Negative:
    * Remove by not Vendor: `TestHandler_RemoveNocIntermediateCert_SenderNotVendor`
    * Remove not existing certificated (Subject/SKID): `TestHandler_RemoveNocIntermediateCert_CertificateDoesNotExist`
    * Remove not existing certificated (Subject/SKID + SerialNumber):
      `TestHandler_RemoveNocIntermediateCert_InvalidSerialNumber`
    * Remove NOC root certificate: `TestHandler_RemoveNocIntermediateCert_ForRoot`
    * Remove DA certificate: `TestHandler_RemoveNocIntermediateCert_ForDaCertificate`
    * Remove by other Vendor with different VID: `TestHandler_RemoveNocIntermediateCert_ByOtherVendor`