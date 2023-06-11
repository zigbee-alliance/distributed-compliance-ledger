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

func TestPkiRevocationDistributionPointsByIssuerSubjectKeyIDQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNPkiRevocationDistributionPointsByIssuerSubjectKeyID(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest
		response *types.QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest{
				IssuerSubjectKeyID: msgs[0].IssuerSubjectKeyID,
			},
			response: &types.QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse{PkiRevocationDistributionPointsByIssuerSubjectKeyID: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest{
				IssuerSubjectKeyID: msgs[1].IssuerSubjectKeyID,
			},
			response: &types.QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse{PkiRevocationDistributionPointsByIssuerSubjectKeyID: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest{
				IssuerSubjectKeyID: strconv.Itoa(100000),
			},
			err: status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.PkiRevocationDistributionPointsByIssuerSubjectKeyID(wctx, tc.request)
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
