package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/nullify"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestRejectedNodeQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.ValidatorKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRejectedNode(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetRejectedNodeRequest
		response *types.QueryGetRejectedNodeResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetRejectedNodeRequest{
				Owner: msgs[0].Owner,
			},
			response: &types.QueryGetRejectedNodeResponse{RejectedNode: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetRejectedNodeRequest{
				Owner: msgs[1].Owner,
			},
			response: &types.QueryGetRejectedNodeResponse{RejectedNode: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetRejectedNodeRequest{
				Owner: strconv.Itoa(100000),
			},
			err: status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.RejectedNode(wctx, tc.request)
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

func TestRejectedNodeQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.ValidatorKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRejectedNode(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllRejectedNodeRequest {
		return &types.QueryAllRejectedNodeRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.RejectedNodeAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.RejectedNode), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.RejectedNode),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.RejectedNodeAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.RejectedNode), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.RejectedNode),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.RejectedNodeAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.RejectedNode),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.RejectedNodeAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
