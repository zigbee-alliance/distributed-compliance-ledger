package pki_test

import (
	"testing"

	"github.com/stretchr/testify/require"
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
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PkiKeeper(t)
	pki.InitGenesis(ctx, *k, genesisState)
	got := pki.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.ElementsMatch(t, genesisState.ApprovedCertificatesList, got.ApprovedCertificatesList)
	require.ElementsMatch(t, genesisState.ProposedCertificateList, got.ProposedCertificateList)
	require.ElementsMatch(t, genesisState.ChildCertificatesList, got.ChildCertificatesList)
	require.ElementsMatch(t, genesisState.ProposedCertificateRevocationList, got.ProposedCertificateRevocationList)
	require.ElementsMatch(t, genesisState.RevokedCertificatesList, got.RevokedCertificatesList)
	// this line is used by starport scaffolding # genesis/test/assert
}
