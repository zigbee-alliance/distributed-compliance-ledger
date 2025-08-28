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

	// cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types".

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	sdkstakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	testkeeper "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	dclauthkeeper "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/keeper"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

type TestSetup struct {
	Ctx             sdk.Context
	ValidatorKeeper keeper.Keeper
	DclauthKeeper   dclauthkeeper.Keeper
	Handler         sdk.Handler
	NodeAdmin       sdk.AccAddress
}

func Setup(t *testing.T) TestSetup {
	t.Helper()
	dclauthK, _ := testkeeper.DclauthKeeper(t)
	k, ctx := testkeeper.ValidatorKeeper(t, dclauthK)
	handler := NewHandler(*k)

	ba := authtypes.NewBaseAccount(testconstants.Address1, testconstants.PubKey1, 0, 0)
	account := dclauthtypes.NewAccount(
		ba,
		dclauthtypes.AccountRoles{dclauthtypes.NodeAdmin}, nil, nil, testconstants.VendorID1, testconstants.ProductIDsEmpty,
	)
	dclauthK.SetAccount(ctx, account)

	valAddress, _ := sdk.ValAddressFromBech32(testconstants.ValidatorAddress1)

	validator1, _ := types.NewValidator(
		valAddress,
		testconstants.ValidatorPubKey1,
		types.Description{Moniker: "Validator 1"},
	)
	k.SetValidator(ctx, validator1)

	setup := TestSetup{
		Ctx:             ctx,
		ValidatorKeeper: *k,
		DclauthKeeper:   *dclauthK,
		Handler:         handler,
	}

	return setup
}

func TestHandler_CreateValidator(t *testing.T) {
	setup := Setup(t)

	valAddr := sdk.ValAddress(testconstants.Address1)
	// create validator
	msgCreateValidator, err := types.NewMsgCreateValidator(
		valAddr,
		testconstants.ValidatorPubKey1,
		&types.Description{Moniker: testconstants.ProductName},
	)
	require.NoError(t, err)
	result, err := setup.Handler(setup.Ctx, msgCreateValidator)
	require.NoError(t, err)

	events := result.Events
	require.Equal(t, 2, len(events))
	require.Equal(t, types.EventTypeCreateValidator, events[0].Type)
	require.Equal(t, sdk.EventTypeMessage, events[1].Type)

	// check corresponding records are created
	require.True(t, setup.ValidatorKeeper.IsValidatorPresent(setup.Ctx, valAddr))

	// this record will be added in the end block handler
	require.False(t, setup.ValidatorKeeper.IsLastValidatorPowerPresent(setup.Ctx, valAddr))

	// query validator
	validator, _ := queryValidator(setup, valAddr.String())
	require.Equal(t, msgCreateValidator.Signer, valAddr.String())
	require.Equal(t, msgCreateValidator.PubKey, validator.PubKey)
	require.Equal(t, msgCreateValidator.Description, *validator.Description)
}

func TestHandler_CreateValidator_ByNotNodeAdmin(t *testing.T) {
	setup := Setup(t)

	msgCreateValidator, err := types.NewMsgCreateValidator(
		sdk.ValAddress(testconstants.Address1),
		testconstants.ValidatorPubKey1,
		&types.Description{Moniker: testconstants.ProductName},
	)
	require.NoError(t, err)

	for _, role := range []dclauthtypes.AccountRole{dclauthtypes.CertificationCenter, dclauthtypes.Vendor, dclauthtypes.Trustee} {
		// create signer account
		ba := authtypes.NewBaseAccount(testconstants.Address1, testconstants.PubKey1, 0, 0)
		account := dclauthtypes.NewAccount(ba, dclauthtypes.AccountRoles{role}, nil, nil, testconstants.VendorID1, testconstants.ProductIDsEmpty)
		setup.DclauthKeeper.SetAccount(setup.Ctx, account)

		// try to create validator
		_, err := setup.Handler(setup.Ctx, msgCreateValidator)
		require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
	}
}

func TestHandler_CreateValidator_WithIncorrectPubKey(t *testing.T) {
	setup := Setup(t)

	valAddr := sdk.ValAddress(testconstants.Address1)
	// create validator
	msgCreateValidator, err := types.NewMsgCreateValidator(
		valAddr,
		testconstants.PubKey1, // not a validator pubkey
		&types.Description{Moniker: testconstants.ProductName},
	)
	require.NoError(t, err)

	_, err = setup.Handler(setup.Ctx, msgCreateValidator)
	require.ErrorIs(t, err, sdkstakingtypes.ErrValidatorPubKeyTypeNotSupported)
}

func TestHandler_CreateValidator_TwiceForSameValidatorAddress(t *testing.T) {
	setup := Setup(t)

	// create validator
	msgCreateValidator, err := types.NewMsgCreateValidator(
		sdk.ValAddress(testconstants.Address1),
		testconstants.ValidatorPubKey1,
		&types.Description{Moniker: testconstants.ProductName},
	)
	require.NoError(t, err)
	_, err = setup.Handler(setup.Ctx, msgCreateValidator)
	require.NoError(t, err)

	// create validator
	ba := authtypes.NewBaseAccount(testconstants.Address2, testconstants.PubKey2, 0, 0)
	account := dclauthtypes.NewAccount(ba,
		dclauthtypes.AccountRoles{dclauthtypes.NodeAdmin}, nil, nil, testconstants.VendorID2, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account)

	msgCreateValidator, err = types.NewMsgCreateValidator(
		sdk.ValAddress(testconstants.Address2),
		testconstants.ValidatorPubKey1,
		&types.Description{Moniker: testconstants.ProductName},
	)
	require.NoError(t, err)
	_, err = setup.Handler(setup.Ctx, msgCreateValidator)
	require.ErrorIs(t, err, sdkstakingtypes.ErrValidatorPubKeyExists)
}

func TestHandler_CreateValidator_TwiceForSameValidatorOwner(t *testing.T) {
	setup := Setup(t)

	// create validator
	msgCreateValidator, err := types.NewMsgCreateValidator(
		sdk.ValAddress(testconstants.Address1),
		testconstants.ValidatorPubKey1,
		&types.Description{Moniker: testconstants.ProductName},
	)
	require.NoError(t, err)
	_, err = setup.Handler(setup.Ctx, msgCreateValidator)
	require.NoError(t, err)

	// create validator with different address
	msgCreateValidator2, err := types.NewMsgCreateValidator(
		sdk.ValAddress(testconstants.Address1),
		testconstants.ValidatorPubKey2,
		&types.Description{Moniker: testconstants.ProductName},
	)
	require.NoError(t, err)
	_, err = setup.Handler(setup.Ctx, msgCreateValidator2)
	require.ErrorIs(t, err, sdkstakingtypes.ErrValidatorOwnerExists)
}

func queryValidator(setup TestSetup, owner string) (*types.Validator, error) {
	resp, err := setup.ValidatorKeeper.Validator(
		sdk.WrapSDKContext(setup.Ctx), &types.QueryGetValidatorRequest{Owner: owner},
	)
	if err != nil {
		return nil, err
	}

	return &resp.Validator, nil
}

