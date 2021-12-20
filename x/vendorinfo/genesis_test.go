package vendorinfo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		VendorInfoTypeList: []types.VendorInfoType{
	{
		Index: "0",
},
	{
		Index: "1",
},
},
// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VendorinfoKeeper(t)
	vendorinfo.InitGenesis(ctx, *k, genesisState)
	got := vendorinfo.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.Len(t, got.VendorInfoTypeList, len(genesisState.VendorInfoTypeList))
require.Subset(t, genesisState.VendorInfoTypeList, got.VendorInfoTypeList)
// this line is used by starport scaffolding # genesis/test/assert
}
