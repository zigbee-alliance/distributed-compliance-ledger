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

func (k Keeper) RejectedCertificateAll(c context.Context, req *types.QueryAllRejectedCertificatesRequest) (*types.QueryAllRejectedCertificatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var rejectedCertificates []types.RejectedCertificate
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	rejectedCertificateStore := prefix.NewStore(store, types.KeyPrefix(types.RejectedCertificateKeyPrefix))

	pageRes, err := query.Paginate(rejectedCertificateStore, req.Pagination, func(key []byte, value []byte) error {
		var rejectedCertificate types.RejectedCertificate
		if err := k.cdc.Unmarshal(value, &rejectedCertificate); err != nil {
			return err
		}

		rejectedCertificates = append(rejectedCertificates, rejectedCertificate)

		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRejectedCertificatesResponse{RejectedCertificate: rejectedCertificates, Pagination: pageRes}, nil
}

func (k Keeper) RejectedCertificate(c context.Context, req *types.QueryGetRejectedCertificatesRequest) (*types.QueryGetRejectedCertificatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.NotFound, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetRejectedCertificate(
		ctx,
		req.Subject,
		req.SubjectKeyId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetRejectedCertificatesResponse{RejectedCertificate: val}, nil
}
