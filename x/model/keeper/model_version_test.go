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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

// Prevent strconv unused error.
var _ = strconv.IntSize

func createNModelVersion(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ModelVersion {
	items := make([]types.ModelVersion, n)
	for i := range items {
		items[i].Vid = int32(i)
		items[i].Pid = int32(i)
		items[i].SoftwareVersion = uint32(i)

		keeper.SetModelVersion(ctx, items[i])
	}
	return items
}

func TestModelVersionGet(t *testing.T) {
	keeper, ctx := keepertest.ModelKeeper(t)
	items := createNModelVersion(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetModelVersion(ctx,
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

func TestModelVersionRemove(t *testing.T) {
	keeper, ctx := keepertest.ModelKeeper(t)
	items := createNModelVersion(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveModelVersion(ctx,
			item.Vid,
			item.Pid,
			item.SoftwareVersion,
		)
		_, found := keeper.GetModelVersion(ctx,
			item.Vid,
			item.Pid,
			item.SoftwareVersion,
		)
		require.False(t, found)
	}
}

func TestModelVersionGetAll(t *testing.T) {
	keeper, ctx := keepertest.ModelKeeper(t)
	items := createNModelVersion(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllModelVersion(ctx)),
	)
}
*/
