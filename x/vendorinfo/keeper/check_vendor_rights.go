package keeper

import (
	"fmt"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func checkAddVendorRights(ctx sdk.Context, k Keeper, signer sdk.AccAddress, vid int32) error { //nolint:goconst
	// sender must have Vendor role and VendorID or VendorAdmin role to add a new vendorinfo
	isVendor := k.dclauthKeeper.HasRole(ctx, signer, dclauthtypes.Vendor)
	isVendorAdmin := k.dclauthKeeper.HasRole(ctx, signer, dclauthtypes.VendorAdmin)

	if !isVendor && !isVendorAdmin {
		return errors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("MsgAddVendorInfo transaction should be "+
			"signed by an account with the %s or %s roles", dclauthtypes.Vendor, dclauthtypes.VendorAdmin))
	}

	// Checks if the msg creator is the same as the current owner
	if isVendor && !k.dclauthKeeper.HasVendorID(ctx, signer, vid) {
		return errors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("MsgAddVendorInfo transaction should be "+
			"signed by a vendor account associated with the vendorID %v ", vid))
	}

	return nil
}

func checkUpdateVendorRights(ctx sdk.Context, k Keeper, signer sdk.AccAddress, vid int32) error {
	// sender must have Vendor role and VendorID or VendorAdmin role to update a vendorinfo
	isVendor := k.dclauthKeeper.HasRole(ctx, signer, dclauthtypes.Vendor)
	isVendorAdmin := k.dclauthKeeper.HasRole(ctx, signer, dclauthtypes.VendorAdmin)

	if !isVendor && !isVendorAdmin {
		return errors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("MsgAddVendorInfo transaction should be "+
			"signed by an account with the %s or %s roles", dclauthtypes.Vendor, dclauthtypes.VendorAdmin))
	}

	// Checks if the msg creator is the same as the current owner
	if isVendor && !k.dclauthKeeper.HasVendorID(ctx, signer, vid) {
		return errors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("MsgAddVendorInfo transaction should be "+
			"signed by a vendor account associated with the vendorID %v ", vid))
	}

	return nil
}
