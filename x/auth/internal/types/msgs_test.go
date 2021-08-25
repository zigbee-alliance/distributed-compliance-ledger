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
package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
)

/*
	MsgProposeAddAccount
*/

func TestNewMsgProposeAddAccount(t *testing.T) {
	msg := NewMsgProposeAddAccount(testconstants.Address1, testconstants.Pubkey1Str,
		AccountRoles{}, testconstants.VendorId1, testconstants.Signer)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "propose_add_account")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}

func TestValidateMsgProposeAddAccount(t *testing.T) {
	cases := []struct {
		valid bool
		msg   MsgProposeAddAccount
	}{
		{true, NewMsgProposeAddAccount(testconstants.Address1, testconstants.Pubkey1Str,
			AccountRoles{}, 0, testconstants.Signer)},
		{true, NewMsgProposeAddAccount(testconstants.Address1, testconstants.Pubkey1Str,
			AccountRoles{Vendor, NodeAdmin}, testconstants.VendorId1, testconstants.Signer)},
		{true, NewMsgProposeAddAccount(testconstants.Address1, testconstants.Pubkey1Str,
			AccountRoles{Vendor, NodeAdmin}, testconstants.VendorId1, testconstants.Signer)},
		{false, NewMsgProposeAddAccount(nil, testconstants.Pubkey1Str,
			AccountRoles{}, 0, testconstants.Signer)},
		{false, NewMsgProposeAddAccount(testconstants.Address1, "",
			AccountRoles{}, 0, testconstants.Signer)},
		{false, NewMsgProposeAddAccount(testconstants.Address1, testconstants.Pubkey1Str,
			AccountRoles{"Wrong Role"}, 0, testconstants.Signer)},
		{false, NewMsgProposeAddAccount(testconstants.Address1, testconstants.Pubkey1Str,
			AccountRoles{}, 0, nil)},
	}

	for _, tc := range cases {
		err := tc.msg.ValidateBasic()

		if tc.valid {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
		}
	}
}

func TestMsgProposeAddAccountGetSignBytes(t *testing.T) {
	msg := NewMsgProposeAddAccount(testconstants.Address1, testconstants.Pubkey1Str,
		AccountRoles{}, testconstants.VendorId1, testconstants.Signer)

	expected := `{"type":"auth/ProposeAddAccount","value":{"address":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz",` +
		`"pub_key":"cosmospub1addwnpepq28rlfval9n8khmgqz55mlfwn4rlh0jk80k9n7fvtu4g4u37qtvry76ww9h","roles":[],` +
		`"signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz","vendorId":1000}}`
	require.Equal(t, expected, string(msg.GetSignBytes()))
}

/*
	MsgApproveAddAccount
*/

func TestNewMsgApproveAddAccount(t *testing.T) {
	msg := NewMsgApproveAddAccount(testconstants.Address1, testconstants.Signer)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "approve_add_account")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}

func TestValidateMsgApproveAddAccount(t *testing.T) {
	cases := []struct {
		valid bool
		msg   MsgApproveAddAccount
	}{
		{true, NewMsgApproveAddAccount(testconstants.Address1, testconstants.Signer)},
		{false, NewMsgApproveAddAccount(nil, testconstants.Signer)},
		{false, NewMsgApproveAddAccount(testconstants.Address1, nil)},
	}

	for _, tc := range cases {
		err := tc.msg.ValidateBasic()

		if tc.valid {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
		}
	}
}

func TestMsgApproveAddAccountGetSignBytes(t *testing.T) {
	msg := NewMsgApproveAddAccount(testconstants.Address1, testconstants.Signer)

	expected := `{"type":"auth/ApproveAddAccount","value":{"address":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz",` +
		`"signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz"}}`
	require.Equal(t, expected, string(msg.GetSignBytes()))
}

/*
	MsgProposeRevokeAccount
*/

func TestNewMsgProposeRevokeAccount(t *testing.T) {
	msg := NewMsgProposeRevokeAccount(testconstants.Address1, testconstants.Signer)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "propose_revoke_account")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}

func TestValidateMsgProposeRevokeAccount(t *testing.T) {
	cases := []struct {
		valid bool
		msg   MsgProposeRevokeAccount
	}{
		{true, NewMsgProposeRevokeAccount(testconstants.Address1, testconstants.Signer)},
		{false, NewMsgProposeRevokeAccount(nil, testconstants.Signer)},
		{false, NewMsgProposeRevokeAccount(testconstants.Address1, nil)},
	}

	for _, tc := range cases {
		err := tc.msg.ValidateBasic()

		if tc.valid {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
		}
	}
}

func TestMsgProposeRevokeAccountGetSignBytes(t *testing.T) {
	msg := NewMsgProposeRevokeAccount(testconstants.Address1, testconstants.Signer)

	expected := `{"type":"auth/ProposeRevokeAccount","value":{"address":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz",` +
		`"signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz"}}`
	require.Equal(t, expected, string(msg.GetSignBytes()))
}

/*
	MsgApproveRevokeAccount
*/

func TestNewMsgApproveRevokeAccount(t *testing.T) {
	msg := NewMsgApproveRevokeAccount(testconstants.Address1, testconstants.Signer)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "approve_revoke_account")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}

func TestValidateMsgApproveRevokeAccount(t *testing.T) {
	cases := []struct {
		valid bool
		msg   MsgApproveRevokeAccount
	}{
		{true, NewMsgApproveRevokeAccount(testconstants.Address1, testconstants.Signer)},
		{false, NewMsgApproveRevokeAccount(nil, testconstants.Signer)},
		{false, NewMsgApproveRevokeAccount(testconstants.Address1, nil)},
	}

	for _, tc := range cases {
		err := tc.msg.ValidateBasic()

		if tc.valid {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
		}
	}
}

func TestMsgApproveRevokeAccountGetSignBytes(t *testing.T) {
	msg := NewMsgApproveRevokeAccount(testconstants.Address1, testconstants.Signer)

	expected := `{"type":"auth/ApproveRevokeAccount","value":{"address":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz",` +
		`"signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz"}}`
	require.Equal(t, expected, string(msg.GetSignBytes()))
}
