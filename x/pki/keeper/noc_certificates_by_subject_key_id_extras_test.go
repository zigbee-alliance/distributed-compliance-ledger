package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func TestAddGetRemoveNocCertificatesBySubjectKeyID(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)

	keeper.AddNocCertificatesBySubjectKeyID(ctx, types.NocCertificates{
		SubjectKeyId: "skid",
		Certs: []*types.Certificate{
			{Subject: "s1", SubjectKeyId: "skid"},
		},
	})

	got, found := keeper.GetNocCertificatesBySubjectKeyID(ctx, "skid")
	require.True(t, found)
	require.Len(t, got.Certs, 1)
	require.Equal(t, "s1", got.Certs[0].Subject)

	keeper.RemoveNocCertificatesBySubjectAndSubjectKeyID(ctx, "s1", "skid")

	_, found = keeper.GetNocCertificatesBySubjectKeyID(ctx, "skid")
	require.False(t, found)
}
