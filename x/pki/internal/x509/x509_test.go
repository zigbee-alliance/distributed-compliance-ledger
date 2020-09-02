//nolint:testpackage
package x509

import (
	"testing"

	testconstants "git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"github.com/stretchr/testify/require"
)

func Test_DecodeCertificates(t *testing.T) {
	// decode leaf certificate
	certificate, err := DecodeX509Certificate(testconstants.LeafCertPem)
	require.Nil(t, err)
	require.False(t, certificate.IsSelfSigned())
	require.Equal(t, testconstants.IntermediateSubject, certificate.Issuer)
	require.Equal(t, testconstants.IntermediateSubjectKeyID, certificate.AuthorityKeyID)
	require.Equal(t, testconstants.LeafSerialNumber, certificate.SerialNumber)
	require.Equal(t, testconstants.LeafSubject, certificate.Subject)
	require.Equal(t, testconstants.LeafSubjectKeyID, certificate.SubjectKeyID)

	// decode intermediate certificate
	certificate, err = DecodeX509Certificate(testconstants.IntermediateCertPem)
	require.Nil(t, err)
	require.False(t, certificate.IsSelfSigned())
	require.Equal(t, testconstants.RootSubject, certificate.Issuer)
	require.Equal(t, testconstants.RootSubjectKeyID, certificate.AuthorityKeyID)
	require.Equal(t, testconstants.IntermediateSerialNumber, certificate.SerialNumber)
	require.Equal(t, testconstants.IntermediateSubject, certificate.Subject)
	require.Equal(t, testconstants.IntermediateSubjectKeyID, certificate.SubjectKeyID)

	// decode root certificate
	certificate, err = DecodeX509Certificate(testconstants.RootCertPem)
	require.Nil(t, err)
	require.True(t, certificate.IsSelfSigned())
	require.Equal(t, testconstants.RootSubject, certificate.Issuer)
	require.Equal(t, "", certificate.AuthorityKeyID)
	require.Equal(t, testconstants.RootSerialNumber, certificate.SerialNumber)
	require.Equal(t, testconstants.RootSubject, certificate.Subject)
	require.Equal(t, testconstants.RootSubjectKeyID, certificate.SubjectKeyID)
}

func Test_VerifyLeafCertificate(t *testing.T) {
	certificate, _ := DecodeX509Certificate(testconstants.LeafCertPem)
	parentCertificate, _ := DecodeX509Certificate(testconstants.IntermediateCertPem)
	err := certificate.Verify(parentCertificate)
	require.Nil(t, err)
}

func Test_VerifyRootCertificate(t *testing.T) {
	certificate, _ := DecodeX509Certificate(testconstants.RootCertPem)
	err := certificate.Verify(certificate)
	require.Nil(t, err)
}
