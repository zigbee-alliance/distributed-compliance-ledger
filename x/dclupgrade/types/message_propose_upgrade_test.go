package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
)

func TestMsgProposeUpgrade_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgProposeUpgrade
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgProposeUpgrade{
				Creator: "invalid_address",
				Plan:    testconstants.Plan,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgProposeUpgrade{
				Creator: sample.AccAddress(),
				Plan:    testconstants.Plan,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
