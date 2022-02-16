package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/types"
)

// SetApprovedUpgrade set a specific approvedUpgrade in the store from its index
func (k Keeper) SetApprovedUpgrade(ctx sdk.Context, approvedUpgrade types.ApprovedUpgrade) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApprovedUpgradeKeyPrefix))
	b := k.cdc.MustMarshal(&approvedUpgrade)
	store.Set(types.ApprovedUpgradeKey(
		approvedUpgrade.Name,
	), b)
}

// GetApprovedUpgrade returns a approvedUpgrade from its index
func (k Keeper) GetApprovedUpgrade(
	ctx sdk.Context,
	name string,

) (val types.ApprovedUpgrade, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApprovedUpgradeKeyPrefix))

	b := store.Get(types.ApprovedUpgradeKey(
		name,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveApprovedUpgrade removes a approvedUpgrade from the store
func (k Keeper) RemoveApprovedUpgrade(
	ctx sdk.Context,
	name string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApprovedUpgradeKeyPrefix))
	store.Delete(types.ApprovedUpgradeKey(
		name,
	))
}

// GetAllApprovedUpgrade returns all approvedUpgrade
func (k Keeper) GetAllApprovedUpgrade(ctx sdk.Context) (list []types.ApprovedUpgrade) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApprovedUpgradeKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ApprovedUpgrade
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
