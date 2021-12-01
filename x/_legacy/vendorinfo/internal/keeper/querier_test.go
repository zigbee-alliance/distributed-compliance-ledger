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
	types "github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/internal/types"
)

func TestQuerier_QueryVendorInfo(t *testing.T) {
	setup := Setup()

	// add vendorInfo
	vendorInfo := DefaultVendorInfo()
	setup.VendorInfoKeeper.SetVendorInfo(setup.Ctx, vendorInfo)

	// query vendorInfo
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryVendor, fmt.Sprintf("%v", vendorInfo.VendorID)},
		abci.RequestQuery{},
	)

	var receivedVendorInfo types.VendorInfo
	_ = setup.Cdc.UnmarshalJSON(result, &receivedVendorInfo)

	// check
	require.Equal(t, receivedVendorInfo, vendorInfo)
}

func TestQuerier_QueryVendorInfoForUnknown(t *testing.T) {
	setup := Setup()

	// query model
	result, err := setup.Querier(
		setup.Ctx,
		[]string{QueryVendor, fmt.Sprintf("%v", testconstants.VID)},
		abci.RequestQuery{},
	)

	// check
	require.Nil(t, result)
	require.NotNil(t, err)
	require.Equal(t, types.CodeVendorDoesNotExist, err.Code())
}

func TestQuerier_QueryAllVendorInfos(t *testing.T) {
	setup := Setup()
	count := 5

	// add 5 vendorInfos
	firstID := PopulateStoreWithVendorInfo(setup, count)

	// query all vendorInfos
	params := pagination.NewPaginationParams(0, 0)
	receiveModels := getVendorInfos(setup, params)

	// check
	require.Equal(t, count, receiveModels.Total)
	require.Equal(t, count, len(receiveModels.Vendors))

	for i, item := range receiveModels.Vendors {
		require.Equal(t, uint16(i)+firstID, item.VendorID)
	}
}

func TestQuerier_QueryAllVendorInfosWithPaginationHeaders(t *testing.T) {
	setup := Setup()
	count := 5

	// add 5 models
	firstID := PopulateStoreWithVendorInfo(setup, count)

	// query all models skip=1 take=2
	skip := 1
	take := 2
	params := pagination.NewPaginationParams(skip, take)
	receivedVendorInfos := getVendorInfos(setup, params)

	// check
	require.Equal(t, count, receivedVendorInfos.Total)
	require.Equal(t, take, len(receivedVendorInfos.Vendors))

	for i, item := range receivedVendorInfos.Vendors {
		require.Equal(t, uint16(skip)+uint16(i)+firstID, item.VendorID)
	}
}

func getVendorInfos(setup TestSetup, params pagination.PaginationParams) types.ListVendors {
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllVendors},
		abci.RequestQuery{Data: setup.Cdc.MustMarshalJSON(params)},
	)

	var receiveModels types.ListVendors
	_ = setup.Cdc.UnmarshalJSON(result, &receiveModels)

	return receiveModels
}
