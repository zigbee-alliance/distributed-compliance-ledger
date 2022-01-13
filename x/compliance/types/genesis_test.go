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

package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated complianceInfo",
			genState: &types.GenesisState{
				ComplianceInfoList: []types.ComplianceInfo{
					{
						Vid:               0,
						Pid:               0,
						SoftwareVersion:   0,
						CertificationType: "0",
					},
					{
						Vid:               0,
						Pid:               0,
						SoftwareVersion:   0,
						CertificationType: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated certifiedModel",
			genState: &types.GenesisState{
				CertifiedModelList: []types.CertifiedModel{
					{
						Vid:               0,
						Pid:               0,
						SoftwareVersion:   0,
						CertificationType: "0",
					},
					{
						Vid:               0,
						Pid:               0,
						SoftwareVersion:   0,
						CertificationType: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated revokedModel",
			genState: &types.GenesisState{
				RevokedModelList: []types.RevokedModel{
					{
						Vid:               0,
						Pid:               0,
						SoftwareVersion:   0,
						CertificationType: "0",
					},
					{
						Vid:               0,
						Pid:               0,
						SoftwareVersion:   0,
						CertificationType: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated provisionalModel",
			genState: &types.GenesisState{
				ProvisionalModelList: []types.ProvisionalModel{
					{
						Vid:               0,
						Pid:               0,
						SoftwareVersion:   0,
						CertificationType: "0",
					},
					{
						Vid:               0,
						Pid:               0,
						SoftwareVersion:   0,
						CertificationType: "0",
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
