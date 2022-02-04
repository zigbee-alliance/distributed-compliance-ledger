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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestProposedUpgradeQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.DclupgradeKeeper(t, nil, nil)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNProposedUpgrade(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetProposedUpgradeRequest
		response *types.QueryGetProposedUpgradeResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetProposedUpgradeRequest{
				Name: msgs[0].Plan.Name,
			},
			response: &types.QueryGetProposedUpgradeResponse{ProposedUpgrade: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetProposedUpgradeRequest{
				Name: msgs[1].Plan.Name,
			},
			response: &types.QueryGetProposedUpgradeResponse{ProposedUpgrade: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetProposedUpgradeRequest{
				Name: strconv.Itoa(100000),
			},
			err: status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ProposedUpgrade(wctx, tc.request)
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

func TestProposedUpgradeQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.DclupgradeKeeper(t, nil, nil)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNProposedUpgrade(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllProposedUpgradeRequest {
		return &types.QueryAllProposedUpgradeRequest{
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
			resp, err := keeper.ProposedUpgradeAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ProposedUpgrade), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ProposedUpgrade),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ProposedUpgradeAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ProposedUpgrade), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ProposedUpgrade),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ProposedUpgradeAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.ProposedUpgrade),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ProposedUpgradeAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
