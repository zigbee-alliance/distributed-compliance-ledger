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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestChildCertificatesQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNChildCertificates(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetChildCertificatesRequest
		response *types.QueryGetChildCertificatesResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetChildCertificatesRequest{
				Issuer:         msgs[0].Issuer,
				AuthorityKeyId: msgs[0].AuthorityKeyId,
			},
			response: &types.QueryGetChildCertificatesResponse{ChildCertificates: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetChildCertificatesRequest{
				Issuer:         msgs[1].Issuer,
				AuthorityKeyId: msgs[1].AuthorityKeyId,
			},
			response: &types.QueryGetChildCertificatesResponse{ChildCertificates: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetChildCertificatesRequest{
				Issuer:         strconv.Itoa(100000),
				AuthorityKeyId: strconv.Itoa(100000),
			},
			err: status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ChildCertificates(wctx, tc.request)
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

func TestChildCertificatesQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNChildCertificates(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllChildCertificatesRequest {
		return &types.QueryAllChildCertificatesRequest{
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
			resp, err := keeper.ChildCertificatesAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ChildCertificates), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ChildCertificates),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ChildCertificatesAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ChildCertificates), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ChildCertificates),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ChildCertificatesAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.ChildCertificates),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ChildCertificatesAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
