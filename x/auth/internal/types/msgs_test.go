//nolint:testpackage
package types

// nolint:goimports
import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

/*
	MsgProposeAddAccount
*/

func TestNewMsgProposeAddAccount(t *testing.T) {
	var msg = NewMsgProposeAddAccount(testconstants.Address1, testconstants.Pubkey1Str,
		AccountRoles{}, testconstants.Signer)

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
			AccountRoles{}, testconstants.Signer)},
		{true, NewMsgProposeAddAccount(testconstants.Address1, testconstants.Pubkey1Str,
			AccountRoles{Vendor, NodeAdmin}, testconstants.Signer)},
		{false, NewMsgProposeAddAccount(nil, testconstants.Pubkey1Str,
			AccountRoles{}, testconstants.Signer)},
		{false, NewMsgProposeAddAccount(testconstants.Address1, "",
			AccountRoles{}, testconstants.Signer)},
		{false, NewMsgProposeAddAccount(testconstants.Address1, testconstants.Pubkey1Str,
			AccountRoles{"Wrong Role"}, testconstants.Signer)},
		{false, NewMsgProposeAddAccount(testconstants.Address1, testconstants.Pubkey1Str,
			AccountRoles{}, nil)},
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
	var msg = NewMsgProposeAddAccount(testconstants.Address1, testconstants.Pubkey1Str,
		AccountRoles{}, testconstants.Signer)

	expected := `{"type":"auth/ProposeAddAccount","value":{"address":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz",` +
		`"pub_key":"cosmospub1addwnpepq28rlfval9n8khmgqz55mlfwn4rlh0jk80k9n7fvtu4g4u37qtvry76ww9h","roles":[],` +
		`"signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz"}}`
	require.Equal(t, expected, string(msg.GetSignBytes()))
}

/*
	MsgApproveAddAccount
*/

func TestNewMsgApproveAddAccount(t *testing.T) {
	var msg = NewMsgApproveAddAccount(testconstants.Address1, testconstants.Signer)

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
	var msg = NewMsgApproveAddAccount(testconstants.Address1, testconstants.Signer)

	expected := `{"type":"auth/ApproveAddAccount","value":{"address":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz",` +
		`"signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz"}}`
	require.Equal(t, expected, string(msg.GetSignBytes()))
}
