package keeper

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/internal/types"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/test_constants"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

type TestSetup struct {
	Cdc             *codec.Codec
	Ctx             sdk.Context
	ModelinfoKeeper Keeper
	Querier         sdk.Querier
}

func Setup() TestSetup {
	// Init Codec
	cdc := codec.New()
	sdk.RegisterCodec(cdc)

	// Init KVSore
	db := dbm.NewMemDB()
	dbStore := store.NewCommitMultiStore(db)
	modelinfoKey := sdk.NewKVStoreKey(types.StoreKey)
	dbStore.MountStoreWithDB(modelinfoKey, sdk.StoreTypeIAVL, db)
	_ = dbStore.LoadLatestVersion()

	// Init Keepers
	modelinfoKeeper := NewKeeper(modelinfoKey, cdc)

	// Init Querier
	querier := NewQuerier(modelinfoKeeper)

	// Create context
	ctx := sdk.NewContext(dbStore, abci.Header{ChainID: "zbl-test-chain-id"}, false, log.NewNopLogger())

	setup := TestSetup{
		Cdc:             cdc,
		Ctx:             ctx,
		ModelinfoKeeper: modelinfoKeeper,
		Querier:         querier,
	}
	return setup
}

func DefaultModelInfo() types.ModelInfo {
	return types.ModelInfo{
		VID:                      test_constants.VID,
		PID:                      test_constants.PID,
		CID:                      test_constants.CID,
		Name:                     test_constants.Name,
		Owner:                    test_constants.Owner,
		Description:              test_constants.Description,
		SKU:                      test_constants.Sku,
		FirmwareVersion:          test_constants.FirmwareVersion,
		HardwareVersion:          test_constants.HardwareVersion,
		Custom:                   test_constants.Custom,
		CertificateID:            test_constants.CertificateID,
		CertifiedDate:            test_constants.CertifiedDate,
		TisOrTrpTestingCompleted: test_constants.TisOrTrpTestingCompleted,
	}
}

// add 10 models with same VID and check associated products {VID: 1, PID: 1..count}
func PopulateStoreWithWithModelsHavingSameVendor(setup TestSetup, count int) int16 {
	firstId := int16(1)

	modelInfo := DefaultModelInfo()
	modelInfo.VID = firstId

	for i := firstId; i <= int16(count); i++ {
		// add model info {VID: 1, PID: i}
		modelInfo.PID = i
		setup.ModelinfoKeeper.SetModelInfo(setup.Ctx, modelInfo)
	}
	return firstId
}

// add 10 models with same VID and check associated products {VID: 1..count, PID: 1..count}
func PopulateStoreWithWithModelsHavingDifferentVendor(setup TestSetup, count int) int16 {
	firstId := int16(1)

	modelInfo := DefaultModelInfo()

	for i := firstId; i <= int16(count); i++ {
		// add model info {VID: i, PID: i}
		modelInfo.VID = i
		modelInfo.PID = i
		setup.ModelinfoKeeper.SetModelInfo(setup.Ctx, modelInfo)
	}
	return firstId
}
