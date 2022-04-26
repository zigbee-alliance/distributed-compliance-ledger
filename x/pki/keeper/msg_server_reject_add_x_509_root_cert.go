package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func (k msgServer) RejectAddX509RootCert(goCtx context.Context, msg *types.MsgRejectAddX509RootCert) (*types.MsgRejectAddX509RootCertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRejectAddX509RootCertResponse{}, nil
}
