package types

// import (
// 	"testing"

// 	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
// 	"github.com/stretchr/testify/require"
// 	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
// )

// func TestMsgCreateVendorInfo_ValidateBasic(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		msg  MsgCreateVendorInfo
// 		err  error
// 	}{
// 		{
// 			name: "invalid address",
// 			msg: MsgCreateVendorInfo{
// 				Creator: "invalid_address",
// 			},
// 			err: sdkerrors.ErrInvalidAddress,
// 		}, {
// 			name: "valid address",
// 			msg: MsgCreateVendorInfo{
// 				Creator: sample.AccAddress(),
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			err := tt.msg.ValidateBasic()
// 			if tt.err != nil {
// 				require.ErrorIs(t, err, tt.err)
// 				return
// 			}
// 			require.NoError(t, err)
// 		})
// 	}
// }

// func TestMsgUpdateVendorInfo_ValidateBasic(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		msg  MsgUpdateVendorInfo
// 		err  error
// 	}{
// 		{
// 			name: "invalid address",
// 			msg: MsgUpdateVendorInfo{
// 				Creator: "invalid_address",
// 			},
// 			err: sdkerrors.ErrInvalidAddress,
// 		}, {
// 			name: "valid address",
// 			msg: MsgUpdateVendorInfo{
// 				Creator: sample.AccAddress(),
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			err := tt.msg.ValidateBasic()
// 			if tt.err != nil {
// 				require.ErrorIs(t, err, tt.err)
// 				return
// 			}
// 			require.NoError(t, err)
// 		})
// 	}
// }
