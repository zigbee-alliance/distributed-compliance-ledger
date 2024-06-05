package keeper

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func checkModelRights(ctx sdk.Context, k Keeper, signer sdk.AccAddress, vid int32, pid int32, message string) error {
	// sender must have Vendor role to add new model
	if !k.dclauthKeeper.HasRole(ctx, signer, dclauthtypes.Vendor) {
		return errors.Wrapf(sdkerrors.ErrUnauthorized, "%s transaction should be "+
			"signed by an account with the %s role", message, dclauthtypes.Vendor)
	}

	if !k.dclauthKeeper.HasVendorID(ctx, signer, vid) {
		return errors.Wrapf(sdkerrors.ErrUnauthorized, "%s transaction should be "+
			"signed by a vendor account containing the vendorID %v ", message, vid)
	}

	if !k.dclauthKeeper.HasRightsToChange(ctx, signer, pid) {
		return errors.Wrapf(sdkerrors.ErrUnauthorized, "%s transaction should be "+
			"signed by a vendor account who has rights to modify the productID %v ", message, pid)
	}

	return nil
}
