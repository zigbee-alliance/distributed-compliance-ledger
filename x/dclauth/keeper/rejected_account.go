package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

// SetRejectedAccount set a specific rejectedAccount in the store from its index.
func (k Keeper) SetRejectedAccount(ctx sdk.Context, rejectedAccount types.RejectedAccount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RejectedAccountKeyPrefix))
	b := k.cdc.MustMarshal(&rejectedAccount)
	store.Set(types.RejectedAccountKey(
		rejectedAccount.GetAddress(),
	), b)
}

// GetRejectedAccount returns a rejectedAccount from its index.
func (k Keeper) GetRejectedAccount(
	ctx sdk.Context,
	address sdk.AccAddress,
) (val types.RejectedAccount, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RejectedAccountKeyPrefix))

	b := store.Get(types.RejectedAccountKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// RemoveRejectedAccount removes a rejectedAccount from the store.
func (k Keeper) RemoveRejectedAccount(
	ctx sdk.Context,
	address sdk.AccAddress,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RejectedAccountKeyPrefix))
	store.Delete(types.RejectedAccountKey(
		address,
	))
}

// GetAllRejectedAccount returns all rejectedAccount.
func (k Keeper) GetAllRejectedAccount(ctx sdk.Context) (list []types.RejectedAccount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RejectedAccountKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RejectedAccount
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
