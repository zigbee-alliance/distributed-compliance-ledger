// Copyright 2022 DSR Corporation
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

package keeper_test

/*
import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/nullify"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest/types"
)

// Prevent strconv unused error.
var _ = strconv.IntSize

func createNTestingResults(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.TestingResults {
	items := make([]types.TestingResults, n)
	for i := range items {
		items[i].Vid = int32(i)
		items[i].Pid = int32(i)
		items[i].SoftwareVersion = uint32(i)

		keeper.SetTestingResults(ctx, items[i])
	}
	return items
}

func TestTestingResultsGet(t *testing.T) {
	keeper, ctx := keepertest.CompliancetestKeeper(t)
	items := createNTestingResults(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetTestingResults(ctx,
			item.Vid,
			item.Pid,
			item.SoftwareVersion,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestTestingResultsRemove(t *testing.T) {
	keeper, ctx := keepertest.CompliancetestKeeper(t)
	items := createNTestingResults(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTestingResults(ctx,
			item.Vid,
			item.Pid,
			item.SoftwareVersion,
		)
		_, found := keeper.GetTestingResults(ctx,
			item.Vid,
			item.Pid,
			item.SoftwareVersion,
		)
		require.False(t, found)
	}
}

func TestTestingResultsGetAll(t *testing.T) {
	keeper, ctx := keepertest.CompliancetestKeeper(t)
	items := createNTestingResults(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTestingResults(ctx)),
	)
}
*/
