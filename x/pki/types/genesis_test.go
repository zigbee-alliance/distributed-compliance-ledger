package types_test

/*
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
						SubjectKeyID: "0",
					},
					{
						Subject:      "1",
						SubjectKeyID: "1",
					},
				},
				ProposedCertificateList: []types.ProposedCertificate{
					{
						Subject:      "0",
						SubjectKeyID: "0",
					},
					{
						Subject:      "1",
						SubjectKeyID: "1",
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
						SubjectKeyID: "0",
					},
					{
						Subject:      "1",
						SubjectKeyID: "1",
					},
				},
				RevokedCertificatesList: []types.RevokedCertificates{
					{
						Subject:      "0",
						SubjectKeyID: "0",
					},
					{
						Subject:      "1",
						SubjectKeyID: "1",
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
				RejectedCertificateList: []types.RejectedCertificate{
					{
						Subject:      "0",
						SubjectKeyID: "0",
					},
					{
						Subject:      "1",
						SubjectKeyID: "1",
					},
				},
				PKIRevocationDistributionPointList: []types.PKIRevocationDistributionPoint{
	{
		Vid: 0,
Label: "0",
IssuerSubjectKeyID: "0",
},
	{
		Vid: 1,
Label: "1",
IssuerSubjectKeyID: "1",
},
},
PkiRevocationDistributionPointsByIssuerSubjectKeyIDList: []types.PkiRevocationDistributionPointsByIssuerSubjectKeyID{
	{
		IssuerSubjectKeyID: "0",
},
	{
		IssuerSubjectKeyID: "1",
},
},
NocRootCertificatesList: []types.NocRootCertificates{
	{
		Vid: 0,
},
	{
		Vid: 1,
},
},
NocCertificatesList: []types.NocCertificates{
	{
		Vid: 0,
},
	{
		Vid: 1,
},
},
RevokedNocRootCertificatesList: []types.RevokedNocRootCertificates{
	{
		Subject: "0",
SubjectKeyID: "0",
},
	{
		Subject: "1",
SubjectKeyID: "1",
},
},
NocCertificatesByVidAndSkidList: []types.NocCertificatesByVidAndSkid{
	{
		Vid: 0,
SubjectKeyID: "0",
},
	{
		Vid: 1,
SubjectKeyID: "1",
},
},
NocCertificatesBySubjectKeyIDList: []types.NocCertificatesBySubjectKeyID{
	{
		Vid: 0,
SubjectKeyID: "0",
},
	{
		Vid: 1,
SubjectKeyID: "1",
},
},
NocCertificatesList: []types.NocCertificates{
	{
		Subject: "0",
SubjectKeyID: "0",
},
	{
		Subject: "1",
SubjectKeyID: "1",
},
},
NocCertificatesBySubjectList: []types.NocCertificatesBySubject{
	{
		Subject: "0",
},
	{
		Subject: "1",
},
},
CertificatesList: []types.AllCertificates{
	{
		Subject: "0",
SubjectKeyID: "0",
},
	{
		Subject: "1",
SubjectKeyID: "1",
},
},
RevokedNocIcaCertificatesList: []types.RevokedNocIcaCertificates{
	{
		Subject: "0",
SubjectKeyID: "0",
},
	{
		Subject: "1",
SubjectKeyID: "1",
},
},
AllCertificatesBySubjectList: []types.AllCertificatesBySubject{
	{
		Subject: "0",
},
	{
		Subject: "1",
},
},
AllCertificatesBySubjectKeyIdList: []types.AllCertificatesBySubjectKeyId{
	{
		SubjectKeyId: "0",
},
	{
		SubjectKeyId: "1",
},
},
AllCertificatesBySubjectKeyIdList: []types.AllCertificatesBySubjectKeyID{
	{
		SubjectKeyId: "0",
},
	{
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
						SubjectKeyID: "0",
					},
					{
						Subject:      "0",
						SubjectKeyID: "0",
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
						SubjectKeyID: "0",
					},
					{
						Subject:      "0",
						SubjectKeyID: "0",
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
						SubjectKeyID: "0",
					},
					{
						Subject:      "0",
						SubjectKeyID: "0",
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
						SubjectKeyID: "0",
					},
					{
						Subject:      "0",
						SubjectKeyID: "0",
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
		{
			desc: "duplicated rejectedCertificate",
			genState: &types.GenesisState{
				RejectedCertificateList: []types.RejectedCertificate{
					{
						Subject:      "0",
						SubjectKeyID: "0",
					},
					{
						Subject:      "0",
						SubjectKeyID: "0",
					},
				},
			},
			valid: false,
		},
		{
	desc:     "duplicated pKIRevocationDistributionPoint",
	genState: &types.GenesisState{
		PKIRevocationDistributionPointList: []types.PKIRevocationDistributionPoint{
			{
				Vid: 0,
Label: "0",
IssuerSubjectKeyID: "0",
},
			{
				Vid: 0,
Label: "0",
IssuerSubjectKeyID: "0",
},
		},
	},
	valid:    false,
},
{
	desc:     "duplicated pkiRevocationDistributionPointsByIssuerSubjectKeyID",
	genState: &types.GenesisState{
		PkiRevocationDistributionPointsByIssuerSubjectKeyIDList: []types.PkiRevocationDistributionPointsByIssuerSubjectKeyID{
			{
				IssuerSubjectKeyID: "0",
},
			{
				IssuerSubjectKeyID: "0",
},
		},
	},
	valid:    false,
},
{
	desc:     "duplicated nocRootCertificates",
	genState: &types.GenesisState{
		NocRootCertificatesList: []types.NocRootCertificates{
			{
				Vid: 0,
},
			{
				Vid: 0,
},
		},
	},
	valid:    false,
},
{
	desc:     "duplicated nocCertificates",
	genState: &types.GenesisState{
		NocCertificatesList: []types.NocCertificates{
			{
				Vid: 0,
},
			{
				Vid: 0,
},
		},
	},
	valid:    false,
},
{
	desc:     "duplicated revokedNocRootCertificates",
	genState: &types.GenesisState{
		RevokedNocRootCertificatesList: []types.RevokedNocRootCertificates{
			{
				Subject: "0",
SubjectKeyID: "0",
},
			{
				Subject: "0",
SubjectKeyID: "0",
},
		},
	},
	valid:    false,
},
{
	desc:     "duplicated nocCertificatesByVidAndSkid",
	genState: &types.GenesisState{
		NocCertificatesByVidAndSkidList: []types.NocCertificatesByVidAndSkid{
			{
				Vid: 0,
SubjectKeyID: "0",
},
			{
				Vid: 0,
SubjectKeyID: "0",
},
		},
	},
	valid:    false,
},
{
	desc:     "duplicated nocCertificatesBySubjectKeyId",
	genState: &types.GenesisState{
		NocCertificatesBySubjectKeyIDList: []types.NocCertificatesBySubjectKeyID{
			{
				Vid: 0,
SubjectKeyID: "0",
},
			{
				Vid: 0,
SubjectKeyID: "0",
},
		},
	},
	valid:    false,
},
{
	desc:     "duplicated nocCertificates",
	genState: &types.GenesisState{
		NocCertificatesList: []types.NocCertificates{
			{
				Subject: "0",
SubjectKeyID: "0",
},
			{
				Subject: "0",
SubjectKeyID: "0",
},
		},
	},
	valid:    false,
},
{
	desc:     "duplicated nocCertificatesBySubject",
	genState: &types.GenesisState{
		NocCertificatesBySubjectList: []types.NocCertificatesBySubject{
			{
				Subject: "0",
},
			{
				Subject: "0",
},
		},
	},
	valid:    false,
},
{
	desc:     "duplicated certificates",
	genState: &types.GenesisState{
		CertificatesList: []types.AllCertificates{
			{
				Subject: "0",
SubjectKeyID: "0",
},
			{
				Subject: "0",
SubjectKeyID: "0",
},
		},
	},
	valid:    false,
},
{
	desc:     "duplicated revokedNocIcaCertificates",
	genState: &types.GenesisState{
		RevokedNocIcaCertificatesList: []types.RevokedNocIcaCertificates{
			{
				Subject: "0",
SubjectKeyID: "0",
},
			{
				Subject: "0",
SubjectKeyID: "0",
},
		},
	},
	valid:    false,
},
{
	desc:     "duplicated allCertificatesBySubject",
	genState: &types.GenesisState{
		AllCertificatesBySubjectList: []types.AllCertificatesBySubject{
			{
				Subject: "0",
},
			{
				Subject: "0",
},
		},
	},
	valid:    false,
},
{
	desc:     "duplicated allCertificatesBySubjectKeyId",
	genState: &types.GenesisState{
		AllCertificatesBySubjectKeyIdList: []types.AllCertificatesBySubjectKeyId{
			{
				SubjectKeyId: "0",
},
			{
				SubjectKeyId: "0",
},
		},
	},
	valid:    false,
},
{
	desc:     "duplicated allCertificatesBySubjectKeyId",
	genState: &types.GenesisState{
		AllCertificatesBySubjectKeyIdList: []types.AllCertificatesBySubjectKeyId{
			{
				SubjectKeyId: "0",
},
			{
				SubjectKeyId: "0",
},
		},
	},
	valid:    false,
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
