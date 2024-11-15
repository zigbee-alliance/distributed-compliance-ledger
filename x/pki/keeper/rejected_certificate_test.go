package keeper_test

/*
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

func createNRejectedCertificate(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.RejectedCertificate {
	items := make([]types.RejectedCertificate, n)
	for i := range items {
		items[i].Subject = strconv.Itoa(i)
		items[i].SubjectKeyID = strconv.Itoa(i)

		keeper.SetRejectedCertificate(ctx, items[i])
	}
	return items
}

func TestRejectedCertificateGet(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t)
	items := createNRejectedCertificate(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetRejectedCertificate(ctx,
			item.Subject,
			item.SubjectKeyID,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestRejectedCertificateRemove(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t)
	items := createNRejectedCertificate(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRejectedCertificate(ctx,
			item.Subject,
			item.SubjectKeyID,
		)
		_, found := keeper.GetRejectedCertificate(ctx,
			item.Subject,
			item.SubjectKeyID,
		)
		require.False(t, found)
	}
}

func TestRejectedCertificateGetAll(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t)
	items := createNRejectedCertificate(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllRejectedCertificate(ctx)),
	)
}
*/
