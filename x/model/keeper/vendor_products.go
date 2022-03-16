package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

// SetVendorProducts set a specific vendorProducts in the store from its index.
func (k Keeper) SetVendorProducts(ctx sdk.Context, vendorProducts types.VendorProducts) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VendorProductsKeyPrefix))
	b := k.cdc.MustMarshal(&vendorProducts)
	store.Set(types.VendorProductsKey(
		vendorProducts.Vid,
	), b)
}

// GetVendorProducts returns a vendorProducts from its index.
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

// RemoveVendorProducts removes a vendorProducts from the store.
func (k Keeper) RemoveVendorProducts(
	ctx sdk.Context,
	vid int32,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VendorProductsKeyPrefix))
	store.Delete(types.VendorProductsKey(
		vid,
	))
}

// GetAllVendorProducts returns all vendorProducts.
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

// SetVendorProduct sets a product to existing or new VendorProducts.
func (k Keeper) SetVendorProduct(
	ctx sdk.Context,
	vid int32,
	product types.Product,
) {
	vendorProducts, found := k.GetVendorProducts(ctx, vid)

	if found {
		productFound := false

		for i, value := range vendorProducts.Products {
			if value.Pid == product.Pid {
				vendorProducts.Products[i] = &product
				productFound = true
			}
		}

		if !productFound {
			vendorProducts.Products = append(vendorProducts.Products, &product)
		}
	} else {
		vendorProducts.Vid = vid
		vendorProducts.Products = []*types.Product{&product}
	}

	k.SetVendorProducts(ctx, vendorProducts)
}

// RemoveVendorProduct removes a product from existing VendorProducts.
func (k Keeper) RemoveVendorProduct(
	ctx sdk.Context,
	vid int32,
	pid int32,
) {
	vendorProducts, found := k.GetVendorProducts(ctx, vid)

	if found {
		for i, value := range vendorProducts.Products {
			if value.Pid == pid {
				vendorProducts.Products = append(vendorProducts.Products[:i], vendorProducts.Products[i+1:]...)

				break
			}
		}

		if len(vendorProducts.Products) == 0 {
			k.RemoveVendorProducts(ctx, vid)
		} else {
			k.SetVendorProducts(ctx, vendorProducts)
		}
	}
}
