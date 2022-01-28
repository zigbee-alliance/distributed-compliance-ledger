package compliancetest_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		TestingResultsList: []types.TestingResults{
			{
				Vid:             0,
				Pid:             0,
				SoftwareVersion: 0,
			},
			{
				Vid:             1,
				Pid:             1,
				SoftwareVersion: 1,
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CompliancetestKeeper(t, nil, nil)
	compliancetest.InitGenesis(ctx, *k, genesisState)
	got := compliancetest.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.ElementsMatch(t, genesisState.TestingResultsList, got.TestingResultsList)
	// this line is used by starport scaffolding # genesis/test/assert
}
