package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) NocCertificatesByVidAndSkid(c context.Context, req *types.QueryGetNocCertificatesByVidAndSkidRequest) (*types.QueryGetNocCertificatesByVidAndSkidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetNocCertificatesByVidAndSkid(
		ctx,
		req.Vid,
		req.SubjectKeyId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetNocCertificatesByVidAndSkidResponse{NocCertificatesByVidAndSkid: val}, nil
}
