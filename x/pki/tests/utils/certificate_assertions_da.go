package utils

import (
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func EnsureCertificatePresentInDaCertificateIndexes(
	t *testing.T,
	setup *TestSetup,
	subject string,
	subjectKeyID string,
	serialNumber string,
	isRoot bool,
	skipCheckForSubject bool,
) {
	t.Helper()

	// DaCertificates: Subject and SKID
	approvedCertificates, _ := QueryApprovedCertificates(setup, subject, subjectKeyID)
	require.Equal(t, subject, approvedCertificates.Subject)
	require.Equal(t, subjectKeyID, approvedCertificates.SubjectKeyId)
	require.Len(t, approvedCertificates.Certs, 1)
	require.Equal(t, serialNumber, approvedCertificates.Certs[0].SerialNumber)
	require.Equal(t, isRoot, approvedCertificates.Certs[0].IsRoot)

	if isRoot {
		// DaCertificates: Root Subject and SKID
		approvedRootCertificate, _ := QueryApprovedRootCertificates(setup, subject, subjectKeyID)
		require.Equal(t, subject, approvedRootCertificate.Subject)
		require.Equal(t, subjectKeyID, approvedRootCertificate.SubjectKeyId)
	}

	// DaCertificates: SKID
	certificateBySubjectKeyID, _ := QueryApprovedCertificatesBySubjectKeyID(setup, subjectKeyID)
	require.Len(t, certificateBySubjectKeyID, 1)
	require.Len(t, certificateBySubjectKeyID[0].Certs, 1)
	require.Equal(t, serialNumber, certificateBySubjectKeyID[0].Certs[0].SerialNumber)
	require.Equal(t, isRoot, certificateBySubjectKeyID[0].Certs[0].IsRoot)

	if !skipCheckForSubject {
		// DACertificates: Subject
		certificatesBySubject, _ := QueryApprovedCertificatesBySubject(setup, subject)
		require.Len(t, certificatesBySubject.SubjectKeyIds, 1)
		require.Equal(t, subjectKeyID, certificatesBySubject.SubjectKeyIds[0])
	}
}

func EnsureCertificateNotPresentInDaCertificateIndexes(
	t *testing.T,
	setup *TestSetup,
	subject string,
	subjectKeyID string,
	isRoot bool,
	skipCheckBySubject bool, // TODO: FIX constants and eliminate this condition
	skipCheckBySkid bool,
) {
	t.Helper()

	// DA certificates indexes checks

	// DaCertificates: Subject and SKID
	_, err := QueryApprovedCertificates(setup, subject, subjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))

	if isRoot {
		// DaCertificates: Root Subject and SKID
		_, err := QueryApprovedRootCertificates(setup, subject, subjectKeyID)
		require.Equal(t, codes.NotFound, status.Code(err))
	}

	if !skipCheckBySkid {
		// DaCertificates: SubjectKeyID
		certificatesBySubjectKeyID, _ := QueryApprovedCertificatesBySubjectKeyID(setup, subjectKeyID)
		require.Empty(t, certificatesBySubjectKeyID)
	}

	if !skipCheckBySubject {
		// NocCertificates: Subject
		_, err = QueryApprovedCertificatesBySubject(setup, subject)
		require.Equal(t, codes.NotFound, status.Code(err))
	}
}

func EnsureDaRootCertificateExist(
	t *testing.T,
	setup *TestSetup,
	subject string,
	subjectKeyID string,
	issuer string,
	serialNumber string,
) {
	t.Helper()

	// DA certificates indexes checks
	EnsureCertificatePresentInDaCertificateIndexes(t, setup, subject, subjectKeyID, serialNumber, true, false)

	// All certificates indexes checks
	EnsureGlobalCertificateExist(t, setup, subject, subjectKeyID, serialNumber, false)

	// UniqueCertificate: check that unique certificate key registered
	EnsureUniqueCertificateCertificateExist(t, setup, issuer, serialNumber)
}

func EnsureDaIntermediateCertificateExist(
	t *testing.T,
	setup *TestSetup,
	subject string,
	subjectKeyID string,
	issuer string,
	serialNumber string,
	skipCheckForSubject bool,
) {
	t.Helper()

	// DA certificates indexes checks
	EnsureCertificatePresentInDaCertificateIndexes(t, setup, subject, subjectKeyID, serialNumber, false, skipCheckForSubject)

	// All certificates indexes checks
	EnsureGlobalCertificateExist(t, setup, subject, subjectKeyID, serialNumber, skipCheckForSubject)

	// UniqueCertificate: check that unique certificate key registered
	EnsureUniqueCertificateCertificateExist(t, setup, issuer, serialNumber)
}

func EnsureDaRootCertificateNotExist(
	t *testing.T,
	setup *TestSetup,
	subject string,
	subjectKeyID string,
	issuer string,
	serialNumber string,
	isRevoked bool,
) {
	t.Helper()

	// DA certificates indexes checks
	EnsureCertificateNotPresentInDaCertificateIndexes(t, setup, subject, subjectKeyID, true, false, false)

	// All certificates indexes checks
	EnsureGlobalCertificateNotExist(t, setup, subject, subjectKeyID, false, false)

	// UniqueCertificate: check that unique certificate key registered
	EnsureUniqueCertificateCertificateNotExist(t, setup, issuer, serialNumber, isRevoked)
}

func EnsureDaIntermediateCertificateNotExist(
	t *testing.T,
	setup *TestSetup,
	subject string,
	subjectKeyID string,
	issuer string,
	serialNumber string,
	skipCheckForUniqueness bool,
	skipCheckForSubject bool,
) {
	t.Helper()

	// DA certificates indexes checks
	EnsureCertificateNotPresentInDaCertificateIndexes(t, setup, subject, subjectKeyID, false, skipCheckForSubject, false)

	// All certificates indexes checks
	EnsureGlobalCertificateNotExist(t, setup, subject, subjectKeyID, skipCheckForSubject, false)

	// UniqueCertificate: check that unique certificate key registered
	EnsureUniqueCertificateCertificateNotExist(t, setup, issuer, serialNumber, skipCheckForUniqueness)
}
