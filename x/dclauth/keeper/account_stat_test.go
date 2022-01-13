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
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/nullify"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func createTestAccountStat(keeper *keeper.Keeper, ctx sdk.Context) types.AccountStat {
	item := types.AccountStat{}
	keeper.SetAccountStat(ctx, item)
	return item
}

func TestAccountStatGet(t *testing.T) {
	keeper, ctx := keepertest.DclauthKeeper(t)
	item := createTestAccountStat(keeper, ctx)
	rst, found := keeper.GetAccountStat(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestAccountStatRemove(t *testing.T) {
	keeper, ctx := keepertest.DclauthKeeper(t)
	createTestAccountStat(keeper, ctx)
	keeper.RemoveAccountStat(ctx)
	_, found := keeper.GetAccountStat(ctx)
	require.False(t, found)
}
*/
