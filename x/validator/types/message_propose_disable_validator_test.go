package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
)

func TestMsgProposeDisableValidator_ValidateBasic(t *testing.T) {
	negativeTests := []struct {
		name string
		msg  MsgProposeDisableValidator
		err  error
	}{
		{
			name: "invalid creator address",
			msg: MsgProposeDisableValidator{
				Creator: "invalid_address",
				Address: testconstants.ValidatorAddress1,
				Info:    testconstants.Info,
				Time:    testconstants.Time,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "omitted creator address",
			msg: MsgProposeDisableValidator{
				Creator: "",
				Address: testconstants.ValidatorAddress1,
				Info:    testconstants.Info,
				Time:    testconstants.Time,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid validator address",
			msg: MsgProposeDisableValidator{
				Creator: testconstants.Address1.String(),
				Address: "invalid_address",
				Info:    testconstants.Info,
				Time:    testconstants.Time,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "omitted validator address",
			msg: MsgProposeDisableValidator{
				Creator: testconstants.Address1.String(),
				Address: "",
				Info:    testconstants.Info,
				Time:    testconstants.Time,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
	}

	positiveTests := []struct {
		name string
		msg  MsgProposeDisableValidator
	}{
		{
			name: "valid ProposeDisableValidator message",
			msg: MsgProposeDisableValidator{
				Creator: sample.AccAddress(),
				Address: testconstants.ValidatorAddress1,
				Info:    testconstants.Info,
				Time:    testconstants.Time,
			},
		},
	}
	for _, tt := range positiveTests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			require.NoError(t, err)
		})
	}

	for _, tt := range negativeTests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			require.Error(t, err)
			require.ErrorIs(t, err, tt.err)
		})
	}
}
