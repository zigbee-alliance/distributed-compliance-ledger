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

func createNComplianceInfo(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ComplianceInfo {
	items := make([]types.ComplianceInfo, n)
	for i := range items {
		items[i].Vid = int32(i)
		items[i].Pid = int32(i)
		items[i].SoftwareVersion = uint32(i)
		items[i].CertificationType = strconv.Itoa(i)

		keeper.SetComplianceInfo(ctx, items[i])
	}
	return items
}

func TestComplianceInfoGet(t *testing.T) {
	keeper, ctx := keepertest.ComplianceKeeper(t, nil, nil)
	items := createNComplianceInfo(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetComplianceInfo(ctx,
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

func TestComplianceInfoRemove(t *testing.T) {
	keeper, ctx := keepertest.ComplianceKeeper(t, nil, nil)
	items := createNComplianceInfo(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveComplianceInfo(ctx,
			item.Vid,
			item.Pid,
			item.SoftwareVersion,
			item.CertificationType,
		)
		_, found := keeper.GetComplianceInfo(ctx,
			item.Vid,
			item.Pid,
			item.SoftwareVersion,
			item.CertificationType,
		)
		require.False(t, found)
	}
}

func TestComplianceInfoGetAll(t *testing.T) {
	keeper, ctx := keepertest.ComplianceKeeper(t, nil, nil)
	items := createNComplianceInfo(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllComplianceInfo(ctx)),
	)
}
