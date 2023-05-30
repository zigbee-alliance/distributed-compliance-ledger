package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/pki module sentinel errors.
var (
	ErrProposedCertificateAlreadyExists            = sdkerrors.Register(ModuleName, 401, "proposed certificate already exists")
	ErrProposedCertificateDoesNotExist             = sdkerrors.Register(ModuleName, 402, "proposed certificate does not exist")
	ErrCertificateAlreadyExists                    = sdkerrors.Register(ModuleName, 403, "certificate already exists")
	ErrCertificateDoesNotExist                     = sdkerrors.Register(ModuleName, 404, "certificate does not exist")
	ErrProposedCertificateRevocationAlreadyExists  = sdkerrors.Register(ModuleName, 405, "proposed certificate revocation already exists")
	ErrProposedCertificateRevocationDoesNotExist   = sdkerrors.Register(ModuleName, 406, "proposed certificate revocation does not exist")
	ErrRevokedCertificateDoesNotExist              = sdkerrors.Register(ModuleName, 407, "revoked certificate does not exist")
	ErrInappropriateCertificateType                = sdkerrors.Register(ModuleName, 408, "inappropriate certificate type")
	ErrInvalidCertificate                          = sdkerrors.Register(ModuleName, 409, "invalid certificate")
	ErrInvalidDataDigestType                       = sdkerrors.Register(ModuleName, 410, "invalid data digest type")
	ErrInvalidRevocationType                       = sdkerrors.Register(ModuleName, 411, "invalid revocation type")
	ErrNotEmptyPid                                 = sdkerrors.Register(ModuleName, 412, "pid is not empty")
	ErrNotEmptyVid                                 = sdkerrors.Register(ModuleName, 413, "vid is not empty")
	ErrRootCertificateIsNotSelfSigned              = sdkerrors.Register(ModuleName, 414, "Root certificate is not self-signed")
	ErrCRLSignerCertificatePidNotEqualMsgPid       = sdkerrors.Register(ModuleName, 415, "CRLSignerCertificate pid does not equal message pid")
	ErrCRLSignerCertificateVidNotEqualMsgVid       = sdkerrors.Register(ModuleName, 416, "CRLSignerCertificate vid does not equal message vid")
	ErrCRLSignerCertificateVidNotEqualAccountVid   = sdkerrors.Register(ModuleName, 417, "CRLSignerCertificate vid does not equal account vid")
	ErrNonRootCertificateSelfSigned                = sdkerrors.Register(ModuleName, 418, "Intermediate or leaf certificate must not be self-signed")
	ErrEmptyDataFileSize                           = sdkerrors.Register(ModuleName, 419, "empty data file size")
	ErrEmptyDataDigest                             = sdkerrors.Register(ModuleName, 420, "empty data digest")
	ErrEmptyDataDigestType                         = sdkerrors.Register(ModuleName, 421, "empty data digest type")
	ErrNotEmptyDataDigestType                      = sdkerrors.Register(ModuleName, 422, "not empty data digest type")
	ErrDataFieldPresented                          = sdkerrors.Register(ModuleName, 423, "one or more of DataDigest, DataDigestType, DataFileSize fields presented")
	ErrWrongSubjectKeyIDFormat                     = sdkerrors.Register(ModuleName, 424, "wrong SubjectKeyID format")
	ErrVidNotFound                                 = sdkerrors.Register(ModuleName, 425, "vid not found")
	ErrPidNotFound                                 = sdkerrors.Register(ModuleName, 426, "pid not found")
	ErrPemValuesNotEqual                           = sdkerrors.Register(ModuleName, 427, "pem values of certificates are not equal")
	ErrPkiRevocationDistributionPointAlreadyExists = sdkerrors.Register(ModuleName, 428, "pki revocation distribution point already exists")
	ErrPkiRevocationDistributionPointDoesNotExists = sdkerrors.Register(ModuleName, 429, "pki revocaition distribution point does not exist")
	ErrUnsupportedOperation                        = sdkerrors.Register(ModuleName, 430, "unsupported operation")
	ErrInvalidVidFormat                            = sdkerrors.Register(ModuleName, 431, "invalid vid format")
	ErrInvalidPidFormat                            = sdkerrors.Register(ModuleName, 432, "invalid pid format")
	ErrInvalidDataUrlFormat                        = sdkerrors.Register(ModuleName, 433, "invalid data url format")
)

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

