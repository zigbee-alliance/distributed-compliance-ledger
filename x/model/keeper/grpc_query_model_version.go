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

func (k Keeper) ModelVersionAll(c context.Context, req *types.QueryAllModelVersionRequest) (*types.QueryAllModelVersionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var modelVersions []types.ModelVersion
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	modelVersionStore := prefix.NewStore(store, types.KeyPrefix(types.ModelVersionKeyPrefix))

	pageRes, err := query.Paginate(modelVersionStore, req.Pagination, func(key []byte, value []byte) error {
		var modelVersion types.ModelVersion
		if err := k.cdc.Unmarshal(value, &modelVersion); err != nil {
			return err
		}

		modelVersions = append(modelVersions, modelVersion)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllModelVersionResponse{ModelVersion: modelVersions, Pagination: pageRes}, nil
}

func (k Keeper) ModelVersion(c context.Context, req *types.QueryGetModelVersionRequest) (*types.QueryGetModelVersionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetModelVersion(
		ctx,
		req.Vid,
		req.Pid,
		req.SoftwareVersion,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetModelVersionResponse{ModelVersion: val}, nil
}
