package compliance

//nolint:goimports
import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/go-amino"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

type TestSetup struct {
	Cdc                  *amino.Codec
	Ctx                  sdk.Context
	CompliancetKeeper    Keeper
	CompliancetestKeeper compliancetest.Keeper
	authKeeper           auth.Keeper
	ModelinfoKeeper      modelinfo.Keeper
	Handler              sdk.Handler
	Querier              sdk.Querier
	CertificationCenter  sdk.AccAddress
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

	authKey := sdk.NewKVStoreKey(auth.StoreKey)
	dbStore.MountStoreWithDB(authKey, sdk.StoreTypeIAVL, db)

	modelinfoKey := sdk.NewKVStoreKey(modelinfo.StoreKey)
	dbStore.MountStoreWithDB(modelinfoKey, sdk.StoreTypeIAVL, db)

	compliancetestKey := sdk.NewKVStoreKey(compliancetest.StoreKey)
	dbStore.MountStoreWithDB(compliancetestKey, sdk.StoreTypeIAVL, db)

	_ = dbStore.LoadLatestVersion()

	// Init Keepers
	compliancetKeeper := NewKeeper(complianceKey, cdc)
	compliancetestKeeper := compliancetest.NewKeeper(compliancetestKey, cdc)
	authKeeper := auth.NewKeeper(authKey, cdc)
	modelinfoKeeper := modelinfo.NewKeeper(modelinfoKey, cdc)

	// Create context
	ctx := sdk.NewContext(dbStore, abci.Header{ChainID: testconstants.ChainID}, false, log.NewNopLogger())

	// Create Handler and Querier
	querier := NewQuerier(compliancetKeeper)
	handler := NewHandler(compliancetKeeper, modelinfoKeeper, compliancetestKeeper, authKeeper)

	account := auth.NewAccount(testconstants.Address1, testconstants.PubKey1, auth.AccountRoles{auth.ZBCertificationCenter})
	authKeeper.AssignNumberAndStoreAccount(ctx, account)

	setup := TestSetup{
		Cdc:                  cdc,
		Ctx:                  ctx,
		CompliancetKeeper:    compliancetKeeper,
		CompliancetestKeeper: compliancetestKeeper,
		ModelinfoKeeper:      modelinfoKeeper,
		authKeeper:           authKeeper,
		Handler:              handler,
		Querier:              querier,
		CertificationCenter:  account.Address,
	}

	return setup
}
