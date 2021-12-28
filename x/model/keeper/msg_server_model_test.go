package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

// Prevent strconv unused error.
var _ = strconv.IntSize

func TestModelMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.ModelKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateModel{
			Creator: creator,
			Vid:     int32(i),
			Pid:     int32(i),
		}
		_, err := srv.CreateModel(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetModel(ctx,
			expected.Vid,
			expected.Pid,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestModelMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateModel
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateModel{
				Creator: creator,
				Vid:     0,
				Pid:     0,
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateModel{
				Creator: "B",
				Vid:     0,
				Pid:     0,
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateModel{
				Creator: creator,
				Vid:     100000,
				Pid:     100000,
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.ModelKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateModel{
				Creator: creator,
				Vid:     0,
				Pid:     0,
			}
			_, err := srv.CreateModel(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateModel(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetModel(ctx,
					expected.Vid,
					expected.Pid,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestModelMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteModel
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteModel{
				Creator: creator,
				Vid:     0,
				Pid:     0,
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteModel{
				Creator: "B",
				Vid:     0,
				Pid:     0,
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteModel{
				Creator: creator,
				Vid:     100000,
				Pid:     100000,
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.ModelKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateModel(wctx, &types.MsgCreateModel{
				Creator: creator,
				Vid:     0,
				Pid:     0,
			})
			require.NoError(t, err)
			_, err = srv.DeleteModel(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetModel(ctx,
					tc.request.Vid,
					tc.request.Pid,
				)
				require.False(t, found)
			}
		})
	}
}
