package compliancetest

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/internal/types"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/go-amino"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
	"testing"
)

type TestSetup struct {
	Cdc                  *amino.Codec
	Ctx                  sdk.Context
	CompliancetestKeeper Keeper
	AuthzKeeper          authz.Keeper
	ModelinfoKeeper      modelinfo.Keeper
	Handler              sdk.Handler
	Querier              sdk.Querier
}

func Setup() TestSetup {
	// Init Codec
	cdc := codec.New()
	sdk.RegisterCodec(cdc)

	// Init KVSore
	db := dbm.NewMemDB()

	dbStore := store.NewCommitMultiStore(db)

	complianceKey := sdk.NewKVStoreKey(StoreKey)
	dbStore.MountStoreWithDB(complianceKey, sdk.StoreTypeIAVL, db)

	authzKey := sdk.NewKVStoreKey(authz.StoreKey)
	dbStore.MountStoreWithDB(authzKey, sdk.StoreTypeIAVL, db)

	modelinfoKey := sdk.NewKVStoreKey(modelinfo.StoreKey)
	dbStore.MountStoreWithDB(modelinfoKey, sdk.StoreTypeIAVL, db)

	_ = dbStore.LoadLatestVersion()

	// Init Keepers
	compliancetestKeeper := NewKeeper(complianceKey, cdc)
	authzKeeper := authz.NewKeeper(authzKey, cdc)
	modelinfoKeeper := modelinfo.NewKeeper(modelinfoKey, cdc)

	// Create context
	ctx := sdk.NewContext(dbStore, abci.Header{ChainID: test_constants.ChainId}, false, log.NewNopLogger())

	// Create Handler and Querier
	querier := NewQuerier(compliancetestKeeper)
	handler := NewHandler(compliancetestKeeper, modelinfoKeeper, authzKeeper)

	setup := TestSetup{
		Cdc:                  cdc,
		Ctx:                  ctx,
		CompliancetestKeeper: compliancetestKeeper,
		ModelinfoKeeper:      modelinfoKeeper,
		AuthzKeeper:          authzKeeper,
		Handler:              handler,
		Querier:              querier,
	}

	return setup
}

func (setup TestSetup) TestHouse(address sdk.AccAddress) sdk.AccAddress {
	setup.AuthzKeeper.AssignRole(setup.Ctx, address, authz.TestHouse)
	return address
}

func (setup TestSetup) Vendor(address sdk.AccAddress) sdk.AccAddress {
	setup.AuthzKeeper.AssignRole(setup.Ctx, address, authz.Vendor)
	return address
}

func (setup TestSetup) Administrator(address sdk.AccAddress) sdk.AccAddress {
	setup.AuthzKeeper.AssignRole(setup.Ctx, address, authz.Administrator)
	return address
}

func TestMsgAddTestingResult(signer sdk.AccAddress, vid int16, pid int16) MsgAddTestingResult {
	return MsgAddTestingResult{
		VID:        vid,
		PID:        pid,
		TestResult: test_constants.TestResult,
		TestDate:   test_constants.TestDate,
		Signer:     signer,
	}
}

func CheckTestingResult(t *testing.T, receivedTestingResult types.TestingResultItem, expectedTestingResult types.MsgAddTestingResult) {
	require.Equal(t, receivedTestingResult.Owner, expectedTestingResult.Signer)
	require.Equal(t, receivedTestingResult.TestResult, expectedTestingResult.TestResult)
	require.Equal(t, receivedTestingResult.TestDate, expectedTestingResult.TestDate)
}
