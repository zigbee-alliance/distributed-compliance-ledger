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

	// check if proposed certificate already has reject approval form signer
	if proposedCertificate.HasRejectFrom(signerAddr.String()) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"Certificate associated with subject=%v and subjectKeyID=%v combination "+
				"already has reject approval from=%v",
			msg.Subject, msg.SubjectKeyId, msg.Signer,
		)
	}

	// check if proposed certificate already has approval form signer
	if proposedCertificate.HasApprovalFrom(signerAddr.String()) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"Certificate associated with subject=%v and subjectKeyID=%v combination "+
				"already has approval from=%v",
			msg.Subject, msg.SubjectKeyId, msg.Signer,
		)
	}

	// append approval
	grant := types.Grant{
		Address: signerAddr.String(),
		Time:    msg.Time,
		Info:    msg.Info,
	}
	proposedCertificate.Rejects = append(proposedCertificate.Rejects, &grant)

	// check if proposed certificate has enough approvals
	if len(proposedCertificate.Rejects) == k.CertificateRejectApprovalsCount(ctx, k.dclauthKeeper) {
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
