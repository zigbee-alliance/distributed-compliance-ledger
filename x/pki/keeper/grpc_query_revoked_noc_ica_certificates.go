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

func (k Keeper) RevokedNocIcaCertificatesAll(c context.Context, req *types.QueryAllRevokedNocIcaCertificatesRequest) (*types.QueryAllRevokedNocIcaCertificatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var revokedNocIcaCertificatess []types.RevokedNocIcaCertificates
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	revokedNocIcaCertificatesStore := prefix.NewStore(store, pkitypes.KeyPrefix(types.RevokedNocIcaCertificatesKeyPrefix))

	pageRes, err := query.Paginate(revokedNocIcaCertificatesStore, req.Pagination, func(key []byte, value []byte) error {
		var revokedNocIcaCertificates types.RevokedNocIcaCertificates
		if err := k.cdc.Unmarshal(value, &revokedNocIcaCertificates); err != nil {
			return err
		}

		revokedNocIcaCertificatess = append(revokedNocIcaCertificatess, revokedNocIcaCertificates)

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRevokedNocIcaCertificatesResponse{RevokedNocIcaCertificates: revokedNocIcaCertificatess, Pagination: pageRes}, nil
}

func (k Keeper) RevokedNocIcaCertificates(c context.Context, req *types.QueryGetRevokedNocIcaCertificatesRequest) (*types.QueryGetRevokedNocIcaCertificatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetRevokedNocIcaCertificates(
		ctx,
		req.Subject,
		req.SubjectKeyId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetRevokedNocIcaCertificatesResponse{RevokedNocIcaCertificates: val}, nil
}
