package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/nullify"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestRevokedRootCertificatesQuery(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestRevokedRootCertificates(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetRevokedRootCertificatesRequest
		response *types.QueryGetRevokedRootCertificatesResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetRevokedRootCertificatesRequest{},
			response: &types.QueryGetRevokedRootCertificatesResponse{RevokedRootCertificates: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.RevokedRootCertificates(wctx, tc.request)
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
