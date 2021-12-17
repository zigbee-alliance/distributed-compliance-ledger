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

func (k Keeper) UniqueCertificateAll(c context.Context, req *types.QueryAllUniqueCertificateRequest) (*types.QueryAllUniqueCertificateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var uniqueCertificates []types.UniqueCertificate
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	uniqueCertificateStore := prefix.NewStore(store, types.KeyPrefix(types.UniqueCertificateKeyPrefix))

	pageRes, err := query.Paginate(uniqueCertificateStore, req.Pagination, func(key []byte, value []byte) error {
		var uniqueCertificate types.UniqueCertificate
		if err := k.cdc.Unmarshal(value, &uniqueCertificate); err != nil {
			return err
		}

		uniqueCertificates = append(uniqueCertificates, uniqueCertificate)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllUniqueCertificateResponse{UniqueCertificate: uniqueCertificates, Pagination: pageRes}, nil
}

func (k Keeper) UniqueCertificate(c context.Context, req *types.QueryGetUniqueCertificateRequest) (*types.QueryGetUniqueCertificateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetUniqueCertificate(
		ctx,
		req.Issuer,
		req.SerialNumber,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetUniqueCertificateResponse{UniqueCertificate: val}, nil
}
