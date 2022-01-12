package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
)

func TestValidateMsgCreateModel(t *testing.T) {
	cases := []struct {
		valid bool
		msg   *MsgAddTestingResult
	}{

		{true, newMsgAddTestingResult(1, 1, 1, "1", testconstants.Signer)},
		{true, newMsgAddTestingResult(65535, 65535, 1, "1", testconstants.Signer)},
		// zero PID/VID/SV - OK
		{true, newMsgAddTestingResult(0, 0, 0, "1", testconstants.Signer)},

		// negative VID - not OK
		{true, newMsgAddTestingResult(-1, 1, 1, "1", testconstants.Signer)},
		// negative PID - not OK
		{true, newMsgAddTestingResult(1, -1, 1, "1", testconstants.Signer)},

		// too large VID - not OK
		{true, newMsgAddTestingResult(65535+1, 1, 1, "1", testconstants.Signer)},
		// too large PID - not OK
		{true, newMsgAddTestingResult(1, 65535+1, 1, "1", testconstants.Signer)},
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
