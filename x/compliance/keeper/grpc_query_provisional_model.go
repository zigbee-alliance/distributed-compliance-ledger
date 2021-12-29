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

func (k Keeper) ProvisionalModelAll(c context.Context, req *types.QueryAllProvisionalModelRequest) (*types.QueryAllProvisionalModelResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var provisionalModels []types.ProvisionalModel
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	provisionalModelStore := prefix.NewStore(store, types.KeyPrefix(types.ProvisionalModelKeyPrefix))

	pageRes, err := query.Paginate(provisionalModelStore, req.Pagination, func(key []byte, value []byte) error {
		var provisionalModel types.ProvisionalModel
		if err := k.cdc.Unmarshal(value, &provisionalModel); err != nil {
			return err
		}

		provisionalModels = append(provisionalModels, provisionalModel)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllProvisionalModelResponse{ProvisionalModel: provisionalModels, Pagination: pageRes}, nil
}

func (k Keeper) ProvisionalModel(c context.Context, req *types.QueryGetProvisionalModelRequest) (*types.QueryGetProvisionalModelResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetProvisionalModel(
		ctx,
		req.Vid,
		req.Pid,
		req.SoftwareVersion,
		req.CertificationType,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetProvisionalModelResponse{ProvisionalModel: &val}, nil
}
