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

func TestUniqueCertificate_ExtendedOperations(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	
	tests := []struct {
		name        string
		setup       func() (types.UniqueCertificate, string, string)
		operation   func(types.UniqueCertificate, string, string)
		verify      func(string, string) (types.UniqueCertificate, bool)
		expectFound bool
		description string
	}{
		{
			name: "EmptyIssuerAndSerial",
			setup: func() (types.UniqueCertificate, string, string) {
				emptyCert := types.UniqueCertificate{
					Issuer:       "",
					SerialNumber: "",
				}
				keeper.SetUniqueCertificate(ctx, emptyCert)
				return emptyCert, emptyCert.Issuer, emptyCert.SerialNumber
			},
			operation: func(cert types.UniqueCertificate, issuer, serial string) {
				// No operation needed for this test
			},
			verify: func(issuer, serial string) (types.UniqueCertificate, bool) {
				return keeper.GetUniqueCertificate(ctx, issuer, serial)
			},
			expectFound: true,
			description: "Test with empty issuer and serial number",
		},
		{
			name: "SpecialCharacters",
			setup: func() (types.UniqueCertificate, string, string) {
				specialCert := types.UniqueCertificate{
					Issuer:       "test@issuer.com",
					SerialNumber: "123-456-789",
				}
				keeper.SetUniqueCertificate(ctx, specialCert)
				return specialCert, specialCert.Issuer, specialCert.SerialNumber
			},
			operation: func(cert types.UniqueCertificate, issuer, serial string) {
				// No operation needed for this test
			},
			verify: func(issuer, serial string) (types.UniqueCertificate, bool) {
				return keeper.GetUniqueCertificate(ctx, issuer, serial)
			},
			expectFound: true,
			description: "Test with special characters",
		},
		{
			name: "UpdateCertificate",
			setup: func() (types.UniqueCertificate, string, string) {
				initialCert := types.UniqueCertificate{
					Issuer:       "test-issuer",
					SerialNumber: "test-serial",
				}
				keeper.SetUniqueCertificate(ctx, initialCert)
				
				updatedCert := types.UniqueCertificate{
					Issuer:       "test-issuer",
					SerialNumber: "test-serial",
					// Add some additional data if the struct has more fields
				}
				keeper.SetUniqueCertificate(ctx, updatedCert)
				return updatedCert, updatedCert.Issuer, updatedCert.SerialNumber
			},
			operation: func(cert types.UniqueCertificate, issuer, serial string) {
				// Update already done in setup
			},
			verify: func(issuer, serial string) (types.UniqueCertificate, bool) {
				return keeper.GetUniqueCertificate(ctx, issuer, serial)
			},
			expectFound: true,
			description: "Test certificate updates",
		},
		{
			name: "MultipleOperations",
			setup: func() (types.UniqueCertificate, string, string) {
				cert1 := types.UniqueCertificate{
					Issuer:       "issuer1",
					SerialNumber: "serial1",
				}
				cert2 := types.UniqueCertificate{
					Issuer:       "issuer2",
					SerialNumber: "serial2",
				}
				keeper.SetUniqueCertificate(ctx, cert1)
				keeper.SetUniqueCertificate(ctx, cert2)
				return cert1, cert1.Issuer, cert1.SerialNumber
			},
			operation: func(cert types.UniqueCertificate, issuer, serial string) {
				keeper.RemoveUniqueCertificate(ctx, issuer, serial)
			},
			verify: func(issuer, serial string) (types.UniqueCertificate, bool) {
				return keeper.GetUniqueCertificate(ctx, issuer, serial)
			},
			expectFound: false,
			description: "Test multiple operations in sequence",
		},
		{
			name: "GetAllWithEmpty",
			setup: func() (types.UniqueCertificate, string, string) {
				// Clear any existing certificates
				allCerts := keeper.GetAllUniqueCertificate(ctx)
				for _, cert := range allCerts {
					keeper.RemoveUniqueCertificate(ctx, cert.Issuer, cert.SerialNumber)
				}
				
				cert := types.UniqueCertificate{
					Issuer:       "test-issuer",
					SerialNumber: "test-serial",
				}
				keeper.SetUniqueCertificate(ctx, cert)
				return cert, cert.Issuer, cert.SerialNumber
			},
			operation: func(cert types.UniqueCertificate, issuer, serial string) {
				// Verify GetAll works
				allCerts := keeper.GetAllUniqueCertificate(ctx)
				require.Len(t, allCerts, 1)
				require.Equal(t, cert, allCerts[0])
			},
			verify: func(issuer, serial string) (types.UniqueCertificate, bool) {
				return keeper.GetUniqueCertificate(ctx, issuer, serial)
			},
			expectFound: true,
			description: "Test GetAll with single certificate",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cert, key1, key2 := tt.setup()
			tt.operation(cert, key1, key2)
			retrieved, found := tt.verify(key1, key2)
			
			if tt.expectFound {
				require.True(t, found)
				require.Equal(t, cert, retrieved)
			} else {
				require.False(t, found)
			}
		})
	}
}

func TestUniqueCertificate_NonExistent(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	
	// Test getting non-existent certificate
	_, found := keeper.GetUniqueCertificate(ctx, "non-existent", "non-existent")
	require.False(t, found)
}
