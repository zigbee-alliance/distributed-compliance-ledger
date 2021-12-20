package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
)

func TestMsgCreateVendorInfoType_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateVendorInfoType
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateVendorInfoType{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateVendorInfoType{
				Creator: sample.AccAddress(),
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

func TestMsgUpdateVendorInfoType_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateVendorInfoType
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateVendorInfoType{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateVendorInfoType{
				Creator: sample.AccAddress(),
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

func TestMsgDeleteVendorInfoType_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteVendorInfoType
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteVendorInfoType{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteVendorInfoType{
				Creator: sample.AccAddress(),
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
