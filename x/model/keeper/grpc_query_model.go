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

func (k Keeper) ModelAll(c context.Context, req *types.QueryAllModelRequest) (*types.QueryAllModelResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var models []types.Model
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	modelStore := prefix.NewStore(store, types.KeyPrefix(types.ModelKeyPrefix))

	pageRes, err := query.Paginate(modelStore, req.Pagination, func(key []byte, value []byte) error {
		var model types.Model
		if err := k.cdc.Unmarshal(value, &model); err != nil {
			return err
		}

		models = append(models, model)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllModelResponse{Model: models, Pagination: pageRes}, nil
}

func (k Keeper) Model(c context.Context, req *types.QueryGetModelRequest) (*types.QueryGetModelResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetModel(
		ctx,
		req.Vid,
		req.Pid,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetModelResponse{Model: val}, nil
}
