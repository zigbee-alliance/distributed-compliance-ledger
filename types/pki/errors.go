package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

// x/pki module sentinel errors.
var (
	ErrProposedCertificateAlreadyExists                  = sdkerrors.Register(ModuleName, 401, "proposed certificate already exists")
	ErrProposedCertificateDoesNotExist                   = sdkerrors.Register(ModuleName, 402, "proposed certificate does not exist")
	ErrCertificateAlreadyExists                          = sdkerrors.Register(ModuleName, 403, "certificate already exists")
	ErrCertificateDoesNotExist                           = sdkerrors.Register(ModuleName, 404, "certificate does not exist")
	ErrProposedCertificateRevocationAlreadyExists        = sdkerrors.Register(ModuleName, 405, "proposed certificate revocation already exists")
	ErrProposedCertificateRevocationDoesNotExist         = sdkerrors.Register(ModuleName, 406, "proposed certificate revocation does not exist")
	ErrRevokedCertificateDoesNotExist                    = sdkerrors.Register(ModuleName, 407, "revoked certificate does not exist")
	ErrInappropriateCertificateType                      = sdkerrors.Register(ModuleName, 408, "inappropriate certificate type")
	ErrInvalidCertificate                                = sdkerrors.Register(ModuleName, 409, "invalid certificate")
	ErrInvalidDataDigestType                             = sdkerrors.Register(ModuleName, 410, "invalid data digest type")
	ErrInvalidRevocationType                             = sdkerrors.Register(ModuleName, 411, "invalid revocation type")
	ErrNotEmptyPid                                       = sdkerrors.Register(ModuleName, 412, "pid is not empty")
	ErrNotEmptyVid                                       = sdkerrors.Register(ModuleName, 413, "vid is not empty")
	ErrRootCertificateIsNotSelfSigned                    = sdkerrors.Register(ModuleName, 414, "Root certificate is not self-signed")
	ErrCRLSignerCertificatePidNotEqualRevocationPointPid = sdkerrors.Register(ModuleName, 415, "CRLSignerCertificate pid does not equal revocation point pid")
	ErrCRLSignerCertificateVidNotEqualRevocationPointVid = sdkerrors.Register(ModuleName, 416, "CRLSignerCertificate vid does not equal revocation point pid")
	ErrCRLSignerCertificatePidNotEqualMsgPid             = sdkerrors.Register(ModuleName, 417, "CRLSignerCertificate pid does not equal message pid")
	ErrCRLSignerCertificateVidNotEqualMsgVid             = sdkerrors.Register(ModuleName, 418, "CRLSignerCertificate vid does not equal message vid")
	ErrMessageVidNotEqualAccountVid                      = sdkerrors.Register(ModuleName, 419, "Message vid does not equal the account vid")
	ErrNonRootCertificateSelfSigned                      = sdkerrors.Register(ModuleName, 420, "Intermediate or leaf certificate must not be self-signed")
	ErrEmptyDataFileSize                                 = sdkerrors.Register(ModuleName, 421, "empty data file size")
	ErrEmptyDataDigest                                   = sdkerrors.Register(ModuleName, 422, "empty data digest")
	ErrEmptyDataDigestType                               = sdkerrors.Register(ModuleName, 423, "empty data digest type")
	ErrNotEmptyDataDigestType                            = sdkerrors.Register(ModuleName, 424, "not empty data digest type")
	ErrDataFieldPresented                                = sdkerrors.Register(ModuleName, 425, "one or more of DataDigest, DataDigestType, DataFileSize fields presented")
	ErrWrongSubjectKeyIDFormat                           = sdkerrors.Register(ModuleName, 426, "wrong SubjectKeyID format")
	ErrVidNotFound                                       = sdkerrors.Register(ModuleName, 427, "vid not found")
	ErrPidNotFound                                       = sdkerrors.Register(ModuleName, 428, "pid not found")
	ErrPemValuesNotEqual                                 = sdkerrors.Register(ModuleName, 429, "pem values of certificates are not equal")
	ErrPkiRevocationDistributionPointAlreadyExists       = sdkerrors.Register(ModuleName, 430, "pki revocation distribution point already exists")
	ErrPkiRevocationDistributionPointDoesNotExists       = sdkerrors.Register(ModuleName, 431, "pki revocaition distribution point does not exist")
	ErrUnsupportedOperation                              = sdkerrors.Register(ModuleName, 432, "unsupported operation")
	ErrInvalidVidFormat                                  = sdkerrors.Register(ModuleName, 433, "invalid vid format")
	ErrInvalidPidFormat                                  = sdkerrors.Register(ModuleName, 434, "invalid pid format")
	ErrInvalidDataURLFormat                              = sdkerrors.Register(ModuleName, 435, "invalid data url format")
	ErrCertificateVidNotEqualMsgVid                      = sdkerrors.Register(ModuleName, 436, "certificate's vid is not equal to the message vid")
	ErrMessageVidNotEqualRootCertVid                     = sdkerrors.Register(ModuleName, 437, "Message vid is not equal to ledger's root certificate vid")
	ErrCertNotChainedBack                                = sdkerrors.Register(ModuleName, 438, "Certificate is not chained back to a root certificate on DCL")
	ErrCertVidNotEqualAccountVid                         = sdkerrors.Register(ModuleName, 439, "account's vid is not equal to ledger's certificate vid")
	ErrCertVidNotEqualToRootVid                          = sdkerrors.Register(ModuleName, 440, "certificate's vid is not equal to vid of root certificate ")
)