func TestHandler_ProposedDisableValidatorExists(t *testing.T) {
	setup := Setup(t)

	// create Trustees
	ba1 := authtypes.NewBaseAccount(testconstants.Address1, testconstants.PubKey1, 0, 0)
	account1 := dclauthtypes.NewAccount(ba1,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID1, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account1)

	ba2 := authtypes.NewBaseAccount(testconstants.Address2, testconstants.PubKey2, 0, 0)
	account2 := dclauthtypes.NewAccount(ba2,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID2, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account2)

	ba3 := authtypes.NewBaseAccount(testconstants.Address3, testconstants.PubKey3, 0, 0)
	account3 := dclauthtypes.NewAccount(ba3,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID3, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account3)

	valAddress, err := sdk.ValAddressFromBech32(testconstants.ValidatorAddress1)
	require.NoError(t, err)

	// propose new disable validator
	msgProposeDisableValidator := NewMsgProposeDisableValidator(account1.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgProposeDisableValidator)
	require.NoError(t, err)

	msgProposeDisableValidator.Creator = account2.GetAddress().String()
	// propose the same disable validator
	_, err = setup.Handler(setup.Ctx, msgProposeDisableValidator)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrProposedDisableValidatorAlreadyExists)
}

func TestHandler_ProposeAndRejectDisableValidator(t *testing.T) {
	setup := Setup(t)

	// create Trustees
	ba1 := authtypes.NewBaseAccount(testconstants.Address1, testconstants.PubKey1, 0, 0)
	account1 := dclauthtypes.NewAccount(ba1,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID1, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account1)

	ba2 := authtypes.NewBaseAccount(testconstants.Address2, testconstants.PubKey2, 0, 0)
	account2 := dclauthtypes.NewAccount(ba2,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID2, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account2)

	ba3 := authtypes.NewBaseAccount(testconstants.Address3, testconstants.PubKey3, 0, 0)
	account3 := dclauthtypes.NewAccount(ba3,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID3, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account3)

	valAddress, err := sdk.ValAddressFromBech32(testconstants.ValidatorAddress1)
	require.NoError(t, err)

	// propose disable validator
	msgProposeDisableValidator := NewMsgProposeDisableValidator(account1.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgProposeDisableValidator)
	require.NoError(t, err)

	// reject disable validator
	rejectDisableValidator := types.NewMsgRejectDisableValidator(account1.GetAddress(), valAddress, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectDisableValidator)
	require.NoError(t, err)

	_, found := setup.ValidatorKeeper.GetProposedDisableValidator(setup.Ctx, valAddress.String())
	require.False(t, found)
}

func TestHandler_ProposeAndRejectDisableValidator_ByAnotherTrustee(t *testing.T) {
	setup := Setup(t)

	// create Trustees
	ba1 := authtypes.NewBaseAccount(testconstants.Address1, testconstants.PubKey1, 0, 0)
	account1 := dclauthtypes.NewAccount(ba1,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID1, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account1)

	ba2 := authtypes.NewBaseAccount(testconstants.Address2, testconstants.PubKey2, 0, 0)
	account2 := dclauthtypes.NewAccount(ba2,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID2, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account2)

	ba3 := authtypes.NewBaseAccount(testconstants.Address3, testconstants.PubKey3, 0, 0)
	account3 := dclauthtypes.NewAccount(ba3,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID3, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account3)

	valAddress, err := sdk.ValAddressFromBech32(testconstants.ValidatorAddress1)
	require.NoError(t, err)

	// propose disable validator
	msgProposeDisableValidator := NewMsgProposeDisableValidator(account1.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgProposeDisableValidator)
	require.NoError(t, err)

	// reject disable validator
	rejectDisableValidator := types.NewMsgRejectDisableValidator(account2.GetAddress(), valAddress, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectDisableValidator)
	require.NoError(t, err)

	_, found := setup.ValidatorKeeper.GetProposedDisableValidator(setup.Ctx, valAddress.String())
	require.True(t, found)
}

func TestHandler_OnlyTrusteeCanProposeDisableValidator(t *testing.T) {
	setup := Setup(t)

	// create account with non-trustee roles
	ba1 := authtypes.NewBaseAccount(testconstants.Address1, testconstants.PubKey1, 0, 0)
	account1 := dclauthtypes.NewAccount(ba1,
		dclauthtypes.AccountRoles{dclauthtypes.CertificationCenter, dclauthtypes.Vendor, dclauthtypes.NodeAdmin}, nil, nil, testconstants.VendorID1, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account1)

	valAddress, err := sdk.ValAddressFromBech32(testconstants.ValidatorAddress1)
	require.NoError(t, err)

	// propose new disablevalidator
	msgProposeDisableValidator := NewMsgProposeDisableValidator(account1.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgProposeDisableValidator)
	require.Error(t, err)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)

	// create account with trustee roles
	ba1 = authtypes.NewBaseAccount(testconstants.Address2, testconstants.PubKey2, 0, 0)
	account1 = dclauthtypes.NewAccount(ba1,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID2, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account1)

	// propose new disablevalidator
	msgProposeDisableValidator = NewMsgProposeDisableValidator(account1.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgProposeDisableValidator)
	require.NoError(t, err)
}

func TestHandler_ProposeDisableValidatorWhenSeveralVotesNeeded(t *testing.T) {
	setup := Setup(t)

	// create Trustees
	ba1 := authtypes.NewBaseAccount(testconstants.Address1, testconstants.PubKey1, 0, 0)
	account1 := dclauthtypes.NewAccount(ba1,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID1, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account1)

	ba2 := authtypes.NewBaseAccount(testconstants.Address2, testconstants.PubKey2, 0, 0)
	account2 := dclauthtypes.NewAccount(ba2,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID2, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account2)

	ba3 := authtypes.NewBaseAccount(testconstants.Address3, testconstants.PubKey3, 0, 0)
	account3 := dclauthtypes.NewAccount(ba3,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID3, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account3)

	valAddress, err := sdk.ValAddressFromBech32(testconstants.ValidatorAddress1)
	require.NoError(t, err)

	// propose new disablevalidator
	msgProposeDisableValidator := NewMsgProposeDisableValidator(account1.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgProposeDisableValidator)
	require.NoError(t, err)

	proposedDisableValidator, isFound := setup.ValidatorKeeper.GetProposedDisableValidator(setup.Ctx, msgProposeDisableValidator.Address)
	require.True(t, isFound)
	require.Equal(t, msgProposeDisableValidator.Address, proposedDisableValidator.Address)
	require.Equal(t, msgProposeDisableValidator.Creator, proposedDisableValidator.Creator)
	require.Equal(t, msgProposeDisableValidator.Creator, proposedDisableValidator.Approvals[0].Address)
	require.Equal(t, msgProposeDisableValidator.Info, proposedDisableValidator.Approvals[0].Info)
	require.Equal(t, msgProposeDisableValidator.Time, proposedDisableValidator.Approvals[0].Time)
	require.Equal(t, 1, len(proposedDisableValidator.Approvals))

	_, isFound = setup.ValidatorKeeper.GetDisabledValidator(setup.Ctx, msgProposeDisableValidator.Address)
	require.False(t, isFound)

	valAddress, err = sdk.ValAddressFromBech32(msgProposeDisableValidator.Address)
	require.NoError(t, err)

	validator, isFound := setup.ValidatorKeeper.GetValidator(setup.Ctx, valAddress)
	require.True(t, isFound)
	require.False(t, validator.Jailed)
	require.Equal(t, int32(10), validator.Power)
	require.Equal(t, "", validator.JailedReason)
}

