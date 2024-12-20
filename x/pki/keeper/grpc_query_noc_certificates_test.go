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

func TestNocCertificatesQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNNocCertificates(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetNocCertificatesRequest
		response *types.QueryGetNocCertificatesResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetNocCertificatesRequest{
				Subject:      msgs[0].Subject,
				SubjectKeyId: msgs[0].SubjectKeyId,
			},
			response: &types.QueryGetNocCertificatesResponse{NocCertificates: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetNocCertificatesRequest{
				Subject:      msgs[1].Subject,
				SubjectKeyId: msgs[1].SubjectKeyId,
			},
			response: &types.QueryGetNocCertificatesResponse{NocCertificates: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetNocCertificatesRequest{
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
			response, err := keeper.NocCertificates(wctx, tc.request)
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

func TestQueryNocCertificatesQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNNocCertificates(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool, subjectKeyId string) *types.QueryNocCertificatesRequest {
		return &types.QueryNocCertificatesRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
			SubjectKeyId: subjectKeyId,
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.NocCertificatesAll(wctx, request(nil, uint64(i), uint64(step), false, ""))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.NocCertificates), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.NocCertificates),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.NocCertificatesAll(wctx, request(next, 0, uint64(step), false, ""))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.NocCertificates), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.NocCertificates),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.NocCertificatesAll(wctx, request(nil, 0, 0, true, ""))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.NocCertificates),
		)
	})
	t.Run("By subjectkey-id", func(t *testing.T) {
		resp, err := keeper.NocCertificatesAll(wctx, request(nil, 0, 0, true, "0"))
		require.NoError(t, err)
		require.Equal(t, 1, len(resp.NocCertificates))
		require.Equal(t, msgs[0].SubjectKeyId, resp.NocCertificates[0].SubjectKeyId)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.NocCertificatesAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
