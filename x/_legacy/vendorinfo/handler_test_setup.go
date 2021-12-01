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

package vendorinfo

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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/internal/types"
)

type TestSetup struct {
	Cdc              *amino.Codec
	Ctx              sdk.Context
	VendorInfoKeeper Keeper
	authKeeper       auth.Keeper
	Handler          sdk.Handler
	Querier          sdk.Querier
	Vendor           sdk.AccAddress
	VendorID         uint16
}

func Setup() TestSetup {
	// Init Codec
	cdc := codec.New()
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)

	// Init KVSore
	db := dbm.NewMemDB()

	dbStore := store.NewCommitMultiStore(db)

	vendorInfoKey := sdk.NewKVStoreKey(StoreKey)
	dbStore.MountStoreWithDB(vendorInfoKey, sdk.StoreTypeIAVL, nil)

	authKey := sdk.NewKVStoreKey(auth.StoreKey)
	dbStore.MountStoreWithDB(authKey, sdk.StoreTypeIAVL, nil)

	_ = dbStore.LoadLatestVersion()

	// Init Keepers
	VendorInfoKeeper := NewKeeper(vendorInfoKey, cdc)
	authKeeper := auth.NewKeeper(authKey, cdc)

	// Create context
	ctx := sdk.NewContext(dbStore, abci.Header{ChainID: testconstants.ChainID}, false, log.NewNopLogger())

	// Create Handler and Querier
	querier := NewQuerier(VendorInfoKeeper)
	handler := NewHandler(VendorInfoKeeper, authKeeper)

	account := auth.NewAccount(testconstants.Address1, testconstants.PubKey1,
		auth.AccountRoles{auth.Vendor}, testconstants.VendorID1)
	account.AccountNumber = authKeeper.GetNextAccountNumber(ctx)
	authKeeper.SetAccount(ctx, account)

	setup := TestSetup{
		Cdc:              cdc,
		Ctx:              ctx,
		VendorInfoKeeper: VendorInfoKeeper,
		authKeeper:       authKeeper,
		Handler:          handler,
		Querier:          querier,
		Vendor:           account.Address,
	}

	return setup
}

func getTestVendor() types.VendorInfo {
	vendorInfo := types.VendorInfo{
		VendorID:             testconstants.VendorID1,
		VendorName:           testconstants.VendorName,
		CompanyLegalName:     testconstants.CompanyLegalName,
		CompanyPreferredName: testconstants.CompanyPreferredName,
		VendorLandingPageURL: testconstants.VendorLandingPageURL,
	}

	return vendorInfo
}

func getTestVendorForUpdate() types.VendorInfo {
	vendorInfo := types.VendorInfo{
		VendorID:             testconstants.VendorID1,
		VendorName:           testconstants.VendorName + "-updated",
		CompanyLegalName:     testconstants.CompanyLegalName + "-updated",
		CompanyPreferredName: testconstants.CompanyPreferredName + "-updated",
		VendorLandingPageURL: testconstants.VendorLandingPageURL + "-updated",
	}

	return vendorInfo
}

func TestMsgAddVendorInfo(signer sdk.AccAddress) MsgAddVendorInfo {
	return MsgAddVendorInfo{
		VendorInfo: getTestVendor(),
		Signer:     signer,
	}
}

func TestMsgUpdateVendorInfo(signer sdk.AccAddress) MsgUpdateVendorInfo {
	return MsgUpdateVendorInfo{
		VendorInfo: getTestVendorForUpdate(),
		Signer:     signer,
	}
}
