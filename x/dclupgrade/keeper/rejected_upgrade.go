package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/types"
)

// SetRejectedUpgrade set a specific rejectedUpgrade in the store from its index.
func (k Keeper) SetRejectedUpgrade(ctx sdk.Context, rejectedUpgrade types.RejectedUpgrade) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RejectedUpgradeKeyPrefix))
	b := k.cdc.MustMarshal(&rejectedUpgrade)
	store.Set(types.RejectedUpgradeKey(
		rejectedUpgrade.Plan.Name,
	), b)
}

// GetRejectedUpgrade returns a rejectedUpgrade from its index
func (k Keeper) GetRejectedUpgrade(
	ctx sdk.Context,
	name string,
) (val types.RejectedUpgrade, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RejectedUpgradeKeyPrefix))

	b := store.Get(types.RejectedUpgradeKey(
		name,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRejectedUpgrade removes a rejectedUpgrade from the store
func (k Keeper) RemoveRejectedUpgrade(
	ctx sdk.Context,
	name string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RejectedUpgradeKeyPrefix))
	store.Delete(types.RejectedUpgradeKey(
		name,
	))
}

// GetAllRejectedUpgrade returns all rejectedUpgrade
func (k Keeper) GetAllRejectedUpgrade(ctx sdk.Context) (list []types.RejectedUpgrade) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RejectedUpgradeKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RejectedUpgrade
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
