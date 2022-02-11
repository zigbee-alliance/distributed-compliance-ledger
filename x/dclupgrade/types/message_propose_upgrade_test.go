package types

import (
	"testing"
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
)

func TestMsgProposeUpgrade_ValidateBasic(t *testing.T) {
	// negative test constants

	planNameLen0 := Plan{
		Name:   "",
		Height: 1,
		Info:   "Some info",
	}

	planHeight0 := Plan{
		Name:   "Some plan name",
		Height: 0,
		Info:   "Some info",
	}

	planHeightLess0 := Plan{
		Name:   "Some plan name",
		Height: -1,
		Info:   "Some info",
	}

	planTimeIsNot0 := Plan{
		Name:   "Some plan name",
		Height: 1,
		Info:   "Some info",
		Time:   time.Now(),
	}
	planUpgradedClientStateIsNotNil := Plan{
		Name:                "Some plan name",
		Height:              1,
		Info:                "Some info",
		UpgradedClientState: &codectypes.Any{TypeUrl: "333"},
	}

	// positive test constants

	planNormal := testconstants.Plan

	planInfoLen0 := Plan{
		Name:   "Some plan name",
		Height: 1,
		Info:   "",
	}

	negative_tests := []struct {
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
		},
		{
			name: "omitted address",
			msg: MsgProposeUpgrade{
				Creator: "",
				Plan:    testconstants.Plan,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "Plan name len 0",
			msg: MsgProposeUpgrade{
				Creator: sample.AccAddress(),
				Plan:    planNameLen0,
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "Plan height 0",
			msg: MsgProposeUpgrade{
				Creator: sample.AccAddress(),
				Plan:    planHeight0,
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "Plan height less than 0",
			msg: MsgProposeUpgrade{
				Creator: sample.AccAddress(),
				Plan:    planHeightLess0,
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "Plan time is not zero",
			msg: MsgProposeUpgrade{
				Creator: sample.AccAddress(),
				Plan:    planTimeIsNot0,
			},
			err: *sdkerrors.ErrInvalidRequest,
		},
		{
			name: "Plan upgradedClientState is not nil",
			msg: MsgProposeUpgrade{
				Creator: sample.AccAddress(),
				Plan:    planUpgradedClientStateIsNotNil,
			},
			err: *sdkerrors.ErrInvalidRequest,
		},
	}

	positive_tests := []struct {
		name string
		msg  MsgProposeUpgrade
	}{
		{
			name: "valid msg propose upgrade message",
			msg: MsgProposeUpgrade{
				Creator: sample.AccAddress(),
				Plan:    planNormal,
			},
		},
		{
			name: "info len 0",
			msg: MsgProposeUpgrade{
				Creator: sample.AccAddress(),
				Plan:    planInfoLen0,
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
