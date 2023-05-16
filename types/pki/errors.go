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
	ErrPAANotSelfSigned                            = sdkerrors.Register(ModuleName, 413, "PAA is not self-signed")
	ErrCRLSignerCertificateDoesNotContainPid       = sdkerrors.Register(ModuleName, 414, "CRLSignerCertificate does not contain pid")
	ErrNonPAASelfSigned                            = sdkerrors.Register(ModuleName, 415, "non PAA certificate self signed")
	ErrEmptyDataFileSize                           = sdkerrors.Register(ModuleName, 416, "empty data file size")
	ErrEmptyDataDigest                             = sdkerrors.Register(ModuleName, 417, "empty data digest")
	ErrEmptyDataDigestType                         = sdkerrors.Register(ModuleName, 418, "empty data digest type")
	ErrDataFieldPresented                          = sdkerrors.Register(ModuleName, 419, "data field presented")
	ErrIssuerSubjectKeyIDNotEqualsCertSubjectKeyID = sdkerrors.Register(ModuleName, 420, "provided SubjectKeyID and certificate SubjectKeyID are not equal")
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

func NewErrPAANotSelfSigned(e interface{}) error {
	return sdkerrors.Wrapf(ErrPAANotSelfSigned, "%v",
		e)
}

func NewErrCRLSignerCertificateDoesNotContainPid(e interface{}) error {
	return sdkerrors.Wrapf(ErrCRLSignerCertificateDoesNotContainPid, "%v",
		e)
}

func NewErrNonPAASelfSigned(e interface{}) error {
	return sdkerrors.Wrapf(ErrNonPAASelfSigned, "%v",
		e)
}

func NewErrEmptyDataFileSize(e interface{}) error {
	return sdkerrors.Wrapf(ErrEmptyDataFileSize, "%v",
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

func NewErrIssuerSubjectKeyIDNotEqualsCertSubjectKeyID(e interface{}) error {
	return sdkerrors.Wrapf(ErrIssuerSubjectKeyIDNotEqualsCertSubjectKeyID, "%v",
		e)
}
