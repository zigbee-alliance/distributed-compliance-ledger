//nolint:testpackage
package types

//nolint:goimports
import (
	"testing"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

/*
	MsgCreateValidator
*/

func TestNewMsgCreateValidator(t *testing.T) {
	msg := NewMsgCreateValidator(testconstants.ValidatorAddress1, testconstants.ValidatorPubKey1,
		Description{Name: testconstants.Name}, testconstants.Owner)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "create_validator")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{msg.Signer})
}

func TestValidateMsgCreateValidator(t *testing.T) {
	cases := []struct {
		valid bool
		msg   MsgCreateValidator
	}{
		{true, NewMsgCreateValidator(testconstants.ValidatorAddress1, testconstants.ValidatorPubKey1,
			Description{Name: testconstants.Name}, testconstants.Owner)},
		{false, NewMsgCreateValidator(nil, testconstants.ValidatorPubKey1,
			Description{Name: testconstants.Name}, testconstants.Owner)},
		{false, NewMsgCreateValidator(testconstants.ValidatorAddress1, "",
			Description{Name: testconstants.Name}, testconstants.Owner)},
		{false, NewMsgCreateValidator(testconstants.ValidatorAddress1, testconstants.ValidatorPubKey1,
			Description{}, testconstants.Owner)},
		{false, NewMsgCreateValidator(testconstants.ValidatorAddress1, testconstants.ValidatorPubKey1,
			Description{Name: testconstants.Name}, nil)},
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
	msg := NewMsgCreateValidator(testconstants.ValidatorAddress1, testconstants.ValidatorPubKey1,
		Description{Name: "Test"}, testconstants.Owner)

	expected := `{"type":"validator/CreateValidator","value":{` +
		`"description":{"name":"Test"},"signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz",` +
		`"validator_address":"cosmosvalcons158uwzeqeu7zg332ztuzc5xh9k5uy3h5ttegzxd",` +
		`"validator_pubkey":"cosmosvalconspub1zcjduepqdmmjdfyvh2mrwl8p8wkwp23kh8lvjrd9u45snxqz6te6y6lwk6gqts45r3"}}`
	require.Equal(t, expected, string(msg.GetSignBytes()))
}
