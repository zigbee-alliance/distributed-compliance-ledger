package keeper

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestKeeper_CertificateGetSet(t *testing.T) {
	setup := Setup()

	// check if certificate present
	require.False(t, setup.PkiKeeper.IsCertificatePresent(setup.Ctx, test_constants.LeafSubject, test_constants.LeafSubjectKeyId))

	// no certificate before its created
	require.Panics(t, func() {
		setup.PkiKeeper.GetCertificate(setup.Ctx, test_constants.LeafSubject, test_constants.LeafSubjectKeyId)
	})

	// store certificate
	certificate := types.NewIntermediateCertificate(
		test_constants.LeafCertPem,
		test_constants.LeafSubject,
		test_constants.LeafSubjectKeyId,
		test_constants.LeafSerialNumber,
		test_constants.RootSubjectKeyId,
		test_constants.Address1,
	)

	setup.PkiKeeper.SetCertificate(setup.Ctx, certificate)

	// check if certificate present
	require.True(t, setup.PkiKeeper.IsCertificatePresent(setup.Ctx, test_constants.LeafSubject, test_constants.LeafSubjectKeyId))

	// get certificate
	receivedCertificate := setup.PkiKeeper.GetCertificate(setup.Ctx, test_constants.LeafSubject, certificate.SubjectKeyId)
	require.Equal(t, certificate.SubjectKeyId, receivedCertificate.SubjectKeyId)
	require.Equal(t, certificate.Subject, receivedCertificate.Subject)
	require.Equal(t, certificate.SerialNumber, receivedCertificate.SerialNumber)
	require.Equal(t, certificate.PemCert, receivedCertificate.PemCert)
	require.Equal(t, certificate.Owner, receivedCertificate.Owner)
	require.Equal(t, certificate.RootSubjectId, receivedCertificate.RootSubjectId)
}

func TestKeeper_PendingCertificateGetSet(t *testing.T) {
	setup := Setup()

	// check if pending certificate present
	require.False(t, setup.PkiKeeper.IsProposedCertificatePresent(setup.Ctx, test_constants.RootSubject, test_constants.RootSubjectKeyId))

	// no certificate before its created
	require.Panics(t, func() {
		setup.PkiKeeper.GetProposedCertificate(setup.Ctx, test_constants.RootSubject, test_constants.RootSubjectKeyId)
	})

	// store certificate
	certificate := types.NewProposedCertificate(
		test_constants.RootCertPem,
		test_constants.RootSubject,
		test_constants.RootSubjectKeyId,
		test_constants.RootSerialNumber,
		test_constants.Address1,
	)

	setup.PkiKeeper.SetProposedCertificate(setup.Ctx, certificate)

	// check if certificate present
	require.True(t, setup.PkiKeeper.IsProposedCertificatePresent(setup.Ctx, test_constants.RootSubject, test_constants.RootSubjectKeyId))

	// get certificate
	receivedCertificate := setup.PkiKeeper.GetProposedCertificate(setup.Ctx, test_constants.RootSubject, certificate.SubjectKeyId)
	require.Equal(t, certificate.Subject, receivedCertificate.Subject)
	require.Equal(t, certificate.SubjectKeyId, receivedCertificate.SubjectKeyId)
	require.Equal(t, certificate.SerialNumber, receivedCertificate.SerialNumber)
	require.Equal(t, certificate.PemCert, receivedCertificate.PemCert)
	require.Equal(t, certificate.Owner, receivedCertificate.Owner)
}

func TestKeeper_ChildCertificatesGetSet(t *testing.T) {
	setup := Setup()

	certificate := types.NewRootCertificate(
		test_constants.RootCertPem,
		test_constants.RootSubject,
		test_constants.RootSubjectKeyId,
		test_constants.RootSerialNumber,
		test_constants.Address1,
	)

	// check if child certificates list present
	require.False(t, setup.PkiKeeper.IsChildCertificatesPresent(setup.Ctx, certificate.Subject, certificate.SubjectKeyId))

	// no child certificates before its created
	childCertificates := setup.PkiKeeper.GetChildCertificates(setup.Ctx, certificate.Subject, certificate.SubjectKeyId)
	require.Equal(t, 0, len(childCertificates.ChildCertificates))

	// store child certificates
	certificateChain := types.NewChildCertificates(certificate.Subject, certificate.SubjectKeyId)
	certificateChain.AddChildCertificate(types.NewCertificateIdentifier(test_constants.IntermediateSubject, test_constants.IntermediateSubjectKeyId))

	setup.PkiKeeper.SetChildCertificatesList(setup.Ctx, certificateChain)

	// check if child certificates present
	require.True(t, setup.PkiKeeper.IsChildCertificatesPresent(setup.Ctx, certificate.Subject, certificate.SubjectKeyId))

	// get child certificates
	receivedCertificatesChain := setup.PkiKeeper.GetChildCertificates(setup.Ctx, certificate.Subject, certificate.SubjectKeyId)
	require.Equal(t, certificateChain.SubjectKeyId, receivedCertificatesChain.SubjectKeyId)
	require.Equal(t, certificateChain.ChildCertificates, receivedCertificatesChain.ChildCertificates)

	// store second child
	certificateChain.AddChildCertificate(types.NewCertificateIdentifier(test_constants.LeafSubject, test_constants.LeafSubjectKeyId))

	setup.PkiKeeper.SetChildCertificatesList(setup.Ctx, certificateChain)

	// get child certificates
	receivedCertificatesChain = setup.PkiKeeper.GetChildCertificates(setup.Ctx, certificate.Subject, certificate.SubjectKeyId)
	require.Equal(t, certificateChain.SubjectKeyId, receivedCertificatesChain.SubjectKeyId)
	require.Equal(t, certificateChain.ChildCertificates, receivedCertificatesChain.ChildCertificates)
}

func TestKeeper_CertificateIterator(t *testing.T) {
	setup := Setup()

	count := 9

	// add 3 leaf / 3 root / 3 pending certificates
	PopulateStoreWithMixedCertificates(setup, count)

	// get total count
	totalCertificates := setup.PkiKeeper.CountTotalCertificates(setup.Ctx)
	require.Equal(t, count/3*2, totalCertificates)

	// get iterator
	var expectedRecords []types.Certificate

	setup.PkiKeeper.IterateCertificates(setup.Ctx, "", func(certificate types.Certificate) (stop bool) {
		expectedRecords = append(expectedRecords, certificate)
		return false
	})
	require.Equal(t, count/3*2, len(expectedRecords))
}

func TestKeeper_PendingCertificateIterator(t *testing.T) {
	setup := Setup()

	count := 9

	// add 3 leaf / 3 root / 3 pending certificates
	PopulateStoreWithMixedCertificates(setup, count)

	// get total count
	totalCertificates := setup.PkiKeeper.CountTotalProposedCertificates(setup.Ctx)
	require.Equal(t, count/3, totalCertificates)

	// get iterator
	var expectedRecords []types.ProposedCertificate

	setup.PkiKeeper.IterateProposedCertificates(setup.Ctx, func(certificate types.ProposedCertificate) (stop bool) {
		expectedRecords = append(expectedRecords, certificate)
		return false
	})
	require.Equal(t, count/3, len(expectedRecords))
}
