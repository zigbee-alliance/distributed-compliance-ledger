package ante

import (
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
)

// TODO issue 99: add GetParams and GetModuleAddress stubs in dclauth.AccountKeeper.
type AccountKeeper authante.AccountKeeper

// which is
/*
// AccountKeeper defines the contract needed for AccountKeeper related APIs.
// Interface provides support to use non-sdk AccountKeeper for AnteHandler's decorators.
type AccountKeeper interface {
	GetParams(ctx sdk.Context) (params types.Params)
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) dclauthtypes.Account
	SetAccount(ctx sdk.Context, acc dclauthtypes.Account)
	GetModuleAddress(moduleName string) sdk.AccAddress
}
*/
