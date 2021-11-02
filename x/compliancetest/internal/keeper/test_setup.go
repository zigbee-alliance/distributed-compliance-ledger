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

//nolint:goimports
import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest/internal/types"
)

type TestSetup struct {
	Cdc                  *codec.Codec
	Ctx                  sdk.Context
	CompliancetestKeeper Keeper
	Querier              sdk.Querier
}

func Setup() TestSetup {
	// Init Codec
	cdc := codec.New()
	sdk.RegisterCodec(cdc)

	// Init KVSore
	db := dbm.NewMemDB()
	dbStore := store.NewCommitMultiStore(db)
	compliancetestKey := sdk.NewKVStoreKey(types.StoreKey)
	dbStore.MountStoreWithDB(compliancetestKey, sdk.StoreTypeIAVL, nil)
	_ = dbStore.LoadLatestVersion()

	// Init Keepers
	compliancetestKeeper := NewKeeper(compliancetestKey, cdc)

	// Init Querier
	querier := NewQuerier(compliancetestKeeper)

	// Create context
	ctx := sdk.NewContext(dbStore, abci.Header{ChainID: testconstants.ChainID}, false, log.NewNopLogger())

	setup := TestSetup{
		Cdc:                  cdc,
		Ctx:                  ctx,
		CompliancetestKeeper: compliancetestKeeper,
		Querier:              querier,
	}

	return setup
}

func DefaultTestingResult() types.TestingResult {
	return types.TestingResult{
		VID:                   testconstants.VID,
		PID:                   testconstants.PID,
		SoftwareVersion:       testconstants.SoftwareVersion,
		SoftwareVersionString: testconstants.SoftwareVersionString,
		TestResult:            testconstants.TestResult,
		TestDate:              testconstants.TestDate,
		Owner:                 testconstants.Owner,
	}
}

func CheckTestingResult(t *testing.T, receivedTestingResult types.TestingResultItem,
	expectedTestingResult types.TestingResult) {
	require.Equal(t, receivedTestingResult.Owner, expectedTestingResult.Owner)
	require.Equal(t, receivedTestingResult.TestResult, expectedTestingResult.TestResult)
	require.Equal(t, receivedTestingResult.TestDate, expectedTestingResult.TestDate)
}
