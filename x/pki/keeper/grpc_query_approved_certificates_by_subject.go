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

func (k Keeper) ApprovedCertificatesBySubjectAll(c context.Context, req *types.QueryAllApprovedCertificatesBySubjectRequest) (*types.QueryAllApprovedCertificatesBySubjectResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var approvedCertificatesBySubjects []types.ApprovedCertificatesBySubject
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	approvedCertificatesBySubjectStore := prefix.NewStore(store, types.KeyPrefix(types.ApprovedCertificatesBySubjectKeyPrefix))

	pageRes, err := query.Paginate(approvedCertificatesBySubjectStore, req.Pagination, func(key []byte, value []byte) error {
		var approvedCertificatesBySubject types.ApprovedCertificatesBySubject
		if err := k.cdc.Unmarshal(value, &approvedCertificatesBySubject); err != nil {
			return err
		}

		approvedCertificatesBySubjects = append(approvedCertificatesBySubjects, approvedCertificatesBySubject)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllApprovedCertificatesBySubjectResponse{ApprovedCertificatesBySubject: approvedCertificatesBySubjects, Pagination: pageRes}, nil
}

func (k Keeper) ApprovedCertificatesBySubject(c context.Context, req *types.QueryGetApprovedCertificatesBySubjectRequest) (*types.QueryGetApprovedCertificatesBySubjectResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, _ := k.GetApprovedCertificatesBySubject(
		ctx,
		req.Subject,
	)
	// if !found {
	// 	return nil, status.Error(codes.InvalidArgument, "not found")
	// }

	return &types.QueryGetApprovedCertificatesBySubjectResponse{ApprovedCertificatesBySubject: val}, nil
}
