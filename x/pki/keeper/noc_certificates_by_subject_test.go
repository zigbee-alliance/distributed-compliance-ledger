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

func createNNocCertificatesBySubject(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.NocCertificatesBySubject {
	items := make([]types.NocCertificatesBySubject, n)
	for i := range items {
		items[i].Subject = strconv.Itoa(i)

		keeper.SetNocCertificatesBySubject(ctx, items[i])
	}
	return items
}

func TestNocCertificatesBySubjectGet(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNNocCertificatesBySubject(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetNocCertificatesBySubject(ctx,
			item.Subject,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestNocCertificatesBySubjectRemove(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNNocCertificatesBySubject(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNocCertificatesBySubject(ctx,
			item.Subject,
		)
		_, found := keeper.GetNocCertificatesBySubject(ctx,
			item.Subject,
		)
		require.False(t, found)
	}
}

func TestNocCertificatesBySubjectGetAll(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNNocCertificatesBySubject(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllNocCertificatesBySubject(ctx)),
	)
}
