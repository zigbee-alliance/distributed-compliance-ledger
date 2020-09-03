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
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/validator/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetValidatorSigningInfo(ctx sdk.Context, address sdk.ConsAddress) (info types.ValidatorSigningInfo) {
	store := ctx.KVStore(k.storeKey)
	value := store.Get(types.GetValidatorSigningInfoKey(address))

	if value == nil {
		return
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(value, &info)

	return
}

func (k Keeper) SetValidatorSigningInfo(ctx sdk.Context, info types.ValidatorSigningInfo) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(info)
	store.Set(types.GetValidatorSigningInfoKey(info.Address), bz)
}

func (k Keeper) IterateValidatorSigningInfos(ctx sdk.Context,
	handler func(info types.ValidatorSigningInfo) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, types.ValidatorSigningInfoPrefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var info types.ValidatorSigningInfo

		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &info)

		if handler(info) {
			break
		}
	}
}

func (k Keeper) GetValidatorMissedBlockBitArray(ctx sdk.Context, address sdk.ConsAddress, index int64) (missed bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetValidatorMissedBlockBitArrayKey(address, index))

	if bz == nil {
		// treat empty key as not missed
		return false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &missed)

	return
}

func (k Keeper) SetValidatorMissedBlockBitArray(ctx sdk.Context, address sdk.ConsAddress, index int64, missed bool) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(missed)
	store.Set(types.GetValidatorMissedBlockBitArrayKey(address, index), bz)
}

func (k Keeper) IterateValidatorMissedBlockBitArray(ctx sdk.Context, address sdk.ConsAddress,
	handler func(index int64, missed bool) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	index := int64(0)

	for ; index < types.SignedBlocksWindow; index++ {
		var missed bool

		bz := store.Get(types.GetValidatorMissedBlockBitArrayKey(address, index))

		if bz == nil {
			continue
		}

		k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &missed)

		if handler(index, missed) {
			break
		}
	}
}

func (k Keeper) ClearValidatorMissedBlockBitArray(ctx sdk.Context, address sdk.ConsAddress) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.GetValidatorMissedBlockBitArrayPrefixKey(address))

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		store.Delete(iter.Key())
	}
}
