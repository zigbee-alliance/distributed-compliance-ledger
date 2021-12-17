package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func (k msgServer) RevokeX509Cert(goCtx context.Context, msg *types.MsgRevokeX509Cert) (*types.MsgRevokeX509CertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRevokeX509CertResponse{}, nil
}
