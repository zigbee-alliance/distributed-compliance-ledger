package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

// SetLastValidatorPower set a specific lastValidatorPower in the store from its index
func (k Keeper) SetLastValidatorPower(ctx sdk.Context, lastValidatorPower types.LastValidatorPower) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LastValidatorPowerKeyPrefix))
	b := k.cdc.MustMarshal(&lastValidatorPower)
	store.Set(types.LastValidatorPowerKey(
		lastValidatorPower.Owner,
	), b)
}

// GetLastValidatorPower returns a lastValidatorPower from its index
func (k Keeper) GetLastValidatorPower(
	ctx sdk.Context,
	owner string,

) (val types.LastValidatorPower, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LastValidatorPowerKeyPrefix))

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
	owner string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LastValidatorPowerKeyPrefix))
	store.Delete(types.LastValidatorPowerKey(
		owner,
	))
}

// GetAllLastValidatorPower returns all lastValidatorPower
func (k Keeper) GetAllLastValidatorPower(ctx sdk.Context) (list []types.LastValidatorPower) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LastValidatorPowerKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LastValidatorPower
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
