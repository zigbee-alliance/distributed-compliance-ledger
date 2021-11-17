package dclgenutil_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclgenutil"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclgenutil/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DclgenutilKeeper(t)
	dclgenutil.InitGenesis(ctx, *k, genesisState)
	got := dclgenutil.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
