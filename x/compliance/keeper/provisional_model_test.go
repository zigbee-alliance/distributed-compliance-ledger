package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/nullify"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
)

// Prevent strconv unused error.
var _ = strconv.IntSize

func createNProvisionalModel(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ProvisionalModel {
	items := make([]types.ProvisionalModel, n)
	for i := range items {
		items[i].Vid = int32(i)
		items[i].Pid = int32(i)
		items[i].SoftwareVersion = uint32(i)
		items[i].CertificationType = strconv.Itoa(i)
		items[i].Value = true

		keeper.SetProvisionalModel(ctx, items[i])
	}
	return items
}

func TestProvisionalModelGet(t *testing.T) {
	keeper, ctx := keepertest.ComplianceKeeper(t, nil, nil)
	items := createNProvisionalModel(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetProvisionalModel(ctx,
			item.Vid,
			item.Pid,
			item.SoftwareVersion,
			item.CertificationType,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestProvisionalModelRemove(t *testing.T) {
	keeper, ctx := keepertest.ComplianceKeeper(t, nil, nil)
	items := createNProvisionalModel(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveProvisionalModel(ctx,
			item.Vid,
			item.Pid,
			item.SoftwareVersion,
			item.CertificationType,
		)
		_, found := keeper.GetProvisionalModel(ctx,
			item.Vid,
			item.Pid,
			item.SoftwareVersion,
			item.CertificationType,
		)
		require.False(t, found)
	}
}

func TestProvisionalModelGetAll(t *testing.T) {
	keeper, ctx := keepertest.ComplianceKeeper(t, nil, nil)
	items := createNProvisionalModel(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllProvisionalModel(ctx)),
	)
}
