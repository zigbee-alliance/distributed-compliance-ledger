package types

// DONTCOVER

import (
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/vendorinfo module sentinel errors
var (
	DefaultCodespace string = ModuleName

	CodeVendorDoesNotExist              uint32 = 701
	CodeMissingVendorIDForVendorAccount uint32 = 702
	CodeVendorInfoAlreadyExists         uint32 = 703
)

func ErrVendorInfoDoesNotExist(vendorID uint16) *sdkerrors.Error {
	return sdkerrors.Register(DefaultCodespace, CodeVendorDoesNotExist,
		fmt.Sprintf("Vendor Account with VendorID %v does not exist on the ledger", vendorID))
}

func ErrMissingVendorIDForVendorAccount() *sdkerrors.Error {
	return sdkerrors.Register(DefaultCodespace, CodeMissingVendorIDForVendorAccount,
		"No Vendor ID is provided in the Vendor Role for the new account")
}

func ErrVendorInfoAlreadyExists(vendorID interface{}) *sdkerrors.Error {
	return sdkerrors.Register(DefaultCodespace, CodeVendorInfoAlreadyExists,
		fmt.Sprintf("Vendor info associated with VendorID=%v already exists on the ledger", vendorID))
}
