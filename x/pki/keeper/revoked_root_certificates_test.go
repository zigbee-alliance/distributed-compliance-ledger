package keeper_test

/*
import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/nullify"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func createTestRevokedRootCertificates(keeper *keeper.Keeper, ctx sdk.Context) types.RevokedRootCertificates {
	item := types.RevokedRootCertificates{}
	keeper.SetRevokedRootCertificates(ctx, item)
	return item
}

func TestRevokedRootCertificatesGet(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t)
	item := createTestRevokedRootCertificates(keeper, ctx)
	rst, found := keeper.GetRevokedRootCertificates(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestRevokedRootCertificatesRemove(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t)
	createTestRevokedRootCertificates(keeper, ctx)
	keeper.RemoveRevokedRootCertificates(ctx)
	_, found := keeper.GetRevokedRootCertificates(ctx)
	require.False(t, found)
}
*/
