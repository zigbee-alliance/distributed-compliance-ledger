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

package model_test

/* TODO issue 99
import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		VendorProductsList: []types.VendorProducts{
			{
				Vid: 0,
			},
			{
				Vid: 1,
			},
		},
		ModelList: []types.Model{
			{
				Vid: 0,
				Pid: 0,
			},
			{
				Vid: 1,
				Pid: 1,
			},
		},
		ModelVersionList: []types.ModelVersion{
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
		ModelVersionsList: []types.ModelVersions{
			{
				Vid: 0,
				Pid: 0,
			},
			{
				Vid: 1,
				Pid: 1,
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ModelKeeper(t)
	model.InitGenesis(ctx, *k, genesisState)
	got := model.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.ElementsMatch(t, genesisState.VendorProductsList, got.VendorProductsList)
	require.ElementsMatch(t, genesisState.ModelList, got.ModelList)
	require.ElementsMatch(t, genesisState.ModelVersionList, got.ModelVersionList)
	require.ElementsMatch(t, genesisState.ModelVersionsList, got.ModelVersionsList)
	// this line is used by starport scaffolding # genesis/test/assert
}
*/
