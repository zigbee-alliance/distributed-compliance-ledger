package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
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
				Info:    testconstants.Info,
				Time:    testconstants.Time,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "omitted creator address",
			msg: MsgEnableValidator{
				Creator: "",
				Address: testconstants.ValidatorAddress1,
				Info:    testconstants.Info,
				Time:    testconstants.Time,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid validator address",
			msg: MsgEnableValidator{
				Creator: testconstants.Address1.String(),
				Address: "invalid_address",
				Info:    testconstants.Info,
				Time:    testconstants.Time,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "omitted validator address",
			msg: MsgEnableValidator{
				Creator: testconstants.Address1.String(),
				Address: "",
				Info:    testconstants.Info,
				Time:    testconstants.Time,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "info len > 4096",
			msg: MsgEnableValidator{
				Creator: sample.AccAddress(),
				Address: testconstants.ValidatorAddress1,
				Info:    tmrand.Str(4097),
				Time:    testconstants.Time,
			},
			err: validator.ErrFieldMaxLengthExceeded,
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
