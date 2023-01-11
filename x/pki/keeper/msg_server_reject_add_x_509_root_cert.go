package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func (k msgServer) RejectAddX509RootCert(goCtx context.Context, msg *types.MsgRejectAddX509RootCert) (*types.MsgRejectAddX509RootCertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}

	// get proposed certificate
	proposedCertificate, found := k.GetProposedCertificate(ctx, msg.Subject, msg.SubjectKeyId)
	if !found {
		return nil, types.NewErrProposedCertificateDoesNotExist(msg.Subject, msg.SubjectKeyId)
	}

	// check if signer has root certificate approval role
	if !k.dclauthKeeper.HasRole(ctx, signerAddr, types.RootCertificateApprovalRole) {
		// Remove proposed certificate if there are no rejects and approvals
		if proposedCertificate.Owner == msg.Signer && len(proposedCertificate.Approvals) == 0 && len(proposedCertificate.Rejects) == 0 {
			k.RemoveProposedCertificate(ctx, msg.Subject, msg.SubjectKeyId)
			k.RemoveUniqueCertificate(ctx, proposedCertificate.Subject, proposedCertificate.SerialNumber)

			return &types.MsgRejectAddX509RootCertResponse{}, nil
		}

		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"MsgApproveAddX509RootCert transaction should be signed by an account with the \"%s\" role",
			types.RootCertificateApprovalRole,
		)
	}

	// check if proposed certificate already has reject approval form signer
	if proposedCertificate.HasRejectFrom(signerAddr.String()) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"Certificate associated with subject=%v and subjectKeyID=%v combination "+
				"already has reject approval from=%v",
			msg.Subject, msg.SubjectKeyId, msg.Signer,
		)
	}

	// append approval
	grant := types.Grant{
		Address: signerAddr.String(),
		Time:    msg.Time,
		Info:    msg.Info,
	}

	// check if proposed certificate has approval form signer
	if proposedCertificate.HasApprovalFrom(signerAddr.String()) {
		// Remove proposed certificate if there are no rejects and other approvals
		if proposedCertificate.Owner == msg.Signer && len(proposedCertificate.Approvals) == 1 && len(proposedCertificate.Rejects) == 0 {
			k.RemoveProposedCertificate(ctx, msg.Subject, msg.SubjectKeyId)
			k.RemoveUniqueCertificate(ctx, proposedCertificate.Subject, proposedCertificate.SerialNumber)

			return &types.MsgRejectAddX509RootCertResponse{}, nil
		}

		// Remove approval from the list of approvals
		for i, other := range proposedCertificate.Approvals {
			if other.Address == grant.Address {
				proposedCertificate.Approvals = append(proposedCertificate.Approvals[:i], proposedCertificate.Approvals[i+1:]...)
			}
		}
	}
	proposedCertificate.Rejects = append(proposedCertificate.Rejects, &grant)

	// check if proposed certificate has enough approvals
	if len(proposedCertificate.Rejects) >= k.CertificateRejectApprovalsCount(ctx, k.dclauthKeeper) {
		// create rejected certificate
		rejectedRootCertificate := types.RejectedCertificate{
			Subject:      proposedCertificate.Subject,
			SubjectKeyId: proposedCertificate.SubjectKeyId,
			Certs: []*types.Certificate{
				{
					PemCert:       proposedCertificate.PemCert,
					SerialNumber:  proposedCertificate.SerialNumber,
					Owner:         proposedCertificate.Owner,
					Subject:       proposedCertificate.Subject,
					SubjectAsText: proposedCertificate.SubjectAsText,
					SubjectKeyId:  proposedCertificate.SubjectKeyId,
					Approvals:     proposedCertificate.Approvals,
					Rejects:       proposedCertificate.Rejects,
				},
			},
		}

		k.SetRejectedCertificate(ctx, rejectedRootCertificate)
		k.RemoveProposedCertificate(ctx, msg.Subject, msg.SubjectKeyId)
		k.RemoveUniqueCertificate(ctx, proposedCertificate.Subject, proposedCertificate.SerialNumber)
	} else {
		// update proposed certificate
		k.SetProposedCertificate(ctx, proposedCertificate)
	}

	return &types.MsgRejectAddX509RootCertResponse{}, nil
}
