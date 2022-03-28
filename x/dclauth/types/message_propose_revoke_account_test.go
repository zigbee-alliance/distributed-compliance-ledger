package types

import (
	fmt "fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
)

func TestNewMsgProposeRevokeAccount(t *testing.T) {
	msg := NewMsgProposeRevokeAccount(testconstants.Signer, testconstants.Address1, testconstants.Info)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "propose_revoke_account")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}

func TestValidateMsgProposeRevokedAccount(t *testing.T) {
	positiveTests := []struct {
		valid bool
		msg   *MsgProposeRevokeAccount
	}{
		{
			valid: true,
			msg:   NewMsgProposeRevokeAccount(testconstants.Signer, testconstants.Address1, testconstants.Info),
		},
		{
			valid: true,
			msg:   NewMsgProposeRevokeAccount(testconstants.Signer, testconstants.Address1, ""),
		},
	}

	negativeTests := []struct {
		valid bool
		msg   *MsgProposeRevokeAccount
		err   error
	}{
		{
			valid: false,
			msg:   NewMsgProposeRevokeAccount(testconstants.Signer, nil, testconstants.Info),
			err:   sdkerrors.ErrInvalidAddress,
		},
		{
			valid: false,
			msg:   NewMsgProposeRevokeAccount(nil, testconstants.Address1, testconstants.Info),
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

func TestMsgProposeRevokeAccountGetSignBytes(t *testing.T) {
	msg := NewMsgProposeRevokeAccount(testconstants.Signer, testconstants.Address2, testconstants.Info)
	transcationTime := msg.Time
	expected := fmt.Sprintf(`{"address":"cosmos1nl4uaesk9gtu7su3n89lne6xpa6lq8gljn79rq","info":"Information for Proposal/Approval/Revoke","signer":"cosmos1s5xf3aanx7w84hgplk9z3l90qfpantg6nsmhpf","time":"%v"}`,
		transcationTime)
	require.Equal(t, expected, string(msg.GetSignBytes()))
}
