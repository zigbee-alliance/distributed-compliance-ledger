package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

// SetVendorProducts set a specific vendorProducts in the store from its index
func (k Keeper) SetVendorProducts(ctx sdk.Context, vendorProducts types.VendorProducts) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VendorProductsKeyPrefix))
	b := k.cdc.MustMarshal(&vendorProducts)
	store.Set(types.VendorProductsKey(
		vendorProducts.Vid,
	), b)
}

// GetVendorProducts returns a vendorProducts from its index
func (k Keeper) GetVendorProducts(
	ctx sdk.Context,
	vid int32,

) (val types.VendorProducts, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VendorProductsKeyPrefix))

	b := store.Get(types.VendorProductsKey(
		vid,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveVendorProducts removes a vendorProducts from the store
func (k Keeper) RemoveVendorProducts(
	ctx sdk.Context,
	vid int32,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VendorProductsKeyPrefix))
	store.Delete(types.VendorProductsKey(
		vid,
	))
}

// GetAllVendorProducts returns all vendorProducts
func (k Keeper) GetAllVendorProducts(ctx sdk.Context) (list []types.VendorProducts) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VendorProductsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.VendorProducts
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
