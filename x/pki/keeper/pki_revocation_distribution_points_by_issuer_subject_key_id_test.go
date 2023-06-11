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

func createNPkiRevocationDistributionPointsByIssuerSubjectKeyID(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.PkiRevocationDistributionPointsByIssuerSubjectKeyID {
	items := make([]types.PkiRevocationDistributionPointsByIssuerSubjectKeyID, n)
	for i := range items {
		items[i].IssuerSubjectKeyID = strconv.Itoa(i)

		keeper.SetPkiRevocationDistributionPointsByIssuerSubjectKeyID(ctx, items[i])
	}

	return items
}

func TestPkiRevocationDistributionPointsByIssuerSubjectKeyIDGet(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNPkiRevocationDistributionPointsByIssuerSubjectKeyID(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(ctx,
			item.IssuerSubjectKeyID,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPkiRevocationDistributionPointsByIssuerSubjectKeyIDRemove(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNPkiRevocationDistributionPointsByIssuerSubjectKeyID(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePkiRevocationDistributionPointsByIssuerSubjectKeyID(ctx,
			item.IssuerSubjectKeyID,
		)
		_, found := keeper.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(ctx,
			item.IssuerSubjectKeyID,
		)
		require.False(t, found)
	}
}

func TestPkiRevocationDistributionPointsByIssuerSubjectKeyIDGetAll(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNPkiRevocationDistributionPointsByIssuerSubjectKeyID(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPkiRevocationDistributionPointsByIssuerSubjectKeyID(ctx)),
	)
}
