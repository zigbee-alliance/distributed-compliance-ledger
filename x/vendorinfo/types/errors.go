package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/vendorinfo module sentinel errors.
var (
	DefaultCodespace string = ModuleName

	CodeVendorDoesNotExist              = sdkerrors.Register(ModuleName, 701, "Code vendor does not exist")
	CodeMissingVendorIDForVendorAccount = sdkerrors.Register(ModuleName, 702, "Code missing vendor id for vendor account")
	CodeVendorInfoAlreadyExists         = sdkerrors.Register(ModuleName, 703, "Code vendorinfo already exists")
)
