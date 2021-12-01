package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func (k msgServer) ProposeRevokeAccount(goCtx context.Context, msg *types.MsgProposeRevokeAccount) (*types.MsgProposeRevokeAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgProposeRevokeAccountResponse{}, nil
}
