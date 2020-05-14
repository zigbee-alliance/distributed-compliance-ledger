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
	Account Index
*/
func (k Keeper) NewAccount(ctx sdk.Context, account types.Account) types.Account {
	account.AccountNumber = k.GetNextAccountNumber(ctx)
	return account
}

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

// GetAllAccounts returns all accounts in the store.
func (k Keeper) GetAllAccounts(ctx sdk.Context) (accounts []types.Account) {
	appendAccount := func(acc types.Account) (stop bool) {
		accounts = append(accounts, acc)
		return false
	}
	k.IterateAccounts(ctx, appendAccount)
	return accounts
}

func (k Keeper) SetAccount(ctx sdk.Context, acc types.Account) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryBare(acc)
	store.Set(types.GetAccountKey(acc.Address), bz)
}

func (k Keeper) IsAccountPresent(ctx sdk.Context, acc sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetAccountKey(acc))
}

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

func (k Keeper) HasRole(ctx sdk.Context, addr sdk.AccAddress, roleToCheck types.AccountRole) bool {
	account := k.GetAccount(ctx, addr)

	for _, role := range account.Roles {
		if role == roleToCheck {
			return true
		}
	}

	return false
}

func (k Keeper) AssignRole(ctx sdk.Context, addr sdk.AccAddress, roleToAdd types.AccountRole) {
	if k.HasRole(ctx, addr, roleToAdd) {
		return
	}

	account := k.GetAccount(ctx, addr)
	account.Roles = append(account.Roles, roleToAdd)
	k.SetAccount(ctx, account)
}

func (k Keeper) RevokeRole(ctx sdk.Context, addr sdk.AccAddress, roleToRevoke types.AccountRole) {
	account := k.GetAccount(ctx, addr)
	var filteredRoles []types.AccountRole

	for _, role := range account.Roles {
		if role != roleToRevoke {
			filteredRoles = append(filteredRoles, role)
		}
	}

	account.Roles = filteredRoles
	k.SetAccount(ctx, account)
}

// Check if the AccountRoles is present in the store or not
func (k Keeper) IsAccountRolesPresent(ctx sdk.Context, addr sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(addr.Bytes())
}

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
	Account Number
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