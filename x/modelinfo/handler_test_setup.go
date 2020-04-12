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

func (setup TestSetup) Vendor() sdk.AccAddress {
	acc, _ := sdk.AccAddressFromBech32("cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz")
	setup.AuthzKeeper.AssignRole(setup.Ctx, acc, authz.Vendor)
	return acc
}

func (setup TestSetup) Administrator() sdk.AccAddress {
	acc, _ := sdk.AccAddressFromBech32("cosmos1j8x9urmqs7p44va5p4cu29z6fc3g0cx2c2vxx2")
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
		CertificateID:            test_constants.CertificateID,
		CertifiedDate:            test_constants.CertifiedDate,
		TisOrTrpTestingCompleted: test_constants.TisOrTrpTestingCompleted,
		Signer:                   signer,
	}
}

func TestMsgUpdatedModelInfo(signer sdk.AccAddress) MsgUpdateModelInfo {
	return MsgUpdateModelInfo{
		VID:                         test_constants.VID,
		PID:                         test_constants.PID,
		NewCID:                      test_constants.CID,
		NewDescription:              "New Description",
		NewCustom:                   test_constants.Custom,
		NewCertificateID:            test_constants.CertificateID,
		NewCertifiedDate:            test_constants.CertifiedDate,
		NewTisOrTrpTestingCompleted: test_constants.TisOrTrpTestingCompleted,
		Signer:                      signer,
	}
}
