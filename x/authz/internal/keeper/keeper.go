package keeper

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	// Unexposed key to access store from sdk.Context.
	storeKey sdk.StoreKey

	// The wire codec for binary encoding/decoding.
	cdc *codec.Codec
}

func NewKeeper(storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{storeKey: storeKey, cdc: cdc}
}

// Gets the entire AccountRoles struct.
func (k Keeper) GetAccountRoles(ctx sdk.Context, addr sdk.AccAddress) types.AccountRoles {
	if !k.IsAccountRolesPresent(ctx, addr) {
		return types.NewAccountRoles(addr, []types.AccountRole{})
	}

	store := ctx.KVStore(k.storeKey)
	bz := store.Get(addr.Bytes())

	var accountRoles types.AccountRoles

	k.cdc.MustUnmarshalBinaryBare(bz, &accountRoles)

	return accountRoles
}

// Sets the entire AccountRoles struct.
func (k Keeper) SetAccountRoles(ctx sdk.Context, accountRoles types.AccountRoles) {
	if len(accountRoles.Roles) == 0 {
		k.DeleteAccountRoles(ctx, accountRoles.Address)
		return
	}

	store := ctx.KVStore(k.storeKey)
	store.Set(accountRoles.Address.Bytes(), k.cdc.MustMarshalBinaryBare(accountRoles))
}

// Deletes the AccountRoles from the store.
func (k Keeper) DeleteAccountRoles(ctx sdk.Context, addr sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(addr.Bytes())
}

// Iterate over all AccountRoles.
func (k Keeper) IterateAccountRoles(ctx sdk.Context, process func(accountRoles types.AccountRoles) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, nil)
	defer iter.Close()

	for {
		if !iter.Valid() {
			return
		}

		val := iter.Value()

		var accountRoles types.AccountRoles

		k.cdc.MustUnmarshalBinaryBare(val, &accountRoles)

		if process(accountRoles) {
			return
		}

		iter.Next()
	}
}

func (k Keeper) HasRole(ctx sdk.Context, addr sdk.AccAddress, roleToCheck types.AccountRole) bool {
	accountRoles := k.GetAccountRoles(ctx, addr)

	for _, role := range accountRoles.Roles {
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

	accountRoles := k.GetAccountRoles(ctx, addr)
	accountRoles.Roles = append(accountRoles.Roles, roleToAdd)
	k.SetAccountRoles(ctx, accountRoles)
}

func (k Keeper) RevokeRole(ctx sdk.Context, addr sdk.AccAddress, roleToRevoke types.AccountRole) {
	accountRoles := k.GetAccountRoles(ctx, addr)

	var filteredRoles []types.AccountRole

	for _, role := range accountRoles.Roles {
		if role != roleToRevoke {
			filteredRoles = append(filteredRoles, role)
		}
	}

	accountRoles.Roles = filteredRoles
	k.SetAccountRoles(ctx, accountRoles)
}

// Check if the AccountRoles is present in the store or not.
func (k Keeper) IsAccountRolesPresent(ctx sdk.Context, addr sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(addr.Bytes())
}

func (k Keeper) CountAccounts(ctx sdk.Context, roleToCount types.AccountRole) int {
	res := 0

	k.IterateAccountRoles(ctx, func(accountRoles types.AccountRoles) (stop bool) {
		for _, role := range accountRoles.Roles {
			if role == roleToCount {
				res++
				return false
			}
		}

		return false
	})

	return res
}
