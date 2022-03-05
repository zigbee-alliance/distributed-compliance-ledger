package keeper

import (
	"fmt"
	"math"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey

		dclauthKeeper types.DclauthKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,

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
	return int(math.Round(types.DisableValidatorPercent * float64(k.dclauthKeeper.CountAccountsWithRole(ctx, types.VoteForDisableValidatorRole))))
}
