package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TestIndex struct {
	Key   string
	Count int
}

type TestIndexes struct {
	Present []TestIndex
	Missing []TestIndex
}

type TestCertificate struct {
	PEM            string
	Subject        string
	SubjectKeyID   string
	Issuer         string
	AuthorityKeyID string
	SerialNumber   string
	VID            int32
	IsRoot         bool
}

type ResolvedCertificate struct {
	AllCertificates                    *types.AllCertificates
	AllCertificatesBySubject           *types.AllCertificatesBySubject
	AllCertificatesBySubjectKeyID      []types.AllCertificates
	ApprovedCertificates               *types.ApprovedCertificates
	ApprovedCertificatesBySubject      *types.ApprovedCertificatesBySubject
	ApprovedCertificatesBySubjectKeyID []types.ApprovedCertificates
	ApprovedRootCertificates           *types.CertificateIdentifier
	ProposedCertificate                *types.ProposedCertificate
	RejectedCertificate                *types.RejectedCertificate
	ChildCertificates                  *types.ChildCertificates
	NocCertificates                    *types.NocCertificates
	NocCertificatesBySubject           *types.NocCertificatesBySubject
	NocCertificatesBySubjectKeyID      []types.NocCertificates
	ProposedRevocation                 *types.ProposedCertificateRevocation
	RevokedCertificates                *types.RevokedCertificates
	RevokedNocIcaCertificates          *types.RevokedNocIcaCertificates
	RevokedNocRootCertificates         *types.RevokedNocRootCertificates
}

