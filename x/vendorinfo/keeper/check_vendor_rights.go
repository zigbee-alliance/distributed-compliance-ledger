package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func checkAddVendorRights(ctx sdk.Context, k Keeper, signer sdk.AccAddress, vid int32) error {
	// sender must have Vendor role and VendorID to add new model
	if !k.dclauthKeeper.HasRole(ctx, signer, dclauthtypes.Vendor) {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("MsgAddVendorInfo transaction should be "+
			"signed by an account with the %s role", dclauthtypes.Vendor))
	}

	// Checks if the the msg creator is the same as the current owner
	if !k.dclauthKeeper.HasVendorID(ctx, signer, vid) {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("MsgAddVendorInfo transaction should be "+
			"signed by a vendor account associated with the vendorID %v ", vid))
	}

	return nil
}

func checkUpdateVendorRights(ctx sdk.Context, k Keeper, signer sdk.AccAddress, vid int32) error {
	// Checks if the the msg creator is the same as the current owner
	if !k.dclauthKeeper.HasVendorID(ctx, sdk.AccAddress(signer), vid) {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("MsgAddVendorInfo transaction should be "+
			"signed by a vendor account associated with the vendorID %v ", vid))
	}

	return nil
}
