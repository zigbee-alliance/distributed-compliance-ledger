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

func createNProposedCertificateRevocation(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ProposedCertificateRevocation {
	items := make([]types.ProposedCertificateRevocation, n)
	for i := range items {
		items[i].Subject = strconv.Itoa(i)
		items[i].SubjectKeyId = strconv.Itoa(i)

		keeper.SetProposedCertificateRevocation(ctx, items[i])
	}
	return items
}

func TestProposedCertificateRevocationGet(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t)
	items := createNProposedCertificateRevocation(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetProposedCertificateRevocation(ctx,
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
func TestProposedCertificateRevocationRemove(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t)
	items := createNProposedCertificateRevocation(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveProposedCertificateRevocation(ctx,
			item.Subject,
			item.SubjectKeyId,
		)
		_, found := keeper.GetProposedCertificateRevocation(ctx,
			item.Subject,
			item.SubjectKeyId,
		)
		require.False(t, found)
	}
}

func TestProposedCertificateRevocationGetAll(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t)
	items := createNProposedCertificateRevocation(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllProposedCertificateRevocation(ctx)),
	)
}
