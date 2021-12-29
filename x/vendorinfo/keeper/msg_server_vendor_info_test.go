package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/types"
)

// Prevent strconv unused error.
var _ = strconv.IntSize

func TestVendorInfoMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.VendorinfoKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateVendorInfo{
			Creator:  creator,
			VendorID: int32(i),
		}
		_, err := srv.CreateVendorInfo(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetVendorInfo(ctx,
			expected.VendorID,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestVendorInfoMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateVendorInfo
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateVendorInfo{
				Creator:  creator,
				VendorID: 0,
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateVendorInfo{
				Creator:  "B",
				VendorID: 0,
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateVendorInfo{
				Creator:  creator,
				VendorID: 100000,
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.VendorinfoKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateVendorInfo{
				Creator:  creator,
				VendorID: 0,
			}
			_, err := srv.CreateVendorInfo(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateVendorInfo(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetVendorInfo(ctx,
					expected.VendorID,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestVendorInfoMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteVendorInfo
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteVendorInfo{
				Creator:  creator,
				VendorID: 0,
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteVendorInfo{
				Creator:  "B",
				VendorID: 0,
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteVendorInfo{
				Creator:  creator,
				VendorID: 100000,
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.VendorinfoKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateVendorInfo(wctx, &types.MsgCreateVendorInfo{
				Creator:  creator,
				VendorID: 0,
			})
			require.NoError(t, err)
			_, err = srv.DeleteVendorInfo(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetVendorInfo(ctx,
					tc.request.VendorID,
				)
				require.False(t, found)
			}
		})
	}
}
