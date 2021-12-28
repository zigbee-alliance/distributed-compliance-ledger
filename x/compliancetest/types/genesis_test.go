package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest/types"
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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated testingResults",
			genState: &types.GenesisState{
				TestingResultsList: []types.TestingResults{
					{
						Vid:             0,
						Pid:             0,
						SoftwareVersion: 0,
					},
					{
						Vid:             0,
						Pid:             0,
						SoftwareVersion: 0,
					},
				},
			},
			valid: false,
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