func TestHandler_DisabledValidator(t *testing.T) {
	setup := Setup(t)

	// create Trustees
	ba1 := authtypes.NewBaseAccount(testconstants.Address1, testconstants.PubKey1, 0, 0)
	account1 := dclauthtypes.NewAccount(ba1,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID1, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account1)

	ba2 := authtypes.NewBaseAccount(testconstants.Address2, testconstants.PubKey2, 0, 0)
	account2 := dclauthtypes.NewAccount(ba2,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID2, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account2)

	ba3 := authtypes.NewBaseAccount(testconstants.Address3, testconstants.PubKey3, 0, 0)
	account3 := dclauthtypes.NewAccount(ba3,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID3, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account3)

	valAddress, err := sdk.ValAddressFromBech32(testconstants.ValidatorAddress1)
	require.NoError(t, err)

	// propose and approve new disablevalidators
	msgProposeDisableValidator1 := NewMsgProposeDisableValidator(account1.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgProposeDisableValidator1)
	require.NoError(t, err)

	proposedDisableValidator, isFound := setup.ValidatorKeeper.GetProposedDisableValidator(setup.Ctx, valAddress.String())
	require.True(t, isFound)
	require.Equal(t, msgProposeDisableValidator1.Address, proposedDisableValidator.Address)
	require.Equal(t, msgProposeDisableValidator1.Creator, proposedDisableValidator.Creator)
	require.Equal(t, msgProposeDisableValidator1.Creator, proposedDisableValidator.Approvals[0].Address)
	require.Equal(t, 1, len(proposedDisableValidator.Approvals))

	_, isFound = setup.ValidatorKeeper.GetDisabledValidator(setup.Ctx, valAddress.String())
	require.False(t, isFound)

	msgApproveDisableValidator := NewMsgApproveDisableValidator(account2.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgApproveDisableValidator)
	require.NoError(t, err)

	_, isFound = setup.ValidatorKeeper.GetProposedDisableValidator(setup.Ctx, valAddress.String())
	require.False(t, isFound)

	disabledValidator, isFound := setup.ValidatorKeeper.GetDisabledValidator(setup.Ctx, valAddress.String())
	require.True(t, isFound)
	require.Equal(t, msgProposeDisableValidator1.Address, disabledValidator.Address)
	require.Equal(t, msgProposeDisableValidator1.Creator, disabledValidator.Creator)
	require.Equal(t, msgProposeDisableValidator1.Creator, disabledValidator.Approvals[0].Address)
	require.Equal(t, msgProposeDisableValidator1.Info, disabledValidator.Approvals[0].Info)
	require.Equal(t, msgProposeDisableValidator1.Time, disabledValidator.Approvals[0].Time)
	require.False(t, disabledValidator.DisabledByNodeAdmin)

	validator, isFound := setup.ValidatorKeeper.GetValidator(setup.Ctx, valAddress)
	require.True(t, isFound)
	require.True(t, validator.Jailed)
	require.Equal(t, int32(0), validator.Power)
	require.Equal(t, testconstants.Info, validator.JailedReason)
}

func TestHandler_DisabledValidatorOnPropose(t *testing.T) {
	setup := Setup(t)

	// create Trustees
	ba1 := authtypes.NewBaseAccount(testconstants.Address1, testconstants.PubKey1, 0, 0)
	account1 := dclauthtypes.NewAccount(ba1,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID1, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account1)

	ba2 := authtypes.NewBaseAccount(testconstants.Address2, testconstants.PubKey2, 0, 0)
	account2 := dclauthtypes.NewAccount(ba2,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID2, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account2)

	valAddress, err := sdk.ValAddressFromBech32(testconstants.ValidatorAddress1)
	require.NoError(t, err)

	// propose new disablevalidator
	msgProposeDisableValidator1 := NewMsgProposeDisableValidator(account1.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgProposeDisableValidator1)
	require.NoError(t, err)

	_, isFound := setup.ValidatorKeeper.GetProposedDisableValidator(setup.Ctx, msgProposeDisableValidator1.Address)
	require.True(t, isFound)

	// approve new disablevalidator
	msgApproveDisableValidator := NewMsgApproveDisableValidator(account2.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgApproveDisableValidator)
	require.NoError(t, err)

	disabledValidator, isFound := setup.ValidatorKeeper.GetDisabledValidator(setup.Ctx, msgProposeDisableValidator1.Address)
	require.True(t, isFound)
	require.Equal(t, msgProposeDisableValidator1.Address, disabledValidator.Address)
	require.Equal(t, msgProposeDisableValidator1.Creator, disabledValidator.Creator)
	require.Equal(t, msgProposeDisableValidator1.Creator, disabledValidator.Approvals[0].Address)
	require.Equal(t, msgProposeDisableValidator1.Info, disabledValidator.Approvals[0].Info)
	require.Equal(t, msgProposeDisableValidator1.Time, disabledValidator.Approvals[0].Time)
	require.False(t, disabledValidator.DisabledByNodeAdmin)

	validator, isFound := setup.ValidatorKeeper.GetValidator(setup.Ctx, valAddress)
	require.True(t, isFound)
	require.True(t, validator.Jailed)
	require.Equal(t, int32(0), validator.Power)
	require.Equal(t, testconstants.Info, validator.JailedReason)
}

func TestHandler_OnlyTrusteeCanApproveDisableValidator(t *testing.T) {
	setup := Setup(t)

	// create Trustees
	ba1 := authtypes.NewBaseAccount(testconstants.Address1, testconstants.PubKey1, 0, 0)
	account1 := dclauthtypes.NewAccount(ba1,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID1, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account1)

	ba2 := authtypes.NewBaseAccount(testconstants.Address2, testconstants.PubKey2, 0, 0)
	account2 := dclauthtypes.NewAccount(ba2,
		dclauthtypes.AccountRoles{dclauthtypes.NodeAdmin, dclauthtypes.CertificationCenter, dclauthtypes.Vendor}, nil, nil, testconstants.VendorID2, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account2)

	ba3 := authtypes.NewBaseAccount(testconstants.Address3, testconstants.PubKey3, 0, 0)
	account3 := dclauthtypes.NewAccount(ba3,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID3, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account3)

	ba4 := authtypes.NewBaseAccount(testconstants.Address4, testconstants.PubKey4, 0, 0)
	account4 := dclauthtypes.NewAccount(ba4,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID4, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account4)

	valAddress, err := sdk.ValAddressFromBech32(testconstants.ValidatorAddress1)
	require.NoError(t, err)

	// propose and approve new disablevalidator
	msgProposeDisableValidator := NewMsgProposeDisableValidator(account1.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgProposeDisableValidator)
	require.NoError(t, err)

	msgApproveDisableValidator := NewMsgApproveDisableValidator(account2.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgApproveDisableValidator)
	require.Error(t, err)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestHandler_ProposedDisableValidatorDoesNotExist(t *testing.T) {
	setup := Setup(t)

	// create Trustees
	ba1 := authtypes.NewBaseAccount(testconstants.Address1, testconstants.PubKey1, 0, 0)
	account1 := dclauthtypes.NewAccount(ba1,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID1, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account1)

	msgApproveDisableValidator := NewMsgApproveDisableValidator(account1.GetAddress(), sdk.ValAddress(testconstants.ValidatorAddress1))
	_, err := setup.Handler(setup.Ctx, msgApproveDisableValidator)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrProposedDisableValidatorDoesNotExist)
}

