package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

// SetRevokedAccount set a specific revokedAccount in the store from its index.
func (k Keeper) SetRevokedAccount(ctx sdk.Context, revokedAccount types.RevokedAccount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RevokedAccountKeyPrefix))
	b := k.cdc.MustMarshal(&revokedAccount)
	store.Set(types.RevokedAccountKey(
		revokedAccount.GetAddress(),
	), b)
}

// GetRevokedAccount returns a revokedAccount from its index.
func (k Keeper) GetRevokedAccount(
	ctx sdk.Context,
	address sdk.AccAddress,

) (val types.RevokedAccount, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RevokedAccountKeyPrefix))

	b := store.Get(types.RevokedAccountKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRevokedAccount removes a revokedAccount from the store.
func (k Keeper) RemoveRevokedAccount(
	ctx sdk.Context,
	address sdk.AccAddress,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RevokedAccountKeyPrefix))
	store.Delete(types.RevokedAccountKey(
		address,
	))
}

// GetAllRevokedAccount returns all revokedAccount.
func (k Keeper) GetAllRevokedAccount(ctx sdk.Context) (list []types.RevokedAccount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RevokedAccountKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RevokedAccount
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
