package keeper_test

import (
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// createNValidator stores n validators and returns them in the canonical
// (unmarshalled) form produced by the keeper, so that direct equality with
// query responses holds despite the cached pub key in the proto Any field.
func createNValidator(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Validator {
	items := make([]types.Validator, n)
	for i := range items {
		pk := ed25519.GenPrivKey().PubKey()
		validator, err := types.NewValidator(
			sdk.ValAddress(pk.Address()),
			pk,
			types.Description{Moniker: strconv.Itoa(i)},
		)
		if err != nil {
			panic(err)
		}

		keeper.SetValidator(ctx, validator)
		stored, _ := keeper.GetValidator(ctx, validator.GetOwner())
		items[i] = stored
	}

	return items
}

func TestValidatorQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.ValidatorKeeper(t, nil)
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
				Owner: sdk.ValAddress(ed25519.GenPrivKey().PubKey().Address()).String(),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "EmptyAddress",
			request: &types.QueryGetValidatorRequest{
				Owner: "",
			},
			err: status.Error(codes.InvalidArgument, "validator address cannot be empty"),
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
				require.NoError(t, err)
				require.Equal(t, tc.response, response)
			}
		})
	}

	t.Run("InvalidAddress", func(t *testing.T) {
		_, err := keeper.Validator(wctx, &types.QueryGetValidatorRequest{
			Owner: "not-a-bech32-address",
		})
		require.Error(t, err)
	})
}

func TestValidatorQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.ValidatorKeeper(t, nil)
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
		require.ElementsMatch(t, msgs, resp.Validator)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ValidatorAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
