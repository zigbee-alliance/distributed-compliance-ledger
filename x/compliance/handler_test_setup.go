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
	AuthzKeeper          auth.Keeper
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

	authzKey := sdk.NewKVStoreKey(auth.StoreKey)
	dbStore.MountStoreWithDB(authzKey, sdk.StoreTypeIAVL, db)

	modelinfoKey := sdk.NewKVStoreKey(modelinfo.StoreKey)
	dbStore.MountStoreWithDB(modelinfoKey, sdk.StoreTypeIAVL, db)

	compliancetestKey := sdk.NewKVStoreKey(compliancetest.StoreKey)
	dbStore.MountStoreWithDB(compliancetestKey, sdk.StoreTypeIAVL, db)

	_ = dbStore.LoadLatestVersion()

	// Init Keepers
	compliancetKeeper := NewKeeper(complianceKey, cdc)
	compliancetestKeeper := compliancetest.NewKeeper(compliancetestKey, cdc)
	authzKeeper := auth.NewKeeper(authzKey, cdc)
	modelinfoKeeper := modelinfo.NewKeeper(modelinfoKey, cdc)

	// Create context
	ctx := sdk.NewContext(dbStore, abci.Header{ChainID: testconstants.ChainID}, false, log.NewNopLogger())

	// Create Handler and Querier
	querier := NewQuerier(compliancetKeeper)
	handler := NewHandler(compliancetKeeper, modelinfoKeeper, compliancetestKeeper, authzKeeper)

	certificationCenter := testconstants.Address1
	authzKeeper.AssignRole(ctx, certificationCenter, auth.ZBCertificationCenter)

	setup := TestSetup{
		Cdc:                  cdc,
		Ctx:                  ctx,
		CompliancetKeeper:    compliancetKeeper,
		CompliancetestKeeper: compliancetestKeeper,
		ModelinfoKeeper:      modelinfoKeeper,
		AuthzKeeper:          authzKeeper,
		Handler:              handler,
		Querier:              querier,
		CertificationCenter:  certificationCenter,
	}

	return setup
}
