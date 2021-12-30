package keeper_test

/*
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

func TestModelVersionMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.ModelKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateModelVersion{
			Creator:         creator,
			Vid:             int32(i),
			Pid:             int32(i),
			SoftwareVersion: uint32(i),
		}
		_, err := srv.CreateModelVersion(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetModelVersion(ctx,
			expected.Vid,
			expected.Pid,
			expected.SoftwareVersion,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestModelVersionMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateModelVersion
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateModelVersion{
				Creator:         creator,
				Vid:             0,
				Pid:             0,
				SoftwareVersion: 0,
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateModelVersion{
				Creator:         "B",
				Vid:             0,
				Pid:             0,
				SoftwareVersion: 0,
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateModelVersion{
				Creator:         creator,
				Vid:             100000,
				Pid:             100000,
				SoftwareVersion: 100000,
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.ModelKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateModelVersion{
				Creator:         creator,
				Vid:             0,
				Pid:             0,
				SoftwareVersion: 0,
			}
			_, err := srv.CreateModelVersion(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateModelVersion(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetModelVersion(ctx,
					expected.Vid,
					expected.Pid,
					expected.SoftwareVersion,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}
*/
