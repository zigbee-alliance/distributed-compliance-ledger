package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

func TestMsgApproveUpgrade_ValidateBasic(t *testing.T) {
	negative_tests := []struct {
		name string
		msg  MsgApproveUpgrade
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgApproveUpgrade{
				Creator: "invalid_address",
				Name:    testconstants.Plan.Name,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "Name is not set",
			msg: MsgApproveUpgrade{
				Creator: sample.AccAddress(),
				Name:    "",
			},
			err: validator.ErrRequiredFieldMissing,
		},
	}

	positive_tests := []struct {
		name string
		msg  MsgApproveUpgrade
	}{
		{
			name: "valid msg approve upgrade message",
			msg: MsgApproveUpgrade{
				Creator: sample.AccAddress(),
				Name:    "Test plan example",
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
