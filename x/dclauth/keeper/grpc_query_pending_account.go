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

func (k Keeper) PendingAccountAll(c context.Context, req *types.QueryAllPendingAccountRequest) (*types.QueryAllPendingAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var pendingAccounts []types.PendingAccount
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	pendingAccountStore := prefix.NewStore(store, types.KeyPrefix(types.PendingAccountKeyPrefix))

	pageRes, err := query.Paginate(pendingAccountStore, req.Pagination, func(key []byte, value []byte) error {
		var pendingAccount types.PendingAccount
		if err := k.cdc.Unmarshal(value, &pendingAccount); err != nil {
			return err
		}

		pendingAccounts = append(pendingAccounts, pendingAccount)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPendingAccountResponse{PendingAccount: pendingAccounts, Pagination: pageRes}, nil
}

func (k Keeper) PendingAccount(c context.Context, req *types.QueryGetPendingAccountRequest) (*types.QueryGetPendingAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	addr, err := sdk.AccAddressFromBech32(req.Address)

	if err != nil {
		return nil, err
	}

	val, found := k.GetPendingAccount(
		ctx,
		addr,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetPendingAccountResponse{PendingAccount: val}, nil
}
