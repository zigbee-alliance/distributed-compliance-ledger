package keeper

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
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
	dbStore.MountStoreWithDB(compliancetestKey, sdk.StoreTypeIAVL, db)
	_ = dbStore.LoadLatestVersion()

	// Init Keepers
	compliancetestKeeper := NewKeeper(compliancetestKey, cdc)

	// Init Querier
	querier := NewQuerier(compliancetestKeeper)

	// Create context
	ctx := sdk.NewContext(dbStore, abci.Header{ChainID: test_constants.ChainId}, false, log.NewNopLogger())

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
		VID:        test_constants.VID,
		PID:        test_constants.PID,
		TestResult: test_constants.TestResult,
		Owner:      test_constants.Owner,
	}
}
