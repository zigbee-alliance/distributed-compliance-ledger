package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	Codespace sdk.CodespaceType = ModuleName

	CodeCertificateAlreadyExists       sdk.CodeType = 401
	CodeCertificateDoesNotExist        sdk.CodeType = 402
	CodePendingCertificateDoesNotExist sdk.CodeType = 403
	CodeInappropriateCertificateType   sdk.CodeType = 404
	CodeInvalidCertificate             sdk.CodeType = 405
)

func ErrCertificateAlreadyExists(subject string, subjectKeyId string, serialNumber string) sdk.Error {
	return sdk.NewError(Codespace, CodeCertificateAlreadyExists,
		fmt.Sprintf("X509 certificate associated with the combination of subject=%v, subjectKeyId=%v and serialNumber=%v already exists on the ledger", subject, subjectKeyId, serialNumber))
}

func ErrProposedCertificateDoesNotExist(subject string, subjectKeyId string) sdk.Error {
	return sdk.NewError(Codespace, CodePendingCertificateDoesNotExist,
		fmt.Sprintf("No proposed X509 root certificate associated with the combination of subject=%v and subjectKeyId=%v on the ledger."+
			"The cerificate either does not exists or already approved.", subject, subjectKeyId))
}

func ErrCertificateDoesNotExist(subject string, subjectKeyId string) sdk.Error {
	return sdk.NewError(Codespace, CodeCertificateDoesNotExist,
		fmt.Sprintf("No X509 certificate associated with the combination of subject=%v and subjectKeyId=%v on the ledger", subject, subjectKeyId))
}

func ErrInappropriateCertificateType(error interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeInappropriateCertificateType, fmt.Sprintf("%v", error))
}

func ErrCodeInvalidCertificate(error interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeInvalidCertificate, fmt.Sprintf("%v", error))
}
