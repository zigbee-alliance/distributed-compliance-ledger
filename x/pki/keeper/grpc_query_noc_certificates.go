package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func (k Keeper) NocCertificatesAll(c context.Context, req *types.QueryAllNocCertificatesRequest) (*types.QueryAllNocCertificatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var nocCertificatess []types.NocCertificates
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	nocCertificatesStore := prefix.NewStore(store, pkitypes.KeyPrefix(types.NocCertificatesKeyPrefix))

	pageRes, err := query.Paginate(nocCertificatesStore, req.Pagination, func(key []byte, value []byte) error {
		var nocCertificates types.NocCertificates
		if err := k.cdc.Unmarshal(value, &nocCertificates); err != nil {
			return err
		}

		nocCertificatess = append(nocCertificatess, nocCertificates)

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNocCertificatesResponse{NocCertificates: nocCertificatess, Pagination: pageRes}, nil
}

func (k Keeper) NocCertificates(c context.Context, req *types.QueryGetNocCertificatesRequest) (*types.QueryGetNocCertificatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetNocCertificates(
		ctx,
		req.Vid,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetNocCertificatesResponse{NocCertificates: val}, nil
}
