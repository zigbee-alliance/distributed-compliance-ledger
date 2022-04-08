package keeper

import (
	"fmt"
	"math"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// TODO issue 99: these getters were initially created
//		for tests needs: to link dependent keepers,
//		need to explore the alternatives
func (k Keeper) StoreKey() sdk.StoreKey {
	return k.storeKey
}

func (k Keeper) MemKey() sdk.StoreKey {
	return k.memKey
}

func (k Keeper) AccountApprovalsCount(ctx sdk.Context) int {
	return int(math.Round(types.AccountApprovalsPercent * float64(k.CountAccountsWithRole(ctx, types.Trustee))))
}

func (k Keeper) AccountRejectApprovalsCount(ctx sdk.Context) int {
	return int(math.Round(types.AccountRejectApprovalsPercent) * float64(k.CountAccountsWithRole(ctx, types.Trustee)))
}
