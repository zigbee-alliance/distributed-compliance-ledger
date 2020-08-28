//nolint:testpackage
package keeper

import (
	"testing"

	testconstants "git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/types"
	"github.com/stretchr/testify/require"
)

func TestKeeper_CertificateGetSet(t *testing.T) {
	setup := Setup()

	// check if certificate present
	require.False(t, setup.PkiKeeper.IsApprovedCertificatesPresent(
		setup.Ctx, testconstants.LeafSubject, testconstants.LeafSubjectKeyID))

	// no certificate before its created
	certificates := setup.PkiKeeper.GetApprovedCertificates(
		setup.Ctx, testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	require.Equal(t, 0, len(certificates.Items))

	// store certificate
	certificate := types.NewNonRootCertificate(
		testconstants.LeafCertPem,
		testconstants.LeafSubject,
		testconstants.LeafSubjectKeyID,
		testconstants.LeafSerialNumber,
		testconstants.LeafIssuer,
		testconstants.LeafAuthorityKeyID,
		testconstants.RootSubject,
		testconstants.RootSubjectKeyID,
		testconstants.Address1,
	)

	setup.PkiKeeper.AddApprovedCertificate(setup.Ctx, certificate)

	// check if certificate present
	require.True(t, setup.PkiKeeper.IsApprovedCertificatesPresent(
		setup.Ctx, testconstants.LeafSubject, testconstants.LeafSubjectKeyID))

	// get certificate
	receivedCertificates := setup.PkiKeeper.GetApprovedCertificates(
		setup.Ctx, testconstants.LeafSubject, certificate.SubjectKeyID)
	require.Equal(t, 1, len(receivedCertificates.Items))

	receivedCertificate := receivedCertificates.Items[0]
	require.Equal(t, certificate.PemCert, receivedCertificate.PemCert)
	require.Equal(t, certificate.Subject, receivedCertificate.Subject)
	require.Equal(t, certificate.SubjectKeyID, receivedCertificate.SubjectKeyID)
	require.Equal(t, certificate.SerialNumber, receivedCertificate.SerialNumber)
	require.Equal(t, certificate.Issuer, receivedCertificate.Issuer)
	require.Equal(t, certificate.AuthorityKeyID, receivedCertificate.AuthorityKeyID)
	require.Equal(t, certificate.RootSubject, receivedCertificate.RootSubject)
	require.Equal(t, certificate.RootSubjectKeyID, receivedCertificate.RootSubjectKeyID)
	require.Equal(t, certificate.IsRoot, receivedCertificate.IsRoot)
	require.Equal(t, certificate.Owner, receivedCertificate.Owner)
}

func TestKeeper_ProposedCertificateGetSet(t *testing.T) {
	setup := Setup()

	// check if proposed certificate present
	require.False(t, setup.PkiKeeper.IsProposedCertificatePresent(
		setup.Ctx, testconstants.RootSubject, testconstants.RootSubjectKeyID))

	// no certificate before its created
	require.Panics(t, func() {
		setup.PkiKeeper.GetProposedCertificate(setup.Ctx, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	})

	// store certificate
	certificate := types.NewProposedCertificate(
		testconstants.RootCertPem,
		testconstants.RootSubject,
		testconstants.RootSubjectKeyID,
		testconstants.RootSerialNumber,
		testconstants.Address1,
	)

	setup.PkiKeeper.SetProposedCertificate(setup.Ctx, certificate)

	// check if certificate present
	require.True(t, setup.PkiKeeper.IsProposedCertificatePresent(
		setup.Ctx, testconstants.RootSubject, testconstants.RootSubjectKeyID))

	// get certificate
	receivedCertificate := setup.PkiKeeper.GetProposedCertificate(
		setup.Ctx, testconstants.RootSubject, certificate.SubjectKeyID)
	require.Equal(t, certificate.PemCert, receivedCertificate.PemCert)
	require.Equal(t, certificate.Subject, receivedCertificate.Subject)
	require.Equal(t, certificate.SubjectKeyID, receivedCertificate.SubjectKeyID)
	require.Equal(t, certificate.SerialNumber, receivedCertificate.SerialNumber)
	require.Equal(t, certificate.Owner, receivedCertificate.Owner)
	// Amino marshals empty slices as nulls: https://github.com/tendermint/go-amino/issues/275
	// require.Equal(t, certificate.Approvals, receivedCertificate.Approvals)
}

func TestKeeper_ChildCertificatesGetSet(t *testing.T) {
	setup := Setup()

	// check if child certificates list present
	require.False(t, setup.PkiKeeper.IsChildCertificatesPresent(setup.Ctx,
		testconstants.RootSubject, testconstants.RootSubjectKeyID))

	// no child certificates before its created
	receivedChildCertificates :=
		setup.PkiKeeper.GetChildCertificates(setup.Ctx, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, 0, len(receivedChildCertificates.CertIdentifiers))

	// store child certificates
	childCertificates := types.NewChildCertificates(testconstants.RootSubject, testconstants.RootSubjectKeyID)
	childCertificates.CertIdentifiers = append(childCertificates.CertIdentifiers,
		types.NewCertificateIdentifier(testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID))

	setup.PkiKeeper.SetChildCertificates(setup.Ctx, childCertificates)

	// check if child certificates present
	require.True(t, setup.PkiKeeper.IsChildCertificatesPresent(setup.Ctx,
		testconstants.RootSubject, testconstants.RootSubjectKeyID))

	// get child certificates
	receivedChildCertificates =
		setup.PkiKeeper.GetChildCertificates(setup.Ctx, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, childCertificates.Issuer, receivedChildCertificates.Issuer)
	require.Equal(t, childCertificates.AuthorityKeyID, receivedChildCertificates.AuthorityKeyID)
	require.Equal(t, childCertificates.CertIdentifiers, receivedChildCertificates.CertIdentifiers)

	// store second child
	childCertificates.CertIdentifiers = append(childCertificates.CertIdentifiers,
		types.NewCertificateIdentifier(testconstants.LeafSubject, testconstants.LeafSubjectKeyID))

	setup.PkiKeeper.SetChildCertificates(setup.Ctx, childCertificates)

	// get child certificates
	receivedChildCertificates =
		setup.PkiKeeper.GetChildCertificates(setup.Ctx, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, childCertificates.Issuer, receivedChildCertificates.Issuer)
	require.Equal(t, childCertificates.AuthorityKeyID, receivedChildCertificates.AuthorityKeyID)
	require.Equal(t, childCertificates.CertIdentifiers, receivedChildCertificates.CertIdentifiers)
}

func TestKeeper_CertificateIterator(t *testing.T) {
	setup := Setup()

	// add 3 leaf / 3 root / 3 proposed certificates
	genCerts := setup.PopulateStoreWithMixedCertificates()

	// get iterator
	var iteratedCerts []types.Certificate

	setup.PkiKeeper.IterateApprovedCertificatesRecords(setup.Ctx, "", func(certificates types.Certificates) (stop bool) {
		iteratedCerts = append(iteratedCerts, certificates.Items...)
		return false
	})

	allApproved := CombineCertLists(genCerts.ApprovedRoots, genCerts.ApprovedNonRoots)

	require.Equal(t, len(allApproved), len(iteratedCerts))

	for i := 0; i < len(allApproved); i++ {
		require.Equal(t, allApproved[i].PemCert, iteratedCerts[i].PemCert)
		require.Equal(t, allApproved[i].Subject, iteratedCerts[i].Subject)
		require.Equal(t, allApproved[i].SubjectKeyID, iteratedCerts[i].SubjectKeyID)
		require.Equal(t, allApproved[i].SerialNumber, iteratedCerts[i].SerialNumber)
		require.Equal(t, allApproved[i].Issuer, iteratedCerts[i].Issuer)
		require.Equal(t, allApproved[i].AuthorityKeyID, iteratedCerts[i].AuthorityKeyID)
		require.Equal(t, allApproved[i].RootSubject, iteratedCerts[i].RootSubject)
		require.Equal(t, allApproved[i].RootSubjectKeyID, iteratedCerts[i].RootSubjectKeyID)
		require.Equal(t, allApproved[i].IsRoot, iteratedCerts[i].IsRoot)
		require.Equal(t, allApproved[i].Owner, iteratedCerts[i].Owner)
	}
}

func TestKeeper_ProposedCertificateIterator(t *testing.T) {
	setup := Setup()

	// add 3 leaf / 3 root / 3 proposed certificates
	genCerts := setup.PopulateStoreWithMixedCertificates()

	// get iterator
	var iteratedProposedCerts []types.ProposedCertificate

	setup.PkiKeeper.IterateProposedCertificates(setup.Ctx, func(certificate types.ProposedCertificate) (stop bool) {
		iteratedProposedCerts = append(iteratedProposedCerts, certificate)
		return false
	})

	require.Equal(t, len(genCerts.ProposedRoots), len(iteratedProposedCerts))

	for i := 0; i < len(genCerts.ProposedRoots); i++ {
		require.Equal(t, genCerts.ProposedRoots[i].PemCert, iteratedProposedCerts[i].PemCert)
		require.Equal(t, genCerts.ProposedRoots[i].Subject, iteratedProposedCerts[i].Subject)
		require.Equal(t, genCerts.ProposedRoots[i].SubjectKeyID, iteratedProposedCerts[i].SubjectKeyID)
		require.Equal(t, genCerts.ProposedRoots[i].SerialNumber, iteratedProposedCerts[i].SerialNumber)
		require.Equal(t, genCerts.ProposedRoots[i].Owner, iteratedProposedCerts[i].Owner)
		// Amino marshals empty slices as nulls: https://github.com/tendermint/go-amino/issues/275
		// require.Equal(t, genCerts.ProposedRoots[i].Approvals, iteratedProposedCerts[i].Approvals)
	}
}
