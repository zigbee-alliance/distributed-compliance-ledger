package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) NewVendorInfoAll(c context.Context, req *types.QueryAllNewVendorInfoRequest) (*types.QueryAllNewVendorInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var newVendorInfos []types.NewVendorInfo
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	newVendorInfoStore := prefix.NewStore(store, types.KeyPrefix(types.NewVendorInfoKeyPrefix))

	pageRes, err := query.Paginate(newVendorInfoStore, req.Pagination, func(key []byte, value []byte) error {
		var newVendorInfo types.NewVendorInfo
		if err := k.cdc.Unmarshal(value, &newVendorInfo); err != nil {
			return err
		}

		newVendorInfos = append(newVendorInfos, newVendorInfo)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNewVendorInfoResponse{NewVendorInfo: newVendorInfos, Pagination: pageRes}, nil
}

func (k Keeper) NewVendorInfo(c context.Context, req *types.QueryGetNewVendorInfoRequest) (*types.QueryGetNewVendorInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetNewVendorInfo(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetNewVendorInfoResponse{NewVendorInfo: val}, nil
}
