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

func (k Keeper) RevokedAccountAll(c context.Context, req *types.QueryAllRevokedAccountRequest) (*types.QueryAllRevokedAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var revokedAccounts []types.RevokedAccount
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	revokedAccountStore := prefix.NewStore(store, types.KeyPrefix(types.RevokedAccountKeyPrefix))

	pageRes, err := query.Paginate(revokedAccountStore, req.Pagination, func(key []byte, value []byte) error {
		var revokedAccount types.RevokedAccount
		if err := k.cdc.Unmarshal(value, &revokedAccount); err != nil {
			return err
		}

		revokedAccounts = append(revokedAccounts, revokedAccount)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRevokedAccountResponse{RevokedAccount: revokedAccounts, Pagination: pageRes}, nil
}

func (k Keeper) RevokedAccount(c context.Context, req *types.QueryGetRevokedAccountRequest) (*types.QueryGetRevokedAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.NotFound, "not found")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetRevokedAccount(
		ctx,
		sdk.AccAddress(req.Address),
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetRevokedAccountResponse{RevokedAccount: val}, nil
}
