package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/dclupgrade module sentinel errors.
var (
	ErrProposedUpgradeAlreadyExists = errors.Register(ModuleName, 801, "proposed upgrade already exists")
	ErrProposedUpgradeDoesNotExist  = errors.Register(ModuleName, 802, "proposed upgrade does not exist")
	ErrApprovedUpgradeAlreadyExists = errors.Register(ModuleName, 803, "approved upgrade already exists")
)

func NewErrProposedUpgradeAlreadyExists(name interface{}) error {
	return errors.Wrapf(
		ErrProposedUpgradeAlreadyExists,
		"Proposed upgrade with name=%v already exists on the ledger",
		name,
	)
}

func NewErrProposedUpgradeDoesNotExist(name interface{}) error {
	return errors.Wrapf(
		ErrProposedUpgradeDoesNotExist,
		"Proposed upgrade with name=%v does not exist on the ledger",
		name,
	)
}

func NewErrApprovedUpgradeAlreadyExists(name interface{}) error {
	return errors.Wrapf(
		ErrApprovedUpgradeAlreadyExists,
		"Approved upgrade with name=%v already exists on the ledger",
		name,
	)
}
