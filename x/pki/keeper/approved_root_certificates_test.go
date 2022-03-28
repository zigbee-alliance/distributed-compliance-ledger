package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/nullify"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func createTestApprovedRootCertificates(keeper *keeper.Keeper, ctx sdk.Context) types.ApprovedRootCertificates {
	item := types.ApprovedRootCertificates{}
	keeper.SetApprovedRootCertificates(ctx, item)

	return item
}

func TestApprovedRootCertificatesGet(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	item := createTestApprovedRootCertificates(keeper, ctx)
	rst, found := keeper.GetApprovedRootCertificates(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestApprovedRootCertificatesRemove(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	createTestApprovedRootCertificates(keeper, ctx)
	keeper.RemoveApprovedRootCertificates(ctx)
	_, found := keeper.GetApprovedRootCertificates(ctx)
	require.False(t, found)
}
