package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

// GetAccount returns a account from its index.
func (k Keeper) GetAccountO(
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

// exists to satisfy cosmos AccountKeeper.GetAccount interface
// TODO consider better way.
func (k Keeper) GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI {
	acc, found := k.GetAccountO(ctx, addr)
	if !found {
		return nil
	}

	return acc.BaseAccount
}

// GetAllAccount returns all account.
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

// just a stub to have AccountKeeper.GetModuleAddress API filled.
func (k Keeper) GetModuleAddress(moduleName string) sdk.AccAddress {
	return nil
}
