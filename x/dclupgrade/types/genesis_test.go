package types_test

import (
	"testing"

	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/types"
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
				ProposedUpgradeList: []types.ProposedUpgrade{
					{Plan: upgradetypes.Plan{Name: "0"}},
					{Plan: upgradetypes.Plan{Name: "1"}},
				},
				ApprovedUpgradeList: []types.ApprovedUpgrade{
					{Plan: upgradetypes.Plan{Name: "0"}},
					{Plan: upgradetypes.Plan{Name: "1"}},
				},
				RejectedUpgradeList: []types.RejectedUpgrade{
					{Plan: upgradetypes.Plan{Name: "0"}},
					{Plan: upgradetypes.Plan{Name: "1"}},
				},
			},
			valid: true,
		},
		{
			desc: "duplicated proposedUpgrade",
			genState: &types.GenesisState{
				ProposedUpgradeList: []types.ProposedUpgrade{
					{Plan: upgradetypes.Plan{Name: "0"}},
					{Plan: upgradetypes.Plan{Name: "0"}},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated approvedUpgrade",
			genState: &types.GenesisState{
				ApprovedUpgradeList: []types.ApprovedUpgrade{
					{Plan: upgradetypes.Plan{Name: "0"}},
					{Plan: upgradetypes.Plan{Name: "0"}},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated rejectedUpgrade",
			genState: &types.GenesisState{
				RejectedUpgradeList: []types.RejectedUpgrade{
					{Plan: upgradetypes.Plan{Name: "0"}},
					{Plan: upgradetypes.Plan{Name: "0"}},
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
