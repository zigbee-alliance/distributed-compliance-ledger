package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PkiRevocationDistributionPointAll(c context.Context, req *types.QueryAllPkiRevocationDistributionPointRequest) (*types.QueryAllPkiRevocationDistributionPointResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var pKIRevocationDistributionPoints []types.PkiRevocationDistributionPoint
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	pKIRevocationDistributionPointStore := prefix.NewStore(store, pkitypes.KeyPrefix(types.PkiRevocationDistributionPointKeyPrefix))

	pageRes, err := query.Paginate(pKIRevocationDistributionPointStore, req.Pagination, func(key []byte, value []byte) error {
		var pKIRevocationDistributionPoint types.PkiRevocationDistributionPoint
		if err := k.cdc.Unmarshal(value, &pKIRevocationDistributionPoint); err != nil {
			return err
		}

		pKIRevocationDistributionPoints = append(pKIRevocationDistributionPoints, pKIRevocationDistributionPoint)

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPkiRevocationDistributionPointResponse{PkiRevocationDistributionPoint: pKIRevocationDistributionPoints, Pagination: pageRes}, nil
}

func (k Keeper) PkiRevocationDistributionPoint(c context.Context, req *types.QueryGetPkiRevocationDistributionPointRequest) (*types.QueryGetPkiRevocationDistributionPointResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetPkiRevocationDistributionPoint(
		ctx,
		req.Vid,
		req.Label,
		req.IssuerSubjectKeyID,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetPkiRevocationDistributionPointResponse{PkiRevocationDistributionPoint: val}, nil
}
