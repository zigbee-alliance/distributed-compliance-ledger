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

func (k Keeper) DisabledValidatorAll(c context.Context, req *types.QueryAllDisabledValidatorRequest) (*types.QueryAllDisabledValidatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var disabledValidators []types.DisabledValidator
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	disabledValidatorStore := prefix.NewStore(store, types.KeyPrefix(types.DisabledValidatorKeyPrefix))

	pageRes, err := query.Paginate(disabledValidatorStore, req.Pagination, func(key []byte, value []byte) error {
		var disabledValidator types.DisabledValidator
		if err := k.cdc.Unmarshal(value, &disabledValidator); err != nil {
			return err
		}

		disabledValidators = append(disabledValidators, disabledValidator)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDisabledValidatorResponse{DisabledValidator: disabledValidators, Pagination: pageRes}, nil
}

func (k Keeper) DisabledValidator(c context.Context, req *types.QueryGetDisabledValidatorRequest) (*types.QueryGetDisabledValidatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetDisabledValidator(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetDisabledValidatorResponse{DisabledValidator: val}, nil
}
