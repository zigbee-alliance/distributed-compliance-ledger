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

func (k Keeper) ValidatorMissedBlockBitArrayAll(c context.Context, req *types.QueryAllValidatorMissedBlockBitArrayRequest) (*types.QueryAllValidatorMissedBlockBitArrayResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var validatorMissedBlockBitArrays []types.ValidatorMissedBlockBitArray
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	validatorMissedBlockBitArrayStore := prefix.NewStore(store, types.KeyPrefix(types.ValidatorMissedBlockBitArrayKeyPrefix))

	pageRes, err := query.Paginate(validatorMissedBlockBitArrayStore, req.Pagination, func(key []byte, value []byte) error {
		var validatorMissedBlockBitArray types.ValidatorMissedBlockBitArray
		if err := k.cdc.Unmarshal(value, &validatorMissedBlockBitArray); err != nil {
			return err
		}

		validatorMissedBlockBitArrays = append(validatorMissedBlockBitArrays, validatorMissedBlockBitArray)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllValidatorMissedBlockBitArrayResponse{ValidatorMissedBlockBitArray: validatorMissedBlockBitArrays, Pagination: pageRes}, nil
}

func (k Keeper) ValidatorMissedBlockBitArray(c context.Context, req *types.QueryGetValidatorMissedBlockBitArrayRequest) (*types.QueryGetValidatorMissedBlockBitArrayResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetValidatorMissedBlockBitArray(
		ctx,
		req.Address,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetValidatorMissedBlockBitArrayResponse{ValidatorMissedBlockBitArray: val}, nil
}
