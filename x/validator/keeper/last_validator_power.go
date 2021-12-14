package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkstakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

// SetLastValidatorPower set a specific lastValidatorPower in the store from its index
func (k Keeper) SetLastValidatorPower(ctx sdk.Context, lastValidatorPower types.LastValidatorPower) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LastValidatorPowerKeyPrefix)

	addr, err := sdk.ValAddressFromBech32(lastValidatorPower.Owner)
	if err != nil {
		panic(err)
	}

	b := k.cdc.MustMarshal(&lastValidatorPower)
	store.Set(types.LastValidatorPowerKey(
		addr,
	), b)
}

// Check if the validator power record associated with validator address is present in the store or not.
func (k Keeper) IsLastValidatorPowerPresent(ctx sdk.Context, owner sdk.ValAddress) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LastValidatorPowerKeyPrefix)

	return store.Has(types.LastValidatorPowerKey(owner))
}

// GetLastValidatorPower returns a lastValidatorPower from its index
func (k Keeper) GetLastValidatorPower(
	ctx sdk.Context,
	owner sdk.ValAddress,

) (val types.LastValidatorPower, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LastValidatorPowerKeyPrefix)

	b := store.Get(types.LastValidatorPowerKey(
		owner,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLastValidatorPower removes a lastValidatorPower from the store
func (k Keeper) RemoveLastValidatorPower(
	ctx sdk.Context,
	owner sdk.ValAddress,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LastValidatorPowerKeyPrefix)
	store.Delete(types.LastValidatorPowerKey(
		owner,
	))
}

// GetAllLastValidatorPower returns all lastValidatorPower
func (k Keeper) GetAllLastValidatorPower(ctx sdk.Context) (list []types.LastValidatorPower) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LastValidatorPowerKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LastValidatorPower
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// count total number of active validators.
func (k Keeper) CountLastValidators(ctx sdk.Context) (count int) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LastValidatorPowerKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		count++
	}

	return count
}

// get all active validator set.
func (k Keeper) GetAllLastValidators(ctx sdk.Context) (validators []types.Validator) {
	k.IterateLastValidators(ctx, func(validator types.Validator) (stop bool) {
		validators = append(validators, validator)

		return false
	})

	return validators
}

// iterate through the active validator set and perform the provided function.
func (k Keeper) IterateLastValidators(ctx sdk.Context, process func(validator types.Validator) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LastValidatorPowerKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var vp types.LastValidatorPower
		k.cdc.MustUnmarshal(iterator.Value(), &vp)

		addr, err := sdk.ValAddressFromBech32(vp.Owner)
		if err != nil {
			panic(err)
		}

		validator, found := k.GetValidator(ctx, addr)
		if !found {
			panic(sdkstakingtypes.ErrNoValidatorFound)
		}

		if process(validator) {
			return
		}
	}
}
