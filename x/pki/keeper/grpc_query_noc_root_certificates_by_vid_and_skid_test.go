package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/nullify"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// Prevent strconv unused error.
var _ = strconv.IntSize

func TestNocRootCertificatesByVidAndSkidQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNNocRootCertificatesByVidAndSkid(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetNocRootCertificatesByVidAndSkidRequest
		response *types.QueryGetNocRootCertificatesByVidAndSkidResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetNocRootCertificatesByVidAndSkidRequest{
				Vid:          msgs[0].Vid,
				SubjectKeyId: msgs[0].SubjectKeyId,
			},
			response: &types.QueryGetNocRootCertificatesByVidAndSkidResponse{NocRootCertificatesByVidAndSkid: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetNocRootCertificatesByVidAndSkidRequest{
				Vid:          msgs[1].Vid,
				SubjectKeyId: msgs[1].SubjectKeyId,
			},
			response: &types.QueryGetNocRootCertificatesByVidAndSkidResponse{NocRootCertificatesByVidAndSkid: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetNocRootCertificatesByVidAndSkidRequest{
				Vid:          100000,
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
			response, err := keeper.NocRootCertificatesByVidAndSkid(wctx, tc.request)
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
