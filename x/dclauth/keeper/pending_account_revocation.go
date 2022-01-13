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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

// SetPendingAccountRevocation set a specific pendingAccountRevocation in the store from its index.
func (k Keeper) SetPendingAccountRevocation(ctx sdk.Context, pendingAccountRevocation types.PendingAccountRevocation) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingAccountRevocationKeyPrefix))
	b := k.cdc.MustMarshal(&pendingAccountRevocation)
	addr, _ := sdk.AccAddressFromBech32(pendingAccountRevocation.Address)
	store.Set(types.PendingAccountRevocationKey(
		addr,
	), b)
}

// GetPendingAccountRevocation returns a pendingAccountRevocation from its index.
func (k Keeper) GetPendingAccountRevocation(
	ctx sdk.Context,
	address sdk.AccAddress,

) (val types.PendingAccountRevocation, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingAccountRevocationKeyPrefix))

	b := store.Get(types.PendingAccountRevocationKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// Check if the Pending Account Revocation record associated with an address is present in the store or not.
func (k Keeper) IsPendingAccountRevocationPresent(ctx sdk.Context, address sdk.AccAddress) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingAccountRevocationKeyPrefix))

	return store.Has(types.PendingAccountRevocationKey(
		address,
	))
}

// RemovePendingAccountRevocation removes a pendingAccountRevocation from the store.
func (k Keeper) RemovePendingAccountRevocation(
	ctx sdk.Context,
	address sdk.AccAddress,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingAccountRevocationKeyPrefix))
	store.Delete(types.PendingAccountRevocationKey(
		address,
	))
}

// GetAllPendingAccountRevocation returns all pendingAccountRevocation.
func (k Keeper) GetAllPendingAccountRevocation(ctx sdk.Context) (list []types.PendingAccountRevocation) {
	k.IteratePendingAccountRevocations(ctx, func(acc types.PendingAccountRevocation) (stop bool) {
		list = append(list, acc)
		return false
	})

	return
}

func (k Keeper) IteratePendingAccountRevocations(ctx sdk.Context, cb func(account types.PendingAccountRevocation) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingAccountRevocationKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PendingAccountRevocation
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if cb(val) {
			break
		}
	}
}
