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
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	// Unexposed key to access store from sdk.Context.
	storeKey sdk.StoreKey

	// The wire codec for binary encoding/decoding.
	cdc *codec.Codec
}

func NewKeeper(storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{storeKey: storeKey, cdc: cdc}
}

// Gets the entire TestingResults record for VID/PID combination.
func (k Keeper) GetTestingResults(ctx sdk.Context, vid uint16, pid uint16) types.TestingResults {
	if !k.IsTestingResultsPresents(ctx, vid, pid) {
		return types.NewTestingResults(vid, pid)
	}

	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetTestingResultsKey(vid, pid))

	var testingResults types.TestingResults

	k.cdc.MustUnmarshalBinaryBare(bz, &testingResults)

	return testingResults
}

// Sets the entire TestingResults record for a VID/PID combination.
func (k Keeper) SetTestingResults(ctx sdk.Context, testingResult types.TestingResults) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetTestingResultsKey(
		testingResult.VID, testingResult.PID), k.cdc.MustMarshalBinaryBare(testingResult))
}

// Add single TestingResult for an existing TestingResults record.
func (k Keeper) AddTestingResult(ctx sdk.Context, testingResult types.TestingResult) {
	testingResults := k.GetTestingResults(ctx, testingResult.VID, testingResult.PID)

	testingResults.AddTestingResult(testingResult)

	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetTestingResultsKey(testingResult.VID, testingResult.PID),
		k.cdc.MustMarshalBinaryBare(testingResults))
}

// Check if the TestingResults record is present in the store or not.
func (k Keeper) IsTestingResultsPresents(ctx sdk.Context, vid uint16, pid uint16) bool {
	store := ctx.KVStore(k.storeKey)

	return store.Has(types.GetTestingResultsKey(vid, pid))
}

// Iterate over all TestingResults records.
func (k Keeper) IterateTestingResults(ctx sdk.Context, process func(info types.TestingResults) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, types.TestingResultsPrefix)
	defer iter.Close()

	for {
		if !iter.Valid() {
			return
		}

		val := iter.Value()

		var testingResult types.TestingResults

		k.cdc.MustUnmarshalBinaryBare(val, &testingResult)

		if process(testingResult) {
			return
		}

		iter.Next()
	}
}
