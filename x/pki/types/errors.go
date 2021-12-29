package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/pki module sentinel errors.
var (
	ErrProposedCertificateAlreadyExists           = sdkerrors.Register(ModuleName, 401, "proposed certificate already exists")
	ErrProposedCertificateDoesNotExist            = sdkerrors.Register(ModuleName, 402, "proposed certificate does not exist")
	ErrCertificateAlreadyExists                   = sdkerrors.Register(ModuleName, 403, "certificate already exists")
	ErrCertificateDoesNotExist                    = sdkerrors.Register(ModuleName, 404, "certificate does not exist")
	ErrProposedCertificateRevocationAlreadyExists = sdkerrors.Register(ModuleName, 405, "proposed certificate revocation already exists")
	ErrProposedCertificateRevocationDoesNotExist  = sdkerrors.Register(ModuleName, 406, "proposed certificate revocation does not exist")
	ErrRevokedCertificateDoesNotExist             = sdkerrors.Register(ModuleName, 407, "revoked certificate does not exist")
	ErrInappropriateCertificateType               = sdkerrors.Register(ModuleName, 408, "inappropriate certificate type")
	ErrInvalidCertificate                         = sdkerrors.Register(ModuleName, 409, "invalid certificate")
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

func NewErrCodeInvalidCertificate(e interface{}) error {
	return sdkerrors.Wrapf(ErrInvalidCertificate, "%v",
		e)
}
