package keeper

import (
	"fmt"
	"math"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey storetypes.StoreKey
		memKey   storetypes.StoreKey

		dclauthKeeper types.DclauthKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,

	dclauthKeeper types.DclauthKeeper,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,

		dclauthKeeper: dclauthKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) DisableValidatorApprovalsCount(ctx sdk.Context) int {
	return int(math.Ceil(types.DisableValidatorPercent * float64(k.dclauthKeeper.CountAccountsWithRole(ctx, types.VoteForDisableValidatorRole))))
}

func (k Keeper) DisableValidatorRejectApprovalsCount(ctx sdk.Context) int {
	return k.dclauthKeeper.CountAccountsWithRole(ctx, types.VoteForDisableValidatorRole) - k.DisableValidatorApprovalsCount(ctx) + 1
}
