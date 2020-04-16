package keeper

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

type TestSetup struct {
	Cdc               *codec.Codec
	Ctx               sdk.Context
	CompliancetKeeper Keeper
	Querier           sdk.Querier
}

func Setup() TestSetup {
	// Init Codec
	cdc := codec.New()
	sdk.RegisterCodec(cdc)

	// Init KVSore
	db := dbm.NewMemDB()
	dbStore := store.NewCommitMultiStore(db)
	complianceKey := sdk.NewKVStoreKey(types.StoreKey)
	dbStore.MountStoreWithDB(complianceKey, sdk.StoreTypeIAVL, db)
	_ = dbStore.LoadLatestVersion()

	// Init Keepers
	complianceKeeper := NewKeeper(complianceKey, cdc)

	// Init Querier
	querier := NewQuerier(complianceKeeper)

	// Create context
	ctx := sdk.NewContext(dbStore, abci.Header{ChainID: test_constants.ChainId}, false, log.NewNopLogger())

	setup := TestSetup{
		Cdc:               cdc,
		Ctx:               ctx,
		CompliancetKeeper: complianceKeeper,
		Querier:           querier,
	}
	return setup
}

func DefaultCertifiedModel() types.CertifiedModel {
	return types.CertifiedModel{
		VID:               test_constants.VID,
		PID:               test_constants.PID,
		CertificationDate: test_constants.CertificationDate,
		CertificationType: test_constants.CertificationType,
		Owner:             test_constants.Owner,
	}
}

// add 10 certified models {VID: 1, PID: 1..count}
func PopulateStoreWithCertifiedModels(setup TestSetup, count int) int16 {
	firstId := int16(1)
	model := DefaultCertifiedModel()
	for i := firstId; i <= int16(count); i++ {
		// add model {VID: 1, PID: i}
		model.PID = i
		model.VID = i
		setup.CompliancetKeeper.SetCertifiedModel(setup.Ctx, model)
	}
	return firstId
}
