package compliancetest

//nolint:goimports
import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth"
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
	authKeeper           auth.Keeper
	ModelinfoKeeper      modelinfo.Keeper
	Handler              sdk.Handler
	Querier              sdk.Querier
	TestHouse            sdk.AccAddress
}

func Setup() TestSetup {
	// Init Codec
	cdc := codec.New()
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)

	// Init KVSore
	db := dbm.NewMemDB()

	dbStore := store.NewCommitMultiStore(db)

	complianceKey := sdk.NewKVStoreKey(StoreKey)
	dbStore.MountStoreWithDB(complianceKey, sdk.StoreTypeIAVL, nil)

	authKey := sdk.NewKVStoreKey(auth.StoreKey)
	dbStore.MountStoreWithDB(authKey, sdk.StoreTypeIAVL, nil)

	modelinfoKey := sdk.NewKVStoreKey(modelinfo.StoreKey)
	dbStore.MountStoreWithDB(modelinfoKey, sdk.StoreTypeIAVL, nil)

	_ = dbStore.LoadLatestVersion()

	// Init Keepers
	compliancetestKeeper := NewKeeper(complianceKey, cdc)
	authKeeper := auth.NewKeeper(authKey, cdc)
	modelinfoKeeper := modelinfo.NewKeeper(modelinfoKey, cdc)

	// Create context
	ctx := sdk.NewContext(dbStore, abci.Header{ChainID: testconstants.ChainID}, false, log.NewNopLogger())

	// Create Handler and Querier
	querier := NewQuerier(compliancetestKeeper)
	handler := NewHandler(compliancetestKeeper, modelinfoKeeper, authKeeper)

	account := auth.NewAccount(testconstants.Address1, testconstants.PubKey1, auth.AccountRoles{auth.TestHouse})
	authKeeper.AssignNumberAndStoreAccount(ctx, account)

	setup := TestSetup{
		Cdc:                  cdc,
		Ctx:                  ctx,
		CompliancetestKeeper: compliancetestKeeper,
		ModelinfoKeeper:      modelinfoKeeper,
		authKeeper:           authKeeper,
		Handler:              handler,
		Querier:              querier,
		TestHouse:            account.Address,
	}

	return setup
}

func TestMsgAddTestingResult(signer sdk.AccAddress, vid uint16, pid uint16) MsgAddTestingResult {
	return MsgAddTestingResult{
		VID:        vid,
		PID:        pid,
		TestResult: testconstants.TestResult,
		TestDate:   testconstants.TestDate,
		Signer:     signer,
	}
}

func CheckTestingResult(t *testing.T, receivedTestingResult types.TestingResultItem,
	expectedTestingResult types.MsgAddTestingResult) {
	require.Equal(t, receivedTestingResult.Owner, expectedTestingResult.Signer)
	require.Equal(t, receivedTestingResult.TestResult, expectedTestingResult.TestResult)
	require.Equal(t, receivedTestingResult.TestDate, expectedTestingResult.TestDate)
}
