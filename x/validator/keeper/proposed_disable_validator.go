// Copyright 2020 DSR Corporation
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

	defer func() { _ = iterator.Close() }()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ProposedDisableValidator
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
