package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RevokedModelAll(c context.Context, req *types.QueryAllRevokedModelRequest) (*types.QueryAllRevokedModelResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var revokedModels []types.RevokedModel
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	revokedModelStore := prefix.NewStore(store, types.KeyPrefix(types.RevokedModelKeyPrefix))

	pageRes, err := query.FilteredPaginate(
		revokedModelStore, req.Pagination,
		func(key []byte, value []byte, accumulate bool) (bool, error) {
			var revokedModel types.RevokedModel
			if err := k.cdc.Unmarshal(value, &revokedModel); err != nil {
				return false, err
			}
			if !revokedModel.Value {
				return false, nil
			}
			if accumulate {
				revokedModels = append(revokedModels, revokedModel)
			}

			return true, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRevokedModelResponse{RevokedModel: revokedModels, Pagination: pageRes}, nil
}

func (k Keeper) RevokedModel(c context.Context, req *types.QueryGetRevokedModelRequest) (*types.QueryGetRevokedModelResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetRevokedModel(
		ctx,
		req.Vid,
		req.Pid,
		req.SoftwareVersion,
		req.CertificationType,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetRevokedModelResponse{RevokedModel: val}, nil
}
