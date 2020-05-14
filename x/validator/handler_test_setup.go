package validator

//nolint:goimports
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
	Cdc             *amino.Codec
	Ctx             sdk.Context
	ValidatorKeeper Keeper
	authKeeper      auth.Keeper
	Handler         sdk.Handler
	Querier         sdk.Querier
	NodeAdmin       sdk.AccAddress
}

func Setup() TestSetup {
	// Init Codec
	cdc := codec.New()
	sdk.RegisterCodec(cdc)

	// Init KVSore
	db := dbm.NewMemDB()

	dbStore := store.NewCommitMultiStore(db)

	validatorKey := sdk.NewKVStoreKey(StoreKey)
	dbStore.MountStoreWithDB(validatorKey, sdk.StoreTypeIAVL, db)

	authKey := sdk.NewKVStoreKey(auth.StoreKey)
	dbStore.MountStoreWithDB(authKey, sdk.StoreTypeIAVL, db)

	_ = dbStore.LoadLatestVersion()

	// Init Keepers
	validatorKeeper := NewKeeper(validatorKey, cdc)
	authKeeper := auth.NewKeeper(authKey, cdc)

	// Create context
	ctx := sdk.NewContext(dbStore, abci.Header{ChainID: testconstants.ChainID}, false, log.NewNopLogger())

	// Create Handler and Querier
	querier := NewQuerier(validatorKeeper)
	handler := NewHandler(validatorKeeper, authKeeper)

	authKeeper.AssignRole(ctx, testconstants.Address1, auth.NodeAdmin)

	setup := TestSetup{
		Cdc:             cdc,
		Ctx:             ctx,
		ValidatorKeeper: validatorKeeper,
		authKeeper:      authKeeper,
		Handler:         handler,
		Querier:         querier,
		NodeAdmin:       testconstants.Address1,
	}

	return setup
}
