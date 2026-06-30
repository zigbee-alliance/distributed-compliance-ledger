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

func createNNocCertificatesBySubjectKeyId(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.NocCertificatesBySubjectKeyID {
	items := make([]types.NocCertificatesBySubjectKeyID, n)
	for i := range items {
		items[i].SubjectKeyId = strconv.Itoa(i)

		keeper.SetNocCertificatesBySubjectKeyID(ctx, items[i])
	}

	return items
}

func TestNocCertificatesBySubjectKeyIdGet(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNNocCertificatesBySubjectKeyId(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetNocCertificatesBySubjectKeyID(ctx,
			item.SubjectKeyId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestNocCertificatesBySubjectKeyIdRemove(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNNocCertificatesBySubjectKeyId(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNocCertificatesBySubjectKeyID(ctx,
			item.SubjectKeyId,
		)
		_, found := keeper.GetNocCertificatesBySubjectKeyID(ctx,
			item.SubjectKeyId,
		)
		require.False(t, found)
	}
}

func TestNocCertificatesBySubjectKeyIdGetAll(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNNocCertificatesBySubjectKeyId(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllNocCertificatesBySubjectKeyID(ctx)),
	)
}

func TestAddGetRemoveNocCertificatesBySubjectKeyID(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)

	keeper.AddNocCertificatesBySubjectKeyID(ctx, types.NocCertificates{
		SubjectKeyId: "skid",
		Certs: []*types.Certificate{
			{Subject: "s1", SubjectKeyId: "skid"},
		},
	})

	got, found := keeper.GetNocCertificatesBySubjectKeyID(ctx, "skid")
	require.True(t, found)
	require.Len(t, got.Certs, 1)
	require.Equal(t, "s1", got.Certs[0].Subject)

	keeper.RemoveNocCertificatesBySubjectAndSubjectKeyID(ctx, "s1", "skid")

	_, found = keeper.GetNocCertificatesBySubjectKeyID(ctx, "skid")
	require.False(t, found)
}
