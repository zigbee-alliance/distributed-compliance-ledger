package dclupgrade_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/nullify"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		ProposedUpgradeList: []types.ProposedUpgrade{
			{
				Plan: types.Plan{
					Name: "0",
				},
			},
			{
				Plan: types.Plan{
					Name: "1",
				},
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DclupgradeKeeper(t)
	dclupgrade.InitGenesis(ctx, *k, genesisState)
	got := dclupgrade.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ProposedUpgradeList, got.ProposedUpgradeList)
	// this line is used by starport scaffolding # genesis/test/assert
}
