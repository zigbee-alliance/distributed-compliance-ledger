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

// import (
// 	"strconv"
// 	"testing"

// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	"github.com/stretchr/testify/require"
// 	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
// 	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/keeper"
// 	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/types"
// )

// // Prevent strconv unused error.
// var _ = strconv.IntSize

// func createNVendorInfo(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.VendorInfo {
// 	items := make([]types.VendorInfo, n)
// 	for i := range items {
// 		items[i].VendorID = int32(i)

// 		keeper.SetVendorInfo(ctx, items[i])
// 	}
// 	return items
// }

// func TestVendorInfoGet(t *testing.T) {
// 	keeper, ctx := keepertest.VendorinfoKeeper(t)
// 	items := createNVendorInfo(keeper, ctx, 10)
// 	for _, item := range items {
// 		rst, found := keeper.GetVendorInfo(ctx,
// 			item.VendorID,
// 		)
// 		require.True(t, found)
// 		require.Equal(t, item, rst)
// 	}
// }

// func TestVendorInfoRemove(t *testing.T) {
// 	keeper, ctx := keepertest.VendorinfoKeeper(t)
// 	items := createNVendorInfo(keeper, ctx, 10)
// 	for _, item := range items {
// 		keeper.RemoveVendorInfo(ctx,
// 			item.VendorID,
// 		)
// 		_, found := keeper.GetVendorInfo(ctx,
// 			item.VendorID,
// 		)
// 		require.False(t, found)
// 	}
// }

// func TestVendorInfoGetAll(t *testing.T) {
// 	keeper, ctx := keepertest.VendorinfoKeeper(t)
// 	items := createNVendorInfo(keeper, ctx, 10)
// 	require.ElementsMatch(t, items, keeper.GetAllVendorInfo(ctx))
// }
