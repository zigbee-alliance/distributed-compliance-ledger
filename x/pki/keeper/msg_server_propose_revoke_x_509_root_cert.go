package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func (k msgServer) ProposeRevokeX509RootCert(goCtx context.Context, msg *types.MsgProposeRevokeX509RootCert) (*types.MsgProposeRevokeX509RootCertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if signer has root certificate approval role
	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}
	if !k.dclauthKeeper.HasRole(ctx, signerAddr, types.RootCertificateApprovalRole) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"MsgProposeRevokeX509RootCert transaction should be signed by "+
				"an account with the \"%s\" role",
			types.RootCertificateApprovalRole,
		)
	}

	// check that proposed certificate revocation does not exist yet
	if k.IsProposedCertificateRevocationPresent(ctx, msg.Subject, msg.SubjectKeyId) {
		return nil, types.NewErrProposedCertificateRevocationAlreadyExists(msg.Subject, msg.SubjectKeyId)
	}

	// get corresponding approved certificates
	certificates, found := k.GetApprovedCertificates(ctx, msg.Subject, msg.SubjectKeyId)
	if !found {
		return nil, types.NewErrCertificateDoesNotExist(msg.Subject, msg.SubjectKeyId)
	}

	// fail if certificates are not root
	if !certificates.Certs[0].IsRoot {
		return nil, types.NewErrInappropriateCertificateType(
			fmt.Sprintf("Inappropriate Certificate Type: Certificate with subject=%v and subjectKeyID=%v "+
				"is not a root certificate.", msg.Subject, msg.SubjectKeyId),
		)
	}

	// create new proposed certificate revocation with approval from signer
	grant := types.Grant{
		Address: msg.Signer,
		Time:    msg.Time,
		Info:    msg.Info,
	}
	revocation := types.ProposedCertificateRevocation{
		Subject:      msg.Subject,
		SubjectKeyId: msg.SubjectKeyId,
		Revocations:  []types.Grant{grant},
	}

	// store proposed certificate revocation
	k.SetProposedCertificateRevocation(ctx, revocation)

	return &types.MsgProposeRevokeX509RootCertResponse{}, nil
}
