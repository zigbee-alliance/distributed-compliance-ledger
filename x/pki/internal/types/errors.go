// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	Codespace sdk.CodespaceType = ModuleName

	CodeProposedCertificateAlreadyExists           sdk.CodeType = 401
	CodeProposedCertificateDoesNotExist            sdk.CodeType = 402
	CodeCertificateAlreadyExists                   sdk.CodeType = 403
	CodeCertificateDoesNotExist                    sdk.CodeType = 404
	CodeProposedCertificateRevocationAlreadyExists sdk.CodeType = 405
	CodeProposedCertificateRevocationDoesNotExist  sdk.CodeType = 406
	CodeRevokedCertificateDoesNotExist             sdk.CodeType = 407
	CodeInappropriateCertificateType               sdk.CodeType = 408
	CodeInvalidCertificate                         sdk.CodeType = 409
)

func ErrProposedCertificateAlreadyExists(subject string, subjectKeyID string) sdk.Error {
	return sdk.NewError(Codespace, CodeProposedCertificateAlreadyExists,
		fmt.Sprintf("Proposed X509 root certificate associated with the combination "+
			"of subject=%v and subjectKeyID=%v already exists on the ledger", subject, subjectKeyID))
}

func ErrProposedCertificateDoesNotExist(subject string, subjectKeyID string) sdk.Error {
	return sdk.NewError(Codespace, CodeProposedCertificateDoesNotExist,
		fmt.Sprintf("No proposed X509 root certificate associated "+
			"with the combination of subject=%v and subjectKeyID=%v on the ledger. "+
			"The cerificate either does not exists or already approved.", subject, subjectKeyID))
}

func ErrCertificateAlreadyExists(issuer string, serialNumber string) sdk.Error {
	return sdk.NewError(Codespace, CodeCertificateAlreadyExists,
		fmt.Sprintf("X509 certificate associated with the combination of "+
			"issuer=%v and serialNumber=%v already exists on the ledger", issuer, serialNumber))
}

func ErrCertificateDoesNotExist(subject string, subjectKeyID string) sdk.Error {
	return sdk.NewError(Codespace, CodeCertificateDoesNotExist,
		fmt.Sprintf("No X509 certificate associated with the "+
			"combination of subject=%v and subjectKeyID=%v on the ledger", subject, subjectKeyID))
}

func ErrProposedCertificateRevocationAlreadyExists(subject string, subjectKeyID string) sdk.Error {
	return sdk.NewError(Codespace, CodeProposedCertificateRevocationAlreadyExists,
		fmt.Sprintf("Proposed X509 root certificate revocation associated with the combination "+
			"of subject=%v and subjectKeyID=%v already exists on the ledger", subject, subjectKeyID))
}

func ErrProposedCertificateRevocationDoesNotExist(subject string, subjectKeyID string) sdk.Error {
	return sdk.NewError(Codespace, CodeProposedCertificateRevocationDoesNotExist,
		fmt.Sprintf("No proposed X509 root certificate revocation associated "+
			"with the combination of subject=%v and subjectKeyID=%v on the ledger.", subject, subjectKeyID))
}

func ErrRevokedCertificateDoesNotExist(subject string, subjectKeyID string) sdk.Error {
	return sdk.NewError(Codespace, CodeRevokedCertificateDoesNotExist,
		fmt.Sprintf("No revoked X509 certificate associated with the "+
			"combination of subject=%v and subjectKeyID=%v on the ledger", subject, subjectKeyID))
}

func ErrInappropriateCertificateType(error interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeInappropriateCertificateType, fmt.Sprintf("%v", error))
}

func ErrCodeInvalidCertificate(error interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeInvalidCertificate, fmt.Sprintf("%v", error))
}
