package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/nullify"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Prevent strconv unused error.
var _ = strconv.IntSize

func TestComplianceInfoQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.ComplianceKeeper(t, nil, nil)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNComplianceInfo(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetComplianceInfoRequest
		response *types.QueryGetComplianceInfoResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetComplianceInfoRequest{
				Vid:               msgs[0].Vid,
				Pid:               msgs[0].Pid,
				SoftwareVersion:   msgs[0].SoftwareVersion,
				CertificationType: msgs[0].CertificationType,
			},
			response: &types.QueryGetComplianceInfoResponse{ComplianceInfo: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetComplianceInfoRequest{
				Vid:               msgs[1].Vid,
				Pid:               msgs[1].Pid,
				SoftwareVersion:   msgs[1].SoftwareVersion,
				CertificationType: msgs[1].CertificationType,
			},
			response: &types.QueryGetComplianceInfoResponse{ComplianceInfo: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetComplianceInfoRequest{
				Vid:               100000,
				Pid:               100000,
				SoftwareVersion:   100000,
				CertificationType: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ComplianceInfo(wctx, tc.request)
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

func TestComplianceInfoQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.ComplianceKeeper(t, nil, nil)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNComplianceInfo(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllComplianceInfoRequest {
		return &types.QueryAllComplianceInfoRequest{
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
			resp, err := keeper.ComplianceInfoAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ComplianceInfo), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ComplianceInfo),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ComplianceInfoAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ComplianceInfo), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ComplianceInfo),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ComplianceInfoAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.ComplianceInfo),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ComplianceInfoAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
