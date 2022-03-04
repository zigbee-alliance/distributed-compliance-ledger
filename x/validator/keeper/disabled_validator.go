package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

// SetDisabledValidator set a specific disabledValidator in the store from its index.
func (k Keeper) SetDisabledValidator(ctx sdk.Context, disabledValidator types.DisabledValidator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DisabledValidatorKeyPrefix))
	b := k.cdc.MustMarshal(&disabledValidator)
	store.Set(types.DisabledValidatorKey(
		disabledValidator.Address,
	), b)
}

// GetDisabledValidator returns a disabledValidator from its index.
func (k Keeper) GetDisabledValidator(
	ctx sdk.Context,
	address string,

) (val types.DisabledValidator, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DisabledValidatorKeyPrefix))

	b := store.Get(types.DisabledValidatorKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDisabledValidator removes a disabledValidator from the store.
func (k Keeper) RemoveDisabledValidator(
	ctx sdk.Context,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DisabledValidatorKeyPrefix))
	store.Delete(types.DisabledValidatorKey(
		address,
	))
}

// GetAllDisabledValidator returns all disabledValidator.
func (k Keeper) GetAllDisabledValidator(ctx sdk.Context) (list []types.DisabledValidator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DisabledValidatorKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DisabledValidator
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
