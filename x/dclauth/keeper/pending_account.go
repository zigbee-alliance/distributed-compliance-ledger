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

// SetPendingAccount set a specific pendingAccount in the store from its index.
func (k Keeper) SetPendingAccount(ctx sdk.Context, pendingAccount types.PendingAccount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingAccountKeyPrefix))
	b := k.cdc.MustMarshal(&pendingAccount)
	store.Set(types.PendingAccountKey(
		pendingAccount.GetAddress(),
	), b)
}

// GetPendingAccount returns a pendingAccount from its index.
func (k Keeper) GetPendingAccount(
	ctx sdk.Context,
	address sdk.AccAddress,

) (val types.PendingAccount, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingAccountKeyPrefix))

	b := store.Get(types.PendingAccountKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// Check if the Account record associated with an address is present in the store or not.
func (k Keeper) IsPendingAccountPresent(ctx sdk.Context, address sdk.AccAddress) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingAccountKeyPrefix))

	return store.Has(types.PendingAccountKey(
		address,
	))
}

// RemovePendingAccount removes a pendingAccount from the store.
func (k Keeper) RemovePendingAccount(
	ctx sdk.Context,
	address sdk.AccAddress,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingAccountKeyPrefix))
	store.Delete(types.PendingAccountKey(
		address,
	))
}

// GetAllPendingAccount returns all pendingAccount.
func (k Keeper) GetAllPendingAccount(ctx sdk.Context) (list []types.PendingAccount) {
	k.IteratePendingAccounts(ctx, func(acc types.PendingAccount) (stop bool) {
		list = append(list, acc)
		return false
	})

	return
}

func (k Keeper) IteratePendingAccounts(ctx sdk.Context, cb func(account types.PendingAccount) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingAccountKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PendingAccount
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if cb(val) {
			break
		}
	}
}
