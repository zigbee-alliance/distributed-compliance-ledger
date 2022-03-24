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