func TestHandler_MessageCreatorAlreadyApprovedDisableValidator(t *testing.T) {
	setup := Setup(t)

	// create Trustees
	ba1 := authtypes.NewBaseAccount(testconstants.Address1, testconstants.PubKey1, 0, 0)
	account1 := dclauthtypes.NewAccount(ba1,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID1, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account1)

	ba2 := authtypes.NewBaseAccount(testconstants.Address2, testconstants.PubKey2, 0, 0)
	account2 := dclauthtypes.NewAccount(ba2,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID2, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account2)

	ba3 := authtypes.NewBaseAccount(testconstants.Address3, testconstants.PubKey3, 0, 0)
	account3 := dclauthtypes.NewAccount(ba3,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID3, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account3)

	valAddress, err := sdk.ValAddressFromBech32(testconstants.ValidatorAddress1)
	require.NoError(t, err)

	// propose and approve new disablevalidator
	msgProposeDisableValidator := NewMsgProposeDisableValidator(account1.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgProposeDisableValidator)
	require.NoError(t, err)

	msgApproveDisableValidator := NewMsgApproveDisableValidator(account1.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgApproveDisableValidator)
	require.Error(t, err)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestHandler_DisabledValidatorDoesNotExist(t *testing.T) {
	setup := Setup(t)

	msgEnableValidator := types.NewMsgEnableValidator(sdk.ValAddress(testconstants.Address1))
	_, err := setup.Handler(setup.Ctx, msgEnableValidator)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrDisabledValidatorDoesNotExist)
}

func TestHandler_NodeAdminCanEnableValidatorDisabledByTrustees(t *testing.T) {
	setup := Setup(t)

	// create Trustee and NodeAdmin
	ba1 := authtypes.NewBaseAccount(testconstants.Address1, testconstants.PubKey1, 0, 0)
	account1 := dclauthtypes.NewAccount(ba1,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID1, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account1)

	valAddress, err := sdk.ValAddressFromBech32(testconstants.ValidatorAddress1)
	require.NoError(t, err)

	ba2 := authtypes.NewBaseAccount(sdk.AccAddress(valAddress), testconstants.PubKey2, 0, 0)
	account2 := dclauthtypes.NewAccount(ba2,
		dclauthtypes.AccountRoles{dclauthtypes.NodeAdmin}, nil, nil, testconstants.VendorID2, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account2)

	// propose and approve new disablevalidator (will be approved because of 1 trustee)
	msgProposeDisableValidator := NewMsgProposeDisableValidator(account1.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgProposeDisableValidator)
	require.NoError(t, err)

	msgEnableValidator := types.NewMsgEnableValidator(valAddress)
	_, err = setup.Handler(setup.Ctx, msgEnableValidator)
	require.NoError(t, err)
}

func TestHandler_DisabledValidatorAlreadyExistsPropose(t *testing.T) {
	setup := Setup(t)

	// create Trustees
	ba1 := authtypes.NewBaseAccount(testconstants.Address3, testconstants.PubKey3, 0, 0)
	account1 := dclauthtypes.NewAccount(ba1,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID3, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account1)

	ba2 := authtypes.NewBaseAccount(testconstants.Address2, testconstants.PubKey2, 0, 0)
	account2 := dclauthtypes.NewAccount(ba2,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID2, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account2)

	valAddress, err := sdk.ValAddressFromBech32(testconstants.ValidatorAddress1)
	require.NoError(t, err)

	// propose new disablevalidator
	msgProposeDisableValidator := NewMsgProposeDisableValidator(account1.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgProposeDisableValidator)
	require.NoError(t, err)

	// approve new disablevalidator
	msgApproveDisableValidator := NewMsgApproveDisableValidator(account2.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgApproveDisableValidator)
	require.NoError(t, err)

	msgProposeDisableValidator = NewMsgProposeDisableValidator(account1.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgProposeDisableValidator)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrDisabledValidatorAlreadytExists)
}

// func TestHandler_DisabledValidatorAlreadyExistsDisable(t *testing.T) {
// 	setup := Setup(t)

// ba3 := authtypes.NewBaseAccount(sdk.AccAddress(testconstants.ValidatorAddress1), testconstants.PubKey3, 0, 0)
// account3 := dclauthtypes.NewAccount(ba3,
// 	dclauthtypes.AccountRoles{dclauthtypes.NodeAdmin}, nil, testconstants.VendorID3)
// setup.DclauthKeeper.SetAccount(setup.Ctx, account3)

// valAddress, err := sdk.ValAddressFromBech32(testconstants.ValidatorAddress1)
// require.NoError(t, err)

// 	msgDisableValidator := types.NewMsgDisableValidator(sdk.ValAddress(testconstants.ValidatorAddress1))
// 	_, err := setup.Handler(setup.Ctx, msgDisableValidator)
// 	require.NoError(t, err)

// 	msgDisableValidator = types.NewMsgDisableValidator(sdk.ValAddress(testconstants.ValidatorAddress1))
// 	_, err = setup.Handler(setup.Ctx, msgDisableValidator)
// 	require.Error(t, err)
// 	require.ErrorIs(t, err, types.ErrDisabledValidatorAlreadytExists)
// }

func TestHandler_OwnerNodeAdminCanDisabledValidator(t *testing.T) {
	setup := Setup(t)

	valAddress, err := sdk.ValAddressFromBech32(testconstants.ValidatorAddress1)
	require.NoError(t, err)

	// create Trustee and NodeAdmin
	ba1 := authtypes.NewBaseAccount(sdk.AccAddress(valAddress), testconstants.PubKey2, 0, 0)
	account1 := dclauthtypes.NewAccount(ba1,
		dclauthtypes.AccountRoles{dclauthtypes.NodeAdmin}, nil, nil, testconstants.VendorID2, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account1)

	msgDisableValidator := types.NewMsgDisableValidator(valAddress)
	_, err = setup.Handler(setup.Ctx, msgDisableValidator)
	require.NoError(t, err)

	validator, isFound := setup.ValidatorKeeper.GetDisabledValidator(setup.Ctx, valAddress.String())
	require.True(t, isFound)
	require.Equal(t, valAddress.String(), validator.Address)
	require.True(t, validator.DisabledByNodeAdmin)
}

func TestHandler_OwnerNodeAdminCanEnabledValidator(t *testing.T) {
	setup := Setup(t)

	valAddress, err := sdk.ValAddressFromBech32(testconstants.ValidatorAddress1)
	require.NoError(t, err)

	// create Trustee and NodeAdmin
	ba1 := authtypes.NewBaseAccount(sdk.AccAddress(valAddress), testconstants.PubKey2, 0, 0)
	account1 := dclauthtypes.NewAccount(ba1,
		dclauthtypes.AccountRoles{dclauthtypes.NodeAdmin}, nil, nil, testconstants.VendorID2, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account1)

	// node admin disabled validator
	msgDisableValidator := types.NewMsgDisableValidator(valAddress)
	_, err = setup.Handler(setup.Ctx, msgDisableValidator)
	require.NoError(t, err)

	validator, isFound := setup.ValidatorKeeper.GetDisabledValidator(setup.Ctx, valAddress.String())
	require.True(t, isFound)
	require.Equal(t, valAddress.String(), validator.Address)
	require.True(t, validator.DisabledByNodeAdmin)

	// node admin enabled validator
	msgEnableValidator := types.NewMsgEnableValidator(valAddress)
	_, err = setup.Handler(setup.Ctx, msgEnableValidator)
	require.NoError(t, err)

	_, isFound = setup.ValidatorKeeper.GetDisabledValidator(setup.Ctx, valAddress.String())
	require.False(t, isFound)
}

