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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/internal/types"
)

type TestSetup struct {
	Cdc             *codec.Codec
	Ctx             sdk.Context
	ValidatorKeeper Keeper
	Querier         sdk.Querier
}

func Setup() TestSetup {
	// Init Codec
	cdc := codec.New()
	sdk.RegisterCodec(cdc)

	// Init KVSore
	db := dbm.NewMemDB()
	dbStore := store.NewCommitMultiStore(db)
	validatorKey := sdk.NewKVStoreKey(types.StoreKey)
	dbStore.MountStoreWithDB(validatorKey, sdk.StoreTypeIAVL, nil)
	_ = dbStore.LoadLatestVersion()

	// Init Keepers
	validatorKeeper := NewKeeper(validatorKey, cdc)

	// Init Querier
	querier := NewQuerier(validatorKeeper)

	// Create context
	ctx := sdk.NewContext(dbStore, abci.Header{ChainID: testconstants.ChainID}, false, log.NewNopLogger())

	setup := TestSetup{
		Cdc:             cdc,
		Ctx:             ctx,
		ValidatorKeeper: validatorKeeper,
		Querier:         querier,
	}

	return setup
}

func DefaultValidator() types.Validator {
	return types.NewValidator(
		testconstants.ValidatorAddress1,
		testconstants.ValidatorPubKey1,
		types.Description{Name: testconstants.ProductName},
		testconstants.Owner,
	)
}

func DefaultValidatorPower() types.LastValidatorPower {
	return types.NewLastValidatorPower(testconstants.ValidatorAddress1)
}

func StoreTwoValidators(setup TestSetup) (types.Validator, types.Validator) {
	validator1 := types.NewValidator(testconstants.ValidatorAddress1, testconstants.ValidatorPubKey1,
		types.Description{Name: "Validator 1"}, testconstants.Address1)
	setup.ValidatorKeeper.SetValidator(setup.Ctx, validator1)

	validator2 := types.NewValidator(testconstants.ValidatorAddress2, testconstants.ValidatorPubKey2,
		types.Description{Name: "Validator 2"}, testconstants.Address2)
	setup.ValidatorKeeper.SetValidator(setup.Ctx, validator2)

	return validator1, validator2
}