func NewErrUnauthorizedRole(transactionName string, requiredRole types.AccountRole) error {
	return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
		"%s transaction should be signed by an account with the \"%s\" role", transactionName, requiredRole)
}

func NewErrInvalidAddress(err error) error {
	return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%v)", err)
}

func NewErrProposedCertificateAlreadyExists(subject string, subjectKeyID string) error {
	return sdkerrors.Wrapf(ErrProposedCertificateAlreadyExists,
		"Proposed X509 root certificate associated with the combination "+
			"of subject=%v and subjectKeyID=%v already exists on the ledger",
		subject, subjectKeyID)
}

func NewErrProposedCertificateDoesNotExist(subject string, subjectKeyID string) error {
	return sdkerrors.Wrapf(ErrProposedCertificateDoesNotExist,
		"No proposed X509 root certificate associated "+
			"with the combination of subject=%v and subjectKeyID=%v on the ledger. "+
			"The cerificate either does not exists or already approved.",
		subject, subjectKeyID)
}

func NewErrCertificateAlreadyExists(issuer string, serialNumber string) error {
	return sdkerrors.Wrapf(ErrCertificateAlreadyExists,
		"X509 certificate associated with the combination of "+
			"issuer=%v and serialNumber=%v already exists on the ledger",
		issuer, serialNumber)
}

func NewErrCertificateDoesNotExist(subject string, subjectKeyID string) error {
	return sdkerrors.Wrapf(ErrCertificateDoesNotExist,
		"No X509 certificate associated with the "+
			"combination of subject=%v and subjectKeyID=%v on the ledger",
		subject, subjectKeyID)
}

func NewErrCertificateBySerialNumberDoesNotExist(subject string, subjectKeyID string, serialNumber string) error {
	return sdkerrors.Wrapf(ErrCertificateDoesNotExist,
		"No X509 certificate associated with the "+
			"combination of subject=%v, subjectKeyID=%v and serialNumber=%v on the ledger",
		subject, subjectKeyID, serialNumber)
}

func NewErrRootCertificateDoesNotExist(subject string, subjectKeyID string) error {
	return sdkerrors.Wrapf(ErrCertificateDoesNotExist,
		"No X509 root certificate associated with the "+
			"combination of subject=%s and subjectKeyID=%s on the ledger",
		subject, subjectKeyID,
	)
}

func NewErrProposedCertificateRevocationAlreadyExists(subject string, subjectKeyID string) error {
	return sdkerrors.Wrapf(ErrProposedCertificateRevocationAlreadyExists,
		"Proposed X509 root certificate revocation associated with the combination "+
			"of subject=%v and subjectKeyID=%v already exists on the ledger",
		subject, subjectKeyID)
}

