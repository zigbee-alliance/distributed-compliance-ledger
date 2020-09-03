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

package compliance

import (
	testconstants "git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/go-amino"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

type TestSetup struct {
	Cdc                  *amino.Codec
	Ctx                  sdk.Context
	CompliancetKeeper    Keeper
	CompliancetestKeeper compliancetest.Keeper
	authKeeper           auth.Keeper
	ModelinfoKeeper      modelinfo.Keeper
	Handler              sdk.Handler
	Querier              sdk.Querier
	CertificationCenter  sdk.AccAddress
}

func Setup() TestSetup {
	// Init Codec
	cdc := codec.New()
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)

	// Init KVSore
	db := dbm.NewMemDB()

	dbStore := store.NewCommitMultiStore(db)

	complianceKey := sdk.NewKVStoreKey(StoreKey)
	dbStore.MountStoreWithDB(complianceKey, sdk.StoreTypeIAVL, nil)

	authKey := sdk.NewKVStoreKey(auth.StoreKey)
	dbStore.MountStoreWithDB(authKey, sdk.StoreTypeIAVL, nil)

	modelinfoKey := sdk.NewKVStoreKey(modelinfo.StoreKey)
	dbStore.MountStoreWithDB(modelinfoKey, sdk.StoreTypeIAVL, nil)

	compliancetestKey := sdk.NewKVStoreKey(compliancetest.StoreKey)
	dbStore.MountStoreWithDB(compliancetestKey, sdk.StoreTypeIAVL, nil)

	_ = dbStore.LoadLatestVersion()

	// Init Keepers
	compliancetKeeper := NewKeeper(complianceKey, cdc)
	compliancetestKeeper := compliancetest.NewKeeper(compliancetestKey, cdc)
	authKeeper := auth.NewKeeper(authKey, cdc)
	modelinfoKeeper := modelinfo.NewKeeper(modelinfoKey, cdc)

	// Create context
	ctx := sdk.NewContext(dbStore, abci.Header{ChainID: testconstants.ChainID}, false, log.NewNopLogger())

	// Create Handler and Querier
	querier := NewQuerier(compliancetKeeper)
	handler := NewHandler(compliancetKeeper, modelinfoKeeper, compliancetestKeeper, authKeeper)

	account := auth.NewAccount(testconstants.Address1, testconstants.PubKey1,
		auth.AccountRoles{auth.ZBCertificationCenter})
	account.AccountNumber = authKeeper.GetNextAccountNumber(ctx)
	authKeeper.SetAccount(ctx, account)

	setup := TestSetup{
		Cdc:                  cdc,
		Ctx:                  ctx,
		CompliancetKeeper:    compliancetKeeper,
		CompliancetestKeeper: compliancetestKeeper,
		ModelinfoKeeper:      modelinfoKeeper,
		authKeeper:           authKeeper,
		Handler:              handler,
		Querier:              querier,
		CertificationCenter:  account.Address,
	}

	return setup
}
