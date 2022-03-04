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
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	sdkstakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	testkeeper "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/testdata"
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
	dclauthK, _ := testkeeper.DclauthKeeper(t)
	k, ctx := testkeeper.ValidatorKeeper(t, dclauthK)

	ba := authtypes.NewBaseAccount(testconstants.Address1, testconstants.PubKey1, 0, 0)
	account := dclauthtypes.NewAccount(
		ba,
		dclauthtypes.AccountRoles{dclauthtypes.NodeAdmin}, nil, testconstants.VendorID1,
	)
	dclauthK.SetAccount(ctx, account)

	setup := TestSetup{
		Ctx:             ctx,
		ValidatorKeeper: *k,
		DclauthKeeper:   *dclauthK,
		Handler:         NewHandler(*k),
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
		account := dclauthtypes.NewAccount(ba, dclauthtypes.AccountRoles{role}, nil, testconstants.VendorID1)
		setup.DclauthKeeper.SetAccount(setup.Ctx, account)

		// try to create validator
		_, err := setup.Handler(setup.Ctx, msgCreateValidator)
		require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
	}
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
		dclauthtypes.AccountRoles{dclauthtypes.NodeAdmin}, nil, testconstants.VendorID2)
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
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, testconstants.VendorID1)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account1)

	ba2 := authtypes.NewBaseAccount(testconstants.Address2, testconstants.PubKey2, 0, 0)
	account2 := dclauthtypes.NewAccount(ba2,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, testconstants.VendorID2)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account2)

	// propose new disablevalidator
	msgProposeDisableValidator := NewMsgProposeDisableValidator(account1.GetAddress())
	_, err := setup.Handler(setup.Ctx, msgProposeDisableValidator)
	require.NoError(t, err)

	msgProposeDisableValidator.Creator = account2.GetAddress().String()
	// propose the same disablevalidator
	_, err = setup.Handler(setup.Ctx, msgProposeDisableValidator)
	require.Error(t, err)
	require.True(t, types.ErrProposedDisableValidatorAlreadyExists.Is(err))
}

func TestHandler_OnlyTrusteeCanProposeDisableValidator(t *testing.T) {
	setup := Setup(t)

	// create Trustees
	ba1 := authtypes.NewBaseAccount(testconstants.Address1, testconstants.PubKey1, 0, 0)
	account1 := dclauthtypes.NewAccount(ba1,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee, dclauthtypes.CertificationCenter, dclauthtypes.Vendor}, nil, testconstants.VendorID1)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account1)

	// propose new disablevalidator
	msgProposeDisableValidator := NewMsgProposeDisableValidator(account1.GetAddress())
	_, err := setup.Handler(setup.Ctx, msgProposeDisableValidator)
	require.Error(t, err)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestHandler_ProposeDisableValidatorWhenSeveralVotesNeeded(t *testing.T) {
	setup := Setup(t)

	// create Trustees
	ba1 := authtypes.NewBaseAccount(testconstants.Address1, testconstants.PubKey1, 0, 0)
	account1 := dclauthtypes.NewAccount(ba1,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, testconstants.VendorID1)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account1)

	ba2 := authtypes.NewBaseAccount(testconstants.Address2, testconstants.PubKey2, 0, 0)
	account2 := dclauthtypes.NewAccount(ba2,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, testconstants.VendorID2)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account2)

	ba3 := authtypes.NewBaseAccount(testconstants.Address3, testconstants.PubKey3, 0, 0)
	account3 := dclauthtypes.NewAccount(ba3,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, testconstants.VendorID3)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account3)

	// propose new disablevalidator
	msgProposeDisableValidator := NewMsgProposeDisableValidator(account1.GetAddress())
	_, err := setup.Handler(setup.Ctx, msgProposeDisableValidator)
	require.NoError(t, err)

	proposedDisableValidator, isFound := setup.ValidatorKeeper.GetProposedDisableValidator(setup.Ctx, msgProposeDisableValidator.Address)
	require.True(t, isFound)
	require.Equal(t, msgProposeDisableValidator.Address, proposedDisableValidator.Address)
	require.Equal(t, msgProposeDisableValidator.Creator, proposedDisableValidator.Creator)
	require.Equal(t, msgProposeDisableValidator.Creator, proposedDisableValidator.Approvals[0])

	_, isFound = setup.ValidatorKeeper.GetDisabledValidator(setup.Ctx, msgProposeDisableValidator.Address)
	require.False(t, isFound)
}

