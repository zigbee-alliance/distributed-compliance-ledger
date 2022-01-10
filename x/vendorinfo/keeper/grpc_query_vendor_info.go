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

func (k Keeper) VendorInfoAll(c context.Context, req *types.QueryAllVendorInfoRequest) (*types.QueryAllVendorInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var vendorInfos []types.VendorInfo
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	vendorInfoStore := prefix.NewStore(store, types.KeyPrefix(types.VendorInfoKeyPrefix))

	pageRes, err := query.Paginate(vendorInfoStore, req.Pagination, func(key []byte, value []byte) error {
		var vendorInfo types.VendorInfo
		if err := k.cdc.Unmarshal(value, &vendorInfo); err != nil {
			return err
		}

		vendorInfos = append(vendorInfos, vendorInfo)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVendorInfoResponse{VendorInfo: vendorInfos, Pagination: pageRes}, nil
}

func (k Keeper) VendorInfo(c context.Context, req *types.QueryGetVendorInfoRequest) (*types.QueryGetVendorInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetVendorInfo(
		ctx,
		req.VendorID,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetVendorInfoResponse{VendorInfo: val}, nil
}
