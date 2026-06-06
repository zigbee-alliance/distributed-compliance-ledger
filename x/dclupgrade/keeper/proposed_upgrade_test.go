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

package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/nullify"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/types"
)

// Prevent strconv unused error.
var _ = strconv.IntSize

func createNProposedUpgrade(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ProposedUpgrade {
	items := make([]types.ProposedUpgrade, n)
	for i := range items {
		items[i].Plan.Name = strconv.Itoa(i)

		keeper.SetProposedUpgrade(ctx, items[i])
	}

	return items
}

func TestProposedUpgradeGet(t *testing.T) {
	keeper, ctx := keepertest.DclupgradeKeeper(t, nil, nil)
	items := createNProposedUpgrade(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetProposedUpgrade(ctx,
			item.Plan.Name,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestProposedUpgradeRemove(t *testing.T) {
	keeper, ctx := keepertest.DclupgradeKeeper(t, nil, nil)
	items := createNProposedUpgrade(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveProposedUpgrade(ctx,
			item.Plan.Name,
		)
		_, found := keeper.GetProposedUpgrade(ctx,
			item.Plan.Name,
		)
		require.False(t, found)
	}
}

func TestProposedUpgradeGetAll(t *testing.T) {
	keeper, ctx := keepertest.DclupgradeKeeper(t, nil, nil)
	items := createNProposedUpgrade(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllProposedUpgrade(ctx)),
	)
}
