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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func createTestRevokedRootCertificates(keeper *keeper.Keeper, ctx sdk.Context) types.RevokedRootCertificates {
	item := types.RevokedRootCertificates{}
	keeper.SetRevokedRootCertificates(ctx, item)
	return item
}

func TestRevokedRootCertificatesGet(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t)
	item := createTestRevokedRootCertificates(keeper, ctx)
	rst, found := keeper.GetRevokedRootCertificates(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestRevokedRootCertificatesRemove(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t)
	createTestRevokedRootCertificates(keeper, ctx)
	keeper.RemoveRevokedRootCertificates(ctx)
	_, found := keeper.GetRevokedRootCertificates(ctx)
	require.False(t, found)
}
*/
