package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

// SetValidatorMissedBlockBitArray set a specific validatorMissedBlockBitArray in the store from its index
func (k Keeper) SetValidatorMissedBlockBitArray(ctx sdk.Context, validatorMissedBlockBitArray types.ValidatorMissedBlockBitArray) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorMissedBlockBitArrayKeyPrefix))
	b := k.cdc.MustMarshal(&validatorMissedBlockBitArray)
	store.Set(types.ValidatorMissedBlockBitArrayKey(
		validatorMissedBlockBitArray.Address,
		validatorMissedBlockBitArray.Index,
	), b)
}

// GetValidatorMissedBlockBitArray returns a validatorMissedBlockBitArray from its index
func (k Keeper) GetValidatorMissedBlockBitArray(
	ctx sdk.Context,
	address string,
	index uint64,

) (val types.ValidatorMissedBlockBitArray, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorMissedBlockBitArrayKeyPrefix))

	b := store.Get(types.ValidatorMissedBlockBitArrayKey(
		address,
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveValidatorMissedBlockBitArray removes a validatorMissedBlockBitArray from the store
func (k Keeper) RemoveValidatorMissedBlockBitArray(
	ctx sdk.Context,
	address string,
	index uint64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorMissedBlockBitArrayKeyPrefix))
	store.Delete(types.ValidatorMissedBlockBitArrayKey(
		address,
		index,
	))
}

// GetAllValidatorMissedBlockBitArray returns all validatorMissedBlockBitArray
func (k Keeper) GetAllValidatorMissedBlockBitArray(ctx sdk.Context) (list []types.ValidatorMissedBlockBitArray) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorMissedBlockBitArrayKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ValidatorMissedBlockBitArray
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
