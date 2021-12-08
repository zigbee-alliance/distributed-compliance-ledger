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

// Prevent strconv unused error
var _ = strconv.IntSize

func TestModelVersionsMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.ModelKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateModelVersions{Creator: creator,
			Vid: int32(i),
			Pid: int32(i),
		}
		_, err := srv.CreateModelVersions(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetModelVersions(ctx,
			expected.Vid,
			expected.Pid,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestModelVersionsMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateModelVersions
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateModelVersions{Creator: creator,
				Vid: 0,
				Pid: 0,
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateModelVersions{Creator: "B",
				Vid: 0,
				Pid: 0,
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateModelVersions{Creator: creator,
				Vid: 100000,
				Pid: 100000,
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.ModelKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateModelVersions{Creator: creator,
				Vid: 0,
				Pid: 0,
			}
			_, err := srv.CreateModelVersions(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateModelVersions(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetModelVersions(ctx,
					expected.Vid,
					expected.Pid,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestModelVersionsMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteModelVersions
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteModelVersions{Creator: creator,
				Vid: 0,
				Pid: 0,
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteModelVersions{Creator: "B",
				Vid: 0,
				Pid: 0,
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteModelVersions{Creator: creator,
				Vid: 100000,
				Pid: 100000,
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.ModelKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateModelVersions(wctx, &types.MsgCreateModelVersions{Creator: creator,
				Vid: 0,
				Pid: 0,
			})
			require.NoError(t, err)
			_, err = srv.DeleteModelVersions(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetModelVersions(ctx,
					tc.request.Vid,
					tc.request.Pid,
				)
				require.False(t, found)
			}
		})
	}
}
