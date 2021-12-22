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

func createNProposedCertificate(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ProposedCertificate {
	items := make([]types.ProposedCertificate, n)
	for i := range items {
		items[i].Subject = strconv.Itoa(i)
		items[i].SubjectKeyId = strconv.Itoa(i)

		keeper.SetProposedCertificate(ctx, items[i])
	}
	return items
}

func TestProposedCertificateGet(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t)
	items := createNProposedCertificate(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetProposedCertificate(ctx,
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
func TestProposedCertificateRemove(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t)
	items := createNProposedCertificate(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveProposedCertificate(ctx,
			item.Subject,
			item.SubjectKeyId,
		)
		_, found := keeper.GetProposedCertificate(ctx,
			item.Subject,
			item.SubjectKeyId,
		)
		require.False(t, found)
	}
}

func TestProposedCertificateGetAll(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t)
	items := createNProposedCertificate(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllProposedCertificate(ctx)),
	)
}
