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

func (k Keeper) NocRootCertificatesAll(c context.Context, req *types.QueryAllNocRootCertificatesRequest) (*types.QueryAllNocRootCertificatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var nocRootCertificatess []types.NocRootCertificates
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	nocRootCertificatesStore := prefix.NewStore(store, pkitypes.KeyPrefix(types.NocRootCertificatesKeyPrefix))

	pageRes, err := query.Paginate(nocRootCertificatesStore, req.Pagination, func(key []byte, value []byte) error {
		var nocRootCertificates types.NocRootCertificates
		if err := k.cdc.Unmarshal(value, &nocRootCertificates); err != nil {
			return err
		}

		nocRootCertificatess = append(nocRootCertificatess, nocRootCertificates)

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNocRootCertificatesResponse{NocRootCertificates: nocRootCertificatess, Pagination: pageRes}, nil
}

func (k Keeper) NocRootCertificates(c context.Context, req *types.QueryGetNocRootCertificatesRequest) (*types.QueryGetNocRootCertificatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetNocRootCertificates(
		ctx,
		req.Vid,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetNocRootCertificatesResponse{NocRootCertificates: val}, nil
}
