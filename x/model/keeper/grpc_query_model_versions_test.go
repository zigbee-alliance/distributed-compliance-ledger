package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/nullify"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Prevent strconv unused error.
var _ = strconv.IntSize

func TestModelVersionsQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.ModelKeeper(t, nil)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNModelVersions(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetModelVersionsRequest
		response *types.QueryGetModelVersionsResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetModelVersionsRequest{
				Vid: msgs[0].Vid,
				Pid: msgs[0].Pid,
			},
			response: &types.QueryGetModelVersionsResponse{ModelVersions: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetModelVersionsRequest{
				Vid: msgs[1].Vid,
				Pid: msgs[1].Pid,
			},
			response: &types.QueryGetModelVersionsResponse{ModelVersions: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetModelVersionsRequest{
				Vid: 100000,
				Pid: 100000,
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ModelVersions(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
