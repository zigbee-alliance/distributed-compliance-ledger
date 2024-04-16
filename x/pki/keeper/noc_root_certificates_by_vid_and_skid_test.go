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

func createNNocRootCertificatesByVidAndSkid(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.NocRootCertificatesByVidAndSkid {
	items := make([]types.NocRootCertificatesByVidAndSkid, n)
	for i := range items {
		items[i].Vid = int32(i)
		items[i].SubjectKeyId = strconv.Itoa(i)

		keeper.SetNocRootCertificatesByVidAndSkid(ctx, items[i])
	}
	return items
}

func TestNocRootCertificatesByVidAndSkidGet(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNNocRootCertificatesByVidAndSkid(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetNocRootCertificatesByVidAndSkid(ctx,
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
func TestNocRootCertificatesByVidAndSkidRemove(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNNocRootCertificatesByVidAndSkid(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNocRootCertificatesByVidAndSkid(ctx,
			item.Vid,
			item.SubjectKeyId,
		)
		_, found := keeper.GetNocRootCertificatesByVidAndSkid(ctx,
			item.Vid,
			item.SubjectKeyId,
		)
		require.False(t, found)
	}
}

func TestNocRootCertificatesByVidAndSkidGetAll(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNNocRootCertificatesByVidAndSkid(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllNocRootCertificatesByVidAndSkid(ctx)),
	)
}
