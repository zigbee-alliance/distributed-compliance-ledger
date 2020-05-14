package types

import (
	"encoding/json"

	abci "github.com/tendermint/tendermint/abci/types"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ValidatorKeeper defines the expected validator keeper.
type ValidatorKeeper interface {
	ApplyAndReturnValidatorSetUpdates(sdk.Context) (updates []abci.ValidatorUpdate)
}

type AuthKeeper interface {
	NewAccount(sdk.Context, auth.Account) auth.Account
	SetAccount(sdk.Context, auth.Account)
	IterateAccounts(ctx sdk.Context, process func(auth.Account) (stop bool))
// AccountKeeper defines the expected account keeper.
}

// GenesisAccountsIterator defines the expected iterating genesis accounts object.
type GenesisAccountsIterator interface {
	IterateGenesisAccounts(
		cdc *codec.Codec,
		appGenesis map[string]json.RawMessage,
		iterateFn func(auth.Account) (stop bool),
	)
}
