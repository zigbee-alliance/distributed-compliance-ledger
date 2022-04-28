package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/types"
)

func (k msgServer) RejectUpgrade(goCtx context.Context, msg *types.MsgRejectUpgrade) (*types.MsgRejectUpgradeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRejectUpgradeResponse{}, nil
}
