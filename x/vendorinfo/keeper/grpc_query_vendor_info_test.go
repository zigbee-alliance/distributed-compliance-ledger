package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Prevent strconv unused error.
var _ = strconv.IntSize

func TestVendorInfoQuerySingle(t *testing.T) {
	dclauthKeeper := &DclauthKeeperMock{}
	keeper, ctx := keepertest.VendorinfoKeeper(t, dclauthKeeper)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNVendorInfo(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetVendorInfoRequest
		response *types.QueryGetVendorInfoResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetVendorInfoRequest{
				VendorID: msgs[0].VendorID,
			},
			response: &types.QueryGetVendorInfoResponse{VendorInfo: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetVendorInfoRequest{
				VendorID: msgs[1].VendorID,
			},
			response: &types.QueryGetVendorInfoResponse{VendorInfo: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetVendorInfoRequest{
				VendorID: 100000,
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.VendorInfo(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}

func TestVendorInfoQueryPaginated(t *testing.T) {
	dclauthKeeper := &DclauthKeeperMock{}
	keeper, ctx := keepertest.VendorinfoKeeper(t, dclauthKeeper)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNVendorInfo(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllVendorInfoRequest {
		return &types.QueryAllVendorInfoRequest{
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
			resp, err := keeper.VendorInfoAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.VendorInfo), step)
			require.Subset(t, msgs, resp.VendorInfo)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.VendorInfoAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.VendorInfo), step)
			require.Subset(t, msgs, resp.VendorInfo)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.VendorInfoAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.VendorInfoAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
