package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"

	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/keeper"
)

// Prevent strconv unused error.
var _ = strconv.IntSize

func _createNApprovedCertificates(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ApprovedCertificates {
	items := make([]types.ApprovedCertificates, n)
	for i := range items {
		items[i].Subject = strconv.Itoa(i)
		items[i].SubjectKeyId = strconv.Itoa(i)
		items[i].Certs = []*types.Certificate{{SubjectKeyId: strconv.Itoa(i)}}
		keeper.SetApprovedCertificates(ctx, items[i])
	}

	return items
}

func TestMigrator_Migrate2to3(t *testing.T) {
	_keeper, ctx := keepertest.PkiKeeper(t, nil)
	msg := _createNApprovedCertificates(_keeper, ctx, 5)

	migrator := keeper.NewMigrator(*_keeper)
	err := migrator.Migrate2to3(ctx)
	require.NoError(t, err)

	subjectKeyID := "0"
	list, found := _keeper.GetApprovedCertificatesBySubjectKeyID(ctx, subjectKeyID)
	require.True(t, found)

	require.Equal(t, 1, len(list.Certs))
	require.Equal(t, subjectKeyID, list.SubjectKeyId)
	require.Equal(t, msg[0].Certs, list.Certs)
}
