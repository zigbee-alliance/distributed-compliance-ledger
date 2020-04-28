package keeper

import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	// Unexposed key to access store from sdk.Context
	storeKey sdk.StoreKey

	// The wire codec for binary encoding/decoding
	cdc *codec.Codec
}

const (
	modelInfoPrefix      = "1"
	vendorProductsPrefix = "2"
)

func NewKeeper(storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{storeKey: storeKey, cdc: cdc}
}

// Gets the entire ModelInfo metadata struct for a ModelInfoID
func (k Keeper) GetModelInfo(ctx sdk.Context, vid uint16, pid uint16) types.ModelInfo {
	if !k.IsModelInfoPresent(ctx, vid, pid) {
		panic("ModelInfo does not exist")
	}

	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(ModelInfoId(vid, pid)))

	var device types.ModelInfo

	k.cdc.MustUnmarshalBinaryBare(bz, &device)

	return device
}

// Sets the entire ModelInfo metadata struct for a ModelInfoID
func (k Keeper) SetModelInfo(ctx sdk.Context, model types.ModelInfo) {
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(ModelInfoId(model.VID, model.PID)), k.cdc.MustMarshalBinaryBare(model))

	// Update the list of products associated with vendor
	product := types.Product{
		PID:   model.PID,
		Name:  model.Name,
		Owner: model.Owner,
		SKU:   model.SKU,
	}
	k.AppendVendorProduct(ctx, model.VID, product)
}

// Deletes the ModelInfo from the store
func (k Keeper) DeleteModelInfo(ctx sdk.Context, vid uint16, pid uint16) {
	if !k.IsModelInfoPresent(ctx, vid, pid) {
		panic("ModelInfo does not exist")
	}

	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(ModelInfoId(vid, pid)))

	// Update the list of devices associated with vendor
	k.RemoveVendorProduct(ctx, vid, pid)
}

// Iterate over all ModelInfos
func (k Keeper) IterateModelInfos(ctx sdk.Context, process func(info types.ModelInfo) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, []byte(modelInfoPrefix))
	defer iter.Close()

	for {
		if !iter.Valid() {
			return
		}

		val := iter.Value()

		var modelInfo types.ModelInfo

		k.cdc.MustUnmarshalBinaryBare(val, &modelInfo)

		if process(modelInfo) {
			return
		}

		iter.Next()
	}
}

func (k Keeper) CountTotalModelInfos(ctx sdk.Context) int {
	return k.countTotal(ctx, modelInfoPrefix)
}

// Check if the ModelInfo is present in the store or not
func (k Keeper) IsModelInfoPresent(ctx sdk.Context, vid uint16, pid uint16) bool {
	return k.isRecordPresent(ctx, ModelInfoId(vid, pid))
}

// Id builder for ModelInfo
func ModelInfoId(vid interface{}, pid interface{}) string {
	return fmt.Sprintf("%s:%v:%v", modelInfoPrefix, vid, pid)
}

// Gets the entire VendorProducts struct for a Vendor
func (k Keeper) GetVendorProducts(ctx sdk.Context, vid uint16) types.VendorProducts {
	if !k.IsVendorProductsPresent(ctx, vid) {
		return types.NewVendorProducts(vid)
	}

	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(VendorProductsId(vid)))

	var vendorProducts types.VendorProducts

	k.cdc.MustUnmarshalBinaryBare(bz, &vendorProducts)

	return vendorProducts
}

// Add Product to Vendor
func (k Keeper) AppendVendorProduct(ctx sdk.Context, vid uint16, product types.Product) {
	store := ctx.KVStore(k.storeKey)

	vendorProducts := k.GetVendorProducts(ctx, vid)
	vendorProducts.AddVendorProduct(product)

	store.Set([]byte(VendorProductsId(vid)), k.cdc.MustMarshalBinaryBare(vendorProducts))
}

// Delete Product of Vendor
func (k Keeper) RemoveVendorProduct(ctx sdk.Context, vid uint16, pid uint16) {
	if !k.IsVendorProductsPresent(ctx, vid) {
		panic("VendorProducts does not exist")
	}

	store := ctx.KVStore(k.storeKey)

	vendorProducts := k.GetVendorProducts(ctx, vid)
	vendorProducts.RemoveVendorProduct(pid)

	if vendorProducts.IsEmpty() {
		store.Delete([]byte(VendorProductsId(vid)))
	} else {
		store.Set([]byte(VendorProductsId(vid)), k.cdc.MustMarshalBinaryBare(vendorProducts))
	}
}

// Check if the VendorProducts is present in the store or not
func (k Keeper) IsVendorProductsPresent(ctx sdk.Context, vid uint16) bool {
	return k.isRecordPresent(ctx, VendorProductsId(vid))
}

// Iterate over all VendorProducts
func (k Keeper) IterateVendorProducts(ctx sdk.Context, process func(vendorProducts types.VendorProducts) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, []byte(vendorProductsPrefix))
	defer iter.Close()

	for {
		if !iter.Valid() {
			return
		}

		val := iter.Value()

		var vendorProducts types.VendorProducts

		k.cdc.MustUnmarshalBinaryBare(val, &vendorProducts)

		if process(vendorProducts) {
			return
		}

		iter.Next()
	}
}

func (k Keeper) CountTotalVendorProducts(ctx sdk.Context) int {
	return k.countTotal(ctx, vendorProductsPrefix)
}

// Id builder for VendorProducts
func VendorProductsId(vid interface{}) string {
	return fmt.Sprintf("%s:%v", vendorProductsPrefix, vid)
}

// Check if the record is present in the store or not
func (k Keeper) isRecordPresent(ctx sdk.Context, id string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(id))
}

func (k Keeper) countTotal(ctx sdk.Context, prefix string) int {
	store := ctx.KVStore(k.storeKey)
	res := 0

	iter := sdk.KVStorePrefixIterator(store, []byte(prefix))
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		res++
	}

	return res
}
