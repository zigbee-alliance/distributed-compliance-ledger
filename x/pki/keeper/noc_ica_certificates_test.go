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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// Prevent strconv unused error.
var _ = strconv.IntSize

func createNNocIcaCertificates(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.NocIcaCertificates {
	items := make([]types.NocIcaCertificates, n)
	for i := range items {
		items[i].Vid = int32(i)

		keeper.SetNocIcaCertificates(ctx, items[i])
	}

	return items
}

func TestNocIcaCertificatesGet(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNNocIcaCertificates(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetNocIcaCertificates(ctx,
			item.Vid,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestNocIcaCertificatesRemove(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNNocIcaCertificates(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNocIcaCertificates(ctx,
			item.Vid,
		)
		_, found := keeper.GetNocIcaCertificates(ctx,
			item.Vid,
		)
		require.False(t, found)
	}
}

func TestNocIcaCertificatesGetAll(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNNocIcaCertificates(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllNocIcaCertificates(ctx)),
	)
}
