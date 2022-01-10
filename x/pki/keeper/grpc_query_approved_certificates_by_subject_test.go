package keeper_test

/* TODO issue 99
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

func TestApprovedCertificatesBySubjectQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNApprovedCertificatesBySubject(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetApprovedCertificatesBySubjectRequest
		response *types.QueryGetApprovedCertificatesBySubjectResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetApprovedCertificatesBySubjectRequest{
				Subject: msgs[0].Subject,
			},
			response: &types.QueryGetApprovedCertificatesBySubjectResponse{ApprovedCertificatesBySubject: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetApprovedCertificatesBySubjectRequest{
				Subject: msgs[1].Subject,
			},
			response: &types.QueryGetApprovedCertificatesBySubjectResponse{ApprovedCertificatesBySubject: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetApprovedCertificatesBySubjectRequest{
				Subject: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ApprovedCertificatesBySubject(wctx, tc.request)
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
