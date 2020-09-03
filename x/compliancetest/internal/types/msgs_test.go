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

//nolint:testpackage
package types

//nolint:goimports
import (
	"testing"
	"time"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestNewMsgAddTestingResult(t *testing.T) {
	msg := NewMsgAddTestingResult(testconstants.VID, testconstants.PID, testconstants.TestResult,
		testconstants.TestDate, testconstants.Signer)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "add_testing_result")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}

func TestMsgAddTestingResultValidation(t *testing.T) {
	cases := []struct {
		valid bool
		msg   MsgAddTestingResult
	}{
		{true, NewMsgAddTestingResult(
			testconstants.VID, testconstants.PID, testconstants.TestResult, testconstants.TestDate, testconstants.Signer)},
		{false, NewMsgAddTestingResult(
			0, testconstants.PID, testconstants.TestResult, testconstants.TestDate, testconstants.Signer)},
		{false, NewMsgAddTestingResult(
			testconstants.VID, 0, testconstants.TestResult, testconstants.TestDate, testconstants.Signer)},
		{false, NewMsgAddTestingResult(
			testconstants.VID, testconstants.PID, "", testconstants.TestDate, testconstants.Signer)},
		{false, NewMsgAddTestingResult(
			testconstants.VID, testconstants.PID, testconstants.TestResult, time.Time{}, testconstants.Signer)},
		{false, NewMsgAddTestingResult(
			testconstants.VID, testconstants.PID, testconstants.TestResult, testconstants.TestDate, nil)},
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
	msg := NewMsgAddTestingResult(testconstants.VID, testconstants.PID, testconstants.TestResult,
		testconstants.TestDate, testconstants.Signer)

	expected := `{"type":"compliancetest/AddTestingResult","value":{"pid":22,"signer":` +
		`"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz",` +
		`"test_date":"2020-02-02T02:00:00Z","test_result":"http://test.result.com","vid":1}}`

	require.Equal(t, expected, string(msg.GetSignBytes()))
}
