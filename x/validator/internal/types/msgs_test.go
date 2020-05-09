package types

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

/*
	MsgCreateValidator
*/

func TestNewMsgCreateValidator(t *testing.T) {
	var msg = NewMsgCreateValidator(test_constants.ValAddress1, test_constants.ConsensusPubKey1, Description{Name: test_constants.Name})

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "create_validator")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{sdk.AccAddress(test_constants.ValAddress1)})
}

func TestValidateMsgCreateValidator(t *testing.T) {
	cases := []struct {
		valid bool
		msg   MsgCreateValidator
	}{
		{true, NewMsgCreateValidator(test_constants.ValAddress1, test_constants.ConsensusPubKey1, Description{Name: test_constants.Name})},
		{false, NewMsgCreateValidator(nil, test_constants.PubKey, Description{Name: test_constants.Name})},
		{false, NewMsgCreateValidator(test_constants.ValAddress1, "", Description{Name: test_constants.Name})},
		{false, NewMsgCreateValidator(test_constants.ValAddress1, test_constants.PubKey, Description{})},
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

func TestMsgCreateValidatorGetSignBytes(t *testing.T) {
	var msg = NewMsgCreateValidator(test_constants.ValAddress1, test_constants.ConsensusPubKey1, Description{Name: "Test"})
	res := msg.GetSignBytes()

	expected := `{"type":"validator/CreateValidator","value":{"description":{"details":"","identity":"",` +
		`"name":"Test","website":""},"pubkey":"cosmosvalconspub1zcjduepqdmmjdfyvh2mrwl8p8wkwp23kh8lvjrd9u45snxqz6te6y6lwk6gqts45r3",` +
		`"validator_address":"cosmosvaloper18gcwk73gtt84aeatqdh7yfesmz9956l0zw8lfw"}}`
	require.Equal(t, expected, string(res))
}
