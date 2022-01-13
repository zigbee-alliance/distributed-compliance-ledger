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

// SetModelVersion set a specific modelVersion in the store from its index.
func (k Keeper) SetModelVersion(ctx sdk.Context, modelVersion types.ModelVersion) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ModelVersionKeyPrefix))
	b := k.cdc.MustMarshal(&modelVersion)
	store.Set(types.ModelVersionKey(
		modelVersion.Vid,
		modelVersion.Pid,
		modelVersion.SoftwareVersion,
	), b)
}

// GetModelVersion returns a modelVersion from its index.
func (k Keeper) GetModelVersion(
	ctx sdk.Context,
	vid int32,
	pid int32,
	softwareVersion uint32,

) (val types.ModelVersion, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ModelVersionKeyPrefix))

	b := store.Get(types.ModelVersionKey(
		vid,
		pid,
		softwareVersion,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveModelVersion removes a modelVersion from the store.
func (k Keeper) RemoveModelVersion(
	ctx sdk.Context,
	vid int32,
	pid int32,
	softwareVersion uint32,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ModelVersionKeyPrefix))
	store.Delete(types.ModelVersionKey(
		vid,
		pid,
		softwareVersion,
	))
}

// GetAllModelVersion returns all modelVersion.
func (k Keeper) GetAllModelVersion(ctx sdk.Context) (list []types.ModelVersion) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ModelVersionKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ModelVersion
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
