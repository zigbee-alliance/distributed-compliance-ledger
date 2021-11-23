package validator_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		ValidatorList: []types.Validator{
			{
				Owner: "0",
			},
			{
				Owner: "1",
			},
		},
		LastValidatorPowerList: []types.LastValidatorPower{
			{
				Owner: "0",
			},
			{
				Owner: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ValidatorKeeper(t)
	validator.InitGenesis(ctx, *k, genesisState)
	got := validator.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.Len(t, got.ValidatorList, len(genesisState.ValidatorList))
	require.Subset(t, genesisState.ValidatorList, got.ValidatorList)
	require.Len(t, got.LastValidatorPowerList, len(genesisState.LastValidatorPowerList))
	require.Subset(t, genesisState.LastValidatorPowerList, got.LastValidatorPowerList)
	// this line is used by starport scaffolding # genesis/test/assert
}
