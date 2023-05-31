package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func (k msgServer) DeletePkiRevocationDistributionPoint(goCtx context.Context, msg *types.MsgDeletePkiRevocationDistributionPoint) (*types.MsgDeletePkiRevocationDistributionPointResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}

	signerAccount, _ := k.dclauthKeeper.GetAccountO(ctx, signerAddr)

	// check if signer has vendor role
	if !k.dclauthKeeper.HasRole(ctx, signerAddr, dclauthtypes.Vendor) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"MsgDeletePkiRevocationDistributionPoint transaction should be signed by an account with the \"%s\" role",
			dclauthtypes.Vendor,
		)
	}

	pkiRevocationDistributionPoint, isFound := k.GetPkiRevocationDistributionPoint(ctx, msg.Vid, msg.Label, msg.IssuerSubjectKeyID)
	if !isFound {
		return nil, pkitypes.NewErrPkiRevocationDistributionPointDoesNotExists("PKI revocation distribution point does not exist")
	}

	if signerAccount.VendorID != pkiRevocationDistributionPoint.Vid {
		return nil, pkitypes.NewErrCRLSignerCertificateVidNotEqualAccountVid("CRL signer Certificate's vid must be equal to signer account's vid")
	}

	k.RemovePkiRevocationDistributionPoint(ctx, msg.Vid, msg.Label, msg.IssuerSubjectKeyID)

	return &types.MsgDeletePkiRevocationDistributionPointResponse{}, nil
}
