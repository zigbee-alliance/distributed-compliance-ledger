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

func (k Keeper) ValidatorOwnerAll(c context.Context, req *types.QueryAllValidatorOwnerRequest) (*types.QueryAllValidatorOwnerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var validatorOwners []types.ValidatorOwner
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	validatorOwnerStore := prefix.NewStore(store, types.KeyPrefix(types.ValidatorOwnerKeyPrefix))

	pageRes, err := query.Paginate(validatorOwnerStore, req.Pagination, func(key []byte, value []byte) error {
		var validatorOwner types.ValidatorOwner
		if err := k.cdc.Unmarshal(value, &validatorOwner); err != nil {
			return err
		}

		validatorOwners = append(validatorOwners, validatorOwner)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllValidatorOwnerResponse{ValidatorOwner: validatorOwners, Pagination: pageRes}, nil
}

func (k Keeper) ValidatorOwner(c context.Context, req *types.QueryGetValidatorOwnerRequest) (*types.QueryGetValidatorOwnerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetValidatorOwner(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetValidatorOwnerResponse{ValidatorOwner: val}, nil
}
