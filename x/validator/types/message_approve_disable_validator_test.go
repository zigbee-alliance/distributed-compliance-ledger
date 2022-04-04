package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
)

func TestMsgApproveDisableValidator_ValidateBasic(t *testing.T) {
	negative_tests := []struct {
		name string
		msg  MsgApproveDisableValidator
		err  error
	}{
		{
			name: "invalid creator address",
			msg: MsgApproveDisableValidator{
				Creator: "invalid_address",
				Address: testconstants.ValidatorAddress1,
				Info:    testconstants.Info,
				Time:    testconstants.Time,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "omitted creator address",
			msg: MsgApproveDisableValidator{
				Creator: "",
				Address: testconstants.ValidatorAddress1,
				Info:    testconstants.Info,
				Time:    testconstants.Time,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid validator address",
			msg: MsgApproveDisableValidator{
				Creator: testconstants.Address1.String(),
				Address: "invalid_address",
				Info:    testconstants.Info,
				Time:    testconstants.Time,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "omitted validator address",
			msg: MsgApproveDisableValidator{
				Creator: testconstants.Address1.String(),
				Address: "",
				Info:    testconstants.Info,
				Time:    testconstants.Time,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
	}

	positive_tests := []struct {
		name string
		msg  MsgApproveDisableValidator
	}{
		{
			name: "valid ApproveDisableValidator message",
			msg: MsgApproveDisableValidator{
				Creator: sample.AccAddress(),
				Address: testconstants.ValidatorAddress1,
				Info:    testconstants.Info,
				Time:    testconstants.Time,
			},
		},
	}
	for _, tt := range positive_tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			require.NoError(t, err)
		})
	}

	for _, tt := range negative_tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			require.Error(t, err)
			require.ErrorIs(t, err, tt.err)
		})
	}
}
