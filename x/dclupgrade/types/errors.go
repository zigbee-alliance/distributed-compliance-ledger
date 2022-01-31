package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/dclupgrade module sentinel errors
var (
	ErrProposedUpgradeAlreadyExists = sdkerrors.Register(ModuleName, 801, "proposed upgrade already exists")
	ErrProposedUpgradeDoesNotExist  = sdkerrors.Register(ModuleName, 802, "proposed upgrade does not exist")
)

func NewErrProposedUpgradeAlreadyExists(name interface{}) error {
	return sdkerrors.Wrapf(
		ErrProposedUpgradeAlreadyExists,
		"Proposed upgrade with name=%v already exists on the ledger",
		name,
	)
}

func NewErrProposedUpgradeDoesNotExist(name interface{}) error {
	return sdkerrors.Wrapf(
		ErrProposedUpgradeDoesNotExist,
		"Proposed upgrade with name=%v does not exist on the ledger",
		name,
	)
}
