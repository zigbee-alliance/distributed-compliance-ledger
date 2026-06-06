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

package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func (k msgServer) RevokeX509Cert(goCtx context.Context, msg *types.MsgRevokeX509Cert) (*types.MsgRevokeX509CertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, pkitypes.NewErrInvalidAddress(err)
	}

	// check if signer has vendor role
	if !k.dclauthKeeper.HasRole(ctx, signerAddr, dclauthtypes.Vendor) {
		return nil, pkitypes.NewErrUnauthorizedRole("MsgRevokeX509Cert", dclauthtypes.Vendor)
	}

	certificates, _ := k.GetApprovedCertificates(ctx, msg.Subject, msg.SubjectKeyId)
	if len(certificates.Certs) == 0 {
		return nil, pkitypes.NewErrCertificateDoesNotExist(msg.Subject, msg.SubjectKeyId)
	}

	if certificates.Certs[0].IsRoot {
		return nil, pkitypes.NewErrMessageExpectedNonRoot(msg.Subject, msg.SubjectKeyId)
	}

	if err := k.EnsureVidMatches(ctx, certificates.Certs[0].Owner, msg.Signer); err != nil {
		return nil, err
	}

	certIdentifier := types.CertificateIdentifier{
		Subject:      msg.Subject,
		SubjectKeyId: msg.SubjectKeyId,
	}

	if msg.SerialNumber != "" {
		err = k.revokeDaCertificateBySerialNumber(ctx, msg.SerialNumber, certificates)
		if err != nil {
			return nil, err
		}
	} else {
		k.revokeDaCertificate(ctx, certIdentifier, certificates)
	}

	if msg.RevokeChild {
		// Remove certificate identifier from issuer's ChildCertificates record
		k.RevokeApprovedChildCertificates(ctx, certIdentifier.Subject, certIdentifier.SubjectKeyId)
	}

	return &types.MsgRevokeX509CertResponse{}, nil
}

func (k msgServer) revokeDaCertificate(ctx sdk.Context, certID types.CertificateIdentifier, certificates types.ApprovedCertificates) {
	// Revoke certificates with given subject/subjectKeyID
	k.AddRevokedCertificates(ctx, types.RevokedCertificates(certificates))
	// Remove certificate from da list
	k.RemoveDaCertificate(ctx, certID.Subject, certID.SubjectKeyId, false)
	// Remove certificate identifier from issuer's ChildCertificates record
	k.RemoveChildCertificate(ctx, certificates.Certs[0].Issuer, certificates.Certs[0].AuthorityKeyId, certID)
}

func (k msgServer) revokeDaCertificateBySerialNumber(
	ctx sdk.Context,
	serialNumber string,
	certificates types.ApprovedCertificates,
) error {
	cert, found := FindCertificateInList(serialNumber, &certificates.Certs)
	if !found {
		return pkitypes.NewErrCertificateBySerialNumberDoesNotExist(certificates.Subject, certificates.SubjectKeyId, serialNumber)
	}

	k.AddRevokedCertificates(ctx, types.RevokedCertificates{
		Subject:       cert.Subject,
		SubjectKeyId:  cert.SubjectKeyId,
		Certs:         []*types.Certificate{cert},
		SchemaVersion: cert.SchemaVersion,
	})

	k.RemoveDaCertificateBySerialNumber(
		ctx,
		certificates.Subject,
		certificates.SubjectKeyId,
		&certificates,
		cert.SerialNumber,
		cert.Issuer,
		false,
	)

	if len(certificates.Certs) == 0 {
		k.RemoveChildCertificate(ctx, cert.Issuer, cert.AuthorityKeyId, types.CertificateIdentifier{
			Subject:      certificates.Subject,
			SubjectKeyId: certificates.SubjectKeyId,
		})
	}

	return nil
}
