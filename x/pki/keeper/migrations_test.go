// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

func _createNApprovedCertificatesBySubject(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ApprovedCertificatesBySubject {
	items := make([]types.ApprovedCertificatesBySubject, n)
	for i := range items {
		items[i].Subject = strconv.Itoa(i)
		items[i].SubjectKeyIds = []string{strconv.Itoa(i)}
		keeper.SetApprovedCertificatesBySubject(ctx, items[i])
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

func TestMigrator_Migrate3to4(t *testing.T) {
	_keeper, ctx := keepertest.PkiKeeper(t, nil)
	msg := _createNApprovedCertificates(_keeper, ctx, 5)
	_createNApprovedCertificatesBySubject(_keeper, ctx, 5)

	migrator := keeper.NewMigrator(*_keeper)
	err := migrator.Migrate3to4(ctx)
	require.NoError(t, err)

	// check that all certificates migrated
	subject := "0"
	subjectKeyID := "0"
	list, found := _keeper.GetAllCertificates(ctx, subject, subjectKeyID)
	require.True(t, found)
	require.Equal(t, 1, len(list.Certs))
	require.Equal(t, subjectKeyID, list.SubjectKeyId)
	require.Equal(t, msg[0].Certs, list.Certs)

	allList := _keeper.GetAllAllCertificates(ctx)
	require.Equal(t, 5, len(allList))

	// check that all certificates by subject migrated
	subjList, found := _keeper.GetAllCertificatesBySubject(ctx, subject)
	require.True(t, found)
	require.Equal(t, subject, subjList.Subject)
	require.Equal(t, 1, len(subjList.SubjectKeyIds))

	// check that all certificates by subject key id migrated
	certificatesBySubjectKeyId, found := _keeper.GetAllCertificatesBySubjectKeyID(ctx, subjectKeyID)
	require.True(t, found)
	require.Equal(t, 1, len(certificatesBySubjectKeyId.Certs))
	require.Equal(t, subjectKeyID, certificatesBySubjectKeyId.SubjectKeyId)
	require.Equal(t, msg[0].Certs, certificatesBySubjectKeyId.Certs)
}
