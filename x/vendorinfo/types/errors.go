package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/vendorinfo module sentinel errors.
var (
	DefaultCodespace = ModuleName

	CodeVendorDoesNotExist              = errors.Register(ModuleName, 701, "Code vendor does not exist")
	CodeMissingVendorIDForVendorAccount = errors.Register(ModuleName, 702, "Code missing vendor id for vendor account")
	CodeVendorInfoAlreadyExists         = errors.Register(ModuleName, 703, "Code vendorinfo already exists")
)