func NewErrProposedCertificateRevocationDoesNotExist(subject string, subjectKeyID string) error {
	return sdkerrors.Wrapf(ErrProposedCertificateRevocationDoesNotExist,
		"No proposed X509 root certificate revocation associated "+
			"with the combination of subject=%v and subjectKeyID=%v on the ledger.",
		subject, subjectKeyID)
}

func NewErrRevokedCertificateDoesNotExist(subject string, subjectKeyID string) error {
	return sdkerrors.Wrapf(ErrRevokedCertificateDoesNotExist,
		"No revoked X509 certificate associated with the "+
			"combination of subject=%v and subjectKeyID=%v on the ledger",
		subject, subjectKeyID)
}

func NewErrInappropriateCertificateType(e interface{}) error {
	return sdkerrors.Wrapf(ErrInappropriateCertificateType, "%v",
		e)
}

func NewErrInvalidCertificate(e interface{}) error {
	return sdkerrors.Wrapf(ErrInvalidCertificate, "%v",
		e)
}

func NewErrInvalidDataDigestType(dataDigestType uint32, allowedDataDigestTypes []uint32) error {
	return sdkerrors.Wrapf(ErrInvalidDataDigestType,
		"Invalid DataDigestType: %d. Supported types are: %v", dataDigestType, allowedDataDigestTypes)
}

func NewErrInvalidRevocationType(revocationType uint32, allowedRevocationTypes []uint32) error {
	return sdkerrors.Wrapf(ErrInvalidRevocationType,
		"Invalid RevocationType: %d. Supported types are: %v", revocationType, allowedRevocationTypes)
}

func NewErrNotEmptyPidForRootCertificate() error {
	return sdkerrors.Wrapf(ErrNotEmptyPid,
		"Product ID (pid) must be empty for root certificates")
}

func NewErrNotEmptyPidForNonRootCertificate() error {
	return sdkerrors.Wrapf(ErrNotEmptyPid,
		"Product ID (pid) must be empty when it is not found in non-root CRL Signer Certificate")
}

func NewErrNotEmptyVid(e interface{}) error {
	return sdkerrors.Wrapf(ErrNotEmptyVid, "%v",
		e)
}

func NewErrRootCertificateIsNotSelfSigned() error {
	return sdkerrors.Wrapf(
		ErrRootCertificateIsNotSelfSigned,
		"Provided root certificate must be self-signed",
	)
}

func NewErrNonRootCertificateSelfSigned() error {
	return sdkerrors.Wrapf(
		ErrNonRootCertificateSelfSigned,
		"Provided non-root certificate must not be self-signed",
	)
}

func NewErrUnauthorizedCertIssuer(subject string, subjectKeyID string) error {
	return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
		"Issuer and authorityKeyID of new certificate with subject=%v and subjectKeyID=%v "+
			"must be the same as ones of existing certificates with the same subject and subjectKeyID",
		subject, subjectKeyID)
}

func NewErrUnauthorizedCertOwner(subject string, subjectKeyID string) error {
	return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
		"Only owner of existing certificates with subject=%v and subjectKeyID=%v "+
			"can add new certificate with the same subject and subjectKeyID",
		subject, subjectKeyID)
}

func NewErrProvidedNocCertButExistingNotNoc(subject string, subjectKeyID string) error {
	return sdkerrors.Wrapf(ErrInappropriateCertificateType,
		"The existing certificate with the same combination of subject (%v) and subjectKeyID (%v) is not a NOC certificate",
		subject, subjectKeyID)
}

func NewErrProvidedNotNocCertButExistingNoc(subject string, subjectKeyID string) error {
	return sdkerrors.Wrapf(ErrInappropriateCertificateType,
		"The existing certificate with the same combination of subject (%v) and subjectKeyID (%v) is a NOC certificate",
		subject, subjectKeyID)
}

func NewErrExistingCertVidNotEqualAccountVid(subject string, subjectKeyID string, vid int32) error {
	return sdkerrors.Wrapf(ErrCertVidNotEqualAccountVid,
		"Certificate with the same combination of subject=%v and subjectKeyID=%v "+
			"has already been published by another vendor with VID=%d.",
		subject, subjectKeyID, vid)
}

