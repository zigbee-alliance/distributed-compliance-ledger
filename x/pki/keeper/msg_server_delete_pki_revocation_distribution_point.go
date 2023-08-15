package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func (k msgServer) DeletePkiRevocationDistributionPoint(goCtx context.Context, msg *types.MsgDeletePkiRevocationDistributionPoint) (*types.MsgDeletePkiRevocationDistributionPointResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, pkitypes.NewErrInvalidAddress(err)
	}

	// check if signer has vendor role
	if !k.dclauthKeeper.HasRole(ctx, signerAddr, dclauthtypes.Vendor) {
		return nil, pkitypes.NewErrUnauthorizedRole("MsgDeletePkiRevocationDistributionPoint", dclauthtypes.Vendor)
	}

	// compare VID in message and Vendor acount
	signerAccount, _ := k.dclauthKeeper.GetAccountO(ctx, signerAddr)
	if signerAccount.VendorID != msg.Vid {
		return nil, pkitypes.NewErrMessageVidNotEqualAccountVid(msg.Vid, signerAccount.VendorID)
	}

	pkiRevocationDistributionPoint, isFound := k.GetPkiRevocationDistributionPoint(ctx, msg.Vid, msg.Label, msg.IssuerSubjectKeyID)
	if !isFound {
		return nil, pkitypes.NewErrPkiRevocationDistributionPointDoesNotExists(msg.Vid, msg.Label, msg.IssuerSubjectKeyID)
	}

	k.RemovePkiRevocationDistributionPoint(ctx, msg.Vid, msg.Label, msg.IssuerSubjectKeyID)
	k.RemovePkiRevocationDistributionPointBySubjectKeyID(ctx, pkiRevocationDistributionPoint)

	return &types.MsgDeletePkiRevocationDistributionPointResponse{}, nil
}
