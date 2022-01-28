package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

type DclauthKeeper interface {
	// Methods imported from dclauth should be defined here
	HasRole(ctx sdk.Context, addr sdk.AccAddress, roleToCheck dclauthtypes.AccountRole) bool
}

type UpgradeKeeper interface {
	// Methods imported from upgrade should be defined here
	ScheduleUpgrade(ctx sdk.Context, plan upgradetypes.Plan) error
}
