package types_test

/* FIXME issue 99

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
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
// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated validator",
			genState: &types.GenesisState{
				ValidatorList: []types.Validator{
					{
						Owner: "0",
					},
					{
						Owner: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated lastValidatorPower",
			genState: &types.GenesisState{
				LastValidatorPowerList: []types.LastValidatorPower{
					{
						Owner: "0",
					},
					{
						Owner: "0",
					},
				},
			},
			valid: false,
		},
		{
	desc:     "duplicated proposedDisableValidator",
	genState: &types.GenesisState{
		ProposedDisableValidatorList: []types.ProposedDisableValidator{
			{
				Address: "0",
},
			{
				Address: "0",
},
		},
	},
	valid:    false,
},
{
	desc:     "duplicated disabledValidator",
	genState: &types.GenesisState{
		DisabledValidatorList: []types.DisabledValidator{
			{
				Address: "0",
},
			{
				Address: "0",
},
		},
	},
	valid:    false,
},
// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
*/
