package keeper

import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/validator/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
)

// keeper of the validator store
type Keeper struct {
	// Unexposed key to access store from sdk.Context
	storeKey sdk.StoreKey

	// The wire codec for binary encoding/decoding
	cdc *codec.Codec
}

func NewKeeper(key sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		storeKey: key,
		cdc:      cdc,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

/*
	Validator by Validator Consensus Address
*/

// Gets the entire Validator record associated with a validator address
func (k Keeper) GetValidator(ctx sdk.Context, addr sdk.ConsAddress) (validator types.Validator) {
	store := ctx.KVStore(k.storeKey)
	value := store.Get(types.GetValidatorKey(addr))
	if value == nil {
		panic(fmt.Sprintf("validator record not found for address: %X\n", addr))
	}

	validator = types.MustUnmarshalBinaryBareValidator(k.cdc, value)
	return validator
}

// Sets the entire Validator record for a validator address
func (k Keeper) SetValidator(ctx sdk.Context, validator types.Validator) {
	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalValidator(k.cdc, validator)
	store.Set(types.GetValidatorKey(validator.Address), bz)
}

// Check if the Validator record associated with a validator address is present in the store or not
func (k Keeper) IsValidatorPresent(ctx sdk.Context, addr sdk.ConsAddress) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetValidatorKey(addr))
}

// get the set of all validators
func (k Keeper) GetAllValidators(ctx sdk.Context) (validators []types.Validator) {
	k.IterateValidators(ctx, func(validator types.Validator) (stop bool) {
		validators = append(validators, validator)
		return false
	})
	return validators
}

// iterate over validators and apply function
func (k Keeper) IterateValidators(ctx sdk.Context, process func(validator types.Validator) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, types.ValidatorPrefix)
	defer iter.Close()

	for {
		if !iter.Valid() {
			return
		}

		validator := types.MustUnmarshalBinaryBareValidator(k.cdc, iter.Value())

		if process(validator) {
			return
		}

		iter.Next()
	}
}

// Slash a validator for an infraction. So it will be removed from Tendermint validator set
func (k Keeper) Slash(ctx sdk.Context, consAddr sdk.ConsAddress) {
	validator := k.GetValidator(ctx, consAddr)

	// Zeroing validator's weight
	validator.Power = types.ZeroPower
	k.SetValidator(ctx, validator)
}

// jail a validator
func (k Keeper) Jail(ctx sdk.Context, consAddr sdk.ConsAddress, reason string) {
	validator := k.GetValidator(ctx, consAddr)

	if validator.Jailed {
		k.Logger(ctx).Error(fmt.Sprintf("Cannot jail already jailed validator, validator: %v\n", validator))
		return
	}

	validator.Jailed = true
	validator.JailedReason = reason
	k.SetValidator(ctx, validator)
}

// unjail a validator
func (k Keeper) Unjail(ctx sdk.Context, consAddr sdk.ConsAddress) {
	validator := k.GetValidator(ctx, consAddr)

	if !validator.Jailed {
		k.Logger(ctx).Error(fmt.Sprintf("Cannot unjail already unjailed validator, validator: %v\n", validator))
		return
	}

	validator.Jailed = false
	k.SetValidator(ctx, validator)
}

/*
	Last state Validator Index
*/
// Gets validator power in the last state by the given validator address
func (k Keeper) GetLastValidatorPower(ctx sdk.Context, address sdk.ConsAddress) (power types.LastValidatorPower) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetValidatorLastPowerKey(address))
	if bz == nil {
		return power
	}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &power)
	return power
}

// Sets validator power
func (k Keeper) SetLastValidatorPower(ctx sdk.Context, validator types.LastValidatorPower) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(validator)
	store.Set(types.GetValidatorLastPowerKey(validator.ConsensusAddress), bz)
}

// Check if the validator power record associated with validator address is present in the store or not
func (k Keeper) IsLastValidatorPowerPresent(ctx sdk.Context, address sdk.ConsAddress) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetValidatorLastPowerKey(address))
}

// Delete validator power
func (k Keeper) DeleteLastValidatorPower(ctx sdk.Context, address sdk.ConsAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetValidatorLastPowerKey(address))
}

// Get active validator set
func (k Keeper) GetLastValidatorPowers(ctx sdk.Context) []types.LastValidatorPower {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.ValidatorLastPowerPrefix)
	defer iter.Close()

	var lastValidators []types.LastValidatorPower

	for {
		if !iter.Valid() {
			break
		}

		var validator types.LastValidatorPower

		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &validator)

		lastValidators = append(lastValidators, validator)

		iter.Next()
	}

	return lastValidators
}

func (k Keeper) GetAllLastValidators(ctx sdk.Context) (validators []types.Validator) {
	k.IterateLastValidators(ctx, func(validator types.Validator) (stop bool) {
		validators = append(validators, validator)
		return false
	})
	return validators
}

// Iterate through the active validator set and perform the provided function
func (k Keeper) IterateLastValidators(ctx sdk.Context, process func(validator types.Validator) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.ValidatorLastPowerPrefix)

	defer iter.Close()
	for {
		if !iter.Valid() {
			return
		}

		addr := sdk.ConsAddress(iter.Key()[1:])
		validator := k.GetValidator(ctx, addr)

		if process(validator) {
			return
		}

		iter.Next()
	}
}
