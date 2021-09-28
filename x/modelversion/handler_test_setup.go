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

package modelversion

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/go-amino"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/modelversion/internal/types"
)

type TestSetup struct {
	Cdc                *amino.Codec
	Ctx                sdk.Context
	ModelVersionKeeper Keeper
	ModelKeeper        model.Keeper
	AuthKeeper         auth.Keeper
	Handler            sdk.Handler
	Querier            sdk.Querier
	Vendor             sdk.AccAddress
	VendorId           uint16
}

func Setup() TestSetup {
	// Init Codec
	cdc := codec.New()
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)

	// Init KVSore
	db := dbm.NewMemDB()

	dbStore := store.NewCommitMultiStore(db)

	modelKey := sdk.NewKVStoreKey(model.StoreKey)
	dbStore.MountStoreWithDB(modelKey, sdk.StoreTypeIAVL, nil)

	authKey := sdk.NewKVStoreKey(auth.StoreKey)
	dbStore.MountStoreWithDB(authKey, sdk.StoreTypeIAVL, nil)

	modelVersionKey := sdk.NewKVStoreKey(StoreKey)
	dbStore.MountStoreWithDB(modelVersionKey, sdk.StoreTypeIAVL, nil)

	_ = dbStore.LoadLatestVersion()

	// Init Keepers
	modelKeeper := model.NewKeeper(modelKey, cdc)
	authKeeper := auth.NewKeeper(authKey, cdc)
	modelVersionKeeper := NewKeeper(modelVersionKey, cdc)

	// Create context
	ctx := sdk.NewContext(dbStore, abci.Header{ChainID: testconstants.ChainID}, false, log.NewNopLogger())

	// Create Handler and Querier
	querier := NewQuerier(modelVersionKeeper)
	handler := NewHandler(modelVersionKeeper, authKeeper, modelKeeper)

	account := auth.NewAccount(testconstants.Address1, testconstants.PubKey1, auth.AccountRoles{auth.Vendor}, testconstants.VendorId1)
	account.AccountNumber = authKeeper.GetNextAccountNumber(ctx)
	authKeeper.SetAccount(ctx, account)

	setup := TestSetup{
		Cdc:                cdc,
		Ctx:                ctx,
		ModelKeeper:        modelKeeper,
		ModelVersionKeeper: modelVersionKeeper,
		AuthKeeper:         authKeeper,
		Handler:            handler,
		Querier:            querier,
		Vendor:             account.Address,
	}

	return setup
}

func getTestModelVersion() types.ModelVersion {
	return ModelVersion{
		VID:                          testconstants.VendorId1,
		PID:                          testconstants.PID,
		SoftwareVersion:              testconstants.SoftwareVersion,
		SoftwareVersionString:        testconstants.SoftwareVersionString,
		CDVersionNumber:              testconstants.CDVersionNumber,
		FirmwareDigests:              testconstants.FirmwareDigests,
		SoftwareVersionValid:         testconstants.SoftwareVersionValid,
		OtaURL:                       testconstants.OtaURL,
		OtaFileSize:                  testconstants.OtaFileSize,
		OtaChecksum:                  testconstants.OtaChecksum,
		OtaChecksumType:              testconstants.OtaChecksumType,
		MinApplicableSoftwareVersion: testconstants.MinApplicableSoftwareVersion,
		MaxApplicableSoftwareVersion: testconstants.MaxApplicableSoftwareVersion,
		ReleaseNotesURL:              testconstants.ReleaseNotesURL,
	}
}

func getTestModelVersionForUpdate() types.ModelVersion {
	return ModelVersion{
		VID:                   testconstants.VendorId1,
		PID:                   testconstants.PID,
		SoftwareVersion:       testconstants.SoftwareVersion,
		SoftwareVersionString: testconstants.SoftwareVersionString,
		ReleaseNotesURL:       testconstants.ReleaseNotesURL + "/updated",
	}
}

func TestMsgAddModelVersion(signer sdk.AccAddress) MsgAddModelVersion {
	return MsgAddModelVersion{
		ModelVersion: getTestModelVersion(),
		Signer:       signer,
	}
}

func TestMsgUpdateModelVersion(signer sdk.AccAddress) MsgUpdateModelVersion {
	return MsgUpdateModelVersion{
		ModelVersion: getTestModelVersionForUpdate(),
		Signer:       signer,
	}
}
