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

func (k Keeper) RejectedDisableValidatorAll(c context.Context, req *types.QueryAllRejectedDisableValidatorRequest) (*types.QueryAllRejectedDisableValidatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var rejectedValidators []types.RejectedDisableValidator
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	rejectedNodeStore := prefix.NewStore(store, types.KeyPrefix(types.RejectedNodeKeyPrefix))

	pageRes, err := query.Paginate(rejectedNodeStore, req.Pagination, func(key []byte, value []byte) error {
		var rejectedNode types.RejectedDisableValidator
		if err := k.cdc.Unmarshal(value, &rejectedNode); err != nil {
			return err
		}

		rejectedValidators = append(rejectedValidators, rejectedNode)

		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRejectedDisableValidatorResponse{RejectedValidator: rejectedValidators, Pagination: pageRes}, nil
}

func (k Keeper) RejectedDisableValidator(c context.Context, req *types.QueryGetRejectedDisableValidatorRequest) (*types.QueryGetRejectedDisableValidatorResponse, error) {
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

	return &types.QueryGetRejectedDisableValidatorResponse{RejectedValidator: val}, nil
}
