package pki

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/go-amino"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

type TestSetup struct {
	Cdc         *amino.Codec
	Ctx         sdk.Context
	PkiKeeper   Keeper
	AuthzKeeper authz.Keeper
	Handler     sdk.Handler
	Querier     sdk.Querier
	Trustee     sdk.AccAddress
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

	authzKey := sdk.NewKVStoreKey(authz.StoreKey)
	dbStore.MountStoreWithDB(authzKey, sdk.StoreTypeIAVL, db)

	_ = dbStore.LoadLatestVersion()

	// Init Keepers
	pkiKeeper := NewKeeper(pkiKey, cdc)
	authzKeeper := authz.NewKeeper(authzKey, cdc)

	// Create context
	ctx := sdk.NewContext(dbStore, abci.Header{ChainID: test_constants.ChainId}, false, log.NewNopLogger())

	// Create Handler and Querier
	querier := NewQuerier(pkiKeeper)
	handler := NewHandler(pkiKeeper, authzKeeper)

	trustee := test_constants.Address2
	authzKeeper.AssignRole(ctx, trustee, authz.Trustee)

	setup := TestSetup{
		Cdc:         cdc,
		Ctx:         ctx,
		PkiKeeper:   pkiKeeper,
		AuthzKeeper: authzKeeper,
		Handler:     handler,
		Querier:     querier,
		Trustee:     trustee,
	}

	return setup
}
