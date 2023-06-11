package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PkiRevocationDistributionPointsByIssuerSubjectKeyID(c context.Context, req *types.QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest) (*types.QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(
		ctx,
		req.IssuerSubjectKeyID,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse{PkiRevocationDistributionPointsByIssuerSubjectKeyID: val}, nil
}
