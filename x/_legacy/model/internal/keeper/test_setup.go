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
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/internal/types"
)

type TestSetup struct {
	Cdc         *codec.Codec
	Ctx         sdk.Context
	ModelKeeper Keeper
	Querier     sdk.Querier
}

func Setup() TestSetup {
	// Init Codec
	cdc := codec.New()
	sdk.RegisterCodec(cdc)

	// Init KVSore
	db := dbm.NewMemDB()
	dbStore := store.NewCommitMultiStore(db)
	modelKey := sdk.NewKVStoreKey(types.StoreKey)
	dbStore.MountStoreWithDB(modelKey, sdk.StoreTypeIAVL, nil)
	_ = dbStore.LoadLatestVersion()

	// Init Keepers
	modelKeeper := NewKeeper(modelKey, cdc)

	// Init Querier
	querier := NewQuerier(modelKeeper)

	// Create context
	ctx := sdk.NewContext(dbStore, abci.Header{ChainID: "dcl-test-chain-id"}, false, log.NewNopLogger())

	setup := TestSetup{
		Cdc:         cdc,
		Ctx:         ctx,
		ModelKeeper: modelKeeper,
		Querier:     querier,
	}

	return setup
}

func DefaultModel() types.Model {
	model := types.Model{
		VID:          testconstants.VID,
		PID:          testconstants.PID,
		DeviceTypeID: testconstants.DeviceTypeID,
		ProductName:  testconstants.ProductName,
		ProductLabel: testconstants.ProductLabel,
		PartNumber:   testconstants.PartNumber,
	}

	return model
}

// add 10 models with same VID and check associated products {VID: 1, PID: 1..count}.
func PopulateStoreWithModelsHavingSameVendor(setup TestSetup, count int) uint16 {
	firstID := uint16(1)

	model := DefaultModel()
	model.VID = firstID

	for i := firstID; i <= uint16(count); i++ {
		// add model info {VID: 1, PID: i}
		model.PID = i
		setup.ModelKeeper.SetModel(setup.Ctx, model)
	}

	return firstID
}

// add 10 models with same VID and check associated products {VID: 1..count, PID: 1..count}.
func PopulateStoreWithModelsHavingDifferentVendor(setup TestSetup, count int) uint16 {
	firstID := uint16(1)

	model := DefaultModel()

	for i := firstID; i <= uint16(count); i++ {
		// add model info {VID: i, PID: i}
		model.VID = i
		model.PID = i
		setup.ModelKeeper.SetModel(setup.Ctx, model)
	}

	return firstID
}
