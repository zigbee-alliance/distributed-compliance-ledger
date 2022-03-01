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

func (k Keeper) ApprovedUpgradeAll(c context.Context, req *types.QueryAllApprovedUpgradeRequest) (*types.QueryAllApprovedUpgradeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var approvedUpgrades []types.ApprovedUpgrade
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	approvedUpgradeStore := prefix.NewStore(store, types.KeyPrefix(types.ApprovedUpgradeKeyPrefix))

	pageRes, err := query.Paginate(approvedUpgradeStore, req.Pagination, func(key []byte, value []byte) error {
		var approvedUpgrade types.ApprovedUpgrade
		if err := k.cdc.Unmarshal(value, &approvedUpgrade); err != nil {
			return err
		}

		approvedUpgrades = append(approvedUpgrades, approvedUpgrade)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllApprovedUpgradeResponse{ApprovedUpgrade: approvedUpgrades, Pagination: pageRes}, nil
}

func (k Keeper) ApprovedUpgrade(c context.Context, req *types.QueryGetApprovedUpgradeRequest) (*types.QueryGetApprovedUpgradeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetApprovedUpgrade(
		ctx,
		req.Name,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetApprovedUpgradeResponse{ApprovedUpgrade: val}, nil
}
