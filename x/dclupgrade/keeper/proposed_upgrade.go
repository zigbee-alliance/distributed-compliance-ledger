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

// SetProposedUpgrade set a specific proposedUpgrade in the store from its index.
func (k Keeper) SetProposedUpgrade(ctx sdk.Context, proposedUpgrade types.ProposedUpgrade) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposedUpgradeKeyPrefix))
	b := k.cdc.MustMarshal(&proposedUpgrade)
	store.Set(types.ProposedUpgradeKey(
		proposedUpgrade.Plan.Name,
	), b)
}

// GetProposedUpgrade returns a proposedUpgrade from its index.
func (k Keeper) GetProposedUpgrade(
	ctx sdk.Context,
	name string,
) (val types.ProposedUpgrade, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposedUpgradeKeyPrefix))

	b := store.Get(types.ProposedUpgradeKey(
		name,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// RemoveProposedUpgrade removes a proposedUpgrade from the store.
func (k Keeper) RemoveProposedUpgrade(
	ctx sdk.Context,
	name string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposedUpgradeKeyPrefix))
	store.Delete(types.ProposedUpgradeKey(
		name,
	))
}

// GetAllProposedUpgrade returns all proposedUpgrade.
func (k Keeper) GetAllProposedUpgrade(ctx sdk.Context) (list []types.ProposedUpgrade) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposedUpgradeKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer func() { _ = iterator.Close() }()
	for ; iterator.Valid(); iterator.Next() {
		var val types.ProposedUpgrade
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
