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
package types_test

import (
	"testing"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

/*
	dclauthtypes.MsgProposeAddAccount
*/

func NewMsgProposeAddAccountWrapper(
	t *testing.T,
	signer sdk.AccAddress,
	address sdk.AccAddress,
	pubKey cryptotypes.PubKey,
	roles dclauthtypes.AccountRoles,
	vendorID uint64,
) *dclauthtypes.MsgProposeAddAccount {
	msg, err := dclauthtypes.NewMsgProposeAddAccount(signer, address, pubKey, roles, vendorID)
	require.NoError(t, err)
	return msg
}

func TestNewMsgProposeAddAccount(t *testing.T) {
	msg := NewMsgProposeAddAccountWrapper(
		t,
		testconstants.Signer,
		testconstants.Address1, testconstants.PubKey1,
		dclauthtypes.AccountRoles{}, testconstants.VendorID1,
	)

	require.Equal(t, msg.Route(), dclauthtypes.RouterKey)
	require.Equal(t, msg.Type(), "propose_add_account")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}

func TestValidateMsgProposeAddAccount(t *testing.T) {
	cases := []struct {
		valid bool
		msg   *dclauthtypes.MsgProposeAddAccount
	}{
		{false, NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, testconstants.PubKey1,
			dclauthtypes.AccountRoles{}, 0)}, // no roles provided
		{true, NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, testconstants.PubKey1,
			dclauthtypes.AccountRoles{dclauthtypes.NodeAdmin}, 0)},
		{true, NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, testconstants.PubKey1,
			dclauthtypes.AccountRoles{dclauthtypes.Vendor, dclauthtypes.NodeAdmin}, testconstants.VendorID1)},
		{true, NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, testconstants.PubKey1,
			dclauthtypes.AccountRoles{dclauthtypes.Vendor, dclauthtypes.NodeAdmin}, testconstants.VendorID1)},
		{false, NewMsgProposeAddAccountWrapper(t, testconstants.Signer, nil, testconstants.PubKey1,
			dclauthtypes.AccountRoles{dclauthtypes.NodeAdmin}, 0)},
		//{false, NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, "",
		//	dclauthtypes.AccountRoles{}, 0)},
		{false, NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, testconstants.PubKey1,
			dclauthtypes.AccountRoles{"Wrong Role"}, 0)},
		{false, NewMsgProposeAddAccountWrapper(t, nil, testconstants.Address1, testconstants.PubKey1,
			dclauthtypes.AccountRoles{dclauthtypes.NodeAdmin}, 0)},
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
	msg := NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address2, testconstants.PubKey2,
		dclauthtypes.AccountRoles{}, testconstants.VendorID1)

	expected := `{"address":"cosmos1nl4uaesk9gtu7su3n89lne6xpa6lq8gljn79rq","pubKey":{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A2wJ7uOEE5Zm04K52czFTXfDj1qF2mholzi1zOJVlKlr"}` +
		`,"roles":[],"signer":"cosmos1s5xf3aanx7w84hgplk9z3l90qfpantg6nsmhpf","vendorID":"1000"}`

	require.Equal(t, expected, string(msg.GetSignBytes()))
}

/*
	dclauthtypes.MsgApproveAddAccount
*/

func TestNewMsgApproveAddAccount(t *testing.T) {
	msg := dclauthtypes.NewMsgApproveAddAccount(testconstants.Signer, testconstants.Address1)

	require.Equal(t, msg.Route(), dclauthtypes.RouterKey)
	require.Equal(t, msg.Type(), "approve_add_account")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}

func TestValidateMsgApproveAddAccount(t *testing.T) {
	cases := []struct {
		valid bool
		msg   *dclauthtypes.MsgApproveAddAccount
	}{
		{true, dclauthtypes.NewMsgApproveAddAccount(testconstants.Signer, testconstants.Address1)},
		{false, dclauthtypes.NewMsgApproveAddAccount(testconstants.Signer, nil)},
		{false, dclauthtypes.NewMsgApproveAddAccount(nil, testconstants.Address1)},
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
	msg := dclauthtypes.NewMsgApproveAddAccount(testconstants.Signer, testconstants.Address2)

	expected := `{"address":"cosmos1nl4uaesk9gtu7su3n89lne6xpa6lq8gljn79rq","signer":"cosmos1s5xf3aanx7w84hgplk9z3l90qfpantg6nsmhpf"}`
	require.Equal(t, expected, string(msg.GetSignBytes()))
}

/*
	MsgProposeRevokeAccount
*/

func TestNewMsgProposeRevokeAccount(t *testing.T) {
	msg := dclauthtypes.NewMsgProposeRevokeAccount(testconstants.Signer, testconstants.Address1)

	require.Equal(t, msg.Route(), dclauthtypes.RouterKey)
	require.Equal(t, msg.Type(), "propose_revoke_account")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}

func TestValidateMsgProposeRevokeAccount(t *testing.T) {
	cases := []struct {
		valid bool
		msg   *dclauthtypes.MsgProposeRevokeAccount
	}{
		{true, dclauthtypes.NewMsgProposeRevokeAccount(testconstants.Signer, testconstants.Address1)},
		{false, dclauthtypes.NewMsgProposeRevokeAccount(testconstants.Signer, nil)},
		{false, dclauthtypes.NewMsgProposeRevokeAccount(nil, testconstants.Address1)},
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
	msg := dclauthtypes.NewMsgProposeRevokeAccount(testconstants.Signer, testconstants.Address2)

	expected := `{"address":"cosmos1nl4uaesk9gtu7su3n89lne6xpa6lq8gljn79rq","signer":"cosmos1s5xf3aanx7w84hgplk9z3l90qfpantg6nsmhpf"}`
	require.Equal(t, expected, string(msg.GetSignBytes()))
}

/*
	dclauthtypes.MsgApproveRevokeAccount
*/

func TestNewMsgApproveRevokeAccount(t *testing.T) {
	msg := dclauthtypes.NewMsgApproveRevokeAccount(testconstants.Signer, testconstants.Address1)

	require.Equal(t, msg.Route(), dclauthtypes.RouterKey)
	require.Equal(t, msg.Type(), "approve_revoke_account")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}

func TestValidateMsgApproveRevokeAccount(t *testing.T) {
	cases := []struct {
		valid bool
		msg   *dclauthtypes.MsgApproveRevokeAccount
	}{
		{true, dclauthtypes.NewMsgApproveRevokeAccount(testconstants.Signer, testconstants.Address1)},
		{false, dclauthtypes.NewMsgApproveRevokeAccount(testconstants.Signer, nil)},
		{false, dclauthtypes.NewMsgApproveRevokeAccount(nil, testconstants.Address1)},
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
	msg := dclauthtypes.NewMsgApproveRevokeAccount(testconstants.Signer, testconstants.Address2)
	expected := `{"address":"cosmos1nl4uaesk9gtu7su3n89lne6xpa6lq8gljn79rq","signer":"cosmos1s5xf3aanx7w84hgplk9z3l90qfpantg6nsmhpf"}`
	require.Equal(t, expected, string(msg.GetSignBytes()))
}
