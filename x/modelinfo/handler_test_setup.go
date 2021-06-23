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

package modelinfo

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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/modelinfo/internal/types"
)

type TestSetup struct {
	Cdc             *amino.Codec
	Ctx             sdk.Context
	ModelinfoKeeper Keeper
	authKeeper      auth.Keeper
	Handler         sdk.Handler
	Querier         sdk.Querier
	Vendor          sdk.AccAddress
}

func Setup() TestSetup {
	// Init Codec
	cdc := codec.New()
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)

	// Init KVSore
	db := dbm.NewMemDB()

	dbStore := store.NewCommitMultiStore(db)

	modelinfoKey := sdk.NewKVStoreKey(StoreKey)
	dbStore.MountStoreWithDB(modelinfoKey, sdk.StoreTypeIAVL, nil)

	authKey := sdk.NewKVStoreKey(auth.StoreKey)
	dbStore.MountStoreWithDB(authKey, sdk.StoreTypeIAVL, nil)

	_ = dbStore.LoadLatestVersion()

	// Init Keepers
	modelinfoKeeper := NewKeeper(modelinfoKey, cdc)
	authKeeper := auth.NewKeeper(authKey, cdc)

	// Create context
	ctx := sdk.NewContext(dbStore, abci.Header{ChainID: testconstants.ChainID}, false, log.NewNopLogger())

	// Create Handler and Querier
	querier := NewQuerier(modelinfoKeeper)
	handler := NewHandler(modelinfoKeeper, authKeeper)

	account := auth.NewAccount(testconstants.Address1, testconstants.PubKey1, auth.AccountRoles{auth.Vendor})
	account.AccountNumber = authKeeper.GetNextAccountNumber(ctx)
	authKeeper.SetAccount(ctx, account)

	setup := TestSetup{
		Cdc:             cdc,
		Ctx:             ctx,
		ModelinfoKeeper: modelinfoKeeper,
		authKeeper:      authKeeper,
		Handler:         handler,
		Querier:         querier,
		Vendor:          account.Address,
	}

	return setup
}

func getTestModel() types.Model {
	return Model{
		VID:                                      testconstants.VID,
		PID:                                      testconstants.PID,
		CID:                                      testconstants.CID,
		Name:                                     testconstants.Name,
		Description:                              testconstants.Description,
		SKU:                                      testconstants.SKU,
		SoftwareVersion:                          testconstants.SoftwareVersion,
		SoftwareVersionString:                    testconstants.SoftwareVersionString,
		HardwareVersion:                          testconstants.HardwareVersion,
		HardwareVersionString:                    testconstants.HardwareVersionString,
		CDVersionNumber:                          testconstants.CDVersionNumber,
		FirmwareDigests:                          testconstants.FirmwareDigests,
		Revoked:                                  testconstants.Revoked,
		OtaURL:                                   testconstants.OtaURL,
		OtaChecksum:                              testconstants.OtaChecksum,
		OtaChecksumType:                          testconstants.OtaChecksumType,
		OtaBlob:                                  testconstants.OtaBlob,
		CommissioningCustomFlow:                  testconstants.CommissioningCustomFlow,
		CommissioningCustomFlowUrl:               testconstants.CommissioningCustomFlowUrl,
		CommissioningModeInitialStepsHint:        testconstants.CommissioningModeInitialStepsHint,
		CommissioningModeInitialStepsInstruction: testconstants.CommissioningModeInitialStepsInstruction,
		CommissioningModeSecondaryStepsHint:      testconstants.CommissioningModeSecondaryStepsHint,
		CommissioningModeSecondaryStepsInstruction: testconstants.CommissioningModeSecondaryStepsInstruction,
		ReleaseNotesUrl: testconstants.ReleaseNotesUrl,
		UserManualUrl:   testconstants.UserManualUrl,
		SupportUrl:      testconstants.SupportUrl,
		ProductURL:      testconstants.ProductURL,
		ChipBlob:        testconstants.ChipBlob,
		VendorBlob:      testconstants.VendorBlob,
	}
}

func getTestModelForUpdate() types.Model {

	return Model{
		VID:                        testconstants.VID,
		PID:                        testconstants.PID,
		CID:                        testconstants.CID + 1,
		Description:                "New Description",
		OtaURL:                     "http://ota.firmware.com/new",
		CDVersionNumber:            testconstants.CDVersionNumber,
		FirmwareDigests:            testconstants.FirmwareDigests,
		Revoked:                    testconstants.Revoked,
		OtaChecksum:                testconstants.OtaChecksum,
		OtaChecksumType:            testconstants.OtaChecksumType,
		OtaBlob:                    testconstants.OtaBlob,
		CommissioningCustomFlowUrl: testconstants.CommissioningCustomFlowUrl,
		ReleaseNotesUrl:            testconstants.ReleaseNotesUrl,
		UserManualUrl:              testconstants.UserManualUrl,
		SupportUrl:                 testconstants.SupportUrl,
		ProductURL:                 testconstants.ProductURL,
		ChipBlob:                   testconstants.ChipBlob,
		VendorBlob:                 testconstants.VendorBlob,
	}
}
func TestMsgAddModelInfo(signer sdk.AccAddress) MsgAddModelInfo {
	return MsgAddModelInfo{
		Model:  getTestModel(),
		Signer: signer,
	}
}

func TestMsgUpdateModelInfo(signer sdk.AccAddress) MsgUpdateModelInfo {

	return MsgUpdateModelInfo{
		Model:  getTestModelForUpdate(),
		Signer: signer,
	}
}
