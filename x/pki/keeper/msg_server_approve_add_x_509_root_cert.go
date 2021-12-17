package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func (k msgServer) ApproveAddX509RootCert(goCtx context.Context, msg *types.MsgApproveAddX509RootCert) (*types.MsgApproveAddX509RootCertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgApproveAddX509RootCertResponse{}, nil
}
