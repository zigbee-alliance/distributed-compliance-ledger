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

func TestNocCertificatesBySubjectKeyID(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)

	// Test setting and getting NOC certificates by subject key ID
	subjectKeyID := "test-noc-key-id"
	nocCerts := types.NocCertificatesBySubjectKeyID{
		SubjectKeyId: subjectKeyID,
		Certs: []*types.Certificate{
			{
				Subject:      "test-noc-subject",
				SubjectKeyId: subjectKeyID,
			},
		},
	}

	keeper.SetNocCertificatesBySubjectKeyID(ctx, nocCerts)

	// Test getting the certificate
	retrieved, found := keeper.GetNocCertificatesBySubjectKeyID(ctx, subjectKeyID)
	require.True(t, found)
	require.Equal(t, nocCerts, retrieved)

	// Test removing the certificate
	keeper.RemoveNocCertificatesBySubjectKeyID(ctx, subjectKeyID)
	_, found = keeper.GetNocCertificatesBySubjectKeyID(ctx, subjectKeyID)
	require.False(t, found)
}

func TestNocCertificates_EdgeCases(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)

	// Test with empty certificates
	emptyCert := types.NocCertificates{
		Subject:      "empty-noc-subject",
		SubjectKeyId: "empty-noc-key-id",
		Certs:        []*types.Certificate(nil),
	}

	keeper.SetNocCertificates(ctx, emptyCert)
	retrieved, found := keeper.GetNocCertificates(ctx, emptyCert.Subject, emptyCert.SubjectKeyId)
	require.True(t, found)
	require.Equal(t, emptyCert, retrieved)

	// Test with nil certificates
	nilCert := types.NocCertificates{
		Subject:      "nil-noc-subject",
		SubjectKeyId: "nil-noc-key-id",
		Certs:        nil,
	}

	keeper.SetNocCertificates(ctx, nilCert)
	retrieved, found = keeper.GetNocCertificates(ctx, nilCert.Subject, nilCert.SubjectKeyId)
	require.True(t, found)
	require.Equal(t, nilCert, retrieved)
}

func TestNocCertificates_Update(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)

	// Create certificate with multiple certs
	cert := types.NocCertificates{
		Subject:      "update-noc-subject",
		SubjectKeyId: "update-noc-key-id",
		Certs: []*types.Certificate{
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

	keeper.SetNocCertificates(ctx, cert)

	// Verify the certificate
	retrieved, found := keeper.GetNocCertificates(ctx, cert.Subject, cert.SubjectKeyId)
	require.True(t, found)
	require.Equal(t, cert, retrieved)
	require.Len(t, retrieved.Certs, 2)
}

func TestNocCertificates_MultipleOperations(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)

	// Test multiple operations in sequence
	cert1 := types.NocCertificates{
		Subject:      "noc-issuer1",
		SubjectKeyId: "noc-serial1",
		Certs: []*types.Certificate{
			{
				Subject:      "noc-issuer1",
				SubjectKeyId: "noc-serial1",
			},
		},
	}

	cert2 := types.NocCertificates{
		Subject:      "noc-issuer2",
		SubjectKeyId: "noc-serial2",
		Certs: []*types.Certificate{
			{
				Subject:      "noc-issuer2",
				SubjectKeyId: "noc-serial2",
			},
		},
	}

	// Set both certificates
	keeper.SetNocCertificates(ctx, cert1)
	keeper.SetNocCertificates(ctx, cert2)

	// Verify both exist
	retrieved1, found1 := keeper.GetNocCertificates(ctx, cert1.Subject, cert1.SubjectKeyId)
	require.True(t, found1)
	require.Equal(t, cert1, retrieved1)

	retrieved2, found2 := keeper.GetNocCertificates(ctx, cert2.Subject, cert2.SubjectKeyId)
	require.True(t, found2)
	require.Equal(t, cert2, retrieved2)

	// Remove one certificate
	keeper.RemoveNocCertificates(ctx, cert1.Subject, cert1.SubjectKeyId)

	// Verify first is removed, second still exists
	_, found1 = keeper.GetNocCertificates(ctx, cert1.Subject, cert1.SubjectKeyId)
	require.False(t, found1)

	retrieved2, found2 = keeper.GetNocCertificates(ctx, cert2.Subject, cert2.SubjectKeyId)
	require.True(t, found2)
	require.Equal(t, cert2, retrieved2)
}

func TestNocCertificates_NonExistent(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)

	// Test getting non-existent NOC certificates
	_, found := keeper.GetNocCertificates(ctx, "non-existent", "non-existent")
	require.False(t, found)

	// Test getting non-existent NOC certificates by subject key ID
	_, found = keeper.GetNocCertificatesBySubjectKeyID(ctx, "non-existent")
	require.False(t, found)
}
