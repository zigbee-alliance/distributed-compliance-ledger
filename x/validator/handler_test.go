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

//nolint:testpackage
package validator

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	constants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/internal/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/internal/types"
)

func TestHandler_CreateValidator(t *testing.T) {
	setup := Setup()

	// create validator
	msgCreateValidator := types.NewMsgCreateValidator(constants.ValidatorAddress1, constants.ValidatorPubKey1,
		types.Description{Name: constants.ProductName}, constants.Address1)
	result := setup.Handler(setup.Ctx, msgCreateValidator)
	require.Equal(t, sdk.CodeOK, result.Code)

	events := result.Events.ToABCIEvents()
	require.Equal(t, 2, len(events))
	require.Equal(t, types.EventTypeCreateValidator, events[0].Type)
	require.Equal(t, sdk.EventTypeMessage, events[1].Type)

	// check corresponding records are created
	require.True(t, setup.ValidatorKeeper.IsValidatorPresent(setup.Ctx, msgCreateValidator.Address))

	// this record will be added in the end block handler
	require.False(t, setup.ValidatorKeeper.IsLastValidatorPowerPresent(setup.Ctx, msgCreateValidator.Address))

	// query validator
	validator, _ := queryValidator(setup, msgCreateValidator.Address)
	require.Equal(t, msgCreateValidator.Address, validator.Address)
	require.Equal(t, msgCreateValidator.PubKey, validator.PubKey)
	require.Equal(t, msgCreateValidator.Description, validator.Description)
}

func TestHandler_CreateValidator_ByNotNodeAdmin(t *testing.T) {
	setup := Setup()

	msgCreateValidator := types.NewMsgCreateValidator(constants.ValidatorAddress1, constants.ValidatorPubKey1,
		types.Description{Name: constants.ProductName}, constants.Address1)

	for _, role := range []auth.AccountRole{auth.TestHouse, auth.ZBCertificationCenter, auth.Vendor, auth.Trustee} {
		// create signer account
		account := auth.NewAccount(constants.Address1, constants.PubKey1, auth.AccountRoles{role}, constants.VendorId1)
		setup.authKeeper.SetAccount(setup.Ctx, account)

		// try to create validator
		result := setup.Handler(setup.Ctx, msgCreateValidator)
		require.Equal(t, sdk.CodeUnauthorized, result.Code)
	}
}

func TestHandler_CreateValidator_TwiceForSameValidatorAddress(t *testing.T) {
	setup := Setup()

	// create validator
	msgCreateValidator := types.NewMsgCreateValidator(constants.ValidatorAddress1, constants.ValidatorPubKey1,
		types.Description{Name: constants.ProductName}, constants.Address1)
	result := setup.Handler(setup.Ctx, msgCreateValidator)
	require.Equal(t, sdk.CodeOK, result.Code)

	// create validator
	account := auth.NewAccount(constants.Address2, constants.PubKey2, auth.AccountRoles{auth.NodeAdmin}, constants.VendorId2)
	setup.authKeeper.SetAccount(setup.Ctx, account)

	msgCreateValidator = types.NewMsgCreateValidator(constants.ValidatorAddress1, constants.ValidatorPubKey1,
		types.Description{Name: constants.ProductName}, constants.Address2)
	result = setup.Handler(setup.Ctx, msgCreateValidator)
	require.Equal(t, types.CodeValidatorAlreadyExist, result.Code)
}

func TestHandler_CreateValidator_TwiceForSameValidatorOwner(t *testing.T) {
	setup := Setup()

	// create validator
	msgCreateValidator := types.NewMsgCreateValidator(constants.ValidatorAddress1, constants.ValidatorPubKey1,
		types.Description{Name: constants.ProductName}, constants.Address1)
	result := setup.Handler(setup.Ctx, msgCreateValidator)
	require.Equal(t, sdk.CodeOK, result.Code)

	// create validator with different address
	msgCreateValidator2 := types.NewMsgCreateValidator(constants.ValidatorAddress2, constants.ValidatorPubKey2,
		types.Description{Name: constants.ProductName}, constants.Address1)
	result = setup.Handler(setup.Ctx, msgCreateValidator2)
	require.Equal(t, types.CodeAccountAlreadyHasNode, result.Code)
}

func queryValidator(setup TestSetup, address sdk.ConsAddress) (*types.Validator, sdk.Error) {
	// query validator
	result, err := setup.Querier(
		setup.Ctx,
		[]string{keeper.QueryValidator, address.String()},
		abci.RequestQuery{},
	)
	if err != nil {
		return nil, err
	}

	var validator types.Validator

	setup.Cdc.MustUnmarshalJSON(result, &validator)

	return &validator, nil
}
