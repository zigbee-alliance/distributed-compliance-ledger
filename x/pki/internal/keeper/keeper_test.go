//nolint:testpackage
package keeper

//nolint:goimports
import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestKeeper_CertificateGetSet(t *testing.T) {
	setup := Setup()

	// check if certificate present
	require.False(t, setup.PkiKeeper.IsCertificatePresent(
		setup.Ctx, testconstants.LeafSubject, testconstants.LeafSubjectKeyID))

	// no certificate before its created
	certificates := setup.PkiKeeper.GetCertificates(
		setup.Ctx, testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	require.Equal(t, 0, len(certificates.Items))

	// store certificate
	certificate := types.NewIntermediateCertificate(
		testconstants.LeafCertPem,
		testconstants.LeafSubject,
		testconstants.LeafSubjectKeyID,
		testconstants.LeafSerialNumber,
		testconstants.RootSubject,
		testconstants.RootSubjectKeyID,
		testconstants.Address1,
	)

	setup.PkiKeeper.SetCertificate(setup.Ctx, certificate)

	// check if certificate present
	require.True(t, setup.PkiKeeper.IsCertificatePresent(
		setup.Ctx, testconstants.LeafSubject, testconstants.LeafSubjectKeyID))

	// get certificate
	receivedCertificates := setup.PkiKeeper.GetCertificates(
		setup.Ctx, testconstants.LeafSubject, certificate.SubjectKeyID)
	require.Equal(t, 1, len(receivedCertificates.Items))

	receivedCertificate := receivedCertificates.Items[0]
	require.Equal(t, certificate.SubjectKeyID, receivedCertificate.SubjectKeyID)
	require.Equal(t, certificate.Subject, receivedCertificate.Subject)
	require.Equal(t, certificate.SerialNumber, receivedCertificate.SerialNumber)
	require.Equal(t, certificate.PemCert, receivedCertificate.PemCert)
	require.Equal(t, certificate.Owner, receivedCertificate.Owner)
	require.Equal(t, certificate.RootSubjectKeyID, receivedCertificate.RootSubjectKeyID)
}

func TestKeeper_PendingCertificateGetSet(t *testing.T) {
	setup := Setup()

	// check if pending certificate present
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
	require.Equal(t, certificate.Subject, receivedCertificate.Subject)
	require.Equal(t, certificate.SubjectKeyID, receivedCertificate.SubjectKeyID)
	require.Equal(t, certificate.SerialNumber, receivedCertificate.SerialNumber)
	require.Equal(t, certificate.PemCert, receivedCertificate.PemCert)
	require.Equal(t, certificate.Owner, receivedCertificate.Owner)
}

func TestKeeper_ChildCertificatesGetSet(t *testing.T) {
	setup := Setup()

	certificate := types.NewRootCertificate(
		testconstants.RootCertPem,
		testconstants.RootSubject,
		testconstants.RootSubjectKeyID,
		testconstants.RootSerialNumber,
		testconstants.Address1,
	)

	// check if child certificates list present
	require.False(t, setup.PkiKeeper.IsChildCertificatesPresent(setup.Ctx, certificate.Subject, certificate.SubjectKeyID))

	// no child certificates before its created
	childCertificates := setup.PkiKeeper.GetChildCertificates(setup.Ctx, certificate.Subject, certificate.SubjectKeyID)
	require.Equal(t, 0, len(childCertificates.ChildCertificates))

	// store child certificates
	certificateChain := types.NewChildCertificates(certificate.Subject, certificate.SubjectKeyID)
	certificateChain.AddChildCertificate(
		types.NewCertificateIdentifier(testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID))

	setup.PkiKeeper.SetChildCertificatesList(setup.Ctx, certificateChain)

	// check if child certificates present
	require.True(t,
		setup.PkiKeeper.IsChildCertificatesPresent(setup.Ctx, certificate.Subject, certificate.SubjectKeyID))

	// get child certificates
	receivedCertificatesChain :=
		setup.PkiKeeper.GetChildCertificates(setup.Ctx, certificate.Subject, certificate.SubjectKeyID)
	require.Equal(t, certificateChain.SubjectKeyID, receivedCertificatesChain.SubjectKeyID)
	require.Equal(t, certificateChain.ChildCertificates, receivedCertificatesChain.ChildCertificates)

	// store second child
	certificateChain.AddChildCertificate(
		types.NewCertificateIdentifier(testconstants.LeafSubject, testconstants.LeafSubjectKeyID))

	setup.PkiKeeper.SetChildCertificatesList(setup.Ctx, certificateChain)

	// get child certificates
	receivedCertificatesChain =
		setup.PkiKeeper.GetChildCertificates(setup.Ctx, certificate.Subject, certificate.SubjectKeyID)
	require.Equal(t, certificateChain.SubjectKeyID, receivedCertificatesChain.SubjectKeyID)
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

	setup.PkiKeeper.IterateCertificates(setup.Ctx, "", func(certificates types.Certificates) (stop bool) {
		expectedRecords = append(expectedRecords, certificates.Items...)
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
