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

func createNPkiRevocationDistributionPointsByIssuerSubjectKeyId(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.PkiRevocationDistributionPointsByIssuerSubjectKeyId {
	items := make([]types.PkiRevocationDistributionPointsByIssuerSubjectKeyId, n)
	for i := range items {
		items[i].IssuerSubjectKeyId = strconv.Itoa(i)

		keeper.SetPkiRevocationDistributionPointsByIssuerSubjectKeyId(ctx, items[i])
	}
	return items
}

func TestPkiRevocationDistributionPointsByIssuerSubjectKeyIdGet(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNPkiRevocationDistributionPointsByIssuerSubjectKeyId(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyId(ctx,
			item.IssuerSubjectKeyId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPkiRevocationDistributionPointsByIssuerSubjectKeyIdRemove(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNPkiRevocationDistributionPointsByIssuerSubjectKeyId(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePkiRevocationDistributionPointsByIssuerSubjectKeyId(ctx,
			item.IssuerSubjectKeyId,
		)
		_, found := keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyId(ctx,
			item.IssuerSubjectKeyId,
		)
		require.False(t, found)
	}
}

func TestPkiRevocationDistributionPointsByIssuerSubjectKeyIdGetAll(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNPkiRevocationDistributionPointsByIssuerSubjectKeyId(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPkiRevocationDistributionPointsByIssuerSubjectKeyId(ctx)),
	)
}
