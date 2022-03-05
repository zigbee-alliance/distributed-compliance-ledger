package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
)

func TestMsgEnableValidator_ValidateBasic(t *testing.T) {
	negative_tests := []struct {
		name string
		msg  MsgEnableValidator
		err  error
	}{
		{
			name: "invalid creator address",
			msg: MsgEnableValidator{
				Creator: "invalid_address",
				Address: testconstants.ValidatorAddress1,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "omitted creator address",
			msg: MsgEnableValidator{
				Creator: "",
				Address: testconstants.ValidatorAddress1,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid validator address",
			msg: MsgEnableValidator{
				Creator: testconstants.Address1.String(),
				Address: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "omitted validator address",
			msg: MsgEnableValidator{
				Creator: testconstants.Address1.String(),
				Address: "",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
	}

	positive_tests := []struct {
		name string
		msg  MsgEnableValidator
	}{
		{
			name: "valid EnableValidator message",
			msg: MsgEnableValidator{
				Creator: sample.AccAddress(),
				Address: testconstants.ValidatorAddress1,
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