func TestHandler_DisabledValidator(t *testing.T) {
	setup := Setup(t)

	// create Trustees
	ba1 := authtypes.NewBaseAccount(testconstants.Address1, testconstants.PubKey1, 0, 0)
	account1 := dclauthtypes.NewAccount(ba1,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, testconstants.VendorID1)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account1)

	ba2 := authtypes.NewBaseAccount(testconstants.Address2, testconstants.PubKey2, 0, 0)
	account2 := dclauthtypes.NewAccount(ba2,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, testconstants.VendorID2)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account2)

	ba3 := authtypes.NewBaseAccount(testconstants.Address3, testconstants.PubKey3, 0, 0)
	account3 := dclauthtypes.NewAccount(ba3,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, testconstants.VendorID3)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account3)

	// propose and approve new disablevalidators
	msgProposeDisableValidator1 := NewMsgProposeDisableValidator(account1.GetAddress())
	_, err := setup.Handler(setup.Ctx, msgProposeDisableValidator1)
	require.NoError(t, err)

	msgApproveDisableValidator := NewMsgApproveDisableValidator(account2.GetAddress())
	_, err = setup.Handler(setup.Ctx, msgApproveDisableValidator)
	require.NoError(t, err)

	_, isFound := setup.ValidatorKeeper.GetProposedDisableValidator(setup.Ctx, msgProposeDisableValidator1.Address)
	require.False(t, isFound)

	disabledValidator, isFound := setup.ValidatorKeeper.GetDisabledValidator(setup.Ctx, msgProposeDisableValidator1.Address)
	require.True(t, isFound)
	require.Equal(t, msgProposeDisableValidator1.Address, disabledValidator.Address)
	require.Equal(t, msgProposeDisableValidator1.Creator, disabledValidator.Creator)
	require.Equal(t, msgProposeDisableValidator1.Creator, disabledValidator.Approvals[0])
	require.False(t, disabledValidator.DisabledByNodeAdmin)
}

func TestHandler_ProposeDisableValidatorWhenMoreVotesNeeded(t *testing.T) {
	setup := Setup(t)

	// create Trustees
	ba1 := authtypes.NewBaseAccount(testconstants.Address1, testconstants.PubKey1, 0, 0)
	account1 := dclauthtypes.NewAccount(ba1,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, testconstants.VendorID1)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account1)

	ba2 := authtypes.NewBaseAccount(testconstants.Address2, testconstants.PubKey2, 0, 0)
	account2 := dclauthtypes.NewAccount(ba2,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, testconstants.VendorID2)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account2)

	ba3 := authtypes.NewBaseAccount(testconstants.Address3, testconstants.PubKey3, 0, 0)
	account3 := dclauthtypes.NewAccount(ba3,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, testconstants.VendorID3)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account3)

	ba4 := authtypes.NewBaseAccount(testdata.GenerateAccAddress(), testconstants.PubKey4, 0, 0)
	account4 := dclauthtypes.NewAccount(ba4,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, 4000)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account4)

	// propose and approve new disablevalidators
	msgProposeDisableValidator := NewMsgProposeDisableValidator(account1.GetAddress())
	_, err := setup.Handler(setup.Ctx, msgProposeDisableValidator)
	require.NoError(t, err)

	msgApproveDisableValidator := NewMsgApproveDisableValidator(account2.GetAddress())
	_, err = setup.Handler(setup.Ctx, msgApproveDisableValidator)
	require.NoError(t, err)

	proposedDisableValidator, isFound := setup.ValidatorKeeper.GetProposedDisableValidator(setup.Ctx, msgProposeDisableValidator.Address)
	require.True(t, isFound)
	require.Equal(t, msgProposeDisableValidator.Address, proposedDisableValidator.Address)
	require.Equal(t, msgProposeDisableValidator.Creator, proposedDisableValidator.Creator)
	require.Equal(t, msgProposeDisableValidator.Creator, proposedDisableValidator.Approvals[0])
	require.Equal(t, 2, len(proposedDisableValidator.Approvals))

	_, isFound = setup.ValidatorKeeper.GetDisabledValidator(setup.Ctx, msgProposeDisableValidator.Address)
	require.False(t, isFound)
}

