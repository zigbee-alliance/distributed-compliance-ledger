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

func (k Keeper) ValidatorAll(c context.Context, req *types.QueryAllValidatorRequest) (*types.QueryAllValidatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var validators []types.Validator
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	validatorStore := prefix.NewStore(store, types.KeyPrefix(types.ValidatorKeyPrefix))

	// TODO issue 99: add filtering per State support here
	pageRes, err := query.Paginate(validatorStore, req.Pagination, func(key []byte, value []byte) error {
		var validator types.Validator
		if err := k.cdc.Unmarshal(value, &validator); err != nil {
			return err
		}

		validators = append(validators, validator)

		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllValidatorResponse{Validator: validators, Pagination: pageRes}, nil
}

func (k Keeper) Validator(c context.Context, req *types.QueryGetValidatorRequest) (*types.QueryGetValidatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if req.Owner == "" {
		return nil, status.Error(codes.InvalidArgument, "validator address cannot be empty")
	}

	valAddr, err := sdk.ValAddressFromBech32(req.Owner)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetValidator(
		ctx,
		valAddr,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetValidatorResponse{Validator: val}, nil
}
