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

func (k Keeper) ProposedUpgradeAll(c context.Context, req *types.QueryAllProposedUpgradeRequest) (*types.QueryAllProposedUpgradeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var proposedUpgrades []types.ProposedUpgrade
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	proposedUpgradeStore := prefix.NewStore(store, types.KeyPrefix(types.ProposedUpgradeKeyPrefix))

	pageRes, err := query.Paginate(proposedUpgradeStore, req.Pagination, func(key []byte, value []byte) error {
		var proposedUpgrade types.ProposedUpgrade
		if err := k.cdc.Unmarshal(value, &proposedUpgrade); err != nil {
			return err
		}

		proposedUpgrades = append(proposedUpgrades, proposedUpgrade)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllProposedUpgradeResponse{ProposedUpgrade: proposedUpgrades, Pagination: pageRes}, nil
}

func (k Keeper) ProposedUpgrade(c context.Context, req *types.QueryGetProposedUpgradeRequest) (*types.QueryGetProposedUpgradeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetProposedUpgrade(
		ctx,
		req.Name,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetProposedUpgradeResponse{ProposedUpgrade: val}, nil
}
