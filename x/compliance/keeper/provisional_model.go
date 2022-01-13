package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
)

// SetProvisionalModel set a specific provisionalModel in the store from its index.
func (k Keeper) SetProvisionalModel(ctx sdk.Context, provisionalModel types.ProvisionalModel) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProvisionalModelKeyPrefix))
	b := k.cdc.MustMarshal(&provisionalModel)
	store.Set(types.ProvisionalModelKey(
		provisionalModel.Vid,
		provisionalModel.Pid,
		provisionalModel.SoftwareVersion,
		provisionalModel.CertificationType,
	), b)
}

// GetProvisionalModel returns a provisionalModel from its index.
func (k Keeper) GetProvisionalModel(
	ctx sdk.Context,
	vid int32,
	pid int32,
	softwareVersion uint32,
	certificationType string,

) (val types.ProvisionalModel, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProvisionalModelKeyPrefix))

	b := store.Get(types.ProvisionalModelKey(
		vid,
		pid,
		softwareVersion,
		certificationType,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveProvisionalModel removes a provisionalModel from the store.
func (k Keeper) RemoveProvisionalModel(
	ctx sdk.Context,
	vid int32,
	pid int32,
	softwareVersion uint32,
	certificationType string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProvisionalModelKeyPrefix))
	store.Delete(types.ProvisionalModelKey(
		vid,
		pid,
		softwareVersion,
		certificationType,
	))
}

// GetAllProvisionalModel returns all provisionalModel.
func (k Keeper) GetAllProvisionalModel(ctx sdk.Context) (list []types.ProvisionalModel) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProvisionalModelKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ProvisionalModel
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
