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
	Cdc         *amino.Codec
	Ctx         sdk.Context
	PkiKeeper   Keeper
	AuthzKeeper auth.Keeper
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

	authzKey := sdk.NewKVStoreKey(auth.StoreKey)
	dbStore.MountStoreWithDB(authzKey, sdk.StoreTypeIAVL, db)

	_ = dbStore.LoadLatestVersion()

	// Init Keepers
	pkiKeeper := NewKeeper(pkiKey, cdc)
	authzKeeper := auth.NewKeeper(authzKey, cdc)

	// Create context
	ctx := sdk.NewContext(dbStore, abci.Header{ChainID: testconstants.ChainID}, false, log.NewNopLogger())

	// Create Handler and Querier
	querier := NewQuerier(pkiKeeper)
	handler := NewHandler(pkiKeeper, authzKeeper)

	trustee := testconstants.Address2
	authzKeeper.AssignRole(ctx, trustee, auth.Trustee)

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
