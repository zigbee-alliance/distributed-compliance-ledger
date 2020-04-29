package authnext

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/params/subspace"
	"github.com/tendermint/go-amino"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

type TestSetup struct {
	Cdc         *amino.Codec
	Ctx         sdk.Context
	Keeper      AccountKeeper
	AuthzKeeper authz.Keeper
	Handler     sdk.Handler
	Querier     sdk.Querier
}

func Setup() TestSetup {
	// Init Codec
	cdc := codec.New()
	sdk.RegisterCodec(cdc)
	types.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)

	// Init KVSore
	db := dbm.NewMemDB()

	dbStore := store.NewCommitMultiStore(db)

	storeKey := sdk.NewKVStoreKey(StoreKey)
	authCapKey := sdk.NewKVStoreKey("authCapKey")
	keyParams := sdk.NewKVStoreKey("subspace")
	tkeyParams := sdk.NewTransientStoreKey("transient_subspace")
	authzKey := sdk.NewTransientStoreKey(authz.StoreKey)

	dbStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	dbStore.MountStoreWithDB(authCapKey, sdk.StoreTypeIAVL, db)
	dbStore.MountStoreWithDB(keyParams, sdk.StoreTypeIAVL, db)
	dbStore.MountStoreWithDB(tkeyParams, sdk.StoreTypeTransient, db)
	dbStore.MountStoreWithDB(authzKey, sdk.StoreTypeIAVL, db)

	_ = dbStore.LoadLatestVersion()

	ps := subspace.NewSubspace(cdc, keyParams, tkeyParams, types.DefaultParamspace)
	accountKeeper := auth.NewAccountKeeper(cdc, storeKey, ps, types.ProtoBaseAccount)

	authzKeeper := authz.NewKeeper(authzKey, cdc)

	// Create context
	ctx := sdk.NewContext(dbStore, abci.Header{ChainID: test_constants.ChainId}, false, log.NewNopLogger())

	// Create Handler and Querier
	querier := NewQuerier(accountKeeper, authzKeeper, cdc)
	handler := NewHandler(accountKeeper, authzKeeper, cdc)

	setup := TestSetup{
		Cdc:         cdc,
		Ctx:         ctx,
		Keeper:      accountKeeper,
		AuthzKeeper: authzKeeper,
		Handler:     handler,
		Querier:     querier,
	}

	return setup
}
