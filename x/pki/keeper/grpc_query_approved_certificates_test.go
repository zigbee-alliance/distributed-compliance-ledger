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

func TestApprovedCertificatesQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNApprovedCertificates(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetApprovedCertificatesRequest
		response *types.QueryGetApprovedCertificatesResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetApprovedCertificatesRequest{
				Subject:      msgs[0].Subject,
				SubjectKeyId: msgs[0].SubjectKeyId,
			},
			response: &types.QueryGetApprovedCertificatesResponse{ApprovedCertificates: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetApprovedCertificatesRequest{
				Subject:      msgs[1].Subject,
				SubjectKeyId: msgs[1].SubjectKeyId,
			},
			response: &types.QueryGetApprovedCertificatesResponse{ApprovedCertificates: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetApprovedCertificatesRequest{
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
			response, err := keeper.ApprovedCertificates(wctx, tc.request)
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

func TestApprovedCertificatesQueryAll(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNApprovedCertificates(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool, subjectKeyId string) *types.QueryAllApprovedCertificatesRequest {
		return &types.QueryAllApprovedCertificatesRequest{
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
			resp, err := keeper.ApprovedCertificatesAll(wctx, request(nil, uint64(i), uint64(step), false, ""))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ApprovedCertificates), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ApprovedCertificates),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ApprovedCertificatesAll(wctx, request(next, 0, uint64(step), false, ""))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ApprovedCertificates), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ApprovedCertificates),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ApprovedCertificatesAll(wctx, request(nil, 0, 0, true, ""))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.ApprovedCertificates),
		)
	})
	t.Run("By subjectkey-id", func(t *testing.T) {
		resp, err := keeper.ApprovedCertificatesAll(wctx, request(nil, 0, 0, true, msgs[1].SubjectKeyId))
		require.NoError(t, err)
		require.Equal(t, 1, len(resp.ApprovedCertificates))
		require.Equal(t, msgs[1].SubjectKeyId, resp.ApprovedCertificates[0].SubjectKeyId)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ApprovedCertificatesAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

func TestApprovedCertificatesQueryAll_Pagination(t *testing.T) {
	k, ctx := keepertest.PkiKeeper(t, nil)
	wctx := sdk.WrapSDKContext(ctx)
	certs := createNApprovedCertificates(k, ctx, 5)
	subjectKeyId := certs[0].SubjectKeyId

	tests := []struct {
		name         string
		req          *types.QueryAllApprovedCertificatesRequest
		expectCount  int
		expectTotal  uint64
		expectPaged  bool
		expectErr    bool
		expectErrMsg string
	}{
		{
			name: "WithSubjectKeyIdFilter",
			req: &types.QueryAllApprovedCertificatesRequest{
				SubjectKeyId: subjectKeyId,
				Pagination: &query.PageRequest{Limit: 10},
			},
			expectCount: 1,
		},
		{
			name: "AllCertificates",
			req: &types.QueryAllApprovedCertificatesRequest{
				Pagination: &query.PageRequest{Limit: 10},
			},
			expectCount: 5,
		},
		{
			name: "Pagination",
			req: &types.QueryAllApprovedCertificatesRequest{
				Pagination: &query.PageRequest{Limit: 3},
			},
			expectCount: 3,
			expectPaged: true,
		},
		{
			name: "CountTotal",
			req: &types.QueryAllApprovedCertificatesRequest{
				Pagination: &query.PageRequest{Limit: 10, CountTotal: true},
			},
			expectCount: 5,
			expectTotal: 5,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := k.ApprovedCertificatesAll(wctx, tc.req)

			if tc.expectErr {
				require.Error(t, err)
				if tc.expectErrMsg != "" {
					require.Contains(t, err.Error(), tc.expectErrMsg)
				}
				return
			}

			require.NoError(t, err)
			require.Len(t, resp.ApprovedCertificates, tc.expectCount)

			if tc.expectPaged {
				require.NotNil(t, resp.Pagination)
			}

			if tc.expectTotal > 0 {
				require.Equal(t, tc.expectTotal, resp.Pagination.Total)
			}
		})
	}
}
