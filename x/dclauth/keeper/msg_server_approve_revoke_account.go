package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func (k msgServer) ApproveRevokeAccount(goCtx context.Context, msg *types.MsgApproveRevokeAccount) (*types.MsgApproveRevokeAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgApproveRevokeAccountResponse{}, nil
}
