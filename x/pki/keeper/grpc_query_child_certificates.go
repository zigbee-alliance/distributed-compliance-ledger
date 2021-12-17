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

func (k Keeper) ChildCertificatesAll(c context.Context, req *types.QueryAllChildCertificatesRequest) (*types.QueryAllChildCertificatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var childCertificatess []types.ChildCertificates
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	childCertificatesStore := prefix.NewStore(store, types.KeyPrefix(types.ChildCertificatesKeyPrefix))

	pageRes, err := query.Paginate(childCertificatesStore, req.Pagination, func(key []byte, value []byte) error {
		var childCertificates types.ChildCertificates
		if err := k.cdc.Unmarshal(value, &childCertificates); err != nil {
			return err
		}

		childCertificatess = append(childCertificatess, childCertificates)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllChildCertificatesResponse{ChildCertificates: childCertificatess, Pagination: pageRes}, nil
}

func (k Keeper) ChildCertificates(c context.Context, req *types.QueryGetChildCertificatesRequest) (*types.QueryGetChildCertificatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetChildCertificates(
		ctx,
		req.Issuer,
		req.AuthorityKeyId,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetChildCertificatesResponse{ChildCertificates: val}, nil
}
