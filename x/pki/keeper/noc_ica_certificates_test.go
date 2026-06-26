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

func createNNocIcaCertificates(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.NocIcaCertificates {
	items := make([]types.NocIcaCertificates, n)
	for i := range items {
		items[i].Vid = int32(i)

		keeper.SetNocIcaCertificates(ctx, items[i])
	}

	return items
}

func TestNocIcaCertificatesGet(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNNocIcaCertificates(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetNocIcaCertificates(ctx,
			item.Vid,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestNocIcaCertificatesRemove(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNNocIcaCertificates(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNocIcaCertificates(ctx,
			item.Vid,
		)
		_, found := keeper.GetNocIcaCertificates(ctx,
			item.Vid,
		)
		require.False(t, found)
	}
}

func TestNocIcaCertificatesGetAll(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNNocIcaCertificates(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllNocIcaCertificates(ctx)),
	)
}

func TestGetNocIcaCertificatesBySubjectAndSKID(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	keeper.SetNocIcaCertificates(ctx, types.NocIcaCertificates{
		Vid: 1,
		Certs: []*types.Certificate{
			{Subject: "s1", SubjectKeyId: "k1"},
			{Subject: "s2", SubjectKeyId: "k2"},
		},
	})

	res, found := keeper.GetNocIcaCertificatesBySubjectAndSKID(ctx, 1, "s1", "k1")
	require.True(t, found)
	require.Len(t, res.Certs, 1)
	require.Equal(t, "s1", res.Certs[0].Subject)

	_, found = keeper.GetNocIcaCertificatesBySubjectAndSKID(ctx, 1, "nope", "nope")
	require.False(t, found)

	_, found = keeper.GetNocIcaCertificatesBySubjectAndSKID(ctx, 99, "s1", "k1")
	require.False(t, found)
}
