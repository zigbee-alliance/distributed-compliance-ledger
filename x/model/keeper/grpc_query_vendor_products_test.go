package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/nullify"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Prevent strconv unused error.
var _ = strconv.IntSize

func TestVendorProductsQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.ModelKeeper(t, nil)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNVendorProducts(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetVendorProductsRequest
		response *types.QueryGetVendorProductsResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetVendorProductsRequest{
				Vid: msgs[0].Vid,
			},
			response: &types.QueryGetVendorProductsResponse{VendorProducts: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetVendorProductsRequest{
				Vid: msgs[1].Vid,
			},
			response: &types.QueryGetVendorProductsResponse{VendorProducts: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetVendorProductsRequest{
				Vid: 100000,
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.VendorProducts(wctx, tc.request)
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
