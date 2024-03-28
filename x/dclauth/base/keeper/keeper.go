package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	basetypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey storetypes.StoreKey
		memKey   storetypes.StoreKey

		paramSubspace paramtypes.Subspace
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", basetypes.ModuleName))
}

func (k Keeper) decodeAccount(bz []byte) basetypes.AccountI {
	acc, err := k.UnmarshalAccount(bz)
	if err != nil {
		panic(err)
	}

	return acc
}

// UnmarshalAccount returns an Account interface from raw encoded account
// bytes of a Proto-based Account type.
func (k Keeper) UnmarshalAccount(bz []byte) (basetypes.AccountI, error) {
	var acc basetypes.AccountI

	return acc, k.cdc.UnmarshalInterface(bz, &acc)
}
