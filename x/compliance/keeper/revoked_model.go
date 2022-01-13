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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
)

// SetRevokedModel set a specific revokedModel in the store from its index
func (k Keeper) SetRevokedModel(ctx sdk.Context, revokedModel types.RevokedModel) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RevokedModelKeyPrefix))
	b := k.cdc.MustMarshal(&revokedModel)
	store.Set(types.RevokedModelKey(
		revokedModel.Vid,
		revokedModel.Pid,
		revokedModel.SoftwareVersion,
		revokedModel.CertificationType,
	), b)
}

// GetRevokedModel returns a revokedModel from its index
func (k Keeper) GetRevokedModel(
	ctx sdk.Context,
	vid int32,
	pid int32,
	softwareVersion uint32,
	certificationType string,

) (val types.RevokedModel, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RevokedModelKeyPrefix))

	b := store.Get(types.RevokedModelKey(
		vid,
		pid,
		softwareVersion,
		certificationType,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRevokedModel removes a revokedModel from the store
func (k Keeper) RemoveRevokedModel(
	ctx sdk.Context,
	vid int32,
	pid int32,
	softwareVersion uint32,
	certificationType string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RevokedModelKeyPrefix))
	store.Delete(types.RevokedModelKey(
		vid,
		pid,
		softwareVersion,
		certificationType,
	))
}

// GetAllRevokedModel returns all revokedModel
func (k Keeper) GetAllRevokedModel(ctx sdk.Context) (list []types.RevokedModel) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RevokedModelKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RevokedModel
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
