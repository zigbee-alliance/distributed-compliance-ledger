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

func TestApprovedCertificatesBySubjectKeyID(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	
	// Test setting and getting approved certificates by subject key ID
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
	
	// Test getting the certificate
	retrieved, found := keeper.GetApprovedCertificatesBySubjectKeyID(ctx, subjectKeyID)
	require.True(t, found)
	require.Equal(t, approvedCerts, retrieved)
	
	// Test removing the certificate
	keeper.RemoveApprovedCertificatesBySubjectKeyID(ctx, subjectKeyID)
	_, found = keeper.GetApprovedCertificatesBySubjectKeyID(ctx, subjectKeyID)
	require.False(t, found)
}

func TestApprovedCertificatesBySubject(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	
	// Test setting and getting approved certificates by subject
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
	
	// Test getting the certificate
	retrieved, found := keeper.GetApprovedCertificatesBySubject(ctx, subject)
	require.True(t, found)
	require.Equal(t, approvedCerts, retrieved)
	
	// Test removing the certificate
	keeper.RemoveApprovedCertificatesBySubject(ctx, subject)
	_, found = keeper.GetApprovedCertificatesBySubject(ctx, subject)
	require.False(t, found)
}

func TestApprovedCertificates_EdgeCases(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	
	// Test with empty certificates
	emptyCert := types.ApprovedCertificates{
		Subject:      "empty-subject",
		SubjectKeyId: "empty-key-id",
		Certs:        []*types.CertificateIdentifier{},
	}
	
	keeper.SetApprovedCertificates(ctx, emptyCert)
	retrieved, found := keeper.GetApprovedCertificates(ctx, emptyCert.Subject, emptyCert.SubjectKeyId)
	require.True(t, found)
	require.Equal(t, emptyCert, retrieved)
	
	// Test with nil certificates
	nilCert := types.ApprovedCertificates{
		Subject:      "nil-subject",
		SubjectKeyId: "nil-key-id",
		Certs:        nil,
	}
	
	keeper.SetApprovedCertificates(ctx, nilCert)
	retrieved, found = keeper.GetApprovedCertificates(ctx, nilCert.Subject, nilCert.SubjectKeyId)
	require.True(t, found)
	require.Equal(t, nilCert, retrieved)
}

func TestApprovedCertificates_NonExistent(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	
	// Test getting non-existent certificates
	_, found := keeper.GetApprovedCertificates(ctx, "non-existent", "non-existent")
	require.False(t, found)
	
	_, found = keeper.GetApprovedCertificatesBySubjectKeyID(ctx, "non-existent")
	require.False(t, found)
	
	_, found = keeper.GetApprovedCertificatesBySubject(ctx, "non-existent")
	require.False(t, found)
}

func TestApprovedCertificates_Update(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	
	// Create initial certificate
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
	
	// Update the certificate
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
	
	// Verify the update
	retrieved, found := keeper.GetApprovedCertificates(ctx, updatedCert.Subject, updatedCert.SubjectKeyId)
	require.True(t, found)
	require.Equal(t, updatedCert, retrieved)
	require.Len(t, retrieved.Certs, 2)
}
