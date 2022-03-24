package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

type DclauthKeeper interface {
	// Methods imported from dclauth should be defined here
	HasRole(ctx sdk.Context, addr sdk.AccAddress, roleToCheck types.AccountRole) bool
}
