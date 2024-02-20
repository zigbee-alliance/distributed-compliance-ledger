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

// Prevent strconv unused error
var _ = strconv.IntSize

func createNNocRootCertificates(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.NocRootCertificates {
	items := make([]types.NocRootCertificates, n)
	for i := range items {
		items[i].Vid = int32(i)

		keeper.SetNocRootCertificates(ctx, items[i])
	}

	return items
}

func TestNocRootCertificatesGet(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNNocRootCertificates(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetNocRootCertificates(ctx,
			item.Vid,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestNocRootCertificatesRemove(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNNocRootCertificates(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNocRootCertificates(ctx,
			item.Vid,
		)
		_, found := keeper.GetNocRootCertificates(ctx,
			item.Vid,
		)
		require.False(t, found)
	}
}

func TestNocRootCertificatesGetAll(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNNocRootCertificates(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllNocRootCertificates(ctx)),
	)
}
