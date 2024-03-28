package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
)

func TestMsgProposeUpgrade_ValidateBasic(t *testing.T) {
	negativeTests := []struct {
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
	}

	positiveTests := []struct {
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