func TestHandler_OnlyTrusteeCanApproveDisableValidator(t *testing.T) {
	setup := Setup(t)

	// create Trustees
	ba1 := authtypes.NewBaseAccount(testconstants.Address1, testconstants.PubKey1, 0, 0)
	account1 := dclauthtypes.NewAccount(ba1,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee, dclauthtypes.CertificationCenter, dclauthtypes.Vendor}, nil, testconstants.VendorID1)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account1)

	ba2 := authtypes.NewBaseAccount(testconstants.Address2, testconstants.PubKey2, 0, 0)
	account2 := dclauthtypes.NewAccount(ba2,
		dclauthtypes.AccountRoles{dclauthtypes.NodeAdmin, dclauthtypes.CertificationCenter, dclauthtypes.Vendor}, nil, testconstants.VendorID2)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account2)

	ba3 := authtypes.NewBaseAccount(testconstants.Address3, testconstants.PubKey3, 0, 0)
	account3 := dclauthtypes.NewAccount(ba3,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, testconstants.VendorID3)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account3)

	// propose and approve new disablevalidator
	msgProposeDisableValidator := NewMsgProposeDisableValidator(account1.GetAddress())
	_, err := setup.Handler(setup.Ctx, msgProposeDisableValidator)
	require.NoError(t, err)

	msgApproveDisableValidator := NewMsgApproveDisableValidator(account2.GetAddress())
	_, err = setup.Handler(setup.Ctx, msgApproveDisableValidator)
	require.Error(t, err)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestHandler_ProposedDisableValidatorDoesNotExist(t *testing.T) {
	setup := Setup(t)

	// create Trustees
	ba1 := authtypes.NewBaseAccount(testconstants.Address1, testconstants.PubKey1, 0, 0)
	account1 := dclauthtypes.NewAccount(ba1,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee, dclauthtypes.CertificationCenter, dclauthtypes.Vendor}, nil, testconstants.VendorID1)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account1)

	msgApproveDisableValidator := NewMsgApproveDisableValidator(account1.GetAddress())
	_, err := setup.Handler(setup.Ctx, msgApproveDisableValidator)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrProposedDisableValidatorDoesNotExist)
}

func TestHandler_MessageCreatorAlreadyApprovedDisableValidator(t *testing.T) {
	setup := Setup(t)

	// create Trustees
	ba1 := authtypes.NewBaseAccount(testconstants.Address1, testconstants.PubKey1, 0, 0)
	account1 := dclauthtypes.NewAccount(ba1,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, testconstants.VendorID1)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account1)

	ba2 := authtypes.NewBaseAccount(testconstants.Address2, testconstants.PubKey2, 0, 0)
	account2 := dclauthtypes.NewAccount(ba2,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, testconstants.VendorID2)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account2)

	ba3 := authtypes.NewBaseAccount(testconstants.Address3, testconstants.PubKey3, 0, 0)
	account3 := dclauthtypes.NewAccount(ba3,
		dclauthtypes.AccountRoles{dclauthtypes.Trustee}, nil, testconstants.VendorID3)
	setup.DclauthKeeper.SetAccount(setup.Ctx, account3)

	// propose and approve new disablevalidator
	msgProposeDisableValidator := NewMsgProposeDisableValidator(account1.GetAddress())
	_, err := setup.Handler(setup.Ctx, msgProposeDisableValidator)
	require.NoError(t, err)

	msgApproveDisableValidator := NewMsgApproveDisableValidator(account1.GetAddress())
	_, err = setup.Handler(setup.Ctx, msgApproveDisableValidator)
	require.Error(t, err)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func NewMsgProposeDisableValidator(signer sdk.AccAddress) *types.MsgProposeDisableValidator {
	return &types.MsgProposeDisableValidator{
		Creator: signer.String(),
		Address: testconstants.ValidatorAddress1,
	}
}

func NewMsgApproveDisableValidator(signer sdk.AccAddress) *types.MsgApproveDisableValidator {
	return &types.MsgApproveDisableValidator{
		Creator: signer.String(),
		Address: testconstants.ValidatorAddress1,
	}
}
