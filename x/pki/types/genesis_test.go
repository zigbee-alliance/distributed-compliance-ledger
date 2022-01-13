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

/* TODO issue 99
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
				UniqueCertificateList: []types.UniqueCertificate{
					{
						Issuer:       "0",
						SerialNumber: "0",
					},
					{
						Issuer:       "1",
						SerialNumber: "1",
					},
				},
				ApprovedRootCertificates: &types.ApprovedRootCertificates{
					Certs: []*types.CertificateIdentifier{},
				},
				RevokedRootCertificates: &types.RevokedRootCertificates{
					Certs: []*types.CertificateIdentifier{},
				},
				ApprovedCertificatesBySubjectList: []types.ApprovedCertificatesBySubject{
					{
						Subject: "0",
					},
					{
						Subject: "1",
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
		{
			desc: "duplicated uniqueCertificate",
			genState: &types.GenesisState{
				UniqueCertificateList: []types.UniqueCertificate{
					{
						Issuer:       "0",
						SerialNumber: "0",
					},
					{
						Issuer:       "0",
						SerialNumber: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated approvedCertificatesBySubject",
			genState: &types.GenesisState{
				ApprovedCertificatesBySubjectList: []types.ApprovedCertificatesBySubject{
					{
						Subject: "0",
					},
					{
						Subject: "0",
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
