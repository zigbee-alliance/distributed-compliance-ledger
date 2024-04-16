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

func (k Keeper) NocIcaCertificatesAll(c context.Context, req *types.QueryAllNocIcaCertificatesRequest) (*types.QueryAllNocIcaCertificatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var nocIcaCertificates []types.NocIcaCertificates
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	nocIcaCertificatesStore := prefix.NewStore(store, pkitypes.KeyPrefix(types.NocIcaCertificatesKeyPrefix))

	pageRes, err := query.Paginate(nocIcaCertificatesStore, req.Pagination, func(key []byte, value []byte) error {
		var nocCertificates types.NocIcaCertificates
		if err := k.cdc.Unmarshal(value, &nocCertificates); err != nil {
			return err
		}

		nocIcaCertificates = append(nocIcaCertificates, nocCertificates)

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNocIcaCertificatesResponse{NocIcaCertificates: nocIcaCertificates, Pagination: pageRes}, nil
}

func (k Keeper) NocIcaCertificates(c context.Context, req *types.QueryGetNocIcaCertificatesRequest) (*types.QueryGetNocIcaCertificatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetNocIcaCertificates(
		ctx,
		req.Vid,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetNocIcaCertificatesResponse{NocIcaCertificates: val}, nil
}
