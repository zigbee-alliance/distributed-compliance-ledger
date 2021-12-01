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

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/pagination"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/internal/types"
)

func TestQuerier_QueryModel(t *testing.T) {
	setup := Setup()

	// add model
	model := DefaultModel()
	setup.ModelKeeper.SetModel(setup.Ctx, model)

	// query model
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryModel, fmt.Sprintf("%v", model.VID), fmt.Sprintf("%v", model.PID)},
		abci.RequestQuery{},
	)

	var receivedModel types.Model
	_ = setup.Cdc.UnmarshalJSON(result, &receivedModel)

	// check
	require.Equal(t, receivedModel, model)
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
	require.Equal(t, types.CodeModelDoesNotExist, err.Code())
}

func TestQuerier_QueryAllModels(t *testing.T) {
	setup := Setup()
	count := 5

	// add 5 models
	firstID := PopulateStoreWithModelsHavingDifferentVendor(setup, count)

	// query all models
	params := pagination.NewPaginationParams(0, 0)
	receiveModels := getModels(setup, params)

	// check
	require.Equal(t, count, receiveModels.Total)
	require.Equal(t, count, len(receiveModels.Items))

	for i, item := range receiveModels.Items {
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
	receiveModels := getModels(setup, params)

	// check
	require.Equal(t, count, receiveModels.Total)
	require.Equal(t, take, len(receiveModels.Items))

	for i, item := range receiveModels.Items {
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
	receiveModels := getVendors(setup, params)

	// check
	require.Equal(t, count, receiveModels.Total)
	require.Equal(t, count, len(receiveModels.Items))

	for i, item := range receiveModels.Items {
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
	receiveModels := getVendors(setup, params)

	// check
	expectedCount := 1
	require.Equal(t, expectedCount, receiveModels.Total)
	require.Equal(t, expectedCount, len(receiveModels.Items))
	require.Equal(t, firstID, receiveModels.Items[0].VID)
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
	receiveModels := getVendors(setup, params)

	// check
	require.Equal(t, count, receiveModels.Total)
	require.Equal(t, take, len(receiveModels.Items))

	for i, item := range receiveModels.Items {
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

func getModels(setup TestSetup, params pagination.PaginationParams) types.ListModelItems {
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllModels},
		abci.RequestQuery{Data: setup.Cdc.MustMarshalJSON(params)},
	)

	var receiveModels types.ListModelItems
	_ = setup.Cdc.UnmarshalJSON(result, &receiveModels)

	return receiveModels
}

func getVendors(setup TestSetup, params pagination.PaginationParams) types.ListVendorItems {
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryVendors},
		abci.RequestQuery{Data: setup.Cdc.MustMarshalJSON(params)},
	)

	var receiveModels types.ListVendorItems
	_ = setup.Cdc.UnmarshalJSON(result, &receiveModels)

	return receiveModels
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
