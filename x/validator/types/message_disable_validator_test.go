package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
)

func TestMsgDisableValidator_ValidateBasic(t *testing.T) {
	negative_tests := []struct {
		name string
		msg  MsgDisableValidator
		err  error
	}{
		{
			name: "invalid creator address",
			msg: MsgDisableValidator{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "omitted creator address",
			msg: MsgDisableValidator{
				Creator: "",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
	}

	positive_tests := []struct {
		name string
		msg  MsgDisableValidator
	}{
		{
			name: "valid DisableValidator message",
			msg: MsgDisableValidator{
				Creator: testconstants.ValidatorAddress1,
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
