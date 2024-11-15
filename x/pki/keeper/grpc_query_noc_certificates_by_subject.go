package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) NocCertificatesBySubject(c context.Context, req *types.QueryGetNocCertificatesBySubjectRequest) (*types.QueryGetNocCertificatesBySubjectResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetNocCertificatesBySubject(
		ctx,
		req.Subject,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetNocCertificatesBySubjectResponse{NocCertificatesBySubject: val}, nil
}
