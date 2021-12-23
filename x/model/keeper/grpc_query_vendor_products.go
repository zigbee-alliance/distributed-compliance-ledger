package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) VendorProductsAll(c context.Context, req *types.QueryAllVendorProductsRequest) (*types.QueryAllVendorProductsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var vendorProductss []types.VendorProducts
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	vendorProductsStore := prefix.NewStore(store, types.KeyPrefix(types.VendorProductsKeyPrefix))

	pageRes, err := query.Paginate(vendorProductsStore, req.Pagination, func(key []byte, value []byte) error {
		var vendorProducts types.VendorProducts
		if err := k.cdc.Unmarshal(value, &vendorProducts); err != nil {
			return err
		}

		vendorProductss = append(vendorProductss, vendorProducts)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVendorProductsResponse{VendorProducts: vendorProductss, Pagination: pageRes}, nil
}

func (k Keeper) VendorProducts(c context.Context, req *types.QueryGetVendorProductsRequest) (*types.QueryGetVendorProductsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetVendorProducts(
		ctx,
		req.Vid,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetVendorProductsResponse{VendorProducts: val}, nil
}
