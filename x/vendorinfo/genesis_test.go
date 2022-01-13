package vendorinfo_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/types"
)

type DclauthKeeperMock struct {
	mock.Mock
}

func (m *DclauthKeeperMock) HasRole(
	ctx sdk.Context,
	addr sdk.AccAddress,
	roleToCheck dclauthtypes.AccountRole,
) bool {
	args := m.Called(ctx, addr, roleToCheck)
	return args.Bool(0)
}

func (m *DclauthKeeperMock) HasVendorID(
	ctx sdk.Context,
	addr sdk.AccAddress,
	vid int32,
) bool {
	args := m.Called(ctx, addr, vid)
	return args.Bool(0)
}

var _ types.DclauthKeeper = &DclauthKeeperMock{}

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		VendorInfoList: []types.VendorInfo{
			{
				VendorID: 0,
			},
			{
				VendorID: 1,
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}
	dclauthKeeper := &DclauthKeeperMock{}

	k, ctx := keepertest.VendorinfoKeeper(t, dclauthKeeper)
	vendorinfo.InitGenesis(ctx, *k, genesisState)
	got := vendorinfo.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.Len(t, got.VendorInfoList, len(genesisState.VendorInfoList))
	require.Subset(t, genesisState.VendorInfoList, got.VendorInfoList)
	// this line is used by starport scaffolding # genesis/test/assert
}
