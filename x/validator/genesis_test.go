package validator_test

/* TODO issue 99
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
		ProposedDisableValidatorList: []types.ProposedDisableValidator{
		{
			Address: "0",
},
		{
			Address: "1",
},
	},
	DisabledValidatorList: []types.DisabledValidator{
		{
			Address: "0",
},
		{
			Address: "1",
},
	},
	// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ValidatorKeeper(t, nil)
	validator.InitGenesis(ctx, *k, genesisState)
	got := validator.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.Len(t, got.ValidatorList, len(genesisState.ValidatorList))
	require.Subset(t, genesisState.ValidatorList, got.ValidatorList)
	require.Len(t, got.LastValidatorPowerList, len(genesisState.LastValidatorPowerList))
	require.Subset(t, genesisState.LastValidatorPowerList, got.LastValidatorPowerList)
	require.ElementsMatch(t, genesisState.ProposedDisableValidatorList, got.ProposedDisableValidatorList)
require.ElementsMatch(t, genesisState.DisabledValidatorList, got.DisabledValidatorList)
// this line is used by starport scaffolding # genesis/test/assert
}
*/
