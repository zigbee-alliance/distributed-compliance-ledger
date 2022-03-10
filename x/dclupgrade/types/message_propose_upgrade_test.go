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
	negative_tests := []struct {
		name string
		msg  MsgProposeUpgrade
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgProposeUpgrade{
				Creator: "invalid_address",
				Plan: Plan{
					Name:   testconstants.UpgradePlanName,
					Height: testconstants.UpgradePlanHeight,
					Info:   testconstants.UpgradePlanInfo,
				},
				Info: testconstants.Info,
				Time: testconstants.Time,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "omitted address",
			msg: MsgProposeUpgrade{
				Creator: "",
				Plan: Plan{
					Name:   testconstants.UpgradePlanName,
					Height: testconstants.UpgradePlanHeight,
					Info:   testconstants.UpgradePlanInfo,
				},
				Info: testconstants.Info,
				Time: testconstants.Time,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "plan name is not set",
			msg: MsgProposeUpgrade{
				Creator: sample.AccAddress(),
				Plan: Plan{
					Name:   "",
					Height: testconstants.UpgradePlanHeight,
					Info:   testconstants.UpgradePlanInfo,
				},
				Info: testconstants.Info,
				Time: testconstants.Time,
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "plan height is 0",
			msg: MsgProposeUpgrade{
				Creator: sample.AccAddress(),
				Plan: Plan{
					Name:   testconstants.UpgradePlanName,
					Height: 0,
					Info:   testconstants.UpgradePlanInfo,
				},
				Info: testconstants.Info,
				Time: testconstants.Time,
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "plan height is less than 0",
			msg: MsgProposeUpgrade{
				Creator: sample.AccAddress(),
				Plan: Plan{
					Name:   testconstants.UpgradePlanName,
					Height: -1,
					Info:   testconstants.UpgradePlanInfo,
				},
				Info: testconstants.Info,
				Time: testconstants.Time,
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "plan time is not zero",
			msg: MsgProposeUpgrade{
				Creator: sample.AccAddress(),
				Plan: Plan{
					Name:   testconstants.UpgradePlanName,
					Height: testconstants.UpgradePlanHeight,
					Info:   testconstants.UpgradePlanInfo,
					Time:   time.Now(),
				},
				Info: testconstants.Info,
				Time: testconstants.Time,
			},
			err: *sdkerrors.ErrInvalidRequest,
		},
		{
			name: "Plan upgradedClientState is not nil",
			msg: MsgProposeUpgrade{
				Creator: sample.AccAddress(),
				Plan: Plan{
					Name:                testconstants.UpgradePlanName,
					Height:              testconstants.UpgradePlanHeight,
					Info:                testconstants.UpgradePlanInfo,
					UpgradedClientState: &codectypes.Any{TypeUrl: "333"},
				},
				Info: testconstants.Info,
				Time: testconstants.Time,
			},
			err: *sdkerrors.ErrInvalidRequest,
		},
	}

	positive_tests := []struct {
		name string
		msg  MsgProposeUpgrade
	}{
		{
			name: "valid MsgProposeUpgrade message",
			msg: MsgProposeUpgrade{
				Creator: sample.AccAddress(),
				Plan: Plan{
					Name:   testconstants.UpgradePlanName,
					Height: testconstants.UpgradePlanHeight,
					Info:   testconstants.UpgradePlanInfo,
				},
				Info: testconstants.Info,
				Time: testconstants.Time,
			},
		},
		{
			name: "info is not set",
			msg: MsgProposeUpgrade{
				Creator: sample.AccAddress(),
				Plan: Plan{
					Name:   testconstants.UpgradePlanName,
					Height: testconstants.UpgradePlanHeight,
					Info:   testconstants.UpgradePlanInfo,
				},
				Info: "",
				Time: testconstants.Time,
			},
		},
		{
			name: "plan height is greater than 0",
			msg: MsgProposeUpgrade{
				Creator: sample.AccAddress(),
				Plan: Plan{
					Name:   testconstants.UpgradePlanName,
					Height: 1,
					Info:   testconstants.UpgradePlanInfo,
				},
				Info: testconstants.Info,
				Time: testconstants.Time,
			},
		},
		{
			name: "plan info is not set",
			msg: MsgProposeUpgrade{
				Creator: sample.AccAddress(),
				Plan: Plan{
					Name:   testconstants.UpgradePlanName,
					Height: testconstants.UpgradePlanHeight,
					Info:   "",
				},
				Info: testconstants.Info,
				Time: testconstants.Time,
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