func TestHandler_TrusteeDisabledValidatorOwnerNodeAdminCanEnableValidator(t *testing.T) {
	setup := Setup(t)

	// create Trustees
	ba1 := authtypes.NewBaseAccount(testconstants.Address1, testconstants.PubKey1, 0, 0)
	account1 := dclauthtypes.NewAccount(ba1,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID1, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account1)

	ba2 := authtypes.NewBaseAccount(testconstants.Address2, testconstants.PubKey2, 0, 0)
	account2 := dclauthtypes.NewAccount(ba2,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID2, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account2)

	ba3 := authtypes.NewBaseAccount(testconstants.Address3, testconstants.PubKey3, 0, 0)
	account3 := dclauthtypes.NewAccount(ba3,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID3, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account3)

	valAddress, err := sdk.ValAddressFromBech32(testconstants.ValidatorAddress1)
	require.NoError(t, err)

	// create NodeAdmin
	ba4 := authtypes.NewBaseAccount(sdk.AccAddress(valAddress), testconstants.PubKey4, 0, 0)
	account4 := dclauthtypes.NewAccount(ba4,
		dclauthtypes.AccountRoles{dclauthtypes.NodeAdmin}, nil, nil, testconstants.VendorID4, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account4)

	// propose new disable validator
	msgProposeDisableValidator1 := NewMsgProposeDisableValidator(account1.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgProposeDisableValidator1)
	require.NoError(t, err)

	// approve new disable validator
	msgProposeDisableValidator2 := NewMsgApproveDisableValidator(account2.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgProposeDisableValidator2)
	require.NoError(t, err)

	// owner node admin can enable disabled validator
	msgEnableValidator := types.NewMsgEnableValidator(valAddress)
	_, err = setup.Handler(setup.Ctx, msgEnableValidator)
	require.NoError(t, err)

	_, isFound := setup.ValidatorKeeper.GetDisabledValidator(setup.Ctx, valAddress.String())
	require.False(t, isFound)
}

func TestHandler_OwnerNodeAdminDisabledValidatorAndNodeAdminCanNotAddNewValidator(t *testing.T) {
	setup := Setup(t)

	valAddress, err := sdk.ValAddressFromBech32(testconstants.ValidatorAddress1)
	require.NoError(t, err)

	// create NodeAdmin
	ba1 := authtypes.NewBaseAccount(sdk.AccAddress(valAddress), testconstants.PubKey1, 0, 0)
	account1 := dclauthtypes.NewAccount(ba1,
		dclauthtypes.AccountRoles{dclauthtypes.NodeAdmin}, nil, nil, testconstants.VendorID1, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account1)

	// node admin disable validator
	msgDisableValidator := types.NewMsgDisableValidator(valAddress)
	_, err = setup.Handler(setup.Ctx, msgDisableValidator)
	require.NoError(t, err)

	msgCreateValidator, err := types.NewMsgCreateValidator(
		valAddress,
		testconstants.ValidatorPubKey2,
		&types.Description{Moniker: testconstants.ProductName},
	)
	require.NoError(t, err)

	// node admin try to add a new validator
	_, err = setup.Handler(setup.Ctx, msgCreateValidator)
	require.Error(t, err)
}

func TestHandler_TrusteeDisabledValidatorAndOwnerNodeAdminCanNotAddNewValidator(t *testing.T) {
	setup := Setup(t)

	// create Trustees
	ba1 := authtypes.NewBaseAccount(testconstants.Address1, testconstants.PubKey1, 0, 0)
	account1 := dclauthtypes.NewAccount(ba1,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID1, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account1)

	ba2 := authtypes.NewBaseAccount(testconstants.Address2, testconstants.PubKey2, 0, 0)
	account2 := dclauthtypes.NewAccount(ba2,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID2, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account2)

	ba3 := authtypes.NewBaseAccount(testconstants.Address3, testconstants.PubKey3, 0, 0)
	account3 := dclauthtypes.NewAccount(ba3,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID3, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account3)

	valAddress, err := sdk.ValAddressFromBech32(testconstants.ValidatorAddress1)
	require.NoError(t, err)

	// create NodeAdmin
	ba4 := authtypes.NewBaseAccount(sdk.AccAddress(valAddress), testconstants.PubKey4, 0, 0)
	account4 := dclauthtypes.NewAccount(ba4,
		dclauthtypes.AccountRoles{dclauthtypes.NodeAdmin}, nil, nil, testconstants.VendorID4, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account4)

	// propose new disable validator
	msgProposeDisableValidator1 := NewMsgProposeDisableValidator(account1.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgProposeDisableValidator1)
	require.NoError(t, err)

	// approve new disable validator
	msgProposeDisableValidator2 := NewMsgApproveDisableValidator(account2.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgProposeDisableValidator2)
	require.NoError(t, err)

	_, isFound := setup.ValidatorKeeper.GetDisabledValidator(setup.Ctx, valAddress.String())
	require.True(t, isFound)

	msgCreateValidator, err := types.NewMsgCreateValidator(
		valAddress,
		testconstants.ValidatorPubKey2,
		&types.Description{Moniker: testconstants.ProductName},
	)
	require.NoError(t, err)

	// node admin try to add a new validator
	_, err = setup.Handler(setup.Ctx, msgCreateValidator)
	require.Error(t, err)
}

func TestHandler_RejectDisableValidator_TwoRejectApprovalsAreNeeded(t *testing.T) {
	setup := Setup(t)

	// create 3 Trustees accounts
	ba1 := authtypes.NewBaseAccount(testconstants.Address1, testconstants.PubKey1, 0, 0)
	account1 := dclauthtypes.NewAccount(ba1,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID1, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account1)

	ba2 := authtypes.NewBaseAccount(testconstants.Address2, testconstants.PubKey2, 0, 0)
	account2 := dclauthtypes.NewAccount(ba2,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID2, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account2)

	ba3 := authtypes.NewBaseAccount(testconstants.Address3, testconstants.PubKey3, 0, 0)
	account3 := dclauthtypes.NewAccount(ba3,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID3, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account3)

	valAddress, err := sdk.ValAddressFromBech32(testconstants.ValidatorAddress1)
	require.NoError(t, err)

	// account1 (Trustee) propose disable validator
	msgProposeDisableValidator := NewMsgProposeDisableValidator(account1.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgProposeDisableValidator)
	require.NoError(t, err)

	// account2 (Trustee) reject disable validator
	msgRejectDisableValidator := NewMsgRejectDisableValidator(account2.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgRejectDisableValidator)
	require.NoError(t, err)

	// validator should be in the entity <Proposed Disable Validator>, because we haven't enough reject approvals
	proposedDisableValidator, isFound := setup.ValidatorKeeper.GetProposedDisableValidator(setup.Ctx, valAddress.String())
	require.True(t, isFound)
	require.Equal(t, msgRejectDisableValidator.Address, proposedDisableValidator.Address)
	require.Equal(t, account1.Address, proposedDisableValidator.Approvals[0].Address)
	require.Equal(t, account2.Address, proposedDisableValidator.Rejects[0].Address)

	// account3 (Trustee) reject disable validator
	msgRejectDisableValidator = NewMsgRejectDisableValidator(account3.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgRejectDisableValidator)
	require.NoError(t, err)

	// validator should not be in the entity <Proposed Disable Validator>, because we have enough reject approvals
	_, isFound = setup.ValidatorKeeper.GetProposedDisableValidator(setup.Ctx, valAddress.String())
	require.False(t, isFound)

	// validator should be in the entity <Rejected Disable Validator>, because we have enough reject approvals
	rejectedDisableValidator, isFound := setup.ValidatorKeeper.GetRejectedNode(setup.Ctx, msgRejectDisableValidator.Address)
	require.True(t, isFound)
	require.Equal(t, msgRejectDisableValidator.Address, proposedDisableValidator.Address)
	require.Equal(t, account1.Address, rejectedDisableValidator.Approvals[0].Address)
	require.Equal(t, account2.Address, rejectedDisableValidator.Rejects[0].Address)
	require.Equal(t, account3.Address, rejectedDisableValidator.Rejects[1].Address)
}