func NewErrRootCertVidNotEqualToAccountVidOrCertVid(rootVID int32, accountVID int32, certVID int32) error {
	if rootVID != certVID {
		return sdkerrors.Wrapf(ErrCertVidNotEqualToRootVid,
			"Root certificate is VID scoped: An intermediate certificate must be also VID scoped to the same VID as a root one: "+
				"Root certificate's VID = %v, Certificate's VID = %v",
			rootVID, certVID)
	}

	return sdkerrors.Wrapf(ErrCertVidNotEqualAccountVid,
		"Root certificate is VID scoped: "+
			"Only a Vendor associated with this VID can add an intermediate certificate: "+
			"Root certificate's VID = %v, Account VID = %v",
		rootVID, accountVID)
}

func NewErrAccountVidNotEqualToCertVid(accountVID int32, certVID int32) error {
	return sdkerrors.Wrapf(ErrCertVidNotEqualAccountVid,
		"Intermediate certificate is VID scoped: Only a Vendor associated with this VID can add an intermediate certificate: "+
			"Account VID = %v, Certificate's VID = %v",
		accountVID, certVID)
}

func NewErrCRLSignerCertificatePidNotEqualMsgPid(certificatePid int32, messagePid int32) error {
	return sdkerrors.Wrapf(
		ErrCRLSignerCertificatePidNotEqualMsgPid,
		"CRL Signer Certificate's pid=%d must be equal to the provided pid=%d in the message",
		certificatePid, messagePid,
	)
}

func NewErrCRLSignerCertificateVidNotEqualMsgVid(certificateVid int32, messageVid int32) error {
	return sdkerrors.Wrapf(
		ErrCRLSignerCertificateVidNotEqualMsgVid,
		"CRL Signer Certificate's vid=%d must be equal to the provided vid=%d in the message",
		certificateVid, messageVid,
	)
}

func NewErrMessageVidNotEqualRootCertVid(vid1 int32, vid2 int32) error {
	return sdkerrors.Wrapf(ErrMessageVidNotEqualRootCertVid,
		"Message vid=%d is not equal to ledger's root certificate vid=%d",
		vid1, vid2)
}

func NewErrCRLSignerCertificatePidNotEqualRevocationPointPid(certificatePid int32, revocationPointPid int32) error {
	return sdkerrors.Wrapf(
		ErrCRLSignerCertificatePidNotEqualRevocationPointPid,
		"CRL Signer Certificate's pid=%d must be equal to the provided pid=%d in the reovocation point",
		certificatePid, revocationPointPid)
}

func NewErrCRLSignerCertificateVidNotEqualRevocationPointVid(vid1 int32, vid2 int32) error {
	return sdkerrors.Wrapf(ErrCRLSignerCertificateVidNotEqualRevocationPointVid,
		"CRL Signer Certificate's vid=%d must be equal to the provided vid=%d in the reovocation point", vid1, vid2)
}

func NewErrNonEmptyDataDigest() error {
	return sdkerrors.Wrapf(ErrEmptyDataFileSize, "Data Digest must be provided only if Data File Size is provided")
}

func NewErrNotEmptyDataDigestType() error {
	return sdkerrors.Wrapf(ErrNotEmptyDataDigestType, "Data Digest Type must be provided only if Data Digest is provided")
}

func NewErrEmptyDataDigest() error {
	return sdkerrors.Wrapf(ErrEmptyDataDigest, "Data Digest must be provided if Data File Size is provided")
}

func NewErrEmptyDataDigestType() error {
	return sdkerrors.Wrapf(ErrEmptyDataDigestType, "Data Digest Type must be provided if Data Digest is provided")
}

func NewErrDataFieldPresented(revocationType uint32) error {
	return sdkerrors.Wrapf(
		ErrDataFieldPresented,
		"Data Digest, Data File Size and Data Digest Type must be omitted for Revocation Type %d", revocationType,
	)
}

func NewErrWrongSubjectKeyIDFormat() error {
	return sdkerrors.Wrapf(
		ErrWrongSubjectKeyIDFormat,
		"Wrong IssuerSubjectKeyID format. It must consist of even number of uppercase hexadecimal characters ([0-9A-F]), "+
			"with no whitespace and no non-hexadecimal characters",
	)
}

