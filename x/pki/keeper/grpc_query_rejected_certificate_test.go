package keeper_test

/*
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

func TestRejectedCertificateQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRejectedCertificate(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetRejectedCertificateRequest
		response *types.QueryGetRejectedCertificateResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetRejectedCertificateRequest{
				Subject:      msgs[0].Subject,
				SubjectKeyID: msgs[0].SubjectKeyID,
			},
			response: &types.QueryGetRejectedCertificateResponse{RejectedCertificate: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetRejectedCertificateRequest{
				Subject:      msgs[1].Subject,
				SubjectKeyID: msgs[1].SubjectKeyID,
			},
			response: &types.QueryGetRejectedCertificateResponse{RejectedCertificate: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetRejectedCertificateRequest{
				Subject:      strconv.Itoa(100000),
				SubjectKeyID: strconv.Itoa(100000),
			},
			err: status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.RejectedCertificate(wctx, tc.request)
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

func TestRejectedCertificateQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRejectedCertificate(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllRejectedCertificateRequest {
		return &types.QueryAllRejectedCertificateRequest{
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
			resp, err := keeper.RejectedCertificateAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.RejectedCertificate), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.RejectedCertificate),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.RejectedCertificateAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.RejectedCertificate), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.RejectedCertificate),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.RejectedCertificateAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.RejectedCertificate),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.RejectedCertificateAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
*/
