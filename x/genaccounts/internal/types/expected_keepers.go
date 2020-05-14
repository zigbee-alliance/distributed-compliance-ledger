package types

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AccountKeeper defines the expected account keeper (noalias)
type AuthKeeper interface {
	NewAccount(sdk.Context, auth.Account) auth.Account
	SetAccount(sdk.Context, auth.Account)
	IterateAccounts(ctx sdk.Context, process func(auth.Account) (stop bool))
}
