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

package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
)

func TestNewMsgApproveAddAccount(t *testing.T) {
	msg := NewMsgApproveAddAccount(testconstants.Signer, testconstants.Address1, testconstants.Info)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "approve_add_account")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}

func TestValidateMsgApproveAddAccount(t *testing.T) {
	positiveTests := []struct {
		valid bool
		msg   *MsgApproveAddAccount
	}{
		{
			valid: true,
			msg:   NewMsgApproveAddAccount(testconstants.Signer, testconstants.Address1, testconstants.Info),
		},
		{
			valid: true,
			msg:   NewMsgApproveAddAccount(testconstants.Signer, testconstants.Address1, ""),
		},
	}

	negativeTests := []struct {
		valid bool
		msg   *MsgApproveAddAccount
		err   error
	}{
		{
			valid: false,
			msg:   NewMsgApproveAddAccount(testconstants.Signer, nil, testconstants.Info),
			err:   sdkerrors.ErrInvalidAddress,
		},
		{
			valid: false,
			msg:   NewMsgApproveAddAccount(nil, testconstants.Address1, testconstants.Info),
			err:   sdkerrors.ErrInvalidAddress,
		},
	}

	for _, tt := range positiveTests {
		err := tt.msg.ValidateBasic()

		if tt.valid {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
		}
	}

	for _, tt := range negativeTests {
		err := tt.msg.ValidateBasic()

		if tt.valid {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
			require.ErrorIs(t, err, tt.err)
		}
	}
}

func TestMsgApproveAddAccountGetSignBytes(t *testing.T) {
	msg := NewMsgProposeRevokeAccount(testconstants.Signer, testconstants.Address1, testconstants.Info)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "propose_revoke_account")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}
