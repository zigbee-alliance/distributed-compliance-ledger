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

func (k Keeper) RevokedCertificatesAll(c context.Context, req *types.QueryAllRevokedCertificatesRequest) (*types.QueryAllRevokedCertificatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var revokedCertificatess []types.RevokedCertificates
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	revokedCertificatesStore := prefix.NewStore(store, types.KeyPrefix(types.RevokedCertificatesKeyPrefix))

	pageRes, err := query.Paginate(revokedCertificatesStore, req.Pagination, func(key []byte, value []byte) error {
		var revokedCertificates types.RevokedCertificates
		if err := k.cdc.Unmarshal(value, &revokedCertificates); err != nil {
			return err
		}

		revokedCertificatess = append(revokedCertificatess, revokedCertificates)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRevokedCertificatesResponse{RevokedCertificates: revokedCertificatess, Pagination: pageRes}, nil
}

func (k Keeper) RevokedCertificates(c context.Context, req *types.QueryGetRevokedCertificatesRequest) (*types.QueryGetRevokedCertificatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetRevokedCertificates(
		ctx,
		req.Subject,
		req.SubjectKeyId,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetRevokedCertificatesResponse{RevokedCertificates: val}, nil
}
