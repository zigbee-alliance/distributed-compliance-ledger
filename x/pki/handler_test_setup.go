package pki

// nolint:goimports
import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/go-amino"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

type TestSetup struct {
	Cdc        *amino.Codec
	Ctx        sdk.Context
	PkiKeeper  Keeper
	authKeeper auth.Keeper
	Handler    sdk.Handler
	Querier    sdk.Querier
	Trustee    sdk.AccAddress
}

func Setup() TestSetup {
	// Init Codec
	cdc := codec.New()
	sdk.RegisterCodec(cdc)

	// Init KVSore
	db := dbm.NewMemDB()

	dbStore := store.NewCommitMultiStore(db)

	pkiKey := sdk.NewKVStoreKey(StoreKey)
	dbStore.MountStoreWithDB(pkiKey, sdk.StoreTypeIAVL, db)

	authKey := sdk.NewKVStoreKey(auth.StoreKey)
	dbStore.MountStoreWithDB(authKey, sdk.StoreTypeIAVL, db)

	_ = dbStore.LoadLatestVersion()

	// Init Keepers
	pkiKeeper := NewKeeper(pkiKey, cdc)
	authKeeper := auth.NewKeeper(authKey, cdc)

	// Create context
	ctx := sdk.NewContext(dbStore, abci.Header{ChainID: testconstants.ChainID}, false, log.NewNopLogger())

	// Create Handler and Querier
	querier := NewQuerier(pkiKeeper)
	handler := NewHandler(pkiKeeper, authKeeper)

	account := auth.NewAccount(testconstants.Address2, testconstants.PubKey2, auth.AccountRoles{auth.Trustee})
	authKeeper.AssignNumberAndStoreAccount(ctx, account)

	setup := TestSetup{
		Cdc:        cdc,
		Ctx:        ctx,
		PkiKeeper:  pkiKeeper,
		authKeeper: authKeeper,
		Handler:    handler,
		Querier:    querier,
		Trustee:    account.Address,
	}

	return setup
}
