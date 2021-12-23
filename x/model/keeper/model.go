package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

// SetModel set a specific model in the store from its index
func (k Keeper) SetModel(ctx sdk.Context, model types.Model) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ModelKeyPrefix))
	b := k.cdc.MustMarshal(&model)
	store.Set(types.ModelKey(
		model.Vid,
		model.Pid,
	), b)
}

// GetModel returns a model from its index
func (k Keeper) GetModel(
	ctx sdk.Context,
	vid int32,
	pid int32,

) (val types.Model, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ModelKeyPrefix))

	b := store.Get(types.ModelKey(
		vid,
		pid,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveModel removes a model from the store
func (k Keeper) RemoveModel(
	ctx sdk.Context,
	vid int32,
	pid int32,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ModelKeyPrefix))
	store.Delete(types.ModelKey(
		vid,
		pid,
	))
}

// GetAllModel returns all model
func (k Keeper) GetAllModel(ctx sdk.Context) (list []types.Model) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ModelKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Model
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
