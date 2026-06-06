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

// SetRejectedUpgrade set a specific rejectedUpgrade in the store from its index.
func (k Keeper) SetRejectedUpgrade(ctx sdk.Context, rejectedUpgrade types.RejectedUpgrade) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RejectedUpgradeKeyPrefix))
	b := k.cdc.MustMarshal(&rejectedUpgrade)
	store.Set(types.RejectedUpgradeKey(
		rejectedUpgrade.Plan.Name,
	), b)
}

// GetRejectedUpgrade returns a rejectedUpgrade from its index.
func (k Keeper) GetRejectedUpgrade(
	ctx sdk.Context,
	name string,
) (val types.RejectedUpgrade, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RejectedUpgradeKeyPrefix))

	b := store.Get(types.RejectedUpgradeKey(
		name,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// RemoveRejectedUpgrade removes a rejectedUpgrade from the store.
func (k Keeper) RemoveRejectedUpgrade(
	ctx sdk.Context,
	name string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RejectedUpgradeKeyPrefix))
	store.Delete(types.RejectedUpgradeKey(
		name,
	))
}

// GetAllRejectedUpgrade returns all rejectedUpgrade.
func (k Keeper) GetAllRejectedUpgrade(ctx sdk.Context) (list []types.RejectedUpgrade) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RejectedUpgradeKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer func() { _ = iterator.Close() }()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RejectedUpgrade
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
