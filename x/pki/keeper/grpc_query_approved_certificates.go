package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ApprovedCertificatesAll(c context.Context, req *types.QueryAllApprovedCertificatesRequest) (*types.QueryAllApprovedCertificatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var approvedCertificatess []types.ApprovedCertificates
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	approvedCertificatesStore := prefix.NewStore(store, types.KeyPrefix(types.ApprovedCertificatesKeyPrefix))

	pageRes, err := query.Paginate(approvedCertificatesStore, req.Pagination, func(key []byte, value []byte) error {
		var approvedCertificates types.ApprovedCertificates
		if err := k.cdc.Unmarshal(value, &approvedCertificates); err != nil {
			return err
		}

		approvedCertificatess = append(approvedCertificatess, approvedCertificates)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllApprovedCertificatesResponse{ApprovedCertificates: approvedCertificatess, Pagination: pageRes}, nil
}

func (k Keeper) ApprovedCertificates(c context.Context, req *types.QueryGetApprovedCertificatesRequest) (*types.QueryGetApprovedCertificatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetApprovedCertificates(
		ctx,
		req.Subject,
		req.SubjectKeyId,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetApprovedCertificatesResponse{ApprovedCertificates: val}, nil
}
