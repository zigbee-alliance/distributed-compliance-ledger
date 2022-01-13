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

// SetModelVersions set a specific modelVersions in the store from its index.
func (k Keeper) SetModelVersions(ctx sdk.Context, modelVersions types.ModelVersions) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ModelVersionsKeyPrefix))
	b := k.cdc.MustMarshal(&modelVersions)
	store.Set(types.ModelVersionsKey(
		modelVersions.Vid,
		modelVersions.Pid,
	), b)
}

// GetModelVersions returns a modelVersions from its index.
func (k Keeper) GetModelVersions(
	ctx sdk.Context,
	vid int32,
	pid int32,

) (val types.ModelVersions, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ModelVersionsKeyPrefix))

	b := store.Get(types.ModelVersionsKey(
		vid,
		pid,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveModelVersions removes a modelVersions from the store.
func (k Keeper) RemoveModelVersions(
	ctx sdk.Context,
	vid int32,
	pid int32,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ModelVersionsKeyPrefix))
	store.Delete(types.ModelVersionsKey(
		vid,
		pid,
	))
}

// GetAllModelVersions returns all modelVersions.
func (k Keeper) GetAllModelVersions(ctx sdk.Context) (list []types.ModelVersions) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ModelVersionsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ModelVersions
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// AddModelVersion adds a softwareVersion to existing or new ModelVersions.
func (k Keeper) AddModelVersion(ctx sdk.Context, vid int32, pid int32, softwareVersion uint32) {
	modelVersions, found := k.GetModelVersions(ctx, vid, pid)

	if found {
		for _, value := range modelVersions.SoftwareVersions {
			if value == softwareVersion {
				return
			}
		}

		modelVersions.SoftwareVersions = append(modelVersions.SoftwareVersions, softwareVersion)
	} else {
		modelVersions.Vid = vid
		modelVersions.Pid = pid
		modelVersions.SoftwareVersions = []uint32{softwareVersion}
	}

	k.SetModelVersions(ctx, modelVersions)
}
