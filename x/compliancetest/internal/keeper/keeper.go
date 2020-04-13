package keeper

import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	// Unexposed key to access store from sdk.Context
	storeKey sdk.StoreKey

	// The wire codec for binary encoding/decoding
	cdc *codec.Codec
}

const (
	testingResultPrefix = "tr"
)

func NewKeeper(storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{storeKey: storeKey, cdc: cdc}
}

// Gets the entire TestingResults record for VID/PID combination
func (k Keeper) GetTestingResults(ctx sdk.Context, vid int16, pid int16) types.TestingResults {
	if !k.IsTestingResultsPresents(ctx, vid, pid) {
		return types.NewTestingResults(vid, pid)
	}

	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(TestingResultId(vid, pid)))

	var testingResults types.TestingResults
	k.cdc.MustUnmarshalBinaryBare(bz, &testingResults)
	return testingResults
}

// Sets the entire TestingResults struct for a TestingResultID
func (k Keeper) SetTestingResults(ctx sdk.Context, testingResult types.TestingResults) {
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(TestingResultId(testingResult.VID, testingResult.PID)), k.cdc.MustMarshalBinaryBare(testingResult))
}

// Add TestingResult for an existing TestingResults record
func (k Keeper) AddTestingResult(ctx sdk.Context, testingResult types.TestingResult) sdk.Error {
	testingResults := k.GetTestingResults(ctx, testingResult.VID, testingResult.PID)

	if testingResults.ContainsTestingResult(testingResult.Owner) {
		return types.ErrTestingResultAlreadyExists()
	}

	testingResults.AddTestingResult(testingResult)

	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(TestingResultId(testingResult.VID, testingResult.PID)), k.cdc.MustMarshalBinaryBare(testingResults))

	return nil
}

// Check if the TestingResults record is present in the store or not
func (k Keeper) IsTestingResultsPresents(ctx sdk.Context, vid int16, pid int16) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(TestingResultId(vid, pid)))
}

// Check if the TestingResult is present in the store or not
func (k Keeper) IsTestingResultPresents(ctx sdk.Context, vid int16, pid int16, owner sdk.AccAddress) bool {
	testingResults := k.GetTestingResults(ctx, vid, pid)
	return testingResults.ContainsTestingResult(owner)
}

// Iterate over all TestingResults
func (k Keeper) IterateTestingResults(ctx sdk.Context, process func(info types.TestingResults) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, []byte(testingResultPrefix))
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

// Id builder for TestingResultItem
func TestingResultId(vid interface{}, pid interface{}) string {
	return fmt.Sprintf("%s:%v:%v", testingResultPrefix, vid, pid)
}
