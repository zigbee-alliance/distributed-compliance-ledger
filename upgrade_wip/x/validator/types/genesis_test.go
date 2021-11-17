package types_test

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
						Address: "0",
					},
					{
						Address: "1",
					},
				},
				LastValidatorPowerList: []types.LastValidatorPower{
					{
						ConsensusAddress: "0",
					},
					{
						ConsensusAddress: "1",
					},
				},
				ValidatorSigningInfoList: []types.ValidatorSigningInfo{
					{
						Address: "0",
					},
					{
						Address: "1",
					},
				},
				ValidatorMissedBlockBitArrayList: []types.ValidatorMissedBlockBitArray{
					{
						Address: "0",
						Index:   0,
					},
					{
						Address: "1",
						Index:   1,
					},
				},
				ValidatorOwnerList: []types.ValidatorOwner{
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
						Address: "0",
					},
					{
						Address: "0",
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
						ConsensusAddress: "0",
					},
					{
						ConsensusAddress: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated validatorSigningInfo",
			genState: &types.GenesisState{
				ValidatorSigningInfoList: []types.ValidatorSigningInfo{
					{
						Address: "0",
					},
					{
						Address: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated validatorMissedBlockBitArray",
			genState: &types.GenesisState{
				ValidatorMissedBlockBitArrayList: []types.ValidatorMissedBlockBitArray{
					{
						Address: "0",
						Index:   0,
					},
					{
						Address: "0",
						Index:   0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated validatorOwner",
			genState: &types.GenesisState{
				ValidatorOwnerList: []types.ValidatorOwner{
					{
						Address: "0",
					},
					{
						Address: "0",
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
