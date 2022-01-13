// Copyright 2022 DSR Corporation
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

package compliancetest_test

/*
import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
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
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CompliancetestKeeper(t)
	compliancetest.InitGenesis(ctx, *k, genesisState)
	got := compliancetest.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.ElementsMatch(t, genesisState.TestingResultsList, got.TestingResultsList)
	// this line is used by starport scaffolding # genesis/test/assert
}
*/
