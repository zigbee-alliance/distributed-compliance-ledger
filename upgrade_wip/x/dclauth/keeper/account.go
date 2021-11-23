package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

// SetAccount set a specific account in the store from its index
func (k Keeper) SetAccount(ctx sdk.Context, account types.Account) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountKeyPrefix))
	b := k.cdc.MustMarshal(&account)
	store.Set(types.AccountKey(
		account.GetAddress(),
	), b)
}

// GetAccount returns a account from its index
func (k Keeper) GetAccount(
	ctx sdk.Context,
	address sdk.AccAddress,

) (val types.Account, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountKeyPrefix))

	b := store.Get(types.AccountKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// Check if the Account record associated with an address is present in the store or not.
func (k Keeper) IsAccountPresent(ctx sdk.Context, address sdk.AccAddress) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountKeyPrefix))

	return store.Has(types.AccountKey(
		address,
	))
}

// RemoveAccount removes a account from the store
func (k Keeper) RemoveAccount(
	ctx sdk.Context,
	address sdk.AccAddress,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountKeyPrefix))
	store.Delete(types.AccountKey(
		address,
	))
}

// GetAllAccount returns all account
func (k Keeper) GetAllAccount(ctx sdk.Context) (list []types.Account) {
	k.IterateAccounts(ctx, func(acc types.Account) (stop bool) {
		list = append(list, acc)
		return false
	})

	return
}

func (k Keeper) IterateAccounts(ctx sdk.Context, cb func(account types.Account) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Account
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if cb(val) {
			break
		}
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

// Check if account has vendorID association.
func (k Keeper) HasVendorID(ctx sdk.Context, addr sdk.AccAddress, vid uint16) bool {
	account := k.GetAccount(ctx, addr)

	if account.VendorID == vid {
		return true
	} else {
		return false
	}
}

// Count account with assigned role.
func (k Keeper) CountAccountsWithRole(ctx sdk.Context, roleToCount types.AccountRole) int {
	res := 0

	k.IterateAccounts(ctx, func(acc types.Account) (stop bool) {
		for _, role := range acc.Roles {
			if role == roleToCount {
				res++

				return false
			}
		}

		return false
	})

	return res
}
