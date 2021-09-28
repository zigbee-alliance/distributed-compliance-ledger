// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/internal/types"
)

type Keeper struct {
	// Unexposed key to access store from sdk.Context.
	storeKey sdk.StoreKey

	// The wire codec for binary encoding/decoding.
	cdc *codec.Codec
}

func NewKeeper(storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{storeKey: storeKey, cdc: cdc}
}

// Gets the entire Model struct for a ModelID.
func (k Keeper) GetModel(ctx sdk.Context, vid uint16, pid uint16) types.Model {
	if !k.IsModelPresent(ctx, vid, pid) {
		panic("Model does not exist")
	}

	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetModelKey(vid, pid))

	var device types.Model

	k.cdc.MustUnmarshalBinaryBare(bz, &device)

	return device
}

// Sets the entire Model metadata struct for a ModelID.
func (k Keeper) SetModel(ctx sdk.Context, model types.Model) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetModelKey(model.VID, model.PID), k.cdc.MustMarshalBinaryBare(model))

	// Update the list of products associated with vendor.
	product := types.Product{
		PID:        model.PID,
		Name:       model.ProductName,
		PartNumber: model.PartNumber,
	}
	k.AppendVendorProduct(ctx, model.VID, product)
}

// Deletes the Model from the store.
func (k Keeper) DeleteModel(ctx sdk.Context, vid uint16, pid uint16) {
	if !k.IsModelPresent(ctx, vid, pid) {
		panic("Model does not exist")
	}

	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetModelKey(vid, pid))

	// Update the list of devices associated with vendor.
	k.RemoveVendorProduct(ctx, vid, pid)
}

// Iterate over all Models.
func (k Keeper) IterateModels(ctx sdk.Context, process func(info types.Model) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, types.ModelPrefix)
	defer iter.Close()

	for {
		if !iter.Valid() {
			return
		}

		val := iter.Value()

		var model types.Model

		k.cdc.MustUnmarshalBinaryBare(val, &model)

		if process(model) {
			return
		}

		iter.Next()
	}
}

func (k Keeper) CountTotalModels(ctx sdk.Context) int {
	return k.countTotal(ctx, types.ModelPrefix)
}

// Check if the Model is present in the store or not.
func (k Keeper) IsModelPresent(ctx sdk.Context, vid uint16, pid uint16) bool {
	return k.isRecordPresent(ctx, types.GetModelKey(vid, pid))
}

// Gets the entire VendorProducts struct for a Vendor.
func (k Keeper) GetVendorProducts(ctx sdk.Context, vid uint16) types.VendorProducts {
	if !k.IsVendorProductsPresent(ctx, vid) {
		return types.NewVendorProducts(vid)
	}

	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetVendorProductsKey(vid))

	var vendorProducts types.VendorProducts

	k.cdc.MustUnmarshalBinaryBare(bz, &vendorProducts)

	return vendorProducts
}

// Add Product to Vendor.
func (k Keeper) AppendVendorProduct(ctx sdk.Context, vid uint16, product types.Product) {
	store := ctx.KVStore(k.storeKey)

	vendorProducts := k.GetVendorProducts(ctx, vid)
	vendorProducts.AddVendorProduct(product)

	store.Set(types.GetVendorProductsKey(vid), k.cdc.MustMarshalBinaryBare(vendorProducts))
}

// Delete Product of Vendor.
func (k Keeper) RemoveVendorProduct(ctx sdk.Context, vid uint16, pid uint16) {
	if !k.IsVendorProductsPresent(ctx, vid) {
		panic("VendorProducts does not exist")
	}

	store := ctx.KVStore(k.storeKey)

	vendorProducts := k.GetVendorProducts(ctx, vid)
	vendorProducts.RemoveVendorProduct(pid)

	if vendorProducts.IsEmpty() {
		store.Delete(types.GetVendorProductsKey(vid))
	} else {
		store.Set(types.GetVendorProductsKey(vid), k.cdc.MustMarshalBinaryBare(vendorProducts))
	}
}

// Check if the VendorProducts is present in the store or not.
func (k Keeper) IsVendorProductsPresent(ctx sdk.Context, vid uint16) bool {
	return k.isRecordPresent(ctx, types.GetVendorProductsKey(vid))
}

// Iterate over all VendorProducts.
func (k Keeper) IterateVendorProducts(ctx sdk.Context, process func(vendorProducts types.VendorProducts) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, types.VendorProductsPrefix)
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
	return k.countTotal(ctx, types.VendorProductsPrefix)
}

// Check if the record is present in the store or not.
func (k Keeper) isRecordPresent(ctx sdk.Context, id []byte) bool {
	store := ctx.KVStore(k.storeKey)

	return store.Has(id)
}

func (k Keeper) countTotal(ctx sdk.Context, prefix []byte) int {
	store := ctx.KVStore(k.storeKey)
	res := 0

	iter := sdk.KVStorePrefixIterator(store, prefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		res++
	}

	return res
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
