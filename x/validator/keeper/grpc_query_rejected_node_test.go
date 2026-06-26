package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/nullify"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

// createNRejectedDisableValidator stores n rejected-disable-validator records.
// SetRejectedNode keys by the (bech32 valoper) Address, silently skipping
// records with an invalid address, so valid sample addresses are required.
func createNRejectedDisableValidator(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.RejectedDisableValidator {
	items := make([]types.RejectedDisableValidator, n)
	for i := range items {
		items[i].Address = sample.ValAddress()

		keeper.SetRejectedNode(ctx, items[i])
	}

	return items
}

func TestRejectedDisableValidatorQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.ValidatorKeeper(t, nil)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRejectedDisableValidator(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetRejectedDisableValidatorRequest
		response *types.QueryGetRejectedDisableValidatorResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetRejectedDisableValidatorRequest{
				Owner: msgs[0].Address,
			},
			response: &types.QueryGetRejectedDisableValidatorResponse{RejectedValidator: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetRejectedDisableValidatorRequest{
				Owner: msgs[1].Address,
			},
			response: &types.QueryGetRejectedDisableValidatorResponse{RejectedValidator: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetRejectedDisableValidatorRequest{
				Owner: sample.ValAddress(),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.RejectedDisableValidator(wctx, tc.request)
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

func TestRejectedDisableValidatorQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.ValidatorKeeper(t, nil)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRejectedDisableValidator(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllRejectedDisableValidatorRequest {
		return &types.QueryAllRejectedDisableValidatorRequest{
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
			resp, err := keeper.RejectedDisableValidatorAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.RejectedValidator), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.RejectedValidator),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.RejectedDisableValidatorAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.RejectedValidator), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.RejectedValidator),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.RejectedDisableValidatorAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.RejectedValidator),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.RejectedDisableValidatorAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
