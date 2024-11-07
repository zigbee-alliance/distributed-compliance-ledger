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

func createNAllCertificatesBySubject(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.AllCertificatesBySubject {
	items := make([]types.AllCertificatesBySubject, n)
	for i := range items {
		items[i].Subject = strconv.Itoa(i)

		keeper.SetAllCertificatesBySubject(ctx, items[i])
	}
	return items
}

func TestAllCertificatesBySubjectGet(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNAllCertificatesBySubject(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetAllCertificatesBySubject(ctx,
			item.Subject,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestAllCertificatesBySubjectRemove(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNAllCertificatesBySubject(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAllCertificatesBySubject(ctx,
			item.Subject,
		)
		_, found := keeper.GetAllCertificatesBySubject(ctx,
			item.Subject,
		)
		require.False(t, found)
	}
}

func TestAllCertificatesBySubjectGetAll(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNAllCertificatesBySubject(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAllCertificatesBySubject(ctx)),
	)
}
