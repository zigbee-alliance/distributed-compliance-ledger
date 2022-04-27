package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

// SetRejectedNode set a specific rejectedNode in the store from its index.
func (k Keeper) SetRejectedNode(ctx sdk.Context, rejectedDisableValidator types.RejectedDisableValidator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RejectedNodeKeyPrefix))
	b := k.cdc.MustMarshal(&rejectedDisableValidator)
	store.Set(types.RejectedNodeKey(
		rejectedDisableValidator.GetAddress(),
	), b)
}

// GetRejectedNode returns a rejectedNode from its index.
func (k Keeper) GetRejectedNode(
	ctx sdk.Context,
	owner string,
) (val types.RejectedDisableValidator, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RejectedNodeKeyPrefix))

	ownerValAddress, _ := sdk.ValAddressFromBech32(owner)
	b := store.Get(types.RejectedNodeKey(
		ownerValAddress,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// RemoveRejectedNode removes a rejectedNode from the store.
func (k Keeper) RemoveRejectedNode(
	ctx sdk.Context,
	owner string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RejectedNodeKeyPrefix))
	ownerValAddress, _ := sdk.ValAddressFromBech32(owner)
	store.Delete(types.RejectedNodeKey(
		ownerValAddress,
	))
}

// GetAllRejectedNode returns all rejectedNode.
func (k Keeper) GetAllRejectedDisableValidator(ctx sdk.Context) (list []types.RejectedDisableValidator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RejectedNodeKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RejectedDisableValidator
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
