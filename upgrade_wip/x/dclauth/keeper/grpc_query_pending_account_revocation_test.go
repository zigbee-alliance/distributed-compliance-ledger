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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestPendingAccountRevocationQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.DclauthKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNPendingAccountRevocation(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetPendingAccountRevocationRequest
		response *types.QueryGetPendingAccountRevocationResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetPendingAccountRevocationRequest{
				Address: msgs[0].Address,
			},
			response: &types.QueryGetPendingAccountRevocationResponse{PendingAccountRevocation: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetPendingAccountRevocationRequest{
				Address: msgs[1].Address,
			},
			response: &types.QueryGetPendingAccountRevocationResponse{PendingAccountRevocation: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetPendingAccountRevocationRequest{
				Address: strconv.Itoa(100000),
			},
			err: status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.PendingAccountRevocation(wctx, tc.request)
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

func TestPendingAccountRevocationQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.DclauthKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNPendingAccountRevocation(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllPendingAccountRevocationRequest {
		return &types.QueryAllPendingAccountRevocationRequest{
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
			resp, err := keeper.PendingAccountRevocationAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.PendingAccountRevocation), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.PendingAccountRevocation),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.PendingAccountRevocationAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.PendingAccountRevocation), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.PendingAccountRevocation),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.PendingAccountRevocationAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.PendingAccountRevocation),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.PendingAccountRevocationAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
