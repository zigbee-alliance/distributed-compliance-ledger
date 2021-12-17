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

func createNApprovedCertificates(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ApprovedCertificates {
	items := make([]types.ApprovedCertificates, n)
	for i := range items {
		items[i].Subject = strconv.Itoa(i)
		items[i].SubjectKeyId = strconv.Itoa(i)

		keeper.SetApprovedCertificates(ctx, items[i])
	}
	return items
}

func TestApprovedCertificatesGet(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t)
	items := createNApprovedCertificates(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetApprovedCertificates(ctx,
			item.Subject,
			item.SubjectKeyId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestApprovedCertificatesRemove(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t)
	items := createNApprovedCertificates(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveApprovedCertificates(ctx,
			item.Subject,
			item.SubjectKeyId,
		)
		_, found := keeper.GetApprovedCertificates(ctx,
			item.Subject,
			item.SubjectKeyId,
		)
		require.False(t, found)
	}
}

func TestApprovedCertificatesGetAll(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t)
	items := createNApprovedCertificates(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllApprovedCertificates(ctx)),
	)
}
