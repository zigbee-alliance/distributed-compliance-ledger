package keeper

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/test_constants"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
	"testing"
)

type testSetup struct {
	ctx              sdk.Context
	complianceKeeper Keeper
}

func setup() testSetup {
	// Init Codec
	cdc := codec.New()
	sdk.RegisterCodec(cdc)

	// Init KVSore
	db := dbm.NewMemDB()

	dbStore := store.NewCommitMultiStore(db)

	complianceKey := sdk.NewKVStoreKey(types.StoreKey)
	dbStore.MountStoreWithDB(complianceKey, sdk.StoreTypeIAVL, db)

	dbStore.LoadLatestVersion()

	// Init Keepers
	complianceKeeper := NewKeeper(complianceKey, cdc)

	// Create context
	ctx := sdk.NewContext(dbStore, abci.Header{ChainID: "zbl-test-chain-id"}, false, log.NewNopLogger())

	// Create handler (optional)

	setup := testSetup{ctx: ctx, complianceKeeper: complianceKeeper,}

	return setup
}

func TestModelInfoKeeperGetSet(t *testing.T) {
	setup := setup()

	// no model before its created
	require.Panics(t, func() {
		setup.complianceKeeper.GetModelInfo(setup.ctx, test_constants.Id)
	})

	// create model
	setup.complianceKeeper.SetModelInfo(setup.ctx, types.ModelInfo{
		ID:                       test_constants.Id,
		Name:                     test_constants.Name,
		Owner:                    test_constants.Owner,
		Description:              test_constants.Description,
		SKU:                      test_constants.Sku,
		FirmwareVersion:          test_constants.FirmwareVersion,
		HardwareVersion:          test_constants.HardwareVersion,
		CertificateID:            test_constants.CertificateID,
		CertifiedDate:            test_constants.CertifiedDate,
		TisOrTrpTestingCompleted: false,
	})

	// check if model present
	require.True(t, setup.complianceKeeper.IsModelInfoPresent(setup.ctx, test_constants.Id))

	// get model info
	modelInfo := setup.complianceKeeper.GetModelInfo(setup.ctx, test_constants.Id)
	require.NotNil(t, modelInfo)
	require.Equal(t, test_constants.Name, modelInfo.Name)
	require.Equal(t, test_constants.Owner, modelInfo.Owner)
}

func TestModelInfoKeeperIterator(t *testing.T) {
	setup := setup()

	// create model
	setup.complianceKeeper.SetModelInfo(setup.ctx, types.ModelInfo{
		ID:                       test_constants.Id,
		Name:                     test_constants.Name,
		Owner:                    test_constants.Owner,
		Description:              test_constants.Description,
		SKU:                      test_constants.Sku,
		FirmwareVersion:          test_constants.FirmwareVersion,
		HardwareVersion:          test_constants.HardwareVersion,
		CertificateID:            test_constants.CertificateID,
		CertifiedDate:            test_constants.CertifiedDate,
		TisOrTrpTestingCompleted: false,
	})

	// get total count
	totalModes := setup.complianceKeeper.CountTotal(setup.ctx)
	require.Equal(t, 1, totalModes)

	// get iterator
	var expectedRecords []types.ModelInfo

	setup.complianceKeeper.IterateModelInfos(setup.ctx, func(modelInfo types.ModelInfo) (stop bool) {
		expectedRecords = append(expectedRecords, modelInfo)
		return false
	})
	require.Equal(t, 1, len(expectedRecords))
}

func TestModelInfoKeeperDelete(t *testing.T) {
	setup := setup()

	// create model
	setup.complianceKeeper.SetModelInfo(setup.ctx, types.ModelInfo{
		ID:                       test_constants.Id,
		Name:                     test_constants.Name,
		Owner:                    test_constants.Owner,
		Description:              test_constants.Description,
		SKU:                      test_constants.Sku,
		FirmwareVersion:          test_constants.FirmwareVersion,
		HardwareVersion:          test_constants.HardwareVersion,
		CertificateID:            test_constants.CertificateID,
		CertifiedDate:            test_constants.CertifiedDate,
		TisOrTrpTestingCompleted: false,
	})

	// check if model present
	require.True(t, setup.complianceKeeper.IsModelInfoPresent(setup.ctx, test_constants.Id))

	// delete model
	setup.complianceKeeper.DeleteModelInfo(setup.ctx, test_constants.Id)

	// check if model present
	require.False(t, setup.complianceKeeper.IsModelInfoPresent(setup.ctx, test_constants.Id))

	// try to get model info
	require.Panics(t, func() {
		setup.complianceKeeper.GetModelInfo(setup.ctx, test_constants.Id)
	})
}
