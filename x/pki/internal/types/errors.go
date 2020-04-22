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
	CodeInappropriateCertificateType   sdk.CodeType = 403
)

func ErrCertificateAlreadyExists(subject string, subjectKeyId string) sdk.Error {
	return sdk.NewError(Codespace, CodeCertificateAlreadyExists,
		fmt.Sprintf("X509 certificate associated with the combination of subject=%v and subjectKeyId=%v already exists on the ledger", subject, subjectKeyId))
}

func ErrProposedCertificateDoesNotExist(subject string, subjectKeyId string) sdk.Error {
	return sdk.NewError(Codespace, CodePendingCertificateDoesNotExist,
		fmt.Sprintf("No proposed X509 root certificate associated with the combination of subject=%v and subjectKeyId=%v on the ledger."+
			"The serificate either does not exists or already approved.", subject, subjectKeyId))
}

func ErrCertificateDoesNotExist(subject string, subjectKeyId string) sdk.Error {
	return sdk.NewError(Codespace, CodeCertificateDoesNotExist,
		fmt.Sprintf("No X509 certificate associated with the combination of subject=%v and subjectKeyId=%v on the ledger", subject, subjectKeyId))
}
