package pki_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
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
			Certs: []*types.CertificateIdentifier{
				{
					Subject:      testconstants.IntermediateSubject,
					SubjectKeyId: testconstants.IntermediateSubjectKeyID,
				},
			},
		},
		RevokedRootCertificates: &types.RevokedRootCertificates{
			Certs: []*types.CertificateIdentifier{
				{
					Subject:      testconstants.IntermediateSubject,
					SubjectKeyId: testconstants.IntermediateSubjectKeyID,
				},
			},
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
				SubjectKeyId: "0",
			},
			{
				Subject:      "1",
				SubjectKeyId: "1",
			},
		},
		PkiRevocationDistributionPointList: []types.PkiRevocationDistributionPoint{
			{
				Vid:                0,
				Label:              "0",
				IssuerSubjectKeyID: "0",
			},
			{
				Vid:                1,
				Label:              "1",
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
		ApprovedCertificatesBySubjectKeyIdList: []types.ApprovedCertificatesBySubjectKeyId{
			{
				SubjectKeyId: "0",
			},
			{
				SubjectKeyId: "1",
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
		NocIcaCertificatesList: []types.NocIcaCertificates{
			{
				Vid: 0,
			},
			{
				Vid: 1,
			},
		},
		RevokedNocRootCertificatesList: []types.RevokedNocRootCertificates{
			{
				Subject:      "0",
				SubjectKeyId: "0",
			},
			{
				Subject:      "1",
				SubjectKeyId: "1",
			},
		},
		NocCertificatesByVidAndSkidList: []types.NocCertificatesByVidAndSkid{
			{
				Vid:          0,
				SubjectKeyId: "0",
			},
			{
				Vid:          1,
				SubjectKeyId: "1",
			},
		},
		NocCertificatesBySubjectKeyIdList: []types.NocCertificatesBySubjectKeyId{
			{
				SubjectKeyId: "0",
			},
			{
				SubjectKeyId: "1",
			},
		},
		NocCertificatesList: []types.NocCertificates{
			{
				Subject:      "0",
				SubjectKeyId: "0",
			},
			{
				Subject:      "1",
				SubjectKeyId: "1",
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
				Subject:      "0",
				SubjectKeyId: "0",
			},
			{
				Subject:      "1",
				SubjectKeyId: "1",
			},
		},
		RevokedNocIcaCertificatesList: []types.RevokedNocIcaCertificates{
			{
				Subject:      "0",
				SubjectKeyId: "0",
			},
			{
				Subject:      "1",
				SubjectKeyId: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PkiKeeper(t, nil)
	pki.InitGenesis(ctx, *k, genesisState)
	got := pki.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.ElementsMatch(t, genesisState.ApprovedCertificatesList, got.ApprovedCertificatesList)
	require.ElementsMatch(t, genesisState.ProposedCertificateList, got.ProposedCertificateList)
	require.ElementsMatch(t, genesisState.ChildCertificatesList, got.ChildCertificatesList)
	require.ElementsMatch(t, genesisState.ProposedCertificateRevocationList, got.ProposedCertificateRevocationList)
	require.ElementsMatch(t, genesisState.RevokedCertificatesList, got.RevokedCertificatesList)
	require.ElementsMatch(t, genesisState.UniqueCertificateList, got.UniqueCertificateList)
	require.Equal(t, genesisState.ApprovedRootCertificates, got.ApprovedRootCertificates)
	require.Equal(t, genesisState.RevokedRootCertificates, got.RevokedRootCertificates)
	require.ElementsMatch(t, genesisState.ApprovedCertificatesBySubjectList, got.ApprovedCertificatesBySubjectList)
	require.ElementsMatch(t, genesisState.RejectedCertificateList, got.RejectedCertificateList)
	require.ElementsMatch(t, genesisState.PkiRevocationDistributionPointList, got.PkiRevocationDistributionPointList)
	require.ElementsMatch(t, genesisState.PkiRevocationDistributionPointsByIssuerSubjectKeyIDList, got.PkiRevocationDistributionPointsByIssuerSubjectKeyIDList)
	require.ElementsMatch(t, genesisState.ApprovedCertificatesBySubjectKeyIdList, got.ApprovedCertificatesBySubjectKeyIdList)
	require.ElementsMatch(t, genesisState.NocRootCertificatesList, got.NocRootCertificatesList)
	require.ElementsMatch(t, genesisState.NocIcaCertificatesList, got.NocIcaCertificatesList)
	require.ElementsMatch(t, genesisState.RevokedNocRootCertificatesList, got.RevokedNocRootCertificatesList)
	require.ElementsMatch(t, genesisState.NocCertificatesByVidAndSkidList, got.NocCertificatesByVidAndSkidList)
	require.ElementsMatch(t, genesisState.NocCertificatesBySubjectKeyIdList, got.NocCertificatesBySubjectKeyIdList)
	require.ElementsMatch(t, genesisState.NocCertificatesList, got.NocCertificatesList)
	require.ElementsMatch(t, genesisState.NocCertificatesBySubjectList, got.NocCertificatesBySubjectList)
	require.ElementsMatch(t, genesisState.CertificatesList, got.CertificatesList)
	require.ElementsMatch(t, genesisState.RevokedNocIcaCertificatesList, got.RevokedNocIcaCertificatesList)
	// this line is used by starport scaffolding # genesis/test/assert
}
