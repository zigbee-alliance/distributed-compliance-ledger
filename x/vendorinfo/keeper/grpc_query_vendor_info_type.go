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

func (k Keeper) VendorInfoTypeAll(c context.Context, req *types.QueryAllVendorInfoTypeRequest) (*types.QueryAllVendorInfoTypeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var vendorInfoTypes []types.VendorInfoType
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	vendorInfoTypeStore := prefix.NewStore(store, types.KeyPrefix(types.VendorInfoTypeKeyPrefix))

	pageRes, err := query.Paginate(vendorInfoTypeStore, req.Pagination, func(key []byte, value []byte) error {
		var vendorInfoType types.VendorInfoType
		if err := k.cdc.Unmarshal(value, &vendorInfoType); err != nil {
			return err
		}

		vendorInfoTypes = append(vendorInfoTypes, vendorInfoType)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVendorInfoTypeResponse{VendorInfoType: vendorInfoTypes, Pagination: pageRes}, nil
}

func (k Keeper) VendorInfoType(c context.Context, req *types.QueryGetVendorInfoTypeRequest) (*types.QueryGetVendorInfoTypeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetVendorInfoType(
		ctx,
		req.VendorID,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetVendorInfoTypeResponse{VendorInfoType: val}, nil
}