func NewErrInvalidDataDigestType(e interface{}) error {
	return sdkerrors.Wrapf(ErrInvalidDataDigestType, "%v",
		e)
}

func NewErrInvalidRevocationType(e interface{}) error {
	return sdkerrors.Wrapf(ErrInvalidRevocationType, "%v",
		e)
}

func NewErrNotEmptyPid(e interface{}) error {
	return sdkerrors.Wrapf(ErrNotEmptyPid, "%v",
		e)
}

func NewErrNotEmptyVid(e interface{}) error {
	return sdkerrors.Wrapf(ErrNotEmptyVid, "%v",
		e)
}

func NewErrRootCertificateIsNotSelfSigned(e interface{}) error {
	return sdkerrors.Wrapf(ErrRootCertificateIsNotSelfSigned, "%v",
		e)
}

func NewErrCRLSignerCertificatePidNotEqualMsgPid(e interface{}) error {
	return sdkerrors.Wrapf(ErrCRLSignerCertificatePidNotEqualMsgPid, "%v",
		e)
}

func NewErrCRLSignerCertificateVidNotEqualMsgVid(e interface{}) error {
	return sdkerrors.Wrapf(ErrCRLSignerCertificateVidNotEqualMsgVid, "%v",
		e)
}

func NewErrNonRootCertificateSelfSigned(e interface{}) error {
	return sdkerrors.Wrapf(ErrNonRootCertificateSelfSigned, "%v",
		e)
}

func NewErrNonEmptyDataDigest(e interface{}) error {
	return sdkerrors.Wrapf(ErrEmptyDataFileSize, "%v",
		e)
}

func NewErrNotEmptyDataDigestType(e interface{}) error {
	return sdkerrors.Wrapf(ErrNotEmptyDataDigestType, "%v",
		e)
}

func NewErrEmptyDataDigest(e interface{}) error {
	return sdkerrors.Wrapf(ErrEmptyDataDigest, "%v",
		e)
}

func NewErrEmptyDataDigestType(e interface{}) error {
	return sdkerrors.Wrapf(ErrEmptyDataDigestType, "%v",
		e)
}

func NewErrDataFieldPresented(e interface{}) error {
	return sdkerrors.Wrapf(ErrDataFieldPresented, "%v",
		e)
}

func NewErrWrongSubjectKeyIDFormat(e interface{}) error {
	return sdkerrors.Wrapf(ErrWrongSubjectKeyIDFormat, "%v",
		e)
}

func NewErrVidNotFound(e interface{}) error {
	return sdkerrors.Wrapf(ErrVidNotFound, "%v",
		e)
}

func NewErrPidNotFound(e interface{}) error {
	return sdkerrors.Wrapf(ErrPidNotFound, "%v",
		e)
}

func NewErrPemValuesNotEqual(e interface{}) error {
	return sdkerrors.Wrapf(ErrPemValuesNotEqual, "%v", e)
}

func NewErrPkiRevocationDistributionPointAlreadyExists(e interface{}) error {
	return sdkerrors.Wrapf(ErrPkiRevocationDistributionPointAlreadyExists, "%v", e)
}

func NewErrPkiRevocationDistributionPointDoesNotExists(e interface{}) error {
	return sdkerrors.Wrapf(ErrPkiRevocationDistributionPointDoesNotExists, "%v", e)
}

func NewErrCRLSignerCertificateVidNotEqualAccountVid(e interface{}) error {
	return sdkerrors.Wrapf(ErrCRLSignerCertificateVidNotEqualAccountVid, "%v", e)
}

func NewErrUnsupportedOperation(e interface{}) error {
	return sdkerrors.Wrapf(ErrUnsupportedOperation, "%v", e)
}

func NewErrInvalidVidFormat(e interface{}) error {
	return sdkerrors.Wrapf(ErrInvalidVidFormat, "%v", e)
}

func NewErrInvalidPidFormat(e interface{}) error {
	return sdkerrors.Wrapf(ErrInvalidPidFormat, "%v", e)
}

func NewErrInvalidDataUrlFormat(e interface{}) error {
	return sdkerrors.Wrapf(ErrInvalidDataUrlFormat, "%v", e)
}
