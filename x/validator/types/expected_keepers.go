package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

type DclauthKeeper interface {
	// Methods imported from dclauth should be defined here
	HasRole(ctx sdk.Context, addr sdk.AccAddress, roleToCheck dclauthtypes.AccountRole) bool
	CountAccountsWithRole(ctx sdk.Context, roleToCount dclauthtypes.AccountRole) int
	GetAccountO(ctx sdk.Context, address sdk.AccAddress) (val dclauthtypes.Account, found bool)
	SetRevokedAccount(ctx sdk.Context, revokedAccount dclauthtypes.RevokedAccount)
	RemoveAccount(ctx sdk.Context, address sdk.AccAddress)
	MoveAccountToRevokeAccount(
		ctx sdk.Context, accAddr sdk.AccAddress, approvals []*dclauthtypes.Grant, reason dclauthtypes.RevokedAccount_Reason,
	) (*dclauthtypes.RevokedAccount, error)
}
