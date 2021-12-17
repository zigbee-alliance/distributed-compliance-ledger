package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
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
				ApprovedCertificatesList: []types.ApprovedCertificates{
					{
						Subject:      "0",
						SubjectKeyId: "0",
					},
					{
						Subject:      "1",
						SubjectKeyId: "1",
					},
				},
				ProposedCertificateList: []types.ProposedCertificate{
					{
						Subject:      "0",
						SubjectKeyId: "0",
					},
					{
						Subject:      "1",
						SubjectKeyId: "1",
					},
				},
				ChildCertificatesList: []types.ChildCertificates{
					{
						Issuer:         "0",
						AuthorityKeyId: "0",
					},
					{
						Issuer:         "1",
						AuthorityKeyId: "1",
					},
				},
				ProposedCertificateRevocationList: []types.ProposedCertificateRevocation{
					{
						Subject:      "0",
						SubjectKeyId: "0",
					},
					{
						Subject:      "1",
						SubjectKeyId: "1",
					},
				},
				RevokedCertificatesList: []types.RevokedCertificates{
					{
						Subject:      "0",
						SubjectKeyId: "0",
					},
					{
						Subject:      "1",
						SubjectKeyId: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated approvedCertificates",
			genState: &types.GenesisState{
				ApprovedCertificatesList: []types.ApprovedCertificates{
					{
						Subject:      "0",
						SubjectKeyId: "0",
					},
					{
						Subject:      "0",
						SubjectKeyId: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated proposedCertificate",
			genState: &types.GenesisState{
				ProposedCertificateList: []types.ProposedCertificate{
					{
						Subject:      "0",
						SubjectKeyId: "0",
					},
					{
						Subject:      "0",
						SubjectKeyId: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated childCertificates",
			genState: &types.GenesisState{
				ChildCertificatesList: []types.ChildCertificates{
					{
						Issuer:         "0",
						AuthorityKeyId: "0",
					},
					{
						Issuer:         "0",
						AuthorityKeyId: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated proposedCertificateRevocation",
			genState: &types.GenesisState{
				ProposedCertificateRevocationList: []types.ProposedCertificateRevocation{
					{
						Subject:      "0",
						SubjectKeyId: "0",
					},
					{
						Subject:      "0",
						SubjectKeyId: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated revokedCertificates",
			genState: &types.GenesisState{
				RevokedCertificatesList: []types.RevokedCertificates{
					{
						Subject:      "0",
						SubjectKeyId: "0",
					},
					{
						Subject:      "0",
						SubjectKeyId: "0",
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
