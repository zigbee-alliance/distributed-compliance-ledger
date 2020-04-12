package keeper

import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/internal/types"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/test_constants"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"testing"
)

func TestQuerier_QueryModel(t *testing.T) {
	setup := Setup()

	// add model
	modelInfo := DefaultModelInfo()
	setup.ModelinfoKeeper.SetModelInfo(setup.Ctx, modelInfo)

	// query model
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryModel, fmt.Sprintf("%v", modelInfo.VID), fmt.Sprintf("%v", modelInfo.PID)},
		abci.RequestQuery{},
	)

	var receivedModelInfo types.ModelInfo
	_ = setup.Cdc.UnmarshalJSON(result, &receivedModelInfo)

	// check
	require.Equal(t, receivedModelInfo.VID, modelInfo.VID)
	require.Equal(t, receivedModelInfo.PID, modelInfo.PID)
	require.Equal(t, receivedModelInfo.CID, modelInfo.CID)
	require.Equal(t, receivedModelInfo.Name, modelInfo.Name)
	require.Equal(t, receivedModelInfo.Description, modelInfo.Description)
}

func TestQuerier_QueryModelForUnknown(t *testing.T) {
	setup := Setup()

	// query model
	result, err := setup.Querier(
		setup.Ctx,
		[]string{QueryModel, fmt.Sprintf("%v", test_constants.VID), fmt.Sprintf("%v", test_constants.PID)},
		abci.RequestQuery{},
	)

	// check
	require.Nil(t, result)
	require.NotNil(t, err)
	require.Equal(t, types.CodeModelInfoDoesNotExist, err.Code())
}

func TestQuerier_QueryAllModels(t *testing.T) {
	setup := Setup()
	count := 5

	// add 5 models
	firstId := PopulateStoreWithWithModelsHavingDifferentVendor(setup, count)

	// query all models
	params := types.NewPaginationParams(0, 0)
	receiveModelInfos := getModels(setup, params)

	// check
	require.Equal(t, count, receiveModelInfos.Total)
	require.Equal(t, count, len(receiveModelInfos.Items))

	for i, item := range receiveModelInfos.Items {
		require.Equal(t, int16(i)+firstId, item.VID)
		require.Equal(t, int16(i)+firstId, item.PID)
	}
}

func TestQuerier_QueryAllModelsWithPaginationHeaders(t *testing.T) {
	setup := Setup()
	count := 5

	// add 5 models
	firstId := PopulateStoreWithWithModelsHavingDifferentVendor(setup, count)

	// query all models skip=1 take=2
	skip := 1
	take := 2
	params := types.NewPaginationParams(skip, take)
	receiveModelInfos := getModels(setup, params)

	// check
	require.Equal(t, count, receiveModelInfos.Total)
	require.Equal(t, take, len(receiveModelInfos.Items))

	for i, item := range receiveModelInfos.Items {
		require.Equal(t, int16(skip)+int16(i)+firstId, item.VID)
	}
}

func TestQuerier_QueryVendorsForModelsHaveDifferentVendors(t *testing.T) {
	setup := Setup()

	count := 5

	// add 5 models with different vendors
	firstId := PopulateStoreWithWithModelsHavingDifferentVendor(setup, count)

	params := types.NewPaginationParams(0, 0)

	// query all vendors
	receiveModelInfos := getVendors(setup, params)

	// check
	require.Equal(t, count, receiveModelInfos.Total)
	require.Equal(t, count, len(receiveModelInfos.Items))

	for i, item := range receiveModelInfos.Items {
		require.Equal(t, int16(i)+firstId, item.VID)
	}
}

func TestQuerier_QueryVendorsForModelsHaveSameVendor(t *testing.T) {
	setup := Setup()
	count := 5

	// add 5 models with same vendors
	firstId := PopulateStoreWithWithModelsHavingSameVendor(setup, count)

	params := types.NewPaginationParams(0, 0)

	// query all vendors
	receiveModelInfos := getVendors(setup, params)

	// check
	expectedCount := 1
	require.Equal(t, expectedCount, receiveModelInfos.Total)
	require.Equal(t, expectedCount, len(receiveModelInfos.Items))
	require.Equal(t, firstId, receiveModelInfos.Items[0].VID)
}

func TestQuerier_QueryVendorsWithPaginationHeaders(t *testing.T) {
	setup := Setup()
	count := 5

	// add 5 models with different vendor
	firstId := PopulateStoreWithWithModelsHavingDifferentVendor(setup, count)

	skip := 1
	take := 2
	params := types.NewPaginationParams(skip, take)

	// query vendors skip=1, take=2
	receiveModelInfos := getVendors(setup, params)

	// check
	require.Equal(t, count, receiveModelInfos.Total)
	require.Equal(t, take, len(receiveModelInfos.Items))

	for i, item := range receiveModelInfos.Items {
		require.Equal(t, int16(skip)+int16(i)+firstId, item.VID)
	}
}

func TestQuerier_QueryVendorModels(t *testing.T) {
	setup := Setup()
	count := 5

	// add 5 models with same vendors
	firstId := PopulateStoreWithWithModelsHavingSameVendor(setup, count)

	params := types.NewPaginationParams(0, 0)

	// query all models
	receivedVendorModels := getVendorModels(setup, firstId, params)

	// check
	require.Equal(t, count, receivedVendorModels.Total)
	require.Equal(t, count, len(receivedVendorModels.Items))

	for i, item := range receivedVendorModels.Items {
		require.Equal(t, firstId, item.VID)
		require.Equal(t, int16(i)+firstId, item.PID)
	}
}

func TestQuerier_QueryVendorModelsWithPaginationHeaders(t *testing.T) {
	setup := Setup()
	count := 5

	// add 5 models with same vendors
	firstId := PopulateStoreWithWithModelsHavingSameVendor(setup, count)

	skip := 1
	take := 2
	params := types.NewPaginationParams(skip, take)

	// query vendor models skip=1 take=2
	receivedVendorModels := getVendorModels(setup, firstId, params)

	// check
	require.Equal(t, count, receivedVendorModels.Total)
	require.Equal(t, take, len(receivedVendorModels.Items))

	for i, item := range receivedVendorModels.Items {
		require.Equal(t, firstId, item.VID)
		require.Equal(t, int16(skip)+int16(i)+firstId, item.PID)
	}
}

func getModels(setup TestSetup, params types.PaginationParams) types.LisModelInfoItems {
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllModels},
		abci.RequestQuery{Data: setup.Cdc.MustMarshalJSON(params)},
	)

	var receiveModelInfos types.LisModelInfoItems
	_ = setup.Cdc.UnmarshalJSON(result, &receiveModelInfos)

	return receiveModelInfos
}

func getVendors(setup TestSetup, params types.PaginationParams) types.LisVendorItems {
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryVendors},
		abci.RequestQuery{Data: setup.Cdc.MustMarshalJSON(params)},
	)

	var receiveModelInfos types.LisVendorItems
	_ = setup.Cdc.UnmarshalJSON(result, &receiveModelInfos)

	return receiveModelInfos
}

func getVendorModels(setup TestSetup, vid int16, params types.PaginationParams) types.LisModelInfoItems {
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryVendorModels, fmt.Sprintf("%v", vid)},
		abci.RequestQuery{Data: setup.Cdc.MustMarshalJSON(params)},
	)

	var receivedVendorModels types.LisModelInfoItems
	_ = setup.Cdc.UnmarshalJSON(result, &receivedVendorModels)
	return receivedVendorModels
}
