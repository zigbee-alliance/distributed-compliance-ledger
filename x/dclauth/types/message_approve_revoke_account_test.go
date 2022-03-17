package types

import (
	fmt "fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
)

func TestNewMsgApproveRevokeAccount(t *testing.T) {
	msg := NewMsgApproveRevokeAccount(testconstants.Signer, testconstants.Address1, testconstants.Info)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "approve_revoke_account")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}

func TestValidateMsgApproveRevokeAccount(t *testing.T) {
	tests := []struct {
		valid bool
		msg   *MsgApproveRevokeAccount
	}{
		{
			valid: true,
			msg:   NewMsgApproveRevokeAccount(testconstants.Signer, testconstants.Address1, testconstants.Info),
		},
		{
			valid: true,
			msg:   NewMsgApproveRevokeAccount(testconstants.Signer, testconstants.Address1, ""),
		},
		{
			valid: false,
			msg:   NewMsgApproveRevokeAccount(testconstants.Signer, nil, testconstants.Info),
		},
		{
			valid: false,
			msg:   NewMsgApproveRevokeAccount(nil, testconstants.Address1, testconstants.Info),
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

func TestMsgApproveRevokeAccountGetSignBytes(t *testing.T) {
	msg := NewMsgApproveRevokeAccount(testconstants.Signer, testconstants.Address2, "Sample Information")
	transcationTime := msg.Time
	expected := fmt.Sprintf(`{"address":"cosmos1nl4uaesk9gtu7su3n89lne6xpa6lq8gljn79rq","info":"Sample Information","signer":"cosmos1s5xf3aanx7w84hgplk9z3l90qfpantg6nsmhpf","time":"%v"}`,
		transcationTime)
	require.Equal(t, expected, string(msg.GetSignBytes()))
}
