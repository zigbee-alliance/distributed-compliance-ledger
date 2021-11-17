package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

// SetValidatorSigningInfo set a specific validatorSigningInfo in the store from its index
func (k Keeper) SetValidatorSigningInfo(ctx sdk.Context, validatorSigningInfo types.ValidatorSigningInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorSigningInfoKeyPrefix))
	b := k.cdc.MustMarshal(&validatorSigningInfo)
	store.Set(types.ValidatorSigningInfoKey(
		validatorSigningInfo.Address,
	), b)
}

// GetValidatorSigningInfo returns a validatorSigningInfo from its index
func (k Keeper) GetValidatorSigningInfo(
	ctx sdk.Context,
	address string,

) (val types.ValidatorSigningInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorSigningInfoKeyPrefix))

	b := store.Get(types.ValidatorSigningInfoKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveValidatorSigningInfo removes a validatorSigningInfo from the store
func (k Keeper) RemoveValidatorSigningInfo(
	ctx sdk.Context,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorSigningInfoKeyPrefix))
	store.Delete(types.ValidatorSigningInfoKey(
		address,
	))
}

// GetAllValidatorSigningInfo returns all validatorSigningInfo
func (k Keeper) GetAllValidatorSigningInfo(ctx sdk.Context) (list []types.ValidatorSigningInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorSigningInfoKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ValidatorSigningInfo
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
