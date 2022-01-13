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

package vendorinfo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		VendorInfoList: []types.VendorInfo{
			{
				VendorID: 0,
			},
			{
				VendorID: 1,
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VendorinfoKeeper(t)
	vendorinfo.InitGenesis(ctx, *k, genesisState)
	got := vendorinfo.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.Len(t, got.VendorInfoList, len(genesisState.VendorInfoList))
	require.Subset(t, genesisState.VendorInfoList, got.VendorInfoList)
	// this line is used by starport scaffolding # genesis/test/assert
}
