package x509

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_DecodeCertificates(t *testing.T) {
	// decode leaf certificate
	certificate, err := DecodeX509Certificate(test_constants.LeafCertPem)
	require.Nil(t, err)
	require.Equal(t, test_constants.IntermediateSubject, certificate.Issuer)
	require.Equal(t, test_constants.IntermediateSubjectKeyId, certificate.AuthorityKeyId)
	require.Equal(t, test_constants.LeafSerialNumber, certificate.SerialNumber)
	require.Equal(t, test_constants.LeafSubject, certificate.Subject)
	require.Equal(t, test_constants.LeafSubjectKeyId, certificate.SubjectKeyId)
	require.False(t, certificate.IsRootCertificate())

	// decode intermediate certificate
	certificate, err = DecodeX509Certificate(test_constants.IntermediateCertPem)
	require.Nil(t, err)
	require.Equal(t, test_constants.RootSubject, certificate.Issuer)
	require.Equal(t, test_constants.RootSubjectKeyId, certificate.AuthorityKeyId)
	require.Equal(t, test_constants.IntermediateSerialNumber, certificate.SerialNumber)
	require.Equal(t, test_constants.IntermediateSubject, certificate.Subject)
	require.Equal(t, test_constants.IntermediateSubjectKeyId, certificate.SubjectKeyId)
	require.False(t, certificate.IsRootCertificate())

	// decode root certificate
	certificate, err = DecodeX509Certificate(test_constants.RootCertPem)
	require.Nil(t, err)
	require.True(t, certificate.IsRootCertificate())
	require.Equal(t, test_constants.RootSubject, certificate.Issuer)
	require.Equal(t, "", certificate.AuthorityKeyId)
	require.Equal(t, test_constants.RootSerialNumber, certificate.SerialNumber)
	require.Equal(t, test_constants.RootSubject, certificate.Subject)
	require.Equal(t, test_constants.RootSubjectKeyId, certificate.SubjectKeyId)
	require.True(t, certificate.IsRootCertificate())
}

func Test_VerifyLeafCertificate(t *testing.T) {
	certificate, _ := DecodeX509Certificate(test_constants.LeafCertPem)
	parentCertificate, _ := DecodeX509Certificate(test_constants.IntermediateCertPem)
	err := certificate.VerifyX509Certificate(parentCertificate.Certificate)
	require.Nil(t, err)
}

func Test_VerifyRootCertificate(t *testing.T) {
	certificate, _ := DecodeX509Certificate(test_constants.RootCertPem)
	err := certificate.VerifyX509Certificate(certificate.Certificate)
	require.Nil(t, err)
}
