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
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/internal/types"
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
func (k Keeper) GetVendorInfo(ctx sdk.Context, vid uint16) types.VendorInfo {
	if !k.IsVendorInfoPresent(ctx, vid) {
		panic("VendorInfo does not exist")
	}

	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetVendorInfoKey(vid))

	var device types.VendorInfo

	k.cdc.MustUnmarshalBinaryBare(bz, &device)

	return device
}

// Sets the entire VendorInfo metadata struct for a VendorID.
func (k Keeper) SetVendorInfo(ctx sdk.Context, vendorInfo types.VendorInfo) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetVendorInfoKey(vendorInfo.VendorId), k.cdc.MustMarshalBinaryBare(vendorInfo))
}

// Iterate over all VendorInfos.
func (k Keeper) IterateVendorInfos(ctx sdk.Context, process func(info types.VendorInfo) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, types.VendorPrefix)
	defer iter.Close()

	for {
		if !iter.Valid() {
			return
		}

		val := iter.Value()

		var vendorInfo types.VendorInfo

		k.cdc.MustUnmarshalBinaryBare(val, &vendorInfo)

		if process(vendorInfo) {
			return
		}

		iter.Next()
	}
}

func (k Keeper) CountTotalVendorInfos(ctx sdk.Context) int {
	return k.countTotal(ctx, types.VendorPrefix)
}

// Check if the VendorInfo is present in the store or not.
func (k Keeper) IsVendorInfoPresent(ctx sdk.Context, vid uint16) bool {
	return k.isRecordPresent(ctx, types.GetVendorInfoKey(vid))
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
