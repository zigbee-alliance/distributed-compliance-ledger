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
	// this line is used by starport scaffolding # genesis/test/assert
}
*/
