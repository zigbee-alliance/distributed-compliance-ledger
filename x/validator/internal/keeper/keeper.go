package keeper

import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/validator/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// keeper of the validator store
type Keeper struct {
	// Unexposed key to access store from sdk.Context
	storeKey sdk.StoreKey

	// The wire codec for binary encoding/decoding
	cdc *codec.Codec
}

const (
	ValidatorPrefix                   = "1"
	validatorByConsensusAddressPrefix = "2"
	lastValidatorPowerKeyPrefix       = "3"
)

func NewKeeper(key sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		storeKey: key,
		cdc:      cdc,
	}
}

/*
	Validator by Validator address
*/

// get a single validator
func (k Keeper) GetValidator(ctx sdk.Context, addr sdk.ValAddress) (validator types.Validator) {
	store := ctx.KVStore(k.storeKey)
	value := store.Get(ValidatorId(addr))
	if value == nil {
		panic(fmt.Sprintf("validator record not found for address: %X\n", addr))
	}

	validator = types.MustUnmarshalValidator(k.cdc, value)
	return validator
}

func (k Keeper) SetValidator(ctx sdk.Context, validator types.Validator) {
	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalValidator(k.cdc, validator)
	store.Set(ValidatorId(validator.OperatorAddress), bz)
}

func (k Keeper) IsValidatorPresent(ctx sdk.Context, addr sdk.ValAddress) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(ValidatorId(addr))
}

// get the set of all validators with no limits, used during genesis dump
func (k Keeper) GetAllValidators(ctx sdk.Context) (validators []types.Validator) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(ValidatorPrefix))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		validator := types.MustUnmarshalValidator(k.cdc, iterator.Value())
		validators = append(validators, validator)
	}
	return validators
}

func (k Keeper) IterateValidators(ctx sdk.Context, process func(validator types.Validator) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, []byte(ValidatorPrefix))
	defer iter.Close()

	for {
		if !iter.Valid() {
			return
		}

		val := iter.Value()

		var validator types.Validator

		k.cdc.MustUnmarshalBinaryBare(val, &validator)

		if process(validator) {
			return
		}

		iter.Next()
	}
}

func ValidatorId(operatorAddr sdk.ValAddress) []byte {
	return []byte(fmt.Sprintf("%s:%v", ValidatorPrefix, operatorAddr))
}

/*
	Validator by Consensus address
*/
func (k Keeper) GetValidatorByConsAddr(ctx sdk.Context, consAddr sdk.ConsAddress) (validator types.Validator) {
	store := ctx.KVStore(k.storeKey)
	opAddr := store.Get(ValidatorByConsAddrId(consAddr))
	if opAddr == nil {
		panic(fmt.Errorf("validator with consensus-OperatorAddress %s not found", consAddr))
	}
	return k.GetValidator(ctx, opAddr)
}

func (k Keeper) SetValidatorByConsAddr(ctx sdk.Context, validator types.Validator) {
	store := ctx.KVStore(k.storeKey)
	consAddr := sdk.ConsAddress(validator.GetConsPubKey().Address())
	store.Set(ValidatorByConsAddrId(consAddr), validator.OperatorAddress)
}

func (k Keeper) IsValidatorByConsAddrPresent(ctx sdk.Context, consAddr sdk.ConsAddress) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(ValidatorByConsAddrId(consAddr))
}

func ValidatorByConsAddrId(addr sdk.ConsAddress) []byte {
	return []byte(fmt.Sprintf("%s:%v", validatorByConsensusAddressPrefix, addr))
}

/*
	Last state Validator Index
*/
func (k Keeper) GetLastValidatorPower(ctx sdk.Context, operator sdk.ValAddress) (power int64) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(LastValidatorPowerId(operator))
	if bz == nil {
		return 0
	}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &power)
	return power
}

func (k Keeper) SetLastValidatorPower(ctx sdk.Context, operator sdk.ValAddress, power int64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(power)
	store.Set(LastValidatorPowerId(operator), bz)
}

// returns an iterator for the consensus validators in the last block
func (k Keeper) LastValidatorPowersIterator(ctx sdk.Context) (iterator sdk.Iterator) {
	store := ctx.KVStore(k.storeKey)
	iterator = sdk.KVStorePrefixIterator(store, []byte(lastValidatorPowerKeyPrefix))
	return iterator
}

// Iterate over last validator powers.
func (k Keeper) IterateLastValidatorPowers(ctx sdk.Context, handler func(operator sdk.ValAddress, power int64) (stop bool)) {
	iter := k.LastValidatorPowersIterator(ctx)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		addr := sdk.ValAddress(iter.Key()[1:])
		var power int64
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &power)
		if handler(addr, power) {
			break
		}
	}
}

func (k Keeper) GetLastValidatorPowers(ctx sdk.Context) []types.LastValidatorPower {
	var lastValidatorPowers []types.LastValidatorPower
	k.IterateLastValidatorPowers(ctx, func(addr sdk.ValAddress, power int64) (stop bool) {
		lastValidatorPowers = append(lastValidatorPowers, types.LastValidatorPower{Address: addr, Power: power})
		return false
	})
	return lastValidatorPowers
}

// iterate through the active validator set and perform the provided function
func (k Keeper) IterateLastValidators(ctx sdk.Context, fn func(index int64, validator types.Validator) (stop bool)) {
	iter := k.LastValidatorPowersIterator(ctx)
	defer iter.Close()
	i := int64(0)
	for ; iter.Valid(); iter.Next() {
		addr := sdk.ValAddress(iter.Key()[1:])
		validator := k.GetValidator(ctx, addr)
		stop := fn(i, validator)
		if stop {
			break
		}
		i++
	}
}

func LastValidatorPowerId(addr sdk.ValAddress) []byte {
	return []byte(fmt.Sprintf("%s:%v", lastValidatorPowerKeyPrefix, addr))
}
