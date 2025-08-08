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

func TestApprovedCertificatesQueryPaginated(t *testing.T) {
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
		resp, err := keeper.ApprovedCertificatesAll(wctx, request(nil, 0, 0, true, "0"))
		require.NoError(t, err)
		require.Equal(t, 1, len(resp.ApprovedCertificates))
		require.Equal(t, msgs[0].SubjectKeyId, resp.ApprovedCertificates[0].SubjectKeyId)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ApprovedCertificatesAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

func TestApprovedCertificatesQuery_ExtendedScenarios(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	wctx := sdk.WrapSDKContext(ctx)
	
	tests := []struct {
		name        string
		setup       func() *types.QueryGetApprovedCertificatesRequest
		expectError bool
		errorMsg    string
		verify      func(*types.QueryGetApprovedCertificatesResponse, error)
	}{
		{
			name: "EmptyCertificates",
			setup: func() *types.QueryGetApprovedCertificatesRequest {
				emptyCert := createNApprovedCertificates(keeper, ctx, 1)[0]
				emptyCert.Certs = []*types.CertificateIdentifier{}
				keeper.SetApprovedCertificates(ctx, emptyCert)
				
				return &types.QueryGetApprovedCertificatesRequest{
					Subject:      emptyCert.Subject,
					SubjectKeyId: emptyCert.SubjectKeyId,
				}
			},
			expectError: false,
			verify: func(response *types.QueryGetApprovedCertificatesResponse, err error) {
				require.NoError(t, err)
				require.Empty(t, response.ApprovedCertificates.Certs)
			},
		},
		{
			name: "NilRequest",
			setup: func() *types.QueryGetApprovedCertificatesRequest {
				return nil
			},
			expectError: true,
			verify: func(response *types.QueryGetApprovedCertificatesResponse, err error) {
				require.Error(t, err)
			},
		},
		{
			name: "EmptySubjectAndKeyId",
			setup: func() *types.QueryGetApprovedCertificatesRequest {
				return &types.QueryGetApprovedCertificatesRequest{
					Subject:      "",
					SubjectKeyId: "",
				}
			},
			expectError: true,
			verify: func(response *types.QueryGetApprovedCertificatesResponse, err error) {
				require.Error(t, err)
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := tt.setup()
			response, err := keeper.ApprovedCertificates(wctx, request)
			tt.verify(response, err)
		})
	}
}

func TestApprovedCertificatesQueryAll_ExtendedScenarios(t *testing.T) {
	keeper, ctx := keepertest.PkiKeeper(t, nil)
	wctx := sdk.WrapSDKContext(ctx)
	
	tests := []struct {
		name        string
		setup       func() *types.QueryAllApprovedCertificatesRequest
		expectError bool
		verify      func(*types.QueryAllApprovedCertificatesResponse, error)
	}{
		{
			name: "EmptyResult",
			setup: func() *types.QueryAllApprovedCertificatesRequest {
				// Clear existing certificates
				allCerts := keeper.GetAllApprovedCertificates(ctx)
				for _, cert := range allCerts {
					keeper.RemoveApprovedCertificates(ctx, cert.Subject, cert.SubjectKeyId)
				}
				
				return &types.QueryAllApprovedCertificatesRequest{
					Pagination: &query.PageRequest{
						Limit: 10,
					},
				}
			},
			expectError: false,
			verify: func(response *types.QueryAllApprovedCertificatesResponse, err error) {
				require.NoError(t, err)
				require.Empty(t, response.ApprovedCertificates)
			},
		},
		{
			name: "WithSubjectKeyIdFilter",
			setup: func() *types.QueryAllApprovedCertificatesRequest {
				// Create certificates with different subject key IDs
				certs := createNApprovedCertificates(keeper, ctx, 3)
				
				return &types.QueryAllApprovedCertificatesRequest{
					Pagination: &query.PageRequest{
						Limit: 10,
					},
					SubjectKeyId: certs[0].SubjectKeyId,
				}
			},
			expectError: false,
			verify: func(response *types.QueryAllApprovedCertificatesResponse, err error) {
				require.NoError(t, err)
				require.Len(t, response.ApprovedCertificates, 1)
			},
		},
		{
			name: "Pagination",
			setup: func() *types.QueryAllApprovedCertificatesRequest {
				// Create multiple certificates
				createNApprovedCertificates(keeper, ctx, 10)
				
				return &types.QueryAllApprovedCertificatesRequest{
					Pagination: &query.PageRequest{
						Limit: 5,
					},
				}
			},
			expectError: false,
			verify: func(response *types.QueryAllApprovedCertificatesResponse, err error) {
				require.NoError(t, err)
				require.Len(t, response.ApprovedCertificates, 5)
				require.NotNil(t, response.Pagination)
			},
		},
		{
			name: "CountTotal",
			setup: func() *types.QueryAllApprovedCertificatesRequest {
				// Create certificates
				createNApprovedCertificates(keeper, ctx, 5)
				
				return &types.QueryAllApprovedCertificatesRequest{
					Pagination: &query.PageRequest{
						Limit:      10,
						CountTotal: true,
					},
				}
			},
			expectError: false,
			verify: func(response *types.QueryAllApprovedCertificatesResponse, err error) {
				require.NoError(t, err)
				require.Equal(t, uint64(5), response.Pagination.Total)
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := tt.setup()
			response, err := keeper.ApprovedCertificatesAll(wctx, request)
			tt.verify(response, err)
		})
	}
}
