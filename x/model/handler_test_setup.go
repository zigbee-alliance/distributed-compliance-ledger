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

package model

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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/internal/types"
)

type TestSetup struct {
	Cdc         *amino.Codec
	Ctx         sdk.Context
	ModelKeeper Keeper
	authKeeper  auth.Keeper
	Handler     sdk.Handler
	Querier     sdk.Querier
	Vendor      sdk.AccAddress
	VendorId    uint16
}

func Setup() TestSetup {
	// Init Codec
	cdc := codec.New()
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)

	// Init KVSore
	db := dbm.NewMemDB()

	dbStore := store.NewCommitMultiStore(db)

	modelKey := sdk.NewKVStoreKey(StoreKey)
	dbStore.MountStoreWithDB(modelKey, sdk.StoreTypeIAVL, nil)

	authKey := sdk.NewKVStoreKey(auth.StoreKey)
	dbStore.MountStoreWithDB(authKey, sdk.StoreTypeIAVL, nil)

	_ = dbStore.LoadLatestVersion()

	// Init Keepers
	modelKeeper := NewKeeper(modelKey, cdc)
	authKeeper := auth.NewKeeper(authKey, cdc)

	// Create context
	ctx := sdk.NewContext(dbStore, abci.Header{ChainID: testconstants.ChainID}, false, log.NewNopLogger())

	// Create Handler and Querier
	querier := NewQuerier(modelKeeper)
	handler := NewHandler(modelKeeper, authKeeper)

	account := auth.NewAccount(testconstants.Address1, testconstants.PubKey1, auth.AccountRoles{auth.Vendor}, testconstants.VendorId1)
	account.AccountNumber = authKeeper.GetNextAccountNumber(ctx)
	authKeeper.SetAccount(ctx, account)

	setup := TestSetup{
		Cdc:         cdc,
		Ctx:         ctx,
		ModelKeeper: modelKeeper,
		authKeeper:  authKeeper,
		Handler:     handler,
		Querier:     querier,
		Vendor:      account.Address,
	}

	return setup
}

func getTestModel() types.Model {
	return Model{
		VID:                                      testconstants.VendorId1,
		PID:                                      testconstants.PID,
		DeviceTypeID:                             testconstants.DeviceTypeID,
		ProductName:                              testconstants.ProductName,
		ProductLabel:                             testconstants.ProductLabel,
		PartNumber:                               testconstants.PartNumber,
		CommissioningCustomFlow:                  testconstants.CommissioningCustomFlow,
		CommissioningCustomFlowURL:               testconstants.CommissioningCustomFlowURL,
		CommissioningModeInitialStepsHint:        testconstants.CommissioningModeInitialStepsHint,
		CommissioningModeInitialStepsInstruction: testconstants.CommissioningModeInitialStepsInstruction,
		CommissioningModeSecondaryStepsHint:      testconstants.CommissioningModeSecondaryStepsHint,
		CommissioningModeSecondaryStepsInstruction: testconstants.CommissioningModeSecondaryStepsInstruction,
		UserManualURL: testconstants.UserManualURL,
		SupportURL:    testconstants.SupportURL,
		ProductURL:    testconstants.ProductURL,
	}
}

func getTestModelForUpdate() types.Model {
	return Model{
		VID:                        testconstants.VendorId1,
		PID:                        testconstants.PID,
		DeviceTypeID:               testconstants.DeviceTypeID + 1,
		ProductLabel:               "New Description",
		CommissioningCustomFlowURL: testconstants.CommissioningCustomFlowURL,
		UserManualURL:              testconstants.UserManualURL,
		SupportURL:                 testconstants.SupportURL,
		ProductURL:                 testconstants.ProductURL,
	}
}

func TestMsgAddModel(signer sdk.AccAddress) MsgAddModel {
	return MsgAddModel{
		Model:  getTestModel(),
		Signer: signer,
	}
}

func TestMsgUpdateModel(signer sdk.AccAddress) MsgUpdateModel {
	return MsgUpdateModel{
		Model:  getTestModelForUpdate(),
		Signer: signer,
	}
}
