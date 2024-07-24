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

func createNNocCertificatesByVidAndSkid(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.NocCertificatesByVidAndSkid {
	items := make([]types.NocCertificatesByVidAndSkid, n)
	for i := range items {
		items[i].Vid = int32(i)
		items[i].SubjectKeyId = strconv.Itoa(i)

		keeper.SetNocCertificatesByVidAndSkid(ctx, items[i])
	}

	return items
}

func TestNocCertificatesByVidAndSkidGet(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNNocCertificatesByVidAndSkid(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetNocCertificatesByVidAndSkid(ctx,
			item.Vid,
			item.SubjectKeyId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestNocCertificatesByVidAndSkidRemove(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNNocCertificatesByVidAndSkid(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNocCertificatesByVidAndSkid(ctx,
			item.Vid,
			item.SubjectKeyId,
		)
		_, found := keeper.GetNocCertificatesByVidAndSkid(ctx,
			item.Vid,
			item.SubjectKeyId,
		)
		require.False(t, found)
	}
}

func TestNocCertificatesByVidAndSkidGetAll(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNNocCertificatesByVidAndSkid(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllNocCertificatesByVidAndSkid(ctx)),
	)
}
