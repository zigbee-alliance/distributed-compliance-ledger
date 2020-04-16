package types

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewMsgAddTestingResult(t *testing.T) {
	var msg = NewMsgAddTestingResult(test_constants.VID, test_constants.PID, test_constants.TestResult,
		test_constants.TestDate, test_constants.Signer)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "add_testing_result")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{test_constants.Signer})
}

func TestMsgAddTestingResultValidation(t *testing.T) {
	cases := []struct {
		valid bool
		msg   MsgAddTestingResult
	}{
		{true, NewMsgAddTestingResult(
			test_constants.VID, test_constants.PID, test_constants.TestResult, test_constants.TestDate, test_constants.Signer)},
		{false, NewMsgAddTestingResult(
			0, test_constants.PID, test_constants.TestResult, test_constants.TestDate, test_constants.Signer)},
		{false, NewMsgAddTestingResult(
			test_constants.VID, 0, test_constants.TestResult, test_constants.TestDate, test_constants.Signer)},
		{false, NewMsgAddTestingResult(
			test_constants.VID, test_constants.PID, "", test_constants.TestDate, test_constants.Signer)},
		{false, NewMsgAddTestingResult(
			test_constants.VID, test_constants.PID, test_constants.TestResult, time.Time{}, test_constants.Signer)},
		{false, NewMsgAddTestingResult(
			test_constants.VID, test_constants.PID, test_constants.TestResult, test_constants.TestDate, nil)},
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

func TestMsgAddTestingResultGetSignBytes(t *testing.T) {
	var msg = NewMsgAddTestingResult(test_constants.VID, test_constants.PID, test_constants.TestResult,
		test_constants.TestDate, test_constants.Signer)
	res := msg.GetSignBytes()

	expected := `{"type":"compliancetest/AddTestingResult","value":{"pid":22,"signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz",` +
		`"test_result":"http://test.result.com","vid":1}}`

	require.Equal(t, expected, string(res))
}
