package utils

import (
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func EnsureCertificatePresentInNocCertificateIndexes(
	t *testing.T,
	setup *TestSetup,
	subject string,
	subjectKeyID string,
	serialNumber string,
	vid int32,
	isRoot bool,
	skipCheckByVid bool,
) {
	t.Helper()

	// Noc certificates indexes checks

	// NocCertificates: Subject and SKID
	nocCertificate, _ := QueryNocCertificates(setup, subject, subjectKeyID)
	require.Equal(t, subject, nocCertificate.Subject)
	require.Equal(t, subjectKeyID, nocCertificate.SubjectKeyId)
	require.Equal(t, testconstants.SchemaVersion, nocCertificate.SchemaVersion)
	require.Len(t, nocCertificate.Certs, 1)
	require.Equal(t, serialNumber, nocCertificate.Certs[0].SerialNumber)

	// NocCertificates: SubjectKeyID
	nocCertificatesBySubjectKeyID, _ := QueryNocCertificatesBySubjectKeyID(setup, subjectKeyID)
	require.Len(t, nocCertificatesBySubjectKeyID, 1)
	require.Len(t, nocCertificatesBySubjectKeyID[0].Certs, 1)
	require.Equal(t, serialNumber, nocCertificatesBySubjectKeyID[0].Certs[0].SerialNumber)

	// NocCertificates: Subject
	nocCertificatesBySubject, _ := QueryNocCertificatesBySubject(setup, subject)
	require.Equal(t, subject, nocCertificatesBySubject.Subject)
	require.Len(t, nocCertificatesBySubject.SubjectKeyIds, 1)
	require.Equal(t, subjectKeyID, nocCertificatesBySubject.SubjectKeyIds[0])

	// NocCertificates: VID and SKID
	nocCertificateByVidAndSkid, _ := QueryNocCertificatesByVidAndSkid(setup, vid, subjectKeyID)
	require.Equal(t, vid, nocCertificateByVidAndSkid.Vid)
	require.Len(t, nocCertificateByVidAndSkid.Certs, 1)
	require.Equal(t, subjectKeyID, nocCertificateByVidAndSkid.SubjectKeyId)

	if skipCheckByVid {
		return
	}

	// NocCertificates: VID
	if isRoot {
		nocRootCertificate, _ := QueryNocRootCertificates(setup, vid)
		require.Equal(t, vid, nocRootCertificate.Vid)
		require.Len(t, nocRootCertificate.Certs, 1)
	} else {
		nocIcaCertificate, _ := QueryNocIcaCertificatesByVid(setup, vid)
		require.Equal(t, vid, nocIcaCertificate.Vid)
		require.Len(t, nocIcaCertificate.Certs, 1)
	}
}

func EnsureCertificateNotPresentInNocCertificateIndexes(
	t *testing.T,
	setup *TestSetup,
	subject string,
	subjectKeyID string,
	vid int32,
	isRoot bool,
	skipCheckByVid bool,
) {
	t.Helper()

	// Noc certificates indexes checks

	// NocCertificates: Subject and SKID
	_, err := QueryNocCertificates(setup, subject, subjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))

	// NocCertificates: SubjectKeyID
	certificatesBySubjectKeyID, _ := QueryNocCertificatesBySubjectKeyID(setup, subjectKeyID)
	require.Empty(t, certificatesBySubjectKeyID)

	// NocCertificates: Subject
	_, err = QueryNocCertificatesBySubject(setup, subject)
	require.Equal(t, codes.NotFound, status.Code(err))

	// NocCertificates: VID and SKID
	_, err = QueryNocCertificatesByVidAndSkid(setup, vid, subjectKeyID)
	require.Equal(t, codes.NotFound, status.Code(err))

	// NocCertificates: VID
	if skipCheckByVid {
		return
	}

	if isRoot {
		_, err = QueryNocRootCertificates(setup, vid)
		require.Equal(t, codes.NotFound, status.Code(err))
	} else {
		_, err = QueryNocIcaCertificatesByVid(setup, vid)
		require.Equal(t, codes.NotFound, status.Code(err))
	}
}

func EnsureNocRootCertificateExist(
	t *testing.T,
	setup *TestSetup,
	subject string,
	subjectKeyID string,
	issuer string,
	serialNumber string,
	vid int32,
) {
	t.Helper()

	// Noc certificates indexes checks
	EnsureCertificatePresentInNocCertificateIndexes(t, setup, subject, subjectKeyID, serialNumber, vid, true, false)

	// All certificates indexes checks
	EnsureGlobalCertificateExist(t, setup, subject, subjectKeyID, serialNumber, false)

	// UniqueCertificate: check that unique certificate key registered
	EnsureUniqueCertificateCertificateExist(t, setup, issuer, serialNumber)
}

func EnsureNocIntermediateCertificateExist(
	t *testing.T,
	setup *TestSetup,
	subject string,
	subjectKeyID string,
	issuer string,
	serialNumber string,
	vid int32,
	skipCheckByVid bool,
) {
	t.Helper()

	// Noc certificates indexes checks
	EnsureCertificatePresentInNocCertificateIndexes(t, setup, subject, subjectKeyID, serialNumber, vid, false, skipCheckByVid)

	// All certificates indexes checks
	EnsureGlobalCertificateExist(t, setup, subject, subjectKeyID, serialNumber, false)

	// UniqueCertificate: check that unique certificate key registered
	EnsureUniqueCertificateCertificateExist(t, setup, issuer, serialNumber)
}

func EnsureNocIntermediateCertificateNotExist(
	t *testing.T,
	setup *TestSetup,
	subject string,
	subjectKeyID string,
	issuer string,
	serialNumber string,
	vid int32,
	skipCheckByVid bool,
	skipCheckForUniqueness bool,
) {
	t.Helper()

	// Noc certificates indexes checks
	EnsureCertificateNotPresentInNocCertificateIndexes(t, setup, subject, subjectKeyID, vid, false, skipCheckByVid)

	// All certificates indexes checks
	EnsureGlobalCertificateNotExist(t, setup, subject, subjectKeyID, false, false)

	// UniqueCertificate: check that unique certificate key registered
	EnsureUniqueCertificateCertificateNotExist(t, setup, issuer, serialNumber, skipCheckForUniqueness)
}

func EnsureNocRootCertificateNotExist(
	t *testing.T,
	setup *TestSetup,
	subject string,
	subjectKeyID string,
	issuer string,
	serialNumber string,
	vid int32,
	skipCheckByVid bool,
	skipCheckForUniqueness bool,
) {
	t.Helper()

	// Noc certificates indexes checks
	EnsureCertificateNotPresentInNocCertificateIndexes(t, setup, subject, subjectKeyID, vid, true, skipCheckByVid)

	// All certificates indexes checks
	EnsureGlobalCertificateNotExist(t, setup, subject, subjectKeyID, false, false)

	// UniqueCertificate: check that unique certificate key registered
	EnsureUniqueCertificateCertificateNotExist(t, setup, issuer, serialNumber, skipCheckForUniqueness)
}
