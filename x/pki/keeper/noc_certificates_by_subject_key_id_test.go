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

func createNNocCertificatesBySubjectKeyId(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.NocCertificatesBySubjectKeyId {
	items := make([]types.NocCertificatesBySubjectKeyId, n)
	for i := range items {
		items[i].SubjectKeyId = strconv.Itoa(i)

		keeper.SetNocCertificatesBySubjectKeyId(ctx, items[i])
	}
	return items
}

func TestNocCertificatesBySubjectKeyIdGet(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNNocCertificatesBySubjectKeyId(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetNocCertificatesBySubjectKeyId(ctx,
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
		keeper.RemoveNocCertificatesBySubjectKeyId(ctx,
			item.SubjectKeyId,
		)
		_, found := keeper.GetNocCertificatesBySubjectKeyId(ctx,
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
		nullify.Fill(keeper.GetAllNocCertificatesBySubjectKeyId(ctx)),
	)
}
