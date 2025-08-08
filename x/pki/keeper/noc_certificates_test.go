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

// Prevent strconv unused error
var _ = strconv.IntSize

func createNNocCertificates(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.NocCertificates {
	items := make([]types.NocCertificates, n)
	for i := range items {
		items[i].Subject = strconv.Itoa(i)
		items[i].SubjectKeyId = strconv.Itoa(i)

		keeper.SetNocCertificates(ctx, items[i])
		keeper.SetNocCertificatesBySubjectKeyID(ctx, types.NocCertificatesBySubjectKeyID{
			SubjectKeyId: items[i].SubjectKeyId,
			Certs:        items[i].Certs,
		})
	}

	return items
}

func TestNocCertificatesGet(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNNocCertificates(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetNocCertificates(ctx,
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
func TestNocCertificatesRemove(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNNocCertificates(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNocCertificates(ctx,
			item.Subject,
			item.SubjectKeyId,
		)
		_, found := keeper.GetNocCertificates(ctx,
			item.Subject,
			item.SubjectKeyId,
		)
		require.False(t, found)
	}
}

func TestNocCertificatesGetAll(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNNocCertificates(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllNocCertificates(ctx)),
	)
}

func TestNocCertificates_ExtendedOperations(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	
	tests := []struct {
		name        string
		setup       func() (types.NocCertificates, string, string)
		operation   func(types.NocCertificates, string, string)
		verify      func(string, string) (types.NocCertificates, bool)
		expectFound bool
		description string
	}{
		{
			name: "BySubjectKeyID",
			setup: func() (types.NocCertificates, string, string) {
				subjectKeyID := "test-noc-key-id"
				nocCerts := types.NocCertificatesBySubjectKeyID{
					SubjectKeyId: subjectKeyID,
					Certs: []*types.CertificateIdentifier{
						{
							Subject:      "test-noc-subject",
							SubjectKeyId: subjectKeyID,
						},
					},
				}
				keeper.SetNocCertificatesBySubjectKeyID(ctx, nocCerts)
				return types.NocCertificates{}, subjectKeyID, ""
			},
			operation: func(cert types.NocCertificates, subjectKeyID, _ string) {
				keeper.RemoveNocCertificatesBySubjectKeyID(ctx, subjectKeyID)
			},
			verify: func(subjectKeyID, _ string) (types.NocCertificates, bool) {
				return keeper.GetNocCertificatesBySubjectKeyID(ctx, subjectKeyID)
			},
			expectFound: false,
			description: "Test setting and getting NOC certificates by subject key ID",
		},
		{
			name: "EmptyCertificates",
			setup: func() (types.NocCertificates, string, string) {
				emptyCert := types.NocCertificates{
					Subject:      "empty-noc-subject",
					SubjectKeyId: "empty-noc-key-id",
					Certs:        []*types.CertificateIdentifier{},
				}
				keeper.SetNocCertificates(ctx, emptyCert)
				return emptyCert, emptyCert.Subject, emptyCert.SubjectKeyId
			},
			operation: func(cert types.NocCertificates, subject, subjectKeyId string) {
				// No operation needed for this test
			},
			verify: func(subject, subjectKeyId string) (types.NocCertificates, bool) {
				return keeper.GetNocCertificates(ctx, subject, subjectKeyId)
			},
			expectFound: true,
			description: "Test with empty certificates",
		},
		{
			name: "NilCertificates",
			setup: func() (types.NocCertificates, string, string) {
				nilCert := types.NocCertificates{
					Subject:      "nil-noc-subject",
					SubjectKeyId: "nil-noc-key-id",
					Certs:        nil,
				}
				keeper.SetNocCertificates(ctx, nilCert)
				return nilCert, nilCert.Subject, nilCert.SubjectKeyId
			},
			operation: func(cert types.NocCertificates, subject, subjectKeyId string) {
				// No operation needed for this test
			},
			verify: func(subject, subjectKeyId string) (types.NocCertificates, bool) {
				return keeper.GetNocCertificates(ctx, subject, subjectKeyId)
			},
			expectFound: true,
			description: "Test with nil certificates",
		},
		{
			name: "UpdateCertificates",
			setup: func() (types.NocCertificates, string, string) {
				initialCert := types.NocCertificates{
					Subject:      "update-noc-subject",
					SubjectKeyId: "update-noc-key-id",
					Certs: []*types.CertificateIdentifier{
						{
							Subject:      "update-noc-subject",
							SubjectKeyId: "update-noc-key-id",
						},
					},
				}
				keeper.SetNocCertificates(ctx, initialCert)
				
				updatedCert := types.NocCertificates{
					Subject:      "update-noc-subject",
					SubjectKeyId: "update-noc-key-id",
					Certs: []*types.CertificateIdentifier{
						{
							Subject:      "update-noc-subject",
							SubjectKeyId: "update-noc-key-id",
						},
						{
							Subject:      "update-noc-subject-2",
							SubjectKeyId: "update-noc-key-id-2",
						},
					},
				}
				keeper.SetNocCertificates(ctx, updatedCert)
				return updatedCert, updatedCert.Subject, updatedCert.SubjectKeyId
			},
			operation: func(cert types.NocCertificates, subject, subjectKeyId string) {
				// Update already done in setup
			},
			verify: func(subject, subjectKeyId string) (types.NocCertificates, bool) {
				return keeper.GetNocCertificates(ctx, subject, subjectKeyId)
			},
			expectFound: true,
			description: "Test certificate updates",
		},
		{
			name: "MultipleOperations",
			setup: func() (types.NocCertificates, string, string) {
				cert1 := types.NocCertificates{
					Subject:      "noc-issuer1",
					SubjectKeyId: "noc-serial1",
					Certs: []*types.CertificateIdentifier{
						{
							Subject:      "noc-issuer1",
							SubjectKeyId: "noc-serial1",
						},
					},
				}
				cert2 := types.NocCertificates{
					Subject:      "noc-issuer2",
					SubjectKeyId: "noc-serial2",
					Certs: []*types.CertificateIdentifier{
						{
							Subject:      "noc-issuer2",
							SubjectKeyId: "noc-serial2",
						},
					},
				}
				keeper.SetNocCertificates(ctx, cert1)
				keeper.SetNocCertificates(ctx, cert2)
				return cert1, cert1.Subject, cert1.SubjectKeyId
			},
			operation: func(cert types.NocCertificates, subject, subjectKeyId string) {
				keeper.RemoveNocCertificates(ctx, subject, subjectKeyId)
			},
			verify: func(subject, subjectKeyId string) (types.NocCertificates, bool) {
				return keeper.GetNocCertificates(ctx, subject, subjectKeyId)
			},
			expectFound: false,
			description: "Test multiple operations in sequence",
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

func TestNocCertificates_NonExistent(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	
	tests := []struct {
		name        string
		subject     string
		subjectKeyId string
		operation   func(string, string) (types.NocCertificates, bool)
	}{
		{
			name:        "GetNocCertificates",
			subject:     "non-existent",
			subjectKeyId: "non-existent",
			operation:   func(subject, subjectKeyId string) (types.NocCertificates, bool) {
				return keeper.GetNocCertificates(ctx, subject, subjectKeyId)
			},
		},
		{
			name:        "GetNocCertificatesBySubjectKeyID",
			subject:     "",
			subjectKeyId: "non-existent",
			operation:   func(_, subjectKeyId string) (types.NocCertificates, bool) {
				return keeper.GetNocCertificatesBySubjectKeyID(ctx, subjectKeyId)
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
