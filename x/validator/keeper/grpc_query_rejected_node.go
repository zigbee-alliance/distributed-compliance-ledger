package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RejectedDisableNodeAll(c context.Context, req *types.QueryAllRejectedDisableNodeRequest) (*types.QueryAllRejectedDisableNodeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var rejectedNodes []types.RejectedDisableNode
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	rejectedNodeStore := prefix.NewStore(store, types.KeyPrefix(types.RejectedNodeKeyPrefix))

	pageRes, err := query.Paginate(rejectedNodeStore, req.Pagination, func(key []byte, value []byte) error {
		var rejectedNode types.RejectedDisableNode
		if err := k.cdc.Unmarshal(value, &rejectedNode); err != nil {
			return err
		}

		rejectedNodes = append(rejectedNodes, rejectedNode)

		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRejectedDisableNodeResponse{RejectedNode: rejectedNodes, Pagination: pageRes}, nil
}

func (k Keeper) RejectedDisableNode(c context.Context, req *types.QueryGetRejectedDisableNodeRequest) (*types.QueryGetRejectedDisableNodeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetRejectedNode(
		ctx,
		req.Owner,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetRejectedDisableNodeResponse{RejectedNode: val}, nil
}
