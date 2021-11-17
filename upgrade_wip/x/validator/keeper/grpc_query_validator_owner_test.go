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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestValidatorOwnerQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.ValidatorKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNValidatorOwner(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetValidatorOwnerRequest
		response *types.QueryGetValidatorOwnerResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetValidatorOwnerRequest{
				Address: msgs[0].Address,
			},
			response: &types.QueryGetValidatorOwnerResponse{ValidatorOwner: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetValidatorOwnerRequest{
				Address: msgs[1].Address,
			},
			response: &types.QueryGetValidatorOwnerResponse{ValidatorOwner: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetValidatorOwnerRequest{
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
			response, err := keeper.ValidatorOwner(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}

func TestValidatorOwnerQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.ValidatorKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNValidatorOwner(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllValidatorOwnerRequest {
		return &types.QueryAllValidatorOwnerRequest{
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
			resp, err := keeper.ValidatorOwnerAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ValidatorOwner), step)
			require.Subset(t, msgs, resp.ValidatorOwner)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ValidatorOwnerAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ValidatorOwner), step)
			require.Subset(t, msgs, resp.ValidatorOwner)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ValidatorOwnerAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ValidatorOwnerAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
