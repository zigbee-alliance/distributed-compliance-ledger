package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/nullify"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func createTestRevokedRootCertificates(keeper *keeper.Keeper, ctx sdk.Context) types.RevokedRootCertificates { //nolint:unparam
	item := types.RevokedRootCertificates{}
	keeper.SetRevokedRootCertificates(ctx, item)

	return item
}

func TestRevokedRootCertificatesGet(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	item := createTestRevokedRootCertificates(keeper, ctx)
	rst, found := keeper.GetRevokedRootCertificates(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestRevokedRootCertificatesRemove(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	createTestRevokedRootCertificates(keeper, ctx)
	keeper.RemoveRevokedRootCertificates(ctx)
	_, found := keeper.GetRevokedRootCertificates(ctx)
	require.False(t, found)
}

func TestRevokedRootCertificateAddRemove(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)

	certID := types.CertificateIdentifier{Subject: "subj", SubjectKeyId: "skid"}
	other := types.CertificateIdentifier{Subject: "subj2", SubjectKeyId: "skid2"}

	keeper.AddRevokedRootCertificate(ctx, certID)
	keeper.AddRevokedRootCertificate(ctx, certID) // duplicate is ignored
	keeper.AddRevokedRootCertificate(ctx, other)

	stored, found := keeper.GetRevokedRootCertificates(ctx)
	require.True(t, found)
	require.Len(t, stored.Certs, 2)

	// Removing an absent identifier is a no-op.
	keeper.RemoveRevokedRootCertificate(ctx, types.CertificateIdentifier{Subject: "nope", SubjectKeyId: "nope"})
	stored, _ = keeper.GetRevokedRootCertificates(ctx)
	require.Len(t, stored.Certs, 2)

	keeper.RemoveRevokedRootCertificate(ctx, certID)
	stored, _ = keeper.GetRevokedRootCertificates(ctx)
	require.Len(t, stored.Certs, 1)
	require.Equal(t, other, *stored.Certs[0])
}
