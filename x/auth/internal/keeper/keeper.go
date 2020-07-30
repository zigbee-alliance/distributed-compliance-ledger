package keeper

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
	Main Account Index
*/
// Assign sequence number to account object.
func (k Keeper) NewAccountWithNumber(ctx sdk.Context, account types.Account) types.Account {
	account.AccountNumber = k.GetNextAccountNumber(ctx)
	return account
}

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

// Assign account number and store it.
func (k Keeper) AssignNumberAndStoreAccount(ctx sdk.Context, account types.Account) {
	store := ctx.KVStore(k.storeKey)
	account.AccountNumber = k.GetNextAccountNumber(ctx)
	bz := k.cdc.MustMarshalBinaryBare(account)
	store.Set(types.GetAccountKey(account.Address), bz)
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

/*
	Proposed Account
*/

// Gets the Proposed Account record associated with an address.
func (k Keeper) GetProposedAccount(ctx sdk.Context, address sdk.AccAddress) types.PendingAccount {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetPendingAccountKey(address))
	if bz == nil {
		panic("Proposed Account does not exist")
	}

	var cert types.PendingAccount
	k.cdc.MustUnmarshalBinaryBare(bz, &cert)
	return cert
}

// Sets Proposed Account record for an address.
func (k Keeper) SetProposedAccount(ctx sdk.Context, account types.PendingAccount) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetPendingAccountKey(account.Address), k.cdc.MustMarshalBinaryBare(account))
}

// Check if the Proposed Account record associated with an address is present in the store or not.
func (k Keeper) IsProposedAccountPresent(ctx sdk.Context, address sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetPendingAccountKey(address))
}

// Count total Proposed Accounts.
func (k Keeper) CountTotalProposedAccounts(ctx sdk.Context) int {
	return k.countTotal(ctx, types.PendingAccountPrefix)
}

// Iterate over all Proposed Accounts.
func (k Keeper) IterateProposedAccounts(ctx sdk.Context, process func(info types.PendingAccount) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, types.PendingAccountPrefix)
	defer iter.Close()

	for {
		if !iter.Valid() {
			return
		}

		val := iter.Value()

		var account types.PendingAccount

		k.cdc.MustUnmarshalBinaryBare(val, &account)

		if process(account) {
			return
		}

		iter.Next()
	}
}

// Deletes the Proposed Account from the store.
func (k Keeper) DeleteProposedAccount(ctx sdk.Context, address sdk.AccAddress) {
	if !k.IsProposedAccountPresent(ctx, address) {
		panic("Proposed Account does not exist")
	}

	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetPendingAccountKey(address))
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

/*
	Common functions
*/
func (k Keeper) countTotal(ctx sdk.Context, prefix []byte) int {
	store := ctx.KVStore(k.storeKey)
	res := 0

	iter := sdk.KVStorePrefixIterator(store, prefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		res++
	}

	return res
}
