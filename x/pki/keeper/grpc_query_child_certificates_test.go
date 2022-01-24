package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/nullify"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Prevent strconv unused error.
var _ = strconv.IntSize

func TestChildCertificatesQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
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
			err: status.Error(codes.NotFound, "not found"),
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
