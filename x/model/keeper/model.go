// Copyright 2022 DSR Corporation
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
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

// SetModel set a specific model in the store from its index.
func (k Keeper) SetModel(ctx sdk.Context, model types.Model) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ModelKeyPrefix))
	b := k.cdc.MustMarshal(&model)
	store.Set(types.ModelKey(
		model.Vid,
		model.Pid,
	), b)
}

// GetModel returns a model from its index.
func (k Keeper) GetModel(
	ctx sdk.Context,
	vid int32,
	pid int32,

) (val types.Model, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ModelKeyPrefix))

	b := store.Get(types.ModelKey(
		vid,
		pid,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveModel removes a model from the store.
func (k Keeper) RemoveModel(
	ctx sdk.Context,
	vid int32,
	pid int32,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ModelKeyPrefix))
	store.Delete(types.ModelKey(
		vid,
		pid,
	))
}

// GetAllModel returns all model.
func (k Keeper) GetAllModel(ctx sdk.Context) (list []types.Model) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ModelKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Model
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
