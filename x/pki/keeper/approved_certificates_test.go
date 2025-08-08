package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/nullify"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// Prevent strconv unused error.
var _ = strconv.IntSize

func createNApprovedCertificates(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ApprovedCertificates {
	items := make([]types.ApprovedCertificates, n)
	for i := range items {
		items[i].Subject = strconv.Itoa(i)
		items[i].SubjectKeyId = strconv.Itoa(i)

		keeper.SetApprovedCertificates(ctx, items[i])
		keeper.SetApprovedCertificatesBySubjectKeyID(ctx, types.ApprovedCertificatesBySubjectKeyId{
			SubjectKeyId: items[i].SubjectKeyId,
			Certs:        items[i].Certs,
		})
	}

	return items
}

func TestApprovedCertificatesGet(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNApprovedCertificates(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetApprovedCertificates(ctx,
			item.Subject,
			item.SubjectKeyId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestApprovedCertificatesRemove(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNApprovedCertificates(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveApprovedCertificates(ctx,
			item.Subject,
			item.SubjectKeyId,
		)
		_, found := keeper.GetApprovedCertificates(ctx,
			item.Subject,
			item.SubjectKeyId,
		)
		require.False(t, found)
	}
}

func TestApprovedCertificatesGetAll(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNApprovedCertificates(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllApprovedCertificates(ctx)),
	)
}

func TestApprovedCertificates_ExtendedOperations(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	
	tests := []struct {
		name        string
		setup       func() (types.ApprovedCertificates, string, string)
		operation   func(types.ApprovedCertificates, string, string)
		verify      func(string, string) (types.ApprovedCertificates, bool)
		expectFound bool
		description string
	}{
		{
			name: "BySubjectKeyID",
			setup: func() (types.ApprovedCertificates, string, string) {
				subjectKeyID := "test-key-id"
				approvedCerts := types.ApprovedCertificatesBySubjectKeyId{
					SubjectKeyId: subjectKeyID,
					Certs: []*types.CertificateIdentifier{
						{
							Subject:      "test-subject",
							SubjectKeyId: subjectKeyID,
						},
					},
				}
				keeper.SetApprovedCertificatesBySubjectKeyID(ctx, approvedCerts)
				return types.ApprovedCertificates{}, subjectKeyID, ""
			},
			operation: func(cert types.ApprovedCertificates, subjectKeyID, _ string) {
				keeper.RemoveApprovedCertificatesBySubjectKeyID(ctx, subjectKeyID)
			},
			verify: func(subjectKeyID, _ string) (types.ApprovedCertificates, bool) {
				return keeper.GetApprovedCertificatesBySubjectKeyID(ctx, subjectKeyID)
			},
			expectFound: false,
			description: "Test setting and getting approved certificates by subject key ID",
		},
		{
			name: "BySubject",
			setup: func() (types.ApprovedCertificates, string, string) {
				subject := "test-subject"
				approvedCerts := types.ApprovedCertificatesBySubject{
					Subject: subject,
					Certs: []*types.CertificateIdentifier{
						{
							Subject:      subject,
							SubjectKeyId: "test-key-id",
						},
					},
				}
				keeper.SetApprovedCertificatesBySubject(ctx, approvedCerts)
				return types.ApprovedCertificates{}, "", subject
			},
			operation: func(cert types.ApprovedCertificates, _, subject string) {
				keeper.RemoveApprovedCertificatesBySubject(ctx, subject)
			},
			verify: func(_, subject string) (types.ApprovedCertificates, bool) {
				return keeper.GetApprovedCertificatesBySubject(ctx, subject)
			},
			expectFound: false,
			description: "Test setting and getting approved certificates by subject",
		},
		{
			name: "EmptyCertificates",
			setup: func() (types.ApprovedCertificates, string, string) {
				emptyCert := types.ApprovedCertificates{
					Subject:      "empty-subject",
					SubjectKeyId: "empty-key-id",
					Certs:        []*types.CertificateIdentifier{},
				}
				keeper.SetApprovedCertificates(ctx, emptyCert)
				return emptyCert, emptyCert.Subject, emptyCert.SubjectKeyId
			},
			operation: func(cert types.ApprovedCertificates, subject, subjectKeyId string) {
				// No operation needed for this test
			},
			verify: func(subject, subjectKeyId string) (types.ApprovedCertificates, bool) {
				return keeper.GetApprovedCertificates(ctx, subject, subjectKeyId)
			},
			expectFound: true,
			description: "Test with empty certificates",
		},
		{
			name: "NilCertificates",
			setup: func() (types.ApprovedCertificates, string, string) {
				nilCert := types.ApprovedCertificates{
					Subject:      "nil-subject",
					SubjectKeyId: "nil-key-id",
					Certs:        nil,
				}
				keeper.SetApprovedCertificates(ctx, nilCert)
				return nilCert, nilCert.Subject, nilCert.SubjectKeyId
			},
			operation: func(cert types.ApprovedCertificates, subject, subjectKeyId string) {
				// No operation needed for this test
			},
			verify: func(subject, subjectKeyId string) (types.ApprovedCertificates, bool) {
				return keeper.GetApprovedCertificates(ctx, subject, subjectKeyId)
			},
			expectFound: true,
			description: "Test with nil certificates",
		},
		{
			name: "UpdateCertificates",
			setup: func() (types.ApprovedCertificates, string, string) {
				initialCert := types.ApprovedCertificates{
					Subject:      "update-subject",
					SubjectKeyId: "update-key-id",
					Certs: []*types.CertificateIdentifier{
						{
							Subject:      "update-subject",
							SubjectKeyId: "update-key-id",
						},
					},
				}
				keeper.SetApprovedCertificates(ctx, initialCert)
				
				updatedCert := types.ApprovedCertificates{
					Subject:      "update-subject",
					SubjectKeyId: "update-key-id",
					Certs: []*types.CertificateIdentifier{
						{
							Subject:      "update-subject",
							SubjectKeyId: "update-key-id",
						},
						{
							Subject:      "update-subject-2",
							SubjectKeyId: "update-key-id-2",
						},
					},
				}
				keeper.SetApprovedCertificates(ctx, updatedCert)
				return updatedCert, updatedCert.Subject, updatedCert.SubjectKeyId
			},
			operation: func(cert types.ApprovedCertificates, subject, subjectKeyId string) {
				// Update already done in setup
			},
			verify: func(subject, subjectKeyId string) (types.ApprovedCertificates, bool) {
				return keeper.GetApprovedCertificates(ctx, subject, subjectKeyId)
			},
			expectFound: true,
			description: "Test certificate updates",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cert, key1, key2 := tt.setup()
			tt.operation(cert, key1, key2)
			retrieved, found := tt.verify(key1, key2)
			
			if tt.expectFound {
				require.True(t, found)
				if tt.name == "UpdateCertificates" {
					require.Len(t, retrieved.Certs, 2)
				}
			} else {
				require.False(t, found)
			}
		})
	}
}

func TestApprovedCertificates_NonExistent(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	
	tests := []struct {
		name        string
		subject     string
		subjectKeyId string
		operation   func(string, string) (types.ApprovedCertificates, bool)
	}{
		{
			name:        "GetApprovedCertificates",
			subject:     "non-existent",
			subjectKeyId: "non-existent",
			operation:   func(subject, subjectKeyId string) (types.ApprovedCertificates, bool) {
				return keeper.GetApprovedCertificates(ctx, subject, subjectKeyId)
			},
		},
		{
			name:        "GetApprovedCertificatesBySubjectKeyID",
			subject:     "",
			subjectKeyId: "non-existent",
			operation:   func(_, subjectKeyId string) (types.ApprovedCertificates, bool) {
				return keeper.GetApprovedCertificatesBySubjectKeyID(ctx, subjectKeyId)
			},
		},
		{
			name:        "GetApprovedCertificatesBySubject",
			subject:     "non-existent",
			subjectKeyId: "",
			operation:   func(subject, _ string) (types.ApprovedCertificates, bool) {
				return keeper.GetApprovedCertificatesBySubject(ctx, subject)
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, found := tt.operation(tt.subject, tt.subjectKeyId)
			require.False(t, found)
		})
	}
}
