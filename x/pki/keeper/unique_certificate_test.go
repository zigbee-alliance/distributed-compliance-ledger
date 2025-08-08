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

func createNUniqueCertificate(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.UniqueCertificate {
	items := make([]types.UniqueCertificate, n)
	for i := range items {
		items[i].Issuer = strconv.Itoa(i)
		items[i].SerialNumber = strconv.Itoa(i)

		keeper.SetUniqueCertificate(ctx, items[i])
	}

	return items
}

func TestUniqueCertificateGet(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNUniqueCertificate(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetUniqueCertificate(ctx,
			item.Issuer,
			item.SerialNumber,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestUniqueCertificateRemove(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNUniqueCertificate(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveUniqueCertificate(ctx,
			item.Issuer,
			item.SerialNumber,
		)
		_, found := keeper.GetUniqueCertificate(ctx,
			item.Issuer,
			item.SerialNumber,
		)
		require.False(t, found)
	}
}

func TestUniqueCertificateGetAll(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	items := createNUniqueCertificate(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllUniqueCertificate(ctx)),
	)
}

func TestUniqueCertificate_EdgeCases(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	
	// Test with empty issuer and serial number
	emptyCert := types.UniqueCertificate{
		Issuer:       "",
		SerialNumber: "",
	}
	
	keeper.SetUniqueCertificate(ctx, emptyCert)
	retrieved, found := keeper.GetUniqueCertificate(ctx, emptyCert.Issuer, emptyCert.SerialNumber)
	require.True(t, found)
	require.Equal(t, emptyCert, retrieved)
	
	// Test with special characters
	specialCert := types.UniqueCertificate{
		Issuer:       "test@issuer.com",
		SerialNumber: "123-456-789",
	}
	
	keeper.SetUniqueCertificate(ctx, specialCert)
	retrieved, found = keeper.GetUniqueCertificate(ctx, specialCert.Issuer, specialCert.SerialNumber)
	require.True(t, found)
	require.Equal(t, specialCert, retrieved)
}

func TestUniqueCertificate_NonExistent(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	
	// Test getting non-existent certificate
	_, found := keeper.GetUniqueCertificate(ctx, "non-existent", "non-existent")
	require.False(t, found)
}

func TestUniqueCertificate_Update(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	
	// Create initial certificate
	initialCert := types.UniqueCertificate{
		Issuer:       "test-issuer",
		SerialNumber: "test-serial",
	}
	
	keeper.SetUniqueCertificate(ctx, initialCert)
	
	// Update the certificate (same key, different data)
	updatedCert := types.UniqueCertificate{
		Issuer:       "test-issuer",
		SerialNumber: "test-serial",
		// Add some additional data if the struct has more fields
	}
	
	keeper.SetUniqueCertificate(ctx, updatedCert)
	
	// Verify the update
	retrieved, found := keeper.GetUniqueCertificate(ctx, updatedCert.Issuer, updatedCert.SerialNumber)
	require.True(t, found)
	require.Equal(t, updatedCert, retrieved)
}

func TestUniqueCertificate_MultipleOperations(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	
	// Test multiple operations in sequence
	cert1 := types.UniqueCertificate{
		Issuer:       "issuer1",
		SerialNumber: "serial1",
	}
	
	cert2 := types.UniqueCertificate{
		Issuer:       "issuer2",
		SerialNumber: "serial2",
	}
	
	// Set both certificates
	keeper.SetUniqueCertificate(ctx, cert1)
	keeper.SetUniqueCertificate(ctx, cert2)
	
	// Verify both exist
	retrieved1, found1 := keeper.GetUniqueCertificate(ctx, cert1.Issuer, cert1.SerialNumber)
	require.True(t, found1)
	require.Equal(t, cert1, retrieved1)
	
	retrieved2, found2 := keeper.GetUniqueCertificate(ctx, cert2.Issuer, cert2.SerialNumber)
	require.True(t, found2)
	require.Equal(t, cert2, retrieved2)
	
	// Remove one certificate
	keeper.RemoveUniqueCertificate(ctx, cert1.Issuer, cert1.SerialNumber)
	
	// Verify first is removed, second still exists
	_, found1 = keeper.GetUniqueCertificate(ctx, cert1.Issuer, cert1.SerialNumber)
	require.False(t, found1)
	
	retrieved2, found2 = keeper.GetUniqueCertificate(ctx, cert2.Issuer, cert2.SerialNumber)
	require.True(t, found2)
	require.Equal(t, cert2, retrieved2)
}

func TestUniqueCertificate_GetAllWithEmpty(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	
	// Test GetAll with no certificates
	allCerts := keeper.GetAllUniqueCertificate(ctx)
	require.Empty(t, allCerts)
	
	// Add one certificate and test GetAll
	cert := types.UniqueCertificate{
		Issuer:       "test-issuer",
		SerialNumber: "test-serial",
	}
	keeper.SetUniqueCertificate(ctx, cert)
	
	allCerts = keeper.GetAllUniqueCertificate(ctx)
	require.Len(t, allCerts, 1)
	require.Equal(t, cert, allCerts[0])
}
