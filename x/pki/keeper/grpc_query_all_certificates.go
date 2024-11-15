package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func (k Keeper) CertificatesAll(c context.Context, req *types.QueryAllCertificatesRequest) (*types.QueryAllCertificatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var (
		certificatess []types.AllCertificates
		pageRes       *query.PageResponse
		err           error
	)
	ctx := sdk.UnwrapSDKContext(c)

	if req.SubjectKeyId != "" {
		aprCerts, found := k.GetAllCertificatesBySubjectKeyID(
			ctx,
			req.SubjectKeyId,
		)

		if found {
			certificatess = append(certificatess, types.AllCertificates{
				SubjectKeyId: aprCerts.SubjectKeyId,
				Certs:        aprCerts.Certs,
			})
		}
		pageRes = &query.PageResponse{Total: 1}
	} else {
		store := ctx.KVStore(k.storeKey)
		approvedCertificatesStore := prefix.NewStore(store, pkitypes.KeyPrefix(types.ApprovedCertificatesKeyPrefix))

		pageRes, err = query.Paginate(approvedCertificatesStore, req.Pagination, func(key []byte, value []byte) error {
			var certificates types.AllCertificates
			if err := k.cdc.Unmarshal(value, &certificates); err != nil {
				return err
			}

			certificatess = append(certificatess, certificates)

			return nil
		})
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &types.QueryAllCertificatesResponse{Certificates: certificatess, Pagination: pageRes}, nil
}

func (k Keeper) Certificates(c context.Context, req *types.QueryGetCertificatesRequest) (*types.QueryGetCertificatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetAllCertificates(
		ctx,
		req.Subject,
		req.SubjectKeyId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetCertificatesResponse{Certificates: val}, nil
}
