package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func TestGenesisState_Validate(t *testing.T) {
	valAddr1 := sample.ValAddress()
	valAddr2 := sample.ValAddress()

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
					{Owner: valAddr1},
					{Owner: valAddr2},
				},
				LastValidatorPowerList: []types.LastValidatorPower{
					{Owner: valAddr1},
					{Owner: valAddr2},
				},
				ProposedDisableValidatorList: []types.ProposedDisableValidator{
					{Address: "0"},
					{Address: "1"},
				},
				DisabledValidatorList: []types.DisabledValidator{
					{Address: "0"},
					{Address: "1"},
				},
				RejectedValidatorList: []types.RejectedDisableValidator{
					{Address: valAddr1},
					{Address: valAddr2},
				},
			},
			valid: true,
		},
		{
			desc: "invalid validator owner",
			genState: &types.GenesisState{
				ValidatorList: []types.Validator{
					{Owner: "invalid-address"},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated validator",
			genState: &types.GenesisState{
				ValidatorList: []types.Validator{
					{Owner: valAddr1},
					{Owner: valAddr1},
				},
			},
			valid: false,
		},
		{
			desc: "invalid lastValidatorPower owner",
			genState: &types.GenesisState{
				LastValidatorPowerList: []types.LastValidatorPower{
					{Owner: "invalid-address"},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated lastValidatorPower",
			genState: &types.GenesisState{
				LastValidatorPowerList: []types.LastValidatorPower{
					{Owner: valAddr1},
					{Owner: valAddr1},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated proposedDisableValidator",
			genState: &types.GenesisState{
				ProposedDisableValidatorList: []types.ProposedDisableValidator{
					{Address: "0"},
					{Address: "0"},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated disabledValidator",
			genState: &types.GenesisState{
				DisabledValidatorList: []types.DisabledValidator{
					{Address: "0"},
					{Address: "0"},
				},
			},
			valid: false,
		},
		{
			desc: "invalid rejectedValidator address",
			genState: &types.GenesisState{
				RejectedValidatorList: []types.RejectedDisableValidator{
					{Address: "invalid-address"},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated rejectedValidator",
			genState: &types.GenesisState{
				RejectedValidatorList: []types.RejectedDisableValidator{
					{Address: valAddr1},
					{Address: valAddr1},
				},
			},
			valid: false,
		},
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
