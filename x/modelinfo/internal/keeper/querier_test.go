// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//nolint:testpackage
package keeper

//nolint:goimports
import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/pagination"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/modelinfo/internal/types"
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
		[]string{QueryModel, fmt.Sprintf("%v", testconstants.VID), fmt.Sprintf("%v", testconstants.PID)},
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
	firstID := PopulateStoreWithModelsHavingDifferentVendor(setup, count)

	// query all models
	params := pagination.NewPaginationParams(0, 0)
	receiveModelInfos := getModels(setup, params)

	// check
	require.Equal(t, count, receiveModelInfos.Total)
	require.Equal(t, count, len(receiveModelInfos.Items))

	for i, item := range receiveModelInfos.Items {
		require.Equal(t, uint16(i)+firstID, item.VID)
		require.Equal(t, uint16(i)+firstID, item.PID)
	}
}

func TestQuerier_QueryAllModelsWithPaginationHeaders(t *testing.T) {
	setup := Setup()
	count := 5

	// add 5 models
	firstID := PopulateStoreWithModelsHavingDifferentVendor(setup, count)

	// query all models skip=1 take=2
	skip := 1
	take := 2
	params := pagination.NewPaginationParams(skip, take)
	receiveModelInfos := getModels(setup, params)

	// check
	require.Equal(t, count, receiveModelInfos.Total)
	require.Equal(t, take, len(receiveModelInfos.Items))

	for i, item := range receiveModelInfos.Items {
		require.Equal(t, uint16(skip)+uint16(i)+firstID, item.VID)
	}
}

func TestQuerier_QueryVendorsForModelsHaveDifferentVendors(t *testing.T) {
	setup := Setup()

	count := 5

	// add 5 models with different vendors
	firstID := PopulateStoreWithModelsHavingDifferentVendor(setup, count)

	params := pagination.NewPaginationParams(0, 0)

	// query all vendors
	receiveModelInfos := getVendors(setup, params)

	// check
	require.Equal(t, count, receiveModelInfos.Total)
	require.Equal(t, count, len(receiveModelInfos.Items))

	for i, item := range receiveModelInfos.Items {
		require.Equal(t, uint16(i)+firstID, item.VID)
	}
}

func TestQuerier_QueryVendorsForModelsHaveSameVendor(t *testing.T) {
	setup := Setup()
	count := 5

	// add 5 models with same vendors
	firstID := PopulateStoreWithModelsHavingSameVendor(setup, count)

	params := pagination.NewPaginationParams(0, 0)

	// query all vendors
	receiveModelInfos := getVendors(setup, params)

	// check
	expectedCount := 1
	require.Equal(t, expectedCount, receiveModelInfos.Total)
	require.Equal(t, expectedCount, len(receiveModelInfos.Items))
	require.Equal(t, firstID, receiveModelInfos.Items[0].VID)
}

func TestQuerier_QueryVendorsWithPaginationHeaders(t *testing.T) {
	setup := Setup()
	count := 5

	// add 5 models with different vendor
	firstID := PopulateStoreWithModelsHavingDifferentVendor(setup, count)

	skip := 1
	take := 2
	params := pagination.NewPaginationParams(skip, take)

	// query vendors skip=1, take=2
	receiveModelInfos := getVendors(setup, params)

	// check
	require.Equal(t, count, receiveModelInfos.Total)
	require.Equal(t, take, len(receiveModelInfos.Items))

	for i, item := range receiveModelInfos.Items {
		require.Equal(t, uint16(skip)+uint16(i)+firstID, item.VID)
	}
}

func TestQuerier_QueryVendorModels(t *testing.T) {
	setup := Setup()
	count := 5

	// add 5 models with same vendors
	firstID := PopulateStoreWithModelsHavingSameVendor(setup, count)

	// query all models
	receivedVendorModels := getVendorModels(setup, firstID)

	// check
	require.Equal(t, count, len(receivedVendorModels.Products))

	for i, item := range receivedVendorModels.Products {
		require.Equal(t, uint16(i)+firstID, item.PID)
	}
}

func getModels(setup TestSetup, params pagination.PaginationParams) types.LisModelInfoItems {
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllModels},
		abci.RequestQuery{Data: setup.Cdc.MustMarshalJSON(params)},
	)

	var receiveModelInfos types.LisModelInfoItems
	_ = setup.Cdc.UnmarshalJSON(result, &receiveModelInfos)

	return receiveModelInfos
}

func getVendors(setup TestSetup, params pagination.PaginationParams) types.ListVendorItems {
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryVendors},
		abci.RequestQuery{Data: setup.Cdc.MustMarshalJSON(params)},
	)

	var receiveModelInfos types.ListVendorItems
	_ = setup.Cdc.UnmarshalJSON(result, &receiveModelInfos)

	return receiveModelInfos
}

func getVendorModels(setup TestSetup, vid uint16) types.VendorProducts {
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryVendorModels, fmt.Sprintf("%v", vid)},
		abci.RequestQuery{},
	)

	var receivedVendorModels types.VendorProducts
	_ = setup.Cdc.UnmarshalJSON(result, &receivedVendorModels)

	return receivedVendorModels
}
