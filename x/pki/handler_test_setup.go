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

package pki

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
)

type TestSetup struct {
	Cdc        *amino.Codec
	Ctx        sdk.Context
	PkiKeeper  Keeper
	AuthKeeper auth.Keeper
	Handler    sdk.Handler
	Querier    sdk.Querier
	Trustee    sdk.AccAddress
}

func Setup() TestSetup {
	// Init Codec
	cdc := codec.New()
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)

	// Init KVSore
	db := dbm.NewMemDB()

	dbStore := store.NewCommitMultiStore(db)

	pkiKey := sdk.NewKVStoreKey(StoreKey)
	dbStore.MountStoreWithDB(pkiKey, sdk.StoreTypeIAVL, nil)

	authKey := sdk.NewKVStoreKey(auth.StoreKey)
	dbStore.MountStoreWithDB(authKey, sdk.StoreTypeIAVL, nil)

	_ = dbStore.LoadLatestVersion()

	// Init Keepers
	pkiKeeper := NewKeeper(pkiKey, cdc)
	authKeeper := auth.NewKeeper(authKey, cdc)

	// Create context
	ctx := sdk.NewContext(dbStore, abci.Header{ChainID: testconstants.ChainID}, false, log.NewNopLogger())

	// Create Handler and Querier
	querier := NewQuerier(pkiKeeper)
	handler := NewHandler(pkiKeeper, authKeeper)

	account := auth.NewAccount(testconstants.Address2, testconstants.PubKey2, auth.AccountRoles{auth.Trustee}, testconstants.VendorId2)
	account.AccountNumber = authKeeper.GetNextAccountNumber(ctx)
	authKeeper.SetAccount(ctx, account)

	setup := TestSetup{
		Cdc:        cdc,
		Ctx:        ctx,
		PkiKeeper:  pkiKeeper,
		AuthKeeper: authKeeper,
		Handler:    handler,
		Querier:    querier,
		Trustee:    account.Address,
	}

	return setup
}
