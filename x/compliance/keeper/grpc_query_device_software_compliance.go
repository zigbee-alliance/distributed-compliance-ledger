package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) DeviceSoftwareComplianceAll(c context.Context, req *types.QueryAllDeviceSoftwareComplianceRequest) (*types.QueryAllDeviceSoftwareComplianceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var deviceSoftwareCompliances []types.DeviceSoftwareCompliance
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	deviceSoftwareComplianceStore := prefix.NewStore(store, types.KeyPrefix(types.DeviceSoftwareComplianceKeyPrefix))

	pageRes, err := query.Paginate(deviceSoftwareComplianceStore, req.Pagination, func(key []byte, value []byte) error {
		var deviceSoftwareCompliance types.DeviceSoftwareCompliance
		if err := k.cdc.Unmarshal(value, &deviceSoftwareCompliance); err != nil {
			return err
		}

		deviceSoftwareCompliances = append(deviceSoftwareCompliances, deviceSoftwareCompliance)

		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDeviceSoftwareComplianceResponse{DeviceSoftwareCompliance: deviceSoftwareCompliances, Pagination: pageRes}, nil
}

func (k Keeper) DeviceSoftwareCompliance(c context.Context, req *types.QueryGetDeviceSoftwareComplianceRequest) (*types.QueryGetDeviceSoftwareComplianceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetDeviceSoftwareCompliance(
		ctx,
		req.CDCertificateId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetDeviceSoftwareComplianceResponse{DeviceSoftwareCompliance: val}, nil
}
