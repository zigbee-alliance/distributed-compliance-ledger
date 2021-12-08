package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/types"
)

// SetNewVendorInfo set a specific newVendorInfo in the store from its index
func (k Keeper) SetNewVendorInfo(ctx sdk.Context, newVendorInfo types.NewVendorInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NewVendorInfoKeyPrefix))
	b := k.cdc.MustMarshal(&newVendorInfo)
	store.Set(types.NewVendorInfoKey(
		newVendorInfo.Index,
	), b)
}

// GetNewVendorInfo returns a newVendorInfo from its index
func (k Keeper) GetNewVendorInfo(
	ctx sdk.Context,
	index string,

) (val types.NewVendorInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NewVendorInfoKeyPrefix))

	b := store.Get(types.NewVendorInfoKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNewVendorInfo removes a newVendorInfo from the store
func (k Keeper) RemoveNewVendorInfo(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NewVendorInfoKeyPrefix))
	store.Delete(types.NewVendorInfoKey(
		index,
	))
}

// GetAllNewVendorInfo returns all newVendorInfo
func (k Keeper) GetAllNewVendorInfo(ctx sdk.Context) (list []types.NewVendorInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NewVendorInfoKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NewVendorInfo
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
