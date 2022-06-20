package keeper_test

/*
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

// Prevent strconv unused error
var _ = strconv.IntSize

func createNDeviceSoftwareCompliance(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.DeviceSoftwareCompliance {
	items := make([]types.DeviceSoftwareCompliance, n)
	for i := range items {
		items[i].CDCertificateId = strconv.Itoa(i)

		keeper.SetDeviceSoftwareCompliance(ctx, items[i])
	}
	return items
}

func TestDeviceSoftwareComplianceGet(t *testing.T) {
	keeper, ctx := keepertest.ComplianceKeeper(t)
	items := createNDeviceSoftwareCompliance(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetDeviceSoftwareCompliance(ctx,
			item.CDCertificateId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestDeviceSoftwareComplianceRemove(t *testing.T) {
	keeper, ctx := keepertest.ComplianceKeeper(t)
	items := createNDeviceSoftwareCompliance(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDeviceSoftwareCompliance(ctx,
			item.CdCertificateId,
		)
		_, found := keeper.GetDeviceSoftwareCompliance(ctx,
			item.CdCertificateId,
		)
		require.False(t, found)
	}
}

func TestDeviceSoftwareComplianceGetAll(t *testing.T) {
	keeper, ctx := keepertest.ComplianceKeeper(t)
	items := createNDeviceSoftwareCompliance(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDeviceSoftwareCompliance(ctx)),
	)
}
*/
