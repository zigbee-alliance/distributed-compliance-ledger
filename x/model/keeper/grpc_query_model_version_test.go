package keeper_test

/*
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

func TestModelVersionQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.ModelKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNModelVersion(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetModelVersionRequest
		response *types.QueryGetModelVersionResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetModelVersionRequest{
				Vid:             msgs[0].Vid,
				Pid:             msgs[0].Pid,
				SoftwareVersion: msgs[0].SoftwareVersion,
			},
			response: &types.QueryGetModelVersionResponse{ModelVersion: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetModelVersionRequest{
				Vid:             msgs[1].Vid,
				Pid:             msgs[1].Pid,
				SoftwareVersion: msgs[1].SoftwareVersion,
			},
			response: &types.QueryGetModelVersionResponse{ModelVersion: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetModelVersionRequest{
				Vid:             100000,
				Pid:             100000,
				SoftwareVersion: 100000,
			},
			err: status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ModelVersion(wctx, tc.request)
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
*/
