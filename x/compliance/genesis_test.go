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

package compliance_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		ComplianceInfoList: []types.ComplianceInfo{
			{
				Vid:               0,
				Pid:               0,
				SoftwareVersion:   0,
				CertificationType: "0",
			},
			{
				Vid:               1,
				Pid:               1,
				SoftwareVersion:   1,
				CertificationType: "1",
			},
		},
		CertifiedModelList: []types.CertifiedModel{
			{
				Vid:               0,
				Pid:               0,
				SoftwareVersion:   0,
				CertificationType: "0",
			},
			{
				Vid:               1,
				Pid:               1,
				SoftwareVersion:   1,
				CertificationType: "1",
			},
		},
		RevokedModelList: []types.RevokedModel{
			{
				Vid:               0,
				Pid:               0,
				SoftwareVersion:   0,
				CertificationType: "0",
			},
			{
				Vid:               1,
				Pid:               1,
				SoftwareVersion:   1,
				CertificationType: "1",
			},
		},
		ProvisionalModelList: []types.ProvisionalModel{
			{
				Vid:               0,
				Pid:               0,
				SoftwareVersion:   0,
				CertificationType: "0",
			},
			{
				Vid:               1,
				Pid:               1,
				SoftwareVersion:   1,
				CertificationType: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ComplianceKeeper(t)
	compliance.InitGenesis(ctx, *k, genesisState)
	got := compliance.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.ElementsMatch(t, genesisState.ComplianceInfoList, got.ComplianceInfoList)
	require.ElementsMatch(t, genesisState.CertifiedModelList, got.CertifiedModelList)
	require.ElementsMatch(t, genesisState.RevokedModelList, got.RevokedModelList)
	require.ElementsMatch(t, genesisState.ProvisionalModelList, got.ProvisionalModelList)
	// this line is used by starport scaffolding # genesis/test/assert
}
