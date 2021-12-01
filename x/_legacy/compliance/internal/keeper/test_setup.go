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

package keeper

//nolint:goimports
import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/internal/types"
)

type TestSetup struct {
	Cdc               *codec.Codec
	Ctx               sdk.Context
	CompliancetKeeper Keeper
	Querier           sdk.Querier
}

func Setup() TestSetup {
	// Init Codec
	cdc := codec.New()
	sdk.RegisterCodec(cdc)

	// Init KVSore
	db := dbm.NewMemDB()
	dbStore := store.NewCommitMultiStore(db)
	complianceKey := sdk.NewKVStoreKey(types.StoreKey)
	dbStore.MountStoreWithDB(complianceKey, sdk.StoreTypeIAVL, nil)
	_ = dbStore.LoadLatestVersion()

	// Init Keepers
	complianceKeeper := NewKeeper(complianceKey, cdc)

	// Init Querier
	querier := NewQuerier(complianceKeeper)

	// Create context
	ctx := sdk.NewContext(dbStore, abci.Header{ChainID: testconstants.ChainID}, false, log.NewNopLogger())

	setup := TestSetup{
		Cdc:               cdc,
		Ctx:               ctx,
		CompliancetKeeper: complianceKeeper,
		Querier:           querier,
	}

	return setup
}

func DefaultCertifiedModel() types.ComplianceInfo {
	return types.NewCertifiedComplianceInfo(
		testconstants.VID,
		testconstants.PID,
		testconstants.SoftwareVersion,
		testconstants.SoftwareVersionString,
		types.CertificationType(testconstants.CertificationType),
		testconstants.CertificationDate,
		testconstants.EmptyString,
		testconstants.Owner,
	)
}

func DefaultRevokedModel() types.ComplianceInfo {
	return types.NewRevokedComplianceInfo(
		testconstants.VID,
		testconstants.PID,
		testconstants.SoftwareVersion,
		testconstants.SoftwareVersionString,
		types.CertificationType(testconstants.CertificationType),
		testconstants.RevocationDate,
		testconstants.RevocationReason,
		testconstants.Owner,
	)
}

// add n=count/2 certified and count-n revoked models {VID: 1, PID: 1..count}.
func PopulateStoreWithMixedModels(setup TestSetup, count int) uint16 {
	firstID := uint16(1)
	n := count / 2

	certifiedModel := DefaultCertifiedModel()
	PopulateStoreWithModels(setup, firstID, n, certifiedModel)

	revokedModel := DefaultRevokedModel()
	PopulateStoreWithModels(setup, firstID+uint16(n), count, revokedModel)

	return firstID
}

// add n models {VID: 1, PID: 1..count}.
func PopulateStoreWithModels(setup TestSetup, firstID uint16, count int, complianceInfo types.ComplianceInfo) uint16 {
	for i := firstID; i <= uint16(count); i++ {
		// add model {VID: 1, PID: i}
		complianceInfo.PID = i
		complianceInfo.VID = i
		complianceInfo.SoftwareVersion = uint32(i)

		setup.CompliancetKeeper.SetComplianceInfo(setup.Ctx, complianceInfo)
	}

	return firstID
}

func CheckComplianceInfo(t *testing.T, expected types.ComplianceInfo, received types.ComplianceInfo) {
	require.Equal(t, expected.VID, received.VID)
	require.Equal(t, expected.PID, received.PID)
	require.Equal(t, expected.SoftwareVersion, received.SoftwareVersion)
	require.Equal(t, expected.SoftwareVersionString, received.SoftwareVersionString)
	require.Equal(t, expected.CertificationType, received.CertificationType)
	require.Equal(t, expected.Date, received.Date)
	require.Equal(t, expected.Reason, received.Reason)
}
