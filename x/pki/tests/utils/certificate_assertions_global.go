package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func EnsureUniqueCertificateCertificateExist(
	t *testing.T,
	setup *TestSetup,
	issuer string,
	serialNumber string,
) {
	t.Helper()

	// UniqueCertificate: check that unique certificate key registered
	require.True(t, setup.Keeper.IsUniqueCertificatePresent(
		setup.Ctx, issuer, serialNumber))
}

func EnsureUniqueCertificateCertificateNotExist(
	t *testing.T,
	setup *TestSetup,
	issuer string,
	serialNumber string,
	skipCheck bool,
) {
	t.Helper()

	if !skipCheck {
		// UniqueCertificate: check that unique certificate key registered
		found := setup.Keeper.IsUniqueCertificatePresent(setup.Ctx, issuer, serialNumber)
		require.False(t, found)
	}
}

func EnsureGlobalCertificateExist(
	t *testing.T,
	setup *TestSetup,
	subject string,
	subjectKeyID string,
	serialNumber string,
	skipCheckForSubject bool, // TODO: FIX constants and eliminate this condition
) {
	t.Helper()

	// AllCertificate: Subject and SKID
	allCertificate, err := QueryAllCertificates(setup, subject, subjectKeyID)
	require.NoError(t, err)
	require.Equal(t, subject, allCertificate.Subject)
	require.Equal(t, subjectKeyID, allCertificate.SubjectKeyId)
	require.Len(t, allCertificate.Certs, 1)
	require.Equal(t, serialNumber, allCertificate.Certs[0].SerialNumber)

	// AllCertificate: SKID
	certificateBySubjectKeyID, _ := QueryAllCertificatesBySubjectKeyID(setup, subjectKeyID)
	require.Len(t, certificateBySubjectKeyID, 1)
	require.Len(t, certificateBySubjectKeyID[0].Certs, 1)
	require.Equal(t, serialNumber, certificateBySubjectKeyID[0].Certs[0].SerialNumber)

	if !skipCheckForSubject {
		// AllCertificate: Subject
		certificatesBySubject, _ := QueryAllCertificatesBySubject(setup, subject)
		require.Len(t, certificatesBySubject.SubjectKeyIds, 1)
		require.Equal(t, subjectKeyID, certificatesBySubject.SubjectKeyIds[0])
	}
}

func EnsureGlobalCertificateNotExist(
	t *testing.T,
	setup *TestSetup,
	subject string,
	subjectKeyID string,
	skipCheckForSubject bool, // TODO: FIX constants and eliminate this condition
	skipCheckForSkid bool,
) {
	t.Helper()

	// All certificates indexes checks

	// AllCertificate: Subject and SKID
	_, err := QueryAllCertificates(setup, subject, subjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))

	if !skipCheckForSkid {
		// AllCertificate: SKID
		certificatesBySubjectKeyID, _ := QueryAllCertificatesBySubjectKeyID(setup, subjectKeyID)
		require.Empty(t, certificatesBySubjectKeyID)
	}

	if !skipCheckForSubject {
		// AllCertificate: Subject
		_, err = QueryAllCertificatesBySubject(setup, subject)
		require.Equal(t, codes.NotFound, status.Code(err))
	}
}

func EnsureChildCertificateExist(
	t *testing.T,
	setup *TestSetup,
	subject string,
	subjectKeyID string,
	issuer string,
	authorityKeyID string,
) {
	t.Helper()

	issuerChildren, _ := QueryChildCertificates(setup, subject, subjectKeyID)
	require.Len(t, issuerChildren.CertIds, 1)

	certID := types.CertificateIdentifier{
		Subject:      issuer,
		SubjectKeyId: authorityKeyID,
	}
	require.Equal(t, &certID, issuerChildren.CertIds[0])
}
