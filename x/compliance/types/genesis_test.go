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
