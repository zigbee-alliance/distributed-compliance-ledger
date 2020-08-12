package types

//nolint:goimports
import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	abci "github.com/tendermint/tendermint/abci/types"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ValidatorKeeper defines the expected validator keeper.
type ValidatorKeeper interface {
	ApplyAndReturnValidatorSetUpdates(sdk.Context) (updates []abci.ValidatorUpdate)
}

// AccountKeeper defines the expected account keeper.
type AuthKeeper interface {
	GetNextAccountNumber(sdk.Context) uint64
	SetAccount(sdk.Context, auth.Account)
	IterateAccounts(ctx sdk.Context, process func(auth.Account) (stop bool))
}

// GenesisAccountsIterator defines the expected iterating genesis accounts object.
type GenesisAccountsIterator interface {
	IterateGenesisAccounts(
		cdc *codec.Codec,
		appGenesis map[string]json.RawMessage,
		iterateFn func(auth.Account) (stop bool),
	)
}
