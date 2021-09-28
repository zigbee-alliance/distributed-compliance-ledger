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

package validator

//nolint:goimports
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
	Cdc             *amino.Codec
	Ctx             sdk.Context
	ValidatorKeeper Keeper
	authKeeper      auth.Keeper
	Handler         sdk.Handler
	Querier         sdk.Querier
	NodeAdmin       sdk.AccAddress
}

func Setup() TestSetup {
	// Init Codec
	cdc := codec.New()
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)

	// Init KVSore
	db := dbm.NewMemDB()

	dbStore := store.NewCommitMultiStore(db)

	validatorKey := sdk.NewKVStoreKey(StoreKey)
	dbStore.MountStoreWithDB(validatorKey, sdk.StoreTypeIAVL, nil)

	authKey := sdk.NewKVStoreKey(auth.StoreKey)
	dbStore.MountStoreWithDB(authKey, sdk.StoreTypeIAVL, nil)

	_ = dbStore.LoadLatestVersion()

	// Init Keepers
	validatorKeeper := NewKeeper(validatorKey, cdc)
	authKeeper := auth.NewKeeper(authKey, cdc)

	// Create context
	ctx := sdk.NewContext(dbStore, abci.Header{ChainID: testconstants.ChainID}, false, log.NewNopLogger())

	// Create Handler and Querier
	querier := NewQuerier(validatorKeeper)
	handler := NewHandler(validatorKeeper, authKeeper)

	account := auth.NewAccount(testconstants.Address1, testconstants.PubKey1, auth.AccountRoles{auth.NodeAdmin}, testconstants.VendorId1)
	authKeeper.SetAccount(ctx, account)

	setup := TestSetup{
		Cdc:             cdc,
		Ctx:             ctx,
		ValidatorKeeper: validatorKeeper,
		authKeeper:      authKeeper,
		Handler:         handler,
		Querier:         querier,
		NodeAdmin:       account.Address,
	}

	return setup
}
