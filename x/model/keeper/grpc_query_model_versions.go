package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ModelVersionsAll(c context.Context, req *types.QueryAllModelVersionsRequest) (*types.QueryAllModelVersionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var modelVersionss []types.ModelVersions
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	modelVersionsStore := prefix.NewStore(store, types.KeyPrefix(types.ModelVersionsKeyPrefix))

	pageRes, err := query.Paginate(modelVersionsStore, req.Pagination, func(key []byte, value []byte) error {
		var modelVersions types.ModelVersions
		if err := k.cdc.Unmarshal(value, &modelVersions); err != nil {
			return err
		}

		modelVersionss = append(modelVersionss, modelVersions)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllModelVersionsResponse{ModelVersions: modelVersionss, Pagination: pageRes}, nil
}

func (k Keeper) ModelVersions(c context.Context, req *types.QueryGetModelVersionsRequest) (*types.QueryGetModelVersionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetModelVersions(
		ctx,
		req.Vid,
		req.Pid,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetModelVersionsResponse{ModelVersions: val}, nil
}
