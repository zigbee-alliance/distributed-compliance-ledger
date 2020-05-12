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
	var msg = NewMsgCreateValidator(test_constants.ConsensusAddress1, test_constants.ConsensusPubKey1,
		Description{Name: test_constants.Name}, test_constants.Owner)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "create_validator")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{msg.Signer})
}

func TestValidateMsgCreateValidator(t *testing.T) {
	cases := []struct {
		valid bool
		msg   MsgCreateValidator
	}{
		{true, NewMsgCreateValidator(test_constants.ConsensusAddress1, test_constants.ConsensusPubKey1,
			Description{Name: test_constants.Name}, test_constants.Owner)},
		{false, NewMsgCreateValidator(nil, test_constants.PubKey,
			Description{Name: test_constants.Name}, test_constants.Owner)},
		{false, NewMsgCreateValidator(test_constants.ConsensusAddress1, "",
			Description{Name: test_constants.Name}, test_constants.Owner)},
		{false, NewMsgCreateValidator(test_constants.ConsensusAddress1, test_constants.PubKey,
			Description{}, test_constants.Owner)},
		{false, NewMsgCreateValidator(test_constants.ConsensusAddress1, test_constants.PubKey,
			Description{Name: test_constants.Name}, nil)},
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
	var msg = NewMsgCreateValidator(test_constants.ConsensusAddress1, test_constants.ConsensusPubKey1,
		Description{Name: "Test"}, test_constants.Owner)
	res := msg.GetSignBytes()

	expected := `{"type":"validator/CreateValidator","value":{"address":"cosmosvalcons158uwzeqeu7zg332ztuzc5xh9k5uy3h5ttegzxd",` +
		`"description":{"name":"Test"},"pubkey":"cosmosvalconspub1zcjduepqdmmjdfyvh2mrwl8p8wkwp23kh8lvjrd9u45snxqz6te6y6lwk6gqts45r3",` +
		`"signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz"}}`
	require.Equal(t, expected, string(res))
}
