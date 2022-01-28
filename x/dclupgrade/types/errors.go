package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/dclupgrade module sentinel errors
var (
	ErrProposedUpgradeAlreadyExists = sdkerrors.Register(ModuleName, 1101, "proposed upgrade already exists")
)

func NewErrProposedUpgradeAlreadyExists(name interface{}) error {
	return sdkerrors.Wrapf(
		ErrProposedUpgradeAlreadyExists,
		"Proposed upgrade with name=%v already exists on the ledger",
		name,
	)
}
