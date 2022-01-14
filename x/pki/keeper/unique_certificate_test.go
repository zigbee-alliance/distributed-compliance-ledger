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

// Prevent strconv unused error.
var _ = strconv.IntSize

func createNUniqueCertificate(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.UniqueCertificate {
	items := make([]types.UniqueCertificate, n)
	for i := range items {
		items[i].Issuer = strconv.Itoa(i)
		items[i].SerialNumber = strconv.Itoa(i)

		keeper.SetUniqueCertificate(ctx, items[i])
	}
	return items
}

func TestUniqueCertificateGet(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t)
	items := createNUniqueCertificate(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetUniqueCertificate(ctx,
			item.Issuer,
			item.SerialNumber,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestUniqueCertificateRemove(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t)
	items := createNUniqueCertificate(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveUniqueCertificate(ctx,
			item.Issuer,
			item.SerialNumber,
		)
		_, found := keeper.GetUniqueCertificate(ctx,
			item.Issuer,
			item.SerialNumber,
		)
		require.False(t, found)
	}
}

func TestUniqueCertificateGetAll(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t)
	items := createNUniqueCertificate(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllUniqueCertificate(ctx)),
	)
}
*/
