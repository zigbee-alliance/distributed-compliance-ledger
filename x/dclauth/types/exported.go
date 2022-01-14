package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GenesisAccount defines a genesis account interface that allows account
// address retrieval.
type GenesisAccount interface {
	GetAddress() sdk.AccAddress
	Validate() error
}
