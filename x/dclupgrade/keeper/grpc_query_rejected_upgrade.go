package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RejectedUpgradeAll(c context.Context, req *types.QueryAllRejectedUpgradeRequest) (*types.QueryAllRejectedUpgradeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var rejectedUpgrades []types.RejectedUpgrade
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	rejectedUpgradeStore := prefix.NewStore(store, types.KeyPrefix(types.RejectedUpgradeKeyPrefix))

	pageRes, err := query.Paginate(rejectedUpgradeStore, req.Pagination, func(key []byte, value []byte) error {
		var rejectedUpgrade types.RejectedUpgrade
		if err := k.cdc.Unmarshal(value, &rejectedUpgrade); err != nil {
			return err
		}

		rejectedUpgrades = append(rejectedUpgrades, rejectedUpgrade)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRejectedUpgradeResponse{RejectedUpgrade: rejectedUpgrades, Pagination: pageRes}, nil
}

func (k Keeper) RejectedUpgrade(c context.Context, req *types.QueryGetRejectedUpgradeRequest) (*types.QueryGetRejectedUpgradeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetRejectedUpgrade(
		ctx,
		req.Name,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetRejectedUpgradeResponse{RejectedUpgrade: val}, nil
}
