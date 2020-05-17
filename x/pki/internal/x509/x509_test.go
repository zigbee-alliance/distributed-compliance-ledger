//nolint:testpackage
package x509

// nolint:goimports
import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_DecodeCertificates(t *testing.T) {
	// decode leaf certificate
	certificate, err := DecodeX509Certificate(testconstants.LeafCertPem)
	require.Nil(t, err)
	require.Equal(t, testconstants.IntermediateSubject, certificate.Issuer)
	require.Equal(t, testconstants.IntermediateSubjectKeyID, certificate.AuthorityKeyID)
	require.Equal(t, testconstants.LeafSerialNumber, certificate.SerialNumber)
	require.Equal(t, testconstants.LeafSubject, certificate.Subject)
	require.Equal(t, testconstants.LeafSubjectKeyID, certificate.SubjectKeyID)
	require.False(t, certificate.IsRootCertificate())

	// decode intermediate certificate
	certificate, err = DecodeX509Certificate(testconstants.IntermediateCertPem)
	require.Nil(t, err)
	require.Equal(t, testconstants.RootSubject, certificate.Issuer)
	require.Equal(t, testconstants.RootSubjectKeyID, certificate.AuthorityKeyID)
	require.Equal(t, testconstants.IntermediateSerialNumber, certificate.SerialNumber)
	require.Equal(t, testconstants.IntermediateSubject, certificate.Subject)
	require.Equal(t, testconstants.IntermediateSubjectKeyID, certificate.SubjectKeyID)
	require.False(t, certificate.IsRootCertificate())

	// decode root certificate
	certificate, err = DecodeX509Certificate(testconstants.RootCertPem)
	require.Nil(t, err)
	require.True(t, certificate.IsRootCertificate())
	require.Equal(t, testconstants.RootSubject, certificate.Issuer)
	require.Equal(t, "", certificate.AuthorityKeyID)
	require.Equal(t, testconstants.RootSerialNumber, certificate.SerialNumber)
	require.Equal(t, testconstants.RootSubject, certificate.Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, certificate.SubjectKeyID)
	require.True(t, certificate.IsRootCertificate())
}

func Test_VerifyLeafCertificate(t *testing.T) {
	certificate, _ := DecodeX509Certificate(testconstants.LeafCertPem)
	parentCertificate, _ := DecodeX509Certificate(testconstants.IntermediateCertPem)
	err := certificate.VerifyX509Certificate(parentCertificate.Certificate)
	require.Nil(t, err)
}

func Test_VerifyRootCertificate(t *testing.T) {
	certificate, _ := DecodeX509Certificate(testconstants.RootCertPem)
	err := certificate.VerifyX509Certificate(certificate.Certificate)
	require.Nil(t, err)
}
