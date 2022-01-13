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

// SetProvisionalModel set a specific provisionalModel in the store from its index.
func (k Keeper) SetProvisionalModel(ctx sdk.Context, provisionalModel types.ProvisionalModel) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProvisionalModelKeyPrefix))
	b := k.cdc.MustMarshal(&provisionalModel)
	store.Set(types.ProvisionalModelKey(
		provisionalModel.Vid,
		provisionalModel.Pid,
		provisionalModel.SoftwareVersion,
		provisionalModel.CertificationType,
	), b)
}

// GetProvisionalModel returns a provisionalModel from its index.
func (k Keeper) GetProvisionalModel(
	ctx sdk.Context,
	vid int32,
	pid int32,
	softwareVersion uint32,
	certificationType string,

) (val types.ProvisionalModel, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProvisionalModelKeyPrefix))

	b := store.Get(types.ProvisionalModelKey(
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

// RemoveProvisionalModel removes a provisionalModel from the store.
func (k Keeper) RemoveProvisionalModel(
	ctx sdk.Context,
	vid int32,
	pid int32,
	softwareVersion uint32,
	certificationType string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProvisionalModelKeyPrefix))
	store.Delete(types.ProvisionalModelKey(
		vid,
		pid,
		softwareVersion,
		certificationType,
	))
}

// GetAllProvisionalModel returns all provisionalModel.
func (k Keeper) GetAllProvisionalModel(ctx sdk.Context) (list []types.ProvisionalModel) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProvisionalModelKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ProvisionalModel
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
