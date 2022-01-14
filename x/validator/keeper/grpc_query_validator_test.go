package keeper_test

/* FIXME issue 99

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestValidatorQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.ValidatorKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNValidator(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetValidatorRequest
		response *types.QueryGetValidatorResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetValidatorRequest{
				Owner: msgs[0].Owner,
			},
			response: &types.QueryGetValidatorResponse{Validator: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetValidatorRequest{
				Owner: msgs[1].Owner,
			},
			response: &types.QueryGetValidatorResponse{Validator: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetValidatorRequest{
				Owner: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Validator(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}

func TestValidatorQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.ValidatorKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNValidator(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllValidatorRequest {
		return &types.QueryAllValidatorRequest{
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
			resp, err := keeper.ValidatorAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Validator), step)
			require.Subset(t, msgs, resp.Validator)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ValidatorAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Validator), step)
			require.Subset(t, msgs, resp.Validator)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ValidatorAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ValidatorAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
*/
