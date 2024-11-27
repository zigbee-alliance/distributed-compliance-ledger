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
	Exist bool
	Count int
}

type TestCertificate struct {
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
	AllCertificatesBySubjectKeyId      []types.AllCertificates
	ApprovedCertificates               *types.ApprovedCertificates
	ApprovedCertificatesBySubject      *types.ApprovedCertificatesBySubject
	ApprovedCertificatesBySubjectKeyId []types.ApprovedCertificates
	ApprovedRootCertificates           *types.CertificateIdentifier
	ProposedCertificate                *types.ProposedCertificate
	RejectedCertificate                *types.RejectedCertificate
	ChildCertificates                  *types.ChildCertificates
	NocCertificates                    *types.NocCertificates
	NocCertificatesBySubject           *types.NocCertificatesBySubject
	NocCertificatesBySubjectKeyId      []types.NocCertificates
	ProposedRevocation                 *types.ProposedCertificateRevocation
}

func CheckCertificateStateIndexes(
	t *testing.T,
	setup *TestSetup,
	certificate TestCertificate,
	indexes []TestIndex,
) ResolvedCertificate {
	var resolvedCertificate ResolvedCertificate

	for _, index := range indexes {
		if index.Key == types.AllCertificatesKeyPrefix {
			if index.Exist {
				certificates, _ := QueryAllCertificates(setup, certificate.Subject, certificate.SubjectKeyID)
				require.Equal(t, certificate.Subject, certificates.Subject)
				require.Equal(t, certificate.SubjectKeyID, certificates.SubjectKeyId)
				require.Len(t, certificates.Certs, GetExpectedCount(index))
				require.Equal(t, certificate.IsRoot, certificates.Certs[0].IsRoot)
				resolvedCertificate.AllCertificates = certificates
			} else {
				_, err := QueryAllCertificates(setup, certificate.Subject, certificate.SubjectKeyID)
				require.Equal(t, codes.NotFound, status.Code(err))
			}
		}
		if index.Key == types.AllCertificatesBySubjectKeyPrefix {
			if index.Exist {
				certificatesBySubject, _ := QueryAllCertificatesBySubject(setup, certificate.Subject)
				require.Len(t, certificatesBySubject.SubjectKeyIds, GetExpectedCount(index))
				require.Equal(t, certificate.SubjectKeyID, certificatesBySubject.SubjectKeyIds[0])
				resolvedCertificate.AllCertificatesBySubject = certificatesBySubject
			} else {
				_, err := QueryAllCertificatesBySubject(setup, certificate.Subject)
				require.Equal(t, codes.NotFound, status.Code(err))
			}
		}
		if index.Key == types.AllCertificatesBySubjectKeyIDKeyPrefix {
			if index.Exist {
				certificateBySubjectKeyID, _ := QueryAllCertificatesBySubjectKeyID(setup, certificate.SubjectKeyID)
				require.Len(t, certificateBySubjectKeyID[0].Certs, GetExpectedCount(index))
				require.Equal(t, certificate.IsRoot, certificateBySubjectKeyID[0].Certs[0].IsRoot)
				resolvedCertificate.AllCertificatesBySubjectKeyId = certificateBySubjectKeyID
			} else {
				certificatesBySubjectKeyID, _ := QueryAllCertificatesBySubjectKeyID(setup, certificate.SubjectKeyID)
				require.Empty(t, certificatesBySubjectKeyID)
			}
		}
		if index.Key == types.ApprovedCertificatesKeyPrefix {
			if index.Exist {
				certificates, _ := QueryApprovedCertificates(setup, certificate.Subject, certificate.SubjectKeyID)
				require.Equal(t, certificate.Subject, certificates.Subject)
				require.Equal(t, certificate.SubjectKeyID, certificates.SubjectKeyId)
				require.Len(t, certificates.Certs, GetExpectedCount(index))
				require.Equal(t, certificate.IsRoot, certificates.Certs[0].IsRoot)
				resolvedCertificate.ApprovedCertificates = certificates
			} else {
				_, err := QueryApprovedCertificates(setup, certificate.Subject, certificate.SubjectKeyID)
				require.Equal(t, codes.NotFound, status.Code(err))
			}
		}
		if index.Key == types.ApprovedCertificatesBySubjectKeyPrefix {
			if index.Exist {
				certificatesBySubject, _ := QueryApprovedCertificatesBySubject(setup, certificate.Subject)
				require.Len(t, certificatesBySubject.SubjectKeyIds, GetExpectedCount(index))
				require.Equal(t, certificate.SubjectKeyID, certificatesBySubject.SubjectKeyIds[0])
				resolvedCertificate.ApprovedCertificatesBySubject = certificatesBySubject
			} else {
				_, err := QueryApprovedCertificatesBySubject(setup, certificate.Subject)
				require.Equal(t, codes.NotFound, status.Code(err))
			}
		}
		if index.Key == types.ApprovedCertificatesBySubjectKeyIDKeyPrefix {
			if index.Exist {
				approvedCertificatesBySkid, _ := QueryApprovedCertificatesBySubjectKeyID(setup, certificate.SubjectKeyID)
				require.Len(t, approvedCertificatesBySkid, 1)
				require.Len(t, approvedCertificatesBySkid[0].Certs, GetExpectedCount(index))
				require.Equal(t, certificate.IsRoot, approvedCertificatesBySkid[0].Certs[0].IsRoot)
				resolvedCertificate.ApprovedCertificatesBySubjectKeyId = approvedCertificatesBySkid
			} else {
				certificatesBySubjectKeyID, _ := QueryApprovedCertificatesBySubjectKeyID(setup, certificate.SubjectKeyID)
				require.Empty(t, certificatesBySubjectKeyID)
			}
		}
		if index.Key == types.ApprovedRootCertificatesKeyPrefix {
			if index.Exist {
				approvedRootCertificate, _ := QueryApprovedRootCertificates(setup, certificate.Subject, certificate.SubjectKeyID)
				require.Equal(t, certificate.Subject, approvedRootCertificate.Subject)
				require.Equal(t, certificate.SubjectKeyID, approvedRootCertificate.SubjectKeyId)
				resolvedCertificate.ApprovedRootCertificates = approvedRootCertificate
			} else {
				_, err := QueryApprovedRootCertificates(setup, certificate.Subject, certificate.SubjectKeyID)
				require.Equal(t, codes.NotFound, status.Code(err))
			}
		}
		if index.Key == types.ProposedCertificateKeyPrefix {
			if index.Exist {
				proposedCertificate, _ := QueryProposedCertificate(setup, certificate.Subject, certificate.SubjectKeyID)
				require.Equal(t, certificate.Subject, proposedCertificate.Subject)
				require.Equal(t, certificate.SubjectKeyID, proposedCertificate.SubjectKeyId)
				resolvedCertificate.ProposedCertificate = proposedCertificate
			} else {
				_, err := QueryProposedCertificate(setup, certificate.Subject, certificate.SubjectKeyID)
				require.Equal(t, codes.NotFound, status.Code(err))
			}
		}
		if index.Key == types.RejectedCertificateKeyPrefix {
			if index.Exist {
				rejectedCertificate, _ := QueryRejectedCertificates(setup, certificate.Subject, certificate.SubjectKeyID)
				require.Equal(t, certificate.Subject, rejectedCertificate.Subject)
				require.Equal(t, certificate.SubjectKeyID, rejectedCertificate.SubjectKeyId)
				require.Len(t, rejectedCertificate.Certs, GetExpectedCount(index))
				resolvedCertificate.RejectedCertificate = rejectedCertificate
			} else {
				_, err := QueryRejectedCertificates(setup, certificate.Subject, certificate.SubjectKeyID)
				require.Equal(t, codes.NotFound, status.Code(err))
			}
		}
		if index.Key == types.ChildCertificatesKeyPrefix {
			if index.Exist {
				issuerChildren, _ := QueryChildCertificates(setup, certificate.Issuer, certificate.AuthorityKeyID)
				require.Len(t, issuerChildren.CertIds, GetExpectedCount(index))
				certID := types.CertificateIdentifier{
					Subject:      certificate.Subject,
					SubjectKeyId: certificate.SubjectKeyID,
				}
				require.Equal(t, &certID, issuerChildren.CertIds[0])
				resolvedCertificate.ChildCertificates = issuerChildren
			} else {
				_, err := QueryChildCertificates(setup, certificate.Issuer, certificate.AuthorityKeyID)
				require.Equal(t, codes.NotFound, status.Code(err))
			}
		}
		if index.Key == types.UniqueCertificateKeyPrefix {
			require.Equal(t, index.Exist, setup.Keeper.IsUniqueCertificatePresent(
				setup.Ctx, certificate.Issuer, certificate.SerialNumber))
		}
		if index.Key == types.NocCertificatesKeyPrefix {
			if index.Exist {
				certificates, _ := QueryNocCertificates(setup, certificate.Subject, certificate.SubjectKeyID)
				require.Equal(t, certificate.Subject, certificates.Subject)
				require.Equal(t, certificate.SubjectKeyID, certificates.SubjectKeyId)
				require.Len(t, certificates.Certs, GetExpectedCount(index))
				resolvedCertificate.NocCertificates = certificates
			} else {
				_, err := QueryNocCertificates(setup, certificate.Subject, certificate.SubjectKeyID)
				require.Equal(t, codes.NotFound, status.Code(err))
			}
		}
		if index.Key == types.NocCertificatesBySubjectKeyIDKeyPrefix {
			if index.Exist {
				nocCertificatesBySkid, _ := QueryNocCertificatesBySubjectKeyID(setup, certificate.SubjectKeyID)
				require.Len(t, nocCertificatesBySkid, 1)
				require.Len(t, nocCertificatesBySkid[0].Certs, GetExpectedCount(index))
				require.Equal(t, certificate.IsRoot, nocCertificatesBySkid[0].Certs[0].IsRoot)
				resolvedCertificate.NocCertificatesBySubjectKeyId = nocCertificatesBySkid
			} else {
				certificatesBySubjectKeyID, _ := QueryNocCertificatesBySubjectKeyID(setup, certificate.SubjectKeyID)
				require.Empty(t, certificatesBySubjectKeyID)
			}
		}
		if index.Key == types.NocCertificatesBySubjectKeyPrefix {
			if index.Exist {
				nocCertificatesBySubject, _ := QueryNocCertificatesBySubject(setup, certificate.Subject)
				require.Len(t, nocCertificatesBySubject.SubjectKeyIds, GetExpectedCount(index))
				require.Equal(t, certificate.SubjectKeyID, nocCertificatesBySubject.SubjectKeyIds[0])
				resolvedCertificate.NocCertificatesBySubject = nocCertificatesBySubject
			} else {
				_, err := QueryNocCertificatesBySubject(setup, certificate.Subject)
				require.Equal(t, codes.NotFound, status.Code(err))
			}
		}
		if index.Key == types.NocCertificatesByVidAndSkidKeyPrefix {
			if index.Exist {
				nocCertificatesByVidAndSkid, _ := QueryNocCertificatesByVidAndSkid(setup, certificate.VID, certificate.SubjectKeyID)
				require.Equal(t, certificate.VID, nocCertificatesByVidAndSkid.Vid)
				require.Len(t, nocCertificatesByVidAndSkid.Certs, GetExpectedCount(index))
				require.Equal(t, certificate.SubjectKeyID, nocCertificatesByVidAndSkid.SubjectKeyId)
			} else {
				_, err := QueryNocCertificatesByVidAndSkid(setup, certificate.VID, certificate.SubjectKeyID)
				require.Equal(t, codes.NotFound, status.Code(err))
			}
		}
		if index.Key == types.NocRootCertificatesKeyPrefix {
			if index.Exist {
				nocRootCertificatesByVid, _ := QueryNocRootCertificatesByVid(setup, certificate.VID)
				require.Equal(t, certificate.VID, nocRootCertificatesByVid.Vid)
				require.Len(t, nocRootCertificatesByVid.Certs, GetExpectedCount(index))
			} else {
				_, err := QueryNocRootCertificatesByVid(setup, certificate.VID)
				require.Equal(t, codes.NotFound, status.Code(err))
			}
		}
		if index.Key == types.NocIcaCertificatesKeyPrefix {
			if index.Exist {
				nocIcaCertificatesBy, _ := QueryNocIcaCertificatesByVid(setup, certificate.VID)
				require.Equal(t, certificate.VID, nocIcaCertificatesBy.Vid)
				require.Len(t, nocIcaCertificatesBy.Certs, GetExpectedCount(index))
			} else {
				_, err := QueryNocIcaCertificatesByVid(setup, certificate.VID)
				require.Equal(t, codes.NotFound, status.Code(err))
			}
		}
		if index.Key == types.RevokedNocIcaCertificatesKeyPrefix {
			require.Equal(t, index.Exist, setup.Keeper.IsRevokedNocIcaCertificatePresent(
				setup.Ctx, certificate.Subject, certificate.SubjectKeyID))
		}
		if index.Key == types.RevokedNocRootCertificatesKeyPrefix {
			require.Equal(t, index.Exist, setup.Keeper.IsRevokedNocRootCertificatePresent(
				setup.Ctx, certificate.Subject, certificate.SubjectKeyID))
		}
		if index.Key == types.RevokedCertificatesKeyPrefix {
			require.Equal(t, index.Exist, setup.Keeper.IsRevokedCertificatePresent(
				setup.Ctx, certificate.Subject, certificate.SubjectKeyID))
		}
		if index.Key == types.ProposedCertificateRevocationKeyPrefix {
			if index.Exist {
				proposedRevocation, _ := QueryProposedCertificateRevocation(
					setup,
					certificate.Subject,
					certificate.SubjectKeyID,
					certificate.SerialNumber,
				)
				resolvedCertificate.ProposedRevocation = proposedRevocation
			} else {
				_, err := QueryProposedCertificateRevocation(
					setup,
					certificate.Subject,
					certificate.SubjectKeyID,
					certificate.SerialNumber,
				)
				require.Equal(t, codes.NotFound, status.Code(err))
			}
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
