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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/modelversion/internal/types"
)

type TestSetup struct {
	Cdc                *codec.Codec
	Ctx                sdk.Context
	ModelKeeper        model.Keeper
	ModelVersionKeeper Keeper
	Querier            sdk.Querier
}

func Setup() TestSetup {
	// Init Codec
	cdc := codec.New()
	sdk.RegisterCodec(cdc)

	// Init KVSore
	db := dbm.NewMemDB()
	dbStore := store.NewCommitMultiStore(db)
	modelKey := sdk.NewKVStoreKey(model.StoreKey)
	modelVersionKey := sdk.NewKVStoreKey(types.StoreKey)
	dbStore.MountStoreWithDB(modelKey, sdk.StoreTypeIAVL, nil)
	dbStore.MountStoreWithDB(modelVersionKey, sdk.StoreTypeIAVL, nil)
	_ = dbStore.LoadLatestVersion()

	// Init Keepers
	modelKeeper := model.NewKeeper(modelKey, cdc)
	modelVersionKeeper := NewKeeper(modelVersionKey, cdc)

	// Init Querier
	querier := NewQuerier(modelVersionKeeper)

	// Create context
	ctx := sdk.NewContext(dbStore, abci.Header{ChainID: "dcl-test-chain-id"}, false, log.NewNopLogger())

	setup := TestSetup{
		Cdc:                cdc,
		Ctx:                ctx,
		ModelKeeper:        modelKeeper,
		ModelVersionKeeper: modelVersionKeeper,
		Querier:            querier,
	}

	return setup
}

func DefaultModelVersion() types.ModelVersion {
	return types.ModelVersion{
		VID:                          testconstants.VID,
		PID:                          testconstants.PID,
		SoftwareVersion:              testconstants.SoftwareVersion,
		SoftwareVersionString:        testconstants.SoftwareVersionString,
		CDVersionNumber:              testconstants.CDVersionNumber,
		FirmwareDigests:              testconstants.FirmwareDigests,
		SoftwareVersionValid:         testconstants.SoftwareVersionValid,
		OtaURL:                       testconstants.OtaURL,
		OtaFileSize:                  testconstants.OtaFileSize,
		OtaChecksum:                  testconstants.OtaChecksum,
		OtaChecksumType:              testconstants.OtaChecksumType,
		MinApplicableSoftwareVersion: testconstants.MinApplicableSoftwareVersion,
		MaxApplicableSoftwareVersion: testconstants.MaxApplicableSoftwareVersion,
		ReleaseNotesURL:              testconstants.ReleaseNotesURL,
	}
}

func DefaultModel() model.Model {
	model := model.Model{
		VID:          testconstants.VID,
		PID:          testconstants.PID,
		DeviceTypeID: testconstants.DeviceTypeID,
		ProductName:  testconstants.ProductName,
		ProductLabel: testconstants.ProductLabel,
		PartNumber:   testconstants.PartNumber,
	}

	return model
}

// add 10 model versions with same VID/PID and check associated versions
func PopulateStoreWithModelVersions(setup TestSetup, count int) (uint16, uint16) {
	firstID := uint32(1)

	model := DefaultModel()
	setup.ModelKeeper.SetModel(setup.Ctx, model)
	modelVersion := DefaultModelVersion()
	for i := firstID; i <= uint32(count); i++ {
		// add model versions info {VID: 1, PID: i}
		modelVersion.SoftwareVersion = i
		setup.ModelVersionKeeper.SetModelVersion(setup.Ctx, modelVersion)

	}
	return modelVersion.VID, modelVersion.PID
}

// add 10 model versions with different VID/PID and check associated versions
func PopulateStoreWithModelsHavingDifferentVendor(setup TestSetup, count int) uint16 {
	firstID := uint16(1)

	model := DefaultModel()
	modelVersion := DefaultModelVersion()

	for i := firstID; i <= uint16(count); i++ {
		// add model info {VID: i, PID: i}
		model.VID = i
		model.PID = i
		modelVersion.VID = i
		modelVersion.PID = i
		modelVersion.SoftwareVersion = uint32(i)
		setup.ModelKeeper.SetModel(setup.Ctx, model)
		setup.ModelVersionKeeper.SetModelVersion(setup.Ctx, modelVersion)
	}

	return firstID
}

func AddModel(setup TestSetup, vid uint16, pid uint16) {
	model := model.Model{
		VID:          vid,
		PID:          pid,
		DeviceTypeID: testconstants.DeviceTypeID,
		ProductName:  testconstants.ProductName,
		ProductLabel: testconstants.ProductLabel,
		PartNumber:   testconstants.PartNumber,
	}

	setup.ModelKeeper.SetModel(setup.Ctx, model)

}

func AddModelVersion(setup TestSetup) types.ModelVersion {
	modelVersion := DefaultModelVersion()
	setup.ModelVersionKeeper.SetModelVersion(setup.Ctx, DefaultModelVersion())
	return modelVersion
}
