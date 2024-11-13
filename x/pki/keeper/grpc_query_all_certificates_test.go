package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/nullify"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Prevent strconv unused error.
var _ = strconv.IntSize

func TestCertificatesQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNCertificates(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetCertificatesRequest
		response *types.QueryGetCertificatesResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetCertificatesRequest{
				Subject:      msgs[0].Subject,
				SubjectKeyId: msgs[0].SubjectKeyId,
			},
			response: &types.QueryGetCertificatesResponse{Certificates: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetCertificatesRequest{
				Subject:      msgs[1].Subject,
				SubjectKeyId: msgs[1].SubjectKeyId,
			},
			response: &types.QueryGetCertificatesResponse{Certificates: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetCertificatesRequest{
				Subject:      strconv.Itoa(100000),
				SubjectKeyId: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Certificates(wctx, tc.request)
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

func TestCertificatesQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNCertificates(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllCertificatesRequest {
		return &types.QueryAllCertificatesRequest{
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
			resp, err := keeper.CertificatesAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Certificates), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Certificates),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.CertificatesAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Certificates), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Certificates),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.CertificatesAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Certificates),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ApprovedCertificatesAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
