package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
)

// SetVendorInfoType set a specific vendorInfoType in the store from its index
func (k Keeper) SetVendorInfoType(ctx sdk.Context, vendorInfoType types.VendorInfoType) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VendorInfoTypeKeyPrefix))
	b := k.cdc.MustMarshal(&vendorInfoType)
	store.Set(types.VendorInfoTypeKey(
        vendorInfoType.Index,
    ), b)
}

// GetVendorInfoType returns a vendorInfoType from its index
func (k Keeper) GetVendorInfoType(
    ctx sdk.Context,
    index string,
    
) (val types.VendorInfoType, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VendorInfoTypeKeyPrefix))

	b := store.Get(types.VendorInfoTypeKey(
        index,
    ))
    if b == nil {
        return val, false
    }

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveVendorInfoType removes a vendorInfoType from the store
func (k Keeper) RemoveVendorInfoType(
    ctx sdk.Context,
    index string,
    
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VendorInfoTypeKeyPrefix))
	store.Delete(types.VendorInfoTypeKey(
	    index,
    ))
}

// GetAllVendorInfoType returns all vendorInfoType
func (k Keeper) GetAllVendorInfoType(ctx sdk.Context) (list []types.VendorInfoType) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VendorInfoTypeKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.VendorInfoType
		k.cdc.MustUnmarshal(iterator.Value(), &val)
        list = append(list, val)
	}

    return
}
