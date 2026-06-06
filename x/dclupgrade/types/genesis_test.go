// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types_test

/*
import (
	"testing"

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
				ApprovedUpgradeList: []types.ApprovedUpgrade{
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
				RejectedUpgradeList: []types.RejectedUpgrade{
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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated proposedUpgrade",
			genState: &types.GenesisState{
				ProposedUpgradeList: []types.ProposedUpgrade{
					{
						Plan: types.Plan{
							Name: "0",
						},
					},
					{
						Plan: types.Plan{
							Name: "0",
						},
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated approvedUpgrade",
			genState: &types.GenesisState{
				ApprovedUpgradeList: []types.ApprovedUpgrade{
					{
						Plan: types.Plan{
							Name: "0",
						},
					},
					{
						Plan: types.Plan{
							Name: "0",
						},
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated rejectedUpgrade",
			genState: &types.GenesisState{
				RejectedUpgradeList: []types.RejectedUpgrade{
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
*/
