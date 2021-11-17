package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

// SetValidatorOwner set a specific validatorOwner in the store from its index
func (k Keeper) SetValidatorOwner(ctx sdk.Context, validatorOwner types.ValidatorOwner) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorOwnerKeyPrefix))
	b := k.cdc.MustMarshal(&validatorOwner)
	store.Set(types.ValidatorOwnerKey(
		validatorOwner.Address,
	), b)
}

// GetValidatorOwner returns a validatorOwner from its index
func (k Keeper) GetValidatorOwner(
	ctx sdk.Context,
	address string,

) (val types.ValidatorOwner, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorOwnerKeyPrefix))

	b := store.Get(types.ValidatorOwnerKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveValidatorOwner removes a validatorOwner from the store
func (k Keeper) RemoveValidatorOwner(
	ctx sdk.Context,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorOwnerKeyPrefix))
	store.Delete(types.ValidatorOwnerKey(
		address,
	))
}

// GetAllValidatorOwner returns all validatorOwner
func (k Keeper) GetAllValidatorOwner(ctx sdk.Context) (list []types.ValidatorOwner) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorOwnerKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ValidatorOwner
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