func NewErrVidNotFound(e interface{}) error {
	return sdkerrors.Wrapf(ErrVidNotFound, "%v",
		e)
}

func NewErrPidNotFoundInCertificateButProvidedInRevocationPoint() error {
	return sdkerrors.Wrapf(
		ErrPidNotFound,
		"Product ID (pid) not found in CRL Signer Certificate when it is provided in the revocation point",
	)
}

func NewErrPidNotFoundInMessage(certificatePid int32) error {
	return sdkerrors.Wrapf(
		ErrPidNotFound,
		"Product ID (pid) must be provided when pid=%d in non-root CRL Signer Certificate", certificatePid,
	)
}

func NewErrPemValuesNotEqual(subject string, subjectKeyID string) error {
	return sdkerrors.Wrapf(
		ErrPemValuesNotEqual,
		"PEM values of the CRL signer certificate and a certificate found by subject=%s and subjectKeyID=%s are not equal",
		subject, subjectKeyID,
	)
}

func NewErrPkiRevocationDistributionPointWithVidAndLabelAlreadyExists(vid int32, label string, issuerSubjectKeyID string) error {
	return sdkerrors.Wrapf(
		ErrPkiRevocationDistributionPointAlreadyExists,
		"PKI revocation distribution point associated with vid=%d and label=%s already exist for issuerSubjectKeyID=%s",
		vid, label, issuerSubjectKeyID,
	)
}

func NewErrPkiRevocationDistributionPointWithDataURLAlreadyExists(dataURL string, issuerSubjectKeyID string) error {
	return sdkerrors.Wrapf(
		ErrPkiRevocationDistributionPointAlreadyExists,
		"PKI revocation distribution point associated with dataUrl=%s already exist for issuerSubjectKeyID=%s",
		dataURL, issuerSubjectKeyID,
	)
}

func NewErrPkiRevocationDistributionPointDoesNotExists(vid int32, label string, issuerSubjectKeyID string) error {
	return sdkerrors.Wrapf(
		ErrPkiRevocationDistributionPointDoesNotExists,
		"PKI revocation distribution point associated with vid=%d and label=%s does not exist for issuerSubjectKeyID=%s",
		vid, label, issuerSubjectKeyID,
	)
}

func NewErrMessageVidNotEqualAccountVid(msgVid int32, accountVid int32) error {
	return sdkerrors.Wrapf(ErrMessageVidNotEqualAccountVid, "Message vid=%d is not equal to account vid=%d", msgVid, accountVid)
}

func NewErrMessageRemoveRoot(subject string, subjectKeyID string) error {
	return sdkerrors.Wrapf(ErrInappropriateCertificateType, "Inappropriate Certificate Type: Certificate with subject=%s and subjectKeyID=%s "+
		"is a root certificate.", subject, subjectKeyID,
	)
}

func NewErrMessageOnlyOwnerCanExecute(command string) error {
	return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "Only owner can revoke certificate using %s", command)
}

func NewErrUnsupportedOperation(e interface{}) error {
	return sdkerrors.Wrapf(ErrUnsupportedOperation, "%v", e)
}

func NewErrInvalidVidFormat(e interface{}) error {
	return sdkerrors.Wrapf(ErrInvalidVidFormat, "Could not parse vid: %v", e)
}

func NewErrInvalidPidFormat(e interface{}) error {
	return sdkerrors.Wrapf(ErrInvalidPidFormat, "Could not parse pid: %v", e)
}

func NewErrInvalidDataURLSchema() error {
	return sdkerrors.Wrapf(ErrInvalidDataURLFormat, "Data Url must start with https:// or http://")
}

func NewErrCertificateVidNotEqualMsgVid(e interface{}) error {
	return sdkerrors.Wrapf(ErrCertificateVidNotEqualMsgVid, "%v", e)
}

func NewErrCertNotChainedBack() error {
	return sdkerrors.Wrapf(ErrCertNotChainedBack, "CRL Signer Certificate is not chained back to root certificate on DCL")
}
