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

func createAllCertificatesBySubjectKeyID(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.AllCertificatesBySubjectKeyId {
	items := make([]types.AllCertificatesBySubjectKeyId, n)
	for i := range items {
		items[i].SubjectKeyId = strconv.Itoa(i)

		keeper.SetAllCertificatesBySubjectKeyID(ctx, items[i])
	}

	return items
}

func TestAllCertificatesBySubjectKeyIdGet(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createAllCertificatesBySubjectKeyID(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetAllCertificatesBySubjectKeyID(ctx,
			item.SubjectKeyId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestAllCertificatesBySubjectKeyIdRemove(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createAllCertificatesBySubjectKeyID(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAllCertificatesBySubjectKeyID(ctx,
			"",
			item.SubjectKeyId,
		)
		_, found := keeper.GetAllCertificatesBySubjectKeyID(ctx,
			item.SubjectKeyId,
		)
		require.False(t, found)
	}
}

func TestAllCertificatesBySubjectKeyIdGetAll(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createAllCertificatesBySubjectKeyID(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAllCertificatesBySubjectKeyID(ctx)),
	)
}
