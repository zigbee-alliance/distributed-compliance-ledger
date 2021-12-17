package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func (k msgServer) ProposeRevokeX509RootCert(goCtx context.Context, msg *types.MsgProposeRevokeX509RootCert) (*types.MsgProposeRevokeX509RootCertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgProposeRevokeX509RootCertResponse{}, nil
}