//nolint:gocyclo
func CheckCertificateStateIndexes(
	t *testing.T,
	setup *TestSetup,
	certificate TestCertificate,
	indexes TestIndexes,
) ResolvedCertificate {
	t.Helper()

	var resolvedCertificate ResolvedCertificate

	for _, index := range indexes.Present {
		if index.Key == types.AllCertificatesKeyPrefix {
			certificates, _ := QueryAllCertificates(setup, certificate.Subject, certificate.SubjectKeyID)
			require.Equal(t, certificate.Subject, certificates.Subject)
			require.Equal(t, certificate.SubjectKeyID, certificates.SubjectKeyId)
			require.Len(t, certificates.Certs, GetExpectedCount(index))
			require.Equal(t, certificate.IsRoot, certificates.Certs[0].IsRoot)
			resolvedCertificate.AllCertificates = certificates
		}
		if index.Key == types.AllCertificatesBySubjectKeyPrefix {
			certificatesBySubject, _ := QueryAllCertificatesBySubject(setup, certificate.Subject)
			require.Len(t, certificatesBySubject.SubjectKeyIds, GetExpectedCount(index))
			require.Equal(t, certificate.SubjectKeyID, certificatesBySubject.SubjectKeyIds[0])
			resolvedCertificate.AllCertificatesBySubject = certificatesBySubject
		}
		if index.Key == types.AllCertificatesBySubjectKeyIDKeyPrefix {
			certificateBySubjectKeyID, _ := QueryAllCertificatesBySubjectKeyID(setup, certificate.SubjectKeyID)
			require.Len(t, certificateBySubjectKeyID[0].Certs, GetExpectedCount(index))
			require.Equal(t, certificate.IsRoot, certificateBySubjectKeyID[0].Certs[0].IsRoot)
			resolvedCertificate.AllCertificatesBySubjectKeyID = certificateBySubjectKeyID
		}
		if index.Key == types.ApprovedCertificatesKeyPrefix {
			certificates, _ := QueryApprovedCertificates(setup, certificate.Subject, certificate.SubjectKeyID)
			require.Equal(t, certificate.Subject, certificates.Subject)
			require.Equal(t, certificate.SubjectKeyID, certificates.SubjectKeyId)
			require.Len(t, certificates.Certs, GetExpectedCount(index))
			require.Equal(t, certificate.IsRoot, certificates.Certs[0].IsRoot)
			resolvedCertificate.ApprovedCertificates = certificates
		}
		if index.Key == types.ApprovedCertificatesBySubjectKeyPrefix {
			certificatesBySubject, _ := QueryApprovedCertificatesBySubject(setup, certificate.Subject)
			require.Len(t, certificatesBySubject.SubjectKeyIds, GetExpectedCount(index))
			require.Equal(t, certificate.SubjectKeyID, certificatesBySubject.SubjectKeyIds[0])
			resolvedCertificate.ApprovedCertificatesBySubject = certificatesBySubject
		}
		if index.Key == types.ApprovedCertificatesBySubjectKeyIDKeyPrefix {
			approvedCertificatesBySkid, _ := QueryApprovedCertificatesBySubjectKeyID(setup, certificate.SubjectKeyID)
			require.Len(t, approvedCertificatesBySkid, 1)
			require.Len(t, approvedCertificatesBySkid[0].Certs, GetExpectedCount(index))
			require.Equal(t, certificate.IsRoot, approvedCertificatesBySkid[0].Certs[0].IsRoot)
			resolvedCertificate.ApprovedCertificatesBySubjectKeyID = approvedCertificatesBySkid
		}
		if index.Key == types.ApprovedRootCertificatesKeyPrefix {
			approvedRootCertificate, _ := QueryApprovedRootCertificates(setup, certificate.Subject, certificate.SubjectKeyID)
			require.Equal(t, certificate.Subject, approvedRootCertificate.Subject)
			require.Equal(t, certificate.SubjectKeyID, approvedRootCertificate.SubjectKeyId)
			resolvedCertificate.ApprovedRootCertificates = approvedRootCertificate
		}
		if index.Key == types.ProposedCertificateKeyPrefix {
			proposedCertificate, _ := QueryProposedCertificate(setup, certificate.Subject, certificate.SubjectKeyID)
			require.Equal(t, certificate.Subject, proposedCertificate.Subject)
			require.Equal(t, certificate.SubjectKeyID, proposedCertificate.SubjectKeyId)
			resolvedCertificate.ProposedCertificate = proposedCertificate
		}
		if index.Key == types.RejectedCertificateKeyPrefix {
			rejectedCertificate, _ := QueryRejectedCertificates(setup, certificate.Subject, certificate.SubjectKeyID)
			require.Equal(t, certificate.Subject, rejectedCertificate.Subject)
			require.Equal(t, certificate.SubjectKeyID, rejectedCertificate.SubjectKeyId)
			require.Len(t, rejectedCertificate.Certs, GetExpectedCount(index))
			resolvedCertificate.RejectedCertificate = rejectedCertificate
		}
		if index.Key == types.ChildCertificatesKeyPrefix {
			issuerChildren, _ := QueryChildCertificates(setup, certificate.Issuer, certificate.AuthorityKeyID)
			require.Len(t, issuerChildren.CertIds, GetExpectedCount(index))
			certID := types.CertificateIdentifier{
				Subject:      certificate.Subject,
				SubjectKeyId: certificate.SubjectKeyID,
			}
			require.Equal(t, &certID, issuerChildren.CertIds[0])
			resolvedCertificate.ChildCertificates = issuerChildren
		}
		if index.Key == types.UniqueCertificateKeyPrefix {
			require.True(t, setup.Keeper.IsUniqueCertificatePresent(
				setup.Ctx, certificate.Issuer, certificate.SerialNumber))
		}
		if index.Key == types.NocCertificatesKeyPrefix {
			certificates, _ := QueryNocCertificates(setup, certificate.Subject, certificate.SubjectKeyID)
			require.Equal(t, certificate.Subject, certificates.Subject)
			require.Equal(t, certificate.SubjectKeyID, certificates.SubjectKeyId)
			require.Len(t, certificates.Certs, GetExpectedCount(index))
			resolvedCertificate.NocCertificates = certificates
		}
		if index.Key == types.NocCertificatesBySubjectKeyIDKeyPrefix {
			nocCertificatesBySkid, _ := QueryNocCertificatesBySubjectKeyID(setup, certificate.SubjectKeyID)
			require.Len(t, nocCertificatesBySkid, 1)
			require.Len(t, nocCertificatesBySkid[0].Certs, GetExpectedCount(index))
			require.Equal(t, certificate.IsRoot, nocCertificatesBySkid[0].Certs[0].IsRoot)
			resolvedCertificate.NocCertificatesBySubjectKeyID = nocCertificatesBySkid
		}
		if index.Key == types.NocCertificatesBySubjectKeyPrefix {
			nocCertificatesBySubject, _ := QueryNocCertificatesBySubject(setup, certificate.Subject)
			require.Len(t, nocCertificatesBySubject.SubjectKeyIds, GetExpectedCount(index))
			require.Equal(t, certificate.SubjectKeyID, nocCertificatesBySubject.SubjectKeyIds[0])
			resolvedCertificate.NocCertificatesBySubject = nocCertificatesBySubject
		}
		if index.Key == types.NocCertificatesByVidAndSkidKeyPrefix {
			nocCertificatesByVidAndSkid, _ := QueryNocCertificatesByVidAndSkid(setup, certificate.VID, certificate.SubjectKeyID)
			require.Equal(t, certificate.VID, nocCertificatesByVidAndSkid.Vid)
			require.Len(t, nocCertificatesByVidAndSkid.Certs, GetExpectedCount(index))
			require.Equal(t, certificate.SubjectKeyID, nocCertificatesByVidAndSkid.SubjectKeyId)
		}
		if index.Key == types.NocRootCertificatesKeyPrefix {
			nocRootCertificatesByVid, _ := QueryNocRootCertificatesByVid(setup, certificate.VID)
			require.Equal(t, certificate.VID, nocRootCertificatesByVid.Vid)
			require.Len(t, nocRootCertificatesByVid.Certs, GetExpectedCount(index))
		}
		if index.Key == types.NocIcaCertificatesKeyPrefix {
			nocIcaCertificatesBy, _ := QueryNocIcaCertificatesByVid(setup, certificate.VID)
			require.Equal(t, certificate.VID, nocIcaCertificatesBy.Vid)
			require.Len(t, nocIcaCertificatesBy.Certs, GetExpectedCount(index))
		}
		if index.Key == types.RevokedNocIcaCertificatesKeyPrefix {
			revokedNocIcaCertificates, _ := QueryNocRevokedIcaCertificates(setup, certificate.Subject, certificate.SubjectKeyID)
			require.Len(t, revokedNocIcaCertificates.Certs, GetExpectedCount(index))
			require.Equal(t, certificate.Subject, revokedNocIcaCertificates.Subject)
			require.Equal(t, certificate.SubjectKeyID, revokedNocIcaCertificates.SubjectKeyId)
			resolvedCertificate.RevokedNocIcaCertificates = revokedNocIcaCertificates
		}
		if index.Key == types.RevokedNocRootCertificatesKeyPrefix {
			revokedNocRootCertificates, _ := QueryNocRevokedRootCertificates(setup, certificate.Subject, certificate.SubjectKeyID)
			require.Len(t, revokedNocRootCertificates.Certs, GetExpectedCount(index))
			require.Equal(t, certificate.Subject, revokedNocRootCertificates.Subject)
			require.Equal(t, certificate.SubjectKeyID, revokedNocRootCertificates.SubjectKeyId)
			resolvedCertificate.RevokedNocRootCertificates = revokedNocRootCertificates
		}
		if index.Key == types.RevokedCertificatesKeyPrefix {
			revokedCertificates, _ := QueryRevokedCertificates(setup, certificate.Subject, certificate.SubjectKeyID)
			require.Len(t, revokedCertificates.Certs, GetExpectedCount(index))
			require.Equal(t, certificate.Subject, revokedCertificates.Subject)
			require.Equal(t, certificate.SubjectKeyID, revokedCertificates.SubjectKeyId)
			resolvedCertificate.RevokedCertificates = revokedCertificates
		}
		if index.Key == types.ProposedCertificateRevocationKeyPrefix {
			proposedRevocation, _ := QueryProposedCertificateRevocation(
				setup,
				certificate.Subject,
				certificate.SubjectKeyID,
				certificate.SerialNumber,
			)
			resolvedCertificate.ProposedRevocation = proposedRevocation
		}
	}

	for _, index := range indexes.Missing {
		if index.Key == types.AllCertificatesKeyPrefix {
			_, err := QueryAllCertificates(setup, certificate.Subject, certificate.SubjectKeyID)
			require.Equal(t, codes.NotFound, status.Code(err))
		}
		if index.Key == types.AllCertificatesBySubjectKeyPrefix {
			_, err := QueryAllCertificatesBySubject(setup, certificate.Subject)
			require.Equal(t, codes.NotFound, status.Code(err))
		}
		if index.Key == types.AllCertificatesBySubjectKeyIDKeyPrefix {
			certificatesBySubjectKeyID, _ := QueryAllCertificatesBySubjectKeyID(setup, certificate.SubjectKeyID)
			require.Empty(t, certificatesBySubjectKeyID)
		}
		if index.Key == types.ApprovedCertificatesKeyPrefix {
			_, err := QueryApprovedCertificates(setup, certificate.Subject, certificate.SubjectKeyID)
			require.Equal(t, codes.NotFound, status.Code(err))
		}
		if index.Key == types.ApprovedCertificatesBySubjectKeyPrefix {
			_, err := QueryApprovedCertificatesBySubject(setup, certificate.Subject)
			require.Equal(t, codes.NotFound, status.Code(err))
		}
		if index.Key == types.ApprovedCertificatesBySubjectKeyIDKeyPrefix {
			certificatesBySubjectKeyID, _ := QueryApprovedCertificatesBySubjectKeyID(setup, certificate.SubjectKeyID)
			require.Empty(t, certificatesBySubjectKeyID)
		}
		if index.Key == types.ApprovedRootCertificatesKeyPrefix {
			_, err := QueryApprovedRootCertificates(setup, certificate.Subject, certificate.SubjectKeyID)
			require.Equal(t, codes.NotFound, status.Code(err))
		}
		if index.Key == types.ProposedCertificateKeyPrefix {
			_, err := QueryProposedCertificate(setup, certificate.Subject, certificate.SubjectKeyID)
			require.Equal(t, codes.NotFound, status.Code(err))
		}
		if index.Key == types.RejectedCertificateKeyPrefix {
			_, err := QueryRejectedCertificates(setup, certificate.Subject, certificate.SubjectKeyID)
			require.Equal(t, codes.NotFound, status.Code(err))
		}
		if index.Key == types.ChildCertificatesKeyPrefix {
			_, err := QueryChildCertificates(setup, certificate.Issuer, certificate.AuthorityKeyID)
			require.Equal(t, codes.NotFound, status.Code(err))
		}
		if index.Key == types.UniqueCertificateKeyPrefix {
			require.False(t, setup.Keeper.IsUniqueCertificatePresent(
				setup.Ctx, certificate.Issuer, certificate.SerialNumber))
		}
		if index.Key == types.NocCertificatesKeyPrefix {
			_, err := QueryNocCertificates(setup, certificate.Subject, certificate.SubjectKeyID)
			require.Equal(t, codes.NotFound, status.Code(err))
		}
		if index.Key == types.NocCertificatesBySubjectKeyIDKeyPrefix {
			certificatesBySubjectKeyID, _ := QueryNocCertificatesBySubjectKeyID(setup, certificate.SubjectKeyID)
			require.Empty(t, certificatesBySubjectKeyID)
		}
		if index.Key == types.NocCertificatesBySubjectKeyPrefix {
			_, err := QueryNocCertificatesBySubject(setup, certificate.Subject)
			require.Equal(t, codes.NotFound, status.Code(err))
		}
		if index.Key == types.NocCertificatesByVidAndSkidKeyPrefix {
			_, err := QueryNocCertificatesByVidAndSkid(setup, certificate.VID, certificate.SubjectKeyID)
			require.Equal(t, codes.NotFound, status.Code(err))
		}
		if index.Key == types.NocRootCertificatesKeyPrefix {
			_, err := QueryNocRootCertificatesByVid(setup, certificate.VID)
			require.Equal(t, codes.NotFound, status.Code(err))
		}
		if index.Key == types.NocIcaCertificatesKeyPrefix {
			_, err := QueryNocIcaCertificatesByVid(setup, certificate.VID)
			require.Equal(t, codes.NotFound, status.Code(err))
		}
		if index.Key == types.RevokedNocIcaCertificatesKeyPrefix {
			_, err := QueryNocRevokedIcaCertificates(setup, certificate.Subject, certificate.SubjectKeyID)
			require.Equal(t, codes.NotFound, status.Code(err))
		}
		if index.Key == types.RevokedNocRootCertificatesKeyPrefix {
			_, err := QueryNocRevokedRootCertificates(setup, certificate.Subject, certificate.SubjectKeyID)
			require.Equal(t, codes.NotFound, status.Code(err))
		}
		if index.Key == types.RevokedCertificatesKeyPrefix {
			_, err := QueryRevokedCertificates(setup, certificate.Subject, certificate.SubjectKeyID)
			require.Equal(t, codes.NotFound, status.Code(err))
		}
		if index.Key == types.ProposedCertificateRevocationKeyPrefix {
			_, err := QueryProposedCertificateRevocation(
				setup,
				certificate.Subject,
				certificate.SubjectKeyID,
				certificate.SerialNumber,
			)
			require.Equal(t, codes.NotFound, status.Code(err))
		}
	}

	return resolvedCertificate
}

func GetExpectedCount(index TestIndex) int {
	count := index.Count
	if index.Count == 0 {
		count = 1
	}

	return count
}
