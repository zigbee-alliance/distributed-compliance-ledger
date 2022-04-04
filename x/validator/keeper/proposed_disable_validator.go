package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

// SetProposedDisableValidator set a specific proposedDisableValidator in the store from its index.
func (k Keeper) SetProposedDisableValidator(ctx sdk.Context, proposedDisableValidator types.ProposedDisableValidator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposedDisableValidatorKeyPrefix))
	b := k.cdc.MustMarshal(&proposedDisableValidator)
	store.Set(types.ProposedDisableValidatorKey(
		proposedDisableValidator.Address,
	), b)
}

// GetProposedDisableValidator returns a proposedDisableValidator from its index.
func (k Keeper) GetProposedDisableValidator(
	ctx sdk.Context,
	address string,

) (val types.ProposedDisableValidator, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposedDisableValidatorKeyPrefix))

	b := store.Get(types.ProposedDisableValidatorKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveProposedDisableValidator removes a proposedDisableValidator from the store.
func (k Keeper) RemoveProposedDisableValidator(
	ctx sdk.Context,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposedDisableValidatorKeyPrefix))
	store.Delete(types.ProposedDisableValidatorKey(
		address,
	))
}

// GetAllProposedDisableValidator returns all proposedDisableValidator.
func (k Keeper) GetAllProposedDisableValidator(ctx sdk.Context) (list []types.ProposedDisableValidator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposedDisableValidatorKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ProposedDisableValidator
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
