// Copyright 2022 DSR Corporation
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

package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/x509"
)

func (k msgServer) AddX509Cert(goCtx context.Context, msg *types.MsgAddX509Cert) (*types.MsgAddX509CertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// decode pem certificate
	x509Certificate, err := x509.DecodeX509Certificate(msg.Cert)
	if err != nil {
		return nil, types.NewErrInvalidCertificate(err)
	}

	// fail if certificate is self-signed
	if x509Certificate.IsSelfSigned() {
		return nil, types.NewErrInappropriateCertificateType(
			"Inappropriate Certificate Type: Passed certificate is self-signed, " +
				"so it cannot be added to the system as a non-root certificate. " +
				"To propose adding a root certificate please use `PROPOSE_ADD_X509_ROOT_CERT` transaction.")
	}

	// check if certificate with Issuer/Serial Number combination already exists
	if k.IsUniqueCertificatePresent(ctx, x509Certificate.Issuer, x509Certificate.SerialNumber) {
		return nil, types.NewErrCertificateAlreadyExists(x509Certificate.Issuer, x509Certificate.SerialNumber)
	}

	// Get list of certificates for Subject / Subject Key Id combination
	certificates, found := k.GetApprovedCertificates(ctx, x509Certificate.Subject, x509Certificate.SubjectKeyID)
	if found {
		// Issuer and authorityKeyID must be the same as ones of exisiting certificates with the same subject and
		// subjectKeyID. Since new certificate is not self-signed, we have to ensure that the exisiting certificates
		// are not self-signed too, consequently are non-root certificates, before to match issuer and authorityKeyID.
		if certificates.Certs[0].IsRoot || x509Certificate.Issuer != certificates.Certs[0].Issuer ||
			x509Certificate.AuthorityKeyID != certificates.Certs[0].AuthorityKeyId {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
				"Issuer and authorityKeyID of new certificate with subject=%v and subjectKeyID=%v "+
					"must be the same as ones of existing certificates with the same subject and subjectKeyID",
				x509Certificate.Subject, x509Certificate.SubjectKeyID,
			)
		}

		// signer must be same as owner of existing certificates
		if msg.Signer != certificates.Certs[0].Owner {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
				"Only owner of existing certificates with subject=%v and subjectKeyID=%v "+
					"can add new certificate with the same subject and subjectKeyID",
				x509Certificate.Subject, x509Certificate.SubjectKeyID,
			)
		}
	}

	// Valid certificate chain must be built for new certificate
	rootCertificateSubject, rootCertificateSubjectKeyID, err := k.verifyCertificate(ctx, x509Certificate)
	if err != nil {
		return nil, err
	}

	// create new certificate
	certificate := types.NewNonRootCertificate(
		msg.Cert,
		x509Certificate.Subject,
		x509Certificate.SubjectKeyID,
		x509Certificate.SerialNumber,
		x509Certificate.Issuer,
		x509Certificate.AuthorityKeyID,
		rootCertificateSubject,
		rootCertificateSubjectKeyID,
		msg.Signer,
	)

	// append new certificate to list of certificates with the same Subject/SubjectKeyId combination and store updated list
	k.AddApprovedCertificate(ctx, certificate)

	// add the certificate identifier to the issuer's Child Certificates record
	certificateIdentifier := types.CertificateIdentifier{
		Subject:      certificate.Subject,
		SubjectKeyId: certificate.SubjectKeyId,
	}
	k.AddChildCertificate(ctx, certificate.Issuer, certificate.AuthorityKeyId, certificateIdentifier)

	// register the unique certificate key
	uniqueCertificate := types.UniqueCertificate{
		Issuer:       x509Certificate.Issuer,
		SerialNumber: x509Certificate.SerialNumber,
		Present:      true,
	}
	k.SetUniqueCertificate(ctx, uniqueCertificate)

	// add to subject -> subject key ID map
	k.AddApprovedCertificateBySubject(ctx, certificate.Subject, certificate.SubjectKeyId)

	return &types.MsgAddX509CertResponse{}, nil
}
