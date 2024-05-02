package keeper

import (
	"fmt"
	"math"

	"github.com/cometbft/cometbft/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey storetypes.StoreKey
		memKey   storetypes.StoreKey

		dclauthKeeper types.DclauthKeeper
		upgradeKeeper types.UpgradeKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,

	dclauthKeeper types.DclauthKeeper,
	upgradeKeeper types.UpgradeKeeper,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,

		dclauthKeeper: dclauthKeeper,
		upgradeKeeper: upgradeKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) UpgradeApprovalsCount(ctx sdk.Context) int {
	return int(math.Ceil(types.UpgradeApprovalsPercent * float64(k.dclauthKeeper.CountAccountsWithRole(ctx, types.UpgradeApprovalRole))))
}

func (k Keeper) UpgradeRejectsCount(ctx sdk.Context) int {
	return k.dclauthKeeper.CountAccountsWithRole(ctx, types.UpgradeApprovalRole) - k.UpgradeApprovalsCount(ctx) + 1
}
