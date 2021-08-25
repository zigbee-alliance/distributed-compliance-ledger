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
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth/internal/types"
)

type Keeper struct {
	// Unexposed key to access store from sdk.Context
	storeKey sdk.StoreKey

	// The wire codec for binary encoding/decoding
	cdc *codec.Codec
}

func NewKeeper(storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{storeKey: storeKey, cdc: cdc}
}

/*
	Account
*/
// Get the Account record associated with an address.
func (k Keeper) GetAccount(ctx sdk.Context, addr sdk.AccAddress) (acc types.Account) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetAccountKey(addr))

	if bz == nil {
		panic("Account does not exist")
	}

	err := k.cdc.UnmarshalBinaryBare(bz, &acc)
	if err != nil {
		panic(err)
	}

	return
}

// Get all Account records from the store.
func (k Keeper) GetAllAccounts(ctx sdk.Context) (accounts []types.Account) {
	appendAccount := func(acc types.Account) (stop bool) {
		accounts = append(accounts, acc)

		return false
	}
	k.IterateAccounts(ctx, appendAccount)

	return accounts
}

// Set Account record for an address.
func (k Keeper) SetAccount(ctx sdk.Context, acc types.Account) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryBare(acc)
	store.Set(types.GetAccountKey(acc.Address), bz)
}

// Check if the Account record associated with an address is present in the store or not.
func (k Keeper) IsAccountPresent(ctx sdk.Context, acc sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)

	return store.Has(types.GetAccountKey(acc))
}

// Iterate over all stored accounts.
func (k Keeper) IterateAccounts(ctx sdk.Context, process func(types.Account) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.AccountPrefix)

	defer iter.Close()

	for {
		if !iter.Valid() {
			return
		}

		val := iter.Value()

		var account types.Account

		k.cdc.MustUnmarshalBinaryBare(val, &account)

		if process(account) {
			return
		}

		iter.Next()
	}
}

// Check if account has assigned role.
func (k Keeper) HasRole(ctx sdk.Context, addr sdk.AccAddress, roleToCheck types.AccountRole) bool {
	account := k.GetAccount(ctx, addr)

	for _, role := range account.Roles {
		if role == roleToCheck {
			return true
		}
	}

	return false
}

// Check if account has vendorId association.
func (k Keeper) HasVendorId(ctx sdk.Context, addr sdk.AccAddress, vid uint16) bool {
	account := k.GetAccount(ctx, addr)

	if account.VendorId == vid {
		return true
	} else {
		return false
	}
}

// Count account with assigned role.
func (k Keeper) CountAccountsWithRole(ctx sdk.Context, roleToCount types.AccountRole) int {
	res := 0

	k.IterateAccounts(ctx, func(account types.Account) (stop bool) {
		for _, role := range account.Roles {
			if role == roleToCount {
				res++

				return false
			}
		}

		return false
	})

	return res
}

// Deletes the Account from the store.
func (k Keeper) DeleteAccount(ctx sdk.Context, address sdk.AccAddress) {
	if !k.IsAccountPresent(ctx, address) {
		panic("Account does not exist")
	}

	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetAccountKey(address))
}

/*
	Pending Account
*/
// Gets the Pending Account record associated with an address.
func (k Keeper) GetPendingAccount(ctx sdk.Context, address sdk.AccAddress) types.PendingAccount {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetPendingAccountKey(address))

	if bz == nil {
		panic("Pending Account does not exist")
	}

	var pendAcc types.PendingAccount

	k.cdc.MustUnmarshalBinaryBare(bz, &pendAcc)

	return pendAcc
}

// Sets Pending Account record for an address.
func (k Keeper) SetPendingAccount(ctx sdk.Context, pendAcc types.PendingAccount) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetPendingAccountKey(pendAcc.Address), k.cdc.MustMarshalBinaryBare(pendAcc))
}

// Check if the Pending Account record associated with an address is present in the store or not.
func (k Keeper) IsPendingAccountPresent(ctx sdk.Context, address sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)

	return store.Has(types.GetPendingAccountKey(address))
}

// Iterate over all Pending Accounts.
func (k Keeper) IteratePendingAccounts(ctx sdk.Context, process func(info types.PendingAccount) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, types.PendingAccountPrefix)
	defer iter.Close()

	for {
		if !iter.Valid() {
			return
		}

		val := iter.Value()

		var pendAcc types.PendingAccount

		k.cdc.MustUnmarshalBinaryBare(val, &pendAcc)

		if process(pendAcc) {
			return
		}

		iter.Next()
	}
}

// Deletes the Pending Account from the store.
func (k Keeper) DeletePendingAccount(ctx sdk.Context, address sdk.AccAddress) {
	if !k.IsPendingAccountPresent(ctx, address) {
		panic("Pending Account does not exist")
	}

	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetPendingAccountKey(address))
}

/*
	Pending Account Revocation
*/
// Gets the Pending Account Revocation record associated with an address.
func (k Keeper) GetPendingAccountRevocation(ctx sdk.Context, address sdk.AccAddress) types.PendingAccountRevocation {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetPendingAccountRevocationKey(address))

	if bz == nil {
		panic("Pending Account Revocation does not exist")
	}

	var revoc types.PendingAccountRevocation

	k.cdc.MustUnmarshalBinaryBare(bz, &revoc)

	return revoc
}

// Sets Pending Account Revocation record for an address.
func (k Keeper) SetPendingAccountRevocation(ctx sdk.Context, revoc types.PendingAccountRevocation) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetPendingAccountRevocationKey(revoc.Address), k.cdc.MustMarshalBinaryBare(revoc))
}

// Check if the Pending Account Revocation record associated with an address is present in the store or not.
func (k Keeper) IsPendingAccountRevocationPresent(ctx sdk.Context, address sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)

	return store.Has(types.GetPendingAccountRevocationKey(address))
}

// Iterate over all Pending Account Revocations.
func (k Keeper) IteratePendingAccountRevocations(ctx sdk.Context,
	process func(info types.PendingAccountRevocation) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, types.PendingAccountRevocationPrefix)
	defer iter.Close()

	for {
		if !iter.Valid() {
			return
		}

		val := iter.Value()

		var revoc types.PendingAccountRevocation

		k.cdc.MustUnmarshalBinaryBare(val, &revoc)

		if process(revoc) {
			return
		}

		iter.Next()
	}
}

// Deletes the Pending Account Revocation from the store.
func (k Keeper) DeletePendingAccountRevocation(ctx sdk.Context, address sdk.AccAddress) {
	if !k.IsPendingAccountRevocationPresent(ctx, address) {
		panic("Pending Account Revocation does not exist")
	}

	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetPendingAccountRevocationKey(address))
}

/*
	Account Number Counter
*/
func (k Keeper) GetNextAccountNumber(ctx sdk.Context) (accNumber uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.AccountNumberCounterKey)

	if bz == nil {
		accNumber = 0
	} else {
		k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &accNumber)
	}

	bz = k.cdc.MustMarshalBinaryLengthPrefixed(accNumber + 1)
	store.Set(types.AccountNumberCounterKey, bz)

	return
}
