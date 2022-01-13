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
)

func (k msgServer) ApproveAddX509RootCert(goCtx context.Context, msg *types.MsgApproveAddX509RootCert) (*types.MsgApproveAddX509RootCertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}

	// check if signer has root certificate approval role
	if !k.dclauthKeeper.HasRole(ctx, signerAddr, types.RootCertificateApprovalRole) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"MsgApproveAddX509RootCert transaction should be signed by an account with the \"%s\" role",
			types.RootCertificateApprovalRole,
		)
	}

	// get proposed certificate
	proposedCertificate, found := k.GetProposedCertificate(ctx, msg.Subject, msg.SubjectKeyId)
	if !found {
		return nil, types.NewErrProposedCertificateDoesNotExist(msg.Subject, msg.SubjectKeyId)
	}

	// check if proposed certificate already has approval form signer
	if proposedCertificate.HasApprovalFrom(signerAddr) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"Certificate associated with subject=%v and subjectKeyID=%v combination "+
				"already has approval from=%v",
			msg.Subject, msg.SubjectKeyId, msg.Signer,
		)
	}

	// append approval
	proposedCertificate.Approvals = append(proposedCertificate.Approvals, signerAddr.String())

	// check if proposed certificate has enough approvals
	if len(proposedCertificate.Approvals) == types.RootCertificateApprovals {
		// create approved certificate
		rootCertificate := types.NewRootCertificate(
			proposedCertificate.PemCert,
			proposedCertificate.Subject,
			proposedCertificate.SubjectKeyId,
			proposedCertificate.SerialNumber,
			proposedCertificate.Owner,
		)

		// add approved certificate to stored list of certificates with the same Subject/SubjectKeyId combination
		k.AddApprovedCertificate(ctx, rootCertificate)

		// delete proposed certificate
		k.RemoveProposedCertificate(ctx, msg.Subject, msg.SubjectKeyId)

		// add to root certificates index
		certID := types.CertificateIdentifier{
			Subject:      rootCertificate.Subject,
			SubjectKeyId: rootCertificate.SubjectKeyId,
		}
		k.AddApprovedRootCertificate(ctx, certID)

		// add to subject -> subject key ID map
		k.AddApprovedCertificateBySubject(ctx, rootCertificate.Subject, rootCertificate.SubjectKeyId)
	} else {
		// update proposed certificate
		k.SetProposedCertificate(ctx, proposedCertificate)
	}

	return &types.MsgApproveAddX509RootCertResponse{}, nil
}
