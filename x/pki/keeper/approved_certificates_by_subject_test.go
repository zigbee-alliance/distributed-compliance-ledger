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

func createNApprovedCertificatesBySubject(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ApprovedCertificatesBySubject {
	items := make([]types.ApprovedCertificatesBySubject, n)
	for i := range items {
		items[i].Subject = strconv.Itoa(i)

		keeper.SetApprovedCertificatesBySubject(ctx, items[i])
	}
	return items
}

func TestApprovedCertificatesBySubjectGet(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t)
	items := createNApprovedCertificatesBySubject(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetApprovedCertificatesBySubject(ctx,
			item.Subject,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestApprovedCertificatesBySubjectRemove(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t)
	items := createNApprovedCertificatesBySubject(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveApprovedCertificatesBySubject(ctx,
			item.Subject,
		)
		_, found := keeper.GetApprovedCertificatesBySubject(ctx,
			item.Subject,
		)
		require.False(t, found)
	}
}

func TestApprovedCertificatesBySubjectGetAll(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t)
	items := createNApprovedCertificatesBySubject(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllApprovedCertificatesBySubject(ctx)),
	)
}
