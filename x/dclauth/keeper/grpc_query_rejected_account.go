package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RejectedAccountAll(c context.Context, req *types.QueryAllRejectedAccountRequest) (*types.QueryAllRejectedAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.NotFound, "not found")
	}

	var rejectedAccounts []types.RejectedAccount
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	rejectedAccountStore := prefix.NewStore(store, types.KeyPrefix(types.RejectedAccountKeyPrefix))

	pageRes, err := query.Paginate(rejectedAccountStore, req.Pagination, func(key []byte, value []byte) error {
		var rejectedAccount types.RejectedAccount
		if err := k.cdc.Unmarshal(value, &rejectedAccount); err != nil {
			return err
		}

		rejectedAccounts = append(rejectedAccounts, rejectedAccount)

		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRejectedAccountResponse{RejectedAccount: rejectedAccounts, Pagination: pageRes}, nil
}

func (k Keeper) RejectedAccount(c context.Context, req *types.QueryGetRejectedAccountRequest) (*types.QueryGetRejectedAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.NotFound, "not found")
	}
	ctx := sdk.UnwrapSDKContext(c)

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}

	val, found := k.GetRejectedAccount(
		ctx, addr,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetRejectedAccountResponse{RejectedAccount: val}, nil
}
