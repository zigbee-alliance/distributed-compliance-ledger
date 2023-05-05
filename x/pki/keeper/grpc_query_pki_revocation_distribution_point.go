package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PKIRevocationDistributionPointAll(c context.Context, req *types.QueryAllPKIRevocationDistributionPointRequest) (*types.QueryAllPKIRevocationDistributionPointResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var pKIRevocationDistributionPoints []types.PKIRevocationDistributionPoint
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	pKIRevocationDistributionPointStore := prefix.NewStore(store, types.KeyPrefix(types.PKIRevocationDistributionPointKeyPrefix))

	pageRes, err := query.Paginate(pKIRevocationDistributionPointStore, req.Pagination, func(key []byte, value []byte) error {
		var pKIRevocationDistributionPoint types.PKIRevocationDistributionPoint
		if err := k.cdc.Unmarshal(value, &pKIRevocationDistributionPoint); err != nil {
			return err
		}

		pKIRevocationDistributionPoints = append(pKIRevocationDistributionPoints, pKIRevocationDistributionPoint)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPKIRevocationDistributionPointResponse{PKIRevocationDistributionPoint: pKIRevocationDistributionPoints, Pagination: pageRes}, nil
}

func (k Keeper) PKIRevocationDistributionPoint(c context.Context, req *types.QueryGetPKIRevocationDistributionPointRequest) (*types.QueryGetPKIRevocationDistributionPointResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetPKIRevocationDistributionPoint(
	    ctx,
	    req.Vid,
        req.Label,
        req.IssuerSubjectKeyID,
        )
	if !found {
	    return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetPKIRevocationDistributionPointResponse{PKIRevocationDistributionPoint: val}, nil
}