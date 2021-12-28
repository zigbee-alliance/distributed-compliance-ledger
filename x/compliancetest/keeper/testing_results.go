package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest/types"
)

// SetTestingResults set a specific testingResults in the store from its index
func (k Keeper) SetTestingResults(ctx sdk.Context, testingResults types.TestingResults) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TestingResultsKeyPrefix))
	b := k.cdc.MustMarshal(&testingResults)
	store.Set(types.TestingResultsKey(
		testingResults.Vid,
		testingResults.Pid,
		testingResults.SoftwareVersion,
	), b)
}

// GetTestingResults returns a testingResults from its index
func (k Keeper) GetTestingResults(
	ctx sdk.Context,
	vid int32,
	pid int32,
	softwareVersion uint32,

) (val types.TestingResults, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TestingResultsKeyPrefix))

	b := store.Get(types.TestingResultsKey(
		vid,
		pid,
		softwareVersion,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTestingResults removes a testingResults from the store
func (k Keeper) RemoveTestingResults(
	ctx sdk.Context,
	vid int32,
	pid int32,
	softwareVersion uint32,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TestingResultsKeyPrefix))
	store.Delete(types.TestingResultsKey(
		vid,
		pid,
		softwareVersion,
	))
}

// GetAllTestingResults returns all testingResults
func (k Keeper) GetAllTestingResults(ctx sdk.Context) (list []types.TestingResults) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TestingResultsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TestingResults
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// Add a testing result to the list of results
func (k Keeper) AppendTestingResult(ctx sdk.Context, testingResult types.TestingResult) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TestingResultsKeyPrefix))

	testingResultsBytes := store.Get(types.TestingResultsKey(
		testingResult.Vid,
		testingResult.Pid,
		testingResult.SoftwareVersion,
	))
	var testingResults types.TestingResults

	if testingResultsBytes == nil {
		testingResults = types.TestingResults{
			Vid:             testingResult.Vid,
			Pid:             testingResult.Pid,
			SoftwareVersion: testingResult.SoftwareVersion,
			Results:         []*types.TestingResult{},
		}
	} else {
		k.cdc.MustUnmarshal(testingResultsBytes, &testingResults)
	}

	testingResults.Results = append(testingResults.Results, &testingResult)

	b := k.cdc.MustMarshal(&testingResults)
	store.Set(types.TestingResultsKey(
		testingResult.Vid,
		testingResult.Pid,
		testingResult.SoftwareVersion,
	), b)
}
