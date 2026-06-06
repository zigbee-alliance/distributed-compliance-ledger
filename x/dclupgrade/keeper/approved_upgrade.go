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
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/types"
)

// SetApprovedUpgrade set a specific approvedUpgrade in the store from its index.
func (k Keeper) SetApprovedUpgrade(ctx sdk.Context, approvedUpgrade types.ApprovedUpgrade) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApprovedUpgradeKeyPrefix))
	b := k.cdc.MustMarshal(&approvedUpgrade)
	store.Set(types.ApprovedUpgradeKey(
		approvedUpgrade.Plan.Name,
	), b)
}

// GetApprovedUpgrade returns a approvedUpgrade from its index.
func (k Keeper) GetApprovedUpgrade(
	ctx sdk.Context,
	name string,
) (val types.ApprovedUpgrade, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApprovedUpgradeKeyPrefix))

	b := store.Get(types.ApprovedUpgradeKey(
		name,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// RemoveApprovedUpgrade removes a approvedUpgrade from the store.
func (k Keeper) RemoveApprovedUpgrade(
	ctx sdk.Context,
	name string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApprovedUpgradeKeyPrefix))
	store.Delete(types.ApprovedUpgradeKey(
		name,
	))
}

// GetAllApprovedUpgrade returns all approvedUpgrade.
func (k Keeper) GetAllApprovedUpgrade(ctx sdk.Context) (list []types.ApprovedUpgrade) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApprovedUpgradeKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer func() { _ = iterator.Close() }()
	for ; iterator.Valid(); iterator.Next() {
		var val types.ApprovedUpgrade
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
