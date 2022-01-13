// Copyright 2022 DSR Corporation
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

package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
)

func TestValidateMsgAddTestingResul(t *testing.T) {
	cases := []struct {
		valid bool
		msg   *MsgAddTestingResult
	}{

		{true, newMsgAddTestingResult(1, 1, 1, "1", testconstants.Signer)},
		{true, newMsgAddTestingResult(65535, 65535, 1, "1", testconstants.Signer)},

		// zero SV - OK
		{true, newMsgAddTestingResult(1, 1, 0, "1", testconstants.Signer)},

		// zero PID/VID - not OK
		{false, newMsgAddTestingResult(1, 0, 1, "1", testconstants.Signer)},
		{false, newMsgAddTestingResult(0, 1, 1, "1", testconstants.Signer)},

		// negative VID/PID - not OK
		{false, newMsgAddTestingResult(-1, 1, 1, "1", testconstants.Signer)},
		{false, newMsgAddTestingResult(1, -1, 1, "1", testconstants.Signer)},

		// too large VID/PID - not OK
		{false, newMsgAddTestingResult(65535+1, 1, 1, "1", testconstants.Signer)},
		{false, newMsgAddTestingResult(1, 65535+1, 1, "1", testconstants.Signer)},
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

func newMsgAddTestingResult(
	vid int32,
	pid int32,
	softwareVersion uint32,
	softwareVersionString string,
	signer sdk.AccAddress,
) *MsgAddTestingResult {

	return &MsgAddTestingResult{
		Signer:                signer.String(),
		Vid:                   vid,
		Pid:                   pid,
		SoftwareVersion:       softwareVersion,
		SoftwareVersionString: softwareVersionString,
		TestResult:            testconstants.TestResult,
		TestDate:              testconstants.TestDate,
	}
}
