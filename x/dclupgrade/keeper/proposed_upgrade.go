package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/types"
)

// SetProposedUpgrade set a specific proposedUpgrade in the store from its index.
func (k Keeper) SetProposedUpgrade(ctx sdk.Context, proposedUpgrade types.ProposedUpgrade) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposedUpgradeKeyPrefix))
	b := k.cdc.MustMarshal(&proposedUpgrade)
	store.Set(types.ProposedUpgradeKey(
		proposedUpgrade.Plan.Name,
	), b)
}

// GetProposedUpgrade returns a proposedUpgrade from its index.
func (k Keeper) GetProposedUpgrade(
	ctx sdk.Context,
	name string,

) (val types.ProposedUpgrade, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposedUpgradeKeyPrefix))

	b := store.Get(types.ProposedUpgradeKey(
		name,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// RemoveProposedUpgrade removes a proposedUpgrade from the store.
func (k Keeper) RemoveProposedUpgrade(
	ctx sdk.Context,
	name string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposedUpgradeKeyPrefix))
	store.Delete(types.ProposedUpgradeKey(
		name,
	))
}

// GetAllProposedUpgrade returns all proposedUpgrade.
func (k Keeper) GetAllProposedUpgrade(ctx sdk.Context) (list []types.ProposedUpgrade) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposedUpgradeKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var val types.ProposedUpgrade
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
