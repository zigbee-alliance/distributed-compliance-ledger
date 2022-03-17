package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
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
	tests := []struct {
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
		{
			valid: false,
			msg:   NewMsgApproveAddAccount(testconstants.Signer, nil, testconstants.Info),
		},
		{
			valid: false,
			msg:   NewMsgApproveAddAccount(nil, testconstants.Address1, testconstants.Info),
		},
	}

	for _, tt := range tests {
		err := tt.msg.ValidateBasic()

		if tt.valid {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
		}
	}
}

func TestMsgApproveAddAccountGetSignBytes(t *testing.T) {
	msg := NewMsgProposeRevokeAccount(testconstants.Signer, testconstants.Address1, testconstants.Info)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "propose_revoke_account")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}
