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
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

// SetAccountStat set accountStat in the store
// NOTE:
//	- the API is needed by genesis logic
//	- but it shouldn't be used in run-time, so makes sense to take care about better way
func (k Keeper) SetAccountStat(ctx sdk.Context, accountStat types.AccountStat) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountStatKey))
	b := k.cdc.MustMarshal(&accountStat)
	store.Set([]byte{0}, b)
}

// GetAccountStat returns accountStat.
func (k Keeper) GetAccountStat(ctx sdk.Context) (val types.AccountStat, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountStatKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

/*
// RemoveAccountStat removes accountStat from the store
func (k Keeper) RemoveAccountStat(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountStatKey))
	store.Delete([]byte{0})
}
*/

func (k Keeper) GetNextAccountNumber(ctx sdk.Context) (accNumber uint64) {
	accountStat, found := k.GetAccountStat(ctx)

	if !found {
		accountStat = types.AccountStat{
			Number: 0,
		}
	}

	accNumber = accountStat.Number

	accountStat.Number += 1

	// that logic is not exposed as API intentionally
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountStatKey))
	b := k.cdc.MustMarshal(&accountStat)
	store.Set([]byte{0}, b)

	return
}
