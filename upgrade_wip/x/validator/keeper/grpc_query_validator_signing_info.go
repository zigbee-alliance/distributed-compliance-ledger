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

func (k Keeper) ValidatorSigningInfoAll(c context.Context, req *types.QueryAllValidatorSigningInfoRequest) (*types.QueryAllValidatorSigningInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var validatorSigningInfos []types.ValidatorSigningInfo
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	validatorSigningInfoStore := prefix.NewStore(store, types.KeyPrefix(types.ValidatorSigningInfoKeyPrefix))

	pageRes, err := query.Paginate(validatorSigningInfoStore, req.Pagination, func(key []byte, value []byte) error {
		var validatorSigningInfo types.ValidatorSigningInfo
		if err := k.cdc.Unmarshal(value, &validatorSigningInfo); err != nil {
			return err
		}

		validatorSigningInfos = append(validatorSigningInfos, validatorSigningInfo)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllValidatorSigningInfoResponse{ValidatorSigningInfo: validatorSigningInfos, Pagination: pageRes}, nil
}

func (k Keeper) ValidatorSigningInfo(c context.Context, req *types.QueryGetValidatorSigningInfoRequest) (*types.QueryGetValidatorSigningInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetValidatorSigningInfo(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetValidatorSigningInfoResponse{ValidatorSigningInfo: val}, nil
}
