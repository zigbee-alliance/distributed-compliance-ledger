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

func createNPkiRevocationDistributionPoint(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.PkiRevocationDistributionPoint {
	items := make([]types.PkiRevocationDistributionPoint, n)
	for i := range items {
		items[i].Vid = uint64(i)
		items[i].Label = strconv.Itoa(i)
		items[i].IssuerSubjectKeyID = strconv.Itoa(i)

		keeper.SetPkiRevocationDistributionPoint(ctx, items[i])
	}
	return items
}

func TestPkiRevocationDistributionPointGet(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNPkiRevocationDistributionPoint(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPkiRevocationDistributionPoint(ctx,
			item.Vid,
			item.Label,
			item.IssuerSubjectKeyID,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPkiRevocationDistributionPointRemove(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNPkiRevocationDistributionPoint(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePkiRevocationDistributionPoint(ctx,
			item.Vid,
			item.Label,
			item.IssuerSubjectKeyID,
		)
		_, found := keeper.GetPkiRevocationDistributionPoint(ctx,
			item.Vid,
			item.Label,
			item.IssuerSubjectKeyID,
		)
		require.False(t, found)
	}
}

func TestPkiRevocationDistributionPointGetAll(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNPkiRevocationDistributionPoint(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPkiRevocationDistributionPoint(ctx)),
	)
}
