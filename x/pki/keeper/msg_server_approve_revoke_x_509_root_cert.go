package keeper

import (
	"context"

	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func (k msgServer) ApproveRevokeX509RootCert(goCtx context.Context, msg *types.MsgApproveRevokeX509RootCert) (*types.MsgApproveRevokeX509RootCertResponse, error) {
	//_ := sdk.UnwrapSDKContext(goCtx)

	return &types.MsgApproveRevokeX509RootCertResponse{}, nil
}
