package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

// SetPendingAccount set a specific pendingAccount in the store from its index
func (k Keeper) SetPendingAccount(ctx sdk.Context, pendingAccount types.PendingAccount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingAccountKeyPrefix))
	b := k.cdc.MustMarshal(&pendingAccount)
	store.Set(types.PendingAccountKey(
		pendingAccount.Address,
	), b)
}

// GetPendingAccount returns a pendingAccount from its index
func (k Keeper) GetPendingAccount(
	ctx sdk.Context,
	address string,

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

// RemovePendingAccount removes a pendingAccount from the store
func (k Keeper) RemovePendingAccount(
	ctx sdk.Context,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingAccountKeyPrefix))
	store.Delete(types.PendingAccountKey(
		address,
	))
}

// GetAllPendingAccount returns all pendingAccount
func (k Keeper) GetAllPendingAccount(ctx sdk.Context) (list []types.PendingAccount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingAccountKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PendingAccount
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
