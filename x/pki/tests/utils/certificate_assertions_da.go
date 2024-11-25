package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func EnsureCertificatePresentInDaCertificateIndexes(
	t *testing.T,
	setup *TestSetup,
	subject string,
	subjectKeyID string,
	serialNumber string,
	isRoot bool,
	skipCheckForSubject bool,
) *types.ApprovedCertificates {
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

	return approvedCertificates
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
) *types.ApprovedCertificates {
	t.Helper()

	// DA certificates indexes checks
	certificate := EnsureCertificatePresentInDaCertificateIndexes(t, setup, subject, subjectKeyID, serialNumber, true, false)

	// All certificates indexes checks
	EnsureGlobalCertificateExist(t, setup, subject, subjectKeyID, serialNumber, false)

	// UniqueCertificate: check that unique certificate key registered
	EnsureUniqueCertificateCertificateExist(t, setup, issuer, serialNumber)

	return certificate
}

func EnsureDaIntermediateCertificateExist(
	t *testing.T,
	setup *TestSetup,
	subject string,
	subjectKeyID string,
	issuer string,
	serialNumber string,
	skipCheckForSubject bool,
) *types.ApprovedCertificates {
	t.Helper()

	// DA certificates indexes checks
	certificate := EnsureCertificatePresentInDaCertificateIndexes(t, setup, subject, subjectKeyID, serialNumber, false, skipCheckForSubject)

	// All certificates indexes checks
	EnsureGlobalCertificateExist(t, setup, subject, subjectKeyID, serialNumber, skipCheckForSubject)

	// UniqueCertificate: check that unique certificate key registered
	EnsureUniqueCertificateCertificateExist(t, setup, issuer, serialNumber)

	return certificate
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

func EnsureProposedDaRootCertificateExist(
	t *testing.T,
	setup *TestSetup,
	subject string,
	subjectKeyID string,
	serialNumber string,
) *types.ProposedCertificate {
	t.Helper()

	proposedCertificate, _ := QueryProposedCertificate(setup, subject, subjectKeyID)
	require.Equal(t, subject, proposedCertificate.Subject)
	require.Equal(t, subjectKeyID, proposedCertificate.SubjectKeyId)
	require.Equal(t, serialNumber, proposedCertificate.SerialNumber)

	return proposedCertificate
}

func EnsureRejectedDaRootCertificateExist(
	t *testing.T,
	setup *TestSetup,
	subject string,
	subjectKeyID string,
) *types.Certificate {
	t.Helper()

	proposedCertificate, _ := QueryRejectedCertificates(setup, subject, subjectKeyID)
	require.Equal(t, subject, proposedCertificate.Subject)
	require.Equal(t, subjectKeyID, proposedCertificate.SubjectKeyId)
	require.Len(t, proposedCertificate.Certs, 1)

	return proposedCertificate.Certs[0]
}
