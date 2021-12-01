package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

// SetPendingAccountRevocation set a specific pendingAccountRevocation in the store from its index
func (k Keeper) SetPendingAccountRevocation(ctx sdk.Context, pendingAccountRevocation types.PendingAccountRevocation) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingAccountRevocationKeyPrefix))
	b := k.cdc.MustMarshal(&pendingAccountRevocation)
	store.Set(types.PendingAccountRevocationKey(
		pendingAccountRevocation.Address,
	), b)
}

// GetPendingAccountRevocation returns a pendingAccountRevocation from its index
func (k Keeper) GetPendingAccountRevocation(
	ctx sdk.Context,
	address string,

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

// RemovePendingAccountRevocation removes a pendingAccountRevocation from the store
func (k Keeper) RemovePendingAccountRevocation(
	ctx sdk.Context,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingAccountRevocationKeyPrefix))
	store.Delete(types.PendingAccountRevocationKey(
		address,
	))
}

// GetAllPendingAccountRevocation returns all pendingAccountRevocation
func (k Keeper) GetAllPendingAccountRevocation(ctx sdk.Context) (list []types.PendingAccountRevocation) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingAccountRevocationKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PendingAccountRevocation
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