func TestHandler_RejectDisableValidator_ByNotTrustee(t *testing.T) {
	setup := Setup(t)

	// create 3 Trustee accounts
	ba1 := authtypes.NewBaseAccount(testconstants.Address1, testconstants.PubKey1, 0, 0)
	account1 := dclauthtypes.NewAccount(ba1,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID1, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account1)

	ba2 := authtypes.NewBaseAccount(testconstants.Address2, testconstants.PubKey2, 0, 0)
	account2 := dclauthtypes.NewAccount(ba2,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID2, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account2)

	ba3 := authtypes.NewBaseAccount(testconstants.Address3, testconstants.PubKey3, 0, 0)
	account3 := dclauthtypes.NewAccount(ba3,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID3, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account3)

	// create vendor account
	ba4 := authtypes.NewBaseAccount(testconstants.Address4, testconstants.PubKey4, 0, 0)
	account4 := dclauthtypes.NewAccount(ba4,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor}, nil, nil, testconstants.VendorID4, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account4)

	valAddress, err := sdk.ValAddressFromBech32(testconstants.ValidatorAddress1)
	require.NoError(t, err)

	// account1 (Trustee) propose disable validator
	msgProposeDisableValidator := NewMsgProposeDisableValidator(account1.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgProposeDisableValidator)
	require.NoError(t, err)

	// ensure that account is in <Proposed Disable Validator>
	_, isFound := setup.ValidatorKeeper.GetProposedDisableValidator(setup.Ctx, valAddress.String())
	require.True(t, isFound)

	// not Trustee account try to reject disable validator
	msgRejectDisableValidator := NewMsgRejectDisableValidator(account4.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgRejectDisableValidator)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestHandler_RejectDisableValidator_ForUnknownAccount(t *testing.T) {
	setup := Setup(t)

	// create Trustee accounts
	ba1 := authtypes.NewBaseAccount(testconstants.Address1, testconstants.PubKey1, 0, 0)
	account1 := dclauthtypes.NewAccount(ba1,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID1, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account1)

	valAddress, err := sdk.ValAddressFromBech32(testconstants.ValidatorAddress1)
	require.NoError(t, err)

	// reject disable from unknown account
	msgRejectDisableValidator := NewMsgRejectDisableValidator(testconstants.Address4, valAddress)
	_, err = setup.Handler(setup.Ctx, msgRejectDisableValidator)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestHandler_Duplicate_RejectDisableValidatorFromTheSameTrustee(t *testing.T) {
	setup := Setup(t)

	// create 3 Trustee accounts
	ba1 := authtypes.NewBaseAccount(testconstants.Address1, testconstants.PubKey1, 0, 0)
	account1 := dclauthtypes.NewAccount(ba1,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID1, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account1)

	ba2 := authtypes.NewBaseAccount(testconstants.Address2, testconstants.PubKey2, 0, 0)
	account2 := dclauthtypes.NewAccount(ba2,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID2, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account2)

	ba3 := authtypes.NewBaseAccount(testconstants.Address3, testconstants.PubKey3, 0, 0)
	account3 := dclauthtypes.NewAccount(ba3,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID3, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account3)

	valAddress, err := sdk.ValAddressFromBech32(testconstants.ValidatorAddress1)
	require.NoError(t, err)

	// account1 (Trustee) propose disable validator
	msgProposeDisableValidator := NewMsgProposeDisableValidator(account1.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgProposeDisableValidator)
	require.NoError(t, err)

	// ensure that account is in <Proposed Disable Validator>
	_, isFound := setup.ValidatorKeeper.GetProposedDisableValidator(setup.Ctx, valAddress.String())
	require.True(t, isFound)

	// account2 (Trustee) reject disable validator
	msgRejectDisableValidator := NewMsgRejectDisableValidator(account2.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgRejectDisableValidator)
	require.NoError(t, err)

	// validator should be in the entity <Proposed Disable Validator>, because we haven't enough reject approvals
	_, isFound = setup.ValidatorKeeper.GetProposedDisableValidator(setup.Ctx, valAddress.String())
	require.True(t, isFound)

	// account2 (Trustee) try to second time reject disable validator
	msgRejectDisableValidator = NewMsgRejectDisableValidator(account2.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgRejectDisableValidator)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestHandler_ApproveDisableValidatorAndRejectDisableValidator_FromTheSamerTrustee(t *testing.T) {
	setup := Setup(t)

	// create 4 Trustee accounts
	ba1 := authtypes.NewBaseAccount(testconstants.Address1, testconstants.PubKey1, 0, 0)
	account1 := dclauthtypes.NewAccount(ba1,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID1, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account1)

	ba2 := authtypes.NewBaseAccount(testconstants.Address2, testconstants.PubKey2, 0, 0)
	account2 := dclauthtypes.NewAccount(ba2,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID2, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account2)

	ba3 := authtypes.NewBaseAccount(testconstants.Address3, testconstants.PubKey3, 0, 0)
	account3 := dclauthtypes.NewAccount(ba3,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID3, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account3)

	ba4 := authtypes.NewBaseAccount(testconstants.Address4, testconstants.PubKey4, 0, 0)
	account4 := dclauthtypes.NewAccount(ba4,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID4, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account4)

	valAddress, err := sdk.ValAddressFromBech32(testconstants.ValidatorAddress1)
	require.NoError(t, err)

	// account1 (Trustee) propose disable validator
	msgProposeDisableValidator := NewMsgProposeDisableValidator(account1.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgProposeDisableValidator)
	require.NoError(t, err)

	// ensure that account is in <Proposed Disable Validator>
	_, isFound := setup.ValidatorKeeper.GetProposedDisableValidator(setup.Ctx, valAddress.String())
	require.True(t, isFound)

	// account2 (Trustee) approve disable validator
	msgApproveDisableValidator := NewMsgApproveDisableValidator(account2.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgApproveDisableValidator)
	require.NoError(t, err)

	// ensure that account is in <Proposed Disable Validator>
	proposedDisableValidator, isFound := setup.ValidatorKeeper.GetProposedDisableValidator(setup.Ctx, valAddress.String())
	require.True(t, isFound)
	prevRejectsLen := len(proposedDisableValidator.Rejects)
	prevApprovalsLen := len(proposedDisableValidator.Approvals)
	// account2 (Trustee) try rejects disable validator
	msgRejectDisableValidator := NewMsgRejectDisableValidator(account2.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgRejectDisableValidator)
	require.NoError(t, err)

	proposedDisableValidator, isFound = setup.ValidatorKeeper.GetProposedDisableValidator(setup.Ctx, valAddress.String())
	require.True(t, isFound)
	require.Equal(t, len(proposedDisableValidator.Rejects), prevRejectsLen+1)
	require.Equal(t, len(proposedDisableValidator.Approvals), prevApprovalsLen-1)
}

func TestHandler_RejectDisableValidatorAndApproveDisableValidator_FromTheSameTrustee(t *testing.T) {
	setup := Setup(t)

	// create 4 Trustee accounts
	ba1 := authtypes.NewBaseAccount(testconstants.Address1, testconstants.PubKey1, 0, 0)
	account1 := dclauthtypes.NewAccount(ba1,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID1, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account1)

	ba2 := authtypes.NewBaseAccount(testconstants.Address2, testconstants.PubKey2, 0, 0)
	account2 := dclauthtypes.NewAccount(ba2,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID2, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account2)

	ba3 := authtypes.NewBaseAccount(testconstants.Address3, testconstants.PubKey3, 0, 0)
	account3 := dclauthtypes.NewAccount(ba3,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID3, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account3)

	ba4 := authtypes.NewBaseAccount(testconstants.Address4, testconstants.PubKey4, 0, 0)
	account4 := dclauthtypes.NewAccount(ba4,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID4, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account4)

	valAddress, err := sdk.ValAddressFromBech32(testconstants.ValidatorAddress1)
	require.NoError(t, err)

	// account1 (Trustee) propose disable validator
	msgProposeDisableValidator := NewMsgProposeDisableValidator(account1.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgProposeDisableValidator)
	require.NoError(t, err)

	// ensure that account is in <Proposed Disable Validator>
	_, isFound := setup.ValidatorKeeper.GetProposedDisableValidator(setup.Ctx, valAddress.String())
	require.True(t, isFound)

	// account2 (Trustee) reject disable validator
	msgRejectDisableValidator := NewMsgRejectDisableValidator(account2.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgRejectDisableValidator)
	require.NoError(t, err)

	// ensure that account is in <Proposed Disable Validator>
	proposedDisableValidator, isFound := setup.ValidatorKeeper.GetProposedDisableValidator(setup.Ctx, valAddress.String())
	require.True(t, isFound)
	prevRejectsLen := len(proposedDisableValidator.Rejects)
	prevApprovalsLen := len(proposedDisableValidator.Approvals)

	// account2 (Trustee) try approve disable validator
	msgApproveDisableValidator := NewMsgApproveDisableValidator(account2.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgApproveDisableValidator)
	require.NoError(t, err)

	proposedDisableValidator, isFound = setup.ValidatorKeeper.GetProposedDisableValidator(setup.Ctx, valAddress.String())
	require.True(t, isFound)
	require.Equal(t, len(proposedDisableValidator.Rejects), prevRejectsLen-1)
	require.Equal(t, len(proposedDisableValidator.Approvals), prevApprovalsLen+1)
}

func TestHandler_DoubleTimeRejectDisableValidator(t *testing.T) {
	setup := Setup(t)

	// create 3 Trustee accounts
	ba1 := authtypes.NewBaseAccount(testconstants.Address1, testconstants.PubKey1, 0, 0)
	account1 := dclauthtypes.NewAccount(ba1,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID1, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account1)

	ba2 := authtypes.NewBaseAccount(testconstants.Address2, testconstants.PubKey2, 0, 0)
	account2 := dclauthtypes.NewAccount(ba2,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID2, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account2)

	ba3 := authtypes.NewBaseAccount(testconstants.Address3, testconstants.PubKey3, 0, 0)
	account3 := dclauthtypes.NewAccount(ba3,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID3, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account3)

	valAddress, err := sdk.ValAddressFromBech32(testconstants.ValidatorAddress1)
	require.NoError(t, err)

	// account1 (Trustee) propose disable validator
	msgProposeDisableValidator := NewMsgProposeDisableValidator(account1.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgProposeDisableValidator)
	require.NoError(t, err)

	// ensure that account is in <Proposed Disable Validator>
	_, isFound := setup.ValidatorKeeper.GetProposedDisableValidator(setup.Ctx, valAddress.String())
	require.True(t, isFound)

	// account2 (Trustee) rejects disable validator
	msgRejectDisableValidator := NewMsgRejectDisableValidator(account2.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgRejectDisableValidator)
	require.NoError(t, err)

	// ensure that account is in <Proposed Disable Validator>
	proposedDisableValidator, isFound := setup.ValidatorKeeper.GetProposedDisableValidator(setup.Ctx, valAddress.String())
	require.True(t, isFound)

	// account3 (Trustee) rejects disable same validator
	msgRejectDisableValidator = NewMsgRejectDisableValidator(account3.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgRejectDisableValidator)
	require.NoError(t, err)

	// validator should not be in the entity <Proposed Disable Validator>, because we have enough reject approvals
	_, isFound = setup.ValidatorKeeper.GetProposedDisableValidator(setup.Ctx, valAddress.String())
	require.False(t, isFound)

	// validator should be in the entity <Rejected Disable Validator>, because we have enough reject approvals
	rejectedDisableValidator, isFound := setup.ValidatorKeeper.GetRejectedNode(setup.Ctx, msgRejectDisableValidator.Address)
	require.True(t, isFound)
	require.Equal(t, msgRejectDisableValidator.Address, proposedDisableValidator.Address)
	require.Equal(t, account1.Address, rejectedDisableValidator.Approvals[0].Address)
	require.Equal(t, account2.Address, rejectedDisableValidator.Rejects[0].Address)
	require.Equal(t, account3.Address, rejectedDisableValidator.Rejects[1].Address)

	// account1 (Trustee) proposes disable validator
	msgProposeDisableValidator = NewMsgProposeDisableValidator(account1.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgProposeDisableValidator)
	require.NoError(t, err)

	// ensure that account is in <Proposed Disable Validator>
	_, isFound = setup.ValidatorKeeper.GetProposedDisableValidator(setup.Ctx, valAddress.String())
	require.True(t, isFound)

	// ensure that account not exist in <Rejected Disable Validator>
	_, isFound = setup.ValidatorKeeper.GetRejectedNode(setup.Ctx, valAddress.String())
	require.False(t, isFound)

	// account3 (Trustee) rejects disable same validator
	msgRejectDisableValidator = NewMsgRejectDisableValidator(account3.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgRejectDisableValidator)
	require.NoError(t, err)

	// ensure that account is in <Proposed Disable Validator>
	proposedDisableValidator, isFound = setup.ValidatorKeeper.GetProposedDisableValidator(setup.Ctx, valAddress.String())
	require.True(t, isFound)

	// account2 (Trustee) rejects disable validator
	msgRejectDisableValidator = NewMsgRejectDisableValidator(account2.GetAddress(), valAddress)
	_, err = setup.Handler(setup.Ctx, msgRejectDisableValidator)
	require.NoError(t, err)

	// validator should not be in the entity <Proposed Disable Validator>, because we have enough reject approvals
	_, isFound = setup.ValidatorKeeper.GetProposedDisableValidator(setup.Ctx, valAddress.String())
	require.False(t, isFound)

	// validator should be in the entity <Rejected Disable Validator>, because we have enough reject approvals
	rejectedDisableValidator, isFound = setup.ValidatorKeeper.GetRejectedNode(setup.Ctx, msgRejectDisableValidator.Address)
	require.True(t, isFound)
	require.Equal(t, msgRejectDisableValidator.Address, proposedDisableValidator.Address)
	require.Equal(t, account1.Address, rejectedDisableValidator.Approvals[0].Address)
	require.Equal(t, account3.Address, rejectedDisableValidator.Rejects[0].Address)
	require.Equal(t, account2.Address, rejectedDisableValidator.Rejects[1].Address)
}

func TestHandler_RejectDisableValidator_TwoRejectApprovalsAreNeeded_FiveTrustees(t *testing.T) {
	setup := Setup(t)

	// we have 5 trustees: 1 approval comes from propose => we need 2 rejects to make validator disabled

	// create Trustees
	ba1 := authtypes.NewBaseAccount(testconstants.Address1, testconstants.PubKey1, 0, 0)
	account1 := dclauthtypes.NewAccount(ba1,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID1, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account1)

	ba2 := authtypes.NewBaseAccount(testconstants.Address2, testconstants.PubKey2, 0, 0)
	account2 := dclauthtypes.NewAccount(ba2,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID2, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account2)

	ba3 := authtypes.NewBaseAccount(testconstants.Address3, testconstants.PubKey3, 0, 0)
	account3 := dclauthtypes.NewAccount(ba3,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID3, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account3)

	ba4 := authtypes.NewBaseAccount(testconstants.Address4, testconstants.PubKey4, 0, 0)
	account4 := dclauthtypes.NewAccount(ba4,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID4, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account4)

	_, ba5PubKey, ba5Address := testdata.KeyTestPubAddr()
	ba5 := authtypes.NewBaseAccount(ba5Address, ba5PubKey, 0, 0)
	account5 := dclauthtypes.NewAccount(ba5,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID4, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account5)

	valAddress, err := sdk.ValAddressFromBech32(testconstants.ValidatorAddress1)
	require.NoError(t, err)

	// propose disable validator by account Trustee1
	proposeDisableValidator := types.NewMsgProposeDisableValidator(ba1.GetAddress(), valAddress, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, proposeDisableValidator)
	require.NoError(t, err)

	// reject disable validator by account Trustee2
	rejectDisableValidator := types.NewMsgRejectDisableValidator(ba2.GetAddress(), valAddress, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectDisableValidator)
	require.NoError(t, err)

	// disable validator should be in the entity <Proposed Disable validator>, because we haven't enough reject approvals
	proposedDisableValidator, found := setup.ValidatorKeeper.GetProposedDisableValidator(setup.Ctx, valAddress.String())
	require.True(t, found)

	// check proposed disable validator
	require.Equal(t, proposeDisableValidator.Address, proposedDisableValidator.Address)
	require.Equal(t, proposeDisableValidator.Creator, proposedDisableValidator.Creator)

	// reject disable validator by account Trustee3
	rejectDisableValidator = types.NewMsgRejectDisableValidator(ba3.GetAddress(), valAddress, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectDisableValidator)
	require.NoError(t, err)

	// validator should not be in the entity <Disabled Validator>, because we have enough reject approvals
	_, found = setup.ValidatorKeeper.GetValidator(setup.Ctx, valAddress)
	require.True(t, found)

	_, found = setup.ValidatorKeeper.GetDisabledValidator(setup.Ctx, valAddress.String())
	require.False(t, found)
}

func TestHandler_ApproveDisableValidator_FourApprovalsAreNeeded_FiveTrustees(t *testing.T) {
	setup := Setup(t)

	// we have 5 trustees: 1 approval comes from propose => we need 3 more to disable validator

	// create Trustees
	ba1 := authtypes.NewBaseAccount(testconstants.Address1, testconstants.PubKey1, 0, 0)
	account1 := dclauthtypes.NewAccount(ba1,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID1, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account1)

	ba2 := authtypes.NewBaseAccount(testconstants.Address2, testconstants.PubKey2, 0, 0)
	account2 := dclauthtypes.NewAccount(ba2,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID2, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account2)

	ba3 := authtypes.NewBaseAccount(testconstants.Address3, testconstants.PubKey3, 0, 0)
	account3 := dclauthtypes.NewAccount(ba3,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID3, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account3)

	ba4 := authtypes.NewBaseAccount(testconstants.Address4, testconstants.PubKey4, 0, 0)
	account4 := dclauthtypes.NewAccount(ba4,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID4, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account4)

	_, ba5PubKey, ba5Address := testdata.KeyTestPubAddr()
	ba5 := authtypes.NewBaseAccount(ba5Address, ba5PubKey, 0, 0)
	account5 := dclauthtypes.NewAccount(ba5,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, nil, testconstants.VendorID4, testconstants.ProductIDsEmpty)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account5)

	valAddress, err := sdk.ValAddressFromBech32(testconstants.ValidatorAddress1)
	require.NoError(t, err)

	// propose disable validator by account Trustee1
	proposeDisableValidator := types.NewMsgProposeDisableValidator(ba1.GetAddress(), valAddress, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, proposeDisableValidator)
	require.NoError(t, err)

	// approve disable validator by account Trustee2
	approveDisableValidator := types.NewMsgApproveDisableValidator(ba2.GetAddress(), valAddress, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveDisableValidator)
	require.NoError(t, err)

	// approve disable validator by account Trustee3
	approveDisableValidator = types.NewMsgApproveDisableValidator(ba3.GetAddress(), valAddress, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveDisableValidator)
	require.NoError(t, err)

	// reject disable validator by account Trustee4
	rejectDisableValidator := types.NewMsgRejectDisableValidator(ba4.GetAddress(), valAddress, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, rejectDisableValidator)
	require.NoError(t, err)

	// disable validator should be in the entity <Proposed Disable validator>, because we haven't enough reject approvals
	proposedDisableValidator, found := setup.ValidatorKeeper.GetProposedDisableValidator(setup.Ctx, valAddress.String())
	require.True(t, found)

	// check proposed disable validator
	require.Equal(t, proposeDisableValidator.Address, proposedDisableValidator.Address)
	require.Equal(t, proposeDisableValidator.Creator, proposedDisableValidator.Creator)

	// approve disable validator by account Trustee5
	approveDisableValidator = types.NewMsgApproveDisableValidator(ba5.GetAddress(), valAddress, testconstants.Info)
	_, err = setup.Handler(setup.Ctx, approveDisableValidator)
	require.NoError(t, err)

	// validator should be in the entity <Disabled Validator>, because we have enough reject approvals
	_, found = setup.ValidatorKeeper.GetValidator(setup.Ctx, valAddress)
	require.True(t, found)

	_, found = setup.ValidatorKeeper.GetDisabledValidator(setup.Ctx, valAddress.String())
	require.True(t, found)
}

func NewMsgProposeDisableValidator(signer sdk.AccAddress, address sdk.ValAddress) *types.MsgProposeDisableValidator {
	return &types.MsgProposeDisableValidator{
		Creator: signer.String(),
		Address: address.String(),
		Time:    testconstants.Time,
		Info:    testconstants.Info,
	}
}

func NewMsgApproveDisableValidator(signer sdk.AccAddress, address sdk.ValAddress) *types.MsgApproveDisableValidator {
	return &types.MsgApproveDisableValidator{
		Creator: signer.String(),
		Address: address.String(),
		Time:    testconstants.Time,
		Info:    testconstants.Info,
	}
}

func NewMsgRejectDisableValidator(signer sdk.AccAddress, address sdk.ValAddress) *types.MsgRejectDisableValidator {
	return &types.MsgRejectDisableValidator{
		Creator: signer.String(),
		Address: address.String(),
		Time:    testconstants.Time,
		Info:    testconstants.Info,
	}
}
