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

func createNPKIRevocationDistributionPoint(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.PKIRevocationDistributionPoint {
	items := make([]types.PKIRevocationDistributionPoint, n)
	for i := range items {
		items[i].Vid = uint64(i)
		items[i].Label = strconv.Itoa(i)
		items[i].IssuerSubjectKeyID = strconv.Itoa(i)

		keeper.SetPKIRevocationDistributionPoint(ctx, items[i])
	}
	return items
}

func TestPKIRevocationDistributionPointGet(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNPKIRevocationDistributionPoint(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPKIRevocationDistributionPoint(ctx,
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
func TestPKIRevocationDistributionPointRemove(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNPKIRevocationDistributionPoint(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePKIRevocationDistributionPoint(ctx,
			item.Vid,
			item.Label,
			item.IssuerSubjectKeyID,
		)
		_, found := keeper.GetPKIRevocationDistributionPoint(ctx,
			item.Vid,
			item.Label,
			item.IssuerSubjectKeyID,
		)
		require.False(t, found)
	}
}

func TestPKIRevocationDistributionPointGetAll(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNPKIRevocationDistributionPoint(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPKIRevocationDistributionPoint(ctx)),
	)
}
