package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func (k msgServer) AssignVid(goCtx context.Context, msg *types.MsgAssignVid) (*types.MsgAssignVidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}

	// check if signer has vendor role
	if !k.dclauthKeeper.HasRole(ctx, signerAddr, dclauthtypes.Vendor) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"AssignVid transaction should be signed by an account with the \"%s\" role",
			dclauthtypes.Vendor,
		)
	}

	// get corresponding approved certificates
	certificates, found := k.GetApprovedCertificates(ctx, msg.Subject, msg.SubjectKeyId)
	if !found {
		return nil, pkitypes.NewErrCertificateDoesNotExist(msg.Subject, msg.SubjectKeyId)
	}

	rootCertificate := certificates.Certs[0]

	// fail if certificates are not root
	if !rootCertificate.IsRoot {
		return nil, pkitypes.NewErrInappropriateCertificateType(
			fmt.Sprintf("Inappropriate Certificate Type: Certificate with subject=%v and subjectKeyID=%v "+
				"is not a root certificate.", msg.Subject, msg.SubjectKeyId),
		)
	}

	// check that the certificate VID has not been set
	if rootCertificate.Vid != 0 {
		return nil, pkitypes.NewErrNotEmptyVid("Vendor ID (VID) is already present in the certificate")
	}

	rootCertificate.Vid = msg.Vid

	k.SetApprovedCertificates(ctx, certificates)

	return &types.MsgAssignVidResponse{}, nil
}
