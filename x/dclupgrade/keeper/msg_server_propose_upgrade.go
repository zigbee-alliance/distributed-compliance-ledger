package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/types"
)

func (k msgServer) ProposeUpgrade(goCtx context.Context, msg *types.MsgProposeUpgrade) (*types.MsgProposeUpgradeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgProposeUpgradeResponse{}, nil
}
