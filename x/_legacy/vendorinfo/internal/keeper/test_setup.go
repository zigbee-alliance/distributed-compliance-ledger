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

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/internal/types"
)

type TestSetup struct {
	Cdc              *codec.Codec
	Ctx              sdk.Context
	VendorInfoKeeper Keeper
	Querier          sdk.Querier
}

func Setup() TestSetup {
	// Init Codec
	cdc := codec.New()
	sdk.RegisterCodec(cdc)

	// Init KVSore
	db := dbm.NewMemDB()
	dbStore := store.NewCommitMultiStore(db)
	vendorInfoKey := sdk.NewKVStoreKey(types.StoreKey)
	dbStore.MountStoreWithDB(vendorInfoKey, sdk.StoreTypeIAVL, nil)
	_ = dbStore.LoadLatestVersion()

	// Init Keepers
	vendorInfoKeeper := NewKeeper(vendorInfoKey, cdc)

	// Init Querier
	querier := NewQuerier(vendorInfoKeeper)

	// Create context
	ctx := sdk.NewContext(dbStore, abci.Header{ChainID: "dcl-test-chain-id"}, false, log.NewNopLogger())

	setup := TestSetup{
		Cdc:              cdc,
		Ctx:              ctx,
		VendorInfoKeeper: vendorInfoKeeper,
		Querier:          querier,
	}

	return setup
}

func DefaultVendorInfo() types.VendorInfo {
	vendorInfo := types.VendorInfo{
		VendorID:             testconstants.VID,
		VendorName:           testconstants.VendorName,
		CompanyLegalName:     testconstants.CompanyLegalName,
		CompanyPreferredName: testconstants.CompanyPreferredName,
		VendorLandingPageURL: testconstants.VendorLandingPageURL,
	}

	return vendorInfo
}

// add vendor info multiple counts.
func PopulateStoreWithVendorInfo(setup TestSetup, count int) uint16 {
	firstID := uint16(1)

	vendorInfo := DefaultVendorInfo()
	vendorInfo.VendorID = firstID

	for i := firstID; i <= uint16(count); i++ {
		// add model info {VID: 1, PID: i}
		vendorInfo.VendorID = i
		setup.VendorInfoKeeper.SetVendorInfo(setup.Ctx, vendorInfo)
	}

	return firstID
}
