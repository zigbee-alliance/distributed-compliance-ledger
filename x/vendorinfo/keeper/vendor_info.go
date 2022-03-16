package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/types"
)

// SetVendorInfo set a specific vendorInfo in the store from its index.
func (k Keeper) SetVendorInfo(ctx sdk.Context, vendorInfo types.VendorInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VendorInfoKeyPrefix))
	b := k.cdc.MustMarshal(&vendorInfo)
	store.Set(types.VendorInfoKey(
		vendorInfo.VendorID,
	), b)
}

// GetVendorInfo returns a vendorInfo from its index.
func (k Keeper) GetVendorInfo(
	ctx sdk.Context,
	vendorID int32,

) (val types.VendorInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VendorInfoKeyPrefix))

	b := store.Get(types.VendorInfoKey(
		vendorID,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// RemoveVendorInfo removes a vendorInfo from the store.
func (k Keeper) RemoveVendorInfo(
	ctx sdk.Context,
	vendorID int32,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VendorInfoKeyPrefix))
	store.Delete(types.VendorInfoKey(
		vendorID,
	))
}

// GetAllVendorInfo returns all vendorInfo.
func (k Keeper) GetAllVendorInfo(ctx sdk.Context) (list []types.VendorInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VendorInfoKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var val types.VendorInfo
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
