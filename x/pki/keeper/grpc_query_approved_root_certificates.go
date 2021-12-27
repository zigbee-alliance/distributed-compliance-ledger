package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ApprovedRootCertificates(c context.Context, req *types.QueryGetApprovedRootCertificatesRequest) (*types.QueryGetApprovedRootCertificatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, _ := k.GetApprovedRootCertificates(ctx)

	// Return empty list if not found
	return &types.QueryGetApprovedRootCertificatesResponse{ApprovedRootCertificates: val}, nil
}
