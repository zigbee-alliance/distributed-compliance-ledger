package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

// SetModelVersion set a specific modelVersion in the store from its index
func (k Keeper) SetModelVersion(ctx sdk.Context, modelVersion types.ModelVersion) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ModelVersionKeyPrefix))
	b := k.cdc.MustMarshal(&modelVersion)
	store.Set(types.ModelVersionKey(
		modelVersion.Vid,
		modelVersion.Pid,
		modelVersion.SoftwareVersion,
	), b)
}

// GetModelVersion returns a modelVersion from its index
func (k Keeper) GetModelVersion(
	ctx sdk.Context,
	vid int32,
	pid int32,
	softwareVersion uint64,

) (val types.ModelVersion, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ModelVersionKeyPrefix))

	b := store.Get(types.ModelVersionKey(
		vid,
		pid,
		softwareVersion,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveModelVersion removes a modelVersion from the store
func (k Keeper) RemoveModelVersion(
	ctx sdk.Context,
	vid int32,
	pid int32,
	softwareVersion uint64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ModelVersionKeyPrefix))
	store.Delete(types.ModelVersionKey(
		vid,
		pid,
		softwareVersion,
	))
}

// GetAllModelVersion returns all modelVersion
func (k Keeper) GetAllModelVersion(ctx sdk.Context) (list []types.ModelVersion) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ModelVersionKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ModelVersion
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
