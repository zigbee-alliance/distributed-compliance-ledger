// Copyright 2022 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

// SetValidator set a specific validator in the store from its index.
func (k Keeper) SetValidator(ctx sdk.Context, validator types.Validator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorKeyPrefix))
	b := k.cdc.MustMarshal(&validator)
	store.Set(types.ValidatorKey(
		validator.GetOwner(),
	), b)
}

// Check if the Validator record associated with a validator address is present in the store or not.
func (k Keeper) IsValidatorPresent(ctx sdk.Context, owner sdk.ValAddress) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorKeyPrefix))

	return store.Has(types.ValidatorKey(owner))
}

// GetValidator returns a validator from its index.
func (k Keeper) GetValidator(
	ctx sdk.Context,
	owner sdk.ValAddress,

) (val types.Validator, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorKeyPrefix))

	b := store.Get(types.ValidatorKey(
		owner,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) mustGetValidator(ctx sdk.Context, owner sdk.ValAddress) types.Validator {
	validator, found := k.GetValidator(ctx, owner)
	if !found {
		panic(fmt.Sprintf("validator record not found for address: %X\n", owner))
	}

	return validator
}

// RemoveValidator removes a validator from the store.
func (k Keeper) RemoveValidator(
	ctx sdk.Context,
	owner sdk.ValAddress,

) {
	validator, found := k.GetValidator(ctx, owner)
	if !found {
		return
	}

	valConsAddr, err := validator.GetConsAddr()
	if err != nil {
		// TODO ??? issue 99: the best way to deal with that
		panic(err)
	} else {
		store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorByConsAddrKeyPrefix))
		store.Delete(types.ValidatorByConsAddrKey(
			valConsAddr,
		))
	}

	// FIXME issue 99: owner should be a key here
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LastValidatorPowerKeyPrefix))
	store.Delete(types.LastValidatorPowerKey(owner))

	store = prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorKeyPrefix))
	store.Delete(types.ValidatorKey(owner))

	// TODO call hooks ???
}

// validator index.
func (k Keeper) SetValidatorByConsAddr(ctx sdk.Context, validator types.Validator) error {
	consAddr, err := validator.GetConsAddr()
	if err != nil {
		return err
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorByConsAddrKeyPrefix))
	store.Set(types.ValidatorByConsAddrKey(
		consAddr,
	), validator.GetOwner())

	return nil
}

// get a single validator by consensus address.
func (k Keeper) GetValidatorByConsAddr(ctx sdk.Context, consAddr sdk.ConsAddress) (validator types.Validator, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorByConsAddrKeyPrefix))

	owner := store.Get(types.ValidatorByConsAddrKey(
		consAddr,
	))
	if owner == nil {
		return validator, false
	}

	return k.GetValidator(ctx, owner)
}

func (k Keeper) mustGetValidatorByConsAddr(ctx sdk.Context, consAddr sdk.ConsAddress) types.Validator {
	validator, found := k.GetValidatorByConsAddr(ctx, consAddr)
	if !found {
		panic(fmt.Errorf("validator with consensus-Address %s not found", consAddr))
	}

	return validator
}

// GetAllValidator returns all validator.
func (k Keeper) GetAllValidator(ctx sdk.Context) (list []types.Validator) {
	k.IterateValidators(ctx, func(validator types.Validator) (stop bool) {
		list = append(list, validator)

		return false
	})

	return
}

// iterate over validators and apply function.
func (k Keeper) IterateValidators(ctx sdk.Context, process func(validator types.Validator) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Validator
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if process(val) {
			return
		}
	}
}
