package modelinfo

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/test_constants"
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
	ModelinfoKeeper Keeper
	AuthzKeeper     authz.Keeper
	Handler         sdk.Handler
	Querier         sdk.Querier
}

func Setup() TestSetup {
	// Init Codec
	cdc := codec.New()
	sdk.RegisterCodec(cdc)

	// Init KVSore
	db := dbm.NewMemDB()

	dbStore := store.NewCommitMultiStore(db)

	modelinfoKey := sdk.NewKVStoreKey(StoreKey)
	dbStore.MountStoreWithDB(modelinfoKey, sdk.StoreTypeIAVL, db)

	authzKey := sdk.NewKVStoreKey(authz.StoreKey)
	dbStore.MountStoreWithDB(authzKey, sdk.StoreTypeIAVL, db)

	_ = dbStore.LoadLatestVersion()

	// Init Keepers
	modelinfoKeeper := NewKeeper(modelinfoKey, cdc)
	authzKeeper := authz.NewKeeper(authzKey, cdc)

	// Create context
	ctx := sdk.NewContext(dbStore, abci.Header{ChainID: test_constants.ChainId}, false, log.NewNopLogger())

	// Create Handler and Querier
	querier := NewQuerier(modelinfoKeeper)
	handler := NewHandler(modelinfoKeeper, authzKeeper)

	setup := TestSetup{
		Cdc:             cdc,
		Ctx:             ctx,
		ModelinfoKeeper: modelinfoKeeper,
		AuthzKeeper:     authzKeeper,
		Handler:         handler,
		Querier:         querier,
	}

	return setup
}

func (setup TestSetup) Vendor(acc sdk.AccAddress) sdk.AccAddress {
	setup.AuthzKeeper.AssignRole(setup.Ctx, acc, authz.Vendor)
	return acc
}

func (setup TestSetup) Administrator(acc sdk.AccAddress) sdk.AccAddress {
	setup.AuthzKeeper.AssignRole(setup.Ctx, acc, authz.Administrator)
	return acc
}

func TestMsgAddModelInfo(signer sdk.AccAddress) MsgAddModelInfo {
	return MsgAddModelInfo{
		VID:                      test_constants.VID,
		PID:                      test_constants.PID,
		CID:                      test_constants.CID,
		Name:                     test_constants.Name,
		Description:              test_constants.Description,
		SKU:                      test_constants.Sku,
		FirmwareVersion:          test_constants.FirmwareVersion,
		HardwareVersion:          test_constants.HardwareVersion,
		Custom:                   test_constants.Custom,
		TisOrTrpTestingCompleted: test_constants.TisOrTrpTestingCompleted,
		Signer:                   signer,
	}
}

func TestMsgUpdatedModelInfo(signer sdk.AccAddress) MsgUpdateModelInfo {
	return MsgUpdateModelInfo{
		VID:                      test_constants.VID,
		PID:                      test_constants.PID,
		CID:                      test_constants.CID + 1,
		Description:              "New Description",
		Custom:                   "New Custom Data",
		TisOrTrpTestingCompleted: test_constants.TisOrTrpTestingCompleted,
		Signer:                   signer,
	}
}
