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
	fmt "fmt"
	"testing"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
)

/*
	MsgProposeAddAccount
*/

func NewMsgProposeAddAccountWrapper(
	t *testing.T,
	signer sdk.AccAddress,
	address sdk.AccAddress,
	pubKey cryptotypes.PubKey,
	roles AccountRoles,
	vendorID int32,
) *MsgProposeAddAccount {
	msg, err := NewMsgProposeAddAccount(signer, address, pubKey, roles, vendorID, testconstants.Info)
	require.NoError(t, err)
	return msg
}

func TestNewMsgProposeAddAccount(t *testing.T) {
	msg := NewMsgProposeAddAccountWrapper(
		t,
		testconstants.Signer,
		testconstants.Address1, testconstants.PubKey1,
		AccountRoles{}, testconstants.VendorID1,
	)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "propose_add_account")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}

func TestValidateMsgProposeAddAccount(t *testing.T) {
	cases := []struct {
		valid bool
		msg   *MsgProposeAddAccount
	}{
		{false, NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, testconstants.PubKey1,
			AccountRoles{}, 1)}, // no roles provided
		{true, NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, testconstants.PubKey1,
			AccountRoles{NodeAdmin}, 1)},
		{true, NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, testconstants.PubKey1,
			AccountRoles{Vendor, NodeAdmin}, testconstants.VendorID1)},

		// zero VID with Vendor role - error - can not create Vendor with vid=0 (reserved)
		{false, NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, testconstants.PubKey1,
			AccountRoles{Vendor, NodeAdmin}, 0)},

		// zero VID without Vendor role - no error
		{true, NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, testconstants.PubKey1,
			AccountRoles{NodeAdmin}, 0)},

		// negative VID - error
		{false, NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, testconstants.PubKey1,
			AccountRoles{Vendor, NodeAdmin}, -1)},

		// too large VID - error
		{false, NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, testconstants.PubKey1,
			AccountRoles{Vendor, NodeAdmin}, 65535+1)},

		{true, NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, testconstants.PubKey1,
			AccountRoles{Vendor, NodeAdmin}, testconstants.VendorID1)},
		{false, NewMsgProposeAddAccountWrapper(t, testconstants.Signer, nil, testconstants.PubKey1,
			AccountRoles{NodeAdmin}, 1)},
		// {false, NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, "",
		//	AccountRoles{}, 1)},
		{false, NewMsgProposeAddAccountWrapper(t, testconstants.Signer, testconstants.Address1, testconstants.PubKey1,
			AccountRoles{"Wrong Role"}, 1)},
		{false, NewMsgProposeAddAccountWrapper(t, nil, testconstants.Address1, testconstants.PubKey1,
			AccountRoles{NodeAdmin}, 1)},
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
		AccountRoles{}, testconstants.VendorID1)
	transcationTime := msg.Time
	expected := fmt.Sprintf(`{"address":"cosmos1nl4uaesk9gtu7su3n89lne6xpa6lq8gljn79rq","info":"Information for Proposal/Approval/Revoke","pubKey":{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A2wJ7uOEE5Zm04K52czFTXfDj1qF2mholzi1zOJVlKlr"},"roles":[],"signer":"cosmos1s5xf3aanx7w84hgplk9z3l90qfpantg6nsmhpf","time":"%v","vendorID":1000}`,
		transcationTime)

	require.Equal(t, expected, string(msg.GetSignBytes()))
}

/*
	MsgApproveAddAccount
*/

func TestNewMsgApproveAddAccount(t *testing.T) {
	msg := NewMsgApproveAddAccount(testconstants.Signer, testconstants.Address1, testconstants.Info)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "approve_add_account")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}

func TestValidateMsgApproveAddAccount(t *testing.T) {
	cases := []struct {
		valid bool
		msg   *MsgApproveAddAccount
	}{
		{true, NewMsgApproveAddAccount(testconstants.Signer, testconstants.Address1, testconstants.Info)},
		{true, NewMsgApproveAddAccount(testconstants.Signer, testconstants.Address1, "")},
		{false, NewMsgApproveAddAccount(testconstants.Signer, nil, testconstants.Info)},
		{false, NewMsgApproveAddAccount(nil, testconstants.Address1, testconstants.Info)},
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
	msg := NewMsgApproveAddAccount(testconstants.Signer, testconstants.Address2, testconstants.Info)
	//nolint:goconst
	transcationTime := msg.Time
	expected := fmt.Sprintf(`{"address":"cosmos1nl4uaesk9gtu7su3n89lne6xpa6lq8gljn79rq","info":"Information for Proposal/Approval/Revoke","signer":"cosmos1s5xf3aanx7w84hgplk9z3l90qfpantg6nsmhpf","time":"%v"}`,
		transcationTime)
	require.Equal(t, expected, string(msg.GetSignBytes()))
}

/*
	MsgProposeRevokeAccount
*/

func TestNewMsgProposeRevokeAccount(t *testing.T) {
	msg := NewMsgProposeRevokeAccount(testconstants.Signer, testconstants.Address1, testconstants.Info)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "propose_revoke_account")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}

func TestValidateMsgProposeRevokeAccount(t *testing.T) {
	cases := []struct {
		valid bool
		msg   *MsgProposeRevokeAccount
	}{
		{true, NewMsgProposeRevokeAccount(testconstants.Signer, testconstants.Address1, testconstants.Info)},
		{true, NewMsgProposeRevokeAccount(testconstants.Signer, testconstants.Address1, "")},
		{false, NewMsgProposeRevokeAccount(testconstants.Signer, nil, testconstants.Info)},
		{false, NewMsgProposeRevokeAccount(nil, testconstants.Address1, testconstants.Info)},
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
	msg := NewMsgProposeRevokeAccount(testconstants.Signer, testconstants.Address2, testconstants.Info)
	transcationTime := msg.Time
	expected := fmt.Sprintf(`{"address":"cosmos1nl4uaesk9gtu7su3n89lne6xpa6lq8gljn79rq","info":"Information for Proposal/Approval/Revoke","signer":"cosmos1s5xf3aanx7w84hgplk9z3l90qfpantg6nsmhpf","time":"%v"}`,
		transcationTime)
	require.Equal(t, expected, string(msg.GetSignBytes()))
}

/*
	MsgApproveRevokeAccount
*/

func TestNewMsgApproveRevokeAccount(t *testing.T) {
	msg := NewMsgApproveRevokeAccount(testconstants.Signer, testconstants.Address1, testconstants.Info)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "approve_revoke_account")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}

func TestValidateMsgApproveRevokeAccount(t *testing.T) {
	cases := []struct {
		valid bool
		msg   *MsgApproveRevokeAccount
	}{
		{true, NewMsgApproveRevokeAccount(testconstants.Signer, testconstants.Address1, testconstants.Info)},
		{true, NewMsgApproveRevokeAccount(testconstants.Signer, testconstants.Address1, "")},
		{false, NewMsgApproveRevokeAccount(testconstants.Signer, nil, testconstants.Info)},
		{false, NewMsgApproveRevokeAccount(nil, testconstants.Address1, testconstants.Info)},
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
	msg := NewMsgApproveRevokeAccount(testconstants.Signer, testconstants.Address2, "Sample Information")
	transcationTime := msg.Time
	expected := fmt.Sprintf(`{"address":"cosmos1nl4uaesk9gtu7su3n89lne6xpa6lq8gljn79rq","info":"Sample Information","signer":"cosmos1s5xf3aanx7w84hgplk9z3l90qfpantg6nsmhpf","time":"%v"}`,
		transcationTime)
	require.Equal(t, expected, string(msg.GetSignBytes()))
}
