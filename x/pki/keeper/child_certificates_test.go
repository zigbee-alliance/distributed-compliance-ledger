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

func createNChildCertificates(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ChildCertificates {
	items := make([]types.ChildCertificates, n)
	for i := range items {
		items[i].Issuer = strconv.Itoa(i)
		items[i].AuthorityKeyId = strconv.Itoa(i)

		keeper.SetChildCertificates(ctx, items[i])
	}
	return items
}

func TestChildCertificatesGet(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t)
	items := createNChildCertificates(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetChildCertificates(ctx,
			item.Issuer,
			item.AuthorityKeyId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestChildCertificatesRemove(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t)
	items := createNChildCertificates(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveChildCertificates(ctx,
			item.Issuer,
			item.AuthorityKeyId,
		)
		_, found := keeper.GetChildCertificates(ctx,
			item.Issuer,
			item.AuthorityKeyId,
		)
		require.False(t, found)
	}
}

func TestChildCertificatesGetAll(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t)
	items := createNChildCertificates(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllChildCertificates(ctx)),
	)
}
