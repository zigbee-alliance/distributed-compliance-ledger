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

func (k Keeper) ProposedDisableValidatorAll(c context.Context, req *types.QueryAllProposedDisableValidatorRequest) (*types.QueryAllProposedDisableValidatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var proposedDisableValidators []types.ProposedDisableValidator
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	proposedDisableValidatorStore := prefix.NewStore(store, types.KeyPrefix(types.ProposedDisableValidatorKeyPrefix))

	pageRes, err := query.Paginate(proposedDisableValidatorStore, req.Pagination, func(key []byte, value []byte) error {
		var proposedDisableValidator types.ProposedDisableValidator
		if err := k.cdc.Unmarshal(value, &proposedDisableValidator); err != nil {
			return err
		}

		proposedDisableValidators = append(proposedDisableValidators, proposedDisableValidator)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllProposedDisableValidatorResponse{ProposedDisableValidator: proposedDisableValidators, Pagination: pageRes}, nil
}

func (k Keeper) ProposedDisableValidator(c context.Context, req *types.QueryGetProposedDisableValidatorRequest) (*types.QueryGetProposedDisableValidatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetProposedDisableValidator(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetProposedDisableValidatorResponse{ProposedDisableValidator: val}, nil
}
