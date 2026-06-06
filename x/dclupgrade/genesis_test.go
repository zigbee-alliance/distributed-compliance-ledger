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
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DclupgradeKeeper(t, nil, nil)
	dclupgrade.InitGenesis(ctx, *k, genesisState)
	got := dclupgrade.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ProposedUpgradeList, got.ProposedUpgradeList)
	require.ElementsMatch(t, genesisState.ApprovedUpgradeList, got.ApprovedUpgradeList)
	require.ElementsMatch(t, genesisState.RejectedUpgradeList, got.RejectedUpgradeList)
	// this line is used by starport scaffolding # genesis/test/assert
}
